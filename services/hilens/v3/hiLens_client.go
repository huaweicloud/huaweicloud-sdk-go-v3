package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/hilens/v3/model"
)

type HiLensClient struct {
	HcClient *http_client.HcHttpClient
}

func NewHiLensClient(hcClient *http_client.HcHttpClient) *HiLensClient {
	return &HiLensClient{HcClient: hcClient}
}

func HiLensClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddDeploymentNodes 批量部署
//
// 通过指定设备id列表或者设备标签将应用部署下发到多个设备上。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) AddDeploymentNodes(request *model.AddDeploymentNodesRequest) (*model.AddDeploymentNodesResponse, error) {
	requestDef := GenReqDefForAddDeploymentNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDeploymentNodesResponse), nil
	}
}

// AddDeploymentNodesInvoker 批量部署
func (c *HiLensClient) AddDeploymentNodesInvoker(request *model.AddDeploymentNodesRequest) *AddDeploymentNodesInvoker {
	requestDef := GenReqDefForAddDeploymentNodes()
	return &AddDeploymentNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateNodeTags 批量添加节点标签
//
// 专业版HiLens控制台标签管理，用户选择多个设备，批量添加多个标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) BatchCreateNodeTags(request *model.BatchCreateNodeTagsRequest) (*model.BatchCreateNodeTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateNodeTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateNodeTagsResponse), nil
	}
}

// BatchCreateNodeTagsInvoker 批量添加节点标签
func (c *HiLensClient) BatchCreateNodeTagsInvoker(request *model.BatchCreateNodeTagsRequest) *BatchCreateNodeTagsInvoker {
	requestDef := GenReqDefForBatchCreateNodeTags()
	return &BatchCreateNodeTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConfigMap 创建配置项
//
// 创建配置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateConfigMap(request *model.CreateConfigMapRequest) (*model.CreateConfigMapResponse, error) {
	requestDef := GenReqDefForCreateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConfigMapResponse), nil
	}
}

// CreateConfigMapInvoker 创建配置项
func (c *HiLensClient) CreateConfigMapInvoker(request *model.CreateConfigMapRequest) *CreateConfigMapInvoker {
	requestDef := GenReqDefForCreateConfigMap()
	return &CreateConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDeployment 创建应用部署
//
// 创建应用部署。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateDeployment(request *model.CreateDeploymentRequest) (*model.CreateDeploymentResponse, error) {
	requestDef := GenReqDefForCreateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDeploymentResponse), nil
	}
}

// CreateDeploymentInvoker 创建应用部署
func (c *HiLensClient) CreateDeploymentInvoker(request *model.CreateDeploymentRequest) *CreateDeploymentInvoker {
	requestDef := GenReqDefForCreateDeployment()
	return &CreateDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNode 注册设备
//
// 填写设备信息，将设备注册到HiLens专业版控制台上。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateNode(request *model.CreateNodeRequest) (*model.CreateNodeResponse, error) {
	requestDef := GenReqDefForCreateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodeResponse), nil
	}
}

// CreateNodeInvoker 注册设备
func (c *HiLensClient) CreateNodeInvoker(request *model.CreateNodeRequest) *CreateNodeInvoker {
	requestDef := GenReqDefForCreateNode()
	return &CreateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrderForm 创建免费技能订单
//
// 创建免费技能订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateOrderForm(request *model.CreateOrderFormRequest) (*model.CreateOrderFormResponse, error) {
	requestDef := GenReqDefForCreateOrderForm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrderFormResponse), nil
	}
}

// CreateOrderFormInvoker 创建免费技能订单
func (c *HiLensClient) CreateOrderFormInvoker(request *model.CreateOrderFormRequest) *CreateOrderFormInvoker {
	requestDef := GenReqDefForCreateOrderForm()
	return &CreateOrderFormInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateResourceTags 添加资源标签
//
// 专业版HiLens控制台标签管理，添加对应资源的标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateResourceTags(request *model.CreateResourceTagsRequest) (*model.CreateResourceTagsResponse, error) {
	requestDef := GenReqDefForCreateResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateResourceTagsResponse), nil
	}
}

