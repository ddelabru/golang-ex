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

package msort

import "sync"

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

func Msort[T ordered](a []T) {
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
		Msort(a[0:half])
	}()
	go func() {
		defer wg.Done()
		Msort(a[half:])
	}()
	wg.Wait()
	var left, right = make([]T, half), make([]T, len(a)-half)
	copy(left, a[0:half])
	copy(right, a[half:])
	merge(left, right, a)
}
