package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

type VpcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewVpcClient(hcClient *httpclient.HcHttpClient) *VpcClient {
	return &VpcClient{HcClient: hcClient}
}

func VpcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AcceptVpcPeering 接受对等连接请求
//
// 租户A名下的VPC申请和租户B的VPC建立对等连接，需要等待租户B接受该请求。此接口用于租户接受其他租户发起的对等连接请求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) AcceptVpcPeering(request *model.AcceptVpcPeeringRequest) (*model.AcceptVpcPeeringResponse, error) {
	requestDef := GenReqDefForAcceptVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptVpcPeeringResponse), nil
	}
}

// AcceptVpcPeeringInvoker 接受对等连接请求
func (c *VpcClient) AcceptVpcPeeringInvoker(request *model.AcceptVpcPeeringRequest) *AcceptVpcPeeringInvoker {
	requestDef := GenReqDefForAcceptVpcPeering()
	return &AcceptVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRouteTable 子网关联路由表
//
// 路由表关联子网。子网关联路由表A后，再关联B，不需要先跟路由表A解关联再关联路由表B
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) AssociateRouteTable(request *model.AssociateRouteTableRequest) (*model.AssociateRouteTableResponse, error) {
	requestDef := GenReqDefForAssociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouteTableResponse), nil
	}
}

// AssociateRouteTableInvoker 子网关联路由表
func (c *VpcClient) AssociateRouteTableInvoker(request *model.AssociateRouteTableRequest) *AssociateRouteTableInvoker {
	requestDef := GenReqDefForAssociateRouteTable()
	return &AssociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSecurityGroupTags 批量创建安全组资源标签
//
// 为指定的安全组资源实例批量添加标签。
// 此接口为幂等接口：创建时如果请求体中存在重复key则报错。创建时，不允许设置重复key数据，如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchCreateSecurityGroupTags(request *model.BatchCreateSecurityGroupTagsRequest) (*model.BatchCreateSecurityGroupTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateSecurityGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSecurityGroupTagsResponse), nil
	}
}

// BatchCreateSecurityGroupTagsInvoker 批量创建安全组资源标签
func (c *VpcClient) BatchCreateSecurityGroupTagsInvoker(request *model.BatchCreateSecurityGroupTagsRequest) *BatchCreateSecurityGroupTagsInvoker {
	requestDef := GenReqDefForBatchCreateSecurityGroupTags()
	return &BatchCreateSecurityGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateSubnetTags 批量创建子网资源标签
//
// 为指定的子网资源实例批量添加标签。
// 此接口为幂等接口：创建时如果请求体中存在重复key则报错。创建时，不允许设置重复key数据，如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchCreateSubnetTags(request *model.BatchCreateSubnetTagsRequest) (*model.BatchCreateSubnetTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateSubnetTagsResponse), nil
	}
}

// BatchCreateSubnetTagsInvoker 批量创建子网资源标签
func (c *VpcClient) BatchCreateSubnetTagsInvoker(request *model.BatchCreateSubnetTagsRequest) *BatchCreateSubnetTagsInvoker {
	requestDef := GenReqDefForBatchCreateSubnetTags()
	return &BatchCreateSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecurityGroupTags 批量删除安全组资源标签
//
// 为指定的安全组资源实例批量删除标签
// 此接口为幂等接口：删除时，如果删除的标签不存在，默认处理成功；删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchDeleteSecurityGroupTags(request *model.BatchDeleteSecurityGroupTagsRequest) (*model.BatchDeleteSecurityGroupTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecurityGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecurityGroupTagsResponse), nil
	}
}

// BatchDeleteSecurityGroupTagsInvoker 批量删除安全组资源标签
func (c *VpcClient) BatchDeleteSecurityGroupTagsInvoker(request *model.BatchDeleteSecurityGroupTagsRequest) *BatchDeleteSecurityGroupTagsInvoker {
	requestDef := GenReqDefForBatchDeleteSecurityGroupTags()
	return &BatchDeleteSecurityGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSubnetTags 批量删除子网资源标签
//
// 为指定的子网资源实例批量删除标签
// 此接口为幂等接口：删除时，如果删除的标签不存在，默认处理成功；删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchDeleteSubnetTags(request *model.BatchDeleteSubnetTagsRequest) (*model.BatchDeleteSubnetTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSubnetTagsResponse), nil
	}
}

// BatchDeleteSubnetTagsInvoker 批量删除子网资源标签
func (c *VpcClient) BatchDeleteSubnetTagsInvoker(request *model.BatchDeleteSubnetTagsRequest) *BatchDeleteSubnetTagsInvoker {
	requestDef := GenReqDefForBatchDeleteSubnetTags()
	return &BatchDeleteSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlowLog 创建流日志
//
// 创建流日志。
// 流日志功能可以记录虚拟私有云中的流量信息，帮助您检查和优化安全组和网络ACL防火墙控制规则、监控网络流量、进行网络攻击分析等。
// VPC流日志功能需要与云日志服务LTS结合使用，请先在云日志服务中创建日志组和日志主题，然后再创建VPC流日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateFlowLog(request *model.CreateFlowLogRequest) (*model.CreateFlowLogResponse, error) {
	requestDef := GenReqDefForCreateFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlowLogResponse), nil
	}
}

// CreateFlowLogInvoker 创建流日志
func (c *VpcClient) CreateFlowLogInvoker(request *model.CreateFlowLogRequest) *CreateFlowLogInvoker {
	requestDef := GenReqDefForCreateFlowLog()
	return &CreateFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePort 创建端口
//
// 创建端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreatePort(request *model.CreatePortRequest) (*model.CreatePortResponse, error) {
	requestDef := GenReqDefForCreatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePortResponse), nil
	}
}

// CreatePortInvoker 创建端口
func (c *VpcClient) CreatePortInvoker(request *model.CreatePortRequest) *CreatePortInvoker {
	requestDef := GenReqDefForCreatePort()
	return &CreatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRouteTable 创建路由表
//
// 创建路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateRouteTable(request *model.CreateRouteTableRequest) (*model.CreateRouteTableResponse, error) {
	requestDef := GenReqDefForCreateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRouteTableResponse), nil
	}
}

// CreateRouteTableInvoker 创建路由表
func (c *VpcClient) CreateRouteTableInvoker(request *model.CreateRouteTableRequest) *CreateRouteTableInvoker {
	requestDef := GenReqDefForCreateRouteTable()
	return &CreateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroup 创建安全组
//
// 创建安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateSecurityGroup(request *model.CreateSecurityGroupRequest) (*model.CreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupResponse), nil
	}
}

// CreateSecurityGroupInvoker 创建安全组
func (c *VpcClient) CreateSecurityGroupInvoker(request *model.CreateSecurityGroupRequest) *CreateSecurityGroupInvoker {
	requestDef := GenReqDefForCreateSecurityGroup()
	return &CreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroupRule 创建安全组规则
//
// 创建安全组规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateSecurityGroupRule(request *model.CreateSecurityGroupRuleRequest) (*model.CreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

// CreateSecurityGroupRuleInvoker 创建安全组规则
func (c *VpcClient) CreateSecurityGroupRuleInvoker(request *model.CreateSecurityGroupRuleRequest) *CreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForCreateSecurityGroupRule()
	return &CreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityGroupTag 创建安全组资源标签
//
// 给指定安全组资源实例增加标签信息。
// 此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateSecurityGroupTag(request *model.CreateSecurityGroupTagRequest) (*model.CreateSecurityGroupTagResponse, error) {
	requestDef := GenReqDefForCreateSecurityGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityGroupTagResponse), nil
	}
}

// CreateSecurityGroupTagInvoker 创建安全组资源标签
func (c *VpcClient) CreateSecurityGroupTagInvoker(request *model.CreateSecurityGroupTagRequest) *CreateSecurityGroupTagInvoker {
	requestDef := GenReqDefForCreateSecurityGroupTag()
	return &CreateSecurityGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubnet 创建子网
//
// 创建子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateSubnet(request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	requestDef := GenReqDefForCreateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetResponse), nil
	}
}

// CreateSubnetInvoker 创建子网
func (c *VpcClient) CreateSubnetInvoker(request *model.CreateSubnetRequest) *CreateSubnetInvoker {
	requestDef := GenReqDefForCreateSubnet()
	return &CreateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubnetTag 创建子网资源标签
//
// 给指定子网资源实例增加标签信息。
// 此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateSubnetTag(request *model.CreateSubnetTagRequest) (*model.CreateSubnetTagResponse, error) {
	requestDef := GenReqDefForCreateSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubnetTagResponse), nil
	}
}

