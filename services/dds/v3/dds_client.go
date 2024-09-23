package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dds/v3/model"
)

type DdsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDdsClient(hcClient *httpclient.HcHttpClient) *DdsClient {
	return &DdsClient{HcClient: hcClient}
}

func DdsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddReadonlyNode 实例新增只读节点
//
// DDS副本集实例新增只读节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) AddReadonlyNode(request *model.AddReadonlyNodeRequest) (*model.AddReadonlyNodeResponse, error) {
	requestDef := GenReqDefForAddReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddReadonlyNodeResponse), nil
	}
}

// AddReadonlyNodeInvoker 实例新增只读节点
func (c *DdsClient) AddReadonlyNodeInvoker(request *model.AddReadonlyNodeRequest) *AddReadonlyNodeInvoker {
	requestDef := GenReqDefForAddReadonlyNode()
	return &AddReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddShardingNode 扩容集群实例的节点数量
//
// 扩容指定集群实例的节点数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ChangeOpsWindow 设置可维护时间段
//
// 修改用户允许启动某项对数据库实例运行有影响的任务的时间范围，例如操作系统升级和数据库软件版本升级的时间窗。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ChangeOpsWindow(request *model.ChangeOpsWindowRequest) (*model.ChangeOpsWindowResponse, error) {
	requestDef := GenReqDefForChangeOpsWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeOpsWindowResponse), nil
	}
}

// ChangeOpsWindowInvoker 设置可维护时间段
func (c *DdsClient) ChangeOpsWindowInvoker(request *model.ChangeOpsWindowRequest) *ChangeOpsWindowInvoker {
	requestDef := GenReqDefForChangeOpsWindow()
	return &ChangeOpsWindowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckPassword 检查数据库密码
//
// 检查数据库密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// CheckWeakPassword 检查弱密码
//
// 检查弱密码
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) CheckWeakPassword(request *model.CheckWeakPasswordRequest) (*model.CheckWeakPasswordResponse, error) {
	requestDef := GenReqDefForCheckWeakPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckWeakPasswordResponse), nil
	}
}

// CheckWeakPasswordInvoker 检查弱密码
func (c *DdsClient) CheckWeakPasswordInvoker(request *model.CheckWeakPasswordRequest) *CheckWeakPasswordInvoker {
	requestDef := GenReqDefForCheckWeakPassword()
	return &CheckWeakPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareConfiguration 参数模板比较
//
// 比较两个参数模板之间的差异。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) CompareConfiguration(request *model.CompareConfigurationRequest) (*model.CompareConfigurationResponse, error) {
	requestDef := GenReqDefForCompareConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareConfigurationResponse), nil
	}
}

// CompareConfigurationInvoker 参数模板比较
func (c *DdsClient) CompareConfigurationInvoker(request *model.CompareConfigurationRequest) *CompareConfigurationInvoker {
	requestDef := GenReqDefForCompareConfiguration()
	return &CompareConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyConfiguration 复制参数模板
//
// 复制参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) CopyConfiguration(request *model.CopyConfigurationRequest) (*model.CopyConfigurationResponse, error) {
	requestDef := GenReqDefForCopyConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyConfigurationResponse), nil
	}
}

// CopyConfigurationInvoker 复制参数模板
func (c *DdsClient) CopyConfigurationInvoker(request *model.CopyConfigurationRequest) *CopyConfigurationInvoker {
	requestDef := GenReqDefForCopyConfiguration()
	return &CopyConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfiguration 创建参数模板
//
// 创建参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// CreateIp 创建集群的Shard/Config IP
//
// 创建集群的Shard/Config IP
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) CreateIp(request *model.CreateIpRequest) (*model.CreateIpResponse, error) {
	requestDef := GenReqDefForCreateIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpResponse), nil
	}
}

