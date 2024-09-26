package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpcep/v1/model"
)

type VpcepClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewVpcepClient(hcClient *httpclient.HcHttpClient) *VpcepClient {
	return &VpcepClient{HcClient: hcClient}
}

func VpcepClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AcceptOrRejectEndpoint 接受或拒绝终端节点的连接
//
// 接受或者拒绝终端节点连接到当前的终端节点服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) AcceptOrRejectEndpoint(request *model.AcceptOrRejectEndpointRequest) (*model.AcceptOrRejectEndpointResponse, error) {
	requestDef := GenReqDefForAcceptOrRejectEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptOrRejectEndpointResponse), nil
	}
}

// AcceptOrRejectEndpointInvoker 接受或拒绝终端节点的连接
func (c *VpcepClient) AcceptOrRejectEndpointInvoker(request *model.AcceptOrRejectEndpointRequest) *AcceptOrRejectEndpointInvoker {
	requestDef := GenReqDefForAcceptOrRejectEndpoint()
	return &AcceptOrRejectEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddOrRemoveServicePermissions 批量添加或移除终端节点服务的白名单
//
// 批量添加或移除当前用户下终端节点服务的白名单。
// 说明
// 本账号默认在自身用户的终端节点服务的白名单中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) AddOrRemoveServicePermissions(request *model.AddOrRemoveServicePermissionsRequest) (*model.AddOrRemoveServicePermissionsResponse, error) {
	requestDef := GenReqDefForAddOrRemoveServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddOrRemoveServicePermissionsResponse), nil
	}
}

// AddOrRemoveServicePermissionsInvoker 批量添加或移除终端节点服务的白名单
func (c *VpcepClient) AddOrRemoveServicePermissionsInvoker(request *model.AddOrRemoveServicePermissionsRequest) *AddOrRemoveServicePermissionsInvoker {
	requestDef := GenReqDefForAddOrRemoveServicePermissions()
	return &AddOrRemoveServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddEndpointServicePermissions 批量添加终端节点服务的白名单
//
// 批量添加当前用户下终端节点服务的白名单，支持添加描述信息。
// 说明
// 本账号默认在自身用户的终端节点服务的白名单中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) BatchAddEndpointServicePermissions(request *model.BatchAddEndpointServicePermissionsRequest) (*model.BatchAddEndpointServicePermissionsResponse, error) {
	requestDef := GenReqDefForBatchAddEndpointServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddEndpointServicePermissionsResponse), nil
	}
}

// BatchAddEndpointServicePermissionsInvoker 批量添加终端节点服务的白名单
func (c *VpcepClient) BatchAddEndpointServicePermissionsInvoker(request *model.BatchAddEndpointServicePermissionsRequest) *BatchAddEndpointServicePermissionsInvoker {
	requestDef := GenReqDefForBatchAddEndpointServicePermissions()
	return &BatchAddEndpointServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveEndpointServicePermissions 批量删除终端节点服务的白名单
//
// 批量删除当前用户下终端节点服务的白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) BatchRemoveEndpointServicePermissions(request *model.BatchRemoveEndpointServicePermissionsRequest) (*model.BatchRemoveEndpointServicePermissionsResponse, error) {
	requestDef := GenReqDefForBatchRemoveEndpointServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveEndpointServicePermissionsResponse), nil
	}
}

// BatchRemoveEndpointServicePermissionsInvoker 批量删除终端节点服务的白名单
func (c *VpcepClient) BatchRemoveEndpointServicePermissionsInvoker(request *model.BatchRemoveEndpointServicePermissionsRequest) *BatchRemoveEndpointServicePermissionsInvoker {
	requestDef := GenReqDefForBatchRemoveEndpointServicePermissions()
	return &BatchRemoveEndpointServicePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpoint 创建终端节点
//
// 创建终端节点，以便访问终端节点服务。
//
// 该接口为异步接口，调用成功会返回200状态码，说明请求已正常下发。
// 通常创建终端节点需要1~2分钟，可以通过[查询终端节点详情](https://support.huaweicloud.com/api-vpcep/ListEndpointInfoDetails.html)查看创建结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

// CreateEndpointInvoker 创建终端节点
func (c *VpcepClient) CreateEndpointInvoker(request *model.CreateEndpointRequest) *CreateEndpointInvoker {
	requestDef := GenReqDefForCreateEndpoint()
	return &CreateEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEndpointService 创建终端节点服务
//
// 创建终端节点服务，允许其他用户创建终端节点连接您创建的终端节点服务，
// 使用您所提供的服务。
//
// 该接口为异步接口，调用成功会返回200状态码，说明请求已正常下发。
// 通常创建终端节点服务需要1~2分钟，可以通过[查询终端节点服务详情](https://support.huaweicloud.com/api-vpcep/ListServiceDetails.html)查看创建结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) CreateEndpointService(request *model.CreateEndpointServiceRequest) (*model.CreateEndpointServiceResponse, error) {
	requestDef := GenReqDefForCreateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointServiceResponse), nil
	}
}