// CreateResourceTagsInvoker 添加资源标签
func (c *HiLensClient) CreateResourceTagsInvoker(request *model.CreateResourceTagsRequest) *CreateResourceTagsInvoker {
	requestDef := GenReqDefForCreateResourceTags()
	return &CreateResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecret 创建密钥
//
// 创建密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateSecret(request *model.CreateSecretRequest) (*model.CreateSecretResponse, error) {
	requestDef := GenReqDefForCreateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecretResponse), nil
	}
}

// CreateSecretInvoker 创建密钥
func (c *HiLensClient) CreateSecretInvoker(request *model.CreateSecretRequest) *CreateSecretInvoker {
	requestDef := GenReqDefForCreateSecret()
	return &CreateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTask 创建作业
//
// 创建作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// CreateTaskInvoker 创建作业
func (c *HiLensClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkSpace 创建工作空间
//
// 创建一个工作空间，其中工作空间名不能与已有的重复
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) CreateWorkSpace(request *model.CreateWorkSpaceRequest) (*model.CreateWorkSpaceResponse, error) {
	requestDef := GenReqDefForCreateWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkSpaceResponse), nil
	}
}

// CreateWorkSpaceInvoker 创建工作空间
func (c *HiLensClient) CreateWorkSpaceInvoker(request *model.CreateWorkSpaceRequest) *CreateWorkSpaceInvoker {
	requestDef := GenReqDefForCreateWorkSpace()
	return &CreateWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteConfigMap 删除配置项
//
// 根据配置项id删除某个配置项
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteConfigMap(request *model.DeleteConfigMapRequest) (*model.DeleteConfigMapResponse, error) {
	requestDef := GenReqDefForDeleteConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteConfigMapResponse), nil
	}
}

// DeleteConfigMapInvoker 删除配置项
func (c *HiLensClient) DeleteConfigMapInvoker(request *model.DeleteConfigMapRequest) *DeleteConfigMapInvoker {
	requestDef := GenReqDefForDeleteConfigMap()
	return &DeleteConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDeployment 删除应用部署
//
// 删除指定应用部署。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteDeployment(request *model.DeleteDeploymentRequest) (*model.DeleteDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDeploymentResponse), nil
	}
}

// DeleteDeploymentInvoker 删除应用部署
func (c *HiLensClient) DeleteDeploymentInvoker(request *model.DeleteDeploymentRequest) *DeleteDeploymentInvoker {
	requestDef := GenReqDefForDeleteDeployment()
	return &DeleteDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNode 删除设备
//
// 删除专业版HiLens控制台上的设备，并与端侧的设备进行解绑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteNode(request *model.DeleteNodeRequest) (*model.DeleteNodeResponse, error) {
	requestDef := GenReqDefForDeleteNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodeResponse), nil
	}
}

// DeleteNodeInvoker 删除设备
func (c *HiLensClient) DeleteNodeInvoker(request *model.DeleteNodeRequest) *DeleteNodeInvoker {
	requestDef := GenReqDefForDeleteNode()
	return &DeleteNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePod 删除应用实例
//
// 删除指定实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeletePod(request *model.DeletePodRequest) (*model.DeletePodResponse, error) {
	requestDef := GenReqDefForDeletePod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePodResponse), nil
	}
}

// DeletePodInvoker 删除应用实例
func (c *HiLensClient) DeletePodInvoker(request *model.DeletePodRequest) *DeletePodInvoker {
	requestDef := GenReqDefForDeletePod()
	return &DeletePodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteResourceTag 删除资源标签
//
// 专业版HiLens控制台标签管理，删除对应资源的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteResourceTag(request *model.DeleteResourceTagRequest) (*model.DeleteResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteResourceTagResponse), nil
	}
}

