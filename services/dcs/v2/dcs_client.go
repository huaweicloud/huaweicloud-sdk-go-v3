package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"
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

// 批量添加或删除标签
//
// 为指定实例批量添加标签，或批量删除标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) BatchCreateOrDeleteTags(request *model.BatchCreateOrDeleteTagsRequest) (*model.BatchCreateOrDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateOrDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateOrDeleteTagsResponse), nil
	}
}

// 批量删除实例
//
// 批量删除多个缓存实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) BatchDeleteInstances(request *model.BatchDeleteInstancesRequest) (*model.BatchDeleteInstancesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInstancesResponse), nil
	}
}

// 批量查询实例节点信息
//
// 批量查询指定项目所有实例的节点信息、有效实例个数及节点个数。
// 创建中实例不返回节点信息。
// 仅支持Redis4.0和Redis5.0实例查询
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) BatchShowNodesInformation(request *model.BatchShowNodesInformationRequest) (*model.BatchShowNodesInformationResponse, error) {
	requestDef := GenReqDefForBatchShowNodesInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowNodesInformationResponse), nil
	}
}

// 批量停止数据迁移任务
//
// 批量停止数据迁移任务，接口响应成功，仅表示下发任务成功。查询到迁移任务状态为TERMINATED时，即停止成功。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) BatchStopMigrationTasks(request *model.BatchStopMigrationTasksRequest) (*model.BatchStopMigrationTasksResponse, error) {
	requestDef := GenReqDefForBatchStopMigrationTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopMigrationTasksResponse), nil
	}
}

// 主备切换
//
// 切换实例主备节点，只有主备实例支持该操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ChangeMasterStandby(request *model.ChangeMasterStandbyRequest) (*model.ChangeMasterStandbyResponse, error) {
	requestDef := GenReqDefForChangeMasterStandby()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeMasterStandbyResponse), nil
	}
}

// 备份指定实例
//
// 备份指定的缓存实例。
// &gt; 只有主备和集群类型的缓存实例支持备份恢复操作，单机实例不支持备份恢复操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CopyInstance(request *model.CopyInstanceRequest) (*model.CopyInstanceResponse, error) {
	requestDef := GenReqDefForCopyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyInstanceResponse), nil
	}
}

// 创建大key分析任务
//
// 为Redis实例创建大key分析任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateBigkeyScanTask(request *model.CreateBigkeyScanTaskRequest) (*model.CreateBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBigkeyScanTaskResponse), nil
	}
}

// 创建实例诊断任务
//
// 诊断指定的缓存实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateDiagnosisTask(request *model.CreateDiagnosisTaskRequest) (*model.CreateDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskResponse), nil
	}
}

// 创建热key分析任务
//
// 创建热key分析任务，Redis 3.0 不支持热key分析。
//
// 热key分析需要将缓存实例配置参数maxmemory-policy设置为allkeys-lfu或volatile-lfu。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateHotkeyScanTask(request *model.CreateHotkeyScanTaskRequest) (*model.CreateHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForCreateHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHotkeyScanTaskResponse), nil
	}
}

// 创建缓存实例
//
// 创建缓存实例，该接口创建的缓存实例支持按需计费和包周期两种方式。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// 创建数据迁移任务
//
// 创建数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateMigrationTask(request *model.CreateMigrationTaskRequest) (*model.CreateMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMigrationTaskResponse), nil
	}
}

// 创建在线数据迁移任务
//
// 创建在线数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateOnlineMigrationTask(request *model.CreateOnlineMigrationTaskRequest) (*model.CreateOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForCreateOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOnlineMigrationTaskResponse), nil
	}
}

// 采集Redis运行日志
//
// 采集Redis运行日志。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateRedislog(request *model.CreateRedislogRequest) (*model.CreateRedislogResponse, error) {
	requestDef := GenReqDefForCreateRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogResponse), nil
	}
}

