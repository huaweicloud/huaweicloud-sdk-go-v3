package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dli/v1/model"
)

type DliClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDliClient(hcClient *http_client.HcHttpClient) *DliClient {
	return &DliClient{HcClient: hcClient}
}

func DliClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateJobTemplates 创建作业模板
//
// 该API用于创建作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateJobTemplates(request *model.CreateJobTemplatesRequest) (*model.CreateJobTemplatesResponse, error) {
	requestDef := GenReqDefForCreateJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobTemplatesResponse), nil
	}
}

// CreateJobTemplatesInvoker 创建作业模板
func (c *DliClient) CreateJobTemplatesInvoker(request *model.CreateJobTemplatesRequest) *CreateJobTemplatesInvoker {
	requestDef := GenReqDefForCreateJobTemplates()
	return &CreateJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlTemplates 存储指定SQL语句
//
// 该API用于存储指定的SQL语句，后续可以重复使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateSqlTemplates(request *model.CreateSqlTemplatesRequest) (*model.CreateSqlTemplatesResponse, error) {
	requestDef := GenReqDefForCreateSqlTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlTemplatesResponse), nil
	}
}

// CreateSqlTemplatesInvoker 存储指定SQL语句
func (c *DliClient) CreateSqlTemplatesInvoker(request *model.CreateSqlTemplatesRequest) *CreateSqlTemplatesInvoker {
	requestDef := GenReqDefForCreateSqlTemplates()
	return &CreateSqlTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSqlTemplates 批量删除SQL模板
//
// 该API用于批量删除SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteSqlTemplates(request *model.DeleteSqlTemplatesRequest) (*model.DeleteSqlTemplatesResponse, error) {
	requestDef := GenReqDefForDeleteSqlTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSqlTemplatesResponse), nil
	}
}

// DeleteSqlTemplatesInvoker 批量删除SQL模板
func (c *DliClient) DeleteSqlTemplatesInvoker(request *model.DeleteSqlTemplatesRequest) *DeleteSqlTemplatesInvoker {
	requestDef := GenReqDefForDeleteSqlTemplates()
	return &DeleteSqlTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobTemplates 查询作业模板列表
//
// 该API用于查询作业模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListJobTemplates(request *model.ListJobTemplatesRequest) (*model.ListJobTemplatesResponse, error) {
	requestDef := GenReqDefForListJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobTemplatesResponse), nil
	}
}

// ListJobTemplatesInvoker 查询作业模板列表
func (c *DliClient) ListJobTemplatesInvoker(request *model.ListJobTemplatesRequest) *ListJobTemplatesInvoker {
	requestDef := GenReqDefForListJobTemplates()
	return &ListJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobTemplate 获取作业模板
//
// 该API用于获取作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowJobTemplate(request *model.ShowJobTemplateRequest) (*model.ShowJobTemplateResponse, error) {
	requestDef := GenReqDefForShowJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobTemplateResponse), nil
	}
}

// ShowJobTemplateInvoker 获取作业模板
func (c *DliClient) ShowJobTemplateInvoker(request *model.ShowJobTemplateRequest) *ShowJobTemplateInvoker {
	requestDef := GenReqDefForShowJobTemplate()
	return &ShowJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlSampleTemplates 查询所有SQL样例模板
//
// 该API用于查询所有SQL样例模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSqlSampleTemplates(request *model.ShowSqlSampleTemplatesRequest) (*model.ShowSqlSampleTemplatesResponse, error) {
	requestDef := GenReqDefForShowSqlSampleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlSampleTemplatesResponse), nil
	}
}

// ShowSqlSampleTemplatesInvoker 查询所有SQL样例模板
func (c *DliClient) ShowSqlSampleTemplatesInvoker(request *model.ShowSqlSampleTemplatesRequest) *ShowSqlSampleTemplatesInvoker {
	requestDef := GenReqDefForShowSqlSampleTemplates()
	return &ShowSqlSampleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlTemplates 查看所有SQL模板
//
// 该API用查看用户保存的所有SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSqlTemplates(request *model.ShowSqlTemplatesRequest) (*model.ShowSqlTemplatesResponse, error) {
	requestDef := GenReqDefForShowSqlTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlTemplatesResponse), nil
	}
}

// ShowSqlTemplatesInvoker 查看所有SQL模板
func (c *DliClient) ShowSqlTemplatesInvoker(request *model.ShowSqlTemplatesRequest) *ShowSqlTemplatesInvoker {
	requestDef := GenReqDefForShowSqlTemplates()
	return &ShowSqlTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJobTemplates 修改作业模板
//
// 该API用于修改作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateJobTemplates(request *model.UpdateJobTemplatesRequest) (*model.UpdateJobTemplatesResponse, error) {
	requestDef := GenReqDefForUpdateJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobTemplatesResponse), nil
	}
}

// UpdateJobTemplatesInvoker 修改作业模板
func (c *DliClient) UpdateJobTemplatesInvoker(request *model.UpdateJobTemplatesRequest) *UpdateJobTemplatesInvoker {
	requestDef := GenReqDefForUpdateJobTemplates()
	return &UpdateJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSqlTemplates 更新SQL模板
//
// 该API用于更新SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateSqlTemplates(request *model.UpdateSqlTemplatesRequest) (*model.UpdateSqlTemplatesResponse, error) {
	requestDef := GenReqDefForUpdateSqlTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlTemplatesResponse), nil
	}
}

// UpdateSqlTemplatesInvoker 更新SQL模板
func (c *DliClient) UpdateSqlTemplatesInvoker(request *model.UpdateSqlTemplatesRequest) *UpdateSqlTemplatesInvoker {
	requestDef := GenReqDefForUpdateSqlTemplates()
	return &UpdateSqlTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateConnectionQueue 绑定队列
//
// 该API用于在已创建的增强型跨源中绑定队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) AssociateConnectionQueue(request *model.AssociateConnectionQueueRequest) (*model.AssociateConnectionQueueResponse, error) {
	requestDef := GenReqDefForAssociateConnectionQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateConnectionQueueResponse), nil
	}
}

// AssociateConnectionQueueInvoker 绑定队列
func (c *DliClient) AssociateConnectionQueueInvoker(request *model.AssociateConnectionQueueRequest) *AssociateConnectionQueueInvoker {
	requestDef := GenReqDefForAssociateConnectionQueue()
	return &AssociateConnectionQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateQueueToElasticResourcePool 关联队列到弹性资源池
//
// 关联队列到弹性资源池
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) AssociateQueueToElasticResourcePool(request *model.AssociateQueueToElasticResourcePoolRequest) (*model.AssociateQueueToElasticResourcePoolResponse, error) {
	requestDef := GenReqDefForAssociateQueueToElasticResourcePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateQueueToElasticResourcePoolResponse), nil
	}
}

// AssociateQueueToElasticResourcePoolInvoker 关联队列到弹性资源池
func (c *DliClient) AssociateQueueToElasticResourcePoolInvoker(request *model.AssociateQueueToElasticResourcePoolRequest) *AssociateQueueToElasticResourcePoolInvoker {
	requestDef := GenReqDefForAssociateQueueToElasticResourcePool()
	return &AssociateQueueToElasticResourcePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeResource 数据赋权（用户、项目）
//
// 该API用于将DLI资源权限赋给、回收、更新指定的其他用户或项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) AuthorizeResource(request *model.AuthorizeResourceRequest) (*model.AuthorizeResourceResponse, error) {
	requestDef := GenReqDefForAuthorizeResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeResourceResponse), nil
	}
}

// AuthorizeResourceInvoker 数据赋权（用户、项目）
func (c *DliClient) AuthorizeResourceInvoker(request *model.AuthorizeResourceRequest) *AuthorizeResourceInvoker {
	requestDef := GenReqDefForAuthorizeResource()
	return &AuthorizeResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteQueuePlans 批量删除队列定时扩缩容计划
//
// 该API用于批量删除队列定时扩缩容计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) BatchDeleteQueuePlans(request *model.BatchDeleteQueuePlansRequest) (*model.BatchDeleteQueuePlansResponse, error) {
	requestDef := GenReqDefForBatchDeleteQueuePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteQueuePlansResponse), nil
	}
}

// BatchDeleteQueuePlansInvoker 批量删除队列定时扩缩容计划
func (c *DliClient) BatchDeleteQueuePlansInvoker(request *model.BatchDeleteQueuePlansRequest) *BatchDeleteQueuePlansInvoker {
	requestDef := GenReqDefForBatchDeleteQueuePlans()
	return &BatchDeleteQueuePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAuthorization 数据赋权（用户）
//
// 该API用于将数据库或数据表的数据权限赋给指定的其他用户。
// 说明：
// 被赋权用户所在用户组的所属区域需具有Tenant Guest权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ChangeAuthorization(request *model.ChangeAuthorizationRequest) (*model.ChangeAuthorizationResponse, error) {
	requestDef := GenReqDefForChangeAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAuthorizationResponse), nil
	}
}

// ChangeAuthorizationInvoker 数据赋权（用户）
func (c *DliClient) ChangeAuthorizationInvoker(request *model.ChangeAuthorizationRequest) *ChangeAuthorizationInvoker {
	requestDef := GenReqDefForChangeAuthorization()
	return &ChangeAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeQueuePlan 修改队列定时扩缩容计划
//
// 该API用于修改指定ID的队列定时扩缩容计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ChangeQueuePlan(request *model.ChangeQueuePlanRequest) (*model.ChangeQueuePlanResponse, error) {
	requestDef := GenReqDefForChangeQueuePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeQueuePlanResponse), nil
	}
}

