package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspace/v2/model"
)

type WorkspaceClient struct {
	HcClient *http_client.HcHttpClient
}

func NewWorkspaceClient(hcClient *http_client.HcHttpClient) *WorkspaceClient {
	return &WorkspaceClient{HcClient: hcClient}
}

func WorkspaceClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListAvailabilityZones 查询可用分区列表
//
// 该接口用于查询云桌面支持的可用分区列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListAvailabilityZones(request *model.ListAvailabilityZonesRequest) (*model.ListAvailabilityZonesResponse, error) {
	requestDef := GenReqDefForListAvailabilityZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailabilityZonesResponse), nil
	}
}

// ListAvailabilityZonesInvoker 查询可用分区列表
func (c *WorkspaceClient) ListAvailabilityZonesInvoker(request *model.ListAvailabilityZonesRequest) *ListAvailabilityZonesInvoker {
	requestDef := GenReqDefForListAvailabilityZones()
	return &ListAvailabilityZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportUserLoginInfoNew 导出连接记录
//
// 该接口用于导出连接记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ExportUserLoginInfoNew(request *model.ExportUserLoginInfoNewRequest) (*model.ExportUserLoginInfoNewResponse, error) {
	requestDef := GenReqDefForExportUserLoginInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportUserLoginInfoNewResponse), nil
	}
}

// ExportUserLoginInfoNewInvoker 导出连接记录
func (c *WorkspaceClient) ExportUserLoginInfoNewInvoker(request *model.ExportUserLoginInfoNewRequest) *ExportUserLoginInfoNewInvoker {
	requestDef := GenReqDefForExportUserLoginInfoNew()
	return &ExportUserLoginInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryOnlineInfoNew 查询登录人数
//
// 该接口用于查询登录人数，注意查询参数不可全空。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListHistoryOnlineInfoNew(request *model.ListHistoryOnlineInfoNewRequest) (*model.ListHistoryOnlineInfoNewResponse, error) {
	requestDef := GenReqDefForListHistoryOnlineInfoNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOnlineInfoNewResponse), nil
	}
}

// ListHistoryOnlineInfoNewInvoker 查询登录人数
func (c *WorkspaceClient) ListHistoryOnlineInfoNewInvoker(request *model.ListHistoryOnlineInfoNewRequest) *ListHistoryOnlineInfoNewInvoker {
	requestDef := GenReqDefForListHistoryOnlineInfoNew()
	return &ListHistoryOnlineInfoNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLoginRecordsNew 查询登录信息
//
// 该接口用于查询登录信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListLoginRecordsNew(request *model.ListLoginRecordsNewRequest) (*model.ListLoginRecordsNewResponse, error) {
	requestDef := GenReqDefForListLoginRecordsNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLoginRecordsNewResponse), nil
	}
}

// ListLoginRecordsNewInvoker 查询登录信息
func (c *WorkspaceClient) ListLoginRecordsNewInvoker(request *model.ListLoginRecordsNewRequest) *ListLoginRecordsNewInvoker {
	requestDef := GenReqDefForListLoginRecordsNew()
	return &ListLoginRecordsNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteDesktops 批量删除桌面
//
// 批量删除桌面，删除后无法恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteDesktops(request *model.BatchDeleteDesktopsRequest) (*model.BatchDeleteDesktopsResponse, error) {
	requestDef := GenReqDefForBatchDeleteDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteDesktopsResponse), nil
	}
}

// BatchDeleteDesktopsInvoker 批量删除桌面
func (c *WorkspaceClient) BatchDeleteDesktopsInvoker(request *model.BatchDeleteDesktopsRequest) *BatchDeleteDesktopsInvoker {
	requestDef := GenReqDefForBatchDeleteDesktops()
	return &BatchDeleteDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRunDesktops 操作桌面
//
// 批量操作桌面，用于批量开机、关机和重启。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) BatchRunDesktops(request *model.BatchRunDesktopsRequest) (*model.BatchRunDesktopsResponse, error) {
	requestDef := GenReqDefForBatchRunDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRunDesktopsResponse), nil
	}
}

// BatchRunDesktopsInvoker 操作桌面
func (c *WorkspaceClient) BatchRunDesktopsInvoker(request *model.BatchRunDesktopsRequest) *BatchRunDesktopsInvoker {
	requestDef := GenReqDefForBatchRunDesktops()
	return &BatchRunDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktop 创建桌面
//
// 创建桌面，并将此桌面分配给用户，当桌面创建成功后用户可以登录使用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktop(request *model.CreateDesktopRequest) (*model.CreateDesktopResponse, error) {
	requestDef := GenReqDefForCreateDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopResponse), nil
	}
}