// 获取日志下载链接
//
// 获取日志下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) CreateRedislogDownloadLink(request *model.CreateRedislogDownloadLinkRequest) (*model.CreateRedislogDownloadLinkResponse, error) {
	requestDef := GenReqDefForCreateRedislogDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRedislogDownloadLinkResponse), nil
	}
}

// 删除后台任务
//
// 删除后台任务
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteBackgroundTask(request *model.DeleteBackgroundTaskRequest) (*model.DeleteBackgroundTaskResponse, error) {
	requestDef := GenReqDefForDeleteBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackgroundTaskResponse), nil
	}
}

// 删除备份文件
//
// 删除缓存实例已备份的文件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteBackupFile(request *model.DeleteBackupFileRequest) (*model.DeleteBackupFileResponse, error) {
	requestDef := GenReqDefForDeleteBackupFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupFileResponse), nil
	}
}

// 删除大key分析记录
//
// 删除大key分析记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteBigkeyScanTask(request *model.DeleteBigkeyScanTaskRequest) (*model.DeleteBigkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteBigkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBigkeyScanTaskResponse), nil
	}
}

// 删除热key分析任务
//
// 删除热key分析任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteHotkeyScanTask(request *model.DeleteHotkeyScanTaskRequest) (*model.DeleteHotkeyScanTaskResponse, error) {
	requestDef := GenReqDefForDeleteHotkeyScanTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHotkeyScanTaskResponse), nil
	}
}

// 域名摘除IP
//
// 将只读副本的IP从域名中摘除，摘除成功后，只读域名不会再解析到该副本IP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteIpFromDomainName(request *model.DeleteIpFromDomainNameRequest) (*model.DeleteIpFromDomainNameResponse, error) {
	requestDef := GenReqDefForDeleteIpFromDomainName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIpFromDomainNameResponse), nil
	}
}

// 删除数据迁移任务
//
// 删除数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteMigrationTask(request *model.DeleteMigrationTaskRequest) (*model.DeleteMigrationTaskResponse, error) {
	requestDef := GenReqDefForDeleteMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMigrationTaskResponse), nil
	}
}

// 删除实例
//
// 删除指定的缓存实例，释放该实例的所有资源。
//
// &gt; 如果是删除按需资源，请按照本章节执行；如果是删除包周期资源，即退订，请参考[退订包周期资源](https://support.huaweicloud.com/api-oce/zh-cn_topic_0082522030.html#section2)。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) DeleteSingleInstance(request *model.DeleteSingleInstanceRequest) (*model.DeleteSingleInstanceResponse, error) {
	requestDef := GenReqDefForDeleteSingleInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSingleInstanceResponse), nil
	}
}

// 查询可用区信息
//
// 查询所在局点的可用区信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListAvailableZones(request *model.ListAvailableZonesRequest) (*model.ListAvailableZonesResponse, error) {
	requestDef := GenReqDefForListAvailableZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableZonesResponse), nil
	}
}

// 查询后台任务列表
//
// 查询后台任务列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListBackgroundTask(request *model.ListBackgroundTaskRequest) (*model.ListBackgroundTaskResponse, error) {
	requestDef := GenReqDefForListBackgroundTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackgroundTaskResponse), nil
	}
}

// 获取备份文件下载链接
//
// 获取指定实例的备份文件下载链接，下载备份文件。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListBackupFileLinks(request *model.ListBackupFileLinksRequest) (*model.ListBackupFileLinksResponse, error) {
	requestDef := GenReqDefForListBackupFileLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupFileLinksResponse), nil
	}
}

// 查询实例备份信息
//
// 查询指定缓存实例的备份信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListBackupRecords(request *model.ListBackupRecordsRequest) (*model.ListBackupRecordsResponse, error) {
	requestDef := GenReqDefForListBackupRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupRecordsResponse), nil
	}
}

// 查询大key分析任务列表
//
// 查询大key分析任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListBigkeyScanTasks(request *model.ListBigkeyScanTasksRequest) (*model.ListBigkeyScanTasksResponse, error) {
	requestDef := GenReqDefForListBigkeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBigkeyScanTasksResponse), nil
	}
}

