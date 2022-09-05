package main

import (
	"fmt"

	"github.com/nikolazifra/go-samples/hello1034"
)

func main() {
	fmt.Println(hello1034.Hello())

	for i, ch := range "Japan 日本" {
		fmt.Printf("%d:%q ", i, ch)
	}
}
