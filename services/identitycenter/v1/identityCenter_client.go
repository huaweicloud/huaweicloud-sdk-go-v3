package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type IdentityCenterClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterClient(hcClient *httpclient.HcHttpClient) *IdentityCenterClient {
	return &IdentityCenterClient{HcClient: hcClient}
}

func IdentityCenterClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAccountAssignment 创建账号分配
//
// 使用指定的权限集为指定账号分配对应主体的访问权限，主体可为IAM身份中心用户或用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateAccountAssignment(request *model.CreateAccountAssignmentRequest) (*model.CreateAccountAssignmentResponse, error) {
	requestDef := GenReqDefForCreateAccountAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccountAssignmentResponse), nil
	}
}

// CreateAccountAssignmentInvoker 创建账号分配
func (c *IdentityCenterClient) CreateAccountAssignmentInvoker(request *model.CreateAccountAssignmentRequest) *CreateAccountAssignmentInvoker {
	requestDef := GenReqDefForCreateAccountAssignment()
	return &CreateAccountAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccountAssignment 删除账号分配
//
// 使用指定的权限集从指定的账号删除主体的访问权限，主体可为IAM身份中心用户或用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteAccountAssignment(request *model.DeleteAccountAssignmentRequest) (*model.DeleteAccountAssignmentResponse, error) {
	requestDef := GenReqDefForDeleteAccountAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccountAssignmentResponse), nil
	}
}

// DeleteAccountAssignmentInvoker 删除账号分配
func (c *IdentityCenterClient) DeleteAccountAssignmentInvoker(request *model.DeleteAccountAssignmentRequest) *DeleteAccountAssignmentInvoker {
	requestDef := GenReqDefForDeleteAccountAssignment()
	return &DeleteAccountAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeAccountAssignmentCreationStatus 查询账号分配创建状态详情
//
// 根据请求ID，查询指定IAM身份中心实例下的账号分配创建状态详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeAccountAssignmentCreationStatus(request *model.DescribeAccountAssignmentCreationStatusRequest) (*model.DescribeAccountAssignmentCreationStatusResponse, error) {
	requestDef := GenReqDefForDescribeAccountAssignmentCreationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeAccountAssignmentCreationStatusResponse), nil
	}
}

// DescribeAccountAssignmentCreationStatusInvoker 查询账号分配创建状态详情
func (c *IdentityCenterClient) DescribeAccountAssignmentCreationStatusInvoker(request *model.DescribeAccountAssignmentCreationStatusRequest) *DescribeAccountAssignmentCreationStatusInvoker {
	requestDef := GenReqDefForDescribeAccountAssignmentCreationStatus()
	return &DescribeAccountAssignmentCreationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeAccountAssignmentDeletionStatus 查询账号分配删除状态详情
//
// 根据请求ID，查询指定IAM身份中心实例下的账号分配删除状态详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeAccountAssignmentDeletionStatus(request *model.DescribeAccountAssignmentDeletionStatusRequest) (*model.DescribeAccountAssignmentDeletionStatusResponse, error) {
	requestDef := GenReqDefForDescribeAccountAssignmentDeletionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeAccountAssignmentDeletionStatusResponse), nil
	}
}

// DescribeAccountAssignmentDeletionStatusInvoker 查询账号分配删除状态详情
func (c *IdentityCenterClient) DescribeAccountAssignmentDeletionStatusInvoker(request *model.DescribeAccountAssignmentDeletionStatusRequest) *DescribeAccountAssignmentDeletionStatusInvoker {
	requestDef := GenReqDefForDescribeAccountAssignmentDeletionStatus()
	return &DescribeAccountAssignmentDeletionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentCreationStatus 列出账号分配创建状态
//
// 查询指定IAM身份中心实例下的账号分配的创建状态列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListAccountAssignmentCreationStatus(request *model.ListAccountAssignmentCreationStatusRequest) (*model.ListAccountAssignmentCreationStatusResponse, error) {
	requestDef := GenReqDefForListAccountAssignmentCreationStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountAssignmentCreationStatusResponse), nil
	}
}

