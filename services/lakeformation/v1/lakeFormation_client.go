package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/lakeformation/v1/model"
)

type LakeFormationClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewLakeFormationClient(hcClient *httpclient.HcHttpClient) *LakeFormationClient {
	return &LakeFormationClient{HcClient: hcClient}
}

func LakeFormationClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ApplyForAccess 申请接入服务
//
// 申请接入服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ApplyForAccess(request *model.ApplyForAccessRequest) (*model.ApplyForAccessResponse, error) {
	requestDef := GenReqDefForApplyForAccess()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyForAccessResponse), nil
	}
}

// ApplyForAccessInvoker 申请接入服务
func (c *LakeFormationClient) ApplyForAccessInvoker(request *model.ApplyForAccessRequest) *ApplyForAccessInvoker {
	requestDef := GenReqDefForApplyForAccess()
	return &ApplyForAccessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAuthorizeInterface 批量授权
//
// 批量授权接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchAuthorizeInterface(request *model.BatchAuthorizeInterfaceRequest) (*model.BatchAuthorizeInterfaceResponse, error) {
	requestDef := GenReqDefForBatchAuthorizeInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAuthorizeInterfaceResponse), nil
	}
}

// BatchAuthorizeInterfaceInvoker 批量授权
func (c *LakeFormationClient) BatchAuthorizeInterfaceInvoker(request *model.BatchAuthorizeInterfaceRequest) *BatchAuthorizeInterfaceInvoker {
	requestDef := GenReqDefForBatchAuthorizeInterface()
	return &BatchAuthorizeInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCancelAuthorizationInterface 取消批量授权
//
// 批量取消授权接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchCancelAuthorizationInterface(request *model.BatchCancelAuthorizationInterfaceRequest) (*model.BatchCancelAuthorizationInterfaceResponse, error) {
	requestDef := GenReqDefForBatchCancelAuthorizationInterface()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCancelAuthorizationInterfaceResponse), nil
	}
}

// BatchCancelAuthorizationInterfaceInvoker 取消批量授权
func (c *LakeFormationClient) BatchCancelAuthorizationInterfaceInvoker(request *model.BatchCancelAuthorizationInterfaceRequest) *BatchCancelAuthorizationInterfaceInvoker {
	requestDef := GenReqDefForBatchCancelAuthorizationInterface()
	return &BatchCancelAuthorizationInterfaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckPermission 批量鉴权
//
// 批量鉴权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchCheckPermission(request *model.BatchCheckPermissionRequest) (*model.BatchCheckPermissionResponse, error) {
	requestDef := GenReqDefForBatchCheckPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckPermissionResponse), nil
	}
}

// BatchCheckPermissionInvoker 批量鉴权
func (c *LakeFormationClient) BatchCheckPermissionInvoker(request *model.BatchCheckPermissionRequest) *BatchCheckPermissionInvoker {
	requestDef := GenReqDefForBatchCheckPermission()
	return &BatchCheckPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAccessClient 创建服务接入客户端
//
// 创建服务接入客户端。
// 其他限制：
//   同一个实例下默认最多创建20个接入客户端。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateAccessClient(request *model.CreateAccessClientRequest) (*model.CreateAccessClientResponse, error) {
	requestDef := GenReqDefForCreateAccessClient()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAccessClientResponse), nil
	}
}

// CreateAccessClientInvoker 创建服务接入客户端
func (c *LakeFormationClient) CreateAccessClientInvoker(request *model.CreateAccessClientRequest) *CreateAccessClientInvoker {
	requestDef := GenReqDefForCreateAccessClient()
	return &CreateAccessClientInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccessClient 删除服务接入客户端
//
// 根据ID删除服务接入客户端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteAccessClient(request *model.DeleteAccessClientRequest) (*model.DeleteAccessClientResponse, error) {
	requestDef := GenReqDefForDeleteAccessClient()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessClientResponse), nil
	}
}

// DeleteAccessClientInvoker 删除服务接入客户端
func (c *LakeFormationClient) DeleteAccessClientInvoker(request *model.DeleteAccessClientRequest) *DeleteAccessClientInvoker {
	requestDef := GenReqDefForDeleteAccessClient()
	return &DeleteAccessClientInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessClientInfos 获取服务接入客户端信息列表
//
// 根据LakeFormation实例获取服务实例相关的接入客户端信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListAccessClientInfos(request *model.ListAccessClientInfosRequest) (*model.ListAccessClientInfosResponse, error) {
	requestDef := GenReqDefForListAccessClientInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessClientInfosResponse), nil
	}
}

// ListAccessClientInfosInvoker 获取服务接入客户端信息列表
func (c *LakeFormationClient) ListAccessClientInfosInvoker(request *model.ListAccessClientInfosRequest) *ListAccessClientInfosInvoker {
	requestDef := GenReqDefForListAccessClientInfos()
	return &ListAccessClientInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessInfos 获取服务接入信息
//
// 根据LakeFormation实例获取服务实例相关的接入信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListAccessInfos(request *model.ListAccessInfosRequest) (*model.ListAccessInfosResponse, error) {
	requestDef := GenReqDefForListAccessInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessInfosResponse), nil
	}
}

// ListAccessInfosInvoker 获取服务接入信息
func (c *LakeFormationClient) ListAccessInfosInvoker(request *model.ListAccessInfosRequest) *ListAccessInfosInvoker {
	requestDef := GenReqDefForListAccessInfos()
	return &ListAccessInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInterfaces 查询策略
//
// 通过过滤条件查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListInterfaces(request *model.ListInterfacesRequest) (*model.ListInterfacesResponse, error) {
	requestDef := GenReqDefForListInterfaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInterfacesResponse), nil
	}
}

// ListInterfacesInvoker 查询策略
func (c *LakeFormationClient) ListInterfacesInvoker(request *model.ListInterfacesRequest) *ListInterfacesInvoker {
	requestDef := GenReqDefForListInterfaces()
	return &ListInterfacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPolicy 获取同步权限策略
//
// 分页获取同步权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListPolicy(request *model.ListPolicyRequest) (*model.ListPolicyResponse, error) {
	requestDef := GenReqDefForListPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyResponse), nil
	}
}

// ListPolicyInvoker 获取同步权限策略
func (c *LakeFormationClient) ListPolicyInvoker(request *model.ListPolicyRequest) *ListPolicyInvoker {
	requestDef := GenReqDefForListPolicy()
	return &ListPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessClient 获取服务接入客户端详情
//
// 根据ID获取服务接入客户端详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowAccessClient(request *model.ShowAccessClientRequest) (*model.ShowAccessClientResponse, error) {
	requestDef := GenReqDefForShowAccessClient()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessClientResponse), nil
	}
}

// ShowAccessClientInvoker 获取服务接入客户端详情
func (c *LakeFormationClient) ShowAccessClientInvoker(request *model.ShowAccessClientRequest) *ShowAccessClientInvoker {
	requestDef := GenReqDefForShowAccessClient()
	return &ShowAccessClientInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSyncPolicy 获取同步权限策略
//
// 获取同步权限策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowSyncPolicy(request *model.ShowSyncPolicyRequest) (*model.ShowSyncPolicyResponse, error) {
	requestDef := GenReqDefForShowSyncPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSyncPolicyResponse), nil
	}
}

