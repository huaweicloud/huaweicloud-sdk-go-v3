package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
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

// AddAlarmRuleResources 批量增加告警规则资源
//
// 批量增加告警规则资源(资源分组类型的告警规则不支持)，资源分组类型的修改请使用资源分组管理相关接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) AddAlarmRuleResources(request *model.AddAlarmRuleResourcesRequest) (*model.AddAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForAddAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAlarmRuleResourcesResponse), nil
	}
}

// AddAlarmRuleResourcesInvoker 批量增加告警规则资源
func (c *CesClient) AddAlarmRuleResourcesInvoker(request *model.AddAlarmRuleResourcesRequest) *AddAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForAddAlarmRuleResources()
	return &AddAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAlarmRules 批量删除告警规则
//
// 批量删除告警规则V2接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) BatchDeleteAlarmRules(request *model.BatchDeleteAlarmRulesRequest) (*model.BatchDeleteAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAlarmRulesResponse), nil
	}
}

// BatchDeleteAlarmRulesInvoker 批量删除告警规则
func (c *CesClient) BatchDeleteAlarmRulesInvoker(request *model.BatchDeleteAlarmRulesRequest) *BatchDeleteAlarmRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAlarmRules()
	return &BatchDeleteAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableAlarmRules 批量启停告警规则
//
// 批量启停告警规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) BatchEnableAlarmRules(request *model.BatchEnableAlarmRulesRequest) (*model.BatchEnableAlarmRulesResponse, error) {
	requestDef := GenReqDefForBatchEnableAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAlarmRulesResponse), nil
	}
}

// BatchEnableAlarmRulesInvoker 批量启停告警规则
func (c *CesClient) BatchEnableAlarmRulesInvoker(request *model.BatchEnableAlarmRulesRequest) *BatchEnableAlarmRulesInvoker {
	requestDef := GenReqDefForBatchEnableAlarmRules()
	return &BatchEnableAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlarmRules 创建告警规则
//
// 创建告警规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) CreateAlarmRules(request *model.CreateAlarmRulesRequest) (*model.CreateAlarmRulesResponse, error) {
	requestDef := GenReqDefForCreateAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlarmRulesResponse), nil
	}
}

// CreateAlarmRulesInvoker 创建告警规则
func (c *CesClient) CreateAlarmRulesInvoker(request *model.CreateAlarmRulesRequest) *CreateAlarmRulesInvoker {
	requestDef := GenReqDefForCreateAlarmRules()
	return &CreateAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlarmRuleResources 批量删除告警规则资源
//
// 批量删除告警规则资源（资源分组类型的告警规则不支持），资源分组类型的修改请使用资源分组管理相关接口
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) DeleteAlarmRuleResources(request *model.DeleteAlarmRuleResourcesRequest) (*model.DeleteAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForDeleteAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlarmRuleResourcesResponse), nil
	}
}

// DeleteAlarmRuleResourcesInvoker 批量删除告警规则资源
func (c *CesClient) DeleteAlarmRuleResourcesInvoker(request *model.DeleteAlarmRuleResourcesRequest) *DeleteAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForDeleteAlarmRuleResources()
	return &DeleteAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAgentDimensionInfo 查询主机监控维度指标信息
//
// 根据ECS/BMS资源ID查询磁盘、挂载点、进程、显卡、RAID控制器维度指标信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) ListAgentDimensionInfo(request *model.ListAgentDimensionInfoRequest) (*model.ListAgentDimensionInfoResponse, error) {
	requestDef := GenReqDefForListAgentDimensionInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAgentDimensionInfoResponse), nil
	}
}

// ListAgentDimensionInfoInvoker 查询主机监控维度指标信息
func (c *CesClient) ListAgentDimensionInfoInvoker(request *model.ListAgentDimensionInfoRequest) *ListAgentDimensionInfoInvoker {
	requestDef := GenReqDefForListAgentDimensionInfo()
	return &ListAgentDimensionInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmHistories 查询告警记录列表
//
// 查询告警记录列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}

// ListAlarmHistoriesInvoker 查询告警记录列表
func (c *CesClient) ListAlarmHistoriesInvoker(request *model.ListAlarmHistoriesRequest) *ListAlarmHistoriesInvoker {
	requestDef := GenReqDefForListAlarmHistories()
	return &ListAlarmHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRulePolicies 查询告警规则策略列表
//
// 根据告警规则ID查询策略列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) ListAlarmRulePolicies(request *model.ListAlarmRulePoliciesRequest) (*model.ListAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForListAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulePoliciesResponse), nil
	}
}

// ListAlarmRulePoliciesInvoker 查询告警规则策略列表
func (c *CesClient) ListAlarmRulePoliciesInvoker(request *model.ListAlarmRulePoliciesRequest) *ListAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForListAlarmRulePolicies()
	return &ListAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRuleResources 查询告警规则资源列表
//
// 根据告警规则ID查询告警规则资源列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) ListAlarmRuleResources(request *model.ListAlarmRuleResourcesRequest) (*model.ListAlarmRuleResourcesResponse, error) {
	requestDef := GenReqDefForListAlarmRuleResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRuleResourcesResponse), nil
	}
}

// ListAlarmRuleResourcesInvoker 查询告警规则资源列表
func (c *CesClient) ListAlarmRuleResourcesInvoker(request *model.ListAlarmRuleResourcesRequest) *ListAlarmRuleResourcesInvoker {
	requestDef := GenReqDefForListAlarmRuleResources()
	return &ListAlarmRuleResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlarmRules 查询告警规则列表
//
// 查询告警规则列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) ListAlarmRules(request *model.ListAlarmRulesRequest) (*model.ListAlarmRulesResponse, error) {
	requestDef := GenReqDefForListAlarmRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmRulesResponse), nil
	}
}

// ListAlarmRulesInvoker 查询告警规则列表
func (c *CesClient) ListAlarmRulesInvoker(request *model.ListAlarmRulesRequest) *ListAlarmRulesInvoker {
	requestDef := GenReqDefForListAlarmRules()
	return &ListAlarmRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlarmRulePolicies 修改告警规则策略(全量修改)
//
// 修改告警规则策略(全量修改)
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CesClient) UpdateAlarmRulePolicies(request *model.UpdateAlarmRulePoliciesRequest) (*model.UpdateAlarmRulePoliciesResponse, error) {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlarmRulePoliciesResponse), nil
	}
}

// UpdateAlarmRulePoliciesInvoker 修改告警规则策略(全量修改)
func (c *CesClient) UpdateAlarmRulePoliciesInvoker(request *model.UpdateAlarmRulePoliciesRequest) *UpdateAlarmRulePoliciesInvoker {
	requestDef := GenReqDefForUpdateAlarmRulePolicies()
	return &UpdateAlarmRulePoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
