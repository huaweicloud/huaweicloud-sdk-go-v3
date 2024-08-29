package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdbforopengauss/v3/model"
)

type GaussDBforopenGaussClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGaussDBforopenGaussClient(hcClient *httpclient.HcHttpClient) *GaussDBforopenGaussClient {
	return &GaussDBforopenGaussClient{HcClient: hcClient}
}

func GaussDBforopenGaussClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddInstanceTags 添加实例标签。
//
// 对指定实例添加用户标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AddInstanceTags(request *model.AddInstanceTagsRequest) (*model.AddInstanceTagsResponse, error) {
	requestDef := GenReqDefForAddInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddInstanceTagsResponse), nil
	}
}

// AddInstanceTagsInvoker 添加实例标签。
func (c *GaussDBforopenGaussClient) AddInstanceTagsInvoker(request *model.AddInstanceTagsRequest) *AddInstanceTagsInvoker {
	requestDef := GenReqDefForAddInstanceTags()
	return &AddInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AllowDbPrivileges 授权数据库帐号
//
// 在指定实例的数据库中, 设置帐号的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AllowDbPrivileges(request *model.AllowDbPrivilegesRequest) (*model.AllowDbPrivilegesResponse, error) {
	requestDef := GenReqDefForAllowDbPrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowDbPrivilegesResponse), nil
	}
}

// AllowDbPrivilegesInvoker 授权数据库帐号
func (c *GaussDBforopenGaussClient) AllowDbPrivilegesInvoker(request *model.AllowDbPrivilegesRequest) *AllowDbPrivilegesInvoker {
	requestDef := GenReqDefForAllowDbPrivileges()
	return &AllowDbPrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachEip 绑定/解绑弹性公网IP
//
// 实例下的节点绑定弹性公网IP/解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

// AttachEipInvoker 绑定/解绑弹性公网IP
func (c *GaussDBforopenGaussClient) AttachEipInvoker(request *model.AttachEipRequest) *AttachEipInvoker {
	requestDef := GenReqDefForAttachEip()
	return &AttachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmRestoredData 备份恢复到目标实例数据后执行数据确认
//
// 确认备份恢复到目标实例的数据正常。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ConfirmRestoredData(request *model.ConfirmRestoredDataRequest) (*model.ConfirmRestoredDataResponse, error) {
	requestDef := GenReqDefForConfirmRestoredData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmRestoredDataResponse), nil
	}
}

// ConfirmRestoredDataInvoker 备份恢复到目标实例数据后执行数据确认
func (c *GaussDBforopenGaussClient) ConfirmRestoredDataInvoker(request *model.ConfirmRestoredDataRequest) *ConfirmRestoredDataInvoker {
	requestDef := GenReqDefForConfirmRestoredData()
	return &ConfirmRestoredDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyConfiguration 复制参数模板
//
// 复制参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

// CopyConfigurationInvoker 复制参数模板
func (c *GaussDBforopenGaussClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigurationTemplate 创建参数模板
//
// 创建参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateConfigurationTemplate(request *model.CreateConfigurationTemplateRequest) (*model.CreateConfigurationTemplateResponse, error) {
	requestDef := GenReqDefForCreateConfigurationTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationTemplateResponse), nil
	}
}

// CreateConfigurationTemplateInvoker 创建参数模板
func (c *GaussDBforopenGaussClient) CreateConfigurationTemplateInvoker(request *model.CreateConfigurationTemplateRequest) *CreateConfigurationTemplateInvoker {
	requestDef := GenReqDefForCreateConfigurationTemplate()
	return &CreateConfigurationTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabase 创建数据库
//
// 在指定实例中创建数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDatabase(request *model.CreateDatabaseRequest) (*model.CreateDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResponse), nil
	}
}

// CreateDatabaseInvoker 创建数据库
func (c *GaussDBforopenGaussClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseSchemas 创建数据库SCHEMA
//
// 在指定实例的数据库中, 创建数据库schema。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDatabaseSchemas(request *model.CreateDatabaseSchemasRequest) (*model.CreateDatabaseSchemasResponse, error) {
	requestDef := GenReqDefForCreateDatabaseSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseSchemasResponse), nil
	}
}

// CreateDatabaseSchemasInvoker 创建数据库SCHEMA
func (c *GaussDBforopenGaussClient) CreateDatabaseSchemasInvoker(request *model.CreateDatabaseSchemasRequest) *CreateDatabaseSchemasInvoker {
	requestDef := GenReqDefForCreateDatabaseSchemas()
	return &CreateDatabaseSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDbInstance 创建数据库实例
//
// 创建数据库实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDbInstance(request *model.CreateDbInstanceRequest) (*model.CreateDbInstanceResponse, error) {
	requestDef := GenReqDefForCreateDbInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbInstanceResponse), nil
	}
}

// CreateDbInstanceInvoker 创建数据库实例
func (c *GaussDBforopenGaussClient) CreateDbInstanceInvoker(request *model.CreateDbInstanceRequest) *CreateDbInstanceInvoker {
	requestDef := GenReqDefForCreateDbInstance()
	return &CreateDbInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDbUser 创建数据库用户
//
// 在指定实例中创建数据库用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDbUser(request *model.CreateDbUserRequest) (*model.CreateDbUserResponse, error) {
	requestDef := GenReqDefForCreateDbUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbUserResponse), nil
	}
}

// CreateDbUserInvoker 创建数据库用户
func (c *GaussDBforopenGaussClient) CreateDbUserInvoker(request *model.CreateDbUserRequest) *CreateDbUserInvoker {
	requestDef := GenReqDefForCreateDbUser()
	return &CreateDbUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussDbInstance 创建数据库实例
//
// 创建数据库实例，仅支持IAM5的新平面认证方式（AK/SK认证方式）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateGaussDbInstance(request *model.CreateGaussDbInstanceRequest) (*model.CreateGaussDbInstanceResponse, error) {
	requestDef := GenReqDefForCreateGaussDbInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussDbInstanceResponse), nil
	}
}

// CreateGaussDbInstanceInvoker 创建数据库实例
func (c *GaussDBforopenGaussClient) CreateGaussDbInstanceInvoker(request *model.CreateGaussDbInstanceRequest) *CreateGaussDbInstanceInvoker {
	requestDef := GenReqDefForCreateGaussDbInstance()
	return &CreateGaussDbInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建数据库实例
//
// 创建数据库企业版和集中式实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建数据库实例
func (c *GaussDBforopenGaussClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManualBackup 创建手动备份
//
// 创建手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateManualBackup(request *model.CreateManualBackupRequest) (*model.CreateManualBackupResponse, error) {
	requestDef := GenReqDefForCreateManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualBackupResponse), nil
	}
}

