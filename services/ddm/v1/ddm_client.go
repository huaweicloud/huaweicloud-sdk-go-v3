package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ddm/v1/model"
)

type DdmClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDdmClient(hcClient *httpclient.HcHttpClient) *DdmClient {
	return &DdmClient{HcClient: hcClient}
}

func DdmClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ChangeDatabaseVersion 变更内核版本
//
// 变更内核版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ChangeDatabaseVersion(request *model.ChangeDatabaseVersionRequest) (*model.ChangeDatabaseVersionResponse, error) {
	requestDef := GenReqDefForChangeDatabaseVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDatabaseVersionResponse), nil
	}
}

// ChangeDatabaseVersionInvoker 变更内核版本
func (c *DdmClient) ChangeDatabaseVersionInvoker(request *model.ChangeDatabaseVersionRequest) *ChangeDatabaseVersionInvoker {
	requestDef := GenReqDefForChangeDatabaseVersion()
	return &ChangeDatabaseVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDdmConfigurations 创建参数组
//
// 创建参数组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateDdmConfigurations(request *model.CreateDdmConfigurationsRequest) (*model.CreateDdmConfigurationsResponse, error) {
	requestDef := GenReqDefForCreateDdmConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDdmConfigurationsResponse), nil
	}
}

// CreateDdmConfigurationsInvoker 创建参数组
func (c *DdmClient) CreateDdmConfigurationsInvoker(request *model.CreateDdmConfigurationsRequest) *CreateDdmConfigurationsInvoker {
	requestDef := GenReqDefForCreateDdmConfigurations()
	return &CreateDdmConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfiguration 删除参数组
//
// 删除参数组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

// DeleteConfigurationInvoker 删除参数组
func (c *DdmClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseAvailableVersions 查询可变更内核版本
//
// 查询可变更内核版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDatabaseAvailableVersions(request *model.ListDatabaseAvailableVersionsRequest) (*model.ListDatabaseAvailableVersionsResponse, error) {
	requestDef := GenReqDefForListDatabaseAvailableVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseAvailableVersionsResponse), nil
	}
}

// ListDatabaseAvailableVersionsInvoker 查询可变更内核版本
func (c *DdmClient) ListDatabaseAvailableVersionsInvoker(request *model.ListDatabaseAvailableVersionsRequest) *ListDatabaseAvailableVersionsInvoker {
	requestDef := GenReqDefForListDatabaseAvailableVersions()
	return &ListDatabaseAvailableVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDdmConfigurations 获取参数模板列表
//
// 获取参数模板列表，包括所有DDM的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDdmConfigurations(request *model.ListDdmConfigurationsRequest) (*model.ListDdmConfigurationsResponse, error) {
	requestDef := GenReqDefForListDdmConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDdmConfigurationsResponse), nil
	}
}

// ListDdmConfigurationsInvoker 获取参数模板列表
func (c *DdmClient) ListDdmConfigurationsInvoker(request *model.ListDdmConfigurationsRequest) *ListDdmConfigurationsInvoker {
	requestDef := GenReqDefForListDdmConfigurations()
	return &ListDdmConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyConfiguration 修改实例参数
//
// 修改实例参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ModifyConfiguration(request *model.ModifyConfigurationRequest) (*model.ModifyConfigurationResponse, error) {
	requestDef := GenReqDefForModifyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyConfigurationResponse), nil
	}
}

// ModifyConfigurationInvoker 修改实例参数
func (c *DdmClient) ModifyConfigurationInvoker(request *model.ModifyConfigurationRequest) *ModifyConfigurationInvoker {
	requestDef := GenReqDefForModifyConfiguration()
	return &ModifyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollBackDatabaseVersion 回滚内核版本
//
// 回滚内核版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RollBackDatabaseVersion(request *model.RollBackDatabaseVersionRequest) (*model.RollBackDatabaseVersionResponse, error) {
	requestDef := GenReqDefForRollBackDatabaseVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollBackDatabaseVersionResponse), nil
	}
}

// RollBackDatabaseVersionInvoker 回滚内核版本
func (c *DdmClient) RollBackDatabaseVersionInvoker(request *model.RollBackDatabaseVersionRequest) *RollBackDatabaseVersionInvoker {
	requestDef := GenReqDefForRollBackDatabaseVersion()
	return &RollBackDatabaseVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfiguration 获取指定参数模板的参数
//
// 获取指定参数模板的参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowConfiguration(request *model.ShowConfigurationRequest) (*model.ShowConfigurationResponse, error) {
	requestDef := GenReqDefForShowConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationResponse), nil
	}
}

// ShowConfigurationInvoker 获取指定参数模板的参数
func (c *DdmClient) ShowConfigurationInvoker(request *model.ShowConfigurationRequest) *ShowConfigurationInvoker {
	requestDef := GenReqDefForShowConfiguration()
	return &ShowConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRiskInfo 内核版本风险提醒
//
// 内核版本风险提醒
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowRiskInfo(request *model.ShowRiskInfoRequest) (*model.ShowRiskInfoResponse, error) {
	requestDef := GenReqDefForShowRiskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRiskInfoResponse), nil
	}
}

// ShowRiskInfoInvoker 内核版本风险提醒
func (c *DdmClient) ShowRiskInfoInvoker(request *model.ShowRiskInfoRequest) *ShowRiskInfoInvoker {
	requestDef := GenReqDefForShowRiskInfo()
	return &ShowRiskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersion 查询API版本列表
//
// 查询API版本列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

// ListApiVersionInvoker 查询API版本列表
func (c *DdmClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteNodes 批量删除实例的节点
//
// 批量删除实例的节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) BatchDeleteNodes(request *model.BatchDeleteNodesRequest) (*model.BatchDeleteNodesResponse, error) {
	requestDef := GenReqDefForBatchDeleteNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteNodesResponse), nil
	}
}

// BatchDeleteNodesInvoker 批量删除实例的节点
func (c *DdmClient) BatchDeleteNodesInvoker(request *model.BatchDeleteNodesRequest) *BatchDeleteNodesInvoker {
	requestDef := GenReqDefForBatchDeleteNodes()
	return &BatchDeleteNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindEip 绑定弹性公网IP
//
// 绑定弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) BindEip(request *model.BindEipRequest) (*model.BindEipResponse, error) {
	requestDef := GenReqDefForBindEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindEipResponse), nil
	}
}

// BindEipInvoker 绑定弹性公网IP
func (c *DdmClient) BindEipInvoker(request *model.BindEipRequest) *BindEipInvoker {
	requestDef := GenReqDefForBindEip()
	return &BindEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelMigration 取消分片变更
//
// 取消分片变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CancelMigration(request *model.CancelMigrationRequest) (*model.CancelMigrationResponse, error) {
	requestDef := GenReqDefForCancelMigration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelMigrationResponse), nil
	}
}

// CancelMigrationInvoker 取消分片变更
func (c *DdmClient) CancelMigrationInvoker(request *model.CancelMigrationRequest) *CancelMigrationInvoker {
	requestDef := GenReqDefForCancelMigration()
	return &CancelMigrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeStrategy 修改切换路由策略
//
// 修改切换路由策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ChangeStrategy(request *model.ChangeStrategyRequest) (*model.ChangeStrategyResponse, error) {
	requestDef := GenReqDefForChangeStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeStrategyResponse), nil
	}
}

