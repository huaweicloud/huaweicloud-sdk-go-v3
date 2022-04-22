package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/css/v1/model"
)

type CssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCssClient(hcClient *http_client.HcHttpClient) *CssClient {
	return &CssClient{HcClient: hcClient}
}

func CssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 设置自动创建快照策略
//
// 该接口用于设置自动创建快照，默认一天创建一个快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateAutoCreatePolicy(request *model.CreateAutoCreatePolicyRequest) (*model.CreateAutoCreatePolicyResponse, error) {
	requestDef := GenReqDefForCreateAutoCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoCreatePolicyResponse), nil
	}
}

// 开启公网访问
//
// 该接口用于开启公网访问。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateBindPublic(request *model.CreateBindPublicRequest) (*model.CreateBindPublicResponse, error) {
	requestDef := GenReqDefForCreateBindPublic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBindPublicResponse), nil
	}
}

// 创建集群
//
// 该接口用于创建集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// 添加指定集群标签
//
// 该接口用于给指定集群添加标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateClustersTags(request *model.CreateClustersTagsRequest) (*model.CreateClustersTagsResponse, error) {
	requestDef := GenReqDefForCreateClustersTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClustersTagsResponse), nil
	}
}

// 加载自定义词库
//
// 该接口用于加载存放于OBS的自定义词库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateLoadIkThesaurus(request *model.CreateLoadIkThesaurusRequest) (*model.CreateLoadIkThesaurusResponse, error) {
	requestDef := GenReqDefForCreateLoadIkThesaurus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLoadIkThesaurusResponse), nil
	}
}

// 备份日志
//
// 该接口用于备份日志。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateLogBackup(request *model.CreateLogBackupRequest) (*model.CreateLogBackupResponse, error) {
	requestDef := GenReqDefForCreateLogBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLogBackupResponse), nil
	}
}

// 手动创建快照
//
// 该接口用于手动创建一个快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

// 删除集群
//
// 此接口用于删除集群。集群删除将释放此集群的所有资源，包括客户数据。为了安全起见，请确保为这个集群创建快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// 删除集群标签
//
// 此接口用于删除集群标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteClustersTags(request *model.DeleteClustersTagsRequest) (*model.DeleteClustersTagsResponse, error) {
	requestDef := GenReqDefForDeleteClustersTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClustersTagsResponse), nil
	}
}

// 删除自定义词库
//
// 该接口用于删除自定义词库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteIkThesaurus(request *model.DeleteIkThesaurusRequest) (*model.DeleteIkThesaurusResponse, error) {
	requestDef := GenReqDefForDeleteIkThesaurus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIkThesaurusResponse), nil
	}
}

// 删除快照
//
// 该接口用于删除快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

// 查询集群列表
//
// 该接口用于查询并显示集群列表以及集群的状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListClustersDetails(request *model.ListClustersDetailsRequest) (*model.ListClustersDetailsResponse, error) {
	requestDef := GenReqDefForListClustersDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersDetailsResponse), nil
	}
}

// 查询所有标签
//
// 该接口用于查询指定region下的所有标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListClustersTags(request *model.ListClustersTagsRequest) (*model.ListClustersTagsResponse, error) {
	requestDef := GenReqDefForListClustersTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersTagsResponse), nil
	}
}

// 获取实例规格列表
//
// 该接口用于查询并显示支持的实例规格对应的ID。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// 查询作业列表
//
// 该接口用于查询具体某个集群的日志任务记录列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListLogsJob(request *model.ListLogsJobRequest) (*model.ListLogsJobResponse, error) {
	requestDef := GenReqDefForListLogsJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLogsJobResponse), nil
	}
}

// 查询快照列表
//
// 该接口用于查询集群的所有快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// 获取参数配置列表
//
// 该接口用于获取参数配置列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListYmls(request *model.ListYmlsRequest) (*model.ListYmlsResponse, error) {
	requestDef := GenReqDefForListYmls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListYmlsResponse), nil
	}
}

// 获取参数配置任务列表
//
// 该接口用于获取参数配置任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListYmlsJob(request *model.ListYmlsJobRequest) (*model.ListYmlsJobResponse, error) {
	requestDef := GenReqDefForListYmlsJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListYmlsJobResponse), nil
	}
}