// CreateManualBackupInvoker 创建手动备份
func (c *GaussDBforopenGaussClient) CreateManualBackupInvoker(request *model.CreateManualBackupRequest) *CreateManualBackupInvoker {
	requestDef := GenReqDefForCreateManualBackup()
	return &CreateManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRestoreInstance 恢复到新实例
//
// 根据备份恢复新实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateRestoreInstance(request *model.CreateRestoreInstanceRequest) (*model.CreateRestoreInstanceResponse, error) {
	requestDef := GenReqDefForCreateRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRestoreInstanceResponse), nil
	}
}

// CreateRestoreInstanceInvoker 恢复到新实例
func (c *GaussDBforopenGaussClient) CreateRestoreInstanceInvoker(request *model.CreateRestoreInstanceRequest) *CreateRestoreInstanceInvoker {
	requestDef := GenReqDefForCreateRestoreInstance()
	return &CreateRestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSlowLogDownload 创建慢日志下载信息
//
// 创建慢日志下载信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateSlowLogDownload(request *model.CreateSlowLogDownloadRequest) (*model.CreateSlowLogDownloadResponse, error) {
	requestDef := GenReqDefForCreateSlowLogDownload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSlowLogDownloadResponse), nil
	}
}

// CreateSlowLogDownloadInvoker 创建慢日志下载信息
func (c *GaussDBforopenGaussClient) CreateSlowLogDownloadInvoker(request *model.CreateSlowLogDownloadRequest) *CreateSlowLogDownloadInvoker {
	requestDef := GenReqDefForCreateSlowLogDownload()
	return &CreateSlowLogDownloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfiguration 删除参数模板
//
// 删除参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

// DeleteConfigurationInvoker 删除参数模板
func (c *GaussDBforopenGaussClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabase 删除数据库
//
// 删除指定实例的数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteDatabase(request *model.DeleteDatabaseRequest) (*model.DeleteDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResponse), nil
	}
}

// DeleteDatabaseInvoker 删除数据库
func (c *GaussDBforopenGaussClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除数据库实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *GaussDBforopenGaussClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstanceTag 删除实例标签
//
// 删除实例标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteInstanceTag(request *model.DeleteInstanceTagRequest) (*model.DeleteInstanceTagResponse, error) {
	requestDef := GenReqDefForDeleteInstanceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceTagResponse), nil
	}
}

// DeleteInstanceTagInvoker 删除实例标签
func (c *GaussDBforopenGaussClient) DeleteInstanceTagInvoker(request *model.DeleteInstanceTagRequest) *DeleteInstanceTagInvoker {
	requestDef := GenReqDefForDeleteInstanceTag()
	return &DeleteInstanceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJob 删除任务记录
//
// 删除任务记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteJob(request *model.DeleteJobRequest) (*model.DeleteJobResponse, error) {
	requestDef := GenReqDefForDeleteJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResponse), nil
	}
}

// DeleteJobInvoker 删除任务记录
func (c *GaussDBforopenGaussClient) DeleteJobInvoker(request *model.DeleteJobRequest) *DeleteJobInvoker {
	requestDef := GenReqDefForDeleteJob()
	return &DeleteJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteManualBackup 删除手动备份
//
// 删除手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteManualBackup(request *model.DeleteManualBackupRequest) (*model.DeleteManualBackupResponse, error) {
	requestDef := GenReqDefForDeleteManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManualBackupResponse), nil
	}
}

// DeleteManualBackupInvoker 删除手动备份
func (c *GaussDBforopenGaussClient) DeleteManualBackupInvoker(request *model.DeleteManualBackupRequest) *DeleteManualBackupInvoker {
	requestDef := GenReqDefForDeleteManualBackup()
	return &DeleteManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadBackup 获取备份下载链接
//
// 获取备份下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DownloadBackup(request *model.DownloadBackupRequest) (*model.DownloadBackupResponse, error) {
	requestDef := GenReqDefForDownloadBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadBackupResponse), nil
	}
}

// DownloadBackupInvoker 获取备份下载链接
func (c *GaussDBforopenGaussClient) DownloadBackupInvoker(request *model.DownloadBackupRequest) *DownloadBackupInvoker {
	requestDef := GenReqDefForDownloadBackup()
	return &DownloadBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InstallKernelPlugin 安装插件
//
// 安装插件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) InstallKernelPlugin(request *model.InstallKernelPluginRequest) (*model.InstallKernelPluginResponse, error) {
	requestDef := GenReqDefForInstallKernelPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InstallKernelPluginResponse), nil
	}
}

// InstallKernelPluginInvoker 安装插件
func (c *GaussDBforopenGaussClient) InstallKernelPluginInvoker(request *model.InstallKernelPluginRequest) *InstallKernelPluginInvoker {
	requestDef := GenReqDefForInstallKernelPlugin()
	return &InstallKernelPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicableInstances 查询可应用实例列表
//
// 查询可应用当前参数组模板的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListApplicableInstances(request *model.ListApplicableInstancesRequest) (*model.ListApplicableInstancesResponse, error) {
	requestDef := GenReqDefForListApplicableInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicableInstancesResponse), nil
	}
}

// ListApplicableInstancesInvoker 查询可应用实例列表
func (c *GaussDBforopenGaussClient) ListApplicableInstancesInvoker(request *model.ListApplicableInstancesRequest) *ListApplicableInstancesInvoker {
	requestDef := GenReqDefForListApplicableInstances()
	return &ListApplicableInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppliedHistories 查询参数模板的应用记录
//
// 查询参数模板的应用记录，以实例级别为维度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListAppliedHistories(request *model.ListAppliedHistoriesRequest) (*model.ListAppliedHistoriesResponse, error) {
	requestDef := GenReqDefForListAppliedHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppliedHistoriesResponse), nil
	}
}

// ListAppliedHistoriesInvoker 查询参数模板的应用记录
func (c *GaussDBforopenGaussClient) ListAppliedHistoriesInvoker(request *model.ListAppliedHistoriesRequest) *ListAppliedHistoriesInvoker {
	requestDef := GenReqDefForListAppliedHistories()
	return &ListAppliedHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableFlavors 查询实例可变更规格
//
// 查询实例可变更规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListAvailableFlavors(request *model.ListAvailableFlavorsRequest) (*model.ListAvailableFlavorsResponse, error) {
	requestDef := GenReqDefForListAvailableFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableFlavorsResponse), nil
	}
}

// ListAvailableFlavorsInvoker 查询实例可变更规格
func (c *GaussDBforopenGaussClient) ListAvailableFlavorsInvoker(request *model.ListAvailableFlavorsRequest) *ListAvailableFlavorsInvoker {
	requestDef := GenReqDefForListAvailableFlavors()
	return &ListAvailableFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups 查询备份列表
//
// 获取备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// ListBackupsInvoker 查询备份列表
func (c *GaussDBforopenGaussClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBindedEips 查询实例已绑定EIP列表
//
// 查询实例已绑定EIP列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListBindedEips(request *model.ListBindedEipsRequest) (*model.ListBindedEipsResponse, error) {
	requestDef := GenReqDefForListBindedEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBindedEipsResponse), nil
	}
}

