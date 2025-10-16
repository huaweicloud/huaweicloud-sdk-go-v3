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

// AddHbaConfs 新增客户端接入认证配置
//
// 新增客户端接入认证配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AddHbaConfs(request *model.AddHbaConfsRequest) (*model.AddHbaConfsResponse, error) {
	requestDef := GenReqDefForAddHbaConfs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddHbaConfsResponse), nil
	}
}

// AddHbaConfsInvoker 新增客户端接入认证配置
func (c *GaussDBforopenGaussClient) AddHbaConfsInvoker(request *model.AddHbaConfsRequest) *AddHbaConfsInvoker {
	requestDef := GenReqDefForAddHbaConfs()
	return &AddHbaConfsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// AllowDbRolePrivileges 授权数据库角色
//
// 在数据库中设置角色的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AllowDbRolePrivileges(request *model.AllowDbRolePrivilegesRequest) (*model.AllowDbRolePrivilegesResponse, error) {
	requestDef := GenReqDefForAllowDbRolePrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AllowDbRolePrivilegesResponse), nil
	}
}

// AllowDbRolePrivilegesInvoker 授权数据库角色
func (c *GaussDBforopenGaussClient) AllowDbRolePrivilegesInvoker(request *model.AllowDbRolePrivilegesRequest) *AllowDbRolePrivilegesInvoker {
	requestDef := GenReqDefForAllowDbRolePrivileges()
	return &AllowDbRolePrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// AuthorizeBackupDownload 授权备份文件下载
//
// 授权租户使用OBS Browser+方式下载备份文件，支持实例级、表级的全量备份及差量备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) AuthorizeBackupDownload(request *model.AuthorizeBackupDownloadRequest) (*model.AuthorizeBackupDownloadResponse, error) {
	requestDef := GenReqDefForAuthorizeBackupDownload()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeBackupDownloadResponse), nil
	}
}

// AuthorizeBackupDownloadInvoker 授权备份文件下载
func (c *GaussDBforopenGaussClient) AuthorizeBackupDownloadInvoker(request *model.AuthorizeBackupDownloadRequest) *AuthorizeBackupDownloadInvoker {
	requestDef := GenReqDefForAuthorizeBackupDownload()
	return &AuthorizeBackupDownloadInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSetBackupPolicy 批量设置自动备份策略
//
// 批量设置自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) BatchSetBackupPolicy(request *model.BatchSetBackupPolicyRequest) (*model.BatchSetBackupPolicyResponse, error) {
	requestDef := GenReqDefForBatchSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSetBackupPolicyResponse), nil
	}
}

// BatchSetBackupPolicyInvoker 批量设置自动备份策略
func (c *GaussDBforopenGaussClient) BatchSetBackupPolicyInvoker(request *model.BatchSetBackupPolicyRequest) *BatchSetBackupPolicyInvoker {
	requestDef := GenReqDefForBatchSetBackupPolicy()
	return &BatchSetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowUpgradeCandidateVersions 查询批量实例可升级的版本和升级类型
//
// 查询批量实例可升级的版本和升级类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) BatchShowUpgradeCandidateVersions(request *model.BatchShowUpgradeCandidateVersionsRequest) (*model.BatchShowUpgradeCandidateVersionsResponse, error) {
	requestDef := GenReqDefForBatchShowUpgradeCandidateVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowUpgradeCandidateVersionsResponse), nil
	}
}

// BatchShowUpgradeCandidateVersionsInvoker 查询批量实例可升级的版本和升级类型
func (c *GaussDBforopenGaussClient) BatchShowUpgradeCandidateVersionsInvoker(request *model.BatchShowUpgradeCandidateVersionsRequest) *BatchShowUpgradeCandidateVersionsInvoker {
	requestDef := GenReqDefForBatchShowUpgradeCandidateVersions()
	return &BatchShowUpgradeCandidateVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelScheduleTask 取消定时任务
//
// 取消定时任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CancelScheduleTask(request *model.CancelScheduleTaskRequest) (*model.CancelScheduleTaskResponse, error) {
	requestDef := GenReqDefForCancelScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelScheduleTaskResponse), nil
	}
}

// CancelScheduleTaskInvoker 取消定时任务
func (c *GaussDBforopenGaussClient) CancelScheduleTaskInvoker(request *model.CancelScheduleTaskRequest) *CancelScheduleTaskInvoker {
	requestDef := GenReqDefForCancelScheduleTask()
	return &CancelScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateCrossCloudConstructDisaster 搭建容灾关系
//
// 搭建容灾关系（从主实例端下发）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateCrossCloudConstructDisaster(request *model.CreateCrossCloudConstructDisasterRequest) (*model.CreateCrossCloudConstructDisasterResponse, error) {
	requestDef := GenReqDefForCreateCrossCloudConstructDisaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCrossCloudConstructDisasterResponse), nil
	}
}

// CreateCrossCloudConstructDisasterInvoker 搭建容灾关系
func (c *GaussDBforopenGaussClient) CreateCrossCloudConstructDisasterInvoker(request *model.CreateCrossCloudConstructDisasterRequest) *CreateCrossCloudConstructDisasterInvoker {
	requestDef := GenReqDefForCreateCrossCloudConstructDisaster()
	return &CreateCrossCloudConstructDisasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateDatabaseInstance 创建数据库实例
//
// 创建数据库实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDatabaseInstance(request *model.CreateDatabaseInstanceRequest) (*model.CreateDatabaseInstanceResponse, error) {
	requestDef := GenReqDefForCreateDatabaseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseInstanceResponse), nil
	}
}

// CreateDatabaseInstanceInvoker 创建数据库实例
func (c *GaussDBforopenGaussClient) CreateDatabaseInstanceInvoker(request *model.CreateDatabaseInstanceRequest) *CreateDatabaseInstanceInvoker {
	requestDef := GenReqDefForCreateDatabaseInstance()
	return &CreateDatabaseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateDbRole 创建数据库角色
//
// 创建数据库角色。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateDbRole(request *model.CreateDbRoleRequest) (*model.CreateDbRoleResponse, error) {
	requestDef := GenReqDefForCreateDbRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDbRoleResponse), nil
	}
}

// CreateDbRoleInvoker 创建数据库角色
func (c *GaussDBforopenGaussClient) CreateDbRoleInvoker(request *model.CreateDbRoleRequest) *CreateDbRoleInvoker {
	requestDef := GenReqDefForCreateDbRole()
	return &CreateDbRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateReadonlyNodes 创建只读节点
//
// 创建只读节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateReadonlyNodes(request *model.CreateReadonlyNodesRequest) (*model.CreateReadonlyNodesResponse, error) {
	requestDef := GenReqDefForCreateReadonlyNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReadonlyNodesResponse), nil
	}
}

// CreateReadonlyNodesInvoker 创建只读节点
func (c *GaussDBforopenGaussClient) CreateReadonlyNodesInvoker(request *model.CreateReadonlyNodesRequest) *CreateReadonlyNodesInvoker {
	requestDef := GenReqDefForCreateReadonlyNodes()
	return &CreateReadonlyNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateScheduleTask 批量实例内核版本定时升级
//
// 批量实例内核版本定时升级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateScheduleTask(request *model.CreateScheduleTaskRequest) (*model.CreateScheduleTaskResponse, error) {
	requestDef := GenReqDefForCreateScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduleTaskResponse), nil
	}
}

// CreateScheduleTaskInvoker 批量实例内核版本定时升级
func (c *GaussDBforopenGaussClient) CreateScheduleTaskInvoker(request *model.CreateScheduleTaskRequest) *CreateScheduleTaskInvoker {
	requestDef := GenReqDefForCreateScheduleTask()
	return &CreateScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteDatabaseSchema 删除数据库SCHEMA
//
// 删除数据库schema。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteDatabaseSchema(request *model.DeleteDatabaseSchemaRequest) (*model.DeleteDatabaseSchemaResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseSchema()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseSchemaResponse), nil
	}
}

// DeleteDatabaseSchemaInvoker 删除数据库SCHEMA
func (c *GaussDBforopenGaussClient) DeleteDatabaseSchemaInvoker(request *model.DeleteDatabaseSchemaRequest) *DeleteDatabaseSchemaInvoker {
	requestDef := GenReqDefForDeleteDatabaseSchema()
	return &DeleteDatabaseSchemaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDisasterRecord 删除容灾记录
//
// 删除容灾记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteDisasterRecord(request *model.DeleteDisasterRecordRequest) (*model.DeleteDisasterRecordResponse, error) {
	requestDef := GenReqDefForDeleteDisasterRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDisasterRecordResponse), nil
	}
}

// DeleteDisasterRecordInvoker 删除容灾记录
func (c *GaussDBforopenGaussClient) DeleteDisasterRecordInvoker(request *model.DeleteDisasterRecordRequest) *DeleteDisasterRecordInvoker {
	requestDef := GenReqDefForDeleteDisasterRecord()
	return &DeleteDisasterRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHbaConfs 删除客户端接入认证配置
//
// 删除客户端接入认证配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteHbaConfs(request *model.DeleteHbaConfsRequest) (*model.DeleteHbaConfsResponse, error) {
	requestDef := GenReqDefForDeleteHbaConfs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHbaConfsResponse), nil
	}
}

// DeleteHbaConfsInvoker 删除客户端接入认证配置
func (c *GaussDBforopenGaussClient) DeleteHbaConfsInvoker(request *model.DeleteHbaConfsRequest) *DeleteHbaConfsInvoker {
	requestDef := GenReqDefForDeleteHbaConfs()
	return &DeleteHbaConfsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteReadonlyNodes 删除只读节点
//
// 删除只读节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteReadonlyNodes(request *model.DeleteReadonlyNodesRequest) (*model.DeleteReadonlyNodesResponse, error) {
	requestDef := GenReqDefForDeleteReadonlyNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReadonlyNodesResponse), nil
	}
}

// DeleteReadonlyNodesInvoker 删除只读节点
func (c *GaussDBforopenGaussClient) DeleteReadonlyNodesInvoker(request *model.DeleteReadonlyNodesRequest) *DeleteReadonlyNodesInvoker {
	requestDef := GenReqDefForDeleteReadonlyNodes()
	return &DeleteReadonlyNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduleTask 删除定时任务信息
//
// 删除定时任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteScheduleTask(request *model.DeleteScheduleTaskRequest) (*model.DeleteScheduleTaskResponse, error) {
	requestDef := GenReqDefForDeleteScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduleTaskResponse), nil
	}
}

// DeleteScheduleTaskInvoker 删除定时任务信息
func (c *GaussDBforopenGaussClient) DeleteScheduleTaskInvoker(request *model.DeleteScheduleTaskRequest) *DeleteScheduleTaskInvoker {
	requestDef := GenReqDefForDeleteScheduleTask()
	return &DeleteScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSharding 分片缩容
//
// 数据库分片缩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteSharding(request *model.DeleteShardingRequest) (*model.DeleteShardingResponse, error) {
	requestDef := GenReqDefForDeleteSharding()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteShardingResponse), nil
	}
}

// DeleteShardingInvoker 分片缩容
func (c *GaussDBforopenGaussClient) DeleteShardingInvoker(request *model.DeleteShardingRequest) *DeleteShardingInvoker {
	requestDef := GenReqDefForDeleteSharding()
	return &DeleteShardingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ExecuteCrossCloudDisasterDataCacheEnd 主实例结束容灾日志保持
//
// 结束stream流式容灾的日志保持功能，目前只有stream流容灾支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterDataCacheEnd(request *model.ExecuteCrossCloudDisasterDataCacheEndRequest) (*model.ExecuteCrossCloudDisasterDataCacheEndResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterDataCacheEnd()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterDataCacheEndResponse), nil
	}
}

