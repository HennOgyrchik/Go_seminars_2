package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running")

	con, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("line recieved: ", line)

		upperline := strings.ToUpper(line)
		if _, err := con.Write([]byte(upperline)); err != nil {
			log.Fatal(err)
		}
	}
}
