package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
)

type CseClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCseClient(hcClient *httpclient.HcHttpClient) *CseClient {
	return &CseClient{HcClient: hcClient}
}

func CseClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CreateEngine 创建微服务引擎
//
// 创建微服务引擎，支持创建ServiceComb引擎专享版、注册配置中心、应用网关（公测）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateEngine(request *model.CreateEngineRequest) (*model.CreateEngineResponse, error) {
	requestDef := GenReqDefForCreateEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEngineResponse), nil
	}
}

// CreateEngineInvoker 创建微服务引擎
func (c *CseClient) CreateEngineInvoker(request *model.CreateEngineRequest) *CreateEngineInvoker {
	requestDef := GenReqDefForCreateEngine()
	return &CreateEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGovernancePolicy 创建治理策略
//
// 创建治理策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateGovernancePolicy(request *model.CreateGovernancePolicyRequest) (*model.CreateGovernancePolicyResponse, error) {
	requestDef := GenReqDefForCreateGovernancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGovernancePolicyResponse), nil
	}
}

// CreateGovernancePolicyInvoker 创建治理策略
func (c *CseClient) CreateGovernancePolicyInvoker(request *model.CreateGovernancePolicyRequest) *CreateGovernancePolicyInvoker {
	requestDef := GenReqDefForCreateGovernancePolicy()
	return &CreateGovernancePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMicroserviceRouteRule 创建灰度发布策略
//
// 创建灰度发布策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateMicroserviceRouteRule(request *model.CreateMicroserviceRouteRuleRequest) (*model.CreateMicroserviceRouteRuleResponse, error) {
	requestDef := GenReqDefForCreateMicroserviceRouteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMicroserviceRouteRuleResponse), nil
	}
}

// CreateMicroserviceRouteRuleInvoker 创建灰度发布策略
func (c *CseClient) CreateMicroserviceRouteRuleInvoker(request *model.CreateMicroserviceRouteRuleRequest) *CreateMicroserviceRouteRuleInvoker {
	requestDef := GenReqDefForCreateMicroserviceRouteRule()
	return &CreateMicroserviceRouteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteEngine 删除微服务引擎
//
// 删除微服务引擎。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteEngine(request *model.DeleteEngineRequest) (*model.DeleteEngineResponse, error) {
	requestDef := GenReqDefForDeleteEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteEngineResponse), nil
	}
}

// DeleteEngineInvoker 删除微服务引擎
func (c *CseClient) DeleteEngineInvoker(request *model.DeleteEngineRequest) *DeleteEngineInvoker {
	requestDef := GenReqDefForDeleteEngine()
	return &DeleteEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGovernancePolicy 删除治理策略
//
// 删除治理策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteGovernancePolicy(request *model.DeleteGovernancePolicyRequest) (*model.DeleteGovernancePolicyResponse, error) {
	requestDef := GenReqDefForDeleteGovernancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGovernancePolicyResponse), nil
	}
}

// DeleteGovernancePolicyInvoker 删除治理策略
func (c *CseClient) DeleteGovernancePolicyInvoker(request *model.DeleteGovernancePolicyRequest) *DeleteGovernancePolicyInvoker {
	requestDef := GenReqDefForDeleteGovernancePolicy()
	return &DeleteGovernancePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMicroserviceRouteRule 删除灰度发布策略
//
// 删除灰度发布策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteMicroserviceRouteRule(request *model.DeleteMicroserviceRouteRuleRequest) (*model.DeleteMicroserviceRouteRuleResponse, error) {
	requestDef := GenReqDefForDeleteMicroserviceRouteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMicroserviceRouteRuleResponse), nil
	}
}

// DeleteMicroserviceRouteRuleInvoker 删除灰度发布策略
func (c *CseClient) DeleteMicroserviceRouteRuleInvoker(request *model.DeleteMicroserviceRouteRuleRequest) *DeleteMicroserviceRouteRuleInvoker {
	requestDef := GenReqDefForDeleteMicroserviceRouteRule()
	return &DeleteMicroserviceRouteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DownloadKie 导出kie配置
//
// 导出kie配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DownloadKie(request *model.DownloadKieRequest) (*model.DownloadKieResponse, error) {
	requestDef := GenReqDefForDownloadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DownloadKieResponse), nil
	}
}

