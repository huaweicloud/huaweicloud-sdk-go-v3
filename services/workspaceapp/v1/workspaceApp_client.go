package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspaceapp/v1/model"
)

type WorkspaceAppClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewWorkspaceAppClient(hcClient *httpclient.HcHttpClient) *WorkspaceAppClient {
	return &WorkspaceAppClient{HcClient: hcClient}
}

func WorkspaceAppClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AuthorizeObs 获取上传至OBS桶的临时ak/sk
//
// 获取上传至OBS桶的临时ak/sk。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) AuthorizeObs(request *model.AuthorizeObsRequest) (*model.AuthorizeObsResponse, error) {
	requestDef := GenReqDefForAuthorizeObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeObsResponse), nil
	}
}

// AuthorizeObsInvoker 获取上传至OBS桶的临时ak/sk
func (c *WorkspaceAppClient) AuthorizeObsInvoker(request *model.AuthorizeObsRequest) *AuthorizeObsInvoker {
	requestDef := GenReqDefForAuthorizeObs()
	return &AuthorizeObsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteWarehouseApp 批量删除应用仓库中的指定应用
//
// 批量删除应用仓库中的指定应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteWarehouseApp(request *model.BatchDeleteWarehouseAppRequest) (*model.BatchDeleteWarehouseAppResponse, error) {
	requestDef := GenReqDefForBatchDeleteWarehouseApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteWarehouseAppResponse), nil
	}
}

// BatchDeleteWarehouseAppInvoker 批量删除应用仓库中的指定应用
func (c *WorkspaceAppClient) BatchDeleteWarehouseAppInvoker(request *model.BatchDeleteWarehouseAppRequest) *BatchDeleteWarehouseAppInvoker {
	requestDef := GenReqDefForBatchDeleteWarehouseApp()
	return &BatchDeleteWarehouseAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindAppWarehouseBucket 添加用户应用仓库桶及桶授权
//
// 添加用户应用仓库桶及桶授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BindAppWarehouseBucket(request *model.BindAppWarehouseBucketRequest) (*model.BindAppWarehouseBucketResponse, error) {
	requestDef := GenReqDefForBindAppWarehouseBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindAppWarehouseBucketResponse), nil
	}
}

// BindAppWarehouseBucketInvoker 添加用户应用仓库桶及桶授权
func (c *WorkspaceAppClient) BindAppWarehouseBucketInvoker(request *model.BindAppWarehouseBucketRequest) *BindAppWarehouseBucketInvoker {
	requestDef := GenReqDefForBindAppWarehouseBucket()
	return &BindAppWarehouseBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBucketOrAcl 添加桶或者桶授权
//
// 添加桶或者桶授权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateBucketOrAcl(request *model.CreateBucketOrAclRequest) (*model.CreateBucketOrAclResponse, error) {
	requestDef := GenReqDefForCreateBucketOrAcl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBucketOrAclResponse), nil
	}
}

// CreateBucketOrAclInvoker 添加桶或者桶授权
func (c *WorkspaceAppClient) CreateBucketOrAclInvoker(request *model.CreateBucketOrAclRequest) *CreateBucketOrAclInvoker {
	requestDef := GenReqDefForCreateBucketOrAcl()
	return &CreateBucketOrAclInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWarehouseApp 在应用仓库中新增应用
//
// 在应用仓库中新增应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateWarehouseApp(request *model.CreateWarehouseAppRequest) (*model.CreateWarehouseAppResponse, error) {
	requestDef := GenReqDefForCreateWarehouseApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWarehouseAppResponse), nil
	}
}

// CreateWarehouseAppInvoker 在应用仓库中新增应用
func (c *WorkspaceAppClient) CreateWarehouseAppInvoker(request *model.CreateWarehouseAppRequest) *CreateWarehouseAppInvoker {
	requestDef := GenReqDefForCreateWarehouseApp()
	return &CreateWarehouseAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWarehouseApp 删除应用仓库中的指定应用
//
// 删除应用仓库中的指定应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteWarehouseApp(request *model.DeleteWarehouseAppRequest) (*model.DeleteWarehouseAppResponse, error) {
	requestDef := GenReqDefForDeleteWarehouseApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWarehouseAppResponse), nil
	}
}

// DeleteWarehouseAppInvoker 删除应用仓库中的指定应用
func (c *WorkspaceAppClient) DeleteWarehouseAppInvoker(request *model.DeleteWarehouseAppRequest) *DeleteWarehouseAppInvoker {
	requestDef := GenReqDefForDeleteWarehouseApp()
	return &DeleteWarehouseAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWarehouseApps 查询租户应用仓库中的应用列表
//
// 查询租户应用仓库中的应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListWarehouseApps(request *model.ListWarehouseAppsRequest) (*model.ListWarehouseAppsResponse, error) {
	requestDef := GenReqDefForListWarehouseApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWarehouseAppsResponse), nil
	}
}

// ListWarehouseAppsInvoker 查询租户应用仓库中的应用列表
func (c *WorkspaceAppClient) ListWarehouseAppsInvoker(request *model.ListWarehouseAppsRequest) *ListWarehouseAppsInvoker {
	requestDef := GenReqDefForListWarehouseApps()
	return &ListWarehouseAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppWarehouseBucket 查询用户应用仓库桶
//
// 查询用户应用仓库桶
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowAppWarehouseBucket(request *model.ShowAppWarehouseBucketRequest) (*model.ShowAppWarehouseBucketResponse, error) {
	requestDef := GenReqDefForShowAppWarehouseBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppWarehouseBucketResponse), nil
	}
}

// ShowAppWarehouseBucketInvoker 查询用户应用仓库桶
func (c *WorkspaceAppClient) ShowAppWarehouseBucketInvoker(request *model.ShowAppWarehouseBucketRequest) *ShowAppWarehouseBucketInvoker {
	requestDef := GenReqDefForShowAppWarehouseBucket()
	return &ShowAppWarehouseBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWarehouseApp 修改应用仓库中的指定应用信息
//
// 修改应用仓库中的指定应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateWarehouseApp(request *model.UpdateWarehouseAppRequest) (*model.UpdateWarehouseAppResponse, error) {
	requestDef := GenReqDefForUpdateWarehouseApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWarehouseAppResponse), nil
	}
}

// UpdateWarehouseAppInvoker 修改应用仓库中的指定应用信息
func (c *WorkspaceAppClient) UpdateWarehouseAppInvoker(request *model.UpdateWarehouseAppRequest) *UpdateWarehouseAppInvoker {
	requestDef := GenReqDefForUpdateWarehouseApp()
	return &UpdateWarehouseAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadWarehouseAppIcon 在应用仓库中上传图标文件
//
// 在应用仓库中上传图标文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UploadWarehouseAppIcon(request *model.UploadWarehouseAppIconRequest) (*model.UploadWarehouseAppIconResponse, error) {
	requestDef := GenReqDefForUploadWarehouseAppIcon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadWarehouseAppIconResponse), nil
	}
}

// UploadWarehouseAppIconInvoker 在应用仓库中上传图标文件
func (c *WorkspaceAppClient) UploadWarehouseAppIconInvoker(request *model.UploadWarehouseAppIconRequest) *UploadWarehouseAppIconInvoker {
	requestDef := GenReqDefForUploadWarehouseAppIcon()
	return &UploadWarehouseAppIconInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisableApp 批量禁用应用
//
// 批量禁用应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDisableApp(request *model.BatchDisableAppRequest) (*model.BatchDisableAppResponse, error) {
	requestDef := GenReqDefForBatchDisableApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableAppResponse), nil
	}
}

// BatchDisableAppInvoker 批量禁用应用
func (c *WorkspaceAppClient) BatchDisableAppInvoker(request *model.BatchDisableAppRequest) *BatchDisableAppInvoker {
	requestDef := GenReqDefForBatchDisableApp()
	return &BatchDisableAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableApp 批量启用应用
//
// 批量启用应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchEnableApp(request *model.BatchEnableAppRequest) (*model.BatchEnableAppResponse, error) {
	requestDef := GenReqDefForBatchEnableApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAppResponse), nil
	}
}

// BatchEnableAppInvoker 批量启用应用
func (c *WorkspaceAppClient) BatchEnableAppInvoker(request *model.BatchEnableAppRequest) *BatchEnableAppInvoker {
	requestDef := GenReqDefForBatchEnableApp()
	return &BatchEnableAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppIcon 删除自定义应用图标
//
// 删除自定义应用应用图标，恢复使用默认应用图标，重复执行会按照成功处理(响应200)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteAppIcon(request *model.DeleteAppIconRequest) (*model.DeleteAppIconResponse, error) {
	requestDef := GenReqDefForDeleteAppIcon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppIconResponse), nil
	}
}

// DeleteAppIconInvoker 删除自定义应用图标
func (c *WorkspaceAppClient) DeleteAppIconInvoker(request *model.DeleteAppIconRequest) *DeleteAppIconInvoker {
	requestDef := GenReqDefForDeleteAppIcon()
	return &DeleteAppIconInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublishedApp 查询已发布应用
//
// 查询已发布的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPublishedApp(request *model.ListPublishedAppRequest) (*model.ListPublishedAppResponse, error) {
	requestDef := GenReqDefForListPublishedApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublishedAppResponse), nil
	}
}

// ListPublishedAppInvoker 查询已发布应用
func (c *WorkspaceAppClient) ListPublishedAppInvoker(request *model.ListPublishedAppRequest) *ListPublishedAppInvoker {
	requestDef := GenReqDefForListPublishedApp()
	return &ListPublishedAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishApp 发布应用
//
// 批量发布应用，不允许发布同名的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) PublishApp(request *model.PublishAppRequest) (*model.PublishAppResponse, error) {
	requestDef := GenReqDefForPublishApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishAppResponse), nil
	}
}

// PublishAppInvoker 发布应用
func (c *WorkspaceAppClient) PublishAppInvoker(request *model.PublishAppRequest) *PublishAppInvoker {
	requestDef := GenReqDefForPublishApp()
	return &PublishAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppDetail 查询应用详细信息
//
// 查询应用详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowAppDetail(request *model.ShowAppDetailRequest) (*model.ShowAppDetailResponse, error) {
	requestDef := GenReqDefForShowAppDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppDetailResponse), nil
	}
}

// ShowAppDetailInvoker 查询应用详细信息
func (c *WorkspaceAppClient) ShowAppDetailInvoker(request *model.ShowAppDetailRequest) *ShowAppDetailInvoker {
	requestDef := GenReqDefForShowAppDetail()
	return &ShowAppDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublishableApp 可发布应用列表
//
// 查询应用组下可发布的应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowPublishableApp(request *model.ShowPublishableAppRequest) (*model.ShowPublishableAppResponse, error) {
	requestDef := GenReqDefForShowPublishableApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublishableAppResponse), nil
	}
}

// ShowPublishableAppInvoker 可发布应用列表
func (c *WorkspaceAppClient) ShowPublishableAppInvoker(request *model.ShowPublishableAppRequest) *ShowPublishableAppInvoker {
	requestDef := GenReqDefForShowPublishableApp()
	return &ShowPublishableAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnpublishApp 批量取消应用发布
//
// 批量取消应用发布。
// &gt; - 批量取消应用组下已经发布的应用，应用对应的授权会一起删除，重复执行会按照成功处理(响应200)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UnpublishApp(request *model.UnpublishAppRequest) (*model.UnpublishAppResponse, error) {
	requestDef := GenReqDefForUnpublishApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnpublishAppResponse), nil
	}
}

// UnpublishAppInvoker 批量取消应用发布
func (c *WorkspaceAppClient) UnpublishAppInvoker(request *model.UnpublishAppRequest) *UnpublishAppInvoker {
	requestDef := GenReqDefForUnpublishApp()
	return &UnpublishAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 修改应用信息
//
// 编辑修改应用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 修改应用信息
func (c *WorkspaceAppClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePreBootPolicy 批量设置应用预启动
//
// 批量设置应用预启动。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdatePreBootPolicy(request *model.UpdatePreBootPolicyRequest) (*model.UpdatePreBootPolicyResponse, error) {
	requestDef := GenReqDefForUpdatePreBootPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePreBootPolicyResponse), nil
	}
}

// UpdatePreBootPolicyInvoker 批量设置应用预启动
func (c *WorkspaceAppClient) UpdatePreBootPolicyInvoker(request *model.UpdatePreBootPolicyRequest) *UpdatePreBootPolicyInvoker {
	requestDef := GenReqDefForUpdatePreBootPolicy()
	return &UpdatePreBootPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAppIcon 修改自定义应用图标
//
// 修改自定义应用图标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UploadAppIcon(request *model.UploadAppIconRequest) (*model.UploadAppIconResponse, error) {
	requestDef := GenReqDefForUploadAppIcon()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAppIconResponse), nil
	}
}

