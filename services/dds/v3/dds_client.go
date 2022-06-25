package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"
)

type DdsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDdsClient(hcClient *http_client.HcHttpClient) *DdsClient {
	return &DdsClient{HcClient: hcClient}
}

func DdsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddShardingNode 扩容集群实例的节点数量
//
// 扩容指定集群实例的节点数量。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) AddShardingNode(request *model.AddShardingNodeRequest) (*model.AddShardingNodeResponse, error) {
	requestDef := GenReqDefForAddShardingNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddShardingNodeResponse), nil
	}
}

// AddShardingNodeInvoker 扩容集群实例的节点数量
func (c *DdsClient) AddShardingNodeInvoker(request *model.AddShardingNodeRequest) *AddShardingNodeInvoker {
	requestDef := GenReqDefForAddShardingNode()
	return &AddShardingNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachEip 绑定弹性公网IP
//
// 为实例下的节点绑定弹性公网IP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

// AttachEipInvoker 绑定弹性公网IP
func (c *DdsClient) AttachEipInvoker(request *model.AttachEipRequest) *AttachEipInvoker {
	requestDef := GenReqDefForAttachEip()
	return &AttachEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachInternalIp 修改实例内网地址
//
// 修改实例的内网地址
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) AttachInternalIp(request *model.AttachInternalIpRequest) (*model.AttachInternalIpResponse, error) {
	requestDef := GenReqDefForAttachInternalIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachInternalIpResponse), nil
	}
}

// AttachInternalIpInvoker 修改实例内网地址
func (c *DdsClient) AttachInternalIpInvoker(request *model.AttachInternalIpRequest) *AttachInternalIpInvoker {
	requestDef := GenReqDefForAttachInternalIp()
	return &AttachInternalIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchTagAction 批量添加或删除资源标签
//
// 批量添加或删除指定实例的标签。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

// BatchTagActionInvoker 批量添加或删除资源标签
func (c *DdsClient) BatchTagActionInvoker(request *model.BatchTagActionRequest) *BatchTagActionInvoker {
	requestDef := GenReqDefForBatchTagAction()
	return &BatchTagActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelEip 解绑弹性公网IP
//
// 解绑实例下节点已经绑定的弹性公网IP。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CancelEip(request *model.CancelEipRequest) (*model.CancelEipResponse, error) {
	requestDef := GenReqDefForCancelEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelEipResponse), nil
	}
}

// CancelEipInvoker 解绑弹性公网IP
func (c *DdsClient) CancelEipInvoker(request *model.CancelEipRequest) *CancelEipInvoker {
	requestDef := GenReqDefForCancelEip()
	return &CancelEipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckPassword 检查数据库密码
//
// 检查数据库密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CheckPassword(request *model.CheckPasswordRequest) (*model.CheckPasswordResponse, error) {
	requestDef := GenReqDefForCheckPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPasswordResponse), nil
	}
}

// CheckPasswordInvoker 检查数据库密码
func (c *DdsClient) CheckPasswordInvoker(request *model.CheckPasswordRequest) *CheckPasswordInvoker {
	requestDef := GenReqDefForCheckPassword()
	return &CheckPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfiguration 创建参数模板
//
// 创建参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateConfiguration(request *model.CreateConfigurationRequest) (*model.CreateConfigurationResponse, error) {
	requestDef := GenReqDefForCreateConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigurationResponse), nil
	}
}

// CreateConfigurationInvoker 创建参数模板
func (c *DdsClient) CreateConfigurationInvoker(request *model.CreateConfigurationRequest) *CreateConfigurationInvoker {
	requestDef := GenReqDefForCreateConfiguration()
	return &CreateConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseRole 创建数据库角色
//
// 创建数据库角色。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateDatabaseRole(request *model.CreateDatabaseRoleRequest) (*model.CreateDatabaseRoleResponse, error) {
	requestDef := GenReqDefForCreateDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseRoleResponse), nil
	}
}

// CreateDatabaseRoleInvoker 创建数据库角色
func (c *DdsClient) CreateDatabaseRoleInvoker(request *model.CreateDatabaseRoleRequest) *CreateDatabaseRoleInvoker {
	requestDef := GenReqDefForCreateDatabaseRole()
	return &CreateDatabaseRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabaseUser 创建数据库用户
//
// 创建数据库用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateDatabaseUser(request *model.CreateDatabaseUserRequest) (*model.CreateDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseUserResponse), nil
	}
}

