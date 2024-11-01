/* Main invocation for Go programming exercises

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

	"github.com/ddelabru/msort"
)

func main() {
	a := [...]int{8, 2, 3, 9, 1, 7, 4, 5, 0, 6}
	fmt.Println(a)
	msort.Msort(a[0:])
	fmt.Println(a)
}