// ExecuteCrossCloudDisasterDataCacheEndInvoker 主实例结束容灾日志保持
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterDataCacheEndInvoker(request *model.ExecuteCrossCloudDisasterDataCacheEndRequest) *ExecuteCrossCloudDisasterDataCacheEndInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterDataCacheEnd()
	return &ExecuteCrossCloudDisasterDataCacheEndInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterDataCacheStart 开始日志保持
//
// 主实例开始容灾日志保持，目前只有stream流容灾支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterDataCacheStart(request *model.ExecuteCrossCloudDisasterDataCacheStartRequest) (*model.ExecuteCrossCloudDisasterDataCacheStartResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterDataCacheStart()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterDataCacheStartResponse), nil
	}
}

// ExecuteCrossCloudDisasterDataCacheStartInvoker 开始日志保持
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterDataCacheStartInvoker(request *model.ExecuteCrossCloudDisasterDataCacheStartRequest) *ExecuteCrossCloudDisasterDataCacheStartInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterDataCacheStart()
	return &ExecuteCrossCloudDisasterDataCacheStartInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterEndSimulation 结束容灾演练
//
// 灾备实例结束容灾演练，目前只有stream流容灾支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterEndSimulation(request *model.ExecuteCrossCloudDisasterEndSimulationRequest) (*model.ExecuteCrossCloudDisasterEndSimulationResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterEndSimulation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterEndSimulationResponse), nil
	}
}

// ExecuteCrossCloudDisasterEndSimulationInvoker 结束容灾演练
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterEndSimulationInvoker(request *model.ExecuteCrossCloudDisasterEndSimulationRequest) *ExecuteCrossCloudDisasterEndSimulationInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterEndSimulation()
	return &ExecuteCrossCloudDisasterEndSimulationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterRecoveryFailover 备实例容灾升主
//
// 容灾升主failover（灾备实例端下发）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterRecoveryFailover(request *model.ExecuteCrossCloudDisasterRecoveryFailoverRequest) (*model.ExecuteCrossCloudDisasterRecoveryFailoverResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterRecoveryFailover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterRecoveryFailoverResponse), nil
	}
}

// ExecuteCrossCloudDisasterRecoveryFailoverInvoker 备实例容灾升主
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterRecoveryFailoverInvoker(request *model.ExecuteCrossCloudDisasterRecoveryFailoverRequest) *ExecuteCrossCloudDisasterRecoveryFailoverInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterRecoveryFailover()
	return &ExecuteCrossCloudDisasterRecoveryFailoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterRestore 重建容灾关系
//
// 流容灾备升主选择支持容灾回切，实现容灾关系的重建任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterRestore(request *model.ExecuteCrossCloudDisasterRestoreRequest) (*model.ExecuteCrossCloudDisasterRestoreResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterRestore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterRestoreResponse), nil
	}
}

// ExecuteCrossCloudDisasterRestoreInvoker 重建容灾关系
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterRestoreInvoker(request *model.ExecuteCrossCloudDisasterRestoreRequest) *ExecuteCrossCloudDisasterRestoreInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterRestore()
	return &ExecuteCrossCloudDisasterRestoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterStartSimulation 开始容灾演练
//
// 开始容灾演练，目前只有stream流容灾支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterStartSimulation(request *model.ExecuteCrossCloudDisasterStartSimulationRequest) (*model.ExecuteCrossCloudDisasterStartSimulationResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterStartSimulation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterStartSimulationResponse), nil
	}
}

// ExecuteCrossCloudDisasterStartSimulationInvoker 开始容灾演练
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterStartSimulationInvoker(request *model.ExecuteCrossCloudDisasterStartSimulationRequest) *ExecuteCrossCloudDisasterStartSimulationInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterStartSimulation()
	return &ExecuteCrossCloudDisasterStartSimulationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudDisasterSwitchover 灾备实例主从切换
//
// 容灾switchover（可在主备任一一端下发）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterSwitchover(request *model.ExecuteCrossCloudDisasterSwitchoverRequest) (*model.ExecuteCrossCloudDisasterSwitchoverResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudDisasterSwitchover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudDisasterSwitchoverResponse), nil
	}
}

// ExecuteCrossCloudDisasterSwitchoverInvoker 灾备实例主从切换
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudDisasterSwitchoverInvoker(request *model.ExecuteCrossCloudDisasterSwitchoverRequest) *ExecuteCrossCloudDisasterSwitchoverInvoker {
	requestDef := GenReqDefForExecuteCrossCloudDisasterSwitchover()
	return &ExecuteCrossCloudDisasterSwitchoverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCrossCloudReleaseDisaster 解除容灾关系
//
// 解除容灾（从容灾主集群下发）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudReleaseDisaster(request *model.ExecuteCrossCloudReleaseDisasterRequest) (*model.ExecuteCrossCloudReleaseDisasterResponse, error) {
	requestDef := GenReqDefForExecuteCrossCloudReleaseDisaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCrossCloudReleaseDisasterResponse), nil
	}
}

// ExecuteCrossCloudReleaseDisasterInvoker 解除容灾关系
func (c *GaussDBforopenGaussClient) ExecuteCrossCloudReleaseDisasterInvoker(request *model.ExecuteCrossCloudReleaseDisasterRequest) *ExecuteCrossCloudReleaseDisasterInvoker {
	requestDef := GenReqDefForExecuteCrossCloudReleaseDisaster()
	return &ExecuteCrossCloudReleaseDisasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSlowSql 导出表信息
//
// 导出表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ExportSlowSql(request *model.ExportSlowSqlRequest) (*model.ExportSlowSqlResponse, error) {
	requestDef := GenReqDefForExportSlowSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSlowSqlResponse), nil
	}
}

// ExportSlowSqlInvoker 导出表信息
func (c *GaussDBforopenGaussClient) ExportSlowSqlInvoker(request *model.ExportSlowSqlRequest) *ExportSlowSqlInvoker {
	requestDef := GenReqDefForExportSlowSql()
	return &ExportSlowSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListBackupsDetails 查询备份列表
//
// 获取备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListBackupsDetails(request *model.ListBackupsDetailsRequest) (*model.ListBackupsDetailsResponse, error) {
	requestDef := GenReqDefForListBackupsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsDetailsResponse), nil
	}
}

// ListBackupsDetailsInvoker 查询备份列表
func (c *GaussDBforopenGaussClient) ListBackupsDetailsInvoker(request *model.ListBackupsDetailsRequest) *ListBackupsDetailsInvoker {
	requestDef := GenReqDefForListBackupsDetails()
	return &ListBackupsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 查询协调节点列表。
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

// ListDatabaseInstances 查询数据库实例列表/查询实例详情
//
// 查询数据库实例列表/查询实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseInstances(request *model.ListDatabaseInstancesRequest) (*model.ListDatabaseInstancesResponse, error) {
	requestDef := GenReqDefForListDatabaseInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseInstancesResponse), nil
	}
}

// ListDatabaseInstancesInvoker 查询数据库实例列表/查询实例详情
func (c *GaussDBforopenGaussClient) ListDatabaseInstancesInvoker(request *model.ListDatabaseInstancesRequest) *ListDatabaseInstancesInvoker {
	requestDef := GenReqDefForListDatabaseInstances()
	return &ListDatabaseInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseRoles 查询数据库角色列表
//
// 查询数据库角色列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseRoles(request *model.ListDatabaseRolesRequest) (*model.ListDatabaseRolesResponse, error) {
	requestDef := GenReqDefForListDatabaseRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseRolesResponse), nil
	}
}

// ListDatabaseRolesInvoker 查询数据库角色列表
func (c *GaussDBforopenGaussClient) ListDatabaseRolesInvoker(request *model.ListDatabaseRolesRequest) *ListDatabaseRolesInvoker {
	requestDef := GenReqDefForListDatabaseRoles()
	return &ListDatabaseRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseSchemaTables 查询数据库表列表
//
// 查询指定实例的数据库表列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseSchemaTables(request *model.ListDatabaseSchemaTablesRequest) (*model.ListDatabaseSchemaTablesResponse, error) {
	requestDef := GenReqDefForListDatabaseSchemaTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseSchemaTablesResponse), nil
	}
}

// ListDatabaseSchemaTablesInvoker 查询数据库表列表
func (c *GaussDBforopenGaussClient) ListDatabaseSchemaTablesInvoker(request *model.ListDatabaseSchemaTablesRequest) *ListDatabaseSchemaTablesInvoker {
	requestDef := GenReqDefForListDatabaseSchemaTables()
	return &ListDatabaseSchemaTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListDatabaseVersions 查询数据库引擎的版本
//
// 查询数据库引擎的版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseVersions(request *model.ListDatabaseVersionsRequest) (*model.ListDatabaseVersionsResponse, error) {
	requestDef := GenReqDefForListDatabaseVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseVersionsResponse), nil
	}
}

// ListDatabaseVersionsInvoker 查询数据库引擎的版本
func (c *GaussDBforopenGaussClient) ListDatabaseVersionsInvoker(request *model.ListDatabaseVersionsRequest) *ListDatabaseVersionsInvoker {
	requestDef := GenReqDefForListDatabaseVersions()
	return &ListDatabaseVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseVolume 查询数据库占用空间大小列表
//
// 查询数据库占用空间大小列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseVolume(request *model.ListDatabaseVolumeRequest) (*model.ListDatabaseVolumeResponse, error) {
	requestDef := GenReqDefForListDatabaseVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseVolumeResponse), nil
	}
}

// ListDatabaseVolumeInvoker 查询数据库占用空间大小列表
func (c *GaussDBforopenGaussClient) ListDatabaseVolumeInvoker(request *model.ListDatabaseVolumeRequest) *ListDatabaseVolumeInvoker {
	requestDef := GenReqDefForListDatabaseVolume()
	return &ListDatabaseVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseVolumeSummary 查询数据盘空间概况
//
// 查询数据盘空间概况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatabaseVolumeSummary(request *model.ListDatabaseVolumeSummaryRequest) (*model.ListDatabaseVolumeSummaryResponse, error) {
	requestDef := GenReqDefForListDatabaseVolumeSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseVolumeSummaryResponse), nil
	}
}

// ListDatabaseVolumeSummaryInvoker 查询数据盘空间概况
func (c *GaussDBforopenGaussClient) ListDatabaseVolumeSummaryInvoker(request *model.ListDatabaseVolumeSummaryRequest) *ListDatabaseVolumeSummaryInvoker {
	requestDef := GenReqDefForListDatabaseVolumeSummary()
	return &ListDatabaseVolumeSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListDatastoresDetails 查询引擎列表
//
// 查询引擎列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDatastoresDetails(request *model.ListDatastoresDetailsRequest) (*model.ListDatastoresDetailsResponse, error) {
	requestDef := GenReqDefForListDatastoresDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoresDetailsResponse), nil
	}
}

