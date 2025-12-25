package v5

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v5/model"
)

type IamClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIamClient(hcClient *httpclient.HcHttpClient) *IamClient {
	return &IamClient{HcClient: hcClient}
}

func IamClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials,basic.Credentials,v5.IamCredentials")
	return builder
}

// GetAccountSummaryV5 获取此账号中IAM实体使用情况和IAM配额的摘要信息
//
// 该接口可以用于获取此账号中IAM实体使用情况和IAM配额的摘要信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetAccountSummaryV5(request *model.GetAccountSummaryV5Request) (*model.GetAccountSummaryV5Response, error) {
	requestDef := GenReqDefForGetAccountSummaryV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAccountSummaryV5Response), nil
	}
}

// GetAccountSummaryV5Invoker 获取此账号中IAM实体使用情况和IAM配额的摘要信息
func (c *IamClient) GetAccountSummaryV5Invoker(request *model.GetAccountSummaryV5Request) *GetAccountSummaryV5Invoker {
	requestDef := GenReqDefForGetAccountSummaryV5()
	return &GetAccountSummaryV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetAsymmetricSignatureSwitchV5 获取账号非对称签名开关状态
//
// 该接口用于获取账号非对称签名开关的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetAsymmetricSignatureSwitchV5(request *model.GetAsymmetricSignatureSwitchV5Request) (*model.GetAsymmetricSignatureSwitchV5Response, error) {
	requestDef := GenReqDefForGetAsymmetricSignatureSwitchV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAsymmetricSignatureSwitchV5Response), nil
	}
}

// GetAsymmetricSignatureSwitchV5Invoker 获取账号非对称签名开关状态
func (c *IamClient) GetAsymmetricSignatureSwitchV5Invoker(request *model.GetAsymmetricSignatureSwitchV5Request) *GetAsymmetricSignatureSwitchV5Invoker {
	requestDef := GenReqDefForGetAsymmetricSignatureSwitchV5()
	return &GetAsymmetricSignatureSwitchV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetFeatureStatusV5 获取此账号的功能状态
//
// 该接口可以用于获取此账号的功能状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetFeatureStatusV5(request *model.GetFeatureStatusV5Request) (*model.GetFeatureStatusV5Response, error) {
	requestDef := GenReqDefForGetFeatureStatusV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetFeatureStatusV5Response), nil
	}
}

// GetFeatureStatusV5Invoker 获取此账号的功能状态
func (c *IamClient) GetFeatureStatusV5Invoker(request *model.GetFeatureStatusV5Request) *GetFeatureStatusV5Invoker {
	requestDef := GenReqDefForGetFeatureStatusV5()
	return &GetFeatureStatusV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAsymmetricSignatureSwitchV5 设置账号开启或关闭非对称签名
//
// 该接口用于设置账号开启或关闭非对称签名功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) SetAsymmetricSignatureSwitchV5(request *model.SetAsymmetricSignatureSwitchV5Request) (*model.SetAsymmetricSignatureSwitchV5Response, error) {
	requestDef := GenReqDefForSetAsymmetricSignatureSwitchV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAsymmetricSignatureSwitchV5Response), nil
	}
}

// SetAsymmetricSignatureSwitchV5Invoker 设置账号开启或关闭非对称签名
func (c *IamClient) SetAsymmetricSignatureSwitchV5Invoker(request *model.SetAsymmetricSignatureSwitchV5Request) *SetAsymmetricSignatureSwitchV5Invoker {
	requestDef := GenReqDefForSetAsymmetricSignatureSwitchV5()
	return &SetAsymmetricSignatureSwitchV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgencyV5 创建信任委托
//
// 该接口可以用于创建信任委托。
//
// &gt; 信任委托只能授予身份策略且仅兼容支持身份策略的云服务，详情见[[支持身份策略与信任委托的云服务列表](https://support.huaweicloud.com/productdesc-iam/iam_01_1224.html)](tag:hws)[[支持身份策略与信任委托的云服务列表](https://support.huaweicloud.com/intl/zh-cn/productdesc-iam/iam_01_1224.html)](tag:hws_hk)[[支持身份策略与信任委托的云服务列表](https://support.huaweicloud.com/eu/productdesc-iam/iam_01_1224.html)](tag:hws_eu)[《统一身份认证用户指南》的“支持身份策略与信任委托的云服务列表”章节](tag:fcs,fcs_vm,ctc,hws_test,g42,tm)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateAgencyV5(request *model.CreateAgencyV5Request) (*model.CreateAgencyV5Response, error) {
	requestDef := GenReqDefForCreateAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyV5Response), nil
	}
}

// CreateAgencyV5Invoker 创建信任委托
func (c *IamClient) CreateAgencyV5Invoker(request *model.CreateAgencyV5Request) *CreateAgencyV5Invoker {
	requestDef := GenReqDefForCreateAgencyV5()
	return &CreateAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServiceLinkedAgencyV5 创建服务关联委托
//
// 该接口可以用于创建服务关联委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateServiceLinkedAgencyV5(request *model.CreateServiceLinkedAgencyV5Request) (*model.CreateServiceLinkedAgencyV5Response, error) {
	requestDef := GenReqDefForCreateServiceLinkedAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceLinkedAgencyV5Response), nil
	}
}

// CreateServiceLinkedAgencyV5Invoker 创建服务关联委托
func (c *IamClient) CreateServiceLinkedAgencyV5Invoker(request *model.CreateServiceLinkedAgencyV5Request) *CreateServiceLinkedAgencyV5Invoker {
	requestDef := GenReqDefForCreateServiceLinkedAgencyV5()
	return &CreateServiceLinkedAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgencyV5 删除信任委托
//
// 该接口可以用于删除信任委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteAgencyV5(request *model.DeleteAgencyV5Request) (*model.DeleteAgencyV5Response, error) {
	requestDef := GenReqDefForDeleteAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgencyV5Response), nil
	}
}