// ShowSyncPolicyInvoker 获取同步权限策略
func (c *LakeFormationClient) ShowSyncPolicyInvoker(request *model.ShowSyncPolicyRequest) *ShowSyncPolicyInvoker {
	requestDef := GenReqDefForShowSyncPolicy()
	return &ShowSyncPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccessClient 更新服务接入客户端
//
// 根据ID更新服务接入客户端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateAccessClient(request *model.UpdateAccessClientRequest) (*model.UpdateAccessClientResponse, error) {
	requestDef := GenReqDefForUpdateAccessClient()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccessClientResponse), nil
	}
}

// UpdateAccessClientInvoker 更新服务接入客户端
func (c *LakeFormationClient) UpdateAccessClientInvoker(request *model.UpdateAccessClientRequest) *UpdateAccessClientInvoker {
	requestDef := GenReqDefForUpdateAccessClient()
	return &UpdateAccessClientInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgency 创建委托
//
// 创建委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateAgency(request *model.CreateAgencyRequest) (*model.CreateAgencyResponse, error) {
	requestDef := GenReqDefForCreateAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgencyResponse), nil
	}
}

// CreateAgencyInvoker 创建委托
func (c *LakeFormationClient) CreateAgencyInvoker(request *model.CreateAgencyRequest) *CreateAgencyInvoker {
	requestDef := GenReqDefForCreateAgency()
	return &CreateAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgency 删除委托
//
// 删除委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteAgency(request *model.DeleteAgencyRequest) (*model.DeleteAgencyResponse, error) {
	requestDef := GenReqDefForDeleteAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgencyResponse), nil
	}
}

// DeleteAgencyInvoker 删除委托
func (c *LakeFormationClient) DeleteAgencyInvoker(request *model.DeleteAgencyRequest) *DeleteAgencyInvoker {
	requestDef := GenReqDefForDeleteAgency()
	return &DeleteAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgency 委托查询
//
// 委托查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowAgency(request *model.ShowAgencyRequest) (*model.ShowAgencyResponse, error) {
	requestDef := GenReqDefForShowAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgencyResponse), nil
	}
}

// ShowAgencyInvoker 委托查询
func (c *LakeFormationClient) ShowAgencyInvoker(request *model.ShowAgencyRequest) *ShowAgencyInvoker {
	requestDef := GenReqDefForShowAgency()
	return &ShowAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCatalog 创建catalog
//
// 创建catalog，会在catalog下创建默认数据库，默认数据库名称为：default
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateCatalog(request *model.CreateCatalogRequest) (*model.CreateCatalogResponse, error) {
	requestDef := GenReqDefForCreateCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCatalogResponse), nil
	}
}

// CreateCatalogInvoker 创建catalog
func (c *LakeFormationClient) CreateCatalogInvoker(request *model.CreateCatalogRequest) *CreateCatalogInvoker {
	requestDef := GenReqDefForCreateCatalog()
	return &CreateCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCatalog 删除catalog对象
//
// 删除空的catalog对象。
// 删除前需要保证catalog下只有默认的数据库，且默认数据库下没有表对象，否则删除失败。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteCatalog(request *model.DeleteCatalogRequest) (*model.DeleteCatalogResponse, error) {
	requestDef := GenReqDefForDeleteCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCatalogResponse), nil
	}
}

// DeleteCatalogInvoker 删除catalog对象
func (c *LakeFormationClient) DeleteCatalogInvoker(request *model.DeleteCatalogRequest) *DeleteCatalogInvoker {
	requestDef := GenReqDefForDeleteCatalog()
	return &DeleteCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogs 列举catalog信息
//
// 根据catalog名字的通配符列举catalog的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListCatalogs(request *model.ListCatalogsRequest) (*model.ListCatalogsResponse, error) {
	requestDef := GenReqDefForListCatalogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogsResponse), nil
	}
}

// ListCatalogsInvoker 列举catalog信息
func (c *LakeFormationClient) ListCatalogsInvoker(request *model.ListCatalogsRequest) *ListCatalogsInvoker {
	requestDef := GenReqDefForListCatalogs()
	return &ListCatalogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCatalog 获取catalog信息
//
// 获取catalog信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowCatalog(request *model.ShowCatalogRequest) (*model.ShowCatalogResponse, error) {
	requestDef := GenReqDefForShowCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCatalogResponse), nil
	}
}

// ShowCatalogInvoker 获取catalog信息
func (c *LakeFormationClient) ShowCatalogInvoker(request *model.ShowCatalogRequest) *ShowCatalogInvoker {
	requestDef := GenReqDefForShowCatalog()
	return &ShowCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCatalog 修改catalog信息
//
// 修改catalog信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateCatalog(request *model.UpdateCatalogRequest) (*model.UpdateCatalogResponse, error) {
	requestDef := GenReqDefForUpdateCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCatalogResponse), nil
	}
}

// UpdateCatalogInvoker 修改catalog信息
func (c *LakeFormationClient) UpdateCatalogInvoker(request *model.UpdateCatalogRequest) *UpdateCatalogInvoker {
	requestDef := GenReqDefForUpdateCatalog()
	return &UpdateCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigs 获取所有用户可见的租户面配置
//
// 获取所有用户可见的租户面配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListConfigs(request *model.ListConfigsRequest) (*model.ListConfigsResponse, error) {
	requestDef := GenReqDefForListConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigsResponse), nil
	}
}

// ListConfigsInvoker 获取所有用户可见的租户面配置
func (c *LakeFormationClient) ListConfigsInvoker(request *model.ListConfigsRequest) *ListConfigsInvoker {
	requestDef := GenReqDefForListConfigs()
	return &ListConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCredential 获取临时密钥和securityToken
//
// 获取临时密钥和securityToken，失效时间大于等于1小时，请在1小时内更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowCredential(request *model.ShowCredentialRequest) (*model.ShowCredentialResponse, error) {
	requestDef := GenReqDefForShowCredential()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCredentialResponse), nil
	}
}

// ShowCredentialInvoker 获取临时密钥和securityToken
func (c *LakeFormationClient) ShowCredentialInvoker(request *model.ShowCredentialRequest) *ShowCredentialInvoker {
	requestDef := GenReqDefForShowCredential()
	return &ShowCredentialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabase 创建数据库
//
// 创建数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateDatabase(request *model.CreateDatabaseRequest) (*model.CreateDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResponse), nil
	}
}

// CreateDatabaseInvoker 创建数据库
func (c *LakeFormationClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabase 删除数据库
//
// 删除指定数据库，catalog的默认数据库不允许删除。
// cascade: 指定为true时，删除数据库下的表；指定为false时，只能删除空的数据库
// delete_data: 指定为true时，级联删除会将表的数据放入回收站；指定为false时，不删除表数据
// 删除数据库后不支持恢复数据库下的事务表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteDatabase(request *model.DeleteDatabaseRequest) (*model.DeleteDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResponse), nil
	}
}

// DeleteDatabaseInvoker 删除数据库
func (c *LakeFormationClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseNames 列举数据库名称信息
//
// 列举数据库名称信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListDatabaseNames(request *model.ListDatabaseNamesRequest) (*model.ListDatabaseNamesResponse, error) {
	requestDef := GenReqDefForListDatabaseNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseNamesResponse), nil
	}
}

