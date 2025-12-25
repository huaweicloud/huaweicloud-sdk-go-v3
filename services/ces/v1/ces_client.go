package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
)

type CesClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCesClient(hcClient *httpclient.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// BatchListMetricData 批量查询监控数据
//
// 批量查询指定时间范围内指定指标的指定粒度的监控数据，目前最多支持500指标的批量查询。 对于不同的period取值和查询的指标数量，默认的最大查询区间(to-from)不同。 规则为\&quot;指标数量*(to-from)/监控周期&lt;&#x3D;3000\&quot;，若超出阈值，会自动调整from以满足规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) BatchListMetricData(request *model.BatchListMetricDataRequest) (*model.BatchListMetricDataResponse, error) {
	requestDef := GenReqDefForBatchListMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchListMetricDataResponse), nil
	}
}

// BatchListMetricDataInvoker 批量查询监控数据
func (c *CesClient) BatchListMetricDataInvoker(request *model.BatchListMetricDataRequest) *BatchListMetricDataInvoker {
	requestDef := GenReqDefForBatchListMetricData()
	return &BatchListMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarm 创建告警规则（V1）
//
// 创建一条告警规则。创建告警规则V1接口只支持配置单资源单策略规则，建议使用“[创建告警规则（推荐）](CreateAlarmRules.xml)”与前端功能配套使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateAlarm(request *model.CreateAlarmRequest) (*model.CreateAlarmResponse, error) {
	requestDef := GenReqDefForCreateAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmResponse), nil
	}
}

// CreateAlarmInvoker 创建告警规则（V1）
func (c *CesClient) CreateAlarmInvoker(request *model.CreateAlarmRequest) *CreateAlarmInvoker {
	requestDef := GenReqDefForCreateAlarm()
	return &CreateAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarmTemplate 创建自定义告警模板
//
// 创建自定义告警模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateAlarmTemplate(request *model.CreateAlarmTemplateRequest) (*model.CreateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForCreateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmTemplateResponse), nil
	}
}

// CreateAlarmTemplateInvoker 创建自定义告警模板
func (c *CesClient) CreateAlarmTemplateInvoker(request *model.CreateAlarmTemplateRequest) *CreateAlarmTemplateInvoker {
	requestDef := GenReqDefForCreateAlarmTemplate()
	return &CreateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateEvents 上报事件
//
// 上报自定义事件。事件的time、project_id、event_source、event_name、event_type、sub_event_type、event_state、event_level、event_user、resource_id、resource_name字段相同时，则视为同一条事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateEvents(request *model.CreateEventsRequest) (*model.CreateEventsResponse, error) {
	requestDef := GenReqDefForCreateEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEventsResponse), nil
	}
}

// CreateEventsInvoker 上报事件
func (c *CesClient) CreateEventsInvoker(request *model.CreateEventsRequest) *CreateEventsInvoker {
	requestDef := GenReqDefForCreateEvents()
	return &CreateEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMetricData 添加监控数据
//
// 添加一条或多条指标监控数据。约束与限制：
// 1. 单次POST请求消息体大小不能超过512KB，否则请求会被服务端拒绝。
// 2. POST请求发送周期应小于最小聚合周期，否则会出现聚合数据点不连续。例如：聚合周期为5分钟，发送周期为7分钟，则5分钟情况的聚合数据会出现每10分钟才出现一个点。
// 3. POST请求体中数据收集时间（collect_time）的值必须从当前时间的前三天到当前时间后的十分钟之内某一时间，如果不在这个范围内，则不允许插入指标数据。
// 4. 如果指标上报时间（即调用指标上报接口的时间）与数据收集时间（collect_time）之间的延迟超过10分钟，CES在聚合数据时会丢弃此指标数据。您只能查看近2天的原始指标数据，聚合数据中不会显示这些延迟上报的指标。例如，14:20:00调用CES接口上报数据，请求体中的collect_time字段为14:05:00，表示延迟上报了15分钟。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateMetricData(request *model.CreateMetricDataRequest) (*model.CreateMetricDataResponse, error) {
	requestDef := GenReqDefForCreateMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMetricDataResponse), nil
	}
}

