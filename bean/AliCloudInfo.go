package bean

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type AliCloudInfoEntity struct {
	AccountInfo AliCloudData `yaml:"account"`
	ServerInfo  ServerData   `yaml:"server"`
}

type AliCloudData struct {
	Region          string `yaml:"region"`
	AccessKey       string `yaml:"accessKey"`
	AccessSecret    string `yaml:"accessSecret"`
	SecurityGroupId string `yaml:"securityGroupId"`
}

type ServerData struct {
	Host string `yaml:"host"`
}

func (c *AliCloudInfoEntity) GetConf(confPath string) *AliCloudInfoEntity {
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
