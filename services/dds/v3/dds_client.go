package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

//扩容指定集群实例的节点数量。
func (c *DdsClient) AddShardingNode(request *model.AddShardingNodeRequest) (*model.AddShardingNodeResponse, error) {
	requestDef := GenReqDefForAddShardingNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddShardingNodeResponse), nil
	}
}

//为实例下的节点绑定弹性公网IP。
func (c *DdsClient) AttachEip(request *model.AttachEipRequest) (*model.AttachEipResponse, error) {
	requestDef := GenReqDefForAttachEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachEipResponse), nil
	}
}

//修改实例的内网地址
func (c *DdsClient) AttachInternalIp(request *model.AttachInternalIpRequest) (*model.AttachInternalIpResponse, error) {
	requestDef := GenReqDefForAttachInternalIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachInternalIpResponse), nil
	}
}

//批量添加或删除指定实例的标签。
func (c *DdsClient) BatchTagAction(request *model.BatchTagActionRequest) (*model.BatchTagActionResponse, error) {
	requestDef := GenReqDefForBatchTagAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchTagActionResponse), nil
	}
}

//解绑实例下节点已经绑定的弹性公网IP。
func (c *DdsClient) CancelEip(request *model.CancelEipRequest) (*model.CancelEipResponse, error) {
	requestDef := GenReqDefForCancelEip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelEipResponse), nil
	}
}

//检查数据库密码。
func (c *DdsClient) CheckPassword(request *model.CheckPasswordRequest) (*model.CheckPasswordResponse, error) {
	requestDef := GenReqDefForCheckPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckPasswordResponse), nil
	}
}

//创建数据库角色。
func (c *DdsClient) CreateDatabaseRole(request *model.CreateDatabaseRoleRequest) (*model.CreateDatabaseRoleResponse, error) {
	requestDef := GenReqDefForCreateDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseRoleResponse), nil
	}
}

//创建数据库用户。
func (c *DdsClient) CreateDatabaseUser(request *model.CreateDatabaseUserRequest) (*model.CreateDatabaseUserResponse, error) {
	requestDef := GenReqDefForCreateDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseUserResponse), nil
	}
}

//创建文档数据库实例，包括集群实例、副本集实例、以及单节点实例。
func (c *DdsClient) CreateInstance(request *model.CreateInstanceRequest) (*model.CreateInstanceResponse, error) {
	requestDef := GenReqDefForCreateInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInstanceResponse), nil
	}
}

//打开集群的Shard/Config IP开关
func (c *DdsClient) CreateIp(request *model.CreateIpRequest) (*model.CreateIpResponse, error) {
	requestDef := GenReqDefForCreateIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIpResponse), nil
	}
}

//创建数据库实例的手动备份。
func (c *DdsClient) CreateManualBackup(request *model.CreateManualBackupRequest) (*model.CreateManualBackupResponse, error) {
	requestDef := GenReqDefForCreateManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManualBackupResponse), nil
	}
}

//删除数据库角色。
func (c *DdsClient) DeleteDatabaseRole(request *model.DeleteDatabaseRoleRequest) (*model.DeleteDatabaseRoleResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseRoleResponse), nil
	}
}

//删除数据库用户。
func (c *DdsClient) DeleteDatabaseUser(request *model.DeleteDatabaseUserRequest) (*model.DeleteDatabaseUserResponse, error) {
	requestDef := GenReqDefForDeleteDatabaseUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseUserResponse), nil
	}
}

//删除数据库实例。
func (c *DdsClient) DeleteInstance(request *model.DeleteInstanceRequest) (*model.DeleteInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInstanceResponse), nil
	}
}

//删除数据库实例的手动备份。
func (c *DdsClient) DeleteManualBackup(request *model.DeleteManualBackupRequest) (*model.DeleteManualBackupResponse, error) {
	requestDef := GenReqDefForDeleteManualBackup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteManualBackupResponse), nil
	}
}

