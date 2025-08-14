package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterstore/v1/model"
)

type IdentityCenterStoreClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterStoreClient(hcClient *httpclient.HcHttpClient) *IdentityCenterStoreClient {
	return &IdentityCenterStoreClient{HcClient: hcClient}
}

func IdentityCenterStoreClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateGroup 创建用户组
//
// 在指定的身份源中创建一个IAM身份中心用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建用户组
func (c *IdentityCenterStoreClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除用户组
//
// 根据用户组ID，删除对应的IAM身份中心用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除用户组
func (c *IdentityCenterStoreClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeGroup 查询用户组详情
//
// 根据用户组ID，查询IAM身份中心用户组的详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribeGroup(request *model.DescribeGroupRequest) (*model.DescribeGroupResponse, error) {
	requestDef := GenReqDefForDescribeGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeGroupResponse), nil
	}
}

// DescribeGroupInvoker 查询用户组详情
func (c *IdentityCenterStoreClient) DescribeGroupInvoker(request *model.DescribeGroupRequest) *DescribeGroupInvoker {
	requestDef := GenReqDefForDescribeGroup()
	return &DescribeGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeGroups 批量查询指定用户组详情
//
// 批量查询指定用户组详情。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribeGroups(request *model.DescribeGroupsRequest) (*model.DescribeGroupsResponse, error) {
	requestDef := GenReqDefForDescribeGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeGroupsResponse), nil
	}
}

// DescribeGroupsInvoker 批量查询指定用户组详情
func (c *IdentityCenterStoreClient) DescribeGroupsInvoker(request *model.DescribeGroupsRequest) *DescribeGroupsInvoker {
	requestDef := GenReqDefForDescribeGroups()
	return &DescribeGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetGroupId 查询用户组ID
//
// 根据显示名或外部身份源ID，以精确匹配的方式查询用户组ID。显示名和外部身份源ID两种查询方式二选一，不支持同时传入。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) GetGroupId(request *model.GetGroupIdRequest) (*model.GetGroupIdResponse, error) {
	requestDef := GenReqDefForGetGroupId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetGroupIdResponse), nil
	}
}

// GetGroupIdInvoker 查询用户组ID
func (c *IdentityCenterStoreClient) GetGroupIdInvoker(request *model.GetGroupIdRequest) *GetGroupIdInvoker {
	requestDef := GenReqDefForGetGroupId()
	return &GetGroupIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 列出用户组
//
// 查询指定身份源下的IAM身份中心用户组列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 列出用户组
func (c *IdentityCenterStoreClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroup 更新用户组
//
// 根据用户组ID，更新对应IAM身份中心用户组的属性。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdateGroup(request *model.UpdateGroupRequest) (*model.UpdateGroupResponse, error) {
	requestDef := GenReqDefForUpdateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupResponse), nil
	}
}

// UpdateGroupInvoker 更新用户组
func (c *IdentityCenterStoreClient) UpdateGroupInvoker(request *model.UpdateGroupRequest) *UpdateGroupInvoker {
	requestDef := GenReqDefForUpdateGroup()
	return &UpdateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupMembership 绑定用户和组
//
// 将用户添加到用户组中，用户和用户组必须在同一身份源下。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateGroupMembership(request *model.CreateGroupMembershipRequest) (*model.CreateGroupMembershipResponse, error) {
	requestDef := GenReqDefForCreateGroupMembership()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupMembershipResponse), nil
	}
}

// CreateGroupMembershipInvoker 绑定用户和组
func (c *IdentityCenterStoreClient) CreateGroupMembershipInvoker(request *model.CreateGroupMembershipRequest) *CreateGroupMembershipInvoker {
	requestDef := GenReqDefForCreateGroupMembership()
	return &CreateGroupMembershipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupMembership 解绑用户和组
//
// 根据关联关系ID解绑用户和用户组，也就是将用户移出用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteGroupMembership(request *model.DeleteGroupMembershipRequest) (*model.DeleteGroupMembershipResponse, error) {
	requestDef := GenReqDefForDeleteGroupMembership()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupMembershipResponse), nil
	}
}

