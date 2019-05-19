package lns

import (
	"fmt"
	"log"
	"net"
	"time"
)

var ipFormat string = "/work/ips/%s"

func getLocalIp() (string, error) {
	addrSlice, err := net.InterfaceAddrs()
	var IpAddr string
	if nil != err {
		return "", err
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return IpAddr, nil
			}
		}
	}
	return nil, fmt.Errorf("getLocalIpFaild")

}

func RegisterServer() {
	var ipAddr string
	sentinelAddrs := []string{"redis_sentinel_1:26379", "redis_sentinel_1:26379"}
	sentinel = New("registerServer", sentinelAddrs, 0, 1*time.Duration, 1*time.Duration)
	for {
		for i := 0; i < 2; i++ {
			ipAddr, err := getLocalIp()
			if err == nil {
				break
			}
		}

		_, err := sentinel.Do("master", "SET", fmt.Sprintf(ipFormat, ipAddr), "EX", 5)
		if err != nil {
			log.Printf("requery redis faild")
			break
		}
		time.Sleep(time.Duration(2) * time.Second)

	}
}
