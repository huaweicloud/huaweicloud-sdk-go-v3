package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dcs/v2/model"
)

type DcsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDcsClient(hcClient *http_client.HcHttpClient) *DcsClient {
	return &DcsClient{HcClient: hcClient}
}

func DcsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// BatchCreateOrDeleteTags 批量添加或删除标签
//
// 为指定实例批量添加标签，或批量删除标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) BatchCreateOrDeleteTags(request *model.BatchCreateOrDeleteTagsRequest) (*model.BatchCreateOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

// BatchCreateOrDeleteTagsInvoker 批量添加或删除标签
func (c *DcsClient) BatchCreateOrDeleteTagsInvoker(request *model.BatchCreateOrDeleteTagsRequest) *BatchCreateOrDeleteTagsInvoker {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()
	return &BatchCreateOrDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInstances 批量删除实例
//
// 批量删除多个缓存实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

// BatchDeleteInstancesInvoker 批量删除实例
func (c *DcsClient) BatchDeleteInstancesInvoker(request *model.BatchDeleteInstancesRequest) *BatchDeleteInstancesInvoker {
	requestDef := GenReqDefForBatchDeleteInstances()
	return &BatchDeleteInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowNodesInformation 批量查询实例节点信息
//
// 批量查询指定项目所有实例的节点信息、有效实例个数及节点个数。
// 创建中实例不返回节点信息。
// 仅支持Redis4.0和Redis5.0实例查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) BatchShowNodesInformation(request *model.BatchShowNodesInformationRequest) (*model.BatchShowNodesInformationResponse, error) {
	requestDef := GenReqDefForBatchShowNodesInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowNodesInformationResponse), nil
	}
}

// BatchShowNodesInformationInvoker 批量查询实例节点信息
func (c *DcsClient) BatchShowNodesInformationInvoker(request *model.BatchShowNodesInformationRequest) *BatchShowNodesInformationInvoker {
	requestDef := GenReqDefForBatchShowNodesInformation()
	return &BatchShowNodesInformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchStopMigrationTasks 批量停止数据迁移任务
//
// 批量停止数据迁移任务，接口响应成功，仅表示下发任务成功。查询到迁移任务状态为TERMINATED时，即停止成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) BatchStopMigrationTasks(request *model.BatchStopMigrationTasksRequest) (*model.BatchStopMigrationTasksResponse, error) {
	requestDef := GenReqDefForBatchStopMigrationTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopMigrationTasksResponse), nil
	}
}

// BatchStopMigrationTasksInvoker 批量停止数据迁移任务
func (c *DcsClient) BatchStopMigrationTasksInvoker(request *model.BatchStopMigrationTasksRequest) *BatchStopMigrationTasksInvoker {
	requestDef := GenReqDefForBatchStopMigrationTasks()
	return &BatchStopMigrationTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeMasterStandby 主备切换
//
// 切换实例主备节点，只有主备实例支持该操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ChangeMasterStandby(request *model.ChangeMasterStandbyRequest) (*model.ChangeMasterStandbyResponse, error) {
	requestDef := GenReqDefForChangeMasterStandby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeMasterStandbyResponse), nil
	}
}

// ChangeMasterStandbyInvoker 主备切换
func (c *DcsClient) ChangeMasterStandbyInvoker(request *model.ChangeMasterStandbyRequest) *ChangeMasterStandbyInvoker {
	requestDef := GenReqDefForChangeMasterStandby()
	return &ChangeMasterStandbyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyInstance 备份指定实例
//
// 备份指定的缓存实例。
// &gt; 只有主备和集群类型的缓存实例支持备份恢复操作，单机实例不支持备份恢复操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CopyInstance(request *model.CopyInstanceRequest) (*model.CopyInstanceResponse, error) {
	requestDef := GenReqDefForCopyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyInstanceResponse), nil
	}
}

// CopyInstanceInvoker 备份指定实例
func (c *DcsClient) CopyInstanceInvoker(request *model.CopyInstanceRequest) *CopyInstanceInvoker {
	requestDef := GenReqDefForCopyInstance()
	return &CopyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAutoExpireScanTask 创建过期key扫描任务
//
// 创建过期key扫描任务（Redis 3.0 不支持过期key扫描）。
// 过期key扫描会对键空间进行Redis的scan扫描，释放内存中已过期但是由于惰性删除机制而没有释放的内存空间。
// 过期key扫描在主节点上执行，会对实例性能有一定的影响，建议不要在业务高峰期进行。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateAutoExpireScanTask(request *model.CreateAutoExpireScanTaskRequest) (*model.CreateAutoExpireScanTaskResponse, error) {
	requestDef := GenReqDefForCreateAutoExpireScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAutoExpireScanTaskResponse), nil
	}
}

// CreateAutoExpireScanTaskInvoker 创建过期key扫描任务
func (c *DcsClient) CreateAutoExpireScanTaskInvoker(request *model.CreateAutoExpireScanTaskRequest) *CreateAutoExpireScanTaskInvoker {
	requestDef := GenReqDefForCreateAutoExpireScanTask()
	return &CreateAutoExpireScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBigkeyScanTask 创建大key分析任务
//
// 为Redis实例创建大key分析任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateBigkeyScanTask(request *model.CreateBigkeyScanTaskRequest) (*model.CreateBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBigkeyScanTaskResponse), nil
	}
}