// DeleteResourceTagInvoker 删除资源标签
func (c *HiLensClient) DeleteResourceTagInvoker(request *model.DeleteResourceTagRequest) *DeleteResourceTagInvoker {
	requestDef := GenReqDefForDeleteResourceTag()
	return &DeleteResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecret 删除密钥
//
// 删除密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteSecret(request *model.DeleteSecretRequest) (*model.DeleteSecretResponse, error) {
	requestDef := GenReqDefForDeleteSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecretResponse), nil
	}
}

// DeleteSecretInvoker 删除密钥
func (c *HiLensClient) DeleteSecretInvoker(request *model.DeleteSecretRequest) *DeleteSecretInvoker {
	requestDef := GenReqDefForDeleteSecret()
	return &DeleteSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTask 删除作业
//
// 删除作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteTask(request *model.DeleteTaskRequest) (*model.DeleteTaskResponse, error) {
	requestDef := GenReqDefForDeleteTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskResponse), nil
	}
}

// DeleteTaskInvoker 删除作业
func (c *HiLensClient) DeleteTaskInvoker(request *model.DeleteTaskRequest) *DeleteTaskInvoker {
	requestDef := GenReqDefForDeleteTask()
	return &DeleteTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkSpace 删除工作空间
//
// 删除指定ID的工作空间，如果该工作空间下仍有资源，则删除会失败
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) DeleteWorkSpace(request *model.DeleteWorkSpaceRequest) (*model.DeleteWorkSpaceResponse, error) {
	requestDef := GenReqDefForDeleteWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkSpaceResponse), nil
	}
}

// DeleteWorkSpaceInvoker 删除工作空间
func (c *HiLensClient) DeleteWorkSpaceInvoker(request *model.DeleteWorkSpaceRequest) *DeleteWorkSpaceInvoker {
	requestDef := GenReqDefForDeleteWorkSpace()
	return &DeleteWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// FreezeNode 将激活订单与设备解绑
//
// 将激活订单与设备解绑。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) FreezeNode(request *model.FreezeNodeRequest) (*model.FreezeNodeResponse, error) {
	requestDef := GenReqDefForFreezeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.FreezeNodeResponse), nil
	}
}

// FreezeNodeInvoker 将激活订单与设备解绑
func (c *HiLensClient) FreezeNodeInvoker(request *model.FreezeNodeRequest) *FreezeNodeInvoker {
	requestDef := GenReqDefForFreezeNode()
	return &FreezeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConfigMaps 查询配置项列表
//
// 获取配置项详情，以列表形式返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListConfigMaps(request *model.ListConfigMapsRequest) (*model.ListConfigMapsResponse, error) {
	requestDef := GenReqDefForListConfigMaps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConfigMapsResponse), nil
	}
}

// ListConfigMapsInvoker 查询配置项列表
func (c *HiLensClient) ListConfigMapsInvoker(request *model.ListConfigMapsRequest) *ListConfigMapsInvoker {
	requestDef := GenReqDefForListConfigMaps()
	return &ListConfigMapsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirmwares 查询固件列表
//
// 查看指定固件历史版本信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListFirmwares(request *model.ListFirmwaresRequest) (*model.ListFirmwaresResponse, error) {
	requestDef := GenReqDefForListFirmwares()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirmwaresResponse), nil
	}
}

// ListFirmwaresInvoker 查询固件列表
func (c *HiLensClient) ListFirmwaresInvoker(request *model.ListFirmwaresRequest) *ListFirmwaresInvoker {
	requestDef := GenReqDefForListFirmwares()
	return &ListFirmwaresInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPlatformManager 获取运行服务费订单列表
//
// 获取平台管理费列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListPlatformManager(request *model.ListPlatformManagerRequest) (*model.ListPlatformManagerResponse, error) {
	requestDef := GenReqDefForListPlatformManager()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPlatformManagerResponse), nil
	}
}

