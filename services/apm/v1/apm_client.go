package v1

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/apm/v1/model"
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

//获取该用户创建的aksk列表
func (c *ApmClient) ListAkSk(request *model.ListAkSkRequest) (*model.ListAkSkResponse, error) {
	requestDef := GenReqDefForListAkSk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAkSkResponse), nil
	}
}

//该接口用于查询对应用户下的业务。
func (c *ApmClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}

//根据region名称获取该名称下的master服务podlb地址信息
func (c *ApmClient) ShowMasterAddress(request *model.ShowMasterAddressRequest) (*model.ShowMasterAddressResponse, error) {
	requestDef := GenReqDefForShowMasterAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMasterAddressResponse), nil
	}
}
