package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterscim/v1/model"
)

type IdentityCenterSCIMClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterSCIMClient(hcClient *httpclient.HcHttpClient) *IdentityCenterSCIMClient {
	return &IdentityCenterSCIMClient{HcClient: hcClient}
}

func IdentityCenterSCIMClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("v1.IdentityCenterSCIMCredentials")
	return builder
}

// CreateGroup 创建用户组
//
// 使用SCIM协议，同步用户组到IAM身份中心。
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - displayName、schemas为必填项
// - 在单个请求中最多可以添加100个成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建用户组
func (c *IdentityCenterSCIMClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除用户组
//
// 删除现有用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除用户组
func (c *IdentityCenterSCIMClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetGroup 查询用户组详情
//
// 查询现有用户组的详情信息。
// 暂不支持
// IAM身份中心暂不支持此API操作的以下方面。
// - 查询用户组详情接口和列出用户组接口返回空成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) GetGroup(request *model.GetGroupRequest) (*model.GetGroupResponse, error) {
	requestDef := GenReqDefForGetGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetGroupResponse), nil
	}
}

// GetGroupInvoker 查询用户组详情
func (c *IdentityCenterSCIMClient) GetGroupInvoker(request *model.GetGroupRequest) *GetGroupInvoker {
	requestDef := GenReqDefForGetGroup()
	return &GetGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroups 列出用户组
//
// 对现有用户组列表执行筛选查询，最多只能返回50个结果。
// 暂不支持
// IAM身份中心暂不支持此API操作的以下方面。
// - 查询用户组详情接口和列出用户组接口返回空成员列表
//
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 目前，列出用户组接口最多只能返回50个结果
// - 支持的筛选器组合：(displayName)、(id)。请注意，使用id作为单个过滤器虽然有效，但应避免使用，因为已经有一个查询用户组详情接口可用
// - 过滤器中支持的比较运算符：eq
// - 必须按如下方式指定筛选器：&lt;filterAttribute&gt; eq \&quot;&lt;filterValue&gt;\&quot;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) ListGroups(request *model.ListGroupsRequest) (*model.ListGroupsResponse, error) {
	requestDef := GenReqDefForListGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsResponse), nil
	}
}

// ListGroupsInvoker 列出用户组
func (c *IdentityCenterSCIMClient) ListGroupsInvoker(request *model.ListGroupsRequest) *ListGroupsInvoker {
	requestDef := GenReqDefForListGroups()
	return &ListGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PatchGroup 部分更新用户组
//
// 修改现有用户组的部分属性，和管理用户组中的用户。
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 请求中只允许使用displayName、 members和externalId属性
// - 单个请求中最多允许100个成员更改
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) PatchGroup(request *model.PatchGroupRequest) (*model.PatchGroupResponse, error) {
	requestDef := GenReqDefForPatchGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PatchGroupResponse), nil
	}
}

// PatchGroupInvoker 部分更新用户组
func (c *IdentityCenterSCIMClient) PatchGroupInvoker(request *model.PatchGroupRequest) *PatchGroupInvoker {
	requestDef := GenReqDefForPatchGroup()
	return &PatchGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ServiceProviderConfig 查询服务提供商配置
//
// 查询IAM身份中心的SCIM相关配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) ServiceProviderConfig(request *model.ServiceProviderConfigRequest) (*model.ServiceProviderConfigResponse, error) {
	requestDef := GenReqDefForServiceProviderConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ServiceProviderConfigResponse), nil
	}
}

// ServiceProviderConfigInvoker 查询服务提供商配置
func (c *IdentityCenterSCIMClient) ServiceProviderConfigInvoker(request *model.ServiceProviderConfigRequest) *ServiceProviderConfigInvoker {
	requestDef := GenReqDefForServiceProviderConfig()
	return &ServiceProviderConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUser 创建用户
//
// 使用SCIM协议，同步用户到IAM身份中心。
// 暂不支持
// IAM身份中心暂不支持此API操作的以下方面。
// - ims、photos、x509Certificates、entitlements和password属性
// - manager的displayName子属性
// - emails、addresses和phoneNumbers的display子属性
//
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 必须填写userName、displayName、schemas属性，name属性中的familyName、givenName子属性，emails属性中的value、primary、type子属性
// - addresses可以包含字母、重音字符、符号、数字、标点符号、空格（正常和不换行）
// - 我们不支持多值属性中的多个值（例如emails,addresses,phoneNumbers）。只允许单个值
// - emails属性值必须标记为primary
// - 不能指定groups字段
// - userName字段可以包含字母、数字和部分符号_+&#x3D;,.@-
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserResponse), nil
	}
}