// ListPlatformManagerInvoker 获取运行服务费订单列表
func (c *HiLensClient) ListPlatformManagerInvoker(request *model.ListPlatformManagerRequest) *ListPlatformManagerInvoker {
	requestDef := GenReqDefForListPlatformManager()
	return &ListPlatformManagerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceTags 查询某资源类型的标签
//
// 专业版HiLens控制台标签管理，查询某种资源类型的所有标签，返回标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListResourceTags(request *model.ListResourceTagsRequest) (*model.ListResourceTagsResponse, error) {
	requestDef := GenReqDefForListResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceTagsResponse), nil
	}
}

// ListResourceTagsInvoker 查询某资源类型的标签
func (c *HiLensClient) ListResourceTagsInvoker(request *model.ListResourceTagsRequest) *ListResourceTagsInvoker {
	requestDef := GenReqDefForListResourceTags()
	return &ListResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecrets 查询密钥列表
//
// 专业版HiLens控制台密钥管理，根据用户请求条件筛选，查询用户创建的 密钥信息，以列表形式返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListSecrets(request *model.ListSecretsRequest) (*model.ListSecretsResponse, error) {
	requestDef := GenReqDefForListSecrets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecretsResponse), nil
	}
}

// ListSecretsInvoker 查询密钥列表
func (c *HiLensClient) ListSecretsInvoker(request *model.ListSecretsRequest) *ListSecretsInvoker {
	requestDef := GenReqDefForListSecrets()
	return &ListSecretsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTasks 查询作业列表
//
// 查询当前部署下所有作业，返回详情列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListTasks(request *model.ListTasksRequest) (*model.ListTasksResponse, error) {
	requestDef := GenReqDefForListTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTasksResponse), nil
	}
}

// ListTasksInvoker 查询作业列表
func (c *HiLensClient) ListTasksInvoker(request *model.ListTasksRequest) *ListTasksInvoker {
	requestDef := GenReqDefForListTasks()
	return &ListTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkSpaces 获取工作空间列表
//
// 查询用户名下的所有工作空间信息，并返回列表和总条目数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ListWorkSpaces(request *model.ListWorkSpacesRequest) (*model.ListWorkSpacesResponse, error) {
	requestDef := GenReqDefForListWorkSpaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkSpacesResponse), nil
	}
}

// ListWorkSpacesInvoker 获取工作空间列表
func (c *HiLensClient) ListWorkSpacesInvoker(request *model.ListWorkSpacesRequest) *ListWorkSpacesInvoker {
	requestDef := GenReqDefForListWorkSpaces()
	return &ListWorkSpacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetDefaultOrderForm 设置默认订单
//
// 设置默认订单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) SetDefaultOrderForm(request *model.SetDefaultOrderFormRequest) (*model.SetDefaultOrderFormResponse, error) {
	requestDef := GenReqDefForSetDefaultOrderForm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetDefaultOrderFormResponse), nil
	}
}

// SetDefaultOrderFormInvoker 设置默认订单
func (c *HiLensClient) SetDefaultOrderFormInvoker(request *model.SetDefaultOrderFormRequest) *SetDefaultOrderFormInvoker {
	requestDef := GenReqDefForSetDefaultOrderForm()
	return &SetDefaultOrderFormInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConfigMap 查询配置项详情
//
// 根据配置项id查询某个配置项详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowConfigMap(request *model.ShowConfigMapRequest) (*model.ShowConfigMapResponse, error) {
	requestDef := GenReqDefForShowConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConfigMapResponse), nil
	}
}

// ShowConfigMapInvoker 查询配置项详情
func (c *HiLensClient) ShowConfigMapInvoker(request *model.ShowConfigMapRequest) *ShowConfigMapInvoker {
	requestDef := GenReqDefForShowConfigMap()
	return &ShowConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeployment 查询应用部署详情
//
// 获取部署的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowDeployment(request *model.ShowDeploymentRequest) (*model.ShowDeploymentResponse, error) {
	requestDef := GenReqDefForShowDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentResponse), nil
	}
}

