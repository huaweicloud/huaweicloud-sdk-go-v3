package v1

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/common/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"

	"context"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type KvsMultiChannelClient struct {
	KvsClientMap     sync.Map
	OldKvsClientMaps sync.Map
	SdkConfig        *config.KvsSdkConfig
	tokenNum         int32
	weightCount      int32
	lock             sync.Mutex
	cancel           context.CancelFunc
}

func NewKvsMultiChannelClient(configFilePath string) (*KvsMultiChannelClient, error) {
	var ctx context.Context
	var err error
	multiChannelClient := KvsMultiChannelClient{KvsClientMap: sync.Map{}, OldKvsClientMaps: sync.Map{}}
	multiChannelClient.tokenNum = 1
	multiChannelClient.SdkConfig, err = NewKvsSdkConfigManager(configFilePath, &multiChannelClient.KvsClientMap, &multiChannelClient.OldKvsClientMaps)
	if err != nil {
		return nil, err
	}
	ctx, multiChannelClient.cancel = context.WithCancel(context.Background())
	go StartKvsClientHeartbeatKeeper(&multiChannelClient.KvsClientMap, multiChannelClient.SdkConfig.HeartbeatInterval, multiChannelClient.SdkConfig.HeartbeatThreadMaxCount, ctx)
	go cleanOldKvsClientMap(&multiChannelClient.OldKvsClientMaps, int64(multiChannelClient.SdkConfig.ConnectionTimeout+multiChannelClient.SdkConfig.ReadTimeout), ctx)
	return &multiChannelClient, nil
}

func cleanOldKvsClientMap(OldKvsClientMaps *sync.Map, ConnectionTimeout int64, ctx context.Context) {
	ticker := time.NewTicker(time.Duration(ConnectionTimeout) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			OldKvsClientMaps.Range(func(key, value interface{}) bool {
				keyValue, ok := key.(int64)
				if ok && time.Now().UnixMilli() > keyValue+ConnectionTimeout {
					OldKvsClientMaps.Delete(keyValue)
				}
				return true
			})
		}
	}
}

func (client *KvsMultiChannelClient) Close() {
	client.cancel()
}

// CreateTable 创建表
//
// 在指定仓内创建表，表名在仓内唯一；创建表时，指定主键模板及本地二级索引模板及全局二级索引模板。创建表时，如果仓不存在，将会自动创建仓。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.CreateTableResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.CreateTable(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// CreateTableInvoker 创建表
func (client *KvsMultiChannelClient) CreateTableInvoker(request *model.CreateTableRequest) (*CreateTableInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.CreateTableInvoker(request), nil
}

// DescribeTable 查询表
//
// 指定仓查询表属性，如容量，规模，配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) DescribeTable(request *model.DescribeTableRequest) (*model.DescribeTableResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.DescribeTableResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.DescribeTable(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// DescribeTableInvoker 查询表
func (client *KvsMultiChannelClient) DescribeTableInvoker(request *model.DescribeTableRequest) (*DescribeTableInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.DescribeTableInvoker(request), nil
}

// ListStore 列举仓
//
// 一个账户下可以创建最多25个仓，每个仓可以创建最多100个store，响应中一次性返回所有仓名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) ListStore(request *model.ListStoreRequest) (*model.ListStoreResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.ListStoreResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.ListStore(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// ListStoreInvoker 列举仓
func (client *KvsMultiChannelClient) ListStoreInvoker(request *model.ListStoreRequest) (*ListStoreInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.ListStoreInvoker(request), nil
}

// ListTable 列举表
//
// 指定仓列举创建的所有表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) ListTable(request *model.ListTableRequest) (*model.ListTableResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.ListTableResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.ListTable(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// ListTableInvoker 列举表
func (client *KvsMultiChannelClient) ListTableInvoker(request *model.ListTableRequest) (*ListTableInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.ListTableInvoker(request), nil
}

