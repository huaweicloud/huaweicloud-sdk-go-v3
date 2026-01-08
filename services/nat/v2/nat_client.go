package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nat/v2/model"
)

type NatClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewNatClient(hcClient *httpclient.HcHttpClient) *NatClient {
	return &NatClient{HcClient: hcClient}
}

func NatClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateDeleteTransitSubnetTags 批量添加删除中转子网标签
//
// - 为指定中转子网实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理中转子网实例的标签。
// - 一个中转子网上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) BatchCreateDeleteTransitSubnetTags(request *model.BatchCreateDeleteTransitSubnetTagsRequest) (*model.BatchCreateDeleteTransitSubnetTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteTransitSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteTransitSubnetTagsResponse), nil
	}
}

// BatchCreateDeleteTransitSubnetTagsInvoker 批量添加删除中转子网标签
func (c *NatClient) BatchCreateDeleteTransitSubnetTagsInvoker(request *model.BatchCreateDeleteTransitSubnetTagsRequest) *BatchCreateDeleteTransitSubnetTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeleteTransitSubnetTags()
	return &BatchCreateDeleteTransitSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTransitSubnet 创建中转子网
//
// 创建中转子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateTransitSubnet(request *model.CreateTransitSubnetRequest) (*model.CreateTransitSubnetResponse, error) {
	requestDef := GenReqDefForCreateTransitSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitSubnetResponse), nil
	}
}

// CreateTransitSubnetInvoker 创建中转子网
func (c *NatClient) CreateTransitSubnetInvoker(request *model.CreateTransitSubnetRequest) *CreateTransitSubnetInvoker {
	requestDef := GenReqDefForCreateTransitSubnet()
	return &CreateTransitSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTransitSubnetTag 添加中转子网标签
//
// - 为指定中转子网添加标签
// - 一个中转子网上最多有20个标签。
// - 此接口为幂等接口：
// - 创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateTransitSubnetTag(request *model.CreateTransitSubnetTagRequest) (*model.CreateTransitSubnetTagResponse, error) {
	requestDef := GenReqDefForCreateTransitSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitSubnetTagResponse), nil
	}
}

// CreateTransitSubnetTagInvoker 添加中转子网标签
func (c *NatClient) CreateTransitSubnetTagInvoker(request *model.CreateTransitSubnetTagRequest) *CreateTransitSubnetTagInvoker {
	requestDef := GenReqDefForCreateTransitSubnetTag()
	return &CreateTransitSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransitSubnet 删除中转子网
//
// 删除中转子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteTransitSubnet(request *model.DeleteTransitSubnetRequest) (*model.DeleteTransitSubnetResponse, error) {
	requestDef := GenReqDefForDeleteTransitSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitSubnetResponse), nil
	}
}

// DeleteTransitSubnetInvoker 删除中转子网
func (c *NatClient) DeleteTransitSubnetInvoker(request *model.DeleteTransitSubnetRequest) *DeleteTransitSubnetInvoker {
	requestDef := GenReqDefForDeleteTransitSubnet()
	return &DeleteTransitSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransitSubnetTag 删除中转子网标签
//
// - 幂等接口：
// - 删除时，不对标签字符集做校验，调用接口前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，key不能为空或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteTransitSubnetTag(request *model.DeleteTransitSubnetTagRequest) (*model.DeleteTransitSubnetTagResponse, error) {
	requestDef := GenReqDefForDeleteTransitSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitSubnetTagResponse), nil
	}
}

// DeleteTransitSubnetTagInvoker 删除中转子网标签
func (c *NatClient) DeleteTransitSubnetTagInvoker(request *model.DeleteTransitSubnetTagRequest) *DeleteTransitSubnetTagInvoker {
	requestDef := GenReqDefForDeleteTransitSubnetTag()
	return &DeleteTransitSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitSubnet 查询中转子网列表
//
// 查询中转子网列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitSubnet(request *model.ListTransitSubnetRequest) (*model.ListTransitSubnetResponse, error) {
	requestDef := GenReqDefForListTransitSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitSubnetResponse), nil
	}
}

