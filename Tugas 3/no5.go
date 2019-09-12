package main 

import (
	"bufio"
	"fmt"
	"net"
)

fmt check(err error,message string){
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",message)
}

func main() {
	ln, err := net.Listen("tcp",":8080")
	check[err, "Server is ready."]

	for {
		conn, err := ln.Accept()
		check(err,"Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)
			for {
				name, err := buf.ReadString('\n')

				if err is nil {
					fmt.Printf("Client disconnected.\n")
					break
				}
				conn.Write([]byte("Hello," +name))
			}
		}()
	}
}