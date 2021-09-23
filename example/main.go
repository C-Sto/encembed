package main

import "log"

func main() {
	//embeds main as an encrypted resource
	log.Println(string(cats))
}

//go:generate go run github.com/c-sto/encembed -i main.go -decvarname cats
