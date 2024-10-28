/* Simple parallelized mergesort implementation created as a Go programming
   exercise

   Copyright 2024 Dominic Delabruere
   This program is free software: you can redistribute it and/or modify it under
   the terms of the GNU General Public License as published by the Free Software
   Foundation, either version 3 of the License, or (at your option) any later
   version.

   This program is distributed in the hope that it will be useful, but WITHOUT
   ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
   FOR A PARTICULAR PURPOSE. See the GNU General Public License for more
   details.

   You should have received a copy of the GNU General Public License along with
   this program. If not, see <https://www.gnu.org/licenses/>. */

package main

import (
	"fmt"
	"sync"
)

type ordered interface {
	~int | ~float32 | ~float64 | ~byte | ~rune | ~string
}

func merge[T ordered](left []T, right []T, combined []T) {
	i_left, i_right := 0, 0
	for i := 0; i < len(combined); i++ {
		switch {
		case i_left >= len(left) || (i_right < len(right) && right[i_right] < left[i_left]):
			combined[i] = right[i_right]
			i_right++
		default:
			combined[i] = left[i_left]
			i_left++
		}
	}
}

func msort[T ordered](a []T) {
	if len(a) < 2 {
		return
	}
	if len(a) == 2 {
		if a[1] < a[0] {
			a[0], a[1] = a[1], a[0]
		}
		return
	}
	var half int = len(a) / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		msort(a[0:half])
	}()
	go func() {
		defer wg.Done()
		msort(a[half:])
	}()
	wg.Wait()
	var left, right = make([]T, half), make([]T, len(a)-half)
	copy(left, a[0:half])
	copy(right, a[half:])
	merge(left, right, a)
}

func main() {
	a := [...]int{8, 2, 3, 9, 1, 7, 4, 5, 0, 6}
	fmt.Println(a)
	msort(a[0:])
	fmt.Println(a)
}
