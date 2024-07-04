package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type GaussDBClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewGaussDBClient(hcClient *httpclient.HcHttpClient) *GaussDBClient {
	return &GaussDBClient{HcClient: hcClient}
}

func GaussDBClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddDatabasePermission 授予数据库用户数据库权限
//
// 授予云数据库 GaussDB(for MySQL)实例数据库用户数据库权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) AddDatabasePermission(request *model.AddDatabasePermissionRequest) (*model.AddDatabasePermissionResponse, error) {
	requestDef := GenReqDefForAddDatabasePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDatabasePermissionResponse), nil
	}
}

// AddDatabasePermissionInvoker 授予数据库用户数据库权限
func (c *GaussDBClient) AddDatabasePermissionInvoker(request *model.AddDatabasePermissionRequest) *AddDatabasePermissionInvoker {
	requestDef := GenReqDefForAddDatabasePermission()
	return &AddDatabasePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchTagAction 批量添加或删除标签
//
// 批量添加或删除指定实例的标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

// BatchTagActionInvoker 批量添加或删除标签
func (c *GaussDBClient) BatchTagActionInvoker(request *model.BatchTagActionRequest) *BatchTagActionInvoker {
	requestDef := GenReqDefForBatchTagAction()
	return &BatchTagActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelGaussMySqlInstanceEip 解绑弹性公网IP
//
// 实例解绑弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CancelGaussMySqlInstanceEip(request *model.CancelGaussMySqlInstanceEipRequest) (*model.CancelGaussMySqlInstanceEipResponse, error) {
	requestDef := GenReqDefForCancelGaussMySqlInstanceEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelGaussMySqlInstanceEipResponse), nil
	}
}

// CancelGaussMySqlInstanceEipInvoker 解绑弹性公网IP
func (c *GaussDBClient) CancelGaussMySqlInstanceEipInvoker(request *model.CancelGaussMySqlInstanceEipRequest) *CancelGaussMySqlInstanceEipInvoker {
	requestDef := GenReqDefForCancelGaussMySqlInstanceEip()
	return &CancelGaussMySqlInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelScheduleTask 取消定时任务
//
// 取消定时任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CancelScheduleTask(request *model.CancelScheduleTaskRequest) (*model.CancelScheduleTaskResponse, error) {
	requestDef := GenReqDefForCancelScheduleTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelScheduleTaskResponse), nil
	}
}

// CancelScheduleTaskInvoker 取消定时任务
func (c *GaussDBClient) CancelScheduleTaskInvoker(request *model.CancelScheduleTaskRequest) *CancelScheduleTaskInvoker {
	requestDef := GenReqDefForCancelScheduleTask()
	return &CancelScheduleTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeGaussMySqlInstanceSpecification 变更实例规格
//
// 变更数据库实例的规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ChangeGaussMySqlInstanceSpecification(request *model.ChangeGaussMySqlInstanceSpecificationRequest) (*model.ChangeGaussMySqlInstanceSpecificationResponse, error) {
	requestDef := GenReqDefForChangeGaussMySqlInstanceSpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeGaussMySqlInstanceSpecificationResponse), nil
	}
}

// ChangeGaussMySqlInstanceSpecificationInvoker 变更实例规格
func (c *GaussDBClient) ChangeGaussMySqlInstanceSpecificationInvoker(request *model.ChangeGaussMySqlInstanceSpecificationRequest) *ChangeGaussMySqlInstanceSpecificationInvoker {
	requestDef := GenReqDefForChangeGaussMySqlInstanceSpecification()
	return &ChangeGaussMySqlInstanceSpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeGaussMySqlProxySpecification 数据库代理规格变更
//
// 数据库代理规格变更。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ChangeGaussMySqlProxySpecification(request *model.ChangeGaussMySqlProxySpecificationRequest) (*model.ChangeGaussMySqlProxySpecificationResponse, error) {
	requestDef := GenReqDefForChangeGaussMySqlProxySpecification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeGaussMySqlProxySpecificationResponse), nil
	}
}

// ChangeGaussMySqlProxySpecificationInvoker 数据库代理规格变更
func (c *GaussDBClient) ChangeGaussMySqlProxySpecificationInvoker(request *model.ChangeGaussMySqlProxySpecificationRequest) *ChangeGaussMySqlProxySpecificationInvoker {
	requestDef := GenReqDefForChangeGaussMySqlProxySpecification()
	return &ChangeGaussMySqlProxySpecificationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckResource 资源预校验
//
// 资源预校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckResource(request *model.CheckResourceRequest) (*model.CheckResourceResponse, error) {
	requestDef := GenReqDefForCheckResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckResourceResponse), nil
	}
}

// CheckResourceInvoker 资源预校验
func (c *GaussDBClient) CheckResourceInvoker(request *model.CheckResourceRequest) *CheckResourceInvoker {
	requestDef := GenReqDefForCheckResource()
	return &CheckResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyConfigurations 复制参数组
//
// 复制参数组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CopyConfigurations(request *model.CopyConfigurationsRequest) (*model.CopyConfigurationsResponse, error) {
	requestDef := GenReqDefForCopyConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationsResponse), nil
	}
}

// CopyConfigurationsInvoker 复制参数组
func (c *GaussDBClient) CopyConfigurationsInvoker(request *model.CopyConfigurationsRequest) *CopyConfigurationsInvoker {
	requestDef := GenReqDefForCopyConfigurations()
	return &CopyConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyInstanceConfigurations 复制实例参数组
//
// 复制实例参数组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CopyInstanceConfigurations(request *model.CopyInstanceConfigurationsRequest) (*model.CopyInstanceConfigurationsResponse, error) {
	requestDef := GenReqDefForCopyInstanceConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyInstanceConfigurationsResponse), nil
	}
}

// CopyInstanceConfigurationsInvoker 复制实例参数组
func (c *GaussDBClient) CopyInstanceConfigurationsInvoker(request *model.CopyInstanceConfigurationsRequest) *CopyInstanceConfigurationsInvoker {
	requestDef := GenReqDefForCopyInstanceConfigurations()
	return &CopyInstanceConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessControl 设置访问控制规则
//
// 设置访问控制规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateAccessControl(request *model.CreateAccessControlRequest) (*model.CreateAccessControlResponse, error) {
	requestDef := GenReqDefForCreateAccessControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessControlResponse), nil
	}
}

// CreateAccessControlInvoker 设置访问控制规则
func (c *GaussDBClient) CreateAccessControlInvoker(request *model.CreateAccessControlRequest) *CreateAccessControlInvoker {
	requestDef := GenReqDefForCreateAccessControl()
	return &CreateAccessControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlBackup 创建手动备份
//
// 创建手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlBackup(request *model.CreateGaussMySqlBackupRequest) (*model.CreateGaussMySqlBackupResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlBackupResponse), nil
	}
}

// CreateGaussMySqlBackupInvoker 创建手动备份
func (c *GaussDBClient) CreateGaussMySqlBackupInvoker(request *model.CreateGaussMySqlBackupRequest) *CreateGaussMySqlBackupInvoker {
	requestDef := GenReqDefForCreateGaussMySqlBackup()
	return &CreateGaussMySqlBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlConfiguration 创建参数模板
//
// 创建参数模板信息，包含参数模板名称、描述、数据库版本信息、参数值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlConfiguration(request *model.CreateGaussMySqlConfigurationRequest) (*model.CreateGaussMySqlConfigurationResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlConfigurationResponse), nil
	}
}

// CreateGaussMySqlConfigurationInvoker 创建参数模板
func (c *GaussDBClient) CreateGaussMySqlConfigurationInvoker(request *model.CreateGaussMySqlConfigurationRequest) *CreateGaussMySqlConfigurationInvoker {
	requestDef := GenReqDefForCreateGaussMySqlConfiguration()
	return &CreateGaussMySqlConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlDatabase 创建数据库
//
// 创建云数据库 GaussDB(for MySQL)实例数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlDatabase(request *model.CreateGaussMySqlDatabaseRequest) (*model.CreateGaussMySqlDatabaseResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlDatabaseResponse), nil
	}
}

// CreateGaussMySqlDatabaseInvoker 创建数据库
func (c *GaussDBClient) CreateGaussMySqlDatabaseInvoker(request *model.CreateGaussMySqlDatabaseRequest) *CreateGaussMySqlDatabaseInvoker {
	requestDef := GenReqDefForCreateGaussMySqlDatabase()
	return &CreateGaussMySqlDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlDatabaseUser 创建数据库用户
//
// 创建云数据库GaussDB(for MySQL)实例数据库用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlDatabaseUser(request *model.CreateGaussMySqlDatabaseUserRequest) (*model.CreateGaussMySqlDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlDatabaseUserResponse), nil
	}
}

// CreateGaussMySqlDatabaseUserInvoker 创建数据库用户
func (c *GaussDBClient) CreateGaussMySqlDatabaseUserInvoker(request *model.CreateGaussMySqlDatabaseUserRequest) *CreateGaussMySqlDatabaseUserInvoker {
	requestDef := GenReqDefForCreateGaussMySqlDatabaseUser()
	return &CreateGaussMySqlDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlInstance 创建数据库实例
//
// 创建云数据库GaussDB(for MySQL)实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlInstance(request *model.CreateGaussMySqlInstanceRequest) (*model.CreateGaussMySqlInstanceResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlInstanceResponse), nil
	}
}

// CreateGaussMySqlInstanceInvoker 创建数据库实例
func (c *GaussDBClient) CreateGaussMySqlInstanceInvoker(request *model.CreateGaussMySqlInstanceRequest) *CreateGaussMySqlInstanceInvoker {
	requestDef := GenReqDefForCreateGaussMySqlInstance()
	return &CreateGaussMySqlInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlProxy 开启数据库代理
//
// 开启数据库代理，只支持ELB模式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlProxy(request *model.CreateGaussMySqlProxyRequest) (*model.CreateGaussMySqlProxyResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlProxyResponse), nil
	}
}

// CreateGaussMySqlProxyInvoker 开启数据库代理
func (c *GaussDBClient) CreateGaussMySqlProxyInvoker(request *model.CreateGaussMySqlProxyRequest) *CreateGaussMySqlProxyInvoker {
	requestDef := GenReqDefForCreateGaussMySqlProxy()
	return &CreateGaussMySqlProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMySqlReadonlyNode 创建只读节点
//
// 创建只读节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMySqlReadonlyNode(request *model.CreateGaussMySqlReadonlyNodeRequest) (*model.CreateGaussMySqlReadonlyNodeResponse, error) {
	requestDef := GenReqDefForCreateGaussMySqlReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMySqlReadonlyNodeResponse), nil
	}
}

// CreateGaussMySqlReadonlyNodeInvoker 创建只读节点
func (c *GaussDBClient) CreateGaussMySqlReadonlyNodeInvoker(request *model.CreateGaussMySqlReadonlyNodeRequest) *CreateGaussMySqlReadonlyNodeInvoker {
	requestDef := GenReqDefForCreateGaussMySqlReadonlyNode()
	return &CreateGaussMySqlReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGaussMysqlDns 申请内网域名
//
// 申请内网域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateGaussMysqlDns(request *model.CreateGaussMysqlDnsRequest) (*model.CreateGaussMysqlDnsResponse, error) {
	requestDef := GenReqDefForCreateGaussMysqlDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGaussMysqlDnsResponse), nil
	}
}

// CreateGaussMysqlDnsInvoker 申请内网域名
func (c *GaussDBClient) CreateGaussMysqlDnsInvoker(request *model.CreateGaussMysqlDnsRequest) *CreateGaussMysqlDnsInvoker {
	requestDef := GenReqDefForCreateGaussMysqlDns()
	return &CreateGaussMysqlDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLtsConfigs 批量创建LTS日志配置
//
// 批量创建LTS日志配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateLtsConfigs(request *model.CreateLtsConfigsRequest) (*model.CreateLtsConfigsResponse, error) {
	requestDef := GenReqDefForCreateLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLtsConfigsResponse), nil
	}
}

// CreateLtsConfigsInvoker 批量创建LTS日志配置
func (c *GaussDBClient) CreateLtsConfigsInvoker(request *model.CreateLtsConfigsRequest) *CreateLtsConfigsInvoker {
	requestDef := GenReqDefForCreateLtsConfigs()
	return &CreateLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRestoreTables 表级时间点恢复
//
// 表级时间点恢复。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateRestoreTables(request *model.CreateRestoreTablesRequest) (*model.CreateRestoreTablesResponse, error) {
	requestDef := GenReqDefForCreateRestoreTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRestoreTablesResponse), nil
	}
}

// CreateRestoreTablesInvoker 表级时间点恢复
func (c *GaussDBClient) CreateRestoreTablesInvoker(request *model.CreateRestoreTablesRequest) *CreateRestoreTablesInvoker {
	requestDef := GenReqDefForCreateRestoreTables()
	return &CreateRestoreTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabasePermission 删除数据库用户的数据库权限
//
// 删除云数据库 GaussDB(for MySQL)实例数据库用户的数据库权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteDatabasePermission(request *model.DeleteDatabasePermissionRequest) (*model.DeleteDatabasePermissionResponse, error) {
	requestDef := GenReqDefForDeleteDatabasePermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabasePermissionResponse), nil
	}
}

// DeleteDatabasePermissionInvoker 删除数据库用户的数据库权限
func (c *GaussDBClient) DeleteDatabasePermissionInvoker(request *model.DeleteDatabasePermissionRequest) *DeleteDatabasePermissionInvoker {
	requestDef := GenReqDefForDeleteDatabasePermission()
	return &DeleteDatabasePermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlBackup 删除手动备份
//
// 删除手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlBackup(request *model.DeleteGaussMySqlBackupRequest) (*model.DeleteGaussMySqlBackupResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlBackupResponse), nil
	}
}

// DeleteGaussMySqlBackupInvoker 删除手动备份
func (c *GaussDBClient) DeleteGaussMySqlBackupInvoker(request *model.DeleteGaussMySqlBackupRequest) *DeleteGaussMySqlBackupInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlBackup()
	return &DeleteGaussMySqlBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlConfiguration 删除参数模板
//
// 删除指定参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlConfiguration(request *model.DeleteGaussMySqlConfigurationRequest) (*model.DeleteGaussMySqlConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlConfigurationResponse), nil
	}
}

// DeleteGaussMySqlConfigurationInvoker 删除参数模板
func (c *GaussDBClient) DeleteGaussMySqlConfigurationInvoker(request *model.DeleteGaussMySqlConfigurationRequest) *DeleteGaussMySqlConfigurationInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlConfiguration()
	return &DeleteGaussMySqlConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlDatabase 删除数据库
//
// 删除云数据库 GaussDB(for MySQL)实例数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlDatabase(request *model.DeleteGaussMySqlDatabaseRequest) (*model.DeleteGaussMySqlDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlDatabaseResponse), nil
	}
}

// DeleteGaussMySqlDatabaseInvoker 删除数据库
func (c *GaussDBClient) DeleteGaussMySqlDatabaseInvoker(request *model.DeleteGaussMySqlDatabaseRequest) *DeleteGaussMySqlDatabaseInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlDatabase()
	return &DeleteGaussMySqlDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlDatabaseUser 删除数据库用户
//
// 删除云数据库 GaussDB(for MySQL)实例数据库用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlDatabaseUser(request *model.DeleteGaussMySqlDatabaseUserRequest) (*model.DeleteGaussMySqlDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlDatabaseUserResponse), nil
	}
}

// DeleteGaussMySqlDatabaseUserInvoker 删除数据库用户
func (c *GaussDBClient) DeleteGaussMySqlDatabaseUserInvoker(request *model.DeleteGaussMySqlDatabaseUserRequest) *DeleteGaussMySqlDatabaseUserInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlDatabaseUser()
	return &DeleteGaussMySqlDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlInstance 删除/退订数据库实例
//
// 删除/退订数据库实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlInstance(request *model.DeleteGaussMySqlInstanceRequest) (*model.DeleteGaussMySqlInstanceResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlInstanceResponse), nil
	}
}

// DeleteGaussMySqlInstanceInvoker 删除/退订数据库实例
func (c *GaussDBClient) DeleteGaussMySqlInstanceInvoker(request *model.DeleteGaussMySqlInstanceRequest) *DeleteGaussMySqlInstanceInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlInstance()
	return &DeleteGaussMySqlInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlProxy 关闭数据库代理
//
// 关闭数据库代理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlProxy(request *model.DeleteGaussMySqlProxyRequest) (*model.DeleteGaussMySqlProxyResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlProxyResponse), nil
	}
}

// DeleteGaussMySqlProxyInvoker 关闭数据库代理
func (c *GaussDBClient) DeleteGaussMySqlProxyInvoker(request *model.DeleteGaussMySqlProxyRequest) *DeleteGaussMySqlProxyInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlProxy()
	return &DeleteGaussMySqlProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlReadonlyNode 删除/退订只读节点