// CreateIpInvoker 创建集群的Shard/Config IP
func (c *DdsClient) CreateIpInvoker(request *model.CreateIpRequest) *CreateIpInvoker {
	requestDef := GenReqDefForCreateIp()
	return &CreateIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateKillOpRule 创建killOp规则
//
// 创建killOp规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) CreateKillOpRule(request *model.CreateKillOpRuleRequest) (*model.CreateKillOpRuleResponse, error) {
	requestDef := GenReqDefForCreateKillOpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateKillOpRuleResponse), nil
	}
}

// CreateKillOpRuleInvoker 创建killOp规则
func (c *DdsClient) CreateKillOpRuleInvoker(request *model.CreateKillOpRuleRequest) *CreateKillOpRuleInvoker {
	requestDef := GenReqDefForCreateKillOpRule()
	return &CreateKillOpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManualBackup 创建手动备份
//
// 创建数据库实例的手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteAuditLog 删除审计日志
//
// 删除审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) DeleteAuditLog(request *model.DeleteAuditLogRequest) (*model.DeleteAuditLogResponse, error) {
	requestDef := GenReqDefForDeleteAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuditLogResponse), nil
	}
}

// DeleteAuditLogInvoker 删除审计日志
func (c *DdsClient) DeleteAuditLogInvoker(request *model.DeleteAuditLogRequest) *DeleteAuditLogInvoker {
	requestDef := GenReqDefForDeleteAuditLog()
	return &DeleteAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfiguration 删除参数模板
//
// 删除参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteKillOpRuleList 删除killOp规则
//
// 删除killOp规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) DeleteKillOpRuleList(request *model.DeleteKillOpRuleListRequest) (*model.DeleteKillOpRuleListResponse, error) {
	requestDef := GenReqDefForDeleteKillOpRuleList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteKillOpRuleListResponse), nil
	}
}

// DeleteKillOpRuleListInvoker 删除killOp规则
func (c *DdsClient) DeleteKillOpRuleListInvoker(request *model.DeleteKillOpRuleListRequest) *DeleteKillOpRuleListInvoker {
	requestDef := GenReqDefForDeleteKillOpRuleList()
	return &DeleteKillOpRuleListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLtsConfig 解除关联LTS日志流
//
// 将实例日志与LTS日志流解除关联，后台将取消上传实例日志到的LTS日志流里。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) DeleteLtsConfig(request *model.DeleteLtsConfigRequest) (*model.DeleteLtsConfigResponse, error) {
	requestDef := GenReqDefForDeleteLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLtsConfigResponse), nil
	}
}

// DeleteLtsConfigInvoker 解除关联LTS日志流
func (c *DdsClient) DeleteLtsConfigInvoker(request *model.DeleteLtsConfigRequest) *DeleteLtsConfigInvoker {
	requestDef := GenReqDefForDeleteLtsConfig()
	return &DeleteLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteManualBackup 删除手动备份
//
// 删除数据库实例的手动备份。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// DeleteReadonlyNode 删除只读节点
//
// 当副本集添加了只读节点后，需要删除对应的只读节点需要调用此API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) DeleteReadonlyNode(request *model.DeleteReadonlyNodeRequest) (*model.DeleteReadonlyNodeResponse, error) {
	requestDef := GenReqDefForDeleteReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteReadonlyNodeResponse), nil
	}
}

// DeleteReadonlyNodeInvoker 删除只读节点
func (c *DdsClient) DeleteReadonlyNodeInvoker(request *model.DeleteReadonlyNodeRequest) *DeleteReadonlyNodeInvoker {
	requestDef := GenReqDefForDeleteReadonlyNode()
	return &DeleteReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSession 终结实例节点会话
//
// 终结实例节点会话。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ExpandReplicasetNode 扩容副本集实例的节点数量
//
// 扩容指定副本集实例的节点数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ExpandReplicasetNode(request *model.ExpandReplicasetNodeRequest) (*model.ExpandReplicasetNodeResponse, error) {
	requestDef := GenReqDefForExpandReplicasetNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandReplicasetNodeResponse), nil
	}
}

