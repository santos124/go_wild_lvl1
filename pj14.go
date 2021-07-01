// Имеется последовательность строк - 
// (cat, cat, dog, cat, tree) 
// создать для нее собственное множество.
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arrStr := strings.Split(text, ", ")
	fmt.Println(arrStr)
}