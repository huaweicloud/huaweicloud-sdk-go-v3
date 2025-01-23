package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
)

type MockIKvsClient struct {
	ctrl     *gomock.Controller
	recorder *MockIKvsClientMockRecorder
}

type MockIKvsClientMockRecorder struct {
	mock *MockIKvsClient
}

func NewMockIKvsClient(ctrl *gomock.Controller) *MockIKvsClient {
	mock := &MockIKvsClient{ctrl: ctrl}
	mock.recorder = &MockIKvsClientMockRecorder{mock}
	return mock
}

func (m *MockIKvsClient) EXPECT() *MockIKvsClientMockRecorder {
	return m.recorder
}

// CreateTable 创建表
//
// 在指定仓内创建表，表名在仓内唯一；创建表时，指定主键模板及本地二级索引模板及全局二级索引模板。创建表时，如果仓不存在，将会自动创建仓。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) CreateTable(arg0 *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTable", arg0)
	ret0, _ := ret[0].(*model.CreateTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) CreateTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockIKvsClient)(nil).CreateTable), arg0)
}

// CreateTableInvoker 创建表
func (m *MockIKvsClient) CreateTableInvoker(arg0 *model.CreateTableRequest) *CreateTableInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTableInvoker", arg0)
	ret0, _ := ret[0].(*CreateTableInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) CreateTableInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTableInvoker", reflect.TypeOf((*MockIKvsClient)(nil).CreateTableInvoker), arg0)
}

// DescribeTable 查询表
//
// 指定仓查询表属性，如容量，规模，配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) DescribeTable(arg0 *model.DescribeTableRequest) (*model.DescribeTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTable", arg0)
	ret0, _ := ret[0].(*model.DescribeTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) DescribeTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockIKvsClient)(nil).DescribeTable), arg0)
}

// DescribeTableInvoker 查询表
func (m *MockIKvsClient) DescribeTableInvoker(arg0 *model.DescribeTableRequest) *DescribeTableInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTableInvoker", arg0)
	ret0, _ := ret[0].(*DescribeTableInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) DescribeTableInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTableInvoker", reflect.TypeOf((*MockIKvsClient)(nil).DescribeTableInvoker), arg0)
}

// ListStore 列举仓
//
// 一个账户下可以创建最多25个仓，每个仓可以创建最多100个store，响应中一次性返回所有仓名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) ListStore(arg0 *model.ListStoreRequest) (*model.ListStoreResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStore", arg0)
	ret0, _ := ret[0].(*model.ListStoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) ListStore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStore", reflect.TypeOf((*MockIKvsClient)(nil).ListStore), arg0)
}

// ListStoreInvoker 列举仓
func (m *MockIKvsClient) ListStoreInvoker(arg0 *model.ListStoreRequest) *ListStoreInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoreInvoker", arg0)
	ret0, _ := ret[0].(*ListStoreInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) ListStoreInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoreInvoker", reflect.TypeOf((*MockIKvsClient)(nil).ListStoreInvoker), arg0)
}

// ListTable 列举表
//
// 指定仓列举创建的所有表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) ListTable(arg0 *model.ListTableRequest) (*model.ListTableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTable", arg0)
	ret0, _ := ret[0].(*model.ListTableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) ListTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTable", reflect.TypeOf((*MockIKvsClient)(nil).ListTable), arg0)
}

// ListTableInvoker 列举表
func (m *MockIKvsClient) ListTableInvoker(arg0 *model.ListTableRequest) *ListTableInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTableInvoker", arg0)
	ret0, _ := ret[0].(*ListTableInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) ListTableInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableInvoker", reflect.TypeOf((*MockIKvsClient)(nil).ListTableInvoker), arg0)
}

// CheckHealth 网络信道健康检查
//
// 网络信道健康检查，返回response未抛出网络异常即为成功
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) CheckHealth(arg0 *model.CheckHealthRequest) (*model.CheckHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth", arg0)
	ret0, _ := ret[0].(*model.CheckHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) CheckHealth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockIKvsClient)(nil).CheckHealth), arg0)
}

// CheckHealthInvoker 网络信道健康检查
func (m *MockIKvsClient) CheckHealthInvoker(arg0 *model.CheckHealthRequest) *CheckHealthInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealthInvoker", arg0)
	ret0, _ := ret[0].(*CheckHealthInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) CheckHealthInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealthInvoker", reflect.TypeOf((*MockIKvsClient)(nil).CheckHealthInvoker), arg0)
}

// BatchWriteKv 批量写请求
//
// 批量写请求，其中可以携带一或多个表的不同kv的写操作，上传kv/删除kv。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) BatchWriteKv(arg0 *model.BatchWriteKvRequest) (*model.BatchWriteKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchWriteKv", arg0)
	ret0, _ := ret[0].(*model.BatchWriteKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) BatchWriteKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchWriteKv", reflect.TypeOf((*MockIKvsClient)(nil).BatchWriteKv), arg0)
}

