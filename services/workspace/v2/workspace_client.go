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

// BatchDeleteAccessPolicies 删除接入策略
//
// 该接口用于删除指定接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteAccessPolicies(request *model.BatchDeleteAccessPoliciesRequest) (*model.BatchDeleteAccessPoliciesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAccessPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAccessPoliciesResponse), nil
	}
}

// BatchDeleteAccessPoliciesInvoker 删除接入策略
func (c *WorkspaceClient) BatchDeleteAccessPoliciesInvoker(request *model.BatchDeleteAccessPoliciesRequest) *BatchDeleteAccessPoliciesInvoker {
	requestDef := GenReqDefForBatchDeleteAccessPolicies()
	return &BatchDeleteAccessPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessPolicy 创建接入策略
//
// 该接口用于创建接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateAccessPolicy(request *model.CreateAccessPolicyRequest) (*model.CreateAccessPolicyResponse, error) {
	requestDef := GenReqDefForCreateAccessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessPolicyResponse), nil
	}
}

// CreateAccessPolicyInvoker 创建接入策略
func (c *WorkspaceClient) CreateAccessPolicyInvoker(request *model.CreateAccessPolicyRequest) *CreateAccessPolicyInvoker {
	requestDef := GenReqDefForCreateAccessPolicy()
	return &CreateAccessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPolicies 查询接入策略
//
// 该接口用于查询接入策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAccessPolicies(request *model.ListAccessPoliciesRequest) (*model.ListAccessPoliciesResponse, error) {
	requestDef := GenReqDefForListAccessPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPoliciesResponse), nil
	}
}

// ListAccessPoliciesInvoker 查询接入策略
func (c *WorkspaceClient) ListAccessPoliciesInvoker(request *model.ListAccessPoliciesRequest) *ListAccessPoliciesInvoker {
	requestDef := GenReqDefForListAccessPolicies()
	return &ListAccessPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessPolicyObjects 查询指定接入策略的应用对象
//
// 该接口用于查询指定接入策略的应用对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListAccessPolicyObjects(request *model.ListAccessPolicyObjectsRequest) (*model.ListAccessPolicyObjectsResponse, error) {
	requestDef := GenReqDefForListAccessPolicyObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessPolicyObjectsResponse), nil
	}
}

// ListAccessPolicyObjectsInvoker 查询指定接入策略的应用对象
func (c *WorkspaceClient) ListAccessPolicyObjectsInvoker(request *model.ListAccessPolicyObjectsRequest) *ListAccessPolicyObjectsInvoker {
	requestDef := GenReqDefForListAccessPolicyObjects()
	return &ListAccessPolicyObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessPolicyObjects 更新指定接入策略的应用对象
//
// 该接口用于更新指定接入策略的应用对象。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAccessPolicyObjects(request *model.UpdateAccessPolicyObjectsRequest) (*model.UpdateAccessPolicyObjectsResponse, error) {
	requestDef := GenReqDefForUpdateAccessPolicyObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessPolicyObjectsResponse), nil
	}
}

// UpdateAccessPolicyObjectsInvoker 更新指定接入策略的应用对象
func (c *WorkspaceClient) UpdateAccessPolicyObjectsInvoker(request *model.UpdateAccessPolicyObjectsRequest) *UpdateAccessPolicyObjectsInvoker {
	requestDef := GenReqDefForUpdateAccessPolicyObjects()
	return &UpdateAccessPolicyObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssistAuthConfig 查询辅助认证配置
//
// 查询辅助认证的配置信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowAssistAuthConfig(request *model.ShowAssistAuthConfigRequest) (*model.ShowAssistAuthConfigResponse, error) {
	requestDef := GenReqDefForShowAssistAuthConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssistAuthConfigResponse), nil
	}
}

// ShowAssistAuthConfigInvoker 查询辅助认证配置
func (c *WorkspaceClient) ShowAssistAuthConfigInvoker(request *model.ShowAssistAuthConfigRequest) *ShowAssistAuthConfigInvoker {
	requestDef := GenReqDefForShowAssistAuthConfig()
	return &ShowAssistAuthConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAssistAuthMethodConfig 更新辅助认证策略配置
//
// 更新辅助认证策略配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateAssistAuthMethodConfig(request *model.UpdateAssistAuthMethodConfigRequest) (*model.UpdateAssistAuthMethodConfigResponse, error) {
	requestDef := GenReqDefForUpdateAssistAuthMethodConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAssistAuthMethodConfigResponse), nil
	}
}

