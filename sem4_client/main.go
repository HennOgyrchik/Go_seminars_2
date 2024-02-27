package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	d, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if _, err = d.Write([]byte(text)); err != nil {
			log.Fatal(err)
		}

		text, err = bufio.NewReader(d).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(text)
	}
}