//
// 删除/退订实例的只读节点。多可用区模式删除/退订只读节点时，要保证删除/退订后，剩余的只读节点和主节点在不同的可用区中，否则无法删除/退订该只读节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlReadonlyNode(request *model.DeleteGaussMySqlReadonlyNodeRequest) (*model.DeleteGaussMySqlReadonlyNodeResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlReadonlyNodeResponse), nil
	}
}

// DeleteGaussMySqlReadonlyNodeInvoker 删除/退订只读节点
func (c *GaussDBClient) DeleteGaussMySqlReadonlyNodeInvoker(request *model.DeleteGaussMySqlReadonlyNodeRequest) *DeleteGaussMySqlReadonlyNodeInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlReadonlyNode()
	return &DeleteGaussMySqlReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLtsConfigs 批量删除LTS日志配置
//
// 批量删除LTS日志配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteLtsConfigs(request *model.DeleteLtsConfigsRequest) (*model.DeleteLtsConfigsResponse, error) {
	requestDef := GenReqDefForDeleteLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLtsConfigsResponse), nil
	}
}

// DeleteLtsConfigsInvoker 批量删除LTS日志配置
func (c *GaussDBClient) DeleteLtsConfigsInvoker(request *model.DeleteLtsConfigsRequest) *DeleteLtsConfigsInvoker {
	requestDef := GenReqDefForDeleteLtsConfigs()
	return &DeleteLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduleTasK 删除定时任务
//
// 删除定时任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteScheduleTasK(request *model.DeleteScheduleTasKRequest) (*model.DeleteScheduleTasKResponse, error) {
	requestDef := GenReqDefForDeleteScheduleTasK()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduleTasKResponse), nil
	}
}

// DeleteScheduleTasKInvoker 删除定时任务
func (c *GaussDBClient) DeleteScheduleTasKInvoker(request *model.DeleteScheduleTasKRequest) *DeleteScheduleTasKInvoker {
	requestDef := GenReqDefForDeleteScheduleTasK()
	return &DeleteScheduleTasKInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTaskRecord 删除指定任务记录
//
// 删除指定任务记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteTaskRecord(request *model.DeleteTaskRecordRequest) (*model.DeleteTaskRecordResponse, error) {
	requestDef := GenReqDefForDeleteTaskRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskRecordResponse), nil
	}
}

// DeleteTaskRecordInvoker 删除指定任务记录
func (c *GaussDBClient) DeleteTaskRecordInvoker(request *model.DeleteTaskRecordRequest) *DeleteTaskRecordInvoker {
	requestDef := GenReqDefForDeleteTaskRecord()
	return &DeleteTaskRecordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DescribeBackupEncryptStatus 查询实例是否开启备份加密功能
//
// 查询实例是否开启备份加密功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DescribeBackupEncryptStatus(request *model.DescribeBackupEncryptStatusRequest) (*model.DescribeBackupEncryptStatusResponse, error) {
	requestDef := GenReqDefForDescribeBackupEncryptStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DescribeBackupEncryptStatusResponse), nil
	}
}

// DescribeBackupEncryptStatusInvoker 查询实例是否开启备份加密功能
func (c *GaussDBClient) DescribeBackupEncryptStatusInvoker(request *model.DescribeBackupEncryptStatusRequest) *DescribeBackupEncryptStatusInvoker {
	requestDef := GenReqDefForDescribeBackupEncryptStatus()
	return &DescribeBackupEncryptStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandGaussMySqlInstanceVolume 包周期存储扩容
//
// 包周期存储扩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ExpandGaussMySqlInstanceVolume(request *model.ExpandGaussMySqlInstanceVolumeRequest) (*model.ExpandGaussMySqlInstanceVolumeResponse, error) {
	requestDef := GenReqDefForExpandGaussMySqlInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandGaussMySqlInstanceVolumeResponse), nil
	}
}

// ExpandGaussMySqlInstanceVolumeInvoker 包周期存储扩容
func (c *GaussDBClient) ExpandGaussMySqlInstanceVolumeInvoker(request *model.ExpandGaussMySqlInstanceVolumeRequest) *ExpandGaussMySqlInstanceVolumeInvoker {
	requestDef := GenReqDefForExpandGaussMySqlInstanceVolume()
	return &ExpandGaussMySqlInstanceVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandGaussMySqlProxy 扩容数据库代理节点的数量
//
// 扩容数据库代理节点的数量。
// DeC专属云账号暂不支持数据库代理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ExpandGaussMySqlProxy(request *model.ExpandGaussMySqlProxyRequest) (*model.ExpandGaussMySqlProxyResponse, error) {
	requestDef := GenReqDefForExpandGaussMySqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandGaussMySqlProxyResponse), nil
	}
}

// ExpandGaussMySqlProxyInvoker 扩容数据库代理节点的数量
func (c *GaussDBClient) ExpandGaussMySqlProxyInvoker(request *model.ExpandGaussMySqlProxyRequest) *ExpandGaussMySqlProxyInvoker {
	requestDef := GenReqDefForExpandGaussMySqlProxy()
	return &ExpandGaussMySqlProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InvokeGaussMySqlInstanceSwitchOver 手动主备倒换
//
// 用户手动进行主备倒换。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) InvokeGaussMySqlInstanceSwitchOver(request *model.InvokeGaussMySqlInstanceSwitchOverRequest) (*model.InvokeGaussMySqlInstanceSwitchOverResponse, error) {
	requestDef := GenReqDefForInvokeGaussMySqlInstanceSwitchOver()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InvokeGaussMySqlInstanceSwitchOverResponse), nil
	}
}

// InvokeGaussMySqlInstanceSwitchOverInvoker 手动主备倒换
func (c *GaussDBClient) InvokeGaussMySqlInstanceSwitchOverInvoker(request *model.InvokeGaussMySqlInstanceSwitchOverRequest) *InvokeGaussMySqlInstanceSwitchOverInvoker {
	requestDef := GenReqDefForInvokeGaussMySqlInstanceSwitchOver()
	return &InvokeGaussMySqlInstanceSwitchOverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditLogDownloadLink 获取全量SQL的临时下载链接
//
// 获取全量SQL的临时下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListAuditLogDownloadLink(request *model.ListAuditLogDownloadLinkRequest) (*model.ListAuditLogDownloadLinkResponse, error) {
	requestDef := GenReqDefForListAuditLogDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditLogDownloadLinkResponse), nil
	}
}

// ListAuditLogDownloadLinkInvoker 获取全量SQL的临时下载链接
func (c *GaussDBClient) ListAuditLogDownloadLinkInvoker(request *model.ListAuditLogDownloadLinkRequest) *ListAuditLogDownloadLinkInvoker {
	requestDef := GenReqDefForListAuditLogDownloadLink()
	return &ListAuditLogDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationsDifferences 对比参数模板
//
// 比较两个参数模板之间的差异。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListConfigurationsDifferences(request *model.ListConfigurationsDifferencesRequest) (*model.ListConfigurationsDifferencesResponse, error) {
	requestDef := GenReqDefForListConfigurationsDifferences()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsDifferencesResponse), nil
	}
}

// ListConfigurationsDifferencesInvoker 对比参数模板
func (c *GaussDBClient) ListConfigurationsDifferencesInvoker(request *model.ListConfigurationsDifferencesRequest) *ListConfigurationsDifferencesInvoker {
	requestDef := GenReqDefForListConfigurationsDifferences()
	return &ListConfigurationsDifferencesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurationsInstances 查询可应用的实例列表
//
// 查询指定参数模板可被应用的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListConfigurationsInstances(request *model.ListConfigurationsInstancesRequest) (*model.ListConfigurationsInstancesResponse, error) {
	requestDef := GenReqDefForListConfigurationsInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsInstancesResponse), nil
	}
}

// ListConfigurationsInstancesInvoker 查询可应用的实例列表
func (c *GaussDBClient) ListConfigurationsInstancesInvoker(request *model.ListConfigurationsInstancesRequest) *ListConfigurationsInstancesInvoker {
	requestDef := GenReqDefForListConfigurationsInstances()
	return &ListConfigurationsInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnterpriseProjects 查询企业项目
//
// 查询企业项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListEnterpriseProjects(request *model.ListEnterpriseProjectsRequest) (*model.ListEnterpriseProjectsResponse, error) {
	requestDef := GenReqDefForListEnterpriseProjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseProjectsResponse), nil
	}
}

// ListEnterpriseProjectsInvoker 查询企业项目
func (c *GaussDBClient) ListEnterpriseProjectsInvoker(request *model.ListEnterpriseProjectsRequest) *ListEnterpriseProjectsInvoker {
	requestDef := GenReqDefForListEnterpriseProjects()
	return &ListEnterpriseProjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlConfigurations 查询参数模板
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlConfigurations(request *model.ListGaussMySqlConfigurationsRequest) (*model.ListGaussMySqlConfigurationsResponse, error) {
	requestDef := GenReqDefForListGaussMySqlConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlConfigurationsResponse), nil
	}
}

// ListGaussMySqlConfigurationsInvoker 查询参数模板
func (c *GaussDBClient) ListGaussMySqlConfigurationsInvoker(request *model.ListGaussMySqlConfigurationsRequest) *ListGaussMySqlConfigurationsInvoker {
	requestDef := GenReqDefForListGaussMySqlConfigurations()
	return &ListGaussMySqlConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlDatabase 查询数据库列表
//
// 查询 GaussDB(for MySQL)实例数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlDatabase(request *model.ListGaussMySqlDatabaseRequest) (*model.ListGaussMySqlDatabaseResponse, error) {
	requestDef := GenReqDefForListGaussMySqlDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlDatabaseResponse), nil
	}
}

// ListGaussMySqlDatabaseInvoker 查询数据库列表
func (c *GaussDBClient) ListGaussMySqlDatabaseInvoker(request *model.ListGaussMySqlDatabaseRequest) *ListGaussMySqlDatabaseInvoker {
	requestDef := GenReqDefForListGaussMySqlDatabase()
	return &ListGaussMySqlDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlDatabaseCharsets 查询数据库可用字符集
//
// 查询云数据库 GaussDB(for MySQL)实例数据库可用字符集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlDatabaseCharsets(request *model.ListGaussMySqlDatabaseCharsetsRequest) (*model.ListGaussMySqlDatabaseCharsetsResponse, error) {
	requestDef := GenReqDefForListGaussMySqlDatabaseCharsets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlDatabaseCharsetsResponse), nil
	}
}

// ListGaussMySqlDatabaseCharsetsInvoker 查询数据库可用字符集
func (c *GaussDBClient) ListGaussMySqlDatabaseCharsetsInvoker(request *model.ListGaussMySqlDatabaseCharsetsRequest) *ListGaussMySqlDatabaseCharsetsInvoker {
	requestDef := GenReqDefForListGaussMySqlDatabaseCharsets()
	return &ListGaussMySqlDatabaseCharsetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlDatabaseUser 查询数据库用户
//
// 查询云数据库 GaussDB(for MySQL)实例数据库用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlDatabaseUser(request *model.ListGaussMySqlDatabaseUserRequest) (*model.ListGaussMySqlDatabaseUserResponse, error) {
	requestDef := GenReqDefForListGaussMySqlDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlDatabaseUserResponse), nil
	}
}

// ListGaussMySqlDatabaseUserInvoker 查询数据库用户
func (c *GaussDBClient) ListGaussMySqlDatabaseUserInvoker(request *model.ListGaussMySqlDatabaseUserRequest) *ListGaussMySqlDatabaseUserInvoker {
	requestDef := GenReqDefForListGaussMySqlDatabaseUser()
	return &ListGaussMySqlDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlDedicatedResources 查询专属资源池列表
//
// 获取专属资源池列表，包括用户开通的所有专属资源池信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlDedicatedResources(request *model.ListGaussMySqlDedicatedResourcesRequest) (*model.ListGaussMySqlDedicatedResourcesResponse, error) {
	requestDef := GenReqDefForListGaussMySqlDedicatedResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlDedicatedResourcesResponse), nil
	}
}

// ListGaussMySqlDedicatedResourcesInvoker 查询专属资源池列表
func (c *GaussDBClient) ListGaussMySqlDedicatedResourcesInvoker(request *model.ListGaussMySqlDedicatedResourcesRequest) *ListGaussMySqlDedicatedResourcesInvoker {
	requestDef := GenReqDefForListGaussMySqlDedicatedResources()
	return &ListGaussMySqlDedicatedResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlInstanceDetailInfo 批量查询实例详情
//
// 批量查询实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlInstanceDetailInfo(request *model.ListGaussMySqlInstanceDetailInfoRequest) (*model.ListGaussMySqlInstanceDetailInfoResponse, error) {
	requestDef := GenReqDefForListGaussMySqlInstanceDetailInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlInstanceDetailInfoResponse), nil
	}
}

// ListGaussMySqlInstanceDetailInfoInvoker 批量查询实例详情
func (c *GaussDBClient) ListGaussMySqlInstanceDetailInfoInvoker(request *model.ListGaussMySqlInstanceDetailInfoRequest) *ListGaussMySqlInstanceDetailInfoInvoker {
	requestDef := GenReqDefForListGaussMySqlInstanceDetailInfo()
	return &ListGaussMySqlInstanceDetailInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlInstanceDetailInfoUnifyStatus 批量查询实例详情
//
// 批量查询实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlInstanceDetailInfoUnifyStatus(request *model.ListGaussMySqlInstanceDetailInfoUnifyStatusRequest) (*model.ListGaussMySqlInstanceDetailInfoUnifyStatusResponse, error) {
	requestDef := GenReqDefForListGaussMySqlInstanceDetailInfoUnifyStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlInstanceDetailInfoUnifyStatusResponse), nil
	}
}

// ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker 批量查询实例详情
func (c *GaussDBClient) ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker(request *model.ListGaussMySqlInstanceDetailInfoUnifyStatusRequest) *ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker {
	requestDef := GenReqDefForListGaussMySqlInstanceDetailInfoUnifyStatus()
	return &ListGaussMySqlInstanceDetailInfoUnifyStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlInstances 查询实例列表
//
// 根据指定条件查询实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlInstances(request *model.ListGaussMySqlInstancesRequest) (*model.ListGaussMySqlInstancesResponse, error) {
	requestDef := GenReqDefForListGaussMySqlInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlInstancesResponse), nil
	}
}

// ListGaussMySqlInstancesInvoker 查询实例列表
func (c *GaussDBClient) ListGaussMySqlInstancesInvoker(request *model.ListGaussMySqlInstancesRequest) *ListGaussMySqlInstancesInvoker {
	requestDef := GenReqDefForListGaussMySqlInstances()
	return &ListGaussMySqlInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlInstancesUnifyStatus 查询实例列表
//
// 根据指定条件查询实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlInstancesUnifyStatus(request *model.ListGaussMySqlInstancesUnifyStatusRequest) (*model.ListGaussMySqlInstancesUnifyStatusResponse, error) {
	requestDef := GenReqDefForListGaussMySqlInstancesUnifyStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlInstancesUnifyStatusResponse), nil
	}
}

// ListGaussMySqlInstancesUnifyStatusInvoker 查询实例列表
func (c *GaussDBClient) ListGaussMySqlInstancesUnifyStatusInvoker(request *model.ListGaussMySqlInstancesUnifyStatusRequest) *ListGaussMySqlInstancesUnifyStatusInvoker {
	requestDef := GenReqDefForListGaussMySqlInstancesUnifyStatus()
	return &ListGaussMySqlInstancesUnifyStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImmediateJobs 获取即时任务列表
//
// 获取即时任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListImmediateJobs(request *model.ListImmediateJobsRequest) (*model.ListImmediateJobsResponse, error) {
	requestDef := GenReqDefForListImmediateJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImmediateJobsResponse), nil
	}
}

// ListImmediateJobsInvoker 获取即时任务列表
func (c *GaussDBClient) ListImmediateJobsInvoker(request *model.ListImmediateJobsRequest) *ListImmediateJobsInvoker {
	requestDef := GenReqDefForListImmediateJobs()
	return &ListImmediateJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceConfigurations 获取指定实例的参数信息
//
// 获取指定实例的参数信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListInstanceConfigurations(request *model.ListInstanceConfigurationsRequest) (*model.ListInstanceConfigurationsResponse, error) {
	requestDef := GenReqDefForListInstanceConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceConfigurationsResponse), nil
	}
}

