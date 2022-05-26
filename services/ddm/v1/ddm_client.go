package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ddm/v1/model"
)

type DdmClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDdmClient(hcClient *http_client.HcHttpClient) *DdmClient {
	return &DdmClient{HcClient: hcClient}
}

func DdmClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateDatabase 创建DDM逻辑库
//
// 创建DDM逻辑库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateInstance 购买DDM实例
//
// 创建一个DDM实例。
//
// DDM运行于虚拟私有云。申请DDM实例前，需保证有可用的虚拟私有云，并且已配置好子网与安全组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteDatabase 删除DDM逻辑库
//
// 删除指定的逻辑库，释放该逻辑库的所有资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteInstance 删除DDM实例
//
// 删除指定的DDM实例，释放该实例的所有资源。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteUser 删除DDM帐号
//
// 删除指定的DDM实例帐号，如果帐号关联了逻辑库，则对应的关联关系也会删除。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ExpandInstanceNodes DDM实例节点扩容
//
// 对指定的DDM实例的节点个数进行扩容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListAvailableRdsList 查询创建逻辑库可选取的数据库实例列表
//
// 查询创建逻辑库可选取的数据库实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListDatabases 查询DDM逻辑库列表
//
// 查询DDM逻辑库列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListEngines 查询DDM引擎信息
//
// 查询DDM引擎信息详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListInstances 查询DDM实例列表
//
// 查询DDM实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListUsers 查询DDM帐号列表
//
// 查询DDM帐号列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// RebuildConfig DDM表数据重载
//
// DDM实例跨region容灾场景下，针对目标DDM实例实现表数据reload，使数据同步。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ResetUserPassword 重置DDM账号密码
//
// 重置现有DDM帐号的密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// RestartInstance 重启DDM实例
//
// 重启指定的DDM实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowDatabase 查询DDM逻辑库详细信息
//
// 查询指定逻辑库的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowInstance 查询DDM实例详情
//
// 查询指定DDM实例的详细信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowInstanceParam 查询DDM指定实例的参数详情
//
// 查询DDM指定实例的参数详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowNode 查询DDM实例节点详情
//
// 查询DDM实例节点详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShrinkInstanceNodes DDM实例节点缩容
//
// 对指定的DDM实例的节点个数进行缩容。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateDatabaseInfo 同步DN信息
//
// 同步当前DDM实例已关联的所有DN实例配置信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateInstanceSecurityGroup 修改DDM实例安全组
//
// 修改DDM实例安全组。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