// DeleteGroupMembershipInvoker 解绑用户和组
func (c *IdentityCenterStoreClient) DeleteGroupMembershipInvoker(request *model.DeleteGroupMembershipRequest) *DeleteGroupMembershipInvoker {
	requestDef := GenReqDefForDeleteGroupMembership()
	return &DeleteGroupMembershipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeGroupMembership 查询绑定关系详情
//
// 根据关联关系ID，查询此关联关系的详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribeGroupMembership(request *model.DescribeGroupMembershipRequest) (*model.DescribeGroupMembershipResponse, error) {
	requestDef := GenReqDefForDescribeGroupMembership()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeGroupMembershipResponse), nil
	}
}

// DescribeGroupMembershipInvoker 查询绑定关系详情
func (c *IdentityCenterStoreClient) DescribeGroupMembershipInvoker(request *model.DescribeGroupMembershipRequest) *DescribeGroupMembershipInvoker {
	requestDef := GenReqDefForDescribeGroupMembership()
	return &DescribeGroupMembershipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetGroupMembershipId 查询绑定关系ID
//
// 根据用户ID和用户组ID，查询对应的关联关系ID。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) GetGroupMembershipId(request *model.GetGroupMembershipIdRequest) (*model.GetGroupMembershipIdResponse, error) {
	requestDef := GenReqDefForGetGroupMembershipId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetGroupMembershipIdResponse), nil
	}
}

// GetGroupMembershipIdInvoker 查询绑定关系ID
func (c *IdentityCenterStoreClient) GetGroupMembershipIdInvoker(request *model.GetGroupMembershipIdRequest) *GetGroupMembershipIdInvoker {
	requestDef := GenReqDefForGetGroupMembershipId()
	return &GetGroupMembershipIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// IsMemberInGroups 查询用户是否是用户组成员
//
// 根据用户ID和用户组ID列表，查询用户是否为用户组的成员。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) IsMemberInGroups(request *model.IsMemberInGroupsRequest) (*model.IsMemberInGroupsResponse, error) {
	requestDef := GenReqDefForIsMemberInGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.IsMemberInGroupsResponse), nil
	}
}

// IsMemberInGroupsInvoker 查询用户是否是用户组成员
func (c *IdentityCenterStoreClient) IsMemberInGroupsInvoker(request *model.IsMemberInGroupsRequest) *IsMemberInGroupsInvoker {
	requestDef := GenReqDefForIsMemberInGroups()
	return &IsMemberInGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMemberships 列出组中的用户
//
// 根据用户组ID，列出用户组中的用户。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListGroupMemberships(request *model.ListGroupMembershipsRequest) (*model.ListGroupMembershipsResponse, error) {
	requestDef := GenReqDefForListGroupMemberships()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMembershipsResponse), nil
	}
}

// ListGroupMembershipsInvoker 列出组中的用户
func (c *IdentityCenterStoreClient) ListGroupMembershipsInvoker(request *model.ListGroupMembershipsRequest) *ListGroupMembershipsInvoker {
	requestDef := GenReqDefForListGroupMemberships()
	return &ListGroupMembershipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMembershipsForMember 列出用户加入的组
//
// 根据用户ID，列出用户加入的用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListGroupMembershipsForMember(request *model.ListGroupMembershipsForMemberRequest) (*model.ListGroupMembershipsForMemberResponse, error) {
	requestDef := GenReqDefForListGroupMembershipsForMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupMembershipsForMemberResponse), nil
	}
}

// ListGroupMembershipsForMemberInvoker 列出用户加入的组
func (c *IdentityCenterStoreClient) ListGroupMembershipsForMemberInvoker(request *model.ListGroupMembershipsForMemberRequest) *ListGroupMembershipsForMemberInvoker {
	requestDef := GenReqDefForListGroupMembershipsForMember()
	return &ListGroupMembershipsForMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateExternalIdPConfigurationForDirectory 创建外部身份提供商配置
//
// 创建外部身份提供商配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateExternalIdPConfigurationForDirectory(request *model.CreateExternalIdPConfigurationForDirectoryRequest) (*model.CreateExternalIdPConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForCreateExternalIdPConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExternalIdPConfigurationForDirectoryResponse), nil
	}
}