// ExpandReplicasetNodeInvoker 扩容副本集实例的节点数量
func (c *DdsClient) ExpandReplicasetNodeInvoker(request *model.ExpandReplicasetNodeRequest) *ExpandReplicasetNodeInvoker {
	requestDef := GenReqDefForExpandReplicasetNode()
	return &ExpandReplicasetNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppliedInstances 查询可应用的实例
//
// 查询指定参数模板可被应用的实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListAppliedInstances(request *model.ListAppliedInstancesRequest) (*model.ListAppliedInstancesResponse, error) {
	requestDef := GenReqDefForListAppliedInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppliedInstancesResponse), nil
	}
}

// ListAppliedInstancesInvoker 查询可应用的实例
func (c *DdsClient) ListAppliedInstancesInvoker(request *model.ListAppliedInstancesRequest) *ListAppliedInstancesInvoker {
	requestDef := GenReqDefForListAppliedInstances()
	return &ListAppliedInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuditlogLinks 获取审计日志下载链接
//
// 获取审计日志下载链接。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListDatabases 查询数据库列表
//
// 查询数据库列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 查询数据库列表
func (c *DdsClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatastoreVersions 查询数据库版本信息
//
// 查询指定实例类型的数据库版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListLtsConfigs 查询LTS日志配置信息
//
// 查询LTS日志配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListLtsConfigs(request *model.ListLtsConfigsRequest) (*model.ListLtsConfigsResponse, error) {
	requestDef := GenReqDefForListLtsConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsConfigsResponse), nil
	}
}

// ListLtsConfigsInvoker 查询LTS日志配置信息
func (c *DdsClient) ListLtsConfigsInvoker(request *model.ListLtsConfigsRequest) *ListLtsConfigsInvoker {
	requestDef := GenReqDefForListLtsConfigs()
	return &ListLtsConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsErrorLogs 查询数据库错误日志
//
// 查询数据库错误日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListLtsErrorLogs(request *model.ListLtsErrorLogsRequest) (*model.ListLtsErrorLogsResponse, error) {
	requestDef := GenReqDefForListLtsErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsErrorLogsResponse), nil
	}
}

// ListLtsErrorLogsInvoker 查询数据库错误日志
func (c *DdsClient) ListLtsErrorLogsInvoker(request *model.ListLtsErrorLogsRequest) *ListLtsErrorLogsInvoker {
	requestDef := GenReqDefForListLtsErrorLogs()
	return &ListLtsErrorLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLtsSlowLogs 查询数据库慢日志
//
// 查询数据库慢日志信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListLtsSlowLogs(request *model.ListLtsSlowLogsRequest) (*model.ListLtsSlowLogsResponse, error) {
	requestDef := GenReqDefForListLtsSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLtsSlowLogsResponse), nil
	}
}

// ListLtsSlowLogsInvoker 查询数据库慢日志
func (c *DdsClient) ListLtsSlowLogsInvoker(request *model.ListLtsSlowLogsRequest) *ListLtsSlowLogsInvoker {
	requestDef := GenReqDefForListLtsSlowLogs()
	return &ListLtsSlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProjectTags 查询项目标签
//
// 查询指定project ID下实例的所有标签集合。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListRecycleInstances 查询回收站实例列表
//
// 查询回收站实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListRecycleInstances(request *model.ListRecycleInstancesRequest) (*model.ListRecycleInstancesResponse, error) {
	requestDef := GenReqDefForListRecycleInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecycleInstancesResponse), nil
	}
}

// ListRecycleInstancesInvoker 查询回收站实例列表
func (c *DdsClient) ListRecycleInstancesInvoker(request *model.ListRecycleInstancesRequest) *ListRecycleInstancesInvoker {
	requestDef := GenReqDefForListRecycleInstances()
	return &ListRecycleInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRestoreCollections 获取可恢复的数据库集合列表
//
// 获取可恢复的数据库集合列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListSslCertDownloadAddress 获取SSL证书下载地址
//
// 获取SSL证书下载地址
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListSslCertDownloadAddress(request *model.ListSslCertDownloadAddressRequest) (*model.ListSslCertDownloadAddressResponse, error) {
	requestDef := GenReqDefForListSslCertDownloadAddress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSslCertDownloadAddressResponse), nil
	}
}

