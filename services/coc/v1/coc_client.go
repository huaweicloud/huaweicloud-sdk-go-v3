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

// CreatePasswordChangePlan 创建改密计划
//
// 创建改密计划
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreatePasswordChangePlan(request *model.CreatePasswordChangePlanRequest) (*model.CreatePasswordChangePlanResponse, error) {
	requestDef := GenReqDefForCreatePasswordChangePlan()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePasswordChangePlanResponse), nil
	}
}

// CreatePasswordChangePlanInvoker 创建改密计划
func (c *CocClient) CreatePasswordChangePlanInvoker(request *model.CreatePasswordChangePlanRequest) *CreatePasswordChangePlanInvoker {
	requestDef := GenReqDefForCreatePasswordChangePlan()
	return &CreatePasswordChangePlanInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetAccountPassword 主机密码重置
//
// 主机密码重置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ResetAccountPassword(request *model.ResetAccountPasswordRequest) (*model.ResetAccountPasswordResponse, error) {
	requestDef := GenReqDefForResetAccountPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetAccountPasswordResponse), nil
	}
}

// ResetAccountPasswordInvoker 主机密码重置
func (c *CocClient) ResetAccountPasswordInvoker(request *model.ResetAccountPasswordRequest) *ResetAccountPasswordInvoker {
	requestDef := GenReqDefForResetAccountPassword()
	return &ResetAccountPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAccountPassword 回写改密结果
//
// 回写改密结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateAccountPassword(request *model.UpdateAccountPasswordRequest) (*model.UpdateAccountPasswordResponse, error) {
	requestDef := GenReqDefForUpdateAccountPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAccountPasswordResponse), nil
	}
}

// UpdateAccountPasswordInvoker 回写改密结果
func (c *CocClient) UpdateAccountPasswordInvoker(request *model.UpdateAccountPasswordRequest) *UpdateAccountPasswordInvoker {
	requestDef := GenReqDefForUpdateAccountPassword()
	return &UpdateAccountPasswordInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ClearAlarm 批量清除告警
//
// 清除告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ClearAlarm(request *model.ClearAlarmRequest) (*model.ClearAlarmResponse, error) {
	requestDef := GenReqDefForClearAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ClearAlarmResponse), nil
	}
}

// ClearAlarmInvoker 批量清除告警
func (c *CocClient) ClearAlarmInvoker(request *model.ClearAlarmRequest) *ClearAlarmInvoker {
	requestDef := GenReqDefForClearAlarm()
	return &ClearAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandlerAlarm 自动处理告警
//
// 自动处理告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) HandlerAlarm(request *model.HandlerAlarmRequest) (*model.HandlerAlarmResponse, error) {
	requestDef := GenReqDefForHandlerAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandlerAlarmResponse), nil
	}
}

// HandlerAlarmInvoker 自动处理告警
func (c *CocClient) HandlerAlarmInvoker(request *model.HandlerAlarmRequest) *HandlerAlarmInvoker {
	requestDef := GenReqDefForHandlerAlarm()
	return &HandlerAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmHandleHistories 查询告警工单历史
//
// 查询告警工单历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListAlarmHandleHistories(request *model.ListAlarmHandleHistoriesRequest) (*model.ListAlarmHandleHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHandleHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHandleHistoriesResponse), nil
	}
}

// ListAlarmHandleHistoriesInvoker 查询告警工单历史
func (c *CocClient) ListAlarmHandleHistoriesInvoker(request *model.ListAlarmHandleHistoriesRequest) *ListAlarmHandleHistoriesInvoker {
	requestDef := GenReqDefForListAlarmHandleHistories()
	return &ListAlarmHandleHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarm 查询Alarm
//
// Get alarm info by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowAlarm(request *model.ShowAlarmRequest) (*model.ShowAlarmResponse, error) {
	requestDef := GenReqDefForShowAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmResponse), nil
	}
}

// ShowAlarmInvoker 查询Alarm
func (c *CocClient) ShowAlarmInvoker(request *model.ShowAlarmRequest) *ShowAlarmInvoker {
	requestDef := GenReqDefForShowAlarm()
	return &ShowAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TransferAlarmToIncident 批量告警转事件
//
// 批量告警转事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) TransferAlarmToIncident(request *model.TransferAlarmToIncidentRequest) (*model.TransferAlarmToIncidentResponse, error) {
	requestDef := GenReqDefForTransferAlarmToIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TransferAlarmToIncidentResponse), nil
	}
}

// TransferAlarmToIncidentInvoker 批量告警转事件
func (c *CocClient) TransferAlarmToIncidentInvoker(request *model.TransferAlarmToIncidentRequest) *TransferAlarmToIncidentInvoker {
	requestDef := GenReqDefForTransferAlarmToIncident()
	return &TransferAlarmToIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplication 创建应用
//
// 创建应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateApplication(request *model.CreateApplicationRequest) (*model.CreateApplicationResponse, error) {
	requestDef := GenReqDefForCreateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationResponse), nil
	}
}

// CreateApplicationInvoker 创建应用
func (c *CocClient) CreateApplicationInvoker(request *model.CreateApplicationRequest) *CreateApplicationInvoker {
	requestDef := GenReqDefForCreateApplication()
	return &CreateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplication 删除应用
//
// 删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteApplication(request *model.DeleteApplicationRequest) (*model.DeleteApplicationResponse, error) {
	requestDef := GenReqDefForDeleteApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationResponse), nil
	}
}

// DeleteApplicationInvoker 删除应用
func (c *CocClient) DeleteApplicationInvoker(request *model.DeleteApplicationRequest) *DeleteApplicationInvoker {
	requestDef := GenReqDefForDeleteApplication()
	return &DeleteApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplication 修改应用
//
// 修改应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateApplication(request *model.UpdateApplicationRequest) (*model.UpdateApplicationResponse, error) {
	requestDef := GenReqDefForUpdateApplication()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationResponse), nil
	}
}

