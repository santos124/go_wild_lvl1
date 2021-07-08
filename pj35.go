package main
import (
	"fmt"
)
// АААBBCCCDDEEE -> 3A2B3C2D3E
func main() {
	s := "qaabbccc"
	s2 := []byte(s)
	var s3 []byte = []byte{}
	var count int = 0
	for i := 0; i < len(s2); i++ {
		for j := i; j < len(s2); j++ {
			if s2[i] != s2[j]{
				fmt.Printf("s3=%s, s2[i]=%v, i=%d, j=%d, count=%v\n", s3, s2[i], i, j, byte(count + 48))
				s3 = append(s3, byte(count + '0'))
				
				s3 = append(s3, s2[i])
				i = i + count -1
				count = 0
				j = len(s2)
			} else {
				count++
			}
			
		}
		
	}
	fmt.Printf("s=%s, s2=%s, s3=%s\n", s, s2, s3)
}