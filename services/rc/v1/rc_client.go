package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rc/v1/model"
)

type RcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRcClient(hcClient *httpclient.HcHttpClient) *RcClient {
	return &RcClient{HcClient: hcClient}
}

func RcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AddResourcesToGroup 将资源添加到资源组
//
// 将一个或多个资源添加到资源组，需要当前用户有resourcecenter:group:addResource权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) AddResourcesToGroup(request *model.AddResourcesToGroupRequest) (*model.AddResourcesToGroupResponse, error) {
	requestDef := GenReqDefForAddResourcesToGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddResourcesToGroupResponse), nil
	}
}

// AddResourcesToGroupInvoker 将资源添加到资源组
func (c *RcClient) AddResourcesToGroupInvoker(request *model.AddResourcesToGroupRequest) *AddResourcesToGroupInvoker {
	requestDef := GenReqDefForAddResourcesToGroup()
	return &AddResourcesToGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceGroup 创建资源分组
//
// 创建一个资源分组，需要当前用户有resourcecenter:group:create权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) CreateResourceGroup(request *model.CreateResourceGroupRequest) (*model.CreateResourceGroupResponse, error) {
	requestDef := GenReqDefForCreateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceGroupResponse), nil
	}
}

// CreateResourceGroupInvoker 创建资源分组
func (c *RcClient) CreateResourceGroupInvoker(request *model.CreateResourceGroupRequest) *CreateResourceGroupInvoker {
	requestDef := GenReqDefForCreateResourceGroup()
	return &CreateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceGroup 删除一个资源分组
//
// 删除一个资源分组，需要当前用户有resourcecenter:group:delete权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) DeleteResourceGroup(request *model.DeleteResourceGroupRequest) (*model.DeleteResourceGroupResponse, error) {
	requestDef := GenReqDefForDeleteResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceGroupResponse), nil
	}
}

// DeleteResourceGroupInvoker 删除一个资源分组
func (c *RcClient) DeleteResourceGroupInvoker(request *model.DeleteResourceGroupRequest) *DeleteResourceGroupInvoker {
	requestDef := GenReqDefForDeleteResourceGroup()
	return &DeleteResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceGroups 查询资源分组列表
//
// 查询资源分组列表，需要当前用户有resourcecenter:group:list权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ListResourceGroups(request *model.ListResourceGroupsRequest) (*model.ListResourceGroupsResponse, error) {
	requestDef := GenReqDefForListResourceGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceGroupsResponse), nil
	}
}

// ListResourceGroupsInvoker 查询资源分组列表
func (c *RcClient) ListResourceGroupsInvoker(request *model.ListResourceGroupsRequest) *ListResourceGroupsInvoker {
	requestDef := GenReqDefForListResourceGroups()
	return &ListResourceGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveResourceFromGroup 从资源组中移除资源
//
// 从资源组中移除一个资源，需要当前用户有resourcecenter:group:removeResource权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) RemoveResourceFromGroup(request *model.RemoveResourceFromGroupRequest) (*model.RemoveResourceFromGroupResponse, error) {
	requestDef := GenReqDefForRemoveResourceFromGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveResourceFromGroupResponse), nil
	}
}

// RemoveResourceFromGroupInvoker 从资源组中移除资源
func (c *RcClient) RemoveResourceFromGroupInvoker(request *model.RemoveResourceFromGroupRequest) *RemoveResourceFromGroupInvoker {
	requestDef := GenReqDefForRemoveResourceFromGroup()
	return &RemoveResourceFromGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceGroup 查询一个资源分组
//
// 查询一个资源分组，需要当前用户有resourcecenter:group:get权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ShowResourceGroup(request *model.ShowResourceGroupRequest) (*model.ShowResourceGroupResponse, error) {
	requestDef := GenReqDefForShowResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceGroupResponse), nil
	}
}

// ShowResourceGroupInvoker 查询一个资源分组
func (c *RcClient) ShowResourceGroupInvoker(request *model.ShowResourceGroupRequest) *ShowResourceGroupInvoker {
	requestDef := GenReqDefForShowResourceGroup()
	return &ShowResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceGroup 修改一个资源分组
//
// 修改一个资源分组，需要当前用户有resourcecenter:group:update权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) UpdateResourceGroup(request *model.UpdateResourceGroupRequest) (*model.UpdateResourceGroupResponse, error) {
	requestDef := GenReqDefForUpdateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceGroupResponse), nil
	}
}

// UpdateResourceGroupInvoker 修改一个资源分组
func (c *RcClient) UpdateResourceGroupInvoker(request *model.UpdateResourceGroupRequest) *UpdateResourceGroupInvoker {
	requestDef := GenReqDefForUpdateResourceGroup()
	return &UpdateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceRelations 列举资源关系
//
// 指定资源ID，查询该资源与其他资源的关联关系，需要当前用户有resourcecenter::listResourceRelation权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ShowResourceRelations(request *model.ShowResourceRelationsRequest) (*model.ShowResourceRelationsResponse, error) {
	requestDef := GenReqDefForShowResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsResponse), nil
	}
}

