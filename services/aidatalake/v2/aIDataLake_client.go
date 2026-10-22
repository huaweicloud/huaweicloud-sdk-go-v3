package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aidatalake/v2/model"
)

type AIDataLakeClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAIDataLakeClient(hcClient *httpclient.HcHttpClient) *AIDataLakeClient {
	return &AIDataLakeClient{HcClient: hcClient}
}

func AIDataLakeClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddV2TenantImageVersion 添加自定义镜像版本
//
// 为已有镜像添加新版本。用户通过该接口为已注册的镜像添加新的版本。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) AddV2TenantImageVersion(request *model.AddV2TenantImageVersionRequest) (*model.AddV2TenantImageVersionResponse, error) {
	requestDef := GenReqDefForAddV2TenantImageVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddV2TenantImageVersionResponse), nil
	}
}

// AddV2TenantImageVersionInvoker 添加自定义镜像版本
func (c *AIDataLakeClient) AddV2TenantImageVersionInvoker(request *model.AddV2TenantImageVersionRequest) *AddV2TenantImageVersionInvoker {
	requestDef := GenReqDefForAddV2TenantImageVersion()
	return &AddV2TenantImageVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachV2AgencyPolicy 附加策略
//
// 为委托附加策略。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) AttachV2AgencyPolicy(request *model.AttachV2AgencyPolicyRequest) (*model.AttachV2AgencyPolicyResponse, error) {
	requestDef := GenReqDefForAttachV2AgencyPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachV2AgencyPolicyResponse), nil
	}
}

// AttachV2AgencyPolicyInvoker 附加策略
func (c *AIDataLakeClient) AttachV2AgencyPolicyInvoker(request *model.AttachV2AgencyPolicyRequest) *AttachV2AgencyPolicyInvoker {
	requestDef := GenReqDefForAttachV2AgencyPolicy()
	return &AttachV2AgencyPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2Agency 创建服务委托
//
// 创建服务委托，将指定权限策略授权给AI DataLake服务。用户可通过传入支持的委托权限策略，选择性地授予相应权限。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2Agency(request *model.CreateV2AgencyRequest) (*model.CreateV2AgencyResponse, error) {
	requestDef := GenReqDefForCreateV2Agency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2AgencyResponse), nil
	}
}

// CreateV2AgencyInvoker 创建服务委托
func (c *AIDataLakeClient) CreateV2AgencyInvoker(request *model.CreateV2AgencyRequest) *CreateV2AgencyInvoker {
	requestDef := GenReqDefForCreateV2Agency()
	return &CreateV2AgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2Compute 创建计算资源
//
// 创建计算资源。此接口为异步接口，可通过[查询计算资源操作记录列表](ListV2ComputeVersions.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2Compute(request *model.CreateV2ComputeRequest) (*model.CreateV2ComputeResponse, error) {
	requestDef := GenReqDefForCreateV2Compute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2ComputeResponse), nil
	}
}

// CreateV2ComputeInvoker 创建计算资源
func (c *AIDataLakeClient) CreateV2ComputeInvoker(request *model.CreateV2ComputeRequest) *CreateV2ComputeInvoker {
	requestDef := GenReqDefForCreateV2Compute()
	return &CreateV2ComputeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2Endpoint 创建端点
//
// 在工作空间下创建端点。此接口为异步接口，可通过[查询端点详情](ShowV2Endpoint.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2Endpoint(request *model.CreateV2EndpointRequest) (*model.CreateV2EndpointResponse, error) {
	requestDef := GenReqDefForCreateV2Endpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2EndpointResponse), nil
	}
}

// CreateV2EndpointInvoker 创建端点
func (c *AIDataLakeClient) CreateV2EndpointInvoker(request *model.CreateV2EndpointRequest) *CreateV2EndpointInvoker {
	requestDef := GenReqDefForCreateV2Endpoint()
	return &CreateV2EndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2Network 创建网络配置
//
// 创建网络配置，该工作空间下的资源都将使用该网络配置。此接口是同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2Network(request *model.CreateV2NetworkRequest) (*model.CreateV2NetworkResponse, error) {
	requestDef := GenReqDefForCreateV2Network()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2NetworkResponse), nil
	}
}

// CreateV2NetworkInvoker 创建网络配置
func (c *AIDataLakeClient) CreateV2NetworkInvoker(request *model.CreateV2NetworkRequest) *CreateV2NetworkInvoker {
	requestDef := GenReqDefForCreateV2Network()
	return &CreateV2NetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2TenantImageSnapshot 创建镜像快照
//
// 为指定镜像版本创建快照。此接口为异步接口，一个镜像版本最多创建一个快照，创建后状态为Creating，可通过[查询镜像快照列表](ListV2TenantImageSnapshots.xml)获得镜像快照创建结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2TenantImageSnapshot(request *model.CreateV2TenantImageSnapshotRequest) (*model.CreateV2TenantImageSnapshotResponse, error) {
	requestDef := GenReqDefForCreateV2TenantImageSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2TenantImageSnapshotResponse), nil
	}
}

// CreateV2TenantImageSnapshotInvoker 创建镜像快照
func (c *AIDataLakeClient) CreateV2TenantImageSnapshotInvoker(request *model.CreateV2TenantImageSnapshotRequest) *CreateV2TenantImageSnapshotInvoker {
	requestDef := GenReqDefForCreateV2TenantImageSnapshot()
	return &CreateV2TenantImageSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2Workspace 创建工作空间
//
// 创建工作空间。创建工作空间接口存在同名校验，不支持创建重复名称的工作空间，并且同一账户下最多可创建5个工作空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2Workspace(request *model.CreateV2WorkspaceRequest) (*model.CreateV2WorkspaceResponse, error) {
	requestDef := GenReqDefForCreateV2Workspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2WorkspaceResponse), nil
	}
}

// CreateV2WorkspaceInvoker 创建工作空间
func (c *AIDataLakeClient) CreateV2WorkspaceInvoker(request *model.CreateV2WorkspaceRequest) *CreateV2WorkspaceInvoker {
	requestDef := GenReqDefForCreateV2Workspace()
	return &CreateV2WorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2Agency 删除服务委托
//
// 取消授权给AI DataLake服务的委托。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2Agency(request *model.DeleteV2AgencyRequest) (*model.DeleteV2AgencyResponse, error) {
	requestDef := GenReqDefForDeleteV2Agency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2AgencyResponse), nil
	}
}

// DeleteV2AgencyInvoker 删除服务委托
func (c *AIDataLakeClient) DeleteV2AgencyInvoker(request *model.DeleteV2AgencyRequest) *DeleteV2AgencyInvoker {
	requestDef := GenReqDefForDeleteV2Agency()
	return &DeleteV2AgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2Agreement 删除用户签署协议
//
// 删除用户签署的协议。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2Agreement(request *model.DeleteV2AgreementRequest) (*model.DeleteV2AgreementResponse, error) {
	requestDef := GenReqDefForDeleteV2Agreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2AgreementResponse), nil
	}
}

// DeleteV2AgreementInvoker 删除用户签署协议
func (c *AIDataLakeClient) DeleteV2AgreementInvoker(request *model.DeleteV2AgreementRequest) *DeleteV2AgreementInvoker {
	requestDef := GenReqDefForDeleteV2Agreement()
	return &DeleteV2AgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2Compute 删除计算资源
//
// 删除计算资源，通过输入计算资源ID，返回操作结果。此接口为异步接口，可通过[查询计算资源操作记录列表](ListV2ComputeVersions.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2Compute(request *model.DeleteV2ComputeRequest) (*model.DeleteV2ComputeResponse, error) {
	requestDef := GenReqDefForDeleteV2Compute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2ComputeResponse), nil
	}
}

// DeleteV2ComputeInvoker 删除计算资源
func (c *AIDataLakeClient) DeleteV2ComputeInvoker(request *model.DeleteV2ComputeRequest) *DeleteV2ComputeInvoker {
	requestDef := GenReqDefForDeleteV2Compute()
	return &DeleteV2ComputeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2Endpoint 删除端点
//
// 删除端点。此接口为异步接口，可通过[查询端点详情](ShowV2Endpoint.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2Endpoint(request *model.DeleteV2EndpointRequest) (*model.DeleteV2EndpointResponse, error) {
	requestDef := GenReqDefForDeleteV2Endpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2EndpointResponse), nil
	}
}

// DeleteV2EndpointInvoker 删除端点
func (c *AIDataLakeClient) DeleteV2EndpointInvoker(request *model.DeleteV2EndpointRequest) *DeleteV2EndpointInvoker {
	requestDef := GenReqDefForDeleteV2Endpoint()
	return &DeleteV2EndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2TenantImage 删除镜像下所有版本
//
// 删除指定镜像下的所有版本。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2TenantImage(request *model.DeleteV2TenantImageRequest) (*model.DeleteV2TenantImageResponse, error) {
	requestDef := GenReqDefForDeleteV2TenantImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2TenantImageResponse), nil
	}
}

// DeleteV2TenantImageInvoker 删除镜像下所有版本
func (c *AIDataLakeClient) DeleteV2TenantImageInvoker(request *model.DeleteV2TenantImageRequest) *DeleteV2TenantImageInvoker {
	requestDef := GenReqDefForDeleteV2TenantImage()
	return &DeleteV2TenantImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2TenantImageSnapshot 删除镜像快照
//
// 删除指定的镜像快照。创建中的快照不允许删除。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2TenantImageSnapshot(request *model.DeleteV2TenantImageSnapshotRequest) (*model.DeleteV2TenantImageSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteV2TenantImageSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2TenantImageSnapshotResponse), nil
	}
}