// ListDatastoresDetailsInvoker 查询引擎列表
func (c *GaussDBforopenGaussClient) ListDatastoresDetailsInvoker(request *model.ListDatastoresDetailsRequest) *ListDatastoresDetailsInvoker {
	requestDef := GenReqDefForListDatastoresDetails()
	return &ListDatastoresDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListDisasterRecoveryRecord 查询操作记录
//
// 查询容灾操作记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListDisasterRecoveryRecord(request *model.ListDisasterRecoveryRecordRequest) (*model.ListDisasterRecoveryRecordResponse, error) {
	requestDef := GenReqDefForListDisasterRecoveryRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDisasterRecoveryRecordResponse), nil
	}
}

// ListDisasterRecoveryRecordInvoker 查询操作记录
func (c *GaussDBforopenGaussClient) ListDisasterRecoveryRecordInvoker(request *model.ListDisasterRecoveryRecordRequest) *ListDisasterRecoveryRecordInvoker {
	requestDef := GenReqDefForListDisasterRecoveryRecord()
	return &ListDisasterRecoveryRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseProjects 查询企业项目列表
//
// 查询企业项目列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListEnterpriseProjects(request *model.ListEnterpriseProjectsRequest) (*model.ListEnterpriseProjectsResponse, error) {
	requestDef := GenReqDefForListEnterpriseProjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseProjectsResponse), nil
	}
}

// ListEnterpriseProjectsInvoker 查询企业项目列表
func (c *GaussDBforopenGaussClient) ListEnterpriseProjectsInvoker(request *model.ListEnterpriseProjectsRequest) *ListEnterpriseProjectsInvoker {
	requestDef := GenReqDefForListEnterpriseProjects()
	return &ListEnterpriseProjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListFeatures 查询实例特性列表
//
// 查询当前实例高级特性列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListFeatures(request *model.ListFeaturesRequest) (*model.ListFeaturesResponse, error) {
	requestDef := GenReqDefForListFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesResponse), nil
	}
}

// ListFeaturesInvoker 查询实例特性列表
func (c *GaussDBforopenGaussClient) ListFeaturesInvoker(request *model.ListFeaturesRequest) *ListFeaturesInvoker {
	requestDef := GenReqDefForListFeatures()
	return &ListFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListFlavorsDetails 查询数据库规格
//
// 查询数据库的规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListFlavorsDetails(request *model.ListFlavorsDetailsRequest) (*model.ListFlavorsDetailsResponse, error) {
	requestDef := GenReqDefForListFlavorsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsDetailsResponse), nil
	}
}

// ListFlavorsDetailsInvoker 查询数据库规格
func (c *GaussDBforopenGaussClient) ListFlavorsDetailsInvoker(request *model.ListFlavorsDetailsRequest) *ListFlavorsDetailsInvoker {
	requestDef := GenReqDefForListFlavorsDetails()
	return &ListFlavorsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListHbaInfoHistory 查询客户端接入认证配置修改历史
//
// 查询实例的客户端接入认证配置修改历史。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListHbaInfoHistory(request *model.ListHbaInfoHistoryRequest) (*model.ListHbaInfoHistoryResponse, error) {
	requestDef := GenReqDefForListHbaInfoHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHbaInfoHistoryResponse), nil
	}
}

// ListHbaInfoHistoryInvoker 查询客户端接入认证配置修改历史
func (c *GaussDBforopenGaussClient) ListHbaInfoHistoryInvoker(request *model.ListHbaInfoHistoryRequest) *ListHbaInfoHistoryInvoker {
	requestDef := GenReqDefForListHbaInfoHistory()
	return &ListHbaInfoHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHbaInfos 查询客户端接入认证配置
//
// 查询客户端接入认证配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListHbaInfos(request *model.ListHbaInfosRequest) (*model.ListHbaInfosResponse, error) {
	requestDef := GenReqDefForListHbaInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHbaInfosResponse), nil
	}
}

// ListHbaInfosInvoker 查询客户端接入认证配置
func (c *GaussDBforopenGaussClient) ListHbaInfosInvoker(request *model.ListHbaInfosRequest) *ListHbaInfosInvoker {
	requestDef := GenReqDefForListHbaInfos()
	return &ListHbaInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListInstanceEngineDetail 查看实例引擎版本分布
//
// 查看实例引擎版本分布
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListInstanceEngineDetail(request *model.ListInstanceEngineDetailRequest) (*model.ListInstanceEngineDetailResponse, error) {
	requestDef := GenReqDefForListInstanceEngineDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceEngineDetailResponse), nil
	}
}

// ListInstanceEngineDetailInvoker 查看实例引擎版本分布
func (c *GaussDBforopenGaussClient) ListInstanceEngineDetailInvoker(request *model.ListInstanceEngineDetailRequest) *ListInstanceEngineDetailInvoker {
	requestDef := GenReqDefForListInstanceEngineDetail()
	return &ListInstanceEngineDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListKeyViewExecuteNode 查询关键视图下发节点信息
//
// 查询该实例下可用于执行会话查杀功能的节点 (非日志类型的CN或DN节点) 和其所包含的CN、DN组件信息。该接口是会话查杀功能的前提，其返回值中的每一对节点和组件ID是后续查杀会话接口的入参。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListKeyViewExecuteNode(request *model.ListKeyViewExecuteNodeRequest) (*model.ListKeyViewExecuteNodeResponse, error) {
	requestDef := GenReqDefForListKeyViewExecuteNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKeyViewExecuteNodeResponse), nil
	}
}

// ListKeyViewExecuteNodeInvoker 查询关键视图下发节点信息
func (c *GaussDBforopenGaussClient) ListKeyViewExecuteNodeInvoker(request *model.ListKeyViewExecuteNodeRequest) *ListKeyViewExecuteNodeInvoker {
	requestDef := GenReqDefForListKeyViewExecuteNode()
	return &ListKeyViewExecuteNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListKmsTdeKey 查询KMS秘钥列表
//
// 查询当前可使用的kms秘钥列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListKmsTdeKey(request *model.ListKmsTdeKeyRequest) (*model.ListKmsTdeKeyResponse, error) {
	requestDef := GenReqDefForListKmsTdeKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListKmsTdeKeyResponse), nil
	}
}

// ListKmsTdeKeyInvoker 查询KMS秘钥列表
func (c *GaussDBforopenGaussClient) ListKmsTdeKeyInvoker(request *model.ListKmsTdeKeyRequest) *ListKmsTdeKeyInvoker {
	requestDef := GenReqDefForListKmsTdeKey()
	return &ListKmsTdeKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricDatas 查询实例列表指标
//
// 查询实例列表指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListMetricDatas(request *model.ListMetricDatasRequest) (*model.ListMetricDatasResponse, error) {
	requestDef := GenReqDefForListMetricDatas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricDatasResponse), nil
	}
}

// ListMetricDatasInvoker 查询实例列表指标
func (c *GaussDBforopenGaussClient) ListMetricDatasInvoker(request *model.ListMetricDatasRequest) *ListMetricDatasInvoker {
	requestDef := GenReqDefForListMetricDatas()
	return &ListMetricDatasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListParameterGroupTemplates 获取参数模板列表
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListParameterGroupTemplates(request *model.ListParameterGroupTemplatesRequest) (*model.ListParameterGroupTemplatesResponse, error) {
	requestDef := GenReqDefForListParameterGroupTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListParameterGroupTemplatesResponse), nil
	}
}

// ListParameterGroupTemplatesInvoker 获取参数模板列表
func (c *GaussDBforopenGaussClient) ListParameterGroupTemplatesInvoker(request *model.ListParameterGroupTemplatesRequest) *ListParameterGroupTemplatesInvoker {
	requestDef := GenReqDefForListParameterGroupTemplates()
	return &ListParameterGroupTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListPluginInfoList 查询插件列表
//
// 查询插件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListPluginInfoList(request *model.ListPluginInfoListRequest) (*model.ListPluginInfoListResponse, error) {
	requestDef := GenReqDefForListPluginInfoList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginInfoListResponse), nil
	}
}

// ListPluginInfoListInvoker 查询插件列表
func (c *GaussDBforopenGaussClient) ListPluginInfoListInvoker(request *model.ListPluginInfoListRequest) *ListPluginInfoListInvoker {
	requestDef := GenReqDefForListPluginInfoList()
	return &ListPluginInfoListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListReadonlyNodes 查询只读节点信息
//
// 查询实例的只读节点列表，以及实例可支持的最大只读节点数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListReadonlyNodes(request *model.ListReadonlyNodesRequest) (*model.ListReadonlyNodesResponse, error) {
	requestDef := GenReqDefForListReadonlyNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReadonlyNodesResponse), nil
	}
}

// ListReadonlyNodesInvoker 查询只读节点信息
func (c *GaussDBforopenGaussClient) ListReadonlyNodesInvoker(request *model.ListReadonlyNodesRequest) *ListReadonlyNodesInvoker {
	requestDef := GenReqDefForListReadonlyNodes()
	return &ListReadonlyNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRealTimeSessions 查询实时会话
//
// 查询数据库实例节点的实时会话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRealTimeSessions(request *model.ListRealTimeSessionsRequest) (*model.ListRealTimeSessionsResponse, error) {
	requestDef := GenReqDefForListRealTimeSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRealTimeSessionsResponse), nil
	}
}

// ListRealTimeSessionsInvoker 查询实时会话
func (c *GaussDBforopenGaussClient) ListRealTimeSessionsInvoker(request *model.ListRealTimeSessionsRequest) *ListRealTimeSessionsInvoker {
	requestDef := GenReqDefForListRealTimeSessions()
	return &ListRealTimeSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListRecycleInstancesDetails 查询回收站所有引擎实例列表。
//
// 查询回收站所有引擎实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRecycleInstancesDetails(request *model.ListRecycleInstancesDetailsRequest) (*model.ListRecycleInstancesDetailsResponse, error) {
	requestDef := GenReqDefForListRecycleInstancesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesDetailsResponse), nil
	}
}

// ListRecycleInstancesDetailsInvoker 查询回收站所有引擎实例列表。
func (c *GaussDBforopenGaussClient) ListRecycleInstancesDetailsInvoker(request *model.ListRecycleInstancesDetailsRequest) *ListRecycleInstancesDetailsInvoker {
	requestDef := GenReqDefForListRecycleInstancesDetails()
	return &ListRecycleInstancesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListRestorableInstancesDetails 查询可用于备份恢复的实例列表
//
// 查询可用于备份恢复的实例列表，实例信息要符合备份条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListRestorableInstancesDetails(request *model.ListRestorableInstancesDetailsRequest) (*model.ListRestorableInstancesDetailsResponse, error) {
	requestDef := GenReqDefForListRestorableInstancesDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestorableInstancesDetailsResponse), nil
	}
}

// ListRestorableInstancesDetailsInvoker 查询可用于备份恢复的实例列表
func (c *GaussDBforopenGaussClient) ListRestorableInstancesDetailsInvoker(request *model.ListRestorableInstancesDetailsRequest) *ListRestorableInstancesDetailsInvoker {
	requestDef := GenReqDefForListRestorableInstancesDetails()
	return &ListRestorableInstancesDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListScheduleTask 查看定时任务列表
//
// 查看定时任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListScheduleTask(request *model.ListScheduleTaskRequest) (*model.ListScheduleTaskResponse, error) {
	requestDef := GenReqDefForListScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduleTaskResponse), nil
	}
}