// ListBindedEipsInvoker 查询实例已绑定EIP列表
func (c *GaussDBforopenGaussClient) ListBindedEipsInvoker(request *model.ListBindedEipsRequest) *ListBindedEipsInvoker {
	requestDef := GenReqDefForListBindedEips()
	return &ListBindedEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCnInfosBeforeReduce 查询协调节点列表
//
// 查询协调节点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListCnInfosBeforeReduce(request *model.ListCnInfosBeforeReduceRequest) (*model.ListCnInfosBeforeReduceResponse, error) {
	requestDef := GenReqDefForListCnInfosBeforeReduce()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCnInfosBeforeReduceResponse), nil
	}
}

// ListCnInfosBeforeReduceInvoker 查询协调节点列表
func (c *GaussDBforopenGaussClient) ListCnInfosBeforeReduceInvoker(request *model.ListCnInfosBeforeReduceRequest) *ListCnInfosBeforeReduceInvoker {
	requestDef := GenReqDefForListCnInfosBeforeReduce()
	return &ListCnInfosBeforeReduceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListComponentInfos 查询实例的组件列表
//
// 查询实例的组件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListComponentInfos(request *model.ListComponentInfosRequest) (*model.ListComponentInfosResponse, error) {
	requestDef := GenReqDefForListComponentInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListComponentInfosResponse), nil
	}
}

// ListComponentInfosInvoker 查询实例的组件列表
func (c *GaussDBforopenGaussClient) ListComponentInfosInvoker(request *model.ListComponentInfosRequest) *ListComponentInfosInvoker {
	requestDef := GenReqDefForListComponentInfos()
	return &ListComponentInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurations 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// ListConfigurationsInvoker 获取参数模板列表
func (c *GaussDBforopenGaussClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationsDiff 比较两个参数组模板之间的差异
//
// 获取两个参数配置模板的差异列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListConfigurationsDiff(request *model.ListConfigurationsDiffRequest) (*model.ListConfigurationsDiffResponse, error) {
	requestDef := GenReqDefForListConfigurationsDiff()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsDiffResponse), nil
	}
}

// ListConfigurationsDiffInvoker 比较两个参数组模板之间的差异
func (c *GaussDBforopenGaussClient) ListConfigurationsDiffInvoker(request *model.ListConfigurationsDiffRequest) *ListConfigurationsDiffInvoker {
	requestDef := GenReqDefForListConfigurationsDiff()
	return &ListConfigurationsDiffInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseSchemas 查询数据库SCHEMA列表
//
// 查询指定实例的数据库SCHEMA列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseSchemas(request *model.ListDatabaseSchemasRequest) (*model.ListDatabaseSchemasResponse, error) {
	requestDef := GenReqDefForListDatabaseSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseSchemasResponse), nil
	}
}

// ListDatabaseSchemasInvoker 查询数据库SCHEMA列表
func (c *GaussDBforopenGaussClient) ListDatabaseSchemasInvoker(request *model.ListDatabaseSchemasRequest) *ListDatabaseSchemasInvoker {
	requestDef := GenReqDefForListDatabaseSchemas()
	return &ListDatabaseSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabases 查询数据库列表
//
// 查询指定实例中的数据库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 查询数据库列表
func (c *GaussDBforopenGaussClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatastores 查询数据库引擎的版本
//
// 查询指定数据库引擎对应的版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatastores(request *model.ListDatastoresRequest) (*model.ListDatastoresResponse, error) {
	requestDef := GenReqDefForListDatastores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoresResponse), nil
	}
}

// ListDatastoresInvoker 查询数据库引擎的版本
func (c *GaussDBforopenGaussClient) ListDatastoresInvoker(request *model.ListDatastoresRequest) *ListDatastoresInvoker {
	requestDef := GenReqDefForListDatastores()
	return &ListDatastoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbBackups 查询备份列表
//
// 获取备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDbBackups(request *model.ListDbBackupsRequest) (*model.ListDbBackupsResponse, error) {
	requestDef := GenReqDefForListDbBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbBackupsResponse), nil
	}
}

// ListDbBackupsInvoker 查询备份列表
func (c *GaussDBforopenGaussClient) ListDbBackupsInvoker(request *model.ListDbBackupsRequest) *ListDbBackupsInvoker {
	requestDef := GenReqDefForListDbBackups()
	return &ListDbBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbFlavors 查询数据库规格
//
// 查询数据库的规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDbFlavors(request *model.ListDbFlavorsRequest) (*model.ListDbFlavorsResponse, error) {
	requestDef := GenReqDefForListDbFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbFlavorsResponse), nil
	}
}

// ListDbFlavorsInvoker 查询数据库规格
func (c *GaussDBforopenGaussClient) ListDbFlavorsInvoker(request *model.ListDbFlavorsRequest) *ListDbFlavorsInvoker {
	requestDef := GenReqDefForListDbFlavors()
	return &ListDbFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDbUsers 查询数据库用户列表
//
// 在指定实例中查询数据库用户列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDbUsers(request *model.ListDbUsersRequest) (*model.ListDbUsersResponse, error) {
	requestDef := GenReqDefForListDbUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDbUsersResponse), nil
	}
}

// ListDbUsersInvoker 查询数据库用户列表
func (c *GaussDBforopenGaussClient) ListDbUsersInvoker(request *model.ListDbUsersRequest) *ListDbUsersInvoker {
	requestDef := GenReqDefForListDbUsers()
	return &ListDbUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEpsQuotas 查询企业项目配额组
//
// 查询企业项目配额组信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListEpsQuotas(request *model.ListEpsQuotasRequest) (*model.ListEpsQuotasResponse, error) {
	requestDef := GenReqDefForListEpsQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEpsQuotasResponse), nil
	}
}

// ListEpsQuotasInvoker 查询企业项目配额组
func (c *GaussDBforopenGaussClient) ListEpsQuotasInvoker(request *model.ListEpsQuotasRequest) *ListEpsQuotasInvoker {
	requestDef := GenReqDefForListEpsQuotas()
	return &ListEpsQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询数据库规格
//
// 查询数据库的规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询数据库规格
func (c *GaussDBforopenGaussClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussDbDatastores 查询引擎列表
//
// 查询引擎列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListGaussDbDatastores(request *model.ListGaussDbDatastoresRequest) (*model.ListGaussDbDatastoresResponse, error) {
	requestDef := GenReqDefForListGaussDbDatastores()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussDbDatastoresResponse), nil
	}
}

// ListGaussDbDatastoresInvoker 查询引擎列表
func (c *GaussDBforopenGaussClient) ListGaussDbDatastoresInvoker(request *model.ListGaussDbDatastoresRequest) *ListGaussDbDatastoresInvoker {
	requestDef := GenReqDefForListGaussDbDatastores()
	return &ListGaussDbDatastoresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryOperations 查询参数模板的修改历史
//
// 查询参数模板的修改历史记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListHistoryOperations(request *model.ListHistoryOperationsRequest) (*model.ListHistoryOperationsResponse, error) {
	requestDef := GenReqDefForListHistoryOperations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryOperationsResponse), nil
	}
}

