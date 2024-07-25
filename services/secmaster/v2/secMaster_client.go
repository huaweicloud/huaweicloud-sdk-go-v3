package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/secmaster/v2/model"
)

type SecMasterClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewSecMasterClient(hcClient *httpclient.HcHttpClient) *SecMasterClient {
	return &SecMasterClient{HcClient: hcClient}
}

func SecMasterClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchSearchMetricHits 批量查询指标结果
//
// 批量查询指标结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) BatchSearchMetricHits(request *model.BatchSearchMetricHitsRequest) (*model.BatchSearchMetricHitsResponse, error) {
	requestDef := GenReqDefForBatchSearchMetricHits()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSearchMetricHitsResponse), nil
	}
}

// BatchSearchMetricHitsInvoker 批量查询指标结果
func (c *SecMasterClient) BatchSearchMetricHitsInvoker(request *model.BatchSearchMetricHitsRequest) *BatchSearchMetricHitsInvoker {
	requestDef := GenReqDefForBatchSearchMetricHits()
	return &BatchSearchMetricHitsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAlert 更新告警
//
// 编辑告警，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeAlert(request *model.ChangeAlertRequest) (*model.ChangeAlertResponse, error) {
	requestDef := GenReqDefForChangeAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAlertResponse), nil
	}
}

// ChangeAlertInvoker 更新告警
func (c *SecMasterClient) ChangeAlertInvoker(request *model.ChangeAlertRequest) *ChangeAlertInvoker {
	requestDef := GenReqDefForChangeAlert()
	return &ChangeAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIncident 更新事件
//
// 编辑事件，根据实际修改的属性更新，未修改的列不更新
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangeIncident(request *model.ChangeIncidentRequest) (*model.ChangeIncidentResponse, error) {
	requestDef := GenReqDefForChangeIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIncidentResponse), nil
	}
}

// ChangeIncidentInvoker 更新事件
func (c *SecMasterClient) ChangeIncidentInvoker(request *model.ChangeIncidentRequest) *ChangeIncidentInvoker {
	requestDef := GenReqDefForChangeIncident()
	return &ChangeIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangePlaybookInstance 操作剧本实例
//
// 操作剧本实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ChangePlaybookInstance(request *model.ChangePlaybookInstanceRequest) (*model.ChangePlaybookInstanceResponse, error) {
	requestDef := GenReqDefForChangePlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangePlaybookInstanceResponse), nil
	}
}

// ChangePlaybookInstanceInvoker 操作剧本实例
func (c *SecMasterClient) ChangePlaybookInstanceInvoker(request *model.ChangePlaybookInstanceRequest) *ChangePlaybookInstanceInvoker {
	requestDef := GenReqDefForChangePlaybookInstance()
	return &ChangePlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CopyPlaybookVersion 克隆剧本及版本
//
// 克隆剧本及版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CopyPlaybookVersion(request *model.CopyPlaybookVersionRequest) (*model.CopyPlaybookVersionResponse, error) {
	requestDef := GenReqDefForCopyPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CopyPlaybookVersionResponse), nil
	}
}

// CopyPlaybookVersionInvoker 克隆剧本及版本
func (c *SecMasterClient) CopyPlaybookVersionInvoker(request *model.CopyPlaybookVersionRequest) *CopyPlaybookVersionInvoker {
	requestDef := GenReqDefForCopyPlaybookVersion()
	return &CopyPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlert 创建告警
//
// 创建告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlert(request *model.CreateAlertRequest) (*model.CreateAlertResponse, error) {
	requestDef := GenReqDefForCreateAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertResponse), nil
	}
}

// CreateAlertInvoker 创建告警
func (c *SecMasterClient) CreateAlertInvoker(request *model.CreateAlertRequest) *CreateAlertInvoker {
	requestDef := GenReqDefForCreateAlert()
	return &CreateAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRule 创建告警规则
//
// # Create alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlertRule(request *model.CreateAlertRuleRequest) (*model.CreateAlertRuleResponse, error) {
	requestDef := GenReqDefForCreateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleResponse), nil
	}
}

// CreateAlertRuleInvoker 创建告警规则
func (c *SecMasterClient) CreateAlertRuleInvoker(request *model.CreateAlertRuleRequest) *CreateAlertRuleInvoker {
	requestDef := GenReqDefForCreateAlertRule()
	return &CreateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlertRuleSimulation 模拟告警规则
//
// # Simulate alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateAlertRuleSimulation(request *model.CreateAlertRuleSimulationRequest) (*model.CreateAlertRuleSimulationResponse, error) {
	requestDef := GenReqDefForCreateAlertRuleSimulation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlertRuleSimulationResponse), nil
	}
}