// CreateExternalIdPConfigurationForDirectoryInvoker 创建外部身份提供商配置
func (c *IdentityCenterStoreClient) CreateExternalIdPConfigurationForDirectoryInvoker(request *model.CreateExternalIdPConfigurationForDirectoryRequest) *CreateExternalIdPConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForCreateExternalIdPConfigurationForDirectory()
	return &CreateExternalIdPConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteExternalIdPCertificate 删除外部身份提供商证书
//
// 删除外部身份提供商证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteExternalIdPCertificate(request *model.DeleteExternalIdPCertificateRequest) (*model.DeleteExternalIdPCertificateResponse, error) {
	requestDef := GenReqDefForDeleteExternalIdPCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExternalIdPCertificateResponse), nil
	}
}

// DeleteExternalIdPCertificateInvoker 删除外部身份提供商证书
func (c *IdentityCenterStoreClient) DeleteExternalIdPCertificateInvoker(request *model.DeleteExternalIdPCertificateRequest) *DeleteExternalIdPCertificateInvoker {
	requestDef := GenReqDefForDeleteExternalIdPCertificate()
	return &DeleteExternalIdPCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteExternalIdPConfigurationForDirectory 删除外部身份提供商配置
//
// 删除外部身份提供商配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteExternalIdPConfigurationForDirectory(request *model.DeleteExternalIdPConfigurationForDirectoryRequest) (*model.DeleteExternalIdPConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForDeleteExternalIdPConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteExternalIdPConfigurationForDirectoryResponse), nil
	}
}

// DeleteExternalIdPConfigurationForDirectoryInvoker 删除外部身份提供商配置
func (c *IdentityCenterStoreClient) DeleteExternalIdPConfigurationForDirectoryInvoker(request *model.DeleteExternalIdPConfigurationForDirectoryRequest) *DeleteExternalIdPConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForDeleteExternalIdPConfigurationForDirectory()
	return &DeleteExternalIdPConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableExternalIdPConfigurationForDirectory 停用外部身份提供商
//
// 停用外部身份提供商。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DisableExternalIdPConfigurationForDirectory(request *model.DisableExternalIdPConfigurationForDirectoryRequest) (*model.DisableExternalIdPConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForDisableExternalIdPConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableExternalIdPConfigurationForDirectoryResponse), nil
	}
}

// DisableExternalIdPConfigurationForDirectoryInvoker 停用外部身份提供商
func (c *IdentityCenterStoreClient) DisableExternalIdPConfigurationForDirectoryInvoker(request *model.DisableExternalIdPConfigurationForDirectoryRequest) *DisableExternalIdPConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForDisableExternalIdPConfigurationForDirectory()
	return &DisableExternalIdPConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableExternalIdPConfigurationForDirectory 启用外部身份提供商
//
// 启用外部身份提供商。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) EnableExternalIdPConfigurationForDirectory(request *model.EnableExternalIdPConfigurationForDirectoryRequest) (*model.EnableExternalIdPConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForEnableExternalIdPConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableExternalIdPConfigurationForDirectoryResponse), nil
	}
}

// EnableExternalIdPConfigurationForDirectoryInvoker 启用外部身份提供商
func (c *IdentityCenterStoreClient) EnableExternalIdPConfigurationForDirectoryInvoker(request *model.EnableExternalIdPConfigurationForDirectoryRequest) *EnableExternalIdPConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForEnableExternalIdPConfigurationForDirectory()
	return &EnableExternalIdPConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportExternalIdPCertificate 导入外部身份提供商证书
//
// 导入外部身份提供商证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ImportExternalIdPCertificate(request *model.ImportExternalIdPCertificateRequest) (*model.ImportExternalIdPCertificateResponse, error) {
	requestDef := GenReqDefForImportExternalIdPCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportExternalIdPCertificateResponse), nil
	}
}