// ChangeQueuePlanInvoker 修改队列定时扩缩容计划
func (c *DliClient) ChangeQueuePlanInvoker(request *model.ChangeQueuePlanRequest) *ChangeQueuePlanInvoker {
	requestDef := GenReqDefForChangeQueuePlan()
	return &ChangeQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckConnection 创建地址连通性请求
//
// 创建地址连通性请求API接口，往指定集群发送地址连通性测试请求，并将测试地址插入表内
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CheckConnection(request *model.CheckConnectionRequest) (*model.CheckConnectionResponse, error) {
	requestDef := GenReqDefForCheckConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckConnectionResponse), nil
	}
}

// CheckConnectionInvoker 创建地址连通性请求
func (c *DliClient) CheckConnectionInvoker(request *model.CheckConnectionRequest) *CheckConnectionInvoker {
	requestDef := GenReqDefForCheckConnection()
	return &CheckConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthInfo 创建跨源认证
//
// 该API用于创建跨源认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateAuthInfo(request *model.CreateAuthInfoRequest) (*model.CreateAuthInfoResponse, error) {
	requestDef := GenReqDefForCreateAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthInfoResponse), nil
	}
}

// CreateAuthInfoInvoker 创建跨源认证
func (c *DliClient) CreateAuthInfoInvoker(request *model.CreateAuthInfoRequest) *CreateAuthInfoInvoker {
	requestDef := GenReqDefForCreateAuthInfo()
	return &CreateAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatasourceConnection 创建经典型跨源连接
//
// 该API用于创建与其他服务的经典型跨源连接。
// 说明：
// 如果需要了解Console界面的使用方法，可参考经典型跨源连接。
// 系统default队列不支持创建跨源连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateDatasourceConnection(request *model.CreateDatasourceConnectionRequest) (*model.CreateDatasourceConnectionResponse, error) {
	requestDef := GenReqDefForCreateDatasourceConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatasourceConnectionResponse), nil
	}
}

// CreateDatasourceConnectionInvoker 创建经典型跨源连接
func (c *DliClient) CreateDatasourceConnectionInvoker(request *model.CreateDatasourceConnectionRequest) *CreateDatasourceConnectionInvoker {
	requestDef := GenReqDefForCreateDatasourceConnection()
	return &CreateDatasourceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDliAgency 创建DLI委托
//
// 创建DLI委托
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateDliAgency(request *model.CreateDliAgencyRequest) (*model.CreateDliAgencyResponse, error) {
	requestDef := GenReqDefForCreateDliAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDliAgencyResponse), nil
	}
}

// CreateDliAgencyInvoker 创建DLI委托
func (c *DliClient) CreateDliAgencyInvoker(request *model.CreateDliAgencyRequest) *CreateDliAgencyInvoker {
	requestDef := GenReqDefForCreateDliAgency()
	return &CreateDliAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateElasticResourcePool 创建弹性资源池
//
// 创建弹性资源池
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateElasticResourcePool(request *model.CreateElasticResourcePoolRequest) (*model.CreateElasticResourcePoolResponse, error) {
	requestDef := GenReqDefForCreateElasticResourcePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateElasticResourcePoolResponse), nil
	}
}

// CreateElasticResourcePoolInvoker 创建弹性资源池
func (c *DliClient) CreateElasticResourcePoolInvoker(request *model.CreateElasticResourcePoolRequest) *CreateElasticResourcePoolInvoker {
	requestDef := GenReqDefForCreateElasticResourcePool()
	return &CreateElasticResourcePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnhancedConnection 创建增强型跨源连接
//
// 该API用于创建与其他服务的增强型跨源连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateEnhancedConnection(request *model.CreateEnhancedConnectionRequest) (*model.CreateEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForCreateEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnhancedConnectionResponse), nil
	}
}

// CreateEnhancedConnectionInvoker 创建增强型跨源连接
func (c *DliClient) CreateEnhancedConnectionInvoker(request *model.CreateEnhancedConnectionRequest) *CreateEnhancedConnectionInvoker {
	requestDef := GenReqDefForCreateEnhancedConnection()
	return &CreateEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEnhancedConnectionRoutes 创建路由
//
// 该API用于创建跨源需要的路由。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateEnhancedConnectionRoutes(request *model.CreateEnhancedConnectionRoutesRequest) (*model.CreateEnhancedConnectionRoutesResponse, error) {
	requestDef := GenReqDefForCreateEnhancedConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEnhancedConnectionRoutesResponse), nil
	}
}

// CreateEnhancedConnectionRoutesInvoker 创建路由
func (c *DliClient) CreateEnhancedConnectionRoutesInvoker(request *model.CreateEnhancedConnectionRoutesRequest) *CreateEnhancedConnectionRoutesInvoker {
	requestDef := GenReqDefForCreateEnhancedConnectionRoutes()
	return &CreateEnhancedConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalValue 创建DLI全局变量
//
// 创建全局变量。全局变量用于替换SQL作业中的敏感数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateGlobalValue(request *model.CreateGlobalValueRequest) (*model.CreateGlobalValueResponse, error) {
	requestDef := GenReqDefForCreateGlobalValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalValueResponse), nil
	}
}

// CreateGlobalValueInvoker 创建DLI全局变量
func (c *DliClient) CreateGlobalValueInvoker(request *model.CreateGlobalValueRequest) *CreateGlobalValueInvoker {
	requestDef := GenReqDefForCreateGlobalValue()
	return &CreateGlobalValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQueue 创建队列
//
// 该API用于创建队列，该队列将会绑定用户指定的计算资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateQueue(request *model.CreateQueueRequest) (*model.CreateQueueResponse, error) {
	requestDef := GenReqDefForCreateQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQueueResponse), nil
	}
}

// CreateQueueInvoker 创建队列
func (c *DliClient) CreateQueueInvoker(request *model.CreateQueueRequest) *CreateQueueInvoker {
	requestDef := GenReqDefForCreateQueue()
	return &CreateQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQueuePlan 创建队列定时扩缩容计划
//
// 创建队列定时扩缩容计划接口，对指定的队列创建定时规格变更计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateQueuePlan(request *model.CreateQueuePlanRequest) (*model.CreateQueuePlanResponse, error) {
	requestDef := GenReqDefForCreateQueuePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQueuePlanResponse), nil
	}
}

// CreateQueuePlanInvoker 创建队列定时扩缩容计划
func (c *DliClient) CreateQueuePlanInvoker(request *model.CreateQueuePlanRequest) *CreateQueuePlanInvoker {
	requestDef := GenReqDefForCreateQueuePlan()
	return &CreateQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuthInfo 删除跨源认证
//
// 该API用于删除跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteAuthInfo(request *model.DeleteAuthInfoRequest) (*model.DeleteAuthInfoResponse, error) {
	requestDef := GenReqDefForDeleteAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuthInfoResponse), nil
	}
}

// DeleteAuthInfoInvoker 删除跨源认证
func (c *DliClient) DeleteAuthInfoInvoker(request *model.DeleteAuthInfoRequest) *DeleteAuthInfoInvoker {
	requestDef := GenReqDefForDeleteAuthInfo()
	return &DeleteAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatasourceConnection 删除经典型跨源连接
//
// 该API用于删除已创建的经典型跨源连接。
// 说明：
// 创建中的连接，无法删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteDatasourceConnection(request *model.DeleteDatasourceConnectionRequest) (*model.DeleteDatasourceConnectionResponse, error) {
	requestDef := GenReqDefForDeleteDatasourceConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatasourceConnectionResponse), nil
	}
}

// DeleteDatasourceConnectionInvoker 删除经典型跨源连接
func (c *DliClient) DeleteDatasourceConnectionInvoker(request *model.DeleteDatasourceConnectionRequest) *DeleteDatasourceConnectionInvoker {
	requestDef := GenReqDefForDeleteDatasourceConnection()
	return &DeleteDatasourceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteElasticResourcePool 删除弹性资源池
//
// 删除弹性资源池
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteElasticResourcePool(request *model.DeleteElasticResourcePoolRequest) (*model.DeleteElasticResourcePoolResponse, error) {
	requestDef := GenReqDefForDeleteElasticResourcePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteElasticResourcePoolResponse), nil
	}
}

// DeleteElasticResourcePoolInvoker 删除弹性资源池
func (c *DliClient) DeleteElasticResourcePoolInvoker(request *model.DeleteElasticResourcePoolRequest) *DeleteElasticResourcePoolInvoker {
	requestDef := GenReqDefForDeleteElasticResourcePool()
	return &DeleteElasticResourcePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnhancedConnection 删除增强型跨源连接
//
// 该API用于删除已创建的增强型跨源连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteEnhancedConnection(request *model.DeleteEnhancedConnectionRequest) (*model.DeleteEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForDeleteEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnhancedConnectionResponse), nil
	}
}