// CreateAlertRuleSimulationInvoker 模拟告警规则
func (c *SecMasterClient) CreateAlertRuleSimulationInvoker(request *model.CreateAlertRuleSimulationRequest) *CreateAlertRuleSimulationInvoker {
	requestDef := GenReqDefForCreateAlertRuleSimulation()
	return &CreateAlertRuleSimulationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBatchOrderAlerts 告警转事件
//
// 告警转事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateBatchOrderAlerts(request *model.CreateBatchOrderAlertsRequest) (*model.CreateBatchOrderAlertsResponse, error) {
	requestDef := GenReqDefForCreateBatchOrderAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBatchOrderAlertsResponse), nil
	}
}

// CreateBatchOrderAlertsInvoker 告警转事件
func (c *SecMasterClient) CreateBatchOrderAlertsInvoker(request *model.CreateBatchOrderAlertsRequest) *CreateBatchOrderAlertsInvoker {
	requestDef := GenReqDefForCreateBatchOrderAlerts()
	return &CreateBatchOrderAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataobjectRelations 关联Dataobject
//
// 关联Dataobject
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataobjectRelations(request *model.CreateDataobjectRelationsRequest) (*model.CreateDataobjectRelationsResponse, error) {
	requestDef := GenReqDefForCreateDataobjectRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataobjectRelationsResponse), nil
	}
}

// CreateDataobjectRelationsInvoker 关联Dataobject
func (c *SecMasterClient) CreateDataobjectRelationsInvoker(request *model.CreateDataobjectRelationsRequest) *CreateDataobjectRelationsInvoker {
	requestDef := GenReqDefForCreateDataobjectRelations()
	return &CreateDataobjectRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDataspace 创建数据空间
//
// 创建数据空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateDataspace(request *model.CreateDataspaceRequest) (*model.CreateDataspaceResponse, error) {
	requestDef := GenReqDefForCreateDataspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDataspaceResponse), nil
	}
}

// CreateDataspaceInvoker 创建数据空间
func (c *SecMasterClient) CreateDataspaceInvoker(request *model.CreateDataspaceRequest) *CreateDataspaceInvoker {
	requestDef := GenReqDefForCreateDataspace()
	return &CreateDataspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIncident 创建事件
//
// 创建事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateIncident(request *model.CreateIncidentRequest) (*model.CreateIncidentResponse, error) {
	requestDef := GenReqDefForCreateIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIncidentResponse), nil
	}
}

// CreateIncidentInvoker 创建事件
func (c *SecMasterClient) CreateIncidentInvoker(request *model.CreateIncidentRequest) *CreateIncidentInvoker {
	requestDef := GenReqDefForCreateIncident()
	return &CreateIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateIndicator 创建威胁情报
//
// 创建威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateIndicator(request *model.CreateIndicatorRequest) (*model.CreateIndicatorResponse, error) {
	requestDef := GenReqDefForCreateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateIndicatorResponse), nil
	}
}

// CreateIndicatorInvoker 创建威胁情报
func (c *SecMasterClient) CreateIndicatorInvoker(request *model.CreateIndicatorRequest) *CreateIndicatorInvoker {
	requestDef := GenReqDefForCreateIndicator()
	return &CreateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePipe 创建数据管道
//
// 创建数据管道
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePipe(request *model.CreatePipeRequest) (*model.CreatePipeResponse, error) {
	requestDef := GenReqDefForCreatePipe()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePipeResponse), nil
	}
}

// CreatePipeInvoker 创建数据管道
func (c *SecMasterClient) CreatePipeInvoker(request *model.CreatePipeRequest) *CreatePipeInvoker {
	requestDef := GenReqDefForCreatePipe()
	return &CreatePipeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybook 创建剧本
//
// 创建剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybook(request *model.CreatePlaybookRequest) (*model.CreatePlaybookResponse, error) {
	requestDef := GenReqDefForCreatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookResponse), nil
	}
}

// CreatePlaybookInvoker 创建剧本
func (c *SecMasterClient) CreatePlaybookInvoker(request *model.CreatePlaybookRequest) *CreatePlaybookInvoker {
	requestDef := GenReqDefForCreatePlaybook()
	return &CreatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookAction 创建剧本动作
//
// 创建剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookAction(request *model.CreatePlaybookActionRequest) (*model.CreatePlaybookActionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookActionResponse), nil
	}
}

// CreatePlaybookActionInvoker 创建剧本动作
func (c *SecMasterClient) CreatePlaybookActionInvoker(request *model.CreatePlaybookActionRequest) *CreatePlaybookActionInvoker {
	requestDef := GenReqDefForCreatePlaybookAction()
	return &CreatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookApprove 审核剧本
//
// 审核剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookApprove(request *model.CreatePlaybookApproveRequest) (*model.CreatePlaybookApproveResponse, error) {
	requestDef := GenReqDefForCreatePlaybookApprove()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookApproveResponse), nil
	}
}

