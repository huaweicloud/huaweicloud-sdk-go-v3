package v1

import (
	"github.com/fsnotify/fsnotify"
	config "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/common/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"github.com/spf13/viper"
	"sync"
	"time"
)

type KvsSdkConfigManager struct {
	config           *config.KvsSdkConfig
	credential       model.Credential
	endpoints        []model.Endpoint
	KvsClientMap     *sync.Map
	OldKvsClientMaps *sync.Map
	lock             sync.Mutex
}

func NewKvsSdkConfigManager(configFilePath string, KvsClientMap *sync.Map, OldKvsClientMaps *sync.Map) (*config.KvsSdkConfig, error) {
	var configManager KvsSdkConfigManager
	configManager.KvsClientMap = KvsClientMap
	configManager.OldKvsClientMaps = OldKvsClientMaps
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = loadKvsSdkConfig(&configManager)
	if err != nil {
		return nil, err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = loadKvsSdkConfig(&configManager)
	})
	return configManager.config, nil
}

func loadKvsSdkConfig(configManager *KvsSdkConfigManager) error {
	configManager.lock.Lock()
	defer configManager.lock.Unlock()
	newSdkConfig := config.LoadKvsSdkConfig()
	reloadAKSKConfig(configManager, *newSdkConfig)
	err := reloadEndPointConfig(configManager, *newSdkConfig)
	if err != nil {
		return err
	}
	configManager.config = newSdkConfig
	return nil
}

func reloadAKSKConfig(configManager *KvsSdkConfigManager, newSdkConfig config.KvsSdkConfig) {
	if configManager.credential.Ak != newSdkConfig.Credential.Ak ||
		configManager.credential.Sk != newSdkConfig.Credential.Sk {
		configManager.credential = newSdkConfig.Credential

		newKvsClientMap := configManager.KvsClientMap
		configManager.OldKvsClientMaps.Store(time.Now().UnixMilli(), newKvsClientMap)
		configManager.KvsClientMap.Range(func(key, value interface{}) bool {
			configManager.KvsClientMap.Delete(key)
			return true
		})
		configManager.endpoints = make([]model.Endpoint, 0)
	}
}

func reloadEndPointConfig(configManager *KvsSdkConfigManager, newSdkConfig config.KvsSdkConfig) error {
	for _, newEndpoint := range newSdkConfig.EndPointList {
		if !contains(configManager.endpoints, newEndpoint) {
			managedClient, err := NewKvsManagedClient(newEndpoint, newSdkConfig)
			if err != nil {
				return err
			}
			configManager.endpoints = append(configManager.endpoints, newEndpoint)
			configManager.KvsClientMap.Store(int32(len(configManager.endpoints)), managedClient)
		}
	}
	return nil
}

func contains(slice []model.Endpoint, target model.Endpoint) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
