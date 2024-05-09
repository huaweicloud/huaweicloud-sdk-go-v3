package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dli/v1/model"
)

type DliClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDliClient(hcClient *httpclient.HcHttpClient) *DliClient {
	return &DliClient{HcClient: hcClient}
}

func DliClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
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

// AssociateQueueToEnhancedConnection 绑定队列
//
// 该API用于在已创建的增强型跨源中绑定队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) AssociateQueueToEnhancedConnection(request *model.AssociateQueueToEnhancedConnectionRequest) (*model.AssociateQueueToEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForAssociateQueueToEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateQueueToEnhancedConnectionResponse), nil
	}
}

// AssociateQueueToEnhancedConnectionInvoker 绑定队列
func (c *DliClient) AssociateQueueToEnhancedConnectionInvoker(request *model.AssociateQueueToEnhancedConnectionRequest) *AssociateQueueToEnhancedConnectionInvoker {
	requestDef := GenReqDefForAssociateQueueToEnhancedConnection()
	return &AssociateQueueToEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// BatchDeleteQueuePlansInvoker 批量删除队列定时扩缩容计划
func (c *DliClient) BatchDeleteQueuePlansInvoker(request *model.BatchDeleteQueuePlansRequest) *BatchDeleteQueuePlansInvoker {
	requestDef := GenReqDefForBatchDeleteQueuePlans()
	return &BatchDeleteQueuePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateAuthInfoInvoker 创建跨源认证
func (c *DliClient) CreateAuthInfoInvoker(request *model.CreateAuthInfoRequest) *CreateAuthInfoInvoker {
	requestDef := GenReqDefForCreateAuthInfo()
	return &CreateAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnectivityTask 创建地址连通性请求
//
// 创建地址连通性请求API接口，往指定集群发送地址连通性测试请求，并将测试地址插入表内
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateConnectivityTask(request *model.CreateConnectivityTaskRequest) (*model.CreateConnectivityTaskResponse, error) {
	requestDef := GenReqDefForCreateConnectivityTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectivityTaskResponse), nil
	}
}

// CreateConnectivityTaskInvoker 创建地址连通性请求
func (c *DliClient) CreateConnectivityTaskInvoker(request *model.CreateConnectivityTaskRequest) *CreateConnectivityTaskInvoker {
	requestDef := GenReqDefForCreateConnectivityTask()
	return &CreateConnectivityTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateDatasourceConnectionInvoker 创建经典型跨源连接
func (c *DliClient) CreateDatasourceConnectionInvoker(request *model.CreateDatasourceConnectionRequest) *CreateDatasourceConnectionInvoker {
	requestDef := GenReqDefForCreateDatasourceConnection()
	return &CreateDatasourceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateEnhancedConnectionRoutesInvoker 创建路由
func (c *DliClient) CreateEnhancedConnectionRoutesInvoker(request *model.CreateEnhancedConnectionRoutesRequest) *CreateEnhancedConnectionRoutesInvoker {
	requestDef := GenReqDefForCreateEnhancedConnectionRoutes()
	return &CreateEnhancedConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalVariable 创建DLI全局变量
//
// 创建全局变量。全局变量用于替换SQL作业中的敏感数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateGlobalVariable(request *model.CreateGlobalVariableRequest) (*model.CreateGlobalVariableResponse, error) {
	requestDef := GenReqDefForCreateGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalVariableResponse), nil
	}
}

// CreateGlobalVariableInvoker 创建DLI全局变量
func (c *DliClient) CreateGlobalVariableInvoker(request *model.CreateGlobalVariableRequest) *CreateGlobalVariableInvoker {
	requestDef := GenReqDefForCreateGlobalVariable()
	return &CreateGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateJobAuthInfo 创建跨源认证
//
// 该API用于创建跨源认证。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateJobAuthInfo(request *model.CreateJobAuthInfoRequest) (*model.CreateJobAuthInfoResponse, error) {
	requestDef := GenReqDefForCreateJobAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateJobAuthInfoResponse), nil
	}
}

// CreateJobAuthInfoInvoker 创建跨源认证
func (c *DliClient) CreateJobAuthInfoInvoker(request *model.CreateJobAuthInfoRequest) *CreateJobAuthInfoInvoker {
	requestDef := GenReqDefForCreateJobAuthInfo()
	return &CreateJobAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateQueuePlanInvoker 创建队列定时扩缩容计划
func (c *DliClient) CreateQueuePlanInvoker(request *model.CreateQueuePlanRequest) *CreateQueuePlanInvoker {
	requestDef := GenReqDefForCreateQueuePlan()
	return &CreateQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateQueueProperty 新增队列属性
//
// 该接口用于增加队列属性, 一次可增加多个属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateQueueProperty(request *model.CreateQueuePropertyRequest) (*model.CreateQueuePropertyResponse, error) {
	requestDef := GenReqDefForCreateQueueProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateQueuePropertyResponse), nil
	}
}

// CreateQueuePropertyInvoker 新增队列属性
func (c *DliClient) CreateQueuePropertyInvoker(request *model.CreateQueuePropertyRequest) *CreateQueuePropertyInvoker {
	requestDef := GenReqDefForCreateQueueProperty()
	return &CreateQueuePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRouteToEnhancedConnection 创建路由
//
// 该API用于创建跨源需要的路由。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateRouteToEnhancedConnection(request *model.CreateRouteToEnhancedConnectionRequest) (*model.CreateRouteToEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForCreateRouteToEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRouteToEnhancedConnectionResponse), nil
	}
}

// CreateRouteToEnhancedConnectionInvoker 创建路由
func (c *DliClient) CreateRouteToEnhancedConnectionInvoker(request *model.CreateRouteToEnhancedConnectionRequest) *CreateRouteToEnhancedConnectionInvoker {
	requestDef := GenReqDefForCreateRouteToEnhancedConnection()
	return &CreateRouteToEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteAuthInfoInvoker 删除跨源认证
func (c *DliClient) DeleteAuthInfoInvoker(request *model.DeleteAuthInfoRequest) *DeleteAuthInfoInvoker {
	requestDef := GenReqDefForDeleteAuthInfo()
	return &DeleteAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteEnhancedConnectionRoutesInvoker 删除路由
func (c *DliClient) DeleteEnhancedConnectionRoutesInvoker(request *model.DeleteEnhancedConnectionRoutesRequest) *DeleteEnhancedConnectionRoutesInvoker {
	requestDef := GenReqDefForDeleteEnhancedConnectionRoutes()
	return &DeleteEnhancedConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalVariable 删除DLI全局变量
//
// 删除全局变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteGlobalVariable(request *model.DeleteGlobalVariableRequest) (*model.DeleteGlobalVariableResponse, error) {
	requestDef := GenReqDefForDeleteGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalVariableResponse), nil
	}
}

// DeleteGlobalVariableInvoker 删除DLI全局变量
func (c *DliClient) DeleteGlobalVariableInvoker(request *model.DeleteGlobalVariableRequest) *DeleteGlobalVariableInvoker {
	requestDef := GenReqDefForDeleteGlobalVariable()
	return &DeleteGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteJobAuthInfo 删除跨源认证
//
// 该API用于删除跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteJobAuthInfo(request *model.DeleteJobAuthInfoRequest) (*model.DeleteJobAuthInfoResponse, error) {
	requestDef := GenReqDefForDeleteJobAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobAuthInfoResponse), nil
	}
}