// ImportExternalIdPCertificateInvoker 导入外部身份提供商证书
func (c *IdentityCenterStoreClient) ImportExternalIdPCertificateInvoker(request *model.ImportExternalIdPCertificateRequest) *ImportExternalIdPCertificateInvoker {
	requestDef := GenReqDefForImportExternalIdPCertificate()
	return &ImportExternalIdPCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExternalIdPCertificates 列出外部身份提供商证书
//
// 查询外部身份提供商证书列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListExternalIdPCertificates(request *model.ListExternalIdPCertificatesRequest) (*model.ListExternalIdPCertificatesResponse, error) {
	requestDef := GenReqDefForListExternalIdPCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExternalIdPCertificatesResponse), nil
	}
}

// ListExternalIdPCertificatesInvoker 列出外部身份提供商证书
func (c *IdentityCenterStoreClient) ListExternalIdPCertificatesInvoker(request *model.ListExternalIdPCertificatesRequest) *ListExternalIdPCertificatesInvoker {
	requestDef := GenReqDefForListExternalIdPCertificates()
	return &ListExternalIdPCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExternalIdPConfigurationsForDirectory 查询外部身份提供商配置
//
// 查询外部身份提供商配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListExternalIdPConfigurationsForDirectory(request *model.ListExternalIdPConfigurationsForDirectoryRequest) (*model.ListExternalIdPConfigurationsForDirectoryResponse, error) {
	requestDef := GenReqDefForListExternalIdPConfigurationsForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExternalIdPConfigurationsForDirectoryResponse), nil
	}
}

// ListExternalIdPConfigurationsForDirectoryInvoker 查询外部身份提供商配置
func (c *IdentityCenterStoreClient) ListExternalIdPConfigurationsForDirectoryInvoker(request *model.ListExternalIdPConfigurationsForDirectoryRequest) *ListExternalIdPConfigurationsForDirectoryInvoker {
	requestDef := GenReqDefForListExternalIdPConfigurationsForDirectory()
	return &ListExternalIdPConfigurationsForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateExternalIdPConfigurationForDirectory 更新外部身份提供商配置
//
// 更新外部身份提供商配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdateExternalIdPConfigurationForDirectory(request *model.UpdateExternalIdPConfigurationForDirectoryRequest) (*model.UpdateExternalIdPConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForUpdateExternalIdPConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExternalIdPConfigurationForDirectoryResponse), nil
	}
}

// UpdateExternalIdPConfigurationForDirectoryInvoker 更新外部身份提供商配置
func (c *IdentityCenterStoreClient) UpdateExternalIdPConfigurationForDirectoryInvoker(request *model.UpdateExternalIdPConfigurationForDirectoryRequest) *UpdateExternalIdPConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForUpdateExternalIdPConfigurationForDirectory()
	return &UpdateExternalIdPConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribePasswordPolicy 查询自定义密码策略
//
// 查询自定义密码策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribePasswordPolicy(request *model.DescribePasswordPolicyRequest) (*model.DescribePasswordPolicyResponse, error) {
	requestDef := GenReqDefForDescribePasswordPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribePasswordPolicyResponse), nil
	}
}

// DescribePasswordPolicyInvoker 查询自定义密码策略
func (c *IdentityCenterStoreClient) DescribePasswordPolicyInvoker(request *model.DescribePasswordPolicyRequest) *DescribePasswordPolicyInvoker {
	requestDef := GenReqDefForDescribePasswordPolicy()
	return &DescribePasswordPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePasswordPolicy 更新自定义密码策略
//
// 更新自定义密码策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdatePasswordPolicy(request *model.UpdatePasswordPolicyRequest) (*model.UpdatePasswordPolicyResponse, error) {
	requestDef := GenReqDefForUpdatePasswordPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordPolicyResponse), nil
	}
}

// UpdatePasswordPolicyInvoker 更新自定义密码策略
func (c *IdentityCenterStoreClient) UpdatePasswordPolicyInvoker(request *model.UpdatePasswordPolicyRequest) *UpdatePasswordPolicyInvoker {
	requestDef := GenReqDefForUpdatePasswordPolicy()
	return &UpdatePasswordPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSpCertificate 创建服务提供商证书
//
// 创建服务提供商SAML协议签名证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateSpCertificate(request *model.CreateSpCertificateRequest) (*model.CreateSpCertificateResponse, error) {
	requestDef := GenReqDefForCreateSpCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSpCertificateResponse), nil
	}
}

