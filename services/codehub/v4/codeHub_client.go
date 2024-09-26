package v4

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codehub/v4/model"
)

type CodeHubClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCodeHubClient(hcClient *httpclient.HcHttpClient) *CodeHubClient {
	return &CodeHubClient{HcClient: hcClient}
}

func CodeHubClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AssociateGroupUserGroup 关联代码组与成员组
//
// 关联代码组与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateGroupUserGroup(request *model.AssociateGroupUserGroupRequest) (*model.AssociateGroupUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateGroupUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateGroupUserGroupResponse), nil
	}
}

// AssociateGroupUserGroupInvoker 关联代码组与成员组
func (c *CodeHubClient) AssociateGroupUserGroupInvoker(request *model.AssociateGroupUserGroupRequest) *AssociateGroupUserGroupInvoker {
	requestDef := GenReqDefForAssociateGroupUserGroup()
	return &AssociateGroupUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRepositoryUserGroup 关联仓库与成员组
//
// 关联仓库与成员组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AssociateRepositoryUserGroup(request *model.AssociateRepositoryUserGroupRequest) (*model.AssociateRepositoryUserGroupResponse, error) {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRepositoryUserGroupResponse), nil
	}
}

// AssociateRepositoryUserGroupInvoker 关联仓库与成员组
func (c *CodeHubClient) AssociateRepositoryUserGroupInvoker(request *model.AssociateRepositoryUserGroupRequest) *AssociateRepositoryUserGroupInvoker {
	requestDef := GenReqDefForAssociateRepositoryUserGroup()
	return &AssociateRepositoryUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroup 创建代码组
//
// 创建代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建代码组
func (c *CodeHubClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除代码组
//
// 删除代码组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除代码组
func (c *CodeHubClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManageableGroups 项目下拥有创建权限的代码组列表
//
// 项目下拥有创建权限的代码组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListManageableGroups(request *model.ListManageableGroupsRequest) (*model.ListManageableGroupsResponse, error) {
	requestDef := GenReqDefForListManageableGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManageableGroupsResponse), nil
	}
}

// ListManageableGroupsInvoker 项目下拥有创建权限的代码组列表
func (c *CodeHubClient) ListManageableGroupsInvoker(request *model.ListManageableGroupsRequest) *ListManageableGroupsInvoker {
	requestDef := GenReqDefForListManageableGroups()
	return &ListManageableGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroup 获取代码组信息
//
// 获取代码组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ShowGroup(request *model.ShowGroupRequest) (*model.ShowGroupResponse, error) {
	requestDef := GenReqDefForShowGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupResponse), nil
	}
}

// ShowGroupInvoker 获取代码组信息
func (c *CodeHubClient) ShowGroupInvoker(request *model.ShowGroupRequest) *ShowGroupInvoker {
	requestDef := GenReqDefForShowGroup()
	return &ShowGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTrustedIpAddress 添加仓库ip白名单
//
// 添加仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddTrustedIpAddress(request *model.AddTrustedIpAddressRequest) (*model.AddTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTrustedIpAddressResponse), nil
	}
}

// AddTrustedIpAddressInvoker 添加仓库ip白名单
func (c *CodeHubClient) AddTrustedIpAddressInvoker(request *model.AddTrustedIpAddressRequest) *AddTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTrustedIpAddress()
	return &AddTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrustedIpAddress 删除仓库ip白名单
//
// 删除仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteTrustedIpAddress(request *model.DeleteTrustedIpAddressRequest) (*model.DeleteTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrustedIpAddressResponse), nil
	}
}

// DeleteTrustedIpAddressInvoker 删除仓库ip白名单
func (c *CodeHubClient) DeleteTrustedIpAddressInvoker(request *model.DeleteTrustedIpAddressRequest) *DeleteTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTrustedIpAddress()
	return &DeleteTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrustedIpAddresses 获取仓库ip白名单
//
// 获取仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTrustedIpAddresses(request *model.ListTrustedIpAddressesRequest) (*model.ListTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrustedIpAddressesResponse), nil
	}
}

// ListTrustedIpAddressesInvoker 获取仓库ip白名单
func (c *CodeHubClient) ListTrustedIpAddressesInvoker(request *model.ListTrustedIpAddressesRequest) *ListTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTrustedIpAddresses()
	return &ListTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LockRepository 根据仓库短ID锁定仓库