// ListHistoryOperationsInvoker 查询参数模板的修改历史
func (c *GaussDBforopenGaussClient) ListHistoryOperationsInvoker(request *model.ListHistoryOperationsRequest) *ListHistoryOperationsInvoker {
	requestDef := GenReqDefForListHistoryOperations()
	return &ListHistoryOperationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceDetails 查询数据库实例列表/查询实例详情
//
// 查询数据库实例列表/查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstanceDetails(request *model.ListInstanceDetailsRequest) (*model.ListInstanceDetailsResponse, error) {
	requestDef := GenReqDefForListInstanceDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceDetailsResponse), nil
	}
}

// ListInstanceDetailsInvoker 查询数据库实例列表/查询实例详情
func (c *GaussDBforopenGaussClient) ListInstanceDetailsInvoker(request *model.ListInstanceDetailsRequest) *ListInstanceDetailsInvoker {
	requestDef := GenReqDefForListInstanceDetails()
	return &ListInstanceDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceErrorLogs 查询错误日志下载链接
//
// 查询数据库错误日志下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstanceErrorLogs(request *model.ListInstanceErrorLogsRequest) (*model.ListInstanceErrorLogsResponse, error) {
	requestDef := GenReqDefForListInstanceErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceErrorLogsResponse), nil
	}
}

// ListInstanceErrorLogsInvoker 查询错误日志下载链接
func (c *GaussDBforopenGaussClient) ListInstanceErrorLogsInvoker(request *model.ListInstanceErrorLogsRequest) *ListInstanceErrorLogsInvoker {
	requestDef := GenReqDefForListInstanceErrorLogs()
	return &ListInstanceErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询实例标签
//
// 查询指定实例的用户标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询实例标签
func (c *GaussDBforopenGaussClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询数据库实例列表/查询实例详情
//
// 查询数据库实例列表/查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询数据库实例列表/查询实例详情
func (c *GaussDBforopenGaussClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesDetails 查询数据库实例列表/查询实例详情
//
// 查询数据库实例列表/查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstancesDetails(request *model.ListInstancesDetailsRequest) (*model.ListInstancesDetailsResponse, error) {
	requestDef := GenReqDefForListInstancesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesDetailsResponse), nil
	}
}

// ListInstancesDetailsInvoker 查询数据库实例列表/查询实例详情
func (c *GaussDBforopenGaussClient) ListInstancesDetailsInvoker(request *model.ListInstancesDetailsRequest) *ListInstancesDetailsInvoker {
	requestDef := GenReqDefForListInstancesDetails()
	return &ListInstancesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKernelPlugins 查询实例已安装的插件列表
//
// 查询实例已安装的插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListKernelPlugins(request *model.ListKernelPluginsRequest) (*model.ListKernelPluginsResponse, error) {
	requestDef := GenReqDefForListKernelPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKernelPluginsResponse), nil
	}
}

// ListKernelPluginsInvoker 查询实例已安装的插件列表
func (c *GaussDBforopenGaussClient) ListKernelPluginsInvoker(request *model.ListKernelPluginsRequest) *ListKernelPluginsInvoker {
	requestDef := GenReqDefForListKernelPlugins()
	return &ListKernelPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListParamGroupTemplates 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListParamGroupTemplates(request *model.ListParamGroupTemplatesRequest) (*model.ListParamGroupTemplatesResponse, error) {
	requestDef := GenReqDefForListParamGroupTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListParamGroupTemplatesResponse), nil
	}
}

// ListParamGroupTemplatesInvoker 获取参数模板列表
func (c *GaussDBforopenGaussClient) ListParamGroupTemplatesInvoker(request *model.ListParamGroupTemplatesRequest) *ListParamGroupTemplatesInvoker {
	requestDef := GenReqDefForListParamGroupTemplates()
	return &ListParamGroupTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPluginExtensions 查询实例插件拓展信息
//
// 查询实例插件拓展信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListPluginExtensions(request *model.ListPluginExtensionsRequest) (*model.ListPluginExtensionsResponse, error) {
	requestDef := GenReqDefForListPluginExtensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginExtensionsResponse), nil
	}
}

// ListPluginExtensionsInvoker 查询实例插件拓展信息
func (c *GaussDBforopenGaussClient) ListPluginExtensionsInvoker(request *model.ListPluginExtensionsRequest) *ListPluginExtensionsInvoker {
	requestDef := GenReqDefForListPluginExtensions()
	return &ListPluginExtensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPredefinedTags 查询预定义标签
//
// 查询预预定义标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListPredefinedTags(request *model.ListPredefinedTagsRequest) (*model.ListPredefinedTagsResponse, error) {
	requestDef := GenReqDefForListPredefinedTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPredefinedTagsResponse), nil
	}
}

// ListPredefinedTagsInvoker 查询预定义标签
func (c *GaussDBforopenGaussClient) ListPredefinedTagsInvoker(request *model.ListPredefinedTagsRequest) *ListPredefinedTagsInvoker {
	requestDef := GenReqDefForListPredefinedTags()
	return &ListPredefinedTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询项目下所有用户标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *GaussDBforopenGaussClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecycleInstances 查询回收站所有引擎实例列表。
//
// 查询回收站所有引擎实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

// ListRecycleInstancesInvoker 查询回收站所有引擎实例列表。
func (c *GaussDBforopenGaussClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestorableInstances 查询可用于备份恢复的实例列表
//
// 查询可用于备份恢复的实例列表，实例信息要符合备份条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRestorableInstances(request *model.ListRestorableInstancesRequest) (*model.ListRestorableInstancesResponse, error) {
	requestDef := GenReqDefForListRestorableInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestorableInstancesResponse), nil
	}
}

// ListRestorableInstancesInvoker 查询可用于备份恢复的实例列表
func (c *GaussDBforopenGaussClient) ListRestorableInstancesInvoker(request *model.ListRestorableInstancesRequest) *ListRestorableInstancesInvoker {
	requestDef := GenReqDefForListRestorableInstances()
	return &ListRestorableInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreTimes 查询可恢复时间段
//
// 查询可恢复时间段。
// 如果您备份策略中的保存天数设置较长，建议您传入查询日期“date”。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRestoreTimes(request *model.ListRestoreTimesRequest) (*model.ListRestoreTimesResponse, error) {
	requestDef := GenReqDefForListRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimesResponse), nil
	}
}

// ListRestoreTimesInvoker 查询可恢复时间段
func (c *GaussDBforopenGaussClient) ListRestoreTimesInvoker(request *model.ListRestoreTimesRequest) *ListRestoreTimesInvoker {
	requestDef := GenReqDefForListRestoreTimes()
	return &ListRestoreTimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageTypes 查询数据库磁盘类型
//
// 查询指定数据库引擎对应的磁盘类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListStorageTypes(request *model.ListStorageTypesRequest) (*model.ListStorageTypesResponse, error) {
	requestDef := GenReqDefForListStorageTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypesResponse), nil
	}
}