// ListSslCertDownloadAddressInvoker 获取SSL证书下载地址
func (c *DdsClient) ListSslCertDownloadAddressInvoker(request *model.ListSslCertDownloadAddressRequest) *ListSslCertDownloadAddressInvoker {
	requestDef := GenReqDefForListSslCertDownloadAddress()
	return &ListSslCertDownloadAddressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListStorageType 查询数据库磁盘类型
//
// 查询当前区域下的数据库磁盘类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ListTasks 查询任务列表和详情
//
// 根据指定条件查询任务中心中的任务列表和详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询任务列表和详情
func (c *DdsClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateAz 实例可用区迁移
//
// 实例可用区迁移。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ResetConfiguration 重置参数模板
//
// 重置参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ResetConfiguration(request *model.ResetConfigurationRequest) (*model.ResetConfigurationResponse, error) {
	requestDef := GenReqDefForResetConfiguration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetConfigurationResponse), nil
	}
}

// ResetConfigurationInvoker 重置参数模板
func (c *DdsClient) ResetConfigurationInvoker(request *model.ResetConfigurationRequest) *ResetConfigurationInvoker {
	requestDef := GenReqDefForResetConfiguration()
	return &ResetConfigurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetPassword 修改数据库用户密码
//
// 修改数据库用户密码。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// SetRecyclePolicy 设置实例回收站策略
//
// 设置实例回收站策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) SetRecyclePolicy(request *model.SetRecyclePolicyRequest) (*model.SetRecyclePolicyResponse, error) {
	requestDef := GenReqDefForSetRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecyclePolicyResponse), nil
	}
}

// SetRecyclePolicyInvoker 设置实例回收站策略
func (c *DdsClient) SetRecyclePolicyInvoker(request *model.SetRecyclePolicyRequest) *SetRecyclePolicyInvoker {
	requestDef := GenReqDefForSetRecyclePolicy()
	return &SetRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuditlogPolicy 查询审计日志策略
//
// 查询审计日志策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowClientNetwork 查询副本集跨网段访问配置
//
// 查询副本集跨网段访问配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowClientNetwork(request *model.ShowClientNetworkRequest) (*model.ShowClientNetworkResponse, error) {
	requestDef := GenReqDefForShowClientNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClientNetworkResponse), nil
	}
}

// ShowClientNetworkInvoker 查询副本集跨网段访问配置
func (c *DdsClient) ShowClientNetworkInvoker(request *model.ShowClientNetworkRequest) *ShowClientNetworkInvoker {
	requestDef := GenReqDefForShowClientNetwork()
	return &ShowClientNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationAppliedHistory 查询参数模板被应用历史
//
// 查询参数模板应用历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowConfigurationAppliedHistory(request *model.ShowConfigurationAppliedHistoryRequest) (*model.ShowConfigurationAppliedHistoryResponse, error) {
	requestDef := GenReqDefForShowConfigurationAppliedHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationAppliedHistoryResponse), nil
	}
}

// ShowConfigurationAppliedHistoryInvoker 查询参数模板被应用历史
func (c *DdsClient) ShowConfigurationAppliedHistoryInvoker(request *model.ShowConfigurationAppliedHistoryRequest) *ShowConfigurationAppliedHistoryInvoker {
	requestDef := GenReqDefForShowConfigurationAppliedHistory()
	return &ShowConfigurationAppliedHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationModifyHistory 查询参数模板修改历史
//
// 查询参数模板修改历史。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowConfigurationModifyHistory(request *model.ShowConfigurationModifyHistoryRequest) (*model.ShowConfigurationModifyHistoryResponse, error) {
	requestDef := GenReqDefForShowConfigurationModifyHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigurationModifyHistoryResponse), nil
	}
}