// DeleteV2TenantImageSnapshotInvoker 删除镜像快照
func (c *AIDataLakeClient) DeleteV2TenantImageSnapshotInvoker(request *model.DeleteV2TenantImageSnapshotRequest) *DeleteV2TenantImageSnapshotInvoker {
	requestDef := GenReqDefForDeleteV2TenantImageSnapshot()
	return &DeleteV2TenantImageSnapshotInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2TenantImageVersion 删除指定镜像版本
//
// 删除指定镜像版本。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2TenantImageVersion(request *model.DeleteV2TenantImageVersionRequest) (*model.DeleteV2TenantImageVersionResponse, error) {
	requestDef := GenReqDefForDeleteV2TenantImageVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2TenantImageVersionResponse), nil
	}
}

// DeleteV2TenantImageVersionInvoker 删除指定镜像版本
func (c *AIDataLakeClient) DeleteV2TenantImageVersionInvoker(request *model.DeleteV2TenantImageVersionRequest) *DeleteV2TenantImageVersionInvoker {
	requestDef := GenReqDefForDeleteV2TenantImageVersion()
	return &DeleteV2TenantImageVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2Workspace 删除工作空间
//
// 删除工作空间。此接口为同步接口，无配套使用接口。在该工作空间下存在计算资源、端点以及App的情况下无法删除，需要先清除该工作空间下创建的各种资源后才能删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2Workspace(request *model.DeleteV2WorkspaceRequest) (*model.DeleteV2WorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteV2Workspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2WorkspaceResponse), nil
	}
}

// DeleteV2WorkspaceInvoker 删除工作空间
func (c *AIDataLakeClient) DeleteV2WorkspaceInvoker(request *model.DeleteV2WorkspaceRequest) *DeleteV2WorkspaceInvoker {
	requestDef := GenReqDefForDeleteV2Workspace()
	return &DeleteV2WorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachV2AgencyPolicy 移除策略
//
// 为委托移除关联策略。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DetachV2AgencyPolicy(request *model.DetachV2AgencyPolicyRequest) (*model.DetachV2AgencyPolicyResponse, error) {
	requestDef := GenReqDefForDetachV2AgencyPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachV2AgencyPolicyResponse), nil
	}
}

// DetachV2AgencyPolicyInvoker 移除策略
func (c *AIDataLakeClient) DetachV2AgencyPolicyInvoker(request *model.DetachV2AgencyPolicyRequest) *DetachV2AgencyPolicyInvoker {
	requestDef := GenReqDefForDetachV2AgencyPolicy()
	return &DetachV2AgencyPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2AgreementRule 查询系统协议
//
// 查询系统内置的租户协议内容列表，用户可从中了解协议名称、版本等信息。该接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2AgreementRule(request *model.ListV2AgreementRuleRequest) (*model.ListV2AgreementRuleResponse, error) {
	requestDef := GenReqDefForListV2AgreementRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2AgreementRuleResponse), nil
	}
}

// ListV2AgreementRuleInvoker 查询系统协议
func (c *AIDataLakeClient) ListV2AgreementRuleInvoker(request *model.ListV2AgreementRuleRequest) *ListV2AgreementRuleInvoker {
	requestDef := GenReqDefForListV2AgreementRule()
	return &ListV2AgreementRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2AllEndpoints 查询所有端点列表
//
// 列举该租户所有端点列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2AllEndpoints(request *model.ListV2AllEndpointsRequest) (*model.ListV2AllEndpointsResponse, error) {
	requestDef := GenReqDefForListV2AllEndpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2AllEndpointsResponse), nil
	}
}

// ListV2AllEndpointsInvoker 查询所有端点列表
func (c *AIDataLakeClient) ListV2AllEndpointsInvoker(request *model.ListV2AllEndpointsRequest) *ListV2AllEndpointsInvoker {
	requestDef := GenReqDefForListV2AllEndpoints()
	return &ListV2AllEndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2ComputeVersions 查询计算资源操作记录列表
//
// 查询计算资源操作记录列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2ComputeVersions(request *model.ListV2ComputeVersionsRequest) (*model.ListV2ComputeVersionsResponse, error) {
	requestDef := GenReqDefForListV2ComputeVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2ComputeVersionsResponse), nil
	}
}

// ListV2ComputeVersionsInvoker 查询计算资源操作记录列表
func (c *AIDataLakeClient) ListV2ComputeVersionsInvoker(request *model.ListV2ComputeVersionsRequest) *ListV2ComputeVersionsInvoker {
	requestDef := GenReqDefForListV2ComputeVersions()
	return &ListV2ComputeVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2Computes 查询计算资源列表
//
// 查询工作空间下的计算资源列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2Computes(request *model.ListV2ComputesRequest) (*model.ListV2ComputesResponse, error) {
	requestDef := GenReqDefForListV2Computes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2ComputesResponse), nil
	}
}

// ListV2ComputesInvoker 查询计算资源列表
func (c *AIDataLakeClient) ListV2ComputesInvoker(request *model.ListV2ComputesRequest) *ListV2ComputesInvoker {
	requestDef := GenReqDefForListV2Computes()
	return &ListV2ComputesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2Endpoints 查询端点列表
//
// 列举工作空间下的端点列表。用户可使用该接口列举指定工作空间下的端点列表，支持输入名称、id等参数过滤，支持分页查询。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2Endpoints(request *model.ListV2EndpointsRequest) (*model.ListV2EndpointsResponse, error) {
	requestDef := GenReqDefForListV2Endpoints()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2EndpointsResponse), nil
	}
}

// ListV2EndpointsInvoker 查询端点列表
func (c *AIDataLakeClient) ListV2EndpointsInvoker(request *model.ListV2EndpointsRequest) *ListV2EndpointsInvoker {
	requestDef := GenReqDefForListV2Endpoints()
	return &ListV2EndpointsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2Networks 查询网络配置列表
//
// 查询工作空间下的网络配置列表。接口支持分页查询。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2Networks(request *model.ListV2NetworksRequest) (*model.ListV2NetworksResponse, error) {
	requestDef := GenReqDefForListV2Networks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2NetworksResponse), nil
	}
}

// ListV2NetworksInvoker 查询网络配置列表
func (c *AIDataLakeClient) ListV2NetworksInvoker(request *model.ListV2NetworksRequest) *ListV2NetworksInvoker {
	requestDef := GenReqDefForListV2Networks()
	return &ListV2NetworksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2TenantImageSnapshots 查询镜像快照列表
//
// 查询当前租户下的全部镜像快照列表。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2TenantImageSnapshots(request *model.ListV2TenantImageSnapshotsRequest) (*model.ListV2TenantImageSnapshotsResponse, error) {
	requestDef := GenReqDefForListV2TenantImageSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2TenantImageSnapshotsResponse), nil
	}
}

// ListV2TenantImageSnapshotsInvoker 查询镜像快照列表
func (c *AIDataLakeClient) ListV2TenantImageSnapshotsInvoker(request *model.ListV2TenantImageSnapshotsRequest) *ListV2TenantImageSnapshotsInvoker {
	requestDef := GenReqDefForListV2TenantImageSnapshots()
	return &ListV2TenantImageSnapshotsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2TenantImageVersions 查询自定义镜像版本列表
//
// 查询自定义镜像版本列表。用户可查询指定镜像下的所有版本信息。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2TenantImageVersions(request *model.ListV2TenantImageVersionsRequest) (*model.ListV2TenantImageVersionsResponse, error) {
	requestDef := GenReqDefForListV2TenantImageVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2TenantImageVersionsResponse), nil
	}
}

// ListV2TenantImageVersionsInvoker 查询自定义镜像版本列表
func (c *AIDataLakeClient) ListV2TenantImageVersionsInvoker(request *model.ListV2TenantImageVersionsRequest) *ListV2TenantImageVersionsInvoker {
	requestDef := GenReqDefForListV2TenantImageVersions()
	return &ListV2TenantImageVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2TenantImages 查询自定义镜像列表。
//
// 查询自定义镜像列表。用户可查询已创建的自定义镜像列表。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2TenantImages(request *model.ListV2TenantImagesRequest) (*model.ListV2TenantImagesResponse, error) {
	requestDef := GenReqDefForListV2TenantImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2TenantImagesResponse), nil
	}
}

// ListV2TenantImagesInvoker 查询自定义镜像列表。
func (c *AIDataLakeClient) ListV2TenantImagesInvoker(request *model.ListV2TenantImagesRequest) *ListV2TenantImagesInvoker {
	requestDef := GenReqDefForListV2TenantImages()
	return &ListV2TenantImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2Workspaces 查询工作空间列表
//
// 查询工作空间列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2Workspaces(request *model.ListV2WorkspacesRequest) (*model.ListV2WorkspacesResponse, error) {
	requestDef := GenReqDefForListV2Workspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2WorkspacesResponse), nil
	}
}