// CreateDatabaseUserInvoker 创建数据库用户
func (c *DdsClient) CreateDatabaseUserInvoker(request *model.CreateDatabaseUserRequest) *CreateDatabaseUserInvoker {
	requestDef := GenReqDefForCreateDatabaseUser()
	return &CreateDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInstance 创建实例
//
// 创建文档数据库实例，包括集群实例、副本集实例、以及单节点实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

// CreateInstanceInvoker 创建实例
func (c *DdsClient) CreateInstanceInvoker(request *model.CreateInstanceRequest) *CreateInstanceInvoker {
	requestDef := GenReqDefForCreateInstance()
	return &CreateInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIp 打开集群的Shard/Config IP开关
//
// 打开集群的Shard/Config IP开关
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateIp(request *model.CreateIpRequest) (*model.CreateIpResponse, error) {
	requestDef := GenReqDefForCreateIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpResponse), nil
	}
}

// CreateIpInvoker 打开集群的Shard/Config IP开关
func (c *DdsClient) CreateIpInvoker(request *model.CreateIpRequest) *CreateIpInvoker {
	requestDef := GenReqDefForCreateIp()
	return &CreateIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManualBackup 创建手动备份
//
// 创建数据库实例的手动备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) CreateManualBackup(request *model.CreateManualBackupRequest) (*model.CreateManualBackupResponse, error) {
	requestDef := GenReqDefForCreateManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualBackupResponse), nil
	}
}

// CreateManualBackupInvoker 创建手动备份
func (c *DdsClient) CreateManualBackupInvoker(request *model.CreateManualBackupRequest) *CreateManualBackupInvoker {
	requestDef := GenReqDefForCreateManualBackup()
	return &CreateManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfiguration 删除参数模板
//
// 删除参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteConfiguration(request *model.DeleteConfigurationRequest) (*model.DeleteConfigurationResponse, error) {
	requestDef := GenReqDefForDeleteConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigurationResponse), nil
	}
}

// DeleteConfigurationInvoker 删除参数模板
func (c *DdsClient) DeleteConfigurationInvoker(request *model.DeleteConfigurationRequest) *DeleteConfigurationInvoker {
	requestDef := GenReqDefForDeleteConfiguration()
	return &DeleteConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabaseRole 删除数据库角色
//
// 删除数据库角色。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteDatabaseRole(request *model.DeleteDatabaseRoleRequest) (*model.DeleteDatabaseRoleResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseRoleResponse), nil
	}
}

// DeleteDatabaseRoleInvoker 删除数据库角色
func (c *DdsClient) DeleteDatabaseRoleInvoker(request *model.DeleteDatabaseRoleRequest) *DeleteDatabaseRoleInvoker {
	requestDef := GenReqDefForDeleteDatabaseRole()
	return &DeleteDatabaseRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabaseUser 删除数据库用户
//
// 删除数据库用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteDatabaseUser(request *model.DeleteDatabaseUserRequest) (*model.DeleteDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseUserResponse), nil
	}
}

// DeleteDatabaseUserInvoker 删除数据库用户
func (c *DdsClient) DeleteDatabaseUserInvoker(request *model.DeleteDatabaseUserRequest) *DeleteDatabaseUserInvoker {
	requestDef := GenReqDefForDeleteDatabaseUser()
	return &DeleteDatabaseUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInstance 删除实例
//
// 删除数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

// DeleteInstanceInvoker 删除实例
func (c *DdsClient) DeleteInstanceInvoker(request *model.DeleteInstanceRequest) *DeleteInstanceInvoker {
	requestDef := GenReqDefForDeleteInstance()
	return &DeleteInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteManualBackup 删除手动备份
//
// 删除数据库实例的手动备份。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteManualBackup(request *model.DeleteManualBackupRequest) (*model.DeleteManualBackupResponse, error) {
	requestDef := GenReqDefForDeleteManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManualBackupResponse), nil
	}
}

// DeleteManualBackupInvoker 删除手动备份
func (c *DdsClient) DeleteManualBackupInvoker(request *model.DeleteManualBackupRequest) *DeleteManualBackupInvoker {
	requestDef := GenReqDefForDeleteManualBackup()
	return &DeleteManualBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSession 终结实例节点会话
//
// 终结实例节点会话。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DeleteSession(request *model.DeleteSessionRequest) (*model.DeleteSessionResponse, error) {
	requestDef := GenReqDefForDeleteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSessionResponse), nil
	}
}

// DeleteSessionInvoker 终结实例节点会话
func (c *DdsClient) DeleteSessionInvoker(request *model.DeleteSessionRequest) *DeleteSessionInvoker {
	requestDef := GenReqDefForDeleteSession()
	return &DeleteSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadErrorlog 获取错误日志下载链接
//
// 获取错误日志下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DownloadErrorlog(request *model.DownloadErrorlogRequest) (*model.DownloadErrorlogResponse, error) {
	requestDef := GenReqDefForDownloadErrorlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadErrorlogResponse), nil
	}
}

