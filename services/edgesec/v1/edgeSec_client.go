package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/edgesec/v1/model"
)

type EdgeSecClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEdgeSecClient(hcClient *http_client.HcHttpClient) *EdgeSecClient {
	return &EdgeSecClient{HcClient: hcClient}
}

func EdgeSecClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListEdgeSecSubscription 查询边缘安全已购产品
//
// 租户查询边缘安全已购产品
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListEdgeSecSubscription(request *model.ListEdgeSecSubscriptionRequest) (*model.ListEdgeSecSubscriptionResponse, error) {
	requestDef := GenReqDefForListEdgeSecSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeSecSubscriptionResponse), nil
	}
}

// ListEdgeSecSubscriptionInvoker 查询边缘安全已购产品
func (c *EdgeSecClient) ListEdgeSecSubscriptionInvoker(request *model.ListEdgeSecSubscriptionRequest) *ListEdgeSecSubscriptionInvoker {
	requestDef := GenReqDefForListEdgeSecSubscription()
	return &ListEdgeSecSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeDDoSDomains 添加ddos防护域名
//
// 租户添加ddos防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateEdgeDDoSDomains(request *model.CreateEdgeDDoSDomainsRequest) (*model.CreateEdgeDDoSDomainsResponse, error) {
	requestDef := GenReqDefForCreateEdgeDDoSDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeDDoSDomainsResponse), nil
	}
}

// CreateEdgeDDoSDomainsInvoker 添加ddos防护域名
func (c *EdgeSecClient) CreateEdgeDDoSDomainsInvoker(request *model.CreateEdgeDDoSDomainsRequest) *CreateEdgeDDoSDomainsInvoker {
	requestDef := GenReqDefForCreateEdgeDDoSDomains()
	return &CreateEdgeDDoSDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeDDoSDomains 删除ddos防护域名
//
// 租户删除ddos防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteEdgeDDoSDomains(request *model.DeleteEdgeDDoSDomainsRequest) (*model.DeleteEdgeDDoSDomainsResponse, error) {
	requestDef := GenReqDefForDeleteEdgeDDoSDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeDDoSDomainsResponse), nil
	}
}

// DeleteEdgeDDoSDomainsInvoker 删除ddos防护域名
func (c *EdgeSecClient) DeleteEdgeDDoSDomainsInvoker(request *model.DeleteEdgeDDoSDomainsRequest) *DeleteEdgeDDoSDomainsInvoker {
	requestDef := GenReqDefForDeleteEdgeDDoSDomains()
	return &DeleteEdgeDDoSDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeDDoSDomains 查询ddos防护域名
//
// 查询租户ddos防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListEdgeDDoSDomains(request *model.ListEdgeDDoSDomainsRequest) (*model.ListEdgeDDoSDomainsResponse, error) {
	requestDef := GenReqDefForListEdgeDDoSDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeDDoSDomainsResponse), nil
	}
}

// ListEdgeDDoSDomainsInvoker 查询ddos防护域名
func (c *EdgeSecClient) ListEdgeDDoSDomainsInvoker(request *model.ListEdgeDDoSDomainsRequest) *ListEdgeDDoSDomainsInvoker {
	requestDef := GenReqDefForListEdgeDDoSDomains()
	return &ListEdgeDDoSDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStatisticsEvent 查询租户受攻击事件数据
//
// 查询租户受攻击事件数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowStatisticsEvent(request *model.ShowStatisticsEventRequest) (*model.ShowStatisticsEventResponse, error) {
	requestDef := GenReqDefForShowStatisticsEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticsEventResponse), nil
	}
}

// ShowStatisticsEventInvoker 查询租户受攻击事件数据
func (c *EdgeSecClient) ShowStatisticsEventInvoker(request *model.ShowStatisticsEventRequest) *ShowStatisticsEventInvoker {
	requestDef := GenReqDefForShowStatisticsEvent()
	return &ShowStatisticsEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStatisticsTraffic 查询租户流量数据
//
// 查询租户流量数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowStatisticsTraffic(request *model.ShowStatisticsTrafficRequest) (*model.ShowStatisticsTrafficResponse, error) {
	requestDef := GenReqDefForShowStatisticsTraffic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStatisticsTrafficResponse), nil
	}
}

// ShowStatisticsTrafficInvoker 查询租户流量数据
func (c *EdgeSecClient) ShowStatisticsTrafficInvoker(request *model.ShowStatisticsTrafficRequest) *ShowStatisticsTrafficInvoker {
	requestDef := GenReqDefForShowStatisticsTraffic()
	return &ShowStatisticsTrafficInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeDDoSDomains 更新ddos防护域名
//
// 租户更新ddos防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateEdgeDDoSDomains(request *model.UpdateEdgeDDoSDomainsRequest) (*model.UpdateEdgeDDoSDomainsResponse, error) {
	requestDef := GenReqDefForUpdateEdgeDDoSDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeDDoSDomainsResponse), nil
	}
}