// ListDatabaseNamesInvoker 列举数据库名称信息
func (c *LakeFormationClient) ListDatabaseNamesInvoker(request *model.ListDatabaseNamesRequest) *ListDatabaseNamesInvoker {
	requestDef := GenReqDefForListDatabaseNames()
	return &ListDatabaseNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabases 列举数据库信息
//
// 列举数据库信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 列举数据库信息
func (c *LakeFormationClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatabase 获取数据库
//
// 获取数据库
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowDatabase(request *model.ShowDatabaseRequest) (*model.ShowDatabaseResponse, error) {
	requestDef := GenReqDefForShowDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseResponse), nil
	}
}

// ShowDatabaseInvoker 获取数据库
func (c *LakeFormationClient) ShowDatabaseInvoker(request *model.ShowDatabaseRequest) *ShowDatabaseInvoker {
	requestDef := GenReqDefForShowDatabase()
	return &ShowDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatabase 修改数据库属性
//
// 修改数据库属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateDatabase(request *model.UpdateDatabaseRequest) (*model.UpdateDatabaseResponse, error) {
	requestDef := GenReqDefForUpdateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseResponse), nil
	}
}

// UpdateDatabaseInvoker 修改数据库属性
func (c *LakeFormationClient) UpdateDatabaseInvoker(request *model.UpdateDatabaseRequest) *UpdateDatabaseInvoker {
	requestDef := GenReqDefForUpdateDatabase()
	return &UpdateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFunction 创建函数
//
// 创建函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateFunction(request *model.CreateFunctionRequest) (*model.CreateFunctionResponse, error) {
	requestDef := GenReqDefForCreateFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFunctionResponse), nil
	}
}

// CreateFunctionInvoker 创建函数
func (c *LakeFormationClient) CreateFunctionInvoker(request *model.CreateFunctionRequest) *CreateFunctionInvoker {
	requestDef := GenReqDefForCreateFunction()
	return &CreateFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFunction 删除函数
//
// 删除函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteFunction(request *model.DeleteFunctionRequest) (*model.DeleteFunctionResponse, error) {
	requestDef := GenReqDefForDeleteFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFunctionResponse), nil
	}
}

// DeleteFunctionInvoker 删除函数
func (c *LakeFormationClient) DeleteFunctionInvoker(request *model.DeleteFunctionRequest) *DeleteFunctionInvoker {
	requestDef := GenReqDefForDeleteFunction()
	return &DeleteFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllFunctions 列举catalog下所有的函数
//
// 获取catalog下所有的函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListAllFunctions(request *model.ListAllFunctionsRequest) (*model.ListAllFunctionsResponse, error) {
	requestDef := GenReqDefForListAllFunctions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllFunctionsResponse), nil
	}
}

// ListAllFunctionsInvoker 列举catalog下所有的函数
func (c *LakeFormationClient) ListAllFunctionsInvoker(request *model.ListAllFunctionsRequest) *ListAllFunctionsInvoker {
	requestDef := GenReqDefForListAllFunctions()
	return &ListAllFunctionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctionNames 列举库下所有函数名称
//
// 查询数据库下的所有函数名称列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListFunctionNames(request *model.ListFunctionNamesRequest) (*model.ListFunctionNamesResponse, error) {
	requestDef := GenReqDefForListFunctionNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionNamesResponse), nil
	}
}

// ListFunctionNamesInvoker 列举库下所有函数名称
func (c *LakeFormationClient) ListFunctionNamesInvoker(request *model.ListFunctionNamesRequest) *ListFunctionNamesInvoker {
	requestDef := GenReqDefForListFunctionNames()
	return &ListFunctionNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFunctions 列举函数
//
// 列举函数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListFunctions(request *model.ListFunctionsRequest) (*model.ListFunctionsResponse, error) {
	requestDef := GenReqDefForListFunctions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFunctionsResponse), nil
	}
}

// ListFunctionsInvoker 列举函数
func (c *LakeFormationClient) ListFunctionsInvoker(request *model.ListFunctionsRequest) *ListFunctionsInvoker {
	requestDef := GenReqDefForListFunctions()
	return &ListFunctionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFunction 查询函数
//
// 根据函数名称查询函数信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowFunction(request *model.ShowFunctionRequest) (*model.ShowFunctionResponse, error) {
	requestDef := GenReqDefForShowFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFunctionResponse), nil
	}
}

// ShowFunctionInvoker 查询函数
func (c *LakeFormationClient) ShowFunctionInvoker(request *model.ShowFunctionRequest) *ShowFunctionInvoker {
	requestDef := GenReqDefForShowFunction()
	return &ShowFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFunction 修改函数属性
//
// 修改函数属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateFunction(request *model.UpdateFunctionRequest) (*model.UpdateFunctionResponse, error) {
	requestDef := GenReqDefForUpdateFunction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFunctionResponse), nil
	}
}

// UpdateFunctionInvoker 修改函数属性
func (c *LakeFormationClient) UpdateFunctionInvoker(request *model.UpdateFunctionRequest) *UpdateFunctionInvoker {
	requestDef := GenReqDefForUpdateFunction()
	return &UpdateFunctionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeAccessService 接入服务授权
//
// 接入服务授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) AuthorizeAccessService(request *model.AuthorizeAccessServiceRequest) (*model.AuthorizeAccessServiceResponse, error) {
	requestDef := GenReqDefForAuthorizeAccessService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeAccessServiceResponse), nil
	}
}

// AuthorizeAccessServiceInvoker 接入服务授权
func (c *LakeFormationClient) AuthorizeAccessServiceInvoker(request *model.AuthorizeAccessServiceRequest) *AuthorizeAccessServiceInvoker {
	requestDef := GenReqDefForAuthorizeAccessService()
	return &AuthorizeAccessServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAgreement 注册租户协议
//
// 用户授权并委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateAgreement(request *model.CreateAgreementRequest) (*model.CreateAgreementResponse, error) {
	requestDef := GenReqDefForCreateAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAgreementResponse), nil
	}
}

// CreateAgreementInvoker 注册租户协议
func (c *LakeFormationClient) CreateAgreementInvoker(request *model.CreateAgreementRequest) *CreateAgreementInvoker {
	requestDef := GenReqDefForCreateAgreement()
	return &CreateAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAgreement 删除租户协议
//
// 用户取消授权，同时有权限用户删除委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteAgreement(request *model.DeleteAgreementRequest) (*model.DeleteAgreementResponse, error) {
	requestDef := GenReqDefForDeleteAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAgreementResponse), nil
	}
}

// DeleteAgreementInvoker 删除租户协议
func (c *LakeFormationClient) DeleteAgreementInvoker(request *model.DeleteAgreementRequest) *DeleteAgreementInvoker {
	requestDef := GenReqDefForDeleteAgreement()
	return &DeleteAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessService 查询租户当前的接入服务授权
//
// 查询租户当前的接入服务授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowAccessService(request *model.ShowAccessServiceRequest) (*model.ShowAccessServiceResponse, error) {
	requestDef := GenReqDefForShowAccessService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessServiceResponse), nil
	}
}

