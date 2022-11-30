package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/model"
)

type NatClient struct {
	HcClient *http_client.HcHttpClient
}

func NewNatClient(hcClient *http_client.HcHttpClient) *NatClient {
	return &NatClient{HcClient: hcClient}
}

func NatClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateNatGatewayDnatRules 批量创建DNAT规则
//
// 批量创建DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) BatchCreateNatGatewayDnatRules(request *model.BatchCreateNatGatewayDnatRulesRequest) (*model.BatchCreateNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateNatGatewayDnatRulesResponse), nil
	}
}

// BatchCreateNatGatewayDnatRulesInvoker 批量创建DNAT规则
func (c *NatClient) BatchCreateNatGatewayDnatRulesInvoker(request *model.BatchCreateNatGatewayDnatRulesRequest) *BatchCreateNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForBatchCreateNatGatewayDnatRules()
	return &BatchCreateNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNatGatewayDnatRule 创建DNAT规则
//
// 创建DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateNatGatewayDnatRule(request *model.CreateNatGatewayDnatRuleRequest) (*model.CreateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayDnatRuleResponse), nil
	}
}

// CreateNatGatewayDnatRuleInvoker 创建DNAT规则
func (c *NatClient) CreateNatGatewayDnatRuleInvoker(request *model.CreateNatGatewayDnatRuleRequest) *CreateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewayDnatRule()
	return &CreateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNatGatewayDnatRule 删除DNAT规则
//
// 删除指定的DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteNatGatewayDnatRule(request *model.DeleteNatGatewayDnatRuleRequest) (*model.DeleteNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayDnatRuleResponse), nil
	}
}

// DeleteNatGatewayDnatRuleInvoker 删除DNAT规则
func (c *NatClient) DeleteNatGatewayDnatRuleInvoker(request *model.DeleteNatGatewayDnatRuleRequest) *DeleteNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewayDnatRule()
	return &DeleteNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGatewayDnatRules 查询DNAT规则列表
//
// 查询DNAT规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGatewayDnatRules(request *model.ListNatGatewayDnatRulesRequest) (*model.ListNatGatewayDnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewayDnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewayDnatRulesResponse), nil
	}
}

// ListNatGatewayDnatRulesInvoker 查询DNAT规则列表
func (c *NatClient) ListNatGatewayDnatRulesInvoker(request *model.ListNatGatewayDnatRulesRequest) *ListNatGatewayDnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewayDnatRules()
	return &ListNatGatewayDnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNatGatewayDnatRule 查询指定的DNAT规则详情
//
// 查询指定的DNAT规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowNatGatewayDnatRule(request *model.ShowNatGatewayDnatRuleRequest) (*model.ShowNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayDnatRuleResponse), nil
	}
}

// ShowNatGatewayDnatRuleInvoker 查询指定的DNAT规则详情
func (c *NatClient) ShowNatGatewayDnatRuleInvoker(request *model.ShowNatGatewayDnatRuleRequest) *ShowNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewayDnatRule()
	return &ShowNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNatGatewayDnatRule 更新DNAT规则
//
// 更新指定的DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdateNatGatewayDnatRule(request *model.UpdateNatGatewayDnatRuleRequest) (*model.UpdateNatGatewayDnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayDnatRuleResponse), nil
	}
}

// UpdateNatGatewayDnatRuleInvoker 更新DNAT规则
func (c *NatClient) UpdateNatGatewayDnatRuleInvoker(request *model.UpdateNatGatewayDnatRuleRequest) *UpdateNatGatewayDnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewayDnatRule()
	return &UpdateNatGatewayDnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNatGateway 创建公网NAT网关
//
// 创建公网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateNatGateway(request *model.CreateNatGatewayRequest) (*model.CreateNatGatewayResponse, error) {
	requestDef := GenReqDefForCreateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayResponse), nil
	}
}

// CreateNatGatewayInvoker 创建公网NAT网关
func (c *NatClient) CreateNatGatewayInvoker(request *model.CreateNatGatewayRequest) *CreateNatGatewayInvoker {
	requestDef := GenReqDefForCreateNatGateway()
	return &CreateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNatGateway 删除公网NAT网关
//
// 删除公网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteNatGateway(request *model.DeleteNatGatewayRequest) (*model.DeleteNatGatewayResponse, error) {
	requestDef := GenReqDefForDeleteNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayResponse), nil
	}
}

// DeleteNatGatewayInvoker 删除公网NAT网关
func (c *NatClient) DeleteNatGatewayInvoker(request *model.DeleteNatGatewayRequest) *DeleteNatGatewayInvoker {
	requestDef := GenReqDefForDeleteNatGateway()
	return &DeleteNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGateways 查询公网NAT网关列表
//
// 查询公网NAT网关实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGateways(request *model.ListNatGatewaysRequest) (*model.ListNatGatewaysResponse, error) {
	requestDef := GenReqDefForListNatGateways()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaysResponse), nil
	}
}