// ListInstanceConfigurationsInvoker 获取指定实例的参数信息
func (c *GaussDBClient) ListInstanceConfigurationsInvoker(request *model.ListInstanceConfigurationsRequest) *ListInstanceConfigurationsInvoker {
	requestDef := GenReqDefForListInstanceConfigurations()
	return &ListInstanceConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询资源标签
func (c *GaussDBClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsErrorLogDetails 获取错误日志详情列表
//
// 获取指定实例的错误日志详情列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListLtsErrorLogDetails(request *model.ListLtsErrorLogDetailsRequest) (*model.ListLtsErrorLogDetailsResponse, error) {
	requestDef := GenReqDefForListLtsErrorLogDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsErrorLogDetailsResponse), nil
	}
}

// ListLtsErrorLogDetailsInvoker 获取错误日志详情列表
func (c *GaussDBClient) ListLtsErrorLogDetailsInvoker(request *model.ListLtsErrorLogDetailsRequest) *ListLtsErrorLogDetailsInvoker {
	requestDef := GenReqDefForListLtsErrorLogDetails()
	return &ListLtsErrorLogDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsSlowlogDetails 获取慢日志详情列表
//
// 获取指定实例的慢日志详情列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListLtsSlowlogDetails(request *model.ListLtsSlowlogDetailsRequest) (*model.ListLtsSlowlogDetailsResponse, error) {
	requestDef := GenReqDefForListLtsSlowlogDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsSlowlogDetailsResponse), nil
	}
}

// ListLtsSlowlogDetailsInvoker 获取慢日志详情列表
func (c *GaussDBClient) ListLtsSlowlogDetailsInvoker(request *model.ListLtsSlowlogDetailsRequest) *ListLtsSlowlogDetailsInvoker {
	requestDef := GenReqDefForListLtsSlowlogDetails()
	return &ListLtsSlowlogDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListModifyHistory 查询参数修改历史
//
// 查询参数修改历史。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListModifyHistory(request *model.ListModifyHistoryRequest) (*model.ListModifyHistoryResponse, error) {
	requestDef := GenReqDefForListModifyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListModifyHistoryResponse), nil
	}
}

// ListModifyHistoryInvoker 查询参数修改历史
func (c *GaussDBClient) ListModifyHistoryInvoker(request *model.ListModifyHistoryRequest) *ListModifyHistoryInvoker {
	requestDef := GenReqDefForListModifyHistory()
	return &ListModifyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListParamsTemplateApplyHistory 查询参数模板应用记录。
//
// 查询参数模板应用记录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListParamsTemplateApplyHistory(request *model.ListParamsTemplateApplyHistoryRequest) (*model.ListParamsTemplateApplyHistoryResponse, error) {
	requestDef := GenReqDefForListParamsTemplateApplyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListParamsTemplateApplyHistoryResponse), nil
	}
}

// ListParamsTemplateApplyHistoryInvoker 查询参数模板应用记录。
func (c *GaussDBClient) ListParamsTemplateApplyHistoryInvoker(request *model.ListParamsTemplateApplyHistoryRequest) *ListParamsTemplateApplyHistoryInvoker {
	requestDef := GenReqDefForListParamsTemplateApplyHistory()
	return &ListParamsTemplateApplyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询指定project ID下实例的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *GaussDBClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRecycleInstances 查询回收站实例信息
//
// 查询回收站实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

// ListRecycleInstancesInvoker 查询回收站实例信息
func (c *GaussDBClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduleJobs 获取定时任务列表
//
// 获取定时任务列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListScheduleJobs(request *model.ListScheduleJobsRequest) (*model.ListScheduleJobsResponse, error) {
	requestDef := GenReqDefForListScheduleJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduleJobsResponse), nil
	}
}

// ListScheduleJobsInvoker 获取定时任务列表
func (c *GaussDBClient) ListScheduleJobsInvoker(request *model.ListScheduleJobsRequest) *ListScheduleJobsInvoker {
	requestDef := GenReqDefForListScheduleJobs()
	return &ListScheduleJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyBackupEncryptStatus 打开或关闭备份加密
//
// 打开或关闭备份加密。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ModifyBackupEncryptStatus(request *model.ModifyBackupEncryptStatusRequest) (*model.ModifyBackupEncryptStatusResponse, error) {
	requestDef := GenReqDefForModifyBackupEncryptStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyBackupEncryptStatusResponse), nil
	}
}

// ModifyBackupEncryptStatusInvoker 打开或关闭备份加密
func (c *GaussDBClient) ModifyBackupEncryptStatusInvoker(request *model.ModifyBackupEncryptStatusRequest) *ModifyBackupEncryptStatusInvoker {
	requestDef := GenReqDefForModifyBackupEncryptStatus()
	return &ModifyBackupEncryptStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyGaussMySqlProxyRouteMode 设置读写分离路由模式
//
// 设置读写分离路由模式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ModifyGaussMySqlProxyRouteMode(request *model.ModifyGaussMySqlProxyRouteModeRequest) (*model.ModifyGaussMySqlProxyRouteModeResponse, error) {
	requestDef := GenReqDefForModifyGaussMySqlProxyRouteMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyGaussMySqlProxyRouteModeResponse), nil
	}
}

// ModifyGaussMySqlProxyRouteModeInvoker 设置读写分离路由模式
func (c *GaussDBClient) ModifyGaussMySqlProxyRouteModeInvoker(request *model.ModifyGaussMySqlProxyRouteModeRequest) *ModifyGaussMySqlProxyRouteModeInvoker {
	requestDef := GenReqDefForModifyGaussMySqlProxyRouteMode()
	return &ModifyGaussMySqlProxyRouteModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyGaussMysqlDns 修改内网域名
//
// 修改内网域名。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ModifyGaussMysqlDns(request *model.ModifyGaussMysqlDnsRequest) (*model.ModifyGaussMysqlDnsResponse, error) {
	requestDef := GenReqDefForModifyGaussMysqlDns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyGaussMysqlDnsResponse), nil
	}
}

// ModifyGaussMysqlDnsInvoker 修改内网域名
func (c *GaussDBClient) ModifyGaussMysqlDnsInvoker(request *model.ModifyGaussMysqlDnsRequest) *ModifyGaussMysqlDnsInvoker {
	requestDef := GenReqDefForModifyGaussMysqlDns()
	return &ModifyGaussMysqlDnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyNodePriority 修改节点故障倒换优先级。
//
// 修改节点故障倒换优先级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ModifyNodePriority(request *model.ModifyNodePriorityRequest) (*model.ModifyNodePriorityResponse, error) {
	requestDef := GenReqDefForModifyNodePriority()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyNodePriorityResponse), nil
	}
}

// ModifyNodePriorityInvoker 修改节点故障倒换优先级。
func (c *GaussDBClient) ModifyNodePriorityInvoker(request *model.ModifyNodePriorityRequest) *ModifyNodePriorityInvoker {
	requestDef := GenReqDefForModifyNodePriority()
	return &ModifyNodePriorityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenameInstanceNode 批量修改节点名称.
//
// 批量修改节点名称.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RenameInstanceNode(request *model.RenameInstanceNodeRequest) (*model.RenameInstanceNodeResponse, error) {
	requestDef := GenReqDefForRenameInstanceNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenameInstanceNodeResponse), nil
	}
}

// RenameInstanceNodeInvoker 批量修改节点名称.
func (c *GaussDBClient) RenameInstanceNodeInvoker(request *model.RenameInstanceNodeRequest) *RenameInstanceNodeInvoker {
	requestDef := GenReqDefForRenameInstanceNode()
	return &RenameInstanceNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetGaussMySqlDatabasePassword 修改数据库用户密码
//
// 修改云数据库 GaussDB(for MySQL)实例数据库用户密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResetGaussMySqlDatabasePassword(request *model.ResetGaussMySqlDatabasePasswordRequest) (*model.ResetGaussMySqlDatabasePasswordResponse, error) {
	requestDef := GenReqDefForResetGaussMySqlDatabasePassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetGaussMySqlDatabasePasswordResponse), nil
	}
}

// ResetGaussMySqlDatabasePasswordInvoker 修改数据库用户密码
func (c *GaussDBClient) ResetGaussMySqlDatabasePasswordInvoker(request *model.ResetGaussMySqlDatabasePasswordRequest) *ResetGaussMySqlDatabasePasswordInvoker {
	requestDef := GenReqDefForResetGaussMySqlDatabasePassword()
	return &ResetGaussMySqlDatabasePasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetGaussMySqlPassword 重置数据库密码
//
// 重置数据库密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResetGaussMySqlPassword(request *model.ResetGaussMySqlPasswordRequest) (*model.ResetGaussMySqlPasswordResponse, error) {
	requestDef := GenReqDefForResetGaussMySqlPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetGaussMySqlPasswordResponse), nil
	}
}

// ResetGaussMySqlPasswordInvoker 重置数据库密码
func (c *GaussDBClient) ResetGaussMySqlPasswordInvoker(request *model.ResetGaussMySqlPasswordRequest) *ResetGaussMySqlPasswordInvoker {
	requestDef := GenReqDefForResetGaussMySqlPassword()
	return &ResetGaussMySqlPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartGaussMySqlInstance 重启数据库实例
//
// 重启数据库实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestartGaussMySqlInstance(request *model.RestartGaussMySqlInstanceRequest) (*model.RestartGaussMySqlInstanceResponse, error) {
	requestDef := GenReqDefForRestartGaussMySqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartGaussMySqlInstanceResponse), nil
	}
}

// RestartGaussMySqlInstanceInvoker 重启数据库实例
func (c *GaussDBClient) RestartGaussMySqlInstanceInvoker(request *model.RestartGaussMySqlInstanceRequest) *RestartGaussMySqlInstanceInvoker {
	requestDef := GenReqDefForRestartGaussMySqlInstance()
	return &RestartGaussMySqlInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartGaussMySqlNode 节点重启
//
// 节点重启。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestartGaussMySqlNode(request *model.RestartGaussMySqlNodeRequest) (*model.RestartGaussMySqlNodeResponse, error) {
	requestDef := GenReqDefForRestartGaussMySqlNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartGaussMySqlNodeResponse), nil
	}
}

// RestartGaussMySqlNodeInvoker 节点重启
func (c *GaussDBClient) RestartGaussMySqlNodeInvoker(request *model.RestartGaussMySqlNodeRequest) *RestartGaussMySqlNodeInvoker {
	requestDef := GenReqDefForRestartGaussMySqlNode()
	return &RestartGaussMySqlNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartProxyInstance 重启数据库代理.
//
// 重启数据库代理.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestartProxyInstance(request *model.RestartProxyInstanceRequest) (*model.RestartProxyInstanceResponse, error) {
	requestDef := GenReqDefForRestartProxyInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartProxyInstanceResponse), nil
	}
}

// RestartProxyInstanceInvoker 重启数据库代理.
func (c *GaussDBClient) RestartProxyInstanceInvoker(request *model.RestartProxyInstanceRequest) *RestartProxyInstanceInvoker {
	requestDef := GenReqDefForRestartProxyInstance()
	return &RestartProxyInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreOldInstance 备份恢复到当前实例或已有实例
//
// 备份恢复到当前实例或已有实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestoreOldInstance(request *model.RestoreOldInstanceRequest) (*model.RestoreOldInstanceResponse, error) {
	requestDef := GenReqDefForRestoreOldInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreOldInstanceResponse), nil
	}
}

// RestoreOldInstanceInvoker 备份恢复到当前实例或已有实例
func (c *GaussDBClient) RestoreOldInstanceInvoker(request *model.RestoreOldInstanceRequest) *RestoreOldInstanceInvoker {
	requestDef := GenReqDefForRestoreOldInstance()
	return &RestoreOldInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetGaussMySqlProxyWeight 设置读写分离权重
//
// 设置读写分离权重。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SetGaussMySqlProxyWeight(request *model.SetGaussMySqlProxyWeightRequest) (*model.SetGaussMySqlProxyWeightResponse, error) {
	requestDef := GenReqDefForSetGaussMySqlProxyWeight()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetGaussMySqlProxyWeightResponse), nil
	}
}

// SetGaussMySqlProxyWeightInvoker 设置读写分离权重
func (c *GaussDBClient) SetGaussMySqlProxyWeightInvoker(request *model.SetGaussMySqlProxyWeightRequest) *SetGaussMySqlProxyWeightInvoker {
	requestDef := GenReqDefForSetGaussMySqlProxyWeight()
	return &SetGaussMySqlProxyWeightInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetGaussMySqlQuotas 设置租户基于企业项目的资源配额
//
// 设置指定企业项目的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SetGaussMySqlQuotas(request *model.SetGaussMySqlQuotasRequest) (*model.SetGaussMySqlQuotasResponse, error) {
	requestDef := GenReqDefForSetGaussMySqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetGaussMySqlQuotasResponse), nil
	}
}

// SetGaussMySqlQuotasInvoker 设置租户基于企业项目的资源配额
func (c *GaussDBClient) SetGaussMySqlQuotasInvoker(request *model.SetGaussMySqlQuotasRequest) *SetGaussMySqlQuotasInvoker {
	requestDef := GenReqDefForSetGaussMySqlQuotas()
	return &SetGaussMySqlQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetRecyclePolicy 设置回收站策略
//
// 设置回收站策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SetRecyclePolicy(request *model.SetRecyclePolicyRequest) (*model.SetRecyclePolicyResponse, error) {
	requestDef := GenReqDefForSetRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecyclePolicyResponse), nil
	}
}

// SetRecyclePolicyInvoker 设置回收站策略
func (c *GaussDBClient) SetRecyclePolicyInvoker(request *model.SetRecyclePolicyRequest) *SetRecyclePolicyInvoker {
	requestDef := GenReqDefForSetRecyclePolicy()
	return &SetRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditLog 查询全量SQL开关状态
//
// 查询全量SQL开关状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowAuditLog(request *model.ShowAuditLogRequest) (*model.ShowAuditLogResponse, error) {
	requestDef := GenReqDefForShowAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditLogResponse), nil
	}
}

// ShowAuditLogInvoker 查询全量SQL开关状态
func (c *GaussDBClient) ShowAuditLogInvoker(request *model.ShowAuditLogRequest) *ShowAuditLogInvoker {
	requestDef := GenReqDefForShowAuditLog()
	return &ShowAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoScalingHistory 查询自动变配历史记录.
//
// 查询自动变配历史记录.
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowAutoScalingHistory(request *model.ShowAutoScalingHistoryRequest) (*model.ShowAutoScalingHistoryResponse, error) {
	requestDef := GenReqDefForShowAutoScalingHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoScalingHistoryResponse), nil
	}
}

// ShowAutoScalingHistoryInvoker 查询自动变配历史记录.
func (c *GaussDBClient) ShowAutoScalingHistoryInvoker(request *model.ShowAutoScalingHistoryRequest) *ShowAutoScalingHistoryInvoker {
	requestDef := GenReqDefForShowAutoScalingHistory()
	return &ShowAutoScalingHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoScalingPolicy 查询自动变配
//
// 查询自动变配。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowAutoScalingPolicy(request *model.ShowAutoScalingPolicyRequest) (*model.ShowAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForShowAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoScalingPolicyResponse), nil
	}
}

// ShowAutoScalingPolicyInvoker 查询自动变配
func (c *GaussDBClient) ShowAutoScalingPolicyInvoker(request *model.ShowAutoScalingPolicyRequest) *ShowAutoScalingPolicyInvoker {
	requestDef := GenReqDefForShowAutoScalingPolicy()
	return &ShowAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupRestoreTime 查询可恢复时间段
//
// 查询实例的可恢复时间段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowBackupRestoreTime(request *model.ShowBackupRestoreTimeRequest) (*model.ShowBackupRestoreTimeResponse, error) {
	requestDef := GenReqDefForShowBackupRestoreTime()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupRestoreTimeResponse), nil
	}
}

// ShowBackupRestoreTimeInvoker 查询可恢复时间段
func (c *GaussDBClient) ShowBackupRestoreTimeInvoker(request *model.ShowBackupRestoreTimeRequest) *ShowBackupRestoreTimeInvoker {
	requestDef := GenReqDefForShowBackupRestoreTime()
	return &ShowBackupRestoreTimeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDedicatedResourceInfo 查询专属资源信息详情
//
// 查询专属资源信息详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowDedicatedResourceInfo(request *model.ShowDedicatedResourceInfoRequest) (*model.ShowDedicatedResourceInfoResponse, error) {
	requestDef := GenReqDefForShowDedicatedResourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDedicatedResourceInfoResponse), nil
	}
}

// ShowDedicatedResourceInfoInvoker 查询专属资源信息详情
func (c *GaussDBClient) ShowDedicatedResourceInfoInvoker(request *model.ShowDedicatedResourceInfoRequest) *ShowDedicatedResourceInfoInvoker {
	requestDef := GenReqDefForShowDedicatedResourceInfo()
	return &ShowDedicatedResourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlBackupList 查询全量备份列表
//
// 查询全量备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlBackupList(request *model.ShowGaussMySqlBackupListRequest) (*model.ShowGaussMySqlBackupListResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlBackupList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlBackupListResponse), nil
	}
}

