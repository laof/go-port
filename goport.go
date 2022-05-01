package goport

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func InputPort(str string) string {

	fmt.Print("Port default is " + str + ", Input:")

	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		port := in.Text()
		if len(port) >= 4 {
			str = port
		}
	}

	port := ":" + str

	fmt.Println("http://localhost" + port)
	fmt.Println("http://" + GetIP() + port)

	return str
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String()
	}
	return ""

}