// ListAccountAssignmentCreationStatusInvoker 列出账号分配创建状态
func (c *IdentityCenterClient) ListAccountAssignmentCreationStatusInvoker(request *model.ListAccountAssignmentCreationStatusRequest) *ListAccountAssignmentCreationStatusInvoker {
	requestDef := GenReqDefForListAccountAssignmentCreationStatus()
	return &ListAccountAssignmentCreationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentDeletionStatus 列出账号分配删除状态
//
// 查询指定IAM身份中心实例下的账号分配的删除状态列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListAccountAssignmentDeletionStatus(request *model.ListAccountAssignmentDeletionStatusRequest) (*model.ListAccountAssignmentDeletionStatusResponse, error) {
	requestDef := GenReqDefForListAccountAssignmentDeletionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountAssignmentDeletionStatusResponse), nil
	}
}

// ListAccountAssignmentDeletionStatusInvoker 列出账号分配删除状态
func (c *IdentityCenterClient) ListAccountAssignmentDeletionStatusInvoker(request *model.ListAccountAssignmentDeletionStatusRequest) *ListAccountAssignmentDeletionStatusInvoker {
	requestDef := GenReqDefForListAccountAssignmentDeletionStatus()
	return &ListAccountAssignmentDeletionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignments 列出账号和权限集关联的用户或用户组
//
// 列出与指定账号以及指定权限集关联的用户或用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListAccountAssignments(request *model.ListAccountAssignmentsRequest) (*model.ListAccountAssignmentsResponse, error) {
	requestDef := GenReqDefForListAccountAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountAssignmentsResponse), nil
	}
}