// ListV2WorkspacesInvoker 查询工作空间列表
func (c *AIDataLakeClient) ListV2WorkspacesInvoker(request *model.ListV2WorkspacesRequest) *ListV2WorkspacesInvoker {
	requestDef := GenReqDefForListV2Workspaces()
	return &ListV2WorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterV2TenantImage 注册自定义镜像
//
// 注册自定义镜像。用户通过该接口将外部镜像仓库中的镜像注册到租户中，以便后续使用。与[添加自定义镜像版本](AddV2ImageVersion.xml)接口配合完成镜像版本的添加。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RegisterV2TenantImage(request *model.RegisterV2TenantImageRequest) (*model.RegisterV2TenantImageResponse, error) {
	requestDef := GenReqDefForRegisterV2TenantImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterV2TenantImageResponse), nil
	}
}

// RegisterV2TenantImageInvoker 注册自定义镜像
func (c *AIDataLakeClient) RegisterV2TenantImageInvoker(request *model.RegisterV2TenantImageRequest) *RegisterV2TenantImageInvoker {
	requestDef := GenReqDefForRegisterV2TenantImage()
	return &RegisterV2TenantImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2Agency 查询服务委托
//
// 查询授权给AI DataLake服务的委托。用户调用该接口查询已授权给AI DataLake的委托，以及委托中具体的权限内容。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2Agency(request *model.ShowV2AgencyRequest) (*model.ShowV2AgencyResponse, error) {
	requestDef := GenReqDefForShowV2Agency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2AgencyResponse), nil
	}
}

// ShowV2AgencyInvoker 查询服务委托
func (c *AIDataLakeClient) ShowV2AgencyInvoker(request *model.ShowV2AgencyRequest) *ShowV2AgencyInvoker {
	requestDef := GenReqDefForShowV2Agency()
	return &ShowV2AgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2Agreement 查询用户是否签署协议
//
// 查询用户是否签署协议。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2Agreement(request *model.ShowV2AgreementRequest) (*model.ShowV2AgreementResponse, error) {
	requestDef := GenReqDefForShowV2Agreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2AgreementResponse), nil
	}
}

// ShowV2AgreementInvoker 查询用户是否签署协议
func (c *AIDataLakeClient) ShowV2AgreementInvoker(request *model.ShowV2AgreementRequest) *ShowV2AgreementInvoker {
	requestDef := GenReqDefForShowV2Agreement()
	return &ShowV2AgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2Compute 查看计算资源详情
//
// 查看指定计算资源的详细信息。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2Compute(request *model.ShowV2ComputeRequest) (*model.ShowV2ComputeResponse, error) {
	requestDef := GenReqDefForShowV2Compute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2ComputeResponse), nil
	}
}

// ShowV2ComputeInvoker 查看计算资源详情
func (c *AIDataLakeClient) ShowV2ComputeInvoker(request *model.ShowV2ComputeRequest) *ShowV2ComputeInvoker {
	requestDef := GenReqDefForShowV2Compute()
	return &ShowV2ComputeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2ComputesMetrics 查询资源池监控
//
// 查询资源池监控CPU、MEM、NPU、GPU数据。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2ComputesMetrics(request *model.ShowV2ComputesMetricsRequest) (*model.ShowV2ComputesMetricsResponse, error) {
	requestDef := GenReqDefForShowV2ComputesMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2ComputesMetricsResponse), nil
	}
}

// ShowV2ComputesMetricsInvoker 查询资源池监控
func (c *AIDataLakeClient) ShowV2ComputesMetricsInvoker(request *model.ShowV2ComputesMetricsRequest) *ShowV2ComputesMetricsInvoker {
	requestDef := GenReqDefForShowV2ComputesMetrics()
	return &ShowV2ComputesMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2Endpoint 查询端点详情
//
// 列举工作空间下的端点详情。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2Endpoint(request *model.ShowV2EndpointRequest) (*model.ShowV2EndpointResponse, error) {
	requestDef := GenReqDefForShowV2Endpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2EndpointResponse), nil
	}
}

// ShowV2EndpointInvoker 查询端点详情
func (c *AIDataLakeClient) ShowV2EndpointInvoker(request *model.ShowV2EndpointRequest) *ShowV2EndpointInvoker {
	requestDef := GenReqDefForShowV2Endpoint()
	return &ShowV2EndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2EndpointResourceLimit 查询端点配置限制
//
// 查询指定资源池各规格的端点配置限制。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2EndpointResourceLimit(request *model.ShowV2EndpointResourceLimitRequest) (*model.ShowV2EndpointResourceLimitResponse, error) {
	requestDef := GenReqDefForShowV2EndpointResourceLimit()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2EndpointResourceLimitResponse), nil
	}
}

// ShowV2EndpointResourceLimitInvoker 查询端点配置限制
func (c *AIDataLakeClient) ShowV2EndpointResourceLimitInvoker(request *model.ShowV2EndpointResourceLimitRequest) *ShowV2EndpointResourceLimitInvoker {
	requestDef := GenReqDefForShowV2EndpointResourceLimit()
	return &ShowV2EndpointResourceLimitInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2EndpointsMetrics 查询端点监控
//
// 查询端点CPU、MEM、NPU、GPU监控数据。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2EndpointsMetrics(request *model.ShowV2EndpointsMetricsRequest) (*model.ShowV2EndpointsMetricsResponse, error) {
	requestDef := GenReqDefForShowV2EndpointsMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2EndpointsMetricsResponse), nil
	}
}

// ShowV2EndpointsMetricsInvoker 查询端点监控
func (c *AIDataLakeClient) ShowV2EndpointsMetricsInvoker(request *model.ShowV2EndpointsMetricsRequest) *ShowV2EndpointsMetricsInvoker {
	requestDef := GenReqDefForShowV2EndpointsMetrics()
	return &ShowV2EndpointsMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2TenantImage 查询镜像详情
//
// 查询指定镜像的详细信息。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2TenantImage(request *model.ShowV2TenantImageRequest) (*model.ShowV2TenantImageResponse, error) {
	requestDef := GenReqDefForShowV2TenantImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2TenantImageResponse), nil
	}
}

// ShowV2TenantImageInvoker 查询镜像详情
func (c *AIDataLakeClient) ShowV2TenantImageInvoker(request *model.ShowV2TenantImageRequest) *ShowV2TenantImageInvoker {
	requestDef := GenReqDefForShowV2TenantImage()
	return &ShowV2TenantImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2TenantImageVersion 查询自定义镜像版本
//
// 查询自定义镜像版本。用户可查询指定镜像版本的详细信息。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2TenantImageVersion(request *model.ShowV2TenantImageVersionRequest) (*model.ShowV2TenantImageVersionResponse, error) {
	requestDef := GenReqDefForShowV2TenantImageVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2TenantImageVersionResponse), nil
	}
}

// ShowV2TenantImageVersionInvoker 查询自定义镜像版本
func (c *AIDataLakeClient) ShowV2TenantImageVersionInvoker(request *model.ShowV2TenantImageVersionRequest) *ShowV2TenantImageVersionInvoker {
	requestDef := GenReqDefForShowV2TenantImageVersion()
	return &ShowV2TenantImageVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2Workspace 查询工作空间详细信息
//
// 查询工作空间详细信息。用户可通过该接口查询某工作空间详情，输入为工作空间ID，返回对应工作空间详情。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2Workspace(request *model.ShowV2WorkspaceRequest) (*model.ShowV2WorkspaceResponse, error) {
	requestDef := GenReqDefForShowV2Workspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2WorkspaceResponse), nil
	}
}

// ShowV2WorkspaceInvoker 查询工作空间详细信息
func (c *AIDataLakeClient) ShowV2WorkspaceInvoker(request *model.ShowV2WorkspaceRequest) *ShowV2WorkspaceInvoker {
	requestDef := GenReqDefForShowV2Workspace()
	return &ShowV2WorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SignV2Agreement 签署租户协议
//
// 签署租户协议，用户通过该接口签署租户协议。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) SignV2Agreement(request *model.SignV2AgreementRequest) (*model.SignV2AgreementResponse, error) {
	requestDef := GenReqDefForSignV2Agreement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SignV2AgreementResponse), nil
	}
}

// SignV2AgreementInvoker 签署租户协议
func (c *AIDataLakeClient) SignV2AgreementInvoker(request *model.SignV2AgreementRequest) *SignV2AgreementInvoker {
	requestDef := GenReqDefForSignV2Agreement()
	return &SignV2AgreementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2Compute 更新计算资源
//
// 更新指定计算资源的配置信息。此接口为异步接口，可通过[查询计算资源操作记录列表](ListV2ComputeVersions.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateV2Compute(request *model.UpdateV2ComputeRequest) (*model.UpdateV2ComputeResponse, error) {
	requestDef := GenReqDefForUpdateV2Compute()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2ComputeResponse), nil
	}
}

// UpdateV2ComputeInvoker 更新计算资源
func (c *AIDataLakeClient) UpdateV2ComputeInvoker(request *model.UpdateV2ComputeRequest) *UpdateV2ComputeInvoker {
	requestDef := GenReqDefForUpdateV2Compute()
	return &UpdateV2ComputeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2Endpoint 更新端点
//
// 更新端点。此接口为异步接口，可通过[查询端点详情](ShowV2Endpoint.xml)接口获取操作结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateV2Endpoint(request *model.UpdateV2EndpointRequest) (*model.UpdateV2EndpointResponse, error) {
	requestDef := GenReqDefForUpdateV2Endpoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2EndpointResponse), nil
	}
}