// CreateBigkeyScanTaskInvoker 创建大key分析任务
func (c *DcsClient) CreateBigkeyScanTaskInvoker(request *model.CreateBigkeyScanTaskRequest) *CreateBigkeyScanTaskInvoker {
	requestDef := GenReqDefForCreateBigkeyScanTask()
	return &CreateBigkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCustomTemplate 创建自定义模板
//
// 创建自定义模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateCustomTemplate(request *model.CreateCustomTemplateRequest) (*model.CreateCustomTemplateResponse, error) {
	requestDef := GenReqDefForCreateCustomTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomTemplateResponse), nil
	}
}

// CreateCustomTemplateInvoker 创建自定义模板
func (c *DcsClient) CreateCustomTemplateInvoker(request *model.CreateCustomTemplateRequest) *CreateCustomTemplateInvoker {
	requestDef := GenReqDefForCreateCustomTemplate()
	return &CreateCustomTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnosisTask 创建实例诊断任务
//
// 诊断指定的缓存实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateDiagnosisTask(request *model.CreateDiagnosisTaskRequest) (*model.CreateDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskResponse), nil
	}
}

// CreateDiagnosisTaskInvoker 创建实例诊断任务
func (c *DcsClient) CreateDiagnosisTaskInvoker(request *model.CreateDiagnosisTaskRequest) *CreateDiagnosisTaskInvoker {
	requestDef := GenReqDefForCreateDiagnosisTask()
	return &CreateDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHotkeyScanTask 创建热key分析任务
//
// 创建热key分析任务，Redis 3.0 不支持热key分析。
//
// 热key分析需要将缓存实例配置参数maxmemory-policy设置为allkeys-lfu或volatile-lfu。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateHotkeyScanTask(request *model.CreateHotkeyScanTaskRequest) (*model.CreateHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHotkeyScanTaskResponse), nil
	}
}

// CreateHotkeyScanTaskInvoker 创建热key分析任务
func (c *DcsClient) CreateHotkeyScanTaskInvoker(request *model.CreateHotkeyScanTaskRequest) *CreateHotkeyScanTaskInvoker {
	requestDef := GenReqDefForCreateHotkeyScanTask()
	return &CreateHotkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建缓存实例
//
// 创建缓存实例，该接口创建的缓存实例支持按需计费和包周期两种方式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建缓存实例
func (c *DcsClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMigrationTask 创建数据迁移任务
//
// 创建数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateMigrationTask(request *model.CreateMigrationTaskRequest) (*model.CreateMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigrationTaskResponse), nil
	}
}

// CreateMigrationTaskInvoker 创建数据迁移任务
func (c *DcsClient) CreateMigrationTaskInvoker(request *model.CreateMigrationTaskRequest) *CreateMigrationTaskInvoker {
	requestDef := GenReqDefForCreateMigrationTask()
	return &CreateMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOnlineMigrationTask 创建在线数据迁移任务
//
// 创建在线数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateOnlineMigrationTask(request *model.CreateOnlineMigrationTaskRequest) (*model.CreateOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOnlineMigrationTaskResponse), nil
	}
}

// CreateOnlineMigrationTaskInvoker 创建在线数据迁移任务
func (c *DcsClient) CreateOnlineMigrationTaskInvoker(request *model.CreateOnlineMigrationTaskRequest) *CreateOnlineMigrationTaskInvoker {
	requestDef := GenReqDefForCreateOnlineMigrationTask()
	return &CreateOnlineMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRedislog 采集Redis运行日志
//
// 采集Redis运行日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateRedislog(request *model.CreateRedislogRequest) (*model.CreateRedislogResponse, error) {
	requestDef := GenReqDefForCreateRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogResponse), nil
	}
}

// CreateRedislogInvoker 采集Redis运行日志
func (c *DcsClient) CreateRedislogInvoker(request *model.CreateRedislogRequest) *CreateRedislogInvoker {
	requestDef := GenReqDefForCreateRedislog()
	return &CreateRedislogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRedislogDownloadLink 获取日志下载链接
//
// 获取日志下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) CreateRedislogDownloadLink(request *model.CreateRedislogDownloadLinkRequest) (*model.CreateRedislogDownloadLinkResponse, error) {
	requestDef := GenReqDefForCreateRedislogDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogDownloadLinkResponse), nil
	}
}

// CreateRedislogDownloadLinkInvoker 获取日志下载链接
func (c *DcsClient) CreateRedislogDownloadLinkInvoker(request *model.CreateRedislogDownloadLinkRequest) *CreateRedislogDownloadLinkInvoker {
	requestDef := GenReqDefForCreateRedislogDownloadLink()
	return &CreateRedislogDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackgroundTask 删除后台任务
//
// 删除后台任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// DeleteBackgroundTaskInvoker 删除后台任务
func (c *DcsClient) DeleteBackgroundTaskInvoker(request *model.DeleteBackgroundTaskRequest) *DeleteBackgroundTaskInvoker {
	requestDef := GenReqDefForDeleteBackgroundTask()
	return &DeleteBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackupFile 删除备份文件
//
// 删除缓存实例已备份的文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteBackupFile(request *model.DeleteBackupFileRequest) (*model.DeleteBackupFileResponse, error) {
	requestDef := GenReqDefForDeleteBackupFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupFileResponse), nil
	}
}

// DeleteBackupFileInvoker 删除备份文件
func (c *DcsClient) DeleteBackupFileInvoker(request *model.DeleteBackupFileRequest) *DeleteBackupFileInvoker {
	requestDef := GenReqDefForDeleteBackupFile()
	return &DeleteBackupFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBigkeyScanTask 删除大key分析记录
//
// 删除大key分析记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteBigkeyScanTask(request *model.DeleteBigkeyScanTaskRequest) (*model.DeleteBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBigkeyScanTaskResponse), nil
	}
}

