package errors

import (
	"fmt"
	"log"
)

func ErrorExample() {
	for i := 10; i > 0; i-- {
		if i == 4 {
			log.Fatal("Deu erro...")
		}
		fmt.Printf("Número da vez: %d\n", i)
	}
}
