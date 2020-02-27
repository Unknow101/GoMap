package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	host := flag.String("host", "false", "Host to scan")
	full := flag.Bool("f", false, "Scan all 65535 ports")
	flag.Parse()
	if *host == "false" {
		fmt.Println("Usage : You need to specify target")
		return
	}
	scan := 1024
	if *full == true {
		scan = 65535
	}
	start := time.Now()
	myFigure := figure.NewFigure("GoMap", "rectangles", true)
	myFigure.Print()
	fmt.Println("v0.2 version")
	fmt.Println("==========")
	fmt.Println("Scanning " + *host + " for " + strconv.Itoa(scan) + " ports..")
	fmt.Println("==========")
	var wg sync.WaitGroup
	addr := *host
	for i := 1; i <= scan; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf(addr+":%d", j))
			if err != nil {
				return
			}
			fmt.Printf("%d Port open\n", j)
			conn.Close()
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Scan done in %s seconds", elapsed)
}
