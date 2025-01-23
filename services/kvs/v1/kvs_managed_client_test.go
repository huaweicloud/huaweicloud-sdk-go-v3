package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/common/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestManagedKvsClientInitTest(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	viper.SetConfigFile(currentPath + "/test-kvs-sdk.yaml")
	_ = viper.ReadInConfig()
	viper.Set("kvs.sdk.caCertificatePath", currentPath+"/test-ca.pem")

	sdkConfig := config.LoadKvsSdkConfig()
	mangedKvsClient, _ := NewKvsManagedClient(sdkConfig.EndPointList[0], *sdkConfig)

	assert.NotEmpty(t, mangedKvsClient.KvsClient)
	assert.True(t, mangedKvsClient.IsUsable.Load())
	assert.Equal(t, "endpoint1", mangedKvsClient.Endpoint.Name)
	assert.Equal(t, int32(1), mangedKvsClient.Endpoint.Weight)
}