// ListAccountAssignmentsInvoker 列出账号和权限集关联的用户或用户组
func (c *IdentityCenterClient) ListAccountAssignmentsInvoker(request *model.ListAccountAssignmentsRequest) *ListAccountAssignmentsInvoker {
	requestDef := GenReqDefForListAccountAssignments()
	return &ListAccountAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 列出实例
//
// 查询IAM身份中心的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 列出实例
func (c *IdentityCenterClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceAccessControlAttributeConfiguration 启用指定实例的访问控制功能
//
// 启用指定实例的访问控制功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateInstanceAccessControlAttributeConfiguration(request *model.CreateInstanceAccessControlAttributeConfigurationRequest) (*model.CreateInstanceAccessControlAttributeConfigurationResponse, error) {
	requestDef := GenReqDefForCreateInstanceAccessControlAttributeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

// CreateInstanceAccessControlAttributeConfigurationInvoker 启用指定实例的访问控制功能
func (c *IdentityCenterClient) CreateInstanceAccessControlAttributeConfigurationInvoker(request *model.CreateInstanceAccessControlAttributeConfigurationRequest) *CreateInstanceAccessControlAttributeConfigurationInvoker {
	requestDef := GenReqDefForCreateInstanceAccessControlAttributeConfiguration()
	return &CreateInstanceAccessControlAttributeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceAccessControlAttributeConfiguration 解除指定实例的访问控制属性配置
//
// 禁用指定IAM身份中心实例的基于属性的访问控制（ABAC）功能，并删除已配置的所有属性映射。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteInstanceAccessControlAttributeConfiguration(request *model.DeleteInstanceAccessControlAttributeConfigurationRequest) (*model.DeleteInstanceAccessControlAttributeConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteInstanceAccessControlAttributeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

// DeleteInstanceAccessControlAttributeConfigurationInvoker 解除指定实例的访问控制属性配置
func (c *IdentityCenterClient) DeleteInstanceAccessControlAttributeConfigurationInvoker(request *model.DeleteInstanceAccessControlAttributeConfigurationRequest) *DeleteInstanceAccessControlAttributeConfigurationInvoker {
	requestDef := GenReqDefForDeleteInstanceAccessControlAttributeConfiguration()
	return &DeleteInstanceAccessControlAttributeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeInstanceAccessControlAttributeConfiguration 获取指定实例的访问控制属性配置
//
// 返回已配置为与指定IAM身份中心实例的基于属性的访问控制（ABAC）一起使用的IAM身份中心身份源属性列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeInstanceAccessControlAttributeConfiguration(request *model.DescribeInstanceAccessControlAttributeConfigurationRequest) (*model.DescribeInstanceAccessControlAttributeConfigurationResponse, error) {
	requestDef := GenReqDefForDescribeInstanceAccessControlAttributeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

// DescribeInstanceAccessControlAttributeConfigurationInvoker 获取指定实例的访问控制属性配置
func (c *IdentityCenterClient) DescribeInstanceAccessControlAttributeConfigurationInvoker(request *model.DescribeInstanceAccessControlAttributeConfigurationRequest) *DescribeInstanceAccessControlAttributeConfigurationInvoker {
	requestDef := GenReqDefForDescribeInstanceAccessControlAttributeConfiguration()
	return &DescribeInstanceAccessControlAttributeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceAccessControlAttributeConfiguration 更新指定实例的访问控制属性配置
//
// 更新可与IAM身份中心实例一起使用的IAM身份中心身份源属性，以进行基于属性的访问控制（ABAC）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateInstanceAccessControlAttributeConfiguration(request *model.UpdateInstanceAccessControlAttributeConfigurationRequest) (*model.UpdateInstanceAccessControlAttributeConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateInstanceAccessControlAttributeConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceAccessControlAttributeConfigurationResponse), nil
	}
}

// UpdateInstanceAccessControlAttributeConfigurationInvoker 更新指定实例的访问控制属性配置
func (c *IdentityCenterClient) UpdateInstanceAccessControlAttributeConfigurationInvoker(request *model.UpdateInstanceAccessControlAttributeConfigurationRequest) *UpdateInstanceAccessControlAttributeConfigurationInvoker {
	requestDef := GenReqDefForUpdateInstanceAccessControlAttributeConfiguration()
	return &UpdateInstanceAccessControlAttributeConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachManagedPolicyToPermissionSet 添加系统身份策略
//
// 在指定的权限集中添加系统身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) AttachManagedPolicyToPermissionSet(request *model.AttachManagedPolicyToPermissionSetRequest) (*model.AttachManagedPolicyToPermissionSetResponse, error) {
	requestDef := GenReqDefForAttachManagedPolicyToPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachManagedPolicyToPermissionSetResponse), nil
	}
}

// AttachManagedPolicyToPermissionSetInvoker 添加系统身份策略
func (c *IdentityCenterClient) AttachManagedPolicyToPermissionSetInvoker(request *model.AttachManagedPolicyToPermissionSetRequest) *AttachManagedPolicyToPermissionSetInvoker {
	requestDef := GenReqDefForAttachManagedPolicyToPermissionSet()
	return &AttachManagedPolicyToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePermissionSet 创建权限集
//
// 在指定的IAM身份中心实例中创建一个权限集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreatePermissionSet(request *model.CreatePermissionSetRequest) (*model.CreatePermissionSetResponse, error) {
	requestDef := GenReqDefForCreatePermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePermissionSetResponse), nil
	}
}

// CreatePermissionSetInvoker 创建权限集
func (c *IdentityCenterClient) CreatePermissionSetInvoker(request *model.CreatePermissionSetRequest) *CreatePermissionSetInvoker {
	requestDef := GenReqDefForCreatePermissionSet()
	return &CreatePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomPolicyFromPermissionSet 删除自定义身份策略
//
// 删除指定权限集中的自定义身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteCustomPolicyFromPermissionSet(request *model.DeleteCustomPolicyFromPermissionSetRequest) (*model.DeleteCustomPolicyFromPermissionSetResponse, error) {
	requestDef := GenReqDefForDeleteCustomPolicyFromPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomPolicyFromPermissionSetResponse), nil
	}
}

