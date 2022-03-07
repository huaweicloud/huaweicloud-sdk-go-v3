package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type CesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCesClient(hcClient *http_client.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//增加告警资源(资源分组类型的告警规则不支持)
func (c *CesClient) AddAlarmResources(request *model.AddAlarmResourcesRequest) (*model.AddAlarmResourcesResponse, error) {
	requestDef := GenReqDefForAddAlarmResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAlarmResourcesResponse), nil
	}
}

//给自定义资源分组,即非自动企业项目和标签关联,批量增加关联资源
func (c *CesClient) AddResourceGroupsResourcesBatch(request *model.AddResourceGroupsResourcesBatchRequest) (*model.AddResourceGroupsResourcesBatchResponse, error) {
	requestDef := GenReqDefForAddResourceGroupsResourcesBatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddResourceGroupsResourcesBatchResponse), nil
	}
}

//创建告警，接口对应后端API路径open-alarms
func (c *CesClient) CreateAlarm(request *model.CreateAlarmRequest) (*model.CreateAlarmResponse, error) {
	requestDef := GenReqDefForCreateAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmResponse), nil
	}
}

//删除告警规则V2接口，APIG上注册为V2，后端实际上使用的是V1
func (c *CesClient) DeleteAlarm(request *model.DeleteAlarmRequest) (*model.DeleteAlarmResponse, error) {
	requestDef := GenReqDefForDeleteAlarm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmResponse), nil
	}
}

//删除告警资源（资源分组类型的告警规则不支持）
func (c *CesClient) DeleteAlarmResources(request *model.DeleteAlarmResourcesRequest) (*model.DeleteAlarmResourcesResponse, error) {
	requestDef := GenReqDefForDeleteAlarmResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmResourcesResponse), nil
	}
}

//给自定义资源分组,即非自动企业项目和标签关联,批量删除关联资源
func (c *CesClient) DeleteResourceGroupsResourcesBatch(request *model.DeleteResourceGroupsResourcesBatchRequest) (*model.DeleteResourceGroupsResourcesBatchResponse, error) {
	requestDef := GenReqDefForDeleteResourceGroupsResourcesBatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceGroupsResourcesBatchResponse), nil
	}
}

//查询告警历史列表
func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}

//根据告警ID查询告警资源列表
func (c *CesClient) ListAlarmResources(request *model.ListAlarmResourcesRequest) (*model.ListAlarmResourcesResponse, error) {
	requestDef := GenReqDefForListAlarmResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmResourcesResponse), nil
	}
}

//查询告警列表,实际上内部对应的是v3的版本号
func (c *CesClient) ListAlarms(request *model.ListAlarmsRequest) (*model.ListAlarmsResponse, error) {
	requestDef := GenReqDefForListAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmsResponse), nil
	}
}

//修改告警规则基本信息
func (c *CesClient) UpdateAlarmAction(request *model.UpdateAlarmActionRequest) (*model.UpdateAlarmActionResponse, error) {
	requestDef := GenReqDefForUpdateAlarmAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmActionResponse), nil
	}
}
