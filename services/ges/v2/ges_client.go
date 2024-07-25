package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ges/v2/model"
)

type GesClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGesClient(hcClient *httpclient.HcHttpClient) *GesClient {
	return &GesClient{HcClient: hcClient}
}

func GesClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AttachEip2 绑定EIP
//
// 可以通过绑定弹性公网IP（简称EIP）访问GES服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) AttachEip2(request *model.AttachEip2Request) (*model.AttachEip2Response, error) {
	requestDef := GenReqDefForAttachEip2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEip2Response), nil
	}
}

// AttachEip2Invoker 绑定EIP
func (c *GesClient) AttachEip2Invoker(request *model.AttachEip2Request) *AttachEip2Invoker {
	requestDef := GenReqDefForAttachEip2()
	return &AttachEip2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSecurityGroup 切换安全组
//
// 该接口可以在图创建成功后，修改图的安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ChangeSecurityGroup(request *model.ChangeSecurityGroupRequest) (*model.ChangeSecurityGroupResponse, error) {
	requestDef := GenReqDefForChangeSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSecurityGroupResponse), nil
	}
}

// ChangeSecurityGroupInvoker 切换安全组
func (c *GesClient) ChangeSecurityGroupInvoker(request *model.ChangeSecurityGroupRequest) *ChangeSecurityGroupInvoker {
	requestDef := GenReqDefForChangeSecurityGroup()
	return &ChangeSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClearGraph2 清空图
//
// 清空图中所有数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ClearGraph2(request *model.ClearGraph2Request) (*model.ClearGraph2Response, error) {
	requestDef := GenReqDefForClearGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearGraph2Response), nil
	}
}

// ClearGraph2Invoker 清空图
func (c *GesClient) ClearGraph2Invoker(request *model.ClearGraph2Request) *ClearGraph2Invoker {
	requestDef := GenReqDefForClearGraph2()
	return &ClearGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBackup2 新增备份
//
// 新增备份。当前图数据出现错误或故障时，可以启动备份图进行恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateBackup2(request *model.CreateBackup2Request) (*model.CreateBackup2Response, error) {
	requestDef := GenReqDefForCreateBackup2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackup2Response), nil
	}
}

// CreateBackup2Invoker 新增备份
func (c *GesClient) CreateBackup2Invoker(request *model.CreateBackup2Request) *CreateBackup2Invoker {
	requestDef := GenReqDefForCreateBackup2()
	return &CreateBackup2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGraph2 创建图
//
// 创建一个图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateGraph2(request *model.CreateGraph2Request) (*model.CreateGraph2Response, error) {
	requestDef := GenReqDefForCreateGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGraph2Response), nil
	}
}

// CreateGraph2Invoker 创建图
func (c *GesClient) CreateGraph2Invoker(request *model.CreateGraph2Request) *CreateGraph2Invoker {
	requestDef := GenReqDefForCreateGraph2()
	return &CreateGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMetadata2 新增元数据
//
// 新增元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateMetadata2(request *model.CreateMetadata2Request) (*model.CreateMetadata2Response, error) {
	requestDef := GenReqDefForCreateMetadata2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetadata2Response), nil
	}
}

// CreateMetadata2Invoker 新增元数据
func (c *GesClient) CreateMetadata2Invoker(request *model.CreateMetadata2Request) *CreateMetadata2Invoker {
	requestDef := GenReqDefForCreateMetadata2()
	return &CreateMetadata2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup2 删除备份
//
// 删除备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteBackup2(request *model.DeleteBackup2Request) (*model.DeleteBackup2Response, error) {
	requestDef := GenReqDefForDeleteBackup2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackup2Response), nil
	}
}

// DeleteBackup2Invoker 删除备份
func (c *GesClient) DeleteBackup2Invoker(request *model.DeleteBackup2Request) *DeleteBackup2Invoker {
	requestDef := GenReqDefForDeleteBackup2()
	return &DeleteBackup2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGraph2 删除图
//
// 删除一个图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteGraph2(request *model.DeleteGraph2Request) (*model.DeleteGraph2Response, error) {
	requestDef := GenReqDefForDeleteGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGraph2Response), nil
	}
}

// DeleteGraph2Invoker 删除图
func (c *GesClient) DeleteGraph2Invoker(request *model.DeleteGraph2Request) *DeleteGraph2Invoker {
	requestDef := GenReqDefForDeleteGraph2()
	return &DeleteGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMetadata2 删除元数据
//
// 删除元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteMetadata2(request *model.DeleteMetadata2Request) (*model.DeleteMetadata2Response, error) {
	requestDef := GenReqDefForDeleteMetadata2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMetadata2Response), nil
	}
}