//终结实例节点会话。
func (c *DdsClient) DeleteSession(request *model.DeleteSessionRequest) (*model.DeleteSessionResponse, error) {
	requestDef := GenReqDefForDeleteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSessionResponse), nil
	}
}

//获取错误日志下载链接。
func (c *DdsClient) DownloadErrorlog(request *model.DownloadErrorlogRequest) (*model.DownloadErrorlogResponse, error) {
	requestDef := GenReqDefForDownloadErrorlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadErrorlogResponse), nil
	}
}

//获取慢日志下载链接。
func (c *DdsClient) DownloadSlowlog(request *model.DownloadSlowlogRequest) (*model.DownloadSlowlogResponse, error) {
	requestDef := GenReqDefForDownloadSlowlog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadSlowlogResponse), nil
	}
}

//获取审计日志下载链接。
func (c *DdsClient) ListAuditlogLinks(request *model.ListAuditlogLinksRequest) (*model.ListAuditlogLinksResponse, error) {
	requestDef := GenReqDefForListAuditlogLinks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogLinksResponse), nil
	}
}

//获取审计日志列表。
func (c *DdsClient) ListAuditlogs(request *model.ListAuditlogsRequest) (*model.ListAuditlogsResponse, error) {
	requestDef := GenReqDefForListAuditlogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuditlogsResponse), nil
	}
}

//查询实例可迁移到的可用区。
func (c *DdsClient) ListAz2Migrate(request *model.ListAz2MigrateRequest) (*model.ListAz2MigrateResponse, error) {
	requestDef := GenReqDefForListAz2Migrate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAz2MigrateResponse), nil
	}
}

//根据指定条件查询备份列表。
func (c *DdsClient) ListBackups(request *model.ListBackupsRequest) (*model.ListBackupsResponse, error) {
	requestDef := GenReqDefForListBackups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBackupsResponse), nil
	}
}

//查询数据库角色列表。
func (c *DdsClient) ListDatabaseRoles(request *model.ListDatabaseRolesRequest) (*model.ListDatabaseRolesResponse, error) {
	requestDef := GenReqDefForListDatabaseRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseRolesResponse), nil
	}
}

//查询数据库用户列表。
func (c *DdsClient) ListDatabaseUsers(request *model.ListDatabaseUsersRequest) (*model.ListDatabaseUsersResponse, error) {
	requestDef := GenReqDefForListDatabaseUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseUsersResponse), nil
	}
}

//查询指定实例类型的数据库版本信息。
func (c *DdsClient) ListDatastoreVersions(request *model.ListDatastoreVersionsRequest) (*model.ListDatastoreVersionsResponse, error) {
	requestDef := GenReqDefForListDatastoreVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatastoreVersionsResponse), nil
	}
}

//查询数据库错误信息。
func (c *DdsClient) ListErrorLogs(request *model.ListErrorLogsRequest) (*model.ListErrorLogsResponse, error) {
	requestDef := GenReqDefForListErrorLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListErrorLogsResponse), nil
	}
}

//查询指定条件下的所有实例规格信息。
func (c *DdsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

//查询指定实例的标签信息。
func (c *DdsClient) ListInstanceTags(request *model.ListInstanceTagsRequest) (*model.ListInstanceTagsResponse, error) {
	requestDef := GenReqDefForListInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceTagsResponse), nil
	}
}

//根据指定条件查询实例列表。
func (c *DdsClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

//根据标签查询指定的数据库实例。
func (c *DdsClient) ListInstancesByTags(request *model.ListInstancesByTagsRequest) (*model.ListInstancesByTagsResponse, error) {
	requestDef := GenReqDefForListInstancesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesByTagsResponse), nil
	}
}

//查询指定project ID下实例的所有标签集合。
func (c *DdsClient) ListProjectTags(request *model.ListProjectTagsRequest) (*model.ListProjectTagsResponse, error) {
	requestDef := GenReqDefForListProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProjectTagsResponse), nil
	}
}