// CreateMetricDataInvoker 添加监控数据
func (c *CesClient) CreateMetricDataInvoker(request *model.CreateMetricDataRequest) *CreateMetricDataInvoker {
	requestDef := GenReqDefForCreateMetricData()
	return &CreateMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceGroup 创建资源分组
//
// 创建资源分组，资源分组支持将各类资源按照业务集中进行分组管理，可以从分组角度查看监控与告警信息，以提升运维效率。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) CreateResourceGroup(request *model.CreateResourceGroupRequest) (*model.CreateResourceGroupResponse, error) {
	requestDef := GenReqDefForCreateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceGroupResponse), nil
	}
}

// CreateResourceGroupInvoker 创建资源分组
func (c *CesClient) CreateResourceGroupInvoker(request *model.CreateResourceGroupRequest) *CreateResourceGroupInvoker {
	requestDef := GenReqDefForCreateResourceGroup()
	return &CreateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarm 删除告警规则
//
// 删除一条告警规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteAlarm(request *model.DeleteAlarmRequest) (*model.DeleteAlarmResponse, error) {
	requestDef := GenReqDefForDeleteAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmResponse), nil
	}
}

// DeleteAlarmInvoker 删除告警规则
func (c *CesClient) DeleteAlarmInvoker(request *model.DeleteAlarmRequest) *DeleteAlarmInvoker {
	requestDef := GenReqDefForDeleteAlarm()
	return &DeleteAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarmTemplate 删除自定义告警模板
//
// 根据ID删除自定义告警模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteAlarmTemplate(request *model.DeleteAlarmTemplateRequest) (*model.DeleteAlarmTemplateResponse, error) {
	requestDef := GenReqDefForDeleteAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmTemplateResponse), nil
	}
}

// DeleteAlarmTemplateInvoker 删除自定义告警模板
func (c *CesClient) DeleteAlarmTemplateInvoker(request *model.DeleteAlarmTemplateRequest) *DeleteAlarmTemplateInvoker {
	requestDef := GenReqDefForDeleteAlarmTemplate()
	return &DeleteAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceGroup 删除资源分组
//
// 删除一条资源分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) DeleteResourceGroup(request *model.DeleteResourceGroupRequest) (*model.DeleteResourceGroupResponse, error) {
	requestDef := GenReqDefForDeleteResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceGroupResponse), nil
	}
}

// DeleteResourceGroupInvoker 删除资源分组
func (c *CesClient) DeleteResourceGroupInvoker(request *model.DeleteResourceGroupRequest) *DeleteResourceGroupInvoker {
	requestDef := GenReqDefForDeleteResourceGroup()
	return &DeleteResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmHistories 查询告警历史
//
// 查询告警历史列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}

// ListAlarmHistoriesInvoker 查询告警历史
func (c *CesClient) ListAlarmHistoriesInvoker(request *model.ListAlarmHistoriesRequest) *ListAlarmHistoriesInvoker {
	requestDef := GenReqDefForListAlarmHistories()
	return &ListAlarmHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmTemplates 查询自定义告警模板列表
//
// 查询自定义告警模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarmTemplates(request *model.ListAlarmTemplatesRequest) (*model.ListAlarmTemplatesResponse, error) {
	requestDef := GenReqDefForListAlarmTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmTemplatesResponse), nil
	}
}