// CreateSpCertificateInvoker 创建服务提供商证书
func (c *IdentityCenterStoreClient) CreateSpCertificateInvoker(request *model.CreateSpCertificateRequest) *CreateSpCertificateInvoker {
	requestDef := GenReqDefForCreateSpCertificate()
	return &CreateSpCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSpCertificate 删除服务提供商证书
//
// 删除服务提供商SAML协议签名证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteSpCertificate(request *model.DeleteSpCertificateRequest) (*model.DeleteSpCertificateResponse, error) {
	requestDef := GenReqDefForDeleteSpCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSpCertificateResponse), nil
	}
}

// DeleteSpCertificateInvoker 删除服务提供商证书
func (c *IdentityCenterStoreClient) DeleteSpCertificateInvoker(request *model.DeleteSpCertificateRequest) *DeleteSpCertificateInvoker {
	requestDef := GenReqDefForDeleteSpCertificate()
	return &DeleteSpCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetSpConfigurationForDirectory 查询服务提供商配置
//
// 查询服务提供商配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) GetSpConfigurationForDirectory(request *model.GetSpConfigurationForDirectoryRequest) (*model.GetSpConfigurationForDirectoryResponse, error) {
	requestDef := GenReqDefForGetSpConfigurationForDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetSpConfigurationForDirectoryResponse), nil
	}
}

// GetSpConfigurationForDirectoryInvoker 查询服务提供商配置
func (c *IdentityCenterStoreClient) GetSpConfigurationForDirectoryInvoker(request *model.GetSpConfigurationForDirectoryRequest) *GetSpConfigurationForDirectoryInvoker {
	requestDef := GenReqDefForGetSpConfigurationForDirectory()
	return &GetSpConfigurationForDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpCertificates 列出服务提供商证书
//
// 查询服务提供商SAML协议签名证书
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListSpCertificates(request *model.ListSpCertificatesRequest) (*model.ListSpCertificatesResponse, error) {
	requestDef := GenReqDefForListSpCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpCertificatesResponse), nil
	}
}

// ListSpCertificatesInvoker 列出服务提供商证书
func (c *IdentityCenterStoreClient) ListSpCertificatesInvoker(request *model.ListSpCertificatesRequest) *ListSpCertificatesInvoker {
	requestDef := GenReqDefForListSpCertificates()
	return &ListSpCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSpActiveCertificate 激活服务提供商证书
//
// 激活服务提供商SAML协议签名证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdateSpActiveCertificate(request *model.UpdateSpActiveCertificateRequest) (*model.UpdateSpActiveCertificateResponse, error) {
	requestDef := GenReqDefForUpdateSpActiveCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSpActiveCertificateResponse), nil
	}
}

// UpdateSpActiveCertificateInvoker 激活服务提供商证书
func (c *IdentityCenterStoreClient) UpdateSpActiveCertificateInvoker(request *model.UpdateSpActiveCertificateRequest) *UpdateSpActiveCertificateInvoker {
	requestDef := GenReqDefForUpdateSpActiveCertificate()
	return &UpdateSpActiveCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetIdentityStoreSummary 查询身份源配额信息
//
// 查询身份源配额信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) GetIdentityStoreSummary(request *model.GetIdentityStoreSummaryRequest) (*model.GetIdentityStoreSummaryResponse, error) {
	requestDef := GenReqDefForGetIdentityStoreSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetIdentityStoreSummaryResponse), nil
	}
}

// GetIdentityStoreSummaryInvoker 查询身份源配额信息
func (c *IdentityCenterStoreClient) GetIdentityStoreSummaryInvoker(request *model.GetIdentityStoreSummaryRequest) *GetIdentityStoreSummaryInvoker {
	requestDef := GenReqDefForGetIdentityStoreSummary()
	return &GetIdentityStoreSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBearerToken 创建访问令牌
//
// 创建访问令牌。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateBearerToken(request *model.CreateBearerTokenRequest) (*model.CreateBearerTokenResponse, error) {
	requestDef := GenReqDefForCreateBearerToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBearerTokenResponse), nil
	}
}