// ShowResourceRelationsInvoker 列举资源关系
func (c *RcClient) ShowResourceRelationsInvoker(request *model.ShowResourceRelationsRequest) *ShowResourceRelationsInvoker {
	requestDef := GenReqDefForShowResourceRelations()
	return &ShowResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectAllResourcesSummary 列举资源概要
//
// 查询当前帐号的资源概览，需要当前用户有rc::listResourceSummary权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) CollectAllResourcesSummary(request *model.CollectAllResourcesSummaryRequest) (*model.CollectAllResourcesSummaryResponse, error) {
	requestDef := GenReqDefForCollectAllResourcesSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

// CollectAllResourcesSummaryInvoker 列举资源概要
func (c *RcClient) CollectAllResourcesSummaryInvoker(request *model.CollectAllResourcesSummaryRequest) *CollectAllResourcesSummaryInvoker {
	requestDef := GenReqDefForCollectAllResourcesSummary()
	return &CollectAllResourcesSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountAllResources 查询资源数量
//
// 查询资源数量，需要当前用户有resourcecenter::getResourceCount权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) CountAllResources(request *model.CountAllResourcesRequest) (*model.CountAllResourcesResponse, error) {
	requestDef := GenReqDefForCountAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountAllResourcesResponse), nil
	}
}

// CountAllResourcesInvoker 查询资源数量
func (c *RcClient) CountAllResourcesInvoker(request *model.CountAllResourcesRequest) *CountAllResourcesInvoker {
	requestDef := GenReqDefForCountAllResources()
	return &CountAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllProviders 列举所有已对接的云服务
//
// 查询所有已对接RC的云服务、资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ListAllProviders(request *model.ListAllProvidersRequest) (*model.ListAllProvidersResponse, error) {
	requestDef := GenReqDefForListAllProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllProvidersResponse), nil
	}
}

// ListAllProvidersInvoker 列举所有已对接的云服务
func (c *RcClient) ListAllProvidersInvoker(request *model.ListAllProvidersRequest) *ListAllProvidersInvoker {
	requestDef := GenReqDefForListAllProviders()
	return &ListAllProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllResources 列举所有资源
//
// 返回当前用户下所有资源，需要当前用户有resourcecenter::listResource权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ListAllResources(request *model.ListAllResourcesRequest) (*model.ListAllResourcesResponse, error) {
	requestDef := GenReqDefForListAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllResourcesResponse), nil
	}
}

// ListAllResourcesInvoker 列举所有资源
func (c *RcClient) ListAllResourcesInvoker(request *model.ListAllResourcesRequest) *ListAllResourcesInvoker {
	requestDef := GenReqDefForListAllResources()
	return &ListAllResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTags 列举资源标签
//
// 查询当前帐号下所有资源的标签，需要当前用户有resourcecenter::listResourceTag权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

// ListAllTagsInvoker 列举资源标签
func (c *RcClient) ListAllTagsInvoker(request *model.ListAllTagsRequest) *ListAllTagsInvoker {
	requestDef := GenReqDefForListAllTags()
	return &ListAllTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 列举指定类型的资源
//
// 返回当前租户下特定资源类型的资源，需要当前用户有resourcecenter::listResourceByType权限。比如查询云服务器，对应的RC资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 列举指定类型的资源
func (c *RcClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceById 查询指定类型的单个资源
//
// 指定资源ID，返回该资源的详细信息，需要当前用户有resourcecenter::getResourceByType权限。比如查询云服务器，对应的RC资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ShowResourceById(request *model.ShowResourceByIdRequest) (*model.ShowResourceByIdResponse, error) {
	requestDef := GenReqDefForShowResourceById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceByIdResponse), nil
	}
}

// ShowResourceByIdInvoker 查询指定类型的单个资源
func (c *RcClient) ShowResourceByIdInvoker(request *model.ShowResourceByIdRequest) *ShowResourceByIdInvoker {
	requestDef := GenReqDefForShowResourceById()
	return &ShowResourceByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceDetail 查询单个资源
//
// 查询当前帐号下的单个资源，需要当前用户有resourcecenter::getResource权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RcClient) ShowResourceDetail(request *model.ShowResourceDetailRequest) (*model.ShowResourceDetailResponse, error) {
	requestDef := GenReqDefForShowResourceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceDetailResponse), nil
	}
}

// ShowResourceDetailInvoker 查询单个资源
func (c *RcClient) ShowResourceDetailInvoker(request *model.ShowResourceDetailRequest) *ShowResourceDetailInvoker {
	requestDef := GenReqDefForShowResourceDetail()
	return &ShowResourceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