// 修改密码
//
// 该接口用于修改集群密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// 重启集群
//
// 此接口用于重启集群，重启集群将导致业务中断。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

// 恢复快照
//
// 该接口用于手动恢复一个快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) RestoreSnapshot(request *model.RestoreSnapshotRequest) (*model.RestoreSnapshotResponse, error) {
	requestDef := GenReqDefForRestoreSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreSnapshotResponse), nil
	}
}

// 查询集群的自动创建快照策略
//
// 该接口用于查询自动创建快照策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowAutoCreatePolicy(request *model.ShowAutoCreatePolicyRequest) (*model.ShowAutoCreatePolicyResponse, error) {
	requestDef := GenReqDefForShowAutoCreatePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoCreatePolicyResponse), nil
	}
}

// 查询集群详情
//
// 该接口用于查询并显示单个集群详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowClusterDetail(request *model.ShowClusterDetailRequest) (*model.ShowClusterDetailResponse, error) {
	requestDef := GenReqDefForShowClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailResponse), nil
	}
}

// 查询指定集群的标签
//
// 该接口用于查询指定集群的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowClusterTag(request *model.ShowClusterTagRequest) (*model.ShowClusterTagResponse, error) {
	requestDef := GenReqDefForShowClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterTagResponse), nil
	}
}

// 查询日志基础配置
//
// 该接口用于日志基础配置查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowGetLogSetting(request *model.ShowGetLogSettingRequest) (*model.ShowGetLogSettingResponse, error) {
	requestDef := GenReqDefForShowGetLogSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetLogSettingResponse), nil
	}
}

// 查询自定义词库状态
//
// 该接口用于查询自定义词库的加载状态。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowIkThesaurus(request *model.ShowIkThesaurusRequest) (*model.ShowIkThesaurusResponse, error) {
	requestDef := GenReqDefForShowIkThesaurus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIkThesaurusResponse), nil
	}
}

// 查询日志
//
// 该接口用于查询日志信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowLogBackup(request *model.ShowLogBackupRequest) (*model.ShowLogBackupResponse, error) {
	requestDef := GenReqDefForShowLogBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogBackupResponse), nil
	}
}

// 获取终端节点连接
//
// 该接口用于获取终端节点连接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowVpcepConnection(request *model.ShowVpcepConnectionRequest) (*model.ShowVpcepConnectionResponse, error) {
	requestDef := GenReqDefForShowVpcepConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVpcepConnectionResponse), nil
	}
}

// 自动设置集群快照的基础配置（不推荐使用）
//
// 说明：自动设置集群快照接口将会自动创建快照OBS桶和委托。如果有多个集群，每个集群使用这个接口都会创建一个不一样的OBS桶，可能会导致OBS的配额不够，较多的OBS桶也难以维护。建议可以直接使用[修改集群快照的基础配置](https://support.huaweicloud.com/api-css/css_03_0030.html)。
//
// 该接口用于自动设置集群快照的基础配置，包括配置OBS桶和IAM委托。
//
// - “OBS桶”：快照存储的OBS桶位置。
// - “备份路径”：快照在OBS桶中的存放路径。
// - “IAM委托”：由于需要将快照保存在OBS中，所以需要在IAM中设置对应的委托获取对OBS服务的授权。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartAutoSetting(request *model.StartAutoSettingRequest) (*model.StartAutoSettingResponse, error) {
	requestDef := GenReqDefForStartAutoSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAutoSettingResponse), nil
	}
}

// 开启日志自动备份策略
//
// 该接口用于日志自动备份策略开启。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartLogAutoBackupPolicy(request *model.StartLogAutoBackupPolicyRequest) (*model.StartLogAutoBackupPolicyResponse, error) {
	requestDef := GenReqDefForStartLogAutoBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartLogAutoBackupPolicyResponse), nil
	}
}

// 开启日志功能
//
// 该接口用于开启日志功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartLogs(request *model.StartLogsRequest) (*model.StartLogsResponse, error) {
	requestDef := GenReqDefForStartLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartLogsResponse), nil
	}
}