// CreateEndpointServiceInvoker 创建终端节点服务
func (c *VpcepClient) CreateEndpointServiceInvoker(request *model.CreateEndpointServiceRequest) *CreateEndpointServiceInvoker {
	requestDef := GenReqDefForCreateEndpointService()
	return &CreateEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpoint 删除终端节点
//
// 删除终端节点。
//
// 该接口为异步接口，调用成功会返回200状态码，说明请求已正常下发。
// 通常删除终端节点需要1~2分钟，可以通过[查询终端节点详情](https://support.huaweicloud.com/api-vpcep/ListEndpointInfoDetails.html)查看删除结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

// DeleteEndpointInvoker 删除终端节点
func (c *VpcepClient) DeleteEndpointInvoker(request *model.DeleteEndpointRequest) *DeleteEndpointInvoker {
	requestDef := GenReqDefForDeleteEndpoint()
	return &DeleteEndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpointPolicy 删除网关型终端节点策略（待下线）
//
// 删除网关型终端节点策略，该接口待下线，不建议使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) DeleteEndpointPolicy(request *model.DeleteEndpointPolicyRequest) (*model.DeleteEndpointPolicyResponse, error) {
	requestDef := GenReqDefForDeleteEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointPolicyResponse), nil
	}
}

// DeleteEndpointPolicyInvoker 删除网关型终端节点策略（待下线）
func (c *VpcepClient) DeleteEndpointPolicyInvoker(request *model.DeleteEndpointPolicyRequest) *DeleteEndpointPolicyInvoker {
	requestDef := GenReqDefForDeleteEndpointPolicy()
	return &DeleteEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEndpointService 删除终端节点服务
//
// 删除终端节点服务。
//
// 该接口为异步接口，调用成功会返回200状态码，说明请求已正常下发。通常删除终端节点服务需要1~2分钟，可以通过[查询终端节点服务详情](https://support.huaweicloud.com/api-vpcep/ListServiceDetails.html)查看删除结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) DeleteEndpointService(request *model.DeleteEndpointServiceRequest) (*model.DeleteEndpointServiceResponse, error) {
	requestDef := GenReqDefForDeleteEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointServiceResponse), nil
	}
}

// DeleteEndpointServiceInvoker 删除终端节点服务
func (c *VpcepClient) DeleteEndpointServiceInvoker(request *model.DeleteEndpointServiceRequest) *DeleteEndpointServiceInvoker {
	requestDef := GenReqDefForDeleteEndpointService()
	return &DeleteEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointInfoDetails 查询终端节点详情
//
// 查询终端节点的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListEndpointInfoDetails(request *model.ListEndpointInfoDetailsRequest) (*model.ListEndpointInfoDetailsResponse, error) {
	requestDef := GenReqDefForListEndpointInfoDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointInfoDetailsResponse), nil
	}
}

// ListEndpointInfoDetailsInvoker 查询终端节点详情
func (c *VpcepClient) ListEndpointInfoDetailsInvoker(request *model.ListEndpointInfoDetailsRequest) *ListEndpointInfoDetailsInvoker {
	requestDef := GenReqDefForListEndpointInfoDetails()
	return &ListEndpointInfoDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpointService 查询终端节点服务列表
//
// 查询当前用户下的终端节点服务的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListEndpointService(request *model.ListEndpointServiceRequest) (*model.ListEndpointServiceResponse, error) {
	requestDef := GenReqDefForListEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointServiceResponse), nil
	}
}