// DeleteJobAuthInfoInvoker 删除跨源认证
func (c *DliClient) DeleteJobAuthInfoInvoker(request *model.DeleteJobAuthInfoRequest) *DeleteJobAuthInfoInvoker {
	requestDef := GenReqDefForDeleteJobAuthInfo()
	return &DeleteJobAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteJobResource 删除组内资源包
//
// 该API用于删除某个project某个分组下的资源包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteJobResource(request *model.DeleteJobResourceRequest) (*model.DeleteJobResourceResponse, error) {
	requestDef := GenReqDefForDeleteJobResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteJobResourceResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteJobResourceInvoker 删除组内资源包
func (c *DliClient) DeleteJobResourceInvoker(request *model.DeleteJobResourceRequest) *DeleteJobResourceInvoker {
	requestDef := GenReqDefForDeleteJobResource()
	return &DeleteJobResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteQueuePlanInvoker 单个删除队列定时扩缩容计划
func (c *DliClient) DeleteQueuePlanInvoker(request *model.DeleteQueuePlanRequest) *DeleteQueuePlanInvoker {
	requestDef := GenReqDefForDeleteQueuePlan()
	return &DeleteQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteQueueProperty 删除队列的属性
//
// 该接口用于删除队列的属性，一次可删除多个属性值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteQueueProperty(request *model.DeleteQueuePropertyRequest) (*model.DeleteQueuePropertyResponse, error) {
	requestDef := GenReqDefForDeleteQueueProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteQueuePropertyResponse), nil
	}
}

// DeleteQueuePropertyInvoker 删除队列的属性
func (c *DliClient) DeleteQueuePropertyInvoker(request *model.DeleteQueuePropertyRequest) *DeleteQueuePropertyInvoker {
	requestDef := GenReqDefForDeleteQueueProperty()
	return &DeleteQueuePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteRouteFromEnhancedConnection 删除路由
//
// 该API用于删除跨源需要的路由。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteRouteFromEnhancedConnection(request *model.DeleteRouteFromEnhancedConnectionRequest) (*model.DeleteRouteFromEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForDeleteRouteFromEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRouteFromEnhancedConnectionResponse), nil
	}
}

// DeleteRouteFromEnhancedConnectionInvoker 删除路由
func (c *DliClient) DeleteRouteFromEnhancedConnectionInvoker(request *model.DeleteRouteFromEnhancedConnectionRequest) *DeleteRouteFromEnhancedConnectionInvoker {
	requestDef := GenReqDefForDeleteRouteFromEnhancedConnection()
	return &DeleteRouteFromEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateQueueFromEnhancedConnection 解绑队列
//
// 该API用于在增强型跨源中解绑已绑定的队列。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DisassociateQueueFromEnhancedConnection(request *model.DisassociateQueueFromEnhancedConnectionRequest) (*model.DisassociateQueueFromEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForDisassociateQueueFromEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateQueueFromEnhancedConnectionResponse), nil
	}
}

// DisassociateQueueFromEnhancedConnectionInvoker 解绑队列
func (c *DliClient) DisassociateQueueFromEnhancedConnectionInvoker(request *model.DisassociateQueueFromEnhancedConnectionRequest) *DisassociateQueueFromEnhancedConnectionInvoker {
	requestDef := GenReqDefForDisassociateQueueFromEnhancedConnection()
	return &DisassociateQueueFromEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListAuthInfoInvoker 获取跨源认证列表
func (c *DliClient) ListAuthInfoInvoker(request *model.ListAuthInfoRequest) *ListAuthInfoInvoker {
	requestDef := GenReqDefForListAuthInfo()
	return &ListAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizationPrivileges 查看赋权对象的用者权限信息
//
// 获取对象赋权用户的权限信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListAuthorizationPrivileges(request *model.ListAuthorizationPrivilegesRequest) (*model.ListAuthorizationPrivilegesResponse, error) {
	requestDef := GenReqDefForListAuthorizationPrivileges()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizationPrivilegesResponse), nil
	}
}

// ListAuthorizationPrivilegesInvoker 查看赋权对象的用者权限信息
func (c *DliClient) ListAuthorizationPrivilegesInvoker(request *model.ListAuthorizationPrivilegesRequest) *ListAuthorizationPrivilegesInvoker {
	requestDef := GenReqDefForListAuthorizationPrivileges()
	return &ListAuthorizationPrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogs 获取项目下所有catalog信息
//
// 该API获取指定项目下所有catalog信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListCatalogs(request *model.ListCatalogsRequest) (*model.ListCatalogsResponse, error) {
	requestDef := GenReqDefForListCatalogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogsResponse), nil
	}
}

// ListCatalogsInvoker 获取项目下所有catalog信息
func (c *DliClient) ListCatalogsInvoker(request *model.ListCatalogsRequest) *ListCatalogsInvoker {
	requestDef := GenReqDefForListCatalogs()
	return &ListCatalogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListDatabaseUsersInvoker 查看数据库的使用者
func (c *DliClient) ListDatabaseUsersInvoker(request *model.ListDatabaseUsersRequest) *ListDatabaseUsersInvoker {
	requestDef := GenReqDefForListDatabaseUsers()
	return &ListDatabaseUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// ListGlobalVariables 查询DLI全局变量列表
//
// 查询全局变量列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListGlobalVariables(request *model.ListGlobalVariablesRequest) (*model.ListGlobalVariablesResponse, error) {
	requestDef := GenReqDefForListGlobalVariables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalVariablesResponse), nil
	}
}

// ListGlobalVariablesInvoker 查询DLI全局变量列表
func (c *DliClient) ListGlobalVariablesInvoker(request *model.ListGlobalVariablesRequest) *ListGlobalVariablesInvoker {
	requestDef := GenReqDefForListGlobalVariables()
	return &ListGlobalVariablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobAuthInfos 查询增强型跨源授权信息
//
// 该API用于查询跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListJobAuthInfos(request *model.ListJobAuthInfosRequest) (*model.ListJobAuthInfosResponse, error) {
	requestDef := GenReqDefForListJobAuthInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobAuthInfosResponse), nil
	}
}

