package v1

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

func TestReloadEndpointInsert(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	viper.SetConfigFile(configFilePath)
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)

	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2", "endpoint3", "endpoint4"}, []int{1, 3, 2, 1})

	time.Sleep(5 * time.Second)

	assert.Equal(t, int32(4), getMapSize(&multiChannelClient.KvsClientMap))
	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2", "endpoint3"}, []int{1, 3, 2})
}

func TestReloadEndpointRemove(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	viper.SetConfigFile(configFilePath)
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)

	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2"}, []int{1, 3})

	time.Sleep(5 * time.Second)

	assert.Equal(t, int32(3), getMapSize(&multiChannelClient.KvsClientMap))
	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2", "endpoint3"}, []int{1, 3, 2})
}

func TestReloadAkskAndInsertEndpoint(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	viper.SetConfigFile(configFilePath)
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)

	setConfigFile("testAk2", "testSk2", "testStsToken2", []string{"endpoint1", "endpoint2", "endpoint3", "endpoint4"}, []int{1, 3, 2, 1})

	time.Sleep(5 * time.Second)

	assert.Equal(t, int32(4), getMapSize(&multiChannelClient.KvsClientMap))

	value, _ := multiChannelClient.KvsClientMap.Load(int32(1))
	managedKvsClient, ok := value.(*KvsManagedClient)
	if !ok {
		log.Panicln("cast value of KvsClientMap to *KvsManagedClient failed!")
	}
	assert.Equal(t, "testAk2", managedKvsClient.Credential.Ak)
	assert.Equal(t, "testSk2", managedKvsClient.Credential.Sk)
	assert.Equal(t, "testStsToken2", managedKvsClient.Credential.StsToken)
	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2", "endpoint3"}, []int{1, 3, 2})
}

func setConfigFile(Ak, Sk, StsToken string, endpointNameList []string, endpointWeightList []int) {
	viper.Set("kvs.sdk.heartbeat.interval", 1)
	viper.Set("kvs.sdk.heartbeat.threadMaxCount", 3)
	viper.Set("kvs.sdk.configReloadInterval", 2)
	viper.Set("kvs.sdk.apiRetryCount", 3)
	viper.Set("kvs.sdk.credential.ak", Ak)
	viper.Set("kvs.sdk.credential.sk", Sk)
	viper.Set("kvs.sdk.credential.stsToken", StsToken)
	viper.Set("kvs.sdk.regionId", "cn-north-7")
	viper.Set("kvs.sdk.caCertificatePath", "")
	viper.Set("kvs.sdk.projectId", "testProjectId")
	viper.Set("kvs.sdk.connection.timeout", 3)
	viper.Set("kvs.sdk.endpoint.nameList", endpointNameList)
	viper.Set("kvs.sdk.endpoint.weightList", endpointWeightList)
	err := viper.WriteConfig()
	if err != nil {
		log.Printf("error write config file: %v\n", err)
	}
}
