package main

import (
	"net"
	"os"
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
	if len(os.Args) < 2 {
		color.Red("Missing target parameter")
		os.Exit(0)
	} else {
		ItalicCyan("Target : ")
		color.Red(os.Args[1])
	}
	var wg sync.WaitGroup
	now := time.Now()
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := os.Args[1] + ":" + strconv.Itoa(port)
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