// CreateSubnetTagInvoker 创建子网资源标签
func (c *VpcClient) CreateSubnetTagInvoker(request *model.CreateSubnetTagRequest) *CreateSubnetTagInvoker {
	requestDef := GenReqDefForCreateSubnetTag()
	return &CreateSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcPeering 创建对等连接
//
// 创建对等连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateVpcPeering(request *model.CreateVpcPeeringRequest) (*model.CreateVpcPeeringResponse, error) {
	requestDef := GenReqDefForCreateVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcPeeringResponse), nil
	}
}

// CreateVpcPeeringInvoker 创建对等连接
func (c *VpcClient) CreateVpcPeeringInvoker(request *model.CreateVpcPeeringRequest) *CreateVpcPeeringInvoker {
	requestDef := GenReqDefForCreateVpcPeering()
	return &CreateVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlowLog 删除流日志
//
// 删除流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteFlowLog(request *model.DeleteFlowLogRequest) (*model.DeleteFlowLogResponse, error) {
	requestDef := GenReqDefForDeleteFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlowLogResponse), nil
	}
}

// DeleteFlowLogInvoker 删除流日志
func (c *VpcClient) DeleteFlowLogInvoker(request *model.DeleteFlowLogRequest) *DeleteFlowLogInvoker {
	requestDef := GenReqDefForDeleteFlowLog()
	return &DeleteFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePort 删除端口
//
// 删除端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeletePort(request *model.DeletePortRequest) (*model.DeletePortResponse, error) {
	requestDef := GenReqDefForDeletePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePortResponse), nil
	}
}

// DeletePortInvoker 删除端口
func (c *VpcClient) DeletePortInvoker(request *model.DeletePortRequest) *DeletePortInvoker {
	requestDef := GenReqDefForDeletePort()
	return &DeletePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRouteTable 删除路由表
//
// 删除路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteRouteTable(request *model.DeleteRouteTableRequest) (*model.DeleteRouteTableResponse, error) {
	requestDef := GenReqDefForDeleteRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRouteTableResponse), nil
	}
}

// DeleteRouteTableInvoker 删除路由表
func (c *VpcClient) DeleteRouteTableInvoker(request *model.DeleteRouteTableRequest) *DeleteRouteTableInvoker {
	requestDef := GenReqDefForDeleteRouteTable()
	return &DeleteRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroup 删除安全组
//
// 删除安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteSecurityGroup(request *model.DeleteSecurityGroupRequest) (*model.DeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupResponse), nil
	}
}

// DeleteSecurityGroupInvoker 删除安全组
func (c *VpcClient) DeleteSecurityGroupInvoker(request *model.DeleteSecurityGroupRequest) *DeleteSecurityGroupInvoker {
	requestDef := GenReqDefForDeleteSecurityGroup()
	return &DeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroupRule 删除安全组规则
//
// 删除安全组规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteSecurityGroupRule(request *model.DeleteSecurityGroupRuleRequest) (*model.DeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

// DeleteSecurityGroupRuleInvoker 删除安全组规则
func (c *VpcClient) DeleteSecurityGroupRuleInvoker(request *model.DeleteSecurityGroupRuleRequest) *DeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupRule()
	return &DeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityGroupTag 删除安全组资源标签
//
// 删除指定安全组资源实例的标签信息。
// 该接口为幂等接口：删除的key不存在报404，Key不能为空或者空字符串
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteSecurityGroupTag(request *model.DeleteSecurityGroupTagRequest) (*model.DeleteSecurityGroupTagResponse, error) {
	requestDef := GenReqDefForDeleteSecurityGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityGroupTagResponse), nil
	}
}

// DeleteSecurityGroupTagInvoker 删除安全组资源标签
func (c *VpcClient) DeleteSecurityGroupTagInvoker(request *model.DeleteSecurityGroupTagRequest) *DeleteSecurityGroupTagInvoker {
	requestDef := GenReqDefForDeleteSecurityGroupTag()
	return &DeleteSecurityGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubnet 删除子网
//
// 删除子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteSubnet(request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	requestDef := GenReqDefForDeleteSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetResponse), nil
	}
}

// DeleteSubnetInvoker 删除子网
func (c *VpcClient) DeleteSubnetInvoker(request *model.DeleteSubnetRequest) *DeleteSubnetInvoker {
	requestDef := GenReqDefForDeleteSubnet()
	return &DeleteSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubnetTag 删除子网资源标签
//
// 删除指定子网资源实例的标签信息。
// 该接口为幂等接口：删除的key不存在报404，Key不能为空或者空字符串
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteSubnetTag(request *model.DeleteSubnetTagRequest) (*model.DeleteSubnetTagResponse, error) {
	requestDef := GenReqDefForDeleteSubnetTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubnetTagResponse), nil
	}
}

// DeleteSubnetTagInvoker 删除子网资源标签
func (c *VpcClient) DeleteSubnetTagInvoker(request *model.DeleteSubnetTagRequest) *DeleteSubnetTagInvoker {
	requestDef := GenReqDefForDeleteSubnetTag()
	return &DeleteSubnetTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcPeering 删除对等连接
//
// 删除对等连接。
// 可以在在本端或对端任何一端删除对等连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteVpcPeering(request *model.DeleteVpcPeeringRequest) (*model.DeleteVpcPeeringResponse, error) {
	requestDef := GenReqDefForDeleteVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcPeeringResponse), nil
	}
}

// DeleteVpcPeeringInvoker 删除对等连接
func (c *VpcClient) DeleteVpcPeeringInvoker(request *model.DeleteVpcPeeringRequest) *DeleteVpcPeeringInvoker {
	requestDef := GenReqDefForDeleteVpcPeering()
	return &DeleteVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateRouteTable 子网解关联路由表
//
// 子网解关联路由表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DisassociateRouteTable(request *model.DisassociateRouteTableRequest) (*model.DisassociateRouteTableResponse, error) {
	requestDef := GenReqDefForDisassociateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouteTableResponse), nil
	}
}

// DisassociateRouteTableInvoker 子网解关联路由表
func (c *VpcClient) DisassociateRouteTableInvoker(request *model.DisassociateRouteTableRequest) *DisassociateRouteTableInvoker {
	requestDef := GenReqDefForDisassociateRouteTable()
	return &DisassociateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowLogs 查询流日志列表
//
// 查询提交请求的租户的所有流日志列表，并根据过滤条件进行过滤
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListFlowLogs(request *model.ListFlowLogsRequest) (*model.ListFlowLogsResponse, error) {
	requestDef := GenReqDefForListFlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowLogsResponse), nil
	}
}

// ListFlowLogsInvoker 查询流日志列表
func (c *VpcClient) ListFlowLogsInvoker(request *model.ListFlowLogsRequest) *ListFlowLogsInvoker {
	requestDef := GenReqDefForListFlowLogs()
	return &ListFlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPorts 查询端口列表
//
// 查询提交请求的租户的所有端口，单次查询最多返回2000条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListPorts(request *model.ListPortsRequest) (*model.ListPortsResponse, error) {
	requestDef := GenReqDefForListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPortsResponse), nil
	}
}

// ListPortsInvoker 查询端口列表
func (c *VpcClient) ListPortsInvoker(request *model.ListPortsRequest) *ListPortsInvoker {
	requestDef := GenReqDefForListPorts()
	return &ListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRouteTables 查询路由表列表
//
// 查询提交请求的帐户的所有路由表列表，并根据过滤条件进行过滤
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListRouteTables(request *model.ListRouteTablesRequest) (*model.ListRouteTablesResponse, error) {
	requestDef := GenReqDefForListRouteTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRouteTablesResponse), nil
	}
}

// ListRouteTablesInvoker 查询路由表列表
func (c *VpcClient) ListRouteTablesInvoker(request *model.ListRouteTablesRequest) *ListRouteTablesInvoker {
	requestDef := GenReqDefForListRouteTables()
	return &ListRouteTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroupRules 查询安全组规则列表
//
// 查询安全组规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSecurityGroupRules(request *model.ListSecurityGroupRulesRequest) (*model.ListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupRulesResponse), nil
	}
}

// ListSecurityGroupRulesInvoker 查询安全组规则列表
func (c *VpcClient) ListSecurityGroupRulesInvoker(request *model.ListSecurityGroupRulesRequest) *ListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForListSecurityGroupRules()
	return &ListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroupTags 查询安全组项目标签
//
// 查询租户在指定区域和实例类型的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSecurityGroupTags(request *model.ListSecurityGroupTagsRequest) (*model.ListSecurityGroupTagsResponse, error) {
	requestDef := GenReqDefForListSecurityGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupTagsResponse), nil
	}
}

// ListSecurityGroupTagsInvoker 查询安全组项目标签
func (c *VpcClient) ListSecurityGroupTagsInvoker(request *model.ListSecurityGroupTagsRequest) *ListSecurityGroupTagsInvoker {
	requestDef := GenReqDefForListSecurityGroupTags()
	return &ListSecurityGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroups 查询安全组列表
//
// 查询安全组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSecurityGroups(request *model.ListSecurityGroupsRequest) (*model.ListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsResponse), nil
	}
}

