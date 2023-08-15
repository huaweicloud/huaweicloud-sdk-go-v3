package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/ges/v1/model"
)

type GesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGesClient(hcClient *http_client.HcHttpClient) *GesClient {
	return &GesClient{HcClient: hcClient}
}

func GesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AttachEip 绑定EIP(1.0.6)
//
// 可以通过绑定弹性公网IP（简称EIP）访问GES服务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

// AttachEipInvoker 绑定EIP(1.0.6)
func (c *GesClient) AttachEipInvoker(request *model.AttachEipRequest) *AttachEipInvoker {
	requestDef := GenReqDefForAttachEip()
	return &AttachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClearGraph 清空图(2.1.2)
//
// 清空图中所有数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ClearGraph(request *model.ClearGraphRequest) (*model.ClearGraphResponse, error) {
	requestDef := GenReqDefForClearGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearGraphResponse), nil
	}
}

// ClearGraphInvoker 清空图(2.1.2)
func (c *GesClient) ClearGraphInvoker(request *model.ClearGraphRequest) *ClearGraphInvoker {
	requestDef := GenReqDefForClearGraph()
	return &ClearGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBackup 新增备份(1.0.0)
//
// 新增备份。当前图数据出现错误或故障时，可以启动备份图进行恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateBackup(request *model.CreateBackupRequest) (*model.CreateBackupResponse, error) {
	requestDef := GenReqDefForCreateBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackupResponse), nil
	}
}

// CreateBackupInvoker 新增备份(1.0.0)
func (c *GesClient) CreateBackupInvoker(request *model.CreateBackupRequest) *CreateBackupInvoker {
	requestDef := GenReqDefForCreateBackup()
	return &CreateBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGraph 创建图(2.2.2)
//
// 创建一个图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateGraph(request *model.CreateGraphRequest) (*model.CreateGraphResponse, error) {
	requestDef := GenReqDefForCreateGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGraphResponse), nil
	}
}

// CreateGraphInvoker 创建图(2.2.2)
func (c *GesClient) CreateGraphInvoker(request *model.CreateGraphRequest) *CreateGraphInvoker {
	requestDef := GenReqDefForCreateGraph()
	return &CreateGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMetadata 新增元数据(2.1.18)
//
// 新增元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) CreateMetadata(request *model.CreateMetadataRequest) (*model.CreateMetadataResponse, error) {
	requestDef := GenReqDefForCreateMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetadataResponse), nil
	}
}

// CreateMetadataInvoker 新增元数据(2.1.18)
func (c *GesClient) CreateMetadataInvoker(request *model.CreateMetadataRequest) *CreateMetadataInvoker {
	requestDef := GenReqDefForCreateMetadata()
	return &CreateMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup 删除备份(1.0.0)
//
// 删除备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// DeleteBackupInvoker 删除备份(1.0.0)
func (c *GesClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGraph 删除图(1.0.0)
//
// 删除一个图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteGraph(request *model.DeleteGraphRequest) (*model.DeleteGraphResponse, error) {
	requestDef := GenReqDefForDeleteGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGraphResponse), nil
	}
}

// DeleteGraphInvoker 删除图(1.0.0)
func (c *GesClient) DeleteGraphInvoker(request *model.DeleteGraphRequest) *DeleteGraphInvoker {
	requestDef := GenReqDefForDeleteGraph()
	return &DeleteGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMetadata 删除元数据(1.0.2)
//
// 删除元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DeleteMetadata(request *model.DeleteMetadataRequest) (*model.DeleteMetadataResponse, error) {
	requestDef := GenReqDefForDeleteMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMetadataResponse), nil
	}
}

// DeleteMetadataInvoker 删除元数据(1.0.2)
func (c *GesClient) DeleteMetadataInvoker(request *model.DeleteMetadataRequest) *DeleteMetadataInvoker {
	requestDef := GenReqDefForDeleteMetadata()
	return &DeleteMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachEip 解绑EIP(1.0.6)
//
// 当无需继续使用EIP时，您可通过解绑EIP来释放网络资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) DetachEip(request *model.DetachEipRequest) (*model.DetachEipResponse, error) {
	requestDef := GenReqDefForDetachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachEipResponse), nil
	}
}