// ShowDeploymentInvoker 查询应用部署详情
func (c *HiLensClient) ShowDeploymentInvoker(request *model.ShowDeploymentRequest) *ShowDeploymentInvoker {
	requestDef := GenReqDefForShowDeployment()
	return &ShowDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeploymentPods 查询应用实例列表
//
// 获取用户实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowDeploymentPods(request *model.ShowDeploymentPodsRequest) (*model.ShowDeploymentPodsResponse, error) {
	requestDef := GenReqDefForShowDeploymentPods()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentPodsResponse), nil
	}
}

// ShowDeploymentPodsInvoker 查询应用实例列表
func (c *HiLensClient) ShowDeploymentPodsInvoker(request *model.ShowDeploymentPodsRequest) *ShowDeploymentPodsInvoker {
	requestDef := GenReqDefForShowDeploymentPods()
	return &ShowDeploymentPodsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDeployments 查询应用部署列表
//
// 获取部署列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowDeployments(request *model.ShowDeploymentsRequest) (*model.ShowDeploymentsResponse, error) {
	requestDef := GenReqDefForShowDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDeploymentsResponse), nil
	}
}

// ShowDeploymentsInvoker 查询应用部署列表
func (c *HiLensClient) ShowDeploymentsInvoker(request *model.ShowDeploymentsRequest) *ShowDeploymentsInvoker {
	requestDef := GenReqDefForShowDeployments()
	return &ShowDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNode 查询设备详情
//
// 支持查询HiLens专业版控制台上的设备详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowNode(request *model.ShowNodeRequest) (*model.ShowNodeResponse, error) {
	requestDef := GenReqDefForShowNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeResponse), nil
	}
}

// ShowNodeInvoker 查询设备详情
func (c *HiLensClient) ShowNodeInvoker(request *model.ShowNodeRequest) *ShowNodeInvoker {
	requestDef := GenReqDefForShowNode()
	return &ShowNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodeActivationRecords 获取激活记录列表
//
// 获取激活记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowNodeActivationRecords(request *model.ShowNodeActivationRecordsRequest) (*model.ShowNodeActivationRecordsResponse, error) {
	requestDef := GenReqDefForShowNodeActivationRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeActivationRecordsResponse), nil
	}
}

// ShowNodeActivationRecordsInvoker 获取激活记录列表
func (c *HiLensClient) ShowNodeActivationRecordsInvoker(request *model.ShowNodeActivationRecordsRequest) *ShowNodeActivationRecordsInvoker {
	requestDef := GenReqDefForShowNodeActivationRecords()
	return &ShowNodeActivationRecordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodes 查询设备列表
//
// 专业版HiLens控制台设备管理，根据用户请求条件筛选，查询用户注册的设备信息，以列表形式返回。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowNodes(request *model.ShowNodesRequest) (*model.ShowNodesResponse, error) {
	requestDef := GenReqDefForShowNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodesResponse), nil
	}
}

// ShowNodesInvoker 查询设备列表
func (c *HiLensClient) ShowNodesInvoker(request *model.ShowNodesRequest) *ShowNodesInvoker {
	requestDef := GenReqDefForShowNodes()
	return &ShowNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceTags 查询资源标签
//
// 专业版HiLens控制台标签管理，查询具体资源的标签，返回标签列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowResourceTags(request *model.ShowResourceTagsRequest) (*model.ShowResourceTagsResponse, error) {
	requestDef := GenReqDefForShowResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagsResponse), nil
	}
}

// ShowResourceTagsInvoker 查询资源标签
func (c *HiLensClient) ShowResourceTagsInvoker(request *model.ShowResourceTagsRequest) *ShowResourceTagsInvoker {
	requestDef := GenReqDefForShowResourceTags()
	return &ShowResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecret 查询密钥详情
//
// 查询密钥详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowSecret(request *model.ShowSecretRequest) (*model.ShowSecretResponse, error) {
	requestDef := GenReqDefForShowSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecretResponse), nil
	}
}

// ShowSecretInvoker 查询密钥详情
func (c *HiLensClient) ShowSecretInvoker(request *model.ShowSecretRequest) *ShowSecretInvoker {
	requestDef := GenReqDefForShowSecret()
	return &ShowSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSkillInfo 获取技能详情
//
// 获取技能详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowSkillInfo(request *model.ShowSkillInfoRequest) (*model.ShowSkillInfoResponse, error) {
	requestDef := GenReqDefForShowSkillInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSkillInfoResponse), nil
	}
}

