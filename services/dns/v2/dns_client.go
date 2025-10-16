package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dns/v2/model"
)

type DnsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDnsClient(hcClient *httpclient.HcHttpClient) *DnsClient {
	return &DnsClient{HcClient: hcClient}
}

func DnsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AssociateEndpointIpaddress 终端节点绑定IP地址
//
// 终端节点绑定IP地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) AssociateEndpointIpaddress(request *model.AssociateEndpointIpaddressRequest) (*model.AssociateEndpointIpaddressResponse, error) {
	requestDef := GenReqDefForAssociateEndpointIpaddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateEndpointIpaddressResponse), nil
	}
}

// AssociateEndpointIpaddressInvoker 终端节点绑定IP地址
func (c *DnsClient) AssociateEndpointIpaddressInvoker(request *model.AssociateEndpointIpaddressRequest) *AssociateEndpointIpaddressInvoker {
	requestDef := GenReqDefForAssociateEndpointIpaddress()
	return &AssociateEndpointIpaddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateResolverRuleRouter 解析器转发规则关联VPC
//
// 解析器转发规则关联VPC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) AssociateResolverRuleRouter(request *model.AssociateResolverRuleRouterRequest) (*model.AssociateResolverRuleRouterResponse, error) {
	requestDef := GenReqDefForAssociateResolverRuleRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateResolverRuleRouterResponse), nil
	}
}

// AssociateResolverRuleRouterInvoker 解析器转发规则关联VPC
func (c *DnsClient) AssociateResolverRuleRouterInvoker(request *model.AssociateResolverRuleRouterRequest) *AssociateResolverRuleRouterInvoker {
	requestDef := GenReqDefForAssociateResolverRuleRouter()
	return &AssociateResolverRuleRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRouter 在内网域名上关联VPC
//
// 当您的内网域名创建完成后，可以通过调用此接口为内网域名关联新的VPC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) AssociateRouter(request *model.AssociateRouterRequest) (*model.AssociateRouterResponse, error) {
	requestDef := GenReqDefForAssociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouterResponse), nil
	}
}

// AssociateRouterInvoker 在内网域名上关联VPC
func (c *DnsClient) AssociateRouterInvoker(request *model.AssociateRouterRequest) *AssociateRouterInvoker {
	requestDef := GenReqDefForAssociateRouter()
	return &AssociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateTag 为指定实例批量添加或删除标签
//
// 为指定实例批量添加或删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchCreateTag(request *model.BatchCreateTagRequest) (*model.BatchCreateTagResponse, error) {
	requestDef := GenReqDefForBatchCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagResponse), nil
	}
}

// BatchCreateTagInvoker 为指定实例批量添加或删除标签
func (c *DnsClient) BatchCreateTagInvoker(request *model.BatchCreateTagRequest) *BatchCreateTagInvoker {
	requestDef := GenReqDefForBatchCreateTag()
	return &BatchCreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePtrRecords 批量删除反向解析
//
// 批量删除反向解析。本接口为原子操作，所有记录应全部删除成功或全部失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchDeletePtrRecords(request *model.BatchDeletePtrRecordsRequest) (*model.BatchDeletePtrRecordsResponse, error) {
	requestDef := GenReqDefForBatchDeletePtrRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePtrRecordsResponse), nil
	}
}

// BatchDeletePtrRecordsInvoker 批量删除反向解析
func (c *DnsClient) BatchDeletePtrRecordsInvoker(request *model.BatchDeletePtrRecordsRequest) *BatchDeletePtrRecordsInvoker {
	requestDef := GenReqDefForBatchDeletePtrRecords()
	return &BatchDeletePtrRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRecordSetWithLine 批量删除域名下的记录集
//
// 批量删除域名下的记录集，当删除的资源不存在时，则默认删除成功。
// 响应结果中只包含本次实际删除的资源。
// 支持公网域名和内网域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchDeleteRecordSetWithLine(request *model.BatchDeleteRecordSetWithLineRequest) (*model.BatchDeleteRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForBatchDeleteRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRecordSetWithLineResponse), nil
	}
}

// BatchDeleteRecordSetWithLineInvoker 批量删除域名下的记录集
func (c *DnsClient) BatchDeleteRecordSetWithLineInvoker(request *model.BatchDeleteRecordSetWithLineRequest) *BatchDeleteRecordSetWithLineInvoker {
	requestDef := GenReqDefForBatchDeleteRecordSetWithLine()
	return &BatchDeleteRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteRecordSets 批量删除记录集
//
// 批量删除记录集。
// 响应结果中只包含本次实际删除的记录集。
// 支持批量删除公网域名和内网域名的记录集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchDeleteRecordSets(request *model.BatchDeleteRecordSetsRequest) (*model.BatchDeleteRecordSetsResponse, error) {
	requestDef := GenReqDefForBatchDeleteRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteRecordSetsResponse), nil
	}
}

// BatchDeleteRecordSetsInvoker 批量删除记录集
func (c *DnsClient) BatchDeleteRecordSetsInvoker(request *model.BatchDeleteRecordSetsRequest) *BatchDeleteRecordSetsInvoker {
	requestDef := GenReqDefForBatchDeleteRecordSets()
	return &BatchDeleteRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteZones 批量删除域名
//
// 批量删除域名。
// 本接口为原子操作，所有记录应全部删除成功或全部失败。
// 支持公网域名、内网域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchDeleteZones(request *model.BatchDeleteZonesRequest) (*model.BatchDeleteZonesResponse, error) {
	requestDef := GenReqDefForBatchDeleteZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteZonesResponse), nil
	}
}

// BatchDeleteZonesInvoker 批量删除域名
func (c *DnsClient) BatchDeleteZonesInvoker(request *model.BatchDeleteZonesRequest) *BatchDeleteZonesInvoker {
	requestDef := GenReqDefForBatchDeleteZones()
	return &BatchDeleteZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetRecordSetsStatus 批量设置记录集状态
//
// 批量设置记录集状态。
// 响应结果中只包含本次实际更新的记录集。
// 支持公网域名和内网域名的记录集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchSetRecordSetsStatus(request *model.BatchSetRecordSetsStatusRequest) (*model.BatchSetRecordSetsStatusResponse, error) {
	requestDef := GenReqDefForBatchSetRecordSetsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetRecordSetsStatusResponse), nil
	}
}

// BatchSetRecordSetsStatusInvoker 批量设置记录集状态
func (c *DnsClient) BatchSetRecordSetsStatusInvoker(request *model.BatchSetRecordSetsStatusRequest) *BatchSetRecordSetsStatusInvoker {
	requestDef := GenReqDefForBatchSetRecordSetsStatus()
	return &BatchSetRecordSetsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetZonesStatus 批量设置域名状态
//
// 批量设置域名状态。
// 响应结果中只包含本次实际更新的域名。
// 支持公网域名、内网域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchSetZonesStatus(request *model.BatchSetZonesStatusRequest) (*model.BatchSetZonesStatusResponse, error) {
	requestDef := GenReqDefForBatchSetZonesStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetZonesStatusResponse), nil
	}
}

