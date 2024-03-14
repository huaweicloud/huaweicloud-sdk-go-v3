package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ram/v1/model"
)

type RamClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewRamClient(hcClient *httpclient.HcHttpClient) *RamClient {
	return &RamClient{HcClient: hcClient}
}

func RamClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// AssociateResourceSharePermission 绑定或替换共享资源权限
//
// 为资源共享实例中包含的资源类型绑定或替换共享资源权限。 对于资源共享实例中的每一种资源类型，您可以设置唯一权限。仅当资源共享实例中当前没有该资源类型的资源时，您才能绑定新的共享资源权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) AssociateResourceSharePermission(request *model.AssociateResourceSharePermissionRequest) (*model.AssociateResourceSharePermissionResponse, error) {
	requestDef := GenReqDefForAssociateResourceSharePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateResourceSharePermissionResponse), nil
	}
}

// AssociateResourceSharePermissionInvoker 绑定或替换共享资源权限
func (c *RamClient) AssociateResourceSharePermissionInvoker(request *model.AssociateResourceSharePermissionRequest) *AssociateResourceSharePermissionInvoker {
	requestDef := GenReqDefForAssociateResourceSharePermission()
	return &AssociateResourceSharePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateResourceSharePermission 移除共享资源权限
//
// 移除资源共享实例绑定的共享资源权限。权限更改立即生效。只有当目前资源共享实例中没有绑定相关资源类型时，您才能从资源共享实例中移除共享资源权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) DisassociateResourceSharePermission(request *model.DisassociateResourceSharePermissionRequest) (*model.DisassociateResourceSharePermissionResponse, error) {
	requestDef := GenReqDefForDisassociateResourceSharePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateResourceSharePermissionResponse), nil
	}
}

// DisassociateResourceSharePermissionInvoker 移除共享资源权限
func (c *RamClient) DisassociateResourceSharePermissionInvoker(request *model.DisassociateResourceSharePermissionRequest) *DisassociateResourceSharePermissionInvoker {
	requestDef := GenReqDefForDisassociateResourceSharePermission()
	return &DisassociateResourceSharePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceSharePermissions 检索绑定的共享资源权限
//
// 检索资源共享实例关联的共享资源权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListResourceSharePermissions(request *model.ListResourceSharePermissionsRequest) (*model.ListResourceSharePermissionsResponse, error) {
	requestDef := GenReqDefForListResourceSharePermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceSharePermissionsResponse), nil
	}
}

// ListResourceSharePermissionsInvoker 检索绑定的共享资源权限
func (c *RamClient) ListResourceSharePermissionsInvoker(request *model.ListResourceSharePermissionsRequest) *ListResourceSharePermissionsInvoker {
	requestDef := GenReqDefForListResourceSharePermissions()
	return &ListResourceSharePermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuota 查询资源共享的配额
//
// 查询当前账号的资源共享配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListQuota(request *model.ListQuotaRequest) (*model.ListQuotaResponse, error) {
	requestDef := GenReqDefForListQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotaResponse), nil
	}
}

// ListQuotaInvoker 查询资源共享的配额
func (c *RamClient) ListQuotaInvoker(request *model.ListQuotaRequest) *ListQuotaInvoker {
	requestDef := GenReqDefForListQuota()
	return &ListQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTypes 检索云服务资源类型
//
// 查询已对接云服务的资源类型和区域等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListResourceTypes(request *model.ListResourceTypesRequest) (*model.ListResourceTypesResponse, error) {
	requestDef := GenReqDefForListResourceTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTypesResponse), nil
	}
}

// ListResourceTypesInvoker 检索云服务资源类型
func (c *RamClient) ListResourceTypesInvoker(request *model.ListResourceTypesRequest) *ListResourceTypesInvoker {
	requestDef := GenReqDefForListResourceTypes()
	return &ListResourceTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableOrganizationShare 关闭与组织共享
//
// 关闭与组织共享资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) DisableOrganizationShare(request *model.DisableOrganizationShareRequest) (*model.DisableOrganizationShareResponse, error) {
	requestDef := GenReqDefForDisableOrganizationShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableOrganizationShareResponse), nil
	}
}

// DisableOrganizationShareInvoker 关闭与组织共享
func (c *RamClient) DisableOrganizationShareInvoker(request *model.DisableOrganizationShareRequest) *DisableOrganizationShareInvoker {
	requestDef := GenReqDefForDisableOrganizationShare()
	return &DisableOrganizationShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableOrganizationShare 启用与组织共享
//
// 启用与组织共享资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) EnableOrganizationShare(request *model.EnableOrganizationShareRequest) (*model.EnableOrganizationShareResponse, error) {
	requestDef := GenReqDefForEnableOrganizationShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableOrganizationShareResponse), nil
	}
}