// ListEndpointServiceInvoker 查询终端节点服务列表
func (c *VpcepClient) ListEndpointServiceInvoker(request *model.ListEndpointServiceRequest) *ListEndpointServiceInvoker {
	requestDef := GenReqDefForListEndpointService()
	return &ListEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEndpoints 查询终端节点列表
//
// 查询当前用户下的终端节点的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

// ListEndpointsInvoker 查询终端节点列表
func (c *VpcepClient) ListEndpointsInvoker(request *model.ListEndpointsRequest) *ListEndpointsInvoker {
	requestDef := GenReqDefForListEndpoints()
	return &ListEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotaDetails 查询配额
//
// 查询用户的资源配额，包括终端节点服务和终端节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

// ListQuotaDetailsInvoker 查询配额
func (c *VpcepClient) ListQuotaDetailsInvoker(request *model.ListQuotaDetailsRequest) *ListQuotaDetailsInvoker {
	requestDef := GenReqDefForListQuotaDetails()
	return &ListQuotaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceConnections 查询连接终端节点服务的连接列表
//
// 查询连接当前用户下的某一个终端节点服务的连接列表。marker_id是连接的唯一标识。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListServiceConnections(request *model.ListServiceConnectionsRequest) (*model.ListServiceConnectionsResponse, error) {
	requestDef := GenReqDefForListServiceConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceConnectionsResponse), nil
	}
}

// ListServiceConnectionsInvoker 查询连接终端节点服务的连接列表
func (c *VpcepClient) ListServiceConnectionsInvoker(request *model.ListServiceConnectionsRequest) *ListServiceConnectionsInvoker {
	requestDef := GenReqDefForListServiceConnections()
	return &ListServiceConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceDescribeDetails 查询终端节点服务概要
//
// 查询终端节点服务的概要信息， 此接口是供创建终端节点的用户来查询需要连接的终端节点服务信息。 此接口既可以方便其他用户查询到您的终端节点服务概要信息, 又可以避免您的终端节点服务的细节信息暴露给其他用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListServiceDescribeDetails(request *model.ListServiceDescribeDetailsRequest) (*model.ListServiceDescribeDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDescribeDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDescribeDetailsResponse), nil
	}
}

// ListServiceDescribeDetailsInvoker 查询终端节点服务概要
func (c *VpcepClient) ListServiceDescribeDetailsInvoker(request *model.ListServiceDescribeDetailsRequest) *ListServiceDescribeDetailsInvoker {
	requestDef := GenReqDefForListServiceDescribeDetails()
	return &ListServiceDescribeDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceDetails 查询终端节点服务详情
//
// 查询终端节点服务的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListServiceDetails(request *model.ListServiceDetailsRequest) (*model.ListServiceDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDetailsResponse), nil
	}
}

// ListServiceDetailsInvoker 查询终端节点服务详情
func (c *VpcepClient) ListServiceDetailsInvoker(request *model.ListServiceDetailsRequest) *ListServiceDetailsInvoker {
	requestDef := GenReqDefForListServiceDetails()
	return &ListServiceDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServicePermissionsDetails 查询终端节点服务的白名单列表
//
// 查询当前用户下终端节点服务的白名单列表。
// 说明
// 本账号默认在当前用户下终端节点服务的白名单中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListServicePermissionsDetails(request *model.ListServicePermissionsDetailsRequest) (*model.ListServicePermissionsDetailsResponse, error) {
	requestDef := GenReqDefForListServicePermissionsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePermissionsDetailsResponse), nil
	}
}

// ListServicePermissionsDetailsInvoker 查询终端节点服务的白名单列表
func (c *VpcepClient) ListServicePermissionsDetailsInvoker(request *model.ListServicePermissionsDetailsRequest) *ListServicePermissionsDetailsInvoker {
	requestDef := GenReqDefForListServicePermissionsDetails()
	return &ListServicePermissionsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServicePublicDetails 查询公共终端节点服务列表
//
// 查询公共终端节点服务的列表，公共终端节点服务是所有用户可见且可连接的终端节点服务，
// 由运维人员创建，用户可直接使用，但无权创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListServicePublicDetails(request *model.ListServicePublicDetailsRequest) (*model.ListServicePublicDetailsResponse, error) {
	requestDef := GenReqDefForListServicePublicDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePublicDetailsResponse), nil
	}
}