// UpdateApplicationInvoker 修改应用
func (c *CocClient) UpdateApplicationInvoker(request *model.UpdateApplicationRequest) *UpdateApplicationInvoker {
	requestDef := GenReqDefForUpdateApplication()
	return &UpdateApplicationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateApplicationView 批量创建应用，分组，组件
//
// 批量创建应用，分组，组件。
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

// BatchCreateApplicationViewInvoker 批量创建应用，分组，组件
func (c *CocClient) BatchCreateApplicationViewInvoker(request *model.BatchCreateApplicationViewRequest) *BatchCreateApplicationViewInvoker {
	requestDef := GenReqDefForBatchCreateApplicationView()
	return &BatchCreateApplicationViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationView 查询应用、组件、分组名称列表
//
// 查询应用、组件、分组名称列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplicationView(request *model.ListApplicationViewRequest) (*model.ListApplicationViewResponse, error) {
	requestDef := GenReqDefForListApplicationView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationViewResponse), nil
	}
}

// ListApplicationViewInvoker 查询应用、组件、分组名称列表
func (c *CocClient) ListApplicationViewInvoker(request *model.ListApplicationViewRequest) *ListApplicationViewInvoker {
	requestDef := GenReqDefForListApplicationView()
	return &ListApplicationViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAssessTask 创建应用评估任务
//
// 创建应用评估任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateAssessTask(request *model.CreateAssessTaskRequest) (*model.CreateAssessTaskResponse, error) {
	requestDef := GenReqDefForCreateAssessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAssessTaskResponse), nil
	}
}

// CreateAssessTaskInvoker 创建应用评估任务
func (c *CocClient) CreateAssessTaskInvoker(request *model.CreateAssessTaskRequest) *CreateAssessTaskInvoker {
	requestDef := GenReqDefForCreateAssessTask()
	return &CreateAssessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAssessTask 分页查询评估任务列表
//
// 分页查询评估任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListAssessTask(request *model.ListAssessTaskRequest) (*model.ListAssessTaskResponse, error) {
	requestDef := GenReqDefForListAssessTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAssessTaskResponse), nil
	}
}

// ListAssessTaskInvoker 分页查询评估任务列表
func (c *CocClient) ListAssessTaskInvoker(request *model.ListAssessTaskRequest) *ListAssessTaskInvoker {
	requestDef := GenReqDefForListAssessTask()
	return &ListAssessTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateChange UpdateChange 更新变更单
//
// UpdateChange 更新变更单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateChange(request *model.UpdateChangeRequest) (*model.UpdateChangeResponse, error) {
	requestDef := GenReqDefForUpdateChange()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateChangeResponse), nil
	}
}

// UpdateChangeInvoker UpdateChange 更新变更单
func (c *CocClient) UpdateChangeInvoker(request *model.UpdateChangeRequest) *UpdateChangeInvoker {
	requestDef := GenReqDefForUpdateChange()
	return &UpdateChangeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// HandleIncident HandleIncident 处理事件单
//
// HandleIncident 处理事件单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) HandleIncident(request *model.HandleIncidentRequest) (*model.HandleIncidentResponse, error) {
	requestDef := GenReqDefForHandleIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.HandleIncidentResponse), nil
	}
}

// HandleIncidentInvoker HandleIncident 处理事件单
func (c *CocClient) HandleIncidentInvoker(request *model.HandleIncidentRequest) *HandleIncidentInvoker {
	requestDef := GenReqDefForHandleIncident()
	return &HandleIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidents ListIncidents 查询事件单列表
//
// ListIncidents 查询事件单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListIncidents(request *model.ListIncidentsRequest) (*model.ListIncidentsResponse, error) {
	requestDef := GenReqDefForListIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentsResponse), nil
	}
}

// ListIncidentsInvoker ListIncidents 查询事件单列表
func (c *CocClient) ListIncidentsInvoker(request *model.ListIncidentsRequest) *ListIncidentsInvoker {
	requestDef := GenReqDefForListIncidents()
	return &ListIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidentsHistories ListIncidentsHistories 获取事件单历史
//
// ListIncidentsHistories  获取事件单历史
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListIncidentsHistories(request *model.ListIncidentsHistoriesRequest) (*model.ListIncidentsHistoriesResponse, error) {
	requestDef := GenReqDefForListIncidentsHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentsHistoriesResponse), nil
	}
}

// ListIncidentsHistoriesInvoker ListIncidentsHistories 获取事件单历史
func (c *CocClient) ListIncidentsHistoriesInvoker(request *model.ListIncidentsHistoriesRequest) *ListIncidentsHistoriesInvoker {
	requestDef := GenReqDefForListIncidentsHistories()
	return &ListIncidentsHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIncidentTask ShowIncidentTask 获取事件任务
//
// ShowIncidentTask 获取事件任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowIncidentTask(request *model.ShowIncidentTaskRequest) (*model.ShowIncidentTaskResponse, error) {
	requestDef := GenReqDefForShowIncidentTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIncidentTaskResponse), nil
	}
}

// ShowIncidentTaskInvoker ShowIncidentTask 获取事件任务
func (c *CocClient) ShowIncidentTaskInvoker(request *model.ShowIncidentTaskRequest) *ShowIncidentTaskInvoker {
	requestDef := GenReqDefForShowIncidentTask()
	return &ShowIncidentTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateApplicationComponents 创建组件
//
// 创建组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateApplicationComponents(request *model.CreateApplicationComponentsRequest) (*model.CreateApplicationComponentsResponse, error) {
	requestDef := GenReqDefForCreateApplicationComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationComponentsResponse), nil
	}
}

// CreateApplicationComponentsInvoker 创建组件
func (c *CocClient) CreateApplicationComponentsInvoker(request *model.CreateApplicationComponentsRequest) *CreateApplicationComponentsInvoker {
	requestDef := GenReqDefForCreateApplicationComponents()
	return &CreateApplicationComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationComponent 删除组件
//
// 删除组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteApplicationComponent(request *model.DeleteApplicationComponentRequest) (*model.DeleteApplicationComponentResponse, error) {
	requestDef := GenReqDefForDeleteApplicationComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationComponentResponse), nil
	}
}

// DeleteApplicationComponentInvoker 删除组件
func (c *CocClient) DeleteApplicationComponentInvoker(request *model.DeleteApplicationComponentRequest) *DeleteApplicationComponentInvoker {
	requestDef := GenReqDefForDeleteApplicationComponent()
	return &DeleteApplicationComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationComponents 查询组件
//
// 查询组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplicationComponents(request *model.ListApplicationComponentsRequest) (*model.ListApplicationComponentsResponse, error) {
	requestDef := GenReqDefForListApplicationComponents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationComponentsResponse), nil
	}
}

