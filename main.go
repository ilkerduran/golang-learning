package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var arr [][]uint8 
	for i:=0; i<dx; i++ {
		for j:=0; i<dy; j++ {
			arr[i][j] = uint8(i*j)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic(2,2))
}