// ListSecurityGroupsInvoker 查询安全组列表
func (c *VpcClient) ListSecurityGroupsInvoker(request *model.ListSecurityGroupsRequest) *ListSecurityGroupsInvoker {
	requestDef := GenReqDefForListSecurityGroups()
	return &ListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityGroupsByTags 查询安全组资源实例
//
// 使用标签过滤实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSecurityGroupsByTags(request *model.ListSecurityGroupsByTagsRequest) (*model.ListSecurityGroupsByTagsResponse, error) {
	requestDef := GenReqDefForListSecurityGroupsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityGroupsByTagsResponse), nil
	}
}

// ListSecurityGroupsByTagsInvoker 查询安全组资源实例
func (c *VpcClient) ListSecurityGroupsByTagsInvoker(request *model.ListSecurityGroupsByTagsRequest) *ListSecurityGroupsByTagsInvoker {
	requestDef := GenReqDefForListSecurityGroupsByTags()
	return &ListSecurityGroupsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubnetTags 查询子网项目标签
//
// 查询租户在指定区域和实例类型的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSubnetTags(request *model.ListSubnetTagsRequest) (*model.ListSubnetTagsResponse, error) {
	requestDef := GenReqDefForListSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetTagsResponse), nil
	}
}

// ListSubnetTagsInvoker 查询子网项目标签
func (c *VpcClient) ListSubnetTagsInvoker(request *model.ListSubnetTagsRequest) *ListSubnetTagsInvoker {
	requestDef := GenReqDefForListSubnetTags()
	return &ListSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubnets 查询子网列表
//
// 查询子网列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSubnets(request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	requestDef := GenReqDefForListSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsResponse), nil
	}
}

// ListSubnetsInvoker 查询子网列表
func (c *VpcClient) ListSubnetsInvoker(request *model.ListSubnetsRequest) *ListSubnetsInvoker {
	requestDef := GenReqDefForListSubnets()
	return &ListSubnetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubnetsByTags 查询子网资源实例
//
// 使用标签过滤实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListSubnetsByTags(request *model.ListSubnetsByTagsRequest) (*model.ListSubnetsByTagsResponse, error) {
	requestDef := GenReqDefForListSubnetsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubnetsByTagsResponse), nil
	}
}

// ListSubnetsByTagsInvoker 查询子网资源实例
func (c *VpcClient) ListSubnetsByTagsInvoker(request *model.ListSubnetsByTagsRequest) *ListSubnetsByTagsInvoker {
	requestDef := GenReqDefForListSubnetsByTags()
	return &ListSubnetsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcPeerings 查询对等连接列表
//
// 查询提交请求的租户的所有对等连接。根据过滤条件进行过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListVpcPeerings(request *model.ListVpcPeeringsRequest) (*model.ListVpcPeeringsResponse, error) {
	requestDef := GenReqDefForListVpcPeerings()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcPeeringsResponse), nil
	}
}

// ListVpcPeeringsInvoker 查询对等连接列表
func (c *VpcClient) ListVpcPeeringsInvoker(request *model.ListVpcPeeringsRequest) *ListVpcPeeringsInvoker {
	requestDef := GenReqDefForListVpcPeerings()
	return &ListVpcPeeringsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RejectVpcPeering 拒绝对等连接请求
//
// 租户A名下的VPC申请和租户B的VPC建立对等连接，需要等待租户B接受该请求。此接口用于租户拒绝其他租户发起的对等连接请求。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) RejectVpcPeering(request *model.RejectVpcPeeringRequest) (*model.RejectVpcPeeringResponse, error) {
	requestDef := GenReqDefForRejectVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RejectVpcPeeringResponse), nil
	}
}

// RejectVpcPeeringInvoker 拒绝对等连接请求
func (c *VpcClient) RejectVpcPeeringInvoker(request *model.RejectVpcPeeringRequest) *RejectVpcPeeringInvoker {
	requestDef := GenReqDefForRejectVpcPeering()
	return &RejectVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlowLog 查询流日志
//
// 查询流日志详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowFlowLog(request *model.ShowFlowLogRequest) (*model.ShowFlowLogResponse, error) {
	requestDef := GenReqDefForShowFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlowLogResponse), nil
	}
}

// ShowFlowLogInvoker 查询流日志
func (c *VpcClient) ShowFlowLogInvoker(request *model.ShowFlowLogRequest) *ShowFlowLogInvoker {
	requestDef := GenReqDefForShowFlowLog()
	return &ShowFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPort 查询端口
//
// 查询单个端口详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowPort(request *model.ShowPortRequest) (*model.ShowPortResponse, error) {
	requestDef := GenReqDefForShowPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPortResponse), nil
	}
}

// ShowPortInvoker 查询端口
func (c *VpcClient) ShowPortInvoker(request *model.ShowPortRequest) *ShowPortInvoker {
	requestDef := GenReqDefForShowPort()
	return &ShowPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 查询配额
//
// 查询单租户在VPC服务下的网络资源配额，包括vpc配额、子网配额、安全组配额、安全组规则配额、弹性公网IP配额，vpn配额等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 查询配额
func (c *VpcClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRouteTable 查询路由表
//
// 查询路由表详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowRouteTable(request *model.ShowRouteTableRequest) (*model.ShowRouteTableResponse, error) {
	requestDef := GenReqDefForShowRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRouteTableResponse), nil
	}
}

// ShowRouteTableInvoker 查询路由表
func (c *VpcClient) ShowRouteTableInvoker(request *model.ShowRouteTableRequest) *ShowRouteTableInvoker {
	requestDef := GenReqDefForShowRouteTable()
	return &ShowRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityGroup 查询安全组
//
// 查询单个安全组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowSecurityGroup(request *model.ShowSecurityGroupRequest) (*model.ShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupResponse), nil
	}
}

// ShowSecurityGroupInvoker 查询安全组
func (c *VpcClient) ShowSecurityGroupInvoker(request *model.ShowSecurityGroupRequest) *ShowSecurityGroupInvoker {
	requestDef := GenReqDefForShowSecurityGroup()
	return &ShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityGroupRule 查询安全组规则
//
// 查询单个安全组规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowSecurityGroupRule(request *model.ShowSecurityGroupRuleRequest) (*model.ShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

// ShowSecurityGroupRuleInvoker 查询安全组规则
func (c *VpcClient) ShowSecurityGroupRuleInvoker(request *model.ShowSecurityGroupRuleRequest) *ShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForShowSecurityGroupRule()
	return &ShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityGroupTags 查询安全组资源标签
//
// 查询指定安全组实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowSecurityGroupTags(request *model.ShowSecurityGroupTagsRequest) (*model.ShowSecurityGroupTagsResponse, error) {
	requestDef := GenReqDefForShowSecurityGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityGroupTagsResponse), nil
	}
}

// ShowSecurityGroupTagsInvoker 查询安全组资源标签
func (c *VpcClient) ShowSecurityGroupTagsInvoker(request *model.ShowSecurityGroupTagsRequest) *ShowSecurityGroupTagsInvoker {
	requestDef := GenReqDefForShowSecurityGroupTags()
	return &ShowSecurityGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubnet 查询子网
//
// 查询子网详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowSubnet(request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	requestDef := GenReqDefForShowSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetResponse), nil
	}
}

// ShowSubnetInvoker 查询子网
func (c *VpcClient) ShowSubnetInvoker(request *model.ShowSubnetRequest) *ShowSubnetInvoker {
	requestDef := GenReqDefForShowSubnet()
	return &ShowSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSubnetTags 查询子网资源标签
//
// 查询指定子网实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowSubnetTags(request *model.ShowSubnetTagsRequest) (*model.ShowSubnetTagsResponse, error) {
	requestDef := GenReqDefForShowSubnetTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSubnetTagsResponse), nil
	}
}

// ShowSubnetTagsInvoker 查询子网资源标签
func (c *VpcClient) ShowSubnetTagsInvoker(request *model.ShowSubnetTagsRequest) *ShowSubnetTagsInvoker {
	requestDef := GenReqDefForShowSubnetTags()
	return &ShowSubnetTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpcPeering 查询对等连接
//
// 查询对等连接详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowVpcPeering(request *model.ShowVpcPeeringRequest) (*model.ShowVpcPeeringResponse, error) {
	requestDef := GenReqDefForShowVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcPeeringResponse), nil
	}
}

// ShowVpcPeeringInvoker 查询对等连接
func (c *VpcClient) ShowVpcPeeringInvoker(request *model.ShowVpcPeeringRequest) *ShowVpcPeeringInvoker {
	requestDef := GenReqDefForShowVpcPeering()
	return &ShowVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlowLog 更新流日志
//
// 更新流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdateFlowLog(request *model.UpdateFlowLogRequest) (*model.UpdateFlowLogResponse, error) {
	requestDef := GenReqDefForUpdateFlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlowLogResponse), nil
	}
}

