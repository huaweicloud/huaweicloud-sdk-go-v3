package config

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var sdkConfig KvsSdkConfig

func setup() {
	currentPath, err := os.Getwd()
	viper.SetConfigFile(currentPath + "/../test-kvs-sdk.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("error reading config file: %v\n", err)
	}
}

func TestGetAKSKConfig(t *testing.T) {
	setup()
	sdkConfig := LoadKvsSdkConfig()
	assert.Equal(t, "testAk", sdkConfig.Credential.Ak)
	assert.Equal(t, "testSk", sdkConfig.Credential.Sk)
	assert.Equal(t, "testStsToken", sdkConfig.Credential.StsToken)
	assert.Equal(t, "testProjectId", sdkConfig.ProjectId)
}

func TestGetRegionConfig(t *testing.T) {
	setup()
	sdkConfig := LoadKvsSdkConfig()
	assert.Equal(t, "cn-north-7", sdkConfig.RegionId)

	endpoint1 := model.Endpoint{Name: "endpoint1", Weight: 1}
	endpoint2 := model.Endpoint{Name: "endpoint2", Weight: 3}
	endpoint3 := model.Endpoint{Name: "endpoint3", Weight: 2}
	endpointList := []model.Endpoint{endpoint1, endpoint2, endpoint3}

	for i := 0; i < len(sdkConfig.EndPointList); i++ {
		assert.Equal(t, endpointList[i].Name, sdkConfig.EndPointList[i].Name)
		assert.Equal(t, endpointList[i].Weight, sdkConfig.EndPointList[i].Weight)
	}
}

func TestGetClientConfig(t *testing.T) {
	setup()
	sdkConfig := LoadKvsSdkConfig()
	assert.Equal(t, "", sdkConfig.CaCertificatePath)
	assert.Equal(t, 10, sdkConfig.ConnectionTimeout)
	assert.Equal(t, 1, sdkConfig.HeartbeatInterval)
	assert.Equal(t, 5, sdkConfig.ThreadMaxCount)
	assert.Equal(t, 3, sdkConfig.ReloadInterval)
	assert.Equal(t, 3, sdkConfig.ApiRetryCount)
}
