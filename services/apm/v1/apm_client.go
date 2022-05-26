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

// ListAkSk 获取ak-sk
//
// 获取该用户创建的aksk列表
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

// ListAkSkInvoker 获取ak-sk
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

// ShowMasterAddress 查询master地址
//
// 根据region名称获取该名称下的master服务podlb地址信息
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
