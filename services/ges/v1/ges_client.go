package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ges/v1/model"
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

// 绑定EIP(1.0.6)
//
// 可以通过绑定弹性公网IP（简称EIP）访问GES服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

// 清空图(2.1.2)
//
// 清空图中所有数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ClearGraph(request *model.ClearGraphRequest) (*model.ClearGraphResponse, error) {
	requestDef := GenReqDefForClearGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearGraphResponse), nil
	}
}

// 新增备份(1.0.0)
//
// 新增备份。当前图数据出现错误或故障时，可以启动备份图进行恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) CreateBackup(request *model.CreateBackupRequest) (*model.CreateBackupResponse, error) {
	requestDef := GenReqDefForCreateBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBackupResponse), nil
	}
}

// 创建图(2.2.2)
//
// 创建一个图。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) CreateGraph(request *model.CreateGraphRequest) (*model.CreateGraphResponse, error) {
	requestDef := GenReqDefForCreateGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGraphResponse), nil
	}
}

// 新增元数据(2.1.18)
//
// 新增元数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) CreateMetadata(request *model.CreateMetadataRequest) (*model.CreateMetadataResponse, error) {
	requestDef := GenReqDefForCreateMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetadataResponse), nil
	}
}

// 删除备份(1.0.0)
//
// 删除备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// 删除图(1.0.0)
//
// 删除一个图。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) DeleteGraph(request *model.DeleteGraphRequest) (*model.DeleteGraphResponse, error) {
	requestDef := GenReqDefForDeleteGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGraphResponse), nil
	}
}

// 删除元数据(1.0.2)
//
// 删除元数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) DeleteMetadata(request *model.DeleteMetadataRequest) (*model.DeleteMetadataResponse, error) {
	requestDef := GenReqDefForDeleteMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMetadataResponse), nil
	}
}

// 解绑EIP(1.0.6)
//
// 当无需继续使用EIP时，您可通过解绑EIP来释放网络资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) DetachEip(request *model.DetachEipRequest) (*model.DetachEipResponse, error) {
	requestDef := GenReqDefForDetachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachEipResponse), nil
	}
}

// 扩副本(2.2.23)
//
// 扩副本能力允许动态扩容多个从节点，扩容的从节点可以处理读请求，从而提高读请求性能。
// &gt;一万边和百亿边规格的图暂不支持扩副本。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ExpandGraph(request *model.ExpandGraphRequest) (*model.ExpandGraphResponse, error) {
	requestDef := GenReqDefForExpandGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandGraphResponse), nil
	}
}

// 导出图(1.0.5)
//
// 导出图。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ExportGraph(request *model.ExportGraphRequest) (*model.ExportGraphResponse, error) {
	requestDef := GenReqDefForExportGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportGraphResponse), nil
	}
}

// 增量导入图(2.1.14)
//
// 增量导入图数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ImportGraph(request *model.ImportGraphRequest) (*model.ImportGraphResponse, error) {
	requestDef := GenReqDefForImportGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportGraphResponse), nil
	}
}

// 查看所有备份列表(1.0.0)
//
// 查询备份列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// 查看某个图的备份列表(1.0.0)
//
// 查询某个图下的备份列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListGraphBackups(request *model.ListGraphBackupsRequest) (*model.ListGraphBackupsResponse, error) {
	requestDef := GenReqDefForListGraphBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphBackupsResponse), nil
	}
}

// 查询元数据(1.0.2)
//
// 查询某个图下的元数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListGraphMetadatas(request *model.ListGraphMetadatasRequest) (*model.ListGraphMetadatasResponse, error) {
	requestDef := GenReqDefForListGraphMetadatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphMetadatasResponse), nil
	}
}

// 查询图列表(2.1.18)
//
// 查询当前租户所有的图。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListGraphs(request *model.ListGraphsRequest) (*model.ListGraphsResponse, error) {
	requestDef := GenReqDefForListGraphs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGraphsResponse), nil
	}
}

// 查询任务中心(1.1.8)
//
// 查询管理面任务中心。当前创建图、关闭图、启动图、删除图、增加备份、导入图、导出图、升级图等操作为异步任务，该API用于查询这些任务的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// 查询元数据列表(1.0.2)
//
// 查询元数据列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListMetadatas(request *model.ListMetadatasRequest) (*model.ListMetadatasResponse, error) {
	requestDef := GenReqDefForListMetadatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetadatasResponse), nil
	}
}

// 查询配额(1.0.0)
//
// 查询租户配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// 扩容图(2.2.21)
//
// 扩容图规格。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ResizeGraph(request *model.ResizeGraphRequest) (*model.ResizeGraphResponse, error) {
	requestDef := GenReqDefForResizeGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeGraphResponse), nil
	}
}

// 强制重启图(2.2.21)
//
// 强制启动一个图。针对导入、导出 、运行中 、清空中的图。强制重启图，会将该图执行中的异步任务变为失败，然后停止图、启动图到运行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) RestartGraph(request *model.RestartGraphRequest) (*model.RestartGraphResponse, error) {
	requestDef := GenReqDefForRestartGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartGraphResponse), nil
	}
}

// 查询图详情(1.0.0)
//
// 根据图ID查询某个图详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ShowGraph(request *model.ShowGraphRequest) (*model.ShowGraphResponse, error) {
	requestDef := GenReqDefForShowGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGraphResponse), nil
	}
}

// 查询Job状态(1.0.0)-管理面
//
// 查询Job的执行状态。对创建图、关闭图、启动图、删除图、导入图等异步API命令下发后，会返回jobId，通过jobId查询任务的执行状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) ShowJob(request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	requestDef := GenReqDefForShowJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResponse), nil
	}
}

// 启动图(1.0.0)
//
// 启动一个图。暂时不用的图可以先关闭，需要使用时再启动。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) StartGraph(request *model.StartGraphRequest) (*model.StartGraphResponse, error) {
	requestDef := GenReqDefForStartGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartGraphResponse), nil
	}
}

// 关闭图(1.0.0)
//
// 关闭一个图。如果图创建好了，暂时不用可以先关闭，需要使用时再启用。
// &gt;处于关闭状态的图不计算实例费用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) StopGraph(request *model.StopGraphRequest) (*model.StopGraphResponse, error) {
	requestDef := GenReqDefForStopGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopGraphResponse), nil
	}
}

// 升级图(1.0.5)
//
// 升级图。图引擎服务会定期升级版本，用户可根据需要升级图。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) UpgradeGraph(request *model.UpgradeGraphRequest) (*model.UpgradeGraphResponse, error) {
	requestDef := GenReqDefForUpgradeGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeGraphResponse), nil
	}
}

// 从OBS导入元数据(1.0.0)
//
// 从OBS导入元数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GesClient) UploadFromObs(request *model.UploadFromObsRequest) (*model.UploadFromObsResponse, error) {
	requestDef := GenReqDefForUploadFromObs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFromObsResponse), nil
	}
}
