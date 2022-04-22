package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
)

type RmsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewRmsClient(hcClient *http_client.HcHttpClient) *RmsClient {
	return &RmsClient{HcClient: hcClient}
}

func RmsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// 查询资源历史
//
// 查询资源与资源关系的变更历史
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceHistory(request *model.ShowResourceHistoryRequest) (*model.ShowResourceHistoryResponse, error) {
	requestDef := GenReqDefForShowResourceHistory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceHistoryResponse), nil
	}
}

// 创建合规规则
//
// 创建新的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) CreatePolicyAssignments(request *model.CreatePolicyAssignmentsRequest) (*model.CreatePolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForCreatePolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePolicyAssignmentsResponse), nil
	}
}

// 删除合规规则
//
// 根据规则ID删除此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DeletePolicyAssignment(request *model.DeletePolicyAssignmentRequest) (*model.DeletePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDeletePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePolicyAssignmentResponse), nil
	}
}

// 停用合规规则
//
// 根据规则ID停用此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DisablePolicyAssignment(request *model.DisablePolicyAssignmentRequest) (*model.DisablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForDisablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisablePolicyAssignmentResponse), nil
	}
}

// 启用合规规则
//
// 根据规则ID启用此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) EnablePolicyAssignment(request *model.EnablePolicyAssignmentRequest) (*model.EnablePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForEnablePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnablePolicyAssignmentResponse), nil
	}
}

// 列出内置策略
//
// 列出用户的内置策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListBuiltInPolicyDefinitions(request *model.ListBuiltInPolicyDefinitionsRequest) (*model.ListBuiltInPolicyDefinitionsResponse, error) {
	requestDef := GenReqDefForListBuiltInPolicyDefinitions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBuiltInPolicyDefinitionsResponse), nil
	}
}

// 列出合规规则
//
// 列出用户的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyAssignments(request *model.ListPolicyAssignmentsRequest) (*model.ListPolicyAssignmentsResponse, error) {
	requestDef := GenReqDefForListPolicyAssignments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyAssignmentsResponse), nil
	}
}

// 获取规则的合规结果
//
// 根据规则ID查询所有的合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByAssignmentId(request *model.ListPolicyStatesByAssignmentIdRequest) (*model.ListPolicyStatesByAssignmentIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByAssignmentIdResponse), nil
	}
}

// 获取用户的合规结果
//
// 查询用户所有的合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByDomainId(request *model.ListPolicyStatesByDomainIdRequest) (*model.ListPolicyStatesByDomainIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByDomainId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByDomainIdResponse), nil
	}
}

// 获取资源的合规结果
//
// 根据资源ID查询所有合规结果
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListPolicyStatesByResourceId(request *model.ListPolicyStatesByResourceIdRequest) (*model.ListPolicyStatesByResourceIdResponse, error) {
	requestDef := GenReqDefForListPolicyStatesByResourceId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPolicyStatesByResourceIdResponse), nil
	}
}

// 运行合规评估
//
// 根据规则ID评估此规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) RunEvaluationByPolicyAssignmentId(request *model.RunEvaluationByPolicyAssignmentIdRequest) (*model.RunEvaluationByPolicyAssignmentIdResponse, error) {
	requestDef := GenReqDefForRunEvaluationByPolicyAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEvaluationByPolicyAssignmentIdResponse), nil
	}
}

// 查询单个内置策略
//
// 根据策略ID查询单个内置策略
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowBuiltInPolicyDefinition(request *model.ShowBuiltInPolicyDefinitionRequest) (*model.ShowBuiltInPolicyDefinitionResponse, error) {
	requestDef := GenReqDefForShowBuiltInPolicyDefinition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBuiltInPolicyDefinitionResponse), nil
	}
}

