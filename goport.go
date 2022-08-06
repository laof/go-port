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

	for _, val := range GetIP() {
		fmt.Println("http://" + val + port)
	}

	return str
}

func GetIP() []string {
	ips := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
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
		ips = append(ips, ipAddr.IP.String())
	}
	return ips

}

func GetRandomCode() string {
	rand.Seed(time.Now().UnixNano())
	i := time.Now().UnixNano() - rand.Int63n(3000)
	return strconv.FormatInt(i, 10)
}
