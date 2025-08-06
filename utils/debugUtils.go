package utils

import (
	"fmt"
	"os"
)

func DebugPrint(message string){
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf(message)
	}
}