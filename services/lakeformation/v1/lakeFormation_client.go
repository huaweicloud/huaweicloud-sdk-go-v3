package v1

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/lakeformation/v1/model"
)

type LakeFormationClient struct {
	HcClient *http_client.HcHttpClient
}

func NewLakeFormationClient(hcClient *http_client.HcHttpClient) *LakeFormationClient {
	return &LakeFormationClient{HcClient: hcClient}
}

func LakeFormationClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
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

// CreateCatalog 创建catalog
//
// 创建catalog，会在catalog下创建默认数据库，默认数据库名字为：default
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