// ShowConfigurationModifyHistoryInvoker 查询参数模板修改历史
func (c *DdsClient) ShowConfigurationModifyHistoryInvoker(request *model.ShowConfigurationModifyHistoryRequest) *ShowConfigurationModifyHistoryInvoker {
	requestDef := GenReqDefForShowConfigurationModifyHistory()
	return &ShowConfigurationModifyHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigurationParameter 获取参数模板的详情
//
// 获取参数模板的详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowDiskUsage 查询实例磁盘信息
//
// 查询实例磁盘信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowDiskUsage(request *model.ShowDiskUsageRequest) (*model.ShowDiskUsageResponse, error) {
	requestDef := GenReqDefForShowDiskUsage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiskUsageResponse), nil
	}
}

// ShowDiskUsageInvoker 查询实例磁盘信息
func (c *DdsClient) ShowDiskUsageInvoker(request *model.ShowDiskUsageRequest) *ShowDiskUsageInvoker {
	requestDef := GenReqDefForShowDiskUsage()
	return &ShowDiskUsageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEntityConfiguration 获取指定实例的参数信息
//
// 获取指定实例的参数，可以是实例，组，节点的参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowKillOpRuleRuleList 获取killOp规则列表
//
// 获取killOp规则列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowKillOpRuleRuleList(request *model.ShowKillOpRuleRuleListRequest) (*model.ShowKillOpRuleRuleListResponse, error) {
	requestDef := GenReqDefForShowKillOpRuleRuleList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowKillOpRuleRuleListResponse), nil
	}
}

// ShowKillOpRuleRuleListInvoker 获取killOp规则列表
func (c *DdsClient) ShowKillOpRuleRuleListInvoker(request *model.ShowKillOpRuleRuleListRequest) *ShowKillOpRuleRuleListInvoker {
	requestDef := GenReqDefForShowKillOpRuleRuleList()
	return &ShowKillOpRuleRuleListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询单租户在DDS服务下的资源配额，包括单节点实例配额、副本集实例配额、集群实例配额等。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowRecyclePolicy 查询实例回收站策略
//
// 查询实例回收站策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowRecyclePolicy(request *model.ShowRecyclePolicyRequest) (*model.ShowRecyclePolicyResponse, error) {
	requestDef := GenReqDefForShowRecyclePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecyclePolicyResponse), nil
	}
}

// ShowRecyclePolicyInvoker 查询实例回收站策略
func (c *DdsClient) ShowRecyclePolicyInvoker(request *model.ShowRecyclePolicyRequest) *ShowRecyclePolicyInvoker {
	requestDef := GenReqDefForShowRecyclePolicy()
	return &ShowRecyclePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowReplSetName 查询数据库复制集名称
//
// 查询数据库复制集名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowReplSetName(request *model.ShowReplSetNameRequest) (*model.ShowReplSetNameResponse, error) {
	requestDef := GenReqDefForShowReplSetName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowReplSetNameResponse), nil
	}
}

// ShowReplSetNameInvoker 查询数据库复制集名称
func (c *DdsClient) ShowReplSetNameInvoker(request *model.ShowReplSetNameRequest) *ShowReplSetNameInvoker {
	requestDef := GenReqDefForShowReplSetName()
	return &ShowReplSetNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecondLevelMonitoringStatus 查询秒级监控配置
//
// 查询秒级监控配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowSecondLevelMonitoringStatus(request *model.ShowSecondLevelMonitoringStatusRequest) (*model.ShowSecondLevelMonitoringStatusResponse, error) {
	requestDef := GenReqDefForShowSecondLevelMonitoringStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecondLevelMonitoringStatusResponse), nil
	}
}