// DeleteCustomPolicyFromPermissionSetInvoker 删除自定义身份策略
func (c *IdentityCenterClient) DeleteCustomPolicyFromPermissionSetInvoker(request *model.DeleteCustomPolicyFromPermissionSetRequest) *DeleteCustomPolicyFromPermissionSetInvoker {
	requestDef := GenReqDefForDeleteCustomPolicyFromPermissionSet()
	return &DeleteCustomPolicyFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomRoleFromPermissionSet 删除指定权限集中的自定义策略
//
// 删除指定权限集中的自定义策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteCustomRoleFromPermissionSet(request *model.DeleteCustomRoleFromPermissionSetRequest) (*model.DeleteCustomRoleFromPermissionSetResponse, error) {
	requestDef := GenReqDefForDeleteCustomRoleFromPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomRoleFromPermissionSetResponse), nil
	}
}

// DeleteCustomRoleFromPermissionSetInvoker 删除指定权限集中的自定义策略
func (c *IdentityCenterClient) DeleteCustomRoleFromPermissionSetInvoker(request *model.DeleteCustomRoleFromPermissionSetRequest) *DeleteCustomRoleFromPermissionSetInvoker {
	requestDef := GenReqDefForDeleteCustomRoleFromPermissionSet()
	return &DeleteCustomRoleFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePermissionSet 删除权限集
//
// 根据权限集ID，删除指定的权限集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeletePermissionSet(request *model.DeletePermissionSetRequest) (*model.DeletePermissionSetResponse, error) {
	requestDef := GenReqDefForDeletePermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePermissionSetResponse), nil
	}
}

// DeletePermissionSetInvoker 删除权限集
func (c *IdentityCenterClient) DeletePermissionSetInvoker(request *model.DeletePermissionSetRequest) *DeletePermissionSetInvoker {
	requestDef := GenReqDefForDeletePermissionSet()
	return &DeletePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribePermissionSet 查询权限集详情
//
// 根据权限集ID，查询指定权限集的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribePermissionSet(request *model.DescribePermissionSetRequest) (*model.DescribePermissionSetResponse, error) {
	requestDef := GenReqDefForDescribePermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribePermissionSetResponse), nil
	}
}

// DescribePermissionSetInvoker 查询权限集详情
func (c *IdentityCenterClient) DescribePermissionSetInvoker(request *model.DescribePermissionSetRequest) *DescribePermissionSetInvoker {
	requestDef := GenReqDefForDescribePermissionSet()
	return &DescribePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribePermissionSetProvisioningStatus 查询权限集预分配状态详情
//
// 根据请求ID，查询权限集预分配状态的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribePermissionSetProvisioningStatus(request *model.DescribePermissionSetProvisioningStatusRequest) (*model.DescribePermissionSetProvisioningStatusResponse, error) {
	requestDef := GenReqDefForDescribePermissionSetProvisioningStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribePermissionSetProvisioningStatusResponse), nil
	}
}

// DescribePermissionSetProvisioningStatusInvoker 查询权限集预分配状态详情
func (c *IdentityCenterClient) DescribePermissionSetProvisioningStatusInvoker(request *model.DescribePermissionSetProvisioningStatusRequest) *DescribePermissionSetProvisioningStatusInvoker {
	requestDef := GenReqDefForDescribePermissionSetProvisioningStatus()
	return &DescribePermissionSetProvisioningStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachManagedPolicyFromPermissionSet 删除系统身份策略
//
// 删除指定权限集中的系统身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DetachManagedPolicyFromPermissionSet(request *model.DetachManagedPolicyFromPermissionSetRequest) (*model.DetachManagedPolicyFromPermissionSetResponse, error) {
	requestDef := GenReqDefForDetachManagedPolicyFromPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachManagedPolicyFromPermissionSetResponse), nil
	}
}

// DetachManagedPolicyFromPermissionSetInvoker 删除系统身份策略
func (c *IdentityCenterClient) DetachManagedPolicyFromPermissionSetInvoker(request *model.DetachManagedPolicyFromPermissionSetRequest) *DetachManagedPolicyFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachManagedPolicyFromPermissionSet()
	return &DetachManagedPolicyFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetCustomPolicyForPermissionSet 查询自定义身份策略详情
//
// 查询指定权限集中的自定义身份策略详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetCustomPolicyForPermissionSet(request *model.GetCustomPolicyForPermissionSetRequest) (*model.GetCustomPolicyForPermissionSetResponse, error) {
	requestDef := GenReqDefForGetCustomPolicyForPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetCustomPolicyForPermissionSetResponse), nil
	}
}