// ShowAccessServiceInvoker 查询租户当前的接入服务授权
func (c *LakeFormationClient) ShowAccessServiceInvoker(request *model.ShowAccessServiceRequest) *ShowAccessServiceInvoker {
	requestDef := GenReqDefForShowAccessService()
	return &ShowAccessServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgreement 查询租户是否注册协议
//
// 查询租户当前协议和委托信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowAgreement(request *model.ShowAgreementRequest) (*model.ShowAgreementResponse, error) {
	requestDef := GenReqDefForShowAgreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgreementResponse), nil
	}
}

// ShowAgreementInvoker 查询租户是否注册协议
func (c *LakeFormationClient) ShowAgreementInvoker(request *model.ShowAgreementRequest) *ShowAgreementInvoker {
	requestDef := GenReqDefForShowAgreement()
	return &ShowAgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAgreementRule 查询当前系统协议
//
// 查询当前系统协议
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowAgreementRule(request *model.ShowAgreementRuleRequest) (*model.ShowAgreementRuleResponse, error) {
	requestDef := GenReqDefForShowAgreementRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAgreementRuleResponse), nil
	}
}

// ShowAgreementRuleInvoker 查询当前系统协议
func (c *LakeFormationClient) ShowAgreementRuleInvoker(request *model.ShowAgreementRuleRequest) *ShowAgreementRuleInvoker {
	requestDef := GenReqDefForShowAgreementRule()
	return &ShowAgreementRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountMetaObj 元数据数量统计
//
// 元数据数量统计接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CountMetaObj(request *model.CountMetaObjRequest) (*model.CountMetaObjResponse, error) {
	requestDef := GenReqDefForCountMetaObj()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountMetaObjResponse), nil
	}
}

// CountMetaObjInvoker 元数据数量统计
func (c *LakeFormationClient) CountMetaObjInvoker(request *model.CountMetaObjRequest) *CountMetaObjInvoker {
	requestDef := GenReqDefForCountMetaObj()
	return &CountMetaObjInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateLakeFormationInstance 创建实例
//
// 创建一个LakeFormation实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateLakeFormationInstance(request *model.CreateLakeFormationInstanceRequest) (*model.CreateLakeFormationInstanceResponse, error) {
	requestDef := GenReqDefForCreateLakeFormationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateLakeFormationInstanceResponse), nil
	}
}

// CreateLakeFormationInstanceInvoker 创建实例
func (c *LakeFormationClient) CreateLakeFormationInstanceInvoker(request *model.CreateLakeFormationInstanceRequest) *CreateLakeFormationInstanceInvoker {
	requestDef := GenReqDefForCreateLakeFormationInstance()
	return &CreateLakeFormationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLakeFormationInstance 删除实例
//
// 根据实例Id删除LakeFormation实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteLakeFormationInstance(request *model.DeleteLakeFormationInstanceRequest) (*model.DeleteLakeFormationInstanceResponse, error) {
	requestDef := GenReqDefForDeleteLakeFormationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLakeFormationInstanceResponse), nil
	}
}

// DeleteLakeFormationInstanceInvoker 删除实例
func (c *LakeFormationClient) DeleteLakeFormationInstanceInvoker(request *model.DeleteLakeFormationInstanceRequest) *DeleteLakeFormationInstanceInvoker {
	requestDef := GenReqDefForDeleteLakeFormationInstance()
	return &DeleteLakeFormationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLakeFormationInstances 查询实例列表
//
// 查询用户创建的实例列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListLakeFormationInstances(request *model.ListLakeFormationInstancesRequest) (*model.ListLakeFormationInstancesResponse, error) {
	requestDef := GenReqDefForListLakeFormationInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLakeFormationInstancesResponse), nil
	}
}

// ListLakeFormationInstancesInvoker 查询实例列表
func (c *LakeFormationClient) ListLakeFormationInstancesInvoker(request *model.ListLakeFormationInstancesRequest) *ListLakeFormationInstancesInvoker {
	requestDef := GenReqDefForListLakeFormationInstances()
	return &ListLakeFormationInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MoveLakeFormationInstanceOutRecycleBin 恢复实例
//
// 从回收站恢复LakeFormation实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) MoveLakeFormationInstanceOutRecycleBin(request *model.MoveLakeFormationInstanceOutRecycleBinRequest) (*model.MoveLakeFormationInstanceOutRecycleBinResponse, error) {
	requestDef := GenReqDefForMoveLakeFormationInstanceOutRecycleBin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MoveLakeFormationInstanceOutRecycleBinResponse), nil
	}
}

// MoveLakeFormationInstanceOutRecycleBinInvoker 恢复实例
func (c *LakeFormationClient) MoveLakeFormationInstanceOutRecycleBinInvoker(request *model.MoveLakeFormationInstanceOutRecycleBinRequest) *MoveLakeFormationInstanceOutRecycleBinInvoker {
	requestDef := GenReqDefForMoveLakeFormationInstanceOutRecycleBin()
	return &MoveLakeFormationInstanceOutRecycleBinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLakeFormationInstance 查询实例详情
//
// 使用实例Id查询LakeFormation实例详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowLakeFormationInstance(request *model.ShowLakeFormationInstanceRequest) (*model.ShowLakeFormationInstanceResponse, error) {
	requestDef := GenReqDefForShowLakeFormationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLakeFormationInstanceResponse), nil
	}
}

// ShowLakeFormationInstanceInvoker 查询实例详情
func (c *LakeFormationClient) ShowLakeFormationInstanceInvoker(request *model.ShowLakeFormationInstanceRequest) *ShowLakeFormationInstanceInvoker {
	requestDef := GenReqDefForShowLakeFormationInstance()
	return &ShowLakeFormationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLakeFormationInstance 更新实例
//
// 修改LakeFormation实例信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateLakeFormationInstance(request *model.UpdateLakeFormationInstanceRequest) (*model.UpdateLakeFormationInstanceResponse, error) {
	requestDef := GenReqDefForUpdateLakeFormationInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLakeFormationInstanceResponse), nil
	}
}

// UpdateLakeFormationInstanceInvoker 更新实例
func (c *LakeFormationClient) UpdateLakeFormationInstanceInvoker(request *model.UpdateLakeFormationInstanceRequest) *UpdateLakeFormationInstanceInvoker {
	requestDef := GenReqDefForUpdateLakeFormationInstance()
	return &UpdateLakeFormationInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLakeFormationInstanceDefault 设为默认实例
//
// 设为默认实例，只有非默认实例可以设置为默认实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateLakeFormationInstanceDefault(request *model.UpdateLakeFormationInstanceDefaultRequest) (*model.UpdateLakeFormationInstanceDefaultResponse, error) {
	requestDef := GenReqDefForUpdateLakeFormationInstanceDefault()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLakeFormationInstanceDefaultResponse), nil
	}
}

// UpdateLakeFormationInstanceDefaultInvoker 设为默认实例
func (c *LakeFormationClient) UpdateLakeFormationInstanceDefaultInvoker(request *model.UpdateLakeFormationInstanceDefaultRequest) *UpdateLakeFormationInstanceDefaultInvoker {
	requestDef := GenReqDefForUpdateLakeFormationInstanceDefault()
	return &UpdateLakeFormationInstanceDefaultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateLakeFormationInstanceScale 变更实例规格
//
// 变更LakeFormation实例规格
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateLakeFormationInstanceScale(request *model.UpdateLakeFormationInstanceScaleRequest) (*model.UpdateLakeFormationInstanceScaleResponse, error) {
	requestDef := GenReqDefForUpdateLakeFormationInstanceScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateLakeFormationInstanceScaleResponse), nil
	}
}

