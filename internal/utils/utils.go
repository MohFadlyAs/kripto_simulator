package utils

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func TungguUser() {
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	fmt.Scanln()
}