// BatchSetZonesStatusInvoker 批量设置域名状态
func (c *DnsClient) BatchSetZonesStatusInvoker(request *model.BatchSetZonesStatusRequest) *BatchSetZonesStatusInvoker {
	requestDef := GenReqDefForBatchSetZonesStatus()
	return &BatchSetZonesStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateRecordSetWithLine 批量修改记录集
//
// 批量修改记录集。属于原子性操作，请求记录集将全部完成修改，或不做任何修改。
// 仅公网域名支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) BatchUpdateRecordSetWithLine(request *model.BatchUpdateRecordSetWithLineRequest) (*model.BatchUpdateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForBatchUpdateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateRecordSetWithLineResponse), nil
	}
}

// BatchUpdateRecordSetWithLineInvoker 批量修改记录集
func (c *DnsClient) BatchUpdateRecordSetWithLineInvoker(request *model.BatchUpdateRecordSetWithLineRequest) *BatchUpdateRecordSetWithLineInvoker {
	requestDef := GenReqDefForBatchUpdateRecordSetWithLine()
	return &BatchUpdateRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorizeTxtRecord 创建公网子域名授权
//
// 当创建子域名时提示“域名与其他租户冲突，你需要添加TXT授权校验”，通过调用当前接口生成子域名授权的TXT记录验证信息。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// &gt; TXT记录验证信息生成后，请前往主域名所属的DNS服务商处添加相应的TXT类型解析记录，主机记录和记录值与验证信息保持一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateAuthorizeTxtRecord(request *model.CreateAuthorizeTxtRecordRequest) (*model.CreateAuthorizeTxtRecordResponse, error) {
	requestDef := GenReqDefForCreateAuthorizeTxtRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizeTxtRecordResponse), nil
	}
}

// CreateAuthorizeTxtRecordInvoker 创建公网子域名授权
func (c *DnsClient) CreateAuthorizeTxtRecordInvoker(request *model.CreateAuthorizeTxtRecordRequest) *CreateAuthorizeTxtRecordInvoker {
	requestDef := GenReqDefForCreateAuthorizeTxtRecord()
	return &CreateAuthorizeTxtRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorizeTxtRecordVerification 验证公网子域名授权
//
// 用户在主域名所属DNS服务商处添加TXT类型解析记录后，调用当前接口验证子域名授权状态。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateAuthorizeTxtRecordVerification(request *model.CreateAuthorizeTxtRecordVerificationRequest) (*model.CreateAuthorizeTxtRecordVerificationResponse, error) {
	requestDef := GenReqDefForCreateAuthorizeTxtRecordVerification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizeTxtRecordVerificationResponse), nil
	}
}

// CreateAuthorizeTxtRecordVerificationInvoker 验证公网子域名授权
func (c *DnsClient) CreateAuthorizeTxtRecordVerificationInvoker(request *model.CreateAuthorizeTxtRecordVerificationRequest) *CreateAuthorizeTxtRecordVerificationInvoker {
	requestDef := GenReqDefForCreateAuthorizeTxtRecordVerification()
	return &CreateAuthorizeTxtRecordVerificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomLine 创建自定义线路
//
// 创建自定义线路。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateCustomLine(request *model.CreateCustomLineRequest) (*model.CreateCustomLineResponse, error) {
	requestDef := GenReqDefForCreateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomLineResponse), nil
	}
}

// CreateCustomLineInvoker 创建自定义线路
func (c *DnsClient) CreateCustomLineInvoker(request *model.CreateCustomLineRequest) *CreateCustomLineInvoker {
	requestDef := GenReqDefForCreateCustomLine()
	return &CreateCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建终端节点
//
// 创建终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建终端节点
func (c *DnsClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLineGroup 创建线路分组
//
// 创建线路分组。该接口部分区域未上线，如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateLineGroup(request *model.CreateLineGroupRequest) (*model.CreateLineGroupResponse, error) {
	requestDef := GenReqDefForCreateLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLineGroupResponse), nil
	}
}

// CreateLineGroupInvoker 创建线路分组
func (c *DnsClient) CreateLineGroupInvoker(request *model.CreateLineGroupRequest) *CreateLineGroupInvoker {
	requestDef := GenReqDefForCreateLineGroup()
	return &CreateLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivateZone 创建内网域名
//
// 内网域名是指在VPC中生效的域名，内网域名创建后，用户可以将其与私网IP地址相关联，为云服务提供VPC内的内网域名解析服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreatePrivateZone(request *model.CreatePrivateZoneRequest) (*model.CreatePrivateZoneResponse, error) {
	requestDef := GenReqDefForCreatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateZoneResponse), nil
	}
}

// CreatePrivateZoneInvoker 创建内网域名
func (c *DnsClient) CreatePrivateZoneInvoker(request *model.CreatePrivateZoneRequest) *CreatePrivateZoneInvoker {
	requestDef := GenReqDefForCreatePrivateZone()
	return &CreatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePublicZone 创建公网域名
//
// 您在使用华为云云解析服务为自己注册的域名配置DNS解析之前，需要先将域名添加至云解析服务控制台。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreatePublicZone(request *model.CreatePublicZoneRequest) (*model.CreatePublicZoneResponse, error) {
	requestDef := GenReqDefForCreatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicZoneResponse), nil
	}
}

// CreatePublicZoneInvoker 创建公网域名
func (c *DnsClient) CreatePublicZoneInvoker(request *model.CreatePublicZoneRequest) *CreatePublicZoneInvoker {
	requestDef := GenReqDefForCreatePublicZone()
	return &CreatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSetWithBatchLines 批量线路创建记录集
//
// 批量线路创建记录集。属于原子性操作，如果存在一个参数校验不通过，则创建失败。仅公网域名支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSetWithBatchLines(request *model.CreateRecordSetWithBatchLinesRequest) (*model.CreateRecordSetWithBatchLinesResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithBatchLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithBatchLinesResponse), nil
	}
}

// CreateRecordSetWithBatchLinesInvoker 批量线路创建记录集
func (c *DnsClient) CreateRecordSetWithBatchLinesInvoker(request *model.CreateRecordSetWithBatchLinesRequest) *CreateRecordSetWithBatchLinesInvoker {
	requestDef := GenReqDefForCreateRecordSetWithBatchLines()
	return &CreateRecordSetWithBatchLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResolverRule 创建解析器转发规则
//
// 创建解析器转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateResolverRule(request *model.CreateResolverRuleRequest) (*model.CreateResolverRuleResponse, error) {
	requestDef := GenReqDefForCreateResolverRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResolverRuleResponse), nil
	}
}