// BatchWriteKvInvoker 批量写请求
func (m *MockIKvsClient) BatchWriteKvInvoker(arg0 *model.BatchWriteKvRequest) *BatchWriteKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchWriteKvInvoker", arg0)
	ret0, _ := ret[0].(*BatchWriteKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) BatchWriteKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchWriteKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).BatchWriteKvInvoker), arg0)
}

// DeleteKv 删除单个kv
//
// 指定表，指定主键，删除该文档；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) DeleteKv(arg0 *model.DeleteKvRequest) (*model.DeleteKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKv", arg0)
	ret0, _ := ret[0].(*model.DeleteKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) DeleteKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKv", reflect.TypeOf((*MockIKvsClient)(nil).DeleteKv), arg0)
}

// DeleteKvInvoker 删除单个kv
func (m *MockIKvsClient) DeleteKvInvoker(arg0 *model.DeleteKvRequest) *DeleteKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKvInvoker", arg0)
	ret0, _ := ret[0].(*DeleteKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) DeleteKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).DeleteKvInvoker), arg0)
}

// GetKv 查询单个kv
//
// 下载一个kv文档的全部内容，或者部分字段的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) GetKv(arg0 *model.GetKvRequest) (*model.GetKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKv", arg0)
	ret0, _ := ret[0].(*model.GetKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) GetKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKv", reflect.TypeOf((*MockIKvsClient)(nil).GetKv), arg0)
}

// GetKvInvoker 查询单个kv
func (m *MockIKvsClient) GetKvInvoker(arg0 *model.GetKvRequest) *GetKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKvInvoker", arg0)
	ret0, _ := ret[0].(*GetKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) GetKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).GetKvInvoker), arg0)
}

// PutKv 上传单个kv
//
// 指定表，新建kv或覆盖已有kv，且满足表的key schema描述；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) PutKv(arg0 *model.PutKvRequest) (*model.PutKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutKv", arg0)
	ret0, _ := ret[0].(*model.PutKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) PutKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutKv", reflect.TypeOf((*MockIKvsClient)(nil).PutKv), arg0)
}

// PutKvInvoker 上传单个kv
func (m *MockIKvsClient) PutKvInvoker(arg0 *model.PutKvRequest) *PutKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutKvInvoker", arg0)
	ret0, _ := ret[0].(*PutKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) PutKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).PutKvInvoker), arg0)
}

// ScanKv 扫描所有kv
//
// 指定表，扫描表下所有kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) ScanKv(arg0 *model.ScanKvRequest) (*model.ScanKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanKv", arg0)
	ret0, _ := ret[0].(*model.ScanKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) ScanKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanKv", reflect.TypeOf((*MockIKvsClient)(nil).ScanKv), arg0)
}

// ScanKvInvoker 扫描所有kv
func (m *MockIKvsClient) ScanKvInvoker(arg0 *model.ScanKvRequest) *ScanKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanKvInvoker", arg0)
	ret0, _ := ret[0].(*ScanKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) ScanKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).ScanKvInvoker), arg0)
}

// ScanSkeyKv 扫描分区键内kv
//
// 指定表及分区键，携带条件查询kv；允许指定过滤条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) ScanSkeyKv(arg0 *model.ScanSkeyKvRequest) (*model.ScanSkeyKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanSkeyKv", arg0)
	ret0, _ := ret[0].(*model.ScanSkeyKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) ScanSkeyKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanSkeyKv", reflect.TypeOf((*MockIKvsClient)(nil).ScanSkeyKv), arg0)
}

// ScanSkeyKvInvoker 扫描分区键内kv
func (m *MockIKvsClient) ScanSkeyKvInvoker(arg0 *model.ScanSkeyKvRequest) *ScanSkeyKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanSkeyKvInvoker", arg0)
	ret0, _ := ret[0].(*ScanSkeyKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) ScanSkeyKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanSkeyKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).ScanSkeyKvInvoker), arg0)
}

// UpdateKv 更新单个kv
//
// 指定表，指定主键，指定更新文档的部分内容，如果是自描述文档，指定字段名；如果是二进制文档，指定偏移位置和长度；允许指定条件执行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (m *MockIKvsClient) UpdateKv(arg0 *model.UpdateKvRequest) (*model.UpdateKvResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKv", arg0)
	ret0, _ := ret[0].(*model.UpdateKvResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIKvsClientMockRecorder) UpdateKv(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKv", reflect.TypeOf((*MockIKvsClient)(nil).UpdateKv), arg0)
}

// UpdateKvInvoker 更新单个kv
func (m *MockIKvsClient) UpdateKvInvoker(arg0 *model.UpdateKvRequest) *UpdateKvInvoker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKvInvoker", arg0)
	ret0, _ := ret[0].(*UpdateKvInvoker)
	return ret0
}

func (mr *MockIKvsClientMockRecorder) UpdateKvInvoker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKvInvoker", reflect.TypeOf((*MockIKvsClient)(nil).UpdateKvInvoker), arg0)
}
