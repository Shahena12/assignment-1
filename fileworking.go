package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	file, err := os.Create("sha.txt")
	file.WriteString("shahena batch11 software trainee engineer")
	file.Close()

	if err != nil {
		log.Fatal(err)
	}

}

func ReadFile() {

	file, err := ioutil.ReadFile("sha.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
func main() {
	CreateFile()
	ReadFile()

}