// DownloadKieInvoker 导出kie配置
func (c *CseClient) DownloadKieInvoker(request *model.DownloadKieRequest) *DownloadKieInvoker {
	requestDef := GenReqDefForDownloadKie()
	return &DownloadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEngines 查询微服务引擎列表
//
// 查询微服务引擎列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListEngines(request *model.ListEnginesRequest) (*model.ListEnginesResponse, error) {
	requestDef := GenReqDefForListEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEnginesResponse), nil
	}
}

// ListEnginesInvoker 查询微服务引擎列表
func (c *CseClient) ListEnginesInvoker(request *model.ListEnginesRequest) *ListEnginesInvoker {
	requestDef := GenReqDefForListEngines()
	return &ListEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询微服务引擎的规格列表
//
// 查询微服务引擎的规格列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询微服务引擎的规格列表
func (c *CseClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGovernancePolicy 查询指定类型治理策略列表
//
// 查询指定类型治理策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListGovernancePolicy(request *model.ListGovernancePolicyRequest) (*model.ListGovernancePolicyResponse, error) {
	requestDef := GenReqDefForListGovernancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGovernancePolicyResponse), nil
	}
}

// ListGovernancePolicyInvoker 查询指定类型治理策略列表
func (c *CseClient) ListGovernancePolicyInvoker(request *model.ListGovernancePolicyRequest) *ListGovernancePolicyInvoker {
	requestDef := GenReqDefForListGovernancePolicy()
	return &ListGovernancePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGovernancePolicyByPolicyId 查询治理策略详情
//
// 查询治理策略详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListGovernancePolicyByPolicyId(request *model.ListGovernancePolicyByPolicyIdRequest) (*model.ListGovernancePolicyByPolicyIdResponse, error) {
	requestDef := GenReqDefForListGovernancePolicyByPolicyId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGovernancePolicyByPolicyIdResponse), nil
	}
}

// ListGovernancePolicyByPolicyIdInvoker 查询治理策略详情
func (c *CseClient) ListGovernancePolicyByPolicyIdInvoker(request *model.ListGovernancePolicyByPolicyIdRequest) *ListGovernancePolicyByPolicyIdInvoker {
	requestDef := GenReqDefForListGovernancePolicyByPolicyId()
	return &ListGovernancePolicyByPolicyIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGovernancePolicys 查询治理策略列表
//
// 查询治理策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListGovernancePolicys(request *model.ListGovernancePolicysRequest) (*model.ListGovernancePolicysResponse, error) {
	requestDef := GenReqDefForListGovernancePolicys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGovernancePolicysResponse), nil
	}
}

// ListGovernancePolicysInvoker 查询治理策略列表
func (c *CseClient) ListGovernancePolicysInvoker(request *model.ListGovernancePolicysRequest) *ListGovernancePolicysInvoker {
	requestDef := GenReqDefForListGovernancePolicys()
	return &ListGovernancePolicysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMicroserviceRouteRule 查询微服务的灰度发布规则
//
// 查询微服务的灰度发布规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListMicroserviceRouteRule(request *model.ListMicroserviceRouteRuleRequest) (*model.ListMicroserviceRouteRuleResponse, error) {
	requestDef := GenReqDefForListMicroserviceRouteRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMicroserviceRouteRuleResponse), nil
	}
}

// ListMicroserviceRouteRuleInvoker 查询微服务的灰度发布规则
func (c *CseClient) ListMicroserviceRouteRuleInvoker(request *model.ListMicroserviceRouteRuleRequest) *ListMicroserviceRouteRuleInvoker {
	requestDef := GenReqDefForListMicroserviceRouteRule()
	return &ListMicroserviceRouteRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResizeEngine 变更微服务引擎规格
//
// 变更微服务引擎规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ResizeEngine(request *model.ResizeEngineRequest) (*model.ResizeEngineResponse, error) {
	requestDef := GenReqDefForResizeEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeEngineResponse), nil
	}
}

