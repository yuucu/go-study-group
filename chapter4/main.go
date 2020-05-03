package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// go-cutコマンドを実装しよう
func main() {

	delimiter := flag.String("d", ",", "区切り文字を指定してください")
	field := flag.Int("f", 1, "")

	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Usege: go-cut [-d delimiter] file_path")
		os.Exit(1)
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, *delimiter)[*field-1]
		fmt.Println(s)
	}

}
