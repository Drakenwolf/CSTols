package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter target: ")
		input, _ := reader.ReadString('\n')
		target := strings.TrimSpace(input)

		if target == "exit" {
			fmt.Println("Exiting")
			os.Exit(0)
		}

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

		fmt.Println("Scan completed for", target)
	}

}
