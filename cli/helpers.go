package main

import (
	"fmt"
	"log"
)

func getBaseUrl(company string) string {
	return fmt.Sprintf("https://cliente.apdata.com.br/%s/.net/index.ashx", company)
}

func handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