// UpdateV2EndpointInvoker 更新端点
func (c *AIDataLakeClient) UpdateV2EndpointInvoker(request *model.UpdateV2EndpointRequest) *UpdateV2EndpointInvoker {
	requestDef := GenReqDefForUpdateV2Endpoint()
	return &UpdateV2EndpointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2TenantImageVersion 更新自定义镜像版本
//
// 更新自定义镜像版本信息。用户通过该接口更新指定镜像版本的描述、架构、容器软件、分类等信息。与[注册自定义镜像](RegisterV2Image.xml)接口配合完成镜像版本管理。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateV2TenantImageVersion(request *model.UpdateV2TenantImageVersionRequest) (*model.UpdateV2TenantImageVersionResponse, error) {
	requestDef := GenReqDefForUpdateV2TenantImageVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2TenantImageVersionResponse), nil
	}
}

// UpdateV2TenantImageVersionInvoker 更新自定义镜像版本
func (c *AIDataLakeClient) UpdateV2TenantImageVersionInvoker(request *model.UpdateV2TenantImageVersionRequest) *UpdateV2TenantImageVersionInvoker {
	requestDef := GenReqDefForUpdateV2TenantImageVersion()
	return &UpdateV2TenantImageVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2Workspace 更新工作空间
//
// 更新工作空间。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateV2Workspace(request *model.UpdateV2WorkspaceRequest) (*model.UpdateV2WorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateV2Workspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2WorkspaceResponse), nil
	}
}

// UpdateV2WorkspaceInvoker 更新工作空间
func (c *AIDataLakeClient) UpdateV2WorkspaceInvoker(request *model.UpdateV2WorkspaceRequest) *UpdateV2WorkspaceInvoker {
	requestDef := GenReqDefForUpdateV2Workspace()
	return &UpdateV2WorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAuraDataframeJob 取消Aura Dataframe Job
//
// 取消正在运行中的Dataframe Job。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelAuraDataframeJob(request *model.CancelAuraDataframeJobRequest) (*model.CancelAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForCancelAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAuraDataframeJobResponse), nil
	}
}

// CancelAuraDataframeJobInvoker 取消Aura Dataframe Job
func (c *AIDataLakeClient) CancelAuraDataframeJobInvoker(request *model.CancelAuraDataframeJobRequest) *CancelAuraDataframeJobInvoker {
	requestDef := GenReqDefForCancelAuraDataframeJob()
	return &CancelAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuraDataframeJob 创建Aura Dataframe Job
//
// 创建Dataframe Job。
// 用户通过此接口在指定端点创建Dataframe Job。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateAuraDataframeJob(request *model.CreateAuraDataframeJobRequest) (*model.CreateAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForCreateAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuraDataframeJobResponse), nil
	}
}

// CreateAuraDataframeJobInvoker 创建Aura Dataframe Job
func (c *AIDataLakeClient) CreateAuraDataframeJobInvoker(request *model.CreateAuraDataframeJobRequest) *CreateAuraDataframeJobInvoker {
	requestDef := GenReqDefForCreateAuraDataframeJob()
	return &CreateAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraDataframeJob 查看Aura Dataframe Job列表
//
// 查询Aura Dataframe Job列表。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraDataframeJob(request *model.ListAuraDataframeJobRequest) (*model.ListAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForListAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraDataframeJobResponse), nil
	}
}

// ListAuraDataframeJobInvoker 查看Aura Dataframe Job列表
func (c *AIDataLakeClient) ListAuraDataframeJobInvoker(request *model.ListAuraDataframeJobRequest) *ListAuraDataframeJobInvoker {
	requestDef := GenReqDefForListAuraDataframeJob()
	return &ListAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraDataframeJob 查看Aura Dataframe Job详情
//
// 查询Aura Dataframe Job信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraDataframeJob(request *model.ShowAuraDataframeJobRequest) (*model.ShowAuraDataframeJobResponse, error) {
	requestDef := GenReqDefForShowAuraDataframeJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraDataframeJobResponse), nil
	}
}

// ShowAuraDataframeJobInvoker 查看Aura Dataframe Job详情
func (c *AIDataLakeClient) ShowAuraDataframeJobInvoker(request *model.ShowAuraDataframeJobRequest) *ShowAuraDataframeJobInvoker {
	requestDef := GenReqDefForShowAuraDataframeJob()
	return &ShowAuraDataframeJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAuraSqlStatement 取消SQL执行
//
// 取消SQL语句执行。 用户可通过该接口取消SQL语句执行，输入为SQL Session id、statement id，返回操作结果。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelAuraSqlStatement(request *model.CancelAuraSqlStatementRequest) (*model.CancelAuraSqlStatementResponse, error) {
	requestDef := GenReqDefForCancelAuraSqlStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAuraSqlStatementResponse), nil
	}
}

// CancelAuraSqlStatementInvoker 取消SQL执行
func (c *AIDataLakeClient) CancelAuraSqlStatementInvoker(request *model.CancelAuraSqlStatementRequest) *CancelAuraSqlStatementInvoker {
	requestDef := GenReqDefForCancelAuraSqlStatement()
	return &CancelAuraSqlStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CloseAuraSqlSession 关闭Session
//
// 关闭SQL Session。
// 用户可通过此接口关闭SQL Session，输入为Session id，返回操作结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CloseAuraSqlSession(request *model.CloseAuraSqlSessionRequest) (*model.CloseAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForCloseAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseAuraSqlSessionResponse), nil
	}
}

// CloseAuraSqlSessionInvoker 关闭Session
func (c *AIDataLakeClient) CloseAuraSqlSessionInvoker(request *model.CloseAuraSqlSessionRequest) *CloseAuraSqlSessionInvoker {
	requestDef := GenReqDefForCloseAuraSqlSession()
	return &CloseAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuraSqlSession 创建Session
//
// 创建SQL Session。
// 用户通过此接口在指定端点创建SQL Session，通过输入端点id、LakeFormation配置，返回Session信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateAuraSqlSession(request *model.CreateAuraSqlSessionRequest) (*model.CreateAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForCreateAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuraSqlSessionResponse), nil
	}
}

// CreateAuraSqlSessionInvoker 创建Session
func (c *AIDataLakeClient) CreateAuraSqlSessionInvoker(request *model.CreateAuraSqlSessionRequest) *CreateAuraSqlSessionInvoker {
	requestDef := GenReqDefForCreateAuraSqlSession()
	return &CreateAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteAuraSqlStatement 在指定Session中执行SQL
//
// 在指定的Session下执行SQL语句。
// 用户可通过此接口执行SQL语句，输入为SQL语句、SQL Session id、绑定参数等信息，返回语句执行结果或statement id。
// 此接口支持异步和同步两种执行模式，由入参is_sync参数决定。如果是异步执行，配套使用[查看SQL执行结果](ShowAuraV2SqlStatementResult.xml)接口查询语句结果。
// 如果是同步执行，接口会等待一段时间接收结果，如果执行完毕会返回执行结果，如果没有执行完毕，会返回一个运行中状态，后续可以使用[查看SQL执行结果](ShowAuraV2SqlStatementResult.xml)接口查询语句结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ExecuteAuraSqlStatement(request *model.ExecuteAuraSqlStatementRequest) (*model.ExecuteAuraSqlStatementResponse, error) {
	requestDef := GenReqDefForExecuteAuraSqlStatement()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteAuraSqlStatementResponse), nil
	}
}

// ExecuteAuraSqlStatementInvoker 在指定Session中执行SQL
func (c *AIDataLakeClient) ExecuteAuraSqlStatementInvoker(request *model.ExecuteAuraSqlStatementRequest) *ExecuteAuraSqlStatementInvoker {
	requestDef := GenReqDefForExecuteAuraSqlStatement()
	return &ExecuteAuraSqlStatementInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraSessionStatementRecords 查看指定Session下的SQL执行记录
//
// 查询指定Session下的SQL执行记录。
// 输入workspace_id，session_id，statement_id（可选），status（可选），分页查询参数limit和marker；输出此会话下SQL执行记录。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraSessionStatementRecords(request *model.ListAuraSessionStatementRecordsRequest) (*model.ListAuraSessionStatementRecordsResponse, error) {
	requestDef := GenReqDefForListAuraSessionStatementRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraSessionStatementRecordsResponse), nil
	}
}

// ListAuraSessionStatementRecordsInvoker 查看指定Session下的SQL执行记录
func (c *AIDataLakeClient) ListAuraSessionStatementRecordsInvoker(request *model.ListAuraSessionStatementRecordsRequest) *ListAuraSessionStatementRecordsInvoker {
	requestDef := GenReqDefForListAuraSessionStatementRecords()
	return &ListAuraSessionStatementRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraSqlSessions 查看Session列表
//
// 查询SQL Session列表。
// 输入workspace_id，endpoint_name（可选），status（可选），session_id（可选），start_time（可选），end_time（可选）,分页查询参数limit和marker；输出SQL会话列表，包含会话id、会话创建时间、结束时间、会话状态等信息。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraSqlSessions(request *model.ListAuraSqlSessionsRequest) (*model.ListAuraSqlSessionsResponse, error) {
	requestDef := GenReqDefForListAuraSqlSessions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraSqlSessionsResponse), nil
	}
}

