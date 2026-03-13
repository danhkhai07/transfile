package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main(){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Working directory: %v\n", wd)

	conn, err := ln.Accept()
	if err != nil {
		log.Printf("accept error: %v", err)
		return
	}
	defer conn.Close()
	
	dest, _ := os.Create("test_file")
	defer dest.Close()
	written, err := io.Copy(dest, conn)

	fmt.Printf("Gotten: %v\n", written)
}
