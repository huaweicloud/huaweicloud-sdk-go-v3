package v1

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestOldClientMapClean(t *testing.T) {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	multiChannelClient, _ := NewKvsMultiChannelClient(configFilePath)

	setConfigFile("testAk2", "testSk2", "testStsToken2",
		[]string{"endpoint1", "endpoint2", "endpoint3", "endpoint4"},
		[]int{1, 3, 2, 1})

	time.Sleep(20 * time.Second)

	assert.Equal(t, int32(0), getMapSize(&multiChannelClient.OldKvsClientMaps))
	setConfigFile("testAk", "testSk", "testStsToken", []string{"endpoint1", "endpoint2", "endpoint3"}, []int{1, 3, 2})
	multiChannelClient.Close()
	time.Sleep(5 * time.Second)
}

func TestThreeClientCreateTable(t *testing.T) {
	var clientCount1, clientCount2, clientCount3 int

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client1 := NewMockIKvsClient(ctrl)
	client1.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client1.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount1++ }).AnyTimes()
	managedClient1 := KvsManagedClient{
		KvsClient: client1,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)},
		IsUsable:  atomic.Bool{}}
	managedClient1.IsUsable.Store(true)

	client2 := NewMockIKvsClient(ctrl)
	client2.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client2.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount2++ }).AnyTimes()
	managedClient2 := KvsManagedClient{
		KvsClient: client2,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)},
		IsUsable:  atomic.Bool{}}
	managedClient2.IsUsable.Store(true)

	client3 := NewMockIKvsClient(ctrl)
	client3.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client3.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount3++ }).AnyTimes()
	managedClient3 := KvsManagedClient{
		KvsClient: client3,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)},
		IsUsable:  atomic.Bool{}}
	managedClient3.IsUsable.Store(true)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.tokenNum = 1
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)
	newKvsClientMap := sync.Map{}
	newKvsClientMap.Store(int32(1), &managedClient1)
	newKvsClientMap.Store(int32(2), &managedClient2)
	newKvsClientMap.Store(int32(3), &managedClient3)
	multiChannelClient.KvsClientMap = newKvsClientMap

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&multiChannelClient.KvsClientMap, multiChannelClient.SdkConfig.HeartbeatInterval, multiChannelClient.SdkConfig.HeartbeatThreadMaxCount, ctx)
	go cleanOldKvsClientMap(&multiChannelClient.OldKvsClientMaps, int64(multiChannelClient.SdkConfig.ConnectionTimeout), ctx)

	request := &model.CreateTableRequest{}
	for i := 0; i < 7; i++ {
		multiChannelClient.CreateTable(request)
	}
	assert.Equal(t, 2, clientCount1)
	assert.Equal(t, 3, clientCount2)
	assert.Equal(t, 2, clientCount3)
	cancel()
	time.Sleep(5 * time.Second)
}

func TestThreeClientCreateTableWith1Unusable(t *testing.T) {
	var clientCount1, clientCount2, clientCount3 int

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client1 := NewMockIKvsClient(ctrl)
	client1.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client1.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount1++ }).AnyTimes()
	managedClient1 := KvsManagedClient{
		KvsClient: client1,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)},
		IsUsable:  atomic.Bool{}}
	managedClient1.IsUsable.Store(true)

	client2 := NewMockIKvsClient(ctrl)
	err := errors.New("mock socket timeout error")
	client2.EXPECT().CheckHealth(gomock.Any()).Return(nil, err).AnyTimes()
	client2.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount2++ }).AnyTimes()
	managedClient2 := KvsManagedClient{
		KvsClient: client2,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)},
		IsUsable:  atomic.Bool{}}
	managedClient2.IsUsable.Store(true)

	client3 := NewMockIKvsClient(ctrl)
	client3.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client3.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount3++ }).AnyTimes()
	managedClient3 := KvsManagedClient{
		KvsClient: client3,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)},
		IsUsable:  atomic.Bool{}}
	managedClient3.IsUsable.Store(true)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.tokenNum = 1
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)
	newKvsClientMap := sync.Map{}
	newKvsClientMap.Store(int32(1), &managedClient1)
	newKvsClientMap.Store(int32(2), &managedClient2)
	newKvsClientMap.Store(int32(3), &managedClient3)
	multiChannelClient.KvsClientMap = newKvsClientMap

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&multiChannelClient.KvsClientMap, multiChannelClient.SdkConfig.HeartbeatInterval, multiChannelClient.SdkConfig.HeartbeatThreadMaxCount, ctx)
	go cleanOldKvsClientMap(&multiChannelClient.OldKvsClientMaps, int64(multiChannelClient.SdkConfig.ConnectionTimeout), ctx)

	time.Sleep(2 * time.Second)

	request := &model.CreateTableRequest{}
	for i := 0; i < 7; i++ {
		multiChannelClient.CreateTable(request)
	}
	assert.Equal(t, 3, clientCount1)
	assert.Equal(t, 0, clientCount2)
	assert.Equal(t, 4, clientCount3)
	cancel()
	time.Sleep(5 * time.Second)
}