// DeleteEnhancedConnectionInvoker 删除增强型跨源连接
func (c *DliClient) DeleteEnhancedConnectionInvoker(request *model.DeleteEnhancedConnectionRequest) *DeleteEnhancedConnectionInvoker {
	requestDef := GenReqDefForDeleteEnhancedConnection()
	return &DeleteEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEnhancedConnectionRoutes 删除路由
//
// 该API用于删除跨源需要的路由。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteEnhancedConnectionRoutes(request *model.DeleteEnhancedConnectionRoutesRequest) (*model.DeleteEnhancedConnectionRoutesResponse, error) {
	requestDef := GenReqDefForDeleteEnhancedConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEnhancedConnectionRoutesResponse), nil
	}
}

// DeleteEnhancedConnectionRoutesInvoker 删除路由
func (c *DliClient) DeleteEnhancedConnectionRoutesInvoker(request *model.DeleteEnhancedConnectionRoutesRequest) *DeleteEnhancedConnectionRoutesInvoker {
	requestDef := GenReqDefForDeleteEnhancedConnectionRoutes()
	return &DeleteEnhancedConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalValue 删除DLI全局变量
//
// 删除全局变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteGlobalValue(request *model.DeleteGlobalValueRequest) (*model.DeleteGlobalValueResponse, error) {
	requestDef := GenReqDefForDeleteGlobalValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalValueResponse), nil
	}
}

// DeleteGlobalValueInvoker 删除DLI全局变量
func (c *DliClient) DeleteGlobalValueInvoker(request *model.DeleteGlobalValueRequest) *DeleteGlobalValueInvoker {
	requestDef := GenReqDefForDeleteGlobalValue()
	return &DeleteGlobalValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQueue 删除队列
//
// 该API用于删除指定队列。
// 说明：
// 若指定队列正在执行任务，则不允许删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteQueue(request *model.DeleteQueueRequest) (*model.DeleteQueueResponse, error) {
	requestDef := GenReqDefForDeleteQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueueResponse), nil
	}
}

// DeleteQueueInvoker 删除队列
func (c *DliClient) DeleteQueueInvoker(request *model.DeleteQueueRequest) *DeleteQueueInvoker {
	requestDef := GenReqDefForDeleteQueue()
	return &DeleteQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQueuePlan 单个删除队列定时扩缩容计划
//
// 该API用于删除指定ID的队列定时扩缩容计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteQueuePlan(request *model.DeleteQueuePlanRequest) (*model.DeleteQueuePlanResponse, error) {
	requestDef := GenReqDefForDeleteQueuePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueuePlanResponse), nil
	}
}

// DeleteQueuePlanInvoker 单个删除队列定时扩缩容计划
func (c *DliClient) DeleteQueuePlanInvoker(request *model.DeleteQueuePlanRequest) *DeleteQueuePlanInvoker {
	requestDef := GenReqDefForDeleteQueuePlan()
	return &DeleteQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResource 删除组内资源包
//
// 该API用于删除某个project某个分组下的资源包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteResource(request *model.DeleteResourceRequest) (*model.DeleteResourceResponse, error) {
	requestDef := GenReqDefForDeleteResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceResponse), nil
	}
}

// DeleteResourceInvoker 删除组内资源包
func (c *DliClient) DeleteResourceInvoker(request *model.DeleteResourceRequest) *DeleteResourceInvoker {
	requestDef := GenReqDefForDeleteResource()
	return &DeleteResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateConnectionQueue 解绑队列
//
// 该API用于在增强型跨源中解绑已绑定的队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DisassociateConnectionQueue(request *model.DisassociateConnectionQueueRequest) (*model.DisassociateConnectionQueueResponse, error) {
	requestDef := GenReqDefForDisassociateConnectionQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateConnectionQueueResponse), nil
	}
}

// DisassociateConnectionQueueInvoker 解绑队列
func (c *DliClient) DisassociateConnectionQueueInvoker(request *model.DisassociateConnectionQueueRequest) *DisassociateConnectionQueueInvoker {
	requestDef := GenReqDefForDisassociateConnectionQueue()
	return &DisassociateConnectionQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthInfo 获取跨源认证列表
//
// 该API用于查询跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListAuthInfo(request *model.ListAuthInfoRequest) (*model.ListAuthInfoResponse, error) {
	requestDef := GenReqDefForListAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthInfoResponse), nil
	}
}

// ListAuthInfoInvoker 获取跨源认证列表
func (c *DliClient) ListAuthInfoInvoker(request *model.ListAuthInfoRequest) *ListAuthInfoInvoker {
	requestDef := GenReqDefForListAuthInfo()
	return &ListAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabaseUsers 查看数据库的使用者
//
// 该API用于查询可以使用的指定队列的所有用户名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListDatabaseUsers(request *model.ListDatabaseUsersRequest) (*model.ListDatabaseUsersResponse, error) {
	requestDef := GenReqDefForListDatabaseUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabaseUsersResponse), nil
	}
}

// ListDatabaseUsersInvoker 查看数据库的使用者
func (c *DliClient) ListDatabaseUsersInvoker(request *model.ListDatabaseUsersRequest) *ListDatabaseUsersInvoker {
	requestDef := GenReqDefForListDatabaseUsers()
	return &ListDatabaseUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatasourceConnections 查询经典型跨源连接列表
//
// 该API用于查询该用户已创建的经典型跨源连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListDatasourceConnections(request *model.ListDatasourceConnectionsRequest) (*model.ListDatasourceConnectionsResponse, error) {
	requestDef := GenReqDefForListDatasourceConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatasourceConnectionsResponse), nil
	}
}

// ListDatasourceConnectionsInvoker 查询经典型跨源连接列表
func (c *DliClient) ListDatasourceConnectionsInvoker(request *model.ListDatasourceConnectionsRequest) *ListDatasourceConnectionsInvoker {
	requestDef := GenReqDefForListDatasourceConnections()
	return &ListDatasourceConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListElasticResourcePoolQueues 查询弹性资源池所属队列
//
// 查询弹性资源池所属队列
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListElasticResourcePoolQueues(request *model.ListElasticResourcePoolQueuesRequest) (*model.ListElasticResourcePoolQueuesResponse, error) {
	requestDef := GenReqDefForListElasticResourcePoolQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListElasticResourcePoolQueuesResponse), nil
	}
}

// ListElasticResourcePoolQueuesInvoker 查询弹性资源池所属队列
func (c *DliClient) ListElasticResourcePoolQueuesInvoker(request *model.ListElasticResourcePoolQueuesRequest) *ListElasticResourcePoolQueuesInvoker {
	requestDef := GenReqDefForListElasticResourcePoolQueues()
	return &ListElasticResourcePoolQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListElasticResourcePoolScaleRecords 弹性资源池扩缩容历史记录
//
// 查询当前弹性资源池扩缩容历史记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListElasticResourcePoolScaleRecords(request *model.ListElasticResourcePoolScaleRecordsRequest) (*model.ListElasticResourcePoolScaleRecordsResponse, error) {
	requestDef := GenReqDefForListElasticResourcePoolScaleRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListElasticResourcePoolScaleRecordsResponse), nil
	}
}

// ListElasticResourcePoolScaleRecordsInvoker 弹性资源池扩缩容历史记录
func (c *DliClient) ListElasticResourcePoolScaleRecordsInvoker(request *model.ListElasticResourcePoolScaleRecordsRequest) *ListElasticResourcePoolScaleRecordsInvoker {
	requestDef := GenReqDefForListElasticResourcePoolScaleRecords()
	return &ListElasticResourcePoolScaleRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListElasticResourcePools 查询所有弹性资源池
//
// 查询所有弹性资源池
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListElasticResourcePools(request *model.ListElasticResourcePoolsRequest) (*model.ListElasticResourcePoolsResponse, error) {
	requestDef := GenReqDefForListElasticResourcePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListElasticResourcePoolsResponse), nil
	}
}

// ListElasticResourcePoolsInvoker 查询所有弹性资源池
func (c *DliClient) ListElasticResourcePoolsInvoker(request *model.ListElasticResourcePoolsRequest) *ListElasticResourcePoolsInvoker {
	requestDef := GenReqDefForListElasticResourcePools()
	return &ListElasticResourcePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEnhancedConnections 查询增强型跨源连接列表
//
// 该API用于查询该用户已创建的增强型跨源连接列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListEnhancedConnections(request *model.ListEnhancedConnectionsRequest) (*model.ListEnhancedConnectionsResponse, error) {
	requestDef := GenReqDefForListEnhancedConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnhancedConnectionsResponse), nil
	}
}

// ListEnhancedConnectionsInvoker 查询增强型跨源连接列表
func (c *DliClient) ListEnhancedConnectionsInvoker(request *model.ListEnhancedConnectionsRequest) *ListEnhancedConnectionsInvoker {
	requestDef := GenReqDefForListEnhancedConnections()
	return &ListEnhancedConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalValues 查询DLI全局变量列表
//
// 查询全局变量列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListGlobalValues(request *model.ListGlobalValuesRequest) (*model.ListGlobalValuesResponse, error) {
	requestDef := GenReqDefForListGlobalValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalValuesResponse), nil
	}
}

// ListGlobalValuesInvoker 查询DLI全局变量列表
func (c *DliClient) ListGlobalValuesInvoker(request *model.ListGlobalValuesRequest) *ListGlobalValuesInvoker {
	requestDef := GenReqDefForListGlobalValues()
	return &ListGlobalValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueuePlans 查看队列定时扩缩容计划
//
// 查看队列定时扩缩容计划接口，列出指定队列定时规格变更计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListQueuePlans(request *model.ListQueuePlansRequest) (*model.ListQueuePlansResponse, error) {
	requestDef := GenReqDefForListQueuePlans()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueuePlansResponse), nil
	}
}