// DeleteAgencyV5Invoker 删除信任委托
func (c *IamClient) DeleteAgencyV5Invoker(request *model.DeleteAgencyV5Request) *DeleteAgencyV5Invoker {
	requestDef := GenReqDefForDeleteAgencyV5()
	return &DeleteAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceLinkedAgencyV5 删除服务关联委托
//
// 该接口可以用于服务关联委托删除自己。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteServiceLinkedAgencyV5(request *model.DeleteServiceLinkedAgencyV5Request) (*model.DeleteServiceLinkedAgencyV5Response, error) {
	requestDef := GenReqDefForDeleteServiceLinkedAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceLinkedAgencyV5Response), nil
	}
}

// DeleteServiceLinkedAgencyV5Invoker 删除服务关联委托
func (c *IamClient) DeleteServiceLinkedAgencyV5Invoker(request *model.DeleteServiceLinkedAgencyV5Request) *DeleteServiceLinkedAgencyV5Invoker {
	requestDef := GenReqDefForDeleteServiceLinkedAgencyV5()
	return &DeleteServiceLinkedAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetAgencyV5 查询委托或信任委托详情
//
// 该接口可以用于查询委托或信任委托详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetAgencyV5(request *model.GetAgencyV5Request) (*model.GetAgencyV5Response, error) {
	requestDef := GenReqDefForGetAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAgencyV5Response), nil
	}
}

// GetAgencyV5Invoker 查询委托或信任委托详情
func (c *IamClient) GetAgencyV5Invoker(request *model.GetAgencyV5Request) *GetAgencyV5Invoker {
	requestDef := GenReqDefForGetAgencyV5()
	return &GetAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetServiceLinkedAgencyDeletionStatusV5 获取服务关联委托删除状态
//
// 该接口可以用于获取服务关联委托删除状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetServiceLinkedAgencyDeletionStatusV5(request *model.GetServiceLinkedAgencyDeletionStatusV5Request) (*model.GetServiceLinkedAgencyDeletionStatusV5Response, error) {
	requestDef := GenReqDefForGetServiceLinkedAgencyDeletionStatusV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetServiceLinkedAgencyDeletionStatusV5Response), nil
	}
}

// GetServiceLinkedAgencyDeletionStatusV5Invoker 获取服务关联委托删除状态
func (c *IamClient) GetServiceLinkedAgencyDeletionStatusV5Invoker(request *model.GetServiceLinkedAgencyDeletionStatusV5Request) *GetServiceLinkedAgencyDeletionStatusV5Invoker {
	requestDef := GenReqDefForGetServiceLinkedAgencyDeletionStatusV5()
	return &GetServiceLinkedAgencyDeletionStatusV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgenciesV5 查询指定条件下的委托及信任委托列表
//
// 该接口可以用于查询指定条件下的委托及信任委托列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListAgenciesV5(request *model.ListAgenciesV5Request) (*model.ListAgenciesV5Response, error) {
	requestDef := GenReqDefForListAgenciesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgenciesV5Response), nil
	}
}

// ListAgenciesV5Invoker 查询指定条件下的委托及信任委托列表
func (c *IamClient) ListAgenciesV5Invoker(request *model.ListAgenciesV5Request) *ListAgenciesV5Invoker {
	requestDef := GenReqDefForListAgenciesV5()
	return &ListAgenciesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAgencyV5 修改信任委托
//
// 该接口可以用于修改信任委托。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateAgencyV5(request *model.UpdateAgencyV5Request) (*model.UpdateAgencyV5Response, error) {
	requestDef := GenReqDefForUpdateAgencyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAgencyV5Response), nil
	}
}

// UpdateAgencyV5Invoker 修改信任委托
func (c *IamClient) UpdateAgencyV5Invoker(request *model.UpdateAgencyV5Request) *UpdateAgencyV5Invoker {
	requestDef := GenReqDefForUpdateAgencyV5()
	return &UpdateAgencyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTrustPolicyV5 修改信任委托信任策略
//
// 该接口可以用于修改信任委托信任策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateTrustPolicyV5(request *model.UpdateTrustPolicyV5Request) (*model.UpdateTrustPolicyV5Response, error) {
	requestDef := GenReqDefForUpdateTrustPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTrustPolicyV5Response), nil
	}
}

// UpdateTrustPolicyV5Invoker 修改信任委托信任策略
func (c *IamClient) UpdateTrustPolicyV5Invoker(request *model.UpdateTrustPolicyV5Request) *UpdateTrustPolicyV5Invoker {
	requestDef := GenReqDefForUpdateTrustPolicyV5()
	return &UpdateTrustPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetAuthorizationSchemaV5 查询指定服务授权概要
//
// 该接口可以用于查询指定云服务的授权概要。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetAuthorizationSchemaV5(request *model.GetAuthorizationSchemaV5Request) (*model.GetAuthorizationSchemaV5Response, error) {
	requestDef := GenReqDefForGetAuthorizationSchemaV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAuthorizationSchemaV5Response), nil
	}
}

// GetAuthorizationSchemaV5Invoker 查询指定服务授权概要
func (c *IamClient) GetAuthorizationSchemaV5Invoker(request *model.GetAuthorizationSchemaV5Request) *GetAuthorizationSchemaV5Invoker {
	requestDef := GenReqDefForGetAuthorizationSchemaV5()
	return &GetAuthorizationSchemaV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRegisteredServicesForAuthSchemaV5 查询已注册云服务列表
//
// 该接口可以用于查询已注册云服务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListRegisteredServicesForAuthSchemaV5(request *model.ListRegisteredServicesForAuthSchemaV5Request) (*model.ListRegisteredServicesForAuthSchemaV5Response, error) {
	requestDef := GenReqDefForListRegisteredServicesForAuthSchemaV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegisteredServicesForAuthSchemaV5Response), nil
	}
}

// ListRegisteredServicesForAuthSchemaV5Invoker 查询已注册云服务列表
func (c *IamClient) ListRegisteredServicesForAuthSchemaV5Invoker(request *model.ListRegisteredServicesForAuthSchemaV5Request) *ListRegisteredServicesForAuthSchemaV5Invoker {
	requestDef := GenReqDefForListRegisteredServicesForAuthSchemaV5()
	return &ListRegisteredServicesForAuthSchemaV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServicePrincipalsV5 获取全部服务主体
//
// 该接口可以用于获取全部服务主体。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListServicePrincipalsV5(request *model.ListServicePrincipalsV5Request) (*model.ListServicePrincipalsV5Response, error) {
	requestDef := GenReqDefForListServicePrincipalsV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServicePrincipalsV5Response), nil
	}
}