// ShowSkillInfoInvoker 获取技能详情
func (c *HiLensClient) ShowSkillInfoInvoker(request *model.ShowSkillInfoRequest) *ShowSkillInfoInvoker {
	requestDef := GenReqDefForShowSkillInfo()
	return &ShowSkillInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSkillList 查询技能列表
//
// 获取技能列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowSkillList(request *model.ShowSkillListRequest) (*model.ShowSkillListResponse, error) {
	requestDef := GenReqDefForShowSkillList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSkillListResponse), nil
	}
}

// ShowSkillListInvoker 查询技能列表
func (c *HiLensClient) ShowSkillListInvoker(request *model.ShowSkillListRequest) *ShowSkillListInvoker {
	requestDef := GenReqDefForShowSkillList()
	return &ShowSkillListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSkillOrderInfo 查询订单详情
//
// 获取订单详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowSkillOrderInfo(request *model.ShowSkillOrderInfoRequest) (*model.ShowSkillOrderInfoResponse, error) {
	requestDef := GenReqDefForShowSkillOrderInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSkillOrderInfoResponse), nil
	}
}

// ShowSkillOrderInfoInvoker 查询订单详情
func (c *HiLensClient) ShowSkillOrderInfoInvoker(request *model.ShowSkillOrderInfoRequest) *ShowSkillOrderInfoInvoker {
	requestDef := GenReqDefForShowSkillOrderInfo()
	return &ShowSkillOrderInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSkillOrderList 查询订单列表
//
// 获取订单列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowSkillOrderList(request *model.ShowSkillOrderListRequest) (*model.ShowSkillOrderListResponse, error) {
	requestDef := GenReqDefForShowSkillOrderList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSkillOrderListResponse), nil
	}
}

// ShowSkillOrderListInvoker 查询订单列表
func (c *HiLensClient) ShowSkillOrderListInvoker(request *model.ShowSkillOrderListRequest) *ShowSkillOrderListInvoker {
	requestDef := GenReqDefForShowSkillOrderList()
	return &ShowSkillOrderListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTask 查询作业详情
//
// 通过作业ID查询作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowTask(request *model.ShowTaskRequest) (*model.ShowTaskResponse, error) {
	requestDef := GenReqDefForShowTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskResponse), nil
	}
}

// ShowTaskInvoker 查询作业详情
func (c *HiLensClient) ShowTaskInvoker(request *model.ShowTaskRequest) *ShowTaskInvoker {
	requestDef := GenReqDefForShowTask()
	return &ShowTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUpgradeProgress 获取设备固件升级进度
//
// 获取设备固件升级进度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowUpgradeProgress(request *model.ShowUpgradeProgressRequest) (*model.ShowUpgradeProgressResponse, error) {
	requestDef := GenReqDefForShowUpgradeProgress()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUpgradeProgressResponse), nil
	}
}

// ShowUpgradeProgressInvoker 获取设备固件升级进度
func (c *HiLensClient) ShowUpgradeProgressInvoker(request *model.ShowUpgradeProgressRequest) *ShowUpgradeProgressInvoker {
	requestDef := GenReqDefForShowUpgradeProgress()
	return &ShowUpgradeProgressInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkSpace 获取工作空间详情
//
// 获取指定workspace_id的工作空间详情，包括创建时间，描述，创建者等信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) ShowWorkSpace(request *model.ShowWorkSpaceRequest) (*model.ShowWorkSpaceResponse, error) {
	requestDef := GenReqDefForShowWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkSpaceResponse), nil
	}
}