// UpdateFlowLogInvoker 更新流日志
func (c *VpcClient) UpdateFlowLogInvoker(request *model.UpdateFlowLogRequest) *UpdateFlowLogInvoker {
	requestDef := GenReqDefForUpdateFlowLog()
	return &UpdateFlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePort 更新端口
//
// 更新端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdatePort(request *model.UpdatePortRequest) (*model.UpdatePortResponse, error) {
	requestDef := GenReqDefForUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePortResponse), nil
	}
}

// UpdatePortInvoker 更新端口
func (c *VpcClient) UpdatePortInvoker(request *model.UpdatePortRequest) *UpdatePortInvoker {
	requestDef := GenReqDefForUpdatePort()
	return &UpdatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRouteTable 更新路由表
//
// 更新路由表，包括可以更新路由表的名称，描述，以及新增、更新、删除路由条目
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdateRouteTable(request *model.UpdateRouteTableRequest) (*model.UpdateRouteTableResponse, error) {
	requestDef := GenReqDefForUpdateRouteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRouteTableResponse), nil
	}
}

// UpdateRouteTableInvoker 更新路由表
func (c *VpcClient) UpdateRouteTableInvoker(request *model.UpdateRouteTableRequest) *UpdateRouteTableInvoker {
	requestDef := GenReqDefForUpdateRouteTable()
	return &UpdateRouteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubnet 更新子网
//
// 更新子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdateSubnet(request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	requestDef := GenReqDefForUpdateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubnetResponse), nil
	}
}

// UpdateSubnetInvoker 更新子网
func (c *VpcClient) UpdateSubnetInvoker(request *model.UpdateSubnetRequest) *UpdateSubnetInvoker {
	requestDef := GenReqDefForUpdateSubnet()
	return &UpdateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpcPeering 更新对等连接
//
// 更新对等连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdateVpcPeering(request *model.UpdateVpcPeeringRequest) (*model.UpdateVpcPeeringResponse, error) {
	requestDef := GenReqDefForUpdateVpcPeering()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcPeeringResponse), nil
	}
}

// UpdateVpcPeeringInvoker 更新对等连接
func (c *VpcClient) UpdateVpcPeeringInvoker(request *model.UpdateVpcPeeringRequest) *UpdateVpcPeeringInvoker {
	requestDef := GenReqDefForUpdateVpcPeering()
	return &UpdateVpcPeeringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePrivateip 申请私有IP
//
// 申请私有IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreatePrivateip(request *model.CreatePrivateipRequest) (*model.CreatePrivateipResponse, error) {
	requestDef := GenReqDefForCreatePrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateipResponse), nil
	}
}

// CreatePrivateipInvoker 申请私有IP
func (c *VpcClient) CreatePrivateipInvoker(request *model.CreatePrivateipRequest) *CreatePrivateipInvoker {
	requestDef := GenReqDefForCreatePrivateip()
	return &CreatePrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePrivateip 删除私有IP
//
// 删除私有IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeletePrivateip(request *model.DeletePrivateipRequest) (*model.DeletePrivateipResponse, error) {
	requestDef := GenReqDefForDeletePrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateipResponse), nil
	}
}

// DeletePrivateipInvoker 删除私有IP
func (c *VpcClient) DeletePrivateipInvoker(request *model.DeletePrivateipRequest) *DeletePrivateipInvoker {
	requestDef := GenReqDefForDeletePrivateip()
	return &DeletePrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrivateips 查询私有IP列表
//
// 查询指定子网下的私有IP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListPrivateips(request *model.ListPrivateipsRequest) (*model.ListPrivateipsResponse, error) {
	requestDef := GenReqDefForListPrivateips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateipsResponse), nil
	}
}

// ListPrivateipsInvoker 查询私有IP列表
func (c *VpcClient) ListPrivateipsInvoker(request *model.ListPrivateipsRequest) *ListPrivateipsInvoker {
	requestDef := GenReqDefForListPrivateips()
	return &ListPrivateipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkIpAvailabilities 查询网络IP使用情况
//
// 显示一个指定网络中的IPv4地址使用情况。
// 包括此网络中的IP总数以及已用IP总数，以及网络下每一个子网的IP地址总数和可用IP地址总数。
//
// &gt; 须知
//
// - 系统预留地址指的是子网的第1个以及最后4个地址，一般用于网关、DHCP等服务。
// - 这里以及下文描述的IP地址总数、已用IP地址总数不包含系统预留地址。
// - 在分配IP时，用户可以指定系统预留的IP地址。但是不论IP是如何分配的，只要是处于系统预留IP地址段的IP均不会被统计到已用IP地址数目和IP地址总数中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowNetworkIpAvailabilities(request *model.ShowNetworkIpAvailabilitiesRequest) (*model.ShowNetworkIpAvailabilitiesResponse, error) {
	requestDef := GenReqDefForShowNetworkIpAvailabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkIpAvailabilitiesResponse), nil
	}
}

// ShowNetworkIpAvailabilitiesInvoker 查询网络IP使用情况
func (c *VpcClient) ShowNetworkIpAvailabilitiesInvoker(request *model.ShowNetworkIpAvailabilitiesRequest) *ShowNetworkIpAvailabilitiesInvoker {
	requestDef := GenReqDefForShowNetworkIpAvailabilities()
	return &ShowNetworkIpAvailabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPrivateip 查询私有IP
//
// 指定ID查询私有IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowPrivateip(request *model.ShowPrivateipRequest) (*model.ShowPrivateipResponse, error) {
	requestDef := GenReqDefForShowPrivateip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateipResponse), nil
	}
}

// ShowPrivateipInvoker 查询私有IP
func (c *VpcClient) ShowPrivateipInvoker(request *model.ShowPrivateipRequest) *ShowPrivateipInvoker {
	requestDef := GenReqDefForShowPrivateip()
	return &ShowPrivateipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronAddRouterInterface 路由器添加接口
//
// 添加路由器接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronAddRouterInterface(request *model.NeutronAddRouterInterfaceRequest) (*model.NeutronAddRouterInterfaceResponse, error) {
	requestDef := GenReqDefForNeutronAddRouterInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronAddRouterInterfaceResponse), nil
	}
}

// NeutronAddRouterInterfaceInvoker 路由器添加接口
func (c *VpcClient) NeutronAddRouterInterfaceInvoker(request *model.NeutronAddRouterInterfaceRequest) *NeutronAddRouterInterfaceInvoker {
	requestDef := GenReqDefForNeutronAddRouterInterface()
	return &NeutronAddRouterInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateNetwork 创建网络
//
// 创建网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateNetwork(request *model.NeutronCreateNetworkRequest) (*model.NeutronCreateNetworkResponse, error) {
	requestDef := GenReqDefForNeutronCreateNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateNetworkResponse), nil
	}
}

// NeutronCreateNetworkInvoker 创建网络
func (c *VpcClient) NeutronCreateNetworkInvoker(request *model.NeutronCreateNetworkRequest) *NeutronCreateNetworkInvoker {
	requestDef := GenReqDefForNeutronCreateNetwork()
	return &NeutronCreateNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreatePort 创建端口
//
// 创建端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreatePort(request *model.NeutronCreatePortRequest) (*model.NeutronCreatePortResponse, error) {
	requestDef := GenReqDefForNeutronCreatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreatePortResponse), nil
	}
}

// NeutronCreatePortInvoker 创建端口
func (c *VpcClient) NeutronCreatePortInvoker(request *model.NeutronCreatePortRequest) *NeutronCreatePortInvoker {
	requestDef := GenReqDefForNeutronCreatePort()
	return &NeutronCreatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateRouter 创建路由器
//
// 创建路由器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateRouter(request *model.NeutronCreateRouterRequest) (*model.NeutronCreateRouterResponse, error) {
	requestDef := GenReqDefForNeutronCreateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateRouterResponse), nil
	}
}

// NeutronCreateRouterInvoker 创建路由器
func (c *VpcClient) NeutronCreateRouterInvoker(request *model.NeutronCreateRouterRequest) *NeutronCreateRouterInvoker {
	requestDef := GenReqDefForNeutronCreateRouter()
	return &NeutronCreateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateSecurityGroup 创建安全组
//
// 创建安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateSecurityGroup(request *model.NeutronCreateSecurityGroupRequest) (*model.NeutronCreateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronCreateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateSecurityGroupResponse), nil
	}
}

