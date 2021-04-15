package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crthpl/go2v/convert"
)

func main() {
	res, err := convert.Convert(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