// ShowWorkSpaceInvoker 获取工作空间详情
func (c *HiLensClient) ShowWorkSpaceInvoker(request *model.ShowWorkSpaceRequest) *ShowWorkSpaceInvoker {
	requestDef := GenReqDefForShowWorkSpace()
	return &ShowWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAndStopDeployment 暂停、继续部署负载
//
// 启动/暂停应用部署。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) StartAndStopDeployment(request *model.StartAndStopDeploymentRequest) (*model.StartAndStopDeploymentResponse, error) {
	requestDef := GenReqDefForStartAndStopDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAndStopDeploymentResponse), nil
	}
}

// StartAndStopDeploymentInvoker 暂停、继续部署负载
func (c *HiLensClient) StartAndStopDeploymentInvoker(request *model.StartAndStopDeploymentRequest) *StartAndStopDeploymentInvoker {
	requestDef := GenReqDefForStartAndStopDeployment()
	return &StartAndStopDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartAndStopDeploymentPod 启动/停止部署下的指定实例
//
// 启动/停止部署下的指定实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) StartAndStopDeploymentPod(request *model.StartAndStopDeploymentPodRequest) (*model.StartAndStopDeploymentPodResponse, error) {
	requestDef := GenReqDefForStartAndStopDeploymentPod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartAndStopDeploymentPodResponse), nil
	}
}

// StartAndStopDeploymentPodInvoker 启动/停止部署下的指定实例
func (c *HiLensClient) StartAndStopDeploymentPodInvoker(request *model.StartAndStopDeploymentPodRequest) *StartAndStopDeploymentPodInvoker {
	requestDef := GenReqDefForStartAndStopDeploymentPod()
	return &StartAndStopDeploymentPodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchNodeConnection 启停设备
//
// 该API用于启用停用设备。被停用的设备将无法连接到云端服务，重新启用设备恢复连接。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) SwitchNodeConnection(request *model.SwitchNodeConnectionRequest) (*model.SwitchNodeConnectionResponse, error) {
	requestDef := GenReqDefForSwitchNodeConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchNodeConnectionResponse), nil
	}
}

// SwitchNodeConnectionInvoker 启停设备
func (c *HiLensClient) SwitchNodeConnectionInvoker(request *model.SwitchNodeConnectionRequest) *SwitchNodeConnectionInvoker {
	requestDef := GenReqDefForSwitchNodeConnection()
	return &SwitchNodeConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnfreezeNode 使用运行服务费激活设备
//
// 使用运行服务费激活设备。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UnfreezeNode(request *model.UnfreezeNodeRequest) (*model.UnfreezeNodeResponse, error) {
	requestDef := GenReqDefForUnfreezeNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnfreezeNodeResponse), nil
	}
}

// UnfreezeNodeInvoker 使用运行服务费激活设备
func (c *HiLensClient) UnfreezeNodeInvoker(request *model.UnfreezeNodeRequest) *UnfreezeNodeInvoker {
	requestDef := GenReqDefForUnfreezeNode()
	return &UnfreezeNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateConfigMap 更新配置项
//
// 根据配置项id更新配置项信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateConfigMap(request *model.UpdateConfigMapRequest) (*model.UpdateConfigMapResponse, error) {
	requestDef := GenReqDefForUpdateConfigMap()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateConfigMapResponse), nil
	}
}

// UpdateConfigMapInvoker 更新配置项
func (c *HiLensClient) UpdateConfigMapInvoker(request *model.UpdateConfigMapRequest) *UpdateConfigMapInvoker {
	requestDef := GenReqDefForUpdateConfigMap()
	return &UpdateConfigMapInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeployment 更新应用部署
//
// 更新应用部署相关信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateDeployment(request *model.UpdateDeploymentRequest) (*model.UpdateDeploymentResponse, error) {
	requestDef := GenReqDefForUpdateDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentResponse), nil
	}
}

// UpdateDeploymentInvoker 更新应用部署
func (c *HiLensClient) UpdateDeploymentInvoker(request *model.UpdateDeploymentRequest) *UpdateDeploymentInvoker {
	requestDef := GenReqDefForUpdateDeployment()
	return &UpdateDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDeploymentUsingPatch 部分更新应用部署
//
// 更新应用部署部分信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateDeploymentUsingPatch(request *model.UpdateDeploymentUsingPatchRequest) (*model.UpdateDeploymentUsingPatchResponse, error) {
	requestDef := GenReqDefForUpdateDeploymentUsingPatch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDeploymentUsingPatchResponse), nil
	}
}