// 查询实例配置参数
//
// 查询指定实例的配置参数信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// 查询实例诊断任务列表
//
// 查询指定缓存实例诊断任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListDiagnosisTasks(request *model.ListDiagnosisTasksRequest) (*model.ListDiagnosisTasksResponse, error) {
	requestDef := GenReqDefForListDiagnosisTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisTasksResponse), nil
	}
}

// 查询产品规格
//
// 在创建缓存实例时，需要配置订购的产品规格编码（spec_code），可通过该接口查询产品规格，查询条件不选时默认查询全部。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// 查询分片信息
//
// 查询读写分离实例和集群实例的分片和副本信息，其中，读写分离实例仅Redis4.0和Redis5.0的主备实例支持。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListGroupReplicationInfo(request *model.ListGroupReplicationInfoRequest) (*model.ListGroupReplicationInfoResponse, error) {
	requestDef := GenReqDefForListGroupReplicationInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupReplicationInfoResponse), nil
	}
}

// 查询热key分析任务列表
//
// 查询热key分析历史记录。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListHotKeyScanTasks(request *model.ListHotKeyScanTasksRequest) (*model.ListHotKeyScanTasksResponse, error) {
	requestDef := GenReqDefForListHotKeyScanTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHotKeyScanTasksResponse), nil
	}
}

// 查询所有实例列表
//
// 查询租户的缓存实例列表，支持按照条件查询。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// 查询维护时间窗时间段
//
// 查询维护时间窗开始时间和结束时间。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListMaintenanceWindows(request *model.ListMaintenanceWindowsRequest) (*model.ListMaintenanceWindowsResponse, error) {
	requestDef := GenReqDefForListMaintenanceWindows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMaintenanceWindowsResponse), nil
	}
}

// 查询迁移任务列表
//
// 查询迁移任务列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListMigrationTask(request *model.ListMigrationTaskRequest) (*model.ListMigrationTaskResponse, error) {
	requestDef := GenReqDefForListMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMigrationTaskResponse), nil
	}
}

// 查询主维度信息列表
//
// 查询主维度对象列表，主维度ID当前支持dcs_instance_id，dcs_memcached_instance_id。
// &gt; 该接口当前仅在中国华南区开放。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListMonitoredObjects(request *model.ListMonitoredObjectsRequest) (*model.ListMonitoredObjectsResponse, error) {
	requestDef := GenReqDefForListMonitoredObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsResponse), nil
	}
}

// 查询单个主维度下子维度监控对象列表
//
// 查询主维度下子维度监控对象列表，当前支持子维度的主维度ID的有 dcs_instance_id
// &gt; 该接口当前仅在中国华南区开放。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListMonitoredObjectsOfInstance(request *model.ListMonitoredObjectsOfInstanceRequest) (*model.ListMonitoredObjectsOfInstanceResponse, error) {
	requestDef := GenReqDefForListMonitoredObjectsOfInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMonitoredObjectsOfInstanceResponse), nil
	}
}

// 查询实例状态
//
// 查询该租户在当前区域下不同状态的实例数。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListNumberOfInstancesInDifferentStatus(request *model.ListNumberOfInstancesInDifferentStatusRequest) (*model.ListNumberOfInstancesInDifferentStatusResponse, error) {
	requestDef := GenReqDefForListNumberOfInstancesInDifferentStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNumberOfInstancesInDifferentStatusResponse), nil
	}
}

// 查询Redis运行日志列表
//
// 查询Redis运行日志列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListRedislog(request *model.ListRedislogRequest) (*model.ListRedislogResponse, error) {
	requestDef := GenReqDefForListRedislog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRedislogResponse), nil
	}
}

// 查询实例恢复记录
//
// 查询指定缓存实例的恢复记录列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListRestoreRecords(request *model.ListRestoreRecordsRequest) (*model.ListRestoreRecordsResponse, error) {
	requestDef := GenReqDefForListRestoreRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreRecordsResponse), nil
	}
}

