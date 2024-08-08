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

// UpdateWarehouseApp 修改应用仓库中的指定应用信息
//
// 修改应用仓库中的指定应用信息
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

// ShowJob 查询任务的执行状态
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

// ShowJobInvoker 查询任务的执行状态
func (c *WorkspaceAppClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 查询任务的执行状态
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

// ShowJobDetailInvoker 查询任务的执行状态
func (c *WorkspaceAppClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 新增策略组，通过策略组能灵活的控制客户端访问与接入策略，如：文件、剪切板、会话等。
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