// DeleteBigkeyScanTaskInvoker 删除大key分析记录
func (c *DcsClient) DeleteBigkeyScanTaskInvoker(request *model.DeleteBigkeyScanTaskRequest) *DeleteBigkeyScanTaskInvoker {
	requestDef := GenReqDefForDeleteBigkeyScanTask()
	return &DeleteBigkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHotkeyScanTask 删除热key分析任务
//
// 删除热key分析任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteHotkeyScanTask(request *model.DeleteHotkeyScanTaskRequest) (*model.DeleteHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHotkeyScanTaskResponse), nil
	}
}

// DeleteHotkeyScanTaskInvoker 删除热key分析任务
func (c *DcsClient) DeleteHotkeyScanTaskInvoker(request *model.DeleteHotkeyScanTaskRequest) *DeleteHotkeyScanTaskInvoker {
	requestDef := GenReqDefForDeleteHotkeyScanTask()
	return &DeleteHotkeyScanTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIpFromDomainName 域名摘除IP
//
// 将只读副本的IP从域名中摘除，摘除成功后，只读域名不会再解析到该副本IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteIpFromDomainName(request *model.DeleteIpFromDomainNameRequest) (*model.DeleteIpFromDomainNameResponse, error) {
	requestDef := GenReqDefForDeleteIpFromDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpFromDomainNameResponse), nil
	}
}

// DeleteIpFromDomainNameInvoker 域名摘除IP
func (c *DcsClient) DeleteIpFromDomainNameInvoker(request *model.DeleteIpFromDomainNameRequest) *DeleteIpFromDomainNameInvoker {
	requestDef := GenReqDefForDeleteIpFromDomainName()
	return &DeleteIpFromDomainNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMigrationTask 删除数据迁移任务
//
// 删除数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteMigrationTask(request *model.DeleteMigrationTaskRequest) (*model.DeleteMigrationTaskResponse, error) {
	requestDef := GenReqDefForDeleteMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigrationTaskResponse), nil
	}
}

// DeleteMigrationTaskInvoker 删除数据迁移任务
func (c *DcsClient) DeleteMigrationTaskInvoker(request *model.DeleteMigrationTaskRequest) *DeleteMigrationTaskInvoker {
	requestDef := GenReqDefForDeleteMigrationTask()
	return &DeleteMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSingleInstance 删除实例
//
// 删除指定的缓存实例，释放该实例的所有资源。
//
// &gt; 如果是删除按需资源，请按照本章节执行；如果是删除包周期资源，即退订，请参考[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html#section2)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) DeleteSingleInstance(request *model.DeleteSingleInstanceRequest) (*model.DeleteSingleInstanceResponse, error) {
	requestDef := GenReqDefForDeleteSingleInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSingleInstanceResponse), nil
	}
}

// DeleteSingleInstanceInvoker 删除实例
func (c *DcsClient) DeleteSingleInstanceInvoker(request *model.DeleteSingleInstanceRequest) *DeleteSingleInstanceInvoker {
	requestDef := GenReqDefForDeleteSingleInstance()
	return &DeleteSingleInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteClusterSwitchover 集群分片倒换
//
// 集群分片倒换，适用于proxy和cluster实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ExecuteClusterSwitchover(request *model.ExecuteClusterSwitchoverRequest) (*model.ExecuteClusterSwitchoverResponse, error) {
	requestDef := GenReqDefForExecuteClusterSwitchover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteClusterSwitchoverResponse), nil
	}
}

// ExecuteClusterSwitchoverInvoker 集群分片倒换
func (c *DcsClient) ExecuteClusterSwitchoverInvoker(request *model.ExecuteClusterSwitchoverRequest) *ExecuteClusterSwitchoverInvoker {
	requestDef := GenReqDefForExecuteClusterSwitchover()
	return &ExecuteClusterSwitchoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableZones 查询可用区信息
//
// 查询所在局点的可用区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// ListAvailableZonesInvoker 查询可用区信息
func (c *DcsClient) ListAvailableZonesInvoker(request *model.ListAvailableZonesRequest) *ListAvailableZonesInvoker {
	requestDef := GenReqDefForListAvailableZones()
	return &ListAvailableZonesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackgroundTask 查询后台任务列表
//
// 查询后台任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListBackgroundTask(request *model.ListBackgroundTaskRequest) (*model.ListBackgroundTaskResponse, error) {
	requestDef := GenReqDefForListBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTaskResponse), nil
	}
}

// ListBackgroundTaskInvoker 查询后台任务列表
func (c *DcsClient) ListBackgroundTaskInvoker(request *model.ListBackgroundTaskRequest) *ListBackgroundTaskInvoker {
	requestDef := GenReqDefForListBackgroundTask()
	return &ListBackgroundTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackupFileLinks 获取备份文件下载链接
//
// 获取指定实例的备份文件下载链接，下载备份文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListBackupFileLinks(request *model.ListBackupFileLinksRequest) (*model.ListBackupFileLinksResponse, error) {
	requestDef := GenReqDefForListBackupFileLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupFileLinksResponse), nil
	}
}

