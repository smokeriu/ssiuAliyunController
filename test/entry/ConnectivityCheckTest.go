package main

import "ssiuAliyunController/util"

func main() {
	result := util.ConnectivityCheck("tcp", "hadoop", "8088")
	println(result)
}