// DownloadErrorlogInvoker 获取错误日志下载链接
func (c *DdsClient) DownloadErrorlogInvoker(request *model.DownloadErrorlogRequest) *DownloadErrorlogInvoker {
	requestDef := GenReqDefForDownloadErrorlog()
	return &DownloadErrorlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadSlowlog 获取慢日志下载链接
//
// 获取慢日志下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) DownloadSlowlog(request *model.DownloadSlowlogRequest) (*model.DownloadSlowlogResponse, error) {
	requestDef := GenReqDefForDownloadSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadSlowlogResponse), nil
	}
}

// DownloadSlowlogInvoker 获取慢日志下载链接
func (c *DdsClient) DownloadSlowlogInvoker(request *model.DownloadSlowlogRequest) *DownloadSlowlogInvoker {
	requestDef := GenReqDefForDownloadSlowlog()
	return &DownloadSlowlogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditlogLinks 获取审计日志下载链接
//
// 获取审计日志下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListAuditlogLinks(request *model.ListAuditlogLinksRequest) (*model.ListAuditlogLinksResponse, error) {
	requestDef := GenReqDefForListAuditlogLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogLinksResponse), nil
	}
}

// ListAuditlogLinksInvoker 获取审计日志下载链接
func (c *DdsClient) ListAuditlogLinksInvoker(request *model.ListAuditlogLinksRequest) *ListAuditlogLinksInvoker {
	requestDef := GenReqDefForListAuditlogLinks()
	return &ListAuditlogLinksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditlogs 获取审计日志列表
//
// 获取审计日志列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListAuditlogs(request *model.ListAuditlogsRequest) (*model.ListAuditlogsResponse, error) {
	requestDef := GenReqDefForListAuditlogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogsResponse), nil
	}
}

// ListAuditlogsInvoker 获取审计日志列表
func (c *DdsClient) ListAuditlogsInvoker(request *model.ListAuditlogsRequest) *ListAuditlogsInvoker {
	requestDef := GenReqDefForListAuditlogs()
	return &ListAuditlogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAz2Migrate 查询实例可迁移到的可用区
//
// 查询实例可迁移到的可用区。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListAz2Migrate(request *model.ListAz2MigrateRequest) (*model.ListAz2MigrateResponse, error) {
	requestDef := GenReqDefForListAz2Migrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAz2MigrateResponse), nil
	}
}

// ListAz2MigrateInvoker 查询实例可迁移到的可用区
func (c *DdsClient) ListAz2MigrateInvoker(request *model.ListAz2MigrateRequest) *ListAz2MigrateInvoker {
	requestDef := GenReqDefForListAz2Migrate()
	return &ListAz2MigrateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBackups 查询备份列表
//
// 根据指定条件查询备份列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

// ListBackupsInvoker 查询备份列表
func (c *DdsClient) ListBackupsInvoker(request *model.ListBackupsRequest) *ListBackupsInvoker {
	requestDef := GenReqDefForListBackups()
	return &ListBackupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigurations 获取参数模板列表
//
// 获取参数模板列表，包括DDS数据库的所有默认参数模板和用户创建的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListConfigurations(request *model.ListConfigurationsRequest) (*model.ListConfigurationsResponse, error) {
	requestDef := GenReqDefForListConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigurationsResponse), nil
	}
}

// ListConfigurationsInvoker 获取参数模板列表
func (c *DdsClient) ListConfigurationsInvoker(request *model.ListConfigurationsRequest) *ListConfigurationsInvoker {
	requestDef := GenReqDefForListConfigurations()
	return &ListConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseRoles 查询数据库角色列表
//
// 查询数据库角色列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListDatabaseRoles(request *model.ListDatabaseRolesRequest) (*model.ListDatabaseRolesResponse, error) {
	requestDef := GenReqDefForListDatabaseRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseRolesResponse), nil
	}
}

// ListDatabaseRolesInvoker 查询数据库角色列表
func (c *DdsClient) ListDatabaseRolesInvoker(request *model.ListDatabaseRolesRequest) *ListDatabaseRolesInvoker {
	requestDef := GenReqDefForListDatabaseRoles()
	return &ListDatabaseRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseUsers 查询数据库用户列表
//
// 查询数据库用户列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListDatabaseUsers(request *model.ListDatabaseUsersRequest) (*model.ListDatabaseUsersResponse, error) {
	requestDef := GenReqDefForListDatabaseUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseUsersResponse), nil
	}
}

// ListDatabaseUsersInvoker 查询数据库用户列表
func (c *DdsClient) ListDatabaseUsersInvoker(request *model.ListDatabaseUsersRequest) *ListDatabaseUsersInvoker {
	requestDef := GenReqDefForListDatabaseUsers()
	return &ListDatabaseUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatastoreVersions 查询数据库版本信息
//
// 查询指定实例类型的数据库版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListDatastoreVersions(request *model.ListDatastoreVersionsRequest) (*model.ListDatastoreVersionsResponse, error) {
	requestDef := GenReqDefForListDatastoreVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoreVersionsResponse), nil
	}
}