// CreatePlaybookApproveInvoker 审核剧本
func (c *SecMasterClient) CreatePlaybookApproveInvoker(request *model.CreatePlaybookApproveRequest) *CreatePlaybookApproveInvoker {
	requestDef := GenReqDefForCreatePlaybookApprove()
	return &CreatePlaybookApproveInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookRule 创建剧本规则
//
// 创建剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookRule(request *model.CreatePlaybookRuleRequest) (*model.CreatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForCreatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookRuleResponse), nil
	}
}

// CreatePlaybookRuleInvoker 创建剧本规则
func (c *SecMasterClient) CreatePlaybookRuleInvoker(request *model.CreatePlaybookRuleRequest) *CreatePlaybookRuleInvoker {
	requestDef := GenReqDefForCreatePlaybookRule()
	return &CreatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlaybookVersion 创建剧本版本
//
// 创建剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePlaybookVersion(request *model.CreatePlaybookVersionRequest) (*model.CreatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForCreatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePlaybookVersionResponse), nil
	}
}

// CreatePlaybookVersionInvoker 创建剧本版本
func (c *SecMasterClient) CreatePlaybookVersionInvoker(request *model.CreatePlaybookVersionRequest) *CreatePlaybookVersionInvoker {
	requestDef := GenReqDefForCreatePlaybookVersion()
	return &CreatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePostPaidOrder 安全云脑按需订购
//
// 开通安全云脑按需服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreatePostPaidOrder(request *model.CreatePostPaidOrderRequest) (*model.CreatePostPaidOrderResponse, error) {
	requestDef := GenReqDefForCreatePostPaidOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePostPaidOrderResponse), nil
	}
}

// CreatePostPaidOrderInvoker 安全云脑按需订购
func (c *SecMasterClient) CreatePostPaidOrderInvoker(request *model.CreatePostPaidOrderRequest) *CreatePostPaidOrderInvoker {
	requestDef := GenReqDefForCreatePostPaidOrder()
	return &CreatePostPaidOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkspace 新建工作空间
//
// 在使用安全云脑的基线检查、告警管理、安全分析、安全编排等功能前，需要创建工作空间，它可以将资源划分为各个不同的工作场景，避免资源冗余查找不便，影响日常使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) CreateWorkspace(request *model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkspaceResponse), nil
	}
}

// CreateWorkspaceInvoker 新建工作空间
func (c *SecMasterClient) CreateWorkspaceInvoker(request *model.CreateWorkspaceRequest) *CreateWorkspaceInvoker {
	requestDef := GenReqDefForCreateWorkspace()
	return &CreateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlert 删除告警
//
// 删除告警
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlert(request *model.DeleteAlertRequest) (*model.DeleteAlertResponse, error) {
	requestDef := GenReqDefForDeleteAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertResponse), nil
	}
}

// DeleteAlertInvoker 删除告警
func (c *SecMasterClient) DeleteAlertInvoker(request *model.DeleteAlertRequest) *DeleteAlertInvoker {
	requestDef := GenReqDefForDeleteAlert()
	return &DeleteAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlertRule 删除告警规则
//
// # Delete alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteAlertRule(request *model.DeleteAlertRuleRequest) (*model.DeleteAlertRuleResponse, error) {
	requestDef := GenReqDefForDeleteAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlertRuleResponse), nil
	}
}

// DeleteAlertRuleInvoker 删除告警规则
func (c *SecMasterClient) DeleteAlertRuleInvoker(request *model.DeleteAlertRuleRequest) *DeleteAlertRuleInvoker {
	requestDef := GenReqDefForDeleteAlertRule()
	return &DeleteAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataobjectRelations 取消关联Dataobject
//
// 取消关联Dataobject
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteDataobjectRelations(request *model.DeleteDataobjectRelationsRequest) (*model.DeleteDataobjectRelationsResponse, error) {
	requestDef := GenReqDefForDeleteDataobjectRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataobjectRelationsResponse), nil
	}
}

// DeleteDataobjectRelationsInvoker 取消关联Dataobject
func (c *SecMasterClient) DeleteDataobjectRelationsInvoker(request *model.DeleteDataobjectRelationsRequest) *DeleteDataobjectRelationsInvoker {
	requestDef := GenReqDefForDeleteDataobjectRelations()
	return &DeleteDataobjectRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIncident 删除事件
//
// 删除事件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteIncident(request *model.DeleteIncidentRequest) (*model.DeleteIncidentResponse, error) {
	requestDef := GenReqDefForDeleteIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIncidentResponse), nil
	}
}

// DeleteIncidentInvoker 删除事件
func (c *SecMasterClient) DeleteIncidentInvoker(request *model.DeleteIncidentRequest) *DeleteIncidentInvoker {
	requestDef := GenReqDefForDeleteIncident()
	return &DeleteIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteIndicator 删除威胁情报
//
// 删除威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeleteIndicator(request *model.DeleteIndicatorRequest) (*model.DeleteIndicatorResponse, error) {
	requestDef := GenReqDefForDeleteIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteIndicatorResponse), nil
	}
}