// 开启公网访问控制白名单
//
// 该接口用于开启公网访问控制白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartPublicWhitelist(request *model.StartPublicWhitelistRequest) (*model.StartPublicWhitelistResponse, error) {
	requestDef := GenReqDefForStartPublicWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPublicWhitelistResponse), nil
	}
}

// 开启终端节点服务
//
// 该接口用于开启终端节点服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartVpecp(request *model.StartVpecpRequest) (*model.StartVpecpResponse, error) {
	requestDef := GenReqDefForStartVpecp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartVpecpResponse), nil
	}
}

// 关闭日志自动备份策略
//
// 该接口用于日志自动备份策略关闭。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopLogAutoBackupPolicy(request *model.StopLogAutoBackupPolicyRequest) (*model.StopLogAutoBackupPolicyResponse, error) {
	requestDef := GenReqDefForStopLogAutoBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopLogAutoBackupPolicyResponse), nil
	}
}

// 关闭日志功能
//
// 该接口用于关闭日志功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopLogs(request *model.StopLogsRequest) (*model.StopLogsResponse, error) {
	requestDef := GenReqDefForStopLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopLogsResponse), nil
	}
}

// 关闭公网访问控制白名单
//
// 该接口用于关闭公网访问控制白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopPublicWhitelist(request *model.StopPublicWhitelistRequest) (*model.StopPublicWhitelistResponse, error) {
	requestDef := GenReqDefForStopPublicWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPublicWhitelistResponse), nil
	}
}

// 停用快照功能
//
// 该接口用于停用快照功能。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopSnapshot(request *model.StopSnapshotRequest) (*model.StopSnapshotResponse, error) {
	requestDef := GenReqDefForStopSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSnapshotResponse), nil
	}
}

// 关闭终端节点服务
//
// 该接口用于关闭终端节点服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopVpecp(request *model.StopVpecpRequest) (*model.StopVpecpResponse, error) {
	requestDef := GenReqDefForStopVpecp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopVpecpResponse), nil
	}
}

// 批量添加或删除集群标签
//
// 该接口用于批量添加或删除集群标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateBatchClustersTags(request *model.UpdateBatchClustersTagsRequest) (*model.UpdateBatchClustersTagsResponse, error) {
	requestDef := GenReqDefForUpdateBatchClustersTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBatchClustersTagsResponse), nil
	}
}

// 修改集群名称
//
// 该接口用于修改集群名称。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateClusterName(request *model.UpdateClusterNameRequest) (*model.UpdateClusterNameResponse, error) {
	requestDef := GenReqDefForUpdateClusterName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterNameResponse), nil
	}
}

// 扩容集群
//
// 该接口用于集群扩容实例（仅支持扩容elasticsearch实例）。只扩容普通节点，且只针对要扩容的集群实例不存在特殊节点（Master、Client、冷数据节点）的情况。
// 说明：推荐使用[扩容实例的数量和存储容量](https://support.huaweicloud.com/api-css/css_03_0038.html)进行扩容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateExtendCluster(request *model.UpdateExtendClusterRequest) (*model.UpdateExtendClusterResponse, error) {
	requestDef := GenReqDefForUpdateExtendCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExtendClusterResponse), nil
	}
}

// 扩容实例的数量和存储容量
//
// 该接口用于集群扩容不同类型实例的个数以及存储容量。已经存在独立Master、Client、冷数据节点的集群使用该接口扩容。（支持扩容elasticsearch和logstash实例）。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateExtendInstanceStorage(request *model.UpdateExtendInstanceStorageRequest) (*model.UpdateExtendInstanceStorageResponse, error) {
	requestDef := GenReqDefForUpdateExtendInstanceStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExtendInstanceStorageResponse), nil
	}
}

// 变更规格
//
// 该接口用于变更集群规格。只支持变更ess节点类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateFlavor(request *model.UpdateFlavorRequest) (*model.UpdateFlavorResponse, error) {
	requestDef := GenReqDefForUpdateFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlavorResponse), nil
	}
}

// 全规格集群变更
//
// 修改集群规格。支持修改ess， ess-cold， ess-client， ess-master节点类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateFlavorByType(request *model.UpdateFlavorByTypeRequest) (*model.UpdateFlavorByTypeResponse, error) {
	requestDef := GenReqDefForUpdateFlavorByType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlavorByTypeResponse), nil
	}
}