// CreateDesktopInvoker 创建桌面
func (c *WorkspaceClient) CreateDesktopInvoker(request *model.CreateDesktopRequest) *CreateDesktopInvoker {
	requestDef := GenReqDefForCreateDesktop()
	return &CreateDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesktop 删除单个桌面
//
// 删除单个桌面，删除后无法恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktop(request *model.DeleteDesktopRequest) (*model.DeleteDesktopResponse, error) {
	requestDef := GenReqDefForDeleteDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopResponse), nil
	}
}

// DeleteDesktopInvoker 删除单个桌面
func (c *WorkspaceClient) DeleteDesktopInvoker(request *model.DeleteDesktopRequest) *DeleteDesktopInvoker {
	requestDef := GenReqDefForDeleteDesktop()
	return &DeleteDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktops 查询桌面列表
//
// 该接口用于查询桌面虚拟机列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktops(request *model.ListDesktopsRequest) (*model.ListDesktopsResponse, error) {
	requestDef := GenReqDefForListDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsResponse), nil
	}
}

// ListDesktopsInvoker 查询桌面列表
func (c *WorkspaceClient) ListDesktopsInvoker(request *model.ListDesktopsRequest) *ListDesktopsInvoker {
	requestDef := GenReqDefForListDesktops()
	return &ListDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsDetail 查询桌面详情列表
//
// 查询桌面详情信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsDetail(request *model.ListDesktopsDetailRequest) (*model.ListDesktopsDetailResponse, error) {
	requestDef := GenReqDefForListDesktopsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsDetailResponse), nil
	}
}

// ListDesktopsDetailInvoker 查询桌面详情列表
func (c *WorkspaceClient) ListDesktopsDetailInvoker(request *model.ListDesktopsDetailRequest) *ListDesktopsDetailInvoker {
	requestDef := GenReqDefForListDesktopsDetail()
	return &ListDesktopsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeDesktop 变更规格
//
// 变更云桌面规格，只支持变更CPU和内存，不支持变更磁盘，不支持同规格互相变更。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ResizeDesktop(request *model.ResizeDesktopRequest) (*model.ResizeDesktopResponse, error) {
	requestDef := GenReqDefForResizeDesktop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeDesktopResponse), nil
	}
}

// ResizeDesktopInvoker 变更规格
func (c *WorkspaceClient) ResizeDesktopInvoker(request *model.ResizeDesktopRequest) *ResizeDesktopInvoker {
	requestDef := GenReqDefForResizeDesktop()
	return &ResizeDesktopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDesktopDetail 查询单个桌面详情
//
// 指定桌面Id查询详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopDetail(request *model.ShowDesktopDetailRequest) (*model.ShowDesktopDetailResponse, error) {
	requestDef := GenReqDefForShowDesktopDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopDetailResponse), nil
	}
}

// ShowDesktopDetailInvoker 查询单个桌面详情
func (c *WorkspaceClient) ShowDesktopDetailInvoker(request *model.ShowDesktopDetailRequest) *ShowDesktopDetailInvoker {
	requestDef := GenReqDefForShowDesktopDetail()
	return &ShowDesktopDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImages 查询产品镜像列表
//
// 该接口用于查询云桌面支持的产品镜像列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListImages(request *model.ListImagesRequest) (*model.ListImagesResponse, error) {
	requestDef := GenReqDefForListImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImagesResponse), nil
	}
}

// ListImagesInvoker 查询产品镜像列表
func (c *WorkspaceClient) ListImagesInvoker(request *model.ListImagesRequest) *ListImagesInvoker {
	requestDef := GenReqDefForListImages()
	return &ListImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListItaSubJobs 子任务查询
//
// 该接口用于查询异步任务执行情况，比如查询创建桌面的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListItaSubJobs(request *model.ListItaSubJobsRequest) (*model.ListItaSubJobsResponse, error) {
	requestDef := GenReqDefForListItaSubJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListItaSubJobsResponse), nil
	}
}

// ListItaSubJobsInvoker 子任务查询
func (c *WorkspaceClient) ListItaSubJobsInvoker(request *model.ListItaSubJobsRequest) *ListItaSubJobsInvoker {
	requestDef := GenReqDefForListItaSubJobs()
	return &ListItaSubJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品套餐列表
//
// 该接口用于查询云桌面支持的产品套餐列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListProducts(request *model.ListProductsRequest) (*model.ListProductsResponse, error) {
	requestDef := GenReqDefForListProducts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProductsResponse), nil
	}
}

// ListProductsInvoker 查询产品套餐列表
func (c *WorkspaceClient) ListProductsInvoker(request *model.ListProductsRequest) *ListProductsInvoker {
	requestDef := GenReqDefForListProducts()
	return &ListProductsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopUser 创建用户
//
// 该接口用于创建桌面用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) CreateDesktopUser(request *model.CreateDesktopUserRequest) (*model.CreateDesktopUserResponse, error) {
	requestDef := GenReqDefForCreateDesktopUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesktopUserResponse), nil
	}
}