// EnableOrganizationShareInvoker 启用与组织共享
func (c *RamClient) EnableOrganizationShareInvoker(request *model.EnableOrganizationShareRequest) *EnableOrganizationShareInvoker {
	requestDef := GenReqDefForEnableOrganizationShare()
	return &EnableOrganizationShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrganizationShare 检索是否启用与组织共享
//
// 检索是否启用与组织共享资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ShowOrganizationShare(request *model.ShowOrganizationShareRequest) (*model.ShowOrganizationShareResponse, error) {
	requestDef := GenReqDefForShowOrganizationShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrganizationShareResponse), nil
	}
}

// ShowOrganizationShareInvoker 检索是否启用与组织共享
func (c *RamClient) ShowOrganizationShareInvoker(request *model.ShowOrganizationShareRequest) *ShowOrganizationShareInvoker {
	requestDef := GenReqDefForShowOrganizationShare()
	return &ShowOrganizationShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionVersions 获取权限的所有版本
//
// 获取权限的所有版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListPermissionVersions(request *model.ListPermissionVersionsRequest) (*model.ListPermissionVersionsResponse, error) {
	requestDef := GenReqDefForListPermissionVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionVersionsResponse), nil
	}
}

// ListPermissionVersionsInvoker 获取权限的所有版本
func (c *RamClient) ListPermissionVersionsInvoker(request *model.ListPermissionVersionsRequest) *ListPermissionVersionsInvoker {
	requestDef := GenReqDefForListPermissionVersions()
	return &ListPermissionVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissions 检索共享资源权限列表
//
// 检索指定资源类型的共享资源权限列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListPermissions(request *model.ListPermissionsRequest) (*model.ListPermissionsResponse, error) {
	requestDef := GenReqDefForListPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionsResponse), nil
	}
}

// ListPermissionsInvoker 检索共享资源权限列表
func (c *RamClient) ListPermissionsInvoker(request *model.ListPermissionsRequest) *ListPermissionsInvoker {
	requestDef := GenReqDefForListPermissions()
	return &ListPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPermission 检索资源共享权限内容
//
// 检索指定资源类型的共享资源指定版本的权限内容，如果不指定权限版本，则返回默认版本的权限内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ShowPermission(request *model.ShowPermissionRequest) (*model.ShowPermissionResponse, error) {
	requestDef := GenReqDefForShowPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPermissionResponse), nil
	}
}

// ShowPermissionInvoker 检索资源共享权限内容
func (c *RamClient) ShowPermissionInvoker(request *model.ShowPermissionRequest) *ShowPermissionInvoker {
	requestDef := GenReqDefForShowPermission()
	return &ShowPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchSharedPrincipals 检索资源使用者
//
// 检索共享资源的使用者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchSharedPrincipals(request *model.SearchSharedPrincipalsRequest) (*model.SearchSharedPrincipalsResponse, error) {
	requestDef := GenReqDefForSearchSharedPrincipals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchSharedPrincipalsResponse), nil
	}
}

// SearchSharedPrincipalsInvoker 检索资源使用者
func (c *RamClient) SearchSharedPrincipalsInvoker(request *model.SearchSharedPrincipalsRequest) *SearchSharedPrincipalsInvoker {
	requestDef := GenReqDefForSearchSharedPrincipals()
	return &SearchSharedPrincipalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchSharedResources 检索共享的资源
//
// 检索您共享的或共享给您的资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchSharedResources(request *model.SearchSharedResourcesRequest) (*model.SearchSharedResourcesResponse, error) {
	requestDef := GenReqDefForSearchSharedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchSharedResourcesResponse), nil
	}
}

// SearchSharedResourcesInvoker 检索共享的资源
func (c *RamClient) SearchSharedResourcesInvoker(request *model.SearchSharedResourcesRequest) *SearchSharedResourcesInvoker {
	requestDef := GenReqDefForSearchSharedResources()
	return &SearchSharedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceShare 创建资源共享实例
//
// 创建一个资源共享实例。您可以指定需要共享的资源列表，资源使用者列表，以及授予资源使用者的权限列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) CreateResourceShare(request *model.CreateResourceShareRequest) (*model.CreateResourceShareResponse, error) {
	requestDef := GenReqDefForCreateResourceShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceShareResponse), nil
	}
}

// CreateResourceShareInvoker 创建资源共享实例
func (c *RamClient) CreateResourceShareInvoker(request *model.CreateResourceShareRequest) *CreateResourceShareInvoker {
	requestDef := GenReqDefForCreateResourceShare()
	return &CreateResourceShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceShare 删除资源共享实例
//
// 删除指定的资源共享实例。此操作不会删除实体资源，仅停止向其他账号共享资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) DeleteResourceShare(request *model.DeleteResourceShareRequest) (*model.DeleteResourceShareResponse, error) {
	requestDef := GenReqDefForDeleteResourceShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceShareResponse), nil
	}
}

