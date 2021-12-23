package v1

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/vpcep/v1/model"
)

type VpcepClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVpcepClient(hcClient *http_client.HcHttpClient) *VpcepClient {
	return &VpcepClient{HcClient: hcClient}
}

func VpcepClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//功能介绍 接受或者拒绝终端节点连接到当前的终端节点服务。
func (c *VpcepClient) AcceptOrRejectEndpoint(request *model.AcceptOrRejectEndpointRequest) (*model.AcceptOrRejectEndpointResponse, error) {
	requestDef := GenReqDefForAcceptOrRejectEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptOrRejectEndpointResponse), nil
	}
}

//功能介绍 批量添加或移除当前用户下终端节点服务的白名单。 说明 本帐号默认在自身用户的终端节点服务的白名单中。
func (c *VpcepClient) AddOrRemoveServicePermissions(request *model.AddOrRemoveServicePermissionsRequest) (*model.AddOrRemoveServicePermissionsResponse, error) {
	requestDef := GenReqDefForAddOrRemoveServicePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddOrRemoveServicePermissionsResponse), nil
	}
}

//功能介绍 创建终端节点，以便访问终端节点服务。
func (c *VpcepClient) CreateEndpoint(request *model.CreateEndpointRequest) (*model.CreateEndpointResponse, error) {
	requestDef := GenReqDefForCreateEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointResponse), nil
	}
}

//功能介绍 创建终端节点服务，允许其他用户创建终端节点连接您创建的终端节点服务，使用您所提供的服务。 说明 该接口为异步接口，调用成功会返回200状态码，说明请求已正常下发。通常创建终端节点服务需要1~2分钟，可以通过查询终端节点服务详情查看创建结果。
func (c *VpcepClient) CreateEndpointService(request *model.CreateEndpointServiceRequest) (*model.CreateEndpointServiceResponse, error) {
	requestDef := GenReqDefForCreateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEndpointServiceResponse), nil
	}
}

//功能介绍 删除终端节点。
func (c *VpcepClient) DeleteEndpoint(request *model.DeleteEndpointRequest) (*model.DeleteEndpointResponse, error) {
	requestDef := GenReqDefForDeleteEndpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointResponse), nil
	}
}

//功能介绍 删除终端节点服务。
func (c *VpcepClient) DeleteEndpointService(request *model.DeleteEndpointServiceRequest) (*model.DeleteEndpointServiceResponse, error) {
	requestDef := GenReqDefForDeleteEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEndpointServiceResponse), nil
	}
}

//功能介绍 查询终端节点的详细信息。
func (c *VpcepClient) ListEndpointInfoDetails(request *model.ListEndpointInfoDetailsRequest) (*model.ListEndpointInfoDetailsResponse, error) {
	requestDef := GenReqDefForListEndpointInfoDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointInfoDetailsResponse), nil
	}
}

//功能介绍 查询当前用户下的终端节点服务的列表。
func (c *VpcepClient) ListEndpointService(request *model.ListEndpointServiceRequest) (*model.ListEndpointServiceResponse, error) {
	requestDef := GenReqDefForListEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointServiceResponse), nil
	}
}

//功能介绍 查询当前用户下的终端节点的列表。
func (c *VpcepClient) ListEndpoints(request *model.ListEndpointsRequest) (*model.ListEndpointsResponse, error) {
	requestDef := GenReqDefForListEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEndpointsResponse), nil
	}
}

//功能介绍 查询用户的资源配额，包括终端节点服务和终端节点。
func (c *VpcepClient) ListQuotaDetails(request *model.ListQuotaDetailsRequest) (*model.ListQuotaDetailsResponse, error) {
	requestDef := GenReqDefForListQuotaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaDetailsResponse), nil
	}
}

//功能介绍 查询连接当前用户下的某一个终端节点服务的连接列表。marker_id是连接的唯一标识。
func (c *VpcepClient) ListServiceConnections(request *model.ListServiceConnectionsRequest) (*model.ListServiceConnectionsResponse, error) {
	requestDef := GenReqDefForListServiceConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceConnectionsResponse), nil
	}
}