// UpdateAssistAuthMethodConfigInvoker 更新辅助认证策略配置
func (c *WorkspaceClient) UpdateAssistAuthMethodConfigInvoker(request *model.UpdateAssistAuthMethodConfigRequest) *UpdateAssistAuthMethodConfigInvoker {
	requestDef := GenReqDefForUpdateAssistAuthMethodConfig()
	return &UpdateAssistAuthMethodConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailabilityZones 查询可用分区列表
//
// 该接口用于查询云桌面支持的可用分区列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// BatchLogoffDesktops 批量注销桌面
//
// 批量注销桌面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchLogoffDesktops(request *model.BatchLogoffDesktopsRequest) (*model.BatchLogoffDesktopsResponse, error) {
	requestDef := GenReqDefForBatchLogoffDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchLogoffDesktopsResponse), nil
	}
}

// BatchLogoffDesktopsInvoker 批量注销桌面
func (c *WorkspaceClient) BatchLogoffDesktopsInvoker(request *model.BatchLogoffDesktopsRequest) *BatchLogoffDesktopsInvoker {
	requestDef := GenReqDefForBatchLogoffDesktops()
	return &BatchLogoffDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebuildDesktopsSystemDisk 重建桌面
//
// 批量重建桌面系统盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchRebuildDesktopsSystemDisk(request *model.BatchRebuildDesktopsSystemDiskRequest) (*model.BatchRebuildDesktopsSystemDiskResponse, error) {
	requestDef := GenReqDefForBatchRebuildDesktopsSystemDisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebuildDesktopsSystemDiskResponse), nil
	}
}

// BatchRebuildDesktopsSystemDiskInvoker 重建桌面
func (c *WorkspaceClient) BatchRebuildDesktopsSystemDiskInvoker(request *model.BatchRebuildDesktopsSystemDiskRequest) *BatchRebuildDesktopsSystemDiskInvoker {
	requestDef := GenReqDefForBatchRebuildDesktopsSystemDisk()
	return &BatchRebuildDesktopsSystemDiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRunDesktops 操作桌面
//
// 批量操作桌面，用于批量开机、关机和重启。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ChangeDesktopNetwork 切换桌面网络
//
// 切换桌面vpc、子网、ip、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeDesktopNetwork(request *model.ChangeDesktopNetworkRequest) (*model.ChangeDesktopNetworkResponse, error) {
	requestDef := GenReqDefForChangeDesktopNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDesktopNetworkResponse), nil
	}
}

// ChangeDesktopNetworkInvoker 切换桌面网络
func (c *WorkspaceClient) ChangeDesktopNetworkInvoker(request *model.ChangeDesktopNetworkRequest) *ChangeDesktopNetworkInvoker {
	requestDef := GenReqDefForChangeDesktopNetwork()
	return &ChangeDesktopNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktop 创建桌面
//
// 创建桌面，并将此桌面分配给用户，当桌面创建成功后用户可以登录使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowDesktopNetwork 查询桌面网络
//
// 查询桌面vpc、子网、privateIp、EIP、安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowDesktopNetwork(request *model.ShowDesktopNetworkRequest) (*model.ShowDesktopNetworkResponse, error) {
	requestDef := GenReqDefForShowDesktopNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDesktopNetworkResponse), nil
	}
}

// ShowDesktopNetworkInvoker 查询桌面网络
func (c *WorkspaceClient) ShowDesktopNetworkInvoker(request *model.ShowDesktopNetworkRequest) *ShowDesktopNetworkInvoker {
	requestDef := GenReqDefForShowDesktopNetwork()
	return &ShowDesktopNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUnusedDesktops 查询在指定时间段未使用的桌面
//
// 查询在指定时间段未使用的桌面。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUnusedDesktops(request *model.ListUnusedDesktopsRequest) (*model.ListUnusedDesktopsResponse, error) {
	requestDef := GenReqDefForListUnusedDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUnusedDesktopsResponse), nil
	}
}