// DeleteResourceShareInvoker 删除资源共享实例
func (c *RamClient) DeleteResourceShareInvoker(request *model.DeleteResourceShareRequest) *DeleteResourceShareInvoker {
	requestDef := GenReqDefForDeleteResourceShare()
	return &DeleteResourceShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceShares 检索资源共享实例
//
// 检索您创建的或者共享给您的资源共享实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchResourceShares(request *model.SearchResourceSharesRequest) (*model.SearchResourceSharesResponse, error) {
	requestDef := GenReqDefForSearchResourceShares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceSharesResponse), nil
	}
}

// SearchResourceSharesInvoker 检索资源共享实例
func (c *RamClient) SearchResourceSharesInvoker(request *model.SearchResourceSharesRequest) *SearchResourceSharesInvoker {
	requestDef := GenReqDefForSearchResourceShares()
	return &SearchResourceSharesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceShare 更新资源共享实例
//
// 修改资源共享实例的特定属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) UpdateResourceShare(request *model.UpdateResourceShareRequest) (*model.UpdateResourceShareResponse, error) {
	requestDef := GenReqDefForUpdateResourceShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceShareResponse), nil
	}
}

// UpdateResourceShareInvoker 更新资源共享实例
func (c *RamClient) UpdateResourceShareInvoker(request *model.UpdateResourceShareRequest) *UpdateResourceShareInvoker {
	requestDef := GenReqDefForUpdateResourceShare()
	return &UpdateResourceShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateResourceShare 绑定资源使用者和共享资源
//
// 向资源共享实例绑定指定的资源使用者列表或共享资源列表。对于新增的共享资源，有权访问此资源共享实例的资源使用者获得该共享资源的访问权限。对于新增的资源使用者，获得对此资源共享实例中共享资源的访问权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) AssociateResourceShare(request *model.AssociateResourceShareRequest) (*model.AssociateResourceShareResponse, error) {
	requestDef := GenReqDefForAssociateResourceShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateResourceShareResponse), nil
	}
}

// AssociateResourceShareInvoker 绑定资源使用者和共享资源
func (c *RamClient) AssociateResourceShareInvoker(request *model.AssociateResourceShareRequest) *AssociateResourceShareInvoker {
	requestDef := GenReqDefForAssociateResourceShare()
	return &AssociateResourceShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateResourceShare 移除资源使用者和共享资源
//
// 将指定的资源使用者或共享资源从指定的资源共享实例中移除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) DisassociateResourceShare(request *model.DisassociateResourceShareRequest) (*model.DisassociateResourceShareResponse, error) {
	requestDef := GenReqDefForDisassociateResourceShare()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateResourceShareResponse), nil
	}
}

// DisassociateResourceShareInvoker 移除资源使用者和共享资源
func (c *RamClient) DisassociateResourceShareInvoker(request *model.DisassociateResourceShareRequest) *DisassociateResourceShareInvoker {
	requestDef := GenReqDefForDisassociateResourceShare()
	return &DisassociateResourceShareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceShareAssociations 检索绑定的资源使用者和共享资源
//
// 检索您拥有的资源共享实例中绑定的共享资源和资源使用者。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchResourceShareAssociations(request *model.SearchResourceShareAssociationsRequest) (*model.SearchResourceShareAssociationsResponse, error) {
	requestDef := GenReqDefForSearchResourceShareAssociations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceShareAssociationsResponse), nil
	}
}

// SearchResourceShareAssociationsInvoker 检索绑定的资源使用者和共享资源
func (c *RamClient) SearchResourceShareAssociationsInvoker(request *model.SearchResourceShareAssociationsRequest) *SearchResourceShareAssociationsInvoker {
	requestDef := GenReqDefForSearchResourceShareAssociations()
	return &SearchResourceShareAssociationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AcceptResourceShareInvitation 接受共享邀请
//
// 接受来自其他账号的资源共享邀请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) AcceptResourceShareInvitation(request *model.AcceptResourceShareInvitationRequest) (*model.AcceptResourceShareInvitationResponse, error) {
	requestDef := GenReqDefForAcceptResourceShareInvitation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptResourceShareInvitationResponse), nil
	}
}

// AcceptResourceShareInvitationInvoker 接受共享邀请
func (c *RamClient) AcceptResourceShareInvitationInvoker(request *model.AcceptResourceShareInvitationRequest) *AcceptResourceShareInvitationInvoker {
	requestDef := GenReqDefForAcceptResourceShareInvitation()
	return &AcceptResourceShareInvitationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RejectResourceShareInvitation 拒绝共享邀请
//
// 拒绝来自其他账号的资源共享邀请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) RejectResourceShareInvitation(request *model.RejectResourceShareInvitationRequest) (*model.RejectResourceShareInvitationResponse, error) {
	requestDef := GenReqDefForRejectResourceShareInvitation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RejectResourceShareInvitationResponse), nil
	}
}