// 修改日志基础配置
//
// 该接口用于修改日志基础配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateLogSetting(request *model.UpdateLogSettingRequest) (*model.UpdateLogSettingResponse, error) {
	requestDef := GenReqDefForUpdateLogSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLogSettingResponse), nil
	}
}

// 按需集群转包周期
//
// 该接口用于按需集群转包周期集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateOndemandClusterToPeriod(request *model.UpdateOndemandClusterToPeriodRequest) (*model.UpdateOndemandClusterToPeriodResponse, error) {
	requestDef := GenReqDefForUpdateOndemandClusterToPeriod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateOndemandClusterToPeriodResponse), nil
	}
}

// 修改公网访问带宽
//
// 该接口用于修改公网访问带宽。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdatePublicBandWidth(request *model.UpdatePublicBandWidthRequest) (*model.UpdatePublicBandWidthResponse, error) {
	requestDef := GenReqDefForUpdatePublicBandWidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicBandWidthResponse), nil
	}
}

// 缩容集群
//
// 该接口用于集群缩容不同类型实例的个数以及存储容量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateShrinkCluster(request *model.UpdateShrinkClusterRequest) (*model.UpdateShrinkClusterResponse, error) {
	requestDef := GenReqDefForUpdateShrinkCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateShrinkClusterResponse), nil
	}
}

// 指定角色下线
//
// 该接口用于下线集群指定角色。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateShrinkNodes(request *model.UpdateShrinkNodesRequest) (*model.UpdateShrinkNodesResponse, error) {
	requestDef := GenReqDefForUpdateShrinkNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateShrinkNodesResponse), nil
	}
}

// 修改集群快照的基础配置
//
// 该接口用于修改集群快照的基础配置，可修改OBS桶和IAM委托。 说明：如果未开启快照功能，使用该接口后，将会开启快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateSnapshotSetting(request *model.UpdateSnapshotSettingRequest) (*model.UpdateSnapshotSettingResponse, error) {
	requestDef := GenReqDefForUpdateSnapshotSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSnapshotSettingResponse), nil
	}
}

// 关闭公网访问
//
// 该接口用于关闭公网访问。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateUnbindPublic(request *model.UpdateUnbindPublicRequest) (*model.UpdateUnbindPublicResponse, error) {
	requestDef := GenReqDefForUpdateUnbindPublic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUnbindPublicResponse), nil
	}
}

// 更新终端节点连接
//
// 该接口用于更新终端节点连接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateVpcepConnection(request *model.UpdateVpcepConnectionRequest) (*model.UpdateVpcepConnectionResponse, error) {
	requestDef := GenReqDefForUpdateVpcepConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcepConnectionResponse), nil
	}
}

// 修改终端节点服务白名单
//
// 该接口用于修改终端节点服务白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateVpcepWhitelist(request *model.UpdateVpcepWhitelistRequest) (*model.UpdateVpcepWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateVpcepWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVpcepWhitelistResponse), nil
	}
}

// 修改参数配置
//
// 该接口用于修改参数配口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateYmls(request *model.UpdateYmlsRequest) (*model.UpdateYmlsResponse, error) {
	requestDef := GenReqDefForUpdateYmls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateYmlsResponse), nil
	}
}

// 开启Kibana公网访问
//
// 该接口用于开启Kibana公网访问。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartKibanaPublic(request *model.StartKibanaPublicRequest) (*model.StartKibanaPublicResponse, error) {
	requestDef := GenReqDefForStartKibanaPublic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartKibanaPublicResponse), nil
	}
}

// 关闭访问控制
//
// 该接口用于关闭Kibana公网访问控制。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopPublicKibanaWhitelist(request *model.StopPublicKibanaWhitelistRequest) (*model.StopPublicKibanaWhitelistResponse, error) {
	requestDef := GenReqDefForStopPublicKibanaWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPublicKibanaWhitelistResponse), nil
	}
}

// 修改Kibana公网带宽
//
// 该接口用于修改Kibana公网带宽。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateAlterKibana(request *model.UpdateAlterKibanaRequest) (*model.UpdateAlterKibanaResponse, error) {
	requestDef := GenReqDefForUpdateAlterKibana()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlterKibanaResponse), nil
	}
}