// UpdateLakeFormationInstanceScaleInvoker 变更实例规格
func (c *LakeFormationClient) UpdateLakeFormationInstanceScaleInvoker(request *model.UpdateLakeFormationInstanceScaleRequest) *UpdateLakeFormationInstanceScaleInvoker {
	requestDef := GenReqDefForUpdateLakeFormationInstanceScale()
	return &UpdateLakeFormationInstanceScaleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsBuckets 查询OBS桶列表
//
// 查询OBS桶列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListObsBuckets(request *model.ListObsBucketsRequest) (*model.ListObsBucketsResponse, error) {
	requestDef := GenReqDefForListObsBuckets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsBucketsResponse), nil
	}
}

// ListObsBucketsInvoker 查询OBS桶列表
func (c *LakeFormationClient) ListObsBucketsInvoker(request *model.ListObsBucketsRequest) *ListObsBucketsInvoker {
	requestDef := GenReqDefForListObsBuckets()
	return &ListObsBucketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListObsObject 查询obs桶对象列表
//
// 查询obs桶对象列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListObsObject(request *model.ListObsObjectRequest) (*model.ListObsObjectResponse, error) {
	requestDef := GenReqDefForListObsObject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListObsObjectResponse), nil
	}
}

// ListObsObjectInvoker 查询obs桶对象列表
func (c *LakeFormationClient) ListObsObjectInvoker(request *model.ListObsObjectRequest) *ListObsObjectInvoker {
	requestDef := GenReqDefForListObsObject()
	return &ListObsObjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddPartitions 批量添加分区信息
//
// 批量添加分区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) AddPartitions(request *model.AddPartitionsRequest) (*model.AddPartitionsResponse, error) {
	requestDef := GenReqDefForAddPartitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddPartitionsResponse), nil
	}
}

// AddPartitionsInvoker 批量添加分区信息
func (c *LakeFormationClient) AddPartitionsInvoker(request *model.AddPartitionsRequest) *AddPartitionsInvoker {
	requestDef := GenReqDefForAddPartitions()
	return &AddPartitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePartition 批量删除分区信息
//
// 非事务表：如果设置删除数据，立刻删除分区数据路径下的数据。
// 事务表：如果设置删除数据，保留数据在原路径下但对外不可见，待数据超期后统一删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchDeletePartition(request *model.BatchDeletePartitionRequest) (*model.BatchDeletePartitionResponse, error) {
	requestDef := GenReqDefForBatchDeletePartition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePartitionResponse), nil
	}
}

// BatchDeletePartitionInvoker 批量删除分区信息
func (c *LakeFormationClient) BatchDeletePartitionInvoker(request *model.BatchDeletePartitionRequest) *BatchDeletePartitionInvoker {
	requestDef := GenReqDefForBatchDeletePartition()
	return &BatchDeletePartitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePartitionedStatistics 批量清空列表信息
//
// 删除存在分区的数据及统计信息，保留分区的元数据信息。全部存在、部分存在或都不存在，均返回OK
// 非事务表：立刻删除分区路径下的数据。
// 事务表：数据保留但不可见，待被删除数据超期后统一删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchDeletePartitionedStatistics(request *model.BatchDeletePartitionedStatisticsRequest) (*model.BatchDeletePartitionedStatisticsResponse, error) {
	requestDef := GenReqDefForBatchDeletePartitionedStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePartitionedStatisticsResponse), nil
	}
}

// BatchDeletePartitionedStatisticsInvoker 批量清空列表信息
func (c *LakeFormationClient) BatchDeletePartitionedStatisticsInvoker(request *model.BatchDeletePartitionedStatisticsRequest) *BatchDeletePartitionedStatisticsInvoker {
	requestDef := GenReqDefForBatchDeletePartitionedStatistics()
	return &BatchDeletePartitionedStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchListPartitionByValues 批量获取分区信息
//
// 批量获取分区信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchListPartitionByValues(request *model.BatchListPartitionByValuesRequest) (*model.BatchListPartitionByValuesResponse, error) {
	requestDef := GenReqDefForBatchListPartitionByValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListPartitionByValuesResponse), nil
	}
}

// BatchListPartitionByValuesInvoker 批量获取分区信息
func (c *LakeFormationClient) BatchListPartitionByValuesInvoker(request *model.BatchListPartitionByValuesRequest) *BatchListPartitionByValuesInvoker {
	requestDef := GenReqDefForBatchListPartitionByValues()
	return &BatchListPartitionByValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdatePartition 批量修改分区信息
//
// 所有partition必须要全部存在，如果存在某个partition不存在，就返回失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchUpdatePartition(request *model.BatchUpdatePartitionRequest) (*model.BatchUpdatePartitionResponse, error) {
	requestDef := GenReqDefForBatchUpdatePartition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePartitionResponse), nil
	}
}

// BatchUpdatePartitionInvoker 批量修改分区信息
func (c *LakeFormationClient) BatchUpdatePartitionInvoker(request *model.BatchUpdatePartitionRequest) *BatchUpdatePartitionInvoker {
	requestDef := GenReqDefForBatchUpdatePartition()
	return &BatchUpdatePartitionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartitionNames 列举分区值列表
//
// 遍历分区名字列表信息
// 对于事务表，支持基于表的特定版本遍历分区名字列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListPartitionNames(request *model.ListPartitionNamesRequest) (*model.ListPartitionNamesResponse, error) {
	requestDef := GenReqDefForListPartitionNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartitionNamesResponse), nil
	}
}

// ListPartitionNamesInvoker 列举分区值列表
func (c *LakeFormationClient) ListPartitionNamesInvoker(request *model.ListPartitionNamesRequest) *ListPartitionNamesInvoker {
	requestDef := GenReqDefForListPartitionNames()
	return &ListPartitionNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartitionNamesWithoutLimit 列举全量分区值列表
//
// 遍历分区名称列表信息,返回全量的数据。
// 对于事务表，支持基于表的特定版本遍历分区名称列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListPartitionNamesWithoutLimit(request *model.ListPartitionNamesWithoutLimitRequest) (*model.ListPartitionNamesWithoutLimitResponse, error) {
	requestDef := GenReqDefForListPartitionNamesWithoutLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartitionNamesWithoutLimitResponse), nil
	}
}

// ListPartitionNamesWithoutLimitInvoker 列举全量分区值列表
func (c *LakeFormationClient) ListPartitionNamesWithoutLimitInvoker(request *model.ListPartitionNamesWithoutLimitRequest) *ListPartitionNamesWithoutLimitInvoker {
	requestDef := GenReqDefForListPartitionNamesWithoutLimit()
	return &ListPartitionNamesWithoutLimitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPartitions 列举分区信息
//
// 遍历指定数据表下的分区列表，对于事务表，支持基于表的特定版本遍历分区列表。
// 当过滤条件不为空时，优先根据过滤条件筛选过滤，
// 当过滤条件为空且分区值不为空时，再根据分区值筛选过滤，
// 当过滤条件和分区值都为空时，返回指定数据表下所有分区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListPartitions(request *model.ListPartitionsRequest) (*model.ListPartitionsResponse, error) {
	requestDef := GenReqDefForListPartitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartitionsResponse), nil
	}
}