// ListAlarmTemplatesInvoker 查询自定义告警模板列表
func (c *CesClient) ListAlarmTemplatesInvoker(request *model.ListAlarmTemplatesRequest) *ListAlarmTemplatesInvoker {
	requestDef := GenReqDefForListAlarmTemplates()
	return &ListAlarmTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarms 查询告警规则列表（V1）
//
// 查询告警规则列表，可以指定分页条件限制结果数量，可以指定排序规则。如果用本接口去查询多资源多策略的告警规则，也只能返回告警规则的某个策略，建议使用“[查询告警规则列表（推荐）](ListAlarmRules.xml)”与前端功能配套使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

// ListAlarmsInvoker 查询告警规则列表（V1）
func (c *CesClient) ListAlarmsInvoker(request *model.ListAlarmsRequest) *ListAlarmsInvoker {
	requestDef := GenReqDefForListAlarms()
	return &ListAlarmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventDetail 查询某一事件监控详情
//
// 根据事件监控名称，查询该事件发生的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListEventDetail(request *model.ListEventDetailRequest) (*model.ListEventDetailResponse, error) {
	requestDef := GenReqDefForListEventDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventDetailResponse), nil
	}
}

// ListEventDetailInvoker 查询某一事件监控详情
func (c *CesClient) ListEventDetailInvoker(request *model.ListEventDetailRequest) *ListEventDetailInvoker {
	requestDef := GenReqDefForListEventDetail()
	return &ListEventDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 查询事件监控列表
//
// 查询事件监控列表，包括系统事件和自定义事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 查询事件监控列表
func (c *CesClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetrics 查询指标列表
//
// 查询系统当前可监控指标列表，可以指定指标命名空间、指标名称、维度、排序方式，起始记录和最大记录条数过滤查询结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricsResponse), nil
	}
}

// ListMetricsInvoker 查询指标列表
func (c *CesClient) ListMetricsInvoker(request *model.ListMetricsRequest) *ListMetricsInvoker {
	requestDef := GenReqDefForListMetrics()
	return &ListMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceGroup 查询所有资源分组
//
// 查询所创建的所有资源分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ListResourceGroup(request *model.ListResourceGroupRequest) (*model.ListResourceGroupResponse, error) {
	requestDef := GenReqDefForListResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceGroupResponse), nil
	}
}

// ListResourceGroupInvoker 查询所有资源分组
func (c *CesClient) ListResourceGroupInvoker(request *model.ListResourceGroupRequest) *ListResourceGroupInvoker {
	requestDef := GenReqDefForListResourceGroup()
	return &ListResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlarm 查询单条告警规则信息
//
// 根据告警ID查询告警规则信息。告警规则V1接口只支持配置单资源单策略规则，建议使用“[查询告警规则列表（推荐）](ListAlarmRules.xml)”、“[查询告警规则资源列表](ListAlarmRuleResources.xml)”与前端功能配套使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowAlarm(request *model.ShowAlarmRequest) (*model.ShowAlarmResponse, error) {
	requestDef := GenReqDefForShowAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlarmResponse), nil
	}
}

// ShowAlarmInvoker 查询单条告警规则信息
func (c *CesClient) ShowAlarmInvoker(request *model.ShowAlarmRequest) *ShowAlarmInvoker {
	requestDef := GenReqDefForShowAlarm()
	return &ShowAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEventData 查询主机配置数据
//
// 查询指定时间范围指定事件类型的主机配置数据，可以通过参数指定需要查询的数据维度。注意：该接口提供给HANA场景下SAP Monitor查询主机配置数据，其他场景下查不到主机配置数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowEventData(request *model.ShowEventDataRequest) (*model.ShowEventDataResponse, error) {
	requestDef := GenReqDefForShowEventData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEventDataResponse), nil
	}
}

// ShowEventDataInvoker 查询主机配置数据
func (c *CesClient) ShowEventDataInvoker(request *model.ShowEventDataRequest) *ShowEventDataInvoker {
	requestDef := GenReqDefForShowEventData()
	return &ShowEventDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricData 查询监控数据
//
// 查询指定时间范围指定指标的指定粒度的监控数据，可以通过参数指定需要查询的数据维度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowMetricData(request *model.ShowMetricDataRequest) (*model.ShowMetricDataResponse, error) {
	requestDef := GenReqDefForShowMetricData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricDataResponse), nil
	}
}