// DetachEipInvoker 解绑EIP(1.0.6)
func (c *GesClient) DetachEipInvoker(request *model.DetachEipRequest) *DetachEipInvoker {
	requestDef := GenReqDefForDetachEip()
	return &DetachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandGraph 扩副本(2.2.23)
//
// 扩副本能力允许动态扩容多个从节点，扩容的从节点可以处理读请求，从而提高读请求性能。
// &gt;一万边和百亿边规格的图暂不支持扩副本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ExpandGraph(request *model.ExpandGraphRequest) (*model.ExpandGraphResponse, error) {
	requestDef := GenReqDefForExpandGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandGraphResponse), nil
	}
}

// ExpandGraphInvoker 扩副本(2.2.23)
func (c *GesClient) ExpandGraphInvoker(request *model.ExpandGraphRequest) *ExpandGraphInvoker {
	requestDef := GenReqDefForExpandGraph()
	return &ExpandGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportGraph 导出图(1.0.5)
//
// 导出图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ExportGraph(request *model.ExportGraphRequest) (*model.ExportGraphResponse, error) {
	requestDef := GenReqDefForExportGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportGraphResponse), nil
	}
}

// ExportGraphInvoker 导出图(1.0.5)
func (c *GesClient) ExportGraphInvoker(request *model.ExportGraphRequest) *ExportGraphInvoker {
	requestDef := GenReqDefForExportGraph()
	return &ExportGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportGraph 增量导入图(2.1.14)
//
// 增量导入图数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ImportGraph(request *model.ImportGraphRequest) (*model.ImportGraphResponse, error) {
	requestDef := GenReqDefForImportGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportGraphResponse), nil
	}
}

// ImportGraphInvoker 增量导入图(2.1.14)
func (c *GesClient) ImportGraphInvoker(request *model.ImportGraphRequest) *ImportGraphInvoker {
	requestDef := GenReqDefForImportGraph()
	return &ImportGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups 查看所有备份列表(1.0.0)
//
// 查询备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// ListBackupsInvoker 查看所有备份列表(1.0.0)
func (c *GesClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGraphBackups 查看某个图的备份列表(1.0.0)
//
// 查询某个图下的备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListGraphBackups(request *model.ListGraphBackupsRequest) (*model.ListGraphBackupsResponse, error) {
	requestDef := GenReqDefForListGraphBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphBackupsResponse), nil
	}
}

// ListGraphBackupsInvoker 查看某个图的备份列表(1.0.0)
func (c *GesClient) ListGraphBackupsInvoker(request *model.ListGraphBackupsRequest) *ListGraphBackupsInvoker {
	requestDef := GenReqDefForListGraphBackups()
	return &ListGraphBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGraphMetadatas 查询元数据(1.0.2)
//
// 查询某个图下的元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListGraphMetadatas(request *model.ListGraphMetadatasRequest) (*model.ListGraphMetadatasResponse, error) {
	requestDef := GenReqDefForListGraphMetadatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphMetadatasResponse), nil
	}
}

// ListGraphMetadatasInvoker 查询元数据(1.0.2)
func (c *GesClient) ListGraphMetadatasInvoker(request *model.ListGraphMetadatasRequest) *ListGraphMetadatasInvoker {
	requestDef := GenReqDefForListGraphMetadatas()
	return &ListGraphMetadatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGraphs 查询图列表(2.1.18)
//
// 查询当前租户所有的图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListGraphs(request *model.ListGraphsRequest) (*model.ListGraphsResponse, error) {
	requestDef := GenReqDefForListGraphs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphsResponse), nil
	}
}

// ListGraphsInvoker 查询图列表(2.1.18)
func (c *GesClient) ListGraphsInvoker(request *model.ListGraphsRequest) *ListGraphsInvoker {
	requestDef := GenReqDefForListGraphs()
	return &ListGraphsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询任务中心(1.1.8)
//
// 查询管理面任务中心。当前创建图、关闭图、启动图、删除图、增加备份、导入图、导出图、升级图等操作为异步任务，该API用于查询这些任务的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询任务中心(1.1.8)
func (c *GesClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetadatas 查询元数据列表(1.0.2)
//
// 查询元数据列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListMetadatas(request *model.ListMetadatasRequest) (*model.ListMetadatasResponse, error) {
	requestDef := GenReqDefForListMetadatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetadatasResponse), nil
	}
}

// ListMetadatasInvoker 查询元数据列表(1.0.2)
func (c *GesClient) ListMetadatasInvoker(request *model.ListMetadatasRequest) *ListMetadatasInvoker {
	requestDef := GenReqDefForListMetadatas()
	return &ListMetadatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额(1.0.0)
//
// 查询租户配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额(1.0.0)
func (c *GesClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeGraph 扩容图(2.2.21)
//
// 扩容图规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ResizeGraph(request *model.ResizeGraphRequest) (*model.ResizeGraphResponse, error) {
	requestDef := GenReqDefForResizeGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeGraphResponse), nil
	}
}