// ListApplicationComponentsInvoker 查询组件
func (c *CocClient) ListApplicationComponentsInvoker(request *model.ListApplicationComponentsRequest) *ListApplicationComponentsInvoker {
	requestDef := GenReqDefForListApplicationComponents()
	return &ListApplicationComponentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationComponent 修改组件
//
// 修改组件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateApplicationComponent(request *model.UpdateApplicationComponentRequest) (*model.UpdateApplicationComponentResponse, error) {
	requestDef := GenReqDefForUpdateApplicationComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationComponentResponse), nil
	}
}

// UpdateApplicationComponentInvoker 修改组件
func (c *CocClient) UpdateApplicationComponentInvoker(request *model.UpdateApplicationComponentRequest) *UpdateApplicationComponentInvoker {
	requestDef := GenReqDefForUpdateApplicationComponent()
	return &UpdateApplicationComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CountOtherResource 查询线下IDC资源数量
//
// 查询IDC离线资源的数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountOtherResource(request *model.CountOtherResourceRequest) (*model.CountOtherResourceResponse, error) {
	requestDef := GenReqDefForCountOtherResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountOtherResourceResponse), nil
	}
}

// CountOtherResourceInvoker 查询线下IDC资源数量
func (c *CocClient) CountOtherResourceInvoker(request *model.CountOtherResourceRequest) *CountOtherResourceInvoker {
	requestDef := GenReqDefForCountOtherResource()
	return &CountOtherResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportOtherResource 导入IDC离线资源
//
// 管理线下设备提供IDC离线资源纳管功能-导入IDC离线资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ImportOtherResource(request *model.ImportOtherResourceRequest) (*model.ImportOtherResourceResponse, error) {
	requestDef := GenReqDefForImportOtherResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportOtherResourceResponse), nil
	}
}

// ImportOtherResourceInvoker 导入IDC离线资源
func (c *CocClient) ImportOtherResourceInvoker(request *model.ImportOtherResourceRequest) *ImportOtherResourceInvoker {
	requestDef := GenReqDefForImportOtherResource()
	return &ImportOtherResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelDiagnosisTask 取消诊断任务
//
// 取消诊断任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CancelDiagnosisTask(request *model.CancelDiagnosisTaskRequest) (*model.CancelDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCancelDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelDiagnosisTaskResponse), nil
	}
}

// CancelDiagnosisTaskInvoker 取消诊断任务
func (c *CocClient) CancelDiagnosisTaskInvoker(request *model.CancelDiagnosisTaskRequest) *CancelDiagnosisTaskInvoker {
	requestDef := GenReqDefForCancelDiagnosisTask()
	return &CancelDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDiagnosisTask 提交诊断任务
//
// 提交诊断任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateDiagnosisTask(request *model.CreateDiagnosisTaskRequest) (*model.CreateDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForCreateDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDiagnosisTaskResponse), nil
	}
}

// CreateDiagnosisTaskInvoker 提交诊断任务
func (c *CocClient) CreateDiagnosisTaskInvoker(request *model.CreateDiagnosisTaskRequest) *CreateDiagnosisTaskInvoker {
	requestDef := GenReqDefForCreateDiagnosisTask()
	return &CreateDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDiagnosisTasks 查询诊断记录
//
// 查询诊断记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListDiagnosisTasks(request *model.ListDiagnosisTasksRequest) (*model.ListDiagnosisTasksResponse, error) {
	requestDef := GenReqDefForListDiagnosisTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDiagnosisTasksResponse), nil
	}
}

// ListDiagnosisTasksInvoker 查询诊断记录
func (c *CocClient) ListDiagnosisTasksInvoker(request *model.ListDiagnosisTasksRequest) *ListDiagnosisTasksInvoker {
	requestDef := GenReqDefForListDiagnosisTasks()
	return &ListDiagnosisTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryDiagnosisTask 重试诊断任务
//
// 重试诊断任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) RetryDiagnosisTask(request *model.RetryDiagnosisTaskRequest) (*model.RetryDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForRetryDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryDiagnosisTaskResponse), nil
	}
}

// RetryDiagnosisTaskInvoker 重试诊断任务
func (c *CocClient) RetryDiagnosisTaskInvoker(request *model.RetryDiagnosisTaskRequest) *RetryDiagnosisTaskInvoker {
	requestDef := GenReqDefForRetryDiagnosisTask()
	return &RetryDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisNode 查询指定诊断记录下的指定诊断步骤的详情
//
// 查询指定诊断记录下的指定诊断步骤的详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowDiagnosisNode(request *model.ShowDiagnosisNodeRequest) (*model.ShowDiagnosisNodeResponse, error) {
	requestDef := GenReqDefForShowDiagnosisNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisNodeResponse), nil
	}
}

// ShowDiagnosisNodeInvoker 查询指定诊断记录下的指定诊断步骤的详情
func (c *CocClient) ShowDiagnosisNodeInvoker(request *model.ShowDiagnosisNodeRequest) *ShowDiagnosisNodeInvoker {
	requestDef := GenReqDefForShowDiagnosisNode()
	return &ShowDiagnosisNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisSummary 查询批量诊断任务的结果概要
//
// 查询诊断任务的结果概要
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowDiagnosisSummary(request *model.ShowDiagnosisSummaryRequest) (*model.ShowDiagnosisSummaryResponse, error) {
	requestDef := GenReqDefForShowDiagnosisSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisSummaryResponse), nil
	}
}

// ShowDiagnosisSummaryInvoker 查询批量诊断任务的结果概要
func (c *CocClient) ShowDiagnosisSummaryInvoker(request *model.ShowDiagnosisSummaryRequest) *ShowDiagnosisSummaryInvoker {
	requestDef := GenReqDefForShowDiagnosisSummary()
	return &ShowDiagnosisSummaryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDiagnosisTask 查询单个诊断任务详情
//
// 查询单个诊断任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowDiagnosisTask(request *model.ShowDiagnosisTaskRequest) (*model.ShowDiagnosisTaskResponse, error) {
	requestDef := GenReqDefForShowDiagnosisTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDiagnosisTaskResponse), nil
	}
}