// ListJobAuthInfosInvoker 查询增强型跨源授权信息
func (c *DliClient) ListJobAuthInfosInvoker(request *model.ListJobAuthInfosRequest) *ListJobAuthInfosInvoker {
	requestDef := GenReqDefForListJobAuthInfos()
	return &ListJobAuthInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListJobResources 查看分组资源列表
//
// 该API用于查看某个project下的所有资源，其中包含Group。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListJobResources(request *model.ListJobResourcesRequest) (*model.ListJobResourcesResponse, error) {
	requestDef := GenReqDefForListJobResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobResourcesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListJobResourcesInvoker 查看分组资源列表
func (c *DliClient) ListJobResourcesInvoker(request *model.ListJobResourcesRequest) *ListJobResourcesInvoker {
	requestDef := GenReqDefForListJobResources()
	return &ListJobResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListQueuePlansInvoker 查看队列定时扩缩容计划
func (c *DliClient) ListQueuePlansInvoker(request *model.ListQueuePlansRequest) *ListQueuePlansInvoker {
	requestDef := GenReqDefForListQueuePlans()
	return &ListQueuePlansInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueueProperties 获取队列属性
//
// 获取队列配置的属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListQueueProperties(request *model.ListQueuePropertiesRequest) (*model.ListQueuePropertiesResponse, error) {
	requestDef := GenReqDefForListQueueProperties()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueuePropertiesResponse), nil
	}
}

// ListQueuePropertiesInvoker 获取队列属性
func (c *DliClient) ListQueuePropertiesInvoker(request *model.ListQueuePropertiesRequest) *ListQueuePropertiesInvoker {
	requestDef := GenReqDefForListQueueProperties()
	return &ListQueuePropertiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTablePrivilegesInvoker 查看表的用户权限
func (c *DliClient) ListTablePrivilegesInvoker(request *model.ListTablePrivilegesRequest) *ListTablePrivilegesInvoker {
	requestDef := GenReqDefForListTablePrivileges()
	return &ListTablePrivilegesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTableUsersInvoker 查看表的使用者
func (c *DliClient) ListTableUsersInvoker(request *model.ListTableUsersRequest) *ListTableUsersInvoker {
	requestDef := GenReqDefForListTableUsers()
	return &ListTableUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// RegisterAuthorizedQueueInvoker 队列赋权
func (c *DliClient) RegisterAuthorizedQueueInvoker(request *model.RegisterAuthorizedQueueRequest) *RegisterAuthorizedQueueInvoker {
	requestDef := GenReqDefForRegisterAuthorizedQueue()
	return &RegisterAuthorizedQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunAuthorizationAction 数据赋权（用户、项目）
//
// 该API用于将DLI资源权限赋给、回收、更新指定的其他用户或项目。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RunAuthorizationAction(request *model.RunAuthorizationActionRequest) (*model.RunAuthorizationActionResponse, error) {
	requestDef := GenReqDefForRunAuthorizationAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAuthorizationActionResponse), nil
	}
}

// RunAuthorizationActionInvoker 数据赋权（用户、项目）
func (c *DliClient) RunAuthorizationActionInvoker(request *model.RunAuthorizationActionRequest) *RunAuthorizationActionInvoker {
	requestDef := GenReqDefForRunAuthorizationAction()
	return &RunAuthorizationActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunCatalogAction 绑定/解绑catalog映射信息
//
// 该API创建DLI绑定/解绑到lakeformation等服务的元数据目录（CATALOG）相关信息。
// 包含DLI侧CATALOG名称、外部CATALOG名称和类型，类型为预留字段，当前只支持lakeformation。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RunCatalogAction(request *model.RunCatalogActionRequest) (*model.RunCatalogActionResponse, error) {
	requestDef := GenReqDefForRunCatalogAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunCatalogActionResponse), nil
	}
}

// RunCatalogActionInvoker 绑定/解绑catalog映射信息
func (c *DliClient) RunCatalogActionInvoker(request *model.RunCatalogActionRequest) *RunCatalogActionInvoker {
	requestDef := GenReqDefForRunCatalogAction()
	return &RunCatalogActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// RunDataAuthorizationAction 数据赋权（用户）
//
// 该API用于将数据库或数据表的数据权限赋给指定的其他用户。
// 说明：
// 被赋权用户所在用户组的所属区域需具有Tenant Guest权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) RunDataAuthorizationAction(request *model.RunDataAuthorizationActionRequest) (*model.RunDataAuthorizationActionResponse, error) {
	requestDef := GenReqDefForRunDataAuthorizationAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDataAuthorizationActionResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// RunDataAuthorizationActionInvoker 数据赋权（用户）
func (c *DliClient) RunDataAuthorizationActionInvoker(request *model.RunDataAuthorizationActionRequest) *RunDataAuthorizationActionInvoker {
	requestDef := GenReqDefForRunDataAuthorizationAction()
	return &RunDataAuthorizationActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowCatalog 描述catalog信息
//
// 该API用于描述DLI catalog详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowCatalog(request *model.ShowCatalogRequest) (*model.ShowCatalogResponse, error) {
	requestDef := GenReqDefForShowCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCatalogResponse), nil
	}
}

// ShowCatalogInvoker 描述catalog信息
func (c *DliClient) ShowCatalogInvoker(request *model.ShowCatalogRequest) *ShowCatalogInvoker {
	requestDef := GenReqDefForShowCatalog()
	return &ShowCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConnectivityTask 查询指定地址连通性测试详情
//
// 该API用于在连通性测试提交后查询连通性结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowConnectivityTask(request *model.ShowConnectivityTaskRequest) (*model.ShowConnectivityTaskResponse, error) {
	requestDef := GenReqDefForShowConnectivityTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConnectivityTaskResponse), nil
	}
}

// ShowConnectivityTaskInvoker 查询指定地址连通性测试详情
func (c *DliClient) ShowConnectivityTaskInvoker(request *model.ShowConnectivityTaskRequest) *ShowConnectivityTaskInvoker {
	requestDef := GenReqDefForShowConnectivityTask()
	return &ShowConnectivityTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowDatasourceConnectionInvoker 查询经典型跨源连接
func (c *DliClient) ShowDatasourceConnectionInvoker(request *model.ShowDatasourceConnectionRequest) *ShowDatasourceConnectionInvoker {
	requestDef := GenReqDefForShowDatasourceConnection()
	return &ShowDatasourceConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// ShowEnhancedConnectionPrivilege 查询增强型跨源授权信息
//
// 该API用于查询增强型跨源连接授权的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowEnhancedConnectionPrivilege(request *model.ShowEnhancedConnectionPrivilegeRequest) (*model.ShowEnhancedConnectionPrivilegeResponse, error) {
	requestDef := GenReqDefForShowEnhancedConnectionPrivilege()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEnhancedConnectionPrivilegeResponse), nil
	}
}

// ShowEnhancedConnectionPrivilegeInvoker 查询增强型跨源授权信息
func (c *DliClient) ShowEnhancedConnectionPrivilegeInvoker(request *model.ShowEnhancedConnectionPrivilegeRequest) *ShowEnhancedConnectionPrivilegeInvoker {
	requestDef := GenReqDefForShowEnhancedConnectionPrivilege()
	return &ShowEnhancedConnectionPrivilegeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowJobResource 查看组内资源包
//
// 该API用于查看某个project某个分组下的具体资源信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowJobResource(request *model.ShowJobResourceRequest) (*model.ShowJobResourceResponse, error) {
	requestDef := GenReqDefForShowJobResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowJobResourceResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowJobResourceInvoker 查看组内资源包
func (c *DliClient) ShowJobResourceInvoker(request *model.ShowJobResourceRequest) *ShowJobResourceInvoker {
	requestDef := GenReqDefForShowJobResource()
	return &ShowJobResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQueue 查询队列详情
//
// 该API用于列出该project下指定的队列详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowQueue(request *model.ShowQueueRequest) (*model.ShowQueueResponse, error) {
	requestDef := GenReqDefForShowQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQueueResponse), nil
	}
}

