package util

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/sirupsen/logrus"
)

func DeleteOneSecurityRule(client *ecs.Client, securityGroupId, sourceCidrIp, ipProtocol, portRange, priority string) bool {

	request := ecs.CreateRevokeSecurityGroupRequest()
	request.Scheme = "https"
	request.IpProtocol = ipProtocol
	request.SecurityGroupId = securityGroupId
	request.SourceCidrIp = sourceCidrIp
	request.PortRange = portRange
	request.Priority = priority

	response, err := client.RevokeSecurityGroup(request)
	// TODO: 通过response判断操作是否成功
	if err != nil {
		logrus.Error("Get Error when run func: DeleteOneSecurityRule")
		fmt.Print(err.Error())
		return false
	} else {
		logrus.Infof("DeleteOneSecurityRule response is %#v\n", response)
		return true
	}
}

func AddOneSecurityRule(client *ecs.Client, securityGroupId, sourceCidrIp, ipProtocol, portRange, priority, description string) bool {
	request := ecs.CreateAuthorizeSecurityGroupRequest()
	request.Scheme = "https"
	request.SourceCidrIp = sourceCidrIp
	request.SecurityGroupId = securityGroupId
	request.IpProtocol = ipProtocol
	request.PortRange = portRange
	request.Priority = priority
	request.Description = description

	response, err := client.AuthorizeSecurityGroup(request)
	// TODO: 通过response判断操作是否成功
	if err != nil {
		logrus.Error("Get Error when run func: AddOneSecurityRule", err.Error())
		return false
	} else {
		logrus.Infof("AddOneSecurityRule response is %#v\n", response)
		return true
	}
}

func QuerySecurityGroupRule(client *ecs.Client, securityGroupId string) *[]ecs.Permission {
	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.Scheme = "https"
	request.SecurityGroupId = securityGroupId
	response, err := client.DescribeSecurityGroupAttribute(request)
	if err != nil {
		logrus.Error("Query security group rule error. Program will exit.\n", err.Error())
		panic(err)
	}
	return &response.Permissions.Permission
}