// ListServicePrincipalsV5Invoker 获取全部服务主体
func (c *IamClient) ListServicePrincipalsV5Invoker(request *model.ListServicePrincipalsV5Request) *ListServicePrincipalsV5Invoker {
	requestDef := GenReqDefForListServicePrincipalsV5()
	return &ListServicePrincipalsV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddUserToGroupV5 添加IAM用户到用户组
//
// 该接口可以用于添加IAM用户到用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) AddUserToGroupV5(request *model.AddUserToGroupV5Request) (*model.AddUserToGroupV5Response, error) {
	requestDef := GenReqDefForAddUserToGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddUserToGroupV5Response), nil
	}
}

// AddUserToGroupV5Invoker 添加IAM用户到用户组
func (c *IamClient) AddUserToGroupV5Invoker(request *model.AddUserToGroupV5Request) *AddUserToGroupV5Invoker {
	requestDef := GenReqDefForAddUserToGroupV5()
	return &AddUserToGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupV5 创建用户组
//
// 该接口可以用于创建用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateGroupV5(request *model.CreateGroupV5Request) (*model.CreateGroupV5Response, error) {
	requestDef := GenReqDefForCreateGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupV5Response), nil
	}
}

// CreateGroupV5Invoker 创建用户组
func (c *IamClient) CreateGroupV5Invoker(request *model.CreateGroupV5Request) *CreateGroupV5Invoker {
	requestDef := GenReqDefForCreateGroupV5()
	return &CreateGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupV5 删除用户组
//
// 该接口可以用于删除用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteGroupV5(request *model.DeleteGroupV5Request) (*model.DeleteGroupV5Response, error) {
	requestDef := GenReqDefForDeleteGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupV5Response), nil
	}
}

// DeleteGroupV5Invoker 删除用户组
func (c *IamClient) DeleteGroupV5Invoker(request *model.DeleteGroupV5Request) *DeleteGroupV5Invoker {
	requestDef := GenReqDefForDeleteGroupV5()
	return &DeleteGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupsV5 查询用户组列表
//
// 该接口可以用于查询用户组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListGroupsV5(request *model.ListGroupsV5Request) (*model.ListGroupsV5Response, error) {
	requestDef := GenReqDefForListGroupsV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsV5Response), nil
	}
}

// ListGroupsV5Invoker 查询用户组列表
func (c *IamClient) ListGroupsV5Invoker(request *model.ListGroupsV5Request) *ListGroupsV5Invoker {
	requestDef := GenReqDefForListGroupsV5()
	return &ListGroupsV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveUserFromGroupV5 移除用户组中的IAM用户
//
// 该接口可以用于移除用户组中的IAM用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) RemoveUserFromGroupV5(request *model.RemoveUserFromGroupV5Request) (*model.RemoveUserFromGroupV5Response, error) {
	requestDef := GenReqDefForRemoveUserFromGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveUserFromGroupV5Response), nil
	}
}

// RemoveUserFromGroupV5Invoker 移除用户组中的IAM用户
func (c *IamClient) RemoveUserFromGroupV5Invoker(request *model.RemoveUserFromGroupV5Request) *RemoveUserFromGroupV5Invoker {
	requestDef := GenReqDefForRemoveUserFromGroupV5()
	return &RemoveUserFromGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupSummary 查询用户组相关属性
//
// 该接口可以用于查询用户组相关属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowGroupSummary(request *model.ShowGroupSummaryRequest) (*model.ShowGroupSummaryResponse, error) {
	requestDef := GenReqDefForShowGroupSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupSummaryResponse), nil
	}
}

// ShowGroupSummaryInvoker 查询用户组相关属性
func (c *IamClient) ShowGroupSummaryInvoker(request *model.ShowGroupSummaryRequest) *ShowGroupSummaryInvoker {
	requestDef := GenReqDefForShowGroupSummary()
	return &ShowGroupSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGroupV5 查询用户组详情
//
// 该接口可以用于查询用户组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowGroupV5(request *model.ShowGroupV5Request) (*model.ShowGroupV5Response, error) {
	requestDef := GenReqDefForShowGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGroupV5Response), nil
	}
}

// ShowGroupV5Invoker 查询用户组详情
func (c *IamClient) ShowGroupV5Invoker(request *model.ShowGroupV5Request) *ShowGroupV5Invoker {
	requestDef := GenReqDefForShowGroupV5()
	return &ShowGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupV5 修改用户组
//
// 该接口可以用于修改用户组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateGroupV5(request *model.UpdateGroupV5Request) (*model.UpdateGroupV5Response, error) {
	requestDef := GenReqDefForUpdateGroupV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupV5Response), nil
	}
}

// UpdateGroupV5Invoker 修改用户组
func (c *IamClient) UpdateGroupV5Invoker(request *model.UpdateGroupV5Request) *UpdateGroupV5Invoker {
	requestDef := GenReqDefForUpdateGroupV5()
	return &UpdateGroupV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateVirtualMfaDeviceV5 创建MFA设备
//
// 该接口可以用于创建MFA设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateVirtualMfaDeviceV5(request *model.CreateVirtualMfaDeviceV5Request) (*model.CreateVirtualMfaDeviceV5Response, error) {
	requestDef := GenReqDefForCreateVirtualMfaDeviceV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVirtualMfaDeviceV5Response), nil
	}
}

// CreateVirtualMfaDeviceV5Invoker 创建MFA设备
func (c *IamClient) CreateVirtualMfaDeviceV5Invoker(request *model.CreateVirtualMfaDeviceV5Request) *CreateVirtualMfaDeviceV5Invoker {
	requestDef := GenReqDefForCreateVirtualMfaDeviceV5()
	return &CreateVirtualMfaDeviceV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVirtualMfaDeviceV5 删除MFA设备
//
// 该接口可以用于删除MFA设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteVirtualMfaDeviceV5(request *model.DeleteVirtualMfaDeviceV5Request) (*model.DeleteVirtualMfaDeviceV5Response, error) {
	requestDef := GenReqDefForDeleteVirtualMfaDeviceV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVirtualMfaDeviceV5Response), nil
	}
}