// ShowQueueInvoker 查询队列详情
func (c *DliClient) ShowQueueInvoker(request *model.ShowQueueRequest) *ShowQueueInvoker {
	requestDef := GenReqDefForShowQueue()
	return &ShowQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuota 查询配额
//
// 该API用于获取用户配额信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowQuota(request *model.ShowQuotaRequest) (*model.ShowQuotaResponse, error) {
	requestDef := GenReqDefForShowQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotaResponse), nil
	}
}

// ShowQuotaInvoker 查询配额
func (c *DliClient) ShowQuotaInvoker(request *model.ShowQuotaRequest) *ShowQuotaInvoker {
	requestDef := GenReqDefForShowQuota()
	return &ShowQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// UpdateElasticResourcePoolQueue 修改弹性资源池关联的队列优先级
//
// 设置弹性资源池指定队列的扩缩容策略信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateElasticResourcePoolQueue(request *model.UpdateElasticResourcePoolQueueRequest) (*model.UpdateElasticResourcePoolQueueResponse, error) {
	requestDef := GenReqDefForUpdateElasticResourcePoolQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateElasticResourcePoolQueueResponse), nil
	}
}

// UpdateElasticResourcePoolQueueInvoker 修改弹性资源池关联的队列优先级
func (c *DliClient) UpdateElasticResourcePoolQueueInvoker(request *model.UpdateElasticResourcePoolQueueRequest) *UpdateElasticResourcePoolQueueInvoker {
	requestDef := GenReqDefForUpdateElasticResourcePoolQueue()
	return &UpdateElasticResourcePoolQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnhancedConnection 修改增强型跨源主机信息
//
// 该API用于在跨源中修改数据源主机信息，仅支持全量覆盖。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateEnhancedConnection(request *model.UpdateEnhancedConnectionRequest) (*model.UpdateEnhancedConnectionResponse, error) {
	requestDef := GenReqDefForUpdateEnhancedConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnhancedConnectionResponse), nil
	}
}

// UpdateEnhancedConnectionInvoker 修改增强型跨源主机信息
func (c *DliClient) UpdateEnhancedConnectionInvoker(request *model.UpdateEnhancedConnectionRequest) *UpdateEnhancedConnectionInvoker {
	requestDef := GenReqDefForUpdateEnhancedConnection()
	return &UpdateEnhancedConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalVariable 修改DLI全局变量
//
// 修改全局变量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateGlobalVariable(request *model.UpdateGlobalVariableRequest) (*model.UpdateGlobalVariableResponse, error) {
	requestDef := GenReqDefForUpdateGlobalVariable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalVariableResponse), nil
	}
}

// UpdateGlobalVariableInvoker 修改DLI全局变量
func (c *DliClient) UpdateGlobalVariableInvoker(request *model.UpdateGlobalVariableRequest) *UpdateGlobalVariableInvoker {
	requestDef := GenReqDefForUpdateGlobalVariable()
	return &UpdateGlobalVariableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateJobAuthInfo 更新跨源认证
//
// 该API用于更新跨源认证信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateJobAuthInfo(request *model.UpdateJobAuthInfoRequest) (*model.UpdateJobAuthInfoResponse, error) {
	requestDef := GenReqDefForUpdateJobAuthInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobAuthInfoResponse), nil
	}
}

// UpdateJobAuthInfoInvoker 更新跨源认证
func (c *DliClient) UpdateJobAuthInfoInvoker(request *model.UpdateJobAuthInfoRequest) *UpdateJobAuthInfoInvoker {
	requestDef := GenReqDefForUpdateJobAuthInfo()
	return &UpdateJobAuthInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateJobResourceOwner 修改组或者资源包拥有者
//
// 用于修改程序包的owner。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateJobResourceOwner(request *model.UpdateJobResourceOwnerRequest) (*model.UpdateJobResourceOwnerResponse, error) {
	requestDef := GenReqDefForUpdateJobResourceOwner()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateJobResourceOwnerResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateJobResourceOwnerInvoker 修改组或者资源包拥有者
func (c *DliClient) UpdateJobResourceOwnerInvoker(request *model.UpdateJobResourceOwnerRequest) *UpdateJobResourceOwnerInvoker {
	requestDef := GenReqDefForUpdateJobResourceOwner()
	return &UpdateJobResourceOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateQueueCidrInvoker 修改队列网段
func (c *DliClient) UpdateQueueCidrInvoker(request *model.UpdateQueueCidrRequest) *UpdateQueueCidrInvoker {
	requestDef := GenReqDefForUpdateQueueCidr()
	return &UpdateQueueCidrInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateQueuePlan 修改队列定时扩缩容计划
//
// 该API用于修改指定ID的队列定时扩缩容计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateQueuePlan(request *model.UpdateQueuePlanRequest) (*model.UpdateQueuePlanResponse, error) {
	requestDef := GenReqDefForUpdateQueuePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQueuePlanResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateQueuePlanInvoker 修改队列定时扩缩容计划
func (c *DliClient) UpdateQueuePlanInvoker(request *model.UpdateQueuePlanRequest) *UpdateQueuePlanInvoker {
	requestDef := GenReqDefForUpdateQueuePlan()
	return &UpdateQueuePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateQueueProperty 更新队列属性
//
// 更新队列属性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateQueueProperty(request *model.UpdateQueuePropertyRequest) (*model.UpdateQueuePropertyResponse, error) {
	requestDef := GenReqDefForUpdateQueueProperty()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateQueuePropertyResponse), nil
	}
}

// UpdateQueuePropertyInvoker 更新队列属性
func (c *DliClient) UpdateQueuePropertyInvoker(request *model.UpdateQueuePropertyRequest) *UpdateQueuePropertyInvoker {
	requestDef := GenReqDefForUpdateQueueProperty()
	return &UpdateQueuePropertyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadFileJobResources 上传file类型分组资源
//
// 该API用于在project下上传file类型模块。
// 说明： 上传同名file模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadFileJobResources(request *model.UploadFileJobResourcesRequest) (*model.UploadFileJobResourcesResponse, error) {
	requestDef := GenReqDefForUploadFileJobResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadFileJobResourcesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadFileJobResourcesInvoker 上传file类型分组资源
func (c *DliClient) UploadFileJobResourcesInvoker(request *model.UploadFileJobResourcesRequest) *UploadFileJobResourcesInvoker {
	requestDef := GenReqDefForUploadFileJobResources()
	return &UploadFileJobResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadJarJobResources 上传jar类型分组资源
//
// 该API用于在project下上传jar类型分组资源。
// 说明：上传同名资源模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadJarJobResources(request *model.UploadJarJobResourcesRequest) (*model.UploadJarJobResourcesResponse, error) {
	requestDef := GenReqDefForUploadJarJobResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJarJobResourcesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadJarJobResourcesInvoker 上传jar类型分组资源
func (c *DliClient) UploadJarJobResourcesInvoker(request *model.UploadJarJobResourcesRequest) *UploadJarJobResourcesInvoker {
	requestDef := GenReqDefForUploadJarJobResources()
	return &UploadJarJobResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadJobResources 上传分组资源
//
// 该API用于上传分组资源到某个project下。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadJobResources(request *model.UploadJobResourcesRequest) (*model.UploadJobResourcesResponse, error) {
	requestDef := GenReqDefForUploadJobResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadJobResourcesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadJobResourcesInvoker 上传分组资源
func (c *DliClient) UploadJobResourcesInvoker(request *model.UploadJobResourcesRequest) *UploadJobResourcesInvoker {
	requestDef := GenReqDefForUploadJobResources()
	return &UploadJobResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadPythonFileJobResources 上传pyfile类型分组资源
//
// 该API用于在project下的上传pyfile类型模块。
// 说明： 上传同名pyfile类型模块时，新模块将会覆盖旧模块。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UploadPythonFileJobResources(request *model.UploadPythonFileJobResourcesRequest) (*model.UploadPythonFileJobResourcesResponse, error) {
	requestDef := GenReqDefForUploadPythonFileJobResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadPythonFileJobResourcesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UploadPythonFileJobResourcesInvoker 上传pyfile类型分组资源
func (c *DliClient) UploadPythonFileJobResourcesInvoker(request *model.UploadPythonFileJobResourcesRequest) *UploadPythonFileJobResourcesInvoker {
	requestDef := GenReqDefForUploadPythonFileJobResources()
	return &UploadPythonFileJobResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchStopFlinkJobs 批量停止Flink作业
//
// 批量停止正在运行的Flink作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) BatchStopFlinkJobs(request *model.BatchStopFlinkJobsRequest) (*model.BatchStopFlinkJobsResponse, error) {
	requestDef := GenReqDefForBatchStopFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchStopFlinkJobsResponse), nil
	}
}

// BatchStopFlinkJobsInvoker 批量停止Flink作业
func (c *DliClient) BatchStopFlinkJobsInvoker(request *model.BatchStopFlinkJobsRequest) *BatchStopFlinkJobsInvoker {
	requestDef := GenReqDefForBatchStopFlinkJobs()
	return &BatchStopFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkJarJob 新建Flink Jar作业
//
// 用户自定义作业目前支持jar格式，运行在独享集群中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkJarJob(request *model.CreateFlinkJarJobRequest) (*model.CreateFlinkJarJobResponse, error) {
	requestDef := GenReqDefForCreateFlinkJarJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkJarJobResponse), nil
	}
}

// CreateFlinkJarJobInvoker 新建Flink Jar作业
func (c *DliClient) CreateFlinkJarJobInvoker(request *model.CreateFlinkJarJobRequest) *CreateFlinkJarJobInvoker {
	requestDef := GenReqDefForCreateFlinkJarJob()
	return &CreateFlinkJarJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateFlinkSqlJobGraph 生成flink SQL作业的静态流图
//
// 生成flink SQL作业的静态流图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkSqlJobGraph(request *model.CreateFlinkSqlJobGraphRequest) (*model.CreateFlinkSqlJobGraphResponse, error) {
	requestDef := GenReqDefForCreateFlinkSqlJobGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSqlJobGraphResponse), nil
	}
}

