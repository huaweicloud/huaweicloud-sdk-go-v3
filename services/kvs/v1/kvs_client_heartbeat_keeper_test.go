package v1

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/common/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestHeartbeatKeeperAllSuccess(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	viper.SetConfigFile(currentPath + "/test-kvs-sdk.yaml")
	_ = viper.ReadInConfig()
	viper.Set("kvs.sdk.caCertificatePath", currentPath+"/test-ca.pem")

	sdkConfig := config.LoadKvsSdkConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := NewMockIKvsClient(ctrl)
	client.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()

	managedClient1 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)}}

	managedClient2 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)}}

	managedClient3 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)}}

	KvsClientMap := sync.Map{}
	KvsClientMap.Store(int32(1), &managedClient1)
	KvsClientMap.Store(int32(2), &managedClient2)
	KvsClientMap.Store(int32(3), &managedClient3)

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&KvsClientMap, sdkConfig.HeartbeatInterval, sdkConfig.HeartbeatThreadMaxCount, ctx)

	time.Sleep(5 * time.Second)

	assert.True(t, managedClient1.IsUsable.Load())
	assert.True(t, managedClient2.IsUsable.Load())
	assert.True(t, managedClient3.IsUsable.Load())
	cancel()
	time.Sleep(5 * time.Second)
}

func TestHeartbeatKeeperAllFailed(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	viper.SetConfigFile(currentPath + "/test-kvs-sdk.yaml")
	_ = viper.ReadInConfig()
	viper.Set("kvs.sdk.caCertificatePath", currentPath+"/test-ca.pem")

	sdkConfig := config.LoadKvsSdkConfig()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 mock 对象
	client := NewMockIKvsClient(ctrl)
	err = errors.New("mock socket timeout error")
	client.EXPECT().CheckHealth(gomock.Any()).Return(nil, err).AnyTimes()

	managedClient1 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)}}

	managedClient2 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)}}

	managedClient3 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)}}

	KvsClientMap := sync.Map{}
	KvsClientMap.Store(int32(1), &managedClient1)
	KvsClientMap.Store(int32(2), &managedClient2)
	KvsClientMap.Store(int32(3), &managedClient3)

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&KvsClientMap, sdkConfig.HeartbeatInterval, sdkConfig.HeartbeatThreadMaxCount, ctx)

	time.Sleep(5 * time.Second)

	assert.False(t, managedClient1.IsUsable.Load())
	assert.False(t, managedClient2.IsUsable.Load())
	assert.False(t, managedClient3.IsUsable.Load())
	cancel()
	time.Sleep(5 * time.Second)
}