// ShowGaussMySqlBackupListInvoker 查询全量备份列表
func (c *GaussDBClient) ShowGaussMySqlBackupListInvoker(request *model.ShowGaussMySqlBackupListRequest) *ShowGaussMySqlBackupListInvoker {
	requestDef := GenReqDefForShowGaussMySqlBackupList()
	return &ShowGaussMySqlBackupListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlBackupPolicy(request *model.ShowGaussMySqlBackupPolicyRequest) (*model.ShowGaussMySqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlBackupPolicyResponse), nil
	}
}

// ShowGaussMySqlBackupPolicyInvoker 查询自动备份策略
func (c *GaussDBClient) ShowGaussMySqlBackupPolicyInvoker(request *model.ShowGaussMySqlBackupPolicyRequest) *ShowGaussMySqlBackupPolicyInvoker {
	requestDef := GenReqDefForShowGaussMySqlBackupPolicy()
	return &ShowGaussMySqlBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlConfiguration 获取参数模板详情
//
// 获取指定参数模板的参数信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlConfiguration(request *model.ShowGaussMySqlConfigurationRequest) (*model.ShowGaussMySqlConfigurationResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlConfigurationResponse), nil
	}
}

// ShowGaussMySqlConfigurationInvoker 获取参数模板详情
func (c *GaussDBClient) ShowGaussMySqlConfigurationInvoker(request *model.ShowGaussMySqlConfigurationRequest) *ShowGaussMySqlConfigurationInvoker {
	requestDef := GenReqDefForShowGaussMySqlConfiguration()
	return &ShowGaussMySqlConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlEngineVersion 查询数据库引擎的版本
//
// 获取指定数据库引擎对应的数据库版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlEngineVersion(request *model.ShowGaussMySqlEngineVersionRequest) (*model.ShowGaussMySqlEngineVersionResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlEngineVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlEngineVersionResponse), nil
	}
}

// ShowGaussMySqlEngineVersionInvoker 查询数据库引擎的版本
func (c *GaussDBClient) ShowGaussMySqlEngineVersionInvoker(request *model.ShowGaussMySqlEngineVersionRequest) *ShowGaussMySqlEngineVersionInvoker {
	requestDef := GenReqDefForShowGaussMySqlEngineVersion()
	return &ShowGaussMySqlEngineVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlFlavors 查询数据库规格
//
// 获取指定数据库引擎版本对应的规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlFlavors(request *model.ShowGaussMySqlFlavorsRequest) (*model.ShowGaussMySqlFlavorsResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlFlavorsResponse), nil
	}
}

// ShowGaussMySqlFlavorsInvoker 查询数据库规格
func (c *GaussDBClient) ShowGaussMySqlFlavorsInvoker(request *model.ShowGaussMySqlFlavorsRequest) *ShowGaussMySqlFlavorsInvoker {
	requestDef := GenReqDefForShowGaussMySqlFlavors()
	return &ShowGaussMySqlFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlIncrementalBackupList 查询增量备份列表
//
// 查询增量备份列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlIncrementalBackupList(request *model.ShowGaussMySqlIncrementalBackupListRequest) (*model.ShowGaussMySqlIncrementalBackupListResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlIncrementalBackupList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlIncrementalBackupListResponse), nil
	}
}

// ShowGaussMySqlIncrementalBackupListInvoker 查询增量备份列表
func (c *GaussDBClient) ShowGaussMySqlIncrementalBackupListInvoker(request *model.ShowGaussMySqlIncrementalBackupListRequest) *ShowGaussMySqlIncrementalBackupListInvoker {
	requestDef := GenReqDefForShowGaussMySqlIncrementalBackupList()
	return &ShowGaussMySqlIncrementalBackupListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlInstanceInfo 查询实例详情信息
//
// 查询实例详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlInstanceInfo(request *model.ShowGaussMySqlInstanceInfoRequest) (*model.ShowGaussMySqlInstanceInfoResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlInstanceInfoResponse), nil
	}
}

// ShowGaussMySqlInstanceInfoInvoker 查询实例详情信息
func (c *GaussDBClient) ShowGaussMySqlInstanceInfoInvoker(request *model.ShowGaussMySqlInstanceInfoRequest) *ShowGaussMySqlInstanceInfoInvoker {
	requestDef := GenReqDefForShowGaussMySqlInstanceInfo()
	return &ShowGaussMySqlInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlInstanceInfoUnifyStatus 查询实例详情信息
//
// 查询实例详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlInstanceInfoUnifyStatus(request *model.ShowGaussMySqlInstanceInfoUnifyStatusRequest) (*model.ShowGaussMySqlInstanceInfoUnifyStatusResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlInstanceInfoUnifyStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlInstanceInfoUnifyStatusResponse), nil
	}
}

// ShowGaussMySqlInstanceInfoUnifyStatusInvoker 查询实例详情信息
func (c *GaussDBClient) ShowGaussMySqlInstanceInfoUnifyStatusInvoker(request *model.ShowGaussMySqlInstanceInfoUnifyStatusRequest) *ShowGaussMySqlInstanceInfoUnifyStatusInvoker {
	requestDef := GenReqDefForShowGaussMySqlInstanceInfoUnifyStatus()
	return &ShowGaussMySqlInstanceInfoUnifyStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlJobInfo 获取指定ID的任务信息
//
// 获取GaussDB(for MySQL)任务中心指定ID的任务信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlJobInfo(request *model.ShowGaussMySqlJobInfoRequest) (*model.ShowGaussMySqlJobInfoResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlJobInfoResponse), nil
	}
}

// ShowGaussMySqlJobInfoInvoker 获取指定ID的任务信息
func (c *GaussDBClient) ShowGaussMySqlJobInfoInvoker(request *model.ShowGaussMySqlJobInfoRequest) *ShowGaussMySqlJobInfoInvoker {
	requestDef := GenReqDefForShowGaussMySqlJobInfo()
	return &ShowGaussMySqlJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlProjectQuotas 查询租户的实例配额
//
// 获取指定租户的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlProjectQuotas(request *model.ShowGaussMySqlProjectQuotasRequest) (*model.ShowGaussMySqlProjectQuotasResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlProjectQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlProjectQuotasResponse), nil
	}
}

// ShowGaussMySqlProjectQuotasInvoker 查询租户的实例配额
func (c *GaussDBClient) ShowGaussMySqlProjectQuotasInvoker(request *model.ShowGaussMySqlProjectQuotasRequest) *ShowGaussMySqlProjectQuotasInvoker {
	requestDef := GenReqDefForShowGaussMySqlProjectQuotas()
	return &ShowGaussMySqlProjectQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlProxyFlavors 查询数据库代理规格信息
//
// 查询数据库代理规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlProxyFlavors(request *model.ShowGaussMySqlProxyFlavorsRequest) (*model.ShowGaussMySqlProxyFlavorsResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlProxyFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlProxyFlavorsResponse), nil
	}
}

// ShowGaussMySqlProxyFlavorsInvoker 查询数据库代理规格信息
func (c *GaussDBClient) ShowGaussMySqlProxyFlavorsInvoker(request *model.ShowGaussMySqlProxyFlavorsRequest) *ShowGaussMySqlProxyFlavorsInvoker {
	requestDef := GenReqDefForShowGaussMySqlProxyFlavors()
	return &ShowGaussMySqlProxyFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlProxyList 查询数据库代理信息列表
//
// 查询数据库代理信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlProxyList(request *model.ShowGaussMySqlProxyListRequest) (*model.ShowGaussMySqlProxyListResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlProxyList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlProxyListResponse), nil
	}
}

// ShowGaussMySqlProxyListInvoker 查询数据库代理信息列表
func (c *GaussDBClient) ShowGaussMySqlProxyListInvoker(request *model.ShowGaussMySqlProxyListRequest) *ShowGaussMySqlProxyListInvoker {
	requestDef := GenReqDefForShowGaussMySqlProxyList()
	return &ShowGaussMySqlProxyListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlQuotas 查询租户基于企业项目的资源配额
//
// 获取指定企业项目的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlQuotas(request *model.ShowGaussMySqlQuotasRequest) (*model.ShowGaussMySqlQuotasResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlQuotasResponse), nil
	}
}

// ShowGaussMySqlQuotasInvoker 查询租户基于企业项目的资源配额
func (c *GaussDBClient) ShowGaussMySqlQuotasInvoker(request *model.ShowGaussMySqlQuotasRequest) *ShowGaussMySqlQuotasInvoker {
	requestDef := GenReqDefForShowGaussMySqlQuotas()
	return &ShowGaussMySqlQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceDatabaseVersion 查询内核版本信息
//
// 查询内核版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowInstanceDatabaseVersion(request *model.ShowInstanceDatabaseVersionRequest) (*model.ShowInstanceDatabaseVersionResponse, error) {
	requestDef := GenReqDefForShowInstanceDatabaseVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceDatabaseVersionResponse), nil
	}
}

// ShowInstanceDatabaseVersionInvoker 查询内核版本信息
func (c *GaussDBClient) ShowInstanceDatabaseVersionInvoker(request *model.ShowInstanceDatabaseVersionRequest) *ShowInstanceDatabaseVersionInvoker {
	requestDef := GenReqDefForShowInstanceDatabaseVersion()
	return &ShowInstanceDatabaseVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceEip 查询弹性公网IP。
//
// 查询弹性公网IP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowInstanceEip(request *model.ShowInstanceEipRequest) (*model.ShowInstanceEipResponse, error) {
	requestDef := GenReqDefForShowInstanceEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceEipResponse), nil
	}
}

// ShowInstanceEipInvoker 查询弹性公网IP。
func (c *GaussDBClient) ShowInstanceEipInvoker(request *model.ShowInstanceEipRequest) *ShowInstanceEipInvoker {
	requestDef := GenReqDefForShowInstanceEip()
	return &ShowInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceMonitorExtend 查询实例秒级监控
//
// 查询实例秒级监控信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowInstanceMonitorExtend(request *model.ShowInstanceMonitorExtendRequest) (*model.ShowInstanceMonitorExtendResponse, error) {
	requestDef := GenReqDefForShowInstanceMonitorExtend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMonitorExtendResponse), nil
	}
}

// ShowInstanceMonitorExtendInvoker 查询实例秒级监控
func (c *GaussDBClient) ShowInstanceMonitorExtendInvoker(request *model.ShowInstanceMonitorExtendRequest) *ShowInstanceMonitorExtendInvoker {
	requestDef := GenReqDefForShowInstanceMonitorExtend()
	return &ShowInstanceMonitorExtendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIntelligentDiagnosisAbnormalCountOfInstances 获取各指标的异常实例数
//
// 获取各指标的异常实例数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowIntelligentDiagnosisAbnormalCountOfInstances(request *model.ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest) (*model.ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse, error) {
	requestDef := GenReqDefForShowIntelligentDiagnosisAbnormalCountOfInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIntelligentDiagnosisAbnormalCountOfInstancesResponse), nil
	}
}

// ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker 获取各指标的异常实例数
func (c *GaussDBClient) ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker(request *model.ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest) *ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker {
	requestDef := GenReqDefForShowIntelligentDiagnosisAbnormalCountOfInstances()
	return &ShowIntelligentDiagnosisAbnormalCountOfInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIntelligentDiagnosisInstanceInfosPerMetric 获取某个指标的异常实例信息
//
// 获取某个指标的异常实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowIntelligentDiagnosisInstanceInfosPerMetric(request *model.ShowIntelligentDiagnosisInstanceInfosPerMetricRequest) (*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse, error) {
	requestDef := GenReqDefForShowIntelligentDiagnosisInstanceInfosPerMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIntelligentDiagnosisInstanceInfosPerMetricResponse), nil
	}
}

// ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker 获取某个指标的异常实例信息
func (c *GaussDBClient) ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker(request *model.ShowIntelligentDiagnosisInstanceInfosPerMetricRequest) *ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker {
	requestDef := GenReqDefForShowIntelligentDiagnosisInstanceInfosPerMetric()
	return &ShowIntelligentDiagnosisInstanceInfosPerMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLtsConfigs 查询实例LTS日志配置列表
//
// 查询实例LTS日志配置列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowLtsConfigs(request *model.ShowLtsConfigsRequest) (*model.ShowLtsConfigsResponse, error) {
	requestDef := GenReqDefForShowLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLtsConfigsResponse), nil
	}
}

// ShowLtsConfigsInvoker 查询实例LTS日志配置列表
func (c *GaussDBClient) ShowLtsConfigsInvoker(request *model.ShowLtsConfigsRequest) *ShowLtsConfigsInvoker {
	requestDef := GenReqDefForShowLtsConfigs()
	return &ShowLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProxyConfigurations 查询数据库代理内核参数。
//
// 查询数据库代理内核参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowProxyConfigurations(request *model.ShowProxyConfigurationsRequest) (*model.ShowProxyConfigurationsResponse, error) {
	requestDef := GenReqDefForShowProxyConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProxyConfigurationsResponse), nil
	}
}

// ShowProxyConfigurationsInvoker 查询数据库代理内核参数。
func (c *GaussDBClient) ShowProxyConfigurationsInvoker(request *model.ShowProxyConfigurationsRequest) *ShowProxyConfigurationsInvoker {
	requestDef := GenReqDefForShowProxyConfigurations()
	return &ShowProxyConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProxyIpgroup 查询代理实例访问控制
//
// 查询代理实例访问控制
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowProxyIpgroup(request *model.ShowProxyIpgroupRequest) (*model.ShowProxyIpgroupResponse, error) {
	requestDef := GenReqDefForShowProxyIpgroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProxyIpgroupResponse), nil
	}
}

// ShowProxyIpgroupInvoker 查询代理实例访问控制
func (c *GaussDBClient) ShowProxyIpgroupInvoker(request *model.ShowProxyIpgroupRequest) *ShowProxyIpgroupInvoker {
	requestDef := GenReqDefForShowProxyIpgroup()
	return &ShowProxyIpgroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowProxyVersion 查询代理实例小版本
//
// 查询代理实例小版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowProxyVersion(request *model.ShowProxyVersionRequest) (*model.ShowProxyVersionResponse, error) {
	requestDef := GenReqDefForShowProxyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowProxyVersionResponse), nil
	}
}

// ShowProxyVersionInvoker 查询代理实例小版本
func (c *GaussDBClient) ShowProxyVersionInvoker(request *model.ShowProxyVersionRequest) *ShowProxyVersionInvoker {
	requestDef := GenReqDefForShowProxyVersion()
	return &ShowProxyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRecyclePolicy 查询回收站策略
//
// 查询回收站策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

// ShowRecyclePolicyInvoker 查询回收站策略
func (c *GaussDBClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRestoreTables 查询表级时间点恢复可选表
//
// 查询表级时间点恢复可选表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowRestoreTables(request *model.ShowRestoreTablesRequest) (*model.ShowRestoreTablesResponse, error) {
	requestDef := GenReqDefForShowRestoreTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRestoreTablesResponse), nil
	}
}

// ShowRestoreTablesInvoker 查询表级时间点恢复可选表
func (c *GaussDBClient) ShowRestoreTablesInvoker(request *model.ShowRestoreTablesRequest) *ShowRestoreTablesInvoker {
	requestDef := GenReqDefForShowRestoreTables()
	return &ShowRestoreTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSlowlogSensitiveStatus 查询慢日志脱敏状态
//
// 查询慢日志脱敏状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowSlowlogSensitiveStatus(request *model.ShowSlowlogSensitiveStatusRequest) (*model.ShowSlowlogSensitiveStatusResponse, error) {
	requestDef := GenReqDefForShowSlowlogSensitiveStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowlogSensitiveStatusResponse), nil
	}
}

// ShowSlowlogSensitiveStatusInvoker 查询慢日志脱敏状态
func (c *GaussDBClient) ShowSlowlogSensitiveStatusInvoker(request *model.ShowSlowlogSensitiveStatusRequest) *ShowSlowlogSensitiveStatusInvoker {
	requestDef := GenReqDefForShowSlowlogSensitiveStatus()
	return &ShowSlowlogSensitiveStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkGaussMySqlProxy 减少数据库代理节点的数量
//
// 缩容数据库代理节点的数量。
// DeC专属云账号暂不支持数据库代理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShrinkGaussMySqlProxy(request *model.ShrinkGaussMySqlProxyRequest) (*model.ShrinkGaussMySqlProxyResponse, error) {
	requestDef := GenReqDefForShrinkGaussMySqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkGaussMySqlProxyResponse), nil
	}
}

