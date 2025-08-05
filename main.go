package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("📥 Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("📨 Received: %s\n", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Error reading from connection: %v\n", err)
	}

	fmt.Println("🔌 Client disconnected:", conn.RemoteAddr())
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("🚀 Listening on TCP port 808...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("❌ Failed to accept connection: %v\n", err)
			continue
		}
		go handleConnection(conn) // handle multiple clients concurrently
	}
}