// CreateResolverRuleInvoker 创建解析器转发规则
func (c *DnsClient) CreateResolverRuleInvoker(request *model.CreateResolverRuleRequest) *CreateResolverRuleInvoker {
	requestDef := GenReqDefForCreateResolverRule()
	return &CreateResolverRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 为指定实例添加标签
//
// 为指定实例添加标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 为指定实例添加标签
func (c *DnsClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomLine 删除自定义线路
//
// 删除自定义线路。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteCustomLine(request *model.DeleteCustomLineRequest) (*model.DeleteCustomLineResponse, error) {
	requestDef := GenReqDefForDeleteCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomLineResponse), nil
	}
}

// DeleteCustomLineInvoker 删除自定义线路
func (c *DnsClient) DeleteCustomLineInvoker(request *model.DeleteCustomLineRequest) *DeleteCustomLineInvoker {
	requestDef := GenReqDefForDeleteCustomLine()
	return &DeleteCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpoint 删除终端节点
//
// 删除终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

// DeleteEndpointInvoker 删除终端节点
func (c *DnsClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLineGroup 删除线路分组
//
// 删除线路分组。该接口部分区域未上线，如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteLineGroup(request *model.DeleteLineGroupRequest) (*model.DeleteLineGroupResponse, error) {
	requestDef := GenReqDefForDeleteLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLineGroupResponse), nil
	}
}

// DeleteLineGroupInvoker 删除线路分组
func (c *DnsClient) DeleteLineGroupInvoker(request *model.DeleteLineGroupRequest) *DeleteLineGroupInvoker {
	requestDef := GenReqDefForDeleteLineGroup()
	return &DeleteLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivateZone 删除内网域名
//
// 当您的内网域名不再使用时，您可以通过调用此接口将其删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeletePrivateZone(request *model.DeletePrivateZoneRequest) (*model.DeletePrivateZoneResponse, error) {
	requestDef := GenReqDefForDeletePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateZoneResponse), nil
	}
}

// DeletePrivateZoneInvoker 删除内网域名
func (c *DnsClient) DeletePrivateZoneInvoker(request *model.DeletePrivateZoneRequest) *DeletePrivateZoneInvoker {
	requestDef := GenReqDefForDeletePrivateZone()
	return &DeletePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePublicZone 删除公网域名
//
// 当您的公网域名不再使用时，您可以通过调用此接口将其删除。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeletePublicZone(request *model.DeletePublicZoneRequest) (*model.DeletePublicZoneResponse, error) {
	requestDef := GenReqDefForDeletePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicZoneResponse), nil
	}
}

// DeletePublicZoneInvoker 删除公网域名
func (c *DnsClient) DeletePublicZoneInvoker(request *model.DeletePublicZoneRequest) *DeletePublicZoneInvoker {
	requestDef := GenReqDefForDeletePublicZone()
	return &DeletePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResolverRule 删除解析器转发规则
//
// 删除解析器转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteResolverRule(request *model.DeleteResolverRuleRequest) (*model.DeleteResolverRuleResponse, error) {
	requestDef := GenReqDefForDeleteResolverRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResolverRuleResponse), nil
	}
}

// DeleteResolverRuleInvoker 删除解析器转发规则
func (c *DnsClient) DeleteResolverRuleInvoker(request *model.DeleteResolverRuleRequest) *DeleteResolverRuleInvoker {
	requestDef := GenReqDefForDeleteResolverRule()
	return &DeleteResolverRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除资源标签
//
// 删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除资源标签
func (c *DnsClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateEndpointIpaddress 终端节点解绑IP地址
//
// 终端节点解绑IP地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DisassociateEndpointIpaddress(request *model.DisassociateEndpointIpaddressRequest) (*model.DisassociateEndpointIpaddressResponse, error) {
	requestDef := GenReqDefForDisassociateEndpointIpaddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateEndpointIpaddressResponse), nil
	}
}

// DisassociateEndpointIpaddressInvoker 终端节点解绑IP地址
func (c *DnsClient) DisassociateEndpointIpaddressInvoker(request *model.DisassociateEndpointIpaddressRequest) *DisassociateEndpointIpaddressInvoker {
	requestDef := GenReqDefForDisassociateEndpointIpaddress()
	return &DisassociateEndpointIpaddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateResolverRuleRouter 解析器转发规则解关联VPC
//
// 解析器转发规则解关联VPC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DisassociateResolverRuleRouter(request *model.DisassociateResolverRuleRouterRequest) (*model.DisassociateResolverRuleRouterResponse, error) {
	requestDef := GenReqDefForDisassociateResolverRuleRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateResolverRuleRouterResponse), nil
	}
}

// DisassociateResolverRuleRouterInvoker 解析器转发规则解关联VPC
func (c *DnsClient) DisassociateResolverRuleRouterInvoker(request *model.DisassociateResolverRuleRouterRequest) *DisassociateResolverRuleRouterInvoker {
	requestDef := GenReqDefForDisassociateResolverRuleRouter()
	return &DisassociateResolverRuleRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRouter 在内网域名上解关联VPC
//
// 当您的内网域名创建完成后，可以通过调用此接口为内网域名解除已关联的VPC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DisassociateRouter(request *model.DisassociateRouterRequest) (*model.DisassociateRouterResponse, error) {
	requestDef := GenReqDefForDisassociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouterResponse), nil
	}
}

// DisassociateRouterInvoker 在内网域名上解关联VPC
func (c *DnsClient) DisassociateRouterInvoker(request *model.DisassociateRouterRequest) *DisassociateRouterInvoker {
	requestDef := GenReqDefForDisassociateRouter()
	return &DisassociateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersions 查询API版本信息列表
//
// 查询云解析服务支持的所有API版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

// ListApiVersionsInvoker 查询API版本信息列表
func (c *DnsClient) ListApiVersionsInvoker(request *model.ListApiVersionsRequest) *ListApiVersionsInvoker {
	requestDef := GenReqDefForListApiVersions()
	return &ListApiVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomLine 查询自定义线路
//
// 查询自定义线路。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListCustomLine(request *model.ListCustomLineRequest) (*model.ListCustomLineResponse, error) {
	requestDef := GenReqDefForListCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomLineResponse), nil
	}
}

// ListCustomLineInvoker 查询自定义线路
func (c *DnsClient) ListCustomLineInvoker(request *model.ListCustomLineRequest) *ListCustomLineInvoker {
	requestDef := GenReqDefForListCustomLine()
	return &ListCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointIpaddresses 查询IP地址列表
//
// 查询终端节点下的IP地址列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListEndpointIpaddresses(request *model.ListEndpointIpaddressesRequest) (*model.ListEndpointIpaddressesResponse, error) {
	requestDef := GenReqDefForListEndpointIpaddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointIpaddressesResponse), nil
	}
}

