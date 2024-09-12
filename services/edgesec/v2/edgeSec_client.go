package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v2/model"
)

type EdgeSecClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewEdgeSecClient(hcClient *httpclient.HcHttpClient) *EdgeSecClient {
	return &EdgeSecClient{HcClient: hcClient}
}

func EdgeSecClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateDomains 创建防护域名
//
// 创建防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateDomains(request *model.CreateDomainsRequest) (*model.CreateDomainsResponse, error) {
	requestDef := GenReqDefForCreateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainsResponse), nil
	}
}

// CreateDomainsInvoker 创建防护域名
func (c *EdgeSecClient) CreateDomainsInvoker(request *model.CreateDomainsRequest) *CreateDomainsInvoker {
	requestDef := GenReqDefForCreateDomains()
	return &CreateDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomains 删除防护域名
//
// 删除防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

// DeleteDomainsInvoker 删除防护域名
func (c *EdgeSecClient) DeleteDomainsInvoker(request *model.DeleteDomainsRequest) *DeleteDomainsInvoker {
	requestDef := GenReqDefForDeleteDomains()
	return &DeleteDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainDetail 查询防护域名详情
//
// 查询防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDomainDetail(request *model.ShowDomainDetailRequest) (*model.ShowDomainDetailResponse, error) {
	requestDef := GenReqDefForShowDomainDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainDetailResponse), nil
	}
}

// ShowDomainDetailInvoker 查询防护域名详情
func (c *EdgeSecClient) ShowDomainDetailInvoker(request *model.ShowDomainDetailRequest) *ShowDomainDetailInvoker {
	requestDef := GenReqDefForShowDomainDetail()
	return &ShowDomainDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomains 查询防护域名列表
//
// 查询防护域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowDomains(request *model.ShowDomainsRequest) (*model.ShowDomainsResponse, error) {
	requestDef := GenReqDefForShowDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainsResponse), nil
	}
}

// ShowDomainsInvoker 查询防护域名列表
func (c *EdgeSecClient) ShowDomainsInvoker(request *model.ShowDomainsRequest) *ShowDomainsInvoker {
	requestDef := GenReqDefForShowDomains()
	return &ShowDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomains 更新防护域名
//
// 更新防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateDomains(request *model.UpdateDomainsRequest) (*model.UpdateDomainsResponse, error) {
	requestDef := GenReqDefForUpdateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainsResponse), nil
	}
}

// UpdateDomainsInvoker 更新防护域名
func (c *EdgeSecClient) UpdateDomainsInvoker(request *model.UpdateDomainsRequest) *UpdateDomainsInvoker {
	requestDef := GenReqDefForUpdateDomains()
	return &UpdateDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