//
// 根据仓库短ID锁定仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) LockRepository(request *model.LockRepositoryRequest) (*model.LockRepositoryResponse, error) {
	requestDef := GenReqDefForLockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LockRepositoryResponse), nil
	}
}

// LockRepositoryInvoker 根据仓库短ID锁定仓库
func (c *CodeHubClient) LockRepositoryInvoker(request *model.LockRepositoryRequest) *LockRepositoryInvoker {
	requestDef := GenReqDefForLockRepository()
	return &LockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockRepository 根据仓库短ID解锁仓库
//
// 根据仓库短ID解锁仓库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UnlockRepository(request *model.UnlockRepositoryRequest) (*model.UnlockRepositoryResponse, error) {
	requestDef := GenReqDefForUnlockRepository()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockRepositoryResponse), nil
	}
}

// UnlockRepositoryInvoker 根据仓库短ID解锁仓库
func (c *CodeHubClient) UnlockRepositoryInvoker(request *model.UnlockRepositoryRequest) *UnlockRepositoryInvoker {
	requestDef := GenReqDefForUnlockRepository()
	return &UnlockRepositoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrustedIpAddress 修改仓库ip白名单
//
// 修改仓库ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateTrustedIpAddress(request *model.UpdateTrustedIpAddressRequest) (*model.UpdateTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrustedIpAddressResponse), nil
	}
}

// UpdateTrustedIpAddressInvoker 修改仓库ip白名单
func (c *CodeHubClient) UpdateTrustedIpAddressInvoker(request *model.UpdateTrustedIpAddressRequest) *UpdateTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTrustedIpAddress()
	return &UpdateTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTenantTrustedIpAddress 添加租户ip白名单
//
// 添加租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) AddTenantTrustedIpAddress(request *model.AddTenantTrustedIpAddressRequest) (*model.AddTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTenantTrustedIpAddressResponse), nil
	}
}

// AddTenantTrustedIpAddressInvoker 添加租户ip白名单
func (c *CodeHubClient) AddTenantTrustedIpAddressInvoker(request *model.AddTenantTrustedIpAddressRequest) *AddTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForAddTenantTrustedIpAddress()
	return &AddTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTenantTrustedIpAddress 删除租户ip白名单
//
// 删除租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) DeleteTenantTrustedIpAddress(request *model.DeleteTenantTrustedIpAddressRequest) (*model.DeleteTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTenantTrustedIpAddressResponse), nil
	}
}

// DeleteTenantTrustedIpAddressInvoker 删除租户ip白名单
func (c *CodeHubClient) DeleteTenantTrustedIpAddressInvoker(request *model.DeleteTenantTrustedIpAddressRequest) *DeleteTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForDeleteTenantTrustedIpAddress()
	return &DeleteTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantTrustedIpAddresses 获取租户ip白名单
//
// 获取租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) ListTenantTrustedIpAddresses(request *model.ListTenantTrustedIpAddressesRequest) (*model.ListTenantTrustedIpAddressesResponse, error) {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantTrustedIpAddressesResponse), nil
	}
}

// ListTenantTrustedIpAddressesInvoker 获取租户ip白名单
func (c *CodeHubClient) ListTenantTrustedIpAddressesInvoker(request *model.ListTenantTrustedIpAddressesRequest) *ListTenantTrustedIpAddressesInvoker {
	requestDef := GenReqDefForListTenantTrustedIpAddresses()
	return &ListTenantTrustedIpAddressesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTenantTrustedIpAddress 修改租户ip白名单
//
// 修改租户ip白名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CodeHubClient) UpdateTenantTrustedIpAddress(request *model.UpdateTenantTrustedIpAddressRequest) (*model.UpdateTenantTrustedIpAddressResponse, error) {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTenantTrustedIpAddressResponse), nil
	}
}

// UpdateTenantTrustedIpAddressInvoker 修改租户ip白名单
func (c *CodeHubClient) UpdateTenantTrustedIpAddressInvoker(request *model.UpdateTenantTrustedIpAddressRequest) *UpdateTenantTrustedIpAddressInvoker {
	requestDef := GenReqDefForUpdateTenantTrustedIpAddress()
	return &UpdateTenantTrustedIpAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