// CreateUserInvoker 创建用户
func (c *IdentityCenterSCIMClient) CreateUserInvoker(request *model.CreateUserRequest) *CreateUserInvoker {
	requestDef := GenReqDefForCreateUser()
	return &CreateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除用户
//
// 删除现有用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除用户
func (c *IdentityCenterSCIMClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetUser 查询用户详情
//
// 查询现有用户的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) GetUser(request *model.GetUserRequest) (*model.GetUserResponse, error) {
	requestDef := GenReqDefForGetUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetUserResponse), nil
	}
}

// GetUserInvoker 查询用户详情
func (c *IdentityCenterSCIMClient) GetUserInvoker(request *model.GetUserRequest) *GetUserInvoker {
	requestDef := GenReqDefForGetUser()
	return &GetUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 列出用户
//
// 对现有用户列表执行筛选查询，最多只能返回50个结果。
// 暂不支持
// IAM身份中心不支持此API操作的以下方面。
// - startIndex,attributes,excludedAttributes（虽在SCIM协议中列出）
//
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 目前，列出用户接口最多只能返回50个结果
// - 支持的筛选器组合：(userName)、(id)。请注意，使用id作为单个过滤器虽然有效，但应避免使用，因为已经有一个查询用户详情接口可用
// - 过滤器中支持的比较运算符：eq
// - 必须按如下方式指定筛选器：&lt;filterAttribute&gt; eq \&quot;&lt;filterValue&gt;\&quot;
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 列出用户
func (c *IdentityCenterSCIMClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PatchUser 部分更新用户
//
// 修改现有用户的部分属性。
// 暂不支持
// IAM身份中心暂不支持此API操作的以下方面。
// - 对userName或active属性执行多个PATCH操作
// - ims、photos、x509Certificates、entitlements和password属性
// - manager的displayName子属性
// - emails、addresses和phoneNumbers的display子属性
//
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 支持的操作是add、replace和remove
// - 必须指定操作
// - remove操作需要指定属性路径
// - add和replace操作需要指定属性的值
// - 仅允许修改userName、active、externalId、displayName、nickName、profileUrl、title、userType、preferredLanguage、locale、timezone、name、enterprise、emails、addresses和phoneNumbers属性
// - 过滤器中仅支持eq运算符
// - userName或active属性不支持remove操作
// - 我们不支持多值属性中的多个值（例如emails,addresses,phoneNumbers）。这些属性中的每一个都只允许有一个值
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) PatchUser(request *model.PatchUserRequest) (*model.PatchUserResponse, error) {
	requestDef := GenReqDefForPatchUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PatchUserResponse), nil
	}
}

// PatchUserInvoker 部分更新用户
func (c *IdentityCenterSCIMClient) PatchUserInvoker(request *model.PatchUserRequest) *PatchUserInvoker {
	requestDef := GenReqDefForPatchUser()
	return &PatchUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutUser 更新用户
//
// 更新现有用户。
// 暂不支持
// IAM身份中心暂不支持此API操作的以下方面。
// - ims、photos、x509Certificates、entitlements和password属性
// - manager的displayName子属性
// - emails、addresses和phoneNumbers的display子属性
//
// 约束条件
// IAM身份中心对此API操作具有以下约束。
// - 必须填写userName、displayName、schemas属性，name属性中的familyName、givenName子属性，emails属性中的value、primary、type子属性
// - addresses可以包含字母、重音字符、符号、数字、标点符号、空格（正常和不换行）
// - 我们不支持多值属性中的多个值（例如emails,addresses,phoneNumbers）
// - emails属性值必须标记为primary
// - 不能指定groups属性
// - userName字段可以包含字母、数字和部分符号_+&#x3D;,.@-
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterSCIMClient) PutUser(request *model.PutUserRequest) (*model.PutUserResponse, error) {
	requestDef := GenReqDefForPutUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutUserResponse), nil
	}
}

// PutUserInvoker 更新用户
func (c *IdentityCenterSCIMClient) PutUserInvoker(request *model.PutUserRequest) *PutUserInvoker {
	requestDef := GenReqDefForPutUser()
	return &PutUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
