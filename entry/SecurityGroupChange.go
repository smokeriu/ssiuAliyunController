package main

import (
	"flag"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/sirupsen/logrus"
	"ssiuAliyunController/bean"
	"ssiuAliyunController/util"
	"strconv"
	"strings"
)

func main() {
	// init and get config
	configFilePath := flag.String("configPath", "resources/config.yaml", "Get config file path")
	flag.Parse()
	aliCloudEntity := bean.AliCloudInfoEntity{}
	aliCloudInfo := aliCloudEntity.GetConf(*configFilePath)
	accountInfo := aliCloudInfo.AccountInfo
	serverInfo := aliCloudInfo.ServerInfo

	// init alicloud client
	client, _ := util.InitNewClient(accountInfo.Region, accountInfo.AccessKey, accountInfo.AccessSecret)
	// query security group info by SecurityGroupId
	securityGroupRule := util.QuerySecurityGroupRule(client, accountInfo.SecurityGroupId)
	// get current id and useful Permissions
	currentIp := util.GetExternalIp()
	usefulPermissions := getUsefulPermissions(*securityGroupRule, currentIp)

	// Change security ip if disconnected
	changeSecurityIpByResponseIfDisconnected(client, serverInfo.Host, accountInfo.SecurityGroupId, currentIp, &usefulPermissions)
}

// ChangeSecurityIpByResponseIfNeed
// @description 如果连接不通,则修改安全组的源ip
func changeSecurityIpByResponseIfDisconnected(client *ecs.Client, destIp, securityGroupId, newSourceCidrIp string, responses *[]ecs.Permission) {
	for _, respons := range *responses {

		portRange := respons.PortRange
		ipProtocol := respons.IpProtocol

		// check port Connectivity.
		split := strings.Split(portRange, "/")
		startPort, _ := strconv.Atoi(split[0])
		endPort, _ := strconv.Atoi(split[1])

		// 对于无公网ip的用户，可能安全组配置的ip与currentIp不同但仍然能够访问。这种时候不需要修改Ip
		if !util.PortRangeCheck(destIp, ipProtocol, startPort, endPort) {
			// Add Rule only when delete success
			sourceCidrIp := respons.SourceCidrIp
			priority := respons.Priority
			if util.DeleteOneSecurityRule(client, securityGroupId, sourceCidrIp, ipProtocol, portRange, priority) {
				description := respons.Description
				util.AddOneSecurityRule(client, securityGroupId, newSourceCidrIp, ipProtocol, portRange, priority, description)
			}
		} else {
			logrus.Infof("All ports[%d,%d] of the %s can be connected through the %s protocol.\tNo change will happen\n", startPort, endPort, destIp, ipProtocol)
		}
	}
}

// getUsefulPermissions
// @description 移除一些我不关心的permission信息
func getUsefulPermissions(permissions []ecs.Permission, currentIp string) []ecs.Permission {
	index := 0
	for _, permission := range permissions {
		if permission.Description != "挖矿" &&
			permission.SourceCidrIp != "" &&
			permission.SourceCidrIp != "0.0.0.0/0" &&
			permission.SourceCidrIp != currentIp &&
			permission.Direction == "ingress" {
			permissions[index] = permission
			index++
		}
	}
	return permissions[:index]
}