// ListScheduleTaskInvoker 查看定时任务列表
func (c *GaussDBforopenGaussClient) ListScheduleTaskInvoker(request *model.ListScheduleTaskRequest) *ListScheduleTaskInvoker {
	requestDef := GenReqDefForListScheduleTask()
	return &ListScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchemaAndTable 识别SQL文本中的表信息
//
// 识别SQL文本中的表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSchemaAndTable(request *model.ListSchemaAndTableRequest) (*model.ListSchemaAndTableResponse, error) {
	requestDef := GenReqDefForListSchemaAndTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemaAndTableResponse), nil
	}
}

// ListSchemaAndTableInvoker 识别SQL文本中的表信息
func (c *GaussDBforopenGaussClient) ListSchemaAndTableInvoker(request *model.ListSchemaAndTableRequest) *ListSchemaAndTableInvoker {
	requestDef := GenReqDefForListSchemaAndTable()
	return &ListSchemaAndTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionMemoryContext 查询会话内存上下文列表
//
// 查询数据库实例节点的会话内存上下文列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSessionMemoryContext(request *model.ListSessionMemoryContextRequest) (*model.ListSessionMemoryContextResponse, error) {
	requestDef := GenReqDefForListSessionMemoryContext()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionMemoryContextResponse), nil
	}
}

// ListSessionMemoryContextInvoker 查询会话内存上下文列表
func (c *GaussDBforopenGaussClient) ListSessionMemoryContextInvoker(request *model.ListSessionMemoryContextRequest) *ListSessionMemoryContextInvoker {
	requestDef := GenReqDefForListSessionMemoryContext()
	return &ListSessionMemoryContextInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionStatistics 查询实时会话统计
//
// 查询数据库实例节点的实时会话统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSessionStatistics(request *model.ListSessionStatisticsRequest) (*model.ListSessionStatisticsResponse, error) {
	requestDef := GenReqDefForListSessionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionStatisticsResponse), nil
	}
}

// ListSessionStatisticsInvoker 查询实时会话统计
func (c *GaussDBforopenGaussClient) ListSessionStatisticsInvoker(request *model.ListSessionStatisticsRequest) *ListSessionStatisticsInvoker {
	requestDef := GenReqDefForListSessionStatistics()
	return &ListSessionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionTopSqlStatistics 查询实时会话Top SQL统计
//
// 查询实时会话Top SQL统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSessionTopSqlStatistics(request *model.ListSessionTopSqlStatisticsRequest) (*model.ListSessionTopSqlStatisticsResponse, error) {
	requestDef := GenReqDefForListSessionTopSqlStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionTopSqlStatisticsResponse), nil
	}
}

// ListSessionTopSqlStatisticsInvoker 查询实时会话Top SQL统计
func (c *GaussDBforopenGaussClient) ListSessionTopSqlStatisticsInvoker(request *model.ListSessionTopSqlStatisticsRequest) *ListSessionTopSqlStatisticsInvoker {
	requestDef := GenReqDefForListSessionTopSqlStatistics()
	return &ListSessionTopSqlStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessionWaitEventStatistics 查询实时会话Top等待事件统计
//
// 查询实时会话Top等待事件统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSessionWaitEventStatistics(request *model.ListSessionWaitEventStatisticsRequest) (*model.ListSessionWaitEventStatisticsResponse, error) {
	requestDef := GenReqDefForListSessionWaitEventStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionWaitEventStatisticsResponse), nil
	}
}

// ListSessionWaitEventStatisticsInvoker 查询实时会话Top等待事件统计
func (c *GaussDBforopenGaussClient) ListSessionWaitEventStatisticsInvoker(request *model.ListSessionWaitEventStatisticsRequest) *ListSessionWaitEventStatisticsInvoker {
	requestDef := GenReqDefForListSessionWaitEventStatistics()
	return &ListSessionWaitEventStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTableDefinition 查询表定义信息
//
// 查询表定义信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTableDefinition(request *model.ListTableDefinitionRequest) (*model.ListTableDefinitionResponse, error) {
	requestDef := GenReqDefForListTableDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableDefinitionResponse), nil
	}
}

// ListTableDefinitionInvoker 查询表定义信息
func (c *GaussDBforopenGaussClient) ListTableDefinitionInvoker(request *model.ListTableDefinitionRequest) *ListTableDefinitionInvoker {
	requestDef := GenReqDefForListTableDefinition()
	return &ListTableDefinitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableDefinitions 查询指定表的表定义信息列表
//
// 查询指定表的表定义信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTableDefinitions(request *model.ListTableDefinitionsRequest) (*model.ListTableDefinitionsResponse, error) {
	requestDef := GenReqDefForListTableDefinitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableDefinitionsResponse), nil
	}
}

// ListTableDefinitionsInvoker 查询指定表的表定义信息列表
func (c *GaussDBforopenGaussClient) ListTableDefinitionsInvoker(request *model.ListTableDefinitionsRequest) *ListTableDefinitionsInvoker {
	requestDef := GenReqDefForListTableDefinitions()
	return &ListTableDefinitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListTransaction 查询事务列表
//
// 查询数据库实例节点的实时事务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTransaction(request *model.ListTransactionRequest) (*model.ListTransactionResponse, error) {
	requestDef := GenReqDefForListTransaction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTransactionResponse), nil
	}
}

// ListTransactionInvoker 查询事务列表
func (c *GaussDBforopenGaussClient) ListTransactionInvoker(request *model.ListTransactionRequest) *ListTransactionInvoker {
	requestDef := GenReqDefForListTransaction()
	return &ListTransactionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWaitEvent 查询等待事件列表
//
// 查询数据库实例节点的等待事件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListWaitEvent(request *model.ListWaitEventRequest) (*model.ListWaitEventResponse, error) {
	requestDef := GenReqDefForListWaitEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWaitEventResponse), nil
	}
}

// ListWaitEventInvoker 查询等待事件列表
func (c *GaussDBforopenGaussClient) ListWaitEventInvoker(request *model.ListWaitEventRequest) *ListWaitEventInvoker {
	requestDef := GenReqDefForListWaitEvent()
	return &ListWaitEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyAutoEnlargePolicy 修改磁盘自动扩容策略
//
// 修改磁盘自动扩容策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ModifyAutoEnlargePolicy(request *model.ModifyAutoEnlargePolicyRequest) (*model.ModifyAutoEnlargePolicyResponse, error) {
	requestDef := GenReqDefForModifyAutoEnlargePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyAutoEnlargePolicyResponse), nil
	}
}

// ModifyAutoEnlargePolicyInvoker 修改磁盘自动扩容策略
func (c *GaussDBforopenGaussClient) ModifyAutoEnlargePolicyInvoker(request *model.ModifyAutoEnlargePolicyRequest) *ModifyAutoEnlargePolicyInvoker {
	requestDef := GenReqDefForModifyAutoEnlargePolicy()
	return &ModifyAutoEnlargePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ModifyHbaConf 修改客户端接入认证配置
//
// 修改客户端接入认证配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ModifyHbaConf(request *model.ModifyHbaConfRequest) (*model.ModifyHbaConfResponse, error) {
	requestDef := GenReqDefForModifyHbaConf()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyHbaConfResponse), nil
	}
}

// ModifyHbaConfInvoker 修改客户端接入认证配置
func (c *GaussDBforopenGaussClient) ModifyHbaConfInvoker(request *model.ModifyHbaConfRequest) *ModifyHbaConfInvoker {
	requestDef := GenReqDefForModifyHbaConf()
	return &ModifyHbaConfInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyInstancePort 修改指定实例端口号
//
// 修改指定实例端口号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ModifyInstancePort(request *model.ModifyInstancePortRequest) (*model.ModifyInstancePortResponse, error) {
	requestDef := GenReqDefForModifyInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyInstancePortResponse), nil
	}
}

// ModifyInstancePortInvoker 修改指定实例端口号
func (c *GaussDBforopenGaussClient) ModifyInstancePortInvoker(request *model.ModifyInstancePortRequest) *ModifyInstancePortInvoker {
	requestDef := GenReqDefForModifyInstancePort()
	return &ModifyInstancePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ResetDrConfig 重置容灾配置
//
// 重置容灾网络等配置。1.将自动“创建委托”以授权DBS云服务访问VPC资源信息、查询IAAS接口。2.重置实例容灾网络等配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ResetDrConfig(request *model.ResetDrConfigRequest) (*model.ResetDrConfigResponse, error) {
	requestDef := GenReqDefForResetDrConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetDrConfigResponse), nil
	}
}

// ResetDrConfigInvoker 重置容灾配置
func (c *GaussDBforopenGaussClient) ResetDrConfigInvoker(request *model.ResetDrConfigRequest) *ResetDrConfigInvoker {
	requestDef := GenReqDefForResetDrConfig()
	return &ResetDrConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// GaussDB数据库实例规格变更
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

// RestoreHbaInfo 恢复客户端接入认证配置信息
//
// 恢复客户端接入认证配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) RestoreHbaInfo(request *model.RestoreHbaInfoRequest) (*model.RestoreHbaInfoResponse, error) {
	requestDef := GenReqDefForRestoreHbaInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreHbaInfoResponse), nil
	}
}

// RestoreHbaInfoInvoker 恢复客户端接入认证配置信息
func (c *GaussDBforopenGaussClient) RestoreHbaInfoInvoker(request *model.RestoreHbaInfoRequest) *RestoreHbaInfoInvoker {
	requestDef := GenReqDefForRestoreHbaInfo()
	return &RestoreHbaInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
// 查询磁盘自动扩容策略。
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

// ShowAlarmHistoryRecord 获取告警记录历史
//
// 获取告警记录历史。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowAlarmHistoryRecord(request *model.ShowAlarmHistoryRecordRequest) (*model.ShowAlarmHistoryRecordResponse, error) {
	requestDef := GenReqDefForShowAlarmHistoryRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmHistoryRecordResponse), nil
	}
}

// ShowAlarmHistoryRecordInvoker 获取告警记录历史
func (c *GaussDBforopenGaussClient) ShowAlarmHistoryRecordInvoker(request *model.ShowAlarmHistoryRecordRequest) *ShowAlarmHistoryRecordInvoker {
	requestDef := GenReqDefForShowAlarmHistoryRecord()
	return &ShowAlarmHistoryRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarmStatistics 实例告警信息汇总统计
//
// 实例告警信息汇总统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowAlarmStatistics(request *model.ShowAlarmStatisticsRequest) (*model.ShowAlarmStatisticsResponse, error) {
	requestDef := GenReqDefForShowAlarmStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmStatisticsResponse), nil
	}
}

// ShowAlarmStatisticsInvoker 实例告警信息汇总统计
func (c *GaussDBforopenGaussClient) ShowAlarmStatisticsInvoker(request *model.ShowAlarmStatisticsRequest) *ShowAlarmStatisticsInvoker {
	requestDef := GenReqDefForShowAlarmStatistics()
	return &ShowAlarmStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoKillTransactionConfig 获取自动中止事务配置
//
// 获取实例自动中止事务配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowAutoKillTransactionConfig(request *model.ShowAutoKillTransactionConfigRequest) (*model.ShowAutoKillTransactionConfigResponse, error) {
	requestDef := GenReqDefForShowAutoKillTransactionConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoKillTransactionConfigResponse), nil
	}
}

