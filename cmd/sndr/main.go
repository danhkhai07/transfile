package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main(){
	fmt.Printf("IP address to destination: ")
	var ip string
	fmt.Scanf("%s", &ip)
	fmt.Printf("Path to file: ")
	var path string
	fmt.Scanf("%s", &path)

	conn, err := net.Dial("tcp", ip + ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	written, err := io.Copy(conn, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Written: %v\n", written)
}
