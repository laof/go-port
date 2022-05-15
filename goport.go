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

func GetRandomCode() string {
	rand.Seed(time.Now().UnixNano())
	i := time.Now().UnixNano() - rand.Int63n(3000)
	return strconv.FormatInt(i, 10)
}