// ListAuraSqlSessionsInvoker 查看Session列表
func (c *AIDataLakeClient) ListAuraSqlSessionsInvoker(request *model.ListAuraSqlSessionsRequest) *ListAuraSqlSessionsInvoker {
	requestDef := GenReqDefForListAuraSqlSessions()
	return &ListAuraSqlSessionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuraStatementRecords 查看SQL执行记录
//
// 查询SQL执行记录。 查询指定Session下的SQL执行记录。 输入workspace_id，session_id，statement_id（可选），status（可选），分页查询参数limit和marker；输出此会话下SQL执行记录。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListAuraStatementRecords(request *model.ListAuraStatementRecordsRequest) (*model.ListAuraStatementRecordsResponse, error) {
	requestDef := GenReqDefForListAuraStatementRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuraStatementRecordsResponse), nil
	}
}

// ListAuraStatementRecordsInvoker 查看SQL执行记录
func (c *AIDataLakeClient) ListAuraStatementRecordsInvoker(request *model.ListAuraStatementRecordsRequest) *ListAuraStatementRecordsInvoker {
	requestDef := GenReqDefForListAuraStatementRecords()
	return &ListAuraStatementRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraSqlSession 查看Session详情
//
// 查询SQL Session信息。
// 用户可通过此接口查询SQL Session信息，输入为Session id，返回操作结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraSqlSession(request *model.ShowAuraSqlSessionRequest) (*model.ShowAuraSqlSessionResponse, error) {
	requestDef := GenReqDefForShowAuraSqlSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraSqlSessionResponse), nil
	}
}

// ShowAuraSqlSessionInvoker 查看Session详情
func (c *AIDataLakeClient) ShowAuraSqlSessionInvoker(request *model.ShowAuraSqlSessionRequest) *ShowAuraSqlSessionInvoker {
	requestDef := GenReqDefForShowAuraSqlSession()
	return &ShowAuraSqlSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraSqlStatementResult 查看SQL执行结果
//
// 查询SQL语句执行结果。
// 用户可通过该接口查询SQL语句执行结果，输入为SQL Session id、statement id，返回执行结果。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraSqlStatementResult(request *model.ShowAuraSqlStatementResultRequest) (*model.ShowAuraSqlStatementResultResponse, error) {
	requestDef := GenReqDefForShowAuraSqlStatementResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraSqlStatementResultResponse), nil
	}
}

// ShowAuraSqlStatementResultInvoker 查看SQL执行结果
func (c *AIDataLakeClient) ShowAuraSqlStatementResultInvoker(request *model.ShowAuraSqlStatementResultRequest) *ShowAuraSqlStatementResultInvoker {
	requestDef := GenReqDefForShowAuraSqlStatementResult()
	return &ShowAuraSqlStatementResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraStatementOperatorMetrics 查看SQL执行算子监控记录
//
// 查看SQL执行算子监控记录。 用户可通过该接口查看SQL执行算子监控记录，输入为SQL Session id、statement id，返回操作结果。 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraStatementOperatorMetrics(request *model.ShowAuraStatementOperatorMetricsRequest) (*model.ShowAuraStatementOperatorMetricsResponse, error) {
	requestDef := GenReqDefForShowAuraStatementOperatorMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraStatementOperatorMetricsResponse), nil
	}
}

// ShowAuraStatementOperatorMetricsInvoker 查看SQL执行算子监控记录
func (c *AIDataLakeClient) ShowAuraStatementOperatorMetricsInvoker(request *model.ShowAuraStatementOperatorMetricsRequest) *ShowAuraStatementOperatorMetricsInvoker {
	requestDef := GenReqDefForShowAuraStatementOperatorMetrics()
	return &ShowAuraStatementOperatorMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuraStatementQueryMetrics 查看语句监控详情
//
// 查询语句监控详情信息。
// 输入workspace_id，session_id，statement_id，输出此会话下SQL监控数据。
// 此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowAuraStatementQueryMetrics(request *model.ShowAuraStatementQueryMetricsRequest) (*model.ShowAuraStatementQueryMetricsResponse, error) {
	requestDef := GenReqDefForShowAuraStatementQueryMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuraStatementQueryMetricsResponse), nil
	}
}

// ShowAuraStatementQueryMetricsInvoker 查看语句监控详情
func (c *AIDataLakeClient) ShowAuraStatementQueryMetricsInvoker(request *model.ShowAuraStatementQueryMetricsRequest) *ShowAuraStatementQueryMetricsInvoker {
	requestDef := GenReqDefForShowAuraStatementQueryMetrics()
	return &ShowAuraStatementQueryMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobAutoscalingHistories 查询Flink作业自动调优历史
//
// 查询Flink作业自动调优历史记录，包含每次扩缩容操作的详细信息。此接口为同步接口，支持分页查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobAutoscalingHistories(request *model.ListFlinkJobAutoscalingHistoriesRequest) (*model.ListFlinkJobAutoscalingHistoriesResponse, error) {
	requestDef := GenReqDefForListFlinkJobAutoscalingHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobAutoscalingHistoriesResponse), nil
	}
}

// ListFlinkJobAutoscalingHistoriesInvoker 查询Flink作业自动调优历史
func (c *AIDataLakeClient) ListFlinkJobAutoscalingHistoriesInvoker(request *model.ListFlinkJobAutoscalingHistoriesRequest) *ListFlinkJobAutoscalingHistoriesInvoker {
	requestDef := GenReqDefForListFlinkJobAutoscalingHistories()
	return &ListFlinkJobAutoscalingHistoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJobAutoscalingPolicy 查询Flink作业自动调优配置
//
// 查询Flink作业当前的自动调优配置。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJobAutoscalingPolicy(request *model.ShowFlinkJobAutoscalingPolicyRequest) (*model.ShowFlinkJobAutoscalingPolicyResponse, error) {
	requestDef := GenReqDefForShowFlinkJobAutoscalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobAutoscalingPolicyResponse), nil
	}
}

// ShowFlinkJobAutoscalingPolicyInvoker 查询Flink作业自动调优配置
func (c *AIDataLakeClient) ShowFlinkJobAutoscalingPolicyInvoker(request *model.ShowFlinkJobAutoscalingPolicyRequest) *ShowFlinkJobAutoscalingPolicyInvoker {
	requestDef := GenReqDefForShowFlinkJobAutoscalingPolicy()
	return &ShowFlinkJobAutoscalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFlinkJobAutoscalingPolicy 热更新Flink作业自动调优配置
//
// 热更新Flink作业自动调优配置。此接口为异步接口，作业运行中时配置将自动应用，无需重启作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateFlinkJobAutoscalingPolicy(request *model.UpdateFlinkJobAutoscalingPolicyRequest) (*model.UpdateFlinkJobAutoscalingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateFlinkJobAutoscalingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFlinkJobAutoscalingPolicyResponse), nil
	}
}

// UpdateFlinkJobAutoscalingPolicyInvoker 热更新Flink作业自动调优配置
func (c *AIDataLakeClient) UpdateFlinkJobAutoscalingPolicyInvoker(request *model.UpdateFlinkJobAutoscalingPolicyRequest) *UpdateFlinkJobAutoscalingPolicyInvoker {
	requestDef := GenReqDefForUpdateFlinkJobAutoscalingPolicy()
	return &UpdateFlinkJobAutoscalingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelFlinkJob 取消Flink作业执行
//
// 取消Flink作业。此接口为异步接口，响应仅表示取消请求已接收，作业实际停止状态可通过[查询Flink作业状态](ShowFlinkJobState.xml)接口查询。
// 支持多种停止策略：
// - NONE：立即停止不保存状态。
// - STOP_WITH_SAVEPOINT：停止前创建快照。
// - STOP_WITH_DRAIN：Drain模式停止，触发窗口计算后再创建快照停止。
// - FORCE_CANCEL：强制取消，仅当作业已处于CANCELLING状态时可再次触发取消。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelFlinkJob(request *model.CancelFlinkJobRequest) (*model.CancelFlinkJobResponse, error) {
	requestDef := GenReqDefForCancelFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelFlinkJobResponse), nil
	}
}

// CancelFlinkJobInvoker 取消Flink作业执行
func (c *AIDataLakeClient) CancelFlinkJobInvoker(request *model.CancelFlinkJobRequest) *CancelFlinkJobInvoker {
	requestDef := GenReqDefForCancelFlinkJob()
	return &CancelFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobEvents 查询Flink作业事件详情列表
//
// 查询Flink作业的运行事件列表。此接口为同步接口，响应中包含作业运行期间产生的事件详情，可用于问题定位和运行状态追踪。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取作业完整信息，或通过[查询Flink作业状态](ShowFlinkJobState.xml)接口获取作业当前运行状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobEvents(request *model.ListFlinkJobEventsRequest) (*model.ListFlinkJobEventsResponse, error) {
	requestDef := GenReqDefForListFlinkJobEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobEventsResponse), nil
	}
}

// ListFlinkJobEventsInvoker 查询Flink作业事件详情列表
func (c *AIDataLakeClient) ListFlinkJobEventsInvoker(request *model.ListFlinkJobEventsRequest) *ListFlinkJobEventsInvoker {
	requestDef := GenReqDefForListFlinkJobEvents()
	return &ListFlinkJobEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkJobs 查询Flink作业列表
//
// 查询工作空间下Flink作业列表。此接口为同步接口，响应中包含作业的概要信息，支持分页查询和多条件筛选。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取指定作业的完整信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkJobs(request *model.ListFlinkJobsRequest) (*model.ListFlinkJobsResponse, error) {
	requestDef := GenReqDefForListFlinkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkJobsResponse), nil
	}
}