// ShrinkGaussMySqlProxyInvoker 减少数据库代理节点的数量
func (c *GaussDBClient) ShrinkGaussMySqlProxyInvoker(request *model.ShrinkGaussMySqlProxyRequest) *ShrinkGaussMySqlProxyInvoker {
	requestDef := GenReqDefForShrinkGaussMySqlProxy()
	return &ShrinkGaussMySqlProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchAccessControl 开启或关闭访问控制
//
// 开启或关闭访问控制。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SwitchAccessControl(request *model.SwitchAccessControlRequest) (*model.SwitchAccessControlResponse, error) {
	requestDef := GenReqDefForSwitchAccessControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchAccessControlResponse), nil
	}
}

// SwitchAccessControlInvoker 开启或关闭访问控制
func (c *GaussDBClient) SwitchAccessControlInvoker(request *model.SwitchAccessControlRequest) *SwitchAccessControlInvoker {
	requestDef := GenReqDefForSwitchAccessControl()
	return &SwitchAccessControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchGaussMySqlConfiguration 应用参数模板
//
// 指定实例变更参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SwitchGaussMySqlConfiguration(request *model.SwitchGaussMySqlConfigurationRequest) (*model.SwitchGaussMySqlConfigurationResponse, error) {
	requestDef := GenReqDefForSwitchGaussMySqlConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchGaussMySqlConfigurationResponse), nil
	}
}

// SwitchGaussMySqlConfigurationInvoker 应用参数模板
func (c *GaussDBClient) SwitchGaussMySqlConfigurationInvoker(request *model.SwitchGaussMySqlConfigurationRequest) *SwitchGaussMySqlConfigurationInvoker {
	requestDef := GenReqDefForSwitchGaussMySqlConfiguration()
	return &SwitchGaussMySqlConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchGaussMySqlInstanceSsl 开关SSL
//
// 为实例设置SSL数据加密。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SwitchGaussMySqlInstanceSsl(request *model.SwitchGaussMySqlInstanceSslRequest) (*model.SwitchGaussMySqlInstanceSslResponse, error) {
	requestDef := GenReqDefForSwitchGaussMySqlInstanceSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchGaussMySqlInstanceSslResponse), nil
	}
}

// SwitchGaussMySqlInstanceSslInvoker 开关SSL
func (c *GaussDBClient) SwitchGaussMySqlInstanceSslInvoker(request *model.SwitchGaussMySqlInstanceSslRequest) *SwitchGaussMySqlInstanceSslInvoker {
	requestDef := GenReqDefForSwitchGaussMySqlInstanceSsl()
	return &SwitchGaussMySqlInstanceSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchGaussMySqlProxySsl 开关数据库代理SSL
//
// 为数据库代理设置SSL数据加密。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SwitchGaussMySqlProxySsl(request *model.SwitchGaussMySqlProxySslRequest) (*model.SwitchGaussMySqlProxySslResponse, error) {
	requestDef := GenReqDefForSwitchGaussMySqlProxySsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchGaussMySqlProxySslResponse), nil
	}
}

// SwitchGaussMySqlProxySslInvoker 开关数据库代理SSL
func (c *GaussDBClient) SwitchGaussMySqlProxySslInvoker(request *model.SwitchGaussMySqlProxySslRequest) *SwitchGaussMySqlProxySslInvoker {
	requestDef := GenReqDefForSwitchGaussMySqlProxySsl()
	return &SwitchGaussMySqlProxySslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditLog 开启或者关闭全量SQL
//
// 开启或者关闭全量SQL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateAuditLog(request *model.UpdateAuditLogRequest) (*model.UpdateAuditLogResponse, error) {
	requestDef := GenReqDefForUpdateAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditLogResponse), nil
	}
}

// UpdateAuditLogInvoker 开启或者关闭全量SQL
func (c *GaussDBClient) UpdateAuditLogInvoker(request *model.UpdateAuditLogRequest) *UpdateAuditLogInvoker {
	requestDef := GenReqDefForUpdateAuditLog()
	return &UpdateAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAutoScalingPolicy 设置自动变配
//
// 设置自动变配。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateAutoScalingPolicy(request *model.UpdateAutoScalingPolicyRequest) (*model.UpdateAutoScalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateAutoScalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAutoScalingPolicyResponse), nil
	}
}

// UpdateAutoScalingPolicyInvoker 设置自动变配
func (c *GaussDBClient) UpdateAutoScalingPolicyInvoker(request *model.UpdateAutoScalingPolicyRequest) *UpdateAutoScalingPolicyInvoker {
	requestDef := GenReqDefForUpdateAutoScalingPolicy()
	return &UpdateAutoScalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBackupOffsitePolicy 设置跨区备份策略
//
// 设置跨区备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateBackupOffsitePolicy(request *model.UpdateBackupOffsitePolicyRequest) (*model.UpdateBackupOffsitePolicyResponse, error) {
	requestDef := GenReqDefForUpdateBackupOffsitePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBackupOffsitePolicyResponse), nil
	}
}

// UpdateBackupOffsitePolicyInvoker 设置跨区备份策略
func (c *GaussDBClient) UpdateBackupOffsitePolicyInvoker(request *model.UpdateBackupOffsitePolicyRequest) *UpdateBackupOffsitePolicyInvoker {
	requestDef := GenReqDefForUpdateBackupOffsitePolicy()
	return &UpdateBackupOffsitePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlBackupPolicy 设置备份策略
//
// 设置自动备份策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlBackupPolicy(request *model.UpdateGaussMySqlBackupPolicyRequest) (*model.UpdateGaussMySqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlBackupPolicyResponse), nil
	}
}

// UpdateGaussMySqlBackupPolicyInvoker 设置备份策略
func (c *GaussDBClient) UpdateGaussMySqlBackupPolicyInvoker(request *model.UpdateGaussMySqlBackupPolicyRequest) *UpdateGaussMySqlBackupPolicyInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlBackupPolicy()
	return &UpdateGaussMySqlBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlConfiguration 修改参数模板
//
// 修改指定参数模板的参数信息，包括名称、描述、指定参数的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlConfiguration(request *model.UpdateGaussMySqlConfigurationRequest) (*model.UpdateGaussMySqlConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlConfigurationResponse), nil
	}
}

// UpdateGaussMySqlConfigurationInvoker 修改参数模板
func (c *GaussDBClient) UpdateGaussMySqlConfigurationInvoker(request *model.UpdateGaussMySqlConfigurationRequest) *UpdateGaussMySqlConfigurationInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlConfiguration()
	return &UpdateGaussMySqlConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlDatabaseComment 修改数据库备注
//
// 修改云数据库 GaussDB(for MySQL)实例数据库备注。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlDatabaseComment(request *model.UpdateGaussMySqlDatabaseCommentRequest) (*model.UpdateGaussMySqlDatabaseCommentResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlDatabaseComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlDatabaseCommentResponse), nil
	}
}

// UpdateGaussMySqlDatabaseCommentInvoker 修改数据库备注
func (c *GaussDBClient) UpdateGaussMySqlDatabaseCommentInvoker(request *model.UpdateGaussMySqlDatabaseCommentRequest) *UpdateGaussMySqlDatabaseCommentInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlDatabaseComment()
	return &UpdateGaussMySqlDatabaseCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlDatabaseUserComment 修改数据库用户备注
//
// 修改云数据库 GaussDB(for MySQL)实例数据库用户备注。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlDatabaseUserComment(request *model.UpdateGaussMySqlDatabaseUserCommentRequest) (*model.UpdateGaussMySqlDatabaseUserCommentResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlDatabaseUserComment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlDatabaseUserCommentResponse), nil
	}
}

// UpdateGaussMySqlDatabaseUserCommentInvoker 修改数据库用户备注
func (c *GaussDBClient) UpdateGaussMySqlDatabaseUserCommentInvoker(request *model.UpdateGaussMySqlDatabaseUserCommentRequest) *UpdateGaussMySqlDatabaseUserCommentInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlDatabaseUserComment()
	return &UpdateGaussMySqlDatabaseUserCommentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceAlias 修改实例备注
//
// 修改实例备注。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceAlias(request *model.UpdateGaussMySqlInstanceAliasRequest) (*model.UpdateGaussMySqlInstanceAliasResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceAlias()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceAliasResponse), nil
	}
}

// UpdateGaussMySqlInstanceAliasInvoker 修改实例备注
func (c *GaussDBClient) UpdateGaussMySqlInstanceAliasInvoker(request *model.UpdateGaussMySqlInstanceAliasRequest) *UpdateGaussMySqlInstanceAliasInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceAlias()
	return &UpdateGaussMySqlInstanceAliasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceEip 绑定弹性公网IP
//
// 实例绑定弹性公网IP，供外网连接使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceEip(request *model.UpdateGaussMySqlInstanceEipRequest) (*model.UpdateGaussMySqlInstanceEipResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceEipResponse), nil
	}
}

// UpdateGaussMySqlInstanceEipInvoker 绑定弹性公网IP
func (c *GaussDBClient) UpdateGaussMySqlInstanceEipInvoker(request *model.UpdateGaussMySqlInstanceEipRequest) *UpdateGaussMySqlInstanceEipInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceEip()
	return &UpdateGaussMySqlInstanceEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceInternalIp 修改内网地址
//
// 修改实例内网地址。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceInternalIp(request *model.UpdateGaussMySqlInstanceInternalIpRequest) (*model.UpdateGaussMySqlInstanceInternalIpResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceInternalIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceInternalIpResponse), nil
	}
}

// UpdateGaussMySqlInstanceInternalIpInvoker 修改内网地址
func (c *GaussDBClient) UpdateGaussMySqlInstanceInternalIpInvoker(request *model.UpdateGaussMySqlInstanceInternalIpRequest) *UpdateGaussMySqlInstanceInternalIpInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceInternalIp()
	return &UpdateGaussMySqlInstanceInternalIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceName 修改实例名称
//
// 修改实例名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceName(request *model.UpdateGaussMySqlInstanceNameRequest) (*model.UpdateGaussMySqlInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceNameResponse), nil
	}
}

// UpdateGaussMySqlInstanceNameInvoker 修改实例名称
func (c *GaussDBClient) UpdateGaussMySqlInstanceNameInvoker(request *model.UpdateGaussMySqlInstanceNameRequest) *UpdateGaussMySqlInstanceNameInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceName()
	return &UpdateGaussMySqlInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceOpsWindow 设置可维护时间段
//
// 设置可维护时间段。建议将可维护时间段设置在业务低峰期，避免业务在维护过程中异常中断。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceOpsWindow(request *model.UpdateGaussMySqlInstanceOpsWindowRequest) (*model.UpdateGaussMySqlInstanceOpsWindowResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceOpsWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceOpsWindowResponse), nil
	}
}

// UpdateGaussMySqlInstanceOpsWindowInvoker 设置可维护时间段
func (c *GaussDBClient) UpdateGaussMySqlInstanceOpsWindowInvoker(request *model.UpdateGaussMySqlInstanceOpsWindowRequest) *UpdateGaussMySqlInstanceOpsWindowInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceOpsWindow()
	return &UpdateGaussMySqlInstanceOpsWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstancePort 修改实例端口
//
// 修改实例数据库端口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstancePort(request *model.UpdateGaussMySqlInstancePortRequest) (*model.UpdateGaussMySqlInstancePortResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstancePortResponse), nil
	}
}

// UpdateGaussMySqlInstancePortInvoker 修改实例端口
func (c *GaussDBClient) UpdateGaussMySqlInstancePortInvoker(request *model.UpdateGaussMySqlInstancePortRequest) *UpdateGaussMySqlInstancePortInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstancePort()
	return &UpdateGaussMySqlInstancePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceSecurityGroup 修改安全组
//
// 修改指定实例安全组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlInstanceSecurityGroup(request *model.UpdateGaussMySqlInstanceSecurityGroupRequest) (*model.UpdateGaussMySqlInstanceSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlInstanceSecurityGroupResponse), nil
	}
}

// UpdateGaussMySqlInstanceSecurityGroupInvoker 修改安全组
func (c *GaussDBClient) UpdateGaussMySqlInstanceSecurityGroupInvoker(request *model.UpdateGaussMySqlInstanceSecurityGroupRequest) *UpdateGaussMySqlInstanceSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlInstanceSecurityGroup()
	return &UpdateGaussMySqlInstanceSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlQuotas 修改租户基于企业项目的资源配额
//
// 修改指定企业项目的资源配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlQuotas(request *model.UpdateGaussMySqlQuotasRequest) (*model.UpdateGaussMySqlQuotasResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlQuotasResponse), nil
	}
}

// UpdateGaussMySqlQuotasInvoker 修改租户基于企业项目的资源配额
func (c *GaussDBClient) UpdateGaussMySqlQuotasInvoker(request *model.UpdateGaussMySqlQuotasRequest) *UpdateGaussMySqlQuotasInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlQuotas()
	return &UpdateGaussMySqlQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceConfigurations 修改指定实例的参数
//
// 修改指定实例的参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateInstanceConfigurations(request *model.UpdateInstanceConfigurationsRequest) (*model.UpdateInstanceConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateInstanceConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceConfigurationsResponse), nil
	}
}

// UpdateInstanceConfigurationsInvoker 修改指定实例的参数
func (c *GaussDBClient) UpdateInstanceConfigurationsInvoker(request *model.UpdateInstanceConfigurationsRequest) *UpdateInstanceConfigurationsInvoker {
	requestDef := GenReqDefForUpdateInstanceConfigurations()
	return &UpdateInstanceConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceMonitor 设置实例秒级监控
//
// 设置实例秒级监控，包括1秒监控和5秒监控。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateInstanceMonitor(request *model.UpdateInstanceMonitorRequest) (*model.UpdateInstanceMonitorResponse, error) {
	requestDef := GenReqDefForUpdateInstanceMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceMonitorResponse), nil
	}
}

// UpdateInstanceMonitorInvoker 设置实例秒级监控
func (c *GaussDBClient) UpdateInstanceMonitorInvoker(request *model.UpdateInstanceMonitorRequest) *UpdateInstanceMonitorInvoker {
	requestDef := GenReqDefForUpdateInstanceMonitor()
	return &UpdateInstanceMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNewNodeAutoAddSwitch 开启或关闭新增节点自动加入该Proxy
//
// 开启或关闭新增节点自动加入该Proxy。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateNewNodeAutoAddSwitch(request *model.UpdateNewNodeAutoAddSwitchRequest) (*model.UpdateNewNodeAutoAddSwitchResponse, error) {
	requestDef := GenReqDefForUpdateNewNodeAutoAddSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNewNodeAutoAddSwitchResponse), nil
	}
}

// UpdateNewNodeAutoAddSwitchInvoker 开启或关闭新增节点自动加入该Proxy
func (c *GaussDBClient) UpdateNewNodeAutoAddSwitchInvoker(request *model.UpdateNewNodeAutoAddSwitchRequest) *UpdateNewNodeAutoAddSwitchInvoker {
	requestDef := GenReqDefForUpdateNewNodeAutoAddSwitch()
	return &UpdateNewNodeAutoAddSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxyConnectionPoolType 更改数据库代理连接池类型
//
// 更改数据库代理连接池类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateProxyConnectionPoolType(request *model.UpdateProxyConnectionPoolTypeRequest) (*model.UpdateProxyConnectionPoolTypeResponse, error) {
	requestDef := GenReqDefForUpdateProxyConnectionPoolType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxyConnectionPoolTypeResponse), nil
	}
}

// UpdateProxyConnectionPoolTypeInvoker 更改数据库代理连接池类型
func (c *GaussDBClient) UpdateProxyConnectionPoolTypeInvoker(request *model.UpdateProxyConnectionPoolTypeRequest) *UpdateProxyConnectionPoolTypeInvoker {
	requestDef := GenReqDefForUpdateProxyConnectionPoolType()
	return &UpdateProxyConnectionPoolTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxyName 修改代理实例名称
//
// 修改代理实例名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateProxyName(request *model.UpdateProxyNameRequest) (*model.UpdateProxyNameResponse, error) {
	requestDef := GenReqDefForUpdateProxyName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxyNameResponse), nil
	}
}

// UpdateProxyNameInvoker 修改代理实例名称
func (c *GaussDBClient) UpdateProxyNameInvoker(request *model.UpdateProxyNameRequest) *UpdateProxyNameInvoker {
	requestDef := GenReqDefForUpdateProxyName()
	return &UpdateProxyNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxyNewConfigurations 修改代理实例参数
//
// 修改数据库代理参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateProxyNewConfigurations(request *model.UpdateProxyNewConfigurationsRequest) (*model.UpdateProxyNewConfigurationsResponse, error) {
	requestDef := GenReqDefForUpdateProxyNewConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxyNewConfigurationsResponse), nil
	}
}