// ListDatastoreVersionsInvoker 查询数据库版本信息
func (c *DdsClient) ListDatastoreVersionsInvoker(request *model.ListDatastoreVersionsRequest) *ListDatastoreVersionsInvoker {
	requestDef := GenReqDefForListDatastoreVersions()
	return &ListDatastoreVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListErrorLogs 查询数据库错误日志
//
// 查询数据库错误信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListErrorLogs(request *model.ListErrorLogsRequest) (*model.ListErrorLogsResponse, error) {
	requestDef := GenReqDefForListErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorLogsResponse), nil
	}
}

// ListErrorLogsInvoker 查询数据库错误日志
func (c *DdsClient) ListErrorLogsInvoker(request *model.ListErrorLogsRequest) *ListErrorLogsInvoker {
	requestDef := GenReqDefForListErrorLogs()
	return &ListErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavorInfos 查询数据库规格
//
// 查询指定条件下的实例规格信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListFlavorInfos(request *model.ListFlavorInfosRequest) (*model.ListFlavorInfosResponse, error) {
	requestDef := GenReqDefForListFlavorInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorInfosResponse), nil
	}
}

// ListFlavorInfosInvoker 查询数据库规格
func (c *DdsClient) ListFlavorInfosInvoker(request *model.ListFlavorInfosRequest) *ListFlavorInfosInvoker {
	requestDef := GenReqDefForListFlavorInfos()
	return &ListFlavorInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询所有实例规格信息
//
// 查询指定条件下的所有实例规格信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询所有实例规格信息
func (c *DdsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

// ListInstanceTagsInvoker 查询资源标签
func (c *DdsClient) ListInstanceTagsInvoker(request *model.ListInstanceTagsRequest) *ListInstanceTagsInvoker {
	requestDef := GenReqDefForListInstanceTags()
	return &ListInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 查询实例列表和详情
//
// 根据指定条件查询实例列表和详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 查询实例列表和详情
func (c *DdsClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesByTags 查询资源实例
//
// 根据标签查询指定的数据库实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListInstancesByTags(request *model.ListInstancesByTagsRequest) (*model.ListInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagsResponse), nil
	}
}

// ListInstancesByTagsInvoker 查询资源实例
func (c *DdsClient) ListInstancesByTagsInvoker(request *model.ListInstancesByTagsRequest) *ListInstancesByTagsInvoker {
	requestDef := GenReqDefForListInstancesByTags()
	return &ListInstancesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询指定project ID下实例的所有标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

// ListProjectTagsInvoker 查询项目标签
func (c *DdsClient) ListProjectTagsInvoker(request *model.ListProjectTagsRequest) *ListProjectTagsInvoker {
	requestDef := GenReqDefForListProjectTags()
	return &ListProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreCollections 获取可恢复的数据库集合列表
//
// 获取可恢复的数据库集合列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListRestoreCollections(request *model.ListRestoreCollectionsRequest) (*model.ListRestoreCollectionsResponse, error) {
	requestDef := GenReqDefForListRestoreCollections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreCollectionsResponse), nil
	}
}

// ListRestoreCollectionsInvoker 获取可恢复的数据库集合列表
func (c *DdsClient) ListRestoreCollectionsInvoker(request *model.ListRestoreCollectionsRequest) *ListRestoreCollectionsInvoker {
	requestDef := GenReqDefForListRestoreCollections()
	return &ListRestoreCollectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreDatabases 获取可恢复的数据库列表
//
// 获取可恢复的数据库列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListRestoreDatabases(request *model.ListRestoreDatabasesRequest) (*model.ListRestoreDatabasesResponse, error) {
	requestDef := GenReqDefForListRestoreDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreDatabasesResponse), nil
	}
}

// ListRestoreDatabasesInvoker 获取可恢复的数据库列表
func (c *DdsClient) ListRestoreDatabasesInvoker(request *model.ListRestoreDatabasesRequest) *ListRestoreDatabasesInvoker {
	requestDef := GenReqDefForListRestoreDatabases()
	return &ListRestoreDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreTimes 查询可恢复的时间段
//
// 查询实例的可恢复时间段。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListRestoreTimes(request *model.ListRestoreTimesRequest) (*model.ListRestoreTimesResponse, error) {
	requestDef := GenReqDefForListRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimesResponse), nil
	}
}

// ListRestoreTimesInvoker 查询可恢复的时间段
func (c *DdsClient) ListRestoreTimesInvoker(request *model.ListRestoreTimesRequest) *ListRestoreTimesInvoker {
	requestDef := GenReqDefForListRestoreTimes()
	return &ListRestoreTimesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSessions 查询实例节点会话
//
// 查询实例节点会话。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListSessions(request *model.ListSessionsRequest) (*model.ListSessionsResponse, error) {
	requestDef := GenReqDefForListSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionsResponse), nil
	}
}