// ChangeStrategyInvoker 修改切换路由策略
func (c *DdmClient) ChangeStrategyInvoker(request *model.ChangeStrategyRequest) *ChangeStrategyInvoker {
	requestDef := GenReqDefForChangeStrategy()
	return &ChangeStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckMigrateLogicDb 分片变更预校验
//
// 分片变更预校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CheckMigrateLogicDb(request *model.CheckMigrateLogicDbRequest) (*model.CheckMigrateLogicDbResponse, error) {
	requestDef := GenReqDefForCheckMigrateLogicDb()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckMigrateLogicDbResponse), nil
	}
}

// CheckMigrateLogicDbInvoker 分片变更预校验
func (c *DdmClient) CheckMigrateLogicDbInvoker(request *model.CheckMigrateLogicDbRequest) *CheckMigrateLogicDbInvoker {
	requestDef := GenReqDefForCheckMigrateLogicDb()
	return &CheckMigrateLogicDbInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckPreliminaryResults 查询分片变更预校验异步结果
//
// 查询分片变更预校验异步结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CheckPreliminaryResults(request *model.CheckPreliminaryResultsRequest) (*model.CheckPreliminaryResultsResponse, error) {
	requestDef := GenReqDefForCheckPreliminaryResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPreliminaryResultsResponse), nil
	}
}

// CheckPreliminaryResultsInvoker 查询分片变更预校验异步结果
func (c *DdmClient) CheckPreliminaryResultsInvoker(request *model.CheckPreliminaryResultsRequest) *CheckPreliminaryResultsInvoker {
	requestDef := GenReqDefForCheckPreliminaryResults()
	return &CheckPreliminaryResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CleanMigration 清理分片变更
//
// 清理分片变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CleanMigration(request *model.CleanMigrationRequest) (*model.CleanMigrationResponse, error) {
	requestDef := GenReqDefForCleanMigration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CleanMigrationResponse), nil
	}
}

// CleanMigrationInvoker 清理分片变更
func (c *DdmClient) CleanMigrationInvoker(request *model.CleanMigrationRequest) *CleanMigrationInvoker {
	requestDef := GenReqDefForCleanMigration()
	return &CleanMigrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabase 创建DDM逻辑库
//
// 创建DDM逻辑库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateDatabase(request *model.CreateDatabaseRequest) (*model.CreateDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResponse), nil
	}
}

// CreateDatabaseInvoker 创建DDM逻辑库
func (c *DdmClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDdmDatabase 创建DDM逻辑库
//
// 创建DDM逻辑库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateDdmDatabase(request *model.CreateDdmDatabaseRequest) (*model.CreateDdmDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDdmDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDdmDatabaseResponse), nil
	}
}

// CreateDdmDatabaseInvoker 创建DDM逻辑库
func (c *DdmClient) CreateDdmDatabaseInvoker(request *model.CreateDdmDatabaseRequest) *CreateDdmDatabaseInvoker {
	requestDef := GenReqDefForCreateDdmDatabase()
	return &CreateDdmDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDdmInstance 购买创建DDM实例
//
// 购买创建DDM实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateDdmInstance(request *model.CreateDdmInstanceRequest) (*model.CreateDdmInstanceResponse, error) {
	requestDef := GenReqDefForCreateDdmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDdmInstanceResponse), nil
	}
}

// CreateDdmInstanceInvoker 购买创建DDM实例
func (c *DdmClient) CreateDdmInstanceInvoker(request *model.CreateDdmInstanceRequest) *CreateDdmInstanceInvoker {
	requestDef := GenReqDefForCreateDdmInstance()
	return &CreateDdmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroup 创建组
//
// 创建组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateGroup(request *model.CreateGroupRequest) (*model.CreateGroupResponse, error) {
	requestDef := GenReqDefForCreateGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupResponse), nil
	}
}

// CreateGroupInvoker 创建组
func (c *DdmClient) CreateGroupInvoker(request *model.CreateGroupRequest) *CreateGroupInvoker {
	requestDef := GenReqDefForCreateGroup()
	return &CreateGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 购买DDM实例
//
// 创建一个DDM实例。
//
// DDM运行于虚拟私有云。申请DDM实例前，需保证有可用的虚拟私有云，并且已配置好子网与安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 购买DDM实例
func (c *DdmClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUsers 创建DDM帐号
//
// DDM帐号用于连接和管理逻辑库。一个DDM实例最多能创建100个DDM帐号，一个DDM帐号可以关联多个逻辑库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CreateUsers(request *model.CreateUsersRequest) (*model.CreateUsersResponse, error) {
	requestDef := GenReqDefForCreateUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUsersResponse), nil
	}
}

// CreateUsersInvoker 创建DDM帐号
func (c *DdmClient) CreateUsersInvoker(request *model.CreateUsersRequest) *CreateUsersInvoker {
	requestDef := GenReqDefForCreateUsers()
	return &CreateUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBackup 删除备份
//
// 删除备份
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteBackup(request *model.DeleteBackupRequest) (*model.DeleteBackupResponse, error) {
	requestDef := GenReqDefForDeleteBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBackupResponse), nil
	}
}

// DeleteBackupInvoker 删除备份
func (c *DdmClient) DeleteBackupInvoker(request *model.DeleteBackupRequest) *DeleteBackupInvoker {
	requestDef := GenReqDefForDeleteBackup()
	return &DeleteBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabase 删除DDM逻辑库
//
// 删除指定的逻辑库，释放该逻辑库的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteDatabase(request *model.DeleteDatabaseRequest) (*model.DeleteDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResponse), nil
	}
}

// DeleteDatabaseInvoker 删除DDM逻辑库
func (c *DdmClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDdmDatabase 删除逻辑库
//
// 删除指定的逻辑库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteDdmDatabase(request *model.DeleteDdmDatabaseRequest) (*model.DeleteDdmDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDdmDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDdmDatabaseResponse), nil
	}
}

// DeleteDdmDatabaseInvoker 删除逻辑库
func (c *DdmClient) DeleteDdmDatabaseInvoker(request *model.DeleteDdmDatabaseRequest) *DeleteDdmDatabaseInvoker {
	requestDef := GenReqDefForDeleteDdmDatabase()
	return &DeleteDdmDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDdmInstance 删除DDM实例
//
// 删除指定的DDM实例，释放该实例的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteDdmInstance(request *model.DeleteDdmInstanceRequest) (*model.DeleteDdmInstanceResponse, error) {
	requestDef := GenReqDefForDeleteDdmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDdmInstanceResponse), nil
	}
}

// DeleteDdmInstanceInvoker 删除DDM实例
func (c *DdmClient) DeleteDdmInstanceInvoker(request *model.DeleteDdmInstanceRequest) *DeleteDdmInstanceInvoker {
	requestDef := GenReqDefForDeleteDdmInstance()
	return &DeleteDdmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroup 删除实例组
//
// 删除实例组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteGroup(request *model.DeleteGroupRequest) (*model.DeleteGroupResponse, error) {
	requestDef := GenReqDefForDeleteGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupResponse), nil
	}
}

// DeleteGroupInvoker 删除实例组
func (c *DdmClient) DeleteGroupInvoker(request *model.DeleteGroupRequest) *DeleteGroupInvoker {
	requestDef := GenReqDefForDeleteGroup()
	return &DeleteGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除DDM实例
//
// 删除指定的DDM实例，释放该实例的所有资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除DDM实例
func (c *DdmClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNodes 删除实例的节点
//
// 删除实例的节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteNodes(request *model.DeleteNodesRequest) (*model.DeleteNodesResponse, error) {
	requestDef := GenReqDefForDeleteNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodesResponse), nil
	}
}