// ListStorageTypesInvoker 查询数据库磁盘类型
func (c *GaussDBforopenGaussClient) ListStorageTypesInvoker(request *model.ListStorageTypesRequest) *ListStorageTypesInvoker {
	requestDef := GenReqDefForListStorageTypes()
	return &ListStorageTypesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportKernelPlugins 查询支持的插件列表
//
// 查询支持的插件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSupportKernelPlugins(request *model.ListSupportKernelPluginsRequest) (*model.ListSupportKernelPluginsResponse, error) {
	requestDef := GenReqDefForListSupportKernelPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportKernelPluginsResponse), nil
	}
}

// ListSupportKernelPluginsInvoker 查询支持的插件列表
func (c *GaussDBforopenGaussClient) ListSupportKernelPluginsInvoker(request *model.ListSupportKernelPluginsRequest) *ListSupportKernelPluginsInvoker {
	requestDef := GenReqDefForListSupportKernelPlugins()
	return &ListSupportKernelPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询任务列表
//
// 获取任务中心的任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表
func (c *GaussDBforopenGaussClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopIoTraffics 查询Top IO列表
//
// 查询实例数据库进程下的Top IO流量数据，返回与会话信息相关联后的结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTopIoTraffics(request *model.ListTopIoTrafficsRequest) (*model.ListTopIoTrafficsResponse, error) {
	requestDef := GenReqDefForListTopIoTraffics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopIoTrafficsResponse), nil
	}
}

// ListTopIoTrafficsInvoker 查询Top IO列表
func (c *GaussDBforopenGaussClient) ListTopIoTrafficsInvoker(request *model.ListTopIoTrafficsRequest) *ListTopIoTrafficsInvoker {
	requestDef := GenReqDefForListTopIoTraffics()
	return &ListTopIoTrafficsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyEpsQuota 修改企业项目配额
//
// 修改企业项目配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ModifyEpsQuota(request *model.ModifyEpsQuotaRequest) (*model.ModifyEpsQuotaResponse, error) {
	requestDef := GenReqDefForModifyEpsQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyEpsQuotaResponse), nil
	}
}

// ModifyEpsQuotaInvoker 修改企业项目配额
func (c *GaussDBforopenGaussClient) ModifyEpsQuotaInvoker(request *model.ModifyEpsQuotaRequest) *ModifyEpsQuotaInvoker {
	requestDef := GenReqDefForModifyEpsQuota()
	return &ModifyEpsQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetConfiguration 重置参数模板
//
// 重置参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ResetConfiguration(request *model.ResetConfigurationRequest) (*model.ResetConfigurationResponse, error) {
	requestDef := GenReqDefForResetConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConfigurationResponse), nil
	}
}

// ResetConfigurationInvoker 重置参数模板
func (c *GaussDBforopenGaussClient) ResetConfigurationInvoker(request *model.ResetConfigurationRequest) *ResetConfigurationInvoker {
	requestDef := GenReqDefForResetConfiguration()
	return &ResetConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPwd 重置数据库密码。
//
// 重置数据库密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ResetPwd(request *model.ResetPwdRequest) (*model.ResetPwdResponse, error) {
	requestDef := GenReqDefForResetPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPwdResponse), nil
	}
}

// ResetPwdInvoker 重置数据库密码。
func (c *GaussDBforopenGaussClient) ResetPwdInvoker(request *model.ResetPwdRequest) *ResetPwdInvoker {
	requestDef := GenReqDefForResetPwd()
	return &ResetPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstanceFlavor GaussDB数据库实例规格变更
//
// # GaussDB数据库实例规格变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ResizeInstanceFlavor(request *model.ResizeInstanceFlavorRequest) (*model.ResizeInstanceFlavorResponse, error) {
	requestDef := GenReqDefForResizeInstanceFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceFlavorResponse), nil
	}
}

// ResizeInstanceFlavorInvoker GaussDB数据库实例规格变更
func (c *GaussDBforopenGaussClient) ResizeInstanceFlavorInvoker(request *model.ResizeInstanceFlavorRequest) *ResizeInstanceFlavorInvoker {
	requestDef := GenReqDefForResizeInstanceFlavor()
	return &ResizeInstanceFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartInstance 重启数据库实例
//
// 重启数据库实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

// RestartInstanceInvoker 重启数据库实例
func (c *GaussDBforopenGaussClient) RestartInstanceInvoker(request *model.RestartInstanceRequest) *RestartInstanceInvoker {
	requestDef := GenReqDefForRestartInstance()
	return &RestartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreInstance 备份恢复到当前实例
//
// 备份恢复到当前实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

// RestoreInstanceInvoker 备份恢复到当前实例
func (c *GaussDBforopenGaussClient) RestoreInstanceInvoker(request *model.RestoreInstanceRequest) *RestoreInstanceInvoker {
	requestDef := GenReqDefForRestoreInstance()
	return &RestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumePluginExtensions 配置插件拓展能力
//
// 配置插件拓展能力
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ResumePluginExtensions(request *model.ResumePluginExtensionsRequest) (*model.ResumePluginExtensionsResponse, error) {
	requestDef := GenReqDefForResumePluginExtensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumePluginExtensionsResponse), nil
	}
}

// ResumePluginExtensionsInvoker 配置插件拓展能力
func (c *GaussDBforopenGaussClient) ResumePluginExtensionsInvoker(request *model.ResumePluginExtensionsRequest) *ResumePluginExtensionsInvoker {
	requestDef := GenReqDefForResumePluginExtensions()
	return &ResumePluginExtensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunInstanceAction CN横向扩容/DN分片扩容/磁盘扩容
//
// CN横向扩容/DN分片扩容/磁盘扩容
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) RunInstanceAction(request *model.RunInstanceActionRequest) (*model.RunInstanceActionResponse, error) {
	requestDef := GenReqDefForRunInstanceAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunInstanceActionResponse), nil
	}
}

// RunInstanceActionInvoker CN横向扩容/DN分片扩容/磁盘扩容
func (c *GaussDBforopenGaussClient) RunInstanceActionInvoker(request *model.RunInstanceActionRequest) *RunInstanceActionInvoker {
	requestDef := GenReqDefForRunInstanceAction()
	return &RunInstanceActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchAutoEnlargePolicy 查询磁盘自动扩容策略
//
// 查询磁盘自动扩容策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SearchAutoEnlargePolicy(request *model.SearchAutoEnlargePolicyRequest) (*model.SearchAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForSearchAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAutoEnlargePolicyResponse), nil
	}
}

// SearchAutoEnlargePolicyInvoker 查询磁盘自动扩容策略
func (c *GaussDBforopenGaussClient) SearchAutoEnlargePolicyInvoker(request *model.SearchAutoEnlargePolicyRequest) *SearchAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForSearchAutoEnlargePolicy()
	return &SearchAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBackupPolicy 设置自动备份策略。
//
// 设置自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