// ListQueuePlansInvoker 查看队列定时扩缩容计划
func (c *DliClient) ListQueuePlansInvoker(request *model.ListQueuePlansRequest) *ListQueuePlansInvoker {
	requestDef := GenReqDefForListQueuePlans()
	return &ListQueuePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueueUsers 查看队列的使用者
//
// 该API用于查询可以使用的指定队列的所有用户名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListQueueUsers(request *model.ListQueueUsersRequest) (*model.ListQueueUsersResponse, error) {
	requestDef := GenReqDefForListQueueUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueueUsersResponse), nil
	}
}

// ListQueueUsersInvoker 查看队列的使用者
func (c *DliClient) ListQueueUsersInvoker(request *model.ListQueueUsersRequest) *ListQueueUsersInvoker {
	requestDef := GenReqDefForListQueueUsers()
	return &ListQueueUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueues 查询所有队列
//
// 该API用于列出该project下所有的队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListQueues(request *model.ListQueuesRequest) (*model.ListQueuesResponse, error) {
	requestDef := GenReqDefForListQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueuesResponse), nil
	}
}

// ListQueuesInvoker 查询所有队列
func (c *DliClient) ListQueuesInvoker(request *model.ListQueuesRequest) *ListQueuesInvoker {
	requestDef := GenReqDefForListQueues()
	return &ListQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 查看分组资源列表
//
// 该API用于查看某个project下的所有资源，其中包含Group。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 查看分组资源列表
func (c *DliClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTablePrivileges 查看表的用户权限
//
// 该API用于查询指定用户在表上的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListTablePrivileges(request *model.ListTablePrivilegesRequest) (*model.ListTablePrivilegesResponse, error) {
	requestDef := GenReqDefForListTablePrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablePrivilegesResponse), nil
	}
}

// ListTablePrivilegesInvoker 查看表的用户权限
func (c *DliClient) ListTablePrivilegesInvoker(request *model.ListTablePrivilegesRequest) *ListTablePrivilegesInvoker {
	requestDef := GenReqDefForListTablePrivileges()
	return &ListTablePrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableUsers 查看表的使用者
//
// 该API用于查看有权访问指定表或表的列的所有用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListTableUsers(request *model.ListTableUsersRequest) (*model.ListTableUsersResponse, error) {
	requestDef := GenReqDefForListTableUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableUsersResponse), nil
	}
}

// ListTableUsersInvoker 查看表的使用者
func (c *DliClient) ListTableUsersInvoker(request *model.ListTableUsersRequest) *ListTableUsersInvoker {
	requestDef := GenReqDefForListTableUsers()
	return &ListTableUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterAuthorizedQueue 队列赋权
//
// 该API用于与其他用户共享指定的队列，可以给用户赋使用指定的队列的权限或者收回使用权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RegisterAuthorizedQueue(request *model.RegisterAuthorizedQueueRequest) (*model.RegisterAuthorizedQueueResponse, error) {
	requestDef := GenReqDefForRegisterAuthorizedQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterAuthorizedQueueResponse), nil
	}
}

// RegisterAuthorizedQueueInvoker 队列赋权
func (c *DliClient) RegisterAuthorizedQueueInvoker(request *model.RegisterAuthorizedQueueRequest) *RegisterAuthorizedQueueInvoker {
	requestDef := GenReqDefForRegisterAuthorizedQueue()
	return &RegisterAuthorizedQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunQueueAction 重启/扩容/缩容队列
//
// 该功能用于重新启动队列、扩容队列、缩容队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RunQueueAction(request *model.RunQueueActionRequest) (*model.RunQueueActionResponse, error) {
	requestDef := GenReqDefForRunQueueAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunQueueActionResponse), nil
	}
}

// RunQueueActionInvoker 重启/扩容/缩容队列
func (c *DliClient) RunQueueActionInvoker(request *model.RunQueueActionRequest) *RunQueueActionInvoker {
	requestDef := GenReqDefForRunQueueAction()
	return &RunQueueActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatasourceConnection 查询经典型跨源连接
//
// 该API用于查询该用户指定的已创建的经典型跨源连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowDatasourceConnection(request *model.ShowDatasourceConnectionRequest) (*model.ShowDatasourceConnectionResponse, error) {
	requestDef := GenReqDefForShowDatasourceConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatasourceConnectionResponse), nil
	}
}

// ShowDatasourceConnectionInvoker 查询经典型跨源连接
func (c *DliClient) ShowDatasourceConnectionInvoker(request *model.ShowDatasourceConnectionRequest) *ShowDatasourceConnectionInvoker {
	requestDef := GenReqDefForShowDatasourceConnection()
	return &ShowDatasourceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDliAgency 获取dli委托信息
//
// 获取dli委托信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowDliAgency(request *model.ShowDliAgencyRequest) (*model.ShowDliAgencyResponse, error) {
	requestDef := GenReqDefForShowDliAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDliAgencyResponse), nil
	}
}

// ShowDliAgencyInvoker 获取dli委托信息
func (c *DliClient) ShowDliAgencyInvoker(request *model.ShowDliAgencyRequest) *ShowDliAgencyInvoker {
	requestDef := GenReqDefForShowDliAgency()
	return &ShowDliAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnhancedConnection 查询增强型跨源连接
//
// 该API用于查询该用户指定的已创建的增强型跨源连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowEnhancedConnection(request *model.ShowEnhancedConnectionRequest) (*model.ShowEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForShowEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnhancedConnectionResponse), nil
	}
}

// ShowEnhancedConnectionInvoker 查询增强型跨源连接
func (c *DliClient) ShowEnhancedConnectionInvoker(request *model.ShowEnhancedConnectionRequest) *ShowEnhancedConnectionInvoker {
	requestDef := GenReqDefForShowEnhancedConnection()
	return &ShowEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEnhancedPrivilege 查询增强型跨源授权信息
//
// 该API用于查询增强型跨源连接授权的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowEnhancedPrivilege(request *model.ShowEnhancedPrivilegeRequest) (*model.ShowEnhancedPrivilegeResponse, error) {
	requestDef := GenReqDefForShowEnhancedPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnhancedPrivilegeResponse), nil
	}
}

// ShowEnhancedPrivilegeInvoker 查询增强型跨源授权信息
func (c *DliClient) ShowEnhancedPrivilegeInvoker(request *model.ShowEnhancedPrivilegeRequest) *ShowEnhancedPrivilegeInvoker {
	requestDef := GenReqDefForShowEnhancedPrivilege()
	return &ShowEnhancedPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodeConnectivity 查询指定地址连通性测试详情
//
// 该API用于在连通性测试提交后查询连通性结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowNodeConnectivity(request *model.ShowNodeConnectivityRequest) (*model.ShowNodeConnectivityResponse, error) {
	requestDef := GenReqDefForShowNodeConnectivity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeConnectivityResponse), nil
	}
}

// ShowNodeConnectivityInvoker 查询指定地址连通性测试详情
func (c *DliClient) ShowNodeConnectivityInvoker(request *model.ShowNodeConnectivityRequest) *ShowNodeConnectivityInvoker {
	requestDef := GenReqDefForShowNodeConnectivity()
	return &ShowNodeConnectivityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowObjectUser 查看赋权对象的用者权限信息
//
// 获取对象赋权用户的权限信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowObjectUser(request *model.ShowObjectUserRequest) (*model.ShowObjectUserResponse, error) {
	requestDef := GenReqDefForShowObjectUser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowObjectUserResponse), nil
	}
}

// ShowObjectUserInvoker 查看赋权对象的用者权限信息
func (c *DliClient) ShowObjectUserInvoker(request *model.ShowObjectUserRequest) *ShowObjectUserInvoker {
	requestDef := GenReqDefForShowObjectUser()
	return &ShowObjectUserInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQueueDetail 查询队列详情
//
// 该API用于列出该project下指定的队列详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowQueueDetail(request *model.ShowQueueDetailRequest) (*model.ShowQueueDetailResponse, error) {
	requestDef := GenReqDefForShowQueueDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueueDetailResponse), nil
	}
}

// ShowQueueDetailInvoker 查询队列详情
func (c *DliClient) ShowQueueDetailInvoker(request *model.ShowQueueDetailRequest) *ShowQueueDetailInvoker {
	requestDef := GenReqDefForShowQueueDetail()
	return &ShowQueueDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceInfo 查看组内资源包
//
// 该API用于查看某个project某个分组下的具体资源信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowResourceInfo(request *model.ShowResourceInfoRequest) (*model.ShowResourceInfoResponse, error) {
	requestDef := GenReqDefForShowResourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceInfoResponse), nil
	}
}

// ShowResourceInfoInvoker 查看组内资源包
func (c *DliClient) ShowResourceInfoInvoker(request *model.ShowResourceInfoRequest) *ShowResourceInfoInvoker {
	requestDef := GenReqDefForShowResourceInfo()
	return &ShowResourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuthInfo 更新跨源认证
//
// 该API用于更新跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateAuthInfo(request *model.UpdateAuthInfoRequest) (*model.UpdateAuthInfoResponse, error) {
	requestDef := GenReqDefForUpdateAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthInfoResponse), nil
	}
}

