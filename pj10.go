// Дана последовательность температурных колебаний 
// (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5)
// . Объединить данный значения в группы с шагов в 10 
// градусов. Последовательность в подмножностве не важна.
// Пример: (-20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc)

package main
import (
	"fmt"
)
func maxMin(arr []float32) (float32, float32) {
	max := float32(-273.4)
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	min := max
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
func separator(arr []float32) map[int][]float32 {
	fmt.Printf("%v\n", arr)
	m1 := make(map[int][]float32)
	max, min := maxMin(arr)
	for i := float32(-100.0); i < max; i = i + 10.0 {
		tempS := []float32 {}
		if i < 0.0 {
			for j := range arr {
				if arr[j] < i  && arr[j] >= i - 10.0{
					tempS = append(tempS, arr[j])
				}
			}
		}
		if i == 0.0 {
			for j := range arr {
				if arr[j] >= i - 10.0  && arr[j] < i + 10.0{
					tempS = append(tempS, arr[j])
				}
			}
		}
		if i > 0.0 {
			for j := range arr {
				if arr[j] < i + 10.0  && arr[j] >= i {
					tempS = append(tempS, arr[j])
				}
			}
		}
		if len(tempS) > 0 {
			m1[int(i)] = tempS
		}
	}
	return m1
}
func main() {
	arr := []float32 {-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -11.3, -14.0, 5.3, 6.2, -6.6, -7.3}
	fmt.Printf("прогноз погоды %v\n", arr)
	m1 := separator(arr)
	
}