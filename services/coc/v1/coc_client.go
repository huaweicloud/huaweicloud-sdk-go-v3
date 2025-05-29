package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/coc/v1/model"
)

type CocClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCocClient(hcClient *httpclient.HcHttpClient) *CocClient {
	return &CocClient{HcClient: hcClient}
}

func CocClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials,basic.Credentials")
	return builder
}

// ListApplications 查询应用
//
// 查询应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplications(request *model.ListApplicationsRequest) (*model.ListApplicationsResponse, error) {
	requestDef := GenReqDefForListApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationsResponse), nil
	}
}

// ListApplicationsInvoker 查询应用
func (c *CocClient) ListApplicationsInvoker(request *model.ListApplicationsRequest) *ListApplicationsInvoker {
	requestDef := GenReqDefForListApplications()
	return &ListApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationModel 查询下一级的子应用、组件、分组
//
// 查询下一级的子应用、组件、分组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplicationModel(request *model.ListApplicationModelRequest) (*model.ListApplicationModelResponse, error) {
	requestDef := GenReqDefForListApplicationModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationModelResponse), nil
	}
}

// ListApplicationModelInvoker 查询下一级的子应用、组件、分组
func (c *CocClient) ListApplicationModelInvoker(request *model.ListApplicationModelRequest) *ListApplicationModelInvoker {
	requestDef := GenReqDefForListApplicationModel()
	return &ListApplicationModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateApplicationView 批量创建应用视图
//
// 批量创建应用视图
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) BatchCreateApplicationView(request *model.BatchCreateApplicationViewRequest) (*model.BatchCreateApplicationViewResponse, error) {
	requestDef := GenReqDefForBatchCreateApplicationView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateApplicationViewResponse), nil
	}
}

// BatchCreateApplicationViewInvoker 批量创建应用视图
func (c *CocClient) BatchCreateApplicationViewInvoker(request *model.BatchCreateApplicationViewRequest) *BatchCreateApplicationViewInvoker {
	requestDef := GenReqDefForBatchCreateApplicationView()
	return &BatchCreateApplicationViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceCompliant 获取节点合规性报告
//
// 分页获取节点合规性报告
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListInstanceCompliant(request *model.ListInstanceCompliantRequest) (*model.ListInstanceCompliantResponse, error) {
	requestDef := GenReqDefForListInstanceCompliant()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceCompliantResponse), nil
	}
}

// ListInstanceCompliantInvoker 获取节点合规性报告
func (c *CocClient) ListInstanceCompliantInvoker(request *model.ListInstanceCompliantRequest) *ListInstanceCompliantInvoker {
	requestDef := GenReqDefForListInstanceCompliant()
	return &ListInstanceCompliantInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstancePatchItems 分页获取节点补丁详情
//
// 分页获取节点补丁详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowInstancePatchItems(request *model.ShowInstancePatchItemsRequest) (*model.ShowInstancePatchItemsResponse, error) {
	requestDef := GenReqDefForShowInstancePatchItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstancePatchItemsResponse), nil
	}
}

// ShowInstancePatchItemsInvoker 分页获取节点补丁详情
func (c *CocClient) ShowInstancePatchItemsInvoker(request *model.ShowInstancePatchItemsRequest) *ShowInstancePatchItemsInvoker {
	requestDef := GenReqDefForShowInstancePatchItems()
	return &ShowInstancePatchItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportCustomEvent 支持用户自主接入告警数据
//
// 支持租户将自开发的监控系统按照标准化集成至COC，集成后告警会按照标准格式上报至COC告警中心
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateReportCustomEvent(request *model.CreateReportCustomEventRequest) (*model.CreateReportCustomEventResponse, error) {
	requestDef := GenReqDefForCreateReportCustomEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportCustomEventResponse), nil
	}
}

// CreateReportCustomEventInvoker 支持用户自主接入告警数据
func (c *CocClient) CreateReportCustomEventInvoker(request *model.CreateReportCustomEventRequest) *CreateReportCustomEventInvoker {
	requestDef := GenReqDefForCreateReportCustomEvent()
	return &CreateReportCustomEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDocument 创建自定义作业
//
// 创建自定义作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateDocument(request *model.CreateDocumentRequest) (*model.CreateDocumentResponse, error) {
	requestDef := GenReqDefForCreateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocumentResponse), nil
	}
}

// CreateDocumentInvoker 创建自定义作业
func (c *CocClient) CreateDocumentInvoker(request *model.CreateDocumentRequest) *CreateDocumentInvoker {
	requestDef := GenReqDefForCreateDocument()
	return &CreateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDocument 删除自定义作业
//
// 删除自定义作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteDocument(request *model.DeleteDocumentRequest) (*model.DeleteDocumentResponse, error) {
	requestDef := GenReqDefForDeleteDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDocumentResponse), nil
	}
}

// DeleteDocumentInvoker 删除自定义作业
func (c *CocClient) DeleteDocumentInvoker(request *model.DeleteDocumentRequest) *DeleteDocumentInvoker {
	requestDef := GenReqDefForDeleteDocument()
	return &DeleteDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDocument 执行自定义作业
//
// 执行自定义作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecuteDocument(request *model.ExecuteDocumentRequest) (*model.ExecuteDocumentResponse, error) {
	requestDef := GenReqDefForExecuteDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDocumentResponse), nil
	}
}

// ExecuteDocumentInvoker 执行自定义作业
func (c *CocClient) ExecuteDocumentInvoker(request *model.ExecuteDocumentRequest) *ExecuteDocumentInvoker {
	requestDef := GenReqDefForExecuteDocument()
	return &ExecuteDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDocument 查询自定义作业详情
//
// 查询自定义作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetDocument(request *model.GetDocumentRequest) (*model.GetDocumentResponse, error) {
	requestDef := GenReqDefForGetDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDocumentResponse), nil
	}
}

// GetDocumentInvoker 查询自定义作业详情
func (c *CocClient) GetDocumentInvoker(request *model.GetDocumentRequest) *GetDocumentInvoker {
	requestDef := GenReqDefForGetDocument()
	return &GetDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDocumentAtomicInfo 获取原子能力详细
//
// 获取原子能力详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetDocumentAtomicInfo(request *model.GetDocumentAtomicInfoRequest) (*model.GetDocumentAtomicInfoResponse, error) {
	requestDef := GenReqDefForGetDocumentAtomicInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDocumentAtomicInfoResponse), nil
	}
}

// GetDocumentAtomicInfoInvoker 获取原子能力详细
func (c *CocClient) GetDocumentAtomicInfoInvoker(request *model.GetDocumentAtomicInfoRequest) *GetDocumentAtomicInfoInvoker {
	requestDef := GenReqDefForGetDocumentAtomicInfo()
	return &GetDocumentAtomicInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDocumentAtomics 获取原子能力列表
//
// 获取原子能力列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListDocumentAtomics(request *model.ListDocumentAtomicsRequest) (*model.ListDocumentAtomicsResponse, error) {
	requestDef := GenReqDefForListDocumentAtomics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDocumentAtomicsResponse), nil
	}
}

// ListDocumentAtomicsInvoker 获取原子能力列表
func (c *CocClient) ListDocumentAtomicsInvoker(request *model.ListDocumentAtomicsRequest) *ListDocumentAtomicsInvoker {
	requestDef := GenReqDefForListDocumentAtomics()
	return &ListDocumentAtomicsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDocuments 查询自定义作业列表
//
// 查询自定义作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListDocuments(request *model.ListDocumentsRequest) (*model.ListDocumentsResponse, error) {
	requestDef := GenReqDefForListDocuments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDocumentsResponse), nil
	}
}

// ListDocumentsInvoker 查询自定义作业列表
func (c *CocClient) ListDocumentsInvoker(request *model.ListDocumentsRequest) *ListDocumentsInvoker {
	requestDef := GenReqDefForListDocuments()
	return &ListDocumentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDocument 修改自定义作业
//
// 修改自定义作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateDocument(request *model.UpdateDocumentRequest) (*model.UpdateDocumentResponse, error) {
	requestDef := GenReqDefForUpdateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDocumentResponse), nil
	}
}

// UpdateDocumentInvoker 修改自定义作业
func (c *CocClient) UpdateDocumentInvoker(request *model.UpdateDocumentRequest) *UpdateDocumentInvoker {
	requestDef := GenReqDefForUpdateDocument()
	return &UpdateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateReportPrometheusEvent Prometheus事件接入
//
// Prometheus事件接入
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateReportPrometheusEvent(request *model.CreateReportPrometheusEventRequest) (*model.CreateReportPrometheusEventResponse, error) {
	requestDef := GenReqDefForCreateReportPrometheusEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateReportPrometheusEventResponse), nil
	}
}

// CreateReportPrometheusEventInvoker Prometheus事件接入
func (c *CocClient) CreateReportPrometheusEventInvoker(request *model.CreateReportPrometheusEventRequest) *CreateReportPrometheusEventInvoker {
	requestDef := GenReqDefForCreateReportPrometheusEvent()
	return &CreateReportPrometheusEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetExecution 查询作业工单详情
//
// 查询作业工单详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetExecution(request *model.GetExecutionRequest) (*model.GetExecutionResponse, error) {
	requestDef := GenReqDefForGetExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetExecutionResponse), nil
	}
}

// GetExecutionInvoker 查询作业工单详情
func (c *CocClient) GetExecutionInvoker(request *model.GetExecutionRequest) *GetExecutionInvoker {
	requestDef := GenReqDefForGetExecution()
	return &GetExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExecutionInstances 查询工单步骤批次实例
//
// 查询工单步骤批次实例，如脚本分批操作里的ECS实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListExecutionInstances(request *model.ListExecutionInstancesRequest) (*model.ListExecutionInstancesResponse, error) {
	requestDef := GenReqDefForListExecutionInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecutionInstancesResponse), nil
	}
}

// ListExecutionInstancesInvoker 查询工单步骤批次实例
func (c *CocClient) ListExecutionInstancesInvoker(request *model.ListExecutionInstancesRequest) *ListExecutionInstancesInvoker {
	requestDef := GenReqDefForListExecutionInstances()
	return &ListExecutionInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExecutionSteps 查询工单步骤详情
//
// 查询工单步骤详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListExecutionSteps(request *model.ListExecutionStepsRequest) (*model.ListExecutionStepsResponse, error) {
	requestDef := GenReqDefForListExecutionSteps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecutionStepsResponse), nil
	}
}

// ListExecutionStepsInvoker 查询工单步骤详情
func (c *CocClient) ListExecutionStepsInvoker(request *model.ListExecutionStepsRequest) *ListExecutionStepsInvoker {
	requestDef := GenReqDefForListExecutionSteps()
	return &ListExecutionStepsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExecutions 查询作业工单列表
//
// 查询作业工单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListExecutions(request *model.ListExecutionsRequest) (*model.ListExecutionsResponse, error) {
	requestDef := GenReqDefForListExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecutionsResponse), nil
	}
}

// ListExecutionsInvoker 查询作业工单列表
func (c *CocClient) ListExecutionsInvoker(request *model.ListExecutionsRequest) *ListExecutionsInvoker {
	requestDef := GenReqDefForListExecutions()
	return &ListExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OperateExecution 操作工单
//
// 操作工单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) OperateExecution(request *model.OperateExecutionRequest) (*model.OperateExecutionResponse, error) {
	requestDef := GenReqDefForOperateExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.OperateExecutionResponse), nil
	}
}

// OperateExecutionInvoker 操作工单
func (c *CocClient) OperateExecutionInvoker(request *model.OperateExecutionRequest) *OperateExecutionInvoker {
	requestDef := GenReqDefForOperateExecution()
	return &OperateExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCocIncident CreateExternalIncident 创建事件单
//
// CreateExternalIncident 创建事件单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateCocIncident(request *model.CreateCocIncidentRequest) (*model.CreateCocIncidentResponse, error) {
	requestDef := GenReqDefForCreateCocIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCocIncidentResponse), nil
	}
}

// CreateCocIncidentInvoker CreateExternalIncident 创建事件单
func (c *CocClient) CreateCocIncidentInvoker(request *model.CreateCocIncidentRequest) *CreateCocIncidentInvoker {
	requestDef := GenReqDefForCreateCocIncident()
	return &CreateCocIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleCocIncident HandleCocIncident处理事件单
//
// HandleCocIncident 处理事件单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) HandleCocIncident(request *model.HandleCocIncidentRequest) (*model.HandleCocIncidentResponse, error) {
	requestDef := GenReqDefForHandleCocIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleCocIncidentResponse), nil
	}
}

// HandleCocIncidentInvoker HandleCocIncident处理事件单
func (c *CocClient) HandleCocIncidentInvoker(request *model.HandleCocIncidentRequest) *HandleCocIncidentInvoker {
	requestDef := GenReqDefForHandleCocIncident()
	return &HandleCocIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCocTicketOperationHistories GetCocTicketOperationHistories 获取事件单历史
//
// ListCocTicketOperationHistories  获取事件单历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListCocTicketOperationHistories(request *model.ListCocTicketOperationHistoriesRequest) (*model.ListCocTicketOperationHistoriesResponse, error) {
	requestDef := GenReqDefForListCocTicketOperationHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCocTicketOperationHistoriesResponse), nil
	}
}

// ListCocTicketOperationHistoriesInvoker GetCocTicketOperationHistories 获取事件单历史
func (c *CocClient) ListCocTicketOperationHistoriesInvoker(request *model.ListCocTicketOperationHistoriesRequest) *ListCocTicketOperationHistoriesInvoker {
	requestDef := GenReqDefForListCocTicketOperationHistories()
	return &ListCocTicketOperationHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCocIncidentDetail GetCocIncidentDetail 获取事件单详细
//
// ShowCocIncidentDetail  获取事件单详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowCocIncidentDetail(request *model.ShowCocIncidentDetailRequest) (*model.ShowCocIncidentDetailResponse, error) {
	requestDef := GenReqDefForShowCocIncidentDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCocIncidentDetailResponse), nil
	}
}

// ShowCocIncidentDetailInvoker GetCocIncidentDetail 获取事件单详细
func (c *CocClient) ShowCocIncidentDetailInvoker(request *model.ShowCocIncidentDetailRequest) *ShowCocIncidentDetailInvoker {
	requestDef := GenReqDefForShowCocIncidentDetail()
	return &ShowCocIncidentDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCocIssues CreateExternalIssues 创建问题单
//
// CreateExternalIssues 创建问题单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateCocIssues(request *model.CreateCocIssuesRequest) (*model.CreateCocIssuesResponse, error) {
	requestDef := GenReqDefForCreateCocIssues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCocIssuesResponse), nil
	}
}

// CreateCocIssuesInvoker CreateExternalIssues 创建问题单
func (c *CocClient) CreateCocIssuesInvoker(request *model.CreateCocIssuesRequest) *CreateCocIssuesInvoker {
	requestDef := GenReqDefForCreateCocIssues()
	return &CreateCocIssuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCocIssuesDetail GetCocIssuesDetail 获取事件单详细
//
// ShowCocIssuesDetail  获取事件单详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowCocIssuesDetail(request *model.ShowCocIssuesDetailRequest) (*model.ShowCocIssuesDetailResponse, error) {
	requestDef := GenReqDefForShowCocIssuesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCocIssuesDetailResponse), nil
	}
}

// ShowCocIssuesDetailInvoker GetCocIssuesDetail 获取事件单详细
func (c *CocClient) ShowCocIssuesDetailInvoker(request *model.ShowCocIssuesDetailRequest) *ShowCocIssuesDetailInvoker {
	requestDef := GenReqDefForShowCocIssuesDetail()
	return &ShowCocIssuesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorizableTicketsExternal 查询COC可授权单列表
//
// 查询COC可授权单列表（变更单号、事件单号、warroom和告警）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListAuthorizableTicketsExternal(request *model.ListAuthorizableTicketsExternalRequest) (*model.ListAuthorizableTicketsExternalResponse, error) {
	requestDef := GenReqDefForListAuthorizableTicketsExternal()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorizableTicketsExternalResponse), nil
	}
}

// ListAuthorizableTicketsExternalInvoker 查询COC可授权单列表
func (c *CocClient) ListAuthorizableTicketsExternalInvoker(request *model.ListAuthorizableTicketsExternalRequest) *ListAuthorizableTicketsExternalInvoker {
	requestDef := GenReqDefForListAuthorizableTicketsExternal()
	return &ListAuthorizableTicketsExternalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMultiCloudResources 查询用户在云厂商中的资源
//
// 查询用户在云厂商中的资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListMultiCloudResources(request *model.ListMultiCloudResourcesRequest) (*model.ListMultiCloudResourcesResponse, error) {
	requestDef := GenReqDefForListMultiCloudResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMultiCloudResourcesResponse), nil
	}
}

// ListMultiCloudResourcesInvoker 查询用户在云厂商中的资源
func (c *CocClient) ListMultiCloudResourcesInvoker(request *model.ListMultiCloudResourcesRequest) *ListMultiCloudResourcesInvoker {
	requestDef := GenReqDefForListMultiCloudResources()
	return &ListMultiCloudResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountMultiResources 查询用户各种资源总数
//
// 查询用户各种资源总数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountMultiResources(request *model.CountMultiResourcesRequest) (*model.CountMultiResourcesResponse, error) {
	requestDef := GenReqDefForCountMultiResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountMultiResourcesResponse), nil
	}
}

// CountMultiResourcesInvoker 查询用户各种资源总数
func (c *CocClient) CountMultiResourcesInvoker(request *model.CountMultiResourcesRequest) *CountMultiResourcesInvoker {
	requestDef := GenReqDefForCountMultiResources()
	return &CountMultiResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResource 查询用户所有资源
//
// 查询用户所有资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResource(request *model.ListResourceRequest) (*model.ListResourceResponse, error) {
	requestDef := GenReqDefForListResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceResponse), nil
	}
}

// ListResourceInvoker 查询用户所有资源
func (c *CocClient) ListResourceInvoker(request *model.ListResourceRequest) *ListResourceInvoker {
	requestDef := GenReqDefForListResource()
	return &ListResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncResource 从RMS同步用户所有资源
//
// 从RMS同步用户所有资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncResource(request *model.SyncResourceRequest) (*model.SyncResourceResponse, error) {
	requestDef := GenReqDefForSyncResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncResourceResponse), nil
	}
}

// SyncResourceInvoker 从RMS同步用户所有资源
func (c *CocClient) SyncResourceInvoker(request *model.SyncResourceRequest) *SyncResourceInvoker {
	requestDef := GenReqDefForSyncResource()
	return &SyncResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptResourceTags 查询资源标签列表
//
// 查询资源标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScriptResourceTags(request *model.ListScriptResourceTagsRequest) (*model.ListScriptResourceTagsResponse, error) {
	requestDef := GenReqDefForListScriptResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptResourceTagsResponse), nil
	}
}

// ListScriptResourceTagsInvoker 查询资源标签列表
func (c *CocClient) ListScriptResourceTagsInvoker(request *model.ListScriptResourceTagsRequest) *ListScriptResourceTagsInvoker {
	requestDef := GenReqDefForListScriptResourceTags()
	return &ListScriptResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceTags 更新资源标签
//
// 更新资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateResourceTags(request *model.UpdateResourceTagsRequest) (*model.UpdateResourceTagsResponse, error) {
	requestDef := GenReqDefForUpdateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceTagsResponse), nil
	}
}

// UpdateResourceTagsInvoker 更新资源标签
func (c *CocClient) UpdateResourceTagsInvoker(request *model.UpdateResourceTagsRequest) *UpdateResourceTagsInvoker {
	requestDef := GenReqDefForUpdateResourceTags()
	return &UpdateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScheduledTask 新建定时运维
//
// Create Scheduled Task
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateScheduledTask(request *model.CreateScheduledTaskRequest) (*model.CreateScheduledTaskResponse, error) {
	requestDef := GenReqDefForCreateScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScheduledTaskResponse), nil
	}
}

// CreateScheduledTaskInvoker 新建定时运维
func (c *CocClient) CreateScheduledTaskInvoker(request *model.CreateScheduledTaskRequest) *CreateScheduledTaskInvoker {
	requestDef := GenReqDefForCreateScheduledTask()
	return &CreateScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScheduledTask 删除ScheduledTask
//
// Delete scheduled task by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteScheduledTask(request *model.DeleteScheduledTaskRequest) (*model.DeleteScheduledTaskResponse, error) {
	requestDef := GenReqDefForDeleteScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScheduledTaskResponse), nil
	}
}

// DeleteScheduledTaskInvoker 删除ScheduledTask
func (c *CocClient) DeleteScheduledTaskInvoker(request *model.DeleteScheduledTaskRequest) *DeleteScheduledTaskInvoker {
	requestDef := GenReqDefForDeleteScheduledTask()
	return &DeleteScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableScheduledTask 禁用ScheduledTask
//
// Disable scheduled task by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DisableScheduledTask(request *model.DisableScheduledTaskRequest) (*model.DisableScheduledTaskResponse, error) {
	requestDef := GenReqDefForDisableScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableScheduledTaskResponse), nil
	}
}

// DisableScheduledTaskInvoker 禁用ScheduledTask
func (c *CocClient) DisableScheduledTaskInvoker(request *model.DisableScheduledTaskRequest) *DisableScheduledTaskInvoker {
	requestDef := GenReqDefForDisableScheduledTask()
	return &DisableScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableScheduledTask 启用ScheduledTask
//
// Enable scheduled task by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) EnableScheduledTask(request *model.EnableScheduledTaskRequest) (*model.EnableScheduledTaskResponse, error) {
	requestDef := GenReqDefForEnableScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableScheduledTaskResponse), nil
	}
}

// EnableScheduledTaskInvoker 启用ScheduledTask
func (c *CocClient) EnableScheduledTaskInvoker(request *model.EnableScheduledTaskRequest) *EnableScheduledTaskInvoker {
	requestDef := GenReqDefForEnableScheduledTask()
	return &EnableScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTask 查询ScheduledTask列表
//
// Get ScheduledTask infos
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScheduledTask(request *model.ListScheduledTaskRequest) (*model.ListScheduledTaskResponse, error) {
	requestDef := GenReqDefForListScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTaskResponse), nil
	}
}

// ListScheduledTaskInvoker 查询ScheduledTask列表
func (c *CocClient) ListScheduledTaskInvoker(request *model.ListScheduledTaskRequest) *ListScheduledTaskInvoker {
	requestDef := GenReqDefForListScheduledTask()
	return &ListScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledTaskHistory 查询定时运维历史记录
//
// get scheduled task history list
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScheduledTaskHistory(request *model.ListScheduledTaskHistoryRequest) (*model.ListScheduledTaskHistoryResponse, error) {
	requestDef := GenReqDefForListScheduledTaskHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledTaskHistoryResponse), nil
	}
}

// ListScheduledTaskHistoryInvoker 查询定时运维历史记录
func (c *CocClient) ListScheduledTaskHistoryInvoker(request *model.ListScheduledTaskHistoryRequest) *ListScheduledTaskHistoryInvoker {
	requestDef := GenReqDefForListScheduledTaskHistory()
	return &ListScheduledTaskHistoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowScheduledTask 查询ScheduledTask
//
// Get ScheduledTask info by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowScheduledTask(request *model.ShowScheduledTaskRequest) (*model.ShowScheduledTaskResponse, error) {
	requestDef := GenReqDefForShowScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScheduledTaskResponse), nil
	}
}

// ShowScheduledTaskInvoker 查询ScheduledTask
func (c *CocClient) ShowScheduledTaskInvoker(request *model.ShowScheduledTaskRequest) *ShowScheduledTaskInvoker {
	requestDef := GenReqDefForShowScheduledTask()
	return &ShowScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScheduledTask 修改ScheduledTask
//
// Update ScheduledTask
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateScheduledTask(request *model.UpdateScheduledTaskRequest) (*model.UpdateScheduledTaskResponse, error) {
	requestDef := GenReqDefForUpdateScheduledTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScheduledTaskResponse), nil
	}
}

// UpdateScheduledTaskInvoker 修改ScheduledTask
func (c *CocClient) UpdateScheduledTaskInvoker(request *model.UpdateScheduledTaskRequest) *UpdateScheduledTaskInvoker {
	requestDef := GenReqDefForUpdateScheduledTask()
	return &UpdateScheduledTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobBatch 展示批次详情
//
// 查询：批次详情，分页获取批次中的实例列表。
// 过滤条件：分页参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobBatch(request *model.GetScriptJobBatchRequest) (*model.GetScriptJobBatchResponse, error) {
	requestDef := GenReqDefForGetScriptJobBatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobBatchResponse), nil
	}
}

// GetScriptJobBatchInvoker 展示批次详情
func (c *CocClient) GetScriptJobBatchInvoker(request *model.GetScriptJobBatchRequest) *GetScriptJobBatchInvoker {
	requestDef := GenReqDefForGetScriptJobBatch()
	return &GetScriptJobBatchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobInfo 展示脚本工单基本信息
//
// 查询执行：基本信息
// 执行类型、执行名称、创建人、创建时间、结束时间、执行状态、标签（脚本id，脚本名，执行脚本参数，执行用户，超时时长、成功率阈值）
//
// 不同的任务类型消费标签中的不同key
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobInfo(request *model.GetScriptJobInfoRequest) (*model.GetScriptJobInfoResponse, error) {
	requestDef := GenReqDefForGetScriptJobInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobInfoResponse), nil
	}
}

// GetScriptJobInfoInvoker 展示脚本工单基本信息
func (c *CocClient) GetScriptJobInfoInvoker(request *model.GetScriptJobInfoRequest) *GetScriptJobInfoInvoker {
	requestDef := GenReqDefForGetScriptJobInfo()
	return &GetScriptJobInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScriptJobStatistics 展示实例状态统计信息
//
// 查询：实例状态统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScriptJobStatistics(request *model.GetScriptJobStatisticsRequest) (*model.GetScriptJobStatisticsResponse, error) {
	requestDef := GenReqDefForGetScriptJobStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptJobStatisticsResponse), nil
	}
}

// GetScriptJobStatisticsInvoker 展示实例状态统计信息
func (c *CocClient) GetScriptJobStatisticsInvoker(request *model.GetScriptJobStatisticsRequest) *GetScriptJobStatisticsInvoker {
	requestDef := GenReqDefForGetScriptJobStatistics()
	return &GetScriptJobStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptJobBatches 展示批次列表
//
// 查询：批次列表
// 返回：批次index、批次标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScriptJobBatches(request *model.ListScriptJobBatchesRequest) (*model.ListScriptJobBatchesResponse, error) {
	requestDef := GenReqDefForListScriptJobBatches()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptJobBatchesResponse), nil
	}
}

// ListScriptJobBatchesInvoker 展示批次列表
func (c *CocClient) ListScriptJobBatchesInvoker(request *model.ListScriptJobBatchesRequest) *ListScriptJobBatchesInvoker {
	requestDef := GenReqDefForListScriptJobBatches()
	return &ListScriptJobBatchesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScriptJobs 展示工单列表
//
// 查询作业工单列表，分页查询
// 过滤：创建时间开始，创建时间结束、创建人
// 返回：id、脚本名称、区域、创建人、创建时间、结束时间、总耗时、状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScriptJobs(request *model.ListScriptJobsRequest) (*model.ListScriptJobsResponse, error) {
	requestDef := GenReqDefForListScriptJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptJobsResponse), nil
	}
}

// ListScriptJobsInvoker 展示工单列表
func (c *CocClient) ListScriptJobsInvoker(request *model.ListScriptJobsRequest) *ListScriptJobsInvoker {
	requestDef := GenReqDefForListScriptJobs()
	return &ListScriptJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// OperateScriptJob 操作脚本工单
//
// 操作类型：取消实例、跳过批次、取消整个工单、暂停整个工单、继续整个工单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) OperateScriptJob(request *model.OperateScriptJobRequest) (*model.OperateScriptJobResponse, error) {
	requestDef := GenReqDefForOperateScriptJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.OperateScriptJobResponse), nil
	}
}

// OperateScriptJobInvoker 操作脚本工单
func (c *CocClient) OperateScriptJobInvoker(request *model.OperateScriptJobRequest) *OperateScriptJobInvoker {
	requestDef := GenReqDefForOperateScriptJob()
	return &OperateScriptJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AcceptScript 审批待审批的脚本
//
// 功能：审批脚本。
// 约束条件：只有创建脚本填写了审批人，脚本为待审批状态才能审批。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) AcceptScript(request *model.AcceptScriptRequest) (*model.AcceptScriptResponse, error) {
	requestDef := GenReqDefForAcceptScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptScriptResponse), nil
	}
}

// AcceptScriptInvoker 审批待审批的脚本
func (c *CocClient) AcceptScriptInvoker(request *model.AcceptScriptRequest) *AcceptScriptInvoker {
	requestDef := GenReqDefForAcceptScript()
	return &AcceptScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckScriptRisk 评估脚本风险等级
//
// 根据作业内容，对作业评估风险，返回相关分析的结果和信息，结果仅供参考。
// 高危命令指影响系统或服务的正常运行，或造成系统特殊文件被恶意删除或修改命令。 高危命令检测通过校验规则正则匹配脚本内容中是否包含高危命令。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CheckScriptRisk(request *model.CheckScriptRiskRequest) (*model.CheckScriptRiskResponse, error) {
	requestDef := GenReqDefForCheckScriptRisk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckScriptRiskResponse), nil
	}
}

// CheckScriptRiskInvoker 评估脚本风险等级
func (c *CocClient) CheckScriptRiskInvoker(request *model.CheckScriptRiskRequest) *CheckScriptRiskInvoker {
	requestDef := GenReqDefForCheckScriptRisk()
	return &CheckScriptRiskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateScript 创建脚本
//
// 创建作业脚本：自定义脚本
// - 脚本有标签属性，表示是高危脚本。创建时候不需要对脚本进行是否是高危的二次校验。
// - 进行租户隔离；北向接口创建的脚本，审批人字段不填写，默认不需要审批
// - 约束条件：
// - 脚本名称：同一租户下，脚本名称不能重复，最大字符64个字符，支持中文+字母+数字+下划线。
// - 脚本内容最大100kb。
// - 脚本参数个数最多20个。
// - 脚本描述：最大256个字符。
// - 单个参数的参数名称 64个字符，只支持字母+数字+下划线。
// - 单个参数的值最大1024个字符，正则表达式如下：^((?!\\.{2,})[a-zA-Z0-9_\\-/\\.\\*\\x20\\?:\&quot;,&#x3D;+@\\\\\\[\\{\\]\\}])*$。
// - 审批人最多支持5人。
// - 脚本输出的日志总量只支持1MB。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateScript(request *model.CreateScriptRequest) (*model.CreateScriptResponse, error) {
	requestDef := GenReqDefForCreateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateScriptResponse), nil
	}
}

// CreateScriptInvoker 创建脚本
func (c *CocClient) CreateScriptInvoker(request *model.CreateScriptRequest) *CreateScriptInvoker {
	requestDef := GenReqDefForCreateScript()
	return &CreateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteScript 删除自定义脚本
//
// 删除作业脚本：自定义脚本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteScript(request *model.DeleteScriptRequest) (*model.DeleteScriptResponse, error) {
	requestDef := GenReqDefForDeleteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteScriptResponse), nil
	}
}

// DeleteScriptInvoker 删除自定义脚本
func (c *CocClient) DeleteScriptInvoker(request *model.DeleteScriptRequest) *DeleteScriptInvoker {
	requestDef := GenReqDefForDeleteScript()
	return &DeleteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteScript 执行自定义脚本
//
// 执行脚本
//
// 脚本入参、超时时间、执行用户、资源受限
// 脚本入参支持20个。
// 单次下发的机器支持200个。
// 单次批次内机器数量最大10个。
// 最大批次数量为20批。
// 脚本输出的日志总量只支持1MB。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecuteScript(request *model.ExecuteScriptRequest) (*model.ExecuteScriptResponse, error) {
	requestDef := GenReqDefForExecuteScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteScriptResponse), nil
	}
}

// ExecuteScriptInvoker 执行自定义脚本
func (c *CocClient) ExecuteScriptInvoker(request *model.ExecuteScriptRequest) *ExecuteScriptInvoker {
	requestDef := GenReqDefForExecuteScript()
	return &ExecuteScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScript 获取自定义脚本详情
//
// 获取脚本详情
// 约束条件：
// 只能查询自定义脚本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetScript(request *model.GetScriptRequest) (*model.GetScriptResponse, error) {
	requestDef := GenReqDefForGetScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScriptResponse), nil
	}
}

// GetScriptInvoker 获取自定义脚本详情
func (c *CocClient) GetScriptInvoker(request *model.GetScriptRequest) *GetScriptInvoker {
	requestDef := GenReqDefForGetScript()
	return &GetScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstancesBatch 获取自动分批结果
//
// 根据分批策略获取分批结果，只支持自动分批：
// 规则如下：
// 1.单个批次的所有实例必须属于同一个区域；
//      * 2.单个批次的所有实例必须属于同一个可用区；
//      * 3.单个批次的所有实例必须属于同一个应用；
//      * 4.单个批次内同一分组下的实例不超过50%（除分组下仅以一个实例的情况外）；
//      * 5.前三批每批节点数量不超过10。
//      * 6.每批次实例数量不超过10。
//
//    总机器数量为200。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListInstancesBatch(request *model.ListInstancesBatchRequest) (*model.ListInstancesBatchResponse, error) {
	requestDef := GenReqDefForListInstancesBatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesBatchResponse), nil
	}
}

// ListInstancesBatchInvoker 获取自动分批结果
func (c *CocClient) ListInstancesBatchInvoker(request *model.ListInstancesBatchRequest) *ListInstancesBatchInvoker {
	requestDef := GenReqDefForListInstancesBatch()
	return &ListInstancesBatchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScripts 查询脚本列表
//
// 作业脚本列表：自定义脚本
//
// limit最大为100
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListScripts(request *model.ListScriptsRequest) (*model.ListScriptsResponse, error) {
	requestDef := GenReqDefForListScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScriptsResponse), nil
	}
}

// ListScriptsInvoker 查询脚本列表
func (c *CocClient) ListScriptsInvoker(request *model.ListScriptsRequest) *ListScriptsInvoker {
	requestDef := GenReqDefForListScripts()
	return &ListScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateScript 修改脚本
//
// 修改作业脚本：自定义脚本
// 约束条件：
// 脚本名称：同一租户下，脚本名称不能重复，最大字符64个字符，支持中文+字母+数字+下划线。
// 脚本内容最大4096个字符。
// 脚本参数个数最多20个。
// 脚本描述：最大256个字符。
// 单个参数的参数名称 64个字符，只支持字母+数字+下划线。
// 单个参数的值最大1024个字符，正则表达式如下：^((?!.{2,})[a-zA-Z0-9_-/.*\\x20?:\&quot;,&#x3D;+@\\[{]}])*$。
// 修改的脚本如果有审批人，在修改之后，需要再次选择审批人，查询审批
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateScript(request *model.UpdateScriptRequest) (*model.UpdateScriptResponse, error) {
	requestDef := GenReqDefForUpdateScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateScriptResponse), nil
	}
}

// UpdateScriptInvoker 修改脚本
func (c *CocClient) UpdateScriptInvoker(request *model.UpdateScriptRequest) *UpdateScriptInvoker {
	requestDef := GenReqDefForUpdateScript()
	return &UpdateScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecutePublicScript 执行公共脚本
//
// 执行公共脚本
// 脚本入参、超时时间、执行用户、资源受限
// 脚本入参支持20个。
// 单次下发的机器支持200个。
// 单次批次内机器数量最大10个。
// 最大批次数量为20批。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecutePublicScript(request *model.ExecutePublicScriptRequest) (*model.ExecutePublicScriptResponse, error) {
	requestDef := GenReqDefForExecutePublicScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecutePublicScriptResponse), nil
	}
}

// ExecutePublicScriptInvoker 执行公共脚本
func (c *CocClient) ExecutePublicScriptInvoker(request *model.ExecutePublicScriptRequest) *ExecutePublicScriptInvoker {
	requestDef := GenReqDefForExecutePublicScript()
	return &ExecutePublicScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetPublicScript 展示公共脚本详情
//
// 展示公共脚本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) GetPublicScript(request *model.GetPublicScriptRequest) (*model.GetPublicScriptResponse, error) {
	requestDef := GenReqDefForGetPublicScript()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetPublicScriptResponse), nil
	}
}

// GetPublicScriptInvoker 展示公共脚本详情
func (c *CocClient) GetPublicScriptInvoker(request *model.GetPublicScriptRequest) *GetPublicScriptInvoker {
	requestDef := GenReqDefForGetPublicScript()
	return &GetPublicScriptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPublicScripts 获取公共脚本列表
//
// 获取公共脚本列表，分页逻辑：采用limit+marker方式，提高分页效率。用自增id作为marker参数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListPublicScripts(request *model.ListPublicScriptsRequest) (*model.ListPublicScriptsResponse, error) {
	requestDef := GenReqDefForListPublicScripts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicScriptsResponse), nil
	}
}

// ListPublicScriptsInvoker 获取公共脚本列表
func (c *CocClient) ListPublicScriptsInvoker(request *model.ListPublicScriptsRequest) *ListPublicScriptsInvoker {
	requestDef := GenReqDefForListPublicScripts()
	return &ListPublicScriptsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWarRoom 创建租户区WarRoom
//
// 创建租户区WarRoom
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateWarRoom(request *model.CreateWarRoomRequest) (*model.CreateWarRoomResponse, error) {
	requestDef := GenReqDefForCreateWarRoom()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWarRoomResponse), nil
	}
}

// CreateWarRoomInvoker 创建租户区WarRoom
func (c *CocClient) CreateWarRoomInvoker(request *model.CreateWarRoomRequest) *CreateWarRoomInvoker {
	requestDef := GenReqDefForCreateWarRoom()
	return &CreateWarRoomInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWarRooms 查询租户区WarRoom信息列表
//
// 查询租户区WarRoom信息列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListWarRooms(request *model.ListWarRoomsRequest) (*model.ListWarRoomsResponse, error) {
	requestDef := GenReqDefForListWarRooms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWarRoomsResponse), nil
	}
}

// ListWarRoomsInvoker 查询租户区WarRoom信息列表
func (c *CocClient) ListWarRoomsInvoker(request *model.ListWarRoomsRequest) *ListWarRoomsInvoker {
	requestDef := GenReqDefForListWarRooms()
	return &ListWarRoomsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