// DeleteVirtualMfaDeviceV5Invoker 删除MFA设备
func (c *IamClient) DeleteVirtualMfaDeviceV5Invoker(request *model.DeleteVirtualMfaDeviceV5Request) *DeleteVirtualMfaDeviceV5Invoker {
	requestDef := GenReqDefForDeleteVirtualMfaDeviceV5()
	return &DeleteVirtualMfaDeviceV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableMfaDeviceV5 禁用MFA设备
//
// 该接口可以用于禁用指定的MFA设备并删除其与对应IAM用户的关联。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DisableMfaDeviceV5(request *model.DisableMfaDeviceV5Request) (*model.DisableMfaDeviceV5Response, error) {
	requestDef := GenReqDefForDisableMfaDeviceV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableMfaDeviceV5Response), nil
	}
}

// DisableMfaDeviceV5Invoker 禁用MFA设备
func (c *IamClient) DisableMfaDeviceV5Invoker(request *model.DisableMfaDeviceV5Request) *DisableMfaDeviceV5Invoker {
	requestDef := GenReqDefForDisableMfaDeviceV5()
	return &DisableMfaDeviceV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableMfaDeviceV5 启用MFA设备
//
// 该接口可以用于启用指定的MFA设备并将其与指定的IAM用户关联。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) EnableMfaDeviceV5(request *model.EnableMfaDeviceV5Request) (*model.EnableMfaDeviceV5Response, error) {
	requestDef := GenReqDefForEnableMfaDeviceV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableMfaDeviceV5Response), nil
	}
}

// EnableMfaDeviceV5Invoker 启用MFA设备
func (c *IamClient) EnableMfaDeviceV5Invoker(request *model.EnableMfaDeviceV5Request) *EnableMfaDeviceV5Invoker {
	requestDef := GenReqDefForEnableMfaDeviceV5()
	return &EnableMfaDeviceV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMfaDevicesV5 列举全部MFA设备
//
// 该接口可以用于列举全部MFA设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListMfaDevicesV5(request *model.ListMfaDevicesV5Request) (*model.ListMfaDevicesV5Response, error) {
	requestDef := GenReqDefForListMfaDevicesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMfaDevicesV5Response), nil
	}
}

// ListMfaDevicesV5Invoker 列举全部MFA设备
func (c *IamClient) ListMfaDevicesV5Invoker(request *model.ListMfaDevicesV5Request) *ListMfaDevicesV5Invoker {
	requestDef := GenReqDefForListMfaDevicesV5()
	return &ListMfaDevicesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyV5 创建自定义身份策略
//
// 该接口可以用于创建一个默认版本为v1的新自定义身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreatePolicyV5(request *model.CreatePolicyV5Request) (*model.CreatePolicyV5Response, error) {
	requestDef := GenReqDefForCreatePolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyV5Response), nil
	}
}

// CreatePolicyV5Invoker 创建自定义身份策略
func (c *IamClient) CreatePolicyV5Invoker(request *model.CreatePolicyV5Request) *CreatePolicyV5Invoker {
	requestDef := GenReqDefForCreatePolicyV5()
	return &CreatePolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyVersionV5 为指定身份策略创建一个新版本
//
// 该接口可以用于为指定身份策略创建一个新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreatePolicyVersionV5(request *model.CreatePolicyVersionV5Request) (*model.CreatePolicyVersionV5Response, error) {
	requestDef := GenReqDefForCreatePolicyVersionV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyVersionV5Response), nil
	}
}

// CreatePolicyVersionV5Invoker 为指定身份策略创建一个新版本
func (c *IamClient) CreatePolicyVersionV5Invoker(request *model.CreatePolicyVersionV5Request) *CreatePolicyVersionV5Invoker {
	requestDef := GenReqDefForCreatePolicyVersionV5()
	return &CreatePolicyVersionV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyV5 删除自定义身份策略
//
// 该接口可以用于删除一个存在的自定义身份策略，必须确保该自定义身份策略没有附加在任何IAM用户、用户组、委托或信任委托上。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeletePolicyV5(request *model.DeletePolicyV5Request) (*model.DeletePolicyV5Response, error) {
	requestDef := GenReqDefForDeletePolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyV5Response), nil
	}
}

// DeletePolicyV5Invoker 删除自定义身份策略
func (c *IamClient) DeletePolicyV5Invoker(request *model.DeletePolicyV5Request) *DeletePolicyV5Invoker {
	requestDef := GenReqDefForDeletePolicyV5()
	return &DeletePolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyVersionV5 删除指定身份策略版本
//
// 该接口可以用于删除指定身份策略的指定版本。默认身份策略版本不能被删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeletePolicyVersionV5(request *model.DeletePolicyVersionV5Request) (*model.DeletePolicyVersionV5Response, error) {
	requestDef := GenReqDefForDeletePolicyVersionV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyVersionV5Response), nil
	}
}

// DeletePolicyVersionV5Invoker 删除指定身份策略版本
func (c *IamClient) DeletePolicyVersionV5Invoker(request *model.DeletePolicyVersionV5Request) *DeletePolicyVersionV5Invoker {
	requestDef := GenReqDefForDeletePolicyVersionV5()
	return &DeletePolicyVersionV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetPolicyV5 通过身份策略ID获取身份策略
//
// 该接口可以用于通过身份策略ID获取身份策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetPolicyV5(request *model.GetPolicyV5Request) (*model.GetPolicyV5Response, error) {
	requestDef := GenReqDefForGetPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetPolicyV5Response), nil
	}
}

// GetPolicyV5Invoker 通过身份策略ID获取身份策略
func (c *IamClient) GetPolicyV5Invoker(request *model.GetPolicyV5Request) *GetPolicyV5Invoker {
	requestDef := GenReqDefForGetPolicyV5()
	return &GetPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetPolicyVersionV5 查询指定身份策略版本
//
// 该接口可以用于查询指定身份策略的指定版本的相关信息，包括身份策略文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) GetPolicyVersionV5(request *model.GetPolicyVersionV5Request) (*model.GetPolicyVersionV5Response, error) {
	requestDef := GenReqDefForGetPolicyVersionV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetPolicyVersionV5Response), nil
	}
}

// GetPolicyVersionV5Invoker 查询指定身份策略版本
func (c *IamClient) GetPolicyVersionV5Invoker(request *model.GetPolicyVersionV5Request) *GetPolicyVersionV5Invoker {
	requestDef := GenReqDefForGetPolicyVersionV5()
	return &GetPolicyVersionV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoliciesV5 查询所有身份策略
//
// 该接口可以用于查询所有身份策略，包含系统预置身份策略和自定义身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListPoliciesV5(request *model.ListPoliciesV5Request) (*model.ListPoliciesV5Response, error) {
	requestDef := GenReqDefForListPoliciesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoliciesV5Response), nil
	}
}

// ListPoliciesV5Invoker 查询所有身份策略
func (c *IamClient) ListPoliciesV5Invoker(request *model.ListPoliciesV5Request) *ListPoliciesV5Invoker {
	requestDef := GenReqDefForListPoliciesV5()
	return &ListPoliciesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyVersionsV5 查询指定身份策略的所有版本
//
// 该接口可以用于查询指定身份策略的所有版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListPolicyVersionsV5(request *model.ListPolicyVersionsV5Request) (*model.ListPolicyVersionsV5Response, error) {
	requestDef := GenReqDefForListPolicyVersionsV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyVersionsV5Response), nil
	}
}

// ListPolicyVersionsV5Invoker 查询指定身份策略的所有版本
func (c *IamClient) ListPolicyVersionsV5Invoker(request *model.ListPolicyVersionsV5Request) *ListPolicyVersionsV5Invoker {
	requestDef := GenReqDefForListPolicyVersionsV5()
	return &ListPolicyVersionsV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetDefaultPolicyVersionV5 将指定身份策略版本设置为默认版本
//
// 该接口可以用于将指定身份策略的指定版本设置为默认版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) SetDefaultPolicyVersionV5(request *model.SetDefaultPolicyVersionV5Request) (*model.SetDefaultPolicyVersionV5Response, error) {
	requestDef := GenReqDefForSetDefaultPolicyVersionV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDefaultPolicyVersionV5Response), nil
	}
}

// SetDefaultPolicyVersionV5Invoker 将指定身份策略版本设置为默认版本
func (c *IamClient) SetDefaultPolicyVersionV5Invoker(request *model.SetDefaultPolicyVersionV5Request) *SetDefaultPolicyVersionV5Invoker {
	requestDef := GenReqDefForSetDefaultPolicyVersionV5()
	return &SetDefaultPolicyVersionV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachAgencyPolicyV5 为委托或信任委托附加身份策略
//
// 该接口可以用于为指定委托或信任委托附加指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) AttachAgencyPolicyV5(request *model.AttachAgencyPolicyV5Request) (*model.AttachAgencyPolicyV5Response, error) {
	requestDef := GenReqDefForAttachAgencyPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachAgencyPolicyV5Response), nil
	}
}

// AttachAgencyPolicyV5Invoker 为委托或信任委托附加身份策略
func (c *IamClient) AttachAgencyPolicyV5Invoker(request *model.AttachAgencyPolicyV5Request) *AttachAgencyPolicyV5Invoker {
	requestDef := GenReqDefForAttachAgencyPolicyV5()
	return &AttachAgencyPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachGroupPolicyV5 为用户组附加身份策略
//
// 该接口可以用于为指定用户组附加指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) AttachGroupPolicyV5(request *model.AttachGroupPolicyV5Request) (*model.AttachGroupPolicyV5Response, error) {
	requestDef := GenReqDefForAttachGroupPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachGroupPolicyV5Response), nil
	}
}

// AttachGroupPolicyV5Invoker 为用户组附加身份策略
func (c *IamClient) AttachGroupPolicyV5Invoker(request *model.AttachGroupPolicyV5Request) *AttachGroupPolicyV5Invoker {
	requestDef := GenReqDefForAttachGroupPolicyV5()
	return &AttachGroupPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachUserPolicyV5 为IAM用户附加身份策略
//
// 该接口可以用于为指定IAM用户附加指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) AttachUserPolicyV5(request *model.AttachUserPolicyV5Request) (*model.AttachUserPolicyV5Response, error) {
	requestDef := GenReqDefForAttachUserPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachUserPolicyV5Response), nil
	}
}

// AttachUserPolicyV5Invoker 为IAM用户附加身份策略
func (c *IamClient) AttachUserPolicyV5Invoker(request *model.AttachUserPolicyV5Request) *AttachUserPolicyV5Invoker {
	requestDef := GenReqDefForAttachUserPolicyV5()
	return &AttachUserPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachAgencyPolicyV5 从委托或信任委托分离身份策略
//
// 该接口可以用于从指定委托或信任委托中分离指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DetachAgencyPolicyV5(request *model.DetachAgencyPolicyV5Request) (*model.DetachAgencyPolicyV5Response, error) {
	requestDef := GenReqDefForDetachAgencyPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachAgencyPolicyV5Response), nil
	}
}

// DetachAgencyPolicyV5Invoker 从委托或信任委托分离身份策略
func (c *IamClient) DetachAgencyPolicyV5Invoker(request *model.DetachAgencyPolicyV5Request) *DetachAgencyPolicyV5Invoker {
	requestDef := GenReqDefForDetachAgencyPolicyV5()
	return &DetachAgencyPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachGroupPolicyV5 从用户组分离身份策略
//
// 该接口可以用于从指定用户组分离指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DetachGroupPolicyV5(request *model.DetachGroupPolicyV5Request) (*model.DetachGroupPolicyV5Response, error) {
	requestDef := GenReqDefForDetachGroupPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachGroupPolicyV5Response), nil
	}
}

// DetachGroupPolicyV5Invoker 从用户组分离身份策略
func (c *IamClient) DetachGroupPolicyV5Invoker(request *model.DetachGroupPolicyV5Request) *DetachGroupPolicyV5Invoker {
	requestDef := GenReqDefForDetachGroupPolicyV5()
	return &DetachGroupPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachUserPolicyV5 从IAM用户分离身份策略
//
// 该接口可以用于从指定的IAM用户分离指定身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DetachUserPolicyV5(request *model.DetachUserPolicyV5Request) (*model.DetachUserPolicyV5Response, error) {
	requestDef := GenReqDefForDetachUserPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachUserPolicyV5Response), nil
	}
}