// CreateDesktopUserInvoker 创建用户
func (c *WorkspaceClient) CreateDesktopUserInvoker(request *model.CreateDesktopUserRequest) *CreateDesktopUserInvoker {
	requestDef := GenReqDefForCreateDesktopUser()
	return &CreateDesktopUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除指定用户
//
// 删除指定的桌面用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除指定用户
func (c *WorkspaceClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserDetail 查询用户详情信息
//
// 该接口用于查询指定的桌面用户详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListUserDetail(request *model.ListUserDetailRequest) (*model.ListUserDetailResponse, error) {
	requestDef := GenReqDefForListUserDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserDetailResponse), nil
	}
}

// ListUserDetailInvoker 查询用户详情信息
func (c *WorkspaceClient) ListUserDetailInvoker(request *model.ListUserDetailRequest) *ListUserDetailInvoker {
	requestDef := GenReqDefForListUserDetail()
	return &ListUserDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 查询用户列表
//
// 该接口用于查询桌面用户列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 查询用户列表
func (c *WorkspaceClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserInfo 修改用户信息
//
// 该接口用于修改桌面用户信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) UpdateUserInfo(request *model.UpdateUserInfoRequest) (*model.UpdateUserInfoResponse, error) {
	requestDef := GenReqDefForUpdateUserInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserInfoResponse), nil
	}
}

// UpdateUserInfoInvoker 修改用户信息
func (c *WorkspaceClient) UpdateUserInfoInvoker(request *model.UpdateUserInfoRequest) *UpdateUserInfoInvoker {
	requestDef := GenReqDefForUpdateUserInfo()
	return &UpdateUserInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddVolumes 增加桌面磁盘
//
// 增加桌面磁盘。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) AddVolumes(request *model.AddVolumesRequest) (*model.AddVolumesResponse, error) {
	requestDef := GenReqDefForAddVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddVolumesResponse), nil
	}
}

// AddVolumesInvoker 增加桌面磁盘
func (c *WorkspaceClient) AddVolumesInvoker(request *model.AddVolumesRequest) *AddVolumesInvoker {
	requestDef := GenReqDefForAddVolumes()
	return &AddVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandVolumes 扩容桌面磁盘
//
// 扩容桌面磁盘。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ExpandVolumes(request *model.ExpandVolumesRequest) (*model.ExpandVolumesResponse, error) {
	requestDef := GenReqDefForExpandVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandVolumesResponse), nil
	}
}

// ExpandVolumesInvoker 扩容桌面磁盘
func (c *WorkspaceClient) ExpandVolumesInvoker(request *model.ExpandVolumesRequest) *ExpandVolumesInvoker {
	requestDef := GenReqDefForExpandVolumes()
	return &ExpandVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ApplyWorkspace 开通云办公服务
//
// 该接口用于开通云办公服务。
//
// 作为异步接口，调用成功说明云办公服务后台收到了开通请求，但服务是否开通成功需要通过任务查询接口(GET /v2/{project_id}/workspace-sub-jobs)查询该任务的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ApplyWorkspace(request *model.ApplyWorkspaceRequest) (*model.ApplyWorkspaceResponse, error) {
	requestDef := GenReqDefForApplyWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyWorkspaceResponse), nil
	}
}

// ApplyWorkspaceInvoker 开通云办公服务
func (c *WorkspaceClient) ApplyWorkspaceInvoker(request *model.ApplyWorkspaceRequest) *ApplyWorkspaceInvoker {
	requestDef := GenReqDefForApplyWorkspace()
	return &ApplyWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelWorkspace 注销云办公服务
//
// 该接口用于注销云办公服务。注销前请确保当前用户下的云桌面已经删除，注销后无法恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) CancelWorkspace(request *model.CancelWorkspaceRequest) (*model.CancelWorkspaceResponse, error) {
	requestDef := GenReqDefForCancelWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelWorkspaceResponse), nil
	}
}

// CancelWorkspaceInvoker 注销云办公服务
func (c *WorkspaceClient) CancelWorkspaceInvoker(request *model.CancelWorkspaceRequest) *CancelWorkspaceInvoker {
	requestDef := GenReqDefForCancelWorkspace()
	return &CancelWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 查询云办公服务详情
//
// 该接口用于查询云办公服务详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 查询云办公服务详情
func (c *WorkspaceClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 修改云办公服务属性
//
// 该接口目前支持修改云办公服务属性。单次请求仅支持修改一种属性类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *WorkspaceClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 修改云办公服务属性
func (c *WorkspaceClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