// ListEndpointIpaddressesInvoker 查询IP地址列表
func (c *DnsClient) ListEndpointIpaddressesInvoker(request *model.ListEndpointIpaddressesRequest) *ListEndpointIpaddressesInvoker {
	requestDef := GenReqDefForListEndpointIpaddresses()
	return &ListEndpointIpaddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointVpcs 查询终端节点VPC列表
//
// 查询终端节点VPC列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListEndpointVpcs(request *model.ListEndpointVpcsRequest) (*model.ListEndpointVpcsResponse, error) {
	requestDef := GenReqDefForListEndpointVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointVpcsResponse), nil
	}
}

// ListEndpointVpcsInvoker 查询终端节点VPC列表
func (c *DnsClient) ListEndpointVpcsInvoker(request *model.ListEndpointVpcsRequest) *ListEndpointVpcsInvoker {
	requestDef := GenReqDefForListEndpointVpcs()
	return &ListEndpointVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询终端节点列表
//
// 查询终端节点列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询终端节点列表
func (c *DnsClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLineGroups 查询线路分组列表
//
// 查询线路分组列表。该接口部分区域未上线，如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListLineGroups(request *model.ListLineGroupsRequest) (*model.ListLineGroupsResponse, error) {
	requestDef := GenReqDefForListLineGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLineGroupsResponse), nil
	}
}

// ListLineGroupsInvoker 查询线路分组列表
func (c *DnsClient) ListLineGroupsInvoker(request *model.ListLineGroupsRequest) *ListLineGroupsInvoker {
	requestDef := GenReqDefForListLineGroups()
	return &ListLineGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNameServers 查询名称服务器列表
//
// 查询名称服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListNameServers(request *model.ListNameServersRequest) (*model.ListNameServersResponse, error) {
	requestDef := GenReqDefForListNameServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNameServersResponse), nil
	}
}

// ListNameServersInvoker 查询名称服务器列表
func (c *DnsClient) ListNameServersInvoker(request *model.ListNameServersRequest) *ListNameServersInvoker {
	requestDef := GenReqDefForListNameServers()
	return &ListNameServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateZones 查询内网域名列表
//
// 当您的内网域名创建成功后，您可以通过调用此接口查询单个内网域名信息，包括域名、ID、状态、记录集个数、企业项目、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPrivateZones(request *model.ListPrivateZonesRequest) (*model.ListPrivateZonesResponse, error) {
	requestDef := GenReqDefForListPrivateZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateZonesResponse), nil
	}
}

// ListPrivateZonesInvoker 查询内网域名列表
func (c *DnsClient) ListPrivateZonesInvoker(request *model.ListPrivateZonesRequest) *ListPrivateZonesInvoker {
	requestDef := GenReqDefForListPrivateZones()
	return &ListPrivateZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicZones 查询公网域名列表
//
// 当您的公网域名创建成功后，您可以通过调用此接口查询所有公网域名信息，包括域名、ID、状态、记录集个数、企业项目、标签、TTL、创建时间、修改时间、描述等。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPublicZones(request *model.ListPublicZonesRequest) (*model.ListPublicZonesResponse, error) {
	requestDef := GenReqDefForListPublicZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicZonesResponse), nil
	}
}

// ListPublicZonesInvoker 查询公网域名列表
func (c *DnsClient) ListPublicZonesInvoker(request *model.ListPublicZonesRequest) *ListPublicZonesInvoker {
	requestDef := GenReqDefForListPublicZones()
	return &ListPublicZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResolverRules 查询解析器转发规则列表
//
// 查询解析器转发规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListResolverRules(request *model.ListResolverRulesRequest) (*model.ListResolverRulesResponse, error) {
	requestDef := GenReqDefForListResolverRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResolverRulesResponse), nil
	}
}

// ListResolverRulesInvoker 查询解析器转发规则列表
func (c *DnsClient) ListResolverRulesInvoker(request *model.ListResolverRulesRequest) *ListResolverRulesInvoker {
	requestDef := GenReqDefForListResolverRules()
	return &ListResolverRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTag 使用标签查询资源实例
//
// 使用标签查询资源实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListTag(request *model.ListTagRequest) (*model.ListTagResponse, error) {
	requestDef := GenReqDefForListTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResponse), nil
	}
}

// ListTagInvoker 使用标签查询资源实例
func (c *DnsClient) ListTagInvoker(request *model.ListTagRequest) *ListTagInvoker {
	requestDef := GenReqDefForListTag()
	return &ListTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询指定实例类型的所有标签集合
//
// 查询指定实例类型的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询指定实例类型的所有标签集合
func (c *DnsClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetPrivateZoneProxyPattern 设置内网域名的子域名递归解析代理
//
// 当您的内网域名创建成功后，您可以通过调用此接口设置开启或者关闭内网域名的子域名递归解析代理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) SetPrivateZoneProxyPattern(request *model.SetPrivateZoneProxyPatternRequest) (*model.SetPrivateZoneProxyPatternResponse, error) {
	requestDef := GenReqDefForSetPrivateZoneProxyPattern()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetPrivateZoneProxyPatternResponse), nil
	}
}

// SetPrivateZoneProxyPatternInvoker 设置内网域名的子域名递归解析代理
func (c *DnsClient) SetPrivateZoneProxyPatternInvoker(request *model.SetPrivateZoneProxyPatternRequest) *SetPrivateZoneProxyPatternInvoker {
	requestDef := GenReqDefForSetPrivateZoneProxyPattern()
	return &SetPrivateZoneProxyPatternInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiInfo 查询指定版本号的API版本信息
//
// 查询指定版本号的云解析服务API版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowApiInfo(request *model.ShowApiInfoRequest) (*model.ShowApiInfoResponse, error) {
	requestDef := GenReqDefForShowApiInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiInfoResponse), nil
	}
}

// ShowApiInfoInvoker 查询指定版本号的API版本信息
func (c *DnsClient) ShowApiInfoInvoker(request *model.ShowApiInfoRequest) *ShowApiInfoInvoker {
	requestDef := GenReqDefForShowApiInfo()
	return &ShowApiInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthorizeTxtRecord 查询公网子域名授权
//
// 查询已生成的子域名授权TXT记录验证信息。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowAuthorizeTxtRecord(request *model.ShowAuthorizeTxtRecordRequest) (*model.ShowAuthorizeTxtRecordResponse, error) {
	requestDef := GenReqDefForShowAuthorizeTxtRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthorizeTxtRecordResponse), nil
	}
}

// ShowAuthorizeTxtRecordInvoker 查询公网子域名授权
func (c *DnsClient) ShowAuthorizeTxtRecordInvoker(request *model.ShowAuthorizeTxtRecordRequest) *ShowAuthorizeTxtRecordInvoker {
	requestDef := GenReqDefForShowAuthorizeTxtRecord()
	return &ShowAuthorizeTxtRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDomainQuota 查询租户配额
//
// 查询租户在DNS服务下的资源配额，包括公网域名配额、内网域名配额、记录集配额、反向解析配额、自定义线路配额、线路分组配额、入站终端节点配额、出站终端节点配额、转发规则配额等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowDomainQuota(request *model.ShowDomainQuotaRequest) (*model.ShowDomainQuotaResponse, error) {
	requestDef := GenReqDefForShowDomainQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainQuotaResponse), nil
	}
}

