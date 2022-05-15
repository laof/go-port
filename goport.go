package goport

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func InputPort(str string) string {

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

func GetRandomCode() string {
	rand.Seed(time.Now().UnixNano())
	i := time.Now().UnixNano() - rand.Int63n(3000)
	return strconv.FormatInt(i, 10)
}
