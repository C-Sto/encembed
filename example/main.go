package main

import "log"

func main() {
	//embeds main as an encrypted resource
	log.Println(string(embedded()))
}

//go:generate go run github.com/c-sto/encembed -i main.go
