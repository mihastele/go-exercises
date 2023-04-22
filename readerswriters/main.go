package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	reader := strings.NewReader("SAMPLE")
	var newString strings.Builder

	buffer := make([]byte, 4)
	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes %c\n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%v\n", newString.String())

	source := strings.NewReader("SAMPLE")
	buffered := bufio.NewReader(source)
	newString1, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString1)
	} else {
		fmt.Println(err)
	}
}