// NeutronCreateSecurityGroupInvoker 创建安全组
func (c *VpcClient) NeutronCreateSecurityGroupInvoker(request *model.NeutronCreateSecurityGroupRequest) *NeutronCreateSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronCreateSecurityGroup()
	return &NeutronCreateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateSecurityGroupRule 创建安全组规则
//
// 创建安全组规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateSecurityGroupRule(request *model.NeutronCreateSecurityGroupRuleRequest) (*model.NeutronCreateSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronCreateSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateSecurityGroupRuleResponse), nil
	}
}

// NeutronCreateSecurityGroupRuleInvoker 创建安全组规则
func (c *VpcClient) NeutronCreateSecurityGroupRuleInvoker(request *model.NeutronCreateSecurityGroupRuleRequest) *NeutronCreateSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronCreateSecurityGroupRule()
	return &NeutronCreateSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateSubnet 创建子网
//
// 创建子网。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateSubnet(request *model.NeutronCreateSubnetRequest) (*model.NeutronCreateSubnetResponse, error) {
	requestDef := GenReqDefForNeutronCreateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateSubnetResponse), nil
	}
}

// NeutronCreateSubnetInvoker 创建子网
func (c *VpcClient) NeutronCreateSubnetInvoker(request *model.NeutronCreateSubnetRequest) *NeutronCreateSubnetInvoker {
	requestDef := GenReqDefForNeutronCreateSubnet()
	return &NeutronCreateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteNetwork 删除网络
//
// 删除网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteNetwork(request *model.NeutronDeleteNetworkRequest) (*model.NeutronDeleteNetworkResponse, error) {
	requestDef := GenReqDefForNeutronDeleteNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteNetworkResponse), nil
	}
}

// NeutronDeleteNetworkInvoker 删除网络
func (c *VpcClient) NeutronDeleteNetworkInvoker(request *model.NeutronDeleteNetworkRequest) *NeutronDeleteNetworkInvoker {
	requestDef := GenReqDefForNeutronDeleteNetwork()
	return &NeutronDeleteNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeletePort 删除端口
//
// 删除端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeletePort(request *model.NeutronDeletePortRequest) (*model.NeutronDeletePortResponse, error) {
	requestDef := GenReqDefForNeutronDeletePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeletePortResponse), nil
	}
}

// NeutronDeletePortInvoker 删除端口
func (c *VpcClient) NeutronDeletePortInvoker(request *model.NeutronDeletePortRequest) *NeutronDeletePortInvoker {
	requestDef := GenReqDefForNeutronDeletePort()
	return &NeutronDeletePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteRouter 删除路由器
//
// 删除路由器
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteRouter(request *model.NeutronDeleteRouterRequest) (*model.NeutronDeleteRouterResponse, error) {
	requestDef := GenReqDefForNeutronDeleteRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteRouterResponse), nil
	}
}

// NeutronDeleteRouterInvoker 删除路由器
func (c *VpcClient) NeutronDeleteRouterInvoker(request *model.NeutronDeleteRouterRequest) *NeutronDeleteRouterInvoker {
	requestDef := GenReqDefForNeutronDeleteRouter()
	return &NeutronDeleteRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteSecurityGroup 删除安全组
//
// 删除安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteSecurityGroup(request *model.NeutronDeleteSecurityGroupRequest) (*model.NeutronDeleteSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronDeleteSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteSecurityGroupResponse), nil
	}
}

// NeutronDeleteSecurityGroupInvoker 删除安全组
func (c *VpcClient) NeutronDeleteSecurityGroupInvoker(request *model.NeutronDeleteSecurityGroupRequest) *NeutronDeleteSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronDeleteSecurityGroup()
	return &NeutronDeleteSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteSecurityGroupRule 删除安全组规则
//
// 删除安全组规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteSecurityGroupRule(request *model.NeutronDeleteSecurityGroupRuleRequest) (*model.NeutronDeleteSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronDeleteSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteSecurityGroupRuleResponse), nil
	}
}

// NeutronDeleteSecurityGroupRuleInvoker 删除安全组规则
func (c *VpcClient) NeutronDeleteSecurityGroupRuleInvoker(request *model.NeutronDeleteSecurityGroupRuleRequest) *NeutronDeleteSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronDeleteSecurityGroupRule()
	return &NeutronDeleteSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteSubnet 删除子网
//
// 删除子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteSubnet(request *model.NeutronDeleteSubnetRequest) (*model.NeutronDeleteSubnetResponse, error) {
	requestDef := GenReqDefForNeutronDeleteSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteSubnetResponse), nil
	}
}

// NeutronDeleteSubnetInvoker 删除子网
func (c *VpcClient) NeutronDeleteSubnetInvoker(request *model.NeutronDeleteSubnetRequest) *NeutronDeleteSubnetInvoker {
	requestDef := GenReqDefForNeutronDeleteSubnet()
	return &NeutronDeleteSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListNetworks 查询网络列表
//
// 查询提交请求的租户的所有网络，单次查询最多返回2000条数据，超过2000后会返回分页标记。分页查询请参考分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListNetworks(request *model.NeutronListNetworksRequest) (*model.NeutronListNetworksResponse, error) {
	requestDef := GenReqDefForNeutronListNetworks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListNetworksResponse), nil
	}
}

// NeutronListNetworksInvoker 查询网络列表
func (c *VpcClient) NeutronListNetworksInvoker(request *model.NeutronListNetworksRequest) *NeutronListNetworksInvoker {
	requestDef := GenReqDefForNeutronListNetworks()
	return &NeutronListNetworksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListPorts 查询端口列表
//
// 查询提交请求的租户的所有端口，单次查询最多返回2000条数据，超过2000后会返回分页标记。分页查询请参考分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListPorts(request *model.NeutronListPortsRequest) (*model.NeutronListPortsResponse, error) {
	requestDef := GenReqDefForNeutronListPorts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListPortsResponse), nil
	}
}

// NeutronListPortsInvoker 查询端口列表
func (c *VpcClient) NeutronListPortsInvoker(request *model.NeutronListPortsRequest) *NeutronListPortsInvoker {
	requestDef := GenReqDefForNeutronListPorts()
	return &NeutronListPortsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListRouters 查询路由器列表
//
// 查询提交请求的租户有权限操作的所有路由器信息，单次查询最多返回2000条数据，超过2000后会返回分页标记。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListRouters(request *model.NeutronListRoutersRequest) (*model.NeutronListRoutersResponse, error) {
	requestDef := GenReqDefForNeutronListRouters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListRoutersResponse), nil
	}
}

// NeutronListRoutersInvoker 查询路由器列表
func (c *VpcClient) NeutronListRoutersInvoker(request *model.NeutronListRoutersRequest) *NeutronListRoutersInvoker {
	requestDef := GenReqDefForNeutronListRouters()
	return &NeutronListRoutersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListSecurityGroupRules 查询安全组规则列表
//
// 查询提交请求的租户有权限查看的所有安全组规则。单次查询最多返回2000条数据，超过2000后会返回分页标记。分页查询请参考分页查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListSecurityGroupRules(request *model.NeutronListSecurityGroupRulesRequest) (*model.NeutronListSecurityGroupRulesResponse, error) {
	requestDef := GenReqDefForNeutronListSecurityGroupRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListSecurityGroupRulesResponse), nil
	}
}

// NeutronListSecurityGroupRulesInvoker 查询安全组规则列表
func (c *VpcClient) NeutronListSecurityGroupRulesInvoker(request *model.NeutronListSecurityGroupRulesRequest) *NeutronListSecurityGroupRulesInvoker {
	requestDef := GenReqDefForNeutronListSecurityGroupRules()
	return &NeutronListSecurityGroupRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListSecurityGroups 查询安全组列表
//
// 查询提交请求租户的所有安全组，单次查询最多返回2000条数据，超过2000后会返回分页标记。分页查询请参考分页查询 。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListSecurityGroups(request *model.NeutronListSecurityGroupsRequest) (*model.NeutronListSecurityGroupsResponse, error) {
	requestDef := GenReqDefForNeutronListSecurityGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListSecurityGroupsResponse), nil
	}
}

// NeutronListSecurityGroupsInvoker 查询安全组列表
func (c *VpcClient) NeutronListSecurityGroupsInvoker(request *model.NeutronListSecurityGroupsRequest) *NeutronListSecurityGroupsInvoker {
	requestDef := GenReqDefForNeutronListSecurityGroups()
	return &NeutronListSecurityGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListSubnets 查询子网列表
//
// 查询提交请求租户的所有子网，单次查询最多返回2000条数据，超过2000后会返回分页标记。分页查询请参考分页查询 。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListSubnets(request *model.NeutronListSubnetsRequest) (*model.NeutronListSubnetsResponse, error) {
	requestDef := GenReqDefForNeutronListSubnets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListSubnetsResponse), nil
	}
}

