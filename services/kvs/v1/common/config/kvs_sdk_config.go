package config

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"github.com/spf13/viper"
)

type KvsSdkConfig struct {
	HeartbeatInterval       int
	HeartbeatThreadMaxCount int
	ReloadInterval          int
	ApiRetryCount           int
	EndPointList            []model.Endpoint
	Credential              model.Credential
	RegionId                string
	CaCertificatePath       string
	ProjectId               string
	ConnectionTimeout       int
	ReadTimeout             int
}

func LoadKvsSdkConfig() *KvsSdkConfig {
	config := KvsSdkConfig{}
	config.HeartbeatInterval = viper.GetInt("kvs.sdk.heartbeat.interval")
	config.HeartbeatThreadMaxCount = viper.GetInt("kvs.sdk.heartbeat.threadMaxCount")
	config.ReloadInterval = viper.GetInt("kvs.sdk.configReloadInterval")
	config.ApiRetryCount = viper.GetInt("kvs.sdk.apiRetryCount")
	endpointNameList := viper.GetStringSlice("kvs.sdk.endpoint.nameList")
	endpointWeightList := viper.GetIntSlice("kvs.sdk.endpoint.weightList")

	endpointList := make([]model.Endpoint, len(endpointNameList))
	for i := 0; i < len(endpointNameList); i++ {
		endpointList[i].Name = endpointNameList[i]
		if i < len(endpointWeightList) {
			endpointList[i].Weight = int32(endpointWeightList[i])
		} else {
			break
		}
	}
	config.EndPointList = endpointList

	var credential model.Credential
	credential.Ak = viper.GetString("kvs.sdk.credential.ak")
	credential.Sk = viper.GetString("kvs.sdk.credential.sk")
	credential.StsToken = viper.GetString("kvs.sdk.credential.stsToken")
	config.Credential = credential

	config.RegionId = viper.GetString("kvs.sdk.regionId")
	config.CaCertificatePath = viper.GetString("kvs.sdk.caCertificatePath")
	config.ProjectId = viper.GetString("kvs.sdk.projectId")
	config.ConnectionTimeout = viper.GetInt("kvs.sdk.connection.timeout")
	config.ReadTimeout = viper.GetInt("kvs.sdk.read.timeout")
	return &config
}