// ListBackupFileLinksInvoker 获取备份文件下载链接
func (c *DcsClient) ListBackupFileLinksInvoker(request *model.ListBackupFileLinksRequest) *ListBackupFileLinksInvoker {
	requestDef := GenReqDefForListBackupFileLinks()
	return &ListBackupFileLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackupRecords 查询实例备份信息
//
// 查询指定缓存实例的备份信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListBackupRecords(request *model.ListBackupRecordsRequest) (*model.ListBackupRecordsResponse, error) {
	requestDef := GenReqDefForListBackupRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupRecordsResponse), nil
	}
}

// ListBackupRecordsInvoker 查询实例备份信息
func (c *DcsClient) ListBackupRecordsInvoker(request *model.ListBackupRecordsRequest) *ListBackupRecordsInvoker {
	requestDef := GenReqDefForListBackupRecords()
	return &ListBackupRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBigkeyScanTasks 查询大key分析任务列表
//
// 查询大key分析任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListBigkeyScanTasks(request *model.ListBigkeyScanTasksRequest) (*model.ListBigkeyScanTasksResponse, error) {
	requestDef := GenReqDefForListBigkeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBigkeyScanTasksResponse), nil
	}
}

// ListBigkeyScanTasksInvoker 查询大key分析任务列表
func (c *DcsClient) ListBigkeyScanTasksInvoker(request *model.ListBigkeyScanTasksRequest) *ListBigkeyScanTasksInvoker {
	requestDef := GenReqDefForListBigkeyScanTasks()
	return &ListBigkeyScanTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigHistories 查询实例参数修改记录列表
//
// 查询实例的参数修改记录列表，支持按照关键字查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListConfigHistories(request *model.ListConfigHistoriesRequest) (*model.ListConfigHistoriesResponse, error) {
	requestDef := GenReqDefForListConfigHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigHistoriesResponse), nil
	}
}

// ListConfigHistoriesInvoker 查询实例参数修改记录列表
func (c *DcsClient) ListConfigHistoriesInvoker(request *model.ListConfigHistoriesRequest) *ListConfigHistoriesInvoker {
	requestDef := GenReqDefForListConfigHistories()
	return &ListConfigHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigTemplates 查询参数模板列表
//
// 查询租户的参数模板列表，支持按照条件查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListConfigTemplates(request *model.ListConfigTemplatesRequest) (*model.ListConfigTemplatesResponse, error) {
	requestDef := GenReqDefForListConfigTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigTemplatesResponse), nil
	}
}

// ListConfigTemplatesInvoker 查询参数模板列表
func (c *DcsClient) ListConfigTemplatesInvoker(request *model.ListConfigTemplatesRequest) *ListConfigTemplatesInvoker {
	requestDef := GenReqDefForListConfigTemplates()
	return &ListConfigTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurations 查询实例配置参数
//
// 查询指定实例的配置参数信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// ListConfigurationsInvoker 查询实例配置参数
func (c *DcsClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnosisTasks 查询实例诊断任务列表
//
// 查询指定缓存实例诊断任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListDiagnosisTasks(request *model.ListDiagnosisTasksRequest) (*model.ListDiagnosisTasksResponse, error) {
	requestDef := GenReqDefForListDiagnosisTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisTasksResponse), nil
	}
}

// ListDiagnosisTasksInvoker 查询实例诊断任务列表
func (c *DcsClient) ListDiagnosisTasksInvoker(request *model.ListDiagnosisTasksRequest) *ListDiagnosisTasksInvoker {
	requestDef := GenReqDefForListDiagnosisTasks()
	return &ListDiagnosisTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询产品规格
//
// 在创建缓存实例时，需要配置订购的产品规格编码（spec_code），可通过该接口查询产品规格，查询条件不选时默认查询全部。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询产品规格
func (c *DcsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupReplicationInfo 查询分片信息
//
// 查询读写分离实例和集群实例的分片和副本信息，其中，读写分离实例仅Redis4.0和Redis5.0的主备实例支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListGroupReplicationInfo(request *model.ListGroupReplicationInfoRequest) (*model.ListGroupReplicationInfoResponse, error) {
	requestDef := GenReqDefForListGroupReplicationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupReplicationInfoResponse), nil
	}
}

// ListGroupReplicationInfoInvoker 查询分片信息
func (c *DcsClient) ListGroupReplicationInfoInvoker(request *model.ListGroupReplicationInfoRequest) *ListGroupReplicationInfoInvoker {
	requestDef := GenReqDefForListGroupReplicationInfo()
	return &ListGroupReplicationInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHotKeyScanTasks 查询热key分析任务列表
//
// 查询热key分析历史记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListHotKeyScanTasks(request *model.ListHotKeyScanTasksRequest) (*model.ListHotKeyScanTasksResponse, error) {
	requestDef := GenReqDefForListHotKeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHotKeyScanTasksResponse), nil
	}
}

// ListHotKeyScanTasksInvoker 查询热key分析任务列表
func (c *DcsClient) ListHotKeyScanTasksInvoker(request *model.ListHotKeyScanTasksRequest) *ListHotKeyScanTasksInvoker {
	requestDef := GenReqDefForListHotKeyScanTasks()
	return &ListHotKeyScanTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询所有实例列表
//
// 查询租户的缓存实例列表，支持按照条件查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询所有实例列表
func (c *DcsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMaintenanceWindows 查询维护时间窗时间段
//
// 查询维护时间窗开始时间和结束时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListMaintenanceWindows(request *model.ListMaintenanceWindowsRequest) (*model.ListMaintenanceWindowsResponse, error) {
	requestDef := GenReqDefForListMaintenanceWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMaintenanceWindowsResponse), nil
	}
}