// ShowAutoKillTransactionConfigInvoker 获取自动中止事务配置
func (c *GaussDBforopenGaussClient) ShowAutoKillTransactionConfigInvoker(request *model.ShowAutoKillTransactionConfigRequest) *ShowAutoKillTransactionConfigInvoker {
	requestDef := GenReqDefForShowAutoKillTransactionConfig()
	return &ShowAutoKillTransactionConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowCrossCloudDisasterInstanceMonitor 查询实例容灾监控实时状态
//
// 查询实例容灾监控实时状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowCrossCloudDisasterInstanceMonitor(request *model.ShowCrossCloudDisasterInstanceMonitorRequest) (*model.ShowCrossCloudDisasterInstanceMonitorResponse, error) {
	requestDef := GenReqDefForShowCrossCloudDisasterInstanceMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCrossCloudDisasterInstanceMonitorResponse), nil
	}
}

// ShowCrossCloudDisasterInstanceMonitorInvoker 查询实例容灾监控实时状态
func (c *GaussDBforopenGaussClient) ShowCrossCloudDisasterInstanceMonitorInvoker(request *model.ShowCrossCloudDisasterInstanceMonitorRequest) *ShowCrossCloudDisasterInstanceMonitorInvoker {
	requestDef := GenReqDefForShowCrossCloudDisasterInstanceMonitor()
	return &ShowCrossCloudDisasterInstanceMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCrossCloudDisasterRelations 查询容灾关系列表
//
// 查询容灾关系列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowCrossCloudDisasterRelations(request *model.ShowCrossCloudDisasterRelationsRequest) (*model.ShowCrossCloudDisasterRelationsResponse, error) {
	requestDef := GenReqDefForShowCrossCloudDisasterRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCrossCloudDisasterRelationsResponse), nil
	}
}

// ShowCrossCloudDisasterRelationsInvoker 查询容灾关系列表
func (c *GaussDBforopenGaussClient) ShowCrossCloudDisasterRelationsInvoker(request *model.ShowCrossCloudDisasterRelationsRequest) *ShowCrossCloudDisasterRelationsInvoker {
	requestDef := GenReqDefForShowCrossCloudDisasterRelations()
	return &ShowCrossCloudDisasterRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowEpsRemainingQuota 获取企业项目剩余配额
//
// 获取企业项目剩余配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowEpsRemainingQuota(request *model.ShowEpsRemainingQuotaRequest) (*model.ShowEpsRemainingQuotaResponse, error) {
	requestDef := GenReqDefForShowEpsRemainingQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEpsRemainingQuotaResponse), nil
	}
}

// ShowEpsRemainingQuotaInvoker 获取企业项目剩余配额
func (c *GaussDBforopenGaussClient) ShowEpsRemainingQuotaInvoker(request *model.ShowEpsRemainingQuotaRequest) *ShowEpsRemainingQuotaInvoker {
	requestDef := GenReqDefForShowEpsRemainingQuota()
	return &ShowEpsRemainingQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowExpansionParameters 查询扩容优化参数设置模板
//
// 查询扩容优化参数设置模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowExpansionParameters(request *model.ShowExpansionParametersRequest) (*model.ShowExpansionParametersResponse, error) {
	requestDef := GenReqDefForShowExpansionParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowExpansionParametersResponse), nil
	}
}

// ShowExpansionParametersInvoker 查询扩容优化参数设置模板
func (c *GaussDBforopenGaussClient) ShowExpansionParametersInvoker(request *model.ShowExpansionParametersRequest) *ShowExpansionParametersInvoker {
	requestDef := GenReqDefForShowExpansionParameters()
	return &ShowExpansionParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowInstanceMetricData 查询实例指标数据
//
// 查询实例指标数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceMetricData(request *model.ShowInstanceMetricDataRequest) (*model.ShowInstanceMetricDataResponse, error) {
	requestDef := GenReqDefForShowInstanceMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMetricDataResponse), nil
	}
}

// ShowInstanceMetricDataInvoker 查询实例指标数据
func (c *GaussDBforopenGaussClient) ShowInstanceMetricDataInvoker(request *model.ShowInstanceMetricDataRequest) *ShowInstanceMetricDataInvoker {
	requestDef := GenReqDefForShowInstanceMetricData()
	return &ShowInstanceMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowInstanceParamGroupDetail 获取指定实例的参数模板
//
// 获取指定实例的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstanceParamGroupDetail(request *model.ShowInstanceParamGroupDetailRequest) (*model.ShowInstanceParamGroupDetailResponse, error) {
	requestDef := GenReqDefForShowInstanceParamGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceParamGroupDetailResponse), nil
	}
}

// ShowInstanceParamGroupDetailInvoker 获取指定实例的参数模板
func (c *GaussDBforopenGaussClient) ShowInstanceParamGroupDetailInvoker(request *model.ShowInstanceParamGroupDetailRequest) *ShowInstanceParamGroupDetailInvoker {
	requestDef := GenReqDefForShowInstanceParamGroupDetail()
	return &ShowInstanceParamGroupDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowInstancesStatistics 实例状态统计
//
// 实例状态统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowInstancesStatistics(request *model.ShowInstancesStatisticsRequest) (*model.ShowInstancesStatisticsResponse, error) {
	requestDef := GenReqDefForShowInstancesStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstancesStatisticsResponse), nil
	}
}

// ShowInstancesStatisticsInvoker 实例状态统计
func (c *GaussDBforopenGaussClient) ShowInstancesStatisticsInvoker(request *model.ShowInstancesStatisticsRequest) *ShowInstancesStatisticsInvoker {
	requestDef := GenReqDefForShowInstancesStatistics()
	return &ShowInstancesStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowKmsKeyDetail 查询KMS秘钥详情
//
// 查询KMS秘钥详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowKmsKeyDetail(request *model.ShowKmsKeyDetailRequest) (*model.ShowKmsKeyDetailResponse, error) {
	requestDef := GenReqDefForShowKmsKeyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKmsKeyDetailResponse), nil
	}
}

// ShowKmsKeyDetailInvoker 查询KMS秘钥详情
func (c *GaussDBforopenGaussClient) ShowKmsKeyDetailInvoker(request *model.ShowKmsKeyDetailRequest) *ShowKmsKeyDetailInvoker {
	requestDef := GenReqDefForShowKmsKeyDetail()
	return &ShowKmsKeyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricNames 查询指标分组的指标名称
//
// 查询指标分组的指标名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowMetricNames(request *model.ShowMetricNamesRequest) (*model.ShowMetricNamesResponse, error) {
	requestDef := GenReqDefForShowMetricNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricNamesResponse), nil
	}
}

// ShowMetricNamesInvoker 查询指标分组的指标名称
func (c *GaussDBforopenGaussClient) ShowMetricNamesInvoker(request *model.ShowMetricNamesRequest) *ShowMetricNamesInvoker {
	requestDef := GenReqDefForShowMetricNames()
	return &ShowMetricNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowParameterGroupDetail 查询参数模板详情
//
// 根据参数模板ID获取指定参数模板详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowParameterGroupDetail(request *model.ShowParameterGroupDetailRequest) (*model.ShowParameterGroupDetailResponse, error) {
	requestDef := GenReqDefForShowParameterGroupDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowParameterGroupDetailResponse), nil
	}
}

// ShowParameterGroupDetailInvoker 查询参数模板详情
func (c *GaussDBforopenGaussClient) ShowParameterGroupDetailInvoker(request *model.ShowParameterGroupDetailRequest) *ShowParameterGroupDetailInvoker {
	requestDef := GenReqDefForShowParameterGroupDetail()
	return &ShowParameterGroupDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProjectQuotas 查询租户的实例配额
//
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

// ShowRedistributionParameters 查询重分布参数设置模板
//
// 查询重分布参数设置模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowRedistributionParameters(request *model.ShowRedistributionParametersRequest) (*model.ShowRedistributionParametersResponse, error) {
	requestDef := GenReqDefForShowRedistributionParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRedistributionParametersResponse), nil
	}
}

// ShowRedistributionParametersInvoker 查询重分布参数设置模板
func (c *GaussDBforopenGaussClient) ShowRedistributionParametersInvoker(request *model.ShowRedistributionParametersRequest) *ShowRedistributionParametersInvoker {
	requestDef := GenReqDefForShowRedistributionParameters()
	return &ShowRedistributionParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSessionOverview 查询实时会话概览
//
// 查询数据库实例节点的实时会话概览信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSessionOverview(request *model.ShowSessionOverviewRequest) (*model.ShowSessionOverviewResponse, error) {
	requestDef := GenReqDefForShowSessionOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSessionOverviewResponse), nil
	}
}

// ShowSessionOverviewInvoker 查询实时会话概览
func (c *GaussDBforopenGaussClient) ShowSessionOverviewInvoker(request *model.ShowSessionOverviewRequest) *ShowSessionOverviewInvoker {
	requestDef := GenReqDefForShowSessionOverview()
	return &ShowSessionOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShardDiskMessages 获取分片的磁盘使用信息
//
// 获取分片的磁盘使用信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowShardDiskMessages(request *model.ShowShardDiskMessagesRequest) (*model.ShowShardDiskMessagesResponse, error) {
	requestDef := GenReqDefForShowShardDiskMessages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShardDiskMessagesResponse), nil
	}
}

// ShowShardDiskMessagesInvoker 获取分片的磁盘使用信息
func (c *GaussDBforopenGaussClient) ShowShardDiskMessagesInvoker(request *model.ShowShardDiskMessagesRequest) *ShowShardDiskMessagesInvoker {
	requestDef := GenReqDefForShowShardDiskMessages()
	return &ShowShardDiskMessagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowSlowSqlPlan 查询SQL执行计划信息
//
// 查询等待事件的SQL执行计划信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSlowSqlPlan(request *model.ShowSlowSqlPlanRequest) (*model.ShowSlowSqlPlanResponse, error) {
	requestDef := GenReqDefForShowSlowSqlPlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowSqlPlanResponse), nil
	}
}

// ShowSlowSqlPlanInvoker 查询SQL执行计划信息
func (c *GaussDBforopenGaussClient) ShowSlowSqlPlanInvoker(request *model.ShowSlowSqlPlanRequest) *ShowSlowSqlPlanInvoker {
	requestDef := GenReqDefForShowSlowSqlPlan()
	return &ShowSlowSqlPlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSlowSqlStack 查询SQL堆栈信息
//
// 查询等待事件的SQL堆栈信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSlowSqlStack(request *model.ShowSlowSqlStackRequest) (*model.ShowSlowSqlStackResponse, error) {
	requestDef := GenReqDefForShowSlowSqlStack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowSqlStackResponse), nil
	}
}

// ShowSlowSqlStackInvoker 查询SQL堆栈信息
func (c *GaussDBforopenGaussClient) ShowSlowSqlStackInvoker(request *model.ShowSlowSqlStackRequest) *ShowSlowSqlStackInvoker {
	requestDef := GenReqDefForShowSlowSqlStack()
	return &ShowSlowSqlStackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSourceInstanceDetail 根据时间点或者备份文件查询原实例信息
//
// 根据时间点或者备份文件查询原实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSourceInstanceDetail(request *model.ShowSourceInstanceDetailRequest) (*model.ShowSourceInstanceDetailResponse, error) {
	requestDef := GenReqDefForShowSourceInstanceDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSourceInstanceDetailResponse), nil
	}
}