// ListUnusedDesktopsInvoker 查询在指定时间段未使用的桌面
func (c *WorkspaceClient) ListUnusedDesktopsInvoker(request *model.ListUnusedDesktopsRequest) *ListUnusedDesktopsInvoker {
	requestDef := GenReqDefForListUnusedDesktops()
	return &ListUnusedDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsedDesktopInfo 查询使用桌面的时长
//
// 查询使用桌面的时长。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUsedDesktopInfo(request *model.ListUsedDesktopInfoRequest) (*model.ListUsedDesktopInfoResponse, error) {
	requestDef := GenReqDefForListUsedDesktopInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsedDesktopInfoResponse), nil
	}
}

// ListUsedDesktopInfoInvoker 查询使用桌面的时长
func (c *WorkspaceClient) ListUsedDesktopInfoInvoker(request *model.ListUsedDesktopInfoRequest) *ListUsedDesktopInfoInvoker {
	requestDef := GenReqDefForListUsedDesktopInfo()
	return &ListUsedDesktopInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchChangeTags 批量添加删除标签
//
// 为指定桌面批量添加或删除标签。创建时，如果创建的标签已经存在（key相同），则覆盖。删除时，如果删除的标签不存在，默认处理成功
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchChangeTags(request *model.BatchChangeTagsRequest) (*model.BatchChangeTagsResponse, error) {
	requestDef := GenReqDefForBatchChangeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchChangeTagsResponse), nil
	}
}

// BatchChangeTagsInvoker 批量添加删除标签
func (c *WorkspaceClient) BatchChangeTagsInvoker(request *model.BatchChangeTagsRequest) *BatchChangeTagsInvoker {
	requestDef := GenReqDefForBatchChangeTags()
	return &BatchChangeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 创建桌面标签
//
// 该接口用于为桌面创建标签，一个桌面上最多有10个标签。创建时，如果创建的标签已经存在（key相同），则覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 创建桌面标签
func (c *WorkspaceClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除桌面标签
//
// 该接口用于删除桌面标签。删除时，如果删除的标签不存在，默认处理成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除桌面标签
func (c *WorkspaceClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopByTags 使用标签过滤桌面
//
// 使用标签过滤桌面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopByTags(request *model.ListDesktopByTagsRequest) (*model.ListDesktopByTagsResponse, error) {
	requestDef := GenReqDefForListDesktopByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopByTagsResponse), nil
	}
}

// ListDesktopByTagsInvoker 使用标签过滤桌面
func (c *WorkspaceClient) ListDesktopByTagsInvoker(request *model.ListDesktopByTagsRequest) *ListDesktopByTagsInvoker {
	requestDef := GenReqDefForListDesktopByTags()
	return &ListDesktopByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询租户的所有标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *WorkspaceClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTagByDesktopId 查询桌面标签
//
// 查询指定桌面的标签信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowTagByDesktopId(request *model.ShowTagByDesktopIdRequest) (*model.ShowTagByDesktopIdResponse, error) {
	requestDef := GenReqDefForShowTagByDesktopId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagByDesktopIdResponse), nil
	}
}

// ShowTagByDesktopIdInvoker 查询桌面标签
func (c *WorkspaceClient) ShowTagByDesktopIdInvoker(request *model.ShowTagByDesktopIdRequest) *ShowTagByDesktopIdInvoker {
	requestDef := GenReqDefForShowTagByDesktopId()
	return &ShowTagByDesktopIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUserGroups 批量删除用户组
//
// 该接口用于批量删除用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteUserGroups(request *model.BatchDeleteUserGroupsRequest) (*model.BatchDeleteUserGroupsResponse, error) {
	requestDef := GenReqDefForBatchDeleteUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUserGroupsResponse), nil
	}
}

// BatchDeleteUserGroupsInvoker 批量删除用户组
func (c *WorkspaceClient) BatchDeleteUserGroupsInvoker(request *model.BatchDeleteUserGroupsRequest) *BatchDeleteUserGroupsInvoker {
	requestDef := GenReqDefForBatchDeleteUserGroups()
	return &BatchDeleteUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUserGroup 创建用户组
//
// 该接口用于创建用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateUserGroup(request *model.CreateUserGroupRequest) (*model.CreateUserGroupResponse, error) {
	requestDef := GenReqDefForCreateUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUserGroupResponse), nil
	}
}