// ResizeEngineInvoker 变更微服务引擎规格
func (c *CseClient) ResizeEngineInvoker(request *model.ResizeEngineRequest) *ResizeEngineInvoker {
	requestDef := GenReqDefForResizeEngine()
	return &ResizeEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RetryEngine 对微服务引擎进行重试
//
// 对微服务引擎进行重试，当前支持ServiceComb专享版引擎
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) RetryEngine(request *model.RetryEngineRequest) (*model.RetryEngineResponse, error) {
	requestDef := GenReqDefForRetryEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RetryEngineResponse), nil
	}
}

// RetryEngineInvoker 对微服务引擎进行重试
func (c *CseClient) RetryEngineInvoker(request *model.RetryEngineRequest) *RetryEngineInvoker {
	requestDef := GenReqDefForRetryEngine()
	return &RetryEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngine 查询微服务引擎详情
//
// 查询微服务引擎详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowEngine(request *model.ShowEngineRequest) (*model.ShowEngineResponse, error) {
	requestDef := GenReqDefForShowEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineResponse), nil
	}
}

// ShowEngineInvoker 查询微服务引擎详情
func (c *CseClient) ShowEngineInvoker(request *model.ShowEngineRequest) *ShowEngineInvoker {
	requestDef := GenReqDefForShowEngine()
	return &ShowEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineJob 查询微服务引擎任务详情
//
// 查询微服务引擎任务详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowEngineJob(request *model.ShowEngineJobRequest) (*model.ShowEngineJobResponse, error) {
	requestDef := GenReqDefForShowEngineJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineJobResponse), nil
	}
}

// ShowEngineJobInvoker 查询微服务引擎任务详情
func (c *CseClient) ShowEngineJobInvoker(request *model.ShowEngineJobRequest) *ShowEngineJobInvoker {
	requestDef := GenReqDefForShowEngineJob()
	return &ShowEngineJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEngineQuotas 查询微服务引擎配额
//
// 查询微服务引擎配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowEngineQuotas(request *model.ShowEngineQuotasRequest) (*model.ShowEngineQuotasResponse, error) {
	requestDef := GenReqDefForShowEngineQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEngineQuotasResponse), nil
	}
}

// ShowEngineQuotasInvoker 查询微服务引擎配额
func (c *CseClient) ShowEngineQuotasInvoker(request *model.ShowEngineQuotasRequest) *ShowEngineQuotasInvoker {
	requestDef := GenReqDefForShowEngineQuotas()
	return &ShowEngineQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGovernancePolicy 修改治理策略
//
// 修改治理策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UpdateGovernancePolicy(request *model.UpdateGovernancePolicyRequest) (*model.UpdateGovernancePolicyResponse, error) {
	requestDef := GenReqDefForUpdateGovernancePolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGovernancePolicyResponse), nil
	}
}

// UpdateGovernancePolicyInvoker 修改治理策略
func (c *CseClient) UpdateGovernancePolicyInvoker(request *model.UpdateGovernancePolicyRequest) *UpdateGovernancePolicyInvoker {
	requestDef := GenReqDefForUpdateGovernancePolicy()
	return &UpdateGovernancePolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeEngine 升级微服务引擎
//
// 升级微服务引擎
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UpgradeEngine(request *model.UpgradeEngineRequest) (*model.UpgradeEngineResponse, error) {
	requestDef := GenReqDefForUpgradeEngine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeEngineResponse), nil
	}
}

// UpgradeEngineInvoker 升级微服务引擎
func (c *CseClient) UpgradeEngineInvoker(request *model.UpgradeEngineRequest) *UpgradeEngineInvoker {
	requestDef := GenReqDefForUpgradeEngine()
	return &UpgradeEngineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpgradeEngineConfig 更新微服务引擎配置
//
// 更新微服务引擎配置，更新ServiceComb专享版引擎与注册配置中心引擎的配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UpgradeEngineConfig(request *model.UpgradeEngineConfigRequest) (*model.UpgradeEngineConfigResponse, error) {
	requestDef := GenReqDefForUpgradeEngineConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpgradeEngineConfigResponse), nil
	}
}