// ShowMetricDataInvoker 查询监控数据
func (c *CesClient) ShowMetricDataInvoker(request *model.ShowMetricDataRequest) *ShowMetricDataInvoker {
	requestDef := GenReqDefForShowMetricData()
	return &ShowMetricDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQuotas 查询配额
//
// 查询用户可以创建的资源配额总数及当前使用量，当前仅有告警规则一种资源类型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQuotasResponse), nil
	}
}

// ShowQuotasInvoker 查询配额
func (c *CesClient) ShowQuotasInvoker(request *model.ShowQuotasRequest) *ShowQuotasInvoker {
	requestDef := GenReqDefForShowQuotas()
	return &ShowQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceGroup 查询资源分组下的资源
//
// 根据资源分组ID查询资源分组下的资源。此接口已过时，建议使用v2接口 “[查询资源分组下指定服务类别特定维度的资源列表](ListResourceGroupsServicesResources.xml)”
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) ShowResourceGroup(request *model.ShowResourceGroupRequest) (*model.ShowResourceGroupResponse, error) {
	requestDef := GenReqDefForShowResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceGroupResponse), nil
	}
}

// ShowResourceGroupInvoker 查询资源分组下的资源
func (c *CesClient) ShowResourceGroupInvoker(request *model.ShowResourceGroupRequest) *ShowResourceGroupInvoker {
	requestDef := GenReqDefForShowResourceGroup()
	return &ShowResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarm 修改告警规则
//
// 修改告警规则。
// 告警规则V1接口只支持配置单资源单策略规则，建议使用批量增加告警规则资源、批量删除告警规则资源和修改告警规则策略(全量修改)与前端功能配套使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateAlarm(request *model.UpdateAlarmRequest) (*model.UpdateAlarmResponse, error) {
	requestDef := GenReqDefForUpdateAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmResponse), nil
	}
}

// UpdateAlarmInvoker 修改告警规则
func (c *CesClient) UpdateAlarmInvoker(request *model.UpdateAlarmRequest) *UpdateAlarmInvoker {
	requestDef := GenReqDefForUpdateAlarm()
	return &UpdateAlarmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmAction 启停告警规则
//
// 启动或停止一条告警规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateAlarmAction(request *model.UpdateAlarmActionRequest) (*model.UpdateAlarmActionResponse, error) {
	requestDef := GenReqDefForUpdateAlarmAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmActionResponse), nil
	}
}

// UpdateAlarmActionInvoker 启停告警规则
func (c *CesClient) UpdateAlarmActionInvoker(request *model.UpdateAlarmActionRequest) *UpdateAlarmActionInvoker {
	requestDef := GenReqDefForUpdateAlarmAction()
	return &UpdateAlarmActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmTemplate 更新自定义告警模板
//
// 更新自定义告警模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateAlarmTemplate(request *model.UpdateAlarmTemplateRequest) (*model.UpdateAlarmTemplateResponse, error) {
	requestDef := GenReqDefForUpdateAlarmTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmTemplateResponse), nil
	}
}

// UpdateAlarmTemplateInvoker 更新自定义告警模板
func (c *CesClient) UpdateAlarmTemplateInvoker(request *model.UpdateAlarmTemplateRequest) *UpdateAlarmTemplateInvoker {
	requestDef := GenReqDefForUpdateAlarmTemplate()
	return &UpdateAlarmTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateResourceGroup 更新资源分组
//
// 更新资源分组，资源分组支持将各类资源按照业务集中进行分组管理，可以从分组角度查看监控与告警信息，以提升运维效率。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CesClient) UpdateResourceGroup(request *model.UpdateResourceGroupRequest) (*model.UpdateResourceGroupResponse, error) {
	requestDef := GenReqDefForUpdateResourceGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateResourceGroupResponse), nil
	}
}

// UpdateResourceGroupInvoker 更新资源分组
func (c *CesClient) UpdateResourceGroupInvoker(request *model.UpdateResourceGroupRequest) *UpdateResourceGroupInvoker {
	requestDef := GenReqDefForUpdateResourceGroup()
	return &UpdateResourceGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