// DeleteNodesInvoker 删除实例的节点
func (c *DdmClient) DeleteNodesInvoker(request *model.DeleteNodesRequest) *DeleteNodesInvoker {
	requestDef := GenReqDefForDeleteNodes()
	return &DeleteNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUser 删除DDM帐号
//
// 删除指定的DDM实例帐号，如果帐号关联了逻辑库，则对应的关联关系也会删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DeleteUser(request *model.DeleteUserRequest) (*model.DeleteUserResponse, error) {
	requestDef := GenReqDefForDeleteUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUserResponse), nil
	}
}

// DeleteUserInvoker 删除DDM帐号
func (c *DdmClient) DeleteUserInvoker(request *model.DeleteUserRequest) *DeleteUserInvoker {
	requestDef := GenReqDefForDeleteUser()
	return &DeleteUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadSchemaMetadata 导出逻辑库元数据
//
// 导出所有逻辑库物理分片在数据节点上的分布关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) DownloadSchemaMetadata(request *model.DownloadSchemaMetadataRequest) (*model.DownloadSchemaMetadataResponse, error) {
	requestDef := GenReqDefForDownloadSchemaMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadSchemaMetadataResponse), nil
	}
}

// DownloadSchemaMetadataInvoker 导出逻辑库元数据
func (c *DdmClient) DownloadSchemaMetadataInvoker(request *model.DownloadSchemaMetadataRequest) *DownloadSchemaMetadataInvoker {
	requestDef := GenReqDefForDownloadSchemaMetadata()
	return &DownloadSchemaMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteKillLogicalProcesses kill逻辑会话
//
// kill逻辑会话
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ExecuteKillLogicalProcesses(request *model.ExecuteKillLogicalProcessesRequest) (*model.ExecuteKillLogicalProcessesResponse, error) {
	requestDef := GenReqDefForExecuteKillLogicalProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteKillLogicalProcessesResponse), nil
	}
}

// ExecuteKillLogicalProcessesInvoker kill逻辑会话
func (c *DdmClient) ExecuteKillLogicalProcessesInvoker(request *model.ExecuteKillLogicalProcessesRequest) *ExecuteKillLogicalProcessesInvoker {
	requestDef := GenReqDefForExecuteKillLogicalProcesses()
	return &ExecuteKillLogicalProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteKillPhysicalProcesses kill物理会话
//
// kill物理会话
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ExecuteKillPhysicalProcesses(request *model.ExecuteKillPhysicalProcessesRequest) (*model.ExecuteKillPhysicalProcessesResponse, error) {
	requestDef := GenReqDefForExecuteKillPhysicalProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteKillPhysicalProcessesResponse), nil
	}
}

// ExecuteKillPhysicalProcessesInvoker kill物理会话
func (c *DdmClient) ExecuteKillPhysicalProcessesInvoker(request *model.ExecuteKillPhysicalProcessesRequest) *ExecuteKillPhysicalProcessesInvoker {
	requestDef := GenReqDefForExecuteKillPhysicalProcesses()
	return &ExecuteKillPhysicalProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandDdmInstanceNodes DDM实例节点扩容
//
// 对指定的DDM实例的节点个数进行扩容，支持按需实例与包周期实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ExpandDdmInstanceNodes(request *model.ExpandDdmInstanceNodesRequest) (*model.ExpandDdmInstanceNodesResponse, error) {
	requestDef := GenReqDefForExpandDdmInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandDdmInstanceNodesResponse), nil
	}
}

// ExpandDdmInstanceNodesInvoker DDM实例节点扩容
func (c *DdmClient) ExpandDdmInstanceNodesInvoker(request *model.ExpandDdmInstanceNodesRequest) *ExpandDdmInstanceNodesInvoker {
	requestDef := GenReqDefForExpandDdmInstanceNodes()
	return &ExpandDdmInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandInstanceNodes DDM实例节点扩容
//
// 对指定的DDM实例的节点个数进行扩容，支持按需实例与包周期实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ExpandInstanceNodes(request *model.ExpandInstanceNodesRequest) (*model.ExpandInstanceNodesResponse, error) {
	requestDef := GenReqDefForExpandInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandInstanceNodesResponse), nil
	}
}

// ExpandInstanceNodesInvoker DDM实例节点扩容
func (c *DdmClient) ExpandInstanceNodesInvoker(request *model.ExpandInstanceNodesRequest) *ExpandInstanceNodesInvoker {
	requestDef := GenReqDefForExpandInstanceNodes()
	return &ExpandInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableRds 查询创建逻辑库可选取的数据节点实例列表
//
// 查询创建逻辑库可选取的数据节点实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListAvailableRds(request *model.ListAvailableRdsRequest) (*model.ListAvailableRdsResponse, error) {
	requestDef := GenReqDefForListAvailableRds()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableRdsResponse), nil
	}
}

// ListAvailableRdsInvoker 查询创建逻辑库可选取的数据节点实例列表
func (c *DdmClient) ListAvailableRdsInvoker(request *model.ListAvailableRdsRequest) *ListAvailableRdsInvoker {
	requestDef := GenReqDefForListAvailableRds()
	return &ListAvailableRdsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableRdsForMigrate 查询分片变更可选取的数据节点实例列表
//
// 查询分片变更可选取的数据节点实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListAvailableRdsForMigrate(request *model.ListAvailableRdsForMigrateRequest) (*model.ListAvailableRdsForMigrateResponse, error) {
	requestDef := GenReqDefForListAvailableRdsForMigrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableRdsForMigrateResponse), nil
	}
}

// ListAvailableRdsForMigrateInvoker 查询分片变更可选取的数据节点实例列表
func (c *DdmClient) ListAvailableRdsForMigrateInvoker(request *model.ListAvailableRdsForMigrateRequest) *ListAvailableRdsForMigrateInvoker {
	requestDef := GenReqDefForListAvailableRdsForMigrate()
	return &ListAvailableRdsForMigrateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAvailableRdsList 查询创建逻辑库可选取的数据库实例列表
//
// 查询创建逻辑库可选取的数据库实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListAvailableRdsList(request *model.ListAvailableRdsListRequest) (*model.ListAvailableRdsListResponse, error) {
	requestDef := GenReqDefForListAvailableRdsList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAvailableRdsListResponse), nil
	}
}

// ListAvailableRdsListInvoker 查询创建逻辑库可选取的数据库实例列表
func (c *DdmClient) ListAvailableRdsListInvoker(request *model.ListAvailableRdsListRequest) *ListAvailableRdsListInvoker {
	requestDef := GenReqDefForListAvailableRdsList()
	return &ListAvailableRdsListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups 获取备份列表
//
// 获取备份列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// ListBackupsInvoker 获取备份列表
func (c *DdmClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabases 查询DDM逻辑库列表
//
// 查询DDM逻辑库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 查询DDM逻辑库列表
func (c *DdmClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDdmEngines 查询DDM引擎信息
//
// 查询DDM引擎信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDdmEngines(request *model.ListDdmEnginesRequest) (*model.ListDdmEnginesResponse, error) {
	requestDef := GenReqDefForListDdmEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDdmEnginesResponse), nil
	}
}

// ListDdmEnginesInvoker 查询DDM引擎信息
func (c *DdmClient) ListDdmEnginesInvoker(request *model.ListDdmEnginesRequest) *ListDdmEnginesInvoker {
	requestDef := GenReqDefForListDdmEngines()
	return &ListDdmEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDdmFlavors 查询DDM可用区规格信息
//
// 查询DDM可用区规格信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDdmFlavors(request *model.ListDdmFlavorsRequest) (*model.ListDdmFlavorsResponse, error) {
	requestDef := GenReqDefForListDdmFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDdmFlavorsResponse), nil
	}
}