// CreateBearerTokenInvoker 创建访问令牌
func (c *IdentityCenterStoreClient) CreateBearerTokenInvoker(request *model.CreateBearerTokenRequest) *CreateBearerTokenInvoker {
	requestDef := GenReqDefForCreateBearerToken()
	return &CreateBearerTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateProvisioningTenant 启用自动预置
//
// 启用自动预置，开启SCIM协议自动同步功能。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateProvisioningTenant(request *model.CreateProvisioningTenantRequest) (*model.CreateProvisioningTenantResponse, error) {
	requestDef := GenReqDefForCreateProvisioningTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateProvisioningTenantResponse), nil
	}
}

// CreateProvisioningTenantInvoker 启用自动预置
func (c *IdentityCenterStoreClient) CreateProvisioningTenantInvoker(request *model.CreateProvisioningTenantRequest) *CreateProvisioningTenantInvoker {
	requestDef := GenReqDefForCreateProvisioningTenant()
	return &CreateProvisioningTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBearerToken 删除访问令牌
//
// 删除访问令牌。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteBearerToken(request *model.DeleteBearerTokenRequest) (*model.DeleteBearerTokenResponse, error) {
	requestDef := GenReqDefForDeleteBearerToken()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBearerTokenResponse), nil
	}
}

// DeleteBearerTokenInvoker 删除访问令牌
func (c *IdentityCenterStoreClient) DeleteBearerTokenInvoker(request *model.DeleteBearerTokenRequest) *DeleteBearerTokenInvoker {
	requestDef := GenReqDefForDeleteBearerToken()
	return &DeleteBearerTokenInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProvisioningTenant 删除自动预置
//
// 删除自动预置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteProvisioningTenant(request *model.DeleteProvisioningTenantRequest) (*model.DeleteProvisioningTenantResponse, error) {
	requestDef := GenReqDefForDeleteProvisioningTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProvisioningTenantResponse), nil
	}
}

// DeleteProvisioningTenantInvoker 删除自动预置
func (c *IdentityCenterStoreClient) DeleteProvisioningTenantInvoker(request *model.DeleteProvisioningTenantRequest) *DeleteProvisioningTenantInvoker {
	requestDef := GenReqDefForDeleteProvisioningTenant()
	return &DeleteProvisioningTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBearerTokens 列出访问令牌
//
// 查询访问令牌列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListBearerTokens(request *model.ListBearerTokensRequest) (*model.ListBearerTokensResponse, error) {
	requestDef := GenReqDefForListBearerTokens()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBearerTokensResponse), nil
	}
}

// ListBearerTokensInvoker 列出访问令牌
func (c *IdentityCenterStoreClient) ListBearerTokensInvoker(request *model.ListBearerTokensRequest) *ListBearerTokensInvoker {
	requestDef := GenReqDefForListBearerTokens()
	return &ListBearerTokensInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProvisioningTenants 查询自动预置信息
//
// 查询是否开启自动预置，返回具体SCIM自动预配配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListProvisioningTenants(request *model.ListProvisioningTenantsRequest) (*model.ListProvisioningTenantsResponse, error) {
	requestDef := GenReqDefForListProvisioningTenants()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvisioningTenantsResponse), nil
	}
}

// ListProvisioningTenantsInvoker 查询自动预置信息
func (c *IdentityCenterStoreClient) ListProvisioningTenantsInvoker(request *model.ListProvisioningTenantsRequest) *ListProvisioningTenantsInvoker {
	requestDef := GenReqDefForListProvisioningTenants()
	return &ListProvisioningTenantsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSessions 批量删除用户登录会话
//
// 批量删除用户登录会话。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) BatchDeleteSessions(request *model.BatchDeleteSessionsRequest) (*model.BatchDeleteSessionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSessionsResponse), nil
	}
}