// CreateFlinkSqlJobGraphInvoker 生成flink SQL作业的静态流图
func (c *DliClient) CreateFlinkSqlJobGraphInvoker(request *model.CreateFlinkSqlJobGraphRequest) *CreateFlinkSqlJobGraphInvoker {
	requestDef := GenReqDefForCreateFlinkSqlJobGraph()
	return &CreateFlinkSqlJobGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkSqlJobTemplate 新建Flink作业模板
//
// 在DLI服务中新建一个Flink作业模板，最多100个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateFlinkSqlJobTemplate(request *model.CreateFlinkSqlJobTemplateRequest) (*model.CreateFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForCreateFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSqlJobTemplateResponse), nil
	}
}

// CreateFlinkSqlJobTemplateInvoker 新建Flink作业模板
func (c *DliClient) CreateFlinkSqlJobTemplateInvoker(request *model.CreateFlinkSqlJobTemplateRequest) *CreateFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForCreateFlinkSqlJobTemplate()
	return &CreateFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateIefMessageChannelInvoker 创建IEF消息通道
func (c *DliClient) CreateIefMessageChannelInvoker(request *model.CreateIefMessageChannelRequest) *CreateIefMessageChannelInvoker {
	requestDef := GenReqDefForCreateIefMessageChannel()
	return &CreateIefMessageChannelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateIefSystemEventsInvoker IEF系统事件上报
func (c *DliClient) CreateIefSystemEventsInvoker(request *model.CreateIefSystemEventsRequest) *CreateIefSystemEventsInvoker {
	requestDef := GenReqDefForCreateIefSystemEvents()
	return &CreateIefSystemEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// DeleteFlinkSqlJobTemplate 删除Flink作业模板
//
// 删除一个Flink作业模板，即使当前模板正在被作业使用，也允许删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) DeleteFlinkSqlJobTemplate(request *model.DeleteFlinkSqlJobTemplateRequest) (*model.DeleteFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForDeleteFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkSqlJobTemplateResponse), nil
	}
}

// DeleteFlinkSqlJobTemplateInvoker 删除Flink作业模板
func (c *DliClient) DeleteFlinkSqlJobTemplateInvoker(request *model.DeleteFlinkSqlJobTemplateRequest) *DeleteFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForDeleteFlinkSqlJobTemplate()
	return &DeleteFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteFlinkJobSavepoint 触发Flink作业保存点
//
// 触发Flink作业保存点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ExecuteFlinkJobSavepoint(request *model.ExecuteFlinkJobSavepointRequest) (*model.ExecuteFlinkJobSavepointResponse, error) {
	requestDef := GenReqDefForExecuteFlinkJobSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteFlinkJobSavepointResponse), nil
	}
}

// ExecuteFlinkJobSavepointInvoker 触发Flink作业保存点
func (c *DliClient) ExecuteFlinkJobSavepointInvoker(request *model.ExecuteFlinkJobSavepointRequest) *ExecuteFlinkJobSavepointInvoker {
	requestDef := GenReqDefForExecuteFlinkJobSavepoint()
	return &ExecuteFlinkJobSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportFlinkJobs flink作业导出
//
// 通过POST方式，导出flink作业，请求体为JSON格式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ExportFlinkJobs(request *model.ExportFlinkJobsRequest) (*model.ExportFlinkJobsResponse, error) {
	requestDef := GenReqDefForExportFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportFlinkJobsResponse), nil
	}
}

// ExportFlinkJobsInvoker flink作业导出
func (c *DliClient) ExportFlinkJobsInvoker(request *model.ExportFlinkJobsRequest) *ExportFlinkJobsInvoker {
	requestDef := GenReqDefForExportFlinkJobs()
	return &ExportFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportFlinkJobSavepoint 导入Flink作业保存点
//
// 导入Flink作业保存点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ImportFlinkJobSavepoint(request *model.ImportFlinkJobSavepointRequest) (*model.ImportFlinkJobSavepointResponse, error) {
	requestDef := GenReqDefForImportFlinkJobSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFlinkJobSavepointResponse), nil
	}
}