// ResizeGraphInvoker 扩容图(2.2.21)
func (c *GesClient) ResizeGraphInvoker(request *model.ResizeGraphRequest) *ResizeGraphInvoker {
	requestDef := GenReqDefForResizeGraph()
	return &ResizeGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartGraph 强制重启图(2.2.21)
//
// 强制启动一个图。针对导入、导出 、运行中 、清空中的图。强制重启图，会将该图执行中的异步任务变为失败，然后停止图、启动图到运行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) RestartGraph(request *model.RestartGraphRequest) (*model.RestartGraphResponse, error) {
	requestDef := GenReqDefForRestartGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartGraphResponse), nil
	}
}

// RestartGraphInvoker 强制重启图(2.2.21)
func (c *GesClient) RestartGraphInvoker(request *model.RestartGraphRequest) *RestartGraphInvoker {
	requestDef := GenReqDefForRestartGraph()
	return &RestartGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGraph 查询图详情(1.0.0)
//
// 根据图ID查询某个图详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowGraph(request *model.ShowGraphRequest) (*model.ShowGraphResponse, error) {
	requestDef := GenReqDefForShowGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGraphResponse), nil
	}
}

// ShowGraphInvoker 查询图详情(1.0.0)
func (c *GesClient) ShowGraphInvoker(request *model.ShowGraphRequest) *ShowGraphInvoker {
	requestDef := GenReqDefForShowGraph()
	return &ShowGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJob 查询Job状态(1.0.0)-管理面
//
// 查询Job的执行状态。对创建图、关闭图、启动图、删除图、导入图等异步API命令下发后，会返回jobId，通过jobId查询任务的执行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// ShowJobInvoker 查询Job状态(1.0.0)-管理面
func (c *GesClient) ShowJobInvoker(request *model.ShowJobRequest) *ShowJobInvoker {
	requestDef := GenReqDefForShowJob()
	return &ShowJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartGraph 启动图(1.0.0)
//
// 启动一个图。暂时不用的图可以先关闭，需要使用时再启动。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) StartGraph(request *model.StartGraphRequest) (*model.StartGraphResponse, error) {
	requestDef := GenReqDefForStartGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartGraphResponse), nil
	}
}

// StartGraphInvoker 启动图(1.0.0)
func (c *GesClient) StartGraphInvoker(request *model.StartGraphRequest) *StartGraphInvoker {
	requestDef := GenReqDefForStartGraph()
	return &StartGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopGraph 关闭图(1.0.0)
//
// 关闭一个图。如果图创建好了，暂时不用可以先关闭，需要使用时再启用。
// &gt;处于关闭状态的图不计算实例费用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) StopGraph(request *model.StopGraphRequest) (*model.StopGraphResponse, error) {
	requestDef := GenReqDefForStopGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopGraphResponse), nil
	}
}

// StopGraphInvoker 关闭图(1.0.0)
func (c *GesClient) StopGraphInvoker(request *model.StopGraphRequest) *StopGraphInvoker {
	requestDef := GenReqDefForStopGraph()
	return &StopGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeGraph 升级图(1.0.5)
//
// 升级图。图引擎服务会定期升级版本，用户可根据需要升级图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) UpgradeGraph(request *model.UpgradeGraphRequest) (*model.UpgradeGraphResponse, error) {
	requestDef := GenReqDefForUpgradeGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeGraphResponse), nil
	}
}

// UpgradeGraphInvoker 升级图(1.0.5)
func (c *GesClient) UpgradeGraphInvoker(request *model.UpgradeGraphRequest) *UpgradeGraphInvoker {
	requestDef := GenReqDefForUpgradeGraph()
	return &UpgradeGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFromObs 从OBS导入元数据(1.0.0)
//
// 从OBS导入元数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GesClient) UploadFromObs(request *model.UploadFromObsRequest) (*model.UploadFromObsResponse, error) {
	requestDef := GenReqDefForUploadFromObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFromObsResponse), nil
	}
}

// UploadFromObsInvoker 从OBS导入元数据(1.0.0)
func (c *GesClient) UploadFromObsInvoker(request *model.UploadFromObsRequest) *UploadFromObsInvoker {
	requestDef := GenReqDefForUploadFromObs()
	return &UploadFromObsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