// GetCustomPolicyForPermissionSetInvoker 查询自定义身份策略详情
func (c *IdentityCenterClient) GetCustomPolicyForPermissionSetInvoker(request *model.GetCustomPolicyForPermissionSetRequest) *GetCustomPolicyForPermissionSetInvoker {
	requestDef := GenReqDefForGetCustomPolicyForPermissionSet()
	return &GetCustomPolicyForPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetCustomRoleForPermissionSet 获取分配给权限集的自定义策略
//
// 获取分配给权限集的自定义策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetCustomRoleForPermissionSet(request *model.GetCustomRoleForPermissionSetRequest) (*model.GetCustomRoleForPermissionSetResponse, error) {
	requestDef := GenReqDefForGetCustomRoleForPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetCustomRoleForPermissionSetResponse), nil
	}
}

// GetCustomRoleForPermissionSetInvoker 获取分配给权限集的自定义策略
func (c *IdentityCenterClient) GetCustomRoleForPermissionSetInvoker(request *model.GetCustomRoleForPermissionSetRequest) *GetCustomRoleForPermissionSetInvoker {
	requestDef := GenReqDefForGetCustomRoleForPermissionSet()
	return &GetCustomRoleForPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountsForProvisionedPermissionSet 列出权限集关联的账号
//
// 查询与指定权限集关联的账号列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListAccountsForProvisionedPermissionSet(request *model.ListAccountsForProvisionedPermissionSetRequest) (*model.ListAccountsForProvisionedPermissionSetResponse, error) {
	requestDef := GenReqDefForListAccountsForProvisionedPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountsForProvisionedPermissionSetResponse), nil
	}
}

// ListAccountsForProvisionedPermissionSetInvoker 列出权限集关联的账号
func (c *IdentityCenterClient) ListAccountsForProvisionedPermissionSetInvoker(request *model.ListAccountsForProvisionedPermissionSetRequest) *ListAccountsForProvisionedPermissionSetInvoker {
	requestDef := GenReqDefForListAccountsForProvisionedPermissionSet()
	return &ListAccountsForProvisionedPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedPoliciesInPermissionSet 列出系统身份策略
//
// 获取添加到指定权限集中的系统身份策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListManagedPoliciesInPermissionSet(request *model.ListManagedPoliciesInPermissionSetRequest) (*model.ListManagedPoliciesInPermissionSetResponse, error) {
	requestDef := GenReqDefForListManagedPoliciesInPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedPoliciesInPermissionSetResponse), nil
	}
}

// ListManagedPoliciesInPermissionSetInvoker 列出系统身份策略
func (c *IdentityCenterClient) ListManagedPoliciesInPermissionSetInvoker(request *model.ListManagedPoliciesInPermissionSetRequest) *ListManagedPoliciesInPermissionSetInvoker {
	requestDef := GenReqDefForListManagedPoliciesInPermissionSet()
	return &ListManagedPoliciesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSetProvisioningStatus 列出权限集预分配状态
//
// 查询指定实例中的权限集预分配状态列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListPermissionSetProvisioningStatus(request *model.ListPermissionSetProvisioningStatusRequest) (*model.ListPermissionSetProvisioningStatusResponse, error) {
	requestDef := GenReqDefForListPermissionSetProvisioningStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionSetProvisioningStatusResponse), nil
	}
}

// ListPermissionSetProvisioningStatusInvoker 列出权限集预分配状态
func (c *IdentityCenterClient) ListPermissionSetProvisioningStatusInvoker(request *model.ListPermissionSetProvisioningStatusRequest) *ListPermissionSetProvisioningStatusInvoker {
	requestDef := GenReqDefForListPermissionSetProvisioningStatus()
	return &ListPermissionSetProvisioningStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSets 列出权限集
//
// 查询指定实例下的权限集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListPermissionSets(request *model.ListPermissionSetsRequest) (*model.ListPermissionSetsResponse, error) {
	requestDef := GenReqDefForListPermissionSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionSetsResponse), nil
	}
}