// ImportFlinkJobSavepointInvoker 导入Flink作业保存点
func (c *DliClient) ImportFlinkJobSavepointInvoker(request *model.ImportFlinkJobSavepointRequest) *ImportFlinkJobSavepointInvoker {
	requestDef := GenReqDefForImportFlinkJobSavepoint()
	return &ImportFlinkJobSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportFlinkJobs flink作业导入
//
// 通过POST方式，导入flink作业，请求体为JSON格式。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ImportFlinkJobs(request *model.ImportFlinkJobsRequest) (*model.ImportFlinkJobsResponse, error) {
	requestDef := GenReqDefForImportFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportFlinkJobsResponse), nil
	}
}

// ImportFlinkJobsInvoker flink作业导入
func (c *DliClient) ImportFlinkJobsInvoker(request *model.ImportFlinkJobsRequest) *ImportFlinkJobsInvoker {
	requestDef := GenReqDefForImportFlinkJobs()
	return &ImportFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListFlinkSqlJobTemplates 查询Flink作业模板列表
//
// 查询Flink作业模板列表。当前只支持查询用户自定义模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListFlinkSqlJobTemplates(request *model.ListFlinkSqlJobTemplatesRequest) (*model.ListFlinkSqlJobTemplatesResponse, error) {
	requestDef := GenReqDefForListFlinkSqlJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkSqlJobTemplatesResponse), nil
	}
}

// ListFlinkSqlJobTemplatesInvoker 查询Flink作业模板列表
func (c *DliClient) ListFlinkSqlJobTemplatesInvoker(request *model.ListFlinkSqlJobTemplatesRequest) *ListFlinkSqlJobTemplatesInvoker {
	requestDef := GenReqDefForListFlinkSqlJobTemplates()
	return &ListFlinkSqlJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// RegisterBucketInvoker OBS授权给DLI服务
func (c *DliClient) RegisterBucketInvoker(request *model.RegisterBucketRequest) *RegisterBucketInvoker {
	requestDef := GenReqDefForRegisterBucket()
	return &RegisterBucketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// RunIefJobActionCallBackInvoker 边缘Flink作业Action状态回调
func (c *DliClient) RunIefJobActionCallBackInvoker(request *model.RunIefJobActionCallBackRequest) *RunIefJobActionCallBackInvoker {
	requestDef := GenReqDefForRunIefJobActionCallBack()
	return &RunIefJobActionCallBackInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowFlinkJobExecutionGraph 查询Flink作业执行计划
//
// 查询Flink作业执行计划。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowFlinkJobExecutionGraph(request *model.ShowFlinkJobExecutionGraphRequest) (*model.ShowFlinkJobExecutionGraphResponse, error) {
	requestDef := GenReqDefForShowFlinkJobExecutionGraph()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobExecutionGraphResponse), nil
	}
}

// ShowFlinkJobExecutionGraphInvoker 查询Flink作业执行计划
func (c *DliClient) ShowFlinkJobExecutionGraphInvoker(request *model.ShowFlinkJobExecutionGraphRequest) *ShowFlinkJobExecutionGraphInvoker {
	requestDef := GenReqDefForShowFlinkJobExecutionGraph()
	return &ShowFlinkJobExecutionGraphInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowFlinkMetricInvoker 查询Flink作业监控信息
func (c *DliClient) ShowFlinkMetricInvoker(request *model.ShowFlinkMetricRequest) *ShowFlinkMetricInvoker {
	requestDef := GenReqDefForShowFlinkMetric()
	return &ShowFlinkMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkJarJob 更新Flink Jar作业
//
// 更新用户自定义作业，目前支持jar格式，运行在独享集群中。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkJarJob(request *model.UpdateFlinkJarJobRequest) (*model.UpdateFlinkJarJobResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJarJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJarJobResponse), nil
	}
}

// UpdateFlinkJarJobInvoker 更新Flink Jar作业
func (c *DliClient) UpdateFlinkJarJobInvoker(request *model.UpdateFlinkJarJobRequest) *UpdateFlinkJarJobInvoker {
	requestDef := GenReqDefForUpdateFlinkJarJob()
	return &UpdateFlinkJarJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateFlinkJobStatusReport 边缘Flink作业状态信息上报
//
// 该API用于处理边缘Flink作业状态上报信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkJobStatusReport(request *model.UpdateFlinkJobStatusReportRequest) (*model.UpdateFlinkJobStatusReportResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJobStatusReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJobStatusReportResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateFlinkJobStatusReportInvoker 边缘Flink作业状态信息上报
func (c *DliClient) UpdateFlinkJobStatusReportInvoker(request *model.UpdateFlinkJobStatusReportRequest) *UpdateFlinkJobStatusReportInvoker {
	requestDef := GenReqDefForUpdateFlinkJobStatusReport()
	return &UpdateFlinkJobStatusReportInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkSqlJob 更新Flink SQL作业
//
// Stream SQL的语法扩展了Apache Flink SQL。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkSqlJob(request *model.UpdateFlinkSqlJobRequest) (*model.UpdateFlinkSqlJobResponse, error) {
	requestDef := GenReqDefForUpdateFlinkSqlJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkSqlJobResponse), nil
	}
}

// UpdateFlinkSqlJobInvoker 更新Flink SQL作业
func (c *DliClient) UpdateFlinkSqlJobInvoker(request *model.UpdateFlinkSqlJobRequest) *UpdateFlinkSqlJobInvoker {
	requestDef := GenReqDefForUpdateFlinkSqlJob()
	return &UpdateFlinkSqlJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkSqlJobTemplate 更新Flink作业模板
//
// 对DLI服务中已有的Flink作业模板进行更新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateFlinkSqlJobTemplate(request *model.UpdateFlinkSqlJobTemplateRequest) (*model.UpdateFlinkSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForUpdateFlinkSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkSqlJobTemplateResponse), nil
	}
}

// UpdateFlinkSqlJobTemplateInvoker 更新Flink作业模板
func (c *DliClient) UpdateFlinkSqlJobTemplateInvoker(request *model.UpdateFlinkSqlJobTemplateRequest) *UpdateFlinkSqlJobTemplateInvoker {
	requestDef := GenReqDefForUpdateFlinkSqlJobTemplate()
	return &UpdateFlinkSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSparkJob 取消批处理作业
//
// 该API用于取消批处理作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CancelSparkJob(request *model.CancelSparkJobRequest) (*model.CancelSparkJobResponse, error) {
	requestDef := GenReqDefForCancelSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkJobResponse), nil
	}
}

// CancelSparkJobInvoker 取消批处理作业
func (c *DliClient) CancelSparkJobInvoker(request *model.CancelSparkJobRequest) *CancelSparkJobInvoker {
	requestDef := GenReqDefForCancelSparkJob()
	return &CancelSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSparkJob 创建批处理作业
//
// 该API用于在某个队列上创建作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateSparkJob(request *model.CreateSparkJobRequest) (*model.CreateSparkJobResponse, error) {
	requestDef := GenReqDefForCreateSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSparkJobResponse), nil
	}
}