// RejectResourceShareInvitationInvoker 拒绝共享邀请
func (c *RamClient) RejectResourceShareInvitationInvoker(request *model.RejectResourceShareInvitationRequest) *RejectResourceShareInvitationInvoker {
	requestDef := GenReqDefForRejectResourceShareInvitation()
	return &RejectResourceShareInvitationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceShareInvitation 检索共享邀请
//
// 通过条件检索资源共享邀请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchResourceShareInvitation(request *model.SearchResourceShareInvitationRequest) (*model.SearchResourceShareInvitationResponse, error) {
	requestDef := GenReqDefForSearchResourceShareInvitation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceShareInvitationResponse), nil
	}
}

// SearchResourceShareInvitationInvoker 检索共享邀请
func (c *RamClient) SearchResourceShareInvitationInvoker(request *model.SearchResourceShareInvitationRequest) *SearchResourceShareInvitationInvoker {
	requestDef := GenReqDefForSearchResourceShareInvitation()
	return &SearchResourceShareInvitationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateResourceShareTags 资源共享实例增加标签
//
// 资源共享实例增加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) BatchCreateResourceShareTags(request *model.BatchCreateResourceShareTagsRequest) (*model.BatchCreateResourceShareTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateResourceShareTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateResourceShareTagsResponse), nil
	}
}

// BatchCreateResourceShareTagsInvoker 资源共享实例增加标签
func (c *RamClient) BatchCreateResourceShareTagsInvoker(request *model.BatchCreateResourceShareTagsRequest) *BatchCreateResourceShareTagsInvoker {
	requestDef := GenReqDefForBatchCreateResourceShareTags()
	return &BatchCreateResourceShareTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteResourceShareTags 删除资源共享实例的标签
//
// 删除资源共享实例指定的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) BatchDeleteResourceShareTags(request *model.BatchDeleteResourceShareTagsRequest) (*model.BatchDeleteResourceShareTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteResourceShareTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteResourceShareTagsResponse), nil
	}
}

// BatchDeleteResourceShareTagsInvoker 删除资源共享实例的标签
func (c *RamClient) BatchDeleteResourceShareTagsInvoker(request *model.BatchDeleteResourceShareTagsRequest) *BatchDeleteResourceShareTagsInvoker {
	requestDef := GenReqDefForBatchDeleteResourceShareTags()
	return &BatchDeleteResourceShareTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceShareTags 查询已使用的标签列表
//
// 查询资源共享实例已使用标签的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListResourceShareTags(request *model.ListResourceShareTagsRequest) (*model.ListResourceShareTagsResponse, error) {
	requestDef := GenReqDefForListResourceShareTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceShareTagsResponse), nil
	}
}

// ListResourceShareTagsInvoker 查询已使用的标签列表
func (c *RamClient) ListResourceShareTagsInvoker(request *model.ListResourceShareTagsRequest) *ListResourceShareTagsInvoker {
	requestDef := GenReqDefForListResourceShareTags()
	return &ListResourceShareTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceSharesByTags 根据标签信息查询实例列表
//
// 根据标签信息查询资源共享实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) ListResourceSharesByTags(request *model.ListResourceSharesByTagsRequest) (*model.ListResourceSharesByTagsResponse, error) {
	requestDef := GenReqDefForListResourceSharesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceSharesByTagsResponse), nil
	}
}

// ListResourceSharesByTagsInvoker 根据标签信息查询实例列表
func (c *RamClient) ListResourceSharesByTagsInvoker(request *model.ListResourceSharesByTagsRequest) *ListResourceSharesByTagsInvoker {
	requestDef := GenReqDefForListResourceSharesByTags()
	return &ListResourceSharesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchResourceShareCountByTags 根据标签信息查询实例数量
//
// 根据标签信息查询资源共享实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *RamClient) SearchResourceShareCountByTags(request *model.SearchResourceShareCountByTagsRequest) (*model.SearchResourceShareCountByTagsResponse, error) {
	requestDef := GenReqDefForSearchResourceShareCountByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchResourceShareCountByTagsResponse), nil
	}
}

// SearchResourceShareCountByTagsInvoker 根据标签信息查询实例数量
func (c *RamClient) SearchResourceShareCountByTagsInvoker(request *model.SearchResourceShareCountByTagsRequest) *SearchResourceShareCountByTagsInvoker {
	requestDef := GenReqDefForSearchResourceShareCountByTags()
	return &SearchResourceShareCountByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