// ListDdmFlavorsInvoker 查询DDM可用区规格信息
func (c *DdmClient) ListDdmFlavorsInvoker(request *model.ListDdmFlavorsRequest) *ListDdmFlavorsInvoker {
	requestDef := GenReqDefForListDdmFlavors()
	return &ListDdmFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDdms 查询实例列表
//
// 查询实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListDdms(request *model.ListDdmsRequest) (*model.ListDdmsResponse, error) {
	requestDef := GenReqDefForListDdms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDdmsResponse), nil
	}
}

// ListDdmsInvoker 查询实例列表
func (c *DdmClient) ListDdmsInvoker(request *model.ListDdmsRequest) *ListDdmsInvoker {
	requestDef := GenReqDefForListDdms()
	return &ListDdmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngines 查询DDM引擎信息
//
// 查询DDM引擎信息详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListEngines(request *model.ListEnginesRequest) (*model.ListEnginesResponse, error) {
	requestDef := GenReqDefForListEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnginesResponse), nil
	}
}

// ListEnginesInvoker 查询DDM引擎信息
func (c *DdmClient) ListEnginesInvoker(request *model.ListEnginesRequest) *ListEnginesInvoker {
	requestDef := GenReqDefForListEngines()
	return &ListEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询DDM可用区规格信息
//
// 查询DDM可用区规格信息详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询DDM可用区规格信息
func (c *DdmClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroup 获取实例组信息列表
//
// 获取实例组信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListGroup(request *model.ListGroupRequest) (*model.ListGroupResponse, error) {
	requestDef := GenReqDefForListGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupResponse), nil
	}
}

// ListGroupInvoker 获取实例组信息列表
func (c *DdmClient) ListGroupInvoker(request *model.ListGroupRequest) *ListGroupInvoker {
	requestDef := GenReqDefForListGroup()
	return &ListGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询DDM实例列表
//
// 查询DDM实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询DDM实例列表
func (c *DdmClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodes 查询DDM实例节点列表
//
// 查询DDM实例节点列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListNodes(request *model.ListNodesRequest) (*model.ListNodesResponse, error) {
	requestDef := GenReqDefForListNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodesResponse), nil
	}
}

// ListNodesInvoker 查询DDM实例节点列表
func (c *DdmClient) ListNodesInvoker(request *model.ListNodesRequest) *ListNodesInvoker {
	requestDef := GenReqDefForListNodes()
	return &ListNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListReadWriteRatio 读写比例监控
//
// 查询指定时间段内在DDM实例的读写次数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListReadWriteRatio(request *model.ListReadWriteRatioRequest) (*model.ListReadWriteRatioResponse, error) {
	requestDef := GenReqDefForListReadWriteRatio()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListReadWriteRatioResponse), nil
	}
}

// ListReadWriteRatioInvoker 读写比例监控
func (c *DdmClient) ListReadWriteRatioInvoker(request *model.ListReadWriteRatioRequest) *ListReadWriteRatioInvoker {
	requestDef := GenReqDefForListReadWriteRatio()
	return &ListReadWriteRatioInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowLog 慢日志监控
//
// 查询指定时间段内在DDM实例上执行过的慢sql相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListSlowLog(request *model.ListSlowLogRequest) (*model.ListSlowLogResponse, error) {
	requestDef := GenReqDefForListSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogResponse), nil
	}
}

// ListSlowLogInvoker 慢日志监控
func (c *DdmClient) ListSlowLogInvoker(request *model.ListSlowLogRequest) *ListSlowLogInvoker {
	requestDef := GenReqDefForListSlowLog()
	return &ListSlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowLogAnalysis 查询慢日志
//
// 查询慢日志V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListSlowLogAnalysis(request *model.ListSlowLogAnalysisRequest) (*model.ListSlowLogAnalysisResponse, error) {
	requestDef := GenReqDefForListSlowLogAnalysis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogAnalysisResponse), nil
	}
}

// ListSlowLogAnalysisInvoker 查询慢日志
func (c *DdmClient) ListSlowLogAnalysisInvoker(request *model.ListSlowLogAnalysisRequest) *ListSlowLogAnalysisInvoker {
	requestDef := GenReqDefForListSlowLogAnalysis()
	return &ListSlowLogAnalysisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowLogs 慢日志监控
//
// 查询指定时间段内在DDM实例上执行过的慢sql相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

// ListSlowLogsInvoker 慢日志监控
func (c *DdmClient) ListSlowLogsInvoker(request *model.ListSlowLogsRequest) *ListSlowLogsInvoker {
	requestDef := GenReqDefForListSlowLogs()
	return &ListSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询任务列表
//
// 查询任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表
func (c *DdmClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 查询DDM帐号列表
//
// 查询DDM帐号列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 查询DDM帐号列表
func (c *DdmClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateLogicDb 分片变更
//
// 分片变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) MigrateLogicDb(request *model.MigrateLogicDbRequest) (*model.MigrateLogicDbResponse, error) {
	requestDef := GenReqDefForMigrateLogicDb()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateLogicDbResponse), nil
	}
}

// MigrateLogicDbInvoker 分片变更
func (c *DdmClient) MigrateLogicDbInvoker(request *model.MigrateLogicDbRequest) *MigrateLogicDbInvoker {
	requestDef := GenReqDefForMigrateLogicDb()
	return &MigrateLogicDbInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateResults 查询分片变更任务详情
//
// 查询分片变更任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) MigrateResults(request *model.MigrateResultsRequest) (*model.MigrateResultsResponse, error) {
	requestDef := GenReqDefForMigrateResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateResultsResponse), nil
	}
}

// MigrateResultsInvoker 查询分片变更任务详情
func (c *DdmClient) MigrateResultsInvoker(request *model.MigrateResultsRequest) *MigrateResultsInvoker {
	requestDef := GenReqDefForMigrateResults()
	return &MigrateResultsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyEip 修改实例的ELB IP
//
// 修改实例的ELB IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ModifyEip(request *model.ModifyEipRequest) (*model.ModifyEipResponse, error) {
	requestDef := GenReqDefForModifyEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyEipResponse), nil
	}
}

// ModifyEipInvoker 修改实例的ELB IP
func (c *DdmClient) ModifyEipInvoker(request *model.ModifyEipRequest) *ModifyEipInvoker {
	requestDef := GenReqDefForModifyEip()
	return &ModifyEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebuildConfig DDM表数据重载
//
// DDM实例跨region容灾场景下，针对目标DDM实例实现表数据reload，使数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RebuildConfig(request *model.RebuildConfigRequest) (*model.RebuildConfigResponse, error) {
	requestDef := GenReqDefForRebuildConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebuildConfigResponse), nil
	}
}

// RebuildConfigInvoker DDM表数据重载
func (c *DdmClient) RebuildConfigInvoker(request *model.RebuildConfigRequest) *RebuildConfigInvoker {
	requestDef := GenReqDefForRebuildConfig()
	return &RebuildConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetAdministrator DDM管理员账号密码管理
//
// 首次调用时新建DDM管理员帐号并设置密码。后续调用时仅更新管理员密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ResetAdministrator(request *model.ResetAdministratorRequest) (*model.ResetAdministratorResponse, error) {
	requestDef := GenReqDefForResetAdministrator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetAdministratorResponse), nil
	}
}

