package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
)

type CesClient struct {
	hcClient *http_client.HcHttpClient
}

func NewCesClient(hcClient *http_client.HcHttpClient) *CesClient {
	return &CesClient{hcClient: hcClient}
}

func CesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//批量查询指定时间范围内指定指标的指定粒度的监控数据，目前最多支持10指标的批量查询。
func (c *CesClient) BatchListMetricData(request *model.BatchListMetricDataRequest) (*model.BatchListMetricDataResponse, error) {
	requestDef := GenReqDefForBatchListMetricData(request)
	resp, responseDef := GenRespForBatchListMetricData()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//创建一条告警规则。
func (c *CesClient) CreateAlarm(request *model.CreateAlarmRequest) (*model.CreateAlarmResponse, error) {
	requestDef := GenReqDefForCreateAlarm(request)
	resp, responseDef := GenRespForCreateAlarm()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//上报自定义事件。
func (c *CesClient) CreateEvents(request *model.CreateEventsRequest) (*model.CreateEventsResponse, error) {
	requestDef := GenReqDefForCreateEvents(request)
	resp, responseDef := GenRespForCreateEvents()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//添加一条或多条指标监控数据。
func (c *CesClient) CreateMetricData(request *model.CreateMetricDataRequest) (*model.CreateMetricDataResponse, error) {
	requestDef := GenReqDefForCreateMetricData(request)
	resp, responseDef := GenRespForCreateMetricData()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//删除一条告警规则。
func (c *CesClient) DeleteAlarm(request *model.DeleteAlarmRequest) (*model.DeleteAlarmResponse, error) {
	requestDef := GenReqDefForDeleteAlarm(request)
	resp, responseDef := GenRespForDeleteAlarm()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询告警规则列表，可以指定分页条件限制结果数量，可以指定排序规则。
func (c *CesClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms(request)
	resp, responseDef := GenRespForListAlarms()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询系统当前可监控指标列表，可以指定指标命名空间、指标名称、维度、排序方式，起始记录和最大记录条数过滤查询结果。
func (c *CesClient) ListMetrics(request *model.ListMetricsRequest) (*model.ListMetricsResponse, error) {
	requestDef := GenReqDefForListMetrics(request)
	resp, responseDef := GenRespForListMetrics()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//根据告警ID查询告警规则信息。
func (c *CesClient) ShowAlarm(request *model.ShowAlarmRequest) (*model.ShowAlarmResponse, error) {
	requestDef := GenReqDefForShowAlarm(request)
	resp, responseDef := GenRespForShowAlarm()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定时间范围指定事件类型的主机配置数据，可以通过参数指定需要查询的数据维度。注意：该接口提供给HANA场景下SAP Monitor查询主机配置数据，其他场景下查不到主机配置数据。
func (c *CesClient) ShowEventData(request *model.ShowEventDataRequest) (*model.ShowEventDataResponse, error) {
	requestDef := GenReqDefForShowEventData(request)
	resp, responseDef := GenRespForShowEventData()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询指定时间范围指定指标的指定粒度的监控数据，可以通过参数指定需要查询的数据维度。
func (c *CesClient) ShowMetricData(request *model.ShowMetricDataRequest) (*model.ShowMetricDataResponse, error) {
	requestDef := GenReqDefForShowMetricData(request)
	resp, responseDef := GenRespForShowMetricData()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//查询用户可以创建的资源配额总数及当前使用量，当前仅有告警规则一种资源类型。
func (c *CesClient) ShowQuotas(request *model.ShowQuotasRequest) (*model.ShowQuotasResponse, error) {
	requestDef := GenReqDefForShowQuotas(request)
	resp, responseDef := GenRespForShowQuotas()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//启动或停止一条告警规则。
func (c *CesClient) UpdateAlarmAction(request *model.UpdateAlarmActionRequest) (*model.UpdateAlarmActionResponse, error) {
	requestDef := GenReqDefForUpdateAlarmAction(request)
	resp, responseDef := GenRespForUpdateAlarmAction()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
