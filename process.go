package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func lab3(pathFrom string, pathTo string, fileName string) {
	s := strings.Split(fileName, ".")
	newFile := s[0] + ".res"
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if ("./"+file.Name() != pathTo) && !file.IsDir() { //mkdir
			err := os.MkdirAll(pathTo, 0750)
			if err != nil {
				log.Printf("%v", err)
			}
		}
	}

	file, err := os.Open(pathFrom + "/" + fileName) //open
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 1)
	count := 0
	str := ""
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str += string(data[:n])
	}

	for i := 0; i < len(str); i++ {
		if (string(str[i]) == " ") && (string(str[i-1]) == "." || string(str[i-1]) == "!" || string(str[i-1]) == "?") {
			count++
		}
		if ((i + 1) == len(str)) && (string(str[i]) == "." || string(str[i]) == "!" || string(str[i]) == "?") {
			count++
		}
	}
	text := strconv.Itoa(count)
	fileTo, err := os.Create(pathTo + "/" + newFile) //write

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer fileTo.Close()
	fileTo.WriteString(text)
}

func main() {
	pathFrom := os.Args[1]
	pathTo := os.Args[2]

	i := 0
	files, err := ioutil.ReadDir(pathFrom)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		go lab3(pathFrom, pathTo, file.Name())
		i++
	}
	fmt.Print("Total number of processed files: ", i)
	var tmp string
	fmt.Scanln(&tmp)
}
