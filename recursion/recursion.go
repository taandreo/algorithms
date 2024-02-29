/*
Recursion is when a function call itself.

In this example i am doing the factorial (but with addition)

for example:

fat(5) = 5 + 4 + 3 + 2 + 1 = 15

Using recursion we can do that using recursion and defining our stop condition

In this case i when a number is equal to one

Wy can divide the recursion phases in three main steps:

- pre: n + x
- recursion: fat(n + x)
- post: in this case we dont use that
*/

package main

import "fmt"

func fat(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return (n + fat(n-1))
}

func main() {
	fmt.Println(fat(17))
}