// ShowDiagnosisTaskInvoker 查询单个诊断任务详情
func (c *CocClient) ShowDiagnosisTaskInvoker(request *model.ShowDiagnosisTaskRequest) *ShowDiagnosisTaskInvoker {
	requestDef := GenReqDefForShowDiagnosisTask()
	return &ShowDiagnosisTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListEnterpriseProjectCollect 查询企业项目收藏
//
// 查询企业项目收藏。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListEnterpriseProjectCollect(request *model.ListEnterpriseProjectCollectRequest) (*model.ListEnterpriseProjectCollectResponse, error) {
	requestDef := GenReqDefForListEnterpriseProjectCollect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnterpriseProjectCollectResponse), nil
	}
}

// ListEnterpriseProjectCollectInvoker 查询企业项目收藏
func (c *CocClient) ListEnterpriseProjectCollectInvoker(request *model.ListEnterpriseProjectCollectRequest) *ListEnterpriseProjectCollectInvoker {
	requestDef := GenReqDefForListEnterpriseProjectCollect()
	return &ListEnterpriseProjectCollectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateEnterpriseProjectCollect 更新企业项目收藏
//
// 更新企业项目收藏。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateEnterpriseProjectCollect(request *model.UpdateEnterpriseProjectCollectRequest) (*model.UpdateEnterpriseProjectCollectResponse, error) {
	requestDef := GenReqDefForUpdateEnterpriseProjectCollect()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateEnterpriseProjectCollectResponse), nil
	}
}

// UpdateEnterpriseProjectCollectInvoker 更新企业项目收藏
func (c *CocClient) UpdateEnterpriseProjectCollectInvoker(request *model.UpdateEnterpriseProjectCollectRequest) *UpdateEnterpriseProjectCollectInvoker {
	requestDef := GenReqDefForUpdateEnterpriseProjectCollect()
	return &UpdateEnterpriseProjectCollectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListSubTickets 搜索变更工单子单
//
// 搜索变更工单子单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListSubTickets(request *model.ListSubTicketsRequest) (*model.ListSubTicketsResponse, error) {
	requestDef := GenReqDefForListSubTickets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubTicketsResponse), nil
	}
}

// ListSubTicketsInvoker 搜索变更工单子单
func (c *CocClient) ListSubTicketsInvoker(request *model.ListSubTicketsRequest) *ListSubTicketsInvoker {
	requestDef := GenReqDefForListSubTickets()
	return &ListSubTicketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTicketAction 工单操作
//
// 变更单审批、撤销以及问题单的所有操作均通过此接口完成。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ExecuteTicketAction(request *model.ExecuteTicketActionRequest) (*model.ExecuteTicketActionResponse, error) {
	requestDef := GenReqDefForExecuteTicketAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTicketActionResponse), nil
	}
}

// ExecuteTicketActionInvoker 工单操作
func (c *CocClient) ExecuteTicketActionInvoker(request *model.ExecuteTicketActionRequest) *ExecuteTicketActionInvoker {
	requestDef := GenReqDefForExecuteTicketAction()
	return &ExecuteTicketActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTicketOperationHistories 搜索工单历史
//
// List Histories
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListTicketOperationHistories(request *model.ListTicketOperationHistoriesRequest) (*model.ListTicketOperationHistoriesResponse, error) {
	requestDef := GenReqDefForListTicketOperationHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTicketOperationHistoriesResponse), nil
	}
}

// ListTicketOperationHistoriesInvoker 搜索工单历史
func (c *CocClient) ListTicketOperationHistoriesInvoker(request *model.ListTicketOperationHistoriesRequest) *ListTicketOperationHistoriesInvoker {
	requestDef := GenReqDefForListTicketOperationHistories()
	return &ListTicketOperationHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTickets 搜索工单
//
// List ticket
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListTickets(request *model.ListTicketsRequest) (*model.ListTicketsResponse, error) {
	requestDef := GenReqDefForListTickets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTicketsResponse), nil
	}
}

// ListTicketsInvoker 搜索工单
func (c *CocClient) ListTicketsInvoker(request *model.ListTicketsRequest) *ListTicketsInvoker {
	requestDef := GenReqDefForListTickets()
	return &ListTicketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTicketInfo 查询Ticket
//
// Get Ticket info by id
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ShowTicketInfo(request *model.ShowTicketInfoRequest) (*model.ShowTicketInfoResponse, error) {
	requestDef := GenReqDefForShowTicketInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTicketInfoResponse), nil
	}
}

// ShowTicketInfoInvoker 查询Ticket
func (c *CocClient) ShowTicketInfoInvoker(request *model.ShowTicketInfoRequest) *ShowTicketInfoInvoker {
	requestDef := GenReqDefForShowTicketInfo()
	return &ShowTicketInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTicketInfo 删除变更单
//
// 删除变更单，当变更单为撤销状态下，变更单可进行删除操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteTicketInfo(request *model.DeleteTicketInfoRequest) (*model.DeleteTicketInfoResponse, error) {
	requestDef := GenReqDefForDeleteTicketInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTicketInfoResponse), nil
	}
}

// DeleteTicketInfoInvoker 删除变更单
func (c *CocClient) DeleteTicketInfoInvoker(request *model.DeleteTicketInfoRequest) *DeleteTicketInfoInvoker {
	requestDef := GenReqDefForDeleteTicketInfo()
	return &DeleteTicketInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTicket 变更单状态修改
//
// 变更单状态修改，请求路径中的ticket_type为固定值change，且ticket_id传递变更单单号。此接口可操作变更开始、变更结束、变更取消和添加变更结果操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateTicket(request *model.UpdateTicketRequest) (*model.UpdateTicketResponse, error) {
	requestDef := GenReqDefForUpdateTicket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTicketResponse), nil
	}
}