// ShowDomainQuotaInvoker 查询租户配额
func (c *DnsClient) ShowDomainQuotaInvoker(request *model.ShowDomainQuotaRequest) *ShowDomainQuotaInvoker {
	requestDef := GenReqDefForShowDomainQuota()
	return &ShowDomainQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEndpoint 查询终端节点
//
// 查询终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowEndpoint(request *model.ShowEndpointRequest) (*model.ShowEndpointResponse, error) {
	requestDef := GenReqDefForShowEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEndpointResponse), nil
	}
}

// ShowEndpointInvoker 查询终端节点
func (c *DnsClient) ShowEndpointInvoker(request *model.ShowEndpointRequest) *ShowEndpointInvoker {
	requestDef := GenReqDefForShowEndpoint()
	return &ShowEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLineGroup 查询线路分组
//
// 查询线路分组。该接口部分区域未上线，如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowLineGroup(request *model.ShowLineGroupRequest) (*model.ShowLineGroupResponse, error) {
	requestDef := GenReqDefForShowLineGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLineGroupResponse), nil
	}
}

// ShowLineGroupInvoker 查询线路分组
func (c *DnsClient) ShowLineGroupInvoker(request *model.ShowLineGroupRequest) *ShowLineGroupInvoker {
	requestDef := GenReqDefForShowLineGroup()
	return &ShowLineGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateZone 查询内网域名
//
// 当您的内网域名创建成功后，您可以通过调用此接口查询单个内网域名信息，包括域名、ID、状态、记录集个数、企业项目、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPrivateZone(request *model.ShowPrivateZoneRequest) (*model.ShowPrivateZoneResponse, error) {
	requestDef := GenReqDefForShowPrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneResponse), nil
	}
}

// ShowPrivateZoneInvoker 查询内网域名
func (c *DnsClient) ShowPrivateZoneInvoker(request *model.ShowPrivateZoneRequest) *ShowPrivateZoneInvoker {
	requestDef := GenReqDefForShowPrivateZone()
	return &ShowPrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateZoneNameServer 查询内网域名的名称服务器
//
// 当您的内网域名创建成功后，您可以通过调用此接口查询内网域名的名称服务器信息，包括优先级、DNS服务器地址等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPrivateZoneNameServer(request *model.ShowPrivateZoneNameServerRequest) (*model.ShowPrivateZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPrivateZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

// ShowPrivateZoneNameServerInvoker 查询内网域名的名称服务器
func (c *DnsClient) ShowPrivateZoneNameServerInvoker(request *model.ShowPrivateZoneNameServerRequest) *ShowPrivateZoneNameServerInvoker {
	requestDef := GenReqDefForShowPrivateZoneNameServer()
	return &ShowPrivateZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicZone 查询公网域名
//
// 当您的公网域名创建成功后，您可以通过调用此接口查询单个公网域名信息，包括域名、ID、状态、记录集个数、企业项目、标签、TTL、创建时间、修改时间、描述等。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPublicZone(request *model.ShowPublicZoneRequest) (*model.ShowPublicZoneResponse, error) {
	requestDef := GenReqDefForShowPublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneResponse), nil
	}
}

// ShowPublicZoneInvoker 查询公网域名
func (c *DnsClient) ShowPublicZoneInvoker(request *model.ShowPublicZoneRequest) *ShowPublicZoneInvoker {
	requestDef := GenReqDefForShowPublicZone()
	return &ShowPublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicZoneNameServer 查询公网域名的名称服务器
//
// 当您的公网域名创建成功后，您可以通过调用此接口查询公网域名的名称服务器信息，包括主机名、优先级等。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPublicZoneNameServer(request *model.ShowPublicZoneNameServerRequest) (*model.ShowPublicZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPublicZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

// ShowPublicZoneNameServerInvoker 查询公网域名的名称服务器
func (c *DnsClient) ShowPublicZoneNameServerInvoker(request *model.ShowPublicZoneNameServerRequest) *ShowPublicZoneNameServerInvoker {
	requestDef := GenReqDefForShowPublicZoneNameServer()
	return &ShowPublicZoneNameServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResolverRule 查询解析器转发规则
//
// 查询解析器转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowResolverRule(request *model.ShowResolverRuleRequest) (*model.ShowResolverRuleResponse, error) {
	requestDef := GenReqDefForShowResolverRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResolverRuleResponse), nil
	}
}

// ShowResolverRuleInvoker 查询解析器转发规则
func (c *DnsClient) ShowResolverRuleInvoker(request *model.ShowResolverRuleRequest) *ShowResolverRuleInvoker {
	requestDef := GenReqDefForShowResolverRule()
	return &ShowResolverRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTag 查询指定实例的标签信息
//
// 查询指定实例的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

// ShowResourceTagInvoker 查询指定实例的标签信息
func (c *DnsClient) ShowResourceTagInvoker(request *model.ShowResourceTagRequest) *ShowResourceTagInvoker {
	requestDef := GenReqDefForShowResourceTag()
	return &ShowResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCustomLine 修改自定义线路
//
// 修改自定义线路。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateCustomLine(request *model.UpdateCustomLineRequest) (*model.UpdateCustomLineResponse, error) {
	requestDef := GenReqDefForUpdateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomLineResponse), nil
	}
}

// UpdateCustomLineInvoker 修改自定义线路
func (c *DnsClient) UpdateCustomLineInvoker(request *model.UpdateCustomLineRequest) *UpdateCustomLineInvoker {
	requestDef := GenReqDefForUpdateCustomLine()
	return &UpdateCustomLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpoint 修改终端节点
//
// 修改终端节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateEndpoint(request *model.UpdateEndpointRequest) (*model.UpdateEndpointResponse, error) {
	requestDef := GenReqDefForUpdateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointResponse), nil
	}
}

// UpdateEndpointInvoker 修改终端节点
func (c *DnsClient) UpdateEndpointInvoker(request *model.UpdateEndpointRequest) *UpdateEndpointInvoker {
	requestDef := GenReqDefForUpdateEndpoint()
	return &UpdateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLineGroups 修改线路分组
//
// 修改线路分组。该接口部分区域未上线，如需使用请提交工单申请开通。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateLineGroups(request *model.UpdateLineGroupsRequest) (*model.UpdateLineGroupsResponse, error) {
	requestDef := GenReqDefForUpdateLineGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLineGroupsResponse), nil
	}
}

// UpdateLineGroupsInvoker 修改线路分组
func (c *DnsClient) UpdateLineGroupsInvoker(request *model.UpdateLineGroupsRequest) *UpdateLineGroupsInvoker {
	requestDef := GenReqDefForUpdateLineGroups()
	return &UpdateLineGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivateZone 修改内网域名
//
// 当您的内网域名创建成功后，您可以通过调用此接口修改内网域名的基本信息，包括TTL、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePrivateZone(request *model.UpdatePrivateZoneRequest) (*model.UpdatePrivateZoneResponse, error) {
	requestDef := GenReqDefForUpdatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateZoneResponse), nil
	}
}

