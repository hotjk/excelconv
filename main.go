package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Println("excelconv template source target")
		return
	}
	template, source, target := os.Args[1], os.Args[2], os.Args[3]
	domain := New(source, template)
	defer domain.Dispose()
	domain.Convert()
	domain.Save(target)
}