// ListTransitSubnetInvoker 查询中转子网列表
func (c *NatClient) ListTransitSubnetInvoker(request *model.ListTransitSubnetRequest) *ListTransitSubnetInvoker {
	requestDef := GenReqDefForListTransitSubnet()
	return &ListTransitSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitSubnetTags 查询中转子网项目标签
//
// - 查询租户在指定Project的所有中转子网标签集合。
// - 标签管理服务需要能够列出当前租户全部已使用的中转子网标签集合，为打中转子网标签和过滤中转子网实例时提供标签联想功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitSubnetTags(request *model.ListTransitSubnetTagsRequest) (*model.ListTransitSubnetTagsResponse, error) {
	requestDef := GenReqDefForListTransitSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitSubnetTagsResponse), nil
	}
}

// ListTransitSubnetTagsInvoker 查询中转子网项目标签
func (c *NatClient) ListTransitSubnetTagsInvoker(request *model.ListTransitSubnetTagsRequest) *ListTransitSubnetTagsInvoker {
	requestDef := GenReqDefForListTransitSubnetTags()
	return &ListTransitSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitSubnetsByTags 查询中转子网实例
//
// - 使用标签过滤中转子网实例。
// - 标签管理服务需要提供按标签过滤中转子网服务实例并汇总显示在列表中，需要中转子网服务提供查询能力。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitSubnetsByTags(request *model.ListTransitSubnetsByTagsRequest) (*model.ListTransitSubnetsByTagsResponse, error) {
	requestDef := GenReqDefForListTransitSubnetsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitSubnetsByTagsResponse), nil
	}
}

// ListTransitSubnetsByTagsInvoker 查询中转子网实例
func (c *NatClient) ListTransitSubnetsByTagsInvoker(request *model.ListTransitSubnetsByTagsRequest) *ListTransitSubnetsByTagsInvoker {
	requestDef := GenReqDefForListTransitSubnetsByTags()
	return &ListTransitSubnetsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransitSubnet 查询指定的中转子网详情
//
// 查询指定的中转子网详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowTransitSubnet(request *model.ShowTransitSubnetRequest) (*model.ShowTransitSubnetResponse, error) {
	requestDef := GenReqDefForShowTransitSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitSubnetResponse), nil
	}
}

// ShowTransitSubnetInvoker 查询指定的中转子网详情
func (c *NatClient) ShowTransitSubnetInvoker(request *model.ShowTransitSubnetRequest) *ShowTransitSubnetInvoker {
	requestDef := GenReqDefForShowTransitSubnet()
	return &ShowTransitSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransitSubnetTags 查询中转子网标签
//
// - 查询指定中转子网实例的标签信息。
// - 标签管理服务需要使用该接口查询指定中转子网实例的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowTransitSubnetTags(request *model.ShowTransitSubnetTagsRequest) (*model.ShowTransitSubnetTagsResponse, error) {
	requestDef := GenReqDefForShowTransitSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitSubnetTagsResponse), nil
	}
}

// ShowTransitSubnetTagsInvoker 查询中转子网标签
func (c *NatClient) ShowTransitSubnetTagsInvoker(request *model.ShowTransitSubnetTagsRequest) *ShowTransitSubnetTagsInvoker {
	requestDef := GenReqDefForShowTransitSubnetTags()
	return &ShowTransitSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTransitSubnet 更新中转子网
//
// 更新指定的中转子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdateTransitSubnet(request *model.UpdateTransitSubnetRequest) (*model.UpdateTransitSubnetResponse, error) {
	requestDef := GenReqDefForUpdateTransitSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTransitSubnetResponse), nil
	}
}

// UpdateTransitSubnetInvoker 更新中转子网
func (c *NatClient) UpdateTransitSubnetInvoker(request *model.UpdateTransitSubnetRequest) *UpdateTransitSubnetInvoker {
	requestDef := GenReqDefForUpdateTransitSubnet()
	return &UpdateTransitSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreatePrivateDnat 创建DNAT规则
//
// 创建DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreatePrivateDnat(request *model.CreatePrivateDnatRequest) (*model.CreatePrivateDnatResponse, error) {
	requestDef := GenReqDefForCreatePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateDnatResponse), nil
	}
}

// CreatePrivateDnatInvoker 创建DNAT规则
func (c *NatClient) CreatePrivateDnatInvoker(request *model.CreatePrivateDnatRequest) *CreatePrivateDnatInvoker {
	requestDef := GenReqDefForCreatePrivateDnat()
	return &CreatePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeletePrivateDnat 删除DNAT规则
//
// 删除指定的DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeletePrivateDnat(request *model.DeletePrivateDnatRequest) (*model.DeletePrivateDnatResponse, error) {
	requestDef := GenReqDefForDeletePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateDnatResponse), nil
	}
}