// UploadAppIconInvoker 修改自定义应用图标
func (c *WorkspaceAppClient) UploadAppIconInvoker(request *model.UploadAppIconRequest) *UploadAppIconInvoker {
	requestDef := GenReqDefForUploadAppIcon()
	return &UploadAppIconInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InitializeTenant 租户服务激活、初始化
//
// 租户服务激活。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) InitializeTenant(request *model.InitializeTenantRequest) (*model.InitializeTenantResponse, error) {
	requestDef := GenReqDefForInitializeTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InitializeTenantResponse), nil
	}
}

// InitializeTenantInvoker 租户服务激活、初始化
func (c *WorkspaceAppClient) InitializeTenantInvoker(request *model.InitializeTenantRequest) *InitializeTenantInvoker {
	requestDef := GenReqDefForInitializeTenant()
	return &InitializeTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCorpConfigInfo 查询企业系统配置
//
// 配置加载顺序： 查询企业级配置--&gt; 查不到则赋默认阿波罗配置--&gt; 阿波罗没有则不返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListCorpConfigInfo(request *model.ListCorpConfigInfoRequest) (*model.ListCorpConfigInfoResponse, error) {
	requestDef := GenReqDefForListCorpConfigInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCorpConfigInfoResponse), nil
	}
}

// ListCorpConfigInfoInvoker 查询企业系统配置
func (c *WorkspaceAppClient) ListCorpConfigInfoInvoker(request *model.ListCorpConfigInfoRequest) *ListCorpConfigInfoInvoker {
	requestDef := GenReqDefForListCorpConfigInfo()
	return &ListCorpConfigInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantProfile 查询租户信息
//
// 查询租户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListTenantProfile(request *model.ListTenantProfileRequest) (*model.ListTenantProfileResponse, error) {
	requestDef := GenReqDefForListTenantProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantProfileResponse), nil
	}
}

// ListTenantProfileInvoker 查询租户信息
func (c *WorkspaceAppClient) ListTenantProfileInvoker(request *model.ListTenantProfileRequest) *ListTenantProfileInvoker {
	requestDef := GenReqDefForListTenantProfile()
	return &ListTenantProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAppGroup 批量删除应用组
//
// 批量删除应用组,重复执行会按照成功处理(响应200)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteAppGroup(request *model.BatchDeleteAppGroupRequest) (*model.BatchDeleteAppGroupResponse, error) {
	requestDef := GenReqDefForBatchDeleteAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppGroupResponse), nil
	}
}

// BatchDeleteAppGroupInvoker 批量删除应用组
func (c *WorkspaceAppClient) BatchDeleteAppGroupInvoker(request *model.BatchDeleteAppGroupRequest) *BatchDeleteAppGroupInvoker {
	requestDef := GenReqDefForBatchDeleteAppGroup()
	return &BatchDeleteAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppGroup 创建应用组
//
// 该API用于创建应用组。
// &gt; - 应用服务器中安装了不同的应用，这些应用可以组成不同的应用组，进行集中的管理和维护，向用户(组)发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateAppGroup(request *model.CreateAppGroupRequest) (*model.CreateAppGroupResponse, error) {
	requestDef := GenReqDefForCreateAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppGroupResponse), nil
	}
}

// CreateAppGroupInvoker 创建应用组
func (c *WorkspaceAppClient) CreateAppGroupInvoker(request *model.CreateAppGroupRequest) *CreateAppGroupInvoker {
	requestDef := GenReqDefForCreateAppGroup()
	return &CreateAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAppGroup 应用组删除
//
// 删除指定的应用组,重复执行会按照成功处理(响应200)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteAppGroup(request *model.DeleteAppGroupRequest) (*model.DeleteAppGroupResponse, error) {
	requestDef := GenReqDefForDeleteAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppGroupResponse), nil
	}
}

// DeleteAppGroupInvoker 应用组删除
func (c *WorkspaceAppClient) DeleteAppGroupInvoker(request *model.DeleteAppGroupRequest) *DeleteAppGroupInvoker {
	requestDef := GenReqDefForDeleteAppGroup()
	return &DeleteAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateAppGroup 解除服务组关联的所有应用组
//
// 解除服务组关联的所有应用组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DisassociateAppGroup(request *model.DisassociateAppGroupRequest) (*model.DisassociateAppGroupResponse, error) {
	requestDef := GenReqDefForDisassociateAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAppGroupResponse), nil
	}
}

// DisassociateAppGroupInvoker 解除服务组关联的所有应用组
func (c *WorkspaceAppClient) DisassociateAppGroupInvoker(request *model.DisassociateAppGroupRequest) *DisassociateAppGroupInvoker {
	requestDef := GenReqDefForDisassociateAppGroup()
	return &DisassociateAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppGroup 查询应用组
//
// 查询用户创建的应用组，按照名称、授权类型分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAppGroup(request *model.ListAppGroupRequest) (*model.ListAppGroupResponse, error) {
	requestDef := GenReqDefForListAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppGroupResponse), nil
	}
}

// ListAppGroupInvoker 查询应用组
func (c *WorkspaceAppClient) ListAppGroupInvoker(request *model.ListAppGroupRequest) *ListAppGroupInvoker {
	requestDef := GenReqDefForListAppGroup()
	return &ListAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppGroupDetail 查询应用组详情
//
// 查询应用组详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowAppGroupDetail(request *model.ShowAppGroupDetailRequest) (*model.ShowAppGroupDetailResponse, error) {
	requestDef := GenReqDefForShowAppGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppGroupDetailResponse), nil
	}
}

// ShowAppGroupDetailInvoker 查询应用组详情
func (c *WorkspaceAppClient) ShowAppGroupDetailInvoker(request *model.ShowAppGroupDetailRequest) *ShowAppGroupDetailInvoker {
	requestDef := GenReqDefForShowAppGroupDetail()
	return &ShowAppGroupDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAppGroup 修改应用组
//
// 修改应用组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateAppGroup(request *model.UpdateAppGroupRequest) (*model.UpdateAppGroupResponse, error) {
	requestDef := GenReqDefForUpdateAppGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppGroupResponse), nil
	}
}

// UpdateAppGroupInvoker 修改应用组
func (c *WorkspaceAppClient) UpdateAppGroupInvoker(request *model.UpdateAppGroupRequest) *UpdateAppGroupInvoker {
	requestDef := GenReqDefForUpdateAppGroup()
	return &UpdateAppGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrder 创建订单
//
// 创建订单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateOrder(request *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	requestDef := GenReqDefForCreateOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrderResponse), nil
	}
}

// CreateOrderInvoker 创建订单
func (c *WorkspaceAppClient) CreateOrderInvoker(request *model.CreateOrderRequest) *CreateOrderInvoker {
	requestDef := GenReqDefForCreateOrder()
	return &CreateOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProduct 查询云应用套餐
//
// 查询云应用套餐，按照条件过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListProduct(request *model.ListProductRequest) (*model.ListProductResponse, error) {
	requestDef := GenReqDefForListProduct()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductResponse), nil
	}
}

// ListProductInvoker 查询云应用套餐
func (c *WorkspaceAppClient) ListProductInvoker(request *model.ListProductRequest) *ListProductInvoker {
	requestDef := GenReqDefForListProduct()
	return &ListProductInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionType 查询会话套餐列表
//
// 该接口用于查询会话套餐列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListSessionType(request *model.ListSessionTypeRequest) (*model.ListSessionTypeResponse, error) {
	requestDef := GenReqDefForListSessionType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionTypeResponse), nil
	}
}

// ListSessionTypeInvoker 查询会话套餐列表
func (c *WorkspaceAppClient) ListSessionTypeInvoker(request *model.ListSessionTypeRequest) *ListSessionTypeInvoker {
	requestDef := GenReqDefForListSessionType()
	return &ListSessionTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSessionTypes 查询会话套餐列表（已废弃）
//
// 该接口用于查询会话套餐列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowSessionTypes(request *model.ShowSessionTypesRequest) (*model.ShowSessionTypesResponse, error) {
	requestDef := GenReqDefForShowSessionTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSessionTypesResponse), nil
	}
}

// ShowSessionTypesInvoker 查询会话套餐列表（已废弃）
func (c *WorkspaceAppClient) ShowSessionTypesInvoker(request *model.ShowSessionTypesRequest) *ShowSessionTypesInvoker {
	requestDef := GenReqDefForShowSessionTypes()
	return &ShowSessionTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAppGroupAuthorization 增加应用组授权
//
// 应用组添加用户授权，授权后用户就获得应用组下所有已发布应用的权限访问。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) AddAppGroupAuthorization(request *model.AddAppGroupAuthorizationRequest) (*model.AddAppGroupAuthorizationResponse, error) {
	requestDef := GenReqDefForAddAppGroupAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAppGroupAuthorizationResponse), nil
	}
}

// AddAppGroupAuthorizationInvoker 增加应用组授权
func (c *WorkspaceAppClient) AddAppGroupAuthorizationInvoker(request *model.AddAppGroupAuthorizationRequest) *AddAppGroupAuthorizationInvoker {
	requestDef := GenReqDefForAddAppGroupAuthorization()
	return &AddAppGroupAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAppGroupAuthorization 移除应用组授权
//
// 移除应用组内的指定用户的授权，用户授权删除后，用户将没有权限访问应用组内的任何应用。注意：重复执行会按照操作成功处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteAppGroupAuthorization(request *model.BatchDeleteAppGroupAuthorizationRequest) (*model.BatchDeleteAppGroupAuthorizationResponse, error) {
	requestDef := GenReqDefForBatchDeleteAppGroupAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppGroupAuthorizationResponse), nil
	}
}

// BatchDeleteAppGroupAuthorizationInvoker 移除应用组授权
func (c *WorkspaceAppClient) BatchDeleteAppGroupAuthorizationInvoker(request *model.BatchDeleteAppGroupAuthorizationRequest) *BatchDeleteAppGroupAuthorizationInvoker {
	requestDef := GenReqDefForBatchDeleteAppGroupAuthorization()
	return &BatchDeleteAppGroupAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppGroupAuthorization 查询应用组授权记录
//
// 查询应用内已授权的用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAppGroupAuthorization(request *model.ListAppGroupAuthorizationRequest) (*model.ListAppGroupAuthorizationResponse, error) {
	requestDef := GenReqDefForListAppGroupAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppGroupAuthorizationResponse), nil
	}
}

// ListAppGroupAuthorizationInvoker 查询应用组授权记录
func (c *WorkspaceAppClient) ListAppGroupAuthorizationInvoker(request *model.ListAppGroupAuthorizationRequest) *ListAppGroupAuthorizationInvoker {
	requestDef := GenReqDefForListAppGroupAuthorization()
	return &ListAppGroupAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZone 查询可用分区列表
//
// 该接口用于查询支持的可用分区列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAvailabilityZone(request *model.ListAvailabilityZoneRequest) (*model.ListAvailabilityZoneResponse, error) {
	requestDef := GenReqDefForListAvailabilityZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZoneResponse), nil
	}
}

// ListAvailabilityZoneInvoker 查询可用分区列表
func (c *WorkspaceAppClient) ListAvailabilityZoneInvoker(request *model.ListAvailabilityZoneRequest) *ListAvailabilityZoneInvoker {
	requestDef := GenReqDefForListAvailabilityZone()
	return &ListAvailabilityZoneInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAz 查询可用分区列表（按站点分类）
//
// 该接口用于查询支持的可用分区列表，按站点分类。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAz(request *model.ListAzRequest) (*model.ListAzResponse, error) {
	requestDef := GenReqDefForListAz()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAzResponse), nil
	}
}

// ListAzInvoker 查询可用分区列表（按站点分类）
func (c *WorkspaceAppClient) ListAzInvoker(request *model.ListAzRequest) *ListAzInvoker {
	requestDef := GenReqDefForListAz()
	return &ListAzInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteCloudStorage 批量删除云存储
//
// 批量删除云存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteCloudStorage(request *model.BatchDeleteCloudStorageRequest) (*model.BatchDeleteCloudStorageResponse, error) {
	requestDef := GenReqDefForBatchDeleteCloudStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteCloudStorageResponse), nil
	}
}