// ListSessionsInvoker 查询实例节点会话
func (c *DdsClient) ListSessionsInvoker(request *model.ListSessionsRequest) *ListSessionsInvoker {
	requestDef := GenReqDefForListSessions()
	return &ListSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSlowLogs 查询数据库慢日志
//
// 查询数据库慢日志信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

// ListSlowLogsInvoker 查询数据库慢日志
func (c *DdsClient) ListSlowLogsInvoker(request *model.ListSlowLogsRequest) *ListSlowLogsInvoker {
	requestDef := GenReqDefForListSlowLogs()
	return &ListSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageType 查询数据库磁盘类型
//
// 查询当前区域下的数据库磁盘类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListStorageType(request *model.ListStorageTypeRequest) (*model.ListStorageTypeResponse, error) {
	requestDef := GenReqDefForListStorageType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypeResponse), nil
	}
}

// ListStorageTypeInvoker 查询数据库磁盘类型
func (c *DdsClient) ListStorageTypeInvoker(request *model.ListStorageTypeRequest) *ListStorageTypeInvoker {
	requestDef := GenReqDefForListStorageType()
	return &ListStorageTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateAz 实例可用区迁移
//
// 实例可用区迁移。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) MigrateAz(request *model.MigrateAzRequest) (*model.MigrateAzResponse, error) {
	requestDef := GenReqDefForMigrateAz()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateAzResponse), nil
	}
}

// MigrateAzInvoker 实例可用区迁移
func (c *DdsClient) MigrateAzInvoker(request *model.MigrateAzRequest) *MigrateAzInvoker {
	requestDef := GenReqDefForMigrateAz()
	return &MigrateAzInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改数据库用户密码
//
// 修改数据库用户密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// ResetPasswordInvoker 修改数据库用户密码
func (c *DdsClient) ResetPasswordInvoker(request *model.ResetPasswordRequest) *ResetPasswordInvoker {
	requestDef := GenReqDefForResetPassword()
	return &ResetPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstance 变更实例规格
//
// 变更实例的规格。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

// ResizeInstanceInvoker 变更实例规格
func (c *DdsClient) ResizeInstanceInvoker(request *model.ResizeInstanceRequest) *ResizeInstanceInvoker {
	requestDef := GenReqDefForResizeInstance()
	return &ResizeInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeInstanceVolume 扩容实例存储容量
//
// 扩容实例相关的存储容量大小。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ResizeInstanceVolume(request *model.ResizeInstanceVolumeRequest) (*model.ResizeInstanceVolumeResponse, error) {
	requestDef := GenReqDefForResizeInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceVolumeResponse), nil
	}
}

// ResizeInstanceVolumeInvoker 扩容实例存储容量
func (c *DdsClient) ResizeInstanceVolumeInvoker(request *model.ResizeInstanceVolumeRequest) *ResizeInstanceVolumeInvoker {
	requestDef := GenReqDefForResizeInstanceVolume()
	return &ResizeInstanceVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartInstance 重启实例
//
// 重启实例的数据库服务。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

// RestartInstanceInvoker 重启实例
func (c *DdsClient) RestartInstanceInvoker(request *model.RestartInstanceRequest) *RestartInstanceInvoker {
	requestDef := GenReqDefForRestartInstance()
	return &RestartInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreInstance 恢复到当前实例
//
// 恢复到当前实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

// RestoreInstanceInvoker 恢复到当前实例
func (c *DdsClient) RestoreInstanceInvoker(request *model.RestoreInstanceRequest) *RestoreInstanceInvoker {
	requestDef := GenReqDefForRestoreInstance()
	return &RestoreInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreInstanceFromCollection 库表级时间点恢复
//
// 库表级时间点恢复。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) RestoreInstanceFromCollection(request *model.RestoreInstanceFromCollectionRequest) (*model.RestoreInstanceFromCollectionResponse, error) {
	requestDef := GenReqDefForRestoreInstanceFromCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceFromCollectionResponse), nil
	}
}

// RestoreInstanceFromCollectionInvoker 库表级时间点恢复
func (c *DdsClient) RestoreInstanceFromCollectionInvoker(request *model.RestoreInstanceFromCollectionRequest) *RestoreInstanceFromCollectionInvoker {
	requestDef := GenReqDefForRestoreInstanceFromCollection()
	return &RestoreInstanceFromCollectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestoreNewInstance 恢复到新实例
//
// 根据备份恢复新实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) RestoreNewInstance(request *model.RestoreNewInstanceRequest) (*model.RestoreNewInstanceResponse, error) {
	requestDef := GenReqDefForRestoreNewInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreNewInstanceResponse), nil
	}
}