// DeletePrivateDnatInvoker 删除DNAT规则
func (c *NatClient) DeletePrivateDnatInvoker(request *model.DeletePrivateDnatRequest) *DeletePrivateDnatInvoker {
	requestDef := GenReqDefForDeletePrivateDnat()
	return &DeletePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPrivateDnats 查询DNAT规则列表
//
// 查询DNAT规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListPrivateDnats(request *model.ListPrivateDnatsRequest) (*model.ListPrivateDnatsResponse, error) {
	requestDef := GenReqDefForListPrivateDnats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateDnatsResponse), nil
	}
}

// ListPrivateDnatsInvoker 查询DNAT规则列表
func (c *NatClient) ListPrivateDnatsInvoker(request *model.ListPrivateDnatsRequest) *ListPrivateDnatsInvoker {
	requestDef := GenReqDefForListPrivateDnats()
	return &ListPrivateDnatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowPrivateDnat 查询指定的DNAT规则详情
//
// 查询指定的DNAT规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowPrivateDnat(request *model.ShowPrivateDnatRequest) (*model.ShowPrivateDnatResponse, error) {
	requestDef := GenReqDefForShowPrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateDnatResponse), nil
	}
}

// ShowPrivateDnatInvoker 查询指定的DNAT规则详情
func (c *NatClient) ShowPrivateDnatInvoker(request *model.ShowPrivateDnatRequest) *ShowPrivateDnatInvoker {
	requestDef := GenReqDefForShowPrivateDnat()
	return &ShowPrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdatePrivateDnat 更新DNAT规则
//
// 更新指定的DNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdatePrivateDnat(request *model.UpdatePrivateDnatRequest) (*model.UpdatePrivateDnatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateDnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateDnatResponse), nil
	}
}

// UpdatePrivateDnatInvoker 更新DNAT规则
func (c *NatClient) UpdatePrivateDnatInvoker(request *model.UpdatePrivateDnatRequest) *UpdatePrivateDnatInvoker {
	requestDef := GenReqDefForUpdatePrivateDnat()
	return &UpdatePrivateDnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDeleteTransitIpTags 批量添加删除中转IP标签
//
// - 为指定中转IP实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理中转IP实例的标签。
// - 一个中转IP上最多有10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) BatchCreateDeleteTransitIpTags(request *model.BatchCreateDeleteTransitIpTagsRequest) (*model.BatchCreateDeleteTransitIpTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteTransitIpTagsResponse), nil
	}
}

// BatchCreateDeleteTransitIpTagsInvoker 批量添加删除中转IP标签
func (c *NatClient) BatchCreateDeleteTransitIpTagsInvoker(request *model.BatchCreateDeleteTransitIpTagsRequest) *BatchCreateDeleteTransitIpTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeleteTransitIpTags()
	return &BatchCreateDeleteTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTransitIpTag 添加中转IP标签
//
// - 一个中转IP上最多有10个标签。
// - 此接口为幂等接口：
// - 创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateTransitIpTag(request *model.CreateTransitIpTagRequest) (*model.CreateTransitIpTagResponse, error) {
	requestDef := GenReqDefForCreateTransitIpTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitIpTagResponse), nil
	}
}

// CreateTransitIpTagInvoker 添加中转IP标签
func (c *NatClient) CreateTransitIpTagInvoker(request *model.CreateTransitIpTagRequest) *CreateTransitIpTagInvoker {
	requestDef := GenReqDefForCreateTransitIpTag()
	return &CreateTransitIpTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransitIpTag 删除中转IP标签
//
// - 幂等接口：
// - 删除时，不对标签字符集做校验，调用接口前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，key不能为空或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteTransitIpTag(request *model.DeleteTransitIpTagRequest) (*model.DeleteTransitIpTagResponse, error) {
	requestDef := GenReqDefForDeleteTransitIpTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitIpTagResponse), nil
	}
}

// DeleteTransitIpTagInvoker 删除中转IP标签
func (c *NatClient) DeleteTransitIpTagInvoker(request *model.DeleteTransitIpTagRequest) *DeleteTransitIpTagInvoker {
	requestDef := GenReqDefForDeleteTransitIpTag()
	return &DeleteTransitIpTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitIpTags 查询中转IP项目标签
//
// - 查询租户在指定Project和实例类型的所有中转IP标签集合。
// - 标签管理服务需要能够列出当前租户全部已使用的中转IP标签集合，为打中转IP标签和过滤中转IP实例时提供标签联想功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitIpTags(request *model.ListTransitIpTagsRequest) (*model.ListTransitIpTagsResponse, error) {
	requestDef := GenReqDefForListTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpTagsResponse), nil
	}
}