// DetachUserPolicyV5Invoker 从IAM用户分离身份策略
func (c *IamClient) DetachUserPolicyV5Invoker(request *model.DetachUserPolicyV5Request) *DetachUserPolicyV5Invoker {
	requestDef := GenReqDefForDetachUserPolicyV5()
	return &DetachUserPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachedAgencyPoliciesV5 查询指定委托或信任委托附加的所有身份策略
//
// 该接口可用于查询指定委托或信任委托附加的所有身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListAttachedAgencyPoliciesV5(request *model.ListAttachedAgencyPoliciesV5Request) (*model.ListAttachedAgencyPoliciesV5Response, error) {
	requestDef := GenReqDefForListAttachedAgencyPoliciesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachedAgencyPoliciesV5Response), nil
	}
}

// ListAttachedAgencyPoliciesV5Invoker 查询指定委托或信任委托附加的所有身份策略
func (c *IamClient) ListAttachedAgencyPoliciesV5Invoker(request *model.ListAttachedAgencyPoliciesV5Request) *ListAttachedAgencyPoliciesV5Invoker {
	requestDef := GenReqDefForListAttachedAgencyPoliciesV5()
	return &ListAttachedAgencyPoliciesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachedGroupPoliciesV5 查询指定用户组附加的所有身份策略
//
// 该接口可用于查询指定用户组附加的所有身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListAttachedGroupPoliciesV5(request *model.ListAttachedGroupPoliciesV5Request) (*model.ListAttachedGroupPoliciesV5Response, error) {
	requestDef := GenReqDefForListAttachedGroupPoliciesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachedGroupPoliciesV5Response), nil
	}
}

// ListAttachedGroupPoliciesV5Invoker 查询指定用户组附加的所有身份策略
func (c *IamClient) ListAttachedGroupPoliciesV5Invoker(request *model.ListAttachedGroupPoliciesV5Request) *ListAttachedGroupPoliciesV5Invoker {
	requestDef := GenReqDefForListAttachedGroupPoliciesV5()
	return &ListAttachedGroupPoliciesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttachedUserPoliciesV5 查询指定IAM用户附加的所有身份策略
//
// 该接口可用于查询指定IAM用户附加的所有身份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListAttachedUserPoliciesV5(request *model.ListAttachedUserPoliciesV5Request) (*model.ListAttachedUserPoliciesV5Response, error) {
	requestDef := GenReqDefForListAttachedUserPoliciesV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttachedUserPoliciesV5Response), nil
	}
}

// ListAttachedUserPoliciesV5Invoker 查询指定IAM用户附加的所有身份策略
func (c *IamClient) ListAttachedUserPoliciesV5Invoker(request *model.ListAttachedUserPoliciesV5Request) *ListAttachedUserPoliciesV5Invoker {
	requestDef := GenReqDefForListAttachedUserPoliciesV5()
	return &ListAttachedUserPoliciesV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEntitiesForPolicyV5 查询指定身份策略附加的所有实体
//
// 该接口可用于查询指定身份策略附加的所有实体。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListEntitiesForPolicyV5(request *model.ListEntitiesForPolicyV5Request) (*model.ListEntitiesForPolicyV5Response, error) {
	requestDef := GenReqDefForListEntitiesForPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEntitiesForPolicyV5Response), nil
	}
}

// ListEntitiesForPolicyV5Invoker 查询指定身份策略附加的所有实体
func (c *IamClient) ListEntitiesForPolicyV5Invoker(request *model.ListEntitiesForPolicyV5Request) *ListEntitiesForPolicyV5Invoker {
	requestDef := GenReqDefForListEntitiesForPolicyV5()
	return &ListEntitiesForPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTagsV5 删除指定资源的部分标签
//
// 该接口可以用于删除指定资源的部分标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteResourceTagsV5(request *model.DeleteResourceTagsV5Request) (*model.DeleteResourceTagsV5Response, error) {
	requestDef := GenReqDefForDeleteResourceTagsV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagsV5Response), nil
	}
}

// DeleteResourceTagsV5Invoker 删除指定资源的部分标签
func (c *IamClient) DeleteResourceTagsV5Invoker(request *model.DeleteResourceTagsV5Request) *DeleteResourceTagsV5Invoker {
	requestDef := GenReqDefForDeleteResourceTagsV5()
	return &DeleteResourceTagsV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTagsV5 获取指定资源的所有标签
//
// 该接口可以用于获取指定资源的所有标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListResourceTagsV5(request *model.ListResourceTagsV5Request) (*model.ListResourceTagsV5Response, error) {
	requestDef := GenReqDefForListResourceTagsV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsV5Response), nil
	}
}

// ListResourceTagsV5Invoker 获取指定资源的所有标签
func (c *IamClient) ListResourceTagsV5Invoker(request *model.ListResourceTagsV5Request) *ListResourceTagsV5Invoker {
	requestDef := GenReqDefForListResourceTagsV5()
	return &ListResourceTagsV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagResourceV5 为IAM资源打上标签
//
// 该接口可以用于为IAM资源打上标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) TagResourceV5(request *model.TagResourceV5Request) (*model.TagResourceV5Response, error) {
	requestDef := GenReqDefForTagResourceV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagResourceV5Response), nil
	}
}

// TagResourceV5Invoker 为IAM资源打上标签
func (c *IamClient) TagResourceV5Invoker(request *model.TagResourceV5Request) *TagResourceV5Invoker {
	requestDef := GenReqDefForTagResourceV5()
	return &TagResourceV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoginPolicyV5 查询账号登录策略
//
// 该接口可以用于查询账号登录策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowLoginPolicyV5(request *model.ShowLoginPolicyV5Request) (*model.ShowLoginPolicyV5Response, error) {
	requestDef := GenReqDefForShowLoginPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoginPolicyV5Response), nil
	}
}

// ShowLoginPolicyV5Invoker 查询账号登录策略
func (c *IamClient) ShowLoginPolicyV5Invoker(request *model.ShowLoginPolicyV5Request) *ShowLoginPolicyV5Invoker {
	requestDef := GenReqDefForShowLoginPolicyV5()
	return &ShowLoginPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPasswordPolicyV5 查询账号密码策略
//
// 该接口可以用于查询账号密码策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowPasswordPolicyV5(request *model.ShowPasswordPolicyV5Request) (*model.ShowPasswordPolicyV5Response, error) {
	requestDef := GenReqDefForShowPasswordPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPasswordPolicyV5Response), nil
	}
}