// ResetAdministratorInvoker DDM管理员账号密码管理
func (c *DdmClient) ResetAdministratorInvoker(request *model.ResetAdministratorRequest) *ResetAdministratorInvoker {
	requestDef := GenReqDefForResetAdministrator()
	return &ResetAdministratorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetUserPassword 重置DDM账号密码
//
// 重置现有DDM帐号的密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ResetUserPassword(request *model.ResetUserPasswordRequest) (*model.ResetUserPasswordResponse, error) {
	requestDef := GenReqDefForResetUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetUserPasswordResponse), nil
	}
}

// ResetUserPasswordInvoker 重置DDM账号密码
func (c *DdmClient) ResetUserPasswordInvoker(request *model.ResetUserPasswordRequest) *ResetUserPasswordInvoker {
	requestDef := GenReqDefForResetUserPassword()
	return &ResetUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeFlavor 变更DDM实例规格
//
// 变更DDM实例规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ResizeFlavor(request *model.ResizeFlavorRequest) (*model.ResizeFlavorResponse, error) {
	requestDef := GenReqDefForResizeFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeFlavorResponse), nil
	}
}

// ResizeFlavorInvoker 变更DDM实例规格
func (c *DdmClient) ResizeFlavorInvoker(request *model.ResizeFlavorRequest) *ResizeFlavorInvoker {
	requestDef := GenReqDefForResizeFlavor()
	return &ResizeFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartDdmInstance 重启DDM实例
//
// 重启DDM实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RestartDdmInstance(request *model.RestartDdmInstanceRequest) (*model.RestartDdmInstanceResponse, error) {
	requestDef := GenReqDefForRestartDdmInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartDdmInstanceResponse), nil
	}
}

// RestartDdmInstanceInvoker 重启DDM实例
func (c *DdmClient) RestartDdmInstanceInvoker(request *model.RestartDdmInstanceRequest) *RestartDdmInstanceInvoker {
	requestDef := GenReqDefForRestartDdmInstance()
	return &RestartDdmInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartInstance 重启DDM实例
//
// 重启指定的DDM实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

// RestartInstanceInvoker 重启DDM实例
func (c *DdmClient) RestartInstanceInvoker(request *model.RestartInstanceRequest) *RestartInstanceInvoker {
	requestDef := GenReqDefForRestartInstance()
	return &RestartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartNode 重启DDM节点
//
// 重启DDM节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RestartNode(request *model.RestartNodeRequest) (*model.RestartNodeResponse, error) {
	requestDef := GenReqDefForRestartNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartNodeResponse), nil
	}
}

// RestartNodeInvoker 重启DDM节点
func (c *DdmClient) RestartNodeInvoker(request *model.RestartNodeRequest) *RestartNodeInvoker {
	requestDef := GenReqDefForRestartNode()
	return &RestartNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Restore2Exist 恢复到新实例
//
// 恢复到新实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) Restore2Exist(request *model.Restore2ExistRequest) (*model.Restore2ExistResponse, error) {
	requestDef := GenReqDefForRestore2Exist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.Restore2ExistResponse), nil
	}
}

// Restore2ExistInvoker 恢复到新实例
func (c *DdmClient) Restore2ExistInvoker(request *model.Restore2ExistRequest) *Restore2ExistInvoker {
	requestDef := GenReqDefForRestore2Exist()
	return &Restore2ExistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreMetadata 元数据恢复
//
// 元数据恢复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RestoreMetadata(request *model.RestoreMetadataRequest) (*model.RestoreMetadataResponse, error) {
	requestDef := GenReqDefForRestoreMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreMetadataResponse), nil
	}
}

// RestoreMetadataInvoker 元数据恢复
func (c *DdmClient) RestoreMetadataInvoker(request *model.RestoreMetadataRequest) *RestoreMetadataInvoker {
	requestDef := GenReqDefForRestoreMetadata()
	return &RestoreMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryMigration 重试分片变更
//
// 重试分片变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RetryMigration(request *model.RetryMigrationRequest) (*model.RetryMigrationResponse, error) {
	requestDef := GenReqDefForRetryMigration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryMigrationResponse), nil
	}
}

// RetryMigrationInvoker 重试分片变更
func (c *DdmClient) RetryMigrationInvoker(request *model.RetryMigrationRequest) *RetryMigrationInvoker {
	requestDef := GenReqDefForRetryMigration()
	return &RetryMigrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollbackMigration 回滚分片变更
//
// 回滚分片变更
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) RollbackMigration(request *model.RollbackMigrationRequest) (*model.RollbackMigrationResponse, error) {
	requestDef := GenReqDefForRollbackMigration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackMigrationResponse), nil
	}
}

// RollbackMigrationInvoker 回滚分片变更
func (c *DdmClient) RollbackMigrationInvoker(request *model.RollbackMigrationRequest) *RollbackMigrationInvoker {
	requestDef := GenReqDefForRollbackMigration()
	return &RollbackMigrationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvalibleDdms 查询可用于恢复的实例列表
//
// 查询可用于恢复的实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowAvalibleDdms(request *model.ShowAvalibleDdmsRequest) (*model.ShowAvalibleDdmsResponse, error) {
	requestDef := GenReqDefForShowAvalibleDdms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvalibleDdmsResponse), nil
	}
}

// ShowAvalibleDdmsInvoker 查询可用于恢复的实例列表
func (c *DdmClient) ShowAvalibleDdmsInvoker(request *model.ShowAvalibleDdmsRequest) *ShowAvalibleDdmsInvoker {
	requestDef := GenReqDefForShowAvalibleDdms()
	return &ShowAvalibleDdmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvalibleRds 查询可用于时间点恢复的数据节点列表
//
// 查询可用于时间点恢复的数据节点列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowAvalibleRds(request *model.ShowAvalibleRdsRequest) (*model.ShowAvalibleRdsResponse, error) {
	requestDef := GenReqDefForShowAvalibleRds()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvalibleRdsResponse), nil
	}
}

// ShowAvalibleRdsInvoker 查询可用于时间点恢复的数据节点列表
func (c *DdmClient) ShowAvalibleRdsInvoker(request *model.ShowAvalibleRdsRequest) *ShowAvalibleRdsInvoker {
	requestDef := GenReqDefForShowAvalibleRds()
	return &ShowAvalibleRdsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvalibleTime 查询可恢复时间段
//
// 查询可恢复时间段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowAvalibleTime(request *model.ShowAvalibleTimeRequest) (*model.ShowAvalibleTimeResponse, error) {
	requestDef := GenReqDefForShowAvalibleTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvalibleTimeResponse), nil
	}
}

// ShowAvalibleTimeInvoker 查询可恢复时间段
func (c *DdmClient) ShowAvalibleTimeInvoker(request *model.ShowAvalibleTimeRequest) *ShowAvalibleTimeInvoker {
	requestDef := GenReqDefForShowAvalibleTime()
	return &ShowAvalibleTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackup 查询备份详情
//
// 查询备份详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowBackup(request *model.ShowBackupRequest) (*model.ShowBackupResponse, error) {
	requestDef := GenReqDefForShowBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupResponse), nil
	}
}

// ShowBackupInvoker 查询备份详情
func (c *DdmClient) ShowBackupInvoker(request *model.ShowBackupRequest) *ShowBackupInvoker {
	requestDef := GenReqDefForShowBackup()
	return &ShowBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabase 查询DDM逻辑库详细信息
//
// 查询指定逻辑库的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowDatabase(request *model.ShowDatabaseRequest) (*model.ShowDatabaseResponse, error) {
	requestDef := GenReqDefForShowDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseResponse), nil
	}
}