// ListTransitIpTagsInvoker 查询中转IP项目标签
func (c *NatClient) ListTransitIpTagsInvoker(request *model.ListTransitIpTagsRequest) *ListTransitIpTagsInvoker {
	requestDef := GenReqDefForListTransitIpTags()
	return &ListTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitIpsByTags 查询中转IP实例
//
// - 使用标签过滤中转IP实例。
// - 标签管理服务需要提供按标签过滤中转IP服务实例并汇总显示在列表中，需要中转IP服务提供查询能力。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitIpsByTags(request *model.ListTransitIpsByTagsRequest) (*model.ListTransitIpsByTagsResponse, error) {
	requestDef := GenReqDefForListTransitIpsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpsByTagsResponse), nil
	}
}

// ListTransitIpsByTagsInvoker 查询中转IP实例
func (c *NatClient) ListTransitIpsByTagsInvoker(request *model.ListTransitIpsByTagsRequest) *ListTransitIpsByTagsInvoker {
	requestDef := GenReqDefForListTransitIpsByTags()
	return &ListTransitIpsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransitIpTags 查询中转IP标签
//
// - 查询指定中转IP实例的标签信息。
// - 标签管理服务需要使用该接口查询指定中转IP实例的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowTransitIpTags(request *model.ShowTransitIpTagsRequest) (*model.ShowTransitIpTagsResponse, error) {
	requestDef := GenReqDefForShowTransitIpTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitIpTagsResponse), nil
	}
}

// ShowTransitIpTagsInvoker 查询中转IP标签
func (c *NatClient) ShowTransitIpTagsInvoker(request *model.ShowTransitIpTagsRequest) *ShowTransitIpTagsInvoker {
	requestDef := GenReqDefForShowTransitIpTags()
	return &ShowTransitIpTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDeleteNatGatewayTag 批量添加/删除公网NAT网关资源标签
//
// - 为指定公网NAT网关实例批量添加或删除标签。
// - 标签管理服务需要使用该接口批量管理实例的标签。
// - 一个资源上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) BatchCreateDeleteNatGatewayTag(request *model.BatchCreateDeleteNatGatewayTagRequest) (*model.BatchCreateDeleteNatGatewayTagResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteNatGatewayTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteNatGatewayTagResponse), nil
	}
}

// BatchCreateDeleteNatGatewayTagInvoker 批量添加/删除公网NAT网关资源标签
func (c *NatClient) BatchCreateDeleteNatGatewayTagInvoker(request *model.BatchCreateDeleteNatGatewayTagRequest) *BatchCreateDeleteNatGatewayTagInvoker {
	requestDef := GenReqDefForBatchCreateDeleteNatGatewayTag()
	return &BatchCreateDeleteNatGatewayTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDeletePrivateNatTags 批量添加删除私网NAT网关标签
//
// - 为指定私网NAT网关实例批量添加或删除标签
// - 标签管理服务需要使用该接口批量管理私网NAT网关实例的标签。
// - 一个私网NAT网关上最多有10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) BatchCreateDeletePrivateNatTags(request *model.BatchCreateDeletePrivateNatTagsRequest) (*model.BatchCreateDeletePrivateNatTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeletePrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeletePrivateNatTagsResponse), nil
	}
}

// BatchCreateDeletePrivateNatTagsInvoker 批量添加删除私网NAT网关标签
func (c *NatClient) BatchCreateDeletePrivateNatTagsInvoker(request *model.BatchCreateDeletePrivateNatTagsRequest) *BatchCreateDeletePrivateNatTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeletePrivateNatTags()
	return &BatchCreateDeletePrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateNatGatewayTag 添加公网NAT网关资源标签
//
// - 添加公网NAT网关资源标签。一个资源上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateNatGatewayTag(request *model.CreateNatGatewayTagRequest) (*model.CreateNatGatewayTagResponse, error) {
	requestDef := GenReqDefForCreateNatGatewayTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNatGatewayTagResponse), nil
	}
}