// UpdateTicketInvoker 变更单状态修改
func (c *CocClient) UpdateTicketInvoker(request *model.UpdateTicketRequest) *UpdateTicketInvoker {
	requestDef := GenReqDefForUpdateTicket()
	return &UpdateTicketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateExternalCocAttachment 上传附件
//
// 上传附件，创建事件单的场景下，如需上传附件，需要先调用该接口将文件上传到obs。上传成功时，该接口将返回文档唯一id。支持文件类型：.jpg,.png,.docx,.txt,.pdf，且文本大小不超10M
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateExternalCocAttachment(request *model.CreateExternalCocAttachmentRequest) (*model.CreateExternalCocAttachmentResponse, error) {
	requestDef := GenReqDefForCreateExternalCocAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateExternalCocAttachmentResponse), nil
	}
}

// CreateExternalCocAttachmentInvoker 上传附件
func (c *CocClient) CreateExternalCocAttachmentInvoker(request *model.CreateExternalCocAttachmentRequest) *CreateExternalCocAttachmentInvoker {
	requestDef := GenReqDefForCreateExternalCocAttachment()
	return &CreateExternalCocAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListIncidentSimpleTickets 查询简易版事件单列表
//
// 该接口可分页查询到事件单列表信息，分页参数为limit与offset。该接口不支持对事件单进行除分页参数外的条件过滤，且返回的列表字段相对简单，只有事件单标题、事件单单号、描述信息、工单状态、事件级别、企业项目ID、事件单来源、创建人及责任人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListIncidentSimpleTickets(request *model.ListIncidentSimpleTicketsRequest) (*model.ListIncidentSimpleTicketsResponse, error) {
	requestDef := GenReqDefForListIncidentSimpleTickets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentSimpleTicketsResponse), nil
	}
}

// ListIncidentSimpleTicketsInvoker 查询简易版事件单列表
func (c *CocClient) ListIncidentSimpleTicketsInvoker(request *model.ListIncidentSimpleTicketsRequest) *ListIncidentSimpleTicketsInvoker {
	requestDef := GenReqDefForListIncidentSimpleTickets()
	return &ListIncidentSimpleTicketsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateAttachment 上传附件
//
// 上传附件，创建工单（事件单、变更单、问题单）的场景下，如需上传附件，需要先调用该接口将文件上传到obs。上传成功时，该接口将返回文档唯一id。支持文件类型：.jpg,.png,.docx,.txt,.pdf，且文本大小不超10M。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateAttachment(request *model.CreateAttachmentRequest) (*model.CreateAttachmentResponse, error) {
	requestDef := GenReqDefForCreateAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAttachmentResponse), nil
	}
}

// CreateAttachmentInvoker 上传附件
func (c *CocClient) CreateAttachmentInvoker(request *model.CreateAttachmentRequest) *CreateAttachmentInvoker {
	requestDef := GenReqDefForCreateAttachment()
	return &CreateAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTicket 新建工单
//
// 创建变更单或问题单的接口，通过路径参数ticket_type区分需要创建的工单类型。ticket_type为change表示要创建变更单，ticket_type为issues_mgmt为创建问题单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateTicket(request *model.CreateTicketRequest) (*model.CreateTicketResponse, error) {
	requestDef := GenReqDefForCreateTicket()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTicketResponse), nil
	}
}

// CreateTicketInvoker 新建工单
func (c *CocClient) CreateTicketInvoker(request *model.CreateTicketRequest) *CreateTicketInvoker {
	requestDef := GenReqDefForCreateTicket()
	return &CreateTicketInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadAttachment 下载附件
//
// 附件下载操作需基于已上传的附件资源。上传附件时，需调用/v1/{ticket_type}/attachments接口完成上传；成功上传后，可从接口响应中获取doc_id参数。下载附件时，凭借此doc_id再次调用/v1/{ticket_type}/attachments接口，即可获取已上传的对应附件资源，实现附件全生命周期管理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DownloadAttachment(request *model.DownloadAttachmentRequest) (*model.DownloadAttachmentResponse, error) {
	requestDef := GenReqDefForDownloadAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadAttachmentResponse), nil
	}
}

// DownloadAttachmentInvoker 下载附件
func (c *CocClient) DownloadAttachmentInvoker(request *model.DownloadAttachmentRequest) *DownloadAttachmentInvoker {
	requestDef := GenReqDefForDownloadAttachment()
	return &DownloadAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApplicationGroup 创建分组
//
// 创建分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateApplicationGroup(request *model.CreateApplicationGroupRequest) (*model.CreateApplicationGroupResponse, error) {
	requestDef := GenReqDefForCreateApplicationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApplicationGroupResponse), nil
	}
}

// CreateApplicationGroupInvoker 创建分组
func (c *CocClient) CreateApplicationGroupInvoker(request *model.CreateApplicationGroupRequest) *CreateApplicationGroupInvoker {
	requestDef := GenReqDefForCreateApplicationGroup()
	return &CreateApplicationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApplicationGroup 删除分组
//
// 删除分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteApplicationGroup(request *model.DeleteApplicationGroupRequest) (*model.DeleteApplicationGroupResponse, error) {
	requestDef := GenReqDefForDeleteApplicationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApplicationGroupResponse), nil
	}
}

// DeleteApplicationGroupInvoker 删除分组
func (c *CocClient) DeleteApplicationGroupInvoker(request *model.DeleteApplicationGroupRequest) *DeleteApplicationGroupInvoker {
	requestDef := GenReqDefForDeleteApplicationGroup()
	return &DeleteApplicationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApplicationGroups 查询分组
//
// 查询应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListApplicationGroups(request *model.ListApplicationGroupsRequest) (*model.ListApplicationGroupsResponse, error) {
	requestDef := GenReqDefForListApplicationGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplicationGroupsResponse), nil
	}
}

// ListApplicationGroupsInvoker 查询分组
func (c *CocClient) ListApplicationGroupsInvoker(request *model.ListApplicationGroupsRequest) *ListApplicationGroupsInvoker {
	requestDef := GenReqDefForListApplicationGroups()
	return &ListApplicationGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncGroupResource 分组智能同步资源
//
// 分组智能同步资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncGroupResource(request *model.SyncGroupResourceRequest) (*model.SyncGroupResourceResponse, error) {
	requestDef := GenReqDefForSyncGroupResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncGroupResourceResponse), nil
	}
}

// SyncGroupResourceInvoker 分组智能同步资源
func (c *CocClient) SyncGroupResourceInvoker(request *model.SyncGroupResourceRequest) *SyncGroupResourceInvoker {
	requestDef := GenReqDefForSyncGroupResource()
	return &SyncGroupResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApplicationGroup 修改分组
//
// 修改分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateApplicationGroup(request *model.UpdateApplicationGroupRequest) (*model.UpdateApplicationGroupResponse, error) {
	requestDef := GenReqDefForUpdateApplicationGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApplicationGroupResponse), nil
	}
}

// UpdateApplicationGroupInvoker 修改分组
func (c *CocClient) UpdateApplicationGroupInvoker(request *model.UpdateApplicationGroupRequest) *UpdateApplicationGroupInvoker {
	requestDef := GenReqDefForUpdateApplicationGroup()
	return &UpdateApplicationGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupAliResourceRelations 查询分组关联的阿里云资源
//
// 查询分组关联的阿里云资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListGroupAliResourceRelations(request *model.ListGroupAliResourceRelationsRequest) (*model.ListGroupAliResourceRelationsResponse, error) {
	requestDef := GenReqDefForListGroupAliResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupAliResourceRelationsResponse), nil
	}
}

// ListGroupAliResourceRelationsInvoker 查询分组关联的阿里云资源
func (c *CocClient) ListGroupAliResourceRelationsInvoker(request *model.ListGroupAliResourceRelationsRequest) *ListGroupAliResourceRelationsInvoker {
	requestDef := GenReqDefForListGroupAliResourceRelations()
	return &ListGroupAliResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupOtherResourceRelations 查询分组关联的离线资源
//
// 查询分组关联的离线资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListGroupOtherResourceRelations(request *model.ListGroupOtherResourceRelationsRequest) (*model.ListGroupOtherResourceRelationsResponse, error) {
	requestDef := GenReqDefForListGroupOtherResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupOtherResourceRelationsResponse), nil
	}
}

// ListGroupOtherResourceRelationsInvoker 查询分组关联的离线资源
func (c *CocClient) ListGroupOtherResourceRelationsInvoker(request *model.ListGroupOtherResourceRelationsRequest) *ListGroupOtherResourceRelationsInvoker {
	requestDef := GenReqDefForListGroupOtherResourceRelations()
	return &ListGroupOtherResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountGroupRmsResourceRelations 查询分组关联的资源总数
//
// 分组管理多个资源后，可查询分组关联的资源总数， 应用id和分组id不能共存，且必填其中一个。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountGroupRmsResourceRelations(request *model.CountGroupRmsResourceRelationsRequest) (*model.CountGroupRmsResourceRelationsResponse, error) {
	requestDef := GenReqDefForCountGroupRmsResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountGroupRmsResourceRelationsResponse), nil
	}
}

// CountGroupRmsResourceRelationsInvoker 查询分组关联的资源总数
func (c *CocClient) CountGroupRmsResourceRelationsInvoker(request *model.CountGroupRmsResourceRelationsRequest) *CountGroupRmsResourceRelationsInvoker {
	requestDef := GenReqDefForCountGroupRmsResourceRelations()
	return &CountGroupRmsResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGroupRmsResourceRelation 创建分组资源关联
//
// 创建分组资源关联。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateGroupRmsResourceRelation(request *model.CreateGroupRmsResourceRelationRequest) (*model.CreateGroupRmsResourceRelationResponse, error) {
	requestDef := GenReqDefForCreateGroupRmsResourceRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGroupRmsResourceRelationResponse), nil
	}
}

// CreateGroupRmsResourceRelationInvoker 创建分组资源关联
func (c *CocClient) CreateGroupRmsResourceRelationInvoker(request *model.CreateGroupRmsResourceRelationRequest) *CreateGroupRmsResourceRelationInvoker {
	requestDef := GenReqDefForCreateGroupRmsResourceRelation()
	return &CreateGroupRmsResourceRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGroupRmsResourceRelation 解绑分组资源关联
//
// 解绑分组资源关联。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteGroupRmsResourceRelation(request *model.DeleteGroupRmsResourceRelationRequest) (*model.DeleteGroupRmsResourceRelationResponse, error) {
	requestDef := GenReqDefForDeleteGroupRmsResourceRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGroupRmsResourceRelationResponse), nil
	}
}

// DeleteGroupRmsResourceRelationInvoker 解绑分组资源关联
func (c *CocClient) DeleteGroupRmsResourceRelationInvoker(request *model.DeleteGroupRmsResourceRelationRequest) *DeleteGroupRmsResourceRelationInvoker {
	requestDef := GenReqDefForDeleteGroupRmsResourceRelation()
	return &DeleteGroupRmsResourceRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCapacityOrder 查询应用、组件、分组容量数据的排名
//
// 查询应用、组件、分组容量数据的排名，其中应用、组件和分组ID，有且仅有1个有值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListCapacityOrder(request *model.ListCapacityOrderRequest) (*model.ListCapacityOrderResponse, error) {
	requestDef := GenReqDefForListCapacityOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCapacityOrderResponse), nil
	}
}

// ListCapacityOrderInvoker 查询应用、组件、分组容量数据的排名
func (c *CocClient) ListCapacityOrderInvoker(request *model.ListCapacityOrderRequest) *ListCapacityOrderInvoker {
	requestDef := GenReqDefForListCapacityOrder()
	return &ListCapacityOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCapacityView 查询应用的容量数据
//
// 云运维中心支持查看应用、子应用、组件或分组下已关联的资源容量详情，按照资源类型展示资源核心数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListCapacityView(request *model.ListCapacityViewRequest) (*model.ListCapacityViewResponse, error) {
	requestDef := GenReqDefForListCapacityView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCapacityViewResponse), nil
	}
}

// ListCapacityViewInvoker 查询应用的容量数据
func (c *CocClient) ListCapacityViewInvoker(request *model.ListCapacityViewRequest) *ListCapacityViewInvoker {
	requestDef := GenReqDefForListCapacityView()
	return &ListCapacityViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGroupRmsResourceRelations 查询分组关联的华为云资源
//
// 查询分组关联的华为云资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListGroupRmsResourceRelations(request *model.ListGroupRmsResourceRelationsRequest) (*model.ListGroupRmsResourceRelationsResponse, error) {
	requestDef := GenReqDefForListGroupRmsResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGroupRmsResourceRelationsResponse), nil
	}
}

// ListGroupRmsResourceRelationsInvoker 查询分组关联的华为云资源
func (c *CocClient) ListGroupRmsResourceRelationsInvoker(request *model.ListGroupRmsResourceRelationsRequest) *ListGroupRmsResourceRelationsInvoker {
	requestDef := GenReqDefForListGroupRmsResourceRelations()
	return &ListGroupRmsResourceRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGroupRmsResourceRelation 转移资源
//
// 转移资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateGroupRmsResourceRelation(request *model.UpdateGroupRmsResourceRelationRequest) (*model.UpdateGroupRmsResourceRelationResponse, error) {
	requestDef := GenReqDefForUpdateGroupRmsResourceRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGroupRmsResourceRelationResponse), nil
	}
}

// UpdateGroupRmsResourceRelationInvoker 转移资源
func (c *CocClient) UpdateGroupRmsResourceRelationInvoker(request *model.UpdateGroupRmsResourceRelationRequest) *UpdateGroupRmsResourceRelationInvoker {
	requestDef := GenReqDefForUpdateGroupRmsResourceRelation()
	return &UpdateGroupRmsResourceRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CountMultiCloudResources 查询用户在云厂商的资源总数
//
// 查询用户在云厂商（阿里云、AWS、Azure和HCS）的资源总数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountMultiCloudResources(request *model.CountMultiCloudResourcesRequest) (*model.CountMultiCloudResourcesResponse, error) {
	requestDef := GenReqDefForCountMultiCloudResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountMultiCloudResourcesResponse), nil
	}
}

// CountMultiCloudResourcesInvoker 查询用户在云厂商的资源总数
func (c *CocClient) CountMultiCloudResourcesInvoker(request *model.CountMultiCloudResourcesRequest) *CountMultiCloudResourcesInvoker {
	requestDef := GenReqDefForCountMultiCloudResources()
	return &CountMultiCloudResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncMultiCloudResource 手动从云厂商同步用户资源
//
// 手动从云厂商同步用户资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncMultiCloudResource(request *model.SyncMultiCloudResourceRequest) (*model.SyncMultiCloudResourceResponse, error) {
	requestDef := GenReqDefForSyncMultiCloudResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncMultiCloudResourceResponse), nil
	}
}

// SyncMultiCloudResourceInvoker 手动从云厂商同步用户资源
func (c *CocClient) SyncMultiCloudResourceInvoker(request *model.SyncMultiCloudResourceRequest) *SyncMultiCloudResourceInvoker {
	requestDef := GenReqDefForSyncMultiCloudResource()
	return &SyncMultiCloudResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CountResources 查询用户资源总数
//
// 查询用户资源总数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountResources(request *model.CountResourcesRequest) (*model.CountResourcesResponse, error) {
	requestDef := GenReqDefForCountResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourcesResponse), nil
	}
}

// CountResourcesInvoker 查询用户资源总数
func (c *CocClient) CountResourcesInvoker(request *model.CountResourcesRequest) *CountResourcesInvoker {
	requestDef := GenReqDefForCountResources()
	return &CountResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResources 查询用户所有资源
//
// 查询租户所有资源：
//  - 查询租户所有资源等相关信息，便于租户详细了解资源总体情况。
//  - 请求参数provider（云服务名称），type（云资源类型），limit（查询条数）必填，单次最大查询条数：500。
//  - 返回信息包括：资源ID，资源名称，云服务名称，资源类型，项目ID，租户ID，区域ID，企业项目ID，资源标签，资源详细属性，资源ingest属性，uniagentID，uniagent状态，是否托管，是否可运维。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// ListResourcesInvoker 查询用户所有资源
func (c *CocClient) ListResourcesInvoker(request *model.ListResourcesRequest) *ListResourcesInvoker {
	requestDef := GenReqDefForListResources()
	return &ListResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// SyncResourceAgent 同步华为云资源Agent信息
//
// 同步华为云资源Agent信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncResourceAgent(request *model.SyncResourceAgentRequest) (*model.SyncResourceAgentResponse, error) {
	requestDef := GenReqDefForSyncResourceAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncResourceAgentResponse), nil
	}
}

// SyncResourceAgentInvoker 同步华为云资源Agent信息
func (c *CocClient) SyncResourceAgentInvoker(request *model.SyncResourceAgentRequest) *SyncResourceAgentInvoker {
	requestDef := GenReqDefForSyncResourceAgent()
	return &SyncResourceAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateResourceTags 添加资源标签
//
// 添加资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateResourceTags(request *model.CreateResourceTagsRequest) (*model.CreateResourceTagsResponse, error) {
	requestDef := GenReqDefForCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagsResponse), nil
	}
}

// CreateResourceTagsInvoker 添加资源标签
func (c *CocClient) CreateResourceTagsInvoker(request *model.CreateResourceTagsRequest) *CreateResourceTagsInvoker {
	requestDef := GenReqDefForCreateResourceTags()
	return &CreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询资源标签
//
// 查询资源标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询资源标签
func (c *CocClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountResourcesOfResourceView 查询CMDB跨账号资源视图聚合的资源总数
//
// 查询CMDB跨账号资源视图聚合的资源总数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CountResourcesOfResourceView(request *model.CountResourcesOfResourceViewRequest) (*model.CountResourcesOfResourceViewResponse, error) {
	requestDef := GenReqDefForCountResourcesOfResourceView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountResourcesOfResourceViewResponse), nil
	}
}

// CountResourcesOfResourceViewInvoker 查询CMDB跨账号资源视图聚合的资源总数
func (c *CocClient) CountResourcesOfResourceViewInvoker(request *model.CountResourcesOfResourceViewRequest) *CountResourcesOfResourceViewInvoker {
	requestDef := GenReqDefForCountResourcesOfResourceView()
	return &CountResourcesOfResourceViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceViews 创建CMDB跨账号资源视图
//
// 创建CMDB跨账号资源视图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateResourceViews(request *model.CreateResourceViewsRequest) (*model.CreateResourceViewsResponse, error) {
	requestDef := GenReqDefForCreateResourceViews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceViewsResponse), nil
	}
}

// CreateResourceViewsInvoker 创建CMDB跨账号资源视图
func (c *CocClient) CreateResourceViewsInvoker(request *model.CreateResourceViewsRequest) *CreateResourceViewsInvoker {
	requestDef := GenReqDefForCreateResourceViews()
	return &CreateResourceViewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceViews 删除CMDB的跨账号资源管理视图
//
// 删除CMDB的跨账号资源管理视图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteResourceViews(request *model.DeleteResourceViewsRequest) (*model.DeleteResourceViewsResponse, error) {
	requestDef := GenReqDefForDeleteResourceViews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceViewsResponse), nil
	}
}

// DeleteResourceViewsInvoker 删除CMDB的跨账号资源管理视图
func (c *CocClient) DeleteResourceViewsInvoker(request *model.DeleteResourceViewsRequest) *DeleteResourceViewsInvoker {
	requestDef := GenReqDefForDeleteResourceViews()
	return &DeleteResourceViewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceViews 查询CMDB跨账号资源视图
//
// 查询CMDB跨账号资源视图。视图是一组筛选器，用户可以自由配置筛选范围，用于在跨账号场景下访问华为云中的资源信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResourceViews(request *model.ListResourceViewsRequest) (*model.ListResourceViewsResponse, error) {
	requestDef := GenReqDefForListResourceViews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceViewsResponse), nil
	}
}

// ListResourceViewsInvoker 查询CMDB跨账号资源视图
func (c *CocClient) ListResourceViewsInvoker(request *model.ListResourceViewsRequest) *ListResourceViewsInvoker {
	requestDef := GenReqDefForListResourceViews()
	return &ListResourceViewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourcesOfResourceView 查询CMDB跨账号资源视图聚合的资源
//
// 查询CMDB跨账号资源视图聚合的资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListResourcesOfResourceView(request *model.ListResourcesOfResourceViewRequest) (*model.ListResourcesOfResourceViewResponse, error) {
	requestDef := GenReqDefForListResourcesOfResourceView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesOfResourceViewResponse), nil
	}
}

// ListResourcesOfResourceViewInvoker 查询CMDB跨账号资源视图聚合的资源
func (c *CocClient) ListResourcesOfResourceViewInvoker(request *model.ListResourcesOfResourceViewRequest) *ListResourcesOfResourceViewInvoker {
	requestDef := GenReqDefForListResourcesOfResourceView()
	return &ListResourcesOfResourceViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncResourcesOfResourceView 同步CMDB跨账号资源视图
//
// 同步CMDB跨账号资源视图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) SyncResourcesOfResourceView(request *model.SyncResourcesOfResourceViewRequest) (*model.SyncResourcesOfResourceViewResponse, error) {
	requestDef := GenReqDefForSyncResourcesOfResourceView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncResourcesOfResourceViewResponse), nil
	}
}

// SyncResourcesOfResourceViewInvoker 同步CMDB跨账号资源视图
func (c *CocClient) SyncResourcesOfResourceViewInvoker(request *model.SyncResourcesOfResourceViewRequest) *SyncResourcesOfResourceViewInvoker {
	requestDef := GenReqDefForSyncResourcesOfResourceView()
	return &SyncResourcesOfResourceViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceViews 更新CMDB跨账号资源视图
//
// 更新CMDB跨账号资源视图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateResourceViews(request *model.UpdateResourceViewsRequest) (*model.UpdateResourceViewsResponse, error) {
	requestDef := GenReqDefForUpdateResourceViews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceViewsResponse), nil
	}
}

// UpdateResourceViewsInvoker 更新CMDB跨账号资源视图
func (c *CocClient) UpdateResourceViewsInvoker(request *model.UpdateResourceViewsRequest) *UpdateResourceViewsInvoker {
	requestDef := GenReqDefForUpdateResourceViews()
	return &UpdateResourceViewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// CreateVendorAccount 创建云厂商账号
//
// 创建云厂商（阿里云、AWS、Azure和HCS等）账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) CreateVendorAccount(request *model.CreateVendorAccountRequest) (*model.CreateVendorAccountResponse, error) {
	requestDef := GenReqDefForCreateVendorAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateVendorAccountResponse), nil
	}
}

// CreateVendorAccountInvoker 创建云厂商账号
func (c *CocClient) CreateVendorAccountInvoker(request *model.CreateVendorAccountRequest) *CreateVendorAccountInvoker {
	requestDef := GenReqDefForCreateVendorAccount()
	return &CreateVendorAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteVendorAccount 删除云厂商账号
//
// 增加云广商账号，不需要后，可删除云厂商账号。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) DeleteVendorAccount(request *model.DeleteVendorAccountRequest) (*model.DeleteVendorAccountResponse, error) {
	requestDef := GenReqDefForDeleteVendorAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteVendorAccountResponse), nil
	}
}

// DeleteVendorAccountInvoker 删除云厂商账号
func (c *CocClient) DeleteVendorAccountInvoker(request *model.DeleteVendorAccountRequest) *DeleteVendorAccountInvoker {
	requestDef := GenReqDefForDeleteVendorAccount()
	return &DeleteVendorAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListVendorAccount 查询云厂商账号
//
// 查询所有云厂商（阿里云、AWS、Azure和HCS等）账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) ListVendorAccount(request *model.ListVendorAccountRequest) (*model.ListVendorAccountResponse, error) {
	requestDef := GenReqDefForListVendorAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListVendorAccountResponse), nil
	}
}

// ListVendorAccountInvoker 查询云厂商账号
func (c *CocClient) ListVendorAccountInvoker(request *model.ListVendorAccountRequest) *ListVendorAccountInvoker {
	requestDef := GenReqDefForListVendorAccount()
	return &ListVendorAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateVendorAccount 修改云厂商账号
//
// 修改所有云厂商（阿里云、AWS、Azure和HCS等）账号信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CocClient) UpdateVendorAccount(request *model.UpdateVendorAccountRequest) (*model.UpdateVendorAccountResponse, error) {
	requestDef := GenReqDefForUpdateVendorAccount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateVendorAccountResponse), nil
	}
}

// UpdateVendorAccountInvoker 修改云厂商账号
func (c *CocClient) UpdateVendorAccountInvoker(request *model.UpdateVendorAccountRequest) *UpdateVendorAccountInvoker {
	requestDef := GenReqDefForUpdateVendorAccount()
	return &UpdateVendorAccountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