// BatchDeleteCloudStorageInvoker 批量删除云存储
func (c *WorkspaceAppClient) BatchDeleteCloudStorageInvoker(request *model.BatchDeleteCloudStorageRequest) *BatchDeleteCloudStorageInvoker {
	requestDef := GenReqDefForBatchDeleteCloudStorage()
	return &BatchDeleteCloudStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCloudStorage 创建项目配置关联
//
// 创建项目配置关联，目前仅支持关联项目配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateCloudStorage(request *model.CreateCloudStorageRequest) (*model.CreateCloudStorageResponse, error) {
	requestDef := GenReqDefForCreateCloudStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudStorageResponse), nil
	}
}

// CreateCloudStorageInvoker 创建项目配置关联
func (c *WorkspaceAppClient) CreateCloudStorageInvoker(request *model.CreateCloudStorageRequest) *CreateCloudStorageInvoker {
	requestDef := GenReqDefForCreateCloudStorage()
	return &CreateCloudStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserFolderAssignment 创建个人文件夹
//
// 创建个人文件夹，已存在对应目录时，仅更新策略不会重复创建目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateUserFolderAssignment(request *model.CreateUserFolderAssignmentRequest) (*model.CreateUserFolderAssignmentResponse, error) {
	requestDef := GenReqDefForCreateUserFolderAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserFolderAssignmentResponse), nil
	}
}

// CreateUserFolderAssignmentInvoker 创建个人文件夹
func (c *WorkspaceAppClient) CreateUserFolderAssignmentInvoker(request *model.CreateUserFolderAssignmentRequest) *CreateUserFolderAssignmentInvoker {
	requestDef := GenReqDefForCreateUserFolderAssignment()
	return &CreateUserFolderAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCloudStorage 删除云存储
//
// 删除共享存储，只会解除NAS与项目配置之间的关联关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteCloudStorage(request *model.DeleteCloudStorageRequest) (*model.DeleteCloudStorageResponse, error) {
	requestDef := GenReqDefForDeleteCloudStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudStorageResponse), nil
	}
}

// DeleteCloudStorageInvoker 删除云存储
func (c *WorkspaceAppClient) DeleteCloudStorageInvoker(request *model.DeleteCloudStorageRequest) *DeleteCloudStorageInvoker {
	requestDef := GenReqDefForDeleteCloudStorage()
	return &DeleteCloudStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCloudStorageAttachment 删除个人文件夹
//
// 删除个人存储目录，个人目录中的数据也将永久删除且无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteCloudStorageAttachment(request *model.DeleteCloudStorageAttachmentRequest) (*model.DeleteCloudStorageAttachmentResponse, error) {
	requestDef := GenReqDefForDeleteCloudStorageAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudStorageAttachmentResponse), nil
	}
}

// DeleteCloudStorageAttachmentInvoker 删除个人文件夹
func (c *WorkspaceAppClient) DeleteCloudStorageAttachmentInvoker(request *model.DeleteCloudStorageAttachmentRequest) *DeleteCloudStorageAttachmentInvoker {
	requestDef := GenReqDefForDeleteCloudStorageAttachment()
	return &DeleteCloudStorageAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudStorage 查询云存储
//
// 查询云存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListCloudStorage(request *model.ListCloudStorageRequest) (*model.ListCloudStorageResponse, error) {
	requestDef := GenReqDefForListCloudStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudStorageResponse), nil
	}
}

// ListCloudStorageInvoker 查询云存储
func (c *WorkspaceAppClient) ListCloudStorageInvoker(request *model.ListCloudStorageRequest) *ListCloudStorageInvoker {
	requestDef := GenReqDefForListCloudStorage()
	return &ListCloudStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudStorageAssignment 查询个人文件夹列表
//
// 查询个人文件夹列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListCloudStorageAssignment(request *model.ListCloudStorageAssignmentRequest) (*model.ListCloudStorageAssignmentResponse, error) {
	requestDef := GenReqDefForListCloudStorageAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudStorageAssignmentResponse), nil
	}
}

// ListCloudStorageAssignmentInvoker 查询个人文件夹列表
func (c *WorkspaceAppClient) ListCloudStorageAssignmentInvoker(request *model.ListCloudStorageAssignmentRequest) *ListCloudStorageAssignmentInvoker {
	requestDef := GenReqDefForListCloudStorageAssignment()
	return &ListCloudStorageAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectConfigs 查询项目配置列表
//
// 查询项目配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListProjectConfigs(request *model.ListProjectConfigsRequest) (*model.ListProjectConfigsResponse, error) {
	requestDef := GenReqDefForListProjectConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectConfigsResponse), nil
	}
}

// ListProjectConfigsInvoker 查询项目配置列表
func (c *WorkspaceAppClient) ListProjectConfigsInvoker(request *model.ListProjectConfigsRequest) *ListProjectConfigsInvoker {
	requestDef := GenReqDefForListProjectConfigs()
	return &ListProjectConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectConfig 查询项目配置信息
//
// 查询项目配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowProjectConfig(request *model.ShowProjectConfigRequest) (*model.ShowProjectConfigResponse, error) {
	requestDef := GenReqDefForShowProjectConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectConfigResponse), nil
	}
}

// ShowProjectConfigInvoker 查询项目配置信息
func (c *WorkspaceAppClient) ShowProjectConfigInvoker(request *model.ShowProjectConfigRequest) *ShowProjectConfigInvoker {
	requestDef := GenReqDefForShowProjectConfig()
	return &ShowProjectConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCloudUserFolderAssignment 修改个人文件夹
//
// 创建个人文件夹。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateCloudUserFolderAssignment(request *model.UpdateCloudUserFolderAssignmentRequest) (*model.UpdateCloudUserFolderAssignmentResponse, error) {
	requestDef := GenReqDefForUpdateCloudUserFolderAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCloudUserFolderAssignmentResponse), nil
	}
}

// UpdateCloudUserFolderAssignmentInvoker 修改个人文件夹
func (c *WorkspaceAppClient) UpdateCloudUserFolderAssignmentInvoker(request *model.UpdateCloudUserFolderAssignmentRequest) *UpdateCloudUserFolderAssignmentInvoker {
	requestDef := GenReqDefForUpdateCloudUserFolderAssignment()
	return &UpdateCloudUserFolderAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachImageServerApp 分发软件信息至镜像实例
//
// 分发应用软件信息至镜像实例，管理员可以按需下载并安装应用软件。
// * 目前只支持来自云应用仓库的软件信息。
// * 只允许对状态为 &#x60;实例正常运行&#x60;、&#x60;镜像任务结束&#x60; 的实例分发软件信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) AttachImageServerApp(request *model.AttachImageServerAppRequest) (*model.AttachImageServerAppResponse, error) {
	requestDef := GenReqDefForAttachImageServerApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachImageServerAppResponse), nil
	}
}

// AttachImageServerAppInvoker 分发软件信息至镜像实例
func (c *WorkspaceAppClient) AttachImageServerAppInvoker(request *model.AttachImageServerAppRequest) *AttachImageServerAppInvoker {
	requestDef := GenReqDefForAttachImageServerApp()
	return &AttachImageServerAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteImageServer 批量删除镜像实例
//
// 批量删除镜像实例。
// * 忽略不存在的镜像实例，响应正常。
// * 不允许操作状态为 &#x60;创建中&#x60;、&#x60;镜像创建中&#x60;的实例，响应异常。
// * 不支持资源关联发生变化后，请求删除镜像实例关联资源，任务响应异常。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteImageServer(request *model.BatchDeleteImageServerRequest) (*model.BatchDeleteImageServerResponse, error) {
	requestDef := GenReqDefForBatchDeleteImageServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteImageServerResponse), nil
	}
}

// BatchDeleteImageServerInvoker 批量删除镜像实例
func (c *WorkspaceAppClient) BatchDeleteImageServerInvoker(request *model.BatchDeleteImageServerRequest) *BatchDeleteImageServerInvoker {
	requestDef := GenReqDefForBatchDeleteImageServer()
	return &BatchDeleteImageServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImageServer 创建镜像实例
//
// 创建镜像实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateImageServer(request *model.CreateImageServerRequest) (*model.CreateImageServerResponse, error) {
	requestDef := GenReqDefForCreateImageServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageServerResponse), nil
	}
}

// CreateImageServerInvoker 创建镜像实例
func (c *WorkspaceAppClient) CreateImageServerInvoker(request *model.CreateImageServerRequest) *CreateImageServerInvoker {
	requestDef := GenReqDefForCreateImageServer()
	return &CreateImageServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageServers 查询镜像实例列表
//
// 查询镜像实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListImageServers(request *model.ListImageServersRequest) (*model.ListImageServersResponse, error) {
	requestDef := GenReqDefForListImageServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageServersResponse), nil
	}
}

// ListImageServersInvoker 查询镜像实例列表
func (c *WorkspaceAppClient) ListImageServersInvoker(request *model.ListImageServersRequest) *ListImageServersInvoker {
	requestDef := GenReqDefForListImageServers()
	return &ListImageServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLatestAttachedServerApp 查询最近一次分发软件信息列表
//
// 查询最近一次分发软件信息列表，返回ID列表，不包含具体信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListLatestAttachedServerApp(request *model.ListLatestAttachedServerAppRequest) (*model.ListLatestAttachedServerAppResponse, error) {
	requestDef := GenReqDefForListLatestAttachedServerApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLatestAttachedServerAppResponse), nil
	}
}

// ListLatestAttachedServerAppInvoker 查询最近一次分发软件信息列表
func (c *WorkspaceAppClient) ListLatestAttachedServerAppInvoker(request *model.ListLatestAttachedServerAppRequest) *ListLatestAttachedServerAppInvoker {
	requestDef := GenReqDefForListLatestAttachedServerApp()
	return &ListLatestAttachedServerAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RecreateServerImage 构建云应用镜像
//
// 构建云应用镜像。
// * 只允许对状态为 &#x60;实例正常运行&#x60;、&#x60;镜像任务结束&#x60; 的实例构建云应用镜像。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) RecreateServerImage(request *model.RecreateServerImageRequest) (*model.RecreateServerImageResponse, error) {
	requestDef := GenReqDefForRecreateServerImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RecreateServerImageResponse), nil
	}
}

// RecreateServerImageInvoker 构建云应用镜像
func (c *WorkspaceAppClient) RecreateServerImageInvoker(request *model.RecreateServerImageRequest) *RecreateServerImageInvoker {
	requestDef := GenReqDefForRecreateServerImage()
	return &RecreateServerImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageServer 查询指定镜像实例
//
// 查询指定的镜像实例当前这个接口的查询数据和list查询的一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowImageServer(request *model.ShowImageServerRequest) (*model.ShowImageServerResponse, error) {
	requestDef := GenReqDefForShowImageServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageServerResponse), nil
	}
}

// ShowImageServerInvoker 查询指定镜像实例
func (c *WorkspaceAppClient) ShowImageServerInvoker(request *model.ShowImageServerRequest) *ShowImageServerInvoker {
	requestDef := GenReqDefForShowImageServer()
	return &ShowImageServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImageServer 修改镜像实例
//
// 修改镜像实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateImageServer(request *model.UpdateImageServerRequest) (*model.UpdateImageServerResponse, error) {
	requestDef := GenReqDefForUpdateImageServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageServerResponse), nil
	}
}

// UpdateImageServerInvoker 修改镜像实例
func (c *WorkspaceAppClient) UpdateImageServerInvoker(request *model.UpdateImageServerRequest) *UpdateImageServerInvoker {
	requestDef := GenReqDefForUpdateImageServer()
	return &UpdateImageServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAppSubJobs 批量删除子任务
//
// 批量删除子任务，忽略不存在的服务器并且返回成功响应。
// 只能删除以下的两种状态：
// - SUCCESS：成功。
// - FAILED：失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteAppSubJobs(request *model.BatchDeleteAppSubJobsRequest) (*model.BatchDeleteAppSubJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteAppSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAppSubJobsResponse), nil
	}
}

// BatchDeleteAppSubJobsInvoker 批量删除子任务
func (c *WorkspaceAppClient) BatchDeleteAppSubJobsInvoker(request *model.BatchDeleteAppSubJobsRequest) *BatchDeleteAppSubJobsInvoker {
	requestDef := GenReqDefForBatchDeleteAppSubJobs()
	return &BatchDeleteAppSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteImageSubJobs 批量删除镜像子任务
//
// 批量删除子任务，忽略不存在的服务器并且返回成功响应。
// 只能删除以下的两种状态 SUCCESS：成功。 FAILED：失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteImageSubJobs(request *model.BatchDeleteImageSubJobsRequest) (*model.BatchDeleteImageSubJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteImageSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteImageSubJobsResponse), nil
	}
}