// ShowDatabaseInvoker 查询DDM逻辑库详细信息
func (c *DdmClient) ShowDatabaseInvoker(request *model.ShowDatabaseRequest) *ShowDatabaseInvoker {
	requestDef := GenReqDefForShowDatabase()
	return &ShowDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdmJobResult 获取指定ID的任务信息
//
// 获取指定ID的任务信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowDdmJobResult(request *model.ShowDdmJobResultRequest) (*model.ShowDdmJobResultResponse, error) {
	requestDef := GenReqDefForShowDdmJobResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdmJobResultResponse), nil
	}
}

// ShowDdmJobResultInvoker 获取指定ID的任务信息
func (c *DdmClient) ShowDdmJobResultInvoker(request *model.ShowDdmJobResultRequest) *ShowDdmJobResultInvoker {
	requestDef := GenReqDefForShowDdmJobResult()
	return &ShowDdmJobResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdmNodeDetail 查询DDM实例节点详情
//
// 查询DDM实例节点详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowDdmNodeDetail(request *model.ShowDdmNodeDetailRequest) (*model.ShowDdmNodeDetailResponse, error) {
	requestDef := GenReqDefForShowDdmNodeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdmNodeDetailResponse), nil
	}
}

// ShowDdmNodeDetailInvoker 查询DDM实例节点详情
func (c *DdmClient) ShowDdmNodeDetailInvoker(request *model.ShowDdmNodeDetailRequest) *ShowDdmNodeDetailInvoker {
	requestDef := GenReqDefForShowDdmNodeDetail()
	return &ShowDdmNodeDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstance 查询DDM实例详情
//
// 查询指定DDM实例的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowInstance(request *model.ShowInstanceRequest) (*model.ShowInstanceResponse, error) {
	requestDef := GenReqDefForShowInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResponse), nil
	}
}

// ShowInstanceInvoker 查询DDM实例详情
func (c *DdmClient) ShowInstanceInvoker(request *model.ShowInstanceRequest) *ShowInstanceInvoker {
	requestDef := GenReqDefForShowInstance()
	return &ShowInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceDatabase 查询逻辑库详情
//
// 查询逻辑库详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowInstanceDatabase(request *model.ShowInstanceDatabaseRequest) (*model.ShowInstanceDatabaseResponse, error) {
	requestDef := GenReqDefForShowInstanceDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceDatabaseResponse), nil
	}
}

// ShowInstanceDatabaseInvoker 查询逻辑库详情
func (c *DdmClient) ShowInstanceDatabaseInvoker(request *model.ShowInstanceDatabaseRequest) *ShowInstanceDatabaseInvoker {
	requestDef := GenReqDefForShowInstanceDatabase()
	return &ShowInstanceDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceParam 查询DDM指定实例的参数详情
//
// 查询DDM指定实例的参数详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowInstanceParam(request *model.ShowInstanceParamRequest) (*model.ShowInstanceParamResponse, error) {
	requestDef := GenReqDefForShowInstanceParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceParamResponse), nil
	}
}

// ShowInstanceParamInvoker 查询DDM指定实例的参数详情
func (c *DdmClient) ShowInstanceParamInvoker(request *model.ShowInstanceParamRequest) *ShowInstanceParamInvoker {
	requestDef := GenReqDefForShowInstanceParam()
	return &ShowInstanceParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIpGroup 查询访问控制组
//
// 查询访问控制组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowIpGroup(request *model.ShowIpGroupRequest) (*model.ShowIpGroupResponse, error) {
	requestDef := GenReqDefForShowIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIpGroupResponse), nil
	}
}

// ShowIpGroupInvoker 查询访问控制组
func (c *DdmClient) ShowIpGroupInvoker(request *model.ShowIpGroupRequest) *ShowIpGroupInvoker {
	requestDef := GenReqDefForShowIpGroup()
	return &ShowIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalProcesses 查询逻辑会话列表
//
// 查询逻辑会话列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowLogicalProcesses(request *model.ShowLogicalProcessesRequest) (*model.ShowLogicalProcessesResponse, error) {
	requestDef := GenReqDefForShowLogicalProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalProcessesResponse), nil
	}
}

// ShowLogicalProcessesInvoker 查询逻辑会话列表
func (c *DdmClient) ShowLogicalProcessesInvoker(request *model.ShowLogicalProcessesRequest) *ShowLogicalProcessesInvoker {
	requestDef := GenReqDefForShowLogicalProcesses()
	return &ShowLogicalProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMigrationLog 查询分片变更的日志
//
// 查询分片变更的日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowMigrationLog(request *model.ShowMigrationLogRequest) (*model.ShowMigrationLogResponse, error) {
	requestDef := GenReqDefForShowMigrationLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMigrationLogResponse), nil
	}
}

// ShowMigrationLogInvoker 查询分片变更的日志
func (c *DdmClient) ShowMigrationLogInvoker(request *model.ShowMigrationLogRequest) *ShowMigrationLogInvoker {
	requestDef := GenReqDefForShowMigrationLog()
	return &ShowMigrationLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNode 查询DDM实例节点详情
//
// 查询DDM实例节点详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowNode(request *model.ShowNodeRequest) (*model.ShowNodeResponse, error) {
	requestDef := GenReqDefForShowNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeResponse), nil
	}
}

// ShowNodeInvoker 查询DDM实例节点详情
func (c *DdmClient) ShowNodeInvoker(request *model.ShowNodeRequest) *ShowNodeInvoker {
	requestDef := GenReqDefForShowNode()
	return &ShowNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPhysicalProcesses 查询物理会话列表
//
// 查询物理会话列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowPhysicalProcesses(request *model.ShowPhysicalProcessesRequest) (*model.ShowPhysicalProcessesResponse, error) {
	requestDef := GenReqDefForShowPhysicalProcesses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPhysicalProcessesResponse), nil
	}
}

// ShowPhysicalProcessesInvoker 查询物理会话列表
func (c *DdmClient) ShowPhysicalProcessesInvoker(request *model.ShowPhysicalProcessesRequest) *ShowPhysicalProcessesInvoker {
	requestDef := GenReqDefForShowPhysicalProcesses()
	return &ShowPhysicalProcessesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProcessesAuditLog 查询kill会话审计日志
//
// 查询kill会话审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowProcessesAuditLog(request *model.ShowProcessesAuditLogRequest) (*model.ShowProcessesAuditLogResponse, error) {
	requestDef := GenReqDefForShowProcessesAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProcessesAuditLogResponse), nil
	}
}

// ShowProcessesAuditLogInvoker 查询kill会话审计日志
func (c *DdmClient) ShowProcessesAuditLogInvoker(request *model.ShowProcessesAuditLogRequest) *ShowProcessesAuditLogInvoker {
	requestDef := GenReqDefForShowProcessesAuditLog()
	return &ShowProcessesAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPublicIp 获取DDM实例绑定的弹性公网IP信息
//
// 获取DDM实例绑定的弹性公网IP信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowPublicIp(request *model.ShowPublicIpRequest) (*model.ShowPublicIpResponse, error) {
	requestDef := GenReqDefForShowPublicIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicIpResponse), nil
	}
}

