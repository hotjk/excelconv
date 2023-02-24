package main

import (
	"log"
	"os"

	"github.com/hotjk/excelconv/conv"
)

func main() {
	if len(os.Args) != 4 {
		log.Println("excelconv template source target")
		return
	}
	template, source, target := os.Args[1], os.Args[2], os.Args[3]

	domain := conv.New(source, template)
	defer domain.Dispose()
	domain.Convert()
	domain.Save(target)
}