// ListServicePublicDetailsInvoker 查询公共终端节点服务列表
func (c *VpcepClient) ListServicePublicDetailsInvoker(request *model.ListServicePublicDetailsRequest) *ListServicePublicDetailsInvoker {
	requestDef := GenReqDefForListServicePublicDetails()
	return &ListServicePublicDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecifiedVersionDetails 查询指定VPC终端节点接口版本信息
//
// 查询指定VPC终端节点接口版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListSpecifiedVersionDetails(request *model.ListSpecifiedVersionDetailsRequest) (*model.ListSpecifiedVersionDetailsResponse, error) {
	requestDef := GenReqDefForListSpecifiedVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecifiedVersionDetailsResponse), nil
	}
}

// ListSpecifiedVersionDetailsInvoker 查询指定VPC终端节点接口版本信息
func (c *VpcepClient) ListSpecifiedVersionDetailsInvoker(request *model.ListSpecifiedVersionDetailsRequest) *ListSpecifiedVersionDetailsInvoker {
	requestDef := GenReqDefForListSpecifiedVersionDetails()
	return &ListSpecifiedVersionDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVersionDetails 查询VPC终端节点接口版本列表
//
// 查询VPC终端节点接口版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListVersionDetails(request *model.ListVersionDetailsRequest) (*model.ListVersionDetailsResponse, error) {
	requestDef := GenReqDefForListVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionDetailsResponse), nil
	}
}

// ListVersionDetailsInvoker 查询VPC终端节点接口版本列表
func (c *VpcepClient) ListVersionDetailsInvoker(request *model.ListVersionDetailsRequest) *ListVersionDetailsInvoker {
	requestDef := GenReqDefForListVersionDetails()
	return &ListVersionDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointConnectionsDesc 更新终端节点连接描述
//
// 更新终端节点服务连接的终端节点的描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointConnectionsDesc(request *model.UpdateEndpointConnectionsDescRequest) (*model.UpdateEndpointConnectionsDescResponse, error) {
	requestDef := GenReqDefForUpdateEndpointConnectionsDesc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointConnectionsDescResponse), nil
	}
}

// UpdateEndpointConnectionsDescInvoker 更新终端节点连接描述
func (c *VpcepClient) UpdateEndpointConnectionsDescInvoker(request *model.UpdateEndpointConnectionsDescRequest) *UpdateEndpointConnectionsDescInvoker {
	requestDef := GenReqDefForUpdateEndpointConnectionsDesc()
	return &UpdateEndpointConnectionsDescInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointPolicy 修改终端节点策略
//
// 修改终端节点策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointPolicy(request *model.UpdateEndpointPolicyRequest) (*model.UpdateEndpointPolicyResponse, error) {
	requestDef := GenReqDefForUpdateEndpointPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointPolicyResponse), nil
	}
}

// UpdateEndpointPolicyInvoker 修改终端节点策略
func (c *VpcepClient) UpdateEndpointPolicyInvoker(request *model.UpdateEndpointPolicyRequest) *UpdateEndpointPolicyInvoker {
	requestDef := GenReqDefForUpdateEndpointPolicy()
	return &UpdateEndpointPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointRoutetable 修改终端节点的路由表
//
// 修改终端节点的路由表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointRoutetable(request *model.UpdateEndpointRoutetableRequest) (*model.UpdateEndpointRoutetableResponse, error) {
	requestDef := GenReqDefForUpdateEndpointRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointRoutetableResponse), nil
	}
}

// UpdateEndpointRoutetableInvoker 修改终端节点的路由表
func (c *VpcepClient) UpdateEndpointRoutetableInvoker(request *model.UpdateEndpointRoutetableRequest) *UpdateEndpointRoutetableInvoker {
	requestDef := GenReqDefForUpdateEndpointRoutetable()
	return &UpdateEndpointRoutetableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointService 修改终端节点服务
//
// 修改终端节点服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointService(request *model.UpdateEndpointServiceRequest) (*model.UpdateEndpointServiceResponse, error) {
	requestDef := GenReqDefForUpdateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServiceResponse), nil
	}
}