// ListMaintenanceWindowsInvoker 查询维护时间窗时间段
func (c *DcsClient) ListMaintenanceWindowsInvoker(request *model.ListMaintenanceWindowsRequest) *ListMaintenanceWindowsInvoker {
	requestDef := GenReqDefForListMaintenanceWindows()
	return &ListMaintenanceWindowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMigrationTask 查询迁移任务列表
//
// 查询迁移任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListMigrationTask(request *model.ListMigrationTaskRequest) (*model.ListMigrationTaskResponse, error) {
	requestDef := GenReqDefForListMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigrationTaskResponse), nil
	}
}

// ListMigrationTaskInvoker 查询迁移任务列表
func (c *DcsClient) ListMigrationTaskInvoker(request *model.ListMigrationTaskRequest) *ListMigrationTaskInvoker {
	requestDef := GenReqDefForListMigrationTask()
	return &ListMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitoredObjects 查询主维度信息列表
//
// 查询主维度对象列表，主维度ID当前支持dcs_instance_id，dcs_memcached_instance_id。
// &gt; 该接口当前仅在中国华南区开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListMonitoredObjects(request *model.ListMonitoredObjectsRequest) (*model.ListMonitoredObjectsResponse, error) {
	requestDef := GenReqDefForListMonitoredObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsResponse), nil
	}
}

// ListMonitoredObjectsInvoker 查询主维度信息列表
func (c *DcsClient) ListMonitoredObjectsInvoker(request *model.ListMonitoredObjectsRequest) *ListMonitoredObjectsInvoker {
	requestDef := GenReqDefForListMonitoredObjects()
	return &ListMonitoredObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMonitoredObjectsOfInstance 查询单个主维度下子维度监控对象列表
//
// 查询主维度下子维度监控对象列表，当前支持子维度的主维度ID的有 dcs_instance_id
// &gt; 该接口当前仅在中国华南区开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListMonitoredObjectsOfInstance(request *model.ListMonitoredObjectsOfInstanceRequest) (*model.ListMonitoredObjectsOfInstanceResponse, error) {
	requestDef := GenReqDefForListMonitoredObjectsOfInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsOfInstanceResponse), nil
	}
}

// ListMonitoredObjectsOfInstanceInvoker 查询单个主维度下子维度监控对象列表
func (c *DcsClient) ListMonitoredObjectsOfInstanceInvoker(request *model.ListMonitoredObjectsOfInstanceRequest) *ListMonitoredObjectsOfInstanceInvoker {
	requestDef := GenReqDefForListMonitoredObjectsOfInstance()
	return &ListMonitoredObjectsOfInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNumberOfInstancesInDifferentStatus 查询实例状态
//
// 查询该租户在当前区域下不同状态的实例数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListNumberOfInstancesInDifferentStatus(request *model.ListNumberOfInstancesInDifferentStatusRequest) (*model.ListNumberOfInstancesInDifferentStatusResponse, error) {
	requestDef := GenReqDefForListNumberOfInstancesInDifferentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNumberOfInstancesInDifferentStatusResponse), nil
	}
}

// ListNumberOfInstancesInDifferentStatusInvoker 查询实例状态
func (c *DcsClient) ListNumberOfInstancesInDifferentStatusInvoker(request *model.ListNumberOfInstancesInDifferentStatusRequest) *ListNumberOfInstancesInDifferentStatusInvoker {
	requestDef := GenReqDefForListNumberOfInstancesInDifferentStatus()
	return &ListNumberOfInstancesInDifferentStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRedislog 查询Redis运行日志列表
//
// 查询Redis运行日志列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListRedislog(request *model.ListRedislogRequest) (*model.ListRedislogResponse, error) {
	requestDef := GenReqDefForListRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRedislogResponse), nil
	}
}

// ListRedislogInvoker 查询Redis运行日志列表
func (c *DcsClient) ListRedislogInvoker(request *model.ListRedislogRequest) *ListRedislogInvoker {
	requestDef := GenReqDefForListRedislog()
	return &ListRedislogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreRecords 查询实例恢复记录
//
// 查询指定缓存实例的恢复记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListRestoreRecords(request *model.ListRestoreRecordsRequest) (*model.ListRestoreRecordsResponse, error) {
	requestDef := GenReqDefForListRestoreRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreRecordsResponse), nil
	}
}

// ListRestoreRecordsInvoker 查询实例恢复记录
func (c *DcsClient) ListRestoreRecordsInvoker(request *model.ListRestoreRecordsRequest) *ListRestoreRecordsInvoker {
	requestDef := GenReqDefForListRestoreRecords()
	return &ListRestoreRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowlog 查询慢日志
//
// 查询慢日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListSlowlog(request *model.ListSlowlogRequest) (*model.ListSlowlogResponse, error) {
	requestDef := GenReqDefForListSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowlogResponse), nil
	}
}

// ListSlowlogInvoker 查询慢日志
func (c *DcsClient) ListSlowlogInvoker(request *model.ListSlowlogRequest) *ListSlowlogInvoker {
	requestDef := GenReqDefForListSlowlog()
	return &ListSlowlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStatisticsOfRunningInstances 查询运行中实例的统计信息
//
// 查询当前租户下处于“运行中”状态的缓存实例的统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListStatisticsOfRunningInstances(request *model.ListStatisticsOfRunningInstancesRequest) (*model.ListStatisticsOfRunningInstancesResponse, error) {
	requestDef := GenReqDefForListStatisticsOfRunningInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsOfRunningInstancesResponse), nil
	}
}