// UpdateEdgeDDoSDomainsInvoker 更新ddos防护域名
func (c *EdgeSecClient) UpdateEdgeDDoSDomainsInvoker(request *model.UpdateEdgeDDoSDomainsRequest) *UpdateEdgeDDoSDomainsInvoker {
	requestDef := GenReqDefForUpdateEdgeDDoSDomains()
	return &UpdateEdgeDDoSDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyWafPolicy 更新防护策略的域名
//
// 更新防护策略的域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ApplyWafPolicy(request *model.ApplyWafPolicyRequest) (*model.ApplyWafPolicyResponse, error) {
	requestDef := GenReqDefForApplyWafPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyWafPolicyResponse), nil
	}
}

// ApplyWafPolicyInvoker 更新防护策略的域名
func (c *EdgeSecClient) ApplyWafPolicyInvoker(request *model.ApplyWafPolicyRequest) *ApplyWafPolicyInvoker {
	requestDef := GenReqDefForApplyWafPolicy()
	return &ApplyWafPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCertificate 创建证书
//
// 创建证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateCertificate(request *model.CreateCertificateRequest) (*model.CreateCertificateResponse, error) {
	requestDef := GenReqDefForCreateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCertificateResponse), nil
	}
}

// CreateCertificateInvoker 创建证书
func (c *EdgeSecClient) CreateCertificateInvoker(request *model.CreateCertificateRequest) *CreateCertificateInvoker {
	requestDef := GenReqDefForCreateCertificate()
	return &CreateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEdgeWafDomains 创建防护域名
//
// 创建防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateEdgeWafDomains(request *model.CreateEdgeWafDomainsRequest) (*model.CreateEdgeWafDomainsResponse, error) {
	requestDef := GenReqDefForCreateEdgeWafDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEdgeWafDomainsResponse), nil
	}
}

// CreateEdgeWafDomainsInvoker 创建防护域名
func (c *EdgeSecClient) CreateEdgeWafDomainsInvoker(request *model.CreateEdgeWafDomainsRequest) *CreateEdgeWafDomainsInvoker {
	requestDef := GenReqDefForCreateEdgeWafDomains()
	return &CreateEdgeWafDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicy 创建防护策略
//
// 创建防护策略，系统会在生成策略时配置一些默认的配置项，如果需要修改策略的默认配置项需要通过调用更新防护策略接口实现
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreatePolicy(request *model.CreatePolicyRequest) (*model.CreatePolicyResponse, error) {
	requestDef := GenReqDefForCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyResponse), nil
	}
}

// CreatePolicyInvoker 创建防护策略
func (c *EdgeSecClient) CreatePolicyInvoker(request *model.CreatePolicyRequest) *CreatePolicyInvoker {
	requestDef := GenReqDefForCreatePolicy()
	return &CreatePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCertificate 删除证书
//
// 删除证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteCertificate(request *model.DeleteCertificateRequest) (*model.DeleteCertificateResponse, error) {
	requestDef := GenReqDefForDeleteCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCertificateResponse), nil
	}
}

// DeleteCertificateInvoker 删除证书
func (c *EdgeSecClient) DeleteCertificateInvoker(request *model.DeleteCertificateRequest) *DeleteCertificateInvoker {
	requestDef := GenReqDefForDeleteCertificate()
	return &DeleteCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEdgeWafDomains 删除防护域名
//
// 删除防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteEdgeWafDomains(request *model.DeleteEdgeWafDomainsRequest) (*model.DeleteEdgeWafDomainsResponse, error) {
	requestDef := GenReqDefForDeleteEdgeWafDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEdgeWafDomainsResponse), nil
	}
}

// DeleteEdgeWafDomainsInvoker 删除防护域名
func (c *EdgeSecClient) DeleteEdgeWafDomainsInvoker(request *model.DeleteEdgeWafDomainsRequest) *DeleteEdgeWafDomainsInvoker {
	requestDef := GenReqDefForDeleteEdgeWafDomains()
	return &DeleteEdgeWafDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicy 删除防护策略
//
// 删除防护策略，若策略正在使用，则需要先接解除域名与策略的绑定关系才能删除策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeletePolicy(request *model.DeletePolicyRequest) (*model.DeletePolicyResponse, error) {
	requestDef := GenReqDefForDeletePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyResponse), nil
	}
}

