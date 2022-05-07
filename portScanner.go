package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

func main() {
	ItalicCyan := color.New(color.FgCyan).Add(color.Italic).PrintFunc()
	BoldUnderlinedMagenta := color.New(color.FgMagenta).Add(color.Underline).Add(color.Bold).PrintlnFunc()
	BoldGreen := color.New(color.FgGreen).Add(color.Bold).PrintfFunc()
	BoldYellow := color.New(color.FgYellow).Add(color.Bold).PrintFunc()

	BoldUnderlinedMagenta("Simple Port Scanner With Go-Lang")

	var UserInput string
	ItalicCyan("Enter your URL/IP address: ")
	fmt.Scan(&UserInput)

	var wg sync.WaitGroup
	now := time.Now()
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := UserInput + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout("tcp", address, time.Millisecond*40)
			if err != nil {
				return
			}
			conn.Close()
			BoldGreen("Port %d is open\n", port)

		}(port)
		time.Sleep(7 * time.Millisecond)
	}
	wg.Wait()
	totalTime := time.Since(now)
	BoldYellow("Scan finished.Total Time taken : ")
	color.Red(totalTime.String())
}