// ListStatisticsOfRunningInstancesInvoker 查询运行中实例的统计信息
func (c *DcsClient) ListStatisticsOfRunningInstancesInvoker(request *model.ListStatisticsOfRunningInstancesRequest) *ListStatisticsOfRunningInstancesInvoker {
	requestDef := GenReqDefForListStatisticsOfRunningInstances()
	return &ListStatisticsOfRunningInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTagsOfTenant 查询租户所有标签
//
// 查询租户在指定Project中实例类型的所有资源标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ListTagsOfTenant(request *model.ListTagsOfTenantRequest) (*model.ListTagsOfTenantResponse, error) {
	requestDef := GenReqDefForListTagsOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsOfTenantResponse), nil
	}
}

// ListTagsOfTenantInvoker 查询租户所有标签
func (c *DcsClient) ListTagsOfTenantInvoker(request *model.ListTagsOfTenantRequest) *ListTagsOfTenantInvoker {
	requestDef := GenReqDefForListTagsOfTenant()
	return &ListTagsOfTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 重置密码
//
// 重置缓存实例的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 重置密码
func (c *DcsClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 变更实例规格
//
// 用户可以为状态为“运行中”的DCS缓存实例进行规格变更，当前仅能支持按需实例的同副本或分片数量的实例规格变更。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 变更实例规格
func (c *DcsClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartOrFlushInstances 重启实例或清空数据
//
// 重启运行中的DCS缓存实例。
//
// 清空Redis4.0/Redis5.0的实例数据，数据清空后，无法撤销，且无法恢复，请谨慎操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) RestartOrFlushInstances(request *model.RestartOrFlushInstancesRequest) (*model.RestartOrFlushInstancesResponse, error) {
	requestDef := GenReqDefForRestartOrFlushInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartOrFlushInstancesResponse), nil
	}
}

// RestartOrFlushInstancesInvoker 重启实例或清空数据
func (c *DcsClient) RestartOrFlushInstancesInvoker(request *model.RestartOrFlushInstancesRequest) *RestartOrFlushInstancesInvoker {
	requestDef := GenReqDefForRestartOrFlushInstances()
	return &RestartOrFlushInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreInstance 恢复指定实例
//
// 恢复指定的缓存实例。
// &gt; 只有主备和集群类型的缓存实例支持备份恢复操作，单机实例不支持备份恢复操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

// RestoreInstanceInvoker 恢复指定实例
func (c *DcsClient) RestoreInstanceInvoker(request *model.RestoreInstanceRequest) *RestoreInstanceInvoker {
	requestDef := GenReqDefForRestoreInstance()
	return &RestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetOnlineMigrationTask 配置在线数据迁移任务
//
// 配置在线数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) SetOnlineMigrationTask(request *model.SetOnlineMigrationTaskRequest) (*model.SetOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForSetOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetOnlineMigrationTaskResponse), nil
	}
}

// SetOnlineMigrationTaskInvoker 配置在线数据迁移任务
func (c *DcsClient) SetOnlineMigrationTaskInvoker(request *model.SetOnlineMigrationTaskRequest) *SetOnlineMigrationTaskInvoker {
	requestDef := GenReqDefForSetOnlineMigrationTask()
	return &SetOnlineMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBigkeyAutoscanConfig 查询大key自动分析配置
//
// 查询大key自动分析配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowBigkeyAutoscanConfig(request *model.ShowBigkeyAutoscanConfigRequest) (*model.ShowBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyAutoscanConfigResponse), nil
	}
}

// ShowBigkeyAutoscanConfigInvoker 查询大key自动分析配置
func (c *DcsClient) ShowBigkeyAutoscanConfigInvoker(request *model.ShowBigkeyAutoscanConfigRequest) *ShowBigkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForShowBigkeyAutoscanConfig()
	return &ShowBigkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBigkeyScanTaskDetails 查询大key分析详情
//
// 查询大key分析详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowBigkeyScanTaskDetails(request *model.ShowBigkeyScanTaskDetailsRequest) (*model.ShowBigkeyScanTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowBigkeyScanTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyScanTaskDetailsResponse), nil
	}
}

// ShowBigkeyScanTaskDetailsInvoker 查询大key分析详情
func (c *DcsClient) ShowBigkeyScanTaskDetailsInvoker(request *model.ShowBigkeyScanTaskDetailsRequest) *ShowBigkeyScanTaskDetailsInvoker {
	requestDef := GenReqDefForShowBigkeyScanTaskDetails()
	return &ShowBigkeyScanTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisTaskDetails 查询指定诊断报告
//
// 通过报告ID查询诊断报告的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowDiagnosisTaskDetails(request *model.ShowDiagnosisTaskDetailsRequest) (*model.ShowDiagnosisTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowDiagnosisTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisTaskDetailsResponse), nil
	}
}

// ShowDiagnosisTaskDetailsInvoker 查询指定诊断报告
func (c *DcsClient) ShowDiagnosisTaskDetailsInvoker(request *model.ShowDiagnosisTaskDetailsRequest) *ShowDiagnosisTaskDetailsInvoker {
	requestDef := GenReqDefForShowDiagnosisTaskDetails()
	return &ShowDiagnosisTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHotkeyAutoscanConfig 查询热key自动分析配置
//
// 查询热key自动分析配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowHotkeyAutoscanConfig(request *model.ShowHotkeyAutoscanConfigRequest) (*model.ShowHotkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowHotkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyAutoscanConfigResponse), nil
	}
}