// ListFlinkJobsInvoker 查询Flink作业列表
func (c *AIDataLakeClient) ListFlinkJobsInvoker(request *model.ListFlinkJobsRequest) *ListFlinkJobsInvoker {
	requestDef := GenReqDefForListFlinkJobs()
	return &ListFlinkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunFlinkJob 创建Flink作业
//
// 创建并启动Flink作业实例。此接口为异步接口，响应仅表示作业创建请求已接收，作业实际运行状态可通过[查询Flink作业状态](ShowFlinkJobState.xml)接口查询。
// 支持多种作业类型：
// - FLINK_SQL_JOB：Flink SQL作业，通过SQL脚本定义数据处理逻辑。
// - FLINK_JAR_JOB：Flink JAR作业，通过自定义JAR包定义数据处理逻辑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunFlinkJob(request *model.RunFlinkJobRequest) (*model.RunFlinkJobResponse, error) {
	requestDef := GenReqDefForRunFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunFlinkJobResponse), nil
	}
}

// RunFlinkJobInvoker 创建Flink作业
func (c *AIDataLakeClient) RunFlinkJobInvoker(request *model.RunFlinkJobRequest) *RunFlinkJobInvoker {
	requestDef := GenReqDefForRunFlinkJob()
	return &RunFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJob 查看Flink作业详情
//
// 获取Flink作业实例的详细信息。此接口为同步接口，响应中包含作业的完整配置和运行状态信息。
// 可通过[查询Flink作业事件详情列表](ListFlinkJobEvents.xml)接口获取作业运行事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJob(request *model.ShowFlinkJobRequest) (*model.ShowFlinkJobResponse, error) {
	requestDef := GenReqDefForShowFlinkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobResponse), nil
	}
}

// ShowFlinkJobInvoker 查看Flink作业详情
func (c *AIDataLakeClient) ShowFlinkJobInvoker(request *model.ShowFlinkJobRequest) *ShowFlinkJobInvoker {
	requestDef := GenReqDefForShowFlinkJob()
	return &ShowFlinkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkJobState 查询Flink作业状态
//
// 查询Flink作业的运行状态。此接口为同步接口，响应中包含作业当前状态。
// 可通过[查看Flink作业详情](ShowFlinkJob.xml)接口获取作业完整信息，或通过[查询Flink作业事件详情列表](ListFlinkJobEvents.xml)接口获取作业运行事件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkJobState(request *model.ShowFlinkJobStateRequest) (*model.ShowFlinkJobStateResponse, error) {
	requestDef := GenReqDefForShowFlinkJobState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkJobStateResponse), nil
	}
}

// ShowFlinkJobStateInvoker 查询Flink作业状态
func (c *AIDataLakeClient) ShowFlinkJobStateInvoker(request *model.ShowFlinkJobStateRequest) *ShowFlinkJobStateInvoker {
	requestDef := GenReqDefForShowFlinkJobState()
	return &ShowFlinkJobStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFlinkSavepoint 创建Flink作业Savepoint
//
// 为运行中的Flink作业触发Savepoint，用于保存作业当前状态。此接口为异步接口，响应仅表示快照创建请求已接收，快照实际生成状态可通过[查询Flink作业Savepoint列表](ListFlinkSavepoint.xml)接口查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateFlinkSavepoint(request *model.CreateFlinkSavepointRequest) (*model.CreateFlinkSavepointResponse, error) {
	requestDef := GenReqDefForCreateFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFlinkSavepointResponse), nil
	}
}

// CreateFlinkSavepointInvoker 创建Flink作业Savepoint
func (c *AIDataLakeClient) CreateFlinkSavepointInvoker(request *model.CreateFlinkSavepointRequest) *CreateFlinkSavepointInvoker {
	requestDef := GenReqDefForCreateFlinkSavepoint()
	return &CreateFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteFlinkSavepoint 删除Flink作业Savepoint
//
// 删除指定Savepoint的元数据记录，不影响OBS中的实际Savepoint数据。该接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteFlinkSavepoint(request *model.DeleteFlinkSavepointRequest) (*model.DeleteFlinkSavepointResponse, error) {
	requestDef := GenReqDefForDeleteFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteFlinkSavepointResponse), nil
	}
}

// DeleteFlinkSavepointInvoker 删除Flink作业Savepoint
func (c *AIDataLakeClient) DeleteFlinkSavepointInvoker(request *model.DeleteFlinkSavepointRequest) *DeleteFlinkSavepointInvoker {
	requestDef := GenReqDefForDeleteFlinkSavepoint()
	return &DeleteFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlinkSavepoint 查询Flink作业Savepoint列表
//
// 查询指定Flink作业的Savepoint记录，返回分页列表。此接口为同步接口，响应中包含Savepoint的详细信息。
// 可通过[查询Flink作业Savepoint详情](ShowFlinkSavepoint.xml)接口获取指定Savepoint的详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListFlinkSavepoint(request *model.ListFlinkSavepointRequest) (*model.ListFlinkSavepointResponse, error) {
	requestDef := GenReqDefForListFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlinkSavepointResponse), nil
	}
}

// ListFlinkSavepointInvoker 查询Flink作业Savepoint列表
func (c *AIDataLakeClient) ListFlinkSavepointInvoker(request *model.ListFlinkSavepointRequest) *ListFlinkSavepointInvoker {
	requestDef := GenReqDefForListFlinkSavepoint()
	return &ListFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFlinkSavepoint 查询Flink作业Savepoint详情
//
// 获取指定Savepoint的详细信息。此接口为同步接口，响应中包含Savepoint的完整状态和配置信息。
// 可通过[查询Flink作业Savepoint列表](ListFlinkSavepoint.xml)接口获取Savepoint列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowFlinkSavepoint(request *model.ShowFlinkSavepointRequest) (*model.ShowFlinkSavepointResponse, error) {
	requestDef := GenReqDefForShowFlinkSavepoint()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFlinkSavepointResponse), nil
	}
}

// ShowFlinkSavepointInvoker 查询Flink作业Savepoint详情
func (c *AIDataLakeClient) ShowFlinkSavepointInvoker(request *model.ShowFlinkSavepointRequest) *ShowFlinkSavepointInvoker {
	requestDef := GenReqDefForShowFlinkSavepoint()
	return &ShowFlinkSavepointInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateV2RayFederation 创建Ray联邦端点
//
// 创建Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateV2RayFederation(request *model.CreateV2RayFederationRequest) (*model.CreateV2RayFederationResponse, error) {
	requestDef := GenReqDefForCreateV2RayFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateV2RayFederationResponse), nil
	}
}

// CreateV2RayFederationInvoker 创建Ray联邦端点
func (c *AIDataLakeClient) CreateV2RayFederationInvoker(request *model.CreateV2RayFederationRequest) *CreateV2RayFederationInvoker {
	requestDef := GenReqDefForCreateV2RayFederation()
	return &CreateV2RayFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteV2RayFederation 删除Ray联邦端点
//
// 删除Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) DeleteV2RayFederation(request *model.DeleteV2RayFederationRequest) (*model.DeleteV2RayFederationResponse, error) {
	requestDef := GenReqDefForDeleteV2RayFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteV2RayFederationResponse), nil
	}
}

// DeleteV2RayFederationInvoker 删除Ray联邦端点
func (c *AIDataLakeClient) DeleteV2RayFederationInvoker(request *model.DeleteV2RayFederationRequest) *DeleteV2RayFederationInvoker {
	requestDef := GenReqDefForDeleteV2RayFederation()
	return &DeleteV2RayFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2RayFederations 查询Ray联邦端点列表
//
// 查询工作空间下的Ray联邦端点列表。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2RayFederations(request *model.ListV2RayFederationsRequest) (*model.ListV2RayFederationsResponse, error) {
	requestDef := GenReqDefForListV2RayFederations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2RayFederationsResponse), nil
	}
}

// ListV2RayFederationsInvoker 查询Ray联邦端点列表
func (c *AIDataLakeClient) ListV2RayFederationsInvoker(request *model.ListV2RayFederationsRequest) *ListV2RayFederationsInvoker {
	requestDef := GenReqDefForListV2RayFederations()
	return &ListV2RayFederationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowV2RayFederation 查询Ray联邦端点详情
//
// 查询Ray联邦端点详情。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowV2RayFederation(request *model.ShowV2RayFederationRequest) (*model.ShowV2RayFederationResponse, error) {
	requestDef := GenReqDefForShowV2RayFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowV2RayFederationResponse), nil
	}
}

// ShowV2RayFederationInvoker 查询Ray联邦端点详情
func (c *AIDataLakeClient) ShowV2RayFederationInvoker(request *model.ShowV2RayFederationRequest) *ShowV2RayFederationInvoker {
	requestDef := GenReqDefForShowV2RayFederation()
	return &ShowV2RayFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateV2RayFederation 更新Ray联邦端点
//
// 更新Ray联邦端点。此接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) UpdateV2RayFederation(request *model.UpdateV2RayFederationRequest) (*model.UpdateV2RayFederationResponse, error) {
	requestDef := GenReqDefForUpdateV2RayFederation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateV2RayFederationResponse), nil
	}
}

