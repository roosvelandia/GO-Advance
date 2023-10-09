package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 100; i++ {
		// scan port if exist
		// open or closed
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Puerto : &d open \n", i)

	}
}
