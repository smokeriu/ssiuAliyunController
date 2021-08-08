package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

// GetExternalIp
// @description 通过ipw.cn.得到运行主机的外网IP.如果启用了网络代理可能会影响结果
func GetExternalIp() string {
	responseClient, errClient := http.Get("https://ipw.cn/api/ip/myip") // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		panic(errClient)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(responseClient.Body)

	body, _ := ioutil.ReadAll(responseClient.Body)
	return fmt.Sprintf("%s", string(body))
}

// ConnectivityCheck
// @description 检测单个端口的连通性
func ConnectivityCheck(network, ip, port string) bool {
	address := net.JoinHostPort(ip, port)
	// Just check connectivity. So we don't Need to handle the error
	conn, _ := net.DialTimeout(network, address, 10*time.Second)
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	if conn != nil {
		fmt.Println("")
		return true
	} else {
		return false
	}
}

// PortRangeCheck
// @description 检测端口范围的连通性
// @return true -> All ports can connect. false -> One port can't connect
func PortRangeCheck(destIp, ipProtocol string, startPort, endPort int) bool {
	// TODO: parallel run this
	for i := startPort; i < endPort; i++ {
		isSuccess := ConnectivityCheck(ipProtocol, destIp, strconv.Itoa(i))
		if !isSuccess {
			fmt.Printf("Check ip: %s, port: %d. Failed. Will Change to new Ip", destIp, i)
			return false
		}
		i++
	}
	return true
}