// DeleteIndicatorInvoker 删除威胁情报
func (c *SecMasterClient) DeleteIndicatorInvoker(request *model.DeleteIndicatorRequest) *DeleteIndicatorInvoker {
	requestDef := GenReqDefForDeleteIndicator()
	return &DeleteIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybook 删除剧本
//
// 删除剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybook(request *model.DeletePlaybookRequest) (*model.DeletePlaybookResponse, error) {
	requestDef := GenReqDefForDeletePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookResponse), nil
	}
}

// DeletePlaybookInvoker 删除剧本
func (c *SecMasterClient) DeletePlaybookInvoker(request *model.DeletePlaybookRequest) *DeletePlaybookInvoker {
	requestDef := GenReqDefForDeletePlaybook()
	return &DeletePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookAction 删除剧本动作
//
// 删除剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookAction(request *model.DeletePlaybookActionRequest) (*model.DeletePlaybookActionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookActionResponse), nil
	}
}

// DeletePlaybookActionInvoker 删除剧本动作
func (c *SecMasterClient) DeletePlaybookActionInvoker(request *model.DeletePlaybookActionRequest) *DeletePlaybookActionInvoker {
	requestDef := GenReqDefForDeletePlaybookAction()
	return &DeletePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookRule 删除剧本规则
//
// 删除剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookRule(request *model.DeletePlaybookRuleRequest) (*model.DeletePlaybookRuleResponse, error) {
	requestDef := GenReqDefForDeletePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookRuleResponse), nil
	}
}

// DeletePlaybookRuleInvoker 删除剧本规则
func (c *SecMasterClient) DeletePlaybookRuleInvoker(request *model.DeletePlaybookRuleRequest) *DeletePlaybookRuleInvoker {
	requestDef := GenReqDefForDeletePlaybookRule()
	return &DeletePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlaybookVersion 删除剧本版本
//
// 删除剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DeletePlaybookVersion(request *model.DeletePlaybookVersionRequest) (*model.DeletePlaybookVersionResponse, error) {
	requestDef := GenReqDefForDeletePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePlaybookVersionResponse), nil
	}
}

// DeletePlaybookVersionInvoker 删除剧本版本
func (c *SecMasterClient) DeletePlaybookVersionInvoker(request *model.DeletePlaybookVersionRequest) *DeletePlaybookVersionInvoker {
	requestDef := GenReqDefForDeletePlaybookVersion()
	return &DeletePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableAlertRule 停用告警规则
//
// # Disable alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) DisableAlertRule(request *model.DisableAlertRuleRequest) (*model.DisableAlertRuleResponse, error) {
	requestDef := GenReqDefForDisableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableAlertRuleResponse), nil
	}
}

// DisableAlertRuleInvoker 停用告警规则
func (c *SecMasterClient) DisableAlertRuleInvoker(request *model.DisableAlertRuleRequest) *DisableAlertRuleInvoker {
	requestDef := GenReqDefForDisableAlertRule()
	return &DisableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableAlertRule 启用告警规则
//
// # Enable alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) EnableAlertRule(request *model.EnableAlertRuleRequest) (*model.EnableAlertRuleResponse, error) {
	requestDef := GenReqDefForEnableAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableAlertRuleResponse), nil
	}
}

// EnableAlertRuleInvoker 启用告警规则
func (c *SecMasterClient) EnableAlertRuleInvoker(request *model.EnableAlertRuleRequest) *EnableAlertRuleInvoker {
	requestDef := GenReqDefForEnableAlertRule()
	return &EnableAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleMetrics 告警规则总览
//
// # List alert rule metrics
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleMetrics(request *model.ListAlertRuleMetricsRequest) (*model.ListAlertRuleMetricsResponse, error) {
	requestDef := GenReqDefForListAlertRuleMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleMetricsResponse), nil
	}
}

// ListAlertRuleMetricsInvoker 告警规则总览
func (c *SecMasterClient) ListAlertRuleMetricsInvoker(request *model.ListAlertRuleMetricsRequest) *ListAlertRuleMetricsInvoker {
	requestDef := GenReqDefForListAlertRuleMetrics()
	return &ListAlertRuleMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRuleTemplates 列出告警规则模板
//
// # List alert rule templates
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRuleTemplates(request *model.ListAlertRuleTemplatesRequest) (*model.ListAlertRuleTemplatesResponse, error) {
	requestDef := GenReqDefForListAlertRuleTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRuleTemplatesResponse), nil
	}
}