// CreateUserGroupInvoker 创建用户组
func (c *WorkspaceClient) CreateUserGroupInvoker(request *model.CreateUserGroupRequest) *CreateUserGroupInvoker {
	requestDef := GenReqDefForCreateUserGroup()
	return &CreateUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUserGroup 删除用户组
//
// 删除用户组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteUserGroup(request *model.DeleteUserGroupRequest) (*model.DeleteUserGroupResponse, error) {
	requestDef := GenReqDefForDeleteUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserGroupResponse), nil
	}
}

// DeleteUserGroupInvoker 删除用户组
func (c *WorkspaceClient) DeleteUserGroupInvoker(request *model.DeleteUserGroupRequest) *DeleteUserGroupInvoker {
	requestDef := GenReqDefForDeleteUserGroup()
	return &DeleteUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserGroups 查询用户组列表
//
// 查询用户组列表，支持分页。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUserGroups(request *model.ListUserGroupsRequest) (*model.ListUserGroupsResponse, error) {
	requestDef := GenReqDefForListUserGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserGroupsResponse), nil
	}
}

// ListUserGroupsInvoker 查询用户组列表
func (c *WorkspaceClient) ListUserGroupsInvoker(request *model.ListUserGroupsRequest) *ListUserGroupsInvoker {
	requestDef := GenReqDefForListUserGroups()
	return &ListUserGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsersOfGroup 查询用户组中的用户
//
// 该接口用于查询用户组中的用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListUsersOfGroup(request *model.ListUsersOfGroupRequest) (*model.ListUsersOfGroupResponse, error) {
	requestDef := GenReqDefForListUsersOfGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersOfGroupResponse), nil
	}
}

// ListUsersOfGroupInvoker 查询用户组中的用户
func (c *WorkspaceClient) ListUsersOfGroupInvoker(request *model.ListUsersOfGroupRequest) *ListUsersOfGroupInvoker {
	requestDef := GenReqDefForListUsersOfGroup()
	return &ListUsersOfGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunActionsOnGroup 操作用户组
//
// 操作用户组，如添加用户、删除用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) RunActionsOnGroup(request *model.RunActionsOnGroupRequest) (*model.RunActionsOnGroupResponse, error) {
	requestDef := GenReqDefForRunActionsOnGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunActionsOnGroupResponse), nil
	}
}

// RunActionsOnGroupInvoker 操作用户组
func (c *WorkspaceClient) RunActionsOnGroupInvoker(request *model.RunActionsOnGroupRequest) *RunActionsOnGroupInvoker {
	requestDef := GenReqDefForRunActionsOnGroup()
	return &RunActionsOnGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserGroup 修改用户组信息
//
// 该接口用于修改用户组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateUserGroup(request *model.UpdateUserGroupRequest) (*model.UpdateUserGroupResponse, error) {
	requestDef := GenReqDefForUpdateUserGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserGroupResponse), nil
	}
}

// UpdateUserGroupInvoker 修改用户组信息
func (c *WorkspaceClient) UpdateUserGroupInvoker(request *model.UpdateUserGroupRequest) *UpdateUserGroupInvoker {
	requestDef := GenReqDefForUpdateUserGroup()
	return &UpdateUserGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImages 查询产品镜像列表
//
// 该接口用于查询云桌面支持的产品镜像列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ApplyDesktopsInternet 开通桌面上网功能
//
// 开通桌面上网功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ApplyDesktopsInternet(request *model.ApplyDesktopsInternetRequest) (*model.ApplyDesktopsInternetResponse, error) {
	requestDef := GenReqDefForApplyDesktopsInternet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyDesktopsInternetResponse), nil
	}
}

// ApplyDesktopsInternetInvoker 开通桌面上网功能
func (c *WorkspaceClient) ApplyDesktopsInternetInvoker(request *model.ApplyDesktopsInternetRequest) *ApplyDesktopsInternetInvoker {
	requestDef := GenReqDefForApplyDesktopsInternet()
	return &ApplyDesktopsInternetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateDesktopsEip 桌面绑定EIP
//
// 桌面绑定EIP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) AssociateDesktopsEip(request *model.AssociateDesktopsEipRequest) (*model.AssociateDesktopsEipResponse, error) {
	requestDef := GenReqDefForAssociateDesktopsEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateDesktopsEipResponse), nil
	}
}