// CreateNatGatewayTagInvoker 添加公网NAT网关资源标签
func (c *NatClient) CreateNatGatewayTagInvoker(request *model.CreateNatGatewayTagRequest) *CreateNatGatewayTagInvoker {
	requestDef := GenReqDefForCreateNatGatewayTag()
	return &CreateNatGatewayTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivateNat 创建私网NAT网关
//
// 创建私网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreatePrivateNat(request *model.CreatePrivateNatRequest) (*model.CreatePrivateNatResponse, error) {
	requestDef := GenReqDefForCreatePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateNatResponse), nil
	}
}

// CreatePrivateNatInvoker 创建私网NAT网关
func (c *NatClient) CreatePrivateNatInvoker(request *model.CreatePrivateNatRequest) *CreatePrivateNatInvoker {
	requestDef := GenReqDefForCreatePrivateNat()
	return &CreatePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivateNatTag 添加私网NAT网关标签
//
// - 一个私网NAT网关上最多有10个标签。
// - 此接口为幂等接口：
// - 创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreatePrivateNatTag(request *model.CreatePrivateNatTagRequest) (*model.CreatePrivateNatTagResponse, error) {
	requestDef := GenReqDefForCreatePrivateNatTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateNatTagResponse), nil
	}
}

// CreatePrivateNatTagInvoker 添加私网NAT网关标签
func (c *NatClient) CreatePrivateNatTagInvoker(request *model.CreatePrivateNatTagRequest) *CreatePrivateNatTagInvoker {
	requestDef := GenReqDefForCreatePrivateNatTag()
	return &CreatePrivateNatTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteNatGatewayTag 删除公网NAT网关资源标签
//
// - 删除指定公网NAT网关资源实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteNatGatewayTag(request *model.DeleteNatGatewayTagRequest) (*model.DeleteNatGatewayTagResponse, error) {
	requestDef := GenReqDefForDeleteNatGatewayTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNatGatewayTagResponse), nil
	}
}

// DeleteNatGatewayTagInvoker 删除公网NAT网关资源标签
func (c *NatClient) DeleteNatGatewayTagInvoker(request *model.DeleteNatGatewayTagRequest) *DeleteNatGatewayTagInvoker {
	requestDef := GenReqDefForDeleteNatGatewayTag()
	return &DeleteNatGatewayTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivateNat 删除私网NAT网关
//
// 删除私网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeletePrivateNat(request *model.DeletePrivateNatRequest) (*model.DeletePrivateNatResponse, error) {
	requestDef := GenReqDefForDeletePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateNatResponse), nil
	}
}

// DeletePrivateNatInvoker 删除私网NAT网关
func (c *NatClient) DeletePrivateNatInvoker(request *model.DeletePrivateNatRequest) *DeletePrivateNatInvoker {
	requestDef := GenReqDefForDeletePrivateNat()
	return &DeletePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivateNatTag 删除私网NAT网关标签
//
// - 幂等接口：
// - 删除时，不对标签字符集做校验，调用接口前必须要做encodeURI，服务端需要对接口uri做decodeURI。删除的key不存在报404，key不能为空或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeletePrivateNatTag(request *model.DeletePrivateNatTagRequest) (*model.DeletePrivateNatTagResponse, error) {
	requestDef := GenReqDefForDeletePrivateNatTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateNatTagResponse), nil
	}
}

// DeletePrivateNatTagInvoker 删除私网NAT网关标签
func (c *NatClient) DeletePrivateNatTagInvoker(request *model.DeletePrivateNatTagRequest) *DeletePrivateNatTagInvoker {
	requestDef := GenReqDefForDeletePrivateNatTag()
	return &DeletePrivateNatTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGatewayByTag 查询公网NAT网关资源实例
//
// - 使用标签过滤公网NAT网关资源实例。
// - 标签管理服务需要提供按标签过滤公网NAT网关服务实例并汇总显示在列表中，需要公网NAT网关服务提供查询能力。
// - 资源默认按照创建时间倒序，资源tag也按照创建时间倒序。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGatewayByTag(request *model.ListNatGatewayByTagRequest) (*model.ListNatGatewayByTagResponse, error) {
	requestDef := GenReqDefForListNatGatewayByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewayByTagResponse), nil
	}
}

