package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./portscanner <target>")
		os.Exit(1)
	}

	target := os.Args[1]

	ports := []int{80, 443, 8080}
	timeout := 1 * time.Second

	for _, port := range ports {
		conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), timeout)
		if err != nil {
			fmt.Println("Port", port, "is closed on", target)
		} else {
			conn.Close()
			fmt.Println("Port", port, "is open on", target)
		}
	}

	fmt.Println("Scan completed for ports", ports, "on host", target)

}