// ShowSourceInstanceDetailInvoker 根据时间点或者备份文件查询原实例信息
func (c *GaussDBforopenGaussClient) ShowSourceInstanceDetailInvoker(request *model.ShowSourceInstanceDetailRequest) *ShowSourceInstanceDetailInvoker {
	requestDef := GenReqDefForShowSourceInstanceDetail()
	return &ShowSourceInstanceDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowUpgradeCandidateVersionsDetails 查询实例可升级版本
//
// 查询实例可升级版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowUpgradeCandidateVersionsDetails(request *model.ShowUpgradeCandidateVersionsDetailsRequest) (*model.ShowUpgradeCandidateVersionsDetailsResponse, error) {
	requestDef := GenReqDefForShowUpgradeCandidateVersionsDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeCandidateVersionsDetailsResponse), nil
	}
}

// ShowUpgradeCandidateVersionsDetailsInvoker 查询实例可升级版本
func (c *GaussDBforopenGaussClient) ShowUpgradeCandidateVersionsDetailsInvoker(request *model.ShowUpgradeCandidateVersionsDetailsRequest) *ShowUpgradeCandidateVersionsDetailsInvoker {
	requestDef := GenReqDefForShowUpgradeCandidateVersionsDetails()
	return &ShowUpgradeCandidateVersionsDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkCn 协调节点缩容
//
// 协调节点缩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShrinkCn(request *model.ShrinkCnRequest) (*model.ShrinkCnResponse, error) {
	requestDef := GenReqDefForShrinkCn()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkCnResponse), nil
	}
}

// ShrinkCnInvoker 协调节点缩容
func (c *GaussDBforopenGaussClient) ShrinkCnInvoker(request *model.ShrinkCnRequest) *ShrinkCnInvoker {
	requestDef := GenReqDefForShrinkCn()
	return &ShrinkCnInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInstance 启动实例/节点
//
// 启动实例/节点。
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

// StartInstanceInvoker 启动实例/节点
func (c *GaussDBforopenGaussClient) StartInstanceInvoker(request *model.StartInstanceRequest) *StartInstanceInvoker {
	requestDef := GenReqDefForStartInstance()
	return &StartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartMysqlCompatibility 开启M兼容端口服务
//
// 开启指定实例的M兼容端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StartMysqlCompatibility(request *model.StartMysqlCompatibilityRequest) (*model.StartMysqlCompatibilityResponse, error) {
	requestDef := GenReqDefForStartMysqlCompatibility()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartMysqlCompatibilityResponse), nil
	}
}

// StartMysqlCompatibilityInvoker 开启M兼容端口服务
func (c *GaussDBforopenGaussClient) StartMysqlCompatibilityInvoker(request *model.StartMysqlCompatibilityRequest) *StartMysqlCompatibilityInvoker {
	requestDef := GenReqDefForStartMysqlCompatibility()
	return &StartMysqlCompatibilityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// StopFreeSession 结束空闲会话
//
// 结束空闲会话。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopFreeSession(request *model.StopFreeSessionRequest) (*model.StopFreeSessionResponse, error) {
	requestDef := GenReqDefForStopFreeSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopFreeSessionResponse), nil
	}
}

// StopFreeSessionInvoker 结束空闲会话
func (c *GaussDBforopenGaussClient) StopFreeSessionInvoker(request *model.StopFreeSessionRequest) *StopFreeSessionInvoker {
	requestDef := GenReqDefForStopFreeSession()
	return &StopFreeSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInstance 停止实例/节点
//
// 停止数据库，同时支持节点级别的停止操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopInstance(request *model.StopInstanceRequest) (*model.StopInstanceResponse, error) {
	requestDef := GenReqDefForStopInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInstanceResponse), nil
	}
}

// StopInstanceInvoker 停止实例/节点
func (c *GaussDBforopenGaussClient) StopInstanceInvoker(request *model.StopInstanceRequest) *StopInstanceInvoker {
	requestDef := GenReqDefForStopInstance()
	return &StopInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopSession 结束会话
//
// 查杀指定会话列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopSession(request *model.StopSessionRequest) (*model.StopSessionResponse, error) {
	requestDef := GenReqDefForStopSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopSessionResponse), nil
	}
}

// StopSessionInvoker 结束会话
func (c *GaussDBforopenGaussClient) StopSessionInvoker(request *model.StopSessionRequest) *StopSessionInvoker {
	requestDef := GenReqDefForStopSession()
	return &StopSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopTransaction 手动结束事务
//
// 手动结束事务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopTransaction(request *model.StopTransactionRequest) (*model.StopTransactionResponse, error) {
	requestDef := GenReqDefForStopTransaction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTransactionResponse), nil
	}
}

// StopTransactionInvoker 手动结束事务
func (c *GaussDBforopenGaussClient) StopTransactionInvoker(request *model.StopTransactionRequest) *StopTransactionInvoker {
	requestDef := GenReqDefForStopTransaction()
	return &StopTransactionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// SwitchKmsTde 开启透明加密
//
// 开启透明加密。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchKmsTde(request *model.SwitchKmsTdeRequest) (*model.SwitchKmsTdeResponse, error) {
	requestDef := GenReqDefForSwitchKmsTde()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchKmsTdeResponse), nil
	}
}

// SwitchKmsTdeInvoker 开启透明加密
func (c *GaussDBforopenGaussClient) SwitchKmsTdeInvoker(request *model.SwitchKmsTdeRequest) *SwitchKmsTdeInvoker {
	requestDef := GenReqDefForSwitchKmsTde()
	return &SwitchKmsTdeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchReplica 分布式实例一主一备一日志形态切换到一主两备形态
//
// 当前只支持分布式独立部署一主一备一日志切换至一主两备形态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchReplica(request *model.SwitchReplicaRequest) (*model.SwitchReplicaResponse, error) {
	requestDef := GenReqDefForSwitchReplica()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchReplicaResponse), nil
	}
}

// SwitchReplicaInvoker 分布式实例一主一备一日志形态切换到一主两备形态
func (c *GaussDBforopenGaussClient) SwitchReplicaInvoker(request *model.SwitchReplicaRequest) *SwitchReplicaInvoker {
	requestDef := GenReqDefForSwitchReplica()
	return &SwitchReplicaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateExpansionParameters 修改扩容优化参数
//
// 修改扩容优化参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateExpansionParameters(request *model.UpdateExpansionParametersRequest) (*model.UpdateExpansionParametersResponse, error) {
	requestDef := GenReqDefForUpdateExpansionParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateExpansionParametersResponse), nil
	}
}

// UpdateExpansionParametersInvoker 修改扩容优化参数
func (c *GaussDBforopenGaussClient) UpdateExpansionParametersInvoker(request *model.UpdateExpansionParametersRequest) *UpdateExpansionParametersInvoker {
	requestDef := GenReqDefForUpdateExpansionParameters()
	return &UpdateExpansionParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFeatures 开启特性
//
// 打开高级特性开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateFeatures(request *model.UpdateFeaturesRequest) (*model.UpdateFeaturesResponse, error) {
	requestDef := GenReqDefForUpdateFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFeaturesResponse), nil
	}
}

// UpdateFeaturesInvoker 开启特性
func (c *GaussDBforopenGaussClient) UpdateFeaturesInvoker(request *model.UpdateFeaturesRequest) *UpdateFeaturesInvoker {
	requestDef := GenReqDefForUpdateFeatures()
	return &UpdateFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateInstanceVersions 实例内核版本升级
//
// GaussDB实例版本升级。包括灰度升级，就地升级，热补丁升级等三种升级方式。
//
// - 就地升级：
//
//   就地升级需要停止业务进行，会一次性升级集群中所有节点。就地升级需要暂停业务30分钟来升级。
//
// - 灰度升级：
//
//   每个主DN或者CN组件升级就有一次10秒业务中断。升级过程均是先管理面，再数据面，由备到主的升级方式。灰度升级支持升级自动提交和升级待观察两种操作方式。
//
//   - 升级自动提交：所有节点进程一起升级，在升级过程中有大概10秒的业务中断，不阻塞其他业务操作。
//
//   - 升级待观察：升级待观察，将数据库升级过程细分为升级，提交两个阶段。
//
//     - 升级阶段可以根据部署方式细分为按分片或者按AZ的滚动升级。
//
//       - 分布式实例：根据分片数滚动升级。
//       - 集中式实例：根据AZ数进行滚动升级。
//
//     - 提交阶段可以对升级完成后的实例进行业务测试，根据需要可以选择提交升级或者升级回退。
//
//       - 提交升级：提交升级。在升级完成，进入提交阶段时。业务测试正常后提交升级，完成本次升级流程。
//
//       - 升级回退：升级回退，在升级完成，进入提交阶段时。可以根据需要回退本次升级，回退到升级前的版本。
//
// - 热补丁升级
//
//   - 升级自动提交：热补丁自动升级并提交，中间无业务中断。
//
//   - 升级回退：热补丁回退，无业务中断时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateInstanceVersions(request *model.UpdateInstanceVersionsRequest) (*model.UpdateInstanceVersionsResponse, error) {
	requestDef := GenReqDefForUpdateInstanceVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceVersionsResponse), nil
	}
}

// UpdateInstanceVersionsInvoker 实例内核版本升级
func (c *GaussDBforopenGaussClient) UpdateInstanceVersionsInvoker(request *model.UpdateInstanceVersionsRequest) *UpdateInstanceVersionsInvoker {
	requestDef := GenReqDefForUpdateInstanceVersions()
	return &UpdateInstanceVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateMysqlCompatibility 更新/关闭M兼容端口服务
//
// 更新指定实例的M兼容端口服务配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateMysqlCompatibility(request *model.UpdateMysqlCompatibilityRequest) (*model.UpdateMysqlCompatibilityResponse, error) {
	requestDef := GenReqDefForUpdateMysqlCompatibility()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateMysqlCompatibilityResponse), nil
	}
}

// UpdateMysqlCompatibilityInvoker 更新/关闭M兼容端口服务
func (c *GaussDBforopenGaussClient) UpdateMysqlCompatibilityInvoker(request *model.UpdateMysqlCompatibilityRequest) *UpdateMysqlCompatibilityInvoker {
	requestDef := GenReqDefForUpdateMysqlCompatibility()
	return &UpdateMysqlCompatibilityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRedistributionControl 修改重分布参数
//
// 修改重分布参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateRedistributionControl(request *model.UpdateRedistributionControlRequest) (*model.UpdateRedistributionControlResponse, error) {
	requestDef := GenReqDefForUpdateRedistributionControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRedistributionControlResponse), nil
	}
}

// UpdateRedistributionControlInvoker 修改重分布参数
func (c *GaussDBforopenGaussClient) UpdateRedistributionControlInvoker(request *model.UpdateRedistributionControlRequest) *UpdateRedistributionControlInvoker {
	requestDef := GenReqDefForUpdateRedistributionControl()
	return &UpdateRedistributionControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CollectAsp 采集ASP报告
//
// 采集ASP报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CollectAsp(request *model.CollectAspRequest) (*model.CollectAspResponse, error) {
	requestDef := GenReqDefForCollectAsp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectAspResponse), nil
	}
}

// CollectAspInvoker 采集ASP报告
func (c *GaussDBforopenGaussClient) CollectAspInvoker(request *model.CollectAspRequest) *CollectAspInvoker {
	requestDef := GenReqDefForCollectAsp()
	return &CollectAspInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAspInfos 查询asp采集结果
//
// 查询ASP采集结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListAspInfos(request *model.ListAspInfosRequest) (*model.ListAspInfosResponse, error) {
	requestDef := GenReqDefForListAspInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAspInfosResponse), nil
	}
}