// CheckHealth 网络信道健康检查
//
// 网络信道健康检查，返回response未抛出网络异常即为成功
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) CheckHealth(request *model.CheckHealthRequest) (*model.CheckHealthResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.CheckHealthResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.CheckHealth(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// CheckHealthInvoker 网络信道健康检查
func (client *KvsMultiChannelClient) CheckHealthInvoker(request *model.CheckHealthRequest) (*CheckHealthInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.CheckHealthInvoker(request), nil
}

// BatchWriteKv 批量写请求
//
// 批量写请求，其中可以携带一或多个表的不同kv的写操作，上传kv/删除kv。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) BatchWriteKv(request *model.BatchWriteKvRequest) (*model.BatchWriteKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.BatchWriteKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.BatchWriteKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// BatchWriteKvInvoker 批量写请求
func (client *KvsMultiChannelClient) BatchWriteKvInvoker(request *model.BatchWriteKvRequest) (*BatchWriteKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.BatchWriteKvInvoker(request), nil
}

// DeleteKv 删除单个kv
//
// 指定表，指定主键，删除该文档；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) DeleteKv(request *model.DeleteKvRequest) (*model.DeleteKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.DeleteKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.DeleteKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// DeleteKvInvoker 删除单个kv
func (client *KvsMultiChannelClient) DeleteKvInvoker(request *model.DeleteKvRequest) (*DeleteKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.DeleteKvInvoker(request), nil
}

// GetKv 查询单个kv
//
// 下载一个kv文档的全部内容，或者部分字段的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) GetKv(request *model.GetKvRequest) (*model.GetKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.GetKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.GetKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// GetKvInvoker 查询单个kv
func (client *KvsMultiChannelClient) GetKvInvoker(request *model.GetKvRequest) (*GetKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.GetKvInvoker(request), nil
}

// PutKv 上传单个kv
//
// 指定表，新建kv或覆盖已有kv，且满足表的key schema描述；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) PutKv(request *model.PutKvRequest) (*model.PutKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.PutKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.PutKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// PutKvInvoker 上传单个kv
func (client *KvsMultiChannelClient) PutKvInvoker(request *model.PutKvRequest) (*PutKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.PutKvInvoker(request), nil
}

// ScanKv 扫描所有kv
//
// 指定表，扫描表下所有kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) ScanKv(request *model.ScanKvRequest) (*model.ScanKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.ScanKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.ScanKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// ScanKvInvoker 扫描所有kv
func (client *KvsMultiChannelClient) ScanKvInvoker(request *model.ScanKvRequest) (*ScanKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.ScanKvInvoker(request), nil
}

// ScanSkeyKv 扫描分区键内kv
//
// 指定表及分区键，携带条件查询kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) ScanSkeyKv(request *model.ScanSkeyKvRequest) (*model.ScanSkeyKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.ScanSkeyKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.ScanSkeyKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// ScanSkeyKvInvoker 扫描分区键内kv
func (client *KvsMultiChannelClient) ScanSkeyKvInvoker(request *model.ScanSkeyKvRequest) (*ScanSkeyKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.ScanSkeyKvInvoker(request), nil
}

// UpdateKv 更新单个kv
//
// 指定表，指定主键，指定更新文档的部分内容，如果是自描述文档，指定字段名；如果是二进制文档，指定偏移位置和长度；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (client *KvsMultiChannelClient) UpdateKv(request *model.UpdateKvRequest) (*model.UpdateKvResponse, error) {
	var retryCount int
	var managedClient *KvsManagedClient
	var response *model.UpdateKvResponse
	var err error

	for retryCount < client.SdkConfig.ApiRetryCount {
		managedClient, err = getKvsClientByPolling(client, retryCount)
		if err != nil {
			return nil, err
		}
		response, err = managedClient.KvsClient.UpdateKv(request)
		if err != nil {
			retryCount++
			continue
		}
		return response, err
	}
	return response, err
}

// UpdateKvInvoker 更新单个kv
func (client *KvsMultiChannelClient) UpdateKvInvoker(request *model.UpdateKvRequest) (*UpdateKvInvoker, error) {
	managedClient, err := getKvsClientByPolling(client, 0)
	if err != nil {
		return nil, err
	}
	return managedClient.KvsClient.UpdateKvInvoker(request), nil
}

func getKvsClientByPolling(client *KvsMultiChannelClient, retryCount int) (*KvsManagedClient, error) {
	// Exponential Backoff
	if retryCount > 0 {
		time.Sleep(time.Duration(rand.Intn(retryCount*retryCount*1000)) * time.Millisecond)
	}

	var selectCount int32
	clientMapSize := getMapSize(&client.KvsClientMap)
	value, ok := client.KvsClientMap.Load(client.tokenNum)
	if !ok {
		atomic.StoreInt32(&client.tokenNum, 1)
	}
	managedClient, ok := value.(*KvsManagedClient)
	if !ok {
		err := errors.New("trans value of KvsClientMap to *KvsManagedClient failed!")
		return nil, err
	}
	client.lock.Lock()
	defer client.lock.Unlock()
	for {
		if selectCount > clientMapSize {
			err := errors.New("all web channel of Kvs clients are unusable! please wait and retry")
			return nil, err
		}
		selectCount++

		if managedClient.IsUsable.Load() && atomic.LoadInt32(&client.weightCount) < managedClient.Endpoint.Weight {
			atomic.AddInt32(&client.weightCount, 1)
		} else if atomic.LoadInt32(&client.tokenNum) >= clientMapSize {
			atomic.StoreInt32(&client.tokenNum, 1)
			atomic.StoreInt32(&client.weightCount, 1)
		} else {
			atomic.AddInt32(&client.tokenNum, 1)
			atomic.StoreInt32(&client.weightCount, 1)
		}
		value, _ = client.KvsClientMap.Load(client.tokenNum)
		managedClient, ok = value.(*KvsManagedClient)
		if !ok {
			err := errors.New("trans value of KvsClientMap to *KvsManagedClient failed!")
			return nil, err
		}

		if managedClient.IsUsable.Load() {
			break
		}
	}
	return managedClient, nil
}

func getMapSize(m *sync.Map) int32 {
	var size int32
	m.Range(func(key, value interface{}) bool {
		size++
		return true
	})
	return size
}