// UpdateAuthInfoInvoker 更新跨源认证
func (c *DliClient) UpdateAuthInfoInvoker(request *model.UpdateAuthInfoRequest) *UpdateAuthInfoInvoker {
	requestDef := GenReqDefForUpdateAuthInfo()
	return &UpdateAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateElasticResourcePool 修改弹性资源池信息
//
// 修改弹性资源池信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateElasticResourcePool(request *model.UpdateElasticResourcePoolRequest) (*model.UpdateElasticResourcePoolResponse, error) {
	requestDef := GenReqDefForUpdateElasticResourcePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateElasticResourcePoolResponse), nil
	}
}

// UpdateElasticResourcePoolInvoker 修改弹性资源池信息
func (c *DliClient) UpdateElasticResourcePoolInvoker(request *model.UpdateElasticResourcePoolRequest) *UpdateElasticResourcePoolInvoker {
	requestDef := GenReqDefForUpdateElasticResourcePool()
	return &UpdateElasticResourcePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateElasticResourcePoolQueueInfo 修改弹性资源池关联的队列优先级
//
// 设置弹性资源池指定队列的扩缩容策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateElasticResourcePoolQueueInfo(request *model.UpdateElasticResourcePoolQueueInfoRequest) (*model.UpdateElasticResourcePoolQueueInfoResponse, error) {
	requestDef := GenReqDefForUpdateElasticResourcePoolQueueInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateElasticResourcePoolQueueInfoResponse), nil
	}
}

// UpdateElasticResourcePoolQueueInfoInvoker 修改弹性资源池关联的队列优先级
func (c *DliClient) UpdateElasticResourcePoolQueueInfoInvoker(request *model.UpdateElasticResourcePoolQueueInfoRequest) *UpdateElasticResourcePoolQueueInfoInvoker {
	requestDef := GenReqDefForUpdateElasticResourcePoolQueueInfo()
	return &UpdateElasticResourcePoolQueueInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalValue 修改DLI全局变量
//
// 修改全局变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateGlobalValue(request *model.UpdateGlobalValueRequest) (*model.UpdateGlobalValueResponse, error) {
	requestDef := GenReqDefForUpdateGlobalValue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalValueResponse), nil
	}
}

// UpdateGlobalValueInvoker 修改DLI全局变量
func (c *DliClient) UpdateGlobalValueInvoker(request *model.UpdateGlobalValueRequest) *UpdateGlobalValueInvoker {
	requestDef := GenReqDefForUpdateGlobalValue()
	return &UpdateGlobalValueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupOrResourceOwner 修改组或者资源包拥有者
//
// 用于修改程序包的owner。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateGroupOrResourceOwner(request *model.UpdateGroupOrResourceOwnerRequest) (*model.UpdateGroupOrResourceOwnerResponse, error) {
	requestDef := GenReqDefForUpdateGroupOrResourceOwner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupOrResourceOwnerResponse), nil
	}
}

// UpdateGroupOrResourceOwnerInvoker 修改组或者资源包拥有者
func (c *DliClient) UpdateGroupOrResourceOwnerInvoker(request *model.UpdateGroupOrResourceOwnerRequest) *UpdateGroupOrResourceOwnerInvoker {
	requestDef := GenReqDefForUpdateGroupOrResourceOwner()
	return &UpdateGroupOrResourceOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateHostMassage 修改增强型跨源主机信息
//
// 该API用于在跨源中修改数据源主机信息，仅支持全量覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateHostMassage(request *model.UpdateHostMassageRequest) (*model.UpdateHostMassageResponse, error) {
	requestDef := GenReqDefForUpdateHostMassage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateHostMassageResponse), nil
	}
}

// UpdateHostMassageInvoker 修改增强型跨源主机信息
func (c *DliClient) UpdateHostMassageInvoker(request *model.UpdateHostMassageRequest) *UpdateHostMassageInvoker {
	requestDef := GenReqDefForUpdateHostMassage()
	return &UpdateHostMassageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateQueueCidr 修改队列网段
//
// 该功能用于修改包年包月队列网段。
// 说明：
// 如果待修改网段的队列中有正在提交或正在运行的作业，或者改队列已经绑定了增强型跨源，将不支持修改网段操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateQueueCidr(request *model.UpdateQueueCidrRequest) (*model.UpdateQueueCidrResponse, error) {
	requestDef := GenReqDefForUpdateQueueCidr()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQueueCidrResponse), nil
	}
}

// UpdateQueueCidrInvoker 修改队列网段
func (c *DliClient) UpdateQueueCidrInvoker(request *model.UpdateQueueCidrRequest) *UpdateQueueCidrInvoker {
	requestDef := GenReqDefForUpdateQueueCidr()
	return &UpdateQueueCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadFiles 上传file类型分组资源
//
// 该API用于在project下上传file类型模块。
// 说明： 上传同名file模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadFiles(request *model.UploadFilesRequest) (*model.UploadFilesResponse, error) {
	requestDef := GenReqDefForUploadFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFilesResponse), nil
	}
}

// UploadFilesInvoker 上传file类型分组资源
func (c *DliClient) UploadFilesInvoker(request *model.UploadFilesRequest) *UploadFilesInvoker {
	requestDef := GenReqDefForUploadFiles()
	return &UploadFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadJars 上传jar类型分组资源
//
// 该API用于在project下上传jar类型分组资源。
// 说明：上传同名资源模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadJars(request *model.UploadJarsRequest) (*model.UploadJarsResponse, error) {
	requestDef := GenReqDefForUploadJars()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJarsResponse), nil
	}
}

// UploadJarsInvoker 上传jar类型分组资源
func (c *DliClient) UploadJarsInvoker(request *model.UploadJarsRequest) *UploadJarsInvoker {
	requestDef := GenReqDefForUploadJars()
	return &UploadJarsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadPythonFiles 上传pyfile类型分组资源
//
// 该API用于在project下的上传pyfile类型模块。
// 说明： 上传同名pyfile类型模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadPythonFiles(request *model.UploadPythonFilesRequest) (*model.UploadPythonFilesResponse, error) {
	requestDef := GenReqDefForUploadPythonFiles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadPythonFilesResponse), nil
	}
}

// UploadPythonFilesInvoker 上传pyfile类型分组资源
func (c *DliClient) UploadPythonFilesInvoker(request *model.UploadPythonFilesRequest) *UploadPythonFilesInvoker {
	requestDef := GenReqDefForUploadPythonFiles()
	return &UploadPythonFilesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadResources 上传分组资源
//
// 该API用于上传分组资源到某个project下。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadResources(request *model.UploadResourcesRequest) (*model.UploadResourcesResponse, error) {
	requestDef := GenReqDefForUploadResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadResourcesResponse), nil
	}
}

// UploadResourcesInvoker 上传分组资源
func (c *DliClient) UploadResourcesInvoker(request *model.UploadResourcesRequest) *UploadResourcesInvoker {
	requestDef := GenReqDefForUploadResources()
	return &UploadResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteFlinkJobs 批量删除Flink作业
//
// 批量删除任何状态的作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) BatchDeleteFlinkJobs(request *model.BatchDeleteFlinkJobsRequest) (*model.BatchDeleteFlinkJobsResponse, error) {
	requestDef := GenReqDefForBatchDeleteFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteFlinkJobsResponse), nil
	}
}

// BatchDeleteFlinkJobsInvoker 批量删除Flink作业
func (c *DliClient) BatchDeleteFlinkJobsInvoker(request *model.BatchDeleteFlinkJobsRequest) *BatchDeleteFlinkJobsInvoker {
	requestDef := GenReqDefForBatchDeleteFlinkJobs()
	return &BatchDeleteFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRunFlinkJobs 批量运行Flink作业
//
// 触发批量运行Flink作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) BatchRunFlinkJobs(request *model.BatchRunFlinkJobsRequest) (*model.BatchRunFlinkJobsResponse, error) {
	requestDef := GenReqDefForBatchRunFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRunFlinkJobsResponse), nil
	}
}

// BatchRunFlinkJobsInvoker 批量运行Flink作业
func (c *DliClient) BatchRunFlinkJobsInvoker(request *model.BatchRunFlinkJobsRequest) *BatchRunFlinkJobsInvoker {
	requestDef := GenReqDefForBatchRunFlinkJobs()
	return &BatchRunFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeFlinkJobStatusReport 边缘Flink作业状态信息上报
//
// 该API用于处理边缘Flink作业状态上报信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ChangeFlinkJobStatusReport(request *model.ChangeFlinkJobStatusReportRequest) (*model.ChangeFlinkJobStatusReportResponse, error) {
	requestDef := GenReqDefForChangeFlinkJobStatusReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeFlinkJobStatusReportResponse), nil
	}
}

// ChangeFlinkJobStatusReportInvoker 边缘Flink作业状态信息上报
func (c *DliClient) ChangeFlinkJobStatusReportInvoker(request *model.ChangeFlinkJobStatusReportRequest) *ChangeFlinkJobStatusReportInvoker {
	requestDef := GenReqDefForChangeFlinkJobStatusReport()
	return &ChangeFlinkJobStatusReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkJar 新建Flink Jar作业
//
// 用户自定义作业目前支持jar格式，运行在独享集群中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkJar(request *model.CreateFlinkJarRequest) (*model.CreateFlinkJarResponse, error) {
	requestDef := GenReqDefForCreateFlinkJar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkJarResponse), nil
	}
}