// ShowHotkeyAutoscanConfigInvoker 查询热key自动分析配置
func (c *DcsClient) ShowHotkeyAutoscanConfigInvoker(request *model.ShowHotkeyAutoscanConfigRequest) *ShowHotkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForShowHotkeyAutoscanConfig()
	return &ShowHotkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHotkeyTaskDetails 查询热key分析详情
//
// 查询热key分析详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowHotkeyTaskDetails(request *model.ShowHotkeyTaskDetailsRequest) (*model.ShowHotkeyTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowHotkeyTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyTaskDetailsResponse), nil
	}
}

// ShowHotkeyTaskDetailsInvoker 查询热key分析详情
func (c *DcsClient) ShowHotkeyTaskDetailsInvoker(request *model.ShowHotkeyTaskDetailsRequest) *ShowHotkeyTaskDetailsInvoker {
	requestDef := GenReqDefForShowHotkeyTaskDetails()
	return &ShowHotkeyTaskDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询指定实例
//
// 通过实例ID查询实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询指定实例
func (c *DcsClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobInfo 查询租户Job执行结果
//
// 查询租户Job执行结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowJobInfo(request *model.ShowJobInfoRequest) (*model.ShowJobInfoResponse, error) {
	requestDef := GenReqDefForShowJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobInfoResponse), nil
	}
}

// ShowJobInfoInvoker 查询租户Job执行结果
func (c *DcsClient) ShowJobInfoInvoker(request *model.ShowJobInfoRequest) *ShowJobInfoInvoker {
	requestDef := GenReqDefForShowJobInfo()
	return &ShowJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMigrationTask 查询迁移任务详情
//
// 查询迁移任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowMigrationTask(request *model.ShowMigrationTaskRequest) (*model.ShowMigrationTaskResponse, error) {
	requestDef := GenReqDefForShowMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskResponse), nil
	}
}

// ShowMigrationTaskInvoker 查询迁移任务详情
func (c *DcsClient) ShowMigrationTaskInvoker(request *model.ShowMigrationTaskRequest) *ShowMigrationTaskInvoker {
	requestDef := GenReqDefForShowMigrationTask()
	return &ShowMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMigrationTaskStats 查询在线迁移进度明细
//
// 查询在线迁移进度明细。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowMigrationTaskStats(request *model.ShowMigrationTaskStatsRequest) (*model.ShowMigrationTaskStatsResponse, error) {
	requestDef := GenReqDefForShowMigrationTaskStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskStatsResponse), nil
	}
}

// ShowMigrationTaskStatsInvoker 查询在线迁移进度明细
func (c *DcsClient) ShowMigrationTaskStatsInvoker(request *model.ShowMigrationTaskStatsRequest) *ShowMigrationTaskStatsInvoker {
	requestDef := GenReqDefForShowMigrationTaskStats()
	return &ShowMigrationTaskStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotaOfTenant 查询租户配额
//
// 查询租户默认可以创建的实例数和总内存的配额限制，以及可以申请配额的最大值和最小值。不同的租户在不同的区域配额可能不同。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowQuotaOfTenant(request *model.ShowQuotaOfTenantRequest) (*model.ShowQuotaOfTenantResponse, error) {
	requestDef := GenReqDefForShowQuotaOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaOfTenantResponse), nil
	}
}

// ShowQuotaOfTenantInvoker 查询租户配额
func (c *DcsClient) ShowQuotaOfTenantInvoker(request *model.ShowQuotaOfTenantRequest) *ShowQuotaOfTenantInvoker {
	requestDef := GenReqDefForShowQuotaOfTenant()
	return &ShowQuotaOfTenantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTags 查询单个实例标签
//
// 通过实例ID查询标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowTags(request *model.ShowTagsRequest) (*model.ShowTagsResponse, error) {
	requestDef := GenReqDefForShowTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagsResponse), nil
	}
}

// ShowTagsInvoker 查询单个实例标签
func (c *DcsClient) ShowTagsInvoker(request *model.ShowTagsRequest) *ShowTagsInvoker {
	requestDef := GenReqDefForShowTags()
	return &ShowTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopMigrationTask 停止数据迁移任务
//
// 停止数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) StopMigrationTask(request *model.StopMigrationTaskRequest) (*model.StopMigrationTaskResponse, error) {
	requestDef := GenReqDefForStopMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskResponse), nil
	}
}

// StopMigrationTaskInvoker 停止数据迁移任务
func (c *DcsClient) StopMigrationTaskInvoker(request *model.StopMigrationTaskRequest) *StopMigrationTaskInvoker {
	requestDef := GenReqDefForStopMigrationTask()
	return &StopMigrationTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopMigrationTaskSync 同步停止数据迁移任务
//
// 同步停止数据迁移任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) StopMigrationTaskSync(request *model.StopMigrationTaskSyncRequest) (*model.StopMigrationTaskSyncResponse, error) {
	requestDef := GenReqDefForStopMigrationTaskSync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskSyncResponse), nil
	}
}

// StopMigrationTaskSyncInvoker 同步停止数据迁移任务
func (c *DcsClient) StopMigrationTaskSyncInvoker(request *model.StopMigrationTaskSyncRequest) *StopMigrationTaskSyncInvoker {
	requestDef := GenReqDefForStopMigrationTaskSync()
	return &StopMigrationTaskSyncInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBigkeyAutoscanConfig 设置大key自动分析配置
//
// 设置大key自动分析配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateBigkeyAutoscanConfig(request *model.UpdateBigkeyAutoscanConfigRequest) (*model.UpdateBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForUpdateBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBigkeyAutoscanConfigResponse), nil
	}
}

