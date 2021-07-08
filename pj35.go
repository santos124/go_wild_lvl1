package main
import (
	"fmt"
)
// АААBBCCCDDEEE -> 3A2B3C2D3E
func main() {
	s := "АААBBCCCDDEEE"
	s2 := []rune(s)
	var s3 []rune = []rune{}
	var count int = 0
	for i := 0; i < len(s2); i++ {
		for j := i; j < len(s2); j++ {
			if s2[i] != s2[j]{
				s3 = append(s3, rune(count + '0'))
				s3 = append(s3, s2[i])
				i = i + count -1
				count = 0
				j = len(s2)
			} else {
				count++
			}
			if count != 0  && j == len(s2) - 1{
				s3 = append(s3, rune(count + '0'))
				s3 = append(s3, s2[i])
				i = i + count -1
				count = 0
				j = len(s2)
			}
		}
	}
	s4 := string(s3)
	fmt.Printf("s4=%s\n",s4)
}