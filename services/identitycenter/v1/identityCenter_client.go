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

// CreateAccountAssignment 创建账户分配
//
// 使用指定的权限集为指定账户分配对应主体的访问权限，主体可为IdentityCenter用户或IdentityCenter用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// CreateAccountAssignmentInvoker 创建账户分配
func (c *IdentityCenterClient) CreateAccountAssignmentInvoker(request *model.CreateAccountAssignmentRequest) *CreateAccountAssignmentInvoker {
	requestDef := GenReqDefForCreateAccountAssignment()
	return &CreateAccountAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccountAssignment 删除账号分配
//
// 使用指定的权限集从指定的账号删除主体的访问权限，主体可为IAM身份中心用户或用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DescribeAccountAssignmentCreationStatus 查询账户分配创建状态详情
//
// 根据请求ID，查询指定IAM Identity Center实例下的账户分配创建状态详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DescribeAccountAssignmentCreationStatusInvoker 查询账户分配创建状态详情
func (c *IdentityCenterClient) DescribeAccountAssignmentCreationStatusInvoker(request *model.DescribeAccountAssignmentCreationStatusRequest) *DescribeAccountAssignmentCreationStatusInvoker {
	requestDef := GenReqDefForDescribeAccountAssignmentCreationStatus()
	return &DescribeAccountAssignmentCreationStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeAccountAssignmentDeletionStatus 查询账户分配删除状态详情
//
// 根据请求ID，查询指定IAM Identity Center实例下的账户分配删除状态详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DisassociateProfile 解除与用户或组绑定的所有账号授权关联
//
// 解除与用户或组绑定的所有账号授权关联。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DisassociateProfile(request *model.DisassociateProfileRequest) (*model.DisassociateProfileResponse, error) {
	requestDef := GenReqDefForDisassociateProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateProfileResponse), nil
	}
}

// DisassociateProfileInvoker 解除与用户或组绑定的所有账号授权关联
func (c *IdentityCenterClient) DisassociateProfileInvoker(request *model.DisassociateProfileRequest) *DisassociateProfileInvoker {
	requestDef := GenReqDefForDisassociateProfile()
	return &DisassociateProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentCreationStatus 列出账户分配创建状态
//
// 查询指定IAM Identity Center实例下的账户分配的创建状态列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 查询指定IAM Identity Center实例下的账户分配的删除状态列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListAccountAssignments 列出账户和权限集关联的用户或用户组
//
// 列出与指定账户以及指定权限集关联的用户或用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListAccountAssignmentsInvoker 列出账户和权限集关联的用户或用户组
func (c *IdentityCenterClient) ListAccountAssignmentsInvoker(request *model.ListAccountAssignmentsRequest) *ListAccountAssignmentsInvoker {
	requestDef := GenReqDefForListAccountAssignments()
	return &ListAccountAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountAssignmentsForPrincipal 检索与用户或用户组关联的账号列表
//
// 检索与用户或用户组关联的账号列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListAccountAssignmentsForPrincipal(request *model.ListAccountAssignmentsForPrincipalRequest) (*model.ListAccountAssignmentsForPrincipalResponse, error) {
	requestDef := GenReqDefForListAccountAssignmentsForPrincipal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccountAssignmentsForPrincipalResponse), nil
	}
}

// ListAccountAssignmentsForPrincipalInvoker 检索与用户或用户组关联的账号列表
func (c *IdentityCenterClient) ListAccountAssignmentsForPrincipalInvoker(request *model.ListAccountAssignmentsForPrincipalRequest) *ListAccountAssignmentsForPrincipalInvoker {
	requestDef := GenReqDefForListAccountAssignmentsForPrincipal()
	return &ListAccountAssignmentsForPrincipalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplicationInstance 创建应用程序实例
//
// 创建应用程序实例。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateApplicationInstance(request *model.CreateApplicationInstanceRequest) (*model.CreateApplicationInstanceResponse, error) {
	requestDef := GenReqDefForCreateApplicationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationInstanceResponse), nil
	}
}

// CreateApplicationInstanceInvoker 创建应用程序实例
func (c *IdentityCenterClient) CreateApplicationInstanceInvoker(request *model.CreateApplicationInstanceRequest) *CreateApplicationInstanceInvoker {
	requestDef := GenReqDefForCreateApplicationInstance()
	return &CreateApplicationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationInstance 删除应用程序实例
//
// 删除应用程序实例。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteApplicationInstance(request *model.DeleteApplicationInstanceRequest) (*model.DeleteApplicationInstanceResponse, error) {
	requestDef := GenReqDefForDeleteApplicationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationInstanceResponse), nil
	}
}

// DeleteApplicationInstanceInvoker 删除应用程序实例
func (c *IdentityCenterClient) DeleteApplicationInstanceInvoker(request *model.DeleteApplicationInstanceRequest) *DeleteApplicationInstanceInvoker {
	requestDef := GenReqDefForDeleteApplicationInstance()
	return &DeleteApplicationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteProfile 删除应用程序实例与用户或用户组关联关系
//
// 删除应用程序实例与用户或用户组关联关系。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteProfile(request *model.DeleteProfileRequest) (*model.DeleteProfileResponse, error) {
	requestDef := GenReqDefForDeleteProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteProfileResponse), nil
	}
}

// DeleteProfileInvoker 删除应用程序实例与用户或用户组关联关系
func (c *IdentityCenterClient) DeleteProfileInvoker(request *model.DeleteProfileRequest) *DeleteProfileInvoker {
	requestDef := GenReqDefForDeleteProfile()
	return &DeleteProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeApplication 查询应用程序详情
//
// 查询应用程序详情。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeApplication(request *model.DescribeApplicationRequest) (*model.DescribeApplicationResponse, error) {
	requestDef := GenReqDefForDescribeApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeApplicationResponse), nil
	}
}

// DescribeApplicationInvoker 查询应用程序详情
func (c *IdentityCenterClient) DescribeApplicationInvoker(request *model.DescribeApplicationRequest) *DescribeApplicationInvoker {
	requestDef := GenReqDefForDescribeApplication()
	return &DescribeApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeApplicationProvider 查询应用程序提供者详情
//
// 查询应用程序提供者详情。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeApplicationProvider(request *model.DescribeApplicationProviderRequest) (*model.DescribeApplicationProviderResponse, error) {
	requestDef := GenReqDefForDescribeApplicationProvider()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeApplicationProviderResponse), nil
	}
}

// DescribeApplicationProviderInvoker 查询应用程序提供者详情
func (c *IdentityCenterClient) DescribeApplicationProviderInvoker(request *model.DescribeApplicationProviderRequest) *DescribeApplicationProviderInvoker {
	requestDef := GenReqDefForDescribeApplicationProvider()
	return &DescribeApplicationProviderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetApplicationAssignmentConfiguration 查询应用程序分配属性配置
//
// 查询应用程序分配属性配置，目的为用户或者用户组分配对应用程序的访问权限。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetApplicationAssignmentConfiguration(request *model.GetApplicationAssignmentConfigurationRequest) (*model.GetApplicationAssignmentConfigurationResponse, error) {
	requestDef := GenReqDefForGetApplicationAssignmentConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetApplicationAssignmentConfigurationResponse), nil
	}
}

// GetApplicationAssignmentConfigurationInvoker 查询应用程序分配属性配置
func (c *IdentityCenterClient) GetApplicationAssignmentConfigurationInvoker(request *model.GetApplicationAssignmentConfigurationRequest) *GetApplicationAssignmentConfigurationInvoker {
	requestDef := GenReqDefForGetApplicationAssignmentConfiguration()
	return &GetApplicationAssignmentConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetApplicationInstance 查询应用程序实例详情
//
// 查询应用程序实例详情。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetApplicationInstance(request *model.GetApplicationInstanceRequest) (*model.GetApplicationInstanceResponse, error) {
	requestDef := GenReqDefForGetApplicationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetApplicationInstanceResponse), nil
	}
}

// GetApplicationInstanceInvoker 查询应用程序实例详情
func (c *IdentityCenterClient) GetApplicationInstanceInvoker(request *model.GetApplicationInstanceRequest) *GetApplicationInstanceInvoker {
	requestDef := GenReqDefForGetApplicationInstance()
	return &GetApplicationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportApplicationInstanceServiceProviderMetadata 上传应用程序实例元数据文件
//
// 上传应用程序实例元数据文件。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ImportApplicationInstanceServiceProviderMetadata(request *model.ImportApplicationInstanceServiceProviderMetadataRequest) (*model.ImportApplicationInstanceServiceProviderMetadataResponse, error) {
	requestDef := GenReqDefForImportApplicationInstanceServiceProviderMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportApplicationInstanceServiceProviderMetadataResponse), nil
	}
}

// ImportApplicationInstanceServiceProviderMetadataInvoker 上传应用程序实例元数据文件
func (c *IdentityCenterClient) ImportApplicationInstanceServiceProviderMetadataInvoker(request *model.ImportApplicationInstanceServiceProviderMetadataRequest) *ImportApplicationInstanceServiceProviderMetadataInvoker {
	requestDef := GenReqDefForImportApplicationInstanceServiceProviderMetadata()
	return &ImportApplicationInstanceServiceProviderMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationInstances 列出应用程序实例
//
// 列出应用程序实例。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationInstances(request *model.ListApplicationInstancesRequest) (*model.ListApplicationInstancesResponse, error) {
	requestDef := GenReqDefForListApplicationInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationInstancesResponse), nil
	}
}

// ListApplicationInstancesInvoker 列出应用程序实例
func (c *IdentityCenterClient) ListApplicationInstancesInvoker(request *model.ListApplicationInstancesRequest) *ListApplicationInstancesInvoker {
	requestDef := GenReqDefForListApplicationInstances()
	return &ListApplicationInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationProviders 列出应用程序提供者
//
// 查询应用程序提供者列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationProviders(request *model.ListApplicationProvidersRequest) (*model.ListApplicationProvidersResponse, error) {
	requestDef := GenReqDefForListApplicationProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationProvidersResponse), nil
	}
}

// ListApplicationProvidersInvoker 列出应用程序提供者
func (c *IdentityCenterClient) ListApplicationProvidersInvoker(request *model.ListApplicationProvidersRequest) *ListApplicationProvidersInvoker {
	requestDef := GenReqDefForListApplicationProviders()
	return &ListApplicationProvidersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationTemplates 列出应用程序模板
//
// 查询应用程序模板列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationTemplates(request *model.ListApplicationTemplatesRequest) (*model.ListApplicationTemplatesResponse, error) {
	requestDef := GenReqDefForListApplicationTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationTemplatesResponse), nil
	}
}

// ListApplicationTemplatesInvoker 列出应用程序模板
func (c *IdentityCenterClient) ListApplicationTemplatesInvoker(request *model.ListApplicationTemplatesRequest) *ListApplicationTemplatesInvoker {
	requestDef := GenReqDefForListApplicationTemplates()
	return &ListApplicationTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplications 列出应用程序
//
// 查询应用程序列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 列出应用程序
func (c *IdentityCenterClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogApplications 列出应用程序目录中的预置应用模板
//
// 列出应用程序目录中的预置应用模板。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListCatalogApplications(request *model.ListCatalogApplicationsRequest) (*model.ListCatalogApplicationsResponse, error) {
	requestDef := GenReqDefForListCatalogApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogApplicationsResponse), nil
	}
}

// ListCatalogApplicationsInvoker 列出应用程序目录中的预置应用模板
func (c *IdentityCenterClient) ListCatalogApplicationsInvoker(request *model.ListCatalogApplicationsRequest) *ListCatalogApplicationsInvoker {
	requestDef := GenReqDefForListCatalogApplications()
	return &ListCatalogApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProfiles 列出应用程序实例与用户或用户组存在的关联关系
//
// 列出应用程序实例与用户或用户组存在的关联关系。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListProfiles(request *model.ListProfilesRequest) (*model.ListProfilesResponse, error) {
	requestDef := GenReqDefForListProfiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProfilesResponse), nil
	}
}

// ListProfilesInvoker 列出应用程序实例与用户或用户组存在的关联关系
func (c *IdentityCenterClient) ListProfilesInvoker(request *model.ListProfilesRequest) *ListProfilesInvoker {
	requestDef := GenReqDefForListProfiles()
	return &ListProfilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceDisplayData 更新应用程序实例显示信息
//
// 更新应用程序实例显示信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceDisplayData(request *model.UpdateApplicationInstanceDisplayDataRequest) (*model.UpdateApplicationInstanceDisplayDataResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceDisplayData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceDisplayDataResponse), nil
	}
}

// UpdateApplicationInstanceDisplayDataInvoker 更新应用程序实例显示信息
func (c *IdentityCenterClient) UpdateApplicationInstanceDisplayDataInvoker(request *model.UpdateApplicationInstanceDisplayDataRequest) *UpdateApplicationInstanceDisplayDataInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceDisplayData()
	return &UpdateApplicationInstanceDisplayDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceResponseConfiguration 更新应用程序属性配置
//
// 更新应用程序属性配置信息，更新应用程序中的属性映射、中继状态以及会话过期时间。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceResponseConfiguration(request *model.UpdateApplicationInstanceResponseConfigurationRequest) (*model.UpdateApplicationInstanceResponseConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceResponseConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceResponseConfigurationResponse), nil
	}
}

// UpdateApplicationInstanceResponseConfigurationInvoker 更新应用程序属性配置
func (c *IdentityCenterClient) UpdateApplicationInstanceResponseConfigurationInvoker(request *model.UpdateApplicationInstanceResponseConfigurationRequest) *UpdateApplicationInstanceResponseConfigurationInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceResponseConfiguration()
	return &UpdateApplicationInstanceResponseConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceResponseSchemaConfiguration 更新应用程序Schema属性映射配置
//
// 更新应用程序Schema属性映射配置，支持SAML断言中Subject属性映射以及Subject NameID格式。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceResponseSchemaConfiguration(request *model.UpdateApplicationInstanceResponseSchemaConfigurationRequest) (*model.UpdateApplicationInstanceResponseSchemaConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceResponseSchemaConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceResponseSchemaConfigurationResponse), nil
	}
}

// UpdateApplicationInstanceResponseSchemaConfigurationInvoker 更新应用程序Schema属性映射配置
func (c *IdentityCenterClient) UpdateApplicationInstanceResponseSchemaConfigurationInvoker(request *model.UpdateApplicationInstanceResponseSchemaConfigurationRequest) *UpdateApplicationInstanceResponseSchemaConfigurationInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceResponseSchemaConfiguration()
	return &UpdateApplicationInstanceResponseSchemaConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceSecurityConfiguration 更新应用程序实例证书配置
//
// 更新应用程序实例证书配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceSecurityConfiguration(request *model.UpdateApplicationInstanceSecurityConfigurationRequest) (*model.UpdateApplicationInstanceSecurityConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceSecurityConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceSecurityConfigurationResponse), nil
	}
}

// UpdateApplicationInstanceSecurityConfigurationInvoker 更新应用程序实例证书配置
func (c *IdentityCenterClient) UpdateApplicationInstanceSecurityConfigurationInvoker(request *model.UpdateApplicationInstanceSecurityConfigurationRequest) *UpdateApplicationInstanceSecurityConfigurationInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceSecurityConfiguration()
	return &UpdateApplicationInstanceSecurityConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceServiceProviderConfiguration 更新应用程序实例服务提供商配置
//
// 更新应用程序实例服务提供商配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceServiceProviderConfiguration(request *model.UpdateApplicationInstanceServiceProviderConfigurationRequest) (*model.UpdateApplicationInstanceServiceProviderConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceServiceProviderConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceServiceProviderConfigurationResponse), nil
	}
}

// UpdateApplicationInstanceServiceProviderConfigurationInvoker 更新应用程序实例服务提供商配置
func (c *IdentityCenterClient) UpdateApplicationInstanceServiceProviderConfigurationInvoker(request *model.UpdateApplicationInstanceServiceProviderConfigurationRequest) *UpdateApplicationInstanceServiceProviderConfigurationInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceServiceProviderConfiguration()
	return &UpdateApplicationInstanceServiceProviderConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceStatus 更新应用程序实例状态
//
// 更新应用程序实例状态。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceStatus(request *model.UpdateApplicationInstanceStatusRequest) (*model.UpdateApplicationInstanceStatusResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceStatusResponse), nil
	}
}

// UpdateApplicationInstanceStatusInvoker 更新应用程序实例状态
func (c *IdentityCenterClient) UpdateApplicationInstanceStatusInvoker(request *model.UpdateApplicationInstanceStatusRequest) *UpdateApplicationInstanceStatusInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceStatus()
	return &UpdateApplicationInstanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplicationAssignment 应用程序分配用户或用户组
//
// 应用程序分配用户或用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateApplicationAssignment(request *model.CreateApplicationAssignmentRequest) (*model.CreateApplicationAssignmentResponse, error) {
	requestDef := GenReqDefForCreateApplicationAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationAssignmentResponse), nil
	}
}

// CreateApplicationAssignmentInvoker 应用程序分配用户或用户组
func (c *IdentityCenterClient) CreateApplicationAssignmentInvoker(request *model.CreateApplicationAssignmentRequest) *CreateApplicationAssignmentInvoker {
	requestDef := GenReqDefForCreateApplicationAssignment()
	return &CreateApplicationAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationAssignment 删除应用程序已分配用户或用户组
//
// 删除应用程序已分配用户或用户组。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteApplicationAssignment(request *model.DeleteApplicationAssignmentRequest) (*model.DeleteApplicationAssignmentResponse, error) {
	requestDef := GenReqDefForDeleteApplicationAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationAssignmentResponse), nil
	}
}

// DeleteApplicationAssignmentInvoker 删除应用程序已分配用户或用户组
func (c *IdentityCenterClient) DeleteApplicationAssignmentInvoker(request *model.DeleteApplicationAssignmentRequest) *DeleteApplicationAssignmentInvoker {
	requestDef := GenReqDefForDeleteApplicationAssignment()
	return &DeleteApplicationAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationAssignments 查询应用程序已分配的用户或用户组列表
//
// 查询应用程序已分配的用户或用户组列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationAssignments(request *model.ListApplicationAssignmentsRequest) (*model.ListApplicationAssignmentsResponse, error) {
	requestDef := GenReqDefForListApplicationAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationAssignmentsResponse), nil
	}
}

// ListApplicationAssignmentsInvoker 查询应用程序已分配的用户或用户组列表
func (c *IdentityCenterClient) ListApplicationAssignmentsInvoker(request *model.ListApplicationAssignmentsRequest) *ListApplicationAssignmentsInvoker {
	requestDef := GenReqDefForListApplicationAssignments()
	return &ListApplicationAssignmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationAssignmentsForPrincipal 检索与用户或用户组关联的应用程序列表
//
// 检索与用户或用户组关联的应用程序列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationAssignmentsForPrincipal(request *model.ListApplicationAssignmentsForPrincipalRequest) (*model.ListApplicationAssignmentsForPrincipalResponse, error) {
	requestDef := GenReqDefForListApplicationAssignmentsForPrincipal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationAssignmentsForPrincipalResponse), nil
	}
}

// ListApplicationAssignmentsForPrincipalInvoker 检索与用户或用户组关联的应用程序列表
func (c *IdentityCenterClient) ListApplicationAssignmentsForPrincipalInvoker(request *model.ListApplicationAssignmentsForPrincipalRequest) *ListApplicationAssignmentsForPrincipalInvoker {
	requestDef := GenReqDefForListApplicationAssignmentsForPrincipal()
	return &ListApplicationAssignmentsForPrincipalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplicationInstanceCertificate 创建应用程序实例证书
//
// 创建应用程序实例证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateApplicationInstanceCertificate(request *model.CreateApplicationInstanceCertificateRequest) (*model.CreateApplicationInstanceCertificateResponse, error) {
	requestDef := GenReqDefForCreateApplicationInstanceCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationInstanceCertificateResponse), nil
	}
}

// CreateApplicationInstanceCertificateInvoker 创建应用程序实例证书
func (c *IdentityCenterClient) CreateApplicationInstanceCertificateInvoker(request *model.CreateApplicationInstanceCertificateRequest) *CreateApplicationInstanceCertificateInvoker {
	requestDef := GenReqDefForCreateApplicationInstanceCertificate()
	return &CreateApplicationInstanceCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationInstanceCertificate 删除应用程序实例证书
//
// 删除应用程序实例证书。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteApplicationInstanceCertificate(request *model.DeleteApplicationInstanceCertificateRequest) (*model.DeleteApplicationInstanceCertificateResponse, error) {
	requestDef := GenReqDefForDeleteApplicationInstanceCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationInstanceCertificateResponse), nil
	}
}

// DeleteApplicationInstanceCertificateInvoker 删除应用程序实例证书
func (c *IdentityCenterClient) DeleteApplicationInstanceCertificateInvoker(request *model.DeleteApplicationInstanceCertificateRequest) *DeleteApplicationInstanceCertificateInvoker {
	requestDef := GenReqDefForDeleteApplicationInstanceCertificate()
	return &DeleteApplicationInstanceCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationInstanceCertificates 列出应用程序实例证书
//
// 查询应用程序实例证书列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListApplicationInstanceCertificates(request *model.ListApplicationInstanceCertificatesRequest) (*model.ListApplicationInstanceCertificatesResponse, error) {
	requestDef := GenReqDefForListApplicationInstanceCertificates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationInstanceCertificatesResponse), nil
	}
}

// ListApplicationInstanceCertificatesInvoker 列出应用程序实例证书
func (c *IdentityCenterClient) ListApplicationInstanceCertificatesInvoker(request *model.ListApplicationInstanceCertificatesRequest) *ListApplicationInstanceCertificatesInvoker {
	requestDef := GenReqDefForListApplicationInstanceCertificates()
	return &ListApplicationInstanceCertificatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationInstanceActiveCertificate 激活应用程序实例证书
//
// 激活应用程序实例证书，实现证书轮转。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateApplicationInstanceActiveCertificate(request *model.UpdateApplicationInstanceActiveCertificateRequest) (*model.UpdateApplicationInstanceActiveCertificateResponse, error) {
	requestDef := GenReqDefForUpdateApplicationInstanceActiveCertificate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationInstanceActiveCertificateResponse), nil
	}
}

// UpdateApplicationInstanceActiveCertificateInvoker 激活应用程序实例证书
func (c *IdentityCenterClient) UpdateApplicationInstanceActiveCertificateInvoker(request *model.UpdateApplicationInstanceActiveCertificateRequest) *UpdateApplicationInstanceActiveCertificateInvoker {
	requestDef := GenReqDefForUpdateApplicationInstanceActiveCertificate()
	return &UpdateApplicationInstanceActiveCertificateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetSsoConfiguration 查询实例配置信息
//
// 查询IAM身份中心实例配置信息，包括身份认证配置和会话管理配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetSsoConfiguration(request *model.GetSsoConfigurationRequest) (*model.GetSsoConfigurationResponse, error) {
	requestDef := GenReqDefForGetSsoConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetSsoConfigurationResponse), nil
	}
}

// GetSsoConfigurationInvoker 查询实例配置信息
func (c *IdentityCenterClient) GetSsoConfigurationInvoker(request *model.GetSsoConfigurationRequest) *GetSsoConfigurationInvoker {
	requestDef := GenReqDefForGetSsoConfiguration()
	return &GetSsoConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSsoConfiguration 设置实例配置信息
//
// 设置IAM身份中心服务实例配置信息，包括身份认证配置和会话管理配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateSsoConfiguration(request *model.UpdateSsoConfigurationRequest) (*model.UpdateSsoConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateSsoConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSsoConfigurationResponse), nil
	}
}

// UpdateSsoConfigurationInvoker 设置实例配置信息
func (c *IdentityCenterClient) UpdateSsoConfigurationInvoker(request *model.UpdateSsoConfigurationRequest) *UpdateSsoConfigurationInvoker {
	requestDef := GenReqDefForUpdateSsoConfiguration()
	return &UpdateSsoConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlias 自定义访问门户URL
//
// 自定义访问门户URL，默认情况下，您可以使用遵循以下格式的 URL访问门户：idcenter.huaweicloud.com/d-xxxxxxxxxx/portal，您可以按如下方式更改为自定义 URL：idcenter.huaweicloud.com/your_subdomain/portal。设置自定义访问门户URL是一次性操作，无法撤销。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) CreateAlias(request *model.CreateAliasRequest) (*model.CreateAliasResponse, error) {
	requestDef := GenReqDefForCreateAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAliasResponse), nil
	}
}

// CreateAliasInvoker 自定义访问门户URL
func (c *IdentityCenterClient) CreateAliasInvoker(request *model.CreateAliasRequest) *CreateAliasInvoker {
	requestDef := GenReqDefForCreateAlias()
	return &CreateAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIdentityCenter 删除服务实例
//
// 删除IAM Identity Center服务实例。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DeleteIdentityCenter(request *model.DeleteIdentityCenterRequest) (*model.DeleteIdentityCenterResponse, error) {
	requestDef := GenReqDefForDeleteIdentityCenter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIdentityCenterResponse), nil
	}
}

// DeleteIdentityCenterInvoker 删除服务实例
func (c *IdentityCenterClient) DeleteIdentityCenterInvoker(request *model.DeleteIdentityCenterRequest) *DeleteIdentityCenterInvoker {
	requestDef := GenReqDefForDeleteIdentityCenter()
	return &DeleteIdentityCenterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeRegisteredRegions 查询服务实例开通所在区域
//
// 查询IAM身份中心服务实例开通后，具体开通所在区域。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) DescribeRegisteredRegions(request *model.DescribeRegisteredRegionsRequest) (*model.DescribeRegisteredRegionsResponse, error) {
	requestDef := GenReqDefForDescribeRegisteredRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeRegisteredRegionsResponse), nil
	}
}

// DescribeRegisteredRegionsInvoker 查询服务实例开通所在区域
func (c *IdentityCenterClient) DescribeRegisteredRegionsInvoker(request *model.DescribeRegisteredRegionsRequest) *DescribeRegisteredRegionsInvoker {
	requestDef := GenReqDefForDescribeRegisteredRegions()
	return &DescribeRegisteredRegionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetHaConfiguration 查询高可用功能配置
//
// 查询高可用功能配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetHaConfiguration(request *model.GetHaConfigurationRequest) (*model.GetHaConfigurationResponse, error) {
	requestDef := GenReqDefForGetHaConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetHaConfigurationResponse), nil
	}
}

// GetHaConfigurationInvoker 查询高可用功能配置
func (c *IdentityCenterClient) GetHaConfigurationInvoker(request *model.GetHaConfigurationRequest) *GetHaConfigurationInvoker {
	requestDef := GenReqDefForGetHaConfiguration()
	return &GetHaConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetIdentityCenterServiceStatus 查询服务实例状态
//
// 查询IAM Identity Center服务实例状态信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetIdentityCenterServiceStatus(request *model.GetIdentityCenterServiceStatusRequest) (*model.GetIdentityCenterServiceStatusResponse, error) {
	requestDef := GenReqDefForGetIdentityCenterServiceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetIdentityCenterServiceStatusResponse), nil
	}
}

// GetIdentityCenterServiceStatusInvoker 查询服务实例状态
func (c *IdentityCenterClient) GetIdentityCenterServiceStatusInvoker(request *model.GetIdentityCenterServiceStatusRequest) *GetIdentityCenterServiceStatusInvoker {
	requestDef := GenReqDefForGetIdentityCenterServiceStatus()
	return &GetIdentityCenterServiceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIdentityStoreAssociation 获取身份源配置
//
// 获取IAM身份中心服务实例中的身份源配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListIdentityStoreAssociation(request *model.ListIdentityStoreAssociationRequest) (*model.ListIdentityStoreAssociationResponse, error) {
	requestDef := GenReqDefForListIdentityStoreAssociation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIdentityStoreAssociationResponse), nil
	}
}

// ListIdentityStoreAssociationInvoker 获取身份源配置
func (c *IdentityCenterClient) ListIdentityStoreAssociationInvoker(request *model.ListIdentityStoreAssociationRequest) *ListIdentityStoreAssociationInvoker {
	requestDef := GenReqDefForListIdentityStoreAssociation()
	return &ListIdentityStoreAssociationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 列出实例
//
// 查询IAM身份中心的实例列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// RegisterRegion 选择服务实例开通区域
//
// IAM身份中心服务实例开通前，需要选择服务实例具体开通在某一区域。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) RegisterRegion(request *model.RegisterRegionRequest) (*model.RegisterRegionResponse, error) {
	requestDef := GenReqDefForRegisterRegion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterRegionResponse), nil
	}
}

// RegisterRegionInvoker 选择服务实例开通区域
func (c *IdentityCenterClient) RegisterRegionInvoker(request *model.RegisterRegionRequest) *RegisterRegionInvoker {
	requestDef := GenReqDefForRegisterRegion()
	return &RegisterRegionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartIdentityCenter 开通服务实例
//
// 开通IAM Identity Center服务实例。此操作只能由组织的管理账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) StartIdentityCenter(request *model.StartIdentityCenterRequest) (*model.StartIdentityCenterResponse, error) {
	requestDef := GenReqDefForStartIdentityCenter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartIdentityCenterResponse), nil
	}
}

// StartIdentityCenterInvoker 开通服务实例
func (c *IdentityCenterClient) StartIdentityCenterInvoker(request *model.StartIdentityCenterRequest) *StartIdentityCenterInvoker {
	requestDef := GenReqDefForStartIdentityCenter()
	return &StartIdentityCenterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHaConfiguration 更新高可用功能配置
//
// 授权启用或者禁用高可用功能配置。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) UpdateHaConfiguration(request *model.UpdateHaConfigurationRequest) (*model.UpdateHaConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateHaConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHaConfigurationResponse), nil
	}
}