// 获取规则的评估状态
//
// 根据规则ID查询此规则的评估状态
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowEvaluationStateByAssignmentId(request *model.ShowEvaluationStateByAssignmentIdRequest) (*model.ShowEvaluationStateByAssignmentIdResponse, error) {
	requestDef := GenReqDefForShowEvaluationStateByAssignmentId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEvaluationStateByAssignmentIdResponse), nil
	}
}

// 获取单个合规规则
//
// 根据规则ID获取单个规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowPolicyAssignment(request *model.ShowPolicyAssignmentRequest) (*model.ShowPolicyAssignmentResponse, error) {
	requestDef := GenReqDefForShowPolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPolicyAssignmentResponse), nil
	}
}

// 更新合规规则
//
// 更新用户的合规规则
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) UpdatePolicyAssignment(request *model.UpdatePolicyAssignmentRequest) (*model.UpdatePolicyAssignmentResponse, error) {
	requestDef := GenReqDefForUpdatePolicyAssignment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePolicyAssignmentResponse), nil
	}
}

// 查询租户可见的区域
//
// 查询租户可见的区域
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListRegions(request *model.ListRegionsRequest) (*model.ListRegionsResponse, error) {
	requestDef := GenReqDefForListRegions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRegionsResponse), nil
	}
}

// 列举资源关系
//
// 指定资源ID，查询该资源与其他资源的关联关系，可以指定关系方向为\&quot;in\&quot; 或者\&quot;out\&quot;
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceRelations(request *model.ShowResourceRelationsRequest) (*model.ShowResourceRelationsResponse, error) {
	requestDef := GenReqDefForShowResourceRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceRelationsResponse), nil
	}
}

// 列举所有资源
//
// 返回当前租户下所有资源，需要当前用户有rms:resources:list权限。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListAllResources(request *model.ListAllResourcesRequest) (*model.ListAllResourcesResponse, error) {
	requestDef := GenReqDefForListAllResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllResourcesResponse), nil
	}
}

// 列举云服务
//
// 查询RMS支持的云服务、资源、区域列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListProviders(request *model.ListProvidersRequest) (*model.ListProvidersResponse, error) {
	requestDef := GenReqDefForListProviders()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProvidersResponse), nil
	}
}

// 列举指定类型的资源
//
// 返回当前租户下特定资源类型的资源，需要当前用户有rms:resources:list权限。比如查询云服务器，对应的RMS资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。 RMS支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ListResources(request *model.ListResourcesRequest) (*model.ListResourcesResponse, error) {
	requestDef := GenReqDefForListResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourcesResponse), nil
	}
}

// 查询单个资源
//
// 指定资源ID，返回该资源的详细信息，需要当前用户有rms:resources:get权限。比如查询云服务器，对应的RMS资源类型是ecs.cloudservers，其中provider为ecs，type为cloudservers。RMS支持的服务和资源类型参见[支持的服务和区域](https://console.huaweicloud.com/eps/#/resources/supported)。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowResourceById(request *model.ShowResourceByIdRequest) (*model.ShowResourceByIdResponse, error) {
	requestDef := GenReqDefForShowResourceById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceByIdResponse), nil
	}
}

// 创建或更新记录器
//
// 创建或更新资源记录器，只能存在一个资源记录器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) CreateTrackerConfig(request *model.CreateTrackerConfigRequest) (*model.CreateTrackerConfigResponse, error) {
	requestDef := GenReqDefForCreateTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrackerConfigResponse), nil
	}
}

// 删除记录器
//
// 删除资源记录器
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) DeleteTrackerConfig(request *model.DeleteTrackerConfigRequest) (*model.DeleteTrackerConfigResponse, error) {
	requestDef := GenReqDefForDeleteTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrackerConfigResponse), nil
	}
}

// 查询记录器
//
// 查询资源记录器的详细信息
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *RmsClient) ShowTrackerConfig(request *model.ShowTrackerConfigRequest) (*model.ShowTrackerConfigResponse, error) {
	requestDef := GenReqDefForShowTrackerConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrackerConfigResponse), nil
	}
}