// UpdateProxyNewConfigurationsInvoker 修改代理实例参数
func (c *GaussDBClient) UpdateProxyNewConfigurationsInvoker(request *model.UpdateProxyNewConfigurationsRequest) *UpdateProxyNewConfigurationsInvoker {
	requestDef := GenReqDefForUpdateProxyNewConfigurations()
	return &UpdateProxyNewConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxyPort 修改读写分离端口号
//
// 修改读写分离端口号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateProxyPort(request *model.UpdateProxyPortRequest) (*model.UpdateProxyPortResponse, error) {
	requestDef := GenReqDefForUpdateProxyPort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxyPortResponse), nil
	}
}

// UpdateProxyPortInvoker 修改读写分离端口号
func (c *GaussDBClient) UpdateProxyPortInvoker(request *model.UpdateProxyPortRequest) *UpdateProxyPortInvoker {
	requestDef := GenReqDefForUpdateProxyPort()
	return &UpdateProxyPortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateProxySessionConsistence 修改代理会话一致性
//
// 修改代理会话一致性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateProxySessionConsistence(request *model.UpdateProxySessionConsistenceRequest) (*model.UpdateProxySessionConsistenceResponse, error) {
	requestDef := GenReqDefForUpdateProxySessionConsistence()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateProxySessionConsistenceResponse), nil
	}
}

// UpdateProxySessionConsistenceInvoker 修改代理会话一致性
func (c *GaussDBClient) UpdateProxySessionConsistenceInvoker(request *model.UpdateProxySessionConsistenceRequest) *UpdateProxySessionConsistenceInvoker {
	requestDef := GenReqDefForUpdateProxySessionConsistence()
	return &UpdateProxySessionConsistenceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServerlessPolicy 设置Serverless配置策略
//
// 设置Serverless配置策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateServerlessPolicy(request *model.UpdateServerlessPolicyRequest) (*model.UpdateServerlessPolicyResponse, error) {
	requestDef := GenReqDefForUpdateServerlessPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServerlessPolicyResponse), nil
	}
}

// UpdateServerlessPolicyInvoker 设置Serverless配置策略
func (c *GaussDBClient) UpdateServerlessPolicyInvoker(request *model.UpdateServerlessPolicyRequest) *UpdateServerlessPolicyInvoker {
	requestDef := GenReqDefForUpdateServerlessPolicy()
	return &UpdateServerlessPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSlowlogSensitiveSwitch 开启或关闭慢日志脱敏状态
//
// 开启或关闭慢日志脱敏状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateSlowlogSensitiveSwitch(request *model.UpdateSlowlogSensitiveSwitchRequest) (*model.UpdateSlowlogSensitiveSwitchResponse, error) {
	requestDef := GenReqDefForUpdateSlowlogSensitiveSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSlowlogSensitiveSwitchResponse), nil
	}
}

// UpdateSlowlogSensitiveSwitchInvoker 开启或关闭慢日志脱敏状态
func (c *GaussDBClient) UpdateSlowlogSensitiveSwitchInvoker(request *model.UpdateSlowlogSensitiveSwitchRequest) *UpdateSlowlogSensitiveSwitchInvoker {
	requestDef := GenReqDefForUpdateSlowlogSensitiveSwitch()
	return &UpdateSlowlogSensitiveSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTransactionSplitStatus 设置proxy事务拆分
//
// 设置proxy事务拆分。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateTransactionSplitStatus(request *model.UpdateTransactionSplitStatusRequest) (*model.UpdateTransactionSplitStatusResponse, error) {
	requestDef := GenReqDefForUpdateTransactionSplitStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTransactionSplitStatusResponse), nil
	}
}

// UpdateTransactionSplitStatusInvoker 设置proxy事务拆分
func (c *GaussDBClient) UpdateTransactionSplitStatusInvoker(request *model.UpdateTransactionSplitStatusRequest) *UpdateTransactionSplitStatusInvoker {
	requestDef := GenReqDefForUpdateTransactionSplitStatus()
	return &UpdateTransactionSplitStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeGaussMySqlInstanceDatabase 内核版本升级
//
// 内核版本升级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpgradeGaussMySqlInstanceDatabase(request *model.UpgradeGaussMySqlInstanceDatabaseRequest) (*model.UpgradeGaussMySqlInstanceDatabaseResponse, error) {
	requestDef := GenReqDefForUpgradeGaussMySqlInstanceDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeGaussMySqlInstanceDatabaseResponse), nil
	}
}

// UpgradeGaussMySqlInstanceDatabaseInvoker 内核版本升级
func (c *GaussDBClient) UpgradeGaussMySqlInstanceDatabaseInvoker(request *model.UpgradeGaussMySqlInstanceDatabaseRequest) *UpgradeGaussMySqlInstanceDatabaseInvoker {
	requestDef := GenReqDefForUpgradeGaussMySqlInstanceDatabase()
	return &UpgradeGaussMySqlInstanceDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeProxyVersion 升级数据库代理实例内核版本
//
// 升级数据库代理实例内核版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpgradeProxyVersion(request *model.UpgradeProxyVersionRequest) (*model.UpgradeProxyVersionResponse, error) {
	requestDef := GenReqDefForUpgradeProxyVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeProxyVersionResponse), nil
	}
}

// UpgradeProxyVersionInvoker 升级数据库代理实例内核版本
func (c *GaussDBClient) UpgradeProxyVersionInvoker(request *model.UpgradeProxyVersionRequest) *UpgradeProxyVersionInvoker {
	requestDef := GenReqDefForUpgradeProxyVersion()
	return &UpgradeProxyVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckClickHouseDataBaseConfig 数据同步库配置校验
//
// 数据同步库配置校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckClickHouseDataBaseConfig(request *model.CheckClickHouseDataBaseConfigRequest) (*model.CheckClickHouseDataBaseConfigResponse, error) {
	requestDef := GenReqDefForCheckClickHouseDataBaseConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckClickHouseDataBaseConfigResponse), nil
	}
}

// CheckClickHouseDataBaseConfigInvoker 数据同步库配置校验
func (c *GaussDBClient) CheckClickHouseDataBaseConfigInvoker(request *model.CheckClickHouseDataBaseConfigRequest) *CheckClickHouseDataBaseConfigInvoker {
	requestDef := GenReqDefForCheckClickHouseDataBaseConfig()
	return &CheckClickHouseDataBaseConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckClickHouseTableConfig 数据同步表配置校验
//
// 数据同步表配置校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckClickHouseTableConfig(request *model.CheckClickHouseTableConfigRequest) (*model.CheckClickHouseTableConfigResponse, error) {
	requestDef := GenReqDefForCheckClickHouseTableConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckClickHouseTableConfigResponse), nil
	}
}

// CheckClickHouseTableConfigInvoker 数据同步表配置校验
func (c *GaussDBClient) CheckClickHouseTableConfigInvoker(request *model.CheckClickHouseTableConfigRequest) *CheckClickHouseTableConfigInvoker {
	requestDef := GenReqDefForCheckClickHouseTableConfig()
	return &CheckClickHouseTableConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDataBaseConfig HTAP数据同步库配置校验
//
// HTAP数据同步库配置校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckDataBaseConfig(request *model.CheckDataBaseConfigRequest) (*model.CheckDataBaseConfigResponse, error) {
	requestDef := GenReqDefForCheckDataBaseConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDataBaseConfigResponse), nil
	}
}

// CheckDataBaseConfigInvoker HTAP数据同步库配置校验
func (c *GaussDBClient) CheckDataBaseConfigInvoker(request *model.CheckDataBaseConfigRequest) *CheckDataBaseConfigInvoker {
	requestDef := GenReqDefForCheckDataBaseConfig()
	return &CheckDataBaseConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckStarRocksResource StarRocks资源检查
//
// StarRocks资源检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckStarRocksResource(request *model.CheckStarRocksResourceRequest) (*model.CheckStarRocksResourceResponse, error) {
	requestDef := GenReqDefForCheckStarRocksResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckStarRocksResourceResponse), nil
	}
}

// CheckStarRocksResourceInvoker StarRocks资源检查
func (c *GaussDBClient) CheckStarRocksResourceInvoker(request *model.CheckStarRocksResourceRequest) *CheckStarRocksResourceInvoker {
	requestDef := GenReqDefForCheckStarRocksResource()
	return &CheckStarRocksResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckStarrocksParams 参数对比
//
// 对比实例参数和默认模板的差异
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckStarrocksParams(request *model.CheckStarrocksParamsRequest) (*model.CheckStarrocksParamsResponse, error) {
	requestDef := GenReqDefForCheckStarrocksParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckStarrocksParamsResponse), nil
	}
}

// CheckStarrocksParamsInvoker 参数对比
func (c *GaussDBClient) CheckStarrocksParamsInvoker(request *model.CheckStarrocksParamsRequest) *CheckStarrocksParamsInvoker {
	requestDef := GenReqDefForCheckStarrocksParams()
	return &CheckStarrocksParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckTableConfig HTAP数据同步表配置校验
//
// HTAP数据同步表配置校验。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CheckTableConfig(request *model.CheckTableConfigRequest) (*model.CheckTableConfigResponse, error) {
	requestDef := GenReqDefForCheckTableConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTableConfigResponse), nil
	}
}

// CheckTableConfigInvoker HTAP数据同步表配置校验
func (c *GaussDBClient) CheckTableConfigInvoker(request *model.CheckTableConfigRequest) *CheckTableConfigInvoker {
	requestDef := GenReqDefForCheckTableConfig()
	return &CheckTableConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClickHouseDataBaseReplication 创建数据同步
//
// 创建数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateClickHouseDataBaseReplication(request *model.CreateClickHouseDataBaseReplicationRequest) (*model.CreateClickHouseDataBaseReplicationResponse, error) {
	requestDef := GenReqDefForCreateClickHouseDataBaseReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClickHouseDataBaseReplicationResponse), nil
	}
}

// CreateClickHouseDataBaseReplicationInvoker 创建数据同步
func (c *GaussDBClient) CreateClickHouseDataBaseReplicationInvoker(request *model.CreateClickHouseDataBaseReplicationRequest) *CreateClickHouseDataBaseReplicationInvoker {
	requestDef := GenReqDefForCreateClickHouseDataBaseReplication()
	return &CreateClickHouseDataBaseReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClickHouseDatabaseUser 创建数据库账号
//
// 创建数据库账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateClickHouseDatabaseUser(request *model.CreateClickHouseDatabaseUserRequest) (*model.CreateClickHouseDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateClickHouseDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClickHouseDatabaseUserResponse), nil
	}
}

// CreateClickHouseDatabaseUserInvoker 创建数据库账号
func (c *GaussDBClient) CreateClickHouseDatabaseUserInvoker(request *model.CreateClickHouseDatabaseUserRequest) *CreateClickHouseDatabaseUserInvoker {
	requestDef := GenReqDefForCreateClickHouseDatabaseUser()
	return &CreateClickHouseDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateClickHouseInstance 创建实例
//
// 创建实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateClickHouseInstance(request *model.CreateClickHouseInstanceRequest) (*model.CreateClickHouseInstanceResponse, error) {
	requestDef := GenReqDefForCreateClickHouseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClickHouseInstanceResponse), nil
	}
}

// CreateClickHouseInstanceInvoker 创建实例
func (c *GaussDBClient) CreateClickHouseInstanceInvoker(request *model.CreateClickHouseInstanceRequest) *CreateClickHouseInstanceInvoker {
	requestDef := GenReqDefForCreateClickHouseInstance()
	return &CreateClickHouseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStarRocksDataReplication 创建StarRocks数据同步
//
// 创建StarRocks数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateStarRocksDataReplication(request *model.CreateStarRocksDataReplicationRequest) (*model.CreateStarRocksDataReplicationResponse, error) {
	requestDef := GenReqDefForCreateStarRocksDataReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStarRocksDataReplicationResponse), nil
	}
}

// CreateStarRocksDataReplicationInvoker 创建StarRocks数据同步
func (c *GaussDBClient) CreateStarRocksDataReplicationInvoker(request *model.CreateStarRocksDataReplicationRequest) *CreateStarRocksDataReplicationInvoker {
	requestDef := GenReqDefForCreateStarRocksDataReplication()
	return &CreateStarRocksDataReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStarRocksDatabaseUser 创建数据库账号
//
// 创建StarRocks数据库账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateStarRocksDatabaseUser(request *model.CreateStarRocksDatabaseUserRequest) (*model.CreateStarRocksDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateStarRocksDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStarRocksDatabaseUserResponse), nil
	}
}

// CreateStarRocksDatabaseUserInvoker 创建数据库账号
func (c *GaussDBClient) CreateStarRocksDatabaseUserInvoker(request *model.CreateStarRocksDatabaseUserRequest) *CreateStarRocksDatabaseUserInvoker {
	requestDef := GenReqDefForCreateStarRocksDatabaseUser()
	return &CreateStarRocksDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStarrocksInstance 创建StarRocks实例
//
// 创建StarRocks实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) CreateStarrocksInstance(request *model.CreateStarrocksInstanceRequest) (*model.CreateStarrocksInstanceResponse, error) {
	requestDef := GenReqDefForCreateStarrocksInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStarrocksInstanceResponse), nil
	}
}

// CreateStarrocksInstanceInvoker 创建StarRocks实例
func (c *GaussDBClient) CreateStarrocksInstanceInvoker(request *model.CreateStarrocksInstanceRequest) *CreateStarrocksInstanceInvoker {
	requestDef := GenReqDefForCreateStarrocksInstance()
	return &CreateStarrocksInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClickHouseDataBaseConfig 停止修改数据同步
//
// 停止修改数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteClickHouseDataBaseConfig(request *model.DeleteClickHouseDataBaseConfigRequest) (*model.DeleteClickHouseDataBaseConfigResponse, error) {
	requestDef := GenReqDefForDeleteClickHouseDataBaseConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClickHouseDataBaseConfigResponse), nil
	}
}

// DeleteClickHouseDataBaseConfigInvoker 停止修改数据同步
func (c *GaussDBClient) DeleteClickHouseDataBaseConfigInvoker(request *model.DeleteClickHouseDataBaseConfigRequest) *DeleteClickHouseDataBaseConfigInvoker {
	requestDef := GenReqDefForDeleteClickHouseDataBaseConfig()
	return &DeleteClickHouseDataBaseConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClickHouseDataBaseReplication 删除数据同步
//
// 删除数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteClickHouseDataBaseReplication(request *model.DeleteClickHouseDataBaseReplicationRequest) (*model.DeleteClickHouseDataBaseReplicationResponse, error) {
	requestDef := GenReqDefForDeleteClickHouseDataBaseReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClickHouseDataBaseReplicationResponse), nil
	}
}

// DeleteClickHouseDataBaseReplicationInvoker 删除数据同步
func (c *GaussDBClient) DeleteClickHouseDataBaseReplicationInvoker(request *model.DeleteClickHouseDataBaseReplicationRequest) *DeleteClickHouseDataBaseReplicationInvoker {
	requestDef := GenReqDefForDeleteClickHouseDataBaseReplication()
	return &DeleteClickHouseDataBaseReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClickHouseDatabaseUser 删除数据库账户
//
// 删除数据库账户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteClickHouseDatabaseUser(request *model.DeleteClickHouseDatabaseUserRequest) (*model.DeleteClickHouseDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteClickHouseDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClickHouseDatabaseUserResponse), nil
	}
}

// DeleteClickHouseDatabaseUserInvoker 删除数据库账户
func (c *GaussDBClient) DeleteClickHouseDatabaseUserInvoker(request *model.DeleteClickHouseDatabaseUserRequest) *DeleteClickHouseDatabaseUserInvoker {
	requestDef := GenReqDefForDeleteClickHouseDatabaseUser()
	return &DeleteClickHouseDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClickHouseInstance 删除实例
//
// 删除实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteClickHouseInstance(request *model.DeleteClickHouseInstanceRequest) (*model.DeleteClickHouseInstanceResponse, error) {
	requestDef := GenReqDefForDeleteClickHouseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClickHouseInstanceResponse), nil
	}
}

// DeleteClickHouseInstanceInvoker 删除实例
func (c *GaussDBClient) DeleteClickHouseInstanceInvoker(request *model.DeleteClickHouseInstanceRequest) *DeleteClickHouseInstanceInvoker {
	requestDef := GenReqDefForDeleteClickHouseInstance()
	return &DeleteClickHouseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClickHouseLtsConfig 批量解除LTS日志配置
//
// 批量解除LTS日志配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteClickHouseLtsConfig(request *model.DeleteClickHouseLtsConfigRequest) (*model.DeleteClickHouseLtsConfigResponse, error) {
	requestDef := GenReqDefForDeleteClickHouseLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClickHouseLtsConfigResponse), nil
	}
}