// UpdatePrivateZoneInvoker 修改内网域名
func (c *DnsClient) UpdatePrivateZoneInvoker(request *model.UpdatePrivateZoneRequest) *UpdatePrivateZoneInvoker {
	requestDef := GenReqDefForUpdatePrivateZone()
	return &UpdatePrivateZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrivateZoneStatus 设置内网域名状态
//
// 当您的内网域名创建成功后，您可以通过调用此接口设置内网域名的状态，包括暂停、启用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePrivateZoneStatus(request *model.UpdatePrivateZoneStatusRequest) (*model.UpdatePrivateZoneStatusResponse, error) {
	requestDef := GenReqDefForUpdatePrivateZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateZoneStatusResponse), nil
	}
}

// UpdatePrivateZoneStatusInvoker 设置内网域名状态
func (c *DnsClient) UpdatePrivateZoneStatusInvoker(request *model.UpdatePrivateZoneStatusRequest) *UpdatePrivateZoneStatusInvoker {
	requestDef := GenReqDefForUpdatePrivateZoneStatus()
	return &UpdatePrivateZoneStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicZone 修改公网域名
//
// 当您的公网域名创建成功后，您可以通过调用此接口修改公网域名的基本信息，包括TTL、描述等。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePublicZone(request *model.UpdatePublicZoneRequest) (*model.UpdatePublicZoneResponse, error) {
	requestDef := GenReqDefForUpdatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneResponse), nil
	}
}

// UpdatePublicZoneInvoker 修改公网域名
func (c *DnsClient) UpdatePublicZoneInvoker(request *model.UpdatePublicZoneRequest) *UpdatePublicZoneInvoker {
	requestDef := GenReqDefForUpdatePublicZone()
	return &UpdatePublicZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePublicZoneStatus 设置公网域名状态
//
// 当您的公网域名创建成功后，您可以通过调用此接口设置公网域名的状态，包括暂停、启用。
//
// **[公网域名为全局资源，请选择“华北-北京四（cn-north-4）”区域调用。](tag:hws)**
// **[公网域名为全局资源，请选择“亚太-新加坡（ap-southeast-3）”区域调用。](tag:hws_hk)**
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePublicZoneStatus(request *model.UpdatePublicZoneStatusRequest) (*model.UpdatePublicZoneStatusResponse, error) {
	requestDef := GenReqDefForUpdatePublicZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneStatusResponse), nil
	}
}

// UpdatePublicZoneStatusInvoker 设置公网域名状态
func (c *DnsClient) UpdatePublicZoneStatusInvoker(request *model.UpdatePublicZoneStatusRequest) *UpdatePublicZoneStatusInvoker {
	requestDef := GenReqDefForUpdatePublicZoneStatus()
	return &UpdatePublicZoneStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResolverRule 修改解析器转发规则
//
// 修改解析器转发规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateResolverRule(request *model.UpdateResolverRuleRequest) (*model.UpdateResolverRuleResponse, error) {
	requestDef := GenReqDefForUpdateResolverRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResolverRuleResponse), nil
	}
}

// UpdateResolverRuleInvoker 修改解析器转发规则
func (c *DnsClient) UpdateResolverRuleInvoker(request *model.UpdateResolverRuleRequest) *UpdateResolverRuleInvoker {
	requestDef := GenReqDefForUpdateResolverRule()
	return &UpdateResolverRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableDnssecConfig 关闭DNSSEC
//
// 关闭公网域名的DNSSEC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DisableDnssecConfig(request *model.DisableDnssecConfigRequest) (*model.DisableDnssecConfigResponse, error) {
	requestDef := GenReqDefForDisableDnssecConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDnssecConfigResponse), nil
	}
}

// DisableDnssecConfigInvoker 关闭DNSSEC
func (c *DnsClient) DisableDnssecConfigInvoker(request *model.DisableDnssecConfigRequest) *DisableDnssecConfigInvoker {
	requestDef := GenReqDefForDisableDnssecConfig()
	return &DisableDnssecConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDnssecConfig 开启DNSSEC
//
// 开启公网域名的DNSSEC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) EnableDnssecConfig(request *model.EnableDnssecConfigRequest) (*model.EnableDnssecConfigResponse, error) {
	requestDef := GenReqDefForEnableDnssecConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDnssecConfigResponse), nil
	}
}

// EnableDnssecConfigInvoker 开启DNSSEC
func (c *DnsClient) EnableDnssecConfigInvoker(request *model.EnableDnssecConfigRequest) *EnableDnssecConfigInvoker {
	requestDef := GenReqDefForEnableDnssecConfig()
	return &EnableDnssecConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDnssecConfig 查询DNSSEC
//
// 查询公网域名的DNSSEC。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowDnssecConfig(request *model.ShowDnssecConfigRequest) (*model.ShowDnssecConfigResponse, error) {
	requestDef := GenReqDefForShowDnssecConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDnssecConfigResponse), nil
	}
}

// ShowDnssecConfigInvoker 查询DNSSEC
func (c *DnsClient) ShowDnssecConfigInvoker(request *model.ShowDnssecConfigRequest) *ShowDnssecConfigInvoker {
	requestDef := GenReqDefForShowDnssecConfig()
	return &ShowDnssecConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEipRecordSet 设置弹性公网IP的反向解析记录
//
// 设置弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateEipRecordSet(request *model.CreateEipRecordSetRequest) (*model.CreateEipRecordSetResponse, error) {
	requestDef := GenReqDefForCreateEipRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEipRecordSetResponse), nil
	}
}

// CreateEipRecordSetInvoker 设置弹性公网IP的反向解析记录
func (c *DnsClient) CreateEipRecordSetInvoker(request *model.CreateEipRecordSetRequest) *CreateEipRecordSetInvoker {
	requestDef := GenReqDefForCreateEipRecordSet()
	return &CreateEipRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSet 创建记录集
//
// 记录集是指一组资源记录的集合，这些资源记录属于同一域名，用于定义域名支持的解析类型以及解析值。您的域名创建完成后，可以通过调用此接口为域名添加不同类型的记录集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSet(request *model.CreateRecordSetRequest) (*model.CreateRecordSetResponse, error) {
	requestDef := GenReqDefForCreateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetResponse), nil
	}
}

// CreateRecordSetInvoker 创建记录集
func (c *DnsClient) CreateRecordSetInvoker(request *model.CreateRecordSetRequest) *CreateRecordSetInvoker {
	requestDef := GenReqDefForCreateRecordSet()
	return &CreateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordSet 删除记录集
//
// 当您的记录集不再使用时，您可以通过调用此接口将其删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteRecordSet(request *model.DeleteRecordSetRequest) (*model.DeleteRecordSetResponse, error) {
	requestDef := GenReqDefForDeleteRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetResponse), nil
	}
}