// ShowPasswordPolicyV5Invoker 查询账号密码策略
func (c *IamClient) ShowPasswordPolicyV5Invoker(request *model.ShowPasswordPolicyV5Request) *ShowPasswordPolicyV5Invoker {
	requestDef := GenReqDefForShowPasswordPolicyV5()
	return &ShowPasswordPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLoginPolicyV5 修改账号登录策略
//
// 该接口可以用于修改账号登录策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateLoginPolicyV5(request *model.UpdateLoginPolicyV5Request) (*model.UpdateLoginPolicyV5Response, error) {
	requestDef := GenReqDefForUpdateLoginPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoginPolicyV5Response), nil
	}
}

// UpdateLoginPolicyV5Invoker 修改账号登录策略
func (c *IamClient) UpdateLoginPolicyV5Invoker(request *model.UpdateLoginPolicyV5Request) *UpdateLoginPolicyV5Invoker {
	requestDef := GenReqDefForUpdateLoginPolicyV5()
	return &UpdateLoginPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePasswordPolicyV5 修改账号密码策略
//
// 该接口可以用于修改账号密码策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdatePasswordPolicyV5(request *model.UpdatePasswordPolicyV5Request) (*model.UpdatePasswordPolicyV5Response, error) {
	requestDef := GenReqDefForUpdatePasswordPolicyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordPolicyV5Response), nil
	}
}

// UpdatePasswordPolicyV5Invoker 修改账号密码策略
func (c *IamClient) UpdatePasswordPolicyV5Invoker(request *model.UpdatePasswordPolicyV5Request) *UpdatePasswordPolicyV5Invoker {
	requestDef := GenReqDefForUpdatePasswordPolicyV5()
	return &UpdatePasswordPolicyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserV5 创建IAM用户
//
// 该接口可以用于创建IAM用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateUserV5(request *model.CreateUserV5Request) (*model.CreateUserV5Response, error) {
	requestDef := GenReqDefForCreateUserV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserV5Response), nil
	}
}

// CreateUserV5Invoker 创建IAM用户
func (c *IamClient) CreateUserV5Invoker(request *model.CreateUserV5Request) *CreateUserV5Invoker {
	requestDef := GenReqDefForCreateUserV5()
	return &CreateUserV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserV5 删除IAM用户
//
// 该接口可以用于删除指定IAM用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteUserV5(request *model.DeleteUserV5Request) (*model.DeleteUserV5Response, error) {
	requestDef := GenReqDefForDeleteUserV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserV5Response), nil
	}
}

// DeleteUserV5Invoker 删除IAM用户
func (c *IamClient) DeleteUserV5Invoker(request *model.DeleteUserV5Request) *DeleteUserV5Invoker {
	requestDef := GenReqDefForDeleteUserV5()
	return &DeleteUserV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsersV5 查询IAM用户列表
//
// 该接口可以用于查询IAM用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListUsersV5(request *model.ListUsersV5Request) (*model.ListUsersV5Response, error) {
	requestDef := GenReqDefForListUsersV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersV5Response), nil
	}
}

// ListUsersV5Invoker 查询IAM用户列表
func (c *IamClient) ListUsersV5Invoker(request *model.ListUsersV5Request) *ListUsersV5Invoker {
	requestDef := GenReqDefForListUsersV5()
	return &ListUsersV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserLastLoginV5 查询IAM用户最后登录时间
//
// 该接口可以用于查询IAM用户的最后登录时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowUserLastLoginV5(request *model.ShowUserLastLoginV5Request) (*model.ShowUserLastLoginV5Response, error) {
	requestDef := GenReqDefForShowUserLastLoginV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserLastLoginV5Response), nil
	}
}

// ShowUserLastLoginV5Invoker 查询IAM用户最后登录时间
func (c *IamClient) ShowUserLastLoginV5Invoker(request *model.ShowUserLastLoginV5Request) *ShowUserLastLoginV5Invoker {
	requestDef := GenReqDefForShowUserLastLoginV5()
	return &ShowUserLastLoginV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUserV5 查询IAM用户详情
//
// 该接口可以用于查询IAM用户详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowUserV5(request *model.ShowUserV5Request) (*model.ShowUserV5Response, error) {
	requestDef := GenReqDefForShowUserV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUserV5Response), nil
	}
}

// ShowUserV5Invoker 查询IAM用户详情
func (c *IamClient) ShowUserV5Invoker(request *model.ShowUserV5Request) *ShowUserV5Invoker {
	requestDef := GenReqDefForShowUserV5()
	return &ShowUserV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserV5 修改IAM用户信息
//
// 该接口可以用于修改IAM用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateUserV5(request *model.UpdateUserV5Request) (*model.UpdateUserV5Response, error) {
	requestDef := GenReqDefForUpdateUserV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserV5Response), nil
	}
}

// UpdateUserV5Invoker 修改IAM用户信息
func (c *IamClient) UpdateUserV5Invoker(request *model.UpdateUserV5Request) *UpdateUserV5Invoker {
	requestDef := GenReqDefForUpdateUserV5()
	return &UpdateUserV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePasswordV5 修改IAM用户密码
//
// 该接口可以用于IAM用户修改自己的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ChangePasswordV5(request *model.ChangePasswordV5Request) (*model.ChangePasswordV5Response, error) {
	requestDef := GenReqDefForChangePasswordV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePasswordV5Response), nil
	}
}

// ChangePasswordV5Invoker 修改IAM用户密码
func (c *IamClient) ChangePasswordV5Invoker(request *model.ChangePasswordV5Request) *ChangePasswordV5Invoker {
	requestDef := GenReqDefForChangePasswordV5()
	return &ChangePasswordV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessKeyV5 创建永久访问密钥
//
// 该接口可以用于给IAM用户创建永久访问密钥。
//
// 访问密钥（Access Key ID/Secret Access Key，简称AK/SK），是您通过开发工具（API、CLI、SDK）访问的身份凭证，不用于登录控制台。系统通过AK识别访问用户的身份，通过SK进行签名验证，通过加密签名验证可以确保请求的机密性、完整性和请求者身份的正确性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateAccessKeyV5(request *model.CreateAccessKeyV5Request) (*model.CreateAccessKeyV5Response, error) {
	requestDef := GenReqDefForCreateAccessKeyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessKeyV5Response), nil
	}
}