// ShowSecondLevelMonitoringStatusInvoker 查询秒级监控配置
func (c *DdsClient) ShowSecondLevelMonitoringStatusInvoker(request *model.ShowSecondLevelMonitoringStatusRequest) *ShowSecondLevelMonitoringStatusInvoker {
	requestDef := GenReqDefForShowSecondLevelMonitoringStatus()
	return &ShowSecondLevelMonitoringStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowShardingBalancer 查询集群均衡设置
//
// 查询集群均衡设置。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// ShowSlowlogDesensitizationSwitch 查询慢日志明文开关
//
// 查询慢日志明文开关
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowSlowlogDesensitizationSwitch(request *model.ShowSlowlogDesensitizationSwitchRequest) (*model.ShowSlowlogDesensitizationSwitchResponse, error) {
	requestDef := GenReqDefForShowSlowlogDesensitizationSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSlowlogDesensitizationSwitchResponse), nil
	}
}

// ShowSlowlogDesensitizationSwitchInvoker 查询慢日志明文开关
func (c *DdsClient) ShowSlowlogDesensitizationSwitchInvoker(request *model.ShowSlowlogDesensitizationSwitchRequest) *ShowSlowlogDesensitizationSwitchInvoker {
	requestDef := GenReqDefForShowSlowlogDesensitizationSwitch()
	return &ShowSlowlogDesensitizationSwitchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpgradeDuration 查询数据库补丁升级预估时长
//
// 查询数据库补丁升级预估时长
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShowUpgradeDuration(request *model.ShowUpgradeDurationRequest) (*model.ShowUpgradeDurationResponse, error) {
	requestDef := GenReqDefForShowUpgradeDuration()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeDurationResponse), nil
	}
}

// ShowUpgradeDurationInvoker 查询数据库补丁升级预估时长
func (c *DdsClient) ShowUpgradeDurationInvoker(request *model.ShowUpgradeDurationRequest) *ShowUpgradeDurationInvoker {
	requestDef := GenReqDefForShowUpgradeDuration()
	return &ShowUpgradeDurationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShrinkInstanceNodes 删除实例的节点
//
// 删除实例的节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) ShrinkInstanceNodes(request *model.ShrinkInstanceNodesRequest) (*model.ShrinkInstanceNodesResponse, error) {
	requestDef := GenReqDefForShrinkInstanceNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShrinkInstanceNodesResponse), nil
	}
}

// ShrinkInstanceNodesInvoker 删除实例的节点
func (c *DdsClient) ShrinkInstanceNodesInvoker(request *model.ShrinkInstanceNodesRequest) *ShrinkInstanceNodesInvoker {
	requestDef := GenReqDefForShrinkInstanceNodes()
	return &ShrinkInstanceNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopBackup 停止备份
//
// 支持紧急情况下停止备份功能。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) StopBackup(request *model.StopBackupRequest) (*model.StopBackupResponse, error) {
	requestDef := GenReqDefForStopBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopBackupResponse), nil
	}
}

// StopBackupInvoker 停止备份
func (c *DdsClient) StopBackupInvoker(request *model.StopBackupRequest) *StopBackupInvoker {
	requestDef := GenReqDefForStopBackup()
	return &StopBackupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchConfiguration 应用参数模板
//
// 指定实例变更参数模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// SwitchInstancePrimary 强制备节点升主
//
// 支持副本集、shard和config备节点强制升主。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) SwitchInstancePrimary(request *model.SwitchInstancePrimaryRequest) (*model.SwitchInstancePrimaryResponse, error) {
	requestDef := GenReqDefForSwitchInstancePrimary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchInstancePrimaryResponse), nil
	}
}

// SwitchInstancePrimaryInvoker 强制备节点升主
func (c *DdsClient) SwitchInstancePrimaryInvoker(request *model.SwitchInstancePrimaryRequest) *SwitchInstancePrimaryInvoker {
	requestDef := GenReqDefForSwitchInstancePrimary()
	return &SwitchInstancePrimaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSecondLevelMonitoring 开启/关闭秒级监控
//
// 开启或关闭指定实例的秒级监控。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) SwitchSecondLevelMonitoring(request *model.SwitchSecondLevelMonitoringRequest) (*model.SwitchSecondLevelMonitoringResponse, error) {
	requestDef := GenReqDefForSwitchSecondLevelMonitoring()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSecondLevelMonitoringResponse), nil
	}
}