// ListPartitionsInvoker 列举分区信息
func (c *LakeFormationClient) ListPartitionsInvoker(request *model.ListPartitionsRequest) *ListPartitionsInvoker {
	requestDef := GenReqDefForListPartitions()
	return &ListPartitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowPartitionColumnStatistics 批量获取分区的列统计信息
//
// 批量获取分区的列统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchShowPartitionColumnStatistics(request *model.BatchShowPartitionColumnStatisticsRequest) (*model.BatchShowPartitionColumnStatisticsResponse, error) {
	requestDef := GenReqDefForBatchShowPartitionColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowPartitionColumnStatisticsResponse), nil
	}
}

// BatchShowPartitionColumnStatisticsInvoker 批量获取分区的列统计信息
func (c *LakeFormationClient) BatchShowPartitionColumnStatisticsInvoker(request *model.BatchShowPartitionColumnStatisticsRequest) *BatchShowPartitionColumnStatisticsInvoker {
	requestDef := GenReqDefForBatchShowPartitionColumnStatistics()
	return &BatchShowPartitionColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePartitionColumnStatistics 删除分区列的统计信息
//
// 删除分区列的统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeletePartitionColumnStatistics(request *model.DeletePartitionColumnStatisticsRequest) (*model.DeletePartitionColumnStatisticsResponse, error) {
	requestDef := GenReqDefForDeletePartitionColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePartitionColumnStatisticsResponse), nil
	}
}

// DeletePartitionColumnStatisticsInvoker 删除分区列的统计信息
func (c *LakeFormationClient) DeletePartitionColumnStatisticsInvoker(request *model.DeletePartitionColumnStatisticsRequest) *DeletePartitionColumnStatisticsInvoker {
	requestDef := GenReqDefForDeletePartitionColumnStatistics()
	return &DeletePartitionColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetPartitionColumnStatistics 批量设置分区的统计信息
//
// 批量设置分区的统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) SetPartitionColumnStatistics(request *model.SetPartitionColumnStatisticsRequest) (*model.SetPartitionColumnStatisticsResponse, error) {
	requestDef := GenReqDefForSetPartitionColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetPartitionColumnStatisticsResponse), nil
	}
}

// SetPartitionColumnStatisticsInvoker 批量设置分区的统计信息
func (c *LakeFormationClient) SetPartitionColumnStatisticsInvoker(request *model.SetPartitionColumnStatisticsRequest) *SetPartitionColumnStatisticsInvoker {
	requestDef := GenReqDefForSetPartitionColumnStatistics()
	return &SetPartitionColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询用户的配额信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *LakeFormationClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociatePrincipals 将一个或者多个用户/用户组加入角色
//
// 将一个或者多个用户/用户组加入角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) AssociatePrincipals(request *model.AssociatePrincipalsRequest) (*model.AssociatePrincipalsResponse, error) {
	requestDef := GenReqDefForAssociatePrincipals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociatePrincipalsResponse), nil
	}
}

// AssociatePrincipalsInvoker 将一个或者多个用户/用户组加入角色
func (c *LakeFormationClient) AssociatePrincipalsInvoker(request *model.AssociatePrincipalsRequest) *AssociatePrincipalsInvoker {
	requestDef := GenReqDefForAssociatePrincipals()
	return &AssociatePrincipalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRole 创建role
//
// 创建role
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateRole(request *model.CreateRoleRequest) (*model.CreateRoleResponse, error) {
	requestDef := GenReqDefForCreateRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoleResponse), nil
	}
}

// CreateRoleInvoker 创建role
func (c *LakeFormationClient) CreateRoleInvoker(request *model.CreateRoleRequest) *CreateRoleInvoker {
	requestDef := GenReqDefForCreateRole()
	return &CreateRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRole 删除角色
//
// 删除指定角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteRole(request *model.DeleteRoleRequest) (*model.DeleteRoleResponse, error) {
	requestDef := GenReqDefForDeleteRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRoleResponse), nil
	}
}

// DeleteRoleInvoker 删除角色
func (c *LakeFormationClient) DeleteRoleInvoker(request *model.DeleteRoleRequest) *DeleteRoleInvoker {
	requestDef := GenReqDefForDeleteRole()
	return &DeleteRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPrincipals 查询角色下的用户/用户组
//
// 查询角色下的用户/用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListPrincipals(request *model.ListPrincipalsRequest) (*model.ListPrincipalsResponse, error) {
	requestDef := GenReqDefForListPrincipals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrincipalsResponse), nil
	}
}

// ListPrincipalsInvoker 查询角色下的用户/用户组
func (c *LakeFormationClient) ListPrincipalsInvoker(request *model.ListPrincipalsRequest) *ListPrincipalsInvoker {
	requestDef := GenReqDefForListPrincipals()
	return &ListPrincipalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoleNames 列举所有角色名
//
// 查询所有角色名字列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListRoleNames(request *model.ListRoleNamesRequest) (*model.ListRoleNamesResponse, error) {
	requestDef := GenReqDefForListRoleNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRoleNamesResponse), nil
	}
}

// ListRoleNamesInvoker 列举所有角色名
func (c *LakeFormationClient) ListRoleNamesInvoker(request *model.ListRoleNamesRequest) *ListRoleNamesInvoker {
	requestDef := GenReqDefForListRoleNames()
	return &ListRoleNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRoles 根据条件分页列举角色
//
// 返回所有角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListRoles(request *model.ListRolesRequest) (*model.ListRolesResponse, error) {
	requestDef := GenReqDefForListRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRolesResponse), nil
	}
}

// ListRolesInvoker 根据条件分页列举角色
func (c *LakeFormationClient) ListRolesInvoker(request *model.ListRolesRequest) *ListRolesInvoker {
	requestDef := GenReqDefForListRoles()
	return &ListRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokePrincipals 将一个或者多个用户/用户组从角色移除
//
// 将一个或者多个用户/用户组从角色移除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) RevokePrincipals(request *model.RevokePrincipalsRequest) (*model.RevokePrincipalsResponse, error) {
	requestDef := GenReqDefForRevokePrincipals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokePrincipalsResponse), nil
	}
}

// RevokePrincipalsInvoker 将一个或者多个用户/用户组从角色移除
func (c *LakeFormationClient) RevokePrincipalsInvoker(request *model.RevokePrincipalsRequest) *RevokePrincipalsInvoker {
	requestDef := GenReqDefForRevokePrincipals()
	return &RevokePrincipalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRole 获取角色
//
// 获取角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowRole(request *model.ShowRoleRequest) (*model.ShowRoleResponse, error) {
	requestDef := GenReqDefForShowRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRoleResponse), nil
	}
}

// ShowRoleInvoker 获取角色
func (c *LakeFormationClient) ShowRoleInvoker(request *model.ShowRoleRequest) *ShowRoleInvoker {
	requestDef := GenReqDefForShowRole()
	return &ShowRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePrincipals 更新角色中的principals
//
// 更新角色中的principals
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdatePrincipals(request *model.UpdatePrincipalsRequest) (*model.UpdatePrincipalsResponse, error) {
	requestDef := GenReqDefForUpdatePrincipals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrincipalsResponse), nil
	}
}