// AssociateDesktopsEipInvoker 桌面绑定EIP
func (c *WorkspaceClient) AssociateDesktopsEipInvoker(request *model.AssociateDesktopsEipRequest) *AssociateDesktopsEipInvoker {
	requestDef := GenReqDefForAssociateDesktopsEip()
	return &AssociateDesktopsEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisassociateDesktopsEip 批量桌面解绑EIP
//
// 批量桌面解绑EIP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDisassociateDesktopsEip(request *model.BatchDisassociateDesktopsEipRequest) (*model.BatchDisassociateDesktopsEipResponse, error) {
	requestDef := GenReqDefForBatchDisassociateDesktopsEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisassociateDesktopsEipResponse), nil
	}
}

// BatchDisassociateDesktopsEipInvoker 批量桌面解绑EIP
func (c *WorkspaceClient) BatchDisassociateDesktopsEipInvoker(request *model.BatchDisassociateDesktopsEipRequest) *BatchDisassociateDesktopsEipInvoker {
	requestDef := GenReqDefForBatchDisassociateDesktopsEip()
	return &BatchDisassociateDesktopsEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDesktopsEips 查询已绑定桌面和未绑定的Eip
//
// 查询已绑定桌面和未绑定的Eip。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListDesktopsEips(request *model.ListDesktopsEipsRequest) (*model.ListDesktopsEipsResponse, error) {
	requestDef := GenReqDefForListDesktopsEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDesktopsEipsResponse), nil
	}
}

// ListDesktopsEipsInvoker 查询已绑定桌面和未绑定的Eip
func (c *WorkspaceClient) ListDesktopsEipsInvoker(request *model.ListDesktopsEipsRequest) *ListDesktopsEipsInvoker {
	requestDef := GenReqDefForListDesktopsEips()
	return &ListDesktopsEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProducts 查询产品套餐列表
//
// 该接口用于查询云桌面支持的产品套餐列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowQuotas 查询租户配额
//
// Console Framework和WKSDesktopManager调用该接口查询租户配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询租户配额
func (c *WorkspaceClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTerminalsBindingDesktops 增加终端与桌面绑定配置
//
// 增加终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) CreateTerminalsBindingDesktops(request *model.CreateTerminalsBindingDesktopsRequest) (*model.CreateTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForCreateTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTerminalsBindingDesktopsResponse), nil
	}
}

// CreateTerminalsBindingDesktopsInvoker 增加终端与桌面绑定配置
func (c *WorkspaceClient) CreateTerminalsBindingDesktopsInvoker(request *model.CreateTerminalsBindingDesktopsRequest) *CreateTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForCreateTerminalsBindingDesktops()
	return &CreateTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTerminalsBindingDesktops 删除终端与桌面绑定配置
//
// 删除终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteTerminalsBindingDesktops(request *model.DeleteTerminalsBindingDesktopsRequest) (*model.DeleteTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForDeleteTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTerminalsBindingDesktopsResponse), nil
	}
}

// DeleteTerminalsBindingDesktopsInvoker 删除终端与桌面绑定配置
func (c *WorkspaceClient) DeleteTerminalsBindingDesktopsInvoker(request *model.DeleteTerminalsBindingDesktopsRequest) *DeleteTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForDeleteTerminalsBindingDesktops()
	return &DeleteTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTerminalsBindingDesktops 查询终端与桌面绑定配置列表
//
// 查询终端与桌面绑定配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTerminalsBindingDesktops(request *model.ListTerminalsBindingDesktopsRequest) (*model.ListTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForListTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTerminalsBindingDesktopsResponse), nil
	}
}

// ListTerminalsBindingDesktopsInvoker 查询终端与桌面绑定配置列表
func (c *WorkspaceClient) ListTerminalsBindingDesktopsInvoker(request *model.ListTerminalsBindingDesktopsRequest) *ListTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForListTerminalsBindingDesktops()
	return &ListTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTerminalsBindingDesktopsConfig 查询终端与桌面绑定的开关配置信息