// BatchDeleteImageSubJobsInvoker 批量删除镜像子任务
func (c *WorkspaceAppClient) BatchDeleteImageSubJobsInvoker(request *model.BatchDeleteImageSubJobsRequest) *BatchDeleteImageSubJobsInvoker {
	requestDef := GenReqDefForBatchDeleteImageSubJobs()
	return &BatchDeleteImageSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountSubJobs 子任务数量查询
//
// 该接口用于查询异步子任务数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CountSubJobs(request *model.CountSubJobsRequest) (*model.CountSubJobsResponse, error) {
	requestDef := GenReqDefForCountSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountSubJobsResponse), nil
	}
}

// CountSubJobsInvoker 子任务数量查询
func (c *WorkspaceAppClient) CountSubJobsInvoker(request *model.CountSubJobsRequest) *CountSubJobsInvoker {
	requestDef := GenReqDefForCountSubJobs()
	return &CountSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageJobs 查询租户的镜像任务列表
//
// 该接口用于查询租户的异步任务执行情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListImageJobs(request *model.ListImageJobsRequest) (*model.ListImageJobsResponse, error) {
	requestDef := GenReqDefForListImageJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageJobsResponse), nil
	}
}

// ListImageJobsInvoker 查询租户的镜像任务列表
func (c *WorkspaceAppClient) ListImageJobsInvoker(request *model.ListImageJobsRequest) *ListImageJobsInvoker {
	requestDef := GenReqDefForListImageJobs()
	return &ListImageJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageSubJobs 镜像子任务查询
//
// 该接口用于查询异步子任务执行情况,sub_job_ids非空时offset和limit不会生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListImageSubJobs(request *model.ListImageSubJobsRequest) (*model.ListImageSubJobsResponse, error) {
	requestDef := GenReqDefForListImageSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageSubJobsResponse), nil
	}
}

// ListImageSubJobsInvoker 镜像子任务查询
func (c *WorkspaceAppClient) ListImageSubJobsInvoker(request *model.ListImageSubJobsRequest) *ListImageSubJobsInvoker {
	requestDef := GenReqDefForListImageSubJobs()
	return &ListImageSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubJobs 子任务查询
//
// 该接口用于查询异步子任务执行情况,sub_job_ids非空时offset和limit不会生效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListSubJobs(request *model.ListSubJobsRequest) (*model.ListSubJobsResponse, error) {
	requestDef := GenReqDefForListSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubJobsResponse), nil
	}
}

// ListSubJobsInvoker 子任务查询
func (c *WorkspaceAppClient) ListSubJobsInvoker(request *model.ListSubJobsRequest) *ListSubJobsInvoker {
	requestDef := GenReqDefForListSubJobs()
	return &ListSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImageJob 查询镜像任务详情
//
// 该接口用于查询异步任务的执行情况，比如查询创建镜像实例任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowImageJob(request *model.ShowImageJobRequest) (*model.ShowImageJobResponse, error) {
	requestDef := GenReqDefForShowImageJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageJobResponse), nil
	}
}

// ShowImageJobInvoker 查询镜像任务详情
func (c *WorkspaceAppClient) ShowImageJobInvoker(request *model.ShowImageJobRequest) *ShowImageJobInvoker {
	requestDef := GenReqDefForShowImageJob()
	return &ShowImageJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询任务的执行状态（已废弃）
//
// 查询Job的执行状态。
//
// 对于创建云应用服务器命令下发后会返回job_id，通过job_id可以查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询任务的执行状态（已废弃）
func (c *WorkspaceAppClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 查询任务的执行状态详情
//
// 查询Job的执行状态。
//
// 对于创建云服务器、删除云服务器、重建镜像等异步API，下发命令后会返回job_id，通过job_id可以查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 查询任务的执行状态详情
func (c *WorkspaceAppClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizationMailRecord 查询应用组授权邮件发送记录
//
// 查询应用组授权邮件发送记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAuthorizationMailRecord(request *model.ListAuthorizationMailRecordRequest) (*model.ListAuthorizationMailRecordResponse, error) {
	requestDef := GenReqDefForListAuthorizationMailRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizationMailRecordResponse), nil
	}
}

// ListAuthorizationMailRecordInvoker 查询应用组授权邮件发送记录
func (c *WorkspaceAppClient) ListAuthorizationMailRecordInvoker(request *model.ListAuthorizationMailRecordRequest) *ListAuthorizationMailRecordInvoker {
	requestDef := GenReqDefForListAuthorizationMailRecord()
	return &ListAuthorizationMailRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendAuthorizationMail 重发应用组授权邮件（根据授权邮件记录）
//
// 重发应用组授权邮件（根据授权邮件记录）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) SendAuthorizationMail(request *model.SendAuthorizationMailRequest) (*model.SendAuthorizationMailResponse, error) {
	requestDef := GenReqDefForSendAuthorizationMail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendAuthorizationMailResponse), nil
	}
}

// SendAuthorizationMailInvoker 重发应用组授权邮件（根据授权邮件记录）
func (c *WorkspaceAppClient) SendAuthorizationMailInvoker(request *model.SendAuthorizationMailRequest) *SendAuthorizationMailInvoker {
	requestDef := GenReqDefForSendAuthorizationMail()
	return &SendAuthorizationMailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SendAuthorizedMail 重发应用组授权邮件（根据授权记录）
//
// 重发应用组授权邮件（根据授权记录）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) SendAuthorizedMail(request *model.SendAuthorizedMailRequest) (*model.SendAuthorizedMailResponse, error) {
	requestDef := GenReqDefForSendAuthorizedMail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SendAuthorizedMailResponse), nil
	}
}

// SendAuthorizedMailInvoker 重发应用组授权邮件（根据授权记录）
func (c *WorkspaceAppClient) SendAuthorizedMailInvoker(request *model.SendAuthorizedMailRequest) *SendAuthorizedMailInvoker {
	requestDef := GenReqDefForSendAuthorizedMail()
	return &SendAuthorizedMailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePersistentStorage 批量删除WKS存储
//
// 批量删除WKS存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeletePersistentStorage(request *model.BatchDeletePersistentStorageRequest) (*model.BatchDeletePersistentStorageResponse, error) {
	requestDef := GenReqDefForBatchDeletePersistentStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePersistentStorageResponse), nil
	}
}

// BatchDeletePersistentStorageInvoker 批量删除WKS存储
func (c *WorkspaceAppClient) BatchDeletePersistentStorageInvoker(request *model.BatchDeletePersistentStorageRequest) *BatchDeletePersistentStorageInvoker {
	requestDef := GenReqDefForBatchDeletePersistentStorage()
	return &BatchDeletePersistentStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateStoragePolicyStatement 新增或更新存储目录访问权限自定义策略
//
// 新增或更新存储目录访问权限自定义策略(已存在自定义策略时会对已有策略更新)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateOrUpdateStoragePolicyStatement(request *model.CreateOrUpdateStoragePolicyStatementRequest) (*model.CreateOrUpdateStoragePolicyStatementResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateStoragePolicyStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateStoragePolicyStatementResponse), nil
	}
}

// CreateOrUpdateStoragePolicyStatementInvoker 新增或更新存储目录访问权限自定义策略
func (c *WorkspaceAppClient) CreateOrUpdateStoragePolicyStatementInvoker(request *model.CreateOrUpdateStoragePolicyStatementRequest) *CreateOrUpdateStoragePolicyStatementInvoker {
	requestDef := GenReqDefForCreateOrUpdateStoragePolicyStatement()
	return &CreateOrUpdateStoragePolicyStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePersistentStorage 创建WKS存储
//
// 创建WKS存储，目前仅支持创建 SFS3.0 容量型文件系统。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreatePersistentStorage(request *model.CreatePersistentStorageRequest) (*model.CreatePersistentStorageResponse, error) {
	requestDef := GenReqDefForCreatePersistentStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePersistentStorageResponse), nil
	}
}

// CreatePersistentStorageInvoker 创建WKS存储
func (c *WorkspaceAppClient) CreatePersistentStorageInvoker(request *model.CreatePersistentStorageRequest) *CreatePersistentStorageInvoker {
	requestDef := GenReqDefForCreatePersistentStorage()
	return &CreatePersistentStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateShareFolder 创建共享存储目录
//
// 创建共享存储目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateShareFolder(request *model.CreateShareFolderRequest) (*model.CreateShareFolderResponse, error) {
	requestDef := GenReqDefForCreateShareFolder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateShareFolderResponse), nil
	}
}

// CreateShareFolderInvoker 创建共享存储目录
func (c *WorkspaceAppClient) CreateShareFolderInvoker(request *model.CreateShareFolderRequest) *CreateShareFolderInvoker {
	requestDef := GenReqDefForCreateShareFolder()
	return &CreateShareFolderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePersistentStorage 删除WKS存储
//
// 删除共享存储，只会解除NAS与文件系统之间的关联关系，不会删除文件系统和文件系统中的数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeletePersistentStorage(request *model.DeletePersistentStorageRequest) (*model.DeletePersistentStorageResponse, error) {
	requestDef := GenReqDefForDeletePersistentStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePersistentStorageResponse), nil
	}
}

// DeletePersistentStorageInvoker 删除WKS存储
func (c *WorkspaceAppClient) DeletePersistentStorageInvoker(request *model.DeletePersistentStorageRequest) *DeletePersistentStorageInvoker {
	requestDef := GenReqDefForDeletePersistentStorage()
	return &DeletePersistentStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStorageClaim 删除共享目录
//
// 删除共享存储目录。
// &gt; 需要删除绑定的用户及用户组，才能删除共享文目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteStorageClaim(request *model.DeleteStorageClaimRequest) (*model.DeleteStorageClaimResponse, error) {
	requestDef := GenReqDefForDeleteStorageClaim()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStorageClaimResponse), nil
	}
}

// DeleteStorageClaimInvoker 删除共享目录
func (c *WorkspaceAppClient) DeleteStorageClaimInvoker(request *model.DeleteStorageClaimRequest) *DeleteStorageClaimInvoker {
	requestDef := GenReqDefForDeleteStorageClaim()
	return &DeleteStorageClaimInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserStorageAttachment 删除个人存储目录
//
// 删除个人存储目录，个人目录中的数据也将永久删除且无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteUserStorageAttachment(request *model.DeleteUserStorageAttachmentRequest) (*model.DeleteUserStorageAttachmentResponse, error) {
	requestDef := GenReqDefForDeleteUserStorageAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserStorageAttachmentResponse), nil
	}
}

// DeleteUserStorageAttachmentInvoker 删除个人存储目录
func (c *WorkspaceAppClient) DeleteUserStorageAttachmentInvoker(request *model.DeleteUserStorageAttachmentRequest) *DeleteUserStorageAttachmentInvoker {
	requestDef := GenReqDefForDeleteUserStorageAttachment()
	return &DeleteUserStorageAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPersistentStorage 查询WKS存储
//
// 查询WKS存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPersistentStorage(request *model.ListPersistentStorageRequest) (*model.ListPersistentStorageResponse, error) {
	requestDef := GenReqDefForListPersistentStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPersistentStorageResponse), nil
	}
}

// ListPersistentStorageInvoker 查询WKS存储
func (c *WorkspaceAppClient) ListPersistentStorageInvoker(request *model.ListPersistentStorageRequest) *ListPersistentStorageInvoker {
	requestDef := GenReqDefForListPersistentStorage()
	return &ListPersistentStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSfs3Storage 查询SFS3.0存储
//
// 查询SFS3.0存储。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListSfs3Storage(request *model.ListSfs3StorageRequest) (*model.ListSfs3StorageResponse, error) {
	requestDef := GenReqDefForListSfs3Storage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSfs3StorageResponse), nil
	}
}

// ListSfs3StorageInvoker 查询SFS3.0存储
func (c *WorkspaceAppClient) ListSfs3StorageInvoker(request *model.ListSfs3StorageRequest) *ListSfs3StorageInvoker {
	requestDef := GenReqDefForListSfs3Storage()
	return &ListSfs3StorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListShareFolder 查询共享存储目录
//
// 查询共享存储目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListShareFolder(request *model.ListShareFolderRequest) (*model.ListShareFolderResponse, error) {
	requestDef := GenReqDefForListShareFolder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListShareFolderResponse), nil
	}
}