// RestoreNewInstanceInvoker 恢复到新实例
func (c *DdsClient) RestoreNewInstanceInvoker(request *model.RestoreNewInstanceRequest) *RestoreNewInstanceInvoker {
	requestDef := GenReqDefForRestoreNewInstance()
	return &RestoreNewInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetAuditlogPolicy 设置审计日志策略
//
// 设置审计日志策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SetAuditlogPolicy(request *model.SetAuditlogPolicyRequest) (*model.SetAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForSetAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditlogPolicyResponse), nil
	}
}

// SetAuditlogPolicyInvoker 设置审计日志策略
func (c *DdsClient) SetAuditlogPolicyInvoker(request *model.SetAuditlogPolicyRequest) *SetAuditlogPolicyInvoker {
	requestDef := GenReqDefForSetAuditlogPolicy()
	return &SetAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBackupPolicy 设置自动备份策略
//
// 设置自动备份策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

// SetBackupPolicyInvoker 设置自动备份策略
func (c *DdsClient) SetBackupPolicyInvoker(request *model.SetBackupPolicyRequest) *SetBackupPolicyInvoker {
	requestDef := GenReqDefForSetBackupPolicy()
	return &SetBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBalancerSwitch 设置集群均衡开关
//
// 设置集群均衡开关。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SetBalancerSwitch(request *model.SetBalancerSwitchRequest) (*model.SetBalancerSwitchResponse, error) {
	requestDef := GenReqDefForSetBalancerSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerSwitchResponse), nil
	}
}

// SetBalancerSwitchInvoker 设置集群均衡开关
func (c *DdsClient) SetBalancerSwitchInvoker(request *model.SetBalancerSwitchRequest) *SetBalancerSwitchInvoker {
	requestDef := GenReqDefForSetBalancerSwitch()
	return &SetBalancerSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetBalancerWindow 设置集群均衡活动时间窗
//
// 设置集群均衡活动时间窗。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SetBalancerWindow(request *model.SetBalancerWindowRequest) (*model.SetBalancerWindowResponse, error) {
	requestDef := GenReqDefForSetBalancerWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerWindowResponse), nil
	}
}

// SetBalancerWindowInvoker 设置集群均衡活动时间窗
func (c *DdsClient) SetBalancerWindowInvoker(request *model.SetBalancerWindowRequest) *SetBalancerWindowInvoker {
	requestDef := GenReqDefForSetBalancerWindow()
	return &SetBalancerWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditlogPolicy 查询审计日志策略
//
// 查询审计日志策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowAuditlogPolicy(request *model.ShowAuditlogPolicyRequest) (*model.ShowAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForShowAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditlogPolicyResponse), nil
	}
}

// ShowAuditlogPolicyInvoker 查询审计日志策略
func (c *DdsClient) ShowAuditlogPolicyInvoker(request *model.ShowAuditlogPolicyRequest) *ShowAuditlogPolicyInvoker {
	requestDef := GenReqDefForShowAuditlogPolicy()
	return &ShowAuditlogPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupDownloadLink 获取备份下载链接
//
// 获取备份下载链接。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowBackupDownloadLink(request *model.ShowBackupDownloadLinkRequest) (*model.ShowBackupDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowBackupDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

// ShowBackupDownloadLinkInvoker 获取备份下载链接
func (c *DdsClient) ShowBackupDownloadLinkInvoker(request *model.ShowBackupDownloadLinkRequest) *ShowBackupDownloadLinkInvoker {
	requestDef := GenReqDefForShowBackupDownloadLink()
	return &ShowBackupDownloadLinkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

// ShowBackupPolicyInvoker 查询自动备份策略
func (c *DdsClient) ShowBackupPolicyInvoker(request *model.ShowBackupPolicyRequest) *ShowBackupPolicyInvoker {
	requestDef := GenReqDefForShowBackupPolicy()
	return &ShowBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationParameter 获取参数模板的详情
//
// 获取参数模板的详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowConfigurationParameter(request *model.ShowConfigurationParameterRequest) (*model.ShowConfigurationParameterResponse, error) {
	requestDef := GenReqDefForShowConfigurationParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationParameterResponse), nil
	}
}

// ShowConfigurationParameterInvoker 获取参数模板的详情
func (c *DdsClient) ShowConfigurationParameterInvoker(request *model.ShowConfigurationParameterRequest) *ShowConfigurationParameterInvoker {
	requestDef := GenReqDefForShowConfigurationParameter()
	return &ShowConfigurationParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnectionStatistics 查询实例连接数统计信息
//
// 查询客户端IP访问至DDS数据库实例的连接数统计信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowConnectionStatistics(request *model.ShowConnectionStatisticsRequest) (*model.ShowConnectionStatisticsResponse, error) {
	requestDef := GenReqDefForShowConnectionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionStatisticsResponse), nil
	}
}

