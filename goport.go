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
		fmt.Println(err)
		return ""
	}
	var ip string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}