// DeleteClickHouseLtsConfigInvoker 批量解除LTS日志配置
func (c *GaussDBClient) DeleteClickHouseLtsConfigInvoker(request *model.DeleteClickHouseLtsConfigRequest) *DeleteClickHouseLtsConfigInvoker {
	requestDef := GenReqDefForDeleteClickHouseLtsConfig()
	return &DeleteClickHouseLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStarRocksDataReplication 删除StarRocks数据同步
//
// 删除StarRocks数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteStarRocksDataReplication(request *model.DeleteStarRocksDataReplicationRequest) (*model.DeleteStarRocksDataReplicationResponse, error) {
	requestDef := GenReqDefForDeleteStarRocksDataReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStarRocksDataReplicationResponse), nil
	}
}

// DeleteStarRocksDataReplicationInvoker 删除StarRocks数据同步
func (c *GaussDBClient) DeleteStarRocksDataReplicationInvoker(request *model.DeleteStarRocksDataReplicationRequest) *DeleteStarRocksDataReplicationInvoker {
	requestDef := GenReqDefForDeleteStarRocksDataReplication()
	return &DeleteStarRocksDataReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStarRocksDatabaseUser 删除数据库账户
//
// 删除StarRocks数据库账户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteStarRocksDatabaseUser(request *model.DeleteStarRocksDatabaseUserRequest) (*model.DeleteStarRocksDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteStarRocksDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStarRocksDatabaseUserResponse), nil
	}
}

// DeleteStarRocksDatabaseUserInvoker 删除数据库账户
func (c *GaussDBClient) DeleteStarRocksDatabaseUserInvoker(request *model.DeleteStarRocksDatabaseUserRequest) *DeleteStarRocksDatabaseUserInvoker {
	requestDef := GenReqDefForDeleteStarRocksDatabaseUser()
	return &DeleteStarRocksDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStarrocksInstance 删除StarRocks实例
//
// 删除StarRocks实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteStarrocksInstance(request *model.DeleteStarrocksInstanceRequest) (*model.DeleteStarrocksInstanceResponse, error) {
	requestDef := GenReqDefForDeleteStarrocksInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStarrocksInstanceResponse), nil
	}
}

// DeleteStarrocksInstanceInvoker 删除StarRocks实例
func (c *GaussDBClient) DeleteStarrocksInstanceInvoker(request *model.DeleteStarrocksInstanceRequest) *DeleteStarrocksInstanceInvoker {
	requestDef := GenReqDefForDeleteStarrocksInstance()
	return &DeleteStarrocksInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseDataBase 查询数据库列表
//
// 查询数据库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseDataBase(request *model.ListClickHouseDataBaseRequest) (*model.ListClickHouseDataBaseResponse, error) {
	requestDef := GenReqDefForListClickHouseDataBase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseDataBaseResponse), nil
	}
}

// ListClickHouseDataBaseInvoker 查询数据库列表
func (c *GaussDBClient) ListClickHouseDataBaseInvoker(request *model.ListClickHouseDataBaseRequest) *ListClickHouseDataBaseInvoker {
	requestDef := GenReqDefForListClickHouseDataBase()
	return &ListClickHouseDataBaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseDataBaseParameter 查询数据同步的库参数配置
//
// 查询数据同步的库参数配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseDataBaseParameter(request *model.ListClickHouseDataBaseParameterRequest) (*model.ListClickHouseDataBaseParameterResponse, error) {
	requestDef := GenReqDefForListClickHouseDataBaseParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseDataBaseParameterResponse), nil
	}
}

// ListClickHouseDataBaseParameterInvoker 查询数据同步的库参数配置
func (c *GaussDBClient) ListClickHouseDataBaseParameterInvoker(request *model.ListClickHouseDataBaseParameterRequest) *ListClickHouseDataBaseParameterInvoker {
	requestDef := GenReqDefForListClickHouseDataBaseParameter()
	return &ListClickHouseDataBaseParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseDataBaseReplication 查询数据同步信息
//
// 查询数据同步信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseDataBaseReplication(request *model.ListClickHouseDataBaseReplicationRequest) (*model.ListClickHouseDataBaseReplicationResponse, error) {
	requestDef := GenReqDefForListClickHouseDataBaseReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseDataBaseReplicationResponse), nil
	}
}

// ListClickHouseDataBaseReplicationInvoker 查询数据同步信息
func (c *GaussDBClient) ListClickHouseDataBaseReplicationInvoker(request *model.ListClickHouseDataBaseReplicationRequest) *ListClickHouseDataBaseReplicationInvoker {
	requestDef := GenReqDefForListClickHouseDataBaseReplication()
	return &ListClickHouseDataBaseReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseDataBaseReplicationConfig 查看数据同步配置
//
// 查看数据同步配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseDataBaseReplicationConfig(request *model.ListClickHouseDataBaseReplicationConfigRequest) (*model.ListClickHouseDataBaseReplicationConfigResponse, error) {
	requestDef := GenReqDefForListClickHouseDataBaseReplicationConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseDataBaseReplicationConfigResponse), nil
	}
}

// ListClickHouseDataBaseReplicationConfigInvoker 查看数据同步配置
func (c *GaussDBClient) ListClickHouseDataBaseReplicationConfigInvoker(request *model.ListClickHouseDataBaseReplicationConfigRequest) *ListClickHouseDataBaseReplicationConfigInvoker {
	requestDef := GenReqDefForListClickHouseDataBaseReplicationConfig()
	return &ListClickHouseDataBaseReplicationConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseInstance 查询实例详情
//
// 查询实例详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseInstance(request *model.ListClickHouseInstanceRequest) (*model.ListClickHouseInstanceResponse, error) {
	requestDef := GenReqDefForListClickHouseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseInstanceResponse), nil
	}
}

// ListClickHouseInstanceInvoker 查询实例详情
func (c *GaussDBClient) ListClickHouseInstanceInvoker(request *model.ListClickHouseInstanceRequest) *ListClickHouseInstanceInvoker {
	requestDef := GenReqDefForListClickHouseInstance()
	return &ListClickHouseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClickHouseInstanceNode 查询错误日志、慢日志节点信息
//
// 查询错误日志、慢日志节点信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListClickHouseInstanceNode(request *model.ListClickHouseInstanceNodeRequest) (*model.ListClickHouseInstanceNodeResponse, error) {
	requestDef := GenReqDefForListClickHouseInstanceNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClickHouseInstanceNodeResponse), nil
	}
}

// ListClickHouseInstanceNodeInvoker 查询错误日志、慢日志节点信息
func (c *GaussDBClient) ListClickHouseInstanceNodeInvoker(request *model.ListClickHouseInstanceNodeRequest) *ListClickHouseInstanceNodeInvoker {
	requestDef := GenReqDefForListClickHouseInstanceNode()
	return &ListClickHouseInstanceNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHtapDataStore HTAP引擎资源查询
//
// HTAP引擎资源查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListHtapDataStore(request *model.ListHtapDataStoreRequest) (*model.ListHtapDataStoreResponse, error) {
	requestDef := GenReqDefForListHtapDataStore()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHtapDataStoreResponse), nil
	}
}

// ListHtapDataStoreInvoker HTAP引擎资源查询
func (c *GaussDBClient) ListHtapDataStoreInvoker(request *model.ListHtapDataStoreRequest) *ListHtapDataStoreInvoker {
	requestDef := GenReqDefForListHtapDataStore()
	return &ListHtapDataStoreInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHtapFlavor HTAP查询规格信息
//
// HTAP查询规格信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListHtapFlavor(request *model.ListHtapFlavorRequest) (*model.ListHtapFlavorResponse, error) {
	requestDef := GenReqDefForListHtapFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHtapFlavorResponse), nil
	}
}

// ListHtapFlavorInvoker HTAP查询规格信息
func (c *GaussDBClient) ListHtapFlavorInvoker(request *model.ListHtapFlavorRequest) *ListHtapFlavorInvoker {
	requestDef := GenReqDefForListHtapFlavor()
	return &ListHtapFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHtapInstanceInfo 查询HTAP实例列表
//
// 查询HTAP实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListHtapInstanceInfo(request *model.ListHtapInstanceInfoRequest) (*model.ListHtapInstanceInfoResponse, error) {
	requestDef := GenReqDefForListHtapInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHtapInstanceInfoResponse), nil
	}
}

// ListHtapInstanceInfoInvoker 查询HTAP实例列表
func (c *GaussDBClient) ListHtapInstanceInfoInvoker(request *model.ListHtapInstanceInfoRequest) *ListHtapInstanceInfoInvoker {
	requestDef := GenReqDefForListHtapInstanceInfo()
	return &ListHtapInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHtapStorageType 获取HTAP实例存储类型
//
// 获取HTAP实例存储类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListHtapStorageType(request *model.ListHtapStorageTypeRequest) (*model.ListHtapStorageTypeResponse, error) {
	requestDef := GenReqDefForListHtapStorageType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHtapStorageTypeResponse), nil
	}
}

// ListHtapStorageTypeInvoker 获取HTAP实例存储类型
func (c *GaussDBClient) ListHtapStorageTypeInvoker(request *model.ListHtapStorageTypeRequest) *ListHtapStorageTypeInvoker {
	requestDef := GenReqDefForListHtapStorageType()
	return &ListHtapStorageTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStarRocksDataBases 查询StarRocks数据库
//
// 查询StarRocks数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListStarRocksDataBases(request *model.ListStarRocksDataBasesRequest) (*model.ListStarRocksDataBasesResponse, error) {
	requestDef := GenReqDefForListStarRocksDataBases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarRocksDataBasesResponse), nil
	}
}

// ListStarRocksDataBasesInvoker 查询StarRocks数据库
func (c *GaussDBClient) ListStarRocksDataBasesInvoker(request *model.ListStarRocksDataBasesRequest) *ListStarRocksDataBasesInvoker {
	requestDef := GenReqDefForListStarRocksDataBases()
	return &ListStarRocksDataBasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStarRocksDataReplicationConfig 查询StarRocks数据同步配置信息
//
// 查询StarRocks数据同步配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListStarRocksDataReplicationConfig(request *model.ListStarRocksDataReplicationConfigRequest) (*model.ListStarRocksDataReplicationConfigResponse, error) {
	requestDef := GenReqDefForListStarRocksDataReplicationConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarRocksDataReplicationConfigResponse), nil
	}
}

// ListStarRocksDataReplicationConfigInvoker 查询StarRocks数据同步配置信息
func (c *GaussDBClient) ListStarRocksDataReplicationConfigInvoker(request *model.ListStarRocksDataReplicationConfigRequest) *ListStarRocksDataReplicationConfigInvoker {
	requestDef := GenReqDefForListStarRocksDataReplicationConfig()
	return &ListStarRocksDataReplicationConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStarRocksDataReplications 查询StarRocks数据同步状态信息
//
// 查询StarRocks数据同步状态信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListStarRocksDataReplications(request *model.ListStarRocksDataReplicationsRequest) (*model.ListStarRocksDataReplicationsResponse, error) {
	requestDef := GenReqDefForListStarRocksDataReplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarRocksDataReplicationsResponse), nil
	}
}

// ListStarRocksDataReplicationsInvoker 查询StarRocks数据同步状态信息
func (c *GaussDBClient) ListStarRocksDataReplicationsInvoker(request *model.ListStarRocksDataReplicationsRequest) *ListStarRocksDataReplicationsInvoker {
	requestDef := GenReqDefForListStarRocksDataReplications()
	return &ListStarRocksDataReplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStarRocksDbParameters 查询StarRocks数据同步的库参数配置
//
// 查询StarRocks数据同步的库参数配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListStarRocksDbParameters(request *model.ListStarRocksDbParametersRequest) (*model.ListStarRocksDbParametersResponse, error) {
	requestDef := GenReqDefForListStarRocksDbParameters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarRocksDbParametersResponse), nil
	}
}

// ListStarRocksDbParametersInvoker 查询StarRocks数据同步的库参数配置
func (c *GaussDBClient) ListStarRocksDbParametersInvoker(request *model.ListStarRocksDbParametersRequest) *ListStarRocksDbParametersInvoker {
	requestDef := GenReqDefForListStarRocksDbParameters()
	return &ListStarRocksDbParametersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStarrocksInstanceInfo 查询StarRocks实例
//
// 查询StarRocks实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ListStarrocksInstanceInfo(request *model.ListStarrocksInstanceInfoRequest) (*model.ListStarrocksInstanceInfoResponse, error) {
	requestDef := GenReqDefForListStarrocksInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStarrocksInstanceInfoResponse), nil
	}
}

// ListStarrocksInstanceInfoInvoker 查询StarRocks实例
func (c *GaussDBClient) ListStarrocksInstanceInfoInvoker(request *model.ListStarrocksInstanceInfoRequest) *ListStarrocksInstanceInfoInvoker {
	requestDef := GenReqDefForListStarrocksInstanceInfo()
	return &ListStarrocksInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PauseStarRocksDataReplication 暂停StarRocks数据同步
//
// 暂停StarRocks数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) PauseStarRocksDataReplication(request *model.PauseStarRocksDataReplicationRequest) (*model.PauseStarRocksDataReplicationResponse, error) {
	requestDef := GenReqDefForPauseStarRocksDataReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PauseStarRocksDataReplicationResponse), nil
	}
}

// PauseStarRocksDataReplicationInvoker 暂停StarRocks数据同步
func (c *GaussDBClient) PauseStarRocksDataReplicationInvoker(request *model.PauseStarRocksDataReplicationRequest) *PauseStarRocksDataReplicationInvoker {
	requestDef := GenReqDefForPauseStarRocksDataReplication()
	return &PauseStarRocksDataReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootClickHouseInstance 重启实例
//
// 重启实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RebootClickHouseInstance(request *model.RebootClickHouseInstanceRequest) (*model.RebootClickHouseInstanceResponse, error) {
	requestDef := GenReqDefForRebootClickHouseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootClickHouseInstanceResponse), nil
	}
}

// RebootClickHouseInstanceInvoker 重启实例
func (c *GaussDBClient) RebootClickHouseInstanceInvoker(request *model.RebootClickHouseInstanceRequest) *RebootClickHouseInstanceInvoker {
	requestDef := GenReqDefForRebootClickHouseInstance()
	return &RebootClickHouseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeClickHouseFlavor 实例规格变更
//
// 实例规格变更。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResizeClickHouseFlavor(request *model.ResizeClickHouseFlavorRequest) (*model.ResizeClickHouseFlavorResponse, error) {
	requestDef := GenReqDefForResizeClickHouseFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeClickHouseFlavorResponse), nil
	}
}

// ResizeClickHouseFlavorInvoker 实例规格变更
func (c *GaussDBClient) ResizeClickHouseFlavorInvoker(request *model.ResizeClickHouseFlavorRequest) *ResizeClickHouseFlavorInvoker {
	requestDef := GenReqDefForResizeClickHouseFlavor()
	return &ResizeClickHouseFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeClickHouseInstance 实例磁盘扩容
//
// 实例磁盘扩容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResizeClickHouseInstance(request *model.ResizeClickHouseInstanceRequest) (*model.ResizeClickHouseInstanceResponse, error) {
	requestDef := GenReqDefForResizeClickHouseInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeClickHouseInstanceResponse), nil
	}
}

// ResizeClickHouseInstanceInvoker 实例磁盘扩容
func (c *GaussDBClient) ResizeClickHouseInstanceInvoker(request *model.ResizeClickHouseInstanceRequest) *ResizeClickHouseInstanceInvoker {
	requestDef := GenReqDefForResizeClickHouseInstance()
	return &ResizeClickHouseInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeStarRocksFlavor StarRocks实例规格变更
//
// StarRocks实例规格变更。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResizeStarRocksFlavor(request *model.ResizeStarRocksFlavorRequest) (*model.ResizeStarRocksFlavorResponse, error) {
	requestDef := GenReqDefForResizeStarRocksFlavor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeStarRocksFlavorResponse), nil
	}
}

// ResizeStarRocksFlavorInvoker StarRocks实例规格变更
func (c *GaussDBClient) ResizeStarRocksFlavorInvoker(request *model.ResizeStarRocksFlavorRequest) *ResizeStarRocksFlavorInvoker {
	requestDef := GenReqDefForResizeStarRocksFlavor()
	return &ResizeStarRocksFlavorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartStarrocksInstance 重启StarRocks实例
//
// 重启StarRocks实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestartStarrocksInstance(request *model.RestartStarrocksInstanceRequest) (*model.RestartStarrocksInstanceResponse, error) {
	requestDef := GenReqDefForRestartStarrocksInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartStarrocksInstanceResponse), nil
	}
}