// ListShareFolderInvoker 查询共享存储目录
func (c *WorkspaceAppClient) ListShareFolderInvoker(request *model.ListShareFolderRequest) *ListShareFolderInvoker {
	requestDef := GenReqDefForListShareFolder()
	return &ListShareFolderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageAssignment 查询个人存储目录
//
// 查询个人存储目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListStorageAssignment(request *model.ListStorageAssignmentRequest) (*model.ListStorageAssignmentResponse, error) {
	requestDef := GenReqDefForListStorageAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageAssignmentResponse), nil
	}
}

// ListStorageAssignmentInvoker 查询个人存储目录
func (c *WorkspaceAppClient) ListStorageAssignmentInvoker(request *model.ListStorageAssignmentRequest) *ListStorageAssignmentInvoker {
	requestDef := GenReqDefForListStorageAssignment()
	return &ListStorageAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStoragePolicyStatement 查询存储目录访问权限策略
//
// 查询存储目录访问权限策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListStoragePolicyStatement(request *model.ListStoragePolicyStatementRequest) (*model.ListStoragePolicyStatementResponse, error) {
	requestDef := GenReqDefForListStoragePolicyStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStoragePolicyStatementResponse), nil
	}
}

// ListStoragePolicyStatementInvoker 查询存储目录访问权限策略
func (c *WorkspaceAppClient) ListStoragePolicyStatementInvoker(request *model.ListStoragePolicyStatementRequest) *ListStoragePolicyStatementInvoker {
	requestDef := GenReqDefForListStoragePolicyStatement()
	return &ListStoragePolicyStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateShareFolderAssignment 修改共享目录成员
//
// 批量添加或者移除共享目录成员。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateShareFolderAssignment(request *model.UpdateShareFolderAssignmentRequest) (*model.UpdateShareFolderAssignmentResponse, error) {
	requestDef := GenReqDefForUpdateShareFolderAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateShareFolderAssignmentResponse), nil
	}
}

// UpdateShareFolderAssignmentInvoker 修改共享目录成员
func (c *WorkspaceAppClient) UpdateShareFolderAssignmentInvoker(request *model.UpdateShareFolderAssignmentRequest) *UpdateShareFolderAssignmentInvoker {
	requestDef := GenReqDefForUpdateShareFolderAssignment()
	return &UpdateShareFolderAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserFolderAssignment 创建个人存储目录
//
// 创建个人存储目录，已存在对应目录时，仅更新策略不会重复创建目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateUserFolderAssignment(request *model.UpdateUserFolderAssignmentRequest) (*model.UpdateUserFolderAssignmentResponse, error) {
	requestDef := GenReqDefForUpdateUserFolderAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserFolderAssignmentResponse), nil
	}
}

// UpdateUserFolderAssignmentInvoker 创建个人存储目录
func (c *WorkspaceAppClient) UpdateUserFolderAssignmentInvoker(request *model.UpdateUserFolderAssignmentRequest) *UpdateUserFolderAssignmentInvoker {
	requestDef := GenReqDefForUpdateUserFolderAssignment()
	return &UpdateUserFolderAssignmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyGroup 新增策略组
//
// 新增策略组，通过策略组能灵活地控制客户端访问与接入策略，如：文件、剪切板、会话等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreatePolicyGroup(request *model.CreatePolicyGroupRequest) (*model.CreatePolicyGroupResponse, error) {
	requestDef := GenReqDefForCreatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyGroupResponse), nil
	}
}

// CreatePolicyGroupInvoker 新增策略组
func (c *WorkspaceAppClient) CreatePolicyGroupInvoker(request *model.CreatePolicyGroupRequest) *CreatePolicyGroupInvoker {
	requestDef := GenReqDefForCreatePolicyGroup()
	return &CreatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePolicyTemplate 新增策略模板
//
// 新增策略模板。策略模板创建好后，用户在创建策略组的时候，就可以根据已有策略模板按需调整配置，快速完成策略组的创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreatePolicyTemplate(request *model.CreatePolicyTemplateRequest) (*model.CreatePolicyTemplateResponse, error) {
	requestDef := GenReqDefForCreatePolicyTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyTemplateResponse), nil
	}
}

// CreatePolicyTemplateInvoker 新增策略模板
func (c *WorkspaceAppClient) CreatePolicyTemplateInvoker(request *model.CreatePolicyTemplateRequest) *CreatePolicyTemplateInvoker {
	requestDef := GenReqDefForCreatePolicyTemplate()
	return &CreatePolicyTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyGroup 删除策略组
//
// 删除指定策略组，包含策略组对应的策略信息以及应用对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeletePolicyGroup(request *model.DeletePolicyGroupRequest) (*model.DeletePolicyGroupResponse, error) {
	requestDef := GenReqDefForDeletePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyGroupResponse), nil
	}
}

// DeletePolicyGroupInvoker 删除策略组
func (c *WorkspaceAppClient) DeletePolicyGroupInvoker(request *model.DeletePolicyGroupRequest) *DeletePolicyGroupInvoker {
	requestDef := GenReqDefForDeletePolicyGroup()
	return &DeletePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePolicyTemplate 删除策略模板
//
// 删除指定策略模板，包含策略模板对应的策略信息以及应用对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeletePolicyTemplate(request *model.DeletePolicyTemplateRequest) (*model.DeletePolicyTemplateResponse, error) {
	requestDef := GenReqDefForDeletePolicyTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyTemplateResponse), nil
	}
}

// DeletePolicyTemplateInvoker 删除策略模板
func (c *WorkspaceAppClient) DeletePolicyTemplateInvoker(request *model.DeletePolicyTemplateRequest) *DeletePolicyTemplateInvoker {
	requestDef := GenReqDefForDeletePolicyTemplate()
	return &DeletePolicyTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyGroup 查询策略组列表
//
// 查询策略组概要信息列表,包括应用对象和详细策略项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPolicyGroup(request *model.ListPolicyGroupRequest) (*model.ListPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupResponse), nil
	}
}

// ListPolicyGroupInvoker 查询策略组列表
func (c *WorkspaceAppClient) ListPolicyGroupInvoker(request *model.ListPolicyGroupRequest) *ListPolicyGroupInvoker {
	requestDef := GenReqDefForListPolicyGroup()
	return &ListPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyGroupDetailInfo 查询策略组详情列表
//
// 包含策略信息以及应用对象的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPolicyGroupDetailInfo(request *model.ListPolicyGroupDetailInfoRequest) (*model.ListPolicyGroupDetailInfoResponse, error) {
	requestDef := GenReqDefForListPolicyGroupDetailInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyGroupDetailInfoResponse), nil
	}
}

// ListPolicyGroupDetailInfoInvoker 查询策略组详情列表
func (c *WorkspaceAppClient) ListPolicyGroupDetailInfoInvoker(request *model.ListPolicyGroupDetailInfoRequest) *ListPolicyGroupDetailInfoInvoker {
	requestDef := GenReqDefForListPolicyGroupDetailInfo()
	return &ListPolicyGroupDetailInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyOfPolicyGroup 查询策略组中的策略项
//
// 查询指定策略组内的策略项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPolicyOfPolicyGroup(request *model.ListPolicyOfPolicyGroupRequest) (*model.ListPolicyOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForListPolicyOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyOfPolicyGroupResponse), nil
	}
}

// ListPolicyOfPolicyGroupInvoker 查询策略组中的策略项
func (c *WorkspaceAppClient) ListPolicyOfPolicyGroupInvoker(request *model.ListPolicyOfPolicyGroupRequest) *ListPolicyOfPolicyGroupInvoker {
	requestDef := GenReqDefForListPolicyOfPolicyGroup()
	return &ListPolicyOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicyTemplate 查询策略模板列表
//
// 查询策略模板概要信息列表，包含策略信息以及应用对象信息。用户在创建策略组的时候，可以根据已有策略模板按需调整配置，快速完成策略组的创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListPolicyTemplate(request *model.ListPolicyTemplateRequest) (*model.ListPolicyTemplateResponse, error) {
	requestDef := GenReqDefForListPolicyTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyTemplateResponse), nil
	}
}

// ListPolicyTemplateInvoker 查询策略模板列表
func (c *WorkspaceAppClient) ListPolicyTemplateInvoker(request *model.ListPolicyTemplateRequest) *ListPolicyTemplateInvoker {
	requestDef := GenReqDefForListPolicyTemplate()
	return &ListPolicyTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTargetsOfPolicyGroup 查询策略组应用对象
//
// 查询指定策略组所应用的对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListTargetsOfPolicyGroup(request *model.ListTargetsOfPolicyGroupRequest) (*model.ListTargetsOfPolicyGroupResponse, error) {
	requestDef := GenReqDefForListTargetsOfPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTargetsOfPolicyGroupResponse), nil
	}
}

// ListTargetsOfPolicyGroupInvoker 查询策略组应用对象
func (c *WorkspaceAppClient) ListTargetsOfPolicyGroupInvoker(request *model.ListTargetsOfPolicyGroupRequest) *ListTargetsOfPolicyGroupInvoker {
	requestDef := GenReqDefForListTargetsOfPolicyGroup()
	return &ListTargetsOfPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOriginalPolicyInfo 查询初始策略项
//
// 查询初始策略项，初始策略项是所有协议策略配置项的默认配置，用户可以在初始策略项的基础上根据需求修改指定的配置，创建新的策略组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowOriginalPolicyInfo(request *model.ShowOriginalPolicyInfoRequest) (*model.ShowOriginalPolicyInfoResponse, error) {
	requestDef := GenReqDefForShowOriginalPolicyInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOriginalPolicyInfoResponse), nil
	}
}

// ShowOriginalPolicyInfoInvoker 查询初始策略项
func (c *WorkspaceAppClient) ShowOriginalPolicyInfoInvoker(request *model.ShowOriginalPolicyInfoRequest) *ShowOriginalPolicyInfoInvoker {
	requestDef := GenReqDefForShowOriginalPolicyInfo()
	return &ShowOriginalPolicyInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPolicyGroup 查询策略组详情
//
// 根据策略组ID查询策略组详细信息，包含策略信息以及应用对象信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowPolicyGroup(request *model.ShowPolicyGroupRequest) (*model.ShowPolicyGroupResponse, error) {
	requestDef := GenReqDefForShowPolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyGroupResponse), nil
	}
}

// ShowPolicyGroupInvoker 查询策略组详情
func (c *WorkspaceAppClient) ShowPolicyGroupInvoker(request *model.ShowPolicyGroupRequest) *ShowPolicyGroupInvoker {
	requestDef := GenReqDefForShowPolicyGroup()
	return &ShowPolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyGroup 修改策略组
//
// 修改指定策略组的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdatePolicyGroup(request *model.UpdatePolicyGroupRequest) (*model.UpdatePolicyGroupResponse, error) {
	requestDef := GenReqDefForUpdatePolicyGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyGroupResponse), nil
	}
}

// UpdatePolicyGroupInvoker 修改策略组
func (c *WorkspaceAppClient) UpdatePolicyGroupInvoker(request *model.UpdatePolicyGroupRequest) *UpdatePolicyGroupInvoker {
	requestDef := GenReqDefForUpdatePolicyGroup()
	return &UpdatePolicyGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePolicyTemplate 修改策略模板
//
// 修改指定策略模板的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdatePolicyTemplate(request *model.UpdatePolicyTemplateRequest) (*model.UpdatePolicyTemplateResponse, error) {
	requestDef := GenReqDefForUpdatePolicyTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyTemplateResponse), nil
	}
}

// UpdatePolicyTemplateInvoker 修改策略模板
func (c *WorkspaceAppClient) UpdatePolicyTemplateInvoker(request *model.UpdatePolicyTemplateRequest) *UpdatePolicyTemplateInvoker {
	requestDef := GenReqDefForUpdatePolicyTemplate()
	return &UpdatePolicyTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckQuota 配额校验
//
// 配额校验，购买服务器前可用调用该接口，校验本次创建服务器资源是否足够。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CheckQuota(request *model.CheckQuotaRequest) (*model.CheckQuotaResponse, error) {
	requestDef := GenReqDefForCheckQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckQuotaResponse), nil
	}
}

// CheckQuotaInvoker 配额校验
func (c *WorkspaceAppClient) CheckQuotaInvoker(request *model.CheckQuotaRequest) *CheckQuotaInvoker {
	requestDef := GenReqDefForCheckQuota()
	return &CheckQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateScalingPolicy 新增/修改弹性伸缩策略
//
// 新增/修改弹性伸缩策略,仅按需的服务器支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateOrUpdateScalingPolicy(request *model.CreateOrUpdateScalingPolicyRequest) (*model.CreateOrUpdateScalingPolicyResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateScalingPolicyResponse), nil
	}
}