// 查询慢日志
//
// 查询慢日志。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListSlowlog(request *model.ListSlowlogRequest) (*model.ListSlowlogResponse, error) {
	requestDef := GenReqDefForListSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowlogResponse), nil
	}
}

// 查询运行中实例的统计信息
//
// 查询当前租户下处于“运行中”状态的缓存实例的统计信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListStatisticsOfRunningInstances(request *model.ListStatisticsOfRunningInstancesRequest) (*model.ListStatisticsOfRunningInstancesResponse, error) {
	requestDef := GenReqDefForListStatisticsOfRunningInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStatisticsOfRunningInstancesResponse), nil
	}
}

// 查询租户所有标签
//
// 查询租户在指定Project中实例类型的所有资源标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ListTagsOfTenant(request *model.ListTagsOfTenantRequest) (*model.ListTagsOfTenantResponse, error) {
	requestDef := GenReqDefForListTagsOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsOfTenantResponse), nil
	}
}

// 变更实例规格
//
// 用户可以为状态为“运行中”的DCS缓存实例进行规格变更，当前仅能支持按需实例的同副本或分片数量的实例规格变更。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// 重启实例或清空数据
//
// 重启运行中的DCS缓存实例。
//
// 清空Redis4.0/Redis5.0的实例数据，数据清空后，无法撤销，且无法恢复，请谨慎操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) RestartOrFlushInstances(request *model.RestartOrFlushInstancesRequest) (*model.RestartOrFlushInstancesResponse, error) {
	requestDef := GenReqDefForRestartOrFlushInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartOrFlushInstancesResponse), nil
	}
}

// 恢复指定实例
//
// 恢复指定的缓存实例。
// &gt; 只有主备和集群类型的缓存实例支持备份恢复操作，单机实例不支持备份恢复操作。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

// 配置在线数据迁移任务
//
// 配置在线数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) SetOnlineMigrationTask(request *model.SetOnlineMigrationTaskRequest) (*model.SetOnlineMigrationTaskResponse, error) {
	requestDef := GenReqDefForSetOnlineMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetOnlineMigrationTaskResponse), nil
	}
}

// 查询大key自动分析配置
//
// 查询大key自动分析配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowBigkeyAutoscanConfig(request *model.ShowBigkeyAutoscanConfigRequest) (*model.ShowBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyAutoscanConfigResponse), nil
	}
}

// 查询大key分析详情
//
// 查询大key分析详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowBigkeyScanTaskDetails(request *model.ShowBigkeyScanTaskDetailsRequest) (*model.ShowBigkeyScanTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowBigkeyScanTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBigkeyScanTaskDetailsResponse), nil
	}
}

// 查询指定诊断报告
//
// 通过报告ID查询诊断报告的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowDiagnosisTaskDetails(request *model.ShowDiagnosisTaskDetailsRequest) (*model.ShowDiagnosisTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowDiagnosisTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisTaskDetailsResponse), nil
	}
}

// 查询热key自动分析配置
//
// 查询热key自动分析配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowHotkeyAutoscanConfig(request *model.ShowHotkeyAutoscanConfigRequest) (*model.ShowHotkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForShowHotkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyAutoscanConfigResponse), nil
	}
}

// 查询热key分析详情
//
// 查询热key分析详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowHotkeyTaskDetails(request *model.ShowHotkeyTaskDetailsRequest) (*model.ShowHotkeyTaskDetailsResponse, error) {
	requestDef := GenReqDefForShowHotkeyTaskDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHotkeyTaskDetailsResponse), nil
	}
}

// 查询指定实例
//
// 通过实例ID查询实例的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// 查询迁移任务详情
//
// 查询迁移任务详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowMigrationTask(request *model.ShowMigrationTaskRequest) (*model.ShowMigrationTaskResponse, error) {
	requestDef := GenReqDefForShowMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskResponse), nil
	}
}

// 查询在线迁移进度明细
//
// 查询在线迁移进度明细。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowMigrationTaskStats(request *model.ShowMigrationTaskStatsRequest) (*model.ShowMigrationTaskStatsResponse, error) {
	requestDef := GenReqDefForShowMigrationTaskStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationTaskStatsResponse), nil
	}
}