// NeutronListSubnetsInvoker 查询子网列表
func (c *VpcClient) NeutronListSubnetsInvoker(request *model.NeutronListSubnetsRequest) *NeutronListSubnetsInvoker {
	requestDef := GenReqDefForNeutronListSubnets()
	return &NeutronListSubnetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronRemoveRouterInterface 路由器删除接口
//
// 删除路由器接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronRemoveRouterInterface(request *model.NeutronRemoveRouterInterfaceRequest) (*model.NeutronRemoveRouterInterfaceResponse, error) {
	requestDef := GenReqDefForNeutronRemoveRouterInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronRemoveRouterInterfaceResponse), nil
	}
}

// NeutronRemoveRouterInterfaceInvoker 路由器删除接口
func (c *VpcClient) NeutronRemoveRouterInterfaceInvoker(request *model.NeutronRemoveRouterInterfaceRequest) *NeutronRemoveRouterInterfaceInvoker {
	requestDef := GenReqDefForNeutronRemoveRouterInterface()
	return &NeutronRemoveRouterInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowNetwork 查询网络
//
// 查询指定的网络详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowNetwork(request *model.NeutronShowNetworkRequest) (*model.NeutronShowNetworkResponse, error) {
	requestDef := GenReqDefForNeutronShowNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowNetworkResponse), nil
	}
}

// NeutronShowNetworkInvoker 查询网络
func (c *VpcClient) NeutronShowNetworkInvoker(request *model.NeutronShowNetworkRequest) *NeutronShowNetworkInvoker {
	requestDef := GenReqDefForNeutronShowNetwork()
	return &NeutronShowNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowPort 查询端口
//
// 查询端口详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowPort(request *model.NeutronShowPortRequest) (*model.NeutronShowPortResponse, error) {
	requestDef := GenReqDefForNeutronShowPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowPortResponse), nil
	}
}

// NeutronShowPortInvoker 查询端口
func (c *VpcClient) NeutronShowPortInvoker(request *model.NeutronShowPortRequest) *NeutronShowPortInvoker {
	requestDef := GenReqDefForNeutronShowPort()
	return &NeutronShowPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowRouter 查询路由器
//
// 查询路由器详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowRouter(request *model.NeutronShowRouterRequest) (*model.NeutronShowRouterResponse, error) {
	requestDef := GenReqDefForNeutronShowRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowRouterResponse), nil
	}
}

// NeutronShowRouterInvoker 查询路由器
func (c *VpcClient) NeutronShowRouterInvoker(request *model.NeutronShowRouterRequest) *NeutronShowRouterInvoker {
	requestDef := GenReqDefForNeutronShowRouter()
	return &NeutronShowRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowSecurityGroup 查询安全组
//
// 查询安全组详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowSecurityGroup(request *model.NeutronShowSecurityGroupRequest) (*model.NeutronShowSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronShowSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowSecurityGroupResponse), nil
	}
}

// NeutronShowSecurityGroupInvoker 查询安全组
func (c *VpcClient) NeutronShowSecurityGroupInvoker(request *model.NeutronShowSecurityGroupRequest) *NeutronShowSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronShowSecurityGroup()
	return &NeutronShowSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowSecurityGroupRule 查询安全组规则
//
// 查询安全组规则详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowSecurityGroupRule(request *model.NeutronShowSecurityGroupRuleRequest) (*model.NeutronShowSecurityGroupRuleResponse, error) {
	requestDef := GenReqDefForNeutronShowSecurityGroupRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowSecurityGroupRuleResponse), nil
	}
}

// NeutronShowSecurityGroupRuleInvoker 查询安全组规则
func (c *VpcClient) NeutronShowSecurityGroupRuleInvoker(request *model.NeutronShowSecurityGroupRuleRequest) *NeutronShowSecurityGroupRuleInvoker {
	requestDef := GenReqDefForNeutronShowSecurityGroupRule()
	return &NeutronShowSecurityGroupRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowSubnet 查询子网
//
// 查询子网详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowSubnet(request *model.NeutronShowSubnetRequest) (*model.NeutronShowSubnetResponse, error) {
	requestDef := GenReqDefForNeutronShowSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowSubnetResponse), nil
	}
}

// NeutronShowSubnetInvoker 查询子网
func (c *VpcClient) NeutronShowSubnetInvoker(request *model.NeutronShowSubnetRequest) *NeutronShowSubnetInvoker {
	requestDef := GenReqDefForNeutronShowSubnet()
	return &NeutronShowSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateNetwork 更新网络
//
// 更新网络
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateNetwork(request *model.NeutronUpdateNetworkRequest) (*model.NeutronUpdateNetworkResponse, error) {
	requestDef := GenReqDefForNeutronUpdateNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateNetworkResponse), nil
	}
}

// NeutronUpdateNetworkInvoker 更新网络
func (c *VpcClient) NeutronUpdateNetworkInvoker(request *model.NeutronUpdateNetworkRequest) *NeutronUpdateNetworkInvoker {
	requestDef := GenReqDefForNeutronUpdateNetwork()
	return &NeutronUpdateNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdatePort 更新端口
//
// 更新端口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdatePort(request *model.NeutronUpdatePortRequest) (*model.NeutronUpdatePortResponse, error) {
	requestDef := GenReqDefForNeutronUpdatePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdatePortResponse), nil
	}
}

// NeutronUpdatePortInvoker 更新端口
func (c *VpcClient) NeutronUpdatePortInvoker(request *model.NeutronUpdatePortRequest) *NeutronUpdatePortInvoker {
	requestDef := GenReqDefForNeutronUpdatePort()
	return &NeutronUpdatePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateRouter 更新路由器
//
// 更新路由器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateRouter(request *model.NeutronUpdateRouterRequest) (*model.NeutronUpdateRouterResponse, error) {
	requestDef := GenReqDefForNeutronUpdateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateRouterResponse), nil
	}
}

// NeutronUpdateRouterInvoker 更新路由器
func (c *VpcClient) NeutronUpdateRouterInvoker(request *model.NeutronUpdateRouterRequest) *NeutronUpdateRouterInvoker {
	requestDef := GenReqDefForNeutronUpdateRouter()
	return &NeutronUpdateRouterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateSecurityGroup 更新安全组
//
// 更新安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateSecurityGroup(request *model.NeutronUpdateSecurityGroupRequest) (*model.NeutronUpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForNeutronUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateSecurityGroupResponse), nil
	}
}

// NeutronUpdateSecurityGroupInvoker 更新安全组
func (c *VpcClient) NeutronUpdateSecurityGroupInvoker(request *model.NeutronUpdateSecurityGroupRequest) *NeutronUpdateSecurityGroupInvoker {
	requestDef := GenReqDefForNeutronUpdateSecurityGroup()
	return &NeutronUpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateSubnet 更新子网
//
// 更新子网
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateSubnet(request *model.NeutronUpdateSubnetRequest) (*model.NeutronUpdateSubnetResponse, error) {
	requestDef := GenReqDefForNeutronUpdateSubnet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateSubnetResponse), nil
	}
}

// NeutronUpdateSubnetInvoker 更新子网
func (c *VpcClient) NeutronUpdateSubnetInvoker(request *model.NeutronUpdateSubnetRequest) *NeutronUpdateSubnetInvoker {
	requestDef := GenReqDefForNeutronUpdateSubnet()
	return &NeutronUpdateSubnetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronAddFirewallRule 插入网络ACL规则
//
// 插入一条网络ACL规则到某一网络ACL策略中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronAddFirewallRule(request *model.NeutronAddFirewallRuleRequest) (*model.NeutronAddFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronAddFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronAddFirewallRuleResponse), nil
	}
}

// NeutronAddFirewallRuleInvoker 插入网络ACL规则
func (c *VpcClient) NeutronAddFirewallRuleInvoker(request *model.NeutronAddFirewallRuleRequest) *NeutronAddFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronAddFirewallRule()
	return &NeutronAddFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateFirewallGroup 创建网络ACL组
//
// 创建网络ACL组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateFirewallGroup(request *model.NeutronCreateFirewallGroupRequest) (*model.NeutronCreateFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallGroupResponse), nil
	}
}

// NeutronCreateFirewallGroupInvoker 创建网络ACL组
func (c *VpcClient) NeutronCreateFirewallGroupInvoker(request *model.NeutronCreateFirewallGroupRequest) *NeutronCreateFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallGroup()
	return &NeutronCreateFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateFirewallPolicy 创建网络ACL策略
//
// 创建网络ACL策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateFirewallPolicy(request *model.NeutronCreateFirewallPolicyRequest) (*model.NeutronCreateFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallPolicyResponse), nil
	}
}