// DeletePolicyInvoker 删除防护策略
func (c *EdgeSecClient) DeletePolicyInvoker(request *model.DeletePolicyRequest) *DeletePolicyInvoker {
	requestDef := GenReqDefForDeletePolicy()
	return &DeletePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCdnDomains 查询CDN域名列表
//
// 查询CDN域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListCdnDomains(request *model.ListCdnDomainsRequest) (*model.ListCdnDomainsResponse, error) {
	requestDef := GenReqDefForListCdnDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCdnDomainsResponse), nil
	}
}

// ListCdnDomainsInvoker 查询CDN域名列表
func (c *EdgeSecClient) ListCdnDomainsInvoker(request *model.ListCdnDomainsRequest) *ListCdnDomainsInvoker {
	requestDef := GenReqDefForListCdnDomains()
	return &ListCdnDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCertificates 查询证书列表
//
// 查询证书列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListCertificates(request *model.ListCertificatesRequest) (*model.ListCertificatesResponse, error) {
	requestDef := GenReqDefForListCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCertificatesResponse), nil
	}
}

// ListCertificatesInvoker 查询证书列表
func (c *EdgeSecClient) ListCertificatesInvoker(request *model.ListCertificatesRequest) *ListCertificatesInvoker {
	requestDef := GenReqDefForListCertificates()
	return &ListCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEdgeWafDomains 查询WAF防护域名列表
//
// 查询WAF防护域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListEdgeWafDomains(request *model.ListEdgeWafDomainsRequest) (*model.ListEdgeWafDomainsResponse, error) {
	requestDef := GenReqDefForListEdgeWafDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEdgeWafDomainsResponse), nil
	}
}

// ListEdgeWafDomainsInvoker 查询WAF防护域名列表
func (c *EdgeSecClient) ListEdgeWafDomainsInvoker(request *model.ListEdgeWafDomainsRequest) *ListEdgeWafDomainsInvoker {
	requestDef := GenReqDefForListEdgeWafDomains()
	return &ListEdgeWafDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicy 查询防护策略列表
//
// 查询防护策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ListPolicy(request *model.ListPolicyRequest) (*model.ListPolicyResponse, error) {
	requestDef := GenReqDefForListPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyResponse), nil
	}
}

// ListPolicyInvoker 查询防护策略列表
func (c *EdgeSecClient) ListPolicyInvoker(request *model.ListPolicyRequest) *ListPolicyInvoker {
	requestDef := GenReqDefForListPolicy()
	return &ListPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCertificate 查询证书
//
// 查询证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowCertificate(request *model.ShowCertificateRequest) (*model.ShowCertificateResponse, error) {
	requestDef := GenReqDefForShowCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCertificateResponse), nil
	}
}

// ShowCertificateInvoker 查询证书
func (c *EdgeSecClient) ShowCertificateInvoker(request *model.ShowCertificateRequest) *ShowCertificateInvoker {
	requestDef := GenReqDefForShowCertificate()
	return &ShowCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEdgeWafDomains 查询防护域名
//
// 查询防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowEdgeWafDomains(request *model.ShowEdgeWafDomainsRequest) (*model.ShowEdgeWafDomainsResponse, error) {
	requestDef := GenReqDefForShowEdgeWafDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEdgeWafDomainsResponse), nil
	}
}

// ShowEdgeWafDomainsInvoker 查询防护域名
func (c *EdgeSecClient) ShowEdgeWafDomainsInvoker(request *model.ShowEdgeWafDomainsRequest) *ShowEdgeWafDomainsInvoker {
	requestDef := GenReqDefForShowEdgeWafDomains()
	return &ShowEdgeWafDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCertificate 修改证书
//
// 修改证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateCertificate(request *model.UpdateCertificateRequest) (*model.UpdateCertificateResponse, error) {
	requestDef := GenReqDefForUpdateCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCertificateResponse), nil
	}
}

// UpdateCertificateInvoker 修改证书
func (c *EdgeSecClient) UpdateCertificateInvoker(request *model.UpdateCertificateRequest) *UpdateCertificateInvoker {
	requestDef := GenReqDefForUpdateCertificate()
	return &UpdateCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEdgeWafDomains 更新防护域名
//
// 更新防护域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateEdgeWafDomains(request *model.UpdateEdgeWafDomainsRequest) (*model.UpdateEdgeWafDomainsResponse, error) {
	requestDef := GenReqDefForUpdateEdgeWafDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEdgeWafDomainsResponse), nil
	}
}

// UpdateEdgeWafDomainsInvoker 更新防护域名
func (c *EdgeSecClient) UpdateEdgeWafDomainsInvoker(request *model.UpdateEdgeWafDomainsRequest) *UpdateEdgeWafDomainsInvoker {
	requestDef := GenReqDefForUpdateEdgeWafDomains()
	return &UpdateEdgeWafDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