// DeleteRecordSetInvoker 删除记录集
func (c *DnsClient) DeleteRecordSetInvoker(request *model.DeleteRecordSetRequest) *DeleteRecordSetInvoker {
	requestDef := GenReqDefForDeleteRecordSet()
	return &DeleteRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPtrRecords 查询弹性公网IP的反向解析记录列表
//
// 查询弹性公网IP的反向解析记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPtrRecords(request *model.ListPtrRecordsRequest) (*model.ListPtrRecordsResponse, error) {
	requestDef := GenReqDefForListPtrRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPtrRecordsResponse), nil
	}
}

// ListPtrRecordsInvoker 查询弹性公网IP的反向解析记录列表
func (c *DnsClient) ListPtrRecordsInvoker(request *model.ListPtrRecordsRequest) *ListPtrRecordsInvoker {
	requestDef := GenReqDefForListPtrRecords()
	return &ListPtrRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSets 查询租户记录集列表
//
// 当您的记录集创建成功后，您可以通过调用此接口查询指定域名下的所有记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSets(request *model.ListRecordSetsRequest) (*model.ListRecordSetsResponse, error) {
	requestDef := GenReqDefForListRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsResponse), nil
	}
}

// ListRecordSetsInvoker 查询租户记录集列表
func (c *DnsClient) ListRecordSetsInvoker(request *model.ListRecordSetsRequest) *ListRecordSetsInvoker {
	requestDef := GenReqDefForListRecordSets()
	return &ListRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSetsByZone 查询域名下的记录集列表
//
// 当您的记录集创建成功后，您可以通过调用此接口查询指定域名下的所有记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSetsByZone(request *model.ListRecordSetsByZoneRequest) (*model.ListRecordSetsByZoneResponse, error) {
	requestDef := GenReqDefForListRecordSetsByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsByZoneResponse), nil
	}
}

// ListRecordSetsByZoneInvoker 查询域名下的记录集列表
func (c *DnsClient) ListRecordSetsByZoneInvoker(request *model.ListRecordSetsByZoneRequest) *ListRecordSetsByZoneInvoker {
	requestDef := GenReqDefForListRecordSetsByZone()
	return &ListRecordSetsByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestorePtrRecord 将弹性公网IP的反向解析记录恢复为默认值
//
// 将弹性公网IP的反向解析记录恢复为默认值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) RestorePtrRecord(request *model.RestorePtrRecordRequest) (*model.RestorePtrRecordResponse, error) {
	requestDef := GenReqDefForRestorePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestorePtrRecordResponse), nil
	}
}

// RestorePtrRecordInvoker 将弹性公网IP的反向解析记录恢复为默认值
func (c *DnsClient) RestorePtrRecordInvoker(request *model.RestorePtrRecordRequest) *RestorePtrRecordInvoker {
	requestDef := GenReqDefForRestorePtrRecord()
	return &RestorePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPtrRecordSet 查询弹性公网IP的反向解析记录
//
// 查询弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPtrRecordSet(request *model.ShowPtrRecordSetRequest) (*model.ShowPtrRecordSetResponse, error) {
	requestDef := GenReqDefForShowPtrRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPtrRecordSetResponse), nil
	}
}

// ShowPtrRecordSetInvoker 查询弹性公网IP的反向解析记录
func (c *DnsClient) ShowPtrRecordSetInvoker(request *model.ShowPtrRecordSetRequest) *ShowPtrRecordSetInvoker {
	requestDef := GenReqDefForShowPtrRecordSet()
	return &ShowPtrRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSet 查询记录集
//
// 当您的记录集创建成功后，您可以通过调用此接口查询指定域名下的所有记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSet(request *model.ShowRecordSetRequest) (*model.ShowRecordSetResponse, error) {
	requestDef := GenReqDefForShowRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetResponse), nil
	}
}

// ShowRecordSetInvoker 查询记录集
func (c *DnsClient) ShowRecordSetInvoker(request *model.ShowRecordSetRequest) *ShowRecordSetInvoker {
	requestDef := GenReqDefForShowRecordSet()
	return &ShowRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePtrRecord 修改弹性公网IP的反向解析记录
//
// 修改弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePtrRecord(request *model.UpdatePtrRecordRequest) (*model.UpdatePtrRecordResponse, error) {
	requestDef := GenReqDefForUpdatePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePtrRecordResponse), nil
	}
}

// UpdatePtrRecordInvoker 修改弹性公网IP的反向解析记录
func (c *DnsClient) UpdatePtrRecordInvoker(request *model.UpdatePtrRecordRequest) *UpdatePtrRecordInvoker {
	requestDef := GenReqDefForUpdatePtrRecord()
	return &UpdatePtrRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordSet 修改记录集
//
// 当您的记录集创建成功后，您可以通过调用此接口修改记录集的信息，包括域名、类型、记录值、TTL、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateRecordSet(request *model.UpdateRecordSetRequest) (*model.UpdateRecordSetResponse, error) {
	requestDef := GenReqDefForUpdateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetResponse), nil
	}
}

// UpdateRecordSetInvoker 修改记录集
func (c *DnsClient) UpdateRecordSetInvoker(request *model.UpdateRecordSetRequest) *UpdateRecordSetInvoker {
	requestDef := GenReqDefForUpdateRecordSet()
	return &UpdateRecordSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePtr 创建弹性公网IP的反向解析记录
//
// 创建弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreatePtr(request *model.CreatePtrRequest) (*model.CreatePtrResponse, error) {
	requestDef := GenReqDefForCreatePtr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePtrResponse), nil
	}
}

// CreatePtrInvoker 创建弹性公网IP的反向解析记录
func (c *DnsClient) CreatePtrInvoker(request *model.CreatePtrRequest) *CreatePtrInvoker {
	requestDef := GenReqDefForCreatePtr()
	return &CreatePtrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRecordSetWithLine 创建记录集
//
// 记录集是指一组资源记录的集合，这些资源记录属于同一域名，用于定义域名支持的解析类型以及解析值。您的域名创建完成后，可以通过调用此接口为域名添加不同类型的记录集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) CreateRecordSetWithLine(request *model.CreateRecordSetWithLineRequest) (*model.CreateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithLineResponse), nil
	}
}

// CreateRecordSetWithLineInvoker 创建记录集
func (c *DnsClient) CreateRecordSetWithLineInvoker(request *model.CreateRecordSetWithLineRequest) *CreateRecordSetWithLineInvoker {
	requestDef := GenReqDefForCreateRecordSetWithLine()
	return &CreateRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePtr 将弹性公网IP的反向解析记录恢复为默认值
//
// 将弹性公网IP的反向解析记录恢复为默认值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeletePtr(request *model.DeletePtrRequest) (*model.DeletePtrResponse, error) {
	requestDef := GenReqDefForDeletePtr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePtrResponse), nil
	}
}