// CreateFlinkJarInvoker 新建Flink Jar作业
func (c *DliClient) CreateFlinkJarInvoker(request *model.CreateFlinkJarRequest) *CreateFlinkJarInvoker {
	requestDef := GenReqDefForCreateFlinkJar()
	return &CreateFlinkJarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkSqlJob 新建Flink SQL作业
//
// 通过POST方式，提交流式SQL作业，请求体为JSON格式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkSqlJob(request *model.CreateFlinkSqlJobRequest) (*model.CreateFlinkSqlJobResponse, error) {
	requestDef := GenReqDefForCreateFlinkSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSqlJobResponse), nil
	}
}

// CreateFlinkSqlJobInvoker 新建Flink SQL作业
func (c *DliClient) CreateFlinkSqlJobInvoker(request *model.CreateFlinkSqlJobRequest) *CreateFlinkSqlJobInvoker {
	requestDef := GenReqDefForCreateFlinkSqlJob()
	return &CreateFlinkSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkTemplate 新建Flink作业模板
//
// 在DLI服务中新建一个Flink作业模板，最多100个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkTemplate(request *model.CreateFlinkTemplateRequest) (*model.CreateFlinkTemplateResponse, error) {
	requestDef := GenReqDefForCreateFlinkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkTemplateResponse), nil
	}
}

// CreateFlinkTemplateInvoker 新建Flink作业模板
func (c *DliClient) CreateFlinkTemplateInvoker(request *model.CreateFlinkTemplateRequest) *CreateFlinkTemplateInvoker {
	requestDef := GenReqDefForCreateFlinkTemplate()
	return &CreateFlinkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIefMessageChannel 创建IEF消息通道
//
// 该API用于创建IEF消息通道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateIefMessageChannel(request *model.CreateIefMessageChannelRequest) (*model.CreateIefMessageChannelResponse, error) {
	requestDef := GenReqDefForCreateIefMessageChannel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIefMessageChannelResponse), nil
	}
}

// CreateIefMessageChannelInvoker 创建IEF消息通道
func (c *DliClient) CreateIefMessageChannelInvoker(request *model.CreateIefMessageChannelRequest) *CreateIefMessageChannelInvoker {
	requestDef := GenReqDefForCreateIefMessageChannel()
	return &CreateIefMessageChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIefSystemEvents IEF系统事件上报
//
// 该API用于处理IEF系统事件上报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateIefSystemEvents(request *model.CreateIefSystemEventsRequest) (*model.CreateIefSystemEventsResponse, error) {
	requestDef := GenReqDefForCreateIefSystemEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIefSystemEventsResponse), nil
	}
}

// CreateIefSystemEventsInvoker IEF系统事件上报
func (c *DliClient) CreateIefSystemEventsInvoker(request *model.CreateIefSystemEventsRequest) *CreateIefSystemEventsInvoker {
	requestDef := GenReqDefForCreateIefSystemEvents()
	return &CreateIefSystemEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStreamGraph 生成flink SQL作业的静态流图
//
// 生成flink SQL作业的静态流图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateStreamGraph(request *model.CreateStreamGraphRequest) (*model.CreateStreamGraphResponse, error) {
	requestDef := GenReqDefForCreateStreamGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStreamGraphResponse), nil
	}
}

// CreateStreamGraphInvoker 生成flink SQL作业的静态流图
func (c *DliClient) CreateStreamGraphInvoker(request *model.CreateStreamGraphRequest) *CreateStreamGraphInvoker {
	requestDef := GenReqDefForCreateStreamGraph()
	return &CreateStreamGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlinkJob 删除作业
//
// 删除任何状态的作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteFlinkJob(request *model.DeleteFlinkJobRequest) (*model.DeleteFlinkJobResponse, error) {
	requestDef := GenReqDefForDeleteFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkJobResponse), nil
	}
}

// DeleteFlinkJobInvoker 删除作业
func (c *DliClient) DeleteFlinkJobInvoker(request *model.DeleteFlinkJobRequest) *DeleteFlinkJobInvoker {
	requestDef := GenReqDefForDeleteFlinkJob()
	return &DeleteFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlinkTemplate 删除Flink作业模板
//
// 删除一个Flink作业模板，即使当前模板正在被作业使用，也允许删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteFlinkTemplate(request *model.DeleteFlinkTemplateRequest) (*model.DeleteFlinkTemplateResponse, error) {
	requestDef := GenReqDefForDeleteFlinkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkTemplateResponse), nil
	}
}

// DeleteFlinkTemplateInvoker 删除Flink作业模板
func (c *DliClient) DeleteFlinkTemplateInvoker(request *model.DeleteFlinkTemplateRequest) *DeleteFlinkTemplateInvoker {
	requestDef := GenReqDefForDeleteFlinkTemplate()
	return &DeleteFlinkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportFlinkJob flink作业导出
//
// 通过POST方式，导出flink作业，请求体为JSON格式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ExportFlinkJob(request *model.ExportFlinkJobRequest) (*model.ExportFlinkJobResponse, error) {
	requestDef := GenReqDefForExportFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFlinkJobResponse), nil
	}
}

// ExportFlinkJobInvoker flink作业导出
func (c *DliClient) ExportFlinkJobInvoker(request *model.ExportFlinkJobRequest) *ExportFlinkJobInvoker {
	requestDef := GenReqDefForExportFlinkJob()
	return &ExportFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportFlinkJob flink作业导入
//
// 通过POST方式，导入flink作业，请求体为JSON格式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ImportFlinkJob(request *model.ImportFlinkJobRequest) (*model.ImportFlinkJobResponse, error) {
	requestDef := GenReqDefForImportFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFlinkJobResponse), nil
	}
}

// ImportFlinkJobInvoker flink作业导入
func (c *DliClient) ImportFlinkJobInvoker(request *model.ImportFlinkJobRequest) *ImportFlinkJobInvoker {
	requestDef := GenReqDefForImportFlinkJob()
	return &ImportFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobs 查询Flink作业列表
//
// 查询当前用户的作业列表，可以根据作业ID作为ID，查询大于ID或小于ID的限定条数的作业，默认查询全部状态的作业，也可以设定运行中或其他状态条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListFlinkJobs(request *model.ListFlinkJobsRequest) (*model.ListFlinkJobsResponse, error) {
	requestDef := GenReqDefForListFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobsResponse), nil
	}
}

// ListFlinkJobsInvoker 查询Flink作业列表
func (c *DliClient) ListFlinkJobsInvoker(request *model.ListFlinkJobsRequest) *ListFlinkJobsInvoker {
	requestDef := GenReqDefForListFlinkJobs()
	return &ListFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkTemplates 查询Flink作业模板列表
//
// 查询Flink作业模板列表。当前只支持查询用户自定义模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListFlinkTemplates(request *model.ListFlinkTemplatesRequest) (*model.ListFlinkTemplatesResponse, error) {
	requestDef := GenReqDefForListFlinkTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkTemplatesResponse), nil
	}
}

// ListFlinkTemplatesInvoker 查询Flink作业模板列表
func (c *DliClient) ListFlinkTemplatesInvoker(request *model.ListFlinkTemplatesRequest) *ListFlinkTemplatesInvoker {
	requestDef := GenReqDefForListFlinkTemplates()
	return &ListFlinkTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterBucket OBS授权给DLI服务
//
// 用户主动授权OBS桶的操作权限给DLI服务, 用于保存用户作业的checkpoint、作业的运行日志等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RegisterBucket(request *model.RegisterBucketRequest) (*model.RegisterBucketResponse, error) {
	requestDef := GenReqDefForRegisterBucket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterBucketResponse), nil
	}
}

// RegisterBucketInvoker OBS授权给DLI服务
func (c *DliClient) RegisterBucketInvoker(request *model.RegisterBucketRequest) *RegisterBucketInvoker {
	requestDef := GenReqDefForRegisterBucket()
	return &RegisterBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunIefJobActionCallBack 边缘Flink作业Action状态回调
//
// 该API用于处理IEF Flink作业action回调信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RunIefJobActionCallBack(request *model.RunIefJobActionCallBackRequest) (*model.RunIefJobActionCallBackResponse, error) {
	requestDef := GenReqDefForRunIefJobActionCallBack()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunIefJobActionCallBackResponse), nil
	}
}

// RunIefJobActionCallBackInvoker 边缘Flink作业Action状态回调
func (c *DliClient) RunIefJobActionCallBackInvoker(request *model.RunIefJobActionCallBackRequest) *RunIefJobActionCallBackInvoker {
	requestDef := GenReqDefForRunIefJobActionCallBack()
	return &RunIefJobActionCallBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkExecuteGraph 查询Flink作业执行计划
//
// 查询Flink作业执行计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowFlinkExecuteGraph(request *model.ShowFlinkExecuteGraphRequest) (*model.ShowFlinkExecuteGraphResponse, error) {
	requestDef := GenReqDefForShowFlinkExecuteGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkExecuteGraphResponse), nil
	}
}

