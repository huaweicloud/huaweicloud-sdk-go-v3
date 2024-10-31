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
// 在指定的身份源中创建一个IAM身份中心用户组。
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
// 根据用户组ID，删除对应的IAM身份中心用户组。
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
// 根据用户组ID，查询IAM身份中心用户组的详情信息。
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

// GetGroupId 查询用户组ID
//
// 根据显示名或外部身份源ID，以精确匹配的方式查询用户组ID。显示名和外部身份源ID两种查询方式二选一，不支持同时传入。
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
// 查询指定身份源下的IAM身份中心用户组列表。
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
// 根据用户组ID，更新对应IAM身份中心用户组的属性。
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
// 将用户添加到用户组中，用户和用户组必须在同一身份源下。
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
// 根据关联关系ID解绑用户和用户组，也就是将用户移出用户组。
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
// 根据关联关系ID，查询此关联关系的详情信息。
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
// 根据用户ID和用户组ID，查询对应的关联关系ID。
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

// IsMemberInGroups 查询用户是否为用户组成员
//
// 根据用户ID和用户组ID列表，查询用户是否为用户组的成员。
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

// IsMemberInGroupsInvoker 查询用户是否为用户组成员
func (c *IdentityCenterStoreClient) IsMemberInGroupsInvoker(request *model.IsMemberInGroupsRequest) *IsMemberInGroupsInvoker {
	requestDef := GenReqDefForIsMemberInGroups()
	return &IsMemberInGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupMemberships 列出组中的用户
//
// 根据用户组ID，列出用户组中的用户。
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
// 根据用户ID，列出用户加入的用户组。
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

// CreateUser 创建用户
//
// 在指定的身份源中创建一个IAM身份中心用户。
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

// DeleteUser 删除用户
//
// 根据用户ID，删除对应的IAM身份中心用户。
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
// 根据用户ID，查询对应IAM身份中心用户的详情信息。
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

// GetUserId 查询用户ID
//
// 根据用户名或外部身份源ID，以精确匹配的方式查询用户ID。用户名和外部身份源ID两种查询方式二选一，不支持同时传入。
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

// ListUsers 列出用户
//
// 查询指定身份源下的IAM身份中心用户列表。
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

// UpdateUser 更新用户
//
// 根据用户ID，更新对应IAM身份中心用户的属性。
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