// UpdateHaConfigurationInvoker 更新高可用功能配置
func (c *IdentityCenterClient) UpdateHaConfigurationInvoker(request *model.UpdateHaConfigurationRequest) *UpdateHaConfigurationInvoker {
	requestDef := GenReqDefForUpdateHaConfiguration()
	return &UpdateHaConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstanceAccessControlAttributeConfiguration 启用指定实例的访问控制功能
//
// 启用指定实例的访问控制功能。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 禁用指定IAM Identity Center实例的基于属性的访问控制（ABAC）功能，并删除已配置的所有属性映射。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 返回已配置为与指定IAM Identity Center实例的基于属性的访问控制（ABAC）一起使用的IAM Identity Center身份存储属性列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 更新可与IAM Identity Center实例一起使用的IAM Identity Center身份存储属性，以进行基于属性的访问控制（ABAC）。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// GetMfaDeviceManagementForIdentityStore 查询MFA管理配置信息
//
// 查询MFA管理配置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetMfaDeviceManagementForIdentityStore(request *model.GetMfaDeviceManagementForIdentityStoreRequest) (*model.GetMfaDeviceManagementForIdentityStoreResponse, error) {
	requestDef := GenReqDefForGetMfaDeviceManagementForIdentityStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetMfaDeviceManagementForIdentityStoreResponse), nil
	}
}

// GetMfaDeviceManagementForIdentityStoreInvoker 查询MFA管理配置信息
func (c *IdentityCenterClient) GetMfaDeviceManagementForIdentityStoreInvoker(request *model.GetMfaDeviceManagementForIdentityStoreRequest) *GetMfaDeviceManagementForIdentityStoreInvoker {
	requestDef := GenReqDefForGetMfaDeviceManagementForIdentityStore()
	return &GetMfaDeviceManagementForIdentityStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutMfaDeviceManagementForIdentityStore 设置MFA管理设置信息
//
// 设置MFA管理设置信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) PutMfaDeviceManagementForIdentityStore(request *model.PutMfaDeviceManagementForIdentityStoreRequest) (*model.PutMfaDeviceManagementForIdentityStoreResponse, error) {
	requestDef := GenReqDefForPutMfaDeviceManagementForIdentityStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutMfaDeviceManagementForIdentityStoreResponse), nil
	}
}