// ListNatGatewayByTagInvoker 查询公网NAT网关资源实例
func (c *NatClient) ListNatGatewayByTagInvoker(request *model.ListNatGatewayByTagRequest) *ListNatGatewayByTagInvoker {
	requestDef := GenReqDefForListNatGatewayByTag()
	return &ListNatGatewayByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGatewaySpecs 支持创建的NAT网关规格列表
//
// 支持创建的NAT网关规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGatewaySpecs(request *model.ListNatGatewaySpecsRequest) (*model.ListNatGatewaySpecsResponse, error) {
	requestDef := GenReqDefForListNatGatewaySpecs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewaySpecsResponse), nil
	}
}

// ListNatGatewaySpecsInvoker 支持创建的NAT网关规格列表
func (c *NatClient) ListNatGatewaySpecsInvoker(request *model.ListNatGatewaySpecsRequest) *ListNatGatewaySpecsInvoker {
	requestDef := GenReqDefForListNatGatewaySpecs()
	return &ListNatGatewaySpecsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNatGatewayTag 查询公网NAT网关项目标签
//
// - 查询租户在指定项目和公网NAT网关实例类型的所有标签集合。
// - 标签管理服务需要能够列出当前租户全部已使用的标签集合，为各服务Console打标签和过滤实例时提供标签联想功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListNatGatewayTag(request *model.ListNatGatewayTagRequest) (*model.ListNatGatewayTagResponse, error) {
	requestDef := GenReqDefForListNatGatewayTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNatGatewayTagResponse), nil
	}
}

// ListNatGatewayTagInvoker 查询公网NAT网关项目标签
func (c *NatClient) ListNatGatewayTagInvoker(request *model.ListNatGatewayTagRequest) *ListNatGatewayTagInvoker {
	requestDef := GenReqDefForListNatGatewayTag()
	return &ListNatGatewayTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPrivateNatTags 查询私网NAT网关项目标签
//
// - 查询租户在指定Project和实例类型的所有私网NAT网关标签集合。
// - 标签管理服务需要能够列出当前租户全部已使用的私网NAT网关标签集合，为打私网NAT网关标签和过滤私网NAT网关实例时提供标签联想功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListPrivateNatTags(request *model.ListPrivateNatTagsRequest) (*model.ListPrivateNatTagsResponse, error) {
	requestDef := GenReqDefForListPrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatTagsResponse), nil
	}
}

// ListPrivateNatTagsInvoker 查询私网NAT网关项目标签
func (c *NatClient) ListPrivateNatTagsInvoker(request *model.ListPrivateNatTagsRequest) *ListPrivateNatTagsInvoker {
	requestDef := GenReqDefForListPrivateNatTags()
	return &ListPrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateNats 查询私网NAT网关列表
//
// 查询私网NAT网关实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListPrivateNats(request *model.ListPrivateNatsRequest) (*model.ListPrivateNatsResponse, error) {
	requestDef := GenReqDefForListPrivateNats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatsResponse), nil
	}
}

// ListPrivateNatsInvoker 查询私网NAT网关列表
func (c *NatClient) ListPrivateNatsInvoker(request *model.ListPrivateNatsRequest) *ListPrivateNatsInvoker {
	requestDef := GenReqDefForListPrivateNats()
	return &ListPrivateNatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateNatsByTags 查询私网NAT网关实例
//
// - 使用标签过滤私网NAT网关实例。
// - 标签管理服务需要提供按标签过滤私网NAT网关服务实例并汇总显示在列表中，需要私网NAT网关服务提供查询能力。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListPrivateNatsByTags(request *model.ListPrivateNatsByTagsRequest) (*model.ListPrivateNatsByTagsResponse, error) {
	requestDef := GenReqDefForListPrivateNatsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateNatsByTagsResponse), nil
	}
}

// ListPrivateNatsByTagsInvoker 查询私网NAT网关实例
func (c *NatClient) ListPrivateNatsByTagsInvoker(request *model.ListPrivateNatsByTagsRequest) *ListPrivateNatsByTagsInvoker {
	requestDef := GenReqDefForListPrivateNatsByTags()
	return &ListPrivateNatsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecs 查询项目支持的网关规格列表
//
// 查询项目支持的网关规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListSpecs(request *model.ListSpecsRequest) (*model.ListSpecsResponse, error) {
	requestDef := GenReqDefForListSpecs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecsResponse), nil
	}
}