func TestThreeClientCreateTableWith1UnusableInterval(t *testing.T) {
	var clientCount1, clientCount2, clientCount3 int

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client1 := NewMockIKvsClient(ctrl)
	err := errors.New("mock socket timeout error")
	client1.EXPECT().CheckHealth(gomock.Any()).Return(nil, err).AnyTimes()
	client1.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount1++ }).AnyTimes()
	managedClient1 := KvsManagedClient{
		KvsClient: client1,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)},
		IsUsable:  atomic.Bool{}}
	managedClient1.IsUsable.Store(true)

	client2 := NewMockIKvsClient(ctrl)
	client2.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client2.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount2++ }).AnyTimes()
	managedClient2 := KvsManagedClient{
		KvsClient: client2,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)},
		IsUsable:  atomic.Bool{}}
	managedClient2.IsUsable.Store(true)

	client3 := NewMockIKvsClient(ctrl)
	client3.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	client3.EXPECT().CreateTable(gomock.Any()).Do(func(request *model.CreateTableRequest) { clientCount3++ }).AnyTimes()
	managedClient3 := KvsManagedClient{
		KvsClient: client3,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)},
		IsUsable:  atomic.Bool{}}
	managedClient3.IsUsable.Store(true)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.tokenNum = 1
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)
	newKvsClientMap := sync.Map{}
	newKvsClientMap.Store(int32(1), &managedClient1)
	newKvsClientMap.Store(int32(2), &managedClient2)
	newKvsClientMap.Store(int32(3), &managedClient3)
	multiChannelClient.KvsClientMap = newKvsClientMap

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&multiChannelClient.KvsClientMap, multiChannelClient.SdkConfig.HeartbeatInterval, multiChannelClient.SdkConfig.HeartbeatThreadMaxCount, ctx)
	go cleanOldKvsClientMap(&multiChannelClient.OldKvsClientMaps, int64(multiChannelClient.SdkConfig.ConnectionTimeout), ctx)

	request := &model.CreateTableRequest{}
	for i := 0; i < 6; i++ {
		multiChannelClient.CreateTable(request)
	}
	assert.Equal(t, 1, clientCount1)
	assert.Equal(t, 3, clientCount2)
	assert.Equal(t, 2, clientCount3)

	time.Sleep(2 * time.Second)

	for i := 0; i < 6; i++ {
		multiChannelClient.CreateTable(request)
	}
	assert.Equal(t, 1, clientCount1)
	assert.Equal(t, 7, clientCount2)
	assert.Equal(t, 4, clientCount3)
	cancel()
	time.Sleep(5 * time.Second)
}

func TestAllApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := NewMockIKvsClient(ctrl)
	client.EXPECT().CheckHealth(gomock.Any()).Return(nil, nil).AnyTimes()
	managedClient1 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint1", Weight: int32(1)},
		IsUsable:  atomic.Bool{}}
	managedClient1.IsUsable.Store(true)

	managedClient2 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint2", Weight: int32(3)},
		IsUsable:  atomic.Bool{}}
	managedClient2.IsUsable.Store(true)

	managedClient3 := KvsManagedClient{
		KvsClient: client,
		Endpoint:  model.Endpoint{Name: "endpoint3", Weight: int32(2)},
		IsUsable:  atomic.Bool{}}
	managedClient3.IsUsable.Store(true)

	currentPath, err := os.Getwd()
	if err != nil {
		log.Panicln("get currentPath failed!")
	}
	configFilePath := currentPath + "/test-kvs-sdk.yaml"
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.tokenNum = 1
	multiChannelClient.SdkConfig, _ = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)
	newKvsClientMap := sync.Map{}
	newKvsClientMap.Store(int32(1), &managedClient1)
	newKvsClientMap.Store(int32(2), &managedClient2)
	newKvsClientMap.Store(int32(3), &managedClient3)
	multiChannelClient.KvsClientMap = newKvsClientMap

	ctx, cancel := context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&multiChannelClient.KvsClientMap, multiChannelClient.SdkConfig.HeartbeatInterval, multiChannelClient.SdkConfig.HeartbeatThreadMaxCount, ctx)
	go cleanOldKvsClientMap(&multiChannelClient.OldKvsClientMaps, int64(multiChannelClient.SdkConfig.ConnectionTimeout), ctx)

	createTableResponse := &model.CreateTableResponse{}
	client.EXPECT().CreateTable(gomock.Any()).Return(createTableResponse, nil)
	createTableRequest := &model.CreateTableRequest{}
	response1, _ := multiChannelClient.CreateTable(createTableRequest)
	assert.Equal(t, createTableResponse, response1)

	describeTableResponse := &model.DescribeTableResponse{}
	client.EXPECT().DescribeTable(gomock.Any()).Return(describeTableResponse, nil)
	describeTableRequest := &model.DescribeTableRequest{}
	response2, _ := multiChannelClient.DescribeTable(describeTableRequest)
	assert.Equal(t, describeTableResponse, response2)

	listStoreResponse := &model.ListStoreResponse{}
	client.EXPECT().ListStore(gomock.Any()).Return(listStoreResponse, nil)
	listStoreRequest := &model.ListStoreRequest{}
	response3, _ := multiChannelClient.ListStore(listStoreRequest)
	assert.Equal(t, listStoreResponse, response3)

	listTableResponse := &model.ListTableResponse{}
	client.EXPECT().ListTable(gomock.Any()).Return(listTableResponse, nil)
	listTableRequest := &model.ListTableRequest{}
	response4, _ := multiChannelClient.ListTable(listTableRequest)
	assert.Equal(t, listTableResponse, response4)

	batchWriteKvResponse := &model.BatchWriteKvResponse{}
	client.EXPECT().BatchWriteKv(gomock.Any()).Return(batchWriteKvResponse, nil)
	batchWriteKvRequest := &model.BatchWriteKvRequest{}
	response5, _ := multiChannelClient.BatchWriteKv(batchWriteKvRequest)
	assert.Equal(t, batchWriteKvResponse, response5)

	deleteKvResponse := &model.DeleteKvResponse{}
	client.EXPECT().DeleteKv(gomock.Any()).Return(deleteKvResponse, nil)
	deleteKvRequest := &model.DeleteKvRequest{}
	response6, _ := multiChannelClient.DeleteKv(deleteKvRequest)
	assert.Equal(t, deleteKvResponse, response6)

	getKvResponse := &model.GetKvResponse{}
	client.EXPECT().GetKv(gomock.Any()).Return(getKvResponse, nil)
	getKvRequest := &model.GetKvRequest{}
	response7, _ := multiChannelClient.GetKv(getKvRequest)
	assert.Equal(t, getKvResponse, response7)

	putKvResponse := &model.PutKvResponse{}
	client.EXPECT().PutKv(gomock.Any()).Return(putKvResponse, nil)
	putKvRequest := &model.PutKvRequest{}
	response8, _ := multiChannelClient.PutKv(putKvRequest)
	assert.Equal(t, putKvResponse, response8)

	scanKvResponse := &model.ScanKvResponse{}
	client.EXPECT().ScanKv(gomock.Any()).Return(scanKvResponse, nil)
	scanKvRequest := &model.ScanKvRequest{}
	response9, _ := multiChannelClient.ScanKv(scanKvRequest)
	assert.Equal(t, scanKvResponse, response9)

	scanSkeyKvResponse := &model.ScanSkeyKvResponse{}
	client.EXPECT().ScanSkeyKv(gomock.Any()).Return(scanSkeyKvResponse, nil)
	scanSkeyKvRequest := &model.ScanSkeyKvRequest{}
	response10, _ := multiChannelClient.ScanSkeyKv(scanSkeyKvRequest)
	assert.Equal(t, scanSkeyKvResponse, response10)

	updateKvResponse := &model.UpdateKvResponse{}
	client.EXPECT().UpdateKv(gomock.Any()).Return(updateKvResponse, nil)
	updateKvRequest := &model.UpdateKvRequest{}
	response11, _ := multiChannelClient.UpdateKv(updateKvRequest)
	assert.Equal(t, updateKvResponse, response11)
	cancel()
	time.Sleep(5 * time.Second)
}