// ShowFlinkExecuteGraphInvoker 查询Flink作业执行计划
func (c *DliClient) ShowFlinkExecuteGraphInvoker(request *model.ShowFlinkExecuteGraphRequest) *ShowFlinkExecuteGraphInvoker {
	requestDef := GenReqDefForShowFlinkExecuteGraph()
	return &ShowFlinkExecuteGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJob 查询Flink作业详情
//
// 查看一个Flink作业的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowFlinkJob(request *model.ShowFlinkJobRequest) (*model.ShowFlinkJobResponse, error) {
	requestDef := GenReqDefForShowFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobResponse), nil
	}
}

// ShowFlinkJobInvoker 查询Flink作业详情
func (c *DliClient) ShowFlinkJobInvoker(request *model.ShowFlinkJobRequest) *ShowFlinkJobInvoker {
	requestDef := GenReqDefForShowFlinkJob()
	return &ShowFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkMetric 查询Flink作业监控信息
//
// 查询Flink作业监控信息, 支持同时查询多个Flink作业的监控信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowFlinkMetric(request *model.ShowFlinkMetricRequest) (*model.ShowFlinkMetricResponse, error) {
	requestDef := GenReqDefForShowFlinkMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkMetricResponse), nil
	}
}

// ShowFlinkMetricInvoker 查询Flink作业监控信息
func (c *DliClient) ShowFlinkMetricInvoker(request *model.ShowFlinkMetricRequest) *ShowFlinkMetricInvoker {
	requestDef := GenReqDefForShowFlinkMetric()
	return &ShowFlinkMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopFlinkJobs 批量停止Flink作业
//
// 批量停止正在运行的Flink作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) StopFlinkJobs(request *model.StopFlinkJobsRequest) (*model.StopFlinkJobsResponse, error) {
	requestDef := GenReqDefForStopFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopFlinkJobsResponse), nil
	}
}

// StopFlinkJobsInvoker 批量停止Flink作业
func (c *DliClient) StopFlinkJobsInvoker(request *model.StopFlinkJobsRequest) *StopFlinkJobsInvoker {
	requestDef := GenReqDefForStopFlinkJobs()
	return &StopFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkJar 更新Flink Jar作业
//
// 更新用户自定义作业，目前支持jar格式，运行在独享集群中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkJar(request *model.UpdateFlinkJarRequest) (*model.UpdateFlinkJarResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJar()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJarResponse), nil
	}
}

// UpdateFlinkJarInvoker 更新Flink Jar作业
func (c *DliClient) UpdateFlinkJarInvoker(request *model.UpdateFlinkJarRequest) *UpdateFlinkJarInvoker {
	requestDef := GenReqDefForUpdateFlinkJar()
	return &UpdateFlinkJarInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkSql 更新Flink SQL作业
//
// Stream SQL的语法扩展了Apache Flink SQL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkSql(request *model.UpdateFlinkSqlRequest) (*model.UpdateFlinkSqlResponse, error) {
	requestDef := GenReqDefForUpdateFlinkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkSqlResponse), nil
	}
}

// UpdateFlinkSqlInvoker 更新Flink SQL作业
func (c *DliClient) UpdateFlinkSqlInvoker(request *model.UpdateFlinkSqlRequest) *UpdateFlinkSqlInvoker {
	requestDef := GenReqDefForUpdateFlinkSql()
	return &UpdateFlinkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkTemplate 更新Flink作业模板
//
// 对DLI服务中已有的Flink作业模板进行更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkTemplate(request *model.UpdateFlinkTemplateRequest) (*model.UpdateFlinkTemplateResponse, error) {
	requestDef := GenReqDefForUpdateFlinkTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkTemplateResponse), nil
	}
}

// UpdateFlinkTemplateInvoker 更新Flink作业模板
func (c *DliClient) UpdateFlinkTemplateInvoker(request *model.UpdateFlinkTemplateRequest) *UpdateFlinkTemplateInvoker {
	requestDef := GenReqDefForUpdateFlinkTemplate()
	return &UpdateFlinkTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelBatchJob 取消批处理作业
//
// 该API用于取消批处理作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CancelBatchJob(request *model.CancelBatchJobRequest) (*model.CancelBatchJobResponse, error) {
	requestDef := GenReqDefForCancelBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelBatchJobResponse), nil
	}
}

// CancelBatchJobInvoker 取消批处理作业
func (c *DliClient) CancelBatchJobInvoker(request *model.CancelBatchJobRequest) *CancelBatchJobInvoker {
	requestDef := GenReqDefForCancelBatchJob()
	return &CancelBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchJob 创建批处理作业
//
// 该API用于在某个队列上创建作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateBatchJob(request *model.CreateBatchJobRequest) (*model.CreateBatchJobResponse, error) {
	requestDef := GenReqDefForCreateBatchJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchJobResponse), nil
	}
}

// CreateBatchJobInvoker 创建批处理作业
func (c *DliClient) CreateBatchJobInvoker(request *model.CreateBatchJobRequest) *CreateBatchJobInvoker {
	requestDef := GenReqDefForCreateBatchJob()
	return &CreateBatchJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBatches 查询批处理作业列表
//
// 该API用于查询Project下某队列批处理作业的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListBatches(request *model.ListBatchesRequest) (*model.ListBatchesResponse, error) {
	requestDef := GenReqDefForListBatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchesResponse), nil
	}
}

// ListBatchesInvoker 查询批处理作业列表
func (c *DliClient) ListBatchesInvoker(request *model.ListBatchesRequest) *ListBatchesInvoker {
	requestDef := GenReqDefForListBatches()
	return &ListBatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchInfo 查询批处理作业详情
//
// 该API用于根据批处理作业的id查询作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowBatchInfo(request *model.ShowBatchInfoRequest) (*model.ShowBatchInfoResponse, error) {
	requestDef := GenReqDefForShowBatchInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchInfoResponse), nil
	}
}

// ShowBatchInfoInvoker 查询批处理作业详情
func (c *DliClient) ShowBatchInfoInvoker(request *model.ShowBatchInfoRequest) *ShowBatchInfoInvoker {
	requestDef := GenReqDefForShowBatchInfo()
	return &ShowBatchInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchLog 查询批处理作业日志
//
// 该API用于查询批处理作业的后台日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowBatchLog(request *model.ShowBatchLogRequest) (*model.ShowBatchLogResponse, error) {
	requestDef := GenReqDefForShowBatchLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchLogResponse), nil
	}
}

// ShowBatchLogInvoker 查询批处理作业日志
func (c *DliClient) ShowBatchLogInvoker(request *model.ShowBatchLogRequest) *ShowBatchLogInvoker {
	requestDef := GenReqDefForShowBatchLog()
	return &ShowBatchLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBatchState 查询批处理作业状态
//
// 该API用于查询批处理作业的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowBatchState(request *model.ShowBatchStateRequest) (*model.ShowBatchStateResponse, error) {
	requestDef := GenReqDefForShowBatchState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBatchStateResponse), nil
	}
}

// ShowBatchStateInvoker 查询批处理作业状态
func (c *DliClient) ShowBatchStateInvoker(request *model.ShowBatchStateRequest) *ShowBatchStateInvoker {
	requestDef := GenReqDefForShowBatchState()
	return &ShowBatchStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSqlJob 取消作业
//
// 该API用于取消已经提交的作业，若作业已经执行结束或失败则无法取消。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CancelSqlJob(request *model.CancelSqlJobRequest) (*model.CancelSqlJobResponse, error) {
	requestDef := GenReqDefForCancelSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSqlJobResponse), nil
	}
}

// CancelSqlJobInvoker 取消作业
func (c *DliClient) CancelSqlJobInvoker(request *model.CancelSqlJobRequest) *CancelSqlJobInvoker {
	requestDef := GenReqDefForCancelSqlJob()
	return &CancelSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckSql 检查SQL语法
//
// 该API用于检查SQL语法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CheckSql(request *model.CheckSqlRequest) (*model.CheckSqlResponse, error) {
	requestDef := GenReqDefForCheckSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckSqlResponse), nil
	}
}

// CheckSqlInvoker 检查SQL语法
func (c *DliClient) CheckSqlInvoker(request *model.CheckSqlRequest) *CheckSqlInvoker {
	requestDef := GenReqDefForCheckSql()
	return &CheckSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDatabase 创建数据库
//
// 该API用于新增数据库。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateDatabase(request *model.CreateDatabaseRequest) (*model.CreateDatabaseResponse, error) {
	requestDef := GenReqDefForCreateDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseResponse), nil
	}
}

// CreateDatabaseInvoker 创建数据库
func (c *DliClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSqlJob 提交SQL作业
//
// 该API用于通过执行SQL语句的方式向队列提交作业。
//
// 作业包含以下类型：DDL、DCL、IMPORT、QUERY和INSERT。其中，IMPORT与导入数据的功能一致，区别仅在于实现方式不同。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateSqlJob(request *model.CreateSqlJobRequest) (*model.CreateSqlJobResponse, error) {
	requestDef := GenReqDefForCreateSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlJobResponse), nil
	}
}

// CreateSqlJobInvoker 提交SQL作业
func (c *DliClient) CreateSqlJobInvoker(request *model.CreateSqlJobRequest) *CreateSqlJobInvoker {
	requestDef := GenReqDefForCreateSqlJob()
	return &CreateSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTable 创建表
//
// 该API用于创建新的表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateTable(request *model.CreateTableRequest) (*model.CreateTableResponse, error) {
	requestDef := GenReqDefForCreateTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableResponse), nil
	}
}