// CreateAccessKeyV5Invoker 创建永久访问密钥
func (c *IamClient) CreateAccessKeyV5Invoker(request *model.CreateAccessKeyV5Request) *CreateAccessKeyV5Invoker {
	requestDef := GenReqDefForCreateAccessKeyV5()
	return &CreateAccessKeyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLoginProfileV5 创建IAM用户登录信息
//
// 该接口可以用于创建指定IAM用户的登录信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) CreateLoginProfileV5(request *model.CreateLoginProfileV5Request) (*model.CreateLoginProfileV5Response, error) {
	requestDef := GenReqDefForCreateLoginProfileV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoginProfileV5Response), nil
	}
}

// CreateLoginProfileV5Invoker 创建IAM用户登录信息
func (c *IamClient) CreateLoginProfileV5Invoker(request *model.CreateLoginProfileV5Request) *CreateLoginProfileV5Invoker {
	requestDef := GenReqDefForCreateLoginProfileV5()
	return &CreateLoginProfileV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccessKeyV5 删除指定永久访问密钥
//
// 该接口可以用于删除IAM用户的指定永久访问密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteAccessKeyV5(request *model.DeleteAccessKeyV5Request) (*model.DeleteAccessKeyV5Response, error) {
	requestDef := GenReqDefForDeleteAccessKeyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessKeyV5Response), nil
	}
}

// DeleteAccessKeyV5Invoker 删除指定永久访问密钥
func (c *IamClient) DeleteAccessKeyV5Invoker(request *model.DeleteAccessKeyV5Request) *DeleteAccessKeyV5Invoker {
	requestDef := GenReqDefForDeleteAccessKeyV5()
	return &DeleteAccessKeyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLoginProfileV5 删除IAM用户登录信息
//
// 该接口可以用于删除指定IAM用户的登录信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) DeleteLoginProfileV5(request *model.DeleteLoginProfileV5Request) (*model.DeleteLoginProfileV5Response, error) {
	requestDef := GenReqDefForDeleteLoginProfileV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLoginProfileV5Response), nil
	}
}

// DeleteLoginProfileV5Invoker 删除IAM用户登录信息
func (c *IamClient) DeleteLoginProfileV5Invoker(request *model.DeleteLoginProfileV5Request) *DeleteLoginProfileV5Invoker {
	requestDef := GenReqDefForDeleteLoginProfileV5()
	return &DeleteLoginProfileV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessKeysV5 查询所有永久访问密钥
//
// 该接口可以用于查询IAM用户的所有永久访问密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ListAccessKeysV5(request *model.ListAccessKeysV5Request) (*model.ListAccessKeysV5Response, error) {
	requestDef := GenReqDefForListAccessKeysV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessKeysV5Response), nil
	}
}

// ListAccessKeysV5Invoker 查询所有永久访问密钥
func (c *IamClient) ListAccessKeysV5Invoker(request *model.ListAccessKeysV5Request) *ListAccessKeysV5Invoker {
	requestDef := GenReqDefForListAccessKeysV5()
	return &ListAccessKeysV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessKeyLastUsedV5 查询指定永久访问密钥最后使用时间
//
// 该接口可以用于查询IAM用户的指定永久访问密钥的最后使用时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowAccessKeyLastUsedV5(request *model.ShowAccessKeyLastUsedV5Request) (*model.ShowAccessKeyLastUsedV5Response, error) {
	requestDef := GenReqDefForShowAccessKeyLastUsedV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessKeyLastUsedV5Response), nil
	}
}

// ShowAccessKeyLastUsedV5Invoker 查询指定永久访问密钥最后使用时间
func (c *IamClient) ShowAccessKeyLastUsedV5Invoker(request *model.ShowAccessKeyLastUsedV5Request) *ShowAccessKeyLastUsedV5Invoker {
	requestDef := GenReqDefForShowAccessKeyLastUsedV5()
	return &ShowAccessKeyLastUsedV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLoginProfileV5 查询IAM用户登录信息
//
// 该接口可以用于查询指定IAM用户的登录信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) ShowLoginProfileV5(request *model.ShowLoginProfileV5Request) (*model.ShowLoginProfileV5Response, error) {
	requestDef := GenReqDefForShowLoginProfileV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLoginProfileV5Response), nil
	}
}

// ShowLoginProfileV5Invoker 查询IAM用户登录信息
func (c *IamClient) ShowLoginProfileV5Invoker(request *model.ShowLoginProfileV5Request) *ShowLoginProfileV5Invoker {
	requestDef := GenReqDefForShowLoginProfileV5()
	return &ShowLoginProfileV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessKeyV5 修改指定永久访问密钥
//
// 该接口可以用于修改IAM用户的指定永久访问密钥。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateAccessKeyV5(request *model.UpdateAccessKeyV5Request) (*model.UpdateAccessKeyV5Response, error) {
	requestDef := GenReqDefForUpdateAccessKeyV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessKeyV5Response), nil
	}
}

// UpdateAccessKeyV5Invoker 修改指定永久访问密钥
func (c *IamClient) UpdateAccessKeyV5Invoker(request *model.UpdateAccessKeyV5Request) *UpdateAccessKeyV5Invoker {
	requestDef := GenReqDefForUpdateAccessKeyV5()
	return &UpdateAccessKeyV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLoginProfileV5 修改IAM用户登录信息
//
// 该接口可以用于修改指定IAM用户的登录信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IamClient) UpdateLoginProfileV5(request *model.UpdateLoginProfileV5Request) (*model.UpdateLoginProfileV5Response, error) {
	requestDef := GenReqDefForUpdateLoginProfileV5()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLoginProfileV5Response), nil
	}
}

// UpdateLoginProfileV5Invoker 修改IAM用户登录信息
func (c *IamClient) UpdateLoginProfileV5Invoker(request *model.UpdateLoginProfileV5Request) *UpdateLoginProfileV5Invoker {
	requestDef := GenReqDefForUpdateLoginProfileV5()
	return &UpdateLoginProfileV5Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