// ShowConnectionStatisticsInvoker 查询实例连接数统计信息
func (c *DdsClient) ShowConnectionStatisticsInvoker(request *model.ShowConnectionStatisticsRequest) *ShowConnectionStatisticsInvoker {
	requestDef := GenReqDefForShowConnectionStatistics()
	return &ShowConnectionStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEntityConfiguration 获取指定实例的参数信息
//
// 获取指定实例的参数，可以是实例，组，节点的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowEntityConfiguration(request *model.ShowEntityConfigurationRequest) (*model.ShowEntityConfigurationResponse, error) {
	requestDef := GenReqDefForShowEntityConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEntityConfigurationResponse), nil
	}
}

// ShowEntityConfigurationInvoker 获取指定实例的参数信息
func (c *DdsClient) ShowEntityConfigurationInvoker(request *model.ShowEntityConfigurationRequest) *ShowEntityConfigurationInvoker {
	requestDef := GenReqDefForShowEntityConfiguration()
	return &ShowEntityConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobDetail 获取DDS任务中心指定ID的任务信息。
//
// 获取DDS任务中心指定ID的任务信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowJobDetail(request *model.ShowJobDetailRequest) (*model.ShowJobDetailResponse, error) {
	requestDef := GenReqDefForShowJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobDetailResponse), nil
	}
}

// ShowJobDetailInvoker 获取DDS任务中心指定ID的任务信息。
func (c *DdsClient) ShowJobDetailInvoker(request *model.ShowJobDetailRequest) *ShowJobDetailInvoker {
	requestDef := GenReqDefForShowJobDetail()
	return &ShowJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询单租户在DDS服务下的资源配额，包括单节点实例配额、副本集实例配额、集群实例配额等。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *DdsClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShardingBalancer 查询集群均衡设置
//
// 查询集群均衡设置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowShardingBalancer(request *model.ShowShardingBalancerRequest) (*model.ShowShardingBalancerResponse, error) {
	requestDef := GenReqDefForShowShardingBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShardingBalancerResponse), nil
	}
}

// ShowShardingBalancerInvoker 查询集群均衡设置
func (c *DdsClient) ShowShardingBalancerInvoker(request *model.ShowShardingBalancerRequest) *ShowShardingBalancerInvoker {
	requestDef := GenReqDefForShowShardingBalancer()
	return &ShowShardingBalancerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchConfiguration 应用参数模板
//
// 指定实例变更参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SwitchConfiguration(request *model.SwitchConfigurationRequest) (*model.SwitchConfigurationResponse, error) {
	requestDef := GenReqDefForSwitchConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchConfigurationResponse), nil
	}
}

// SwitchConfigurationInvoker 应用参数模板
func (c *DdsClient) SwitchConfigurationInvoker(request *model.SwitchConfigurationRequest) *SwitchConfigurationInvoker {
	requestDef := GenReqDefForSwitchConfiguration()
	return &SwitchConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSlowlogDesensitization 设置慢日志明文开关
//
// 设置实例的慢日志明文开关。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SwitchSlowlogDesensitization(request *model.SwitchSlowlogDesensitizationRequest) (*model.SwitchSlowlogDesensitizationResponse, error) {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSlowlogDesensitizationResponse), nil
	}
}

// SwitchSlowlogDesensitizationInvoker 设置慢日志明文开关
func (c *DdsClient) SwitchSlowlogDesensitizationInvoker(request *model.SwitchSlowlogDesensitizationRequest) *SwitchSlowlogDesensitizationInvoker {
	requestDef := GenReqDefForSwitchSlowlogDesensitization()
	return &SwitchSlowlogDesensitizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSsl 切换SSL开关
//
// 切换实例的SSL开关
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

// SwitchSslInvoker 切换SSL开关
func (c *DdsClient) SwitchSslInvoker(request *model.SwitchSslRequest) *SwitchSslInvoker {
	requestDef := GenReqDefForSwitchSsl()
	return &SwitchSslInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchoverReplicaSet 切换副本集实例的主备节点
//
// 切换副本集实例下的主备节点
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) SwitchoverReplicaSet(request *model.SwitchoverReplicaSetRequest) (*model.SwitchoverReplicaSetResponse, error) {
	requestDef := GenReqDefForSwitchoverReplicaSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchoverReplicaSetResponse), nil
	}
}