// UpgradeEngineConfigInvoker 更新微服务引擎配置
func (c *CseClient) UpgradeEngineConfigInvoker(request *model.UpgradeEngineConfigRequest) *UpgradeEngineConfigInvoker {
	requestDef := GenReqDefForUpgradeEngineConfig()
	return &UpgradeEngineConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadKie 导入kie配置
//
// 导入kie配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UploadKie(request *model.UploadKieRequest) (*model.UploadKieResponse, error) {
	requestDef := GenReqDefForUploadKie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadKieResponse), nil
	}
}

// UploadKieInvoker 导入kie配置
func (c *CseClient) UploadKieInvoker(request *model.UploadKieRequest) *UploadKieInvoker {
	requestDef := GenReqDefForUploadKie()
	return &UploadKieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHttp2Rpc 创建http转rpc方法
//
// 创建http转rpc方法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateHttp2Rpc(request *model.CreateHttp2RpcRequest) (*model.CreateHttp2RpcResponse, error) {
	requestDef := GenReqDefForCreateHttp2Rpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHttp2RpcResponse), nil
	}
}

// CreateHttp2RpcInvoker 创建http转rpc方法
func (c *CseClient) CreateHttp2RpcInvoker(request *model.CreateHttp2RpcRequest) *CreateHttp2RpcInvoker {
	requestDef := GenReqDefForCreateHttp2Rpc()
	return &CreateHttp2RpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePlugin 创建插件
//
// 创建插件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreatePlugin(request *model.CreatePluginRequest) (*model.CreatePluginResponse, error) {
	requestDef := GenReqDefForCreatePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePluginResponse), nil
	}
}

// CreatePluginInvoker 创建插件
func (c *CseClient) CreatePluginInvoker(request *model.CreatePluginRequest) *CreatePluginInvoker {
	requestDef := GenReqDefForCreatePlugin()
	return &CreatePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHttp2Rpc 删除http转rpc方法
//
// 删除http转rpc方法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteHttp2Rpc(request *model.DeleteHttp2RpcRequest) (*model.DeleteHttp2RpcResponse, error) {
	requestDef := GenReqDefForDeleteHttp2Rpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHttp2RpcResponse), nil
	}
}

// DeleteHttp2RpcInvoker 删除http转rpc方法
func (c *CseClient) DeleteHttp2RpcInvoker(request *model.DeleteHttp2RpcRequest) *DeleteHttp2RpcInvoker {
	requestDef := GenReqDefForDeleteHttp2Rpc()
	return &DeleteHttp2RpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePlugin 删除插件
//
// 删除插件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeletePlugin(request *model.DeletePluginRequest) (*model.DeletePluginResponse, error) {
	requestDef := GenReqDefForDeletePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePluginResponse), nil
	}
}

// DeletePluginInvoker 删除插件
func (c *CseClient) DeletePluginInvoker(request *model.DeletePluginRequest) *DeletePluginInvoker {
	requestDef := GenReqDefForDeletePlugin()
	return &DeletePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyHttp2Rpc 修改http转rpc方法
//
// 修改http转rpc方法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ModifyHttp2Rpc(request *model.ModifyHttp2RpcRequest) (*model.ModifyHttp2RpcResponse, error) {
	requestDef := GenReqDefForModifyHttp2Rpc()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyHttp2RpcResponse), nil
	}
}

// ModifyHttp2RpcInvoker 修改http转rpc方法
func (c *CseClient) ModifyHttp2RpcInvoker(request *model.ModifyHttp2RpcRequest) *ModifyHttp2RpcInvoker {
	requestDef := GenReqDefForModifyHttp2Rpc()
	return &ModifyHttp2RpcInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyPlugin 修改插件
//
// 修改插件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ModifyPlugin(request *model.ModifyPluginRequest) (*model.ModifyPluginResponse, error) {
	requestDef := GenReqDefForModifyPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyPluginResponse), nil
	}
}

// ModifyPluginInvoker 修改插件
func (c *CseClient) ModifyPluginInvoker(request *model.ModifyPluginRequest) *ModifyPluginInvoker {
	requestDef := GenReqDefForModifyPlugin()
	return &ModifyPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowHttp2Rpcs 查询http2rpc资源列表
//
// 查询http转rpc资源列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowHttp2Rpcs(request *model.ShowHttp2RpcsRequest) (*model.ShowHttp2RpcsResponse, error) {
	requestDef := GenReqDefForShowHttp2Rpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowHttp2RpcsResponse), nil
	}
}