// CreateTableInvoker 创建表
func (c *DliClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDatabase 删除数据库
//
// 该API用于删除空数据库，若待删除的数据库中存在表，则需先删除其中的所有表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteDatabase(request *model.DeleteDatabaseRequest) (*model.DeleteDatabaseResponse, error) {
	requestDef := GenReqDefForDeleteDatabase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDatabaseResponse), nil
	}
}

// DeleteDatabaseInvoker 删除数据库
func (c *DliClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTable 删除表
//
// 该API用于删除指定的表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteTable(request *model.DeleteTableRequest) (*model.DeleteTableResponse, error) {
	requestDef := GenReqDefForDeleteTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableResponse), nil
	}
}

// DeleteTableInvoker 删除表
func (c *DliClient) DeleteTableInvoker(request *model.DeleteTableRequest) *DeleteTableInvoker {
	requestDef := GenReqDefForDeleteTable()
	return &DeleteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportSqlJobResult 导出查询结果
//
// 该API用于将SQL语句的查询结果导出到OBS对象存储中，只支持导出“QUERY”类型作业的查询结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ExportSqlJobResult(request *model.ExportSqlJobResultRequest) (*model.ExportSqlJobResultResponse, error) {
	requestDef := GenReqDefForExportSqlJobResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportSqlJobResultResponse), nil
	}
}

// ExportSqlJobResultInvoker 导出查询结果
func (c *DliClient) ExportSqlJobResultInvoker(request *model.ExportSqlJobResultRequest) *ExportSqlJobResultInvoker {
	requestDef := GenReqDefForExportSqlJobResult()
	return &ExportSqlJobResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportTable 导出查询结果
//
// 该API用于将SQL语句的查询结果导出到OBS对象存储中，只支持导出“QUERY”类型作业的查询结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ExportTable(request *model.ExportTableRequest) (*model.ExportTableResponse, error) {
	requestDef := GenReqDefForExportTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportTableResponse), nil
	}
}

// ExportTableInvoker 导出查询结果
func (c *DliClient) ExportTableInvoker(request *model.ExportTableRequest) *ExportTableInvoker {
	requestDef := GenReqDefForExportTable()
	return &ExportTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportTable 导入数据
//
// 该API用于将数据从文件导入DLI或OBS表，目前仅支持将OBS上的数据导入DLI或OBS中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ImportTable(request *model.ImportTableRequest) (*model.ImportTableResponse, error) {
	requestDef := GenReqDefForImportTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportTableResponse), nil
	}
}

// ImportTableInvoker 导入数据
func (c *DliClient) ImportTableInvoker(request *model.ImportTableRequest) *ImportTableInvoker {
	requestDef := GenReqDefForImportTable()
	return &ImportTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTables 查询所有表
//
// 该API用于查询指定数据库下符合过滤条件的或所有的表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListAllTables(request *model.ListAllTablesRequest) (*model.ListAllTablesResponse, error) {
	requestDef := GenReqDefForListAllTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTablesResponse), nil
	}
}

// ListAllTablesInvoker 查询所有表
func (c *DliClient) ListAllTablesInvoker(request *model.ListAllTablesRequest) *ListAllTablesInvoker {
	requestDef := GenReqDefForListAllTables()
	return &ListAllTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabases 查询所有数据库
//
// 该API用于查询出所有的数据库信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 查询所有数据库
func (c *DliClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSqlJobs 查询所有作业
//
// 该API用于查询当前工程下面的所有作业的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListSqlJobs(request *model.ListSqlJobsRequest) (*model.ListSqlJobsResponse, error) {
	requestDef := GenReqDefForListSqlJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlJobsResponse), nil
	}
}

// ListSqlJobsInvoker 查询所有作业
func (c *DliClient) ListSqlJobsInvoker(request *model.ListSqlJobsRequest) *ListSqlJobsInvoker {
	requestDef := GenReqDefForListSqlJobs()
	return &ListSqlJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PreviewJobResult 预览SQL作业查询结果
//
// 该API用于在执行SQL查询语句的作业完成后，查看该作业执行的结果。目前仅支持查看“QUERY”类型作业的执行结果。
// 该API只能查看前1000条的结果记录，若要查看全部的结果记录，需要先导出查询结果再进行查看。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) PreviewJobResult(request *model.PreviewJobResultRequest) (*model.PreviewJobResultResponse, error) {
	requestDef := GenReqDefForPreviewJobResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewJobResultResponse), nil
	}
}

// PreviewJobResultInvoker 预览SQL作业查询结果
func (c *DliClient) PreviewJobResultInvoker(request *model.PreviewJobResultRequest) *PreviewJobResultInvoker {
	requestDef := GenReqDefForPreviewJobResult()
	return &PreviewJobResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDescribeTable 描述表信息
//
// 该API用于描述指定表的元数据信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowDescribeTable(request *model.ShowDescribeTableRequest) (*model.ShowDescribeTableResponse, error) {
	requestDef := GenReqDefForShowDescribeTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDescribeTableResponse), nil
	}
}

// ShowDescribeTableInvoker 描述表信息
func (c *DliClient) ShowDescribeTableInvoker(request *model.ShowDescribeTableRequest) *ShowDescribeTableInvoker {
	requestDef := GenReqDefForShowDescribeTable()
	return &ShowDescribeTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowJobProgress 查询作业执行进度信息
//
// 该API用于获取作业执行进度信息，如果作业正在执行，可以获取到子作业的信息，如果作业刚开始或者已经结束，不可以获取到子作业信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowJobProgress(request *model.ShowJobProgressRequest) (*model.ShowJobProgressResponse, error) {
	requestDef := GenReqDefForShowJobProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobProgressResponse), nil
	}
}

// ShowJobProgressInvoker 查询作业执行进度信息
func (c *DliClient) ShowJobProgressInvoker(request *model.ShowJobProgressRequest) *ShowJobProgressInvoker {
	requestDef := GenReqDefForShowJobProgress()
	return &ShowJobProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPartitions 获取分区信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowPartitions(request *model.ShowPartitionsRequest) (*model.ShowPartitionsResponse, error) {
	requestDef := GenReqDefForShowPartitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPartitionsResponse), nil
	}
}

// ShowPartitionsInvoker 获取分区信息列表
func (c *DliClient) ShowPartitionsInvoker(request *model.ShowPartitionsRequest) *ShowPartitionsInvoker {
	requestDef := GenReqDefForShowPartitions()
	return &ShowPartitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlJobDetail 查询作业详细信息
//
// 该API用于查询作业的详细信息，如作业的databasename、tablename、file size和export mode等信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSqlJobDetail(request *model.ShowSqlJobDetailRequest) (*model.ShowSqlJobDetailResponse, error) {
	requestDef := GenReqDefForShowSqlJobDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobDetailResponse), nil
	}
}

// ShowSqlJobDetailInvoker 查询作业详细信息
func (c *DliClient) ShowSqlJobDetailInvoker(request *model.ShowSqlJobDetailRequest) *ShowSqlJobDetailInvoker {
	requestDef := GenReqDefForShowSqlJobDetail()
	return &ShowSqlJobDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSqlJobStatus 查询作业状态
//
// 该API用于在作业提交后查询作业状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSqlJobStatus(request *model.ShowSqlJobStatusRequest) (*model.ShowSqlJobStatusResponse, error) {
	requestDef := GenReqDefForShowSqlJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobStatusResponse), nil
	}
}

// ShowSqlJobStatusInvoker 查询作业状态
func (c *DliClient) ShowSqlJobStatusInvoker(request *model.ShowSqlJobStatusRequest) *ShowSqlJobStatusInvoker {
	requestDef := GenReqDefForShowSqlJobStatus()
	return &ShowSqlJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTableContent 预览表内容
//
// 该API用于用于预览表中前10行的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowTableContent(request *model.ShowTableContentRequest) (*model.ShowTableContentResponse, error) {
	requestDef := GenReqDefForShowTableContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableContentResponse), nil
	}
}

// ShowTableContentInvoker 预览表内容
func (c *DliClient) ShowTableContentInvoker(request *model.ShowTableContentRequest) *ShowTableContentInvoker {
	requestDef := GenReqDefForShowTableContent()
	return &ShowTableContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDatabaseOwner 修改数据库用户
//
// 用于修改数据库的owner。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateDatabaseOwner(request *model.UpdateDatabaseOwnerRequest) (*model.UpdateDatabaseOwnerResponse, error) {
	requestDef := GenReqDefForUpdateDatabaseOwner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDatabaseOwnerResponse), nil
	}
}

// UpdateDatabaseOwnerInvoker 修改数据库用户
func (c *DliClient) UpdateDatabaseOwnerInvoker(request *model.UpdateDatabaseOwnerRequest) *UpdateDatabaseOwnerInvoker {
	requestDef := GenReqDefForUpdateDatabaseOwner()
	return &UpdateDatabaseOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTableOwner 修改表用户
//
// 用于修改表的owner。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateTableOwner(request *model.UpdateTableOwnerRequest) (*model.UpdateTableOwnerResponse, error) {
	requestDef := GenReqDefForUpdateTableOwner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTableOwnerResponse), nil
	}
}

// UpdateTableOwnerInvoker 修改表用户
func (c *DliClient) UpdateTableOwnerInvoker(request *model.UpdateTableOwnerRequest) *UpdateTableOwnerInvoker {
	requestDef := GenReqDefForUpdateTableOwner()
	return &UpdateTableOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