// ShowPublicIpInvoker 获取DDM实例绑定的弹性公网IP信息
func (c *DdmClient) ShowPublicIpInvoker(request *model.ShowPublicIpRequest) *ShowPublicIpInvoker {
	requestDef := GenReqDefForShowPublicIp()
	return &ShowPublicIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRelatedDns 查询实例在恢复时间点关联的数据节点
//
// 查询实例在恢复时间点关联的数据节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowRelatedDns(request *model.ShowRelatedDnsRequest) (*model.ShowRelatedDnsResponse, error) {
	requestDef := GenReqDefForShowRelatedDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRelatedDnsResponse), nil
	}
}

// ShowRelatedDnsInvoker 查询实例在恢复时间点关联的数据节点
func (c *DdmClient) ShowRelatedDnsInvoker(request *model.ShowRelatedDnsRequest) *ShowRelatedDnsInvoker {
	requestDef := GenReqDefForShowRelatedDns()
	return &ShowRelatedDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkInstanceNodes DDM实例节点缩容
//
// 对指定的DDM实例的节点个数进行缩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShrinkInstanceNodes(request *model.ShrinkInstanceNodesRequest) (*model.ShrinkInstanceNodesResponse, error) {
	requestDef := GenReqDefForShrinkInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkInstanceNodesResponse), nil
	}
}

// ShrinkInstanceNodesInvoker DDM实例节点缩容
func (c *DdmClient) ShrinkInstanceNodesInvoker(request *model.ShrinkInstanceNodesRequest) *ShrinkInstanceNodesInvoker {
	requestDef := GenReqDefForShrinkInstanceNodes()
	return &ShrinkInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchIpGroup 创建访问控制组
//
// 创建访问控制组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) SwitchIpGroup(request *model.SwitchIpGroupRequest) (*model.SwitchIpGroupResponse, error) {
	requestDef := GenReqDefForSwitchIpGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchIpGroupResponse), nil
	}
}

// SwitchIpGroupInvoker 创建访问控制组
func (c *DdmClient) SwitchIpGroupInvoker(request *model.SwitchIpGroupRequest) *SwitchIpGroupInvoker {
	requestDef := GenReqDefForSwitchIpGroup()
	return &SwitchIpGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchRoute 切换路由
//
// 切换路由
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) SwitchRoute(request *model.SwitchRouteRequest) (*model.SwitchRouteResponse, error) {
	requestDef := GenReqDefForSwitchRoute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchRouteResponse), nil
	}
}

// SwitchRouteInvoker 切换路由
func (c *DdmClient) SwitchRouteInvoker(request *model.SwitchRouteRequest) *SwitchRouteInvoker {
	requestDef := GenReqDefForSwitchRoute()
	return &SwitchRouteInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSsl 为实例设置SSL数据加密
//
// 为实例设置SSL数据加密。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

// SwitchSslInvoker 为实例设置SSL数据加密
func (c *DdmClient) SwitchSslInvoker(request *model.SwitchSslRequest) *SwitchSslInvoker {
	requestDef := GenReqDefForSwitchSsl()
	return &SwitchSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncDnInformation 同步数据节点
//
// 同步数据节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) SyncDnInformation(request *model.SyncDnInformationRequest) (*model.SyncDnInformationResponse, error) {
	requestDef := GenReqDefForSyncDnInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncDnInformationResponse), nil
	}
}

// SyncDnInformationInvoker 同步数据节点
func (c *DdmClient) SyncDnInformationInvoker(request *model.SyncDnInformationRequest) *SyncDnInformationInvoker {
	requestDef := GenReqDefForSyncDnInformation()
	return &SyncDnInformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindEip 解绑弹性公网IP
//
// 解绑弹性公网IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UnbindEip(request *model.UnbindEipRequest) (*model.UnbindEipResponse, error) {
	requestDef := GenReqDefForUnbindEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindEipResponse), nil
	}
}

// UnbindEipInvoker 解绑弹性公网IP
func (c *DdmClient) UnbindEipInvoker(request *model.UnbindEipRequest) *UnbindEipInvoker {
	requestDef := GenReqDefForUnbindEip()
	return &UnbindEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatabaseInfo 同步DN信息
//
// 同步当前DDM实例已关联的所有DN实例配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateDatabaseInfo(request *model.UpdateDatabaseInfoRequest) (*model.UpdateDatabaseInfoResponse, error) {
	requestDef := GenReqDefForUpdateDatabaseInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseInfoResponse), nil
	}
}

// UpdateDatabaseInfoInvoker 同步DN信息
func (c *DdmClient) UpdateDatabaseInfoInvoker(request *model.UpdateDatabaseInfoRequest) *UpdateDatabaseInfoInvoker {
	requestDef := GenReqDefForUpdateDatabaseInfo()
	return &UpdateDatabaseInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceName 修改DDM实例名称
//
// 修改DDM实例名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

// UpdateInstanceNameInvoker 修改DDM实例名称
func (c *DdmClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceParam 修改DDM实例参数
//
// 修改DDM实例参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateInstanceParam(request *model.UpdateInstanceParamRequest) (*model.UpdateInstanceParamResponse, error) {
	requestDef := GenReqDefForUpdateInstanceParam()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceParamResponse), nil
	}
}

// UpdateInstanceParamInvoker 修改DDM实例参数
func (c *DdmClient) UpdateInstanceParamInvoker(request *model.UpdateInstanceParamRequest) *UpdateInstanceParamInvoker {
	requestDef := GenReqDefForUpdateInstanceParam()
	return &UpdateInstanceParamInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstancePort 修改实例端口
//
// 修改实例端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateInstancePort(request *model.UpdateInstancePortRequest) (*model.UpdateInstancePortResponse, error) {
	requestDef := GenReqDefForUpdateInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstancePortResponse), nil
	}
}

// UpdateInstancePortInvoker 修改实例端口
func (c *DdmClient) UpdateInstancePortInvoker(request *model.UpdateInstancePortRequest) *UpdateInstancePortInvoker {
	requestDef := GenReqDefForUpdateInstancePort()
	return &UpdateInstancePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceSecurityGroup 修改DDM实例安全组
//
// 修改DDM实例安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateInstanceSecurityGroup(request *model.UpdateInstanceSecurityGroupRequest) (*model.UpdateInstanceSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateInstanceSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceSecurityGroupResponse), nil
	}
}

// UpdateInstanceSecurityGroupInvoker 修改DDM实例安全组
func (c *DdmClient) UpdateInstanceSecurityGroupInvoker(request *model.UpdateInstanceSecurityGroupRequest) *UpdateInstanceSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateInstanceSecurityGroup()
	return &UpdateInstanceSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReadAndWriteStrategy 修改DDM已关联的数据库实例的读策略
//
// 修改DDM已关联的数据库实例的读策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateReadAndWriteStrategy(request *model.UpdateReadAndWriteStrategyRequest) (*model.UpdateReadAndWriteStrategyResponse, error) {
	requestDef := GenReqDefForUpdateReadAndWriteStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReadAndWriteStrategyResponse), nil
	}
}

// UpdateReadAndWriteStrategyInvoker 修改DDM已关联的数据库实例的读策略
func (c *DdmClient) UpdateReadAndWriteStrategyInvoker(request *model.UpdateReadAndWriteStrategyRequest) *UpdateReadAndWriteStrategyInvoker {
	requestDef := GenReqDefForUpdateReadAndWriteStrategy()
	return &UpdateReadAndWriteStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUser 修改DDM帐号
//
// 修改现有DDM帐号的权限或者与逻辑库的管理关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUserResponse), nil
	}
}