// CreateSparkJobInvoker 创建批处理作业
func (c *DliClient) CreateSparkJobInvoker(request *model.CreateSparkJobRequest) *CreateSparkJobInvoker {
	requestDef := GenReqDefForCreateSparkJob()
	return &CreateSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSparkJobTemplate 创建作业模板
//
// 该API用于创建作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateSparkJobTemplate(request *model.CreateSparkJobTemplateRequest) (*model.CreateSparkJobTemplateResponse, error) {
	requestDef := GenReqDefForCreateSparkJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSparkJobTemplateResponse), nil
	}
}

// CreateSparkJobTemplateInvoker 创建作业模板
func (c *DliClient) CreateSparkJobTemplateInvoker(request *model.CreateSparkJobTemplateRequest) *CreateSparkJobTemplateInvoker {
	requestDef := GenReqDefForCreateSparkJobTemplate()
	return &CreateSparkJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkJobTemplates 查询作业模板列表
//
// 该API用于查询作业模板列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListSparkJobTemplates(request *model.ListSparkJobTemplatesRequest) (*model.ListSparkJobTemplatesResponse, error) {
	requestDef := GenReqDefForListSparkJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkJobTemplatesResponse), nil
	}
}

// ListSparkJobTemplatesInvoker 查询作业模板列表
func (c *DliClient) ListSparkJobTemplatesInvoker(request *model.ListSparkJobTemplatesRequest) *ListSparkJobTemplatesInvoker {
	requestDef := GenReqDefForListSparkJobTemplates()
	return &ListSparkJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkJobs 查询批处理作业列表
//
// 该API用于查询Project下某队列批处理作业的列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListSparkJobs(request *model.ListSparkJobsRequest) (*model.ListSparkJobsResponse, error) {
	requestDef := GenReqDefForListSparkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkJobsResponse), nil
	}
}

// ListSparkJobsInvoker 查询批处理作业列表
func (c *DliClient) ListSparkJobsInvoker(request *model.ListSparkJobsRequest) *ListSparkJobsInvoker {
	requestDef := GenReqDefForListSparkJobs()
	return &ListSparkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJob 查询批处理作业详情
//
// 该API用于根据批处理作业的id查询作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSparkJob(request *model.ShowSparkJobRequest) (*model.ShowSparkJobResponse, error) {
	requestDef := GenReqDefForShowSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobResponse), nil
	}
}

// ShowSparkJobInvoker 查询批处理作业详情
func (c *DliClient) ShowSparkJobInvoker(request *model.ShowSparkJobRequest) *ShowSparkJobInvoker {
	requestDef := GenReqDefForShowSparkJob()
	return &ShowSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowSparkJobLog 查询批处理作业日志
//
// 该API用于查询批处理作业的后台日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSparkJobLog(request *model.ShowSparkJobLogRequest) (*model.ShowSparkJobLogResponse, error) {
	requestDef := GenReqDefForShowSparkJobLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobLogResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowSparkJobLogInvoker 查询批处理作业日志
func (c *DliClient) ShowSparkJobLogInvoker(request *model.ShowSparkJobLogRequest) *ShowSparkJobLogInvoker {
	requestDef := GenReqDefForShowSparkJobLog()
	return &ShowSparkJobLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJobStatus 查询批处理作业状态
//
// 该API用于查询批处理作业的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSparkJobStatus(request *model.ShowSparkJobStatusRequest) (*model.ShowSparkJobStatusResponse, error) {
	requestDef := GenReqDefForShowSparkJobStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobStatusResponse), nil
	}
}

// ShowSparkJobStatusInvoker 查询批处理作业状态
func (c *DliClient) ShowSparkJobStatusInvoker(request *model.ShowSparkJobStatusRequest) *ShowSparkJobStatusInvoker {
	requestDef := GenReqDefForShowSparkJobStatus()
	return &ShowSparkJobStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJobTemplate 获取作业模板
//
// 该API用于获取作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSparkJobTemplate(request *model.ShowSparkJobTemplateRequest) (*model.ShowSparkJobTemplateResponse, error) {
	requestDef := GenReqDefForShowSparkJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobTemplateResponse), nil
	}
}

// ShowSparkJobTemplateInvoker 获取作业模板
func (c *DliClient) ShowSparkJobTemplateInvoker(request *model.ShowSparkJobTemplateRequest) *ShowSparkJobTemplateInvoker {
	requestDef := GenReqDefForShowSparkJobTemplate()
	return &ShowSparkJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSparkJobTemplate 修改作业模板
//
// 该API用于修改作业模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateSparkJobTemplate(request *model.UpdateSparkJobTemplateRequest) (*model.UpdateSparkJobTemplateResponse, error) {
	requestDef := GenReqDefForUpdateSparkJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSparkJobTemplateResponse), nil
	}
}

// UpdateSparkJobTemplateInvoker 修改作业模板
func (c *DliClient) UpdateSparkJobTemplateInvoker(request *model.UpdateSparkJobTemplateRequest) *UpdateSparkJobTemplateInvoker {
	requestDef := GenReqDefForUpdateSparkJobTemplate()
	return &UpdateSparkJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSqlJobTemplates 批量删除SQL模板
//
// 该API用于批量删除SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) BatchDeleteSqlJobTemplates(request *model.BatchDeleteSqlJobTemplatesRequest) (*model.BatchDeleteSqlJobTemplatesResponse, error) {
	requestDef := GenReqDefForBatchDeleteSqlJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSqlJobTemplatesResponse), nil
	}
}

// BatchDeleteSqlJobTemplatesInvoker 批量删除SQL模板
func (c *DliClient) BatchDeleteSqlJobTemplatesInvoker(request *model.BatchDeleteSqlJobTemplatesRequest) *BatchDeleteSqlJobTemplatesInvoker {
	requestDef := GenReqDefForBatchDeleteSqlJobTemplates()
	return &BatchDeleteSqlJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateSqlJobTemplate 存储指定SQL语句
//
// 该API用于存储指定的SQL语句，后续可以重复使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) CreateSqlJobTemplate(request *model.CreateSqlJobTemplateRequest) (*model.CreateSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForCreateSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSqlJobTemplateResponse), nil
	}
}

// CreateSqlJobTemplateInvoker 存储指定SQL语句
func (c *DliClient) CreateSqlJobTemplateInvoker(request *model.CreateSqlJobTemplateRequest) *CreateSqlJobTemplateInvoker {
	requestDef := GenReqDefForCreateSqlJobTemplate()
	return &CreateSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListSqlJobTemplates 查看所有SQL模板
//
// 该API用查看用户保存的所有SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListSqlJobTemplates(request *model.ListSqlJobTemplatesRequest) (*model.ListSqlJobTemplatesResponse, error) {
	requestDef := GenReqDefForListSqlJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlJobTemplatesResponse), nil
	}
}

