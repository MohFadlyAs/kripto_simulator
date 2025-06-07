package utils

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("")
}

func TungguUser() {
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln()
}