// UpdateBigkeyAutoscanConfigInvoker 设置大key自动分析配置
func (c *DcsClient) UpdateBigkeyAutoscanConfigInvoker(request *model.UpdateBigkeyAutoscanConfigRequest) *UpdateBigkeyAutoscanConfigInvoker {
	requestDef := GenReqDefForUpdateBigkeyAutoscanConfig()
	return &UpdateBigkeyAutoscanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigurations 修改实例配置参数
//
// 为了确保分布式缓存服务发挥出最优性能，您可以根据自己的业务情况对DCS缓存实例的运行参数进行调整。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateConfigurations(request *model.UpdateConfigurationsRequest) (*model.UpdateConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationsResponse), nil
	}
}

// UpdateConfigurationsInvoker 修改实例配置参数
func (c *DcsClient) UpdateConfigurationsInvoker(request *model.UpdateConfigurationsRequest) *UpdateConfigurationsInvoker {
	requestDef := GenReqDefForUpdateConfigurations()
	return &UpdateConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHotkeyAutoScanConfig 设置热key自动分析配置
//
// 设置热key自动分析配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateHotkeyAutoScanConfig(request *model.UpdateHotkeyAutoScanConfigRequest) (*model.UpdateHotkeyAutoScanConfigResponse, error) {
	requestDef := GenReqDefForUpdateHotkeyAutoScanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotkeyAutoScanConfigResponse), nil
	}
}

// UpdateHotkeyAutoScanConfigInvoker 设置热key自动分析配置
func (c *DcsClient) UpdateHotkeyAutoScanConfigInvoker(request *model.UpdateHotkeyAutoScanConfigRequest) *UpdateHotkeyAutoScanConfigInvoker {
	requestDef := GenReqDefForUpdateHotkeyAutoScanConfig()
	return &UpdateHotkeyAutoScanConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstance 修改实例信息
//
// 修改缓存实例的信息，可修改信息包括实例名称、描述、备份策略、维护时间窗开始和结束时间以及安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// UpdateInstanceInvoker 修改实例信息
func (c *DcsClient) UpdateInstanceInvoker(request *model.UpdateInstanceRequest) *UpdateInstanceInvoker {
	requestDef := GenReqDefForUpdateInstance()
	return &UpdateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceBandwidth 变更指定实例的带宽
//
// 变更指定实例的带宽
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateInstanceBandwidth(request *model.UpdateInstanceBandwidthRequest) (*model.UpdateInstanceBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateInstanceBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceBandwidthResponse), nil
	}
}

// UpdateInstanceBandwidthInvoker 变更指定实例的带宽
func (c *DcsClient) UpdateInstanceBandwidthInvoker(request *model.UpdateInstanceBandwidthRequest) *UpdateInstanceBandwidthInvoker {
	requestDef := GenReqDefForUpdateInstanceBandwidth()
	return &UpdateInstanceBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePassword 修改密码
//
// 修改缓存实例的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdatePassword(request *model.UpdatePasswordRequest) (*model.UpdatePasswordResponse, error) {
	requestDef := GenReqDefForUpdatePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordResponse), nil
	}
}

// UpdatePasswordInvoker 修改密码
func (c *DcsClient) UpdatePasswordInvoker(request *model.UpdatePasswordRequest) *UpdatePasswordInvoker {
	requestDef := GenReqDefForUpdatePassword()
	return &UpdatePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSlavePriority 设置备节点优先级
//
// 设置副本优先级，主节点故障时，权重越小的备节点切换为主节点的优先级越高。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateSlavePriority(request *model.UpdateSlavePriorityRequest) (*model.UpdateSlavePriorityResponse, error) {
	requestDef := GenReqDefForUpdateSlavePriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSlavePriorityResponse), nil
	}
}

// UpdateSlavePriorityInvoker 设置备节点优先级
func (c *DcsClient) UpdateSlavePriorityInvoker(request *model.UpdateSlavePriorityRequest) *UpdateSlavePriorityInvoker {
	requestDef := GenReqDefForUpdateSlavePriority()
	return &UpdateSlavePriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpWhitelist 查询指定实例的IP白名单
//
// 查询指定实例的IP白名单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) ShowIpWhitelist(request *model.ShowIpWhitelistRequest) (*model.ShowIpWhitelistResponse, error) {
	requestDef := GenReqDefForShowIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpWhitelistResponse), nil
	}
}

// ShowIpWhitelistInvoker 查询指定实例的IP白名单
func (c *DcsClient) ShowIpWhitelistInvoker(request *model.ShowIpWhitelistRequest) *ShowIpWhitelistInvoker {
	requestDef := GenReqDefForShowIpWhitelist()
	return &ShowIpWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIpWhitelist 设置IP白名单分组
//
// 为指定实例设置IP白名单分组，包含创建、停用、编辑、删除白名单四个功能
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DcsClient) UpdateIpWhitelist(request *model.UpdateIpWhitelistRequest) (*model.UpdateIpWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpWhitelistResponse), nil
	}
}

// UpdateIpWhitelistInvoker 设置IP白名单分组
func (c *DcsClient) UpdateIpWhitelistInvoker(request *model.UpdateIpWhitelistRequest) *UpdateIpWhitelistInvoker {
	requestDef := GenReqDefForUpdateIpWhitelist()
	return &UpdateIpWhitelistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
