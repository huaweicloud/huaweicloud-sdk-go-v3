package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/apm/v1/model"
)

type ApmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewApmClient(hcClient *http_client.HcHttpClient) *ApmClient {
	return &ApmClient{HcClient: hcClient}
}

func ApmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListAkSk 获取ak/sk
//
// 获取该用户创建的ak/sk列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListAkSk(request *model.ListAkSkRequest) (*model.ListAkSkResponse, error) {
	requestDef := GenReqDefForListAkSk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAkSkResponse), nil
	}
}

// ListAkSkInvoker 获取ak/sk
func (c *ApmClient) ListAkSkInvoker(request *model.ListAkSkRequest) *ListAkSkInvoker {
	requestDef := GenReqDefForListAkSk()
	return &ListAkSkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusiness 查询业务列表
//
// 该接口用于查询对应用户下的业务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}

// ListBusinessInvoker 查询业务列表
func (c *ApmClient) ListBusinessInvoker(request *model.ListBusinessRequest) *ListBusinessInvoker {
	requestDef := GenReqDefForListBusiness()
	return &ListBusinessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnvMonitorItem
//
// 查询监控项列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ListEnvMonitorItem(request *model.ListEnvMonitorItemRequest) (*model.ListEnvMonitorItemResponse, error) {
	requestDef := GenReqDefForListEnvMonitorItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnvMonitorItemResponse), nil
	}
}

// ListEnvMonitorItemInvoker
func (c *ApmClient) ListEnvMonitorItemInvoker(request *model.ListEnvMonitorItemRequest) *ListEnvMonitorItemInvoker {
	requestDef := GenReqDefForListEnvMonitorItem()
	return &ListEnvMonitorItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveMonitorItemConfig 保存监控项
//
// 保存监控项
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) SaveMonitorItemConfig(request *model.SaveMonitorItemConfigRequest) (*model.SaveMonitorItemConfigResponse, error) {
	requestDef := GenReqDefForSaveMonitorItemConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveMonitorItemConfigResponse), nil
	}
}

// SaveMonitorItemConfigInvoker 保存监控项
func (c *ApmClient) SaveMonitorItemConfigInvoker(request *model.SaveMonitorItemConfigRequest) *SaveMonitorItemConfigInvoker {
	requestDef := GenReqDefForSaveMonitorItemConfig()
	return &SaveMonitorItemConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchApplication 对指定区域下的应用和环境及其探针情况进行搜索
//
// 对指定区域下的应用和环境及其探针情况进行搜索
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) SearchApplication(request *model.SearchApplicationRequest) (*model.SearchApplicationResponse, error) {
	requestDef := GenReqDefForSearchApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchApplicationResponse), nil
	}
}

// SearchApplicationInvoker 对指定区域下的应用和环境及其探针情况进行搜索
func (c *ApmClient) SearchApplicationInvoker(request *model.SearchApplicationRequest) *SearchApplicationInvoker {
	requestDef := GenReqDefForSearchApplication()
	return &SearchApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMasterAddress 查询master地址
//
// 根据region名称获取该名称下的master服务podlb地址信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *ApmClient) ShowMasterAddress(request *model.ShowMasterAddressRequest) (*model.ShowMasterAddressResponse, error) {
	requestDef := GenReqDefForShowMasterAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterAddressResponse), nil
	}
}

// ShowMasterAddressInvoker 查询master地址
func (c *ApmClient) ShowMasterAddressInvoker(request *model.ShowMasterAddressRequest) *ShowMasterAddressInvoker {
	requestDef := GenReqDefForShowMasterAddress()
	return &ShowMasterAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