// ListAlertRuleTemplatesInvoker 列出告警规则模板
func (c *SecMasterClient) ListAlertRuleTemplatesInvoker(request *model.ListAlertRuleTemplatesRequest) *ListAlertRuleTemplatesInvoker {
	requestDef := GenReqDefForListAlertRuleTemplates()
	return &ListAlertRuleTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlertRules 列出告警规则
//
// # List alert rules
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlertRules(request *model.ListAlertRulesRequest) (*model.ListAlertRulesResponse, error) {
	requestDef := GenReqDefForListAlertRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertRulesResponse), nil
	}
}

// ListAlertRulesInvoker 列出告警规则
func (c *SecMasterClient) ListAlertRulesInvoker(request *model.ListAlertRulesRequest) *ListAlertRulesInvoker {
	requestDef := GenReqDefForListAlertRules()
	return &ListAlertRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlerts 搜索告警列表
//
// 搜索告警列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListAlerts(request *model.ListAlertsRequest) (*model.ListAlertsResponse, error) {
	requestDef := GenReqDefForListAlerts()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlertsResponse), nil
	}
}

// ListAlertsInvoker 搜索告警列表
func (c *SecMasterClient) ListAlertsInvoker(request *model.ListAlertsRequest) *ListAlertsInvoker {
	requestDef := GenReqDefForListAlerts()
	return &ListAlertsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataclass 查询数据类列表
//
// 查询数据类列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataclass(request *model.ListDataclassRequest) (*model.ListDataclassResponse, error) {
	requestDef := GenReqDefForListDataclass()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataclassResponse), nil
	}
}

// ListDataclassInvoker 查询数据类列表
func (c *SecMasterClient) ListDataclassInvoker(request *model.ListDataclassRequest) *ListDataclassInvoker {
	requestDef := GenReqDefForListDataclass()
	return &ListDataclassInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataclassFields 查询字段列表
//
// 查询字段列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataclassFields(request *model.ListDataclassFieldsRequest) (*model.ListDataclassFieldsResponse, error) {
	requestDef := GenReqDefForListDataclassFields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataclassFieldsResponse), nil
	}
}

// ListDataclassFieldsInvoker 查询字段列表
func (c *SecMasterClient) ListDataclassFieldsInvoker(request *model.ListDataclassFieldsRequest) *ListDataclassFieldsInvoker {
	requestDef := GenReqDefForListDataclassFields()
	return &ListDataclassFieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataobjectRelations 查询关联Dataobject列表
//
// 查询关联Dataobject列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListDataobjectRelations(request *model.ListDataobjectRelationsRequest) (*model.ListDataobjectRelationsResponse, error) {
	requestDef := GenReqDefForListDataobjectRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataobjectRelationsResponse), nil
	}
}

// ListDataobjectRelationsInvoker 查询关联Dataobject列表
func (c *SecMasterClient) ListDataobjectRelationsInvoker(request *model.ListDataobjectRelationsRequest) *ListDataobjectRelationsInvoker {
	requestDef := GenReqDefForListDataobjectRelations()
	return &ListDataobjectRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIncidents 搜索事件列表
//
// 搜索事件列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListIncidents(request *model.ListIncidentsRequest) (*model.ListIncidentsResponse, error) {
	requestDef := GenReqDefForListIncidents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIncidentsResponse), nil
	}
}

// ListIncidentsInvoker 搜索事件列表
func (c *SecMasterClient) ListIncidentsInvoker(request *model.ListIncidentsRequest) *ListIncidentsInvoker {
	requestDef := GenReqDefForListIncidents()
	return &ListIncidentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIndicators 查询威胁情报列表
//
// 查询威胁情报列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListIndicators(request *model.ListIndicatorsRequest) (*model.ListIndicatorsResponse, error) {
	requestDef := GenReqDefForListIndicators()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIndicatorsResponse), nil
	}
}

// ListIndicatorsInvoker 查询威胁情报列表
func (c *SecMasterClient) ListIndicatorsInvoker(request *model.ListIndicatorsRequest) *ListIndicatorsInvoker {
	requestDef := GenReqDefForListIndicators()
	return &ListIndicatorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookActions 查询剧本动作
//
// 查询剧本动作列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookActions(request *model.ListPlaybookActionsRequest) (*model.ListPlaybookActionsResponse, error) {
	requestDef := GenReqDefForListPlaybookActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookActionsResponse), nil
	}
}

// ListPlaybookActionsInvoker 查询剧本动作
func (c *SecMasterClient) ListPlaybookActionsInvoker(request *model.ListPlaybookActionsRequest) *ListPlaybookActionsInvoker {
	requestDef := GenReqDefForListPlaybookActions()
	return &ListPlaybookActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookApproves 查询剧本审核结果
//
// 查询剧本审核结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookApproves(request *model.ListPlaybookApprovesRequest) (*model.ListPlaybookApprovesResponse, error) {
	requestDef := GenReqDefForListPlaybookApproves()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookApprovesResponse), nil
	}
}