// NeutronCreateFirewallPolicyInvoker 创建网络ACL策略
func (c *VpcClient) NeutronCreateFirewallPolicyInvoker(request *model.NeutronCreateFirewallPolicyRequest) *NeutronCreateFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallPolicy()
	return &NeutronCreateFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronCreateFirewallRule 创建网络ACL规则
//
// 创建网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronCreateFirewallRule(request *model.NeutronCreateFirewallRuleRequest) (*model.NeutronCreateFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronCreateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronCreateFirewallRuleResponse), nil
	}
}

// NeutronCreateFirewallRuleInvoker 创建网络ACL规则
func (c *VpcClient) NeutronCreateFirewallRuleInvoker(request *model.NeutronCreateFirewallRuleRequest) *NeutronCreateFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronCreateFirewallRule()
	return &NeutronCreateFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteFirewallGroup 删除网络ACL组
//
// 删除网络ACL组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteFirewallGroup(request *model.NeutronDeleteFirewallGroupRequest) (*model.NeutronDeleteFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallGroupResponse), nil
	}
}

// NeutronDeleteFirewallGroupInvoker 删除网络ACL组
func (c *VpcClient) NeutronDeleteFirewallGroupInvoker(request *model.NeutronDeleteFirewallGroupRequest) *NeutronDeleteFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallGroup()
	return &NeutronDeleteFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteFirewallPolicy 删除网络ACL策略
//
// 删除网络ACL策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteFirewallPolicy(request *model.NeutronDeleteFirewallPolicyRequest) (*model.NeutronDeleteFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallPolicyResponse), nil
	}
}

// NeutronDeleteFirewallPolicyInvoker 删除网络ACL策略
func (c *VpcClient) NeutronDeleteFirewallPolicyInvoker(request *model.NeutronDeleteFirewallPolicyRequest) *NeutronDeleteFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallPolicy()
	return &NeutronDeleteFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronDeleteFirewallRule 删除网络ACL规则
//
// 删除网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronDeleteFirewallRule(request *model.NeutronDeleteFirewallRuleRequest) (*model.NeutronDeleteFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronDeleteFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronDeleteFirewallRuleResponse), nil
	}
}

// NeutronDeleteFirewallRuleInvoker 删除网络ACL规则
func (c *VpcClient) NeutronDeleteFirewallRuleInvoker(request *model.NeutronDeleteFirewallRuleRequest) *NeutronDeleteFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronDeleteFirewallRule()
	return &NeutronDeleteFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListFirewallGroups 查询所有网络ACL组
//
// 查询提交请求的租户有权限操作的所有网络ACL组信息。单次查询最多返回2000条数据，超过2000后会返回分页标记。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListFirewallGroups(request *model.NeutronListFirewallGroupsRequest) (*model.NeutronListFirewallGroupsResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallGroupsResponse), nil
	}
}

// NeutronListFirewallGroupsInvoker 查询所有网络ACL组
func (c *VpcClient) NeutronListFirewallGroupsInvoker(request *model.NeutronListFirewallGroupsRequest) *NeutronListFirewallGroupsInvoker {
	requestDef := GenReqDefForNeutronListFirewallGroups()
	return &NeutronListFirewallGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListFirewallPolicies 查询所有网络ACL策略
//
// 查询提交请求的租户有权限操作的所有网络ACL策略信息。单次查询最多返回2000条数据，超过2000后会返回分页标记。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListFirewallPolicies(request *model.NeutronListFirewallPoliciesRequest) (*model.NeutronListFirewallPoliciesResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallPoliciesResponse), nil
	}
}

// NeutronListFirewallPoliciesInvoker 查询所有网络ACL策略
func (c *VpcClient) NeutronListFirewallPoliciesInvoker(request *model.NeutronListFirewallPoliciesRequest) *NeutronListFirewallPoliciesInvoker {
	requestDef := GenReqDefForNeutronListFirewallPolicies()
	return &NeutronListFirewallPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronListFirewallRules 查询所有网络ACL规则
//
// 查询提交请求的租户有权限操作的所有网络ACL规则信息。单次查询最多返回2000条数据，超过2000后会返回分页标记。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronListFirewallRules(request *model.NeutronListFirewallRulesRequest) (*model.NeutronListFirewallRulesResponse, error) {
	requestDef := GenReqDefForNeutronListFirewallRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronListFirewallRulesResponse), nil
	}
}

// NeutronListFirewallRulesInvoker 查询所有网络ACL规则
func (c *VpcClient) NeutronListFirewallRulesInvoker(request *model.NeutronListFirewallRulesRequest) *NeutronListFirewallRulesInvoker {
	requestDef := GenReqDefForNeutronListFirewallRules()
	return &NeutronListFirewallRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronRemoveFirewallRule 移除网络ACL规则
//
// 从某一网络ACL策略中移除一条网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronRemoveFirewallRule(request *model.NeutronRemoveFirewallRuleRequest) (*model.NeutronRemoveFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronRemoveFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronRemoveFirewallRuleResponse), nil
	}
}

// NeutronRemoveFirewallRuleInvoker 移除网络ACL规则
func (c *VpcClient) NeutronRemoveFirewallRuleInvoker(request *model.NeutronRemoveFirewallRuleRequest) *NeutronRemoveFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronRemoveFirewallRule()
	return &NeutronRemoveFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowFirewallGroup 查询特定网络ACL组详情
//
// 查询特定网络ACL组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowFirewallGroup(request *model.NeutronShowFirewallGroupRequest) (*model.NeutronShowFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallGroupResponse), nil
	}
}

// NeutronShowFirewallGroupInvoker 查询特定网络ACL组详情
func (c *VpcClient) NeutronShowFirewallGroupInvoker(request *model.NeutronShowFirewallGroupRequest) *NeutronShowFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronShowFirewallGroup()
	return &NeutronShowFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowFirewallPolicy 查询特定网络ACL策略详情
//
// 查询特定网络ACL策略详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowFirewallPolicy(request *model.NeutronShowFirewallPolicyRequest) (*model.NeutronShowFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallPolicyResponse), nil
	}
}

// NeutronShowFirewallPolicyInvoker 查询特定网络ACL策略详情
func (c *VpcClient) NeutronShowFirewallPolicyInvoker(request *model.NeutronShowFirewallPolicyRequest) *NeutronShowFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronShowFirewallPolicy()
	return &NeutronShowFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronShowFirewallRule 查询特定网络ACL规则
//
// 查询特定网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronShowFirewallRule(request *model.NeutronShowFirewallRuleRequest) (*model.NeutronShowFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronShowFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronShowFirewallRuleResponse), nil
	}
}

// NeutronShowFirewallRuleInvoker 查询特定网络ACL规则
func (c *VpcClient) NeutronShowFirewallRuleInvoker(request *model.NeutronShowFirewallRuleRequest) *NeutronShowFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronShowFirewallRule()
	return &NeutronShowFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateFirewallGroup 更新网络ACL组
//
// 更新网络ACL组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateFirewallGroup(request *model.NeutronUpdateFirewallGroupRequest) (*model.NeutronUpdateFirewallGroupResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallGroupResponse), nil
	}
}

// NeutronUpdateFirewallGroupInvoker 更新网络ACL组
func (c *VpcClient) NeutronUpdateFirewallGroupInvoker(request *model.NeutronUpdateFirewallGroupRequest) *NeutronUpdateFirewallGroupInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallGroup()
	return &NeutronUpdateFirewallGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateFirewallPolicy 更新网络ACL策略
//
// 更新网络ACL策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateFirewallPolicy(request *model.NeutronUpdateFirewallPolicyRequest) (*model.NeutronUpdateFirewallPolicyResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallPolicyResponse), nil
	}
}

// NeutronUpdateFirewallPolicyInvoker 更新网络ACL策略
func (c *VpcClient) NeutronUpdateFirewallPolicyInvoker(request *model.NeutronUpdateFirewallPolicyRequest) *NeutronUpdateFirewallPolicyInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallPolicy()
	return &NeutronUpdateFirewallPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NeutronUpdateFirewallRule 更新网络ACL规则
//
// 更新网络ACL规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) NeutronUpdateFirewallRule(request *model.NeutronUpdateFirewallRuleRequest) (*model.NeutronUpdateFirewallRuleResponse, error) {
	requestDef := GenReqDefForNeutronUpdateFirewallRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NeutronUpdateFirewallRuleResponse), nil
	}
}