//获取可恢复的数据库集合列表。
func (c *DdsClient) ListRestoreCollections(request *model.ListRestoreCollectionsRequest) (*model.ListRestoreCollectionsResponse, error) {
	requestDef := GenReqDefForListRestoreCollections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreCollectionsResponse), nil
	}
}

//获取可恢复的数据库列表。
func (c *DdsClient) ListRestoreDatabases(request *model.ListRestoreDatabasesRequest) (*model.ListRestoreDatabasesResponse, error) {
	requestDef := GenReqDefForListRestoreDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreDatabasesResponse), nil
	}
}

//查询实例的可恢复时间段。
func (c *DdsClient) ListRestoreTimes(request *model.ListRestoreTimesRequest) (*model.ListRestoreTimesResponse, error) {
	requestDef := GenReqDefForListRestoreTimes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRestoreTimesResponse), nil
	}
}

//查询实例节点会话。
func (c *DdsClient) ListSessions(request *model.ListSessionsRequest) (*model.ListSessionsResponse, error) {
	requestDef := GenReqDefForListSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSessionsResponse), nil
	}
}

//查询数据库慢日志信息。
func (c *DdsClient) ListSlowLogs(request *model.ListSlowLogsRequest) (*model.ListSlowLogsResponse, error) {
	requestDef := GenReqDefForListSlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSlowLogsResponse), nil
	}
}

//查询当前区域下的数据库磁盘类型。
func (c *DdsClient) ListStorageType(request *model.ListStorageTypeRequest) (*model.ListStorageTypeResponse, error) {
	requestDef := GenReqDefForListStorageType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListStorageTypeResponse), nil
	}
}

//实例可用区迁移。
func (c *DdsClient) MigrateAz(request *model.MigrateAzRequest) (*model.MigrateAzResponse, error) {
	requestDef := GenReqDefForMigrateAz()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateAzResponse), nil
	}
}

//修改数据库用户密码。
func (c *DdsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

//变更实例的规格。
func (c *DdsClient) ResizeInstance(request *model.ResizeInstanceRequest) (*model.ResizeInstanceResponse, error) {
	requestDef := GenReqDefForResizeInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceResponse), nil
	}
}

//扩容实例相关的存储容量大小。
func (c *DdsClient) ResizeInstanceVolume(request *model.ResizeInstanceVolumeRequest) (*model.ResizeInstanceVolumeResponse, error) {
	requestDef := GenReqDefForResizeInstanceVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeInstanceVolumeResponse), nil
	}
}

//重启实例的数据库服务。
func (c *DdsClient) RestartInstance(request *model.RestartInstanceRequest) (*model.RestartInstanceResponse, error) {
	requestDef := GenReqDefForRestartInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartInstanceResponse), nil
	}
}

//恢复到当前实例。
func (c *DdsClient) RestoreInstance(request *model.RestoreInstanceRequest) (*model.RestoreInstanceResponse, error) {
	requestDef := GenReqDefForRestoreInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceResponse), nil
	}
}

//库表级时间点恢复。
func (c *DdsClient) RestoreInstanceFromCollection(request *model.RestoreInstanceFromCollectionRequest) (*model.RestoreInstanceFromCollectionResponse, error) {
	requestDef := GenReqDefForRestoreInstanceFromCollection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreInstanceFromCollectionResponse), nil
	}
}

//根据备份恢复新实例。
func (c *DdsClient) RestoreNewInstance(request *model.RestoreNewInstanceRequest) (*model.RestoreNewInstanceResponse, error) {
	requestDef := GenReqDefForRestoreNewInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreNewInstanceResponse), nil
	}
}

//设置审计日志策略。
func (c *DdsClient) SetAuditlogPolicy(request *model.SetAuditlogPolicyRequest) (*model.SetAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForSetAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetAuditlogPolicyResponse), nil
	}
}

//设置自动备份策略。
func (c *DdsClient) SetBackupPolicy(request *model.SetBackupPolicyRequest) (*model.SetBackupPolicyResponse, error) {
	requestDef := GenReqDefForSetBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBackupPolicyResponse), nil
	}
}