// ListPlaybookApprovesInvoker 查询剧本审核结果
func (c *SecMasterClient) ListPlaybookApprovesInvoker(request *model.ListPlaybookApprovesRequest) *ListPlaybookApprovesInvoker {
	requestDef := GenReqDefForListPlaybookApproves()
	return &ListPlaybookApprovesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookAuditLogs 查询剧本实例审计日志
//
// 查询剧本实例审计日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookAuditLogs(request *model.ListPlaybookAuditLogsRequest) (*model.ListPlaybookAuditLogsResponse, error) {
	requestDef := GenReqDefForListPlaybookAuditLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookAuditLogsResponse), nil
	}
}

// ListPlaybookAuditLogsInvoker 查询剧本实例审计日志
func (c *SecMasterClient) ListPlaybookAuditLogsInvoker(request *model.ListPlaybookAuditLogsRequest) *ListPlaybookAuditLogsInvoker {
	requestDef := GenReqDefForListPlaybookAuditLogs()
	return &ListPlaybookAuditLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookInstances 查询剧本实例列表
//
// 查询剧本实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookInstances(request *model.ListPlaybookInstancesRequest) (*model.ListPlaybookInstancesResponse, error) {
	requestDef := GenReqDefForListPlaybookInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookInstancesResponse), nil
	}
}

// ListPlaybookInstancesInvoker 查询剧本实例列表
func (c *SecMasterClient) ListPlaybookInstancesInvoker(request *model.ListPlaybookInstancesRequest) *ListPlaybookInstancesInvoker {
	requestDef := GenReqDefForListPlaybookInstances()
	return &ListPlaybookInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybookVersions 查询剧本版本列表
//
// 查询剧本版本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybookVersions(request *model.ListPlaybookVersionsRequest) (*model.ListPlaybookVersionsResponse, error) {
	requestDef := GenReqDefForListPlaybookVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybookVersionsResponse), nil
	}
}

// ListPlaybookVersionsInvoker 查询剧本版本列表
func (c *SecMasterClient) ListPlaybookVersionsInvoker(request *model.ListPlaybookVersionsRequest) *ListPlaybookVersionsInvoker {
	requestDef := GenReqDefForListPlaybookVersions()
	return &ListPlaybookVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlaybooks 查询剧本列表
//
// 查询剧本列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListPlaybooks(request *model.ListPlaybooksRequest) (*model.ListPlaybooksResponse, error) {
	requestDef := GenReqDefForListPlaybooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlaybooksResponse), nil
	}
}

// ListPlaybooksInvoker 查询剧本列表
func (c *SecMasterClient) ListPlaybooksInvoker(request *model.ListPlaybooksRequest) *ListPlaybooksInvoker {
	requestDef := GenReqDefForListPlaybooks()
	return &ListPlaybooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflows 查询流程列表
//
// 查询流程列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

// ListWorkflowsInvoker 查询流程列表
func (c *SecMasterClient) ListWorkflowsInvoker(request *model.ListWorkflowsRequest) *ListWorkflowsInvoker {
	requestDef := GenReqDefForListWorkflows()
	return &ListWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 工作空间列表查询
//
// 工作空间列表查询:可通过工作空间名称、工作空间描述、创建时间等条件对租户的工作空间进行筛选。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 工作空间列表查询
func (c *SecMasterClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchBaseline 搜索基线检查结果列表
//
// 搜索基线检查结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) SearchBaseline(request *model.SearchBaselineRequest) (*model.SearchBaselineResponse, error) {
	requestDef := GenReqDefForSearchBaseline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchBaselineResponse), nil
	}
}

// SearchBaselineInvoker 搜索基线检查结果列表
func (c *SecMasterClient) SearchBaselineInvoker(request *model.SearchBaselineRequest) *SearchBaselineInvoker {
	requestDef := GenReqDefForSearchBaseline()
	return &SearchBaselineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlert 获取告警详情
//
// 获取告警详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlert(request *model.ShowAlertRequest) (*model.ShowAlertResponse, error) {
	requestDef := GenReqDefForShowAlert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertResponse), nil
	}
}

// ShowAlertInvoker 获取告警详情
func (c *SecMasterClient) ShowAlertInvoker(request *model.ShowAlertRequest) *ShowAlertInvoker {
	requestDef := GenReqDefForShowAlert()
	return &ShowAlertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRule 查看告警规则
//
// 查看告警规则 Get alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRule(request *model.ShowAlertRuleRequest) (*model.ShowAlertRuleResponse, error) {
	requestDef := GenReqDefForShowAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleResponse), nil
	}
}