// ListNatGatewaysInvoker 查询公网NAT网关列表
func (c *NatClient) ListNatGatewaysInvoker(request *model.ListNatGatewaysRequest) *ListNatGatewaysInvoker {
	requestDef := GenReqDefForListNatGateways()
	return &ListNatGatewaysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNatGateway 查询指定的公网NAT网关详情
//
// 查询指定的公网NAT网关实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowNatGateway(request *model.ShowNatGatewayRequest) (*model.ShowNatGatewayResponse, error) {
	requestDef := GenReqDefForShowNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayResponse), nil
	}
}

// ShowNatGatewayInvoker 查询指定的公网NAT网关详情
func (c *NatClient) ShowNatGatewayInvoker(request *model.ShowNatGatewayRequest) *ShowNatGatewayInvoker {
	requestDef := GenReqDefForShowNatGateway()
	return &ShowNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNatGateway 更新公网NAT网关
//
// 更新公网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdateNatGateway(request *model.UpdateNatGatewayRequest) (*model.UpdateNatGatewayResponse, error) {
	requestDef := GenReqDefForUpdateNatGateway()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayResponse), nil
	}
}

// UpdateNatGatewayInvoker 更新公网NAT网关
func (c *NatClient) UpdateNatGatewayInvoker(request *model.UpdateNatGatewayRequest) *UpdateNatGatewayInvoker {
	requestDef := GenReqDefForUpdateNatGateway()
	return &UpdateNatGatewayInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNatGatewaySnatRule 创建SNAT规则
//
// 创建SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateNatGatewaySnatRule(request *model.CreateNatGatewaySnatRuleRequest) (*model.CreateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewaySnatRuleResponse), nil
	}
}

// CreateNatGatewaySnatRuleInvoker 创建SNAT规则
func (c *NatClient) CreateNatGatewaySnatRuleInvoker(request *model.CreateNatGatewaySnatRuleRequest) *CreateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForCreateNatGatewaySnatRule()
	return &CreateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNatGatewaySnatRule 删除SNAT规则
//
// 删除指定的SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteNatGatewaySnatRule(request *model.DeleteNatGatewaySnatRuleRequest) (*model.DeleteNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewaySnatRuleResponse), nil
	}
}

// DeleteNatGatewaySnatRuleInvoker 删除SNAT规则
func (c *NatClient) DeleteNatGatewaySnatRuleInvoker(request *model.DeleteNatGatewaySnatRuleRequest) *DeleteNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForDeleteNatGatewaySnatRule()
	return &DeleteNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGatewaySnatRules 查询SNAT规则列表
//
// 查询SNAT规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGatewaySnatRules(request *model.ListNatGatewaySnatRulesRequest) (*model.ListNatGatewaySnatRulesResponse, error) {
	requestDef := GenReqDefForListNatGatewaySnatRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaySnatRulesResponse), nil
	}
}

// ListNatGatewaySnatRulesInvoker 查询SNAT规则列表
func (c *NatClient) ListNatGatewaySnatRulesInvoker(request *model.ListNatGatewaySnatRulesRequest) *ListNatGatewaySnatRulesInvoker {
	requestDef := GenReqDefForListNatGatewaySnatRules()
	return &ListNatGatewaySnatRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNatGatewaySnatRule 查询指定的SNAT规则详情
//
// 查询指定的SNAT规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowNatGatewaySnatRule(request *model.ShowNatGatewaySnatRuleRequest) (*model.ShowNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForShowNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewaySnatRuleResponse), nil
	}
}

// ShowNatGatewaySnatRuleInvoker 查询指定的SNAT规则详情
func (c *NatClient) ShowNatGatewaySnatRuleInvoker(request *model.ShowNatGatewaySnatRuleRequest) *ShowNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForShowNatGatewaySnatRule()
	return &ShowNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNatGatewaySnatRule 更新SNAT规则
//
// 更新指定的SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdateNatGatewaySnatRule(request *model.UpdateNatGatewaySnatRuleRequest) (*model.UpdateNatGatewaySnatRuleResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewaySnatRuleResponse), nil
	}
}

// UpdateNatGatewaySnatRuleInvoker 更新SNAT规则
func (c *NatClient) UpdateNatGatewaySnatRuleInvoker(request *model.UpdateNatGatewaySnatRuleRequest) *UpdateNatGatewaySnatRuleInvoker {
	requestDef := GenReqDefForUpdateNatGatewaySnatRule()
	return &UpdateNatGatewaySnatRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