// SetBackupPolicyInvoker 设置自动备份策略。
func (c *GaussDBforopenGaussClient) SetBackupPolicyInvoker(request *model.SetBackupPolicyRequest) *SetBackupPolicyInvoker {
	requestDef := GenReqDefForSetBackupPolicy()
	return &SetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetDbUserPwd 重置数据库帐号密码
//
// 重置指定数据库帐号的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SetDbUserPwd(request *model.SetDbUserPwdRequest) (*model.SetDbUserPwdResponse, error) {
	requestDef := GenReqDefForSetDbUserPwd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDbUserPwdResponse), nil
	}
}

// SetDbUserPwdInvoker 重置数据库帐号密码
func (c *GaussDBforopenGaussClient) SetDbUserPwdInvoker(request *model.SetDbUserPwdRequest) *SetDbUserPwdInvoker {
	requestDef := GenReqDefForSetDbUserPwd()
	return &SetDbUserPwdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetKernelPluginLicense 配置插件license
//
// 配置插件license
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SetKernelPluginLicense(request *model.SetKernelPluginLicenseRequest) (*model.SetKernelPluginLicenseResponse, error) {
	requestDef := GenReqDefForSetKernelPluginLicense()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetKernelPluginLicenseResponse), nil
	}
}

// SetKernelPluginLicenseInvoker 配置插件license
func (c *GaussDBforopenGaussClient) SetKernelPluginLicenseInvoker(request *model.SetKernelPluginLicenseRequest) *SetKernelPluginLicenseInvoker {
	requestDef := GenReqDefForSetKernelPluginLicense()
	return &SetKernelPluginLicenseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetNewBackupPolicy 设置自动备份策略
//
// 设置自动备份策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SetNewBackupPolicy(request *model.SetNewBackupPolicyRequest) (*model.SetNewBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetNewBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetNewBackupPolicyResponse), nil
	}
}

// SetNewBackupPolicyInvoker 设置自动备份策略
func (c *GaussDBforopenGaussClient) SetNewBackupPolicyInvoker(request *model.SetNewBackupPolicyRequest) *SetNewBackupPolicyInvoker {
	requestDef := GenReqDefForSetNewBackupPolicy()
	return &SetNewBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRecyclePolicy 设置回收站策略
//
// 设置回收站策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SetRecyclePolicy(request *model.SetRecyclePolicyRequest) (*model.SetRecyclePolicyResponse, error) {
	requestDef := GenReqDefForSetRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecyclePolicyResponse), nil
	}
}

// SetRecyclePolicyInvoker 设置回收站策略
func (c *GaussDBforopenGaussClient) SetRecyclePolicyInvoker(request *model.SetRecyclePolicyRequest) *SetRecyclePolicyInvoker {
	requestDef := GenReqDefForSetRecyclePolicy()
	return &SetRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

// ShowBackupPolicyInvoker 查询自动备份策略
func (c *GaussDBforopenGaussClient) ShowBackupPolicyInvoker(request *model.ShowBackupPolicyRequest) *ShowBackupPolicyInvoker {
	requestDef := GenReqDefForShowBackupPolicy()
	return &ShowBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBalanceStatus 查询实例主备平衡状态
//
// 查询实例是否发生过主备切换而导致主机负载不均衡。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowBalanceStatus(request *model.ShowBalanceStatusRequest) (*model.ShowBalanceStatusResponse, error) {
	requestDef := GenReqDefForShowBalanceStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBalanceStatusResponse), nil
	}
}

// ShowBalanceStatusInvoker 查询实例主备平衡状态
func (c *GaussDBforopenGaussClient) ShowBalanceStatusInvoker(request *model.ShowBalanceStatusRequest) *ShowBalanceStatusInvoker {
	requestDef := GenReqDefForShowBalanceStatus()
	return &ShowBalanceStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchUpgradeCandidateVersions 查询批量实例可升级的版本和升级类型。
//
// 查询批量实例可升级的版本和升级类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowBatchUpgradeCandidateVersions(request *model.ShowBatchUpgradeCandidateVersionsRequest) (*model.ShowBatchUpgradeCandidateVersionsResponse, error) {
	requestDef := GenReqDefForShowBatchUpgradeCandidateVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchUpgradeCandidateVersionsResponse), nil
	}
}

// ShowBatchUpgradeCandidateVersionsInvoker 查询批量实例可升级的版本和升级类型。
func (c *GaussDBforopenGaussClient) ShowBatchUpgradeCandidateVersionsInvoker(request *model.ShowBatchUpgradeCandidateVersionsRequest) *ShowBatchUpgradeCandidateVersionsInvoker {
	requestDef := GenReqDefForShowBatchUpgradeCandidateVersions()
	return &ShowBatchUpgradeCandidateVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationDetail 查询参数模板详情
//
// 根据参数模板ID获取指定参数模板详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowConfigurationDetail(request *model.ShowConfigurationDetailRequest) (*model.ShowConfigurationDetailResponse, error) {
	requestDef := GenReqDefForShowConfigurationDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationDetailResponse), nil
	}
}

// ShowConfigurationDetailInvoker 查询参数模板详情
func (c *GaussDBforopenGaussClient) ShowConfigurationDetailInvoker(request *model.ShowConfigurationDetailRequest) *ShowConfigurationDetailInvoker {
	requestDef := GenReqDefForShowConfigurationDetail()
	return &ShowConfigurationDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentForm 查询解决方案模板配置
//
// 根据解决方案模板名称或实例ID查询副本数、分片数、节点数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowDeploymentForm(request *model.ShowDeploymentFormRequest) (*model.ShowDeploymentFormResponse, error) {
	requestDef := GenReqDefForShowDeploymentForm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentFormResponse), nil
	}
}

// ShowDeploymentFormInvoker 查询解决方案模板配置
func (c *GaussDBforopenGaussClient) ShowDeploymentFormInvoker(request *model.ShowDeploymentFormRequest) *ShowDeploymentFormInvoker {
	requestDef := GenReqDefForShowDeploymentForm()
	return &ShowDeploymentFormInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowErrorLogSwitchStatus 查询错误日志采集开关状态
//
// 查询数据库错误日志采集的开关状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowErrorLogSwitchStatus(request *model.ShowErrorLogSwitchStatusRequest) (*model.ShowErrorLogSwitchStatusResponse, error) {
	requestDef := GenReqDefForShowErrorLogSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowErrorLogSwitchStatusResponse), nil
	}
}

// ShowErrorLogSwitchStatusInvoker 查询错误日志采集开关状态
func (c *GaussDBforopenGaussClient) ShowErrorLogSwitchStatusInvoker(request *model.ShowErrorLogSwitchStatusRequest) *ShowErrorLogSwitchStatusInvoker {
	requestDef := GenReqDefForShowErrorLogSwitchStatus()
	return &ShowErrorLogSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceConfiguration 获取指定实例的参数模板
//
// 获取指定实例的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceConfiguration(request *model.ShowInstanceConfigurationRequest) (*model.ShowInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForShowInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceConfigurationResponse), nil
	}
}