// ListSpecsInvoker 查询项目支持的网关规格列表
func (c *NatClient) ListSpecsInvoker(request *model.ListSpecsRequest) *ListSpecsInvoker {
	requestDef := GenReqDefForListSpecs()
	return &ListSpecsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowNatGatewayTag 查询公网NAT网关资源标签
//
// - 查询指定公网NAT网关实例的标签信息。
// - 标签管理服务需要使用该接口查询指定公网NAT网关实例的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowNatGatewayTag(request *model.ShowNatGatewayTagRequest) (*model.ShowNatGatewayTagResponse, error) {
	requestDef := GenReqDefForShowNatGatewayTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNatGatewayTagResponse), nil
	}
}

// ShowNatGatewayTagInvoker 查询公网NAT网关资源标签
func (c *NatClient) ShowNatGatewayTagInvoker(request *model.ShowNatGatewayTagRequest) *ShowNatGatewayTagInvoker {
	requestDef := GenReqDefForShowNatGatewayTag()
	return &ShowNatGatewayTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateNat 查询指定的私网NAT网关详情
//
// 查询指定的私网NAT网关实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowPrivateNat(request *model.ShowPrivateNatRequest) (*model.ShowPrivateNatResponse, error) {
	requestDef := GenReqDefForShowPrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateNatResponse), nil
	}
}

// ShowPrivateNatInvoker 查询指定的私网NAT网关详情
func (c *NatClient) ShowPrivateNatInvoker(request *model.ShowPrivateNatRequest) *ShowPrivateNatInvoker {
	requestDef := GenReqDefForShowPrivateNat()
	return &ShowPrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateNatTags 查询私网NAT网关标签
//
// - 查询指定私网NAT网关实例的标签信息。
// - 标签管理服务需要使用该接口查询指定私网NAT网关实例的全部标签数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowPrivateNatTags(request *model.ShowPrivateNatTagsRequest) (*model.ShowPrivateNatTagsResponse, error) {
	requestDef := GenReqDefForShowPrivateNatTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateNatTagsResponse), nil
	}
}

// ShowPrivateNatTagsInvoker 查询私网NAT网关标签
func (c *NatClient) ShowPrivateNatTagsInvoker(request *model.ShowPrivateNatTagsRequest) *ShowPrivateNatTagsInvoker {
	requestDef := GenReqDefForShowPrivateNatTags()
	return &ShowPrivateNatTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateNatGatewayToPeriod 公网NAT网关按需转包
//
// 公网NAT网关按需转包。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdateNatGatewayToPeriod(request *model.UpdateNatGatewayToPeriodRequest) (*model.UpdateNatGatewayToPeriodResponse, error) {
	requestDef := GenReqDefForUpdateNatGatewayToPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNatGatewayToPeriodResponse), nil
	}
}

// UpdateNatGatewayToPeriodInvoker 公网NAT网关按需转包
func (c *NatClient) UpdateNatGatewayToPeriodInvoker(request *model.UpdateNatGatewayToPeriodRequest) *UpdateNatGatewayToPeriodInvoker {
	requestDef := GenReqDefForUpdateNatGatewayToPeriod()
	return &UpdateNatGatewayToPeriodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivateNat 更新私网NAT网关
//
// 更新私网NAT网关实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdatePrivateNat(request *model.UpdatePrivateNatRequest) (*model.UpdatePrivateNatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateNat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateNatResponse), nil
	}
}

// UpdatePrivateNatInvoker 更新私网NAT网关
func (c *NatClient) UpdatePrivateNatInvoker(request *model.UpdatePrivateNatRequest) *UpdatePrivateNatInvoker {
	requestDef := GenReqDefForUpdatePrivateNat()
	return &UpdatePrivateNatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTransitIp 创建中转IP
//
// 创建中转IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreateTransitIp(request *model.CreateTransitIpRequest) (*model.CreateTransitIpResponse, error) {
	requestDef := GenReqDefForCreateTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTransitIpResponse), nil
	}
}

// CreateTransitIpInvoker 创建中转IP
func (c *NatClient) CreateTransitIpInvoker(request *model.CreateTransitIpRequest) *CreateTransitIpInvoker {
	requestDef := GenReqDefForCreateTransitIp()
	return &CreateTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTransitIp 删除中转IP
//
// 删除中转IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeleteTransitIp(request *model.DeleteTransitIpRequest) (*model.DeleteTransitIpResponse, error) {
	requestDef := GenReqDefForDeleteTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTransitIpResponse), nil
	}
}