//设置集群均衡开关。
func (c *DdsClient) SetBalancerSwitch(request *model.SetBalancerSwitchRequest) (*model.SetBalancerSwitchResponse, error) {
	requestDef := GenReqDefForSetBalancerSwitch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerSwitchResponse), nil
	}
}

//设置集群均衡活动时间窗。
func (c *DdsClient) SetBalancerWindow(request *model.SetBalancerWindowRequest) (*model.SetBalancerWindowResponse, error) {
	requestDef := GenReqDefForSetBalancerWindow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetBalancerWindowResponse), nil
	}
}

//查询审计日志策略。
func (c *DdsClient) ShowAuditlogPolicy(request *model.ShowAuditlogPolicyRequest) (*model.ShowAuditlogPolicyResponse, error) {
	requestDef := GenReqDefForShowAuditlogPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditlogPolicyResponse), nil
	}
}

//获取备份下载链接。
func (c *DdsClient) ShowBackupDownloadLink(request *model.ShowBackupDownloadLinkRequest) (*model.ShowBackupDownloadLinkResponse, error) {
	requestDef := GenReqDefForShowBackupDownloadLink()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupDownloadLinkResponse), nil
	}
}

//查询自动备份策略。
func (c *DdsClient) ShowBackupPolicy(request *model.ShowBackupPolicyRequest) (*model.ShowBackupPolicyResponse, error) {
	requestDef := GenReqDefForShowBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBackupPolicyResponse), nil
	}
}

//查询客户端IP访问至DDS数据库实例的连接数统计信息。
func (c *DdsClient) ShowConnectionStatistics(request *model.ShowConnectionStatisticsRequest) (*model.ShowConnectionStatisticsResponse, error) {
	requestDef := GenReqDefForShowConnectionStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectionStatisticsResponse), nil
	}
}

//查询集群均衡设置。
func (c *DdsClient) ShowShardingBalancer(request *model.ShowShardingBalancerRequest) (*model.ShowShardingBalancerResponse, error) {
	requestDef := GenReqDefForShowShardingBalancer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowShardingBalancerResponse), nil
	}
}

//切换实例的SSL开关
func (c *DdsClient) SwitchSsl(request *model.SwitchSslRequest) (*model.SwitchSslResponse, error) {
	requestDef := GenReqDefForSwitchSsl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchSslResponse), nil
	}
}

//切换副本集实例下的主备节点
func (c *DdsClient) SwitchoverReplicaSet(request *model.SwitchoverReplicaSetRequest) (*model.SwitchoverReplicaSetResponse, error) {
	requestDef := GenReqDefForSwitchoverReplicaSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchoverReplicaSetResponse), nil
	}
}

//修改实例名称
func (c *DdsClient) UpdateInstanceName(request *model.UpdateInstanceNameRequest) (*model.UpdateInstanceNameResponse, error) {
	requestDef := GenReqDefForUpdateInstanceName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceNameResponse), nil
	}
}

//修改数据库实例的端口。
func (c *DdsClient) UpdateInstancePort(request *model.UpdateInstancePortRequest) (*model.UpdateInstancePortResponse, error) {
	requestDef := GenReqDefForUpdateInstancePort()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstancePortResponse), nil
	}
}

//变更实例关联的安全组
func (c *DdsClient) UpdateSecurityGroup(request *model.UpdateSecurityGroupRequest) (*model.UpdateSecurityGroupResponse, error) {
	requestDef := GenReqDefForUpdateSecurityGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityGroupResponse), nil
	}
}

//查询当前支持的API版本信息列表。
func (c *DdsClient) ListApiVersion(request *model.ListApiVersionRequest) (*model.ListApiVersionResponse, error) {
	requestDef := GenReqDefForListApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionResponse), nil
	}
}

//查询指定API版本信息。
func (c *DdsClient) ShowApiVersion(request *model.ShowApiVersionRequest) (*model.ShowApiVersionResponse, error) {
	requestDef := GenReqDefForShowApiVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiVersionResponse), nil
	}
}