// SwitchoverReplicaSetInvoker 切换副本集实例的主备节点
func (c *DdsClient) SwitchoverReplicaSetInvoker(request *model.SwitchoverReplicaSetRequest) *SwitchoverReplicaSetInvoker {
	requestDef := GenReqDefForSwitchoverReplicaSet()
	return &SwitchoverReplicaSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClientNetwork 副本集跨网段访问配置。
//
// 副本集跨网段访问配置。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateClientNetwork(request *model.UpdateClientNetworkRequest) (*model.UpdateClientNetworkResponse, error) {
	requestDef := GenReqDefForUpdateClientNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClientNetworkResponse), nil
	}
}

// UpdateClientNetworkInvoker 副本集跨网段访问配置。
func (c *DdsClient) UpdateClientNetworkInvoker(request *model.UpdateClientNetworkRequest) *UpdateClientNetworkInvoker {
	requestDef := GenReqDefForUpdateClientNetwork()
	return &UpdateClientNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigurationParameter 修改参数模板
//
// 修改指定参数模板的参数信息，包括名称、描述、指定参数的值。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateConfigurationParameter(request *model.UpdateConfigurationParameterRequest) (*model.UpdateConfigurationParameterResponse, error) {
	requestDef := GenReqDefForUpdateConfigurationParameter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigurationParameterResponse), nil
	}
}

// UpdateConfigurationParameterInvoker 修改参数模板
func (c *DdsClient) UpdateConfigurationParameterInvoker(request *model.UpdateConfigurationParameterRequest) *UpdateConfigurationParameterInvoker {
	requestDef := GenReqDefForUpdateConfigurationParameter()
	return &UpdateConfigurationParameterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEntityConfiguration 修改指定实例的参数
//
// 修改指定实例的参数，可以是实例，组，节点的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateEntityConfiguration(request *model.UpdateEntityConfigurationRequest) (*model.UpdateEntityConfigurationResponse, error) {
	requestDef := GenReqDefForUpdateEntityConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEntityConfigurationResponse), nil
	}
}

// UpdateEntityConfigurationInvoker 修改指定实例的参数
func (c *DdsClient) UpdateEntityConfigurationInvoker(request *model.UpdateEntityConfigurationRequest) *UpdateEntityConfigurationInvoker {
	requestDef := GenReqDefForUpdateEntityConfiguration()
	return &UpdateEntityConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceName 修改实例名称
//
// 修改实例名称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

// UpdateInstanceNameInvoker 修改实例名称
func (c *DdsClient) UpdateInstanceNameInvoker(request *model.UpdateInstanceNameRequest) *UpdateInstanceNameInvoker {
	requestDef := GenReqDefForUpdateInstanceName()
	return &UpdateInstanceNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstancePort 修改数据库端口
//
// 修改数据库实例的端口。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateInstancePort(request *model.UpdateInstancePortRequest) (*model.UpdateInstancePortResponse, error) {
	requestDef := GenReqDefForUpdateInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstancePortResponse), nil
	}
}

// UpdateInstancePortInvoker 修改数据库端口
func (c *DdsClient) UpdateInstancePortInvoker(request *model.UpdateInstancePortRequest) *UpdateInstancePortInvoker {
	requestDef := GenReqDefForUpdateInstancePort()
	return &UpdateInstancePortInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInstanceRemark 修改实例备注
//
// 修改实例备注。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateInstanceRemark(request *model.UpdateInstanceRemarkRequest) (*model.UpdateInstanceRemarkResponse, error) {
	requestDef := GenReqDefForUpdateInstanceRemark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceRemarkResponse), nil
	}
}

// UpdateInstanceRemarkInvoker 修改实例备注
func (c *DdsClient) UpdateInstanceRemarkInvoker(request *model.UpdateInstanceRemarkRequest) *UpdateInstanceRemarkInvoker {
	requestDef := GenReqDefForUpdateInstanceRemark()
	return &UpdateInstanceRemarkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityGroup 变更实例安全组
//
// 变更实例关联的安全组
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupResponse), nil
	}
}

// UpdateSecurityGroupInvoker 变更实例安全组
func (c *DdsClient) UpdateSecurityGroupInvoker(request *model.UpdateSecurityGroupRequest) *UpdateSecurityGroupInvoker {
	requestDef := GenReqDefForUpdateSecurityGroup()
	return &UpdateSecurityGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersion 查询当前支持的API版本信息列表
//
// 查询当前支持的API版本信息列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

// ListApiVersionInvoker 查询当前支持的API版本信息列表
func (c *DdsClient) ListApiVersionInvoker(request *model.ListApiVersionRequest) *ListApiVersionInvoker {
	requestDef := GenReqDefForListApiVersion()
	return &ListApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiVersion 查询指定API版本信息
//
// 查询指定API版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DdsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}

// ShowApiVersionInvoker 查询指定API版本信息
func (c *DdsClient) ShowApiVersionInvoker(request *model.ShowApiVersionRequest) *ShowApiVersionInvoker {
	requestDef := GenReqDefForShowApiVersion()
	return &ShowApiVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