// ShowAlertRuleInvoker 查看告警规则
func (c *SecMasterClient) ShowAlertRuleInvoker(request *model.ShowAlertRuleRequest) *ShowAlertRuleInvoker {
	requestDef := GenReqDefForShowAlertRule()
	return &ShowAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlertRuleTemplate 查看告警规则模板
//
// # List alert rule templates
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowAlertRuleTemplate(request *model.ShowAlertRuleTemplateRequest) (*model.ShowAlertRuleTemplateResponse, error) {
	requestDef := GenReqDefForShowAlertRuleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertRuleTemplateResponse), nil
	}
}

// ShowAlertRuleTemplateInvoker 查看告警规则模板
func (c *SecMasterClient) ShowAlertRuleTemplateInvoker(request *model.ShowAlertRuleTemplateRequest) *ShowAlertRuleTemplateInvoker {
	requestDef := GenReqDefForShowAlertRuleTemplate()
	return &ShowAlertRuleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIncident 获取事件详情
//
// 获取事件详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowIncident(request *model.ShowIncidentRequest) (*model.ShowIncidentResponse, error) {
	requestDef := GenReqDefForShowIncident()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIncidentResponse), nil
	}
}

// ShowIncidentInvoker 获取事件详情
func (c *SecMasterClient) ShowIncidentInvoker(request *model.ShowIncidentRequest) *ShowIncidentInvoker {
	requestDef := GenReqDefForShowIncident()
	return &ShowIncidentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowIndicatorDetail 查询威胁情报详情
//
// 查询威胁情报详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowIndicatorDetail(request *model.ShowIndicatorDetailRequest) (*model.ShowIndicatorDetailResponse, error) {
	requestDef := GenReqDefForShowIndicatorDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowIndicatorDetailResponse), nil
	}
}

// ShowIndicatorDetailInvoker 查询威胁情报详情
func (c *SecMasterClient) ShowIndicatorDetailInvoker(request *model.ShowIndicatorDetailRequest) *ShowIndicatorDetailInvoker {
	requestDef := GenReqDefForShowIndicatorDetail()
	return &ShowIndicatorDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybook 查询剧本详情
//
// 查询剧本详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybook(request *model.ShowPlaybookRequest) (*model.ShowPlaybookResponse, error) {
	requestDef := GenReqDefForShowPlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookResponse), nil
	}
}

// ShowPlaybookInvoker 查询剧本详情
func (c *SecMasterClient) ShowPlaybookInvoker(request *model.ShowPlaybookRequest) *ShowPlaybookInvoker {
	requestDef := GenReqDefForShowPlaybook()
	return &ShowPlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookInstance 查询剧本实例详情
//
// # Show playbook instance
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookInstance(request *model.ShowPlaybookInstanceRequest) (*model.ShowPlaybookInstanceResponse, error) {
	requestDef := GenReqDefForShowPlaybookInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookInstanceResponse), nil
	}
}

// ShowPlaybookInstanceInvoker 查询剧本实例详情
func (c *SecMasterClient) ShowPlaybookInstanceInvoker(request *model.ShowPlaybookInstanceRequest) *ShowPlaybookInstanceInvoker {
	requestDef := GenReqDefForShowPlaybookInstance()
	return &ShowPlaybookInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookMonitors 剧本运行监控
//
// 剧本运行监控
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookMonitors(request *model.ShowPlaybookMonitorsRequest) (*model.ShowPlaybookMonitorsResponse, error) {
	requestDef := GenReqDefForShowPlaybookMonitors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookMonitorsResponse), nil
	}
}

// ShowPlaybookMonitorsInvoker 剧本运行监控
func (c *SecMasterClient) ShowPlaybookMonitorsInvoker(request *model.ShowPlaybookMonitorsRequest) *ShowPlaybookMonitorsInvoker {
	requestDef := GenReqDefForShowPlaybookMonitors()
	return &ShowPlaybookMonitorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookRule 查询剧本规则详情
//
// 查询剧本规则详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookRule(request *model.ShowPlaybookRuleRequest) (*model.ShowPlaybookRuleResponse, error) {
	requestDef := GenReqDefForShowPlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookRuleResponse), nil
	}
}

// ShowPlaybookRuleInvoker 查询剧本规则详情
func (c *SecMasterClient) ShowPlaybookRuleInvoker(request *model.ShowPlaybookRuleRequest) *ShowPlaybookRuleInvoker {
	requestDef := GenReqDefForShowPlaybookRule()
	return &ShowPlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookStatistics 剧本数据统计
//
// 剧本统计数据
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookStatistics(request *model.ShowPlaybookStatisticsRequest) (*model.ShowPlaybookStatisticsResponse, error) {
	requestDef := GenReqDefForShowPlaybookStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookStatisticsResponse), nil
	}
}