// ListPermissionSetsInvoker 列出权限集
func (c *IdentityCenterClient) ListPermissionSetsInvoker(request *model.ListPermissionSetsRequest) *ListPermissionSetsInvoker {
	requestDef := GenReqDefForListPermissionSets()
	return &ListPermissionSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSetsProvisionedToAccount 列出分配给账号的权限集
//
// 查询分配给指定账号的权限集列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListPermissionSetsProvisionedToAccount(request *model.ListPermissionSetsProvisionedToAccountRequest) (*model.ListPermissionSetsProvisionedToAccountResponse, error) {
	requestDef := GenReqDefForListPermissionSetsProvisionedToAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionSetsProvisionedToAccountResponse), nil
	}
}

// ListPermissionSetsProvisionedToAccountInvoker 列出分配给账号的权限集
func (c *IdentityCenterClient) ListPermissionSetsProvisionedToAccountInvoker(request *model.ListPermissionSetsProvisionedToAccountRequest) *ListPermissionSetsProvisionedToAccountInvoker {
	requestDef := GenReqDefForListPermissionSetsProvisionedToAccount()
	return &ListPermissionSetsProvisionedToAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ProvisionPermissionSet 预分配权限集
//
// 将指定权限集预分配给指定账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ProvisionPermissionSet(request *model.ProvisionPermissionSetRequest) (*model.ProvisionPermissionSetResponse, error) {
	requestDef := GenReqDefForProvisionPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ProvisionPermissionSetResponse), nil
	}
}

// ProvisionPermissionSetInvoker 预分配权限集
func (c *IdentityCenterClient) ProvisionPermissionSetInvoker(request *model.ProvisionPermissionSetRequest) *ProvisionPermissionSetInvoker {
	requestDef := GenReqDefForProvisionPermissionSet()
	return &ProvisionPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutCustomPolicyToPermissionSet 添加自定义身份策略
//
// 在指定的权限集中添加自定义身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) PutCustomPolicyToPermissionSet(request *model.PutCustomPolicyToPermissionSetRequest) (*model.PutCustomPolicyToPermissionSetResponse, error) {
	requestDef := GenReqDefForPutCustomPolicyToPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutCustomPolicyToPermissionSetResponse), nil
	}
}

// PutCustomPolicyToPermissionSetInvoker 添加自定义身份策略
func (c *IdentityCenterClient) PutCustomPolicyToPermissionSetInvoker(request *model.PutCustomPolicyToPermissionSetRequest) *PutCustomPolicyToPermissionSetInvoker {
	requestDef := GenReqDefForPutCustomPolicyToPermissionSet()
	return &PutCustomPolicyToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutCustomRoleToPermissionSet 将自定义策略附加到权限集
//
// 将自定义策略附加到权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) PutCustomRoleToPermissionSet(request *model.PutCustomRoleToPermissionSetRequest) (*model.PutCustomRoleToPermissionSetResponse, error) {
	requestDef := GenReqDefForPutCustomRoleToPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutCustomRoleToPermissionSetResponse), nil
	}
}

// PutCustomRoleToPermissionSetInvoker 将自定义策略附加到权限集
func (c *IdentityCenterClient) PutCustomRoleToPermissionSetInvoker(request *model.PutCustomRoleToPermissionSetRequest) *PutCustomRoleToPermissionSetInvoker {
	requestDef := GenReqDefForPutCustomRoleToPermissionSet()
	return &PutCustomRoleToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePermissionSet 更新权限集
//
// 根据权限集ID，更新指定权限集的属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdatePermissionSet(request *model.UpdatePermissionSetRequest) (*model.UpdatePermissionSetResponse, error) {
	requestDef := GenReqDefForUpdatePermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePermissionSetResponse), nil
	}
}