//
// 查询终端与桌面绑定的开关配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListTerminalsBindingDesktopsConfig(request *model.ListTerminalsBindingDesktopsConfigRequest) (*model.ListTerminalsBindingDesktopsConfigResponse, error) {
	requestDef := GenReqDefForListTerminalsBindingDesktopsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTerminalsBindingDesktopsConfigResponse), nil
	}
}

// ListTerminalsBindingDesktopsConfigInvoker 查询终端与桌面绑定的开关配置信息
func (c *WorkspaceClient) ListTerminalsBindingDesktopsConfigInvoker(request *model.ListTerminalsBindingDesktopsConfigRequest) *ListTerminalsBindingDesktopsConfigInvoker {
	requestDef := GenReqDefForListTerminalsBindingDesktopsConfig()
	return &ListTerminalsBindingDesktopsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTerminalsBindingDesktops 修改终端与桌面绑定配置
//
// 修改终端与桌面绑定配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTerminalsBindingDesktops(request *model.UpdateTerminalsBindingDesktopsRequest) (*model.UpdateTerminalsBindingDesktopsResponse, error) {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktops()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTerminalsBindingDesktopsResponse), nil
	}
}

// UpdateTerminalsBindingDesktopsInvoker 修改终端与桌面绑定配置
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsInvoker(request *model.UpdateTerminalsBindingDesktopsRequest) *UpdateTerminalsBindingDesktopsInvoker {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktops()
	return &UpdateTerminalsBindingDesktopsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTerminalsBindingDesktopsConfig 设置终端与桌面绑定的开关配置
//
// 设置终端与桌面绑定的开关配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsConfig(request *model.UpdateTerminalsBindingDesktopsConfigRequest) (*model.UpdateTerminalsBindingDesktopsConfigResponse, error) {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktopsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTerminalsBindingDesktopsConfigResponse), nil
	}
}

// UpdateTerminalsBindingDesktopsConfigInvoker 设置终端与桌面绑定的开关配置
func (c *WorkspaceClient) UpdateTerminalsBindingDesktopsConfigInvoker(request *model.UpdateTerminalsBindingDesktopsConfigRequest) *UpdateTerminalsBindingDesktopsConfigInvoker {
	requestDef := GenReqDefForUpdateTerminalsBindingDesktopsConfig()
	return &UpdateTerminalsBindingDesktopsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateUsers 批量创建用户
//
// 该接口用于批量创建用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchCreateUsers(request *model.BatchCreateUsersRequest) (*model.BatchCreateUsersResponse, error) {
	requestDef := GenReqDefForBatchCreateUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateUsersResponse), nil
	}
}

// BatchCreateUsersInvoker 批量创建用户
func (c *WorkspaceClient) BatchCreateUsersInvoker(request *model.BatchCreateUsersRequest) *BatchCreateUsersInvoker {
	requestDef := GenReqDefForBatchCreateUsers()
	return &BatchCreateUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteOtpDevices 解绑OTP设备
//
// 该接口用于解绑用户的OTP设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) BatchDeleteOtpDevices(request *model.BatchDeleteOtpDevicesRequest) (*model.BatchDeleteOtpDevicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteOtpDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteOtpDevicesResponse), nil
	}
}

// BatchDeleteOtpDevicesInvoker 解绑OTP设备
func (c *WorkspaceClient) BatchDeleteOtpDevicesInvoker(request *model.BatchDeleteOtpDevicesRequest) *BatchDeleteOtpDevicesInvoker {
	requestDef := GenReqDefForBatchDeleteOtpDevices()
	return &BatchDeleteOtpDevicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeUserStatus 操作用户
//
// 该接口用于操作用户，包含三种操作：锁定、解锁和重置密码（重置密码建议使用/v2/{project_id}/users/{user_id}/random-password接口，在没有通知方式的情况下必须使用/v2/{project_id}/users/{user_id}/random-password接口）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ChangeUserStatus(request *model.ChangeUserStatusRequest) (*model.ChangeUserStatusResponse, error) {
	requestDef := GenReqDefForChangeUserStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeUserStatusResponse), nil
	}
}

