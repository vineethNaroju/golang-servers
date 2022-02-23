package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReeadFile() {
	content, err := ioutil.ReadFile("../main.go")

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(content))
}

func OsOpenRead() {
	file, err := os.Open("../main.go")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := make([]byte, 16)

	for {
		bytesCount, err := file.Read(buf)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		if bytesCount > 0 {
			fmt.Println(string(buf[:bytesCount]))
			fmt.Println("-------------------------------------------------------------------------------")
		}

	}

	fmt.Println("FILE READING DONE")
}

func WriteeeeFile() {
	// if err := os.WriteFile("./tmp.txt", []byte("Yo bro"), fs.FileMode(os.O_CREATE|os.O_APPEND)); err != nil {
	// 	log.Fatal(err)
	// }

	file, err := os.Create("./yo.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	cnt, err := file.WriteString("\nbruh........")

	fmt.Println(cnt, err)
}

func main() {
	// OsOpenRead()

	WriteeeeFile()
}
