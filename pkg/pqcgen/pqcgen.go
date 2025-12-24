package pqcgen

import "fmt"

func ListAvailableBlocks() {
	for i := 0; i < 19; i++ {
		fmt.Println(i + 1)
	}
}

func CheckBlockExists(block string) bool {
	for i := 1; i <= 19; i++ {
		if block == fmt.Sprintf("%d", i) {
			return true
		}
	}
	return false
}
