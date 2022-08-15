package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/gaussdb/v3/model"
)

type GaussDBClient struct {
	HcClient *http_client.HcHttpClient
}

func NewGaussDBClient(hcClient *http_client.HcHttpClient) *GaussDBClient {
	return &GaussDBClient{HcClient: hcClient}
}

func GaussDBClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddDatabasePermission 授予数据库用户数据库权限
//
// 授予云数据库 GaussDB(for MySQL)实例数据库用户数据库权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ChangeGaussMySqlInstanceSpecification 变更实例规格
//
// 变更数据库实例的规格。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateGaussMySqlBackup 创建手动备份
//
// 创建手动备份
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// CreateGaussMySqlDatabase 创建数据库
//
// 创建云数据库 GaussDB(for MySQL)实例数据库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 创建云数据库 GaussDB(for MySQL)实例数据库用户。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 创建云数据库 GaussDB(for MySQL)实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteDatabasePermission 删除数据库用户的数据库权限
//
// 删除云数据库 GaussDB(for MySQL)实例数据库用户的数据库权限
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteGaussMySqlDatabase 删除数据库
//
// 删除云数据库 GaussDB(for MySQL)实例数据库。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteGaussMySqlInstance 删除实例
//
// 删除数据库实例，不支持删除包周期实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlInstance(request *model.DeleteGaussMySqlInstanceRequest) (*model.DeleteGaussMySqlInstanceResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlInstanceResponse), nil
	}
}

// DeleteGaussMySqlInstanceInvoker 删除实例
func (c *GaussDBClient) DeleteGaussMySqlInstanceInvoker(request *model.DeleteGaussMySqlInstanceRequest) *DeleteGaussMySqlInstanceInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlInstance()
	return &DeleteGaussMySqlInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGaussMySqlProxy 关闭数据库代理
//
// 关闭数据库代理。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// DeleteGaussMySqlReadonlyNode 删除只读节点
//
// 删除实例的只读节点。多可用区模式删除只读节点时，要保证删除后，剩余的只读节点和主节点在不同的可用区中，否则无法删除该只读节点。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) DeleteGaussMySqlReadonlyNode(request *model.DeleteGaussMySqlReadonlyNodeRequest) (*model.DeleteGaussMySqlReadonlyNodeResponse, error) {
	requestDef := GenReqDefForDeleteGaussMySqlReadonlyNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGaussMySqlReadonlyNodeResponse), nil
	}
}

// DeleteGaussMySqlReadonlyNodeInvoker 删除只读节点
func (c *GaussDBClient) DeleteGaussMySqlReadonlyNodeInvoker(request *model.DeleteGaussMySqlReadonlyNodeRequest) *DeleteGaussMySqlReadonlyNodeInvoker {
	requestDef := GenReqDefForDeleteGaussMySqlReadonlyNode()
	return &DeleteGaussMySqlReadonlyNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandGaussMySqlInstanceVolume 包周期存储扩容
//
// 包周期存储扩容
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListGaussMySqlConfigurations 查询参数模板
//
// 获取参数模板列表，包括所有数据库的默认参数模板和用户创建的参数模板。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListGaussMySqlErrorLog 查询数据库错误日志
//
// 查询数据库错误日志。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlErrorLog(request *model.ListGaussMySqlErrorLogRequest) (*model.ListGaussMySqlErrorLogResponse, error) {
	requestDef := GenReqDefForListGaussMySqlErrorLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlErrorLogResponse), nil
	}
}

// ListGaussMySqlErrorLogInvoker 查询数据库错误日志
func (c *GaussDBClient) ListGaussMySqlErrorLogInvoker(request *model.ListGaussMySqlErrorLogRequest) *ListGaussMySqlErrorLogInvoker {
	requestDef := GenReqDefForListGaussMySqlErrorLog()
	return &ListGaussMySqlErrorLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGaussMySqlInstances 查询实例列表
//
// 根据指定条件查询实例列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListGaussMySqlSlowLog 查询数据库慢日志
//
// 查询数据库慢日志
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ListGaussMySqlSlowLog(request *model.ListGaussMySqlSlowLogRequest) (*model.ListGaussMySqlSlowLogResponse, error) {
	requestDef := GenReqDefForListGaussMySqlSlowLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGaussMySqlSlowLogResponse), nil
	}
}

// ListGaussMySqlSlowLogInvoker 查询数据库慢日志
func (c *GaussDBClient) ListGaussMySqlSlowLogInvoker(request *model.ListGaussMySqlSlowLogRequest) *ListGaussMySqlSlowLogInvoker {
	requestDef := GenReqDefForListGaussMySqlSlowLog()
	return &ListGaussMySqlSlowLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceTags 查询资源标签
//
// 查询指定实例的标签信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ListProjectTags 查询项目标签
//
// 查询指定project ID下实例的所有标签集合。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ResetGaussMySqlDatabasePassword 修改数据库用户密码
//
// 修改云数据库 GaussDB(for MySQL)实例数据库用户密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 重置数据库密码
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// SetGaussMySqlProxyWeight 设置读写分离权重
//
// 设置读写分离权重
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowAuditLog 查询审计日志开关状态
//
// 查询审计日志开关状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ShowAuditLog(request *model.ShowAuditLogRequest) (*model.ShowAuditLogResponse, error) {
	requestDef := GenReqDefForShowAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuditLogResponse), nil
	}
}