// UpdateV2RayFederationInvoker 更新Ray联邦端点
func (c *AIDataLakeClient) UpdateV2RayFederationInvoker(request *model.UpdateV2RayFederationRequest) *UpdateV2RayFederationInvoker {
	requestDef := GenReqDefForUpdateV2RayFederation()
	return &UpdateV2RayFederationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelRayJob 取消作业运行
//
// 取消作业运行。主要使用于取消运行Ray Job场景。本接口为异步接口，当前取消作业运行请求下发成功后会返回job_id，此时取消作业运行并没有立即完成，需要通过调用[查看Ray作业详情](ShowRayJob.xml)查询Job状态，当Job状态为CANCELED时代表取消作业运行成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelRayJob(request *model.CancelRayJobRequest) (*model.CancelRayJobResponse, error) {
	requestDef := GenReqDefForCancelRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelRayJobResponse), nil
	}
}

// CancelRayJobInvoker 取消作业运行
func (c *AIDataLakeClient) CancelRayJobInvoker(request *model.CancelRayJobRequest) *CancelRayJobInvoker {
	requestDef := GenReqDefForCancelRayJob()
	return &CancelRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRayJobs 查询Ray作业列表
//
// 列举所有Ray作业。列举工作空间下的作业，分页返回。本接口为同步接口，当前列举所有Ray作业请求下发后会返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListRayJobs(request *model.ListRayJobsRequest) (*model.ListRayJobsResponse, error) {
	requestDef := GenReqDefForListRayJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRayJobsResponse), nil
	}
}

// ListRayJobsInvoker 查询Ray作业列表
func (c *AIDataLakeClient) ListRayJobsInvoker(request *model.ListRayJobsRequest) *ListRayJobsInvoker {
	requestDef := GenReqDefForListRayJobs()
	return &ListRayJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunRayJob 运行作业
//
// 运行Ray作业。本接口为异步接口，当前运行作业请求下发成功后会返回job_id，此时运行作业并没有立即完成，需要通过调用[查看Ray作业详情](ShowRayJob.xml)查询Job状态，当Job状态为SUCCEEDED时代表运行作业成功。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunRayJob(request *model.RunRayJobRequest) (*model.RunRayJobResponse, error) {
	requestDef := GenReqDefForRunRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunRayJobResponse), nil
	}
}

// RunRayJobInvoker 运行作业
func (c *AIDataLakeClient) RunRayJobInvoker(request *model.RunRayJobRequest) *RunRayJobInvoker {
	requestDef := GenReqDefForRunRayJob()
	return &RunRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRayJob 查看Ray作业详情
//
// 查看Ray作业详情。本接口为同步接口，当前查看Ray作业详情请求下发成功后会返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowRayJob(request *model.ShowRayJobRequest) (*model.ShowRayJobResponse, error) {
	requestDef := GenReqDefForShowRayJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRayJobResponse), nil
	}
}

// ShowRayJobInvoker 查看Ray作业详情
func (c *AIDataLakeClient) ShowRayJobInvoker(request *model.ShowRayJobRequest) *ShowRayJobInvoker {
	requestDef := GenReqDefForShowRayJob()
	return &ShowRayJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CloseRaySession 关闭交互式会话
//
// 关闭一个已存在的交互式会话，将会话状态设置为成功。本接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CloseRaySession(request *model.CloseRaySessionRequest) (*model.CloseRaySessionResponse, error) {
	requestDef := GenReqDefForCloseRaySession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CloseRaySessionResponse), nil
	}
}

// CloseRaySessionInvoker 关闭交互式会话
func (c *AIDataLakeClient) CloseRaySessionInvoker(request *model.CloseRaySessionRequest) *CloseRaySessionInvoker {
	requestDef := GenReqDefForCloseRaySession()
	return &CloseRaySessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRaySession 创建交互式会话
//
// 创建一个交互式会话。本接口为同步接口，用于Notebook等交互式工具连接Ray集群。联邦用户认证场景下，会返回含联邦用户身份对应的凭据名称。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CreateRaySession(request *model.CreateRaySessionRequest) (*model.CreateRaySessionResponse, error) {
	requestDef := GenReqDefForCreateRaySession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRaySessionResponse), nil
	}
}

// CreateRaySessionInvoker 创建交互式会话
func (c *AIDataLakeClient) CreateRaySessionInvoker(request *model.CreateRaySessionRequest) *CreateRaySessionInvoker {
	requestDef := GenReqDefForCreateRaySession()
	return &CreateRaySessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSparkJob 取消Spark作业执行
//
// 取消正在执行的Spark作业，此接口为同步接口。调用成功后，作业将被终止执行，直接返回取消结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelSparkJob(request *model.CancelSparkJobRequest) (*model.CancelSparkJobResponse, error) {
	requestDef := GenReqDefForCancelSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkJobResponse), nil
	}
}

// CancelSparkJobInvoker 取消Spark作业执行
func (c *AIDataLakeClient) CancelSparkJobInvoker(request *model.CancelSparkJobRequest) *CancelSparkJobInvoker {
	requestDef := GenReqDefForCancelSparkJob()
	return &CancelSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkJobs 查询Spark作业列表
//
// 查询工作空间下Spark作业列表，此接口为同步接口。支持按作业ID、作业名称、作业状态、作业类型、创建时间等条件过滤查询，支持分页查询，调用成功后直接返回作业列表数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListSparkJobs(request *model.ListSparkJobsRequest) (*model.ListSparkJobsResponse, error) {
	requestDef := GenReqDefForListSparkJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkJobsResponse), nil
	}
}

// ListSparkJobsInvoker 查询Spark作业列表
func (c *AIDataLakeClient) ListSparkJobsInvoker(request *model.ListSparkJobsRequest) *ListSparkJobsInvoker {
	requestDef := GenReqDefForListSparkJobs()
	return &ListSparkJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSparkJob 启动Spark作业
//
// 启动Spark作业，此接口为异步接口。支持Spark Jar作业、Python作业和SQL Script作业。调用该接口后，作业将提交到队列等待执行，返回作业ID后需通过查询作业状态接口确认作业是否成功启动。查询作业状态请参见[查询Spark作业的状态](ShowSparkJobState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunSparkJob(request *model.RunSparkJobRequest) (*model.RunSparkJobResponse, error) {
	requestDef := GenReqDefForRunSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSparkJobResponse), nil
	}
}

// RunSparkJobInvoker 启动Spark作业
func (c *AIDataLakeClient) RunSparkJobInvoker(request *model.RunSparkJobRequest) *RunSparkJobInvoker {
	requestDef := GenReqDefForRunSparkJob()
	return &RunSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJob 查看Spark作业详情
//
// 查看指定Spark作业的详细信息，此接口为同步接口。包括作业ID、作业名称、作业状态、作业配置、资源配置、镜像配置等完整信息，调用成功后直接返回作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkJob(request *model.ShowSparkJobRequest) (*model.ShowSparkJobResponse, error) {
	requestDef := GenReqDefForShowSparkJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobResponse), nil
	}
}

// ShowSparkJobInvoker 查看Spark作业详情
func (c *AIDataLakeClient) ShowSparkJobInvoker(request *model.ShowSparkJobRequest) *ShowSparkJobInvoker {
	requestDef := GenReqDefForShowSparkJob()
	return &ShowSparkJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkJobState 查询Spark作业的状态
//
// 查询Spark作业的状态，此接口为同步接口。可通过作业ID查询Spark作业的当前执行状态，包括排队中、运行中、已成功、已失败等状态，调用后立即返回作业当前状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkJobState(request *model.ShowSparkJobStateRequest) (*model.ShowSparkJobStateResponse, error) {
	requestDef := GenReqDefForShowSparkJobState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkJobStateResponse), nil
	}
}

// ShowSparkJobStateInvoker 查询Spark作业的状态
func (c *AIDataLakeClient) ShowSparkJobStateInvoker(request *model.ShowSparkJobStateRequest) *ShowSparkJobStateInvoker {
	requestDef := GenReqDefForShowSparkJobState()
	return &ShowSparkJobStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkOperatorState 查询Spark异步操作状态
//
// 查询Spark异步操作的执行状态，用于获取异步操作（如重启集群等）的当前状态和执行结果。该接口为同步接口。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkOperatorState(request *model.ShowSparkOperatorStateRequest) (*model.ShowSparkOperatorStateResponse, error) {
	requestDef := GenReqDefForShowSparkOperatorState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkOperatorStateResponse), nil
	}
}

// ShowSparkOperatorStateInvoker 查询Spark异步操作状态
func (c *AIDataLakeClient) ShowSparkOperatorStateInvoker(request *model.ShowSparkOperatorStateRequest) *ShowSparkOperatorStateInvoker {
	requestDef := GenReqDefForShowSparkOperatorState()
	return &ShowSparkOperatorStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelSparkSql 取消SparkSql作业执行
//
// 取消正在运行或排队中的SparkSql作业，此接口为同步接口。只能取消处于QUEUED或RUNNING状态的作业，调用成功后直接返回取消结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CancelSparkSql(request *model.CancelSparkSqlRequest) (*model.CancelSparkSqlResponse, error) {
	requestDef := GenReqDefForCancelSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelSparkSqlResponse), nil
	}
}