// PutMfaDeviceManagementForIdentityStoreInvoker 设置MFA管理设置信息
func (c *IdentityCenterClient) PutMfaDeviceManagementForIdentityStoreInvoker(request *model.PutMfaDeviceManagementForIdentityStoreRequest) *PutMfaDeviceManagementForIdentityStoreInvoker {
	requestDef := GenReqDefForPutMfaDeviceManagementForIdentityStore()
	return &PutMfaDeviceManagementForIdentityStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachManagedPolicyToPermissionSet 添加系统身份策略
//
// 在指定的权限集中添加系统身份策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// AttachManagedRoleToPermissionSet 附加系统策略到权限集
//
// 将系统策略附加到权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// AttachManagedRoleToPermissionSetInvoker 附加系统策略到权限集
func (c *IdentityCenterClient) AttachManagedRoleToPermissionSetInvoker(request *model.AttachManagedRoleToPermissionSetRequest) *AttachManagedRoleToPermissionSetInvoker {
	requestDef := GenReqDefForAttachManagedRoleToPermissionSet()
	return &AttachManagedRoleToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePermissionSet 创建权限集
//
// 在指定的IAM Identity Center实例中创建一个权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DeleteCustomPolicyFromPermissionSet 删除指定权限集中的自定义身份策略
//
// 删除指定权限集中的自定义身份策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DeleteCustomPolicyFromPermissionSetInvoker 删除指定权限集中的自定义身份策略
func (c *IdentityCenterClient) DeleteCustomPolicyFromPermissionSetInvoker(request *model.DeleteCustomPolicyFromPermissionSetRequest) *DeleteCustomPolicyFromPermissionSetInvoker {
	requestDef := GenReqDefForDeleteCustomPolicyFromPermissionSet()
	return &DeleteCustomPolicyFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCustomRoleFromPermissionSet 删除指定权限集中的自定义策略
//
// 删除指定权限集中的自定义策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 根据权限集ID，删除指定的权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 根据权限集ID，查询指定权限集的详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 根据请求Id，查询权限集预分配状态的详情信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DetachManagedPolicyFromPermissionSet 从权限集分离系统身份策略
//
// 将附加的系统身份策略从指定的权限集中分离。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DetachManagedPolicyFromPermissionSetInvoker 从权限集分离系统身份策略
func (c *IdentityCenterClient) DetachManagedPolicyFromPermissionSetInvoker(request *model.DetachManagedPolicyFromPermissionSetRequest) *DetachManagedPolicyFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachManagedPolicyFromPermissionSet()
	return &DetachManagedPolicyFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachManagedRoleFromPermissionSet 从权限集分离系统策略
//
// 将附加的系统策略从指定的权限集中分离。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DetachManagedRoleFromPermissionSetInvoker 从权限集分离系统策略
func (c *IdentityCenterClient) DetachManagedRoleFromPermissionSetInvoker(request *model.DetachManagedRoleFromPermissionSetRequest) *DetachManagedRoleFromPermissionSetInvoker {
	requestDef := GenReqDefForDetachManagedRoleFromPermissionSet()
	return &DetachManagedRoleFromPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetCustomPolicyForPermissionSet 获取分配给权限集的自定义身份策略
//
// 获取分配给权限集的自定义身份策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// GetCustomPolicyForPermissionSetInvoker 获取分配给权限集的自定义身份策略
func (c *IdentityCenterClient) GetCustomPolicyForPermissionSetInvoker(request *model.GetCustomPolicyForPermissionSetRequest) *GetCustomPolicyForPermissionSetInvoker {
	requestDef := GenReqDefForGetCustomPolicyForPermissionSet()
	return &GetCustomPolicyForPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetCustomRoleForPermissionSet 获取分配给权限集的自定义策略
//
// 获取分配给权限集的自定义策略。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// GetPermissionSetSummary 查询权限集配额信息
//
// 查询权限集配额信息。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) GetPermissionSetSummary(request *model.GetPermissionSetSummaryRequest) (*model.GetPermissionSetSummaryResponse, error) {
	requestDef := GenReqDefForGetPermissionSetSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetPermissionSetSummaryResponse), nil
	}
}

// GetPermissionSetSummaryInvoker 查询权限集配额信息
func (c *IdentityCenterClient) GetPermissionSetSummaryInvoker(request *model.GetPermissionSetSummaryRequest) *GetPermissionSetSummaryInvoker {
	requestDef := GenReqDefForGetPermissionSetSummary()
	return &GetPermissionSetSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccountsForProvisionedPermissionSet 列出权限集关联的账号
//
// 查询与指定权限集关联的账户列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListManagedPoliciesInPermissionSet 列出权限集中附加的系统身份策略
//
// 获取附加到指定权限集的系统身份策略列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListManagedPoliciesInPermissionSetInvoker 列出权限集中附加的系统身份策略
func (c *IdentityCenterClient) ListManagedPoliciesInPermissionSetInvoker(request *model.ListManagedPoliciesInPermissionSetRequest) *ListManagedPoliciesInPermissionSetInvoker {
	requestDef := GenReqDefForListManagedPoliciesInPermissionSet()
	return &ListManagedPoliciesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagedRolesInPermissionSet 列出权限集中附加的系统策略
//
// 获取附加到指定权限集的系统策略列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListManagedRolesInPermissionSetInvoker 列出权限集中附加的系统策略
func (c *IdentityCenterClient) ListManagedRolesInPermissionSetInvoker(request *model.ListManagedRolesInPermissionSetRequest) *ListManagedRolesInPermissionSetInvoker {
	requestDef := GenReqDefForListManagedRolesInPermissionSet()
	return &ListManagedRolesInPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissionSetProvisioningStatus 列出权限集预分配状态
//
// 查询指定实例中的权限集预分配状态列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 查询指定实例下的权限集列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListPermissionSetsProvisionedToAccount 列出分配给账户的权限集
//
// 查询分配给指定账户的权限集列表。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// ListPermissionSetsProvisionedToAccountInvoker 列出分配给账户的权限集
func (c *IdentityCenterClient) ListPermissionSetsProvisionedToAccountInvoker(request *model.ListPermissionSetsProvisionedToAccountRequest) *ListPermissionSetsProvisionedToAccountInvoker {
	requestDef := GenReqDefForListPermissionSetsProvisionedToAccount()
	return &ListPermissionSetsProvisionedToAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ProvisionPermissionSet 预分配权限集
//
// 将指定权限集预分配给指定账户。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// PutCustomPolicyToPermissionSet 将自定义身份策略附加到权限集
//
// 将自定义身份策略附加到权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// PutCustomPolicyToPermissionSetInvoker 将自定义身份策略附加到权限集
func (c *IdentityCenterClient) PutCustomPolicyToPermissionSetInvoker(request *model.PutCustomPolicyToPermissionSetRequest) *PutCustomPolicyToPermissionSetInvoker {
	requestDef := GenReqDefForPutCustomPolicyToPermissionSet()
	return &PutCustomPolicyToPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutCustomRoleToPermissionSet 将自定义策略附加到权限集
//
// 将自定义策略附加到权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 根据权限集ID，更新指定权限集的属性。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
// 向指定的资源添加一个或多个标签。目前，您可以将标签附加到实例中的权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DeleteTagResource 从指定资源中删除指定主键标签
//
// 从指定资源中删除具有指定主键的任何标签。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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

// DeleteTagResourceInvoker 从指定资源中删除指定主键标签
func (c *IdentityCenterClient) DeleteTagResourceInvoker(request *model.DeleteTagResourceRequest) *DeleteTagResourceInvoker {
	requestDef := GenReqDefForDeleteTagResource()
	return &DeleteTagResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagResources 列出绑定到指定资源的标签
//
// 列出绑定到指定资源的标签。您可以将标签附加到实例中的权限集。此操作只能由组织的管理账号或作为服务委托管理员的成员账号调用。
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