// ShowAuditLogInvoker 查询审计日志开关状态
func (c *GaussDBClient) ShowAuditLogInvoker(request *model.ShowAuditLogRequest) *ShowAuditLogInvoker {
	requestDef := GenReqDefForShowAuditLog()
	return &ShowAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDedicatedResourceInfo 查询专属资源信息详情
//
// 查询专属资源信息详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowGaussMySqlBackupList 查询备份列表
//
// 查询备份列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlBackupList(request *model.ShowGaussMySqlBackupListRequest) (*model.ShowGaussMySqlBackupListResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlBackupList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlBackupListResponse), nil
	}
}

// ShowGaussMySqlBackupListInvoker 查询备份列表
func (c *GaussDBClient) ShowGaussMySqlBackupListInvoker(request *model.ShowGaussMySqlBackupListRequest) *ShowGaussMySqlBackupListInvoker {
	requestDef := GenReqDefForShowGaussMySqlBackupList()
	return &ShowGaussMySqlBackupListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlBackupPolicy 查询自动备份策略
//
// 查询自动备份策略。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowGaussMySqlEngineVersion 查询数据库引擎的版本
//
// 获取指定数据库引擎对应的数据库版本信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowGaussMySqlInstanceInfo 查询实例详情信息
//
// 查询实例详情信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowGaussMySqlJobInfo 获取指定ID的任务信息
//
// 获取GaussDB(for MySQL)任务中心指定ID的任务信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowGaussMySqlProxy 查询数据库代理信息
//
// 查询数据库代理信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ShowGaussMySqlProxy(request *model.ShowGaussMySqlProxyRequest) (*model.ShowGaussMySqlProxyResponse, error) {
	requestDef := GenReqDefForShowGaussMySqlProxy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGaussMySqlProxyResponse), nil
	}
}

// ShowGaussMySqlProxyInvoker 查询数据库代理信息
func (c *GaussDBClient) ShowGaussMySqlProxyInvoker(request *model.ShowGaussMySqlProxyRequest) *ShowGaussMySqlProxyInvoker {
	requestDef := GenReqDefForShowGaussMySqlProxy()
	return &ShowGaussMySqlProxyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGaussMySqlProxyFlavors 查询数据库代理规格信息
//
// 查询数据库代理规格信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// ShowInstanceMonitorExtend 查询实例秒级监控频率
//
// 查询实例秒级监控频率。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) ShowInstanceMonitorExtend(request *model.ShowInstanceMonitorExtendRequest) (*model.ShowInstanceMonitorExtendResponse, error) {
	requestDef := GenReqDefForShowInstanceMonitorExtend()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceMonitorExtendResponse), nil
	}
}

// ShowInstanceMonitorExtendInvoker 查询实例秒级监控频率
func (c *GaussDBClient) ShowInstanceMonitorExtendInvoker(request *model.ShowInstanceMonitorExtendRequest) *ShowInstanceMonitorExtendInvoker {
	requestDef := GenReqDefForShowInstanceMonitorExtend()
	return &ShowInstanceMonitorExtendInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuditLog 开启或者关闭审计日志
//
// 开启或者关闭审计日志
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) UpdateAuditLog(request *model.UpdateAuditLogRequest) (*model.UpdateAuditLogResponse, error) {
	requestDef := GenReqDefForUpdateAuditLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuditLogResponse), nil
	}
}

// UpdateAuditLogInvoker 开启或者关闭审计日志
func (c *GaussDBClient) UpdateAuditLogInvoker(request *model.UpdateAuditLogRequest) *UpdateAuditLogInvoker {
	requestDef := GenReqDefForUpdateAuditLog()
	return &UpdateAuditLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlBackupPolicy 修改备份策略
//
// 修改备份策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) UpdateGaussMySqlBackupPolicy(request *model.UpdateGaussMySqlBackupPolicyRequest) (*model.UpdateGaussMySqlBackupPolicyResponse, error) {
	requestDef := GenReqDefForUpdateGaussMySqlBackupPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGaussMySqlBackupPolicyResponse), nil
	}
}

// UpdateGaussMySqlBackupPolicyInvoker 修改备份策略
func (c *GaussDBClient) UpdateGaussMySqlBackupPolicyInvoker(request *model.UpdateGaussMySqlBackupPolicyRequest) *UpdateGaussMySqlBackupPolicyInvoker {
	requestDef := GenReqDefForUpdateGaussMySqlBackupPolicy()
	return &UpdateGaussMySqlBackupPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGaussMySqlInstanceName 修改实例名称
//
// 修改实例名称
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateGaussMySqlQuotas 修改租户基于企业项目的资源配额
//
// 修改指定企业项目的资源配额。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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

// UpdateInstanceMonitor 修改实例秒级监控频率
//
// 打开/关闭/修改实例秒级监控。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *GaussDBClient) UpdateInstanceMonitor(request *model.UpdateInstanceMonitorRequest) (*model.UpdateInstanceMonitorResponse, error) {
	requestDef := GenReqDefForUpdateInstanceMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInstanceMonitorResponse), nil
	}
}

// UpdateInstanceMonitorInvoker 修改实例秒级监控频率
func (c *GaussDBClient) UpdateInstanceMonitorInvoker(request *model.UpdateInstanceMonitorRequest) *UpdateInstanceMonitorInvoker {
	requestDef := GenReqDefForUpdateInstanceMonitor()
	return &UpdateInstanceMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlFilterRule 删除SQL限流规则
//
// 删除SQL限流规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 设置SQL限流规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 查询SQL限流开关状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 查询SQL限流规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
// 开启或者关闭SQL限流
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
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