// DeleteMetadata2Invoker 删除元数据
func (c *GesClient) DeleteMetadata2Invoker(request *model.DeleteMetadata2Request) *DeleteMetadata2Invoker {
	requestDef := GenReqDefForDeleteMetadata2()
	return &DeleteMetadata2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachEip2 解绑EIP
//
// 当无需继续使用EIP时，您可通过解绑EIP来释放网络资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DetachEip2(request *model.DetachEip2Request) (*model.DetachEip2Response, error) {
	requestDef := GenReqDefForDetachEip2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachEip2Response), nil
	}
}

// DetachEip2Invoker 解绑EIP
func (c *GesClient) DetachEip2Invoker(request *model.DetachEip2Request) *DetachEip2Invoker {
	requestDef := GenReqDefForDetachEip2()
	return &DetachEip2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandGraph2 扩副本
//
// 扩副本能力允许动态扩容多个从节点，扩容的从节点可以处理读请求，从而提高读请求性能。
// &gt; 1.一万边和百亿边规格的图暂不支持扩副本。
// 2.进行扩副本操作后，不支持变更图规格操作。
// 3.如果要对图进行扩容和扩副本两个操作，需要您先进行变更图规格操作，再进行扩副本操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ExpandGraph2(request *model.ExpandGraph2Request) (*model.ExpandGraph2Response, error) {
	requestDef := GenReqDefForExpandGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandGraph2Response), nil
	}
}

// ExpandGraph2Invoker 扩副本
func (c *GesClient) ExpandGraph2Invoker(request *model.ExpandGraph2Request) *ExpandGraph2Invoker {
	requestDef := GenReqDefForExpandGraph2()
	return &ExpandGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportBackup2 导出备份
//
// 导出备份
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ExportBackup2(request *model.ExportBackup2Request) (*model.ExportBackup2Response, error) {
	requestDef := GenReqDefForExportBackup2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportBackup2Response), nil
	}
}

// ExportBackup2Invoker 导出备份
func (c *GesClient) ExportBackup2Invoker(request *model.ExportBackup2Request) *ExportBackup2Invoker {
	requestDef := GenReqDefForExportBackup2()
	return &ExportBackup2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportGraph2 导出图
//
// 导出图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ExportGraph2(request *model.ExportGraph2Request) (*model.ExportGraph2Response, error) {
	requestDef := GenReqDefForExportGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportGraph2Response), nil
	}
}

// ExportGraph2Invoker 导出图
func (c *GesClient) ExportGraph2Invoker(request *model.ExportGraph2Request) *ExportGraph2Invoker {
	requestDef := GenReqDefForExportGraph2()
	return &ExportGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportBackup2 导入备份
//
// 导入备份
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ImportBackup2(request *model.ImportBackup2Request) (*model.ImportBackup2Response, error) {
	requestDef := GenReqDefForImportBackup2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportBackup2Response), nil
	}
}

// ImportBackup2Invoker 导入备份
func (c *GesClient) ImportBackup2Invoker(request *model.ImportBackup2Request) *ImportBackup2Invoker {
	requestDef := GenReqDefForImportBackup2()
	return &ImportBackup2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportGraph2 增量导入图
//
// 增量导入图数据。
// &gt; 为防止系统重启时，不能正常恢复导入图数据，建议在使用图期间，不要删除存储在OBS中的数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ImportGraph2(request *model.ImportGraph2Request) (*model.ImportGraph2Response, error) {
	requestDef := GenReqDefForImportGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportGraph2Response), nil
	}
}

// ImportGraph2Invoker 增量导入图
func (c *GesClient) ImportGraph2Invoker(request *model.ImportGraph2Request) *ImportGraph2Invoker {
	requestDef := GenReqDefForImportGraph2()
	return &ImportGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups2 查看所有备份列表
//
// 查询备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListBackups2(request *model.ListBackups2Request) (*model.ListBackups2Response, error) {
	requestDef := GenReqDefForListBackups2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackups2Response), nil
	}
}

// ListBackups2Invoker 查看所有备份列表
func (c *GesClient) ListBackups2Invoker(request *model.ListBackups2Request) *ListBackups2Invoker {
	requestDef := GenReqDefForListBackups2()
	return &ListBackups2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGraphBackups2 查看某个图的备份列表
//
// 查询某个图下的备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListGraphBackups2(request *model.ListGraphBackups2Request) (*model.ListGraphBackups2Response, error) {
	requestDef := GenReqDefForListGraphBackups2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphBackups2Response), nil
	}
}

// ListGraphBackups2Invoker 查看某个图的备份列表
func (c *GesClient) ListGraphBackups2Invoker(request *model.ListGraphBackups2Request) *ListGraphBackups2Invoker {
	requestDef := GenReqDefForListGraphBackups2()
	return &ListGraphBackups2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGraphs2 查询图列表
//
// 查询当前租户所有的图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListGraphs2(request *model.ListGraphs2Request) (*model.ListGraphs2Response, error) {
	requestDef := GenReqDefForListGraphs2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphs2Response), nil
	}
}