// 关闭Kibana公网访问
//
// 该接口用于关闭Kibana公网访问。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateCloseKibana(request *model.UpdateCloseKibanaRequest) (*model.UpdateCloseKibanaResponse, error) {
	requestDef := GenReqDefForUpdateCloseKibana()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCloseKibanaResponse), nil
	}
}

// 修改访问控制
//
// 该接口通过修改kibana白名单，修改kibana的访问权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdatePublicKibanaWhitelist(request *model.UpdatePublicKibanaWhitelistRequest) (*model.UpdatePublicKibanaWhitelistResponse, error) {
	requestDef := GenReqDefForUpdatePublicKibanaWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicKibanaWhitelistResponse), nil
	}
}

// 添加到自定义模板
//
// 该接口用于添加到自定义模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) AddFavorite(request *model.AddFavoriteRequest) (*model.AddFavoriteResponse, error) {
	requestDef := GenReqDefForAddFavorite()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddFavoriteResponse), nil
	}
}

// 创建配置文件
//
// 该接口用于创建配置文件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) CreateCnf(request *model.CreateCnfRequest) (*model.CreateCnfResponse, error) {
	requestDef := GenReqDefForCreateCnf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCnfResponse), nil
	}
}

// 删除配置文件
//
// 删除配置文件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteConf(request *model.DeleteConfRequest) (*model.DeleteConfResponse, error) {
	requestDef := GenReqDefForDeleteConf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfResponse), nil
	}
}

// 删除自定义模板
//
// 该接口用于删除自定义模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) DeleteTemplate(request *model.DeleteTemplateRequest) (*model.DeleteTemplateResponse, error) {
	requestDef := GenReqDefForDeleteTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTemplateResponse), nil
	}
}

// 查询操作记录
//
// 该接口用于查询操作记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListActions(request *model.ListActionsRequest) (*model.ListActionsResponse, error) {
	requestDef := GenReqDefForListActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListActionsResponse), nil
	}
}

// 查询配置文件列表
//
// 该接口用于查询配置文件列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListConfs(request *model.ListConfsRequest) (*model.ListConfsResponse, error) {
	requestDef := GenReqDefForListConfs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfsResponse), nil
	}
}

// 查询pipeline列表
//
// 该接口用于查询pipeline列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListPipelines(request *model.ListPipelinesRequest) (*model.ListPipelinesResponse, error) {
	requestDef := GenReqDefForListPipelines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPipelinesResponse), nil
	}
}

// 查询模板列表
//
// 该接口用于查询模板列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ListTemplates(request *model.ListTemplatesRequest) (*model.ListTemplatesResponse, error) {
	requestDef := GenReqDefForListTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTemplatesResponse), nil
	}
}

// 查询配置文件内容
//
// 该接口用于查询配置文件内容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) ShowGetConfDetail(request *model.ShowGetConfDetailRequest) (*model.ShowGetConfDetailResponse, error) {
	requestDef := GenReqDefForShowGetConfDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetConfDetailResponse), nil
	}
}

// 连通性测试
//
// 该接口用于连通性测试。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartConnectivityTest(request *model.StartConnectivityTestRequest) (*model.StartConnectivityTestResponse, error) {
	requestDef := GenReqDefForStartConnectivityTest()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartConnectivityTestResponse), nil
	}
}

// 启动pipeline迁移数据
//
// 该接口用于启动pipeline迁移数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StartPipeline(request *model.StartPipelineRequest) (*model.StartPipelineResponse, error) {
	requestDef := GenReqDefForStartPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartPipelineResponse), nil
	}
}

// 停止pipeline迁移数据
//
// 该接口用于停止pipeline迁移数据。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) StopPipeline(request *model.StopPipelineRequest) (*model.StopPipelineResponse, error) {
	requestDef := GenReqDefForStopPipeline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopPipelineResponse), nil
	}
}

// 更新配置文件
//
// 该接口用于更新配置文件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CssClient) UpdateCnf(request *model.UpdateCnfRequest) (*model.UpdateCnfResponse, error) {
	requestDef := GenReqDefForUpdateCnf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCnfResponse), nil
	}
}
