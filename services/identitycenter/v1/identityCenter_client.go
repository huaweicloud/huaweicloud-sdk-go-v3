package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type IdentityCenterClient struct {
	HcClient *http_client.HcHttpClient
}

func NewIdentityCenterClient(hcClient *http_client.HcHttpClient) *IdentityCenterClient {
	return &IdentityCenterClient{HcClient: hcClient}
}

func IdentityCenterClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAccountAssignment 创建帐户分配
//
// 使用指定的权限集为指定帐户分配对应主体的访问权限，主体可为IdentityCenter用户或IdentityCenter用户组
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

// CreateAccountAssignmentInvoker 创建帐户分配
func (c *IdentityCenterClient) CreateAccountAssignmentInvoker(request *model.CreateAccountAssignmentRequest) *CreateAccountAssignmentInvoker {
	requestDef := GenReqDefForCreateAccountAssignment()
	return &CreateAccountAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccountAssignment 使用指定的权限集从指定帐户删除主体的访问权限
//
// 使用指定的权限集从指定帐户删除主体的访问权限
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

// DeleteAccountAssignmentInvoker 使用指定的权限集从指定帐户删除主体的访问权限
func (c *IdentityCenterClient) DeleteAccountAssignmentInvoker(request *model.DeleteAccountAssignmentRequest) *DeleteAccountAssignmentInvoker {
	requestDef := GenReqDefForDeleteAccountAssignment()
	return &DeleteAccountAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeAccountAssignmentCreationStatus 查询账号分配请求的状态
//
// 根据请求ID，查询指定IAM Identity Center实例下的帐户分配创建状态详情信息
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

// DescribeAccountAssignmentCreationStatusInvoker 查询账号分配请求的状态
func (c *IdentityCenterClient) DescribeAccountAssignmentCreationStatusInvoker(request *model.DescribeAccountAssignmentCreationStatusRequest) *DescribeAccountAssignmentCreationStatusInvoker {
	requestDef := GenReqDefForDescribeAccountAssignmentCreationStatus()
	return &DescribeAccountAssignmentCreationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeAccountAssignmentDeletionStatus 查询账户分配删除状态详情
//
// 根据请求ID，查询指定IAM Identity Center实例下的帐户分配删除状态详情信息
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

// DescribeAccountAssignmentDeletionStatusInvoker 查询账户分配删除状态详情
func (c *IdentityCenterClient) DescribeAccountAssignmentDeletionStatusInvoker(request *model.DescribeAccountAssignmentDeletionStatusRequest) *DescribeAccountAssignmentDeletionStatusInvoker {
	requestDef := GenReqDefForDescribeAccountAssignmentDeletionStatus()
	return &DescribeAccountAssignmentDeletionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentCreationStatus 列出账户分配创建状态
//
// 查询指定IAM Identity Center实例下的帐户分配的创建状态列表
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

// ListAccountAssignmentCreationStatusInvoker 列出账户分配创建状态
func (c *IdentityCenterClient) ListAccountAssignmentCreationStatusInvoker(request *model.ListAccountAssignmentCreationStatusRequest) *ListAccountAssignmentCreationStatusInvoker {
	requestDef := GenReqDefForListAccountAssignmentCreationStatus()
	return &ListAccountAssignmentCreationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentDeletionStatus 列出账户分配删除状态
//
// 查询指定IAM Identity Center实例下的帐户分配的删除状态列表
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

// ListAccountAssignmentDeletionStatusInvoker 列出账户分配删除状态
func (c *IdentityCenterClient) ListAccountAssignmentDeletionStatusInvoker(request *model.ListAccountAssignmentDeletionStatusRequest) *ListAccountAssignmentDeletionStatusInvoker {
	requestDef := GenReqDefForListAccountAssignmentDeletionStatus()
	return &ListAccountAssignmentDeletionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignments 列出与指定权限集以及指定帐户关联的用户/用户组
//
// 列出与指定权限集以及指定帐户关联的用户/用户组
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

// ListAccountAssignmentsInvoker 列出与指定权限集以及指定帐户关联的用户/用户组
func (c *IdentityCenterClient) ListAccountAssignmentsInvoker(request *model.ListAccountAssignmentsRequest) *ListAccountAssignmentsInvoker {
	requestDef := GenReqDefForListAccountAssignments()
	return &ListAccountAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProfileAssociations 查询profile关联的用户或组列表
//
// 查询profile关联的用户或组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListProfileAssociations(request *model.ListProfileAssociationsRequest) (*model.ListProfileAssociationsResponse, error) {
	requestDef := GenReqDefForListProfileAssociations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProfileAssociationsResponse), nil
	}
}

// ListProfileAssociationsInvoker 查询profile关联的用户或组列表
func (c *IdentityCenterClient) ListProfileAssociationsInvoker(request *model.ListProfileAssociationsRequest) *ListProfileAssociationsInvoker {
	requestDef := GenReqDefForListProfileAssociations()
	return &ListProfileAssociationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询实例列表
//
// 查询 IAM Identity Center的资源实例列表.
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

// ListInstancesInvoker 查询实例列表
func (c *IdentityCenterClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachCustomerManagedPolicyToPermissionSet 将指定的客户自定义管理策略附加到指定的权限集
//
// 将指定的客户自定义管理策略附加到指定的权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) AttachCustomerManagedPolicyToPermissionSet(request *model.AttachCustomerManagedPolicyToPermissionSetRequest) (*model.AttachCustomerManagedPolicyToPermissionSetResponse, error) {
	requestDef := GenReqDefForAttachCustomerManagedPolicyToPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachCustomerManagedPolicyToPermissionSetResponse), nil
	}
}

// AttachCustomerManagedPolicyToPermissionSetInvoker 将指定的客户自定义管理策略附加到指定的权限集
func (c *IdentityCenterClient) AttachCustomerManagedPolicyToPermissionSetInvoker(request *model.AttachCustomerManagedPolicyToPermissionSetRequest) *AttachCustomerManagedPolicyToPermissionSetInvoker {
	requestDef := GenReqDefForAttachCustomerManagedPolicyToPermissionSet()
	return &AttachCustomerManagedPolicyToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachManagedPolicyToPermissionSet 将系统管理策略附加到权限集
//
// 将系统管理策略附加到权限集
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

// AttachManagedPolicyToPermissionSetInvoker 将系统管理策略附加到权限集
func (c *IdentityCenterClient) AttachManagedPolicyToPermissionSetInvoker(request *model.AttachManagedPolicyToPermissionSetRequest) *AttachManagedPolicyToPermissionSetInvoker {
	requestDef := GenReqDefForAttachManagedPolicyToPermissionSet()
	return &AttachManagedPolicyToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePermissionSet 在指定的IAM Identity Center实例中创建权限集
//
// 在指定的IAM Identity Center实例中创建一个权限集
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

// CreatePermissionSetInvoker 在指定的IAM Identity Center实例中创建权限集
func (c *IdentityCenterClient) CreatePermissionSetInvoker(request *model.CreatePermissionSetRequest) *CreatePermissionSetInvoker {
	requestDef := GenReqDefForCreatePermissionSet()
	return &CreatePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePermissionSet 删除指定实例的权限集
//
// 删除指定实例的权限集
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

// DeletePermissionSetInvoker 删除指定实例的权限集
func (c *IdentityCenterClient) DeletePermissionSetInvoker(request *model.DeletePermissionSetRequest) *DeletePermissionSetInvoker {
	requestDef := GenReqDefForDeletePermissionSet()
	return &DeletePermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribePermissionSet 查询权限集详情
//
// 根据权限集ID，查询指定权限集的详情信息
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
// 根据请求Id，查询权限集预分配状态的详情信息
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

// DetachCustomerManagedPolicyReferenceFromPermissionSet 将附加的客户自定义管理策略从指定的权限集中分离
//
// 将附加的客户自定义管理策略从指定的权限集中分离
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DetachCustomerManagedPolicyReferenceFromPermissionSet(request *model.DetachCustomerManagedPolicyReferenceFromPermissionSetRequest) (*model.DetachCustomerManagedPolicyReferenceFromPermissionSetResponse, error) {
	requestDef := GenReqDefForDetachCustomerManagedPolicyReferenceFromPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachCustomerManagedPolicyReferenceFromPermissionSetResponse), nil
	}
}

// DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker 将附加的客户自定义管理策略从指定的权限集中分离
func (c *IdentityCenterClient) DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker(request *model.DetachCustomerManagedPolicyReferenceFromPermissionSetRequest) *DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachCustomerManagedPolicyReferenceFromPermissionSet()
	return &DetachCustomerManagedPolicyReferenceFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachManagedPolicyFromPermissionSet 将附加的系统策略从指定的权限集中分离
//
// 将附加的系统策略从指定的权限集中分离
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

// DetachManagedPolicyFromPermissionSetInvoker 将附加的系统策略从指定的权限集中分离
func (c *IdentityCenterClient) DetachManagedPolicyFromPermissionSetInvoker(request *model.DetachManagedPolicyFromPermissionSetRequest) *DetachManagedPolicyFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachManagedPolicyFromPermissionSet()
	return &DetachManagedPolicyFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountsForProvisionedPermissionSet 列出权限集关联的账户
//
// 查询与指定权限集关联的账户列表
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

// ListAccountsForProvisionedPermissionSetInvoker 列出权限集关联的账户
func (c *IdentityCenterClient) ListAccountsForProvisionedPermissionSetInvoker(request *model.ListAccountsForProvisionedPermissionSetRequest) *ListAccountsForProvisionedPermissionSetInvoker {
	requestDef := GenReqDefForListAccountsForProvisionedPermissionSet()
	return &ListAccountsForProvisionedPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCustomerManagedPolicyReferencesInPermissionSet 获取附加到指定权限集的所有客户自定义策略列表
//
// 获取附加到指定权限集的所有客户自定义策略列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListCustomerManagedPolicyReferencesInPermissionSet(request *model.ListCustomerManagedPolicyReferencesInPermissionSetRequest) (*model.ListCustomerManagedPolicyReferencesInPermissionSetResponse, error) {
	requestDef := GenReqDefForListCustomerManagedPolicyReferencesInPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomerManagedPolicyReferencesInPermissionSetResponse), nil
	}
}

// ListCustomerManagedPolicyReferencesInPermissionSetInvoker 获取附加到指定权限集的所有客户自定义策略列表
func (c *IdentityCenterClient) ListCustomerManagedPolicyReferencesInPermissionSetInvoker(request *model.ListCustomerManagedPolicyReferencesInPermissionSetRequest) *ListCustomerManagedPolicyReferencesInPermissionSetInvoker {
	requestDef := GenReqDefForListCustomerManagedPolicyReferencesInPermissionSet()
	return &ListCustomerManagedPolicyReferencesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedPoliciesInPermissionSet 列出权限集中附加的系统管理策略
//
// 列出权限集中附加的系统管理策略
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

// ListManagedPoliciesInPermissionSetInvoker 列出权限集中附加的系统管理策略
func (c *IdentityCenterClient) ListManagedPoliciesInPermissionSetInvoker(request *model.ListManagedPoliciesInPermissionSetRequest) *ListManagedPoliciesInPermissionSetInvoker {
	requestDef := GenReqDefForListManagedPoliciesInPermissionSet()
	return &ListManagedPoliciesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSetProvisioningStatus 列出权限集预分配状态
//
// 查询指定实例中的权限集预分配状态列表
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

// ListPermissionSets 列出指定实例的权限集列表
//
// 列出指定实例的权限集列表
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

// ListPermissionSetsInvoker 列出指定实例的权限集列表
func (c *IdentityCenterClient) ListPermissionSetsInvoker(request *model.ListPermissionSetsRequest) *ListPermissionSetsInvoker {
	requestDef := GenReqDefForListPermissionSets()
	return &ListPermissionSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSetsProvisionedToAccount 列出分配给指定帐户的权限集列表
//
// 列出分配给指定帐户的权限集列表
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

// ListPermissionSetsProvisionedToAccountInvoker 列出分配给指定帐户的权限集列表
func (c *IdentityCenterClient) ListPermissionSetsProvisionedToAccountInvoker(request *model.ListPermissionSetsProvisionedToAccountRequest) *ListPermissionSetsProvisionedToAccountInvoker {
	requestDef := GenReqDefForListPermissionSetsProvisionedToAccount()
	return &ListPermissionSetsProvisionedToAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePermissionSet 更新权限集
//
// 根据权限集ID，删除指定权限集的属性
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
