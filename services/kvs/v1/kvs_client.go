package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
)

type KvsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewKvsClient(hcClient *httpclient.HcHttpClient) *KvsClient {
	return &KvsClient{HcClient: hcClient}
}

func KvsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("basic.Credentials,v1.KvsCredentials")
	return builder
}

// CreateTable 创建表
//
// 在指定仓内创建表，表名在仓内唯一；创建表时，指定主键模板及本地二级索引模板及全局二级索引模板。创建表时，如果仓不存在，将会自动创建仓。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

// CreateTableInvoker 创建表
func (c *KvsClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeTable 查询表
//
// 指定仓查询表属性，如容量，规模，配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) DescribeTable(request *model.DescribeTableRequest) (*model.DescribeTableResponse, error) {
	requestDef := GenReqDefForDescribeTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeTableResponse), nil
	}
}

// DescribeTableInvoker 查询表
func (c *KvsClient) DescribeTableInvoker(request *model.DescribeTableRequest) *DescribeTableInvoker {
	requestDef := GenReqDefForDescribeTable()
	return &DescribeTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStore 列举仓
//
// 一个账户下可以创建最多25个仓，每个仓可以创建最多100个store，响应中一次性返回所有仓名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) ListStore(request *model.ListStoreRequest) (*model.ListStoreResponse, error) {
	requestDef := GenReqDefForListStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoreResponse), nil
	}
}

// ListStoreInvoker 列举仓
func (c *KvsClient) ListStoreInvoker(request *model.ListStoreRequest) *ListStoreInvoker {
	requestDef := GenReqDefForListStore()
	return &ListStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTable 列举表
//
// 指定仓列举创建的所有表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) ListTable(request *model.ListTableRequest) (*model.ListTableResponse, error) {
	requestDef := GenReqDefForListTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableResponse), nil
	}
}

// ListTableInvoker 列举表
func (c *KvsClient) ListTableInvoker(request *model.ListTableRequest) *ListTableInvoker {
	requestDef := GenReqDefForListTable()
	return &ListTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchWriteKv 批量写请求
//
// 批量写请求，其中可以携带一或多个表的不同kv的写操作，上传kv/删除kv。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) BatchWriteKv(request *model.BatchWriteKvRequest) (*model.BatchWriteKvResponse, error) {
	requestDef := GenReqDefForBatchWriteKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchWriteKvResponse), nil
	}
}

// BatchWriteKvInvoker 批量写请求
func (c *KvsClient) BatchWriteKvInvoker(request *model.BatchWriteKvRequest) *BatchWriteKvInvoker {
	requestDef := GenReqDefForBatchWriteKv()
	return &BatchWriteKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteKv 删除单个kv
//
// 指定表，指定主键，删除该文档；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) DeleteKv(request *model.DeleteKvRequest) (*model.DeleteKvResponse, error) {
	requestDef := GenReqDefForDeleteKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKvResponse), nil
	}
}

// DeleteKvInvoker 删除单个kv
func (c *KvsClient) DeleteKvInvoker(request *model.DeleteKvRequest) *DeleteKvInvoker {
	requestDef := GenReqDefForDeleteKv()
	return &DeleteKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetKv 查询单个kv
//
// 下载一个kv文档的全部内容，或者部分字段的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) GetKv(request *model.GetKvRequest) (*model.GetKvResponse, error) {
	requestDef := GenReqDefForGetKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetKvResponse), nil
	}
}

// GetKvInvoker 查询单个kv
func (c *KvsClient) GetKvInvoker(request *model.GetKvRequest) *GetKvInvoker {
	requestDef := GenReqDefForGetKv()
	return &GetKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutKv 上传单个kv
//
// 指定表，新建kv或覆盖已有kv，且满足表的key schema描述；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) PutKv(request *model.PutKvRequest) (*model.PutKvResponse, error) {
	requestDef := GenReqDefForPutKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutKvResponse), nil
	}
}

// PutKvInvoker 上传单个kv
func (c *KvsClient) PutKvInvoker(request *model.PutKvRequest) *PutKvInvoker {
	requestDef := GenReqDefForPutKv()
	return &PutKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ScanKv 扫描所有kv
//
// 指定表，扫描表下所有kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) ScanKv(request *model.ScanKvRequest) (*model.ScanKvResponse, error) {
	requestDef := GenReqDefForScanKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ScanKvResponse), nil
	}
}

// ScanKvInvoker 扫描所有kv
func (c *KvsClient) ScanKvInvoker(request *model.ScanKvRequest) *ScanKvInvoker {
	requestDef := GenReqDefForScanKv()
	return &ScanKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ScanSkeyKv 扫描分区键内kv
//
// 指定表及分区键，携带条件查询kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) ScanSkeyKv(request *model.ScanSkeyKvRequest) (*model.ScanSkeyKvResponse, error) {
	requestDef := GenReqDefForScanSkeyKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ScanSkeyKvResponse), nil
	}
}

// ScanSkeyKvInvoker 扫描分区键内kv
func (c *KvsClient) ScanSkeyKvInvoker(request *model.ScanSkeyKvRequest) *ScanSkeyKvInvoker {
	requestDef := GenReqDefForScanSkeyKv()
	return &ScanSkeyKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateKv 更新单个kv
//
// 指定表，指定主键，指定更新文档的部分内容，如果是自描述文档，指定字段名；如果是二进制文档，指定偏移位置和长度；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *KvsClient) UpdateKv(request *model.UpdateKvRequest) (*model.UpdateKvResponse, error) {
	requestDef := GenReqDefForUpdateKv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKvResponse), nil
	}
}

// UpdateKvInvoker 更新单个kv
func (c *KvsClient) UpdateKvInvoker(request *model.UpdateKvRequest) *UpdateKvInvoker {
	requestDef := GenReqDefForUpdateKv()
	return &UpdateKvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