// ListGraphs2Invoker 查询图列表
func (c *GesClient) ListGraphs2Invoker(request *model.ListGraphs2Request) *ListGraphs2Invoker {
	requestDef := GenReqDefForListGraphs2()
	return &ListGraphs2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs2 查询任务中心
//
// 查询管理面任务中心。当前创建图、关闭图、启动图、删除图、增加备份、导入图、导出图、升级图等操作为异步任务，该API用于查询这些任务的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListJobs2(request *model.ListJobs2Request) (*model.ListJobs2Response, error) {
	requestDef := GenReqDefForListJobs2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobs2Response), nil
	}
}

// ListJobs2Invoker 查询任务中心
func (c *GesClient) ListJobs2Invoker(request *model.ListJobs2Request) *ListJobs2Invoker {
	requestDef := GenReqDefForListJobs2()
	return &ListJobs2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetadatas2 查询元数据列表
//
// 查询元数据列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListMetadatas2(request *model.ListMetadatas2Request) (*model.ListMetadatas2Response, error) {
	requestDef := GenReqDefForListMetadatas2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetadatas2Response), nil
	}
}

// ListMetadatas2Invoker 查询元数据列表
func (c *GesClient) ListMetadatas2Invoker(request *model.ListMetadatas2Request) *ListMetadatas2Invoker {
	requestDef := GenReqDefForListMetadatas2()
	return &ListMetadatas2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas2 查询配额
//
// 查询租户配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListQuotas2(request *model.ListQuotas2Request) (*model.ListQuotas2Response, error) {
	requestDef := GenReqDefForListQuotas2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotas2Response), nil
	}
}

// ListQuotas2Invoker 查询配额
func (c *GesClient) ListQuotas2Invoker(request *model.ListQuotas2Request) *ListQuotas2Invoker {
	requestDef := GenReqDefForListQuotas2()
	return &ListQuotas2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeGraph2 变更图规格
//
// 变更图规格规格。
// &gt; 变更图规格以后所有索引（复合索引和全文索引）都需要重新创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ResizeGraph2(request *model.ResizeGraph2Request) (*model.ResizeGraph2Response, error) {
	requestDef := GenReqDefForResizeGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeGraph2Response), nil
	}
}

// ResizeGraph2Invoker 变更图规格
func (c *GesClient) ResizeGraph2Invoker(request *model.ResizeGraph2Request) *ResizeGraph2Invoker {
	requestDef := GenReqDefForResizeGraph2()
	return &ResizeGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartGraph2 强制重启图
//
// 强制启动一个图。针对导入、导出 、运行中 、清空中的图。强制重启图，会将该图执行中的异步任务变为失败，然后停止图、启动图到运行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) RestartGraph2(request *model.RestartGraph2Request) (*model.RestartGraph2Response, error) {
	requestDef := GenReqDefForRestartGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartGraph2Response), nil
	}
}

// RestartGraph2Invoker 强制重启图
func (c *GesClient) RestartGraph2Invoker(request *model.RestartGraph2Request) *RestartGraph2Invoker {
	requestDef := GenReqDefForRestartGraph2()
	return &RestartGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupDownloadLink 获取备份下载链接
//
// 获取备份下载链接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowBackupDownloadLink(request *model.ShowBackupDownloadLinkRequest) (*model.ShowBackupDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowBackupDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

// ShowBackupDownloadLinkInvoker 获取备份下载链接
func (c *GesClient) ShowBackupDownloadLinkInvoker(request *model.ShowBackupDownloadLinkRequest) *ShowBackupDownloadLinkInvoker {
	requestDef := GenReqDefForShowBackupDownloadLink()
	return &ShowBackupDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGraph2 查询图详情
//
// 根据图ID查询某个图详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowGraph2(request *model.ShowGraph2Request) (*model.ShowGraph2Response, error) {
	requestDef := GenReqDefForShowGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGraph2Response), nil
	}
}

// ShowGraph2Invoker 查询图详情
func (c *GesClient) ShowGraph2Invoker(request *model.ShowGraph2Request) *ShowGraph2Invoker {
	requestDef := GenReqDefForShowGraph2()
	return &ShowGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob2 查询Job状态-管理面
//
// 查询Job的执行状态。对创建图、关闭图、启动图、删除图、导入图等异步API命令下发后，会返回jobId，通过jobId查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowJob2(request *model.ShowJob2Request) (*model.ShowJob2Response, error) {
	requestDef := GenReqDefForShowJob2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJob2Response), nil
	}
}