// CancelSparkSqlInvoker 取消SparkSql作业执行
func (c *AIDataLakeClient) CancelSparkSqlInvoker(request *model.CancelSparkSqlRequest) *CancelSparkSqlInvoker {
	requestDef := GenReqDefForCancelSparkSql()
	return &CancelSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSparkSqls 查询SparkSql作业列表
//
// 查询工作空间下SparkSql作业列表，此接口为同步接口。支持按作业状态、创建时间、SQL片段等条件进行过滤查询，调用成功后直接返回作业列表数据。可通过[查看SparkSql作业详情](ShowSparkSql.xml)接口查看作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListSparkSqls(request *model.ListSparkSqlsRequest) (*model.ListSparkSqlsResponse, error) {
	requestDef := GenReqDefForListSparkSqls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSparkSqlsResponse), nil
	}
}

// ListSparkSqlsInvoker 查询SparkSql作业列表
func (c *AIDataLakeClient) ListSparkSqlsInvoker(request *model.ListSparkSqlsRequest) *ListSparkSqlsInvoker {
	requestDef := GenReqDefForListSparkSqls()
	return &ListSparkSqlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PreviewSparkSqlResult 预览SparkSql作业查询结果
//
// 预览SparkSql作业的查询结果，此接口为同步接口。仅适用于执行成功的DQL类型作业，可查看作业返回的数据内容，调用成功后直接返回查询结果数据。可通过[查看SparkSql作业详情](ShowSparkSql.xml)接口查看作业详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) PreviewSparkSqlResult(request *model.PreviewSparkSqlResultRequest) (*model.PreviewSparkSqlResultResponse, error) {
	requestDef := GenReqDefForPreviewSparkSqlResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PreviewSparkSqlResultResponse), nil
	}
}

// PreviewSparkSqlResultInvoker 预览SparkSql作业查询结果
func (c *AIDataLakeClient) PreviewSparkSqlResultInvoker(request *model.PreviewSparkSqlResultRequest) *PreviewSparkSqlResultInvoker {
	requestDef := GenReqDefForPreviewSparkSqlResult()
	return &PreviewSparkSqlResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RestartSparkSqlCluster 重启SparkSql集群
//
// 重启SparkSql集群，该接口为异步接口，接口调用成功后会返回操作ID（operation_id），您可以通过查询Spark异步操作状态接口查询操作执行结果，详情请参见[查询Spark异步操作状态](ShowSparkOperatorState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RestartSparkSqlCluster(request *model.RestartSparkSqlClusterRequest) (*model.RestartSparkSqlClusterResponse, error) {
	requestDef := GenReqDefForRestartSparkSqlCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartSparkSqlClusterResponse), nil
	}
}

// RestartSparkSqlClusterInvoker 重启SparkSql集群
func (c *AIDataLakeClient) RestartSparkSqlClusterInvoker(request *model.RestartSparkSqlClusterRequest) *RestartSparkSqlClusterInvoker {
	requestDef := GenReqDefForRestartSparkSqlCluster()
	return &RestartSparkSqlClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunSparkSql 执行SparkSql作业
//
// 执行SparkSql作业，此接口为异步接口。接口调用成功后会返回作业ID（statement_id）,您可以通过查询作业状态接口查询作业执行结果，详情请参见[查询SparkSql作业的状态](ShowSparkSqlState.xml)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) RunSparkSql(request *model.RunSparkSqlRequest) (*model.RunSparkSqlResponse, error) {
	requestDef := GenReqDefForRunSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSparkSqlResponse), nil
	}
}

// RunSparkSqlInvoker 执行SparkSql作业
func (c *AIDataLakeClient) RunSparkSqlInvoker(request *model.RunSparkSqlRequest) *RunSparkSqlInvoker {
	requestDef := GenReqDefForRunSparkSql()
	return &RunSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkSql 查看SparkSql作业详情
//
// 查看SparkSql作业的详细信息，此接口为同步接口。包括作业状态、SQL内容、执行参数等，调用成功后直接返回作业详细信息。可通过[执行SparkSql作业](RunSparkSql.xml)接口创建作业，通过[查询SparkSql作业列表](ListSparkSqls.xml)接口查询作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkSql(request *model.ShowSparkSqlRequest) (*model.ShowSparkSqlResponse, error) {
	requestDef := GenReqDefForShowSparkSql()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkSqlResponse), nil
	}
}

// ShowSparkSqlInvoker 查看SparkSql作业详情
func (c *AIDataLakeClient) ShowSparkSqlInvoker(request *model.ShowSparkSqlRequest) *ShowSparkSqlInvoker {
	requestDef := GenReqDefForShowSparkSql()
	return &ShowSparkSqlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSparkSqlState 查询SparkSql作业的状态
//
// 查询SparkSql作业的状态，此接口为同步接口。可通过查询SparkSql作业列表接口获取statement_id，调用成功后直接返回作业当前状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ShowSparkSqlState(request *model.ShowSparkSqlStateRequest) (*model.ShowSparkSqlStateResponse, error) {
	requestDef := GenReqDefForShowSparkSqlState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSparkSqlStateResponse), nil
	}
}

// ShowSparkSqlStateInvoker 查询SparkSql作业的状态
func (c *AIDataLakeClient) ShowSparkSqlStateInvoker(request *model.ShowSparkSqlStateRequest) *ShowSparkSqlStateInvoker {
	requestDef := GenReqDefForShowSparkSqlState()
	return &ShowSparkSqlStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateV2ResourceTags 批量添加资源标签
//
// 资源批量添加标签。此接口为同步接口，无配套使用接口。创建时，不允许设置重复key数据，如果数据库已存在该key，就覆盖value的值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) BatchCreateV2ResourceTags(request *model.BatchCreateV2ResourceTagsRequest) (*model.BatchCreateV2ResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateV2ResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateV2ResourceTagsResponse), nil
	}
}

// BatchCreateV2ResourceTagsInvoker 批量添加资源标签
func (c *AIDataLakeClient) BatchCreateV2ResourceTagsInvoker(request *model.BatchCreateV2ResourceTagsRequest) *BatchCreateV2ResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateV2ResourceTags()
	return &BatchCreateV2ResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteV2ResourceTags 批量删除资源标签
//
// 标签资源批量删除标签。此接口为同步接口，无配套使用接口。删除时，如果删除的标签不存在，默认处理成功，删除时不对标签字符集范围做校验。删除时tags结构体不能缺失，key不能为空或者空字符串。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) BatchDeleteV2ResourceTags(request *model.BatchDeleteV2ResourceTagsRequest) (*model.BatchDeleteV2ResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteV2ResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteV2ResourceTagsResponse), nil
	}
}

// BatchDeleteV2ResourceTagsInvoker 批量删除资源标签
func (c *AIDataLakeClient) BatchDeleteV2ResourceTagsInvoker(request *model.BatchDeleteV2ResourceTagsRequest) *BatchDeleteV2ResourceTagsInvoker {
	requestDef := GenReqDefForBatchDeleteV2ResourceTags()
	return &BatchDeleteV2ResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountV2TagResources 查询资源实例数量
//
// 通过标签查询资源实例数量。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) CountV2TagResources(request *model.CountV2TagResourcesRequest) (*model.CountV2TagResourcesResponse, error) {
	requestDef := GenReqDefForCountV2TagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountV2TagResourcesResponse), nil
	}
}

// CountV2TagResourcesInvoker 查询资源实例数量
func (c *AIDataLakeClient) CountV2TagResourcesInvoker(request *model.CountV2TagResourcesRequest) *CountV2TagResourcesInvoker {
	requestDef := GenReqDefForCountV2TagResources()
	return &CountV2TagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2ProjectTags 查询项目标签
//
// 查询指定项目关联的标签列表。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2ProjectTags(request *model.ListV2ProjectTagsRequest) (*model.ListV2ProjectTagsResponse, error) {
	requestDef := GenReqDefForListV2ProjectTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2ProjectTagsResponse), nil
	}
}

// ListV2ProjectTagsInvoker 查询项目标签
func (c *AIDataLakeClient) ListV2ProjectTagsInvoker(request *model.ListV2ProjectTagsRequest) *ListV2ProjectTagsInvoker {
	requestDef := GenReqDefForListV2ProjectTags()
	return &ListV2ProjectTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2ResourceTags 查询资源标签
//
// 查询资源的标签。用于查询资源的标签，返回对应资源的标签。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2ResourceTags(request *model.ListV2ResourceTagsRequest) (*model.ListV2ResourceTagsResponse, error) {
	requestDef := GenReqDefForListV2ResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2ResourceTagsResponse), nil
	}
}

// ListV2ResourceTagsInvoker 查询资源标签
func (c *AIDataLakeClient) ListV2ResourceTagsInvoker(request *model.ListV2ResourceTagsRequest) *ListV2ResourceTagsInvoker {
	requestDef := GenReqDefForListV2ResourceTags()
	return &ListV2ResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListV2TagResources 查询资源实例列表
//
// 通过标签查询资源列表。此接口为同步接口，无配套使用接口和特殊场景。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AIDataLakeClient) ListV2TagResources(request *model.ListV2TagResourcesRequest) (*model.ListV2TagResourcesResponse, error) {
	requestDef := GenReqDefForListV2TagResources()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListV2TagResourcesResponse), nil
	}
}

// ListV2TagResourcesInvoker 查询资源实例列表
func (c *AIDataLakeClient) ListV2TagResourcesInvoker(request *model.ListV2TagResourcesRequest) *ListV2TagResourcesInvoker {
	requestDef := GenReqDefForListV2TagResources()
	return &ListV2TagResourcesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