// ShowPlaybookStatisticsInvoker 剧本数据统计
func (c *SecMasterClient) ShowPlaybookStatisticsInvoker(request *model.ShowPlaybookStatisticsRequest) *ShowPlaybookStatisticsInvoker {
	requestDef := GenReqDefForShowPlaybookStatistics()
	return &ShowPlaybookStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookTopology 查询剧本拓扑关系
//
// 查询剧本拓扑关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookTopology(request *model.ShowPlaybookTopologyRequest) (*model.ShowPlaybookTopologyResponse, error) {
	requestDef := GenReqDefForShowPlaybookTopology()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookTopologyResponse), nil
	}
}

// ShowPlaybookTopologyInvoker 查询剧本拓扑关系
func (c *SecMasterClient) ShowPlaybookTopologyInvoker(request *model.ShowPlaybookTopologyRequest) *ShowPlaybookTopologyInvoker {
	requestDef := GenReqDefForShowPlaybookTopology()
	return &ShowPlaybookTopologyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlaybookVersion 查询剧本版本详情
//
// # Show playbook version version
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) ShowPlaybookVersion(request *model.ShowPlaybookVersionRequest) (*model.ShowPlaybookVersionResponse, error) {
	requestDef := GenReqDefForShowPlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPlaybookVersionResponse), nil
	}
}

// ShowPlaybookVersionInvoker 查询剧本版本详情
func (c *SecMasterClient) ShowPlaybookVersionInvoker(request *model.ShowPlaybookVersionRequest) *ShowPlaybookVersionInvoker {
	requestDef := GenReqDefForShowPlaybookVersion()
	return &ShowPlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertRule 更新告警规则
//
// # Update alert rule
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateAlertRule(request *model.UpdateAlertRuleRequest) (*model.UpdateAlertRuleResponse, error) {
	requestDef := GenReqDefForUpdateAlertRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertRuleResponse), nil
	}
}

// UpdateAlertRuleInvoker 更新告警规则
func (c *SecMasterClient) UpdateAlertRuleInvoker(request *model.UpdateAlertRuleRequest) *UpdateAlertRuleInvoker {
	requestDef := GenReqDefForUpdateAlertRule()
	return &UpdateAlertRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateIndicator 更新威胁情报
//
// 更新威胁情报
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdateIndicator(request *model.UpdateIndicatorRequest) (*model.UpdateIndicatorResponse, error) {
	requestDef := GenReqDefForUpdateIndicator()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateIndicatorResponse), nil
	}
}

// UpdateIndicatorInvoker 更新威胁情报
func (c *SecMasterClient) UpdateIndicatorInvoker(request *model.UpdateIndicatorRequest) *UpdateIndicatorInvoker {
	requestDef := GenReqDefForUpdateIndicator()
	return &UpdateIndicatorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybook 修改剧本
//
// 修改剧本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybook(request *model.UpdatePlaybookRequest) (*model.UpdatePlaybookResponse, error) {
	requestDef := GenReqDefForUpdatePlaybook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookResponse), nil
	}
}

// UpdatePlaybookInvoker 修改剧本
func (c *SecMasterClient) UpdatePlaybookInvoker(request *model.UpdatePlaybookRequest) *UpdatePlaybookInvoker {
	requestDef := GenReqDefForUpdatePlaybook()
	return &UpdatePlaybookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookAction 更新剧本动作
//
// 更新剧本动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookAction(request *model.UpdatePlaybookActionRequest) (*model.UpdatePlaybookActionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookActionResponse), nil
	}
}

// UpdatePlaybookActionInvoker 更新剧本动作
func (c *SecMasterClient) UpdatePlaybookActionInvoker(request *model.UpdatePlaybookActionRequest) *UpdatePlaybookActionInvoker {
	requestDef := GenReqDefForUpdatePlaybookAction()
	return &UpdatePlaybookActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookRule 更新剧本规则
//
// 更新剧本规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookRule(request *model.UpdatePlaybookRuleRequest) (*model.UpdatePlaybookRuleResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookRuleResponse), nil
	}
}

// UpdatePlaybookRuleInvoker 更新剧本规则
func (c *SecMasterClient) UpdatePlaybookRuleInvoker(request *model.UpdatePlaybookRuleRequest) *UpdatePlaybookRuleInvoker {
	requestDef := GenReqDefForUpdatePlaybookRule()
	return &UpdatePlaybookRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdatePlaybookVersion 更新剧本版本
//
// 更新剧本版本
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *SecMasterClient) UpdatePlaybookVersion(request *model.UpdatePlaybookVersionRequest) (*model.UpdatePlaybookVersionResponse, error) {
	requestDef := GenReqDefForUpdatePlaybookVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePlaybookVersionResponse), nil
	}
}

// UpdatePlaybookVersionInvoker 更新剧本版本
func (c *SecMasterClient) UpdatePlaybookVersionInvoker(request *model.UpdatePlaybookVersionRequest) *UpdatePlaybookVersionInvoker {
	requestDef := GenReqDefForUpdatePlaybookVersion()
	return &UpdatePlaybookVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