// DeletePtrInvoker 将弹性公网IP的反向解析记录恢复为默认值
func (c *DnsClient) DeletePtrInvoker(request *model.DeletePtrRequest) *DeletePtrInvoker {
	requestDef := GenReqDefForDeletePtr()
	return &DeletePtrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRecordSets 删除记录集
//
// 当您的记录集不再使用时，您可以通过调用此接口将其删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) DeleteRecordSets(request *model.DeleteRecordSetsRequest) (*model.DeleteRecordSetsResponse, error) {
	requestDef := GenReqDefForDeleteRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetsResponse), nil
	}
}

// DeleteRecordSetsInvoker 删除记录集
func (c *DnsClient) DeleteRecordSetsInvoker(request *model.DeleteRecordSetsRequest) *DeleteRecordSetsInvoker {
	requestDef := GenReqDefForDeleteRecordSets()
	return &DeleteRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPtrs 查询弹性公网IP的反向解析记录列表
//
// 查询弹性公网IP的反向解析记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPtrs(request *model.ListPtrsRequest) (*model.ListPtrsResponse, error) {
	requestDef := GenReqDefForListPtrs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPtrsResponse), nil
	}
}

// ListPtrsInvoker 查询弹性公网IP的反向解析记录列表
func (c *DnsClient) ListPtrsInvoker(request *model.ListPtrsRequest) *ListPtrsInvoker {
	requestDef := GenReqDefForListPtrs()
	return &ListPtrsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicZoneLines 查询公网域名的线路列表
//
// 公网域名支持设置线路解析，当您的公网域名创建完成并添加记录集时，可通过调用此接口查询公网域名的所有解析线路。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListPublicZoneLines(request *model.ListPublicZoneLinesRequest) (*model.ListPublicZoneLinesResponse, error) {
	requestDef := GenReqDefForListPublicZoneLines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicZoneLinesResponse), nil
	}
}

// ListPublicZoneLinesInvoker 查询公网域名的线路列表
func (c *DnsClient) ListPublicZoneLinesInvoker(request *model.ListPublicZoneLinesRequest) *ListPublicZoneLinesInvoker {
	requestDef := GenReqDefForListPublicZoneLines()
	return &ListPublicZoneLinesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecordSetsWithLine 查询租户记录集列表
//
// 当您的记录集创建成功后，您可以通过调用此接口查询单个记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ListRecordSetsWithLine(request *model.ListRecordSetsWithLineRequest) (*model.ListRecordSetsWithLineResponse, error) {
	requestDef := GenReqDefForListRecordSetsWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsWithLineResponse), nil
	}
}

// ListRecordSetsWithLineInvoker 查询租户记录集列表
func (c *DnsClient) ListRecordSetsWithLineInvoker(request *model.ListRecordSetsWithLineRequest) *ListRecordSetsWithLineInvoker {
	requestDef := GenReqDefForListRecordSetsWithLine()
	return &ListRecordSetsWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRecordSetsStatus 设置记录集状态
//
// 当您的内网域名创建成功后，您可以通过调用此接口设置记录集的状态，包括暂停、启用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) SetRecordSetsStatus(request *model.SetRecordSetsStatusRequest) (*model.SetRecordSetsStatusResponse, error) {
	requestDef := GenReqDefForSetRecordSetsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecordSetsStatusResponse), nil
	}
}

// SetRecordSetsStatusInvoker 设置记录集状态
func (c *DnsClient) SetRecordSetsStatusInvoker(request *model.SetRecordSetsStatusRequest) *SetRecordSetsStatusInvoker {
	requestDef := GenReqDefForSetRecordSetsStatus()
	return &SetRecordSetsStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPtr 查询弹性公网IP的反向解析记录
//
// 查询弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowPtr(request *model.ShowPtrRequest) (*model.ShowPtrResponse, error) {
	requestDef := GenReqDefForShowPtr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPtrResponse), nil
	}
}

// ShowPtrInvoker 查询弹性公网IP的反向解析记录
func (c *DnsClient) ShowPtrInvoker(request *model.ShowPtrRequest) *ShowPtrInvoker {
	requestDef := GenReqDefForShowPtr()
	return &ShowPtrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSetByZone 查询域名下的记录集列表
//
// 当您的记录集创建成功后，您可以通过调用此接口查询单个记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSetByZone(request *model.ShowRecordSetByZoneRequest) (*model.ShowRecordSetByZoneResponse, error) {
	requestDef := GenReqDefForShowRecordSetByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetByZoneResponse), nil
	}
}

// ShowRecordSetByZoneInvoker 查询域名下的记录集列表
func (c *DnsClient) ShowRecordSetByZoneInvoker(request *model.ShowRecordSetByZoneRequest) *ShowRecordSetByZoneInvoker {
	requestDef := GenReqDefForShowRecordSetByZone()
	return &ShowRecordSetByZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecordSetWithLine 查询记录集
//
// 当您的记录集创建成功后，您可以通过调用此接口查询单个记录集信息，包括名称、ID、状态、所属域名、解析记录值、标签、TTL、创建时间、修改时间、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) ShowRecordSetWithLine(request *model.ShowRecordSetWithLineRequest) (*model.ShowRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForShowRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetWithLineResponse), nil
	}
}

// ShowRecordSetWithLineInvoker 查询记录集
func (c *DnsClient) ShowRecordSetWithLineInvoker(request *model.ShowRecordSetWithLineRequest) *ShowRecordSetWithLineInvoker {
	requestDef := GenReqDefForShowRecordSetWithLine()
	return &ShowRecordSetWithLineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePtr 修改弹性公网IP的反向解析记录
//
// 修改弹性公网IP的反向解析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdatePtr(request *model.UpdatePtrRequest) (*model.UpdatePtrResponse, error) {
	requestDef := GenReqDefForUpdatePtr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePtrResponse), nil
	}
}

// UpdatePtrInvoker 修改弹性公网IP的反向解析记录
func (c *DnsClient) UpdatePtrInvoker(request *model.UpdatePtrRequest) *UpdatePtrInvoker {
	requestDef := GenReqDefForUpdatePtr()
	return &UpdatePtrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRecordSets 修改记录集
//
// 当您的记录集创建成功后，您可以通过调用此接口修改记录集的信息，包括域名、类型、记录值、TTL、描述等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DnsClient) UpdateRecordSets(request *model.UpdateRecordSetsRequest) (*model.UpdateRecordSetsResponse, error) {
	requestDef := GenReqDefForUpdateRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetsResponse), nil
	}
}

// UpdateRecordSetsInvoker 修改记录集
func (c *DnsClient) UpdateRecordSetsInvoker(request *model.UpdateRecordSetsRequest) *UpdateRecordSetsInvoker {
	requestDef := GenReqDefForUpdateRecordSets()
	return &UpdateRecordSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