// UpdatePermissionSetInvoker 更新权限集
func (c *IdentityCenterClient) UpdatePermissionSetInvoker(request *model.UpdatePermissionSetRequest) *UpdatePermissionSetInvoker {
	requestDef := GenReqDefForUpdatePermissionSet()
	return &UpdatePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTagResource 为指定资源添加标签
//
// 给指定的资源添加一个或多个标签。当前支持为权限集添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateTagResource(request *model.CreateTagResourceRequest) (*model.CreateTagResourceResponse, error) {
	requestDef := GenReqDefForCreateTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResourceResponse), nil
	}
}

// CreateTagResourceInvoker 为指定资源添加标签
func (c *IdentityCenterClient) CreateTagResourceInvoker(request *model.CreateTagResourceRequest) *CreateTagResourceInvoker {
	requestDef := GenReqDefForCreateTagResource()
	return &CreateTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTagResource 删除指定资源的指定标签
//
// 从指定资源中删除指定的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteTagResource(request *model.DeleteTagResourceRequest) (*model.DeleteTagResourceResponse, error) {
	requestDef := GenReqDefForDeleteTagResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResourceResponse), nil
	}
}

// DeleteTagResourceInvoker 删除指定资源的指定标签
func (c *IdentityCenterClient) DeleteTagResourceInvoker(request *model.DeleteTagResourceRequest) *DeleteTagResourceInvoker {
	requestDef := GenReqDefForDeleteTagResource()
	return &DeleteTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagResources 列出绑定到指定资源的标签
//
// 列出绑定到指定资源的标签。当前支持为权限集添加标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListTagResources(request *model.ListTagResourcesRequest) (*model.ListTagResourcesResponse, error) {
	requestDef := GenReqDefForListTagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResourcesResponse), nil
	}
}

// ListTagResourcesInvoker 列出绑定到指定资源的标签
func (c *IdentityCenterClient) ListTagResourcesInvoker(request *model.ListTagResourcesRequest) *ListTagResourcesInvoker {
	requestDef := GenReqDefForListTagResources()
	return &ListTagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachManagedRoleToPermissionSet 添加系统策略
//
// 在指定的权限集中添加系统策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) AttachManagedRoleToPermissionSet(request *model.AttachManagedRoleToPermissionSetRequest) (*model.AttachManagedRoleToPermissionSetResponse, error) {
	requestDef := GenReqDefForAttachManagedRoleToPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachManagedRoleToPermissionSetResponse), nil
	}
}

// AttachManagedRoleToPermissionSetInvoker 添加系统策略
func (c *IdentityCenterClient) AttachManagedRoleToPermissionSetInvoker(request *model.AttachManagedRoleToPermissionSetRequest) *AttachManagedRoleToPermissionSetInvoker {
	requestDef := GenReqDefForAttachManagedRoleToPermissionSet()
	return &AttachManagedRoleToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachManagedRoleFromPermissionSet 删除系统策略
//
// 删除指定权限集中的系统策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DetachManagedRoleFromPermissionSet(request *model.DetachManagedRoleFromPermissionSetRequest) (*model.DetachManagedRoleFromPermissionSetResponse, error) {
	requestDef := GenReqDefForDetachManagedRoleFromPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachManagedRoleFromPermissionSetResponse), nil
	}
}

// DetachManagedRoleFromPermissionSetInvoker 删除系统策略
func (c *IdentityCenterClient) DetachManagedRoleFromPermissionSetInvoker(request *model.DetachManagedRoleFromPermissionSetRequest) *DetachManagedRoleFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachManagedRoleFromPermissionSet()
	return &DetachManagedRoleFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedRolesInPermissionSet 列出系统策略
//
// 获取添加到指定权限集中的系统策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListManagedRolesInPermissionSet(request *model.ListManagedRolesInPermissionSetRequest) (*model.ListManagedRolesInPermissionSetResponse, error) {
	requestDef := GenReqDefForListManagedRolesInPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagedRolesInPermissionSetResponse), nil
	}
}

// ListManagedRolesInPermissionSetInvoker 列出系统策略
func (c *IdentityCenterClient) ListManagedRolesInPermissionSetInvoker(request *model.ListManagedRolesInPermissionSetRequest) *ListManagedRolesInPermissionSetInvoker {
	requestDef := GenReqDefForListManagedRolesInPermissionSet()
	return &ListManagedRolesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