// UpdatePrincipalsInvoker 更新角色中的principals
func (c *LakeFormationClient) UpdatePrincipalsInvoker(request *model.UpdatePrincipalsRequest) *UpdatePrincipalsInvoker {
	requestDef := GenReqDefForUpdatePrincipals()
	return &UpdatePrincipalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRole 修改角色信息
//
// 修改角色信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateRole(request *model.UpdateRoleRequest) (*model.UpdateRoleResponse, error) {
	requestDef := GenReqDefForUpdateRole()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRoleResponse), nil
	}
}

// UpdateRoleInvoker 修改角色信息
func (c *LakeFormationClient) UpdateRoleInvoker(request *model.UpdateRoleRequest) *UpdateRoleInvoker {
	requestDef := GenReqDefForUpdateRole()
	return &UpdateRoleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSpecs
//
// 查询规格列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListSpecs(request *model.ListSpecsRequest) (*model.ListSpecsResponse, error) {
	requestDef := GenReqDefForListSpecs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSpecsResponse), nil
	}
}

// ListSpecsInvoker
func (c *LakeFormationClient) ListSpecsInvoker(request *model.ListSpecsRequest) *ListSpecsInvoker {
	requestDef := GenReqDefForListSpecs()
	return &ListSpecsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTable 创建表
//
// 创建表操作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

// CreateTableInvoker 创建表
func (c *LakeFormationClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAllTables 清空表的数据
//
// 清空表以及表下所有分区的数据及统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteAllTables(request *model.DeleteAllTablesRequest) (*model.DeleteAllTablesResponse, error) {
	requestDef := GenReqDefForDeleteAllTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAllTablesResponse), nil
	}
}

// DeleteAllTablesInvoker 清空表的数据
func (c *LakeFormationClient) DeleteAllTablesInvoker(request *model.DeleteAllTablesRequest) *DeleteAllTablesInvoker {
	requestDef := GenReqDefForDeleteAllTables()
	return &DeleteAllTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTable 删除表
//
// 删除表及表下的分区
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteTable(request *model.DeleteTableRequest) (*model.DeleteTableResponse, error) {
	requestDef := GenReqDefForDeleteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableResponse), nil
	}
}

// DeleteTableInvoker 删除表
func (c *LakeFormationClient) DeleteTableInvoker(request *model.DeleteTableRequest) *DeleteTableInvoker {
	requestDef := GenReqDefForDeleteTable()
	return &DeleteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableMeta 分页获取表的描述信息
//
// 通过数据库通配符和表通配符，找到符合条件的表并返回表的描述信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListTableMeta(request *model.ListTableMetaRequest) (*model.ListTableMetaResponse, error) {
	requestDef := GenReqDefForListTableMeta()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableMetaResponse), nil
	}
}

// ListTableMetaInvoker 分页获取表的描述信息
func (c *LakeFormationClient) ListTableMetaInvoker(request *model.ListTableMetaRequest) *ListTableMetaInvoker {
	requestDef := GenReqDefForListTableMeta()
	return &ListTableMetaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableNames 列举库下所有表名
//
// 查询数据库下的所有表名字列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListTableNames(request *model.ListTableNamesRequest) (*model.ListTableNamesResponse, error) {
	requestDef := GenReqDefForListTableNames()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableNamesResponse), nil
	}
}

// ListTableNamesInvoker 列举库下所有表名
func (c *LakeFormationClient) ListTableNamesInvoker(request *model.ListTableNamesRequest) *ListTableNamesInvoker {
	requestDef := GenReqDefForListTableNames()
	return &ListTableNamesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTables 根据条件分页列举表信息
//
// 返回数据库下符合查询条件的表的元数据信息，不支持事务操作
// 当表名通配符或表类型不为空时，优先根据表名和类型筛选过滤
// 当表名通配符或表类型为空时，再根据属性筛选过滤
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListTables(request *model.ListTablesRequest) (*model.ListTablesResponse, error) {
	requestDef := GenReqDefForListTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesResponse), nil
	}
}

// ListTablesInvoker 根据条件分页列举表信息
func (c *LakeFormationClient) ListTablesInvoker(request *model.ListTablesRequest) *ListTablesInvoker {
	requestDef := GenReqDefForListTables()
	return &ListTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTablesByName 根据表名列举表信息
//
// 根据表名查询数据库下的表信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListTablesByName(request *model.ListTablesByNameRequest) (*model.ListTablesByNameResponse, error) {
	requestDef := GenReqDefForListTablesByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesByNameResponse), nil
	}
}

// ListTablesByNameInvoker 根据表名列举表信息
func (c *LakeFormationClient) ListTablesByNameInvoker(request *model.ListTablesByNameRequest) *ListTablesByNameInvoker {
	requestDef := GenReqDefForListTablesByName()
	return &ListTablesByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTable 获取表信息
//
// 获取表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ShowTable(request *model.ShowTableRequest) (*model.ShowTableResponse, error) {
	requestDef := GenReqDefForShowTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableResponse), nil
	}
}

// ShowTableInvoker 获取表信息
func (c *LakeFormationClient) ShowTableInvoker(request *model.ShowTableRequest) *ShowTableInvoker {
	requestDef := GenReqDefForShowTable()
	return &ShowTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTable 修改表信息
//
// 修改表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateTable(request *model.UpdateTableRequest) (*model.UpdateTableResponse, error) {
	requestDef := GenReqDefForUpdateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTableResponse), nil
	}
}

// UpdateTableInvoker 修改表信息
func (c *LakeFormationClient) UpdateTableInvoker(request *model.UpdateTableRequest) *UpdateTableInvoker {
	requestDef := GenReqDefForUpdateTable()
	return &UpdateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTableColumnStatistics 删除表的指定列统计信息
//
// 删除表的指定列统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteTableColumnStatistics(request *model.DeleteTableColumnStatisticsRequest) (*model.DeleteTableColumnStatisticsResponse, error) {
	requestDef := GenReqDefForDeleteTableColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableColumnStatisticsResponse), nil
	}
}

// DeleteTableColumnStatisticsInvoker 删除表的指定列统计信息
func (c *LakeFormationClient) DeleteTableColumnStatisticsInvoker(request *model.DeleteTableColumnStatisticsRequest) *DeleteTableColumnStatisticsInvoker {
	requestDef := GenReqDefForDeleteTableColumnStatistics()
	return &DeleteTableColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableColumnStatistics 获取表的列统计信息
//
// 获取表的列统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListTableColumnStatistics(request *model.ListTableColumnStatisticsRequest) (*model.ListTableColumnStatisticsResponse, error) {
	requestDef := GenReqDefForListTableColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableColumnStatisticsResponse), nil
	}
}

// ListTableColumnStatisticsInvoker 获取表的列统计信息
func (c *LakeFormationClient) ListTableColumnStatisticsInvoker(request *model.ListTableColumnStatisticsRequest) *ListTableColumnStatisticsInvoker {
	requestDef := GenReqDefForListTableColumnStatistics()
	return &ListTableColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetTableColumnStatistics 更新表的列统计信息
//
// 更新表的列统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) SetTableColumnStatistics(request *model.SetTableColumnStatisticsRequest) (*model.SetTableColumnStatisticsResponse, error) {
	requestDef := GenReqDefForSetTableColumnStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetTableColumnStatisticsResponse), nil
	}
}