// ShowJob2Invoker 查询Job状态-管理面
func (c *GesClient) ShowJob2Invoker(request *model.ShowJob2Request) *ShowJob2Invoker {
	requestDef := GenReqDefForShowJob2()
	return &ShowJob2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetadata2 查询元数据
//
// 查询某个元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowMetadata2(request *model.ShowMetadata2Request) (*model.ShowMetadata2Response, error) {
	requestDef := GenReqDefForShowMetadata2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetadata2Response), nil
	}
}

// ShowMetadata2Invoker 查询元数据
func (c *GesClient) ShowMetadata2Invoker(request *model.ShowMetadata2Request) *ShowMetadata2Invoker {
	requestDef := GenReqDefForShowMetadata2()
	return &ShowMetadata2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartGraph2 启动图
//
// 启动一个图。暂时不用的图可以先关闭，需要使用时再启动。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) StartGraph2(request *model.StartGraph2Request) (*model.StartGraph2Response, error) {
	requestDef := GenReqDefForStartGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartGraph2Response), nil
	}
}

// StartGraph2Invoker 启动图
func (c *GesClient) StartGraph2Invoker(request *model.StartGraph2Request) *StartGraph2Invoker {
	requestDef := GenReqDefForStartGraph2()
	return &StartGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopGraph2 关闭图
//
// 关闭一个图。如果图创建好了，暂时不用可以先关闭，需要使用时再启用。
// &gt; 处于关闭状态的图不计算实例费用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) StopGraph2(request *model.StopGraph2Request) (*model.StopGraph2Response, error) {
	requestDef := GenReqDefForStopGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopGraph2Response), nil
	}
}

// StopGraph2Invoker 关闭图
func (c *GesClient) StopGraph2Invoker(request *model.StopGraph2Request) *StopGraph2Invoker {
	requestDef := GenReqDefForStopGraph2()
	return &StopGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeGraph2 升级图
//
// 升级图。图引擎服务会定期升级版本，用户可根据需要升级图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) UpgradeGraph2(request *model.UpgradeGraph2Request) (*model.UpgradeGraph2Response, error) {
	requestDef := GenReqDefForUpgradeGraph2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeGraph2Response), nil
	}
}

// UpgradeGraph2Invoker 升级图
func (c *GesClient) UpgradeGraph2Invoker(request *model.UpgradeGraph2Request) *UpgradeGraph2Invoker {
	requestDef := GenReqDefForUpgradeGraph2()
	return &UpgradeGraph2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFromObs2 从OBS导入元数据
//
// 从OBS导入元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) UploadFromObs2(request *model.UploadFromObs2Request) (*model.UploadFromObs2Response, error) {
	requestDef := GenReqDefForUploadFromObs2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFromObs2Response), nil
	}
}

// UploadFromObs2Invoker 从OBS导入元数据
func (c *GesClient) UploadFromObs2Invoker(request *model.UploadFromObs2Request) *UploadFromObs2Invoker {
	requestDef := GenReqDefForUploadFromObs2()
	return &UploadFromObs2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeregisterScenes2 取消订阅场景分析插件
//
// 取消订阅scenes场景应用分析能力，取消订阅后对应scene下的application业务面API将不能使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeregisterScenes2(request *model.DeregisterScenes2Request) (*model.DeregisterScenes2Response, error) {
	requestDef := GenReqDefForDeregisterScenes2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeregisterScenes2Response), nil
	}
}

// DeregisterScenes2Invoker 取消订阅场景分析插件
func (c *GesClient) DeregisterScenes2Invoker(request *model.DeregisterScenes2Request) *DeregisterScenes2Invoker {
	requestDef := GenReqDefForDeregisterScenes2()
	return &DeregisterScenes2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScenes2 查询获取场景应用分析插件
//
// 查询scenes场景下的应用分析能力详情，可以获得对应场景下的application、参数和功能介绍详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListScenes2(request *model.ListScenes2Request) (*model.ListScenes2Response, error) {
	requestDef := GenReqDefForListScenes2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScenes2Response), nil
	}
}

// ListScenes2Invoker 查询获取场景应用分析插件
func (c *GesClient) ListScenes2Invoker(request *model.ListScenes2Request) *ListScenes2Invoker {
	requestDef := GenReqDefForListScenes2()
	return &ListScenes2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterScenes2 订阅场景分析插件
//
// 订阅scenes应用场景分析能力，便于业务面API使用对应功能。
// &gt; 已订阅的不可以重复订阅，需要更新请先取消原有订购，重新订购。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) RegisterScenes2(request *model.RegisterScenes2Request) (*model.RegisterScenes2Response, error) {
	requestDef := GenReqDefForRegisterScenes2()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterScenes2Response), nil
	}
}

// RegisterScenes2Invoker 订阅场景分析插件
func (c *GesClient) RegisterScenes2Invoker(request *model.RegisterScenes2Request) *RegisterScenes2Invoker {
	requestDef := GenReqDefForRegisterScenes2()
	return &RegisterScenes2Invoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