// UpdateEndpointServiceInvoker 修改终端节点服务
func (c *VpcepClient) UpdateEndpointServiceInvoker(request *model.UpdateEndpointServiceRequest) *UpdateEndpointServiceInvoker {
	requestDef := GenReqDefForUpdateEndpointService()
	return &UpdateEndpointServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointServiceName 修改终端节点服务名称
//
// 修改终端节点服务名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointServiceName(request *model.UpdateEndpointServiceNameRequest) (*model.UpdateEndpointServiceNameResponse, error) {
	requestDef := GenReqDefForUpdateEndpointServiceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServiceNameResponse), nil
	}
}

// UpdateEndpointServiceNameInvoker 修改终端节点服务名称
func (c *VpcepClient) UpdateEndpointServiceNameInvoker(request *model.UpdateEndpointServiceNameRequest) *UpdateEndpointServiceNameInvoker {
	requestDef := GenReqDefForUpdateEndpointServiceName()
	return &UpdateEndpointServiceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointServicePermissionDesc 更新终端节点服务白名单描述
//
// 更新当前用户下终端节点服务白名单的描述信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointServicePermissionDesc(request *model.UpdateEndpointServicePermissionDescRequest) (*model.UpdateEndpointServicePermissionDescResponse, error) {
	requestDef := GenReqDefForUpdateEndpointServicePermissionDesc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServicePermissionDescResponse), nil
	}
}

// UpdateEndpointServicePermissionDescInvoker 更新终端节点服务白名单描述
func (c *VpcepClient) UpdateEndpointServicePermissionDescInvoker(request *model.UpdateEndpointServicePermissionDescRequest) *UpdateEndpointServicePermissionDescInvoker {
	requestDef := GenReqDefForUpdateEndpointServicePermissionDesc()
	return &UpdateEndpointServicePermissionDescInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEndpointWhite 更新终端节点
//
// 更新或删除允许访问终端节点的白名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) UpdateEndpointWhite(request *model.UpdateEndpointWhiteRequest) (*model.UpdateEndpointWhiteResponse, error) {
	requestDef := GenReqDefForUpdateEndpointWhite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointWhiteResponse), nil
	}
}

// UpdateEndpointWhiteInvoker 更新终端节点
func (c *VpcepClient) UpdateEndpointWhiteInvoker(request *model.UpdateEndpointWhiteRequest) *UpdateEndpointWhiteInvoker {
	requestDef := GenReqDefForUpdateEndpointWhite()
	return &UpdateEndpointWhiteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddOrRemoveResourceInstance 批量添加或删除资源标签接口
//
// 为指定Endpoint Service或Endpoint批量添加或删除标签。
//   - 一个资源上最多有10个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) BatchAddOrRemoveResourceInstance(request *model.BatchAddOrRemoveResourceInstanceRequest) (*model.BatchAddOrRemoveResourceInstanceResponse, error) {
	requestDef := GenReqDefForBatchAddOrRemoveResourceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddOrRemoveResourceInstanceResponse), nil
	}
}

// BatchAddOrRemoveResourceInstanceInvoker 批量添加或删除资源标签接口
func (c *VpcepClient) BatchAddOrRemoveResourceInstanceInvoker(request *model.BatchAddOrRemoveResourceInstanceRequest) *BatchAddOrRemoveResourceInstanceInvoker {
	requestDef := GenReqDefForBatchAddOrRemoveResourceInstance()
	return &BatchAddOrRemoveResourceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryProjectResourceTags 查询租户资源标签接口
//
// 根据租户ID和资源类型，获取租户下资源的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListQueryProjectResourceTags(request *model.ListQueryProjectResourceTagsRequest) (*model.ListQueryProjectResourceTagsResponse, error) {
	requestDef := GenReqDefForListQueryProjectResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryProjectResourceTagsResponse), nil
	}
}

// ListQueryProjectResourceTagsInvoker 查询租户资源标签接口
func (c *VpcepClient) ListQueryProjectResourceTagsInvoker(request *model.ListQueryProjectResourceTagsRequest) *ListQueryProjectResourceTagsInvoker {
	requestDef := GenReqDefForListQueryProjectResourceTags()
	return &ListQueryProjectResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceInstances 查询资源实例接口
//
// 使用标签过滤查询租户下资源的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VpcepClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}

// ListResourceInstancesInvoker 查询资源实例接口
func (c *VpcepClient) ListResourceInstancesInvoker(request *model.ListResourceInstancesRequest) *ListResourceInstancesInvoker {
	requestDef := GenReqDefForListResourceInstances()
	return &ListResourceInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