// BatchDeleteSessionsInvoker 批量删除用户登录会话
func (c *IdentityCenterStoreClient) BatchDeleteSessionsInvoker(request *model.BatchDeleteSessionsRequest) *BatchDeleteSessionsInvoker {
	requestDef := GenReqDefForBatchDeleteSessions()
	return &BatchDeleteSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListMfaDevicesForUser 列出用户MFA设备
//
// 查询指定用户的MFA设备列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) BatchListMfaDevicesForUser(request *model.BatchListMfaDevicesForUserRequest) (*model.BatchListMfaDevicesForUserResponse, error) {
	requestDef := GenReqDefForBatchListMfaDevicesForUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListMfaDevicesForUserResponse), nil
	}
}

// BatchListMfaDevicesForUserInvoker 列出用户MFA设备
func (c *IdentityCenterStoreClient) BatchListMfaDevicesForUserInvoker(request *model.BatchListMfaDevicesForUserRequest) *BatchListMfaDevicesForUserInvoker {
	requestDef := GenReqDefForBatchListMfaDevicesForUser()
	return &BatchListMfaDevicesForUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 在指定的身份源中创建一个IAM身份中心用户。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *IdentityCenterStoreClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMfaDeviceForUser 删除用户MFA设备
//
// 删除用户绑定的MFA设备。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteMfaDeviceForUser(request *model.DeleteMfaDeviceForUserRequest) (*model.DeleteMfaDeviceForUserResponse, error) {
	requestDef := GenReqDefForDeleteMfaDeviceForUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMfaDeviceForUserResponse), nil
	}
}

// DeleteMfaDeviceForUserInvoker 删除用户MFA设备
func (c *IdentityCenterStoreClient) DeleteMfaDeviceForUserInvoker(request *model.DeleteMfaDeviceForUserRequest) *DeleteMfaDeviceForUserInvoker {
	requestDef := GenReqDefForDeleteMfaDeviceForUser()
	return &DeleteMfaDeviceForUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 根据用户ID，删除对应的IAM身份中心用户。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *IdentityCenterStoreClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeUser 查询用户详情
//
// 根据用户ID，查询对应IAM身份中心用户的详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribeUser(request *model.DescribeUserRequest) (*model.DescribeUserResponse, error) {
	requestDef := GenReqDefForDescribeUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeUserResponse), nil
	}
}

// DescribeUserInvoker 查询用户详情
func (c *IdentityCenterStoreClient) DescribeUserInvoker(request *model.DescribeUserRequest) *DescribeUserInvoker {
	requestDef := GenReqDefForDescribeUser()
	return &DescribeUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeUsers 批量查询指定用户详情
//
// 批量查询指定用户详情。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DescribeUsers(request *model.DescribeUsersRequest) (*model.DescribeUsersResponse, error) {
	requestDef := GenReqDefForDescribeUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeUsersResponse), nil
	}
}

// DescribeUsersInvoker 批量查询指定用户详情
func (c *IdentityCenterStoreClient) DescribeUsersInvoker(request *model.DescribeUsersRequest) *DescribeUsersInvoker {
	requestDef := GenReqDefForDescribeUsers()
	return &DescribeUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableUser 禁用用户
//
// 禁用IAM身份中心的用户。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) DisableUser(request *model.DisableUserRequest) (*model.DisableUserResponse, error) {
	requestDef := GenReqDefForDisableUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableUserResponse), nil
	}
}

// DisableUserInvoker 禁用用户
func (c *IdentityCenterStoreClient) DisableUserInvoker(request *model.DisableUserRequest) *DisableUserInvoker {
	requestDef := GenReqDefForDisableUser()
	return &DisableUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableUser 启用用户
//
// 启用IAM身份中心的用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) EnableUser(request *model.EnableUserRequest) (*model.EnableUserResponse, error) {
	requestDef := GenReqDefForEnableUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableUserResponse), nil
	}
}

// EnableUserInvoker 启用用户
func (c *IdentityCenterStoreClient) EnableUserInvoker(request *model.EnableUserRequest) *EnableUserInvoker {
	requestDef := GenReqDefForEnableUser()
	return &EnableUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetUserId 查询用户ID
//
// 根据用户名或外部身份源ID，以精确匹配的方式查询用户ID。用户名和外部身份源ID两种查询方式二选一，不支持同时传入。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) GetUserId(request *model.GetUserIdRequest) (*model.GetUserIdResponse, error) {
	requestDef := GenReqDefForGetUserId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetUserIdResponse), nil
	}
}