// CreateOrUpdateScalingPolicyInvoker 新增/修改弹性伸缩策略
func (c *WorkspaceAppClient) CreateOrUpdateScalingPolicyInvoker(request *model.CreateOrUpdateScalingPolicyRequest) *CreateOrUpdateScalingPolicyInvoker {
	requestDef := GenReqDefForCreateOrUpdateScalingPolicy()
	return &CreateOrUpdateScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScalingPolicy 删除弹性伸缩策略
//
// 删除弹性伸缩策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteScalingPolicy(request *model.DeleteScalingPolicyRequest) (*model.DeleteScalingPolicyResponse, error) {
	requestDef := GenReqDefForDeleteScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScalingPolicyResponse), nil
	}
}

// DeleteScalingPolicyInvoker 删除弹性伸缩策略
func (c *WorkspaceAppClient) DeleteScalingPolicyInvoker(request *model.DeleteScalingPolicyRequest) *DeleteScalingPolicyInvoker {
	requestDef := GenReqDefForDeleteScalingPolicy()
	return &DeleteScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScalingPolicy 查询服务器组弹性伸缩策略
//
// 查询服务器组弹性伸缩策略,如果服务器未配置策略时响应默认策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowScalingPolicy(request *model.ShowScalingPolicyRequest) (*model.ShowScalingPolicyResponse, error) {
	requestDef := GenReqDefForShowScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScalingPolicyResponse), nil
	}
}

// ShowScalingPolicyInvoker 查询服务器组弹性伸缩策略
func (c *WorkspaceAppClient) ShowScalingPolicyInvoker(request *model.ShowScalingPolicyRequest) *ShowScalingPolicyInvoker {
	requestDef := GenReqDefForShowScalingPolicy()
	return &ShowScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteScheduleTask 批量删除定时任务
//
// 批量删除定时任务，忽略不存在的服务器组并且返回成功响应。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteScheduleTask(request *model.BatchDeleteScheduleTaskRequest) (*model.BatchDeleteScheduleTaskResponse, error) {
	requestDef := GenReqDefForBatchDeleteScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteScheduleTaskResponse), nil
	}
}

// BatchDeleteScheduleTaskInvoker 批量删除定时任务
func (c *WorkspaceAppClient) BatchDeleteScheduleTaskInvoker(request *model.BatchDeleteScheduleTaskRequest) *BatchDeleteScheduleTaskInvoker {
	requestDef := GenReqDefForBatchDeleteScheduleTask()
	return &BatchDeleteScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScheduleTask 新增定时任务
//
// 新增定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateScheduleTask(request *model.CreateScheduleTaskRequest) (*model.CreateScheduleTaskResponse, error) {
	requestDef := GenReqDefForCreateScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduleTaskResponse), nil
	}
}

// CreateScheduleTaskInvoker 新增定时任务
func (c *WorkspaceAppClient) CreateScheduleTaskInvoker(request *model.CreateScheduleTaskRequest) *CreateScheduleTaskInvoker {
	requestDef := GenReqDefForCreateScheduleTask()
	return &CreateScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduleTask 删除任务
//
// 删除任务，忽略不存在的任务并且返回成功响应。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteScheduleTask(request *model.DeleteScheduleTaskRequest) (*model.DeleteScheduleTaskResponse, error) {
	requestDef := GenReqDefForDeleteScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduleTaskResponse), nil
	}
}

// DeleteScheduleTaskInvoker 删除任务
func (c *WorkspaceAppClient) DeleteScheduleTaskInvoker(request *model.DeleteScheduleTaskRequest) *DeleteScheduleTaskInvoker {
	requestDef := GenReqDefForDeleteScheduleTask()
	return &DeleteScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFutureExecutions 未来执行的具体时间列表
//
// 未来执行的具体时间列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListFutureExecutions(request *model.ListFutureExecutionsRequest) (*model.ListFutureExecutionsResponse, error) {
	requestDef := GenReqDefForListFutureExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFutureExecutionsResponse), nil
	}
}

// ListFutureExecutionsInvoker 未来执行的具体时间列表
func (c *WorkspaceAppClient) ListFutureExecutionsInvoker(request *model.ListFutureExecutionsRequest) *ListFutureExecutionsInvoker {
	requestDef := GenReqDefForListFutureExecutions()
	return &ListFutureExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduleTasks 查询定时任务列表
//
// 查询定时任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListScheduleTasks(request *model.ListScheduleTasksRequest) (*model.ListScheduleTasksResponse, error) {
	requestDef := GenReqDefForListScheduleTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduleTasksResponse), nil
	}
}

// ListScheduleTasksInvoker 查询定时任务列表
func (c *WorkspaceAppClient) ListScheduleTasksInvoker(request *model.ListScheduleTasksRequest) *ListScheduleTasksInvoker {
	requestDef := GenReqDefForListScheduleTasks()
	return &ListScheduleTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskExecuteDetail 查询定时任务执行子任务列表
//
// 查询定时任务执行子任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListTaskExecuteDetail(request *model.ListTaskExecuteDetailRequest) (*model.ListTaskExecuteDetailResponse, error) {
	requestDef := GenReqDefForListTaskExecuteDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskExecuteDetailResponse), nil
	}
}

// ListTaskExecuteDetailInvoker 查询定时任务执行子任务列表
func (c *WorkspaceAppClient) ListTaskExecuteDetailInvoker(request *model.ListTaskExecuteDetailRequest) *ListTaskExecuteDetailInvoker {
	requestDef := GenReqDefForListTaskExecuteDetail()
	return &ListTaskExecuteDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTaskExecuteHistory 查询定时任务执行列表
//
// 查询定时任务执行列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListTaskExecuteHistory(request *model.ListTaskExecuteHistoryRequest) (*model.ListTaskExecuteHistoryResponse, error) {
	requestDef := GenReqDefForListTaskExecuteHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTaskExecuteHistoryResponse), nil
	}
}

// ListTaskExecuteHistoryInvoker 查询定时任务执行列表
func (c *WorkspaceAppClient) ListTaskExecuteHistoryInvoker(request *model.ListTaskExecuteHistoryRequest) *ListTaskExecuteHistoryInvoker {
	requestDef := GenReqDefForListTaskExecuteHistory()
	return &ListTaskExecuteHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScheduleTask 查询指定定时任务详情
//
// 查询指定定时任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowScheduleTask(request *model.ShowScheduleTaskRequest) (*model.ShowScheduleTaskResponse, error) {
	requestDef := GenReqDefForShowScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScheduleTaskResponse), nil
	}
}

// ShowScheduleTaskInvoker 查询指定定时任务详情
func (c *WorkspaceAppClient) ShowScheduleTaskInvoker(request *model.ShowScheduleTaskRequest) *ShowScheduleTaskInvoker {
	requestDef := GenReqDefForShowScheduleTask()
	return &ShowScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScheduleTask 修改定时任务
//
// 修改定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateScheduleTask(request *model.UpdateScheduleTaskRequest) (*model.UpdateScheduleTaskResponse, error) {
	requestDef := GenReqDefForUpdateScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduleTaskResponse), nil
	}
}

// UpdateScheduleTaskInvoker 修改定时任务
func (c *WorkspaceAppClient) UpdateScheduleTaskInvoker(request *model.UpdateScheduleTaskRequest) *UpdateScheduleTaskInvoker {
	requestDef := GenReqDefForUpdateScheduleTask()
	return &UpdateScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchChangeServerImage 批量修改服务器的镜像
//
// 批量修改服务器的镜像。
// &gt; - 服务器的镜像和服务器组的镜像不一样时，支持服务器的镜像切换为服务器组的镜像，并且仅允许同等镜像进行切换，例如：同操作系统，免费镜像切换，同源同价的付费镜像切换。如果服务器组的镜像和服务器的镜像为非同等镜像，建议您直接购买新的服务器，删除或者退订老的服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchChangeServerImage(request *model.BatchChangeServerImageRequest) (*model.BatchChangeServerImageResponse, error) {
	requestDef := GenReqDefForBatchChangeServerImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeServerImageResponse), nil
	}
}

// BatchChangeServerImageInvoker 批量修改服务器的镜像
func (c *WorkspaceAppClient) BatchChangeServerImageInvoker(request *model.BatchChangeServerImageRequest) *BatchChangeServerImageInvoker {
	requestDef := GenReqDefForBatchChangeServerImage()
	return &BatchChangeServerImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchChangeServerMaintainMode 标记服务器维护状态
//
// 标记服务器维护状态，处于维护状态中的服务器不再分配流量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchChangeServerMaintainMode(request *model.BatchChangeServerMaintainModeRequest) (*model.BatchChangeServerMaintainModeResponse, error) {
	requestDef := GenReqDefForBatchChangeServerMaintainMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeServerMaintainModeResponse), nil
	}
}

// BatchChangeServerMaintainModeInvoker 标记服务器维护状态
func (c *WorkspaceAppClient) BatchChangeServerMaintainModeInvoker(request *model.BatchChangeServerMaintainModeRequest) *BatchChangeServerMaintainModeInvoker {
	requestDef := GenReqDefForBatchChangeServerMaintainMode()
	return &BatchChangeServerMaintainModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServer 批量删除服务器
//
// 批量删除服务器。
// &gt; - 仅支持删除按需订购的服务器，包周期订购的服务器需要到Console界面进行退订，订单退订成功后服务器将会自动删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteServer(request *model.BatchDeleteServerRequest) (*model.BatchDeleteServerResponse, error) {
	requestDef := GenReqDefForBatchDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerResponse), nil
	}
}

// BatchDeleteServerInvoker 批量删除服务器
func (c *WorkspaceAppClient) BatchDeleteServerInvoker(request *model.BatchDeleteServerRequest) *BatchDeleteServerInvoker {
	requestDef := GenReqDefForBatchDeleteServer()
	return &BatchDeleteServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchMigrateHostsServer 迁移云办公主机下面的服务器到目标云办公主机
//
// 迁移云办公主机下面的服务器到目标云办公主机。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchMigrateHostsServer(request *model.BatchMigrateHostsServerRequest) (*model.BatchMigrateHostsServerResponse, error) {
	requestDef := GenReqDefForBatchMigrateHostsServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchMigrateHostsServerResponse), nil
	}
}

// BatchMigrateHostsServerInvoker 迁移云办公主机下面的服务器到目标云办公主机
func (c *WorkspaceAppClient) BatchMigrateHostsServerInvoker(request *model.BatchMigrateHostsServerRequest) *BatchMigrateHostsServerInvoker {
	requestDef := GenReqDefForBatchMigrateHostsServer()
	return &BatchMigrateHostsServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebootServer 重启服务器
//
// 重启服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchRebootServer(request *model.BatchRebootServerRequest) (*model.BatchRebootServerResponse, error) {
	requestDef := GenReqDefForBatchRebootServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootServerResponse), nil
	}
}

// BatchRebootServerInvoker 重启服务器
func (c *WorkspaceAppClient) BatchRebootServerInvoker(request *model.BatchRebootServerRequest) *BatchRebootServerInvoker {
	requestDef := GenReqDefForBatchRebootServer()
	return &BatchRebootServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchReinstallServer 批量重装服务器
//
// 批量重装服务器。
// &gt; - 使用服务器原有的镜像进行重装，当服务器异常无法恢复时或者服务器运行时间久了，性能下降时，可选择重建服务器。注意：重装时系统盘的数据将会被清理掉。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchReinstallServer(request *model.BatchReinstallServerRequest) (*model.BatchReinstallServerResponse, error) {
	requestDef := GenReqDefForBatchReinstallServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchReinstallServerResponse), nil
	}
}

// BatchReinstallServerInvoker 批量重装服务器
func (c *WorkspaceAppClient) BatchReinstallServerInvoker(request *model.BatchReinstallServerRequest) *BatchReinstallServerInvoker {
	requestDef := GenReqDefForBatchReinstallServer()
	return &BatchReinstallServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRejoinDomain 批量服务器重新加域
//
// 批量服务器重新加入AD域，一般用于解决服务器脱域的情况使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchRejoinDomain(request *model.BatchRejoinDomainRequest) (*model.BatchRejoinDomainResponse, error) {
	requestDef := GenReqDefForBatchRejoinDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRejoinDomainResponse), nil
	}
}

// BatchRejoinDomainInvoker 批量服务器重新加域
func (c *WorkspaceAppClient) BatchRejoinDomainInvoker(request *model.BatchRejoinDomainRequest) *BatchRejoinDomainInvoker {
	requestDef := GenReqDefForBatchRejoinDomain()
	return &BatchRejoinDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStartServer 启动服务器
//
// 启动服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchStartServer(request *model.BatchStartServerRequest) (*model.BatchStartServerResponse, error) {
	requestDef := GenReqDefForBatchStartServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStartServerResponse), nil
	}
}

