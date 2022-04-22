package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 创建DDM逻辑库
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

// 购买DDM实例
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

// 创建DDM帐号
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

// 删除DDM逻辑库
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

// 删除DDM实例
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

// 删除DDM帐号
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

// DDM实例节点扩容
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

// 查询创建逻辑库可选取的数据库实例列表
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

// 查询DDM逻辑库列表
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

// 查询DDM引擎信息
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

// 查询DDM可用区规格信息
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

// 查询DDM实例列表
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

// 查询DDM实例节点列表
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

// 读写比例监控
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

// 慢日志监控
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

// 查询DDM帐号列表
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

// DDM表数据重载
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

// 重置DDM账号密码
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

// 重启DDM实例
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

// 查询DDM逻辑库详细信息
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

// 查询DDM实例详情
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

// 查询DDM指定实例的参数详情
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

// 查询DDM实例节点详情
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

// DDM实例节点缩容
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

// 同步数据库实例信息
//
// 同步当前DDM实例已关的所有数据库实例配置信息。
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

// 修改DDM实例名称
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

// 修改DDM实例参数
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

// 修改DDM实例安全组
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

// 修改DDM已关联的数据库实例的读策略
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

// 修改DDM帐号
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