// ListSqlJobTemplatesInvoker 查看所有SQL模板
func (c *DliClient) ListSqlJobTemplatesInvoker(request *model.ListSqlJobTemplatesRequest) *ListSqlJobTemplatesInvoker {
	requestDef := GenReqDefForListSqlJobTemplates()
	return &ListSqlJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// PreviewSqlJobResult 预览SQL作业查询结果
//
// 该API用于在执行SQL查询语句的作业完成后，查看该作业执行的结果。目前仅支持查看“QUERY”类型作业的执行结果。
// 该API只能查看前1000条的结果记录，若要查看全部的结果记录，需要先导出查询结果再进行查看。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) PreviewSqlJobResult(request *model.PreviewSqlJobResultRequest) (*model.PreviewSqlJobResultResponse, error) {
	requestDef := GenReqDefForPreviewSqlJobResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewSqlJobResultResponse), nil
	}
}

// PreviewSqlJobResultInvoker 预览SQL作业查询结果
func (c *DliClient) PreviewSqlJobResultInvoker(request *model.PreviewSqlJobResultRequest) *PreviewSqlJobResultInvoker {
	requestDef := GenReqDefForPreviewSqlJobResult()
	return &PreviewSqlJobResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ShowSqlJobProgress 查询作业执行进度信息
//
// 该API用于获取作业执行进度信息，如果作业正在执行，可以获取到子作业的信息，如果作业刚开始或者已经结束，不可以获取到子作业信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowSqlJobProgress(request *model.ShowSqlJobProgressRequest) (*model.ShowSqlJobProgressResponse, error) {
	requestDef := GenReqDefForShowSqlJobProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSqlJobProgressResponse), nil
	}
}

// ShowSqlJobProgressInvoker 查询作业执行进度信息
func (c *DliClient) ShowSqlJobProgressInvoker(request *model.ShowSqlJobProgressRequest) *ShowSqlJobProgressInvoker {
	requestDef := GenReqDefForShowSqlJobProgress()
	return &ShowSqlJobProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// UpdateSqlJobTemplate 更新SQL模板
//
// 该API用于更新SQL模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) UpdateSqlJobTemplate(request *model.UpdateSqlJobTemplateRequest) (*model.UpdateSqlJobTemplateResponse, error) {
	requestDef := GenReqDefForUpdateSqlJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSqlJobTemplateResponse), nil
	}
}

// UpdateSqlJobTemplateInvoker 更新SQL模板
func (c *DliClient) UpdateSqlJobTemplateInvoker(request *model.UpdateSqlJobTemplateRequest) *UpdateSqlJobTemplateInvoker {
	requestDef := GenReqDefForUpdateSqlJobTemplate()
	return &UpdateSqlJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateDatabaseInvoker 创建数据库
func (c *DliClient) CreateDatabaseInvoker(request *model.CreateDatabaseRequest) *CreateDatabaseInvoker {
	requestDef := GenReqDefForCreateDatabase()
	return &CreateDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateTableInvoker 创建表
func (c *DliClient) CreateTableInvoker(request *model.CreateTableRequest) *CreateTableInvoker {
	requestDef := GenReqDefForCreateTable()
	return &CreateTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteDatabaseInvoker 删除数据库
func (c *DliClient) DeleteDatabaseInvoker(request *model.DeleteDatabaseRequest) *DeleteDatabaseInvoker {
	requestDef := GenReqDefForDeleteDatabase()
	return &DeleteDatabaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteTableInvoker 删除表
func (c *DliClient) DeleteTableInvoker(request *model.DeleteTableRequest) *DeleteTableInvoker {
	requestDef := GenReqDefForDeleteTable()
	return &DeleteTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ExportTable 导出表数据
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ExportTableInvoker 导出表数据
func (c *DliClient) ExportTableInvoker(request *model.ExportTableRequest) *ExportTableInvoker {
	requestDef := GenReqDefForExportTable()
	return &ExportTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ImportTableInvoker 导入数据
func (c *DliClient) ImportTableInvoker(request *model.ImportTableRequest) *ImportTableInvoker {
	requestDef := GenReqDefForImportTable()
	return &ImportTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListDatabasesInvoker 查询所有数据库
func (c *DliClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListPartitions 获取分区信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListPartitions(request *model.ListPartitionsRequest) (*model.ListPartitionsResponse, error) {
	requestDef := GenReqDefForListPartitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPartitionsResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListPartitionsInvoker 获取分区信息列表
func (c *DliClient) ListPartitionsInvoker(request *model.ListPartitionsRequest) *ListPartitionsInvoker {
	requestDef := GenReqDefForListPartitions()
	return &ListPartitionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListSqlSampleTemplates 查询所有SQL样例模板
//
// 该API用于查询所有SQL样例模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListSqlSampleTemplates(request *model.ListSqlSampleTemplatesRequest) (*model.ListSqlSampleTemplatesResponse, error) {
	requestDef := GenReqDefForListSqlSampleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSqlSampleTemplatesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListSqlSampleTemplatesInvoker 查询所有SQL样例模板
func (c *DliClient) ListSqlSampleTemplatesInvoker(request *model.ListSqlSampleTemplatesRequest) *ListSqlSampleTemplatesInvoker {
	requestDef := GenReqDefForListSqlSampleTemplates()
	return &ListSqlSampleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTables 查询所有表
//
// 该API用于查询指定数据库下符合过滤条件的或所有的表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ListTables(request *model.ListTablesRequest) (*model.ListTablesResponse, error) {
	requestDef := GenReqDefForListTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTablesResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ListTablesInvoker 查询所有表
func (c *DliClient) ListTablesInvoker(request *model.ListTablesRequest) *ListTablesInvoker {
	requestDef := GenReqDefForListTables()
	return &ListTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PreviewTable 预览表内容
//
// 该API用于用于预览表中前10行的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) PreviewTable(request *model.PreviewTableRequest) (*model.PreviewTableResponse, error) {
	requestDef := GenReqDefForPreviewTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewTableResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PreviewTableInvoker 预览表内容
func (c *DliClient) PreviewTableInvoker(request *model.PreviewTableRequest) *PreviewTableInvoker {
	requestDef := GenReqDefForPreviewTable()
	return &PreviewTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowTable 描述表信息
//
// 该API用于描述指定表的元数据信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DliClient) ShowTable(request *model.ShowTableRequest) (*model.ShowTableResponse, error) {
	requestDef := GenReqDefForShowTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ShowTableInvoker 描述表信息
func (c *DliClient) ShowTableInvoker(request *model.ShowTableRequest) *ShowTableInvoker {
	requestDef := GenReqDefForShowTable()
	return &ShowTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateDatabaseOwnerInvoker 修改数据库用户
func (c *DliClient) UpdateDatabaseOwnerInvoker(request *model.UpdateDatabaseOwnerRequest) *UpdateDatabaseOwnerInvoker {
	requestDef := GenReqDefForUpdateDatabaseOwner()
	return &UpdateDatabaseOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
// UpdateTableOwnerInvoker 修改表用户
func (c *DliClient) UpdateTableOwnerInvoker(request *model.UpdateTableOwnerRequest) *UpdateTableOwnerInvoker {
	requestDef := GenReqDefForUpdateTableOwner()
	return &UpdateTableOwnerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
