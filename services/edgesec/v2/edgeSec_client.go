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

// ApplyHttpPolicy 更新防护策略的域名
//
// 更新防护策略的域名
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ApplyHttpPolicy(request *model.ApplyHttpPolicyRequest) (*model.ApplyHttpPolicyResponse, error) {
	requestDef := GenReqDefForApplyHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyHttpPolicyResponse), nil
	}
}

// ApplyHttpPolicyInvoker 更新防护策略的域名
func (c *EdgeSecClient) ApplyHttpPolicyInvoker(request *model.ApplyHttpPolicyRequest) *ApplyHttpPolicyInvoker {
	requestDef := GenReqDefForApplyHttpPolicy()
	return &ApplyHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateHttpPolicy 创建防护策略
//
// 创建防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) CreateHttpPolicy(request *model.CreateHttpPolicyRequest) (*model.CreateHttpPolicyResponse, error) {
	requestDef := GenReqDefForCreateHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttpPolicyResponse), nil
	}
}

// CreateHttpPolicyInvoker 创建防护策略
func (c *EdgeSecClient) CreateHttpPolicyInvoker(request *model.CreateHttpPolicyRequest) *CreateHttpPolicyInvoker {
	requestDef := GenReqDefForCreateHttpPolicy()
	return &CreateHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteHttpPolicy 删除防护策略
//
// 删除防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) DeleteHttpPolicy(request *model.DeleteHttpPolicyRequest) (*model.DeleteHttpPolicyResponse, error) {
	requestDef := GenReqDefForDeleteHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttpPolicyResponse), nil
	}
}

// DeleteHttpPolicyInvoker 删除防护策略
func (c *EdgeSecClient) DeleteHttpPolicyInvoker(request *model.DeleteHttpPolicyRequest) *DeleteHttpPolicyInvoker {
	requestDef := GenReqDefForDeleteHttpPolicy()
	return &DeleteHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowHttpPolicies 查询防护策略列表
//
// 查询防护策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPolicies(request *model.ShowHttpPoliciesRequest) (*model.ShowHttpPoliciesResponse, error) {
	requestDef := GenReqDefForShowHttpPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPoliciesResponse), nil
	}
}

// ShowHttpPoliciesInvoker 查询防护策略列表
func (c *EdgeSecClient) ShowHttpPoliciesInvoker(request *model.ShowHttpPoliciesRequest) *ShowHttpPoliciesInvoker {
	requestDef := GenReqDefForShowHttpPolicies()
	return &ShowHttpPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttpPolicy 查询防护策略
//
// 查询防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) ShowHttpPolicy(request *model.ShowHttpPolicyRequest) (*model.ShowHttpPolicyResponse, error) {
	requestDef := GenReqDefForShowHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttpPolicyResponse), nil
	}
}

// ShowHttpPolicyInvoker 查询防护策略
func (c *EdgeSecClient) ShowHttpPolicyInvoker(request *model.ShowHttpPolicyRequest) *ShowHttpPolicyInvoker {
	requestDef := GenReqDefForShowHttpPolicy()
	return &ShowHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateHttpPolicy 更新防护策略
//
// 更新防护策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpPolicy(request *model.UpdateHttpPolicyRequest) (*model.UpdateHttpPolicyResponse, error) {
	requestDef := GenReqDefForUpdateHttpPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpPolicyResponse), nil
	}
}

// UpdateHttpPolicyInvoker 更新防护策略
func (c *EdgeSecClient) UpdateHttpPolicyInvoker(request *model.UpdateHttpPolicyRequest) *UpdateHttpPolicyInvoker {
	requestDef := GenReqDefForUpdateHttpPolicy()
	return &UpdateHttpPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHttpPolicyRuleStatus 更新防护策略规则开关
//
// 更新防护策略规则开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EdgeSecClient) UpdateHttpPolicyRuleStatus(request *model.UpdateHttpPolicyRuleStatusRequest) (*model.UpdateHttpPolicyRuleStatusResponse, error) {
	requestDef := GenReqDefForUpdateHttpPolicyRuleStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHttpPolicyRuleStatusResponse), nil
	}
}

// UpdateHttpPolicyRuleStatusInvoker 更新防护策略规则开关
func (c *EdgeSecClient) UpdateHttpPolicyRuleStatusInvoker(request *model.UpdateHttpPolicyRuleStatusRequest) *UpdateHttpPolicyRuleStatusInvoker {
	requestDef := GenReqDefForUpdateHttpPolicyRuleStatus()
	return &UpdateHttpPolicyRuleStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