// ShowInstanceConfigurationInvoker 获取指定实例的参数模板
func (c *GaussDBforopenGaussClient) ShowInstanceConfigurationInvoker(request *model.ShowInstanceConfigurationRequest) *ShowInstanceConfigurationInvoker {
	requestDef := GenReqDefForShowInstanceConfiguration()
	return &ShowInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceDisk 查询实例存储空间使用信息
//
// 查询指定实例的存储使用空间和最大空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceDisk(request *model.ShowInstanceDiskRequest) (*model.ShowInstanceDiskResponse, error) {
	requestDef := GenReqDefForShowInstanceDisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceDiskResponse), nil
	}
}

// ShowInstanceDiskInvoker 查询实例存储空间使用信息
func (c *GaussDBforopenGaussClient) ShowInstanceDiskInvoker(request *model.ShowInstanceDiskRequest) *ShowInstanceDiskInvoker {
	requestDef := GenReqDefForShowInstanceDisk()
	return &ShowInstanceDiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceParamGroup 获取指定实例的参数模板
//
// 获取指定实例的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceParamGroup(request *model.ShowInstanceParamGroupRequest) (*model.ShowInstanceParamGroupResponse, error) {
	requestDef := GenReqDefForShowInstanceParamGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceParamGroupResponse), nil
	}
}

// ShowInstanceParamGroupInvoker 获取指定实例的参数模板
func (c *GaussDBforopenGaussClient) ShowInstanceParamGroupInvoker(request *model.ShowInstanceParamGroupRequest) *ShowInstanceParamGroupInvoker {
	requestDef := GenReqDefForShowInstanceParamGroup()
	return &ShowInstanceParamGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceSnapshot 根据时间点或者备份文件查询原实例信息
//
// 根据时间点或者备份文件查询原实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceSnapshot(request *model.ShowInstanceSnapshotRequest) (*model.ShowInstanceSnapshotResponse, error) {
	requestDef := GenReqDefForShowInstanceSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceSnapshotResponse), nil
	}
}

// ShowInstanceSnapshotInvoker 根据时间点或者备份文件查询原实例信息
func (c *GaussDBforopenGaussClient) ShowInstanceSnapshotInvoker(request *model.ShowInstanceSnapshotRequest) *ShowInstanceSnapshotInvoker {
	requestDef := GenReqDefForShowInstanceSnapshot()
	return &ShowInstanceSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 获取指定ID的任务信息。
//
// 获取指定ID的任务信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 获取指定ID的任务信息。
func (c *GaussDBforopenGaussClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectQuotas 查询租户的实例配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowProjectQuotas(request *model.ShowProjectQuotasRequest) (*model.ShowProjectQuotasResponse, error) {
	requestDef := GenReqDefForShowProjectQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProjectQuotasResponse), nil
	}
}

// ShowProjectQuotasInvoker 查询租户的实例配额
func (c *GaussDBforopenGaussClient) ShowProjectQuotasInvoker(request *model.ShowProjectQuotasRequest) *ShowProjectQuotasInvoker {
	requestDef := GenReqDefForShowProjectQuotas()
	return &ShowProjectQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecyclePolicy 查看回收站策略
//
// 查看回收站的回收策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

// ShowRecyclePolicyInvoker 查看回收站策略
func (c *GaussDBforopenGaussClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSlowLogDownload 查询慢日志下载信息
//
// 查询慢日志下载信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSlowLogDownload(request *model.ShowSlowLogDownloadRequest) (*model.ShowSlowLogDownloadResponse, error) {
	requestDef := GenReqDefForShowSlowLogDownload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowLogDownloadResponse), nil
	}
}

// ShowSlowLogDownloadInvoker 查询慢日志下载信息
func (c *GaussDBforopenGaussClient) ShowSlowLogDownloadInvoker(request *model.ShowSlowLogDownloadRequest) *ShowSlowLogDownloadInvoker {
	requestDef := GenReqDefForShowSlowLogDownload()
	return &ShowSlowLogDownloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSslCertDownloadLink 查询实例SSL证书下载地址
//
// 查询实例SSL证书下载地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSslCertDownloadLink(request *model.ShowSslCertDownloadLinkRequest) (*model.ShowSslCertDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowSslCertDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSslCertDownloadLinkResponse), nil
	}
}

// ShowSslCertDownloadLinkInvoker 查询实例SSL证书下载地址
func (c *GaussDBforopenGaussClient) ShowSslCertDownloadLinkInvoker(request *model.ShowSslCertDownloadLinkRequest) *ShowSslCertDownloadLinkInvoker {
	requestDef := GenReqDefForShowSslCertDownloadLink()
	return &ShowSslCertDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpgradeCandidateVersions 查询实例可升级版本
//
// 查询实例可升级版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowUpgradeCandidateVersions(request *model.ShowUpgradeCandidateVersionsRequest) (*model.ShowUpgradeCandidateVersionsResponse, error) {
	requestDef := GenReqDefForShowUpgradeCandidateVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeCandidateVersionsResponse), nil
	}
}

// ShowUpgradeCandidateVersionsInvoker 查询实例可升级版本
func (c *GaussDBforopenGaussClient) ShowUpgradeCandidateVersionsInvoker(request *model.ShowUpgradeCandidateVersionsRequest) *ShowUpgradeCandidateVersionsInvoker {
	requestDef := GenReqDefForShowUpgradeCandidateVersions()
	return &ShowUpgradeCandidateVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInstance 启动数据库
//
// 启动数据库，同时支持节点级别的启动操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StartInstance(request *model.StartInstanceRequest) (*model.StartInstanceResponse, error) {
	requestDef := GenReqDefForStartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInstanceResponse), nil
	}
}

// StartInstanceInvoker 启动数据库
func (c *GaussDBforopenGaussClient) StartInstanceInvoker(request *model.StartInstanceRequest) *StartInstanceInvoker {
	requestDef := GenReqDefForStartInstance()
	return &StartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopBackup 停止备份
//
// 停止进行中的备份，包括全备和差备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopBackup(request *model.StopBackupRequest) (*model.StopBackupResponse, error) {
	requestDef := GenReqDefForStopBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopBackupResponse), nil
	}
}

// StopBackupInvoker 停止备份
func (c *GaussDBforopenGaussClient) StopBackupInvoker(request *model.StopBackupRequest) *StopBackupInvoker {
	requestDef := GenReqDefForStopBackup()
	return &StopBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchConfiguration 应用参数模板
//
// 指定实例变更参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchConfiguration(request *model.SwitchConfigurationRequest) (*model.SwitchConfigurationResponse, error) {
	requestDef := GenReqDefForSwitchConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchConfigurationResponse), nil
	}
}

// SwitchConfigurationInvoker 应用参数模板
func (c *GaussDBforopenGaussClient) SwitchConfigurationInvoker(request *model.SwitchConfigurationRequest) *SwitchConfigurationInvoker {
	requestDef := GenReqDefForSwitchConfiguration()
	return &SwitchConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchShard 分片节点主备切换。
//
// 支持用户对单个或多个DN分片做主备切换，同一分组内只能指定一个新的备节点进行升主操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchShard(request *model.SwitchShardRequest) (*model.SwitchShardResponse, error) {
	requestDef := GenReqDefForSwitchShard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchShardResponse), nil
	}
}