// SwitchSecondLevelMonitoringInvoker 开启/关闭秒级监控
func (c *DdsClient) SwitchSecondLevelMonitoringInvoker(request *model.SwitchSecondLevelMonitoringRequest) *SwitchSecondLevelMonitoringInvoker {
	requestDef := GenReqDefForSwitchSecondLevelMonitoring()
	return &SwitchSecondLevelMonitoringInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchSlowlogDesensitization 设置慢日志明文开关
//
// 设置实例的慢日志明文开关。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpdateKillOpRule 启用/禁用killOp规则
//
// 启用/禁用killOp规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) UpdateKillOpRule(request *model.UpdateKillOpRuleRequest) (*model.UpdateKillOpRuleResponse, error) {
	requestDef := GenReqDefForUpdateKillOpRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateKillOpRuleResponse), nil
	}
}

// UpdateKillOpRuleInvoker 启用/禁用killOp规则
func (c *DdsClient) UpdateKillOpRuleInvoker(request *model.UpdateKillOpRuleRequest) *UpdateKillOpRuleInvoker {
	requestDef := GenReqDefForUpdateKillOpRule()
	return &UpdateKillOpRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLtsConfig 关联LTS日志流
//
// 将实例日志与LTS日志流关联，后台将自动上传实例日志到关联的LTS日志流里。
// 关联成功后，会产生一定费用，具体计费可参考云日志服务（LTS）的定价详情。
// 系统会为当前选择的日志流创建对应日志类型的结构化配置。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) UpdateLtsConfig(request *model.UpdateLtsConfigRequest) (*model.UpdateLtsConfigResponse, error) {
	requestDef := GenReqDefForUpdateLtsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLtsConfigResponse), nil
	}
}

// UpdateLtsConfigInvoker 关联LTS日志流
func (c *DdsClient) UpdateLtsConfigInvoker(request *model.UpdateLtsConfigRequest) *UpdateLtsConfigInvoker {
	requestDef := GenReqDefForUpdateLtsConfig()
	return &UpdateLtsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateReplSetName 修改数据库复制集名称
//
// 修改数据库复制集名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) UpdateReplSetName(request *model.UpdateReplSetNameRequest) (*model.UpdateReplSetNameResponse, error) {
	requestDef := GenReqDefForUpdateReplSetName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateReplSetNameResponse), nil
	}
}

// UpdateReplSetNameInvoker 修改数据库复制集名称
func (c *DdsClient) UpdateReplSetNameInvoker(request *model.UpdateReplSetNameRequest) *UpdateReplSetNameInvoker {
	requestDef := GenReqDefForUpdateReplSetName()
	return &UpdateReplSetNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityGroup 变更实例安全组
//
// 变更实例关联的安全组
//
// Please refer to HUAWEI cloud API Explorer for details.
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

// UpgradeDatabaseVersion 数据库补丁升级
//
// 升级数据库补丁版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DdsClient) UpgradeDatabaseVersion(request *model.UpgradeDatabaseVersionRequest) (*model.UpgradeDatabaseVersionResponse, error) {
	requestDef := GenReqDefForUpgradeDatabaseVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeDatabaseVersionResponse), nil
	}
}

// UpgradeDatabaseVersionInvoker 数据库补丁升级
func (c *DdsClient) UpgradeDatabaseVersionInvoker(request *model.UpgradeDatabaseVersionRequest) *UpgradeDatabaseVersionInvoker {
	requestDef := GenReqDefForUpgradeDatabaseVersion()
	return &UpgradeDatabaseVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiVersion 查询当前支持的API版本信息列表
//
// 查询当前支持的API版本信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
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
// Please refer to HUAWEI cloud API Explorer for details.
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
