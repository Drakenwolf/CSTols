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

func scanPort(port int, target string) {

	timeout := 1 * time.Second

	conn, err := net.DialTimeout("tcp", target+":"+strconv.Itoa(port), timeout)
	if err != nil {
		// fmt.Println("Port", port, "is closed on", target)
	} else {
		conn.Close()
		fmt.Println("Port", port, "is open on", target)
	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	const MAX_PORT int = 65535
	var ports []int

	for {

		fmt.Print("1. Scan all ports 2. Scan specific ports\nEnter option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		fmt.Print("Enter target: ")
		input, _ := reader.ReadString('\n')
		target := strings.TrimSpace(input)

		if target == "exit" {
			fmt.Println("Exiting")
			os.Exit(0)
		}
		if option == "1" {
			for port := 0; port < MAX_PORT; port++ {
				scanPort(port, target)
			}

		} else if option == "2" {
			fmt.Print("Enter ports separated by comma: ")
			input, _ := reader.ReadString('\n')
			portslist := strings.Split(strings.TrimSpace(input), ",")

			for _, portstr := range portslist {
				port, _ := strconv.Atoi(portstr)
				ports = append(ports, port)
			}
			for _, port := range ports {
				scanPort(port, target)
			}

		}
		fmt.Println("Scan completed for", target)

	}

}