// SwitchShardInvoker 分片节点主备切换。
func (c *GaussDBforopenGaussClient) SwitchShardInvoker(request *model.SwitchShardRequest) *SwitchShardInvoker {
	requestDef := GenReqDefForSwitchShard()
	return &SwitchShardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfiguration 修改指定实例的参数
//
// 修改指定实例的参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateInstanceConfiguration(request *model.UpdateInstanceConfigurationRequest) (*model.UpdateInstanceConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationResponse), nil
	}
}

// UpdateInstanceConfigurationInvoker 修改指定实例的参数
func (c *GaussDBforopenGaussClient) UpdateInstanceConfigurationInvoker(request *model.UpdateInstanceConfigurationRequest) *UpdateInstanceConfigurationInvoker {
	requestDef := GenReqDefForUpdateInstanceConfiguration()
	return &UpdateInstanceConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceName 修改实例名称
//
// 修改实例名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

// UpdateInstanceNameInvoker 修改实例名称
func (c *GaussDBforopenGaussClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeInstanceVersion 实例内核版本升级
//
// GaussDB(for openGauss)实例版本升级。包括灰度升级，就地升级，热补丁升级等三种升级方式。
// 就地升级：
// 就地升级需要停止业务进行，会一次性升级集群中所有节点。就地升级需要暂停业务30分钟来升级。
// 灰度升级：
// 升级自动提交：所有节点进程一起升级，在升级过程中有大概10秒的业务中断，不阻塞其他业务操作。
// 升级待观察：升级待观察，将数据库升级过程细分为升级，提交两个阶段。升级阶段可以根据部署方式细分为按分片或者按az的滚动升级，提交阶段可以对升级完成后的实例进行业务测试，根据需要可以选择提交升级，或者升级回退。每个主dn或者cn组件升级就有一次10秒业务中断。升级过程均是先管理面，再数据面，由备到主的升级方式。 分布式实例：根据分片数滚动升级，每次滚动升级可以根据选择的分片数进行指定分片数量的节点进行升级。 主备版实例：根据AZ数进行滚动升级，每次滚动升级可以根据选择的AZ进行1个分区或者多个分区进行升级。
// 提交升级：提交升级。在升级完成，进入提交阶段时。业务测试正常后提交升级，完成本次升级流程。
// 升级回退：升级回退，在升级完成，进入提交阶段时。可以根据需要回退本次升级，回退到升级前的版本。
// 热补丁升级：
// 升级自动提交：热补丁自动升级并提交，中间无业务中断，仅修复产品bug。
// 升级回退：热补丁回退，无业务中断时间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpgradeInstanceVersion(request *model.UpgradeInstanceVersionRequest) (*model.UpgradeInstanceVersionResponse, error) {
	requestDef := GenReqDefForUpgradeInstanceVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeInstanceVersionResponse), nil
	}
}

// UpgradeInstanceVersionInvoker 实例内核版本升级
func (c *GaussDBforopenGaussClient) UpgradeInstanceVersionInvoker(request *model.UpgradeInstanceVersionRequest) *UpgradeInstanceVersionInvoker {
	requestDef := GenReqDefForUpgradeInstanceVersion()
	return &UpgradeInstanceVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeInstancesVersion 批量实例内核版本升级
//
// GaussDB批量实例版本升级。包括灰度升级，就地升级、热补丁升级三种升级方式。
// 就地升级：
// 就地升级需要停止业务进行，会一次性升级集群中所有节点。就地升级需要暂停业务30分钟来升级。
// 灰度升级：
// 升级自动提交：所有节点进程一起升级，在升级过程中有大概10秒的业务中断，不阻塞其他业务操作。
// 升级待观察：升级待观察，将数据库升级过程细分为升级，提交两个阶段。升级阶段可以根据部署方式细分为按分片或者按az的滚动升级，提交阶段可以对升级完成后的实例进行业务测试，根据需要可以选择提交升级，或者升级回退。每个主dn或者cn组件升级就有一次10秒业务中断。升级过程均是先管理面，再数据面，由备到主的升级方式。 分布式实例：根据分片数滚动升级，每次滚动升级可以根据选择的分片数进行指定分片数量的节点进行升级。 主备版实例：根据AZ数进行滚动升级，每次滚动升级可以根据选择的AZ进行1个分区或者多个分区进行升级。
// 热补丁升级：
// 升级自动提交：热补丁自动升级并提交，中间无业务中断，仅修复产品bug。
// 提交升级：提交升级。在升级完成，进入提交阶段时。业务测试正常后提交升级，完成本次升级流程。
// 升级回退：升级回退，在升级完成，进入提交阶段时。可以根据需要回退本次升级，回退到升级前的版本。
// 批量实例可升级版本大于当前所有实例的引擎版本，且选择的所有实例，其升级方式和操作方式要保持一致。
// 若批量实例升级方式是灰度升级，默认升级所有az和分片。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpgradeInstancesVersion(request *model.UpgradeInstancesVersionRequest) (*model.UpgradeInstancesVersionResponse, error) {
	requestDef := GenReqDefForUpgradeInstancesVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeInstancesVersionResponse), nil
	}
}

// UpgradeInstancesVersionInvoker 批量实例内核版本升级
func (c *GaussDBforopenGaussClient) UpgradeInstancesVersionInvoker(request *model.UpgradeInstancesVersionRequest) *UpgradeInstancesVersionInvoker {
	requestDef := GenReqDefForUpgradeInstancesVersion()
	return &UpgradeInstancesVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateParaGroupName 校验参数组名称是否存在
//
// 校验参数组名称是否存在。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ValidateParaGroupName(request *model.ValidateParaGroupNameRequest) (*model.ValidateParaGroupNameResponse, error) {
	requestDef := GenReqDefForValidateParaGroupName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateParaGroupNameResponse), nil
	}
}

// ValidateParaGroupNameInvoker 校验参数组名称是否存在
func (c *GaussDBforopenGaussClient) ValidateParaGroupNameInvoker(request *model.ValidateParaGroupNameRequest) *ValidateParaGroupNameInvoker {
	requestDef := GenReqDefForValidateParaGroupName()
	return &ValidateParaGroupNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateWeakPassword 弱密码校验
//
// 校验数据库root用户密码的安全性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ValidateWeakPassword(request *model.ValidateWeakPasswordRequest) (*model.ValidateWeakPasswordResponse, error) {
	requestDef := GenReqDefForValidateWeakPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateWeakPasswordResponse), nil
	}
}

// ValidateWeakPasswordInvoker 弱密码校验
func (c *GaussDBforopenGaussClient) ValidateWeakPasswordInvoker(request *model.ValidateWeakPasswordRequest) *ValidateWeakPasswordInvoker {
	requestDef := GenReqDefForValidateWeakPassword()
	return &ValidateWeakPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