// 查询租户配额
//
// 查询租户默认可以创建的实例数和总内存的配额限制，以及可以申请配额的最大值和最小值。不同的租户在不同的区域配额可能不同。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowQuotaOfTenant(request *model.ShowQuotaOfTenantRequest) (*model.ShowQuotaOfTenantResponse, error) {
	requestDef := GenReqDefForShowQuotaOfTenant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaOfTenantResponse), nil
	}
}

// 查询单个实例标签
//
// 通过实例ID查询标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowTags(request *model.ShowTagsRequest) (*model.ShowTagsResponse, error) {
	requestDef := GenReqDefForShowTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagsResponse), nil
	}
}

// 停止数据迁移任务
//
// 停止数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) StopMigrationTask(request *model.StopMigrationTaskRequest) (*model.StopMigrationTaskResponse, error) {
	requestDef := GenReqDefForStopMigrationTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskResponse), nil
	}
}

// 同步停止数据迁移任务
//
// 同步停止数据迁移任务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) StopMigrationTaskSync(request *model.StopMigrationTaskSyncRequest) (*model.StopMigrationTaskSyncResponse, error) {
	requestDef := GenReqDefForStopMigrationTaskSync()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopMigrationTaskSyncResponse), nil
	}
}

// 设置大key自动分析配置
//
// 设置大key自动分析配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateBigkeyAutoscanConfig(request *model.UpdateBigkeyAutoscanConfigRequest) (*model.UpdateBigkeyAutoscanConfigResponse, error) {
	requestDef := GenReqDefForUpdateBigkeyAutoscanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBigkeyAutoscanConfigResponse), nil
	}
}

// 修改实例配置参数
//
// 为了确保分布式缓存服务发挥出最优性能，您可以根据自己的业务情况对DCS缓存实例的运行参数进行调整。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateConfigurations(request *model.UpdateConfigurationsRequest) (*model.UpdateConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationsResponse), nil
	}
}

// 设置热key自动分析配置
//
// 设置热key自动分析配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateHotkeyAutoScanConfig(request *model.UpdateHotkeyAutoScanConfigRequest) (*model.UpdateHotkeyAutoScanConfigResponse, error) {
	requestDef := GenReqDefForUpdateHotkeyAutoScanConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHotkeyAutoScanConfigResponse), nil
	}
}

// 修改实例信息
//
// 修改缓存实例的信息，可修改信息包括实例名称、描述、备份策略、维护时间窗开始和结束时间以及安全组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateInstance(request *model.UpdateInstanceRequest) (*model.UpdateInstanceResponse, error) {
	requestDef := GenReqDefForUpdateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceResponse), nil
	}
}

// 修改密码
//
// 修改缓存实例的密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdatePassword(request *model.UpdatePasswordRequest) (*model.UpdatePasswordResponse, error) {
	requestDef := GenReqDefForUpdatePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePasswordResponse), nil
	}
}

// 设置备节点优先级
//
// 设置副本优先级，主节点故障时，权重越小的备节点切换为主节点的优先级越高。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateSlavePriority(request *model.UpdateSlavePriorityRequest) (*model.UpdateSlavePriorityResponse, error) {
	requestDef := GenReqDefForUpdateSlavePriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSlavePriorityResponse), nil
	}
}

// 查询指定实例的IP白名单
//
// 查询指定实例的IP白名单。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) ShowIpWhitelist(request *model.ShowIpWhitelistRequest) (*model.ShowIpWhitelistResponse, error) {
	requestDef := GenReqDefForShowIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpWhitelistResponse), nil
	}
}

// 设置IP白名单分组
//
// 为指定实例设置IP白名单分组，包含创建、停用、编辑、删除白名单四个功能
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DcsClient) UpdateIpWhitelist(request *model.UpdateIpWhitelistRequest) (*model.UpdateIpWhitelistResponse, error) {
	requestDef := GenReqDefForUpdateIpWhitelist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIpWhitelistResponse), nil
	}
}