// DeleteTransitIpInvoker 删除中转IP
func (c *NatClient) DeleteTransitIpInvoker(request *model.DeleteTransitIpRequest) *DeleteTransitIpInvoker {
	requestDef := GenReqDefForDeleteTransitIp()
	return &DeleteTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTransitIps 查询中转IP列表
//
// 查询中转IP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListTransitIps(request *model.ListTransitIpsRequest) (*model.ListTransitIpsResponse, error) {
	requestDef := GenReqDefForListTransitIps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransitIpsResponse), nil
	}
}

// ListTransitIpsInvoker 查询中转IP列表
func (c *NatClient) ListTransitIpsInvoker(request *model.ListTransitIpsRequest) *ListTransitIpsInvoker {
	requestDef := GenReqDefForListTransitIps()
	return &ListTransitIpsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTransitIp 查询指定的中转IP详情
//
// 查询中转IP详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowTransitIp(request *model.ShowTransitIpRequest) (*model.ShowTransitIpResponse, error) {
	requestDef := GenReqDefForShowTransitIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTransitIpResponse), nil
	}
}

// ShowTransitIpInvoker 查询指定的中转IP详情
func (c *NatClient) ShowTransitIpInvoker(request *model.ShowTransitIpRequest) *ShowTransitIpInvoker {
	requestDef := GenReqDefForShowTransitIp()
	return &ShowTransitIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreatePrivateSnat 创建SNAT规则
//
// 创建SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) CreatePrivateSnat(request *model.CreatePrivateSnatRequest) (*model.CreatePrivateSnatResponse, error) {
	requestDef := GenReqDefForCreatePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateSnatResponse), nil
	}
}

// CreatePrivateSnatInvoker 创建SNAT规则
func (c *NatClient) CreatePrivateSnatInvoker(request *model.CreatePrivateSnatRequest) *CreatePrivateSnatInvoker {
	requestDef := GenReqDefForCreatePrivateSnat()
	return &CreatePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeletePrivateSnat 删除SNAT规则
//
// 删除指定的SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) DeletePrivateSnat(request *model.DeletePrivateSnatRequest) (*model.DeletePrivateSnatResponse, error) {
	requestDef := GenReqDefForDeletePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateSnatResponse), nil
	}
}

// DeletePrivateSnatInvoker 删除SNAT规则
func (c *NatClient) DeletePrivateSnatInvoker(request *model.DeletePrivateSnatRequest) *DeletePrivateSnatInvoker {
	requestDef := GenReqDefForDeletePrivateSnat()
	return &DeletePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPrivateSnats 查询SNAT规则列表
//
// 查询SNAT规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ListPrivateSnats(request *model.ListPrivateSnatsRequest) (*model.ListPrivateSnatsResponse, error) {
	requestDef := GenReqDefForListPrivateSnats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateSnatsResponse), nil
	}
}

// ListPrivateSnatsInvoker 查询SNAT规则列表
func (c *NatClient) ListPrivateSnatsInvoker(request *model.ListPrivateSnatsRequest) *ListPrivateSnatsInvoker {
	requestDef := GenReqDefForListPrivateSnats()
	return &ListPrivateSnatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowPrivateSnat 查询指定的SNAT规则详情
//
// 查询指定的SNAT规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) ShowPrivateSnat(request *model.ShowPrivateSnatRequest) (*model.ShowPrivateSnatResponse, error) {
	requestDef := GenReqDefForShowPrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateSnatResponse), nil
	}
}

// ShowPrivateSnatInvoker 查询指定的SNAT规则详情
func (c *NatClient) ShowPrivateSnatInvoker(request *model.ShowPrivateSnatRequest) *ShowPrivateSnatInvoker {
	requestDef := GenReqDefForShowPrivateSnat()
	return &ShowPrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdatePrivateSnat 更新SNAT规则
//
// 更新指定的SNAT规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *NatClient) UpdatePrivateSnat(request *model.UpdatePrivateSnatRequest) (*model.UpdatePrivateSnatResponse, error) {
	requestDef := GenReqDefForUpdatePrivateSnat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateSnatResponse), nil
	}
}

// UpdatePrivateSnatInvoker 更新SNAT规则
func (c *NatClient) UpdatePrivateSnatInvoker(request *model.UpdatePrivateSnatRequest) *UpdatePrivateSnatInvoker {
	requestDef := GenReqDefForUpdatePrivateSnat()
	return &UpdatePrivateSnatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