// BatchStartServerInvoker 启动服务器
func (c *WorkspaceAppClient) BatchStartServerInvoker(request *model.BatchStartServerRequest) *BatchStartServerInvoker {
	requestDef := GenReqDefForBatchStartServer()
	return &BatchStartServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopServer 关闭服务器
//
// 关闭服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchStopServer(request *model.BatchStopServerRequest) (*model.BatchStopServerResponse, error) {
	requestDef := GenReqDefForBatchStopServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopServerResponse), nil
	}
}

// BatchStopServerInvoker 关闭服务器
func (c *WorkspaceAppClient) BatchStopServerInvoker(request *model.BatchStopServerRequest) *BatchStopServerInvoker {
	requestDef := GenReqDefForBatchStopServer()
	return &BatchStopServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateTsvi 批量更新服务器虚拟会话IP配置
//
// 批量更新服务器虚拟会话IP配置，按需重启机器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchUpdateTsvi(request *model.BatchUpdateTsviRequest) (*model.BatchUpdateTsviResponse, error) {
	requestDef := GenReqDefForBatchUpdateTsvi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateTsviResponse), nil
	}
}

// BatchUpdateTsviInvoker 批量更新服务器虚拟会话IP配置
func (c *WorkspaceAppClient) BatchUpdateTsviInvoker(request *model.BatchUpdateTsviRequest) *BatchUpdateTsviInvoker {
	requestDef := GenReqDefForBatchUpdateTsvi()
	return &BatchUpdateTsviInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpgradeHdaVersion 批量升级服务器HDA版本
//
// 批量升级服务器HDA版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchUpgradeHdaVersion(request *model.BatchUpgradeHdaVersionRequest) (*model.BatchUpgradeHdaVersionResponse, error) {
	requestDef := GenReqDefForBatchUpgradeHdaVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpgradeHdaVersionResponse), nil
	}
}

// BatchUpgradeHdaVersionInvoker 批量升级服务器HDA版本
func (c *WorkspaceAppClient) BatchUpgradeHdaVersionInvoker(request *model.BatchUpgradeHdaVersionRequest) *BatchUpgradeHdaVersionInvoker {
	requestDef := GenReqDefForBatchUpgradeHdaVersion()
	return &BatchUpgradeHdaVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeServerImage 修改服务器的镜像
//
// 修改服务器的镜像。
// &gt; - 服务器的镜像和服务器组的镜像不一样时，支持服务器的镜像切换为服务器组的镜像，并且仅允许同等镜像进行切换，例如：同操作系统，免费镜像切换，同源同价的付费镜像切换。如果服务器组的镜像和服务器的镜像为非同等镜像，建议您直接购买新的服务器，删除或者退订老的服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ChangeServerImage(request *model.ChangeServerImageRequest) (*model.ChangeServerImageResponse, error) {
	requestDef := GenReqDefForChangeServerImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeServerImageResponse), nil
	}
}

// ChangeServerImageInvoker 修改服务器的镜像
func (c *WorkspaceAppClient) ChangeServerImageInvoker(request *model.ChangeServerImageRequest) *ChangeServerImageInvoker {
	requestDef := GenReqDefForChangeServerImage()
	return &ChangeServerImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppServers 创建云服务器
//
// 创建云服务器接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateAppServers(request *model.CreateAppServersRequest) (*model.CreateAppServersResponse, error) {
	requestDef := GenReqDefForCreateAppServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppServersResponse), nil
	}
}

// CreateAppServersInvoker 创建云服务器
func (c *WorkspaceAppClient) CreateAppServersInvoker(request *model.CreateAppServersRequest) *CreateAppServersInvoker {
	requestDef := GenReqDefForCreateAppServers()
	return &CreateAppServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServer 删除服务器
//
// 删除服务器，忽略不存在的服务器并且返回成功响应。订单退订成功后调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteServer(request *model.DeleteServerRequest) (*model.DeleteServerResponse, error) {
	requestDef := GenReqDefForDeleteServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerResponse), nil
	}
}

// DeleteServerInvoker 删除服务器
func (c *WorkspaceAppClient) DeleteServerInvoker(request *model.DeleteServerRequest) *DeleteServerInvoker {
	requestDef := GenReqDefForDeleteServer()
	return &DeleteServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessAgentLatestVersion 查询租户的所有HDA最新版本
//
// 查询租户的所有HDA最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAccessAgentLatestVersion(request *model.ListAccessAgentLatestVersionRequest) (*model.ListAccessAgentLatestVersionResponse, error) {
	requestDef := GenReqDefForListAccessAgentLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessAgentLatestVersionResponse), nil
	}
}

// ListAccessAgentLatestVersionInvoker 查询租户的所有HDA最新版本
func (c *WorkspaceAppClient) ListAccessAgentLatestVersionInvoker(request *model.ListAccessAgentLatestVersionRequest) *ListAccessAgentLatestVersionInvoker {
	requestDef := GenReqDefForListAccessAgentLatestVersion()
	return &ListAccessAgentLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerHdaDetails 查询服务器的HDA相关信息
//
// 查询服务器的HDA相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServerHdaDetails(request *model.ListServerHdaDetailsRequest) (*model.ListServerHdaDetailsResponse, error) {
	requestDef := GenReqDefForListServerHdaDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerHdaDetailsResponse), nil
	}
}

// ListServerHdaDetailsInvoker 查询服务器的HDA相关信息
func (c *WorkspaceAppClient) ListServerHdaDetailsInvoker(request *model.ListServerHdaDetailsRequest) *ListServerHdaDetailsInvoker {
	requestDef := GenReqDefForListServerHdaDetails()
	return &ListServerHdaDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerHdaUpgradeRecords 查询服务器的HDA升级跟踪记录
//
// 查询服务器的HDA升级跟踪记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServerHdaUpgradeRecords(request *model.ListServerHdaUpgradeRecordsRequest) (*model.ListServerHdaUpgradeRecordsResponse, error) {
	requestDef := GenReqDefForListServerHdaUpgradeRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerHdaUpgradeRecordsResponse), nil
	}
}

// ListServerHdaUpgradeRecordsInvoker 查询服务器的HDA升级跟踪记录
func (c *WorkspaceAppClient) ListServerHdaUpgradeRecordsInvoker(request *model.ListServerHdaUpgradeRecordsRequest) *ListServerHdaUpgradeRecordsInvoker {
	requestDef := GenReqDefForListServerHdaUpgradeRecords()
	return &ListServerHdaUpgradeRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerMetricData 查询指定时间范围指定指标的指定粒度的监控数据
//
// 查询指定时间范围指定指标的指定粒度的监控数据，可以通过参数指定需要查询的数据维度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServerMetricData(request *model.ListServerMetricDataRequest) (*model.ListServerMetricDataResponse, error) {
	requestDef := GenReqDefForListServerMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerMetricDataResponse), nil
	}
}

// ListServerMetricDataInvoker 查询指定时间范围指定指标的指定粒度的监控数据
func (c *WorkspaceAppClient) ListServerMetricDataInvoker(request *model.ListServerMetricDataRequest) *ListServerMetricDataInvoker {
	requestDef := GenReqDefForListServerMetricData()
	return &ListServerMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServers 查询服务器列表
//
// 查询服务器列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServers(request *model.ListServersRequest) (*model.ListServersResponse, error) {
	requestDef := GenReqDefForListServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServersResponse), nil
	}
}

// ListServersInvoker 查询服务器列表
func (c *WorkspaceAppClient) ListServersInvoker(request *model.ListServersRequest) *ListServersInvoker {
	requestDef := GenReqDefForListServers()
	return &ListServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallServer 重装服务器
//
// 重装服务器。
// &gt; - 使用服务器原有的镜像进行重装，当服务器异常无法恢复时或者服务器运行时间久了，性能下降时，可选择重建服务器。注意：重装时系统盘的数据将会被清理掉。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ReinstallServer(request *model.ReinstallServerRequest) (*model.ReinstallServerResponse, error) {
	requestDef := GenReqDefForReinstallServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallServerResponse), nil
	}
}

// ReinstallServerInvoker 重装服务器
func (c *WorkspaceAppClient) ReinstallServerInvoker(request *model.ReinstallServerRequest) *ReinstallServerInvoker {
	requestDef := GenReqDefForReinstallServer()
	return &ReinstallServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessAgentLatestVersion 查询租户的HDA最新版本
//
// 查询租户的HDA最新版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowAccessAgentLatestVersion(request *model.ShowAccessAgentLatestVersionRequest) (*model.ShowAccessAgentLatestVersionResponse, error) {
	requestDef := GenReqDefForShowAccessAgentLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessAgentLatestVersionResponse), nil
	}
}

// ShowAccessAgentLatestVersionInvoker 查询租户的HDA最新版本
func (c *WorkspaceAppClient) ShowAccessAgentLatestVersionInvoker(request *model.ShowAccessAgentLatestVersionRequest) *ShowAccessAgentLatestVersionInvoker {
	requestDef := GenReqDefForShowAccessAgentLatestVersion()
	return &ShowAccessAgentLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServer 查询指定服务器
//
// 查询指定的服务器当前这个接口的查询数据和list查询的一致。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServer(request *model.ShowServerRequest) (*model.ShowServerResponse, error) {
	requestDef := GenReqDefForShowServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerResponse), nil
	}
}

// ShowServerInvoker 查询指定服务器
func (c *WorkspaceAppClient) ShowServerInvoker(request *model.ShowServerRequest) *ShowServerInvoker {
	requestDef := GenReqDefForShowServer()
	return &ShowServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerMetricData 查询云应用服务器监控信息
//
// 该接口可获取某一计算机在一段时间段(范围：1小时到30天)的数据信息（例如CPU占用率、内存占用率、用户的在线时间段等），最长数据保存时间不能超过180天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerMetricData(request *model.ShowServerMetricDataRequest) (*model.ShowServerMetricDataResponse, error) {
	requestDef := GenReqDefForShowServerMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerMetricDataResponse), nil
	}
}

// ShowServerMetricDataInvoker 查询云应用服务器监控信息
func (c *WorkspaceAppClient) ShowServerMetricDataInvoker(request *model.ShowServerMetricDataRequest) *ShowServerMetricDataInvoker {
	requestDef := GenReqDefForShowServerMetricData()
	return &ShowServerMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerVnc 获取VNC远程登录地址
//
// 获取VNC远程登录地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerVnc(request *model.ShowServerVncRequest) (*model.ShowServerVncResponse, error) {
	requestDef := GenReqDefForShowServerVnc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerVncResponse), nil
	}
}

// ShowServerVncInvoker 获取VNC远程登录地址
func (c *WorkspaceAppClient) ShowServerVncInvoker(request *model.ShowServerVncRequest) *ShowServerVncInvoker {
	requestDef := GenReqDefForShowServerVnc()
	return &ShowServerVncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServer 修改服务器
//
// 修改服务器。
// &gt; - 服务器的状态修改为维护模式后，用户打开应用，选择可用的服务器进行接入的时候，会过滤掉处于维护模式的服务器。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateServer(request *model.UpdateServerRequest) (*model.UpdateServerResponse, error) {
	requestDef := GenReqDefForUpdateServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerResponse), nil
	}
}

// UpdateServerInvoker 修改服务器
func (c *WorkspaceAppClient) UpdateServerInvoker(request *model.UpdateServerRequest) *UpdateServerInvoker {
	requestDef := GenReqDefForUpdateServer()
	return &UpdateServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServerGroup 创建服务器组
//
// 创建服务器组。
// &gt; - 服务器组是一组相同配置的服务器集合，服务器组内的服务器使用同一镜像创建，配置相同，运行相同的应用程序。用户在打开云应用时，会根据调度规则选取组内的一台可用服务器进行连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateServerGroup(request *model.CreateServerGroupRequest) (*model.CreateServerGroupResponse, error) {
	requestDef := GenReqDefForCreateServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServerGroupResponse), nil
	}
}

// CreateServerGroupInvoker 创建服务器组
func (c *WorkspaceAppClient) CreateServerGroupInvoker(request *model.CreateServerGroupRequest) *CreateServerGroupInvoker {
	requestDef := GenReqDefForCreateServerGroup()
	return &CreateServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerGroups 删除服务器组
//
// 删除服务器组。
// - &gt; 删除服务器组之前，需要先删除服务器组内的所有服务器。如果传服务器组已被删除，重复执行删除，则返回成功响应。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteServerGroups(request *model.DeleteServerGroupsRequest) (*model.DeleteServerGroupsResponse, error) {
	requestDef := GenReqDefForDeleteServerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupsResponse), nil
	}
}

