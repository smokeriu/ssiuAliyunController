package util

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func InitNewClient(region, accessKey, accessSecret string) (*ecs.Client, error) {
	return ecs.NewClientWithAccessKey(region, accessKey, accessSecret)
}