// UpdateDeploymentUsingPatchInvoker 部分更新应用部署
func (c *HiLensClient) UpdateDeploymentUsingPatchInvoker(request *model.UpdateDeploymentUsingPatchRequest) *UpdateDeploymentUsingPatchInvoker {
	requestDef := GenReqDefForUpdateDeploymentUsingPatch()
	return &UpdateDeploymentUsingPatchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNode 更新设备信息
//
// 更新设备日志配置，标签以及描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateNode(request *model.UpdateNodeRequest) (*model.UpdateNodeResponse, error) {
	requestDef := GenReqDefForUpdateNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeResponse), nil
	}
}

// UpdateNodeInvoker 更新设备信息
func (c *HiLensClient) UpdateNodeInvoker(request *model.UpdateNodeRequest) *UpdateNodeInvoker {
	requestDef := GenReqDefForUpdateNode()
	return &UpdateNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNodeCert 更新设备证书
//
// 设备出现离线或者证书过期时，可通过该接口更新证书，重新让设备连接到云端
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateNodeCert(request *model.UpdateNodeCertRequest) (*model.UpdateNodeCertResponse, error) {
	requestDef := GenReqDefForUpdateNodeCert()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeCertResponse), nil
	}
}

// UpdateNodeCertInvoker 更新设备证书
func (c *HiLensClient) UpdateNodeCertInvoker(request *model.UpdateNodeCertRequest) *UpdateNodeCertInvoker {
	requestDef := GenReqDefForUpdateNodeCert()
	return &UpdateNodeCertInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNodeFirmware 升级设备固件
//
// 升级设备固件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateNodeFirmware(request *model.UpdateNodeFirmwareRequest) (*model.UpdateNodeFirmwareResponse, error) {
	requestDef := GenReqDefForUpdateNodeFirmware()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNodeFirmwareResponse), nil
	}
}

// UpdateNodeFirmwareInvoker 升级设备固件
func (c *HiLensClient) UpdateNodeFirmwareInvoker(request *model.UpdateNodeFirmwareRequest) *UpdateNodeFirmwareInvoker {
	requestDef := GenReqDefForUpdateNodeFirmware()
	return &UpdateNodeFirmwareInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecret 更新密钥
//
// 更新密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateSecret(request *model.UpdateSecretRequest) (*model.UpdateSecretResponse, error) {
	requestDef := GenReqDefForUpdateSecret()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecretResponse), nil
	}
}

// UpdateSecretInvoker 更新密钥
func (c *HiLensClient) UpdateSecretInvoker(request *model.UpdateSecretRequest) *UpdateSecretInvoker {
	requestDef := GenReqDefForUpdateSecret()
	return &UpdateSecretInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTask 编辑作业
//
// 编辑作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateTask(request *model.UpdateTaskRequest) (*model.UpdateTaskResponse, error) {
	requestDef := GenReqDefForUpdateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskResponse), nil
	}
}

// UpdateTaskInvoker 编辑作业
func (c *HiLensClient) UpdateTaskInvoker(request *model.UpdateTaskRequest) *UpdateTaskInvoker {
	requestDef := GenReqDefForUpdateTask()
	return &UpdateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkSpace 修改工作空间
//
// 更改工作空间信息，暂时只能更改描述
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *HiLensClient) UpdateWorkSpace(request *model.UpdateWorkSpaceRequest) (*model.UpdateWorkSpaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkSpaceResponse), nil
	}
}

// UpdateWorkSpaceInvoker 修改工作空间
func (c *HiLensClient) UpdateWorkSpaceInvoker(request *model.UpdateWorkSpaceRequest) *UpdateWorkSpaceInvoker {
	requestDef := GenReqDefForUpdateWorkSpace()
	return &UpdateWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
