package v1

import "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"

type IKvsClient interface {
	CreateTable(*model.CreateTableRequest) (*model.CreateTableResponse, error)
	CreateTableInvoker(*model.CreateTableRequest) *CreateTableInvoker
	DescribeTable(*model.DescribeTableRequest) (*model.DescribeTableResponse, error)
	DescribeTableInvoker(*model.DescribeTableRequest) *DescribeTableInvoker
	ListStore(*model.ListStoreRequest) (*model.ListStoreResponse, error)
	ListStoreInvoker(*model.ListStoreRequest) *ListStoreInvoker
	ListTable(*model.ListTableRequest) (*model.ListTableResponse, error)
	ListTableInvoker(*model.ListTableRequest) *ListTableInvoker
	CheckHealth(*model.CheckHealthRequest) (*model.CheckHealthResponse, error)
	CheckHealthInvoker(*model.CheckHealthRequest) *CheckHealthInvoker
	BatchWriteKv(*model.BatchWriteKvRequest) (*model.BatchWriteKvResponse, error)
	BatchWriteKvInvoker(*model.BatchWriteKvRequest) *BatchWriteKvInvoker
	DeleteKv(*model.DeleteKvRequest) (*model.DeleteKvResponse, error)
	DeleteKvInvoker(*model.DeleteKvRequest) *DeleteKvInvoker
	GetKv(*model.GetKvRequest) (*model.GetKvResponse, error)
	GetKvInvoker(*model.GetKvRequest) *GetKvInvoker
	PutKv(*model.PutKvRequest) (*model.PutKvResponse, error)
	PutKvInvoker(*model.PutKvRequest) *PutKvInvoker
	ScanKv(*model.ScanKvRequest) (*model.ScanKvResponse, error)
	ScanKvInvoker(*model.ScanKvRequest) *ScanKvInvoker
	ScanSkeyKv(*model.ScanSkeyKvRequest) (*model.ScanSkeyKvResponse, error)
	ScanSkeyKvInvoker(*model.ScanSkeyKvRequest) *ScanSkeyKvInvoker
	UpdateKv(*model.UpdateKvRequest) (*model.UpdateKvResponse, error)
	UpdateKvInvoker(*model.UpdateKvRequest) *UpdateKvInvoker
}