// ListAspInfosInvoker 查询asp采集结果
func (c *GaussDBforopenGaussClient) ListAspInfosInvoker(request *model.ListAspInfosRequest) *ListAspInfosInvoker {
	requestDef := GenReqDefForListAspInfos()
	return &ListAspInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAspStatus 查询ASP开关状态
//
// 查询ASP开关状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowAspStatus(request *model.ShowAspStatusRequest) (*model.ShowAspStatusResponse, error) {
	requestDef := GenReqDefForShowAspStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAspStatusResponse), nil
	}
}

// ShowAspStatusInvoker 查询ASP开关状态
func (c *GaussDBforopenGaussClient) ShowAspStatusInvoker(request *model.ShowAspStatusRequest) *ShowAspStatusInvoker {
	requestDef := GenReqDefForShowAspStatus()
	return &ShowAspStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAspStatus 打开或者关闭ASP开关
//
// 打开或者关闭ASP开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchAspStatus(request *model.SwitchAspStatusRequest) (*model.SwitchAspStatusResponse, error) {
	requestDef := GenReqDefForSwitchAspStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAspStatusResponse), nil
	}
}

// SwitchAspStatusInvoker 打开或者关闭ASP开关
func (c *GaussDBforopenGaussClient) SwitchAspStatusInvoker(request *model.SwitchAspStatusRequest) *SwitchAspStatusInvoker {
	requestDef := GenReqDefForSwitchAspStatus()
	return &SwitchAspStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindLtsConfig 关联LTS日志流
//
// 关联LTS日志流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) BindLtsConfig(request *model.BindLtsConfigRequest) (*model.BindLtsConfigResponse, error) {
	requestDef := GenReqDefForBindLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindLtsConfigResponse), nil
	}
}

// BindLtsConfigInvoker 关联LTS日志流
func (c *GaussDBforopenGaussClient) BindLtsConfigInvoker(request *model.BindLtsConfigRequest) *BindLtsConfigInvoker {
	requestDef := GenReqDefForBindLtsConfig()
	return &BindLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsConfigs 查看LTS日志配置信息
//
// 查看LTS日志配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListLtsConfigs(request *model.ListLtsConfigsRequest) (*model.ListLtsConfigsResponse, error) {
	requestDef := GenReqDefForListLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsConfigsResponse), nil
	}
}

// ListLtsConfigsInvoker 查看LTS日志配置信息
func (c *GaussDBforopenGaussClient) ListLtsConfigsInvoker(request *model.ListLtsConfigsRequest) *ListLtsConfigsInvoker {
	requestDef := GenReqDefForListLtsConfigs()
	return &ListLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindLtsConfig 解除关联LTS日志流
//
// 解除关联LTS日志流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UnbindLtsConfig(request *model.UnbindLtsConfigRequest) (*model.UnbindLtsConfigResponse, error) {
	requestDef := GenReqDefForUnbindLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindLtsConfigResponse), nil
	}
}

// UnbindLtsConfigInvoker 解除关联LTS日志流
func (c *GaussDBforopenGaussClient) UnbindLtsConfigInvoker(request *model.UnbindLtsConfigRequest) *UnbindLtsConfigInvoker {
	requestDef := GenReqDefForUnbindLtsConfig()
	return &UnbindLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLimitTask 创建限流任务
//
// 根据具体范围和类型，进行限流任务的创建
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateLimitTask(request *model.CreateLimitTaskRequest) (*model.CreateLimitTaskResponse, error) {
	requestDef := GenReqDefForCreateLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLimitTaskResponse), nil
	}
}

// CreateLimitTaskInvoker 创建限流任务
func (c *GaussDBforopenGaussClient) CreateLimitTaskInvoker(request *model.CreateLimitTaskRequest) *CreateLimitTaskInvoker {
	requestDef := GenReqDefForCreateLimitTask()
	return &CreateLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlLimitTask 创建SQL限流任务
//
// 根据具体范围和类型，进行限流任务的创建。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateSqlLimitTask(request *model.CreateSqlLimitTaskRequest) (*model.CreateSqlLimitTaskResponse, error) {
	requestDef := GenReqDefForCreateSqlLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlLimitTaskResponse), nil
	}
}

// CreateSqlLimitTaskInvoker 创建SQL限流任务
func (c *GaussDBforopenGaussClient) CreateSqlLimitTaskInvoker(request *model.CreateSqlLimitTaskRequest) *CreateSqlLimitTaskInvoker {
	requestDef := GenReqDefForCreateSqlLimitTask()
	return &CreateSqlLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLimitTask 删除限流任务
//
// 根据task_id进行限流任务的删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteLimitTask(request *model.DeleteLimitTaskRequest) (*model.DeleteLimitTaskResponse, error) {
	requestDef := GenReqDefForDeleteLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLimitTaskResponse), nil
	}
}

// DeleteLimitTaskInvoker 删除限流任务
func (c *GaussDBforopenGaussClient) DeleteLimitTaskInvoker(request *model.DeleteLimitTaskRequest) *DeleteLimitTaskInvoker {
	requestDef := GenReqDefForDeleteLimitTask()
	return &DeleteLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlLimitTask 删除SQL限流任务
//
// 根据task_id进行SQL限流任务的删除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) DeleteSqlLimitTask(request *model.DeleteSqlLimitTaskRequest) (*model.DeleteSqlLimitTaskResponse, error) {
	requestDef := GenReqDefForDeleteSqlLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlLimitTaskResponse), nil
	}
}

// DeleteSqlLimitTaskInvoker 删除SQL限流任务
func (c *GaussDBforopenGaussClient) DeleteSqlLimitTaskInvoker(request *model.DeleteSqlLimitTaskRequest) *DeleteSqlLimitTaskInvoker {
	requestDef := GenReqDefForDeleteSqlLimitTask()
	return &DeleteSqlLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnhanceFullSqlStatistics 查询实例全量SQL聚合统计列表
//
// 查询GaussDB数据库实例全量SQL聚合统计列表，主要统计各唯一SQL ID下的SQL记录总数量以及相关时间指标的平均值。支持增强型条件过滤， 如可以对SQL文本（query字段）进行多条件合并查询，对total_sql_time、sql_count字段进行大于、小于、区间范围等条件的过滤。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListEnhanceFullSqlStatistics(request *model.ListEnhanceFullSqlStatisticsRequest) (*model.ListEnhanceFullSqlStatisticsResponse, error) {
	requestDef := GenReqDefForListEnhanceFullSqlStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnhanceFullSqlStatisticsResponse), nil
	}
}

// ListEnhanceFullSqlStatisticsInvoker 查询实例全量SQL聚合统计列表
func (c *GaussDBforopenGaussClient) ListEnhanceFullSqlStatisticsInvoker(request *model.ListEnhanceFullSqlStatisticsRequest) *ListEnhanceFullSqlStatisticsInvoker {
	requestDef := GenReqDefForListEnhanceFullSqlStatistics()
	return &ListEnhanceFullSqlStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnhanceFullSqls 查询全量单条SQL记录列表
//
// 查询GaussDB数据库实例全量SQL日志记录列表，支持增强型条件过滤。 如可对SQL文本（query字段）进行多条件合并查询，可对db_time、cpu_time、data_io_time及execution_time字段进行大于、小于、区间范围等条件的过滤
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListEnhanceFullSqls(request *model.ListEnhanceFullSqlsRequest) (*model.ListEnhanceFullSqlsResponse, error) {
	requestDef := GenReqDefForListEnhanceFullSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnhanceFullSqlsResponse), nil
	}
}

// ListEnhanceFullSqlsInvoker 查询全量单条SQL记录列表
func (c *GaussDBforopenGaussClient) ListEnhanceFullSqlsInvoker(request *model.ListEnhanceFullSqlsRequest) *ListEnhanceFullSqlsInvoker {
	requestDef := GenReqDefForListEnhanceFullSqls()
	return &ListEnhanceFullSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFullSqlSwitches 查询GaussDB数据库实例全量SQL开关记录列表
//
// 查询GaussDB数据库实例全量SQL开关记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListFullSqlSwitches(request *model.ListFullSqlSwitchesRequest) (*model.ListFullSqlSwitchesResponse, error) {
	requestDef := GenReqDefForListFullSqlSwitches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFullSqlSwitchesResponse), nil
	}
}

// ListFullSqlSwitchesInvoker 查询GaussDB数据库实例全量SQL开关记录列表
func (c *GaussDBforopenGaussClient) ListFullSqlSwitchesInvoker(request *model.ListFullSqlSwitchesRequest) *ListFullSqlSwitchesInvoker {
	requestDef := GenReqDefForListFullSqlSwitches()
	return &ListFullSqlSwitchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLimitTask 根据指定条件查询限流任务列表
//
// 根据指定条件查询限流任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListLimitTask(request *model.ListLimitTaskRequest) (*model.ListLimitTaskResponse, error) {
	requestDef := GenReqDefForListLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLimitTaskResponse), nil
	}
}

// ListLimitTaskInvoker 根据指定条件查询限流任务列表
func (c *GaussDBforopenGaussClient) ListLimitTaskInvoker(request *model.ListLimitTaskRequest) *ListLimitTaskInvoker {
	requestDef := GenReqDefForListLimitTask()
	return &ListLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodeLimitSqlModel 查询节点的sql模板列表
//
// 查询节点的sql模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListNodeLimitSqlModel(request *model.ListNodeLimitSqlModelRequest) (*model.ListNodeLimitSqlModelResponse, error) {
	requestDef := GenReqDefForListNodeLimitSqlModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodeLimitSqlModelResponse), nil
	}
}

// ListNodeLimitSqlModelInvoker 查询节点的sql模板列表
func (c *GaussDBforopenGaussClient) ListNodeLimitSqlModelInvoker(request *model.ListNodeLimitSqlModelRequest) *ListNodeLimitSqlModelInvoker {
	requestDef := GenReqDefForListNodeLimitSqlModel()
	return &ListNodeLimitSqlModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowSqlDetails 查询慢SQL详情
//
// 根据归一化SQLID，查询该SQL模板详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSlowSqlDetails(request *model.ListSlowSqlDetailsRequest) (*model.ListSlowSqlDetailsResponse, error) {
	requestDef := GenReqDefForListSlowSqlDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowSqlDetailsResponse), nil
	}
}

// ListSlowSqlDetailsInvoker 查询慢SQL详情
func (c *GaussDBforopenGaussClient) ListSlowSqlDetailsInvoker(request *model.ListSlowSqlDetailsRequest) *ListSlowSqlDetailsInvoker {
	requestDef := GenReqDefForListSlowSqlDetails()
	return &ListSlowSqlDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowSqls 查询慢SQL列表
//
// 查询慢SQL列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSlowSqls(request *model.ListSlowSqlsRequest) (*model.ListSlowSqlsResponse, error) {
	requestDef := GenReqDefForListSlowSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowSqlsResponse), nil
	}
}

// ListSlowSqlsInvoker 查询慢SQL列表
func (c *GaussDBforopenGaussClient) ListSlowSqlsInvoker(request *model.ListSlowSqlsRequest) *ListSlowSqlsInvoker {
	requestDef := GenReqDefForListSlowSqls()
	return &ListSlowSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlExcuteNodes 查询慢SQL节点信息
//
// 根据实例ID查询慢SQL节点信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSqlExcuteNodes(request *model.ListSqlExcuteNodesRequest) (*model.ListSqlExcuteNodesResponse, error) {
	requestDef := GenReqDefForListSqlExcuteNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlExcuteNodesResponse), nil
	}
}