// DeleteServerGroupsInvoker 删除服务器组
func (c *WorkspaceAppClient) DeleteServerGroupsInvoker(request *model.DeleteServerGroupsRequest) *DeleteServerGroupsInvoker {
	requestDef := GenReqDefForDeleteServerGroups()
	return &DeleteServerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerGroups 查询服务器组列表
//
// 查询服务器组列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServerGroups(request *model.ListServerGroupsRequest) (*model.ListServerGroupsResponse, error) {
	requestDef := GenReqDefForListServerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerGroupsResponse), nil
	}
}

// ListServerGroupsInvoker 查询服务器组列表
func (c *WorkspaceAppClient) ListServerGroupsInvoker(request *model.ListServerGroupsRequest) *ListServerGroupsInvoker {
	requestDef := GenReqDefForListServerGroups()
	return &ListServerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTenantServerGroups 查询租户服务器组基础信息列表
//
// 查询租户服务器组基础信息列表(用于创建应用组时绑定服务器组)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListTenantServerGroups(request *model.ListTenantServerGroupsRequest) (*model.ListTenantServerGroupsResponse, error) {
	requestDef := GenReqDefForListTenantServerGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTenantServerGroupsResponse), nil
	}
}

// ListTenantServerGroupsInvoker 查询租户服务器组基础信息列表
func (c *WorkspaceAppClient) ListTenantServerGroupsInvoker(request *model.ListTenantServerGroupsRequest) *ListTenantServerGroupsInvoker {
	requestDef := GenReqDefForListTenantServerGroups()
	return &ListTenantServerGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerGroup 查询指定服务器组
//
// 查询指定的服务器组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerGroup(request *model.ShowServerGroupRequest) (*model.ShowServerGroupResponse, error) {
	requestDef := GenReqDefForShowServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupResponse), nil
	}
}

// ShowServerGroupInvoker 查询指定服务器组
func (c *WorkspaceAppClient) ShowServerGroupInvoker(request *model.ShowServerGroupRequest) *ShowServerGroupInvoker {
	requestDef := GenReqDefForShowServerGroup()
	return &ShowServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerGroupRestrict 指定租户服务器组限制查询
//
// 指定租户服务器组限制查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerGroupRestrict(request *model.ShowServerGroupRestrictRequest) (*model.ShowServerGroupRestrictResponse, error) {
	requestDef := GenReqDefForShowServerGroupRestrict()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupRestrictResponse), nil
	}
}

// ShowServerGroupRestrictInvoker 指定租户服务器组限制查询
func (c *WorkspaceAppClient) ShowServerGroupRestrictInvoker(request *model.ShowServerGroupRestrictRequest) *ShowServerGroupRestrictInvoker {
	requestDef := GenReqDefForShowServerGroupRestrict()
	return &ShowServerGroupRestrictInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerGroupState 查询指定服务器组内服务器状态
//
// 查询指定的服务器组内服务器状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerGroupState(request *model.ShowServerGroupStateRequest) (*model.ShowServerGroupStateResponse, error) {
	requestDef := GenReqDefForShowServerGroupState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupStateResponse), nil
	}
}

// ShowServerGroupStateInvoker 查询指定服务器组内服务器状态
func (c *WorkspaceAppClient) ShowServerGroupStateInvoker(request *model.ShowServerGroupStateRequest) *ShowServerGroupStateInvoker {
	requestDef := GenReqDefForShowServerGroupState()
	return &ShowServerGroupStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerGroup 修改服务器组
//
// 修改服务器组。
// - &gt; 修改服务器组的镜像，系统盘大小，OU信息后，已创建的服务器配置不变，新添加的服务器会使用新的配置创建。修改最大会话数后，用户接入服务器组时，会按照最新的配置进行路由计算。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) UpdateServerGroup(request *model.UpdateServerGroupRequest) (*model.UpdateServerGroupResponse, error) {
	requestDef := GenReqDefForUpdateServerGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerGroupResponse), nil
	}
}

// UpdateServerGroupInvoker 修改服务器组
func (c *WorkspaceAppClient) UpdateServerGroupInvoker(request *model.UpdateServerGroupRequest) *UpdateServerGroupInvoker {
	requestDef := GenReqDefForUpdateServerGroup()
	return &UpdateServerGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppConnection 查询应用使用记录
//
// 查询应用使用记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListAppConnection(request *model.ListAppConnectionRequest) (*model.ListAppConnectionResponse, error) {
	requestDef := GenReqDefForListAppConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppConnectionResponse), nil
	}
}

// ListAppConnectionInvoker 查询应用使用记录
func (c *WorkspaceAppClient) ListAppConnectionInvoker(request *model.ListAppConnectionRequest) *ListAppConnectionInvoker {
	requestDef := GenReqDefForListAppConnection()
	return &ListAppConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionByUserName 根据用户名查询当前会话
//
// 根据用户名查询当前会话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListSessionByUserName(request *model.ListSessionByUserNameRequest) (*model.ListSessionByUserNameResponse, error) {
	requestDef := GenReqDefForListSessionByUserName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionByUserNameResponse), nil
	}
}

// ListSessionByUserNameInvoker 根据用户名查询当前会话
func (c *WorkspaceAppClient) ListSessionByUserNameInvoker(request *model.ListSessionByUserNameRequest) *ListSessionByUserNameInvoker {
	requestDef := GenReqDefForListSessionByUserName()
	return &ListSessionByUserNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessions 查询用户会话列表
//
// 查询用户会话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListSessions(request *model.ListSessionsRequest) (*model.ListSessionsResponse, error) {
	requestDef := GenReqDefForListSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionsResponse), nil
	}
}

// ListSessionsInvoker 查询用户会话列表
func (c *WorkspaceAppClient) ListSessionsInvoker(request *model.ListSessionsRequest) *ListSessionsInvoker {
	requestDef := GenReqDefForListSessions()
	return &ListSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserConnection 查询用户登录记录
//
// 查询用户登录记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListUserConnection(request *model.ListUserConnectionRequest) (*model.ListUserConnectionResponse, error) {
	requestDef := GenReqDefForListUserConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserConnectionResponse), nil
	}
}

// ListUserConnectionInvoker 查询用户登录记录
func (c *WorkspaceAppClient) ListUserConnectionInvoker(request *model.ListUserConnectionRequest) *ListUserConnectionInvoker {
	requestDef := GenReqDefForListUserConnection()
	return &ListUserConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// LogoffUserSession 用户会话注销
//
// 用户会话注销。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) LogoffUserSession(request *model.LogoffUserSessionRequest) (*model.LogoffUserSessionResponse, error) {
	requestDef := GenReqDefForLogoffUserSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.LogoffUserSessionResponse), nil
	}
}

// LogoffUserSessionInvoker 用户会话注销
func (c *WorkspaceAppClient) LogoffUserSessionInvoker(request *model.LogoffUserSessionRequest) *LogoffUserSessionInvoker {
	requestDef := GenReqDefForLogoffUserSession()
	return &LogoffUserSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateServerGroupTags 批量添加服务器组标签
//
// 此接口为幂等接口：
// 同时对多个服务器组批量添加标签，最大支持100个服务器组，每个服务器组最大20个标签
// 创建时如果请求体中存在重复key则报错。
// 创建时，不允许设置重复key数据,如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchCreateServerGroupTags(request *model.BatchCreateServerGroupTagsRequest) (*model.BatchCreateServerGroupTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateServerGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateServerGroupTagsResponse), nil
	}
}

// BatchCreateServerGroupTagsInvoker 批量添加服务器组标签
func (c *WorkspaceAppClient) BatchCreateServerGroupTagsInvoker(request *model.BatchCreateServerGroupTagsRequest) *BatchCreateServerGroupTagsInvoker {
	requestDef := GenReqDefForBatchCreateServerGroupTags()
	return &BatchCreateServerGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServerGroupTags 批量删除服务器组标签
//
// 此接口为幂等接口：
// 同时对多个服务器组批量删除标签，最大支持100个服务器组，每个服务器组最大20个标签。
// 删除时，如果删除的标签不存在，默认处理成功，删除时不对标签字符集范围做校验。
// 删除时tags结构体不能缺失，key不能为空，或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) BatchDeleteServerGroupTags(request *model.BatchDeleteServerGroupTagsRequest) (*model.BatchDeleteServerGroupTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServerGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServerGroupTagsResponse), nil
	}
}

// BatchDeleteServerGroupTagsInvoker 批量删除服务器组标签
func (c *WorkspaceAppClient) BatchDeleteServerGroupTagsInvoker(request *model.BatchDeleteServerGroupTagsRequest) *BatchDeleteServerGroupTagsInvoker {
	requestDef := GenReqDefForBatchDeleteServerGroupTags()
	return &BatchDeleteServerGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServerGroupTags 添加服务器组标签
//
// 此接口为幂等接口：
// 创建时如果请求体中存在重复key则报错。
// 创建时，不允许设置重复key数据,如果数据库已存在该key，就覆盖value的值。
// 一个服务器组上最多有20个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) CreateServerGroupTags(request *model.CreateServerGroupTagsRequest) (*model.CreateServerGroupTagsResponse, error) {
	requestDef := GenReqDefForCreateServerGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServerGroupTagsResponse), nil
	}
}

// CreateServerGroupTagsInvoker 添加服务器组标签
func (c *WorkspaceAppClient) CreateServerGroupTagsInvoker(request *model.CreateServerGroupTagsRequest) *CreateServerGroupTagsInvoker {
	requestDef := GenReqDefForCreateServerGroupTags()
	return &CreateServerGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServerGroupTags 删除服务器组标签
//
// 此接口为幂等接口：
// 删除时，如果删除的标签不存在，默认处理成功,删除时不对标签字符集范围做校验。
// 删除时tags结构体不能缺失，key不能为空，或者空字符串。
// 支持最多每次删除20个标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) DeleteServerGroupTags(request *model.DeleteServerGroupTagsRequest) (*model.DeleteServerGroupTagsResponse, error) {
	requestDef := GenReqDefForDeleteServerGroupTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServerGroupTagsResponse), nil
	}
}

// DeleteServerGroupTagsInvoker 删除服务器组标签
func (c *WorkspaceAppClient) DeleteServerGroupTagsInvoker(request *model.DeleteServerGroupTagsRequest) *DeleteServerGroupTagsInvoker {
	requestDef := GenReqDefForDeleteServerGroupTags()
	return &DeleteServerGroupTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServerGroupTag 查询租户在所有服务器组上的标签
//
// 查询租户在所有服务器组上的资源标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListServerGroupTag(request *model.ListServerGroupTagRequest) (*model.ListServerGroupTagResponse, error) {
	requestDef := GenReqDefForListServerGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServerGroupTagResponse), nil
	}
}

// ListServerGroupTagInvoker 查询租户在所有服务器组上的标签
func (c *WorkspaceAppClient) ListServerGroupTagInvoker(request *model.ListServerGroupTagRequest) *ListServerGroupTagInvoker {
	requestDef := GenReqDefForListServerGroupTag()
	return &ListServerGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowServerGroupTag 查询服务器组的标签
//
// 查询指定服务器组的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ShowServerGroupTag(request *model.ShowServerGroupTagRequest) (*model.ShowServerGroupTagResponse, error) {
	requestDef := GenReqDefForShowServerGroupTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowServerGroupTagResponse), nil
	}
}

// ShowServerGroupTagInvoker 查询服务器组的标签
func (c *WorkspaceAppClient) ShowServerGroupTagInvoker(request *model.ShowServerGroupTagRequest) *ShowServerGroupTagInvoker {
	requestDef := GenReqDefForShowServerGroupTag()
	return &ShowServerGroupTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVolumeType 查询可用磁盘类型
//
// 该接口用于查询可用磁盘类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceAppClient) ListVolumeType(request *model.ListVolumeTypeRequest) (*model.ListVolumeTypeResponse, error) {
	requestDef := GenReqDefForListVolumeType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVolumeTypeResponse), nil
	}
}

// ListVolumeTypeInvoker 查询可用磁盘类型
func (c *WorkspaceAppClient) ListVolumeTypeInvoker(request *model.ListVolumeTypeRequest) *ListVolumeTypeInvoker {
	requestDef := GenReqDefForListVolumeType()
	return &ListVolumeTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