// SetTableColumnStatisticsInvoker 更新表的列统计信息
func (c *LakeFormationClient) SetTableColumnStatisticsInvoker(request *model.SetTableColumnStatisticsRequest) *SetTableColumnStatisticsInvoker {
	requestDef := GenReqDefForSetTableColumnStatistics()
	return &SetTableColumnStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateConstraint 批量创建列限制条件
//
// 批量创建表的列限制条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchCreateConstraint(request *model.BatchCreateConstraintRequest) (*model.BatchCreateConstraintResponse, error) {
	requestDef := GenReqDefForBatchCreateConstraint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateConstraintResponse), nil
	}
}

// BatchCreateConstraintInvoker 批量创建列限制条件
func (c *LakeFormationClient) BatchCreateConstraintInvoker(request *model.BatchCreateConstraintRequest) *BatchCreateConstraintInvoker {
	requestDef := GenReqDefForBatchCreateConstraint()
	return &BatchCreateConstraintInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConstraint 删除列限制条件
//
// 删除列限制条件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) DeleteConstraint(request *model.DeleteConstraintRequest) (*model.DeleteConstraintResponse, error) {
	requestDef := GenReqDefForDeleteConstraint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConstraintResponse), nil
	}
}

// DeleteConstraintInvoker 删除列限制条件
func (c *LakeFormationClient) DeleteConstraintInvoker(request *model.DeleteConstraintRequest) *DeleteConstraintInvoker {
	requestDef := GenReqDefForDeleteConstraint()
	return &DeleteConstraintInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConstraints 获取列限制条件
//
// 若查询外键信息，需要在参数中填写被引用表的数据库名和表名。如：parent_db&#x3D;db1&amp;parent_tbl&#x3D;tbl1
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListConstraints(request *model.ListConstraintsRequest) (*model.ListConstraintsResponse, error) {
	requestDef := GenReqDefForListConstraints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConstraintsResponse), nil
	}
}

// ListConstraintsInvoker 获取列限制条件
func (c *LakeFormationClient) ListConstraintsInvoker(request *model.ListConstraintsRequest) *ListConstraintsInvoker {
	requestDef := GenReqDefForListConstraints()
	return &ListConstraintsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateLakeFormationInstanceTags 批量更新标签
//
// 为指定实例批量更新标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) BatchUpdateLakeFormationInstanceTags(request *model.BatchUpdateLakeFormationInstanceTagsRequest) (*model.BatchUpdateLakeFormationInstanceTagsResponse, error) {
	requestDef := GenReqDefForBatchUpdateLakeFormationInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateLakeFormationInstanceTagsResponse), nil
	}
}

// BatchUpdateLakeFormationInstanceTagsInvoker 批量更新标签
func (c *LakeFormationClient) BatchUpdateLakeFormationInstanceTagsInvoker(request *model.BatchUpdateLakeFormationInstanceTagsRequest) *BatchUpdateLakeFormationInstanceTagsInvoker {
	requestDef := GenReqDefForBatchUpdateLakeFormationInstanceTags()
	return &BatchUpdateLakeFormationInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListLakeFormationInstanceTags 查询资源标签集合
//
// 查询租户在指定Project中实例类型的所有资源标签集合
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListLakeFormationInstanceTags(request *model.ListLakeFormationInstanceTagsRequest) (*model.ListLakeFormationInstanceTagsResponse, error) {
	requestDef := GenReqDefForListLakeFormationInstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListLakeFormationInstanceTagsResponse), nil
	}
}

// ListLakeFormationInstanceTagsInvoker 查询资源标签集合
func (c *LakeFormationClient) ListLakeFormationInstanceTagsInvoker(request *model.ListLakeFormationInstanceTagsRequest) *ListLakeFormationInstanceTagsInvoker {
	requestDef := GenReqDefForListLakeFormationInstanceTags()
	return &ListLakeFormationInstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateRoles 将多个角色授予User
//
// 将多个角色授予User
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) AssociateRoles(request *model.AssociateRolesRequest) (*model.AssociateRolesResponse, error) {
	requestDef := GenReqDefForAssociateRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRolesResponse), nil
	}
}

// AssociateRolesInvoker 将多个角色授予User
func (c *LakeFormationClient) AssociateRolesInvoker(request *model.AssociateRolesRequest) *AssociateRolesInvoker {
	requestDef := GenReqDefForAssociateRoles()
	return &AssociateRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUserRoles 查询用户的角色列表
//
// 查询用户的角色列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListUserRoles(request *model.ListUserRolesRequest) (*model.ListUserRolesResponse, error) {
	requestDef := GenReqDefForListUserRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUserRolesResponse), nil
	}
}

// ListUserRolesInvoker 查询用户的角色列表
func (c *LakeFormationClient) ListUserRolesInvoker(request *model.ListUserRolesRequest) *ListUserRolesInvoker {
	requestDef := GenReqDefForListUserRoles()
	return &ListUserRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsers 获取用户列表
//
// 获取用户列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListUsers(request *model.ListUsersRequest) (*model.ListUsersResponse, error) {
	requestDef := GenReqDefForListUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsersResponse), nil
	}
}

// ListUsersInvoker 获取用户列表
func (c *LakeFormationClient) ListUsersInvoker(request *model.ListUsersRequest) *ListUsersInvoker {
	requestDef := GenReqDefForListUsers()
	return &ListUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RevokeRoles 将一个或者多个角色从用户移除
//
// 将一个或者多个角色从用户移除
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) RevokeRoles(request *model.RevokeRolesRequest) (*model.RevokeRolesResponse, error) {
	requestDef := GenReqDefForRevokeRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RevokeRolesResponse), nil
	}
}

// RevokeRolesInvoker 将一个或者多个角色从用户移除
func (c *LakeFormationClient) RevokeRolesInvoker(request *model.RevokeRolesRequest) *RevokeRolesInvoker {
	requestDef := GenReqDefForRevokeRoles()
	return &RevokeRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateRoles 更新用户中的角色
//
// 更新用户中的角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) UpdateRoles(request *model.UpdateRolesRequest) (*model.UpdateRolesResponse, error) {
	requestDef := GenReqDefForUpdateRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRolesResponse), nil
	}
}

// UpdateRolesInvoker 更新用户中的角色
func (c *LakeFormationClient) UpdateRolesInvoker(request *model.UpdateRolesRequest) *UpdateRolesInvoker {
	requestDef := GenReqDefForUpdateRoles()
	return &UpdateRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupsForDomain 获取租户的用户组
//
// 获取租户的用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *LakeFormationClient) ListGroupsForDomain(request *model.ListGroupsForDomainRequest) (*model.ListGroupsForDomainResponse, error) {
	requestDef := GenReqDefForListGroupsForDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupsForDomainResponse), nil
	}
}

// ListGroupsForDomainInvoker 获取租户的用户组
func (c *LakeFormationClient) ListGroupsForDomainInvoker(request *model.ListGroupsForDomainRequest) *ListGroupsForDomainInvoker {
	requestDef := GenReqDefForListGroupsForDomain()
	return &ListGroupsForDomainInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