// ChangeUserStatusInvoker 操作用户
func (c *WorkspaceClient) ChangeUserStatusInvoker(request *model.ChangeUserStatusRequest) *ChangeUserStatusInvoker {
	requestDef := GenReqDefForChangeUserStatus()
	return &ChangeUserStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesktopUser 创建用户
//
// 该接口用于创建桌面用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListOtpDevicesByUserId 查询OTP设备
//
// 该接口用于查询相应用户下面的OTP设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ListOtpDevicesByUserId(request *model.ListOtpDevicesByUserIdRequest) (*model.ListOtpDevicesByUserIdResponse, error) {
	requestDef := GenReqDefForListOtpDevicesByUserId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListOtpDevicesByUserIdResponse), nil
	}
}

// ListOtpDevicesByUserIdInvoker 查询OTP设备
func (c *WorkspaceClient) ListOtpDevicesByUserIdInvoker(request *model.ListOtpDevicesByUserIdRequest) *ListOtpDevicesByUserIdInvoker {
	requestDef := GenReqDefForListOtpDevicesByUserId()
	return &ListOtpDevicesByUserIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserDetail 查询用户详情信息
//
// 该接口用于查询指定的桌面用户详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ResetRandomPassword 给用户重置随机密码
//
// 该接口用于给用户重置一个密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ResetRandomPassword(request *model.ResetRandomPasswordRequest) (*model.ResetRandomPasswordResponse, error) {
	requestDef := GenReqDefForResetRandomPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetRandomPasswordResponse), nil
	}
}

// ResetRandomPasswordInvoker 给用户重置随机密码
func (c *WorkspaceClient) ResetRandomPasswordInvoker(request *model.ResetRandomPasswordRequest) *ResetRandomPasswordInvoker {
	requestDef := GenReqDefForResetRandomPassword()
	return &ResetRandomPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUserInfo 修改用户信息
//
// 该接口用于修改桌面用户信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteDesktopVolumes 删除桌面数据盘
//
// 删除桌面数据盘，删除后无法恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) DeleteDesktopVolumes(request *model.DeleteDesktopVolumesRequest) (*model.DeleteDesktopVolumesResponse, error) {
	requestDef := GenReqDefForDeleteDesktopVolumes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesktopVolumesResponse), nil
	}
}

// DeleteDesktopVolumesInvoker 删除桌面数据盘
func (c *WorkspaceClient) DeleteDesktopVolumesInvoker(request *model.DeleteDesktopVolumesRequest) *DeleteDesktopVolumesInvoker {
	requestDef := GenReqDefForDeleteDesktopVolumes()
	return &DeleteDesktopVolumesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandVolumes 扩容桌面磁盘
//
// 扩容桌面磁盘。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowWorkspaceLock 查询云办公服务是否被锁定
//
// 查询云办公服务是否被锁定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) ShowWorkspaceLock(request *model.ShowWorkspaceLockRequest) (*model.ShowWorkspaceLockResponse, error) {
	requestDef := GenReqDefForShowWorkspaceLock()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceLockResponse), nil
	}
}

// ShowWorkspaceLockInvoker 查询云办公服务是否被锁定
func (c *WorkspaceClient) ShowWorkspaceLockInvoker(request *model.ShowWorkspaceLockRequest) *ShowWorkspaceLockInvoker {
	requestDef := GenReqDefForShowWorkspaceLock()
	return &ShowWorkspaceLockInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnlockWorkspace 解除云办公服务锁定状态
//
// 该接口目前支持解除云办公服务锁定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *WorkspaceClient) UnlockWorkspace(request *model.UnlockWorkspaceRequest) (*model.UnlockWorkspaceResponse, error) {
	requestDef := GenReqDefForUnlockWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnlockWorkspaceResponse), nil
	}
}

// UnlockWorkspaceInvoker 解除云办公服务锁定状态
func (c *WorkspaceClient) UnlockWorkspaceInvoker(request *model.UnlockWorkspaceRequest) *UnlockWorkspaceInvoker {
	requestDef := GenReqDefForUnlockWorkspace()
	return &UnlockWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 修改云办公服务属性
//
// 该接口目前支持修改云办公服务属性。单次请求仅支持修改一种属性类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