// UpdateUserInvoker 修改DDM帐号
func (c *DdmClient) UpdateUserInvoker(request *model.UpdateUserRequest) *UpdateUserInvoker {
	requestDef := GenReqDefForUpdateUser()
	return &UpdateUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadSchemaMetadata 导入逻辑库元数据
//
// 导入所有逻辑库物理分片分布关系，以此创建相同物理分片分布关系的逻辑库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) UploadSchemaMetadata(request *model.UploadSchemaMetadataRequest) (*model.UploadSchemaMetadataResponse, error) {
	requestDef := GenReqDefForUploadSchemaMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadSchemaMetadataResponse), nil
	}
}

// UploadSchemaMetadataInvoker 导入逻辑库元数据
func (c *DdmClient) UploadSchemaMetadataInvoker(request *model.UploadSchemaMetadataRequest) *UploadSchemaMetadataInvoker {
	requestDef := GenReqDefForUploadSchemaMetadata()
	return &UploadSchemaMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateWeakPassword 弱密码校验
//
// 弱密码校验
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ValidateWeakPassword(request *model.ValidateWeakPasswordRequest) (*model.ValidateWeakPasswordResponse, error) {
	requestDef := GenReqDefForValidateWeakPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateWeakPasswordResponse), nil
	}
}

// ValidateWeakPasswordInvoker 弱密码校验
func (c *DdmClient) ValidateWeakPasswordInvoker(request *model.ValidateWeakPasswordRequest) *ValidateWeakPasswordInvoker {
	requestDef := GenReqDefForValidateWeakPassword()
	return &ValidateWeakPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDataNodeConnection rds连通性检查V3
//
// rds连通性检查V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CheckDataNodeConnection(request *model.CheckDataNodeConnectionRequest) (*model.CheckDataNodeConnectionResponse, error) {
	requestDef := GenReqDefForCheckDataNodeConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDataNodeConnectionResponse), nil
	}
}

// CheckDataNodeConnectionInvoker rds连通性检查V3
func (c *DdmClient) CheckDataNodeConnectionInvoker(request *model.CheckDataNodeConnectionRequest) *CheckDataNodeConnectionInvoker {
	requestDef := GenReqDefForCheckDataNodeConnection()
	return &CheckDataNodeConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareParameterGroups 比较参数组V3
//
// 比较参数组V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CompareParameterGroups(request *model.CompareParameterGroupsRequest) (*model.CompareParameterGroupsResponse, error) {
	requestDef := GenReqDefForCompareParameterGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareParameterGroupsResponse), nil
	}
}

// CompareParameterGroupsInvoker 比较参数组V3
func (c *DdmClient) CompareParameterGroupsInvoker(request *model.CompareParameterGroupsRequest) *CompareParameterGroupsInvoker {
	requestDef := GenReqDefForCompareParameterGroups()
	return &CompareParameterGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyConfiguration 复制参数组V3
//
// 复制参数组V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

// CopyConfigurationInvoker 复制参数组V3
func (c *DdmClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationApplyHistories 参数组应用记录V3
//
// 参数组应用记录V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListConfigurationApplyHistories(request *model.ListConfigurationApplyHistoriesRequest) (*model.ListConfigurationApplyHistoriesResponse, error) {
	requestDef := GenReqDefForListConfigurationApplyHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationApplyHistoriesResponse), nil
	}
}

// ListConfigurationApplyHistoriesInvoker 参数组应用记录V3
func (c *DdmClient) ListConfigurationApplyHistoriesInvoker(request *model.ListConfigurationApplyHistoriesRequest) *ListConfigurationApplyHistoriesInvoker {
	requestDef := GenReqDefForListConfigurationApplyHistories()
	return &ListConfigurationApplyHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesConfigurations 查询可应用的实例列表V3
//
// 查询可应用的实例列表V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListInstancesConfigurations(request *model.ListInstancesConfigurationsRequest) (*model.ListInstancesConfigurationsResponse, error) {
	requestDef := GenReqDefForListInstancesConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesConfigurationsResponse), nil
	}
}

// ListInstancesConfigurationsInvoker 查询可应用的实例列表V3
func (c *DdmClient) ListInstancesConfigurationsInvoker(request *model.ListInstancesConfigurationsRequest) *ListInstancesConfigurationsInvoker {
	requestDef := GenReqDefForListInstancesConfigurations()
	return &ListInstancesConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetParameterGroup 更新参数组V3
//
// 更新参数组V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ResetParameterGroup(request *model.ResetParameterGroupRequest) (*model.ResetParameterGroupResponse, error) {
	requestDef := GenReqDefForResetParameterGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetParameterGroupResponse), nil
	}
}

// ResetParameterGroupInvoker 更新参数组V3
func (c *DdmClient) ResetParameterGroupInvoker(request *model.ResetParameterGroupRequest) *ResetParameterGroupInvoker {
	requestDef := GenReqDefForResetParameterGroup()
	return &ResetParameterGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDdmDetail 查询实例详情V3
//
// 查询实例详情V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ShowDdmDetail(request *model.ShowDdmDetailRequest) (*model.ShowDdmDetailResponse, error) {
	requestDef := GenReqDefForShowDdmDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDdmDetailResponse), nil
	}
}

// ShowDdmDetailInvoker 查询实例详情V3
func (c *DdmClient) ShowDdmDetailInvoker(request *model.ShowDdmDetailRequest) *ShowDdmDetailInvoker {
	requestDef := GenReqDefForShowDdmDetail()
	return &ShowDdmDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchModifyReadWriteStrategy 批量设置读策略V3
//
// 批量设置读策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) BatchModifyReadWriteStrategy(request *model.BatchModifyReadWriteStrategyRequest) (*model.BatchModifyReadWriteStrategyResponse, error) {
	requestDef := GenReqDefForBatchModifyReadWriteStrategy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchModifyReadWriteStrategyResponse), nil
	}
}

// BatchModifyReadWriteStrategyInvoker 批量设置读策略V3
func (c *DdmClient) BatchModifyReadWriteStrategyInvoker(request *model.BatchModifyReadWriteStrategyRequest) *BatchModifyReadWriteStrategyInvoker {
	requestDef := GenReqDefForBatchModifyReadWriteStrategy()
	return &BatchModifyReadWriteStrategyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifySqlBlackList 修改sql黑名单
//
// 修改sql黑名单V3
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ModifySqlBlackList(request *model.ModifySqlBlackListRequest) (*model.ModifySqlBlackListResponse, error) {
	requestDef := GenReqDefForModifySqlBlackList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifySqlBlackListResponse), nil
	}
}

// ModifySqlBlackListInvoker 修改sql黑名单
func (c *DdmClient) ModifySqlBlackListInvoker(request *model.ModifySqlBlackListRequest) *ModifySqlBlackListInvoker {
	requestDef := GenReqDefForModifySqlBlackList()
	return &ModifySqlBlackListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlBlack 查询sql黑名单V3
//
// 查询sql黑名单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdmClient) ListSqlBlack(request *model.ListSqlBlackRequest) (*model.ListSqlBlackResponse, error) {
	requestDef := GenReqDefForListSqlBlack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlBlackResponse), nil
	}
}

// ListSqlBlackInvoker 查询sql黑名单V3
func (c *DdmClient) ListSqlBlackInvoker(request *model.ListSqlBlackRequest) *ListSqlBlackInvoker {
	requestDef := GenReqDefForListSqlBlack()
	return &ListSqlBlackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