//功能介绍 查询终端节点服务的概要信息，此接口是供创建终端节点的用户来查询需要连接的终端节点服务信息。此接口既可以方便其他用户查询到您的终端节点服务概要信息又可以避免您的终端节点服务的细节信息暴露给其他用户。
func (c *VpcepClient) ListServiceDescribeDetails(request *model.ListServiceDescribeDetailsRequest) (*model.ListServiceDescribeDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDescribeDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDescribeDetailsResponse), nil
	}
}

//功能介绍 查询终端节点服务的详细信息。
func (c *VpcepClient) ListServiceDetails(request *model.ListServiceDetailsRequest) (*model.ListServiceDetailsResponse, error) {
	requestDef := GenReqDefForListServiceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceDetailsResponse), nil
	}
}

//功能介绍 查询当前用户下终端节点服务的白名单列表。 说明 本帐号默认在当前用户下终端节点服务的白名单中。
func (c *VpcepClient) ListServicePermissionsDetails(request *model.ListServicePermissionsDetailsRequest) (*model.ListServicePermissionsDetailsResponse, error) {
	requestDef := GenReqDefForListServicePermissionsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePermissionsDetailsResponse), nil
	}
}

//功能介绍 查询公共终端节点服务的列表，公共终端节点服务是所有用户可见且可连接的终端节点服务，由运维人员创建，用户可直接使用，但无权创建。
func (c *VpcepClient) ListServicePublicDetails(request *model.ListServicePublicDetailsRequest) (*model.ListServicePublicDetailsResponse, error) {
	requestDef := GenReqDefForListServicePublicDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePublicDetailsResponse), nil
	}
}

//功能介绍 查询指定VPC终端节点接口版本信息。
func (c *VpcepClient) ListSpecifiedVersionDetails(request *model.ListSpecifiedVersionDetailsRequest) (*model.ListSpecifiedVersionDetailsResponse, error) {
	requestDef := GenReqDefForListSpecifiedVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecifiedVersionDetailsResponse), nil
	}
}

//功能介绍 查询VPC终端节点接口版本列表。
func (c *VpcepClient) ListVersionDetails(request *model.ListVersionDetailsRequest) (*model.ListVersionDetailsResponse, error) {
	requestDef := GenReqDefForListVersionDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVersionDetailsResponse), nil
	}
}

//功能介绍 修改终端节点路由表。
func (c *VpcepClient) UpdateEndpointRoutetable(request *model.UpdateEndpointRoutetableRequest) (*model.UpdateEndpointRoutetableResponse, error) {
	requestDef := GenReqDefForUpdateEndpointRoutetable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointRoutetableResponse), nil
	}
}

//功能介绍 修改终端节点服务。
func (c *VpcepClient) UpdateEndpointService(request *model.UpdateEndpointServiceRequest) (*model.UpdateEndpointServiceResponse, error) {
	requestDef := GenReqDefForUpdateEndpointService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointServiceResponse), nil
	}
}

//功能介绍 更新或删除允许访问终端节点的白名单。
func (c *VpcepClient) UpdateEndpointWhite(request *model.UpdateEndpointWhiteRequest) (*model.UpdateEndpointWhiteResponse, error) {
	requestDef := GenReqDefForUpdateEndpointWhite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEndpointWhiteResponse), nil
	}
}

//功能介绍 为指定Endpoint Service或Endpoint批量添加或删除标签。 ● 一个资源上最多有10个标签。
func (c *VpcepClient) BatchAddOrRemoveResourceInstance(request *model.BatchAddOrRemoveResourceInstanceRequest) (*model.BatchAddOrRemoveResourceInstanceResponse, error) {
	requestDef := GenReqDefForBatchAddOrRemoveResourceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddOrRemoveResourceInstanceResponse), nil
	}
}

//功能介绍 根据租户ID和资源类型，获取租户下资源的标签。
func (c *VpcepClient) ListQueryProjectResourceTags(request *model.ListQueryProjectResourceTagsRequest) (*model.ListQueryProjectResourceTagsResponse, error) {
	requestDef := GenReqDefForListQueryProjectResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryProjectResourceTagsResponse), nil
	}
}

//功能介绍 使用标签过滤查询租户下资源的实例。
func (c *VpcepClient) ListResourceInstances(request *model.ListResourceInstancesRequest) (*model.ListResourceInstancesResponse, error) {
	requestDef := GenReqDefForListResourceInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceInstancesResponse), nil
	}
}