// ListSqlExcuteNodesInvoker 查询慢SQL节点信息
func (c *GaussDBforopenGaussClient) ListSqlExcuteNodesInvoker(request *model.ListSqlExcuteNodesRequest) *ListSqlExcuteNodesInvoker {
	requestDef := GenReqDefForListSqlExcuteNodes()
	return &ListSqlExcuteNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlLimitTask 根据指定条件查询SQL限流任务列表
//
// 根据指定条件查询SQL限流任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSqlLimitTask(request *model.ListSqlLimitTaskRequest) (*model.ListSqlLimitTaskResponse, error) {
	requestDef := GenReqDefForListSqlLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlLimitTaskResponse), nil
	}
}

// ListSqlLimitTaskInvoker 根据指定条件查询SQL限流任务列表
func (c *GaussDBforopenGaussClient) ListSqlLimitTaskInvoker(request *model.ListSqlLimitTaskRequest) *ListSqlLimitTaskInvoker {
	requestDef := GenReqDefForListSqlLimitTask()
	return &ListSqlLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlTrace 查询GaussDB数据库实例SQL链路
//
// 查询GaussDB数据库实例SQL链路，包含实例上对应组件的链路列表，如dn_6001、dn_6002、cn_5001、cn_5002。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListSqlTrace(request *model.ListSqlTraceRequest) (*model.ListSqlTraceResponse, error) {
	requestDef := GenReqDefForListSqlTrace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlTraceResponse), nil
	}
}

// ListSqlTraceInvoker 查询GaussDB数据库实例SQL链路
func (c *GaussDBforopenGaussClient) ListSqlTraceInvoker(request *model.ListSqlTraceRequest) *ListSqlTraceInvoker {
	requestDef := GenReqDefForListSqlTrace()
	return &ListSqlTraceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalSlowSqlDetail 查询全局慢SQL详情
//
// 根据唯一SQLID，查询指定组件的慢SQL详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowGlobalSlowSqlDetail(request *model.ShowGlobalSlowSqlDetailRequest) (*model.ShowGlobalSlowSqlDetailResponse, error) {
	requestDef := GenReqDefForShowGlobalSlowSqlDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalSlowSqlDetailResponse), nil
	}
}

// ShowGlobalSlowSqlDetailInvoker 查询全局慢SQL详情
func (c *GaussDBforopenGaussClient) ShowGlobalSlowSqlDetailInvoker(request *model.ShowGlobalSlowSqlDetailRequest) *ShowGlobalSlowSqlDetailInvoker {
	requestDef := GenReqDefForShowGlobalSlowSqlDetail()
	return &ShowGlobalSlowSqlDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLimitTask 查询限流任务详情
//
// 查询限流任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowLimitTask(request *model.ShowLimitTaskRequest) (*model.ShowLimitTaskResponse, error) {
	requestDef := GenReqDefForShowLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLimitTaskResponse), nil
	}
}

// ShowLimitTaskInvoker 查询限流任务详情
func (c *GaussDBforopenGaussClient) ShowLimitTaskInvoker(request *model.ShowLimitTaskRequest) *ShowLimitTaskInvoker {
	requestDef := GenReqDefForShowLimitTask()
	return &ShowLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlLimitTask 查询SQL限流任务详情
//
// 查询SQL限流任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSqlLimitTask(request *model.ShowSqlLimitTaskRequest) (*model.ShowSqlLimitTaskResponse, error) {
	requestDef := GenReqDefForShowSqlLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlLimitTaskResponse), nil
	}
}

// ShowSqlLimitTaskInvoker 查询SQL限流任务详情
func (c *GaussDBforopenGaussClient) ShowSqlLimitTaskInvoker(request *model.ShowSqlLimitTaskRequest) *ShowSqlLimitTaskInvoker {
	requestDef := GenReqDefForShowSqlLimitTask()
	return &ShowSqlLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartFullSql 开启GaussDB数据库实例全量SQL功能
//
// 开启全量SQL功能。使实例Agent侧开启内核侧全量SQL能力，持续化采集GaussDB数据库实例上的执行SQL语句，定时批量持久化存储于LTS云服务对应日志流中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StartFullSql(request *model.StartFullSqlRequest) (*model.StartFullSqlResponse, error) {
	requestDef := GenReqDefForStartFullSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartFullSqlResponse), nil
	}
}

// StartFullSqlInvoker 开启GaussDB数据库实例全量SQL功能
func (c *GaussDBforopenGaussClient) StartFullSqlInvoker(request *model.StartFullSqlRequest) *StartFullSqlInvoker {
	requestDef := GenReqDefForStartFullSql()
	return &StartFullSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopFullSql 关闭全量SQL功能
//
// 关闭全量SQL功能。关闭后，实例Agent侧将停止内核侧全量SQL语句记录的实时采集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) StopFullSql(request *model.StopFullSqlRequest) (*model.StopFullSqlResponse, error) {
	requestDef := GenReqDefForStopFullSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopFullSqlResponse), nil
	}
}

// StopFullSqlInvoker 关闭全量SQL功能
func (c *GaussDBforopenGaussClient) StopFullSqlInvoker(request *model.StopFullSqlRequest) *StopFullSqlInvoker {
	requestDef := GenReqDefForStopFullSql()
	return &StopFullSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncLimitData 同步内核侧sql限流数据至管控侧
//
// 同步内核侧sql限流数据至管控侧
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SyncLimitData(request *model.SyncLimitDataRequest) (*model.SyncLimitDataResponse, error) {
	requestDef := GenReqDefForSyncLimitData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncLimitDataResponse), nil
	}
}

// SyncLimitDataInvoker 同步内核侧sql限流数据至管控侧
func (c *GaussDBforopenGaussClient) SyncLimitDataInvoker(request *model.SyncLimitDataRequest) *SyncLimitDataInvoker {
	requestDef := GenReqDefForSyncLimitData()
	return &SyncLimitDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLimitTask 修改限流任务
//
// 根据新的条件进行限流任务的更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateLimitTask(request *model.UpdateLimitTaskRequest) (*model.UpdateLimitTaskResponse, error) {
	requestDef := GenReqDefForUpdateLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLimitTaskResponse), nil
	}
}

// UpdateLimitTaskInvoker 修改限流任务
func (c *GaussDBforopenGaussClient) UpdateLimitTaskInvoker(request *model.UpdateLimitTaskRequest) *UpdateLimitTaskInvoker {
	requestDef := GenReqDefForUpdateLimitTask()
	return &UpdateLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSqlLimitTask 修改SQL限流任务
//
// 根据新的条件进行SQL限流任务的更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) UpdateSqlLimitTask(request *model.UpdateSqlLimitTaskRequest) (*model.UpdateSqlLimitTaskResponse, error) {
	requestDef := GenReqDefForUpdateSqlLimitTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlLimitTaskResponse), nil
	}
}

// UpdateSqlLimitTaskInvoker 修改SQL限流任务
func (c *GaussDBforopenGaussClient) UpdateSqlLimitTaskInvoker(request *model.UpdateSqlLimitTaskRequest) *UpdateSqlLimitTaskInvoker {
	requestDef := GenReqDefForUpdateSqlLimitTask()
	return &UpdateSqlLimitTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlPatch 查询SQL PATCH
//
// 查询SQL PATCH信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowSqlPatch(request *model.ShowSqlPatchRequest) (*model.ShowSqlPatchResponse, error) {
	requestDef := GenReqDefForShowSqlPatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlPatchResponse), nil
	}
}

// ShowSqlPatchInvoker 查询SQL PATCH
func (c *GaussDBforopenGaussClient) ShowSqlPatchInvoker(request *model.ShowSqlPatchRequest) *ShowSqlPatchInvoker {
	requestDef := GenReqDefForShowSqlPatch()
	return &ShowSqlPatchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTopSqls 查询TopSQL列表
//
// 查询TopSQL列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ListTopSqls(request *model.ListTopSqlsRequest) (*model.ListTopSqlsResponse, error) {
	requestDef := GenReqDefForListTopSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTopSqlsResponse), nil
	}
}

// ListTopSqlsInvoker 查询TopSQL列表
func (c *GaussDBforopenGaussClient) ListTopSqlsInvoker(request *model.ListTopSqlsRequest) *ListTopSqlsInvoker {
	requestDef := GenReqDefForListTopSqls()
	return &ListTopSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectWdrSnapshot 采集WDR快照报告
//
// 采集WDR快照报告。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CollectWdrSnapshot(request *model.CollectWdrSnapshotRequest) (*model.CollectWdrSnapshotResponse, error) {
	requestDef := GenReqDefForCollectWdrSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectWdrSnapshotResponse), nil
	}
}

// CollectWdrSnapshotInvoker 采集WDR快照报告
func (c *GaussDBforopenGaussClient) CollectWdrSnapshotInvoker(request *model.CollectWdrSnapshotRequest) *CollectWdrSnapshotInvoker {
	requestDef := GenReqDefForCollectWdrSnapshot()
	return &CollectWdrSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWdrSnapshot 生成WDR快照
//
// 生成WDR快照。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) CreateWdrSnapshot(request *model.CreateWdrSnapshotRequest) (*model.CreateWdrSnapshotResponse, error) {
	requestDef := GenReqDefForCreateWdrSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWdrSnapshotResponse), nil
	}
}

// CreateWdrSnapshotInvoker 生成WDR快照
func (c *GaussDBforopenGaussClient) CreateWdrSnapshotInvoker(request *model.CreateWdrSnapshotRequest) *CreateWdrSnapshotInvoker {
	requestDef := GenReqDefForCreateWdrSnapshot()
	return &CreateWdrSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWdrSnapshotStatus 查询WDR快照开关状态
//
// 查询WDR快照开关状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) ShowWdrSnapshotStatus(request *model.ShowWdrSnapshotStatusRequest) (*model.ShowWdrSnapshotStatusResponse, error) {
	requestDef := GenReqDefForShowWdrSnapshotStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWdrSnapshotStatusResponse), nil
	}
}

// ShowWdrSnapshotStatusInvoker 查询WDR快照开关状态
func (c *GaussDBforopenGaussClient) ShowWdrSnapshotStatusInvoker(request *model.ShowWdrSnapshotStatusRequest) *ShowWdrSnapshotStatusInvoker {
	requestDef := GenReqDefForShowWdrSnapshotStatus()
	return &ShowWdrSnapshotStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchWdrSnapshotStatus 打开或关闭WDR快照开关
//
// 打开或关闭WDR快照开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBforopenGaussClient) SwitchWdrSnapshotStatus(request *model.SwitchWdrSnapshotStatusRequest) (*model.SwitchWdrSnapshotStatusResponse, error) {
	requestDef := GenReqDefForSwitchWdrSnapshotStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchWdrSnapshotStatusResponse), nil
	}
}

// SwitchWdrSnapshotStatusInvoker 打开或关闭WDR快照开关
func (c *GaussDBforopenGaussClient) SwitchWdrSnapshotStatusInvoker(request *model.SwitchWdrSnapshotStatusRequest) *SwitchWdrSnapshotStatusInvoker {
	requestDef := GenReqDefForSwitchWdrSnapshotStatus()
	return &SwitchWdrSnapshotStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