// NeutronUpdateFirewallRuleInvoker 更新网络ACL规则
func (c *VpcClient) NeutronUpdateFirewallRuleInvoker(request *model.NeutronUpdateFirewallRuleRequest) *NeutronUpdateFirewallRuleInvoker {
	requestDef := GenReqDefForNeutronUpdateFirewallRule()
	return &NeutronUpdateFirewallRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersion 查询API版本信息列表
//
// 返回当前API所有可用的版本（仅针对OpenStack原生接口）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

// ListApiVersionInvoker 查询API版本信息列表
func (c *VpcClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateVpcTags 批量创建VPC资源标签
//
// 为指定的VPC资源实例批量添加标签。
// 此接口为幂等接口：创建时如果请求体中存在重复key则报错。创建时，不允许设置重复key数据，如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchCreateVpcTags(request *model.BatchCreateVpcTagsRequest) (*model.BatchCreateVpcTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateVpcTagsResponse), nil
	}
}

// BatchCreateVpcTagsInvoker 批量创建VPC资源标签
func (c *VpcClient) BatchCreateVpcTagsInvoker(request *model.BatchCreateVpcTagsRequest) *BatchCreateVpcTagsInvoker {
	requestDef := GenReqDefForBatchCreateVpcTags()
	return &BatchCreateVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteVpcTags 批量删除VPC资源标签
//
// 为指定的VPC资源实例批量删除标签。
// 此接口为幂等接口：删除时，如果删除的标签不存在，默认处理成功；删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) BatchDeleteVpcTags(request *model.BatchDeleteVpcTagsRequest) (*model.BatchDeleteVpcTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteVpcTagsResponse), nil
	}
}

// BatchDeleteVpcTagsInvoker 批量删除VPC资源标签
func (c *VpcClient) BatchDeleteVpcTagsInvoker(request *model.BatchDeleteVpcTagsRequest) *BatchDeleteVpcTagsInvoker {
	requestDef := GenReqDefForBatchDeleteVpcTags()
	return &BatchDeleteVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpc 创建VPC
//
// 创建虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateVpc(request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	requestDef := GenReqDefForCreateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResponse), nil
	}
}

// CreateVpcInvoker 创建VPC
func (c *VpcClient) CreateVpcInvoker(request *model.CreateVpcRequest) *CreateVpcInvoker {
	requestDef := GenReqDefForCreateVpc()
	return &CreateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcResourceTag 创建VPC资源标签
//
// 给指定VPC资源实例增加标签信息
// 此接口为幂等接口：创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateVpcResourceTag(request *model.CreateVpcResourceTagRequest) (*model.CreateVpcResourceTagResponse, error) {
	requestDef := GenReqDefForCreateVpcResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcResourceTagResponse), nil
	}
}

// CreateVpcResourceTagInvoker 创建VPC资源标签
func (c *VpcClient) CreateVpcResourceTagInvoker(request *model.CreateVpcResourceTagRequest) *CreateVpcResourceTagInvoker {
	requestDef := GenReqDefForCreateVpcResourceTag()
	return &CreateVpcResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVpcRoute 创建VPC路由
//
// 创建路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) CreateVpcRoute(request *model.CreateVpcRouteRequest) (*model.CreateVpcRouteResponse, error) {
	requestDef := GenReqDefForCreateVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVpcRouteResponse), nil
	}
}

// CreateVpcRouteInvoker 创建VPC路由
func (c *VpcClient) CreateVpcRouteInvoker(request *model.CreateVpcRouteRequest) *CreateVpcRouteInvoker {
	requestDef := GenReqDefForCreateVpcRoute()
	return &CreateVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpc 删除VPC
//
// 删除虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteVpc(request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	requestDef := GenReqDefForDeleteVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcResponse), nil
	}
}

// DeleteVpcInvoker 删除VPC
func (c *VpcClient) DeleteVpcInvoker(request *model.DeleteVpcRequest) *DeleteVpcInvoker {
	requestDef := GenReqDefForDeleteVpc()
	return &DeleteVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcRoute 删除VPC路由
//
// 删除路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteVpcRoute(request *model.DeleteVpcRouteRequest) (*model.DeleteVpcRouteResponse, error) {
	requestDef := GenReqDefForDeleteVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcRouteResponse), nil
	}
}

// DeleteVpcRouteInvoker 删除VPC路由
func (c *VpcClient) DeleteVpcRouteInvoker(request *model.DeleteVpcRouteRequest) *DeleteVpcRouteInvoker {
	requestDef := GenReqDefForDeleteVpcRoute()
	return &DeleteVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVpcTag 删除VPC资源标签
//
// 删除指定VPC资源实例的标签信息
// 该接口为幂等接口：删除的key不存在报404，Key不能为空或者空字符串
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) DeleteVpcTag(request *model.DeleteVpcTagRequest) (*model.DeleteVpcTagResponse, error) {
	requestDef := GenReqDefForDeleteVpcTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVpcTagResponse), nil
	}
}

// DeleteVpcTagInvoker 删除VPC资源标签
func (c *VpcClient) DeleteVpcTagInvoker(request *model.DeleteVpcTagRequest) *DeleteVpcTagInvoker {
	requestDef := GenReqDefForDeleteVpcTag()
	return &DeleteVpcTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcRoutes 查询VPC路由列表
//
// 查询提交请求的租户的所有路由列表，并根据过滤条件进行过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListVpcRoutes(request *model.ListVpcRoutesRequest) (*model.ListVpcRoutesResponse, error) {
	requestDef := GenReqDefForListVpcRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcRoutesResponse), nil
	}
}

// ListVpcRoutesInvoker 查询VPC路由列表
func (c *VpcClient) ListVpcRoutesInvoker(request *model.ListVpcRoutesRequest) *ListVpcRoutesInvoker {
	requestDef := GenReqDefForListVpcRoutes()
	return &ListVpcRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcTags 查询VPC项目标签
//
// 查询租户在指定区域和实例类型的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListVpcTags(request *model.ListVpcTagsRequest) (*model.ListVpcTagsResponse, error) {
	requestDef := GenReqDefForListVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcTagsResponse), nil
	}
}

// ListVpcTagsInvoker 查询VPC项目标签
func (c *VpcClient) ListVpcTagsInvoker(request *model.ListVpcTagsRequest) *ListVpcTagsInvoker {
	requestDef := GenReqDefForListVpcTags()
	return &ListVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcs 查询VPC列表
//
// 查询虚拟私有云列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListVpcs(request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	requestDef := GenReqDefForListVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsResponse), nil
	}
}

// ListVpcsInvoker 查询VPC列表
func (c *VpcClient) ListVpcsInvoker(request *model.ListVpcsRequest) *ListVpcsInvoker {
	requestDef := GenReqDefForListVpcs()
	return &ListVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVpcsByTags 查询VPC资源实例
//
// 使用标签过滤实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ListVpcsByTags(request *model.ListVpcsByTagsRequest) (*model.ListVpcsByTagsResponse, error) {
	requestDef := GenReqDefForListVpcsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVpcsByTagsResponse), nil
	}
}

// ListVpcsByTagsInvoker 查询VPC资源实例
func (c *VpcClient) ListVpcsByTagsInvoker(request *model.ListVpcsByTagsRequest) *ListVpcsByTagsInvoker {
	requestDef := GenReqDefForListVpcsByTags()
	return &ListVpcsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpc 查询VPC
//
// 查询虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowVpc(request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	requestDef := GenReqDefForShowVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcResponse), nil
	}
}

// ShowVpcInvoker 查询VPC
func (c *VpcClient) ShowVpcInvoker(request *model.ShowVpcRequest) *ShowVpcInvoker {
	requestDef := GenReqDefForShowVpc()
	return &ShowVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpcRoute 查询VPC路由
//
// 查询路由详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowVpcRoute(request *model.ShowVpcRouteRequest) (*model.ShowVpcRouteResponse, error) {
	requestDef := GenReqDefForShowVpcRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcRouteResponse), nil
	}
}

// ShowVpcRouteInvoker 查询VPC路由
func (c *VpcClient) ShowVpcRouteInvoker(request *model.ShowVpcRouteRequest) *ShowVpcRouteInvoker {
	requestDef := GenReqDefForShowVpcRoute()
	return &ShowVpcRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVpcTags 查询VPC资源标签
//
// 查询指定VPC实例的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) ShowVpcTags(request *model.ShowVpcTagsRequest) (*model.ShowVpcTagsResponse, error) {
	requestDef := GenReqDefForShowVpcTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcTagsResponse), nil
	}
}

// ShowVpcTagsInvoker 查询VPC资源标签
func (c *VpcClient) ShowVpcTagsInvoker(request *model.ShowVpcTagsRequest) *ShowVpcTagsInvoker {
	requestDef := GenReqDefForShowVpcTags()
	return &ShowVpcTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVpc 更新VPC
//
// 更新虚拟私有云。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcClient) UpdateVpc(request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	requestDef := GenReqDefForUpdateVpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcResponse), nil
	}
}

// UpdateVpcInvoker 更新VPC
func (c *VpcClient) UpdateVpcInvoker(request *model.UpdateVpcRequest) *UpdateVpcInvoker {
	requestDef := GenReqDefForUpdateVpc()
	return &UpdateVpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