// ShowHttp2RpcsInvoker 查询http2rpc资源列表
func (c *CseClient) ShowHttp2RpcsInvoker(request *model.ShowHttp2RpcsRequest) *ShowHttp2RpcsInvoker {
	requestDef := GenReqDefForShowHttp2Rpcs()
	return &ShowHttp2RpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPlugins 查询插件列表
//
// 查询插件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowPlugins(request *model.ShowPluginsRequest) (*model.ShowPluginsResponse, error) {
	requestDef := GenReqDefForShowPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginsResponse), nil
	}
}

// ShowPluginsInvoker 查询插件列表
func (c *CseClient) ShowPluginsInvoker(request *model.ShowPluginsRequest) *ShowPluginsInvoker {
	requestDef := GenReqDefForShowPlugins()
	return &ShowPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSinglePlugin 查询单个插件
//
// 查询单个插件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ShowSinglePlugin(request *model.ShowSinglePluginRequest) (*model.ShowSinglePluginResponse, error) {
	requestDef := GenReqDefForShowSinglePlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSinglePluginResponse), nil
	}
}

// ShowSinglePluginInvoker 查询单个插件
func (c *CseClient) ShowSinglePluginInvoker(request *model.ShowSinglePluginRequest) *ShowSinglePluginInvoker {
	requestDef := GenReqDefForShowSinglePlugin()
	return &ShowSinglePluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNacosNamespaces 创建nacos命名空间
//
// 创建nacos命名空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) CreateNacosNamespaces(request *model.CreateNacosNamespacesRequest) (*model.CreateNacosNamespacesResponse, error) {
	requestDef := GenReqDefForCreateNacosNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNacosNamespacesResponse), nil
	}
}

// CreateNacosNamespacesInvoker 创建nacos命名空间
func (c *CseClient) CreateNacosNamespacesInvoker(request *model.CreateNacosNamespacesRequest) *CreateNacosNamespacesInvoker {
	requestDef := GenReqDefForCreateNacosNamespaces()
	return &CreateNacosNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNacosNamespaces 删除nacos命名空间
//
// 删除nacos命名空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) DeleteNacosNamespaces(request *model.DeleteNacosNamespacesRequest) (*model.DeleteNacosNamespacesResponse, error) {
	requestDef := GenReqDefForDeleteNacosNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNacosNamespacesResponse), nil
	}
}

// DeleteNacosNamespacesInvoker 删除nacos命名空间
func (c *CseClient) DeleteNacosNamespacesInvoker(request *model.DeleteNacosNamespacesRequest) *DeleteNacosNamespacesInvoker {
	requestDef := GenReqDefForDeleteNacosNamespaces()
	return &DeleteNacosNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNacosNamespaces 查询nacos命名空间
//
// 查询nacos命名空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) ListNacosNamespaces(request *model.ListNacosNamespacesRequest) (*model.ListNacosNamespacesResponse, error) {
	requestDef := GenReqDefForListNacosNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNacosNamespacesResponse), nil
	}
}

// ListNacosNamespacesInvoker 查询nacos命名空间
func (c *CseClient) ListNacosNamespacesInvoker(request *model.ListNacosNamespacesRequest) *ListNacosNamespacesInvoker {
	requestDef := GenReqDefForListNacosNamespaces()
	return &ListNacosNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNacosNamespaces 更新nacos命名空间
//
// 更新nacos命名空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CseClient) UpdateNacosNamespaces(request *model.UpdateNacosNamespacesRequest) (*model.UpdateNacosNamespacesResponse, error) {
	requestDef := GenReqDefForUpdateNacosNamespaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNacosNamespacesResponse), nil
	}
}

// UpdateNacosNamespacesInvoker 更新nacos命名空间
func (c *CseClient) UpdateNacosNamespacesInvoker(request *model.UpdateNacosNamespacesRequest) *UpdateNacosNamespacesInvoker {
	requestDef := GenReqDefForUpdateNacosNamespaces()
	return &UpdateNacosNamespacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