// GetUserIdInvoker 查询用户ID
func (c *IdentityCenterStoreClient) GetUserIdInvoker(request *model.GetUserIdRequest) *GetUserIdInvoker {
	requestDef := GenReqDefForGetUserId()
	return &GetUserIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessions 列出用户登录会话
//
// 查询指定用户的登录会话信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListSessions(request *model.ListSessionsRequest) (*model.ListSessionsResponse, error) {
	requestDef := GenReqDefForListSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionsResponse), nil
	}
}

// ListSessionsInvoker 列出用户登录会话
func (c *IdentityCenterStoreClient) ListSessionsInvoker(request *model.ListSessionsRequest) *ListSessionsInvoker {
	requestDef := GenReqDefForListSessions()
	return &ListSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 列出用户
//
// 查询指定身份源下的IAM身份中心用户列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 列出用户
func (c *IdentityCenterStoreClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterMfaDeviceForUser 注册MFA设备
//
// 为用户发起注册MFA设备，返回一个MFA注册地址。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) RegisterMfaDeviceForUser(request *model.RegisterMfaDeviceForUserRequest) (*model.RegisterMfaDeviceForUserResponse, error) {
	requestDef := GenReqDefForRegisterMfaDeviceForUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterMfaDeviceForUserResponse), nil
	}
}

// RegisterMfaDeviceForUserInvoker 注册MFA设备
func (c *IdentityCenterStoreClient) RegisterMfaDeviceForUserInvoker(request *model.RegisterMfaDeviceForUserRequest) *RegisterMfaDeviceForUserInvoker {
	requestDef := GenReqDefForRegisterMfaDeviceForUser()
	return &RegisterMfaDeviceForUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwdMode 通过电子邮件发送密码重置链接或生成用户的一次性密码
//
// 通过电子邮件发送密码重置链接或生成用户的一次性密码。。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) ResetPwdMode(request *model.ResetPwdModeRequest) (*model.ResetPwdModeResponse, error) {
	requestDef := GenReqDefForResetPwdMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdModeResponse), nil
	}
}

// ResetPwdModeInvoker 通过电子邮件发送密码重置链接或生成用户的一次性密码
func (c *IdentityCenterStoreClient) ResetPwdModeInvoker(request *model.ResetPwdModeRequest) *ResetPwdModeInvoker {
	requestDef := GenReqDefForResetPwdMode()
	return &ResetPwdModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMfaDeviceForUser 更新MFA设备显示名称
//
// 更新MFA设备显示名称。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdateMfaDeviceForUser(request *model.UpdateMfaDeviceForUserRequest) (*model.UpdateMfaDeviceForUserResponse, error) {
	requestDef := GenReqDefForUpdateMfaDeviceForUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMfaDeviceForUserResponse), nil
	}
}

// UpdateMfaDeviceForUserInvoker 更新MFA设备显示名称
func (c *IdentityCenterStoreClient) UpdateMfaDeviceForUserInvoker(request *model.UpdateMfaDeviceForUserRequest) *UpdateMfaDeviceForUserInvoker {
	requestDef := GenReqDefForUpdateMfaDeviceForUser()
	return &UpdateMfaDeviceForUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 更新用户
//
// 根据用户ID，更新对应IAM身份中心用户的属性。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 更新用户
func (c *IdentityCenterStoreClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// VerifyEmail 验证用户邮箱
//
// 验证用户邮箱。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterStoreClient) VerifyEmail(request *model.VerifyEmailRequest) (*model.VerifyEmailResponse, error) {
	requestDef := GenReqDefForVerifyEmail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.VerifyEmailResponse), nil
	}
}

// VerifyEmailInvoker 验证用户邮箱
func (c *IdentityCenterStoreClient) VerifyEmailInvoker(request *model.VerifyEmailRequest) *VerifyEmailInvoker {
	requestDef := GenReqDefForVerifyEmail()
	return &VerifyEmailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