// RestartStarrocksInstanceInvoker 重启StarRocks实例
func (c *GaussDBClient) RestartStarrocksInstanceInvoker(request *model.RestartStarrocksInstanceRequest) *RestartStarrocksInstanceInvoker {
	requestDef := GenReqDefForRestartStarrocksInstance()
	return &RestartStarrocksInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartStarrocksNode 重启StarRocks节点
//
// 重启StarRocks节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) RestartStarrocksNode(request *model.RestartStarrocksNodeRequest) (*model.RestartStarrocksNodeResponse, error) {
	requestDef := GenReqDefForRestartStarrocksNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartStarrocksNodeResponse), nil
	}
}

// RestartStarrocksNodeInvoker 重启StarRocks节点
func (c *GaussDBClient) RestartStarrocksNodeInvoker(request *model.RestartStarrocksNodeRequest) *RestartStarrocksNodeInvoker {
	requestDef := GenReqDefForRestartStarrocksNode()
	return &RestartStarrocksNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResumeStarRocksDataReplication 恢复StarRocks数据同步
//
// 恢复StarRocks数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ResumeStarRocksDataReplication(request *model.ResumeStarRocksDataReplicationRequest) (*model.ResumeStarRocksDataReplicationResponse, error) {
	requestDef := GenReqDefForResumeStarRocksDataReplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResumeStarRocksDataReplicationResponse), nil
	}
}

// ResumeStarRocksDataReplicationInvoker 恢复StarRocks数据同步
func (c *GaussDBClient) ResumeStarRocksDataReplicationInvoker(request *model.ResumeStarRocksDataReplicationRequest) *ResumeStarRocksDataReplicationInvoker {
	requestDef := GenReqDefForResumeStarRocksDataReplication()
	return &ResumeStarRocksDataReplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClickHouseDatabaseUser 查询数据库账户
//
// 查询数据库账户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowClickHouseDatabaseUser(request *model.ShowClickHouseDatabaseUserRequest) (*model.ShowClickHouseDatabaseUserResponse, error) {
	requestDef := GenReqDefForShowClickHouseDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClickHouseDatabaseUserResponse), nil
	}
}

// ShowClickHouseDatabaseUserInvoker 查询数据库账户
func (c *GaussDBClient) ShowClickHouseDatabaseUserInvoker(request *model.ShowClickHouseDatabaseUserRequest) *ShowClickHouseDatabaseUserInvoker {
	requestDef := GenReqDefForShowClickHouseDatabaseUser()
	return &ShowClickHouseDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClickHouseLtsConfig 查询实例LTS日志配置列表
//
// 查询实例LTS日志配置列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowClickHouseLtsConfig(request *model.ShowClickHouseLtsConfigRequest) (*model.ShowClickHouseLtsConfigResponse, error) {
	requestDef := GenReqDefForShowClickHouseLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClickHouseLtsConfigResponse), nil
	}
}

// ShowClickHouseLtsConfigInvoker 查询实例LTS日志配置列表
func (c *GaussDBClient) ShowClickHouseLtsConfigInvoker(request *model.ShowClickHouseLtsConfigRequest) *ShowClickHouseLtsConfigInvoker {
	requestDef := GenReqDefForShowClickHouseLtsConfig()
	return &ShowClickHouseLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClickHouseSlowLogDetail 查询慢日志
//
// 获取内核慢日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowClickHouseSlowLogDetail(request *model.ShowClickHouseSlowLogDetailRequest) (*model.ShowClickHouseSlowLogDetailResponse, error) {
	requestDef := GenReqDefForShowClickHouseSlowLogDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClickHouseSlowLogDetailResponse), nil
	}
}

// ShowClickHouseSlowLogDetailInvoker 查询慢日志
func (c *GaussDBClient) ShowClickHouseSlowLogDetailInvoker(request *model.ShowClickHouseSlowLogDetailRequest) *ShowClickHouseSlowLogDetailInvoker {
	requestDef := GenReqDefForShowClickHouseSlowLogDetail()
	return &ShowClickHouseSlowLogDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClickHouseSlowLogSensitiveStatus 查询慢日志脱敏状态
//
// 查询慢日志脱敏状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowClickHouseSlowLogSensitiveStatus(request *model.ShowClickHouseSlowLogSensitiveStatusRequest) (*model.ShowClickHouseSlowLogSensitiveStatusResponse, error) {
	requestDef := GenReqDefForShowClickHouseSlowLogSensitiveStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClickHouseSlowLogSensitiveStatusResponse), nil
	}
}

// ShowClickHouseSlowLogSensitiveStatusInvoker 查询慢日志脱敏状态
func (c *GaussDBClient) ShowClickHouseSlowLogSensitiveStatusInvoker(request *model.ShowClickHouseSlowLogSensitiveStatusRequest) *ShowClickHouseSlowLogSensitiveStatusInvoker {
	requestDef := GenReqDefForShowClickHouseSlowLogSensitiveStatus()
	return &ShowClickHouseSlowLogSensitiveStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStarRocksDatabaseUser 查询数据库账户
//
// 查询StarRocks数据库账户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowStarRocksDatabaseUser(request *model.ShowStarRocksDatabaseUserRequest) (*model.ShowStarRocksDatabaseUserResponse, error) {
	requestDef := GenReqDefForShowStarRocksDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStarRocksDatabaseUserResponse), nil
	}
}

// ShowStarRocksDatabaseUserInvoker 查询数据库账户
func (c *GaussDBClient) ShowStarRocksDatabaseUserInvoker(request *model.ShowStarRocksDatabaseUserRequest) *ShowStarRocksDatabaseUserInvoker {
	requestDef := GenReqDefForShowStarRocksDatabaseUser()
	return &ShowStarRocksDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStarrocksParams 查询参数
//
// 按节点类型查询参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowStarrocksParams(request *model.ShowStarrocksParamsRequest) (*model.ShowStarrocksParamsResponse, error) {
	requestDef := GenReqDefForShowStarrocksParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStarrocksParamsResponse), nil
	}
}

// ShowStarrocksParamsInvoker 查询参数
func (c *GaussDBClient) ShowStarrocksParamsInvoker(request *model.ShowStarrocksParamsRequest) *ShowStarrocksParamsInvoker {
	requestDef := GenReqDefForShowStarrocksParams()
	return &ShowStarrocksParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncStarRocksUsers StarRocks实例开启行列分流
//
// StarRocks实例开启行列分流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SyncStarRocksUsers(request *model.SyncStarRocksUsersRequest) (*model.SyncStarRocksUsersResponse, error) {
	requestDef := GenReqDefForSyncStarRocksUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncStarRocksUsersResponse), nil
	}
}

// SyncStarRocksUsersInvoker StarRocks实例开启行列分流
func (c *GaussDBClient) SyncStarRocksUsersInvoker(request *model.SyncStarRocksUsersRequest) *SyncStarRocksUsersInvoker {
	requestDef := GenReqDefForSyncStarRocksUsers()
	return &SyncStarRocksUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClickHouseDataBaseConfig 修改数据同步
//
// 修改数据同步。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateClickHouseDataBaseConfig(request *model.UpdateClickHouseDataBaseConfigRequest) (*model.UpdateClickHouseDataBaseConfigResponse, error) {
	requestDef := GenReqDefForUpdateClickHouseDataBaseConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClickHouseDataBaseConfigResponse), nil
	}
}

// UpdateClickHouseDataBaseConfigInvoker 修改数据同步
func (c *GaussDBClient) UpdateClickHouseDataBaseConfigInvoker(request *model.UpdateClickHouseDataBaseConfigRequest) *UpdateClickHouseDataBaseConfigInvoker {
	requestDef := GenReqDefForUpdateClickHouseDataBaseConfig()
	return &UpdateClickHouseDataBaseConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClickHouseDatabaseUserPassword 修改数据库账号密码
//
// 修改数据库账号密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateClickHouseDatabaseUserPassword(request *model.UpdateClickHouseDatabaseUserPasswordRequest) (*model.UpdateClickHouseDatabaseUserPasswordResponse, error) {
	requestDef := GenReqDefForUpdateClickHouseDatabaseUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClickHouseDatabaseUserPasswordResponse), nil
	}
}

// UpdateClickHouseDatabaseUserPasswordInvoker 修改数据库账号密码
func (c *GaussDBClient) UpdateClickHouseDatabaseUserPasswordInvoker(request *model.UpdateClickHouseDatabaseUserPasswordRequest) *UpdateClickHouseDatabaseUserPasswordInvoker {
	requestDef := GenReqDefForUpdateClickHouseDatabaseUserPassword()
	return &UpdateClickHouseDatabaseUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClickHouseDatabaseUserPermission 修改数据库账号权限
//
// 修改数据库账号权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateClickHouseDatabaseUserPermission(request *model.UpdateClickHouseDatabaseUserPermissionRequest) (*model.UpdateClickHouseDatabaseUserPermissionResponse, error) {
	requestDef := GenReqDefForUpdateClickHouseDatabaseUserPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClickHouseDatabaseUserPermissionResponse), nil
	}
}

// UpdateClickHouseDatabaseUserPermissionInvoker 修改数据库账号权限
func (c *GaussDBClient) UpdateClickHouseDatabaseUserPermissionInvoker(request *model.UpdateClickHouseDatabaseUserPermissionRequest) *UpdateClickHouseDatabaseUserPermissionInvoker {
	requestDef := GenReqDefForUpdateClickHouseDatabaseUserPermission()
	return &UpdateClickHouseDatabaseUserPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClickHouseLtsConfig 批量创建LTS日志配置
//
// 批量创建LTS日志配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateClickHouseLtsConfig(request *model.UpdateClickHouseLtsConfigRequest) (*model.UpdateClickHouseLtsConfigResponse, error) {
	requestDef := GenReqDefForUpdateClickHouseLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClickHouseLtsConfigResponse), nil
	}
}

// UpdateClickHouseLtsConfigInvoker 批量创建LTS日志配置
func (c *GaussDBClient) UpdateClickHouseLtsConfigInvoker(request *model.UpdateClickHouseLtsConfigRequest) *UpdateClickHouseLtsConfigInvoker {
	requestDef := GenReqDefForUpdateClickHouseLtsConfig()
	return &UpdateClickHouseLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClickHouseSlowLogSensitiveStatus 修改慢日志脱敏状态
//
// 修改慢日志脱敏状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateClickHouseSlowLogSensitiveStatus(request *model.UpdateClickHouseSlowLogSensitiveStatusRequest) (*model.UpdateClickHouseSlowLogSensitiveStatusResponse, error) {
	requestDef := GenReqDefForUpdateClickHouseSlowLogSensitiveStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClickHouseSlowLogSensitiveStatusResponse), nil
	}
}

// UpdateClickHouseSlowLogSensitiveStatusInvoker 修改慢日志脱敏状态
func (c *GaussDBClient) UpdateClickHouseSlowLogSensitiveStatusInvoker(request *model.UpdateClickHouseSlowLogSensitiveStatusRequest) *UpdateClickHouseSlowLogSensitiveStatusInvoker {
	requestDef := GenReqDefForUpdateClickHouseSlowLogSensitiveStatus()
	return &UpdateClickHouseSlowLogSensitiveStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStarRocksDatabaseUserPassword 修改数据库账号密码
//
// 修改StarRocks数据库账号密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateStarRocksDatabaseUserPassword(request *model.UpdateStarRocksDatabaseUserPasswordRequest) (*model.UpdateStarRocksDatabaseUserPasswordResponse, error) {
	requestDef := GenReqDefForUpdateStarRocksDatabaseUserPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStarRocksDatabaseUserPasswordResponse), nil
	}
}

// UpdateStarRocksDatabaseUserPasswordInvoker 修改数据库账号密码
func (c *GaussDBClient) UpdateStarRocksDatabaseUserPasswordInvoker(request *model.UpdateStarRocksDatabaseUserPasswordRequest) *UpdateStarRocksDatabaseUserPasswordInvoker {
	requestDef := GenReqDefForUpdateStarRocksDatabaseUserPassword()
	return &UpdateStarRocksDatabaseUserPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStarRocksDatabaseUserPermission 修改数据库账号权限
//
// 修改StarRocks数据库账号权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateStarRocksDatabaseUserPermission(request *model.UpdateStarRocksDatabaseUserPermissionRequest) (*model.UpdateStarRocksDatabaseUserPermissionResponse, error) {
	requestDef := GenReqDefForUpdateStarRocksDatabaseUserPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStarRocksDatabaseUserPermissionResponse), nil
	}
}

// UpdateStarRocksDatabaseUserPermissionInvoker 修改数据库账号权限
func (c *GaussDBClient) UpdateStarRocksDatabaseUserPermissionInvoker(request *model.UpdateStarRocksDatabaseUserPermissionRequest) *UpdateStarRocksDatabaseUserPermissionInvoker {
	requestDef := GenReqDefForUpdateStarRocksDatabaseUserPermission()
	return &UpdateStarRocksDatabaseUserPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStarrocksParams 修改参数
//
// 按节点类型修改节点参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateStarrocksParams(request *model.UpdateStarrocksParamsRequest) (*model.UpdateStarrocksParamsResponse, error) {
	requestDef := GenReqDefForUpdateStarrocksParams()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStarrocksParamsResponse), nil
	}
}

// UpdateStarrocksParamsInvoker 修改参数
func (c *GaussDBClient) UpdateStarrocksParamsInvoker(request *model.UpdateStarrocksParamsRequest) *UpdateStarrocksParamsInvoker {
	requestDef := GenReqDefForUpdateStarrocksParams()
	return &UpdateStarrocksParamsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlFilterRule 删除SQL限流规则
//
// 删除SQL限流规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) DeleteSqlFilterRule(request *model.DeleteSqlFilterRuleRequest) (*model.DeleteSqlFilterRuleResponse, error) {
	requestDef := GenReqDefForDeleteSqlFilterRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlFilterRuleResponse), nil
	}
}

// DeleteSqlFilterRuleInvoker 删除SQL限流规则
func (c *GaussDBClient) DeleteSqlFilterRuleInvoker(request *model.DeleteSqlFilterRuleRequest) *DeleteSqlFilterRuleInvoker {
	requestDef := GenReqDefForDeleteSqlFilterRule()
	return &DeleteSqlFilterRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetSqlFilterRule 设置SQL限流规则
//
// 设置SQL限流规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) SetSqlFilterRule(request *model.SetSqlFilterRuleRequest) (*model.SetSqlFilterRuleResponse, error) {
	requestDef := GenReqDefForSetSqlFilterRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetSqlFilterRuleResponse), nil
	}
}

// SetSqlFilterRuleInvoker 设置SQL限流规则
func (c *GaussDBClient) SetSqlFilterRuleInvoker(request *model.SetSqlFilterRuleRequest) *SetSqlFilterRuleInvoker {
	requestDef := GenReqDefForSetSqlFilterRule()
	return &SetSqlFilterRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlFilterControl 查询SQL限流开关状态
//
// 查询SQL限流开关状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowSqlFilterControl(request *model.ShowSqlFilterControlRequest) (*model.ShowSqlFilterControlResponse, error) {
	requestDef := GenReqDefForShowSqlFilterControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlFilterControlResponse), nil
	}
}

// ShowSqlFilterControlInvoker 查询SQL限流开关状态
func (c *GaussDBClient) ShowSqlFilterControlInvoker(request *model.ShowSqlFilterControlRequest) *ShowSqlFilterControlInvoker {
	requestDef := GenReqDefForShowSqlFilterControl()
	return &ShowSqlFilterControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlFilterRule 查询SQL限流规则
//
// 查询SQL限流规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) ShowSqlFilterRule(request *model.ShowSqlFilterRuleRequest) (*model.ShowSqlFilterRuleResponse, error) {
	requestDef := GenReqDefForShowSqlFilterRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlFilterRuleResponse), nil
	}
}

// ShowSqlFilterRuleInvoker 查询SQL限流规则
func (c *GaussDBClient) ShowSqlFilterRuleInvoker(request *model.ShowSqlFilterRuleRequest) *ShowSqlFilterRuleInvoker {
	requestDef := GenReqDefForShowSqlFilterRule()
	return &ShowSqlFilterRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSqlFilterControl 开启或者关闭SQL限流
//
// 开启或者关闭SQL限流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *GaussDBClient) UpdateSqlFilterControl(request *model.UpdateSqlFilterControlRequest) (*model.UpdateSqlFilterControlResponse, error) {
	requestDef := GenReqDefForUpdateSqlFilterControl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlFilterControlResponse), nil
	}
}

// UpdateSqlFilterControlInvoker 开启或者关闭SQL限流
func (c *GaussDBClient) UpdateSqlFilterControlInvoker(request *model.UpdateSqlFilterControlRequest) *UpdateSqlFilterControlInvoker {
	requestDef := GenReqDefForUpdateSqlFilterControl()
	return &UpdateSqlFilterControlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
