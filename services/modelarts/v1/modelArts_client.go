package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/modelarts/v1/model"
)

type ModelArtsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewModelArtsClient(hcClient *httpclient.HcHttpClient) *ModelArtsClient {
	return &ModelArtsClient{HcClient: hcClient}
}

func ModelArtsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AcceptScheduledEvent 计划事件授权
//
// 计划事件授权接口用于为指定的计划事件分配或调整权限。该接口适用于以下场景：当创建新的计划事件、调整现有计划事件的权限设置或变更权限分配时，用户可通过此接口为指定的计划事件授予或修改权限。使用该接口的前提条件是计划事件已存在且用户具有管理员权限。授权操作完成后，计划事件的权限设置将被更新，相关变更将被记录以便审计。若计划事件不存在、用户无权限操作或授权信息格式不正确，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) AcceptScheduledEvent(request *model.AcceptScheduledEventRequest) (*model.AcceptScheduledEventResponse, error) {
	requestDef := GenReqDefForAcceptScheduledEvent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AcceptScheduledEventResponse), nil
	}
}

// AcceptScheduledEventInvoker 计划事件授权
func (c *ModelArtsClient) AcceptScheduledEventInvoker(request *model.AcceptScheduledEventRequest) *AcceptScheduledEventInvoker {
	requestDef := GenReqDefForAcceptScheduledEvent()
	return &AcceptScheduledEventInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachDynamicStorage 动态挂载Notebook存储
//
// 动态挂载Notebook存储接口支持将存储动态挂载到运行中的Notebook实例的指定文件目录。调用该接口后，系统将在Notebook实例中**异步**挂载指定的存储实例，挂载完成后用户可在容器中以文件系统方式读写存储实例中的文件。若用户无权限访问指定实例或Notebook实例未运行，接口将返回相应的错误信息。
//
// 支持的存储类型：
// - **对象存储 OBS**：适合直接使用OBS桶作为持久化存储进行AI开发和探索场景，但小文件频繁读写性能较差，**模型训练，大文件解压等场景慎用，此类场景可能会导致Notebook文件操作卡顿**。
// - **并行文件系统 PFS**：高性能对象存储文件系统，存储成本低，吞吐量大，能够快速处理高性能计算（HPC）工作负载，**但小文件频繁读写较弱。小文件频繁读写场景可能会导致Notebook文件操作卡顿**
// - **高性能弹性文件服务SFS Turbo**：仅支持专属资源池实例挂载，**挂载前需要在资源池网络管理界面中进行网络关联**，支持多个环境使用，可以在多个开发环境、开发环境和训练之间共享。适合探索、实验等非正式生产场景，但不适合重IO读写模型。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) AttachDynamicStorage(request *model.AttachDynamicStorageRequest) (*model.AttachDynamicStorageResponse, error) {
	requestDef := GenReqDefForAttachDynamicStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachDynamicStorageResponse), nil
	}
}

// AttachDynamicStorageInvoker 动态挂载Notebook存储
func (c *ModelArtsClient) AttachDynamicStorageInvoker(request *model.AttachDynamicStorageRequest) *AttachDynamicStorageInvoker {
	requestDef := GenReqDefForAttachDynamicStorage()
	return &AttachDynamicStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchBindInferApiKeys 批量绑定应用密钥
//
// 本接口用于将生成的多个apikey与指定服务进行批量绑定，用于访问特定服务。调用此接口前，确保已成功创建服务实例，并获取到有效的apikey。绑定成功后，apikey将作为服务调用时的身份验证凭证，确保仅授权用户能够访问该服务。如果尝试绑定已失效或已绑定当前服务的apikey将返回相应的异常信息，提示用户检查apikey的有效性和绑定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchBindInferApiKeys(request *model.BatchBindInferApiKeysRequest) (*model.BatchBindInferApiKeysResponse, error) {
	requestDef := GenReqDefForBatchBindInferApiKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchBindInferApiKeysResponse), nil
	}
}

// BatchBindInferApiKeysInvoker 批量绑定应用密钥
func (c *ModelArtsClient) BatchBindInferApiKeysInvoker(request *model.BatchBindInferApiKeysRequest) *BatchBindInferApiKeysInvoker {
	requestDef := GenReqDefForBatchBindInferApiKeys()
	return &BatchBindInferApiKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchBindPoolNodes 批量为节点绑定逻辑子池
//
// 批量为节点绑定逻辑子池接口用于在物理专属池开启节点绑定功能时，对逻辑子池中的节点进行逻辑子池的换绑操作。该接口适用于以下场景：当需要重新分配资源、调整业务负载或优化资源使用效率时，用户可通过此接口将指定节点从当前逻辑子池迁移到另一个逻辑子池。使用该接口的前提条件是物理专属池已开启节点绑定功能，且目标逻辑子池已存在并具备足够的资源容量。绑定操作完成后，节点将从原逻辑子池解绑并绑定到目标逻辑子池，原逻辑子池的节点数减少，目标逻辑子池的节点数增加。若节点未绑定到任何逻辑子池、目标逻辑子池不存在或资源不足，或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchBindPoolNodes(request *model.BatchBindPoolNodesRequest) (*model.BatchBindPoolNodesResponse, error) {
	requestDef := GenReqDefForBatchBindPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchBindPoolNodesResponse), nil
	}
}

// BatchBindPoolNodesInvoker 批量为节点绑定逻辑子池
func (c *ModelArtsClient) BatchBindPoolNodesInvoker(request *model.BatchBindPoolNodesRequest) *BatchBindPoolNodesInvoker {
	requestDef := GenReqDefForBatchBindPoolNodes()
	return &BatchBindPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreatePoolTags 批量创建资源池标签
//
// 批量创建资源池标签接口用于为指定资源池添加或更新多个标签信息。该接口适用于以下场景：当需要对资源池进行统一分类管理（如成本归属、环境标识）、批量配置元数据（如项目归属、负责人信息）或更新已有标签值时，管理员可通过此接口一次性操作多个标签。使用该接口的前提条件是目标资源池必须已存在且处于可管理状态，调用者需具备资源池标签管理权限，且提交的标签数据需符合格式规范（如key非空、value长度限制）。操作完成后，系统将为资源池添加新标签或覆盖同名标签的值，且不会影响资源池的其他配置属性。若资源池不存在、用户权限不足、标签格式错误或系统服务异常，接口将返回对应的错误信息（如\&quot;404 Not Found\&quot;、\&quot;403 Forbidden\&quot;、\&quot;400 Bad Request\&quot;或\&quot;503 Service Unavailable\&quot;）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchCreatePoolTags(request *model.BatchCreatePoolTagsRequest) (*model.BatchCreatePoolTagsResponse, error) {
	requestDef := GenReqDefForBatchCreatePoolTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreatePoolTagsResponse), nil
	}
}

// BatchCreatePoolTagsInvoker 批量创建资源池标签
func (c *ModelArtsClient) BatchCreatePoolTagsInvoker(request *model.BatchCreatePoolTagsRequest) *BatchCreatePoolTagsInvoker {
	requestDef := GenReqDefForBatchCreatePoolTags()
	return &BatchCreatePoolTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInferIntranetConnections 批量删除内网接入
//
// 本接口用于批量删除指定的内网接入点，适用于需要清理多个不再使用的内网接入点的场景。调用此接口前，确保已具备相应的删除权限，并提供一个有效的内网接入点ID列表。删除成功后，所指定的内网接入点将被彻底移除，不再对任何服务生效。如果提供的内网接入点ID列表中包含无效或已删除的ID，将返回相应的异常信息，提示用户检查ID的有效性。此外，如果调用时出现权限不足或其他系统异常，也将返回相应的异常信息，提示用户检查权限或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchDeleteInferIntranetConnections(request *model.BatchDeleteInferIntranetConnectionsRequest) (*model.BatchDeleteInferIntranetConnectionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteInferIntranetConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInferIntranetConnectionsResponse), nil
	}
}

// BatchDeleteInferIntranetConnectionsInvoker 批量删除内网接入
func (c *ModelArtsClient) BatchDeleteInferIntranetConnectionsInvoker(request *model.BatchDeleteInferIntranetConnectionsRequest) *BatchDeleteInferIntranetConnectionsInvoker {
	requestDef := GenReqDefForBatchDeleteInferIntranetConnections()
	return &BatchDeleteInferIntranetConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteInferServices 删除指定服务列表
//
// 删除指定服务列表功能允许用户批量删除多个服务，适用于需要清理资源、释放计算能力或管理多个服务的场景。使用此功能前，请确保您具备删除服务的权限，并提供有效的服务ID列表。成功执行后，指定的服务将被终止运行并释放相关资源。若服务ID无效、权限不足或服务状态不允许删除，将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchDeleteInferServices(request *model.BatchDeleteInferServicesRequest) (*model.BatchDeleteInferServicesResponse, error) {
	requestDef := GenReqDefForBatchDeleteInferServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteInferServicesResponse), nil
	}
}

// BatchDeleteInferServicesInvoker 删除指定服务列表
func (c *ModelArtsClient) BatchDeleteInferServicesInvoker(request *model.BatchDeleteInferServicesRequest) *BatchDeleteInferServicesInvoker {
	requestDef := GenReqDefForBatchDeleteInferServices()
	return &BatchDeleteInferServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePoolNodes 批量删除节点
//
// 批量删除节点接口用于批量删除指定资源池中的节点。该接口适用于以下场景：当需要清理资源池中的冗余节点、重新分配资源或移除故障节点时，用户可通过此接口批量删除指定的节点。使用该接口的前提条件是资源池已创建且处于可用状态，用户具有删除节点的权限，且资源池中至少保留一个节点。删除操作完成后，指定的节点将被永久移除，资源池中剩余的节点将继续提供服务。若资源池不存在、节点不存在、用户无权限操作或资源池中节点不足，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchDeletePoolNodes(request *model.BatchDeletePoolNodesRequest) (*model.BatchDeletePoolNodesResponse, error) {
	requestDef := GenReqDefForBatchDeletePoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePoolNodesResponse), nil
	}
}

// BatchDeletePoolNodesInvoker 批量删除节点
func (c *ModelArtsClient) BatchDeletePoolNodesInvoker(request *model.BatchDeletePoolNodesRequest) *BatchDeletePoolNodesInvoker {
	requestDef := GenReqDefForBatchDeletePoolNodes()
	return &BatchDeletePoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeletePoolTags 批量删除资源池标签
//
// 批量删除资源标签接口用于移除指定资源上的多个标签信息。该接口适用于以下场景：当需要清理冗余标签（如过期分类、无效元数据）、统一调整资源分类策略或因权限变更需批量移除标签时，管理员可通过此接口一次性删除多个标签。使用该接口的前提条件是目标资源必须已存在且处于可管理状态，调用者需具备资源标签管理权限，且待删除的标签必须已关联至该资源，系统标签管理服务需正常运行。操作完成后，指定标签将从资源中彻底移除，且不会影响资源的其他配置属性。若资源不存在、用户权限不足、标签未关联或系统服务异常，接口将返回对应的错误信息（如\&quot;404 Not Found\&quot;、\&quot;403 Forbidden\&quot;、\&quot;400 Bad Request\&quot;或\&quot;503 Service Unavailable\&quot;）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchDeletePoolTags(request *model.BatchDeletePoolTagsRequest) (*model.BatchDeletePoolTagsResponse, error) {
	requestDef := GenReqDefForBatchDeletePoolTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeletePoolTagsResponse), nil
	}
}

// BatchDeletePoolTagsInvoker 批量删除资源池标签
func (c *ModelArtsClient) BatchDeletePoolTagsInvoker(request *model.BatchDeletePoolTagsRequest) *BatchDeletePoolTagsInvoker {
	requestDef := GenReqDefForBatchDeletePoolTags()
	return &BatchDeletePoolTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchLockPoolNodes 批量对节点功能上锁
//
// 批量对节点功能上锁接口用于批量对指定节点的功能进行上锁操作，被上锁的功能在控制台将无法正常使用。该接口适用于以下场景：当需要临时禁用某些节点的功能以防止误操作、进行系统维护或测试时，用户可通过此接口批量对节点功能进行上锁。使用该接口的前提条件是节点功能已存在且用户具有管理员权限。上锁操作完成后，指定节点的功能将在控制台被禁用，无法进行相关操作。若节点功能不存在、用户无权限操作或请求参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchLockPoolNodes(request *model.BatchLockPoolNodesRequest) (*model.BatchLockPoolNodesResponse, error) {
	requestDef := GenReqDefForBatchLockPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchLockPoolNodesResponse), nil
	}
}

// BatchLockPoolNodesInvoker 批量对节点功能上锁
func (c *ModelArtsClient) BatchLockPoolNodesInvoker(request *model.BatchLockPoolNodesRequest) *BatchLockPoolNodesInvoker {
	requestDef := GenReqDefForBatchLockPoolNodes()
	return &BatchLockPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchMigratePoolNodes 批量迁移节点
//
// 批量迁移节点接口用于在资源池之间批量迁移节点，将节点从一个资源池迁移到另一个资源池。该接口适用于以下场景：当资源池的节点分布不均衡、需要进行集群维护或业务扩展时，用户可通过此接口将指定节点从一个资源池迁移到另一个资源池。使用该接口的前提条件是资源池中至少包含两个节点，且目标资源池具备足够的资源容量（如IP地址等）以接收迁移节点。若资源池只有一个节点、目标集群资源不足、节点状态不支持迁移或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchMigratePoolNodes(request *model.BatchMigratePoolNodesRequest) (*model.BatchMigratePoolNodesResponse, error) {
	requestDef := GenReqDefForBatchMigratePoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchMigratePoolNodesResponse), nil
	}
}

// BatchMigratePoolNodesInvoker 批量迁移节点
func (c *ModelArtsClient) BatchMigratePoolNodesInvoker(request *model.BatchMigratePoolNodesRequest) *BatchMigratePoolNodesInvoker {
	requestDef := GenReqDefForBatchMigratePoolNodes()
	return &BatchMigratePoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRebootPoolNodes 批量重启节点
//
// 批量重启节点接口用于批量重启指定资源池中的节点。该接口适用于以下场景：当需要对资源池中的节点进行系统更新、配置变更、故障恢复或维护操作时，用户可通过此接口批量重启指定的节点。使用该接口的前提条件是资源池已创建且处于可用状态，节点属于该资源池且处于运行状态，且用户具有重启节点的权限。重启操作完成后，指定的节点将被重新启动，资源池中的其他节点将继续正常运行。若资源池不存在、节点不在资源池中、节点未处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchRebootPoolNodes(request *model.BatchRebootPoolNodesRequest) (*model.BatchRebootPoolNodesResponse, error) {
	requestDef := GenReqDefForBatchRebootPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRebootPoolNodesResponse), nil
	}
}

// BatchRebootPoolNodesInvoker 批量重启节点
func (c *ModelArtsClient) BatchRebootPoolNodesInvoker(request *model.BatchRebootPoolNodesRequest) *BatchRebootPoolNodesInvoker {
	requestDef := GenReqDefForBatchRebootPoolNodes()
	return &BatchRebootPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResetPoolNodes 重置节点
//
// 重置节点接口用于将指定节点恢复到初始状态，清除节点上的数据和配置。该接口适用于以下场景：当节点出现故障、配置错误、需要重新部署或进行系统恢复时，用户可通过此接口重置节点，使其恢复到出厂或初始状态。使用该接口的前提条件是节点已存在且用户具有管理员权限。重置操作完成后，节点上的所有数据和配置将被清除，节点将被重新启动并恢复到初始状态。若节点不存在、用户无权限操作或节点处于不可重置状态（如正在运行任务），接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchResetPoolNodes(request *model.BatchResetPoolNodesRequest) (*model.BatchResetPoolNodesResponse, error) {
	requestDef := GenReqDefForBatchResetPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResetPoolNodesResponse), nil
	}
}

// BatchResetPoolNodesInvoker 重置节点
func (c *ModelArtsClient) BatchResetPoolNodesInvoker(request *model.BatchResetPoolNodesRequest) *BatchResetPoolNodesInvoker {
	requestDef := GenReqDefForBatchResetPoolNodes()
	return &BatchResetPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchResizePoolNodes 节点规格变更
//
// 节点规格变更接口用于调整指定节点的规格（如步长），例如将节点从8节点超节点扩容到16节点超节点。该接口适用于以下场景：当需要根据业务需求调整节点的资源容量、优化资源利用率或进行系统升级时，用户可通过此接口变更节点的规格。使用该接口的前提条件是节点已创建且处于可变更状态，目标规格在支持范围内，且用户具有管理员权限。规格变更完成后，节点的资源容量将按新规格调整，相关服务和配置将重新加载以适应新的规格。若节点不存在、节点状态不允许变更、目标规格不支持或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchResizePoolNodes(request *model.BatchResizePoolNodesRequest) (*model.BatchResizePoolNodesResponse, error) {
	requestDef := GenReqDefForBatchResizePoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchResizePoolNodesResponse), nil
	}
}

// BatchResizePoolNodesInvoker 节点规格变更
func (c *ModelArtsClient) BatchResizePoolNodesInvoker(request *model.BatchResizePoolNodesRequest) *BatchResizePoolNodesInvoker {
	requestDef := GenReqDefForBatchResizePoolNodes()
	return &BatchResizePoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUnbindInferApiKeys 批量解绑应用密钥
//
// 本接口用于将已绑定的apikey从指定服务中批量解绑，适用于需要撤销多个apikey对特定服务的访问权限的场景。调用此接口前，确保已获取到需要解绑的多个apikey，并确认这些apikey当前绑定在指定服务上。解绑成功后，这些apikey将不再对指定服务生效，但仍可继续用于其他服务。如果尝试解绑不存在或未绑定到指定服务的apikey，将返回相应的异常信息，提示用户检查apikey的有效性和绑定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchUnbindInferApiKeys(request *model.BatchUnbindInferApiKeysRequest) (*model.BatchUnbindInferApiKeysResponse, error) {
	requestDef := GenReqDefForBatchUnbindInferApiKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnbindInferApiKeysResponse), nil
	}
}

// BatchUnbindInferApiKeysInvoker 批量解绑应用密钥
func (c *ModelArtsClient) BatchUnbindInferApiKeysInvoker(request *model.BatchUnbindInferApiKeysRequest) *BatchUnbindInferApiKeysInvoker {
	requestDef := GenReqDefForBatchUnbindInferApiKeys()
	return &BatchUnbindInferApiKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUnlockPoolNodes 批量对节点功能解锁
//
// 批量对节点功能解锁接口用于批量解除指定节点功能的锁定状态，使被上锁的功能在控制台恢复正常可用状态。该接口适用于以下场景：当需要恢复被锁定的节点功能以正常使用、完成系统维护或测试后，用户可通过此接口批量对节点功能进行解锁。使用该接口的前提条件是节点功能已被上锁且用户具有管理员权限。解锁操作完成后，指定节点的功能将在控制台恢复正常，用户可以正常使用相关功能。若节点功能未被锁定、用户无权限操作或请求参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchUnlockPoolNodes(request *model.BatchUnlockPoolNodesRequest) (*model.BatchUnlockPoolNodesResponse, error) {
	requestDef := GenReqDefForBatchUnlockPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUnlockPoolNodesResponse), nil
	}
}

// BatchUnlockPoolNodesInvoker 批量对节点功能解锁
func (c *ModelArtsClient) BatchUnlockPoolNodesInvoker(request *model.BatchUnlockPoolNodesRequest) *BatchUnlockPoolNodesInvoker {
	requestDef := GenReqDefForBatchUnlockPoolNodes()
	return &BatchUnlockPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdatePoolNodes 批量更新节点
//
// 批量更新节点接口用于同时修改多个节点的配置或属性，支持批量操作时各节点独立执行更新流程。该接口适用于以下场景：当用户需统一升级节点软件版本、批量处理选中节点的资源标签、调整资源分配策略、应用安全补丁或同步配置变更时，可通过此接口批量更新目标节点，确保每个节点的更新过程互不影响。使用该接口的前提条件包括：目标节点已存在且用户具备管理员权限，节点需处于可操作状态（如非锁定或维护中），批量操作时需提供有效的节点列表及更新参数（如配置项、版本号等）作为输入。操作完成后，指定节点将应用新配置并更新状态为可用，原有配置将被覆盖。若节点不存在、用户权限不足、节点状态异常（如正在维护）、更新参数不合规或输入参数缺失，接口将返回对应错误信息（如404未找到节点、403权限拒绝、400参数校验失败等）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchUpdatePoolNodes(request *model.BatchUpdatePoolNodesRequest) (*model.BatchUpdatePoolNodesResponse, error) {
	requestDef := GenReqDefForBatchUpdatePoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdatePoolNodesResponse), nil
	}
}

// BatchUpdatePoolNodesInvoker 批量更新节点
func (c *ModelArtsClient) BatchUpdatePoolNodesInvoker(request *model.BatchUpdatePoolNodesRequest) *BatchUpdatePoolNodesInvoker {
	requestDef := GenReqDefForBatchUpdatePoolNodes()
	return &BatchUpdatePoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindInferApiKey 绑定应用密钥
//
// 本接口用于将生成的apikey与指定服务进行绑定，适用于应用程序需要调用特定服务的场景。调用此接口前，确保已成功创建服务实例，并获取到有效的apikey。绑定成功后，apikey将作为服务调用时的身份验证凭证，确保仅授权用户能够访问该服务。如果尝试绑定已失效的apikey，将返回相应的异常信息，提示用户检查apikey的有效性和绑定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BindInferApiKey(request *model.BindInferApiKeyRequest) (*model.BindInferApiKeyResponse, error) {
	requestDef := GenReqDefForBindInferApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindInferApiKeyResponse), nil
	}
}

// BindInferApiKeyInvoker 绑定应用密钥
func (c *ModelArtsClient) BindInferApiKeyInvoker(request *model.BindInferApiKeyRequest) *BindInferApiKeyInvoker {
	requestDef := GenReqDefForBindInferApiKey()
	return &BindInferApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelInferDeployment 中断服务部署
//
// 中断服务部署接口用于中断处于“升级中”或“部署中”状态的部署，使其快速停止。该接口适用于以下场景：当部署出现严重故障需要立即修复、资源需要快速释放以部署更高优先级的部署，或在测试环境中需要快速迭代时，用户可通过此接口中断指定部署。使用该接口的前提条件是部署当前状态为“升级中”或“部署中”，且用户具有中断部署的权限。若部署为“部署中”状态，执行中断操作，部署状态将变成“停止”，相关资源将被释放，且终端操作将被记录；若部署为“升级中”状态，执行中断操作，部署状态将变成“运行中”。若部署当前状态不是“升级中”或“部署中”，若用户无权限操作，接口将返回相应的错误信息。若部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CancelInferDeployment(request *model.CancelInferDeploymentRequest) (*model.CancelInferDeploymentResponse, error) {
	requestDef := GenReqDefForCancelInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelInferDeploymentResponse), nil
	}
}

// CancelInferDeploymentInvoker 中断服务部署
func (c *ModelArtsClient) CancelInferDeploymentInvoker(request *model.CancelInferDeploymentRequest) *CancelInferDeploymentInvoker {
	requestDef := GenReqDefForCancelInferDeployment()
	return &CancelInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeAlgorithm 更新算法
//
// 更新算法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ChangeAlgorithm(request *model.ChangeAlgorithmRequest) (*model.ChangeAlgorithmResponse, error) {
	requestDef := GenReqDefForChangeAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeAlgorithmResponse), nil
	}
}

// ChangeAlgorithmInvoker 更新算法
func (c *ModelArtsClient) ChangeAlgorithmInvoker(request *model.ChangeAlgorithmRequest) *ChangeAlgorithmInvoker {
	requestDef := GenReqDefForChangeAlgorithm()
	return &ChangeAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeTrainingExperiment 更新训练实验信息
//
// 通过实验ID更新训练实验信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ChangeTrainingExperiment(request *model.ChangeTrainingExperimentRequest) (*model.ChangeTrainingExperimentResponse, error) {
	requestDef := GenReqDefForChangeTrainingExperiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeTrainingExperimentResponse), nil
	}
}

// ChangeTrainingExperimentInvoker 更新训练实验信息
func (c *ModelArtsClient) ChangeTrainingExperimentInvoker(request *model.ChangeTrainingExperimentRequest) *ChangeTrainingExperimentInvoker {
	requestDef := GenReqDefForChangeTrainingExperiment()
	return &ChangeTrainingExperimentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeTrainingJobDescription 更新训练作业描述
//
// 更新训练作业描述。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ChangeTrainingJobDescription(request *model.ChangeTrainingJobDescriptionRequest) (*model.ChangeTrainingJobDescriptionResponse, error) {
	requestDef := GenReqDefForChangeTrainingJobDescription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeTrainingJobDescriptionResponse), nil
	}
}

// ChangeTrainingJobDescriptionInvoker 更新训练作业描述
func (c *ModelArtsClient) ChangeTrainingJobDescriptionInvoker(request *model.ChangeTrainingJobDescriptionRequest) *ChangeTrainingJobDescriptionInvoker {
	requestDef := GenReqDefForChangeTrainingJobDescription()
	return &ChangeTrainingJobDescriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckTrainingExperiment 校验训练实验名称
//
// 校验训练实验名称接口用于新增训练实验前校验训练实验名称是否重复。
// 该接口适用于以下场景：当用户需要创建新的训练实验时，可以通过此接口校验定义的实验名称是否已存在。使用该接口的前提条件是用户具有创建实验的权限。查询操作完成后，将返回实验名称是否重复的结果。若用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CheckTrainingExperiment(request *model.CheckTrainingExperimentRequest) (*model.CheckTrainingExperimentResponse, error) {
	requestDef := GenReqDefForCheckTrainingExperiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckTrainingExperimentResponse), nil
	}
}

// CheckTrainingExperimentInvoker 校验训练实验名称
func (c *ModelArtsClient) CheckTrainingExperimentInvoker(request *model.CheckTrainingExperimentRequest) *CheckTrainingExperimentInvoker {
	requestDef := GenReqDefForCheckTrainingExperiment()
	return &CheckTrainingExperimentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountInferServicesByTags 通过标签查询资源数量
//
// 该接口适用于需要统计和获取符合特定标签或资源名称条件的资源数量的场景，例如在资源管理和监控中，用户可以通过指定标签或资源名称进行精确或模糊查询来统计资源数量。通过调用此接口，用户可以基于多个标签或资源名称进行查询，若不传标签则返回所有资源的总数。用户必须具有足够的权限，且目标资源需存在。查询成功后，返回符合条件的资源总数；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CountInferServicesByTags(request *model.CountInferServicesByTagsRequest) (*model.CountInferServicesByTagsResponse, error) {
	requestDef := GenReqDefForCountInferServicesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountInferServicesByTagsResponse), nil
	}
}

// CountInferServicesByTagsInvoker 通过标签查询资源数量
func (c *ModelArtsClient) CountInferServicesByTagsInvoker(request *model.CountInferServicesByTagsRequest) *CountInferServicesByTagsInvoker {
	requestDef := GenReqDefForCountInferServicesByTags()
	return &CountInferServicesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlgorithm 创建算法
//
// 创建一个算法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateAlgorithm(request *model.CreateAlgorithmRequest) (*model.CreateAlgorithmResponse, error) {
	requestDef := GenReqDefForCreateAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlgorithmResponse), nil
	}
}

// CreateAlgorithmInvoker 创建算法
func (c *ModelArtsClient) CreateAlgorithmInvoker(request *model.CreateAlgorithmRequest) *CreateAlgorithmInvoker {
	requestDef := GenReqDefForCreateAlgorithm()
	return &CreateAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAlgorithmVersionToGallery 创建发布算法资产
//
// 创建发布算法资产接口用于在算法管理中创建并发布新的算法资产。
// 该接口适用于以下场景：当用户开发完成新的算法并希望将其发布为可复用的算法资产时，可以通过此接口创建并发布算法资产。使用该接口的前提条件是用户已登录且具有创建和发布算法资产的权限。创建发布操作完成后，系统将生成新的算法资产，并将其添加到算法资产列表中，用户可以通过算法ID进行管理和调用。若用户无权限操作、算法资产信息不完整或已存在相同名称的算法资产，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateAlgorithmVersionToGallery(request *model.CreateAlgorithmVersionToGalleryRequest) (*model.CreateAlgorithmVersionToGalleryResponse, error) {
	requestDef := GenReqDefForCreateAlgorithmVersionToGallery()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAlgorithmVersionToGalleryResponse), nil
	}
}

// CreateAlgorithmVersionToGalleryInvoker 创建发布算法资产
func (c *ModelArtsClient) CreateAlgorithmVersionToGalleryInvoker(request *model.CreateAlgorithmVersionToGalleryRequest) *CreateAlgorithmVersionToGalleryInvoker {
	requestDef := GenReqDefForCreateAlgorithmVersionToGallery()
	return &CreateAlgorithmVersionToGalleryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAuthorization 配置授权
//
// 配置授权接口用于配置ModelArts的授权。该接口适用于以下场景：当需要为IAM子用户设置访问ModelArts的权限时，管理员可通过此接口配置授权。使用该接口的前提条件是管理员具备IAM系统的Security Administrator权限，并且需要为子用户设置访问密钥。配置完成后，子用户将被授予访问ModelArts资源的权限，从而能够正常使用训练管理、开发环境、数据管理、在线服务等功能。若管理员无权限操作或子用户不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateAuthorization(request *model.CreateAuthorizationRequest) (*model.CreateAuthorizationResponse, error) {
	requestDef := GenReqDefForCreateAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorizationResponse), nil
	}
}

// CreateAuthorizationInvoker 配置授权
func (c *ModelArtsClient) CreateAuthorizationInvoker(request *model.CreateAuthorizationRequest) *CreateAuthorizationInvoker {
	requestDef := GenReqDefForCreateAuthorization()
	return &CreateAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferApiKey 创建应用密钥
//
// 本接口用于在系统中创建一个新的API_KEY，适用于需要为用户或应用程序生成访问凭证的场景。调用此接口前，确保已具备相应的创建权限，并提供必要的参数，如用户ID或应用程序ID。创建成功后，系统将生成一个唯一的API_KEY，并返回该API_KEY的详细信息，包括API_KEY值、创建时间等。如果提供的参数无效或系统中已存在相同的API_KEY，将返回相应的异常信息，提示用户检查输入数据的有效性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferApiKey(request *model.CreateInferApiKeyRequest) (*model.CreateInferApiKeyResponse, error) {
	requestDef := GenReqDefForCreateInferApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferApiKeyResponse), nil
	}
}

// CreateInferApiKeyInvoker 创建应用密钥
func (c *ModelArtsClient) CreateInferApiKeyInvoker(request *model.CreateInferApiKeyRequest) *CreateInferApiKeyInvoker {
	requestDef := GenReqDefForCreateInferApiKey()
	return &CreateInferApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferDeployment 添加部署
//
// 将模型部署为在线服务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferDeployment(request *model.CreateInferDeploymentRequest) (*model.CreateInferDeploymentResponse, error) {
	requestDef := GenReqDefForCreateInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferDeploymentResponse), nil
	}
}

// CreateInferDeploymentInvoker 添加部署
func (c *ModelArtsClient) CreateInferDeploymentInvoker(request *model.CreateInferDeploymentRequest) *CreateInferDeploymentInvoker {
	requestDef := GenReqDefForCreateInferDeployment()
	return &CreateInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferIntranetConnection 创建内网接入
//
// 本接口用于在指定Region中创建内网接入点，适用于需要为应用程序或服务配置内网连接的场景。调用此接口前，确保已具备相应的创建权限，并提供必要的参数，如Region ID、内网接入点名称和网络配置信息。创建成功后，系统将生成一个内网接入点，并返回该接入点的详细信息，包括接入点ID、创建时间、状态等。如果提供的参数无效或内网接入配置冲突，将返回相应的异常信息，提示用户检查输入数据的有效性和配置冲突情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferIntranetConnection(request *model.CreateInferIntranetConnectionRequest) (*model.CreateInferIntranetConnectionResponse, error) {
	requestDef := GenReqDefForCreateInferIntranetConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferIntranetConnectionResponse), nil
	}
}

// CreateInferIntranetConnectionInvoker 创建内网接入
func (c *ModelArtsClient) CreateInferIntranetConnectionInvoker(request *model.CreateInferIntranetConnectionRequest) *CreateInferIntranetConnectionInvoker {
	requestDef := GenReqDefForCreateInferIntranetConnection()
	return &CreateInferIntranetConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferService 创建服务
//
// 将模型部署为在线服务，适用于用户在开发或运维过程中需要将训练好的模型部署为在线服务，以便通过API或HTTP接口提供预测或处理能力的场景。调用此接口前，用户必须具有创建服务的权限，并提供合法的模型镜像路径和完整的服务配置信息（如服务名称、模型镜像路径、资源配置、升级配置等）。调用成功后，系统将成功创建并部署服务，服务状态变为“部署中”，并生成服务的唯一ID返回给用户。服务的详细信息（如状态、创建时间、更新时间等）也会记录在系统中。如果用户没有创建服务的权限，或提供的模型镜像路径不合法，或服务配置信息不完整，调用将返回相应的错误信息。如果系统在部署过程中遇到资源不足或其他内部错误，也将返回错误信息并记录日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferService(request *model.CreateInferServiceRequest) (*model.CreateInferServiceResponse, error) {
	requestDef := GenReqDefForCreateInferService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferServiceResponse), nil
	}
}

// CreateInferServiceInvoker 创建服务
func (c *ModelArtsClient) CreateInferServiceInvoker(request *model.CreateInferServiceRequest) *CreateInferServiceInvoker {
	requestDef := GenReqDefForCreateInferService()
	return &CreateInferServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferServiceTag 添加标签
//
// 该接口适用于需要为资源（如模型、数据集、服务等）添加元数据标签的场景，例如在资源管理或分类中，用户可以通过添加标签来标注资源的用途、状态或其他属性。通过调用此接口，用户可以批量添加标签，如果标签key已存在，则更新其value。用户必须具有足够的权限，且目标资源需存在。添加成功后，资源将包含新的标签信息；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferServiceTag(request *model.CreateInferServiceTagRequest) (*model.CreateInferServiceTagResponse, error) {
	requestDef := GenReqDefForCreateInferServiceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferServiceTagResponse), nil
	}
}

// CreateInferServiceTagInvoker 添加标签
func (c *ModelArtsClient) CreateInferServiceTagInvoker(request *model.CreateInferServiceTagRequest) *CreateInferServiceTagInvoker {
	requestDef := GenReqDefForCreateInferServiceTag()
	return &CreateInferServiceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferTempApiKey 创建临时应用密钥
//
// 本接口用于在系统中创建一个新的临时API_KEY，适用于需要为用户或应用程序生成临时访问凭证的场景。调用此接口前，确保已具备相应的创建权限，并提供必要的参数，如用户ID或应用程序ID。创建成功后，系统将生成一个唯一的API_KEY，并返回该API_KEY的详细信息，包括临时API_KEY值、创建时间等。如果提供的参数无效，将返回相应的异常信息，提示用户检查输入数据的有效性。
// 临时API KEY使用方法：
// **预测接口**加上两个header：
// X-Api-Key-Type&#x3D;temp
// Authorization&#x3D;临时API KEY
// **取值范围：**
// - normal：普通API KEY
// - temp：临时API KEY
// **默认取值：**
// normal。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferTempApiKey(request *model.CreateInferTempApiKeyRequest) (*model.CreateInferTempApiKeyResponse, error) {
	requestDef := GenReqDefForCreateInferTempApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferTempApiKeyResponse), nil
	}
}

// CreateInferTempApiKeyInvoker 创建临时应用密钥
func (c *ModelArtsClient) CreateInferTempApiKeyInvoker(request *model.CreateInferTempApiKeyRequest) *CreateInferTempApiKeyInvoker {
	requestDef := GenReqDefForCreateInferTempApiKey()
	return &CreateInferTempApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateModelArtsAgency 创建ModelArts委托
//
// 创建ModelArts委托接口用于创建包含OBS、SWR、IEF等依赖服务的ModelArts委托。该接口适用于以下场景：当需要配置ModelArts访问OBS、SWR、IEF等服务的权限时，用户可通过此接口创建委托。使用该接口的前提条件是用户具备创建委托的权限，并且需要在IAM系统中具备相应的权限。创建完成后，ModelArts将被授权访问OBS、SWR、IEF等服务，从而能够正常执行数据存储、镜像拉取、模型部署等功能。若用户无权限创建委托或依赖服务未配置，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateModelArtsAgency(request *model.CreateModelArtsAgencyRequest) (*model.CreateModelArtsAgencyResponse, error) {
	requestDef := GenReqDefForCreateModelArtsAgency()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateModelArtsAgencyResponse), nil
	}
}

// CreateModelArtsAgencyInvoker 创建ModelArts委托
func (c *ModelArtsClient) CreateModelArtsAgencyInvoker(request *model.CreateModelArtsAgencyRequest) *CreateModelArtsAgencyInvoker {
	requestDef := GenReqDefForCreateModelArtsAgency()
	return &CreateModelArtsAgencyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNetwork 创建网络资源
//
// 创建网络资源接口用于在系统中创建新的网络资源。该接口适用于以下场景：当需要为业务扩展、资源规划或网络架构调整时，用户可通过此接口创建新的网络资源，如虚拟网络、子网或路由等。使用该接口的前提条件是用户具有管理员权限，并且系统中具备足够的资源支持新网络资源的创建。创建操作完成后，新的网络资源将被成功添加到系统中，并可用于后续的业务配置。若用户无权限、资源不足或输入参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateNetwork(request *model.CreateNetworkRequest) (*model.CreateNetworkResponse, error) {
	requestDef := GenReqDefForCreateNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNetworkResponse), nil
	}
}

// CreateNetworkInvoker 创建网络资源
func (c *ModelArtsClient) CreateNetworkInvoker(request *model.CreateNetworkRequest) *CreateNetworkInvoker {
	requestDef := GenReqDefForCreateNetwork()
	return &CreateNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNodePool 创建节点池
//
// 创建节点池接口用于创建新的节点池。该接口适用于以下场景：当需要扩展计算资源、优化资源分配或部署新的服务时，用户可通过此接口创建指定配置的节点池。使用该接口的前提条件是用户具有管理员权限且节点池的配置参数（如节点数量、规格、网络配置等）已正确设置。创建操作完成后，节点池将被成功创建并处于可用状态，相关节点信息将被记录。若用户无权限操作、配置参数错误或系统资源不足，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateNodePool(request *model.CreateNodePoolRequest) (*model.CreateNodePoolResponse, error) {
	requestDef := GenReqDefForCreateNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNodePoolResponse), nil
	}
}

// CreateNodePoolInvoker 创建节点池
func (c *ModelArtsClient) CreateNodePoolInvoker(request *model.CreateNodePoolRequest) *CreateNodePoolInvoker {
	requestDef := GenReqDefForCreateNodePool()
	return &CreateNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrderId 创建资源池的订单id
//
// 创建资源池订单ID接口用于生成资源池申请的订单标识。该接口适用于以下场景：当用户需要申请新资源池时（如业务扩展、资源不足或临时资源需求），可通过此接口提交按需转包周期订单的创建请求。使用该接口的前提条件是用户需具备资源申请权限，提交的资源池配置参数（如资源类型、容量、周期等）需符合系统校验规则，且当前仅支持按需转包周期订单类型。订单创建成功后，系统将生成唯一订单ID并触发后续资源分配流程，同时记录操作日志。若用户权限不足、配置参数缺失/冲突（如容量超出配额）、订单类型不支持或系统资源不足，接口将返回对应错误码及提示信息，且不会生成订单ID或占用资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateOrderId(request *model.CreateOrderIdRequest) (*model.CreateOrderIdResponse, error) {
	requestDef := GenReqDefForCreateOrderId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrderIdResponse), nil
	}
}

// CreateOrderIdInvoker 创建资源池的订单id
func (c *ModelArtsClient) CreateOrderIdInvoker(request *model.CreateOrderIdRequest) *CreateOrderIdInvoker {
	requestDef := GenReqDefForCreateOrderId()
	return &CreateOrderIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePool 创建资源池
//
// 创建资源池接口用于在系统中创建新的资源池。该接口适用于以下场景：当需要为新业务分配资源、优化资源管理或进行资源隔离时，用户可通过此接口创建新的资源池，用于管理计算、存储、网络等资源。使用该接口的前提条件是用户具有管理员权限，并且系统中具备足够的资源支持新资源池的创建。创建操作完成后，新的资源池将被成功添加到系统中，并处于可用状态，可支持后续的资源分配和管理。若用户无权限、系统资源不足或输入参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreatePool(request *model.CreatePoolRequest) (*model.CreatePoolResponse, error) {
	requestDef := GenReqDefForCreatePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolResponse), nil
	}
}

// CreatePoolInvoker 创建资源池
func (c *ModelArtsClient) CreatePoolInvoker(request *model.CreatePoolRequest) *CreatePoolInvoker {
	requestDef := GenReqDefForCreatePool()
	return &CreatePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreatePoolPlugin 创建插件
//
// 创建插件实例接口用于在系统中创建一个新的插件实例。该接口适用于以下场景：当需要扩展系统功能、部署新的插件、更新现有插件配置或测试插件时，用户可通过此接口创建指定插件的实例。使用该接口的前提条件是插件已存在且用户具有管理员权限或插件管理权限。创建操作完成后，插件实例将被成功创建并处于可用状态，相关配置信息将被记录。若插件不存在、用户无权限操作、配置参数错误或系统资源不足，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreatePoolPlugin(request *model.CreatePoolPluginRequest) (*model.CreatePoolPluginResponse, error) {
	requestDef := GenReqDefForCreatePoolPlugin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePoolPluginResponse), nil
	}
}

// CreatePoolPluginInvoker 创建插件
func (c *ModelArtsClient) CreatePoolPluginInvoker(request *model.CreatePoolPluginRequest) *CreatePoolPluginInvoker {
	requestDef := GenReqDefForCreatePoolPlugin()
	return &CreatePoolPluginInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSaveImageJob 创建训练作业镜像保存任务
//
// 创建训练作业镜像保存任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateSaveImageJob(request *model.CreateSaveImageJobRequest) (*model.CreateSaveImageJobResponse, error) {
	requestDef := GenReqDefForCreateSaveImageJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSaveImageJobResponse), nil
	}
}

// CreateSaveImageJobInvoker 创建训练作业镜像保存任务
func (c *ModelArtsClient) CreateSaveImageJobInvoker(request *model.CreateSaveImageJobRequest) *CreateSaveImageJobInvoker {
	requestDef := GenReqDefForCreateSaveImageJob()
	return &CreateSaveImageJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainJobTags 创建训练作业标签
//
// 创建训练作业标签，支持批量添加，当添加的标签key已存在，则覆盖该标签的value。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateTrainJobTags(request *model.CreateTrainJobTagsRequest) (*model.CreateTrainJobTagsResponse, error) {
	requestDef := GenReqDefForCreateTrainJobTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainJobTagsResponse), nil
	}
}

// CreateTrainJobTagsInvoker 创建训练作业标签
func (c *ModelArtsClient) CreateTrainJobTagsInvoker(request *model.CreateTrainJobTagsRequest) *CreateTrainJobTagsInvoker {
	requestDef := GenReqDefForCreateTrainJobTags()
	return &CreateTrainJobTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingExperiment 创建训练实验
//
// 创建训练实验接口用于在ModelArts平台上创建新的实验分类。
// 该接口适用于以下场景：当用户需要将训练作业放入实验中分类，有序地进行管理，可以通过此接口创建训练实验，常用于多任务的版本管理等场景。使用该接口的前提条件是用户已登录ModelArts平台并具有创建训练实验的权限。创建操作完成后，系统将返回训练实验的详细信息，包括实验ID、当前实验下的训练作业总个数等。若用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateTrainingExperiment(request *model.CreateTrainingExperimentRequest) (*model.CreateTrainingExperimentResponse, error) {
	requestDef := GenReqDefForCreateTrainingExperiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingExperimentResponse), nil
	}
}

// CreateTrainingExperimentInvoker 创建训练实验
func (c *ModelArtsClient) CreateTrainingExperimentInvoker(request *model.CreateTrainingExperimentRequest) *CreateTrainingExperimentInvoker {
	requestDef := GenReqDefForCreateTrainingExperiment()
	return &CreateTrainingExperimentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTrainingJob 创建训练作业
//
// 创建训练作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateTrainingJob(request *model.CreateTrainingJobRequest) (*model.CreateTrainingJobResponse, error) {
	requestDef := GenReqDefForCreateTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTrainingJobResponse), nil
	}
}

// CreateTrainingJobInvoker 创建训练作业
func (c *ModelArtsClient) CreateTrainingJobInvoker(request *model.CreateTrainingJobRequest) *CreateTrainingJobInvoker {
	requestDef := GenReqDefForCreateTrainingJob()
	return &CreateTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkspace 创建工作空间
//
// 创建工作空间（\&quot;default\&quot;为系统预留的默认工作空间名称，不能使用）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkspace(request *model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkspaceResponse), nil
	}
}

// CreateWorkspaceInvoker 创建工作空间
func (c *ModelArtsClient) CreateWorkspaceInvoker(request *model.CreateWorkspaceRequest) *CreateWorkspaceInvoker {
	requestDef := GenReqDefForCreateWorkspace()
	return &CreateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAlgorithm 删除算法
//
// 删除算法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteAlgorithm(request *model.DeleteAlgorithmRequest) (*model.DeleteAlgorithmResponse, error) {
	requestDef := GenReqDefForDeleteAlgorithm()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAlgorithmResponse), nil
	}
}

// DeleteAlgorithmInvoker 删除算法
func (c *ModelArtsClient) DeleteAlgorithmInvoker(request *model.DeleteAlgorithmRequest) *DeleteAlgorithmInvoker {
	requestDef := GenReqDefForDeleteAlgorithm()
	return &DeleteAlgorithmInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuthorizations 删除授权
//
// 删除授权接口用于删除指定用户的授权或删除全量用户的授权。该接口适用于以下场景：当需要撤销特定用户的访问权限或在系统维护时清理所有用户的授权时，管理员可通过此接口删除指定用户的授权或全量用户的授权。使用该接口的前提条件是管理员具备删除授权的权限，并且需要指定要删除授权的用户或选择删除全量用户的授权。删除操作完成后，指定用户的授权将被移除，或所有用户的授权将被清空，用户将无法再访问相关功能。若用户不存在、管理员无权限操作或删除全量授权时系统检测到无管理员权限，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteAuthorizations(request *model.DeleteAuthorizationsRequest) (*model.DeleteAuthorizationsResponse, error) {
	requestDef := GenReqDefForDeleteAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuthorizationsResponse), nil
	}
}

// DeleteAuthorizationsInvoker 删除授权
func (c *ModelArtsClient) DeleteAuthorizationsInvoker(request *model.DeleteAuthorizationsRequest) *DeleteAuthorizationsInvoker {
	requestDef := GenReqDefForDeleteAuthorizations()
	return &DeleteAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImage 删除镜像
//
// 删除镜像接口用于删除镜像对象，对于个人私有镜像可以通过参数一并删除SWR镜像内容。该接口适用于以下场景：当镜像不再需要、配置错误或需要清理资源时，用户可通过此接口删除指定的镜像对象。使用该接口的前提条件是镜像已存在且用户具有删除权限。删除操作完成后，镜像对象将被永久移除，相关资源和配置也将被清理。若镜像不存在、用户无权限操作或镜像正在被使用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteImage(request *model.DeleteImageRequest) (*model.DeleteImageResponse, error) {
	requestDef := GenReqDefForDeleteImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageResponse), nil
	}
}

// DeleteImageInvoker 删除镜像
func (c *ModelArtsClient) DeleteImageInvoker(request *model.DeleteImageRequest) *DeleteImageInvoker {
	requestDef := GenReqDefForDeleteImage()
	return &DeleteImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteImageGroup 删除镜像组
//
// 删除镜像组接口用于删除镜像组内所有的版本对象，对于个人私有镜像可以通过参数一并删除SWR镜像内容。该接口适用于以下场景：当镜像不再需要、配置错误或需要清理资源时，用户可通过此接口删除指定的镜像组对象内所有版本。使用该接口的前提条件是镜像组已存在且用户具有删除权限。删除操作完成后，镜像组内所有版本对象将被永久移除，相关资源和配置也将被清理。若镜像组不存在、用户无权限操作或镜像正在被使用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteImageGroup(request *model.DeleteImageGroupRequest) (*model.DeleteImageGroupResponse, error) {
	requestDef := GenReqDefForDeleteImageGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteImageGroupResponse), nil
	}
}

// DeleteImageGroupInvoker 删除镜像组
func (c *ModelArtsClient) DeleteImageGroupInvoker(request *model.DeleteImageGroupRequest) *DeleteImageGroupInvoker {
	requestDef := GenReqDefForDeleteImageGroup()
	return &DeleteImageGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferApiKey 删除应用密钥
//
// 本接口用于删除指定的apikey，适用于管理员需要撤销对某个应用程序或用户的访问权限的场景。调用此接口前，确保已获取到需要删除的apikey，并确认apikey未在其他服务中使用。删除成功后，该apikey将无法再用于访问任何相关服务。如果尝试删除不存在或已删除的apikey，将返回相应的异常信息，提示用户检查apikey的有效性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferApiKey(request *model.DeleteInferApiKeyRequest) (*model.DeleteInferApiKeyResponse, error) {
	requestDef := GenReqDefForDeleteInferApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferApiKeyResponse), nil
	}
}

// DeleteInferApiKeyInvoker 删除应用密钥
func (c *ModelArtsClient) DeleteInferApiKeyInvoker(request *model.DeleteInferApiKeyRequest) *DeleteInferApiKeyInvoker {
	requestDef := GenReqDefForDeleteInferApiKey()
	return &DeleteInferApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferDeployment 删除服务部署
//
// 该接口适用于删除服务的某个部署。若服务ID、部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferDeployment(request *model.DeleteInferDeploymentRequest) (*model.DeleteInferDeploymentResponse, error) {
	requestDef := GenReqDefForDeleteInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferDeploymentResponse), nil
	}
}

// DeleteInferDeploymentInvoker 删除服务部署
func (c *ModelArtsClient) DeleteInferDeploymentInvoker(request *model.DeleteInferDeploymentRequest) *DeleteInferDeploymentInvoker {
	requestDef := GenReqDefForDeleteInferDeployment()
	return &DeleteInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferDeploymentInstance 删除服务部署的实例
//
// 本接口用于删除指定的单个部署的实例，适用于需要清理或释放不再使用的部署实例资源的场景。调用此接口前，确保已具备相应的删除权限，并提供有效的服务实例ID、部署ID。删除成功后，指定的服务部署实例将被彻底移除，不再对任何请求生效。如果提供的服务实例ID、部署ID无效、服务实例已删除或权限不足，将返回相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferDeploymentInstance(request *model.DeleteInferDeploymentInstanceRequest) (*model.DeleteInferDeploymentInstanceResponse, error) {
	requestDef := GenReqDefForDeleteInferDeploymentInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferDeploymentInstanceResponse), nil
	}
}

// DeleteInferDeploymentInstanceInvoker 删除服务部署的实例
func (c *ModelArtsClient) DeleteInferDeploymentInstanceInvoker(request *model.DeleteInferDeploymentInstanceRequest) *DeleteInferDeploymentInstanceInvoker {
	requestDef := GenReqDefForDeleteInferDeploymentInstance()
	return &DeleteInferDeploymentInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferDeploymentPod 删除Pod
//
// 本接口用于删除指定的单个Pod，适用于需要清理或释放不再使用的Pod资源的场景。调用此接口前，确保已具备相应的删除权限，并提供有效的Pod ID。删除成功后，指定的Pod将被彻底移除，不再对任何服务请求生效。如果提供的Pod ID无效、Pod已删除或权限不足，将返回相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferDeploymentPod(request *model.DeleteInferDeploymentPodRequest) (*model.DeleteInferDeploymentPodResponse, error) {
	requestDef := GenReqDefForDeleteInferDeploymentPod()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferDeploymentPodResponse), nil
	}
}

// DeleteInferDeploymentPodInvoker 删除Pod
func (c *ModelArtsClient) DeleteInferDeploymentPodInvoker(request *model.DeleteInferDeploymentPodRequest) *DeleteInferDeploymentPodInvoker {
	requestDef := GenReqDefForDeleteInferDeploymentPod()
	return &DeleteInferDeploymentPodInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferDeploymentVersion 删除在线服务部署版本
//
// 此接口用于删除指定服务部署的某个在线版本，适用于需要清理不再使用的版本或优化资源管理的场景。
// 请求需包含有效的服务ID、部署ID及版本号。用户必须具有对目标服务部署的管理权限，并且该版本当前未处于活跃状态。
// 删除成功后，指定版本将从在线服务部署中移除，相关资源将被释放。
// 若服务ID、部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；若版本处于活跃状态或有其他依赖，则返回400 Bad Request。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferDeploymentVersion(request *model.DeleteInferDeploymentVersionRequest) (*model.DeleteInferDeploymentVersionResponse, error) {
	requestDef := GenReqDefForDeleteInferDeploymentVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferDeploymentVersionResponse), nil
	}
}

// DeleteInferDeploymentVersionInvoker 删除在线服务部署版本
func (c *ModelArtsClient) DeleteInferDeploymentVersionInvoker(request *model.DeleteInferDeploymentVersionRequest) *DeleteInferDeploymentVersionInvoker {
	requestDef := GenReqDefForDeleteInferDeploymentVersion()
	return &DeleteInferDeploymentVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferServiceTag 删除资源标签
//
// 该接口适用于需要从资源（如模型、数据集、服务等）中移除特定标签的场景，例如在资源管理或分类中，用户可以通过删除标签来调整或清理资源的元数据。通过调用此接口，用户可以批量删除指定的标签。用户必须具有足够的权限，且目标资源需存在。删除成功后，资源将不再包含指定的标签信息；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferServiceTag(request *model.DeleteInferServiceTagRequest) (*model.DeleteInferServiceTagResponse, error) {
	requestDef := GenReqDefForDeleteInferServiceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferServiceTagResponse), nil
	}
}

// DeleteInferServiceTagInvoker 删除资源标签
func (c *ModelArtsClient) DeleteInferServiceTagInvoker(request *model.DeleteInferServiceTagRequest) *DeleteInferServiceTagInvoker {
	requestDef := GenReqDefForDeleteInferServiceTag()
	return &DeleteInferServiceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNetwork 删除网络资源
//
// 删除网络资源接口用于移除指定的网络资源。该接口适用于以下场景：当网络资源不再需要、配置错误或需要清理资源时，用户可通过此接口删除指定的网络资源。使用该接口的前提条件是网络资源已存在且用户具有管理员权限。删除操作完成后，指定的网络资源将被永久移除，相关配置和关联关系也将被清理。若指定的网络资源不存在、用户无权限操作或资源被其他资源依赖，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteNetwork(request *model.DeleteNetworkRequest) (*model.DeleteNetworkResponse, error) {
	requestDef := GenReqDefForDeleteNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNetworkResponse), nil
	}
}

// DeleteNetworkInvoker 删除网络资源
func (c *ModelArtsClient) DeleteNetworkInvoker(request *model.DeleteNetworkRequest) *DeleteNetworkInvoker {
	requestDef := GenReqDefForDeleteNetwork()
	return &DeleteNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNodePool 删除节点池
//
// 删除节点池接口用于移除已创建的节点池，包周期资源池不支持。该接口适用于以下场景：当节点池完成任务、配置错误或需要清理资源时，用户可通过此接口删除指定的节点池。使用该接口的前提条件是节点池已存在且用户具有管理员权限。删除操作完成后，节点池将被永久移除，相关资源和配置也将被清理。若节点池不存在、用户无权限操作或节点池处于不可删除状态（如包周期资源池或节点池正在使用中），接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteNodePool(request *model.DeleteNodePoolRequest) (*model.DeleteNodePoolResponse, error) {
	requestDef := GenReqDefForDeleteNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNodePoolResponse), nil
	}
}

// DeleteNodePoolInvoker 删除节点池
func (c *ModelArtsClient) DeleteNodePoolInvoker(request *model.DeleteNodePoolRequest) *DeleteNodePoolInvoker {
	requestDef := GenReqDefForDeleteNodePool()
	return &DeleteNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeletePool 删除资源池
//
// 删除资源池接口用于移除指定的资源池。该接口适用于以下场景：当资源池不再需要、配置错误或需要清理资源时，用户可通过此接口删除指定的资源池。使用该接口的前提条件是资源池已存在且用户具有管理员权限。删除操作完成后，指定的资源池将被永久移除，相关资源和配置也将被清理。若资源池不存在、用户无权限操作或资源池被其他资源依赖，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
	requestDef := GenReqDefForDeletePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePoolResponse), nil
	}
}

// DeletePoolInvoker 删除资源池
func (c *ModelArtsClient) DeletePoolInvoker(request *model.DeletePoolRequest) *DeletePoolInvoker {
	requestDef := GenReqDefForDeletePool()
	return &DeletePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrainJobTags 删除训练作业标签
//
// 删除训练作业标签，支持批量删除。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteTrainJobTags(request *model.DeleteTrainJobTagsRequest) (*model.DeleteTrainJobTagsResponse, error) {
	requestDef := GenReqDefForDeleteTrainJobTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrainJobTagsResponse), nil
	}
}

// DeleteTrainJobTagsInvoker 删除训练作业标签
func (c *ModelArtsClient) DeleteTrainJobTagsInvoker(request *model.DeleteTrainJobTagsRequest) *DeleteTrainJobTagsInvoker {
	requestDef := GenReqDefForDeleteTrainJobTags()
	return &DeleteTrainJobTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrainingExperiment 删除训练实验
//
// 删除训练实验接口用于移除已创建的训练实验。
// 该接口适用于以下场景：当训练实验完成、配置错误或需要清理资源时，用户可以通过此接口删除指定的训练实验。使用该接口的前提条件是训练实验已存在且用户具有删除该实验的权限。删除操作完成后，训练实验将被永久移除，相关的配置和资源也将被清理。若训练实验不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteTrainingExperiment(request *model.DeleteTrainingExperimentRequest) (*model.DeleteTrainingExperimentResponse, error) {
	requestDef := GenReqDefForDeleteTrainingExperiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrainingExperimentResponse), nil
	}
}

// DeleteTrainingExperimentInvoker 删除训练实验
func (c *ModelArtsClient) DeleteTrainingExperimentInvoker(request *model.DeleteTrainingExperimentRequest) *DeleteTrainingExperimentInvoker {
	requestDef := GenReqDefForDeleteTrainingExperiment()
	return &DeleteTrainingExperimentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTrainingJob 删除训练作业
//
// 删除训练作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteTrainingJob(request *model.DeleteTrainingJobRequest) (*model.DeleteTrainingJobResponse, error) {
	requestDef := GenReqDefForDeleteTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTrainingJobResponse), nil
	}
}

// DeleteTrainingJobInvoker 删除训练作业
func (c *ModelArtsClient) DeleteTrainingJobInvoker(request *model.DeleteTrainingJobRequest) *DeleteTrainingJobInvoker {
	requestDef := GenReqDefForDeleteTrainingJob()
	return &DeleteTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkspace 删除工作空间
//
// 删除工作空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteWorkspace(request *model.DeleteWorkspaceRequest) (*model.DeleteWorkspaceResponse, error) {
	requestDef := GenReqDefForDeleteWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkspaceResponse), nil
	}
}

// DeleteWorkspaceInvoker 删除工作空间
func (c *ModelArtsClient) DeleteWorkspaceInvoker(request *model.DeleteWorkspaceRequest) *DeleteWorkspaceInvoker {
	requestDef := GenReqDefForDeleteWorkspace()
	return &DeleteWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachDynamicStorage 动态卸载Notebook存储
//
// 动态卸载Notebook存储接口用于从运行中的Notebook实例中卸载已挂载的动态存储实例。
//
// 适用场景：用户需要清理或重新组织Notebook实例的挂载资源时，可通过此接口卸载指定的存储实例。使用该接口的前提条件是用户已登录系统并具有访问目标Notebook实例的权限，同时Notebook实例必须处于运行状态且存储实例处于MOUNTED / UNMOUNT_FAILED / MOUNT_FAILED状态。调用该接口后，系统将卸载指定的存储实例，Notebook容器将无法再操作存储中的文件或对象，但存储中的文件或对象保持不变。若用户无权限访问指定实例或Notebook实例未运行，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DetachDynamicStorage(request *model.DetachDynamicStorageRequest) (*model.DetachDynamicStorageResponse, error) {
	requestDef := GenReqDefForDetachDynamicStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachDynamicStorageResponse), nil
	}
}

// DetachDynamicStorageInvoker 动态卸载Notebook存储
func (c *ModelArtsClient) DetachDynamicStorageInvoker(request *model.DetachDynamicStorageRequest) *DetachDynamicStorageInvoker {
	requestDef := GenReqDefForDetachDynamicStorage()
	return &DetachDynamicStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetAuthorizations 查看授权列表
//
// 查看授权列表接口用于查看授权信息。该接口适用于以下场景：当用户需要了解当前的授权情况、审核权限分配或管理权限时，可通过此接口查看授权列表。使用该接口的前提条件是用户具备查看授权的权限。查看操作完成后，将返回授权列表，包括被授权的资源、授权类型以及授权内容等信息。若用户无权限查看或授权列表不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetAuthorizations(request *model.GetAuthorizationsRequest) (*model.GetAuthorizationsResponse, error) {
	requestDef := GenReqDefForGetAuthorizations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetAuthorizationsResponse), nil
	}
}

// GetAuthorizationsInvoker 查看授权列表
func (c *ModelArtsClient) GetAuthorizationsInvoker(request *model.GetAuthorizationsRequest) *GetAuthorizationsInvoker {
	requestDef := GenReqDefForGetAuthorizations()
	return &GetAuthorizationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetHyperinstanceOperation 查询超节点Operation详情
//
// 查询Operation详情接口用于获取指定Operation的详细信息。该接口适用于以下场景：当用户需要了解某个Operation的具体执行情况和状态，以便进行故障排查或操作审计时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询Operation详情的权限，且指定的Operation已存在。查询操作完成后，接口将返回指定Operation的详细信息，包括Operation ID、操作类型、执行状态、开始时间、结束时间、操作结果等。若用户无权限操作、指定的Operation不存在或Operation ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetHyperinstanceOperation(request *model.GetHyperinstanceOperationRequest) (*model.GetHyperinstanceOperationResponse, error) {
	requestDef := GenReqDefForGetHyperinstanceOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetHyperinstanceOperationResponse), nil
	}
}

// GetHyperinstanceOperationInvoker 查询超节点Operation详情
func (c *ModelArtsClient) GetHyperinstanceOperationInvoker(request *model.GetHyperinstanceOperationRequest) *GetHyperinstanceOperationInvoker {
	requestDef := GenReqDefForGetHyperinstanceOperation()
	return &GetHyperinstanceOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAlgorithms 查询算法列表
//
// 查询算法列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListAlgorithms(request *model.ListAlgorithmsRequest) (*model.ListAlgorithmsResponse, error) {
	requestDef := GenReqDefForListAlgorithms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlgorithmsResponse), nil
	}
}

// ListAlgorithmsInvoker 查询算法列表
func (c *ModelArtsClient) ListAlgorithmsInvoker(request *model.ListAlgorithmsRequest) *ListAlgorithmsInvoker {
	requestDef := GenReqDefForListAlgorithms()
	return &ListAlgorithmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDynamicStorages 获取动态挂载存储信息列表
//
// 此接口用于获取指定Notebook实例下挂载的动态存储信息列表。
// 适用场景：用户需要获取指定Notebook实例下挂载的动态存储的存储id、存储类型、挂载路径、挂载状态等信息的场景。若挂载失败，会返回相应错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDynamicStorages(request *model.ListDynamicStoragesRequest) (*model.ListDynamicStoragesResponse, error) {
	requestDef := GenReqDefForListDynamicStorages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDynamicStoragesResponse), nil
	}
}

// ListDynamicStoragesInvoker 获取动态挂载存储信息列表
func (c *ModelArtsClient) ListDynamicStoragesInvoker(request *model.ListDynamicStoragesRequest) *ListDynamicStoragesInvoker {
	requestDef := GenReqDefForListDynamicStorages()
	return &ListDynamicStoragesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEventCategories 获取事件类型列表
//
// 获取事件类型列表接口用于获取训练管理中支持的事件类型列表。
// 该接口适用于以下场景：当用户需要了解训练管理中支持的事件类型，以便在创建或管理训练任务时进行相关配置时，可以通过此接口获取事件类型列表。使用该接口的前提条件是用户已登录且具有访问训练管理的权限。获取操作完成后，响应消息体中将包含所有支持的事件类型及其描述。若用户无权限访问或系统中无事件类型信息，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListEventCategories(request *model.ListEventCategoriesRequest) (*model.ListEventCategoriesResponse, error) {
	requestDef := GenReqDefForListEventCategories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventCategoriesResponse), nil
	}
}

// ListEventCategoriesInvoker 获取事件类型列表
func (c *ModelArtsClient) ListEventCategoriesInvoker(request *model.ListEventCategoriesRequest) *ListEventCategoriesInvoker {
	requestDef := GenReqDefForListEventCategories()
	return &ListEventCategoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEvents 查询事件列表
//
// 查询事件列表接口用于获取系统中记录的事件信息。该接口适用于以下场景：当用户需要监控系统状态、排查问题或进行审计时，可通过此接口查询系统中发生的事件记录。使用该接口的前提条件是用户具有相应的权限，并且系统中已存在事件记录。查询操作完成后，接口将返回事件列表，包含事件ID、类型、时间、描述等信息。若用户无权限、事件记录不存在或查询参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListEvents(request *model.ListEventsRequest) (*model.ListEventsResponse, error) {
	requestDef := GenReqDefForListEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEventsResponse), nil
	}
}

// ListEventsInvoker 查询事件列表
func (c *ModelArtsClient) ListEventsInvoker(request *model.ListEventsRequest) *ListEventsInvoker {
	requestDef := GenReqDefForListEvents()
	return &ListEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImage 查询支持的镜像列表
//
// 查询支持的镜像列表接口用于根据指定条件分页查询满足条件的所有镜像。该接口适用于以下场景：当用户需要查找特定镜像、管理镜像仓库或选择合适的镜像版本进行部署时，可通过此接口获取符合条件的镜像列表。使用该接口的前提条件是镜像仓库已存在且用户具有访问权限。查询操作完成后，将返回满足条件的镜像列表，包括镜像ID、名称、版本、类型、状态、大小和创建时间等详细信息。若镜像仓库不存在、用户无权限访问或查询条件有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListImage(request *model.ListImageRequest) (*model.ListImageResponse, error) {
	requestDef := GenReqDefForListImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageResponse), nil
	}
}

// ListImageInvoker 查询支持的镜像列表
func (c *ModelArtsClient) ListImageInvoker(request *model.ListImageRequest) *ListImageInvoker {
	requestDef := GenReqDefForListImage()
	return &ListImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListImageGroup 查询用户镜像列表
//
// 查询用户镜像列表接口用于查询用户镜像信息概览，以镜像名称作为聚合的信息。该接口适用于以下场景：当用户需要管理多个镜像或了解各镜像的基本信息时，可通过此接口获取镜像列表及其概览信息。使用该接口的前提条件是用户具备镜像管理权限，并且镜像已存在。查询操作完成后，将返回用户所有镜像的列表，包括镜像名称、版本、状态等信息。若镜像不存在或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListImageGroup(request *model.ListImageGroupRequest) (*model.ListImageGroupResponse, error) {
	requestDef := GenReqDefForListImageGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListImageGroupResponse), nil
	}
}

// ListImageGroupInvoker 查询用户镜像列表
func (c *ModelArtsClient) ListImageGroupInvoker(request *model.ListImageGroupRequest) *ListImageGroupInvoker {
	requestDef := GenReqDefForListImageGroup()
	return &ListImageGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferApiKeys 查询应用密钥
//
// 本接口用于查询当前系统中的apikey列表，适用于管理员或用户需要查看和管理apikey的场景。调用此接口前，确保已具备相应的查询权限。返回的列表将包含每个apikey的基本信息，如apikey值、创建时间、绑定的服务等。如果当前系统中没有apikey，将返回空列表或相应的异常信息，提示用户检查查询条件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferApiKeys(request *model.ListInferApiKeysRequest) (*model.ListInferApiKeysResponse, error) {
	requestDef := GenReqDefForListInferApiKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferApiKeysResponse), nil
	}
}

// ListInferApiKeysInvoker 查询应用密钥
func (c *ModelArtsClient) ListInferApiKeysInvoker(request *model.ListInferApiKeysRequest) *ListInferApiKeysInvoker {
	requestDef := GenReqDefForListInferApiKeys()
	return &ListInferApiKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferClusterFlavors 查询支持可切换规格列表
//
// 该接口允许用户查询当前资源实例支持的可切换规格列表，适用于需要调整实例资源配置的场景。使用该接口前，用户需确保已登录并拥有查询权限。执行成功后，用户将获得一个包含各种可切换规格的详细列表，包括规格ID、名称、资源配额等信息，可用于后续的实例规格变更操作。如果用户没有相应的查询权限或资源实例ID无效，接口将返回错误信息，如401 Unauthorized或404 Not Found。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferClusterFlavors(request *model.ListInferClusterFlavorsRequest) (*model.ListInferClusterFlavorsResponse, error) {
	requestDef := GenReqDefForListInferClusterFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferClusterFlavorsResponse), nil
	}
}

// ListInferClusterFlavorsInvoker 查询支持可切换规格列表
func (c *ModelArtsClient) ListInferClusterFlavorsInvoker(request *model.ListInferClusterFlavorsRequest) *ListInferClusterFlavorsInvoker {
	requestDef := GenReqDefForListInferClusterFlavors()
	return &ListInferClusterFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeploymentInstances 查询服务部署实例列表
//
// 本接口用于查询当前[租户](tag:hws,hws_hk,fcs,fcs_super)[资源空间](tag:hcs,hcs_sm)的服务部署实例列表，并支持根据服务部署实例的状态进行筛选，包括运行中和已删除状态，同时支持分页和关键词筛选。适用于需要管理和监控服务实例状态的场景。调用此接口前，确保已具备相应的查询权限，并提供可选的筛选条件和分页参数。返回的列表将包含每个服务部署实例的基本信息，如部署名字、最新更新时间、状态等。如果当前租户没有符合条件的服务实例或提供的参数无效，将返回空列表或相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeploymentInstances(request *model.ListInferDeploymentInstancesRequest) (*model.ListInferDeploymentInstancesResponse, error) {
	requestDef := GenReqDefForListInferDeploymentInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentInstancesResponse), nil
	}
}

// ListInferDeploymentInstancesInvoker 查询服务部署实例列表
func (c *ModelArtsClient) ListInferDeploymentInstancesInvoker(request *model.ListInferDeploymentInstancesRequest) *ListInferDeploymentInstancesInvoker {
	requestDef := GenReqDefForListInferDeploymentInstances()
	return &ListInferDeploymentInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeploymentPodEvents 查询Pod事件
//
// 本接口用于查询指定Pod的Kubernetes事件，适用于需要监控和排查Pod运行状态的场景。调用此接口前，确保已具备相应的查询权限，并提供有效的Pod ID。返回的事件列表将包含每个事件的详细信息，如事件类型、发生次数、事件名称、事件信息、发生时间等。如果提供的Pod ID无效、Pod不存在或权限不足，将返回相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeploymentPodEvents(request *model.ListInferDeploymentPodEventsRequest) (*model.ListInferDeploymentPodEventsResponse, error) {
	requestDef := GenReqDefForListInferDeploymentPodEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentPodEventsResponse), nil
	}
}

// ListInferDeploymentPodEventsInvoker 查询Pod事件
func (c *ModelArtsClient) ListInferDeploymentPodEventsInvoker(request *model.ListInferDeploymentPodEventsRequest) *ListInferDeploymentPodEventsInvoker {
	requestDef := GenReqDefForListInferDeploymentPodEvents()
	return &ListInferDeploymentPodEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeploymentPods 查询服务部署的pod的列表
//
// 本接口用于查询指定服务部署的pod列表，并支持选择是否只获取当前运行中的pod。适用于需要管理和监控服务部署pod状态的场景。调用此接口前，确保已具备相应的查询权限，并提供有效的服务ID、部署ID和可选的筛选参数pod status（如是否只获取当前运行中的pod）。返回的列表将包含每个pod的基本信息，如pod名称、pod所在node的IP、pod所在node名字、pod角色、状态、最近更新时间等。如果指定的服务ID无效或当前服务没有pod，将返回空列表或相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeploymentPods(request *model.ListInferDeploymentPodsRequest) (*model.ListInferDeploymentPodsResponse, error) {
	requestDef := GenReqDefForListInferDeploymentPods()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentPodsResponse), nil
	}
}

// ListInferDeploymentPodsInvoker 查询服务部署的pod的列表
func (c *ModelArtsClient) ListInferDeploymentPodsInvoker(request *model.ListInferDeploymentPodsRequest) *ListInferDeploymentPodsInvoker {
	requestDef := GenReqDefForListInferDeploymentPods()
	return &ListInferDeploymentPodsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeploymentVersions 查询在线服务部署版本列表
//
// 此接口用于获取指定服务部署的版本列表，适用于需要了解当前服务部署可用版本的场景，例如进行版本选择或确认当前版本信息。请求需包含有效的服务ID、部署ID，也可通过排序参数对列表进行排序。用户必须具有对目标服务部署的查看权限。请求成功后，返回该服务部署的所有在线版本信息，包括版本号、发布时间和状态。若服务ID/部署ID无效或用户无权限，则返回400 Bad Request或403 Forbidden；若服务部署无在线版本，则返回空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeploymentVersions(request *model.ListInferDeploymentVersionsRequest) (*model.ListInferDeploymentVersionsResponse, error) {
	requestDef := GenReqDefForListInferDeploymentVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentVersionsResponse), nil
	}
}

// ListInferDeploymentVersionsInvoker 查询在线服务部署版本列表
func (c *ModelArtsClient) ListInferDeploymentVersionsInvoker(request *model.ListInferDeploymentVersionsRequest) *ListInferDeploymentVersionsInvoker {
	requestDef := GenReqDefForListInferDeploymentVersions()
	return &ListInferDeploymentVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeployments 查询服务部署列表
//
// 支持分页和筛选，适用于用户在管理控制台或通过API需要查看特定条件下（如服务状态、名称等）的部署列表的情况。调用此接口前，用户必须具有查询部署列表的权限，并提供合法的分页参数（如页码、每页条数）和筛选条件（如部署状态、名称等）。调用成功后，系统将返回符合筛选条件的部署列表，包含指定页码的数据，并返回总页数和总记录数。如果用户没有查询部署列表的权限，或提供的分页参数和筛选条件不合法，调用将返回相应的错误信息。如果系统在查询过程中遇到内部错误，也将返回错误信息并记录日志。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeployments(request *model.ListInferDeploymentsRequest) (*model.ListInferDeploymentsResponse, error) {
	requestDef := GenReqDefForListInferDeployments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentsResponse), nil
	}
}

// ListInferDeploymentsInvoker 查询服务部署列表
func (c *ModelArtsClient) ListInferDeploymentsInvoker(request *model.ListInferDeploymentsRequest) *ListInferDeploymentsInvoker {
	requestDef := GenReqDefForListInferDeployments()
	return &ListInferDeploymentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferIntranetConnectionApplications 查询当前的内网接入申请列表
//
// 本接口用于查询当前所有的内网接入申请记录，适用于需要管理和监控内网接入申请状态的场景。调用此接口前，确保已具备相应的查询权限。返回的列表将包含每个内网接入申请的基本信息，如申请ID、创建时间、状态、Region ID等。如果当前租户没有内网接入申请记录，将返回空列表。如果调用时出现权限不足或其他系统异常，将返回相应的异常信息，提示用户检查权限或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferIntranetConnectionApplications(request *model.ListInferIntranetConnectionApplicationsRequest) (*model.ListInferIntranetConnectionApplicationsResponse, error) {
	requestDef := GenReqDefForListInferIntranetConnectionApplications()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferIntranetConnectionApplicationsResponse), nil
	}
}

// ListInferIntranetConnectionApplicationsInvoker 查询当前的内网接入申请列表
func (c *ModelArtsClient) ListInferIntranetConnectionApplicationsInvoker(request *model.ListInferIntranetConnectionApplicationsRequest) *ListInferIntranetConnectionApplicationsInvoker {
	requestDef := GenReqDefForListInferIntranetConnectionApplications()
	return &ListInferIntranetConnectionApplicationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferIntranetConnectionReviews 查询当前的内网接入审批列表
//
// 本接口用于查询当前所有的内网接入审批记录，适用于需要管理和监控内网接入审批状态的场景。调用此接口前，确保已具备相应的查询权限。返回的列表将包含每个内网接入审批的基本信息，如审批ID、申请时间、状态（如待审批、已批准、已拒绝）、申请者信息、Region ID等。如果当前租户没有内网接入审批记录，将返回空列表。如果调用时出现权限不足或其他系统异常，将返回相应的异常信息，提示用户检查权限或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferIntranetConnectionReviews(request *model.ListInferIntranetConnectionReviewsRequest) (*model.ListInferIntranetConnectionReviewsResponse, error) {
	requestDef := GenReqDefForListInferIntranetConnectionReviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferIntranetConnectionReviewsResponse), nil
	}
}

// ListInferIntranetConnectionReviewsInvoker 查询当前的内网接入审批列表
func (c *ModelArtsClient) ListInferIntranetConnectionReviewsInvoker(request *model.ListInferIntranetConnectionReviewsRequest) *ListInferIntranetConnectionReviewsInvoker {
	requestDef := GenReqDefForListInferIntranetConnectionReviews()
	return &ListInferIntranetConnectionReviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferServiceEvents 获取在线服务事件列表
//
// 该接口适用于需要监控和管理在线服务事件的场景，例如用户或运维人员需要定期检查服务的日志事件，以及时发现和处理问题。通过调用此接口，用户可以获取当前在线服务的所有事件记录，包括事件类型、事件信息、时间、发生次数等信息。用户必须具有查询服务事件列表的权限，才能成功访问该接口。获取成功后，返回事件列表；若失败，返回具体的错误信息。常见异常包括权限验证错误、服务状态错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferServiceEvents(request *model.ListInferServiceEventsRequest) (*model.ListInferServiceEventsResponse, error) {
	requestDef := GenReqDefForListInferServiceEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferServiceEventsResponse), nil
	}
}

// ListInferServiceEventsInvoker 获取在线服务事件列表
func (c *ModelArtsClient) ListInferServiceEventsInvoker(request *model.ListInferServiceEventsRequest) *ListInferServiceEventsInvoker {
	requestDef := GenReqDefForListInferServiceEvents()
	return &ListInferServiceEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferServiceTags 查询某一类资源下的标签
//
// 该接口适用于需要获取用户当前项目中某一类资源（如指定的Service）的标签信息的场景，例如在资源管理和监控中，用户可以通过查询标签来了解各类资源的分类和属性。通过调用此接口，用户可以获取指定Service在所有工作空间中的标签列表，但无权限的工作空间标签数据将被过滤不返回。用户必须具有足够的权限，且目标资源需存在。查询成功后，返回指定Service的标签列表；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferServiceTags(request *model.ListInferServiceTagsRequest) (*model.ListInferServiceTagsResponse, error) {
	requestDef := GenReqDefForListInferServiceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferServiceTagsResponse), nil
	}
}

// ListInferServiceTagsInvoker 查询某一类资源下的标签
func (c *ModelArtsClient) ListInferServiceTagsInvoker(request *model.ListInferServiceTagsRequest) *ListInferServiceTagsInvoker {
	requestDef := GenReqDefForListInferServiceTags()
	return &ListInferServiceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferServices 查询服务列表
//
// 支持分页和筛选
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferServices(request *model.ListInferServicesRequest) (*model.ListInferServicesResponse, error) {
	requestDef := GenReqDefForListInferServices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferServicesResponse), nil
	}
}

// ListInferServicesInvoker 查询服务列表
func (c *ModelArtsClient) ListInferServicesInvoker(request *model.ListInferServicesRequest) *ListInferServicesInvoker {
	requestDef := GenReqDefForListInferServices()
	return &ListInferServicesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferServicesByTags 通过标签反查资源列表
//
// 该接口适用于需要根据标签或资源名称查找相关资源的场景，例如在资源管理和搜索中，用户可以通过指定标签或进行模糊查询来查找符合特定条件的资源。通过调用此接口，用户可以基于多个标签或资源名称进行精确或模糊查询，若不传标签则返回所有资源。用户必须具有足够的权限，且目标资源需存在。查询成功后，返回符合条件的资源列表；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferServicesByTags(request *model.ListInferServicesByTagsRequest) (*model.ListInferServicesByTagsResponse, error) {
	requestDef := GenReqDefForListInferServicesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferServicesByTagsResponse), nil
	}
}

// ListInferServicesByTagsInvoker 通过标签反查资源列表
func (c *ModelArtsClient) ListInferServicesByTagsInvoker(request *model.ListInferServicesByTagsRequest) *ListInferServicesByTagsInvoker {
	requestDef := GenReqDefForListInferServicesByTags()
	return &ListInferServicesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListJobs 查询任务列表
//
// 查询任务列表接口用于获取当前用户下的任务列表。该接口适用于以下场景：当需要查看任务状态、管理任务进度或统计任务数量时，用户可通过此接口获取当前用户下所有任务的详细信息。使用该接口的前提条件是用户已登录系统且具有查看任务的权限。调用接口成功后，系统将返回当前用户下的任务列表，包括任务ID、名称、状态、创建时间等信息。若用户未登录、无权限访问或系统中未配置任务，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListJobs(request *model.ListJobsRequest) (*model.ListJobsResponse, error) {
	requestDef := GenReqDefForListJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListJobsResponse), nil
	}
}

// ListJobsInvoker 查询任务列表
func (c *ModelArtsClient) ListJobsInvoker(request *model.ListJobsRequest) *ListJobsInvoker {
	requestDef := GenReqDefForListJobs()
	return &ListJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetworks 查询网络资源列表
//
// 查询网络资源列表接口用于获取系统中已创建的网络资源信息。该接口适用于以下场景：当用户需要监控网络状态、进行资源规划、排查网络问题或进行审计时，可通过此接口查询系统中现有的网络资源列表。使用该接口的前提条件是用户具有相应的权限，并且系统中已存在网络资源。查询操作完成后，接口将返回网络资源列表，包含资源ID、类型、状态、创建时间等详细信息。若用户无权限、系统中无网络资源或查询参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListNetworks(request *model.ListNetworksRequest) (*model.ListNetworksResponse, error) {
	requestDef := GenReqDefForListNetworks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworksResponse), nil
	}
}

// ListNetworksInvoker 查询网络资源列表
func (c *ModelArtsClient) ListNetworksInvoker(request *model.ListNetworksRequest) *ListNetworksInvoker {
	requestDef := GenReqDefForListNetworks()
	return &ListNetworksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodePoolNodes 查询节点池的节点列表
//
// 查询节点池的节点列表接口用于获取指定节点池中所有节点的详细信息。该接口适用于以下场景：当需要查看节点池的节点状态、资源使用情况或管理节点资源时，用户可通过此接口获取节点池中节点的详细信息。使用该接口的前提条件是节点池已存在且用户具有访问该节点池的权限。调用接口成功后，系统将返回节点池中所有节点的列表，包括节点ID、名称、状态、IP地址、资源使用情况等详细信息。若节点池不存在、用户无权限访问或节点池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListNodePoolNodes(request *model.ListNodePoolNodesRequest) (*model.ListNodePoolNodesResponse, error) {
	requestDef := GenReqDefForListNodePoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodePoolNodesResponse), nil
	}
}

// ListNodePoolNodesInvoker 查询节点池的节点列表
func (c *ModelArtsClient) ListNodePoolNodesInvoker(request *model.ListNodePoolNodesRequest) *ListNodePoolNodesInvoker {
	requestDef := GenReqDefForListNodePoolNodes()
	return &ListNodePoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNodePools 查询节点池列表
//
// 查询节点池列表接口用于获取指定节点池的列表信息。该接口适用于以下场景：当需要查看节点池的配置、状态或管理节点池资源时，用户可通过此接口获取节点池的详细信息。使用该接口的前提条件是节点池已存在且用户具有管理员权限。调用接口成功后，系统将返回节点池的列表，包括节点池ID、名称、节点数量、状态等详细信息。若节点池不存在、用户无权限操作或节点池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListNodePools(request *model.ListNodePoolsRequest) (*model.ListNodePoolsResponse, error) {
	requestDef := GenReqDefForListNodePools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodePoolsResponse), nil
	}
}

// ListNodePoolsInvoker 查询节点池列表
func (c *ModelArtsClient) ListNodePoolsInvoker(request *model.ListNodePoolsRequest) *ListNodePoolsInvoker {
	requestDef := GenReqDefForListNodePools()
	return &ListNodePoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPluginTemplates 查询插件模板列表
//
// 查询插件模板列表接口用于获取插件模板的基本信息列表。该接口适用于以下场景：当需要浏览或管理插件模板时，用户可通过此接口查询所有可用的插件模板信息，以便选择或进一步操作。使用该接口的前提条件是用户具有访问插件模板的权限，且插件模板服务处于正常运行状态。查询操作完成后，用户将获得插件模板的列表信息，包括模板名称、类型、版本等，便于后续的插件开发或管理。若用户无权限访问、插件模板服务不可用或请求参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListPluginTemplates(request *model.ListPluginTemplatesRequest) (*model.ListPluginTemplatesResponse, error) {
	requestDef := GenReqDefForListPluginTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPluginTemplatesResponse), nil
	}
}

// ListPluginTemplatesInvoker 查询插件模板列表
func (c *ModelArtsClient) ListPluginTemplatesInvoker(request *model.ListPluginTemplatesRequest) *ListPluginTemplatesInvoker {
	requestDef := GenReqDefForListPluginTemplates()
	return &ListPluginTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoolNodes 查询资源池节点列表
//
// 查询资源池节点列表接口用于获取指定资源池中的节点信息列表。该接口适用于以下场景：当需要了解资源池中节点的分布、状态或资源使用情况时，用户可通过此接口查询资源池中的节点列表，以便进行资源监控、分配或管理。使用该接口的前提条件是资源池已创建且处于可用状态，且用户具有访问资源池的权限。查询操作完成后，用户将获得资源池中节点的详细信息，包括节点ID、状态、资源使用情况等，便于后续的资源管理和优化。若资源池不存在、用户无权限访问或请求参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListPoolNodes(request *model.ListPoolNodesRequest) (*model.ListPoolNodesResponse, error) {
	requestDef := GenReqDefForListPoolNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolNodesResponse), nil
	}
}

// ListPoolNodesInvoker 查询资源池节点列表
func (c *ModelArtsClient) ListPoolNodesInvoker(request *model.ListPoolNodesRequest) *ListPoolNodesInvoker {
	requestDef := GenReqDefForListPoolNodes()
	return &ListPoolNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoolPlugins 查询插件列表
//
// 查询插件实例列表接口用于获取系统中已部署的插件实例信息。该接口适用于以下场景：当用户需要查看系统中已部署的插件实例、监控插件运行状态、管理插件配置或进行故障排查时，可通过此接口获取插件实例的详细信息。使用该接口的前提条件是用户具有查询权限且系统中已部署至少一个插件实例。调用该接口后，系统将返回所有插件实例的列表，包括插件名称、类型、状态、版本及部署环境等信息。若用户无权限访问或系统中未部署任何插件实例，接口将返回相应的错误信息或空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListPoolPlugins(request *model.ListPoolPluginsRequest) (*model.ListPoolPluginsResponse, error) {
	requestDef := GenReqDefForListPoolPlugins()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolPluginsResponse), nil
	}
}

// ListPoolPluginsInvoker 查询插件列表
func (c *ModelArtsClient) ListPoolPluginsInvoker(request *model.ListPoolPluginsRequest) *ListPoolPluginsInvoker {
	requestDef := GenReqDefForListPoolPlugins()
	return &ListPoolPluginsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPoolTags 查询资源池的所有标签
//
// 查询资源池所有标签接口用于获取用户当前项目下资源池的所有标签信息，默认查询所有工作空间，但无权限的工作空间不会返回标签数据。该接口适用于以下场景：当需要管理、分类或统计资源池的标签信息时，用户可通过此接口获取资源池的标签数据。使用该接口的前提条件是用户具有访问资源池的权限且资源池已存在。调用接口成功后，系统将返回用户当前项目下所有可访问工作空间的资源池标签信息。若用户无权限访问资源池、资源池不存在或项目未创建，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListPoolTags(request *model.ListPoolTagsRequest) (*model.ListPoolTagsResponse, error) {
	requestDef := GenReqDefForListPoolTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolTagsResponse), nil
	}
}

// ListPoolTagsInvoker 查询资源池的所有标签
func (c *ModelArtsClient) ListPoolTagsInvoker(request *model.ListPoolTagsRequest) *ListPoolTagsInvoker {
	requestDef := GenReqDefForListPoolTags()
	return &ListPoolTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPools 查询资源池列表
//
// 查询资源池列表接口用于获取系统中已创建的资源池信息。该接口适用于以下场景：当用户需要监控资源池状态、进行资源规划、管理资源分配或进行审计时，可通过此接口查询系统中现有的资源池列表。使用该接口的前提条件是用户具有相应的权限，并且系统中已存在资源池。查询操作完成后，接口将返回资源池列表，包含资源池ID、名称、类型、状态、资源配额等详细信息。若用户无权限、系统中无资源池或查询参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListPools(request *model.ListPoolsRequest) (*model.ListPoolsResponse, error) {
	requestDef := GenReqDefForListPools()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPoolsResponse), nil
	}
}

// ListPoolsInvoker 查询资源池列表
func (c *ModelArtsClient) ListPoolsInvoker(request *model.ListPoolsRequest) *ListPoolsInvoker {
	requestDef := GenReqDefForListPools()
	return &ListPoolsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceFlavors 查询资源规格列表
//
// 查询资源规格列表接口用于获取可用的资源规格信息。该接口适用于以下场景：当需要查看或选择资源规格以创建资源池、分配资源或了解可用资源规格时，用户可通过此接口获取资源规格的详细信息。使用该接口的前提条件是用户具有相应的权限（如管理员权限或资源管理权限）。调用接口成功后，系统将返回资源规格的列表，包括规格ID、名称、CPU核数、内存大小、存储容量等详细信息。若用户无权限访问该接口或系统中未配置资源规格，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListResourceFlavors(request *model.ListResourceFlavorsRequest) (*model.ListResourceFlavorsResponse, error) {
	requestDef := GenReqDefForListResourceFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceFlavorsResponse), nil
	}
}

// ListResourceFlavorsInvoker 查询资源规格列表
func (c *ModelArtsClient) ListResourceFlavorsInvoker(request *model.ListResourceFlavorsRequest) *ListResourceFlavorsInvoker {
	requestDef := GenReqDefForListResourceFlavors()
	return &ListResourceFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListScheduledEvents 查询计划事件列表
//
// 查询计划事件列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListScheduledEvents(request *model.ListScheduledEventsRequest) (*model.ListScheduledEventsResponse, error) {
	requestDef := GenReqDefForListScheduledEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListScheduledEventsResponse), nil
	}
}

// ListScheduledEventsInvoker 查询计划事件列表
func (c *ModelArtsClient) ListScheduledEventsInvoker(request *model.ListScheduledEventsRequest) *ListScheduledEventsInvoker {
	requestDef := GenReqDefForListScheduledEvents()
	return &ListScheduledEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrainingExperiments 查询训练实验列表
//
// 查询训练实验列表接口用于获取ModelArts平台上用户已创建的训练实验的列表。
// 该接口适用于以下场景：当用户需要查看所有或部分训练实验的概要信息，如实验名称、描述、创建时间等时，可以通过此接口查询训练实验列表。使用该接口的前提条件是用户已登录ModelArts平台并具有查看训练实验的权限。查询操作完成后，系统将返回符合条件的训练实验列表。若用户无权限操作或查询条件不合法，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListTrainingExperiments(request *model.ListTrainingExperimentsRequest) (*model.ListTrainingExperimentsResponse, error) {
	requestDef := GenReqDefForListTrainingExperiments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrainingExperimentsResponse), nil
	}
}

// ListTrainingExperimentsInvoker 查询训练实验列表
func (c *ModelArtsClient) ListTrainingExperimentsInvoker(request *model.ListTrainingExperimentsRequest) *ListTrainingExperimentsInvoker {
	requestDef := GenReqDefForListTrainingExperiments()
	return &ListTrainingExperimentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrainingJobEvents 获取训练作业事件列表
//
// 获取训练作业事件列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListTrainingJobEvents(request *model.ListTrainingJobEventsRequest) (*model.ListTrainingJobEventsResponse, error) {
	requestDef := GenReqDefForListTrainingJobEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrainingJobEventsResponse), nil
	}
}

// ListTrainingJobEventsInvoker 获取训练作业事件列表
func (c *ModelArtsClient) ListTrainingJobEventsInvoker(request *model.ListTrainingJobEventsRequest) *ListTrainingJobEventsInvoker {
	requestDef := GenReqDefForListTrainingJobEvents()
	return &ListTrainingJobEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrainingJobStages 获取训练作业流程阶段信息列表
//
// 获取训练作业流程阶段信息列表接口用于获取ModelArts平台上指定训练作业的流程阶段信息列表。
// 该接口适用于以下场景：当用户需要查看特定训练作业的流程阶段记录时，可以通过此接口获取阶段信息列表。使用该接口的前提条件是用户已知训练作业ID，并具有查看阶段信息列表的权限。查询操作完成后，平台将返回包含训练作业的阶段信息记录。若训练作业ID不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListTrainingJobStages(request *model.ListTrainingJobStagesRequest) (*model.ListTrainingJobStagesResponse, error) {
	requestDef := GenReqDefForListTrainingJobStages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrainingJobStagesResponse), nil
	}
}

// ListTrainingJobStagesInvoker 获取训练作业流程阶段信息列表
func (c *ModelArtsClient) ListTrainingJobStagesInvoker(request *model.ListTrainingJobStagesRequest) *ListTrainingJobStagesInvoker {
	requestDef := GenReqDefForListTrainingJobStages()
	return &ListTrainingJobStagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrainingJobTasks 查询训练作业的实例历史调度信息
//
// 查询训练作业调度的实例IP、节点IP等信息，可通过schedule_count参数查询具体的某一次调度的实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListTrainingJobTasks(request *model.ListTrainingJobTasksRequest) (*model.ListTrainingJobTasksResponse, error) {
	requestDef := GenReqDefForListTrainingJobTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrainingJobTasksResponse), nil
	}
}

// ListTrainingJobTasksInvoker 查询训练作业的实例历史调度信息
func (c *ModelArtsClient) ListTrainingJobTasksInvoker(request *model.ListTrainingJobTasksRequest) *ListTrainingJobTasksInvoker {
	requestDef := GenReqDefForListTrainingJobTasks()
	return &ListTrainingJobTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTrainingJobs 查询训练作业列表
//
// 根据指定查询条件查询用户创建的训练作业列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListTrainingJobs(request *model.ListTrainingJobsRequest) (*model.ListTrainingJobsResponse, error) {
	requestDef := GenReqDefForListTrainingJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTrainingJobsResponse), nil
	}
}

// ListTrainingJobsInvoker 查询训练作业列表
func (c *ModelArtsClient) ListTrainingJobsInvoker(request *model.ListTrainingJobsRequest) *ListTrainingJobsInvoker {
	requestDef := GenReqDefForListTrainingJobs()
	return &ListTrainingJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkloads 查询资源池作业列表
//
// 查询专属资源池作业列表接口用于获取指定专属资源池中的作业信息列表。该接口适用于以下场景：当需要监控专属资源池的资源使用情况、查看作业状态或管理资源分配时，用户可通过此接口获取专属资源池中作业的详细信息。使用该接口的前提条件是专属资源池已存在且用户具有相应的权限（如管理员权限或资源管理权限）。调用接口成功后，系统将返回专属资源池中作业的列表，包括作业ID、名称、状态、资源使用情况等详细信息。若专属资源池不存在、用户无权限操作或资源池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListWorkloads(request *model.ListWorkloadsRequest) (*model.ListWorkloadsResponse, error) {
	requestDef := GenReqDefForListWorkloads()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkloadsResponse), nil
	}
}

// ListWorkloadsInvoker 查询资源池作业列表
func (c *ModelArtsClient) ListWorkloadsInvoker(request *model.ListWorkloadsRequest) *ListWorkloadsInvoker {
	requestDef := GenReqDefForListWorkloads()
	return &ListWorkloadsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspace 查询工作空间列表
//
// 查询工作空间列表，响应消息体中包含详细信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListWorkspace(request *model.ListWorkspaceRequest) (*model.ListWorkspaceResponse, error) {
	requestDef := GenReqDefForListWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspaceResponse), nil
	}
}

// ListWorkspaceInvoker 查询工作空间列表
func (c *ModelArtsClient) ListWorkspaceInvoker(request *model.ListWorkspaceRequest) *ListWorkspaceInvoker {
	requestDef := GenReqDefForListWorkspace()
	return &ListWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyInferIntranetConnections 修改添加自定义URL申请
//
// 本接口用于修改添加内网自定义URL请求，适用于需要同时更新或者添加多个内网接入点的场景。调用此接口前，确保调用者具备相应的更新权限，提供需要更新的参数，如IP地址、VPC ID、子网ID等。指定的内网接入点将添加新的配置，新的配置将对相关服务生效。如果提供的内网接入点ID列表中包含无效或不存在的ID，接口将返回相应的异常信息，提示用户检查ID的有效性，如果提供的更新参数不符合格式要求（如IP地址格式不正确），接口将返回相应的异常信息，提示用户检查参数的有效性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ModifyInferIntranetConnections(request *model.ModifyInferIntranetConnectionsRequest) (*model.ModifyInferIntranetConnectionsResponse, error) {
	requestDef := GenReqDefForModifyInferIntranetConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyInferIntranetConnectionsResponse), nil
	}
}

// ModifyInferIntranetConnectionsInvoker 修改添加自定义URL申请
func (c *ModelArtsClient) ModifyInferIntranetConnectionsInvoker(request *model.ModifyInferIntranetConnectionsRequest) *ModifyInferIntranetConnectionsInvoker {
	requestDef := GenReqDefForModifyInferIntranetConnections()
	return &ModifyInferIntranetConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// NotifyTrainingJobInformation 训练作业事件上报接口
//
// 训练事件上报给业务面
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) NotifyTrainingJobInformation(request *model.NotifyTrainingJobInformationRequest) (*model.NotifyTrainingJobInformationResponse, error) {
	requestDef := GenReqDefForNotifyTrainingJobInformation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.NotifyTrainingJobInformationResponse), nil
	}
}

// NotifyTrainingJobInformationInvoker 训练作业事件上报接口
func (c *ModelArtsClient) NotifyTrainingJobInformationInvoker(request *model.NotifyTrainingJobInformationRequest) *NotifyTrainingJobInformationInvoker {
	requestDef := GenReqDefForNotifyTrainingJobInformation()
	return &NotifyTrainingJobInformationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PatchNetwork 更新网络资源
//
// 更新网络资源接口用于修改指定网络资源的配置信息。该接口适用于以下场景：当需要调整网络资源的属性、修复配置错误或优化资源设置时，用户可通过此接口更新指定网络资源的详细信息。使用该接口的前提条件是网络资源已存在且用户具有管理员权限。更新操作完成后，指定网络资源的配置信息将被成功修改，系统将反映最新的资源状态和属性。若指定的网络资源不存在、用户无权限操作或输入参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) PatchNetwork(request *model.PatchNetworkRequest) (*model.PatchNetworkResponse, error) {
	requestDef := GenReqDefForPatchNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PatchNetworkResponse), nil
	}
}

// PatchNetworkInvoker 更新网络资源
func (c *ModelArtsClient) PatchNetworkInvoker(request *model.PatchNetworkRequest) *PatchNetworkInvoker {
	requestDef := GenReqDefForPatchNetwork()
	return &PatchNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PatchNodePool 更新节点池
//
// 更新节点池接口用于修改指定节点池的配置信息。该接口适用于以下场景：当需要扩展节点池容量、调整节点规格、优化资源分配或修复节点池配置时，用户可通过此接口更新节点池的相关信息。使用该接口的前提条件是节点池已存在且用户具有管理员权限。更新操作完成后，节点池的配置将被更新，包括节点数量、规格、网络配置等参数。若节点池不存在、用户无权限操作或配置参数错误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) PatchNodePool(request *model.PatchNodePoolRequest) (*model.PatchNodePoolResponse, error) {
	requestDef := GenReqDefForPatchNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PatchNodePoolResponse), nil
	}
}

// PatchNodePoolInvoker 更新节点池
func (c *ModelArtsClient) PatchNodePoolInvoker(request *model.PatchNodePoolRequest) *PatchNodePoolInvoker {
	requestDef := GenReqDefForPatchNodePool()
	return &PatchNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PatchPool 更新资源池
//
// 更新资源池接口用于修改指定资源池的配置和容量。该接口适用于以下场景：当资源池需要扩展容量、调整配置或优化性能时，用户可通过此接口更新资源池的相关信息。使用该接口的前提条件是资源池已存在且用户具有管理员权限。更新操作完成后，资源池的配置和容量将被更新，相关资源和配置也将被调整。若资源池不存在、用户无权限操作或资源池处于不可更新状态，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) PatchPool(request *model.PatchPoolRequest) (*model.PatchPoolResponse, error) {
	requestDef := GenReqDefForPatchPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PatchPoolResponse), nil
	}
}

// PatchPoolInvoker 更新资源池
func (c *ModelArtsClient) PatchPoolInvoker(request *model.PatchPoolRequest) *PatchPoolInvoker {
	requestDef := GenReqDefForPatchPool()
	return &PatchPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RegisterImage 注册自定义镜像
//
// 注册自定义镜像接口用于将用户自定义的镜像注册到ModelArts镜像管理。该接口适用于以下场景：当用户需要将自己的自定义镜像（如特定算法环境、工具链或配置）集成到ModelArts平台时，可通过此接口将镜像注册到镜像管理中以便后续使用。使用该接口的前提条件是用户具备ModelArts镜像管理权限，并且需要提供有效的镜像地址和符合要求的镜像格式。注册操作完成后，自定义镜像将被成功添加到ModelArts镜像列表中，用户可以在后续任务中选择使用该镜像。若镜像地址无效、镜像格式不符合要求或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) RegisterImage(request *model.RegisterImageRequest) (*model.RegisterImageResponse, error) {
	requestDef := GenReqDefForRegisterImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RegisterImageResponse), nil
	}
}

// RegisterImageInvoker 注册自定义镜像
func (c *ModelArtsClient) RegisterImageInvoker(request *model.RegisterImageRequest) *RegisterImageInvoker {
	requestDef := GenReqDefForRegisterImage()
	return &RegisterImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAlgorithmByUuid 查询算法详情
//
// 根据算法id查询指定算法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAlgorithmByUuid(request *model.ShowAlgorithmByUuidRequest) (*model.ShowAlgorithmByUuidResponse, error) {
	requestDef := GenReqDefForShowAlgorithmByUuid()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlgorithmByUuidResponse), nil
	}
}

// ShowAlgorithmByUuidInvoker 查询算法详情
func (c *ModelArtsClient) ShowAlgorithmByUuidInvoker(request *model.ShowAlgorithmByUuidRequest) *ShowAlgorithmByUuidInvoker {
	requestDef := GenReqDefForShowAlgorithmByUuid()
	return &ShowAlgorithmByUuidInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAuthmodeDetail 查询授权模式
//
// 查询授权模式接口用于获取指定资源或功能的授权方式和权限配置信息。该接口适用于以下场景：当系统管理员需要查看资源的访问权限设置、开发者需要验证授权策略配置是否正确，或安全审计人员需要检查授权配置是否符合安全规范时，可通过此接口查询授权模式的详细信息。使用该接口的前提条件是用户具有查询权限且目标资源或功能的授权模式已配置。调用成功后，接口将返回授权模式的类型、规则及权限范围等详细信息。若用户无权限访问该接口，或目标资源的授权模式未配置，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAuthmodeDetail(request *model.ShowAuthmodeDetailRequest) (*model.ShowAuthmodeDetailResponse, error) {
	requestDef := GenReqDefForShowAuthmodeDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAuthmodeDetailResponse), nil
	}
}

// ShowAuthmodeDetailInvoker 查询授权模式
func (c *ModelArtsClient) ShowAuthmodeDetailInvoker(request *model.ShowAuthmodeDetailRequest) *ShowAuthmodeDetailInvoker {
	requestDef := GenReqDefForShowAuthmodeDetail()
	return &ShowAuthmodeDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchParamAnalysisResultPath 获取某个超参敏感度分析图像的路径
//
// 获取某个超参敏感度分析图像的保存路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchParamAnalysisResultPath(request *model.ShowAutoSearchParamAnalysisResultPathRequest) (*model.ShowAutoSearchParamAnalysisResultPathResponse, error) {
	requestDef := GenReqDefForShowAutoSearchParamAnalysisResultPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchParamAnalysisResultPathResponse), nil
	}
}

// ShowAutoSearchParamAnalysisResultPathInvoker 获取某个超参敏感度分析图像的路径
func (c *ModelArtsClient) ShowAutoSearchParamAnalysisResultPathInvoker(request *model.ShowAutoSearchParamAnalysisResultPathRequest) *ShowAutoSearchParamAnalysisResultPathInvoker {
	requestDef := GenReqDefForShowAutoSearchParamAnalysisResultPath()
	return &ShowAutoSearchParamAnalysisResultPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchParamsAnalysis 获取超参敏感度分析结果
//
// 获取超参敏感度分析结果的汇总表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchParamsAnalysis(request *model.ShowAutoSearchParamsAnalysisRequest) (*model.ShowAutoSearchParamsAnalysisResponse, error) {
	requestDef := GenReqDefForShowAutoSearchParamsAnalysis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchParamsAnalysisResponse), nil
	}
}

// ShowAutoSearchParamsAnalysisInvoker 获取超参敏感度分析结果
func (c *ModelArtsClient) ShowAutoSearchParamsAnalysisInvoker(request *model.ShowAutoSearchParamsAnalysisRequest) *ShowAutoSearchParamsAnalysisInvoker {
	requestDef := GenReqDefForShowAutoSearchParamsAnalysis()
	return &ShowAutoSearchParamsAnalysisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchPerTrial 查询超参搜索某个trial的结果
//
// 根据传入的trial_id，查询指定trial的搜索结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchPerTrial(request *model.ShowAutoSearchPerTrialRequest) (*model.ShowAutoSearchPerTrialResponse, error) {
	requestDef := GenReqDefForShowAutoSearchPerTrial()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchPerTrialResponse), nil
	}
}

// ShowAutoSearchPerTrialInvoker 查询超参搜索某个trial的结果
func (c *ModelArtsClient) ShowAutoSearchPerTrialInvoker(request *model.ShowAutoSearchPerTrialRequest) *ShowAutoSearchPerTrialInvoker {
	requestDef := GenReqDefForShowAutoSearchPerTrial()
	return &ShowAutoSearchPerTrialInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchTrialEarlyStop 提前终止自动化搜索作业的某个trial
//
// 提前终止自动化搜索作业的某个trial。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchTrialEarlyStop(request *model.ShowAutoSearchTrialEarlyStopRequest) (*model.ShowAutoSearchTrialEarlyStopResponse, error) {
	requestDef := GenReqDefForShowAutoSearchTrialEarlyStop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchTrialEarlyStopResponse), nil
	}
}

// ShowAutoSearchTrialEarlyStopInvoker 提前终止自动化搜索作业的某个trial
func (c *ModelArtsClient) ShowAutoSearchTrialEarlyStopInvoker(request *model.ShowAutoSearchTrialEarlyStopRequest) *ShowAutoSearchTrialEarlyStopInvoker {
	requestDef := GenReqDefForShowAutoSearchTrialEarlyStop()
	return &ShowAutoSearchTrialEarlyStopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchTrials 查询超参搜索所有trial的结果
//
// 查询超参搜索所有trial的结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchTrials(request *model.ShowAutoSearchTrialsRequest) (*model.ShowAutoSearchTrialsResponse, error) {
	requestDef := GenReqDefForShowAutoSearchTrials()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchTrialsResponse), nil
	}
}

// ShowAutoSearchTrialsInvoker 查询超参搜索所有trial的结果
func (c *ModelArtsClient) ShowAutoSearchTrialsInvoker(request *model.ShowAutoSearchTrialsRequest) *ShowAutoSearchTrialsInvoker {
	requestDef := GenReqDefForShowAutoSearchTrials()
	return &ShowAutoSearchTrialsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchYamlTemplateContent 获取自动化搜索作业yaml模板的内容
//
// 获取自动化搜索作业yaml模板的内容。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchYamlTemplateContent(request *model.ShowAutoSearchYamlTemplateContentRequest) (*model.ShowAutoSearchYamlTemplateContentResponse, error) {
	requestDef := GenReqDefForShowAutoSearchYamlTemplateContent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchYamlTemplateContentResponse), nil
	}
}

// ShowAutoSearchYamlTemplateContentInvoker 获取自动化搜索作业yaml模板的内容
func (c *ModelArtsClient) ShowAutoSearchYamlTemplateContentInvoker(request *model.ShowAutoSearchYamlTemplateContentRequest) *ShowAutoSearchYamlTemplateContentInvoker {
	requestDef := GenReqDefForShowAutoSearchYamlTemplateContent()
	return &ShowAutoSearchYamlTemplateContentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAutoSearchYamlTemplatesInfo 获取自动化搜索作业yaml模板的信息
//
// 获取自动化搜索作业yaml模板的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowAutoSearchYamlTemplatesInfo(request *model.ShowAutoSearchYamlTemplatesInfoRequest) (*model.ShowAutoSearchYamlTemplatesInfoResponse, error) {
	requestDef := GenReqDefForShowAutoSearchYamlTemplatesInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAutoSearchYamlTemplatesInfoResponse), nil
	}
}

// ShowAutoSearchYamlTemplatesInfoInvoker 获取自动化搜索作业yaml模板的信息
func (c *ModelArtsClient) ShowAutoSearchYamlTemplatesInfoInvoker(request *model.ShowAutoSearchYamlTemplatesInfoRequest) *ShowAutoSearchYamlTemplatesInfoInvoker {
	requestDef := GenReqDefForShowAutoSearchYamlTemplatesInfo()
	return &ShowAutoSearchYamlTemplatesInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDynamicStorage 获取动态挂载存储实例详情
//
// 获取动态挂载OBS实例详情接口用于获取已挂载到运行中Notebook实例中的存储实例的详细信息。
//
// 适用场景：用户需要查看Notebook实例中已挂载的存储实例的详细信息时，可通过此接口获取。使用该接口的前提条件是用户已登录系统并具有访问目标Notebook实例的权限。调用该接口后，系统将返回指定Notebook实例中挂载的存储实例的详细信息。若用户无权限访问指定实例，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowDynamicStorage(request *model.ShowDynamicStorageRequest) (*model.ShowDynamicStorageResponse, error) {
	requestDef := GenReqDefForShowDynamicStorage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDynamicStorageResponse), nil
	}
}

// ShowDynamicStorageInvoker 获取动态挂载存储实例详情
func (c *ModelArtsClient) ShowDynamicStorageInvoker(request *model.ShowDynamicStorageRequest) *ShowDynamicStorageInvoker {
	requestDef := GenReqDefForShowDynamicStorage()
	return &ShowDynamicStorageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowImage 查询镜像详情
//
// 查询镜像详情接口用于查询镜像的详细信息。该接口适用于以下场景：当用户需要了解特定镜像的详细信息（如镜像名称、版本、创建时间、大小、状态等）或对镜像执行一些操作时，可通过此接口获取镜像的详细信息。使用该接口的前提条件是用户具备镜像管理权限，并且待查询镜像有效且存在。查询操作完成后，将返回镜像的详细信息，包括镜像ID、名称、版本、创建时间、大小以及状态等。若镜像不存在或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowImage(request *model.ShowImageRequest) (*model.ShowImageResponse, error) {
	requestDef := GenReqDefForShowImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageResponse), nil
	}
}

// ShowImageInvoker 查询镜像详情
func (c *ModelArtsClient) ShowImageInvoker(request *model.ShowImageRequest) *ShowImageInvoker {
	requestDef := GenReqDefForShowImage()
	return &ShowImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferDeployment 查询服务部署详情
//
// 通过服务ID、部署ID查询对应的部署详情，调用者可以通过有效的服务ID、部署ID获取部署的名称、状态、服务实例、配置参数等详细信息。调用者需具有足够的权限，且输入的服务ID、部署ID必须有效。查询成功时返回部署详细信息，查询失败时返回特定的错误码和错误信息。若服务ID或者部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferDeployment(request *model.ShowInferDeploymentRequest) (*model.ShowInferDeploymentResponse, error) {
	requestDef := GenReqDefForShowInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferDeploymentResponse), nil
	}
}

// ShowInferDeploymentInvoker 查询服务部署详情
func (c *ModelArtsClient) ShowInferDeploymentInvoker(request *model.ShowInferDeploymentRequest) *ShowInferDeploymentInvoker {
	requestDef := GenReqDefForShowInferDeployment()
	return &ShowInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferDeploymentVersion 查询在线服务部署版本详情
//
// 此接口用于获取指定服务部署版本的详细信息，适用于需要查看特定版本的详细配置和状态的场景，例如确认版本的功能、性能参数或发布历史。请求需包含有效的服务ID、部署ID及版本号。用户必须具有对目标服务部署的查看权限。请求成功后，返回该版本的详细信息，包括版本号、发布时间、配置参数和状态。若服务ID、部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；若服务部署无该版本信息，则返回404 Not Found。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferDeploymentVersion(request *model.ShowInferDeploymentVersionRequest) (*model.ShowInferDeploymentVersionResponse, error) {
	requestDef := GenReqDefForShowInferDeploymentVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferDeploymentVersionResponse), nil
	}
}

// ShowInferDeploymentVersionInvoker 查询在线服务部署版本详情
func (c *ModelArtsClient) ShowInferDeploymentVersionInvoker(request *model.ShowInferDeploymentVersionRequest) *ShowInferDeploymentVersionInvoker {
	requestDef := GenReqDefForShowInferDeploymentVersion()
	return &ShowInferDeploymentVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferService 查询服务详情
//
// 通过服务ID查询对应的服务详情，调用者可以通过有效的服务ID获取服务的名称、状态、服务实例、配置参数等详细信息。调用者需具有足够的权限，且输入的服务ID必须有效。查询成功时返回服务详细信息，查询失败时返回特定的错误码和错误信息。若服务ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferService(request *model.ShowInferServiceRequest) (*model.ShowInferServiceResponse, error) {
	requestDef := GenReqDefForShowInferService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferServiceResponse), nil
	}
}

// ShowInferServiceInvoker 查询服务详情
func (c *ModelArtsClient) ShowInferServiceInvoker(request *model.ShowInferServiceRequest) *ShowInferServiceInvoker {
	requestDef := GenReqDefForShowInferService()
	return &ShowInferServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferServiceCluster 查询纳管资源池详情
//
// 该接口允许用户通过指定资源池的ID来查询纳管资源池的详细信息，包括实例ID、名称、Flavor规格、实例状态和实例可访问的URL。此功能适用于需要监控或管理云资源的用户，使用该接口前，用户需确保已拥有访问权限及正确的资源池ID。执行成功后，用户将获得所需的实例详情，可用于进一步的资源管理和配置。如果资源池ID无效或用户没有相应的访问权限，接口将返回错误信息，如404 Not Found或401 Unauthorized。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferServiceCluster(request *model.ShowInferServiceClusterRequest) (*model.ShowInferServiceClusterResponse, error) {
	requestDef := GenReqDefForShowInferServiceCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferServiceClusterResponse), nil
	}
}

// ShowInferServiceClusterInvoker 查询纳管资源池详情
func (c *ModelArtsClient) ShowInferServiceClusterInvoker(request *model.ShowInferServiceClusterRequest) *ShowInferServiceClusterInvoker {
	requestDef := GenReqDefForShowInferServiceCluster()
	return &ShowInferServiceClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferServiceTags 查询资源标签
//
// 该接口适用于需要获取资源（如模型、数据集、服务等）的标签信息的场景，例如在资源管理或分类中，用户可以通过查询标签来了解资源的用途、状态或其他属性。通过调用此接口，用户可以通过资源ID获取指定资源的所有标签列表。用户必须具有足够的权限，且目标资源需存在。查询成功后，返回资源的标签列表；若失败，返回具体的错误信息。常见异常包括权限验证错误、资源不存在错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferServiceTags(request *model.ShowInferServiceTagsRequest) (*model.ShowInferServiceTagsResponse, error) {
	requestDef := GenReqDefForShowInferServiceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferServiceTagsResponse), nil
	}
}

// ShowInferServiceTagsInvoker 查询资源标签
func (c *ModelArtsClient) ShowInferServiceTagsInvoker(request *model.ShowInferServiceTagsRequest) *ShowInferServiceTagsInvoker {
	requestDef := GenReqDefForShowInferServiceTags()
	return &ShowInferServiceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetwork 查询网络资源
//
// 查询网络资源接口用于获取指定网络资源的详情信息。该接口适用于以下场景：当用户需要查看特定网络资源的详细配置、状态或属性时，可通过此接口查询指定网络资源的详细信息。使用该接口的前提条件是用户具有相应的权限，并且指定的网络资源已存在于系统中。查询操作完成后，接口将返回指定网络资源的详细信息，包括资源ID、类型、状态、配置参数等。若指定的网络资源不存在、用户无权限操作或输入参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNetwork(request *model.ShowNetworkRequest) (*model.ShowNetworkResponse, error) {
	requestDef := GenReqDefForShowNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkResponse), nil
	}
}

// ShowNetworkInvoker 查询网络资源
func (c *ModelArtsClient) ShowNetworkInvoker(request *model.ShowNetworkRequest) *ShowNetworkInvoker {
	requestDef := GenReqDefForShowNetwork()
	return &ShowNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkAvailableIp 查询网络可用的IP
//
// 查询网络可用的IP接口用于查找指定网络中未被占用的IP地址。该接口适用于以下场景：在网络规划、资源分配或故障排查时，用户需要快速获取可用的IP地址信息。使用该接口的前提条件是用户具有访问目标网络的权限，并且需要提供有效的网络范围（如子网掩码或IP段）。查询完成后，接口将返回指定网络中未被占用的IP地址列表，用户可以根据结果进行IP地址的分配或进一步操作。若网络不可达、权限不足或网络范围有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNetworkAvailableIp(request *model.ShowNetworkAvailableIpRequest) (*model.ShowNetworkAvailableIpResponse, error) {
	requestDef := GenReqDefForShowNetworkAvailableIp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkAvailableIpResponse), nil
	}
}

// ShowNetworkAvailableIpInvoker 查询网络可用的IP
func (c *ModelArtsClient) ShowNetworkAvailableIpInvoker(request *model.ShowNetworkAvailableIpRequest) *ShowNetworkAvailableIpInvoker {
	requestDef := GenReqDefForShowNetworkAvailableIp()
	return &ShowNetworkAvailableIpInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodeConfigTemplate 查询节点配置模板
//
// 查询节点配置模板接口用于获取指定节点配置模板的详细信息。该接口适用于以下场景：当用户需要查看节点配置模板的内容、管理节点配置或进行相关操作时，可通过此接口获取指定节点配置模板的详细信息。使用该接口的前提条件是节点配置模板已存在且用户具有相应的访问权限。调用该接口后，系统将返回指定节点配置模板的详细信息，包括模板名称、版本、配置参数及描述等。若节点配置模板不存在或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNodeConfigTemplate(request *model.ShowNodeConfigTemplateRequest) (*model.ShowNodeConfigTemplateResponse, error) {
	requestDef := GenReqDefForShowNodeConfigTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodeConfigTemplateResponse), nil
	}
}

// ShowNodeConfigTemplateInvoker 查询节点配置模板
func (c *ModelArtsClient) ShowNodeConfigTemplateInvoker(request *model.ShowNodeConfigTemplateRequest) *ShowNodeConfigTemplateInvoker {
	requestDef := GenReqDefForShowNodeConfigTemplate()
	return &ShowNodeConfigTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodePool 查询指定节点池详情
//
// 查询指定节点池详情接口用于获取指定节点池的详细信息。该接口适用于以下场景：当需要查看节点池的配置、状态、资源使用情况或管理节点池时，用户可通过此接口获取节点池的详细信息。使用该接口的前提条件是节点池已存在且用户具有访问该节点池的权限。调用接口成功后，系统将返回节点池的详细信息，包括节点池ID、名称、节点数量、状态、创建时间、配置参数等。若节点池不存在、用户无权限访问或节点池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNodePool(request *model.ShowNodePoolRequest) (*model.ShowNodePoolResponse, error) {
	requestDef := GenReqDefForShowNodePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodePoolResponse), nil
	}
}

// ShowNodePoolInvoker 查询指定节点池详情
func (c *ModelArtsClient) ShowNodePoolInvoker(request *model.ShowNodePoolRequest) *ShowNodePoolInvoker {
	requestDef := GenReqDefForShowNodePool()
	return &ShowNodePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowObsUrlOfTrainingJobLogs 查询训练作业指定任务的日志（OBS链接）
//
// 查询训练作业指定任务的日志（OBS临时链接，有效期5分钟），可全量查看或直接下载。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowObsUrlOfTrainingJobLogs(request *model.ShowObsUrlOfTrainingJobLogsRequest) (*model.ShowObsUrlOfTrainingJobLogsResponse, error) {
	requestDef := GenReqDefForShowObsUrlOfTrainingJobLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowObsUrlOfTrainingJobLogsResponse), nil
	}
}

// ShowObsUrlOfTrainingJobLogsInvoker 查询训练作业指定任务的日志（OBS链接）
func (c *ModelArtsClient) ShowObsUrlOfTrainingJobLogsInvoker(request *model.ShowObsUrlOfTrainingJobLogsRequest) *ShowObsUrlOfTrainingJobLogsInvoker {
	requestDef := GenReqDefForShowObsUrlOfTrainingJobLogs()
	return &ShowObsUrlOfTrainingJobLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOrder 查询订单详情
//
// 查询订单详情接口用于获取指定订单的详细信息。该接口适用于以下场景：当需要查看订单的状态、金额、商品信息或处理订单相关问题时，用户可通过此接口获取订单的详细数据。使用该接口的前提条件是订单已存在且用户具有访问该订单的权限。调用接口成功后，系统将返回订单的详细信息，包括订单号、商品列表、金额、支付状态、下单时间等。若订单不存在、用户无权限访问或订单信息未正确配置，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowOrder(request *model.ShowOrderRequest) (*model.ShowOrderResponse, error) {
	requestDef := GenReqDefForShowOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOrderResponse), nil
	}
}

// ShowOrderInvoker 查询订单详情
func (c *ModelArtsClient) ShowOrderInvoker(request *model.ShowOrderRequest) *ShowOrderInvoker {
	requestDef := GenReqDefForShowOrder()
	return &ShowOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOsConfig 查询OS的配置参数
//
// 查询OS的配置参数接口用于获取ModelArts OS服务的配置参数，如网络网段、用户资源配额等。该接口适用于以下场景：当需要了解当前ModelArts OS服务的网络配置、资源分配情况或进行系统管理时，用户可通过此接口查询相关的配置参数。使用该接口的前提条件是用户具有访问ModelArts OS服务的权限，且服务处于正常运行状态。查询操作完成后，用户将获得指定的配置参数信息，便于进行后续的资源规划或系统优化。若用户无权限访问、服务不可用或请求参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowOsConfig(request *model.ShowOsConfigRequest) (*model.ShowOsConfigResponse, error) {
	requestDef := GenReqDefForShowOsConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOsConfigResponse), nil
	}
}

// ShowOsConfigInvoker 查询OS的配置参数
func (c *ModelArtsClient) ShowOsConfigInvoker(request *model.ShowOsConfigRequest) *ShowOsConfigInvoker {
	requestDef := GenReqDefForShowOsConfig()
	return &ShowOsConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowOsQuota 查询OS的配额
//
// 查询OS配额接口用于获取ModelArts OS服务中部分资源的配额信息，如资源池配额、网络配额等。该接口适用于以下场景：当需要了解资源池或网络资源的使用限制、规划资源分配或监控资源使用情况时，用户可通过此接口获取相关配额信息。使用该接口的前提条件是ModelArts OS服务已部署且用户具有相应的权限（如管理员权限或资源管理权限）。调用接口成功后，系统将返回资源池配额、网络配额等详细信息，帮助用户更好地进行资源规划和管理。若用户无权限访问该接口、服务不可用或配额信息未配置，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowOsQuota(request *model.ShowOsQuotaRequest) (*model.ShowOsQuotaResponse, error) {
	requestDef := GenReqDefForShowOsQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowOsQuotaResponse), nil
	}
}

// ShowOsQuotaInvoker 查询OS的配额
func (c *ModelArtsClient) ShowOsQuotaInvoker(request *model.ShowOsQuotaRequest) *ShowOsQuotaInvoker {
	requestDef := GenReqDefForShowOsQuota()
	return &ShowOsQuotaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPluginTemplate 查询插件模板
//
// 查询插件模板接口用于获取指定插件模板的详细信息。该接口适用于以下场景：当需要了解特定插件模板的配置、功能或使用方式时，用户可通过此接口查询插件模板的详细信息。使用该接口的前提条件是插件模板已存在且用户具有访问权限。查询操作完成后，用户将获得指定插件模板的详细信息，包括模板的配置参数、功能描述等，便于后续的插件开发或配置管理。若插件模板不存在或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPluginTemplate(request *model.ShowPluginTemplateRequest) (*model.ShowPluginTemplateResponse, error) {
	requestDef := GenReqDefForShowPluginTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPluginTemplateResponse), nil
	}
}

// ShowPluginTemplateInvoker 查询插件模板
func (c *ModelArtsClient) ShowPluginTemplateInvoker(request *model.ShowPluginTemplateRequest) *ShowPluginTemplateInvoker {
	requestDef := GenReqDefForShowPluginTemplate()
	return &ShowPluginTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPool 查询资源池
//
// 查询资源池信息接口用于获取指定资源池的详细信息。该接口适用于以下场景：当用户需要查看特定资源池的详细配置、状态、资源使用情况或进行资源管理时，可通过此接口查询指定资源池的详细信息。使用该接口的前提条件是用户具有相应的权限，并且指定的资源池已存在于系统中。查询操作完成后，接口将返回资源池的详细信息，包括资源池ID、名称、类型、状态、资源配额、利用率等。若指定的资源池不存在、用户无权限操作或输入参数有误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPool(request *model.ShowPoolRequest) (*model.ShowPoolResponse, error) {
	requestDef := GenReqDefForShowPool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolResponse), nil
	}
}

// ShowPoolInvoker 查询资源池
func (c *ModelArtsClient) ShowPoolInvoker(request *model.ShowPoolRequest) *ShowPoolInvoker {
	requestDef := GenReqDefForShowPool()
	return &ShowPoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolMonitor 资源池监控
//
// 资源池监控接口用于获取指定资源池的实时或历史监控信息。该接口适用于以下场景：当需要实时查看资源池的资源使用情况、性能状态或历史数据时，用户可通过此接口获取资源池的监控数据。使用该接口的前提条件是资源池已存在且用户具有管理员权限。调用接口成功后，系统将返回资源池的监控信息，包括资源使用率、性能指标及历史趋势等数据。若资源池不存在、用户无权限操作或资源池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolMonitor(request *model.ShowPoolMonitorRequest) (*model.ShowPoolMonitorResponse, error) {
	requestDef := GenReqDefForShowPoolMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolMonitorResponse), nil
	}
}

// ShowPoolMonitorInvoker 资源池监控
func (c *ModelArtsClient) ShowPoolMonitorInvoker(request *model.ShowPoolMonitorRequest) *ShowPoolMonitorInvoker {
	requestDef := GenReqDefForShowPoolMonitor()
	return &ShowPoolMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolNode 查询资源池单个节点详情
//
// 查询资源池中的单个节点接口用于获取指定资源池内单个节点的详细信息。该接口适用于以下场景：当用户需要了解节点资源分配、详细信息时，可通过此接口获取节点的类型、状态、配置参数及关联服务等信息。使用该接口的前提条件是目标资源池已存在且用户具备查看权限，同时需提供有效的资源池标识符作为输入参数。接口调用后，系统将返回资源池中单个节点数据；若资源池不存在、用户权限不足或输入参数无效，接口将返回对应的错误信息（如404未找到资源池或403权限拒绝）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolNode(request *model.ShowPoolNodeRequest) (*model.ShowPoolNodeResponse, error) {
	requestDef := GenReqDefForShowPoolNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolNodeResponse), nil
	}
}

// ShowPoolNodeInvoker 查询资源池单个节点详情
func (c *ModelArtsClient) ShowPoolNodeInvoker(request *model.ShowPoolNodeRequest) *ShowPoolNodeInvoker {
	requestDef := GenReqDefForShowPoolNode()
	return &ShowPoolNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolNodeConfig 查询资源池节点自定义配置
//
// 查询资源池节点自定义配置接口用于获取指定资源池节点的自定义配置信息。该接口适用于以下场景：当需要查看资源池节点的详细配置、优化资源分配或管理节点资源时，用户可通过此接口获取节点的自定义配置数据。使用该接口的前提条件是资源池节点已存在且用户具有访问该节点的权限。调用接口成功后，系统将返回资源池节点的自定义配置信息，包括硬件规格、软件环境、网络设置等详细参数。若资源池节点不存在、用户无权限访问或节点配置信息未正确配置，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolNodeConfig(request *model.ShowPoolNodeConfigRequest) (*model.ShowPoolNodeConfigResponse, error) {
	requestDef := GenReqDefForShowPoolNodeConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolNodeConfigResponse), nil
	}
}

// ShowPoolNodeConfigInvoker 查询资源池节点自定义配置
func (c *ModelArtsClient) ShowPoolNodeConfigInvoker(request *model.ShowPoolNodeConfigRequest) *ShowPoolNodeConfigInvoker {
	requestDef := GenReqDefForShowPoolNodeConfig()
	return &ShowPoolNodeConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolNodeConfigTemplate 查询资源池节点自定义配置模板
//
// 查询资源池节点自定义配置模板接口用于获取节点配置模板的结构定义与参数规范。该接口适用于以下场景：当需要了解节点自定义配置的模板结构（如参数字段、校验规则、示例值）、验证配置模板是否符合规范或进行配置模板选型时，用户可通过此接口获取模板的元数据（如参数说明、类型限制、依赖关系等）。使用该接口的前提条件是目标配置模板必须已注册至系统且处于可访问状态，调用者需具备模板查看权限，且系统配置管理服务正常运行。查询操作完成后，系统将返回模板的完整定义信息（如参数列表、版本号、更新时间等），且不会对模板内容或节点配置产生影响。若模板未注册、用户权限不足或系统服务异常，接口将返回对应的错误信息（如\&quot;404 Not Found\&quot;、\&quot;403 Forbidden\&quot;或\&quot;503 Service Unavailable\&quot;）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolNodeConfigTemplate(request *model.ShowPoolNodeConfigTemplateRequest) (*model.ShowPoolNodeConfigTemplateResponse, error) {
	requestDef := GenReqDefForShowPoolNodeConfigTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolNodeConfigTemplateResponse), nil
	}
}

// ShowPoolNodeConfigTemplateInvoker 查询资源池节点自定义配置模板
func (c *ModelArtsClient) ShowPoolNodeConfigTemplateInvoker(request *model.ShowPoolNodeConfigTemplateRequest) *ShowPoolNodeConfigTemplateInvoker {
	requestDef := GenReqDefForShowPoolNodeConfigTemplate()
	return &ShowPoolNodeConfigTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolRuntimeMetrics 查询资源实时利用率
//
// 查询资源实时利用率接口用于获取当前项目下所有资源池的实时利用率信息。该接口适用于以下场景：当用户需要监控资源使用情况、进行资源优化、容量规划或故障排查时，可通过此接口查询资源池的实时利用率，包括CPU、内存、存储等资源的使用情况。使用该接口的前提条件是用户具有访问该项目的权限，并且资源池已存在且处于运行状态。查询操作完成后，接口将返回资源池的实时利用率数据，包含利用率百分比、资源类型、时间戳等详细信息。若用户无权限、资源池不存在或系统无法获取实时数据，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolRuntimeMetrics(request *model.ShowPoolRuntimeMetricsRequest) (*model.ShowPoolRuntimeMetricsResponse, error) {
	requestDef := GenReqDefForShowPoolRuntimeMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolRuntimeMetricsResponse), nil
	}
}

// ShowPoolRuntimeMetricsInvoker 查询资源实时利用率
func (c *ModelArtsClient) ShowPoolRuntimeMetricsInvoker(request *model.ShowPoolRuntimeMetricsRequest) *ShowPoolRuntimeMetricsInvoker {
	requestDef := GenReqDefForShowPoolRuntimeMetrics()
	return &ShowPoolRuntimeMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolStatistics 资源池统计
//
// 资源池统计接口用于获取指定资源池的统计信息。该接口适用于以下场景：当需要了解资源池的资源使用情况、分配情况或利用率时，用户可通过此接口获取资源池的统计数据。使用该接口的前提条件是资源池已存在且用户具有管理员权限。调用接口成功后，系统将返回资源池的统计信息，包括资源使用总量、已分配量、利用率及资源分配趋势等数据。若资源池不存在、用户无权限操作或资源池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolStatistics(request *model.ShowPoolStatisticsRequest) (*model.ShowPoolStatisticsResponse, error) {
	requestDef := GenReqDefForShowPoolStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolStatisticsResponse), nil
	}
}

// ShowPoolStatisticsInvoker 资源池统计
func (c *ModelArtsClient) ShowPoolStatisticsInvoker(request *model.ShowPoolStatisticsRequest) *ShowPoolStatisticsInvoker {
	requestDef := GenReqDefForShowPoolStatistics()
	return &ShowPoolStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPoolTags 查询资源池上的标签
//
// 查询资源池的标签接口用于获取指定资源池的标签信息。该接口适用于以下场景：当需要查看、管理或统计特定资源池的标签信息时，用户可通过此接口获取资源池的标签数据。使用该接口的前提条件是资源池已存在且用户具有访问该资源池的权限。调用接口成功后，系统将返回指定资源池的标签信息，包括标签键和标签值。若资源池不存在、用户无权限访问或资源池未配置标签，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowPoolTags(request *model.ShowPoolTagsRequest) (*model.ShowPoolTagsResponse, error) {
	requestDef := GenReqDefForShowPoolTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPoolTagsResponse), nil
	}
}

// ShowPoolTagsInvoker 查询资源池上的标签
func (c *ModelArtsClient) ShowPoolTagsInvoker(request *model.ShowPoolTagsRequest) *ShowPoolTagsInvoker {
	requestDef := GenReqDefForShowPoolTags()
	return &ShowPoolTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSaveImageJob 查询训练作业镜像保存任务
//
// 查询训练作业镜像保存任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowSaveImageJob(request *model.ShowSaveImageJobRequest) (*model.ShowSaveImageJobResponse, error) {
	requestDef := GenReqDefForShowSaveImageJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSaveImageJobResponse), nil
	}
}

// ShowSaveImageJobInvoker 查询训练作业镜像保存任务
func (c *ModelArtsClient) ShowSaveImageJobInvoker(request *model.ShowSaveImageJobRequest) *ShowSaveImageJobInvoker {
	requestDef := GenReqDefForShowSaveImageJob()
	return &ShowSaveImageJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSearchAlgorithms 获取支持的超参搜索算法
//
// 获取支持的超参搜索算法。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowSearchAlgorithms(request *model.ShowSearchAlgorithmsRequest) (*model.ShowSearchAlgorithmsResponse, error) {
	requestDef := GenReqDefForShowSearchAlgorithms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSearchAlgorithmsResponse), nil
	}
}

// ShowSearchAlgorithmsInvoker 获取支持的超参搜索算法
func (c *ModelArtsClient) ShowSearchAlgorithmsInvoker(request *model.ShowSearchAlgorithmsRequest) *ShowSearchAlgorithmsInvoker {
	requestDef := GenReqDefForShowSearchAlgorithms()
	return &ShowSearchAlgorithmsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainJobTags 查询训练作业标签
//
// 查询训练作业标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainJobTags(request *model.ShowTrainJobTagsRequest) (*model.ShowTrainJobTagsResponse, error) {
	requestDef := GenReqDefForShowTrainJobTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainJobTagsResponse), nil
	}
}

// ShowTrainJobTagsInvoker 查询训练作业标签
func (c *ModelArtsClient) ShowTrainJobTagsInvoker(request *model.ShowTrainJobTagsRequest) *ShowTrainJobTagsInvoker {
	requestDef := GenReqDefForShowTrainJobTags()
	return &ShowTrainJobTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingExperimentDetails 查询训练实验详情
//
// 查询训练实验详情接口用于获取指定训练实验的详细信息。
// 该接口适用于以下场景：当用户需要查看训练实验的实验名称、描述、创建时间等详细信息时，可以通过此接口获取。使用该接口的前提条件是训练实验已存在且用户具有查看该实验的权限。查询操作完成后，将返回训练实验的详细信息，包括但不限于实验ID、名称、描述、创建时间等。若训练实验不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingExperimentDetails(request *model.ShowTrainingExperimentDetailsRequest) (*model.ShowTrainingExperimentDetailsResponse, error) {
	requestDef := GenReqDefForShowTrainingExperimentDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingExperimentDetailsResponse), nil
	}
}

// ShowTrainingExperimentDetailsInvoker 查询训练实验详情
func (c *ModelArtsClient) ShowTrainingExperimentDetailsInvoker(request *model.ShowTrainingExperimentDetailsRequest) *ShowTrainingExperimentDetailsInvoker {
	requestDef := GenReqDefForShowTrainingExperimentDetails()
	return &ShowTrainingExperimentDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingJobDetails 查询训练作业详情
//
// 查询训练作业详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingJobDetails(request *model.ShowTrainingJobDetailsRequest) (*model.ShowTrainingJobDetailsResponse, error) {
	requestDef := GenReqDefForShowTrainingJobDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingJobDetailsResponse), nil
	}
}

// ShowTrainingJobDetailsInvoker 查询训练作业详情
func (c *ModelArtsClient) ShowTrainingJobDetailsInvoker(request *model.ShowTrainingJobDetailsRequest) *ShowTrainingJobDetailsInvoker {
	requestDef := GenReqDefForShowTrainingJobDetails()
	return &ShowTrainingJobDetailsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingJobEngines 获取训练作业支持的AI预置框架
//
// 获取训练作业支持的AI预置框架。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingJobEngines(request *model.ShowTrainingJobEnginesRequest) (*model.ShowTrainingJobEnginesResponse, error) {
	requestDef := GenReqDefForShowTrainingJobEngines()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingJobEnginesResponse), nil
	}
}

// ShowTrainingJobEnginesInvoker 获取训练作业支持的AI预置框架
func (c *ModelArtsClient) ShowTrainingJobEnginesInvoker(request *model.ShowTrainingJobEnginesRequest) *ShowTrainingJobEnginesInvoker {
	requestDef := GenReqDefForShowTrainingJobEngines()
	return &ShowTrainingJobEnginesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingJobFlavors 获取训练作业支持的公共规格
//
// 获取训练作业支持的公共规格。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingJobFlavors(request *model.ShowTrainingJobFlavorsRequest) (*model.ShowTrainingJobFlavorsResponse, error) {
	requestDef := GenReqDefForShowTrainingJobFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingJobFlavorsResponse), nil
	}
}

// ShowTrainingJobFlavorsInvoker 获取训练作业支持的公共规格
func (c *ModelArtsClient) ShowTrainingJobFlavorsInvoker(request *model.ShowTrainingJobFlavorsRequest) *ShowTrainingJobFlavorsInvoker {
	requestDef := GenReqDefForShowTrainingJobFlavors()
	return &ShowTrainingJobFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingJobLogsPreview 查询训练作业指定任务的日志（预览）
//
// 查询训练作业指定任务的日志（预览）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingJobLogsPreview(request *model.ShowTrainingJobLogsPreviewRequest) (*model.ShowTrainingJobLogsPreviewResponse, error) {
	requestDef := GenReqDefForShowTrainingJobLogsPreview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingJobLogsPreviewResponse), nil
	}
}

// ShowTrainingJobLogsPreviewInvoker 查询训练作业指定任务的日志（预览）
func (c *ModelArtsClient) ShowTrainingJobLogsPreviewInvoker(request *model.ShowTrainingJobLogsPreviewRequest) *ShowTrainingJobLogsPreviewInvoker {
	requestDef := GenReqDefForShowTrainingJobLogsPreview()
	return &ShowTrainingJobLogsPreviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingJobMetrics 查询训练作业指定任务的运行指标
//
// 查询训练作业指定任务的运行指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingJobMetrics(request *model.ShowTrainingJobMetricsRequest) (*model.ShowTrainingJobMetricsResponse, error) {
	requestDef := GenReqDefForShowTrainingJobMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingJobMetricsResponse), nil
	}
}

// ShowTrainingJobMetricsInvoker 查询训练作业指定任务的运行指标
func (c *ModelArtsClient) ShowTrainingJobMetricsInvoker(request *model.ShowTrainingJobMetricsRequest) *ShowTrainingJobMetricsInvoker {
	requestDef := GenReqDefForShowTrainingJobMetrics()
	return &ShowTrainingJobMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTrainingQuotas 获取训练配额
//
// 获取训练配额接口用于查询用户在ModelArts服务中的训练资源配额信息。
// 该接口适用于以下场景：当用户需要了解当前可用的训练资源配额，以便合理规划和管理训练任务时，可以通过此接口获取配额详情。使用该接口的前提条件是用户已登录并具有查看配额的权限。响应消息体中包含用户已创建的训练作业个数、剩余可创建个数等。若用户无权限或配额信息为空，接口将返回相应的错误信息或空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowTrainingQuotas(request *model.ShowTrainingQuotasRequest) (*model.ShowTrainingQuotasResponse, error) {
	requestDef := GenReqDefForShowTrainingQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTrainingQuotasResponse), nil
	}
}

// ShowTrainingQuotasInvoker 获取训练配额
func (c *ModelArtsClient) ShowTrainingQuotasInvoker(request *model.ShowTrainingQuotasRequest) *ShowTrainingQuotasInvoker {
	requestDef := GenReqDefForShowTrainingQuotas()
	return &ShowTrainingQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkloadStatistics 查询专属资源池作业统计信息
//
// 查询专属资源池作业统计信息接口用于获取指定专属资源池中作业的统计信息。该接口适用于以下场景：当需要了解专属资源池中作业的整体运行情况、资源使用效率或作业状态分布时，用户可通过此接口获取统计信息。使用该接口的前提条件是专属资源池已存在且用户具有相应的权限（如管理员权限或资源管理权限）。调用接口成功后，系统将返回专属资源池中作业的统计信息，包括作业总数、运行中作业数、完成作业数、资源使用率等数据。若专属资源池不存在、用户无权限操作或资源池当前不可用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkloadStatistics(request *model.ShowWorkloadStatisticsRequest) (*model.ShowWorkloadStatisticsResponse, error) {
	requestDef := GenReqDefForShowWorkloadStatistics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkloadStatisticsResponse), nil
	}
}

// ShowWorkloadStatisticsInvoker 查询专属资源池作业统计信息
func (c *ModelArtsClient) ShowWorkloadStatisticsInvoker(request *model.ShowWorkloadStatisticsRequest) *ShowWorkloadStatisticsInvoker {
	requestDef := GenReqDefForShowWorkloadStatistics()
	return &ShowWorkloadStatisticsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkspace 查询工作空间详情
//
// 查询工作空间详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkspace(request *model.ShowWorkspaceRequest) (*model.ShowWorkspaceResponse, error) {
	requestDef := GenReqDefForShowWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceResponse), nil
	}
}

// ShowWorkspaceInvoker 查询工作空间详情
func (c *ModelArtsClient) ShowWorkspaceInvoker(request *model.ShowWorkspaceRequest) *ShowWorkspaceInvoker {
	requestDef := GenReqDefForShowWorkspace()
	return &ShowWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkspaceQuotas 查询工作空间配额
//
// 查询工作空间配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkspaceQuotas(request *model.ShowWorkspaceQuotasRequest) (*model.ShowWorkspaceQuotasResponse, error) {
	requestDef := GenReqDefForShowWorkspaceQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceQuotasResponse), nil
	}
}

// ShowWorkspaceQuotasInvoker 查询工作空间配额
func (c *ModelArtsClient) ShowWorkspaceQuotasInvoker(request *model.ShowWorkspaceQuotasRequest) *ShowWorkspaceQuotasInvoker {
	requestDef := GenReqDefForShowWorkspaceQuotas()
	return &ShowWorkspaceQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInferDeployment 启动服务部署
//
// 使部署从“停止”或“失败”状态进入“部署中”状态，适用于用户需要重新启动已停止或启动失败的部署的情况。调用此接口前，部署状态必须为“停止”或“失败”，且用户需具有启动部署的权限。调用成功后，部署状态将变为“部署中”，系统将开始执行部署流程，包括资源准备、配置加载等。如果部署当前状态不是“停止”或“失败”，或用户没有启动部署的权限，调用将返回错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StartInferDeployment(request *model.StartInferDeploymentRequest) (*model.StartInferDeploymentResponse, error) {
	requestDef := GenReqDefForStartInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInferDeploymentResponse), nil
	}
}

// StartInferDeploymentInvoker 启动服务部署
func (c *ModelArtsClient) StartInferDeploymentInvoker(request *model.StartInferDeploymentRequest) *StartInferDeploymentInvoker {
	requestDef := GenReqDefForStartInferDeployment()
	return &StartInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartInferService 启动服务
//
// 使服务从\&quot;停止\&quot;或\&quot;失败\&quot;状态进入\&quot;部署中\&quot;状态，适用于用户需要重新启动已停止或启动失败的服务的情况。调用此接口前，服务状态必须为\&quot;停止\&quot;或\&quot;失败\&quot;，且用户需具有启动服务的权限。调用成功后，服务状态将变为\&quot;部署中\&quot;，系统将开始执行部署流程，包括资源准备、配置加载等。如果服务当前状态不是\&quot;停止\&quot;或\&quot;失败\&quot;，或用户没有启动服务的权限，调用将返回错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StartInferService(request *model.StartInferServiceRequest) (*model.StartInferServiceResponse, error) {
	requestDef := GenReqDefForStartInferService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartInferServiceResponse), nil
	}
}

// StartInferServiceInvoker 启动服务
func (c *ModelArtsClient) StartInferServiceInvoker(request *model.StartInferServiceRequest) *StartInferServiceInvoker {
	requestDef := GenReqDefForStartInferService()
	return &StartInferServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInferDeployment 停止在线服务部署
//
// 停止在线部署功能允许用户在特定状态下主动终止正在运行或处于其他可操作状态的部署实例。该功能适用于需要维护、升级或检测到异常的服务场景，支持在服务处于\&quot;运行中\&quot;、\&quot;部署中\&quot;、\&quot;失败\&quot;或\&quot;告警\&quot;状态时执行停止操作。使用此功能前，请确保部署实例处于可停止状态，并具备相应的API调用权限。成功执行后，部署将进入停止状态，释放相关资源并停止处理新的请求。若部署不在允许停止的状态、调用权限不足或系统内部出现错误，将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopInferDeployment(request *model.StopInferDeploymentRequest) (*model.StopInferDeploymentResponse, error) {
	requestDef := GenReqDefForStopInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInferDeploymentResponse), nil
	}
}

// StopInferDeploymentInvoker 停止在线服务部署
func (c *ModelArtsClient) StopInferDeploymentInvoker(request *model.StopInferDeploymentRequest) *StopInferDeploymentInvoker {
	requestDef := GenReqDefForStopInferDeployment()
	return &StopInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopInferService 停止服务
//
// 使服务从\&quot;运行中\&quot;状态进入\&quot;停止中\&quot;最终变为\&quot;停止\&quot;状态，适用于用户需要停止正在运行的服务以节省资源成本的场景。用户需具有停止服务的权限。调用成功后，服务状态将变为\&quot;停止中\&quot;，系统将开始执行停止流程，包括释放资源、保存状态等。如果用户没有停止服务的权限，调用将返回错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopInferService(request *model.StopInferServiceRequest) (*model.StopInferServiceResponse, error) {
	requestDef := GenReqDefForStopInferService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopInferServiceResponse), nil
	}
}

// StopInferServiceInvoker 停止服务
func (c *ModelArtsClient) StopInferServiceInvoker(request *model.StopInferServiceRequest) *StopInferServiceInvoker {
	requestDef := GenReqDefForStopInferService()
	return &StopInferServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopTrainingJob 终止训练作业
//
// 终止训练作业，只可终止创建中、等待中、运行中的作业。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopTrainingJob(request *model.StopTrainingJobRequest) (*model.StopTrainingJobResponse, error) {
	requestDef := GenReqDefForStopTrainingJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopTrainingJobResponse), nil
	}
}

// StopTrainingJobInvoker 终止训练作业
func (c *ModelArtsClient) StopTrainingJobInvoker(request *model.StopTrainingJobRequest) *StopTrainingJobInvoker {
	requestDef := GenReqDefForStopTrainingJob()
	return &StopTrainingJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchInferDeploymentVersion 切换部署到指定版本
//
// 此接口用于将部署切换到指定版本，适用于需要在不同版本间进行切换以测试或回滚的场景。请求需包含有效的服务ID、部署ID及目标版本号，版本号必须是已发布的有效版本。用户必须具有对目标服务部署的管理权限，并且部署处于运行状态。切换成功后，部署将立即使用新的版本。若服务ID无效、部署ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；若部署状态不允许切换，则返回400 Bad Request。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) SwitchInferDeploymentVersion(request *model.SwitchInferDeploymentVersionRequest) (*model.SwitchInferDeploymentVersionResponse, error) {
	requestDef := GenReqDefForSwitchInferDeploymentVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchInferDeploymentVersionResponse), nil
	}
}

// SwitchInferDeploymentVersionInvoker 切换部署到指定版本
func (c *ModelArtsClient) SwitchInferDeploymentVersionInvoker(request *model.SwitchInferDeploymentVersionRequest) *SwitchInferDeploymentVersionInvoker {
	requestDef := GenReqDefForSwitchInferDeploymentVersion()
	return &SwitchInferDeploymentVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncImage 同步镜像状态
//
// 同步镜像状态接口用于修正镜像状态的异常情况。该接口适用于以下场景：当镜像状态因误操作、网络问题或系统故障等原因出现异常时，用户可通过此接口同步镜像的最新状态。使用该接口的前提条件是镜像已存在且用户具有相应的操作权限。同步操作完成后，镜像的状态将被更新为最新的正确状态，相关资源和配置也将被同步。若镜像不存在、用户无权限操作或同步过程中出现错误，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) SyncImage(request *model.SyncImageRequest) (*model.SyncImageResponse, error) {
	requestDef := GenReqDefForSyncImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncImageResponse), nil
	}
}

// SyncImageInvoker 同步镜像状态
func (c *ModelArtsClient) SyncImageInvoker(request *model.SyncImageRequest) *SyncImageInvoker {
	requestDef := GenReqDefForSyncImage()
	return &SyncImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UnbindInferApiKey 解绑应用密钥
//
// 本接口用于将已绑定的apikey从指定服务中解绑，适用于需要撤销某个apikey对特定服务的访问权限的场景。调用此接口前，确保已获取到需要解绑的apikey，并确认该apikey当前绑定在指定服务上。解绑成功后，该apikey将不再对指定服务生效，但仍可继续用于其他服务。如果尝试解绑不存在或未绑定到指定服务的apikey，将返回相应的异常信息，提示用户检查apikey的有效性和绑定状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UnbindInferApiKey(request *model.UnbindInferApiKeyRequest) (*model.UnbindInferApiKeyResponse, error) {
	requestDef := GenReqDefForUnbindInferApiKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UnbindInferApiKeyResponse), nil
	}
}

// UnbindInferApiKeyInvoker 解绑应用密钥
func (c *ModelArtsClient) UnbindInferApiKeyInvoker(request *model.UnbindInferApiKeyRequest) *UnbindInferApiKeyInvoker {
	requestDef := GenReqDefForUnbindInferApiKey()
	return &UnbindInferApiKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuthMode 更新授权模式
//
// 更新授权模式接口用于修改指定资源或功能的授权方式和权限配置信息。该接口适用于以下场景：当系统管理员需要调整资源的访问权限、开发者需要更新授权策略以适应新的业务需求，或安全审计人员需要修改授权配置以符合新的安全规范时，可通过此接口更新授权模式的详细信息。使用该接口的前提条件是用户具有更新权限且目标资源或功能的授权模式已存在。调用成功后，接口将更新目标资源的授权模式，并返回更新后的授权模式信息。若用户无权限访问该接口，或目标资源的授权模式不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateAuthMode(request *model.UpdateAuthModeRequest) (*model.UpdateAuthModeResponse, error) {
	requestDef := GenReqDefForUpdateAuthMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthModeResponse), nil
	}
}

// UpdateAuthModeInvoker 更新授权模式
func (c *ModelArtsClient) UpdateAuthModeInvoker(request *model.UpdateAuthModeRequest) *UpdateAuthModeInvoker {
	requestDef := GenReqDefForUpdateAuthMode()
	return &UpdateAuthModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateImageGroup 更新镜像组
//
// 更新镜像组接口用于更新镜像组的标签及说明信息。该接口适用于以下场景：当镜像说明需要修改，或者镜像的标签需要修改时，用户可通过此接口修改。使用该接口的前提条件是镜像组已存在且用户具有更新权限。更新操作完成后，镜像组对应的配置文件会。若镜像组不存在、用户无权限操作或镜像正在被使用，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateImageGroup(request *model.UpdateImageGroupRequest) (*model.UpdateImageGroupResponse, error) {
	requestDef := GenReqDefForUpdateImageGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateImageGroupResponse), nil
	}
}

// UpdateImageGroupInvoker 更新镜像组
func (c *ModelArtsClient) UpdateImageGroupInvoker(request *model.UpdateImageGroupRequest) *UpdateImageGroupInvoker {
	requestDef := GenReqDefForUpdateImageGroup()
	return &UpdateImageGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferDeployment 更新服务部署配置
//
// 该接口适用于需要动态调整模型服务部署配置的场景
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferDeployment(request *model.UpdateInferDeploymentRequest) (*model.UpdateInferDeploymentResponse, error) {
	requestDef := GenReqDefForUpdateInferDeployment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferDeploymentResponse), nil
	}
}

// UpdateInferDeploymentInvoker 更新服务部署配置
func (c *ModelArtsClient) UpdateInferDeploymentInvoker(request *model.UpdateInferDeploymentRequest) *UpdateInferDeploymentInvoker {
	requestDef := GenReqDefForUpdateInferDeployment()
	return &UpdateInferDeploymentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferDeploymentScale 手动服务扩缩容
//
// 该接口适用于模型服务实例扩缩容。通过调用此接口，用户可以在原有服务的情况下，对服务进行扩缩容，且不会增加新的版本；包括权限验证错误、服务状态错误和参数验证错误。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferDeploymentScale(request *model.UpdateInferDeploymentScaleRequest) (*model.UpdateInferDeploymentScaleResponse, error) {
	requestDef := GenReqDefForUpdateInferDeploymentScale()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferDeploymentScaleResponse), nil
	}
}

// UpdateInferDeploymentScaleInvoker 手动服务扩缩容
func (c *ModelArtsClient) UpdateInferDeploymentScaleInvoker(request *model.UpdateInferDeploymentScaleRequest) *UpdateInferDeploymentScaleInvoker {
	requestDef := GenReqDefForUpdateInferDeploymentScale()
	return &UpdateInferDeploymentScaleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferIntranetConnection 变更内网申请
//
// 本接口用于对当前租户的内网接入申请进行状态变更操作，支持通过（APPROVE）、拒绝（REJECT）、取消（CANCEL）和重试（RETRY）等操作。适用于需要管理内网接入申请审批流程的场景。调用此接口前，确保已具备相应的变更权限，并提供有效的内网申请ID和所需的操作类型。变更成功后，内网申请的状态将更新为指定的操作结果，并记录相关日志。如果提供的内网申请ID无效、操作类型不支持或权限不足，将返回相应的异常信息，提示用户检查输入数据的有效性和权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferIntranetConnection(request *model.UpdateInferIntranetConnectionRequest) (*model.UpdateInferIntranetConnectionResponse, error) {
	requestDef := GenReqDefForUpdateInferIntranetConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferIntranetConnectionResponse), nil
	}
}

// UpdateInferIntranetConnectionInvoker 变更内网申请
func (c *ModelArtsClient) UpdateInferIntranetConnectionInvoker(request *model.UpdateInferIntranetConnectionRequest) *UpdateInferIntranetConnectionInvoker {
	requestDef := GenReqDefForUpdateInferIntranetConnection()
	return &UpdateInferIntranetConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferService 更新服务配置
//
// 该接口适用于需要动态调整模型服务配置的场景，对模型的性能参数、资源池配置、服务调用配置等进行更新升级。通过调用此接口，用户可以在原有服务的情况下，升级成一个新的服务版本。调用此接口前，服务状态必须为“停止”、“失败”或“运行中”，且用户需具有修改服务的权限。更新成功后，新配置立即生效；若失败，服务保持原有配置并返回错误信息。常见异常包括参数验证错误、权限验证错误和服务状态错误。若服务ID无效、版本号不存在或用户无权限，则返回400 Bad Request或403 Forbidden；若服务状态不允许切换，则返回400 Bad Request。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferService(request *model.UpdateInferServiceRequest) (*model.UpdateInferServiceResponse, error) {
	requestDef := GenReqDefForUpdateInferService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferServiceResponse), nil
	}
}

// UpdateInferServiceInvoker 更新服务配置
func (c *ModelArtsClient) UpdateInferServiceInvoker(request *model.UpdateInferServiceRequest) *UpdateInferServiceInvoker {
	requestDef := GenReqDefForUpdateInferService()
	return &UpdateInferServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 修改工作空间
//
// 修改工作空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 修改工作空间
func (c *ModelArtsClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspaceQuotas 修改工作空间配额
//
// 修改工作空间配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkspaceQuotas(request *model.UpdateWorkspaceQuotasRequest) (*model.UpdateWorkspaceQuotasResponse, error) {
	requestDef := GenReqDefForUpdateWorkspaceQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceQuotasResponse), nil
	}
}

// UpdateWorkspaceQuotasInvoker 修改工作空间配额
func (c *ModelArtsClient) UpdateWorkspaceQuotasInvoker(request *model.UpdateWorkspaceQuotasRequest) *UpdateWorkspaceQuotasInvoker {
	requestDef := GenReqDefForUpdateWorkspaceQuotas()
	return &UpdateWorkspaceQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ValidateAuthorization 鉴权能否使用当前工作空间资源
//
// 鉴权能否使用当前工作空间资源接口用于验证用户是否有权限访问和使用当前工作空间中的资源。该接口适用于以下场景：当用户尝试访问或操作工作空间中的资源时，系统需要确认用户是否具有相应的权限。使用该接口的前提条件是用户已登录且工作空间已存在。鉴权成功后，用户可以正常访问和使用工作空间资源；若鉴权失败，接口将返回相应的错误信息，如用户无权限或工作空间不存在等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ValidateAuthorization(request *model.ValidateAuthorizationRequest) (*model.ValidateAuthorizationResponse, error) {
	requestDef := GenReqDefForValidateAuthorization()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ValidateAuthorizationResponse), nil
	}
}

// ValidateAuthorizationInvoker 鉴权能否使用当前工作空间资源
func (c *ModelArtsClient) ValidateAuthorizationInvoker(request *model.ValidateAuthorizationRequest) *ValidateAuthorizationInvoker {
	requestDef := GenReqDefForValidateAuthorization()
	return &ValidateAuthorizationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferDeploymentHpa 创建自动扩缩容策略
//
// 本接口用于在已部署的服务上创建定时扩缩容策略，适用于需要根据业务负载或特定时间自动调整服务实例个数的场景。调用此接口前，确保服务已成功部署并获取了有效的服务ID，并提供详细的扩缩容策略参数，如扩缩容时间、实例个数范围、条件触发器等。创建成功后，系统将根据设定的策略自动调整服务实例个数，确保服务在指定时间内的性能和可用性。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferDeploymentHpa(request *model.CreateInferDeploymentHpaRequest) (*model.CreateInferDeploymentHpaResponse, error) {
	requestDef := GenReqDefForCreateInferDeploymentHpa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferDeploymentHpaResponse), nil
	}
}

// CreateInferDeploymentHpaInvoker 创建自动扩缩容策略
func (c *ModelArtsClient) CreateInferDeploymentHpaInvoker(request *model.CreateInferDeploymentHpaRequest) *CreateInferDeploymentHpaInvoker {
	requestDef := GenReqDefForCreateInferDeploymentHpa()
	return &CreateInferDeploymentHpaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInferDeploymentHpa 删除自动扩缩容策略
//
// 本接口用于在已部署的服务上删除定时扩缩容策略，适用于需要根据业务负载或特定时间自动删除服务的场景。调用此接口前，确保服务已成功部署并获取了有效的服务ID，部署ID。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。暂时为非开放接口，后端清理服务下的自动扩缩容策略规则使用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteInferDeploymentHpa(request *model.DeleteInferDeploymentHpaRequest) (*model.DeleteInferDeploymentHpaResponse, error) {
	requestDef := GenReqDefForDeleteInferDeploymentHpa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInferDeploymentHpaResponse), nil
	}
}

// DeleteInferDeploymentHpaInvoker 删除自动扩缩容策略
func (c *ModelArtsClient) DeleteInferDeploymentHpaInvoker(request *model.DeleteInferDeploymentHpaRequest) *DeleteInferDeploymentHpaInvoker {
	requestDef := GenReqDefForDeleteInferDeploymentHpa()
	return &DeleteInferDeploymentHpaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInferDeploymentHpaEvents 查看自动扩缩容策略事件
//
// 本接口用于在已部署的服务上查看自动扩缩容策略事件，适用于查看自动扩缩容策略变动历史记录。调用此接口前，确保获取了有效的用户项目ID，服务ID，部署ID。调用成功后，会返回策略事件ID，事件状态，规则执行信息，扩缩容前实例数，扩缩容后实例数，预设目标实例数，执行记录时间。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListInferDeploymentHpaEvents(request *model.ListInferDeploymentHpaEventsRequest) (*model.ListInferDeploymentHpaEventsResponse, error) {
	requestDef := GenReqDefForListInferDeploymentHpaEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInferDeploymentHpaEventsResponse), nil
	}
}

// ListInferDeploymentHpaEventsInvoker 查看自动扩缩容策略事件
func (c *ModelArtsClient) ListInferDeploymentHpaEventsInvoker(request *model.ListInferDeploymentHpaEventsRequest) *ListInferDeploymentHpaEventsInvoker {
	requestDef := GenReqDefForListInferDeploymentHpaEvents()
	return &ListInferDeploymentHpaEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferDeploymentHpa 查看自动扩缩容策略
//
// 本接口用于在已部署的服务上查看自动扩缩容策略。调用此接口前，确保服务已成功部署并获取了有效的服务ID。查询成功后，返回服务对应的策略信息，如规则ID，规则名称，扩缩容类型，扩缩容状态，扩缩容cron表达式，目标实例数等。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferDeploymentHpa(request *model.ShowInferDeploymentHpaRequest) (*model.ShowInferDeploymentHpaResponse, error) {
	requestDef := GenReqDefForShowInferDeploymentHpa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferDeploymentHpaResponse), nil
	}
}

// ShowInferDeploymentHpaInvoker 查看自动扩缩容策略
func (c *ModelArtsClient) ShowInferDeploymentHpaInvoker(request *model.ShowInferDeploymentHpaRequest) *ShowInferDeploymentHpaInvoker {
	requestDef := GenReqDefForShowInferDeploymentHpa()
	return &ShowInferDeploymentHpaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferDeploymentHpa 修改自动扩缩容策略
//
// 本接口用于在已部署的服务上修改定时扩缩容策略，适用于需要根据业务负载或特定时间自动调整服务实例个数的场景。调用此接口前，确保服务已成功部署并获取了有效的服务ID，部署ID，并提供详细的扩缩容策略参数，如扩缩容时间、实例个数范围、条件触发器等。修改成功后，系统将根据设定的策略自动调整服务实例个数，确保服务在指定时间内的性能和可用性。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferDeploymentHpa(request *model.UpdateInferDeploymentHpaRequest) (*model.UpdateInferDeploymentHpaResponse, error) {
	requestDef := GenReqDefForUpdateInferDeploymentHpa()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferDeploymentHpaResponse), nil
	}
}

// UpdateInferDeploymentHpaInvoker 修改自动扩缩容策略
func (c *ModelArtsClient) UpdateInferDeploymentHpaInvoker(request *model.UpdateInferDeploymentHpaRequest) *UpdateInferDeploymentHpaInvoker {
	requestDef := GenReqDefForUpdateInferDeploymentHpa()
	return &UpdateInferDeploymentHpaInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInferHra 创建HRA策略
//
// 本接口用于在已部署且支持HRA策略的服务上创建HRA策略，适用于需要根据业务负载或特定时间自动调整服务实例个数的场景。调用此接口前，确保服务已成功部署并获取了有效的服务ID，并提供详细的hra策略参数，如hra时间、实例个数范围、条件触发器等。创建成功后，系统将根据设定的策略自动调整服务实例个数，确保服务在指定时间内的性能和可用性。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateInferHra(request *model.CreateInferHraRequest) (*model.CreateInferHraResponse, error) {
	requestDef := GenReqDefForCreateInferHra()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInferHraResponse), nil
	}
}

// CreateInferHraInvoker 创建HRA策略
func (c *ModelArtsClient) CreateInferHraInvoker(request *model.CreateInferHraRequest) *CreateInferHraInvoker {
	requestDef := GenReqDefForCreateInferHra()
	return &CreateInferHraInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInferHra 获取推理单元配比检测信息
//
// 本接口用于在已部署的服务上查看推理单元配比检测信息。调用此接口前，确保服务已成功部署并获取了有效的服务ID。查询成功后，返回服务对应的策略信息，如规则ID，规则名称，策略状态，HRA结果状态等。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowInferHra(request *model.ShowInferHraRequest) (*model.ShowInferHraResponse, error) {
	requestDef := GenReqDefForShowInferHra()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInferHraResponse), nil
	}
}

// ShowInferHraInvoker 获取推理单元配比检测信息
func (c *ModelArtsClient) ShowInferHraInvoker(request *model.ShowInferHraRequest) *ShowInferHraInvoker {
	requestDef := GenReqDefForShowInferHra()
	return &ShowInferHraInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInferHra 修改指定部署的HRA策略配置
//
// 本接口用于在已创建HRA策略的服务上修改指定部署的HRA策略配置，适用于需要根据业务负载或特定时间自动调整服务实例个数的场景。调用此接口前，确保服务已成功部署并获取了有效的服务ID，部署ID，并提供详细的hra策略参数，如HRA规则列表、HRA结果状态、策略状态等。修改成功后，系统将根据设定的策略自动调整服务实例个数，确保服务在指定时间内的性能和可用性。如果提供的服务ID无效、参数配置错误或系统资源不足，将返回相应的异常信息，提示用户检查输入数据的有效性或联系技术支持。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateInferHra(request *model.UpdateInferHraRequest) (*model.UpdateInferHraResponse, error) {
	requestDef := GenReqDefForUpdateInferHra()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInferHraResponse), nil
	}
}

// UpdateInferHraInvoker 修改指定部署的HRA策略配置
func (c *ModelArtsClient) UpdateInferHraInvoker(request *model.UpdateInferHraRequest) *UpdateInferHraInvoker {
	requestDef := GenReqDefForUpdateInferHra()
	return &UpdateInferHraInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AttachDevServerVolume Lite Server服务器挂载磁盘
//
// Lite Server服务器挂载磁盘接口用于将额外的磁盘挂载到Lite Server服务器上。该接口适用于以下场景：当用户需要扩展Lite Server服务器的存储空间以满足更大的数据存储需求时，可以通过此接口将指定的磁盘挂载到服务器上。使用该接口的前提条件是Lite Server服务器已创建且处于运行状态、或者停止状态，用户具有挂载磁盘的权限，且指定的磁盘已存在且未被其他服务器使用。挂载操作完成后，磁盘将成功挂载到Lite Server服务器上，用户可以访问和使用新增的存储空间。若Lite Server服务器不存在、指定的磁盘不存在或已被使用，或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) AttachDevServerVolume(request *model.AttachDevServerVolumeRequest) (*model.AttachDevServerVolumeResponse, error) {
	requestDef := GenReqDefForAttachDevServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AttachDevServerVolumeResponse), nil
	}
}

// AttachDevServerVolumeInvoker Lite Server服务器挂载磁盘
func (c *ModelArtsClient) AttachDevServerVolumeInvoker(request *model.AttachDevServerVolumeRequest) *AttachDevServerVolumeInvoker {
	requestDef := GenReqDefForAttachDevServerVolume()
	return &AttachDevServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDevServersAction 批量操作Lite Server实例
//
// 批量操作Lite Server实例接口用于对多个Lite Server实例进行统一操作，如启动、停止、重启或删除等。该接口适用于以下场景：当需要对多个Lite Server实例进行相同的操作，例如在维护期间批量停止实例、更新配置后批量重启实例或清理不再需要的实例时，用户可通过此接口高效地完成批量操作。使用该接口的前提条件是目标Lite Server实例已存在且用户具有相应的操作权限。操作完成后，所有指定的Lite Server实例将根据请求完成相应的状态变更或被移除，相关资源和配置也将被相应调整或清理。若目标Lite Server实例不存在、用户无权限操作或请求参数不正确，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BatchDevServersAction(request *model.BatchDevServersActionRequest) (*model.BatchDevServersActionResponse, error) {
	requestDef := GenReqDefForBatchDevServersAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDevServersActionResponse), nil
	}
}

// BatchDevServersActionInvoker 批量操作Lite Server实例
func (c *ModelArtsClient) BatchDevServersActionInvoker(request *model.BatchDevServersActionRequest) *BatchDevServersActionInvoker {
	requestDef := GenReqDefForBatchDevServersAction()
	return &BatchDevServersActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BindDevServerPublicIP Lite Server服务器绑定EIP
//
// Lite Server服务器绑定的EIP接口用于将弹性公网IP（EIP）绑定到Lite Server服务器上。该接口适用于以下场景：当用户需要为Lite Server服务器分配一个固定的公网IP地址，以便从外部网络访问服务器时，可以通过此接口将指定的EIP绑定到服务器上。使用该接口的前提条件是Lite Server服务器已创建且处于运行状态，用户具有绑定EIP的权限，且指定的EIP已存在且未被其他资源使用。绑定操作完成后，EIP将成功绑定到Lite Server服务器上，服务器可以通过该EIP从外部网络访问。若Lite Server服务器不存在、已处于停止状态、指定的EIP不存在或已被使用，或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) BindDevServerPublicIP(request *model.BindDevServerPublicIpRequest) (*model.BindDevServerPublicIpResponse, error) {
	requestDef := GenReqDefForBindDevServerPublicIP()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BindDevServerPublicIpResponse), nil
	}
}

// BindDevServerPublicIPInvoker Lite Server服务器绑定EIP
func (c *ModelArtsClient) BindDevServerPublicIPInvoker(request *model.BindDevServerPublicIpRequest) *BindDevServerPublicIPInvoker {
	requestDef := GenReqDefForBindDevServerPublicIP()
	return &BindDevServerPublicIPInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeDevServerOS 切换Lite Server服务器操作系统镜像
//
// 切换Lite Server服务器操作系统镜像接口用于更换Lite Server服务器当前使用的操作系统镜像。该接口适用于以下场景：当用户需要更换操作系统以适应不同的开发或测试需求时，可以通过此接口切换指定的Lite Server服务器操作系统镜像。使用该接口的前提条件是Lite Server服务器已存在且处于停止状态，用户具有切换操作系统的权限。切换操作完成后，Lite Server服务器将安装新的操作系统镜像，并重新进入运行状态，若Lite Server服务器不存在、已处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ChangeDevServerOS(request *model.ChangeDevServerOsRequest) (*model.ChangeDevServerOsResponse, error) {
	requestDef := GenReqDefForChangeDevServerOS()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeDevServerOsResponse), nil
	}
}

// ChangeDevServerOSInvoker 切换Lite Server服务器操作系统镜像
func (c *ModelArtsClient) ChangeDevServerOSInvoker(request *model.ChangeDevServerOsRequest) *ChangeDevServerOSInvoker {
	requestDef := GenReqDefForChangeDevServerOS()
	return &ChangeDevServerOSInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeHyperinstanceOS 切换Lite Server超节点服务器操作系统镜像
//
// 切换Lite Server超节点服务器操作系统镜像接口用于更换Lite Server超节点服务器当前使用的操作系统镜像。该接口适用于以下场景：当用户需要更换操作系统以适应不同的开发或测试需求时，可以通过此接口切换指定的Lite Server超节点服务器操作系统镜像。使用该接口的前提条件是Lite Server超节点服务器已存在且处于停止状态，用户具有切换操作系统的权限。切换操作完成后，Lite Server超节点服务器将安装新的操作系统镜像，并重新进入运行状态，若Lite Server超节点服务器不存在、已处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ChangeHyperinstanceOS(request *model.ChangeHyperinstanceOsRequest) (*model.ChangeHyperinstanceOsResponse, error) {
	requestDef := GenReqDefForChangeHyperinstanceOS()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeHyperinstanceOsResponse), nil
	}
}

// ChangeHyperinstanceOSInvoker 切换Lite Server超节点服务器操作系统镜像
func (c *ModelArtsClient) ChangeHyperinstanceOSInvoker(request *model.ChangeHyperinstanceOsRequest) *ChangeHyperinstanceOSInvoker {
	requestDef := GenReqDefForChangeHyperinstanceOS()
	return &ChangeHyperinstanceOSInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDevServer 创建Lite Server
//
// 创建Lite Server接口用于创建LiteServer弹性云服务器、裸金属服务器及超节点服务器。该接口适用于以下场景：用户需要根据业务需求快速部署和配置不同类型的服务器资源。使用该接口的前提条件是用户已登录且具有创建Lite Server的权限，并且需要提供服务器类型、规格、网络配置等必要参数。创建操作完成后，系统将返回新创建的Lite Server实例信息，包括实例ID、状态等。若用户无权限、参数配置错误或资源不足，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateDevServer(request *model.CreateDevServerRequest) (*model.CreateDevServerResponse, error) {
	requestDef := GenReqDefForCreateDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDevServerResponse), nil
	}
}

// CreateDevServerInvoker 创建Lite Server
func (c *ModelArtsClient) CreateDevServerInvoker(request *model.CreateDevServerRequest) *CreateDevServerInvoker {
	requestDef := GenReqDefForCreateDevServer()
	return &CreateDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDevServerJob 创建Lite Server任务
//
// 创建Lite Server任务接口用于在Lite Server上创建新的任务。该接口适用于以下场景：当用户需要在Lite Server上启动新的开发、测试或部署任务时，可以通过此接口创建并配置任务。使用该接口的前提条件是用户具有创建任务的权限，并且提供的任务配置参数符合要求。创建操作完成后，新的Lite Server任务将被成功创建，并返回任务ID和其他相关信息。若用户无权限操作、提供的参数不正确或系统资源不足，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateDevServerJob(request *model.CreateDevServerJobRequest) (*model.CreateDevServerJobResponse, error) {
	requestDef := GenReqDefForCreateDevServerJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDevServerJobResponse), nil
	}
}

// CreateDevServerJobInvoker 创建Lite Server任务
func (c *ModelArtsClient) CreateDevServerJobInvoker(request *model.CreateDevServerJobRequest) *CreateDevServerJobInvoker {
	requestDef := GenReqDefForCreateDevServerJob()
	return &CreateDevServerJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHyperCluster 创建Hyper Cluster
//
// 创建Hyper Cluster接口用于在系统中创建一个新的Hyper Cluster。该接口适用于以下场景：当用户需要使用超节点网络时，可以通过此接口创建Hyper Cluster。使用该接口的前提条件是用户已登录并具有创建Hyper Cluster的权限，且系统中已配置了必要的资源。创建操作完成后，将生成一个新的超节点网络，并返回超节点网络的详细信息，包括ID、名称、子网信息等。若用户无权限操作、系统中缺少必要的资源或配置参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateHyperCluster(request *model.CreateHyperClusterRequest) (*model.CreateHyperClusterResponse, error) {
	requestDef := GenReqDefForCreateHyperCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHyperClusterResponse), nil
	}
}

// CreateHyperClusterInvoker 创建Hyper Cluster
func (c *ModelArtsClient) CreateHyperClusterInvoker(request *model.CreateHyperClusterRequest) *CreateHyperClusterInvoker {
	requestDef := GenReqDefForCreateHyperCluster()
	return &CreateHyperClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateHyperinstanceTags 创建Lite Server超节点标签
//
// 创建Lite Server超节点标签接口用于为Lite Server超节点添加自定义标签。该接口适用于以下场景：当用户需要对Lite Server超节点进行分类管理或标记特定信息时，可以通过此接口为指定的超节点创建标签。使用该接口的前提条件是Lite Server超节点已存在，用户具有创建标签的权限。创建操作完成后，标签将被成功添加到指定的超节点上，用户可以通过标签进行快速查找和管理。若Lite Server超节点不存在、标签已存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateHyperinstanceTags(request *model.CreateHyperinstanceTagsRequest) (*model.CreateHyperinstanceTagsResponse, error) {
	requestDef := GenReqDefForCreateHyperinstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateHyperinstanceTagsResponse), nil
	}
}

// CreateHyperinstanceTagsInvoker 创建Lite Server超节点标签
func (c *ModelArtsClient) CreateHyperinstanceTagsInvoker(request *model.CreateHyperinstanceTagsRequest) *CreateHyperinstanceTagsInvoker {
	requestDef := GenReqDefForCreateHyperinstanceTags()
	return &CreateHyperinstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateRoceNetwork 创建RoCE网络
//
// 创建RoCE网络接口用于在系统中创建一个新的RoCE网络。该接口适用于以下场景：当用户需要为高性能计算或低延迟应用创建专用的RoCE网络时，可以通过此接口创建并配置RoCE网络。使用该接口的前提条件是用户已登录并具有创建RoCE网络的权限，且系统中已配置了必要的网络资源。创建操作完成后，将生成一个新的RoCE网络，并返回网络的详细信息，包括网络ID、子网信息、配置参数等。若用户无权限操作、系统中缺少必要的网络资源或网络配置参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateRoceNetwork(request *model.CreateRoceNetworkRequest) (*model.CreateRoceNetworkResponse, error) {
	requestDef := GenReqDefForCreateRoceNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRoceNetworkResponse), nil
	}
}

// CreateRoceNetworkInvoker 创建RoCE网络
func (c *ModelArtsClient) CreateRoceNetworkInvoker(request *model.CreateRoceNetworkRequest) *CreateRoceNetworkInvoker {
	requestDef := GenReqDefForCreateRoceNetwork()
	return &CreateRoceNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDevServer 删除Lite Server实例
//
// 删除Lite Server实例接口用于移除已创建的Lite Server实例。该接口适用于以下场景：当Lite Server按需实例不再需要使用时或者创建失败的实例以及处于ERROR状态时，用户可通过此接口删除指定的Lite Server实例。使用该接口的前提条件是Lite Server实例已存在且用户具有管理员权限。删除操作完成后，Lite Server实例将被永久移除，相关资源也将被清理。若Lite Server实例不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteDevServer(request *model.DeleteDevServerRequest) (*model.DeleteDevServerResponse, error) {
	requestDef := GenReqDefForDeleteDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDevServerResponse), nil
	}
}

// DeleteDevServerInvoker 删除Lite Server实例
func (c *ModelArtsClient) DeleteDevServerInvoker(request *model.DeleteDevServerRequest) *DeleteDevServerInvoker {
	requestDef := GenReqDefForDeleteDevServer()
	return &DeleteDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDevServerJobs 批量删除Lite Server Job
//
// 批量删除Lite Server Job接口用于批量移除已创建的Lite Server Job。该接口适用于以下场景：当多个Lite Server Job已完成、配置错误或需要清理资源时，用户可以通过此接口批量删除指定的Lite Server Job。使用该接口的前提条件是目标Lite Server Job已存在且用户具有管理员权限。删除操作完成后，指定的Lite Server Job将被永久移除，相关资源和配置也将被清理。若目标Lite Server Job不存在、用户无权限操作或请求参数不正确，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteDevServerJobs(request *model.DeleteDevServerJobsRequest) (*model.DeleteDevServerJobsResponse, error) {
	requestDef := GenReqDefForDeleteDevServerJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDevServerJobsResponse), nil
	}
}

// DeleteDevServerJobsInvoker 批量删除Lite Server Job
func (c *ModelArtsClient) DeleteDevServerJobsInvoker(request *model.DeleteDevServerJobsRequest) *DeleteDevServerJobsInvoker {
	requestDef := GenReqDefForDeleteDevServerJobs()
	return &DeleteDevServerJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHyperCluster 删除Hyper Cluster实例
//
// 删除Hyper Cluster实例接口用于移除已创建的Hyper Cluster。该接口适用于以下场景：当超节点网络配置错误或需要清理资源时，用户可通过此接口删除指定的超节点网络。使用该接口的前提条件是Hyper Cluster实例已存在且用户具有管理员权限。删除操作完成后，超节点网络将被永久移除，相关资源和配置也将被清理。若Hyper Cluster实例不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteHyperCluster(request *model.DeleteHyperClusterRequest) (*model.DeleteHyperClusterResponse, error) {
	requestDef := GenReqDefForDeleteHyperCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHyperClusterResponse), nil
	}
}

// DeleteHyperClusterInvoker 删除Hyper Cluster实例
func (c *ModelArtsClient) DeleteHyperClusterInvoker(request *model.DeleteHyperClusterRequest) *DeleteHyperClusterInvoker {
	requestDef := GenReqDefForDeleteHyperCluster()
	return &DeleteHyperClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHyperinstance 删除Lite Server超节点实例
//
// 删除Lite Server超节点实例接口用于删除按需超节点实例同时移除处于ERROR状态的Lite Server超节点实例。该接口适用于以下场景：当超节点实例因创建失败、或其他原因进入ERROR状态；按需超节点实例，用户可以通过此接口删除指定的超节点实例。使用该接口的前提条件是用户已登录并具有删除超节点实例的权限，且指定的超节点实例是按需且处于运行状态、或者处于ERROR状态。删除操作完成后，指定的超节点实例将被永久移除，相关资源也将被清理。若指定的超节点实例不存在、未处于ERROR状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteHyperinstance(request *model.DeleteHyperinstanceRequest) (*model.DeleteHyperinstanceResponse, error) {
	requestDef := GenReqDefForDeleteHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHyperinstanceResponse), nil
	}
}

// DeleteHyperinstanceInvoker 删除Lite Server超节点实例
func (c *ModelArtsClient) DeleteHyperinstanceInvoker(request *model.DeleteHyperinstanceRequest) *DeleteHyperinstanceInvoker {
	requestDef := GenReqDefForDeleteHyperinstance()
	return &DeleteHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteHyperinstanceTags 删除Lite Server超节点标签
//
// 删除Lite Server超节点标签接口用于移除已创建的Lite Server超节点标签。该接口适用于以下场景：当用户需要清理不再需要的标签或修正标签错误时，可以通过此接口删除指定的超节点标签。使用该接口的前提条件是Lite Server超节点已存在，且该超节点上已存在要删除的标签，用户具有删除标签的权限。删除操作完成后，指定的标签将从超节点上移除，超节点的其他配置和数据保持不变。若Lite Server超节点不存在、标签不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteHyperinstanceTags(request *model.DeleteHyperinstanceTagsRequest) (*model.DeleteHyperinstanceTagsResponse, error) {
	requestDef := GenReqDefForDeleteHyperinstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteHyperinstanceTagsResponse), nil
	}
}

// DeleteHyperinstanceTagsInvoker 删除Lite Server超节点标签
func (c *ModelArtsClient) DeleteHyperinstanceTagsInvoker(request *model.DeleteHyperinstanceTagsRequest) *DeleteHyperinstanceTagsInvoker {
	requestDef := GenReqDefForDeleteHyperinstanceTags()
	return &DeleteHyperinstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DetachDevServerVolume Lite Server服务器卸载磁盘
//
// Lite Server服务器卸载磁盘接口用于从Lite Server服务器上卸载已挂载的磁盘。该接口适用于以下场景：当用户需要释放存储资源或重新分配磁盘时，可以通过此接口卸载指定的磁盘。使用该接口的前提条件是Lite Server服务器已创建且处于运行状态、或者停止状态，用户具有卸载磁盘的权限，且指定的磁盘已挂载到服务器上。卸载操作完成后，磁盘将从Lite Server服务器上成功卸载，用户可以将其挂载到其他服务器或进行其他操作。若Lite Server服务器不存在、指定的磁盘未挂载到服务器上，或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DetachDevServerVolume(request *model.DetachDevServerVolumeRequest) (*model.DetachDevServerVolumeResponse, error) {
	requestDef := GenReqDefForDetachDevServerVolume()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DetachDevServerVolumeResponse), nil
	}
}

// DetachDevServerVolumeInvoker Lite Server服务器卸载磁盘
func (c *ModelArtsClient) DetachDevServerVolumeInvoker(request *model.DetachDevServerVolumeRequest) *DetachDevServerVolumeInvoker {
	requestDef := GenReqDefForDetachDevServerVolume()
	return &DetachDevServerVolumeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDevServerImage 查询Lite Server镜像详情
//
// 查询Lite Server镜像详情接口用于获取指定Lite Server镜像的详细信息。该接口适用于以下场景：当用户需要了解某个Lite Server镜像的具体配置和属性，以便在创建或调整Lite Server实例时选择合适的镜像时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询镜像详情的权限，且指定的镜像已存在。查询操作完成后，接口将返回指定Lite Server镜像的详细信息，包括镜像ID、名称、操作系统、版本、创建时间等。若用户无权限操作、指定的镜像不存在或镜像ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetDevServerImage(request *model.GetDevServerImageRequest) (*model.GetDevServerImageResponse, error) {
	requestDef := GenReqDefForGetDevServerImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDevServerImageResponse), nil
	}
}

// GetDevServerImageInvoker 查询Lite Server镜像详情
func (c *ModelArtsClient) GetDevServerImageInvoker(request *model.GetDevServerImageRequest) *GetDevServerImageInvoker {
	requestDef := GenReqDefForGetDevServerImage()
	return &GetDevServerImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDevServerJob 查询Lite Server Job详情
//
// 查询Lite Server Job详情接口用于获取指定Lite Server Job的详细信息。该接口适用于以下场景：当用户需要查看某个Lite Server Job的执行状态、配置参数、日志信息等详细数据时，可以通过此接口获取相关信息。使用该接口的前提条件是目标Lite Server Job已存在且用户具有查看权限。查询操作完成后，接口将返回指定Lite Server Job的详细信息，包括但不限于Job ID、状态、创建时间、执行时间、配置参数和日志等。若目标Lite Server Job不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetDevServerJob(request *model.GetDevServerJobRequest) (*model.GetDevServerJobResponse, error) {
	requestDef := GenReqDefForGetDevServerJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDevServerJobResponse), nil
	}
}

// GetDevServerJobInvoker 查询Lite Server Job详情
func (c *ModelArtsClient) GetDevServerJobInvoker(request *model.GetDevServerJobRequest) *GetDevServerJobInvoker {
	requestDef := GenReqDefForGetDevServerJob()
	return &GetDevServerJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDevServerJobService 获取Lite Server 部署服务详情
//
// 根据服务id获取Lite Server部署服务详情。该接口适用于以下场景：当用户需要查看部署服务详情，以便查看已部署服务的状态、api等信息时，可以通过此接口获取服务详情。使用该接口的前提条件是用户具有查看服务的权限。查询操作完成后，接口将返回此部署服务的详细信息，包括名称、状态、描述、所用模型、实例详情等信息。若用户无权限操作或无相应id，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetDevServerJobService(request *model.GetDevServerJobServiceRequest) (*model.GetDevServerJobServiceResponse, error) {
	requestDef := GenReqDefForGetDevServerJobService()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDevServerJobServiceResponse), nil
	}
}

// GetDevServerJobServiceInvoker 获取Lite Server 部署服务详情
func (c *ModelArtsClient) GetDevServerJobServiceInvoker(request *model.GetDevServerJobServiceRequest) *GetDevServerJobServiceInvoker {
	requestDef := GenReqDefForGetDevServerJobService()
	return &GetDevServerJobServiceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDevServerJobTemplate 获取Lite Server Job模板详情
//
// 获取Lite Server Job模板详情接口用于获取指定Lite Server Job模板的详细信息。该接口适用于以下场景：当用户需要查看某个特定Job模板的详细配置，以便了解其参数设置、使用说明等信息时，可以通过此接口获取模板详情。查询操作完成后，接口将返回指定模板的详细信息，包括模板ID、名称、描述、配置参数等。若目标模板不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetDevServerJobTemplate(request *model.GetDevServerJobTemplateRequest) (*model.GetDevServerJobTemplateResponse, error) {
	requestDef := GenReqDefForGetDevServerJobTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDevServerJobTemplateResponse), nil
	}
}

// GetDevServerJobTemplateInvoker 获取Lite Server Job模板详情
func (c *ModelArtsClient) GetDevServerJobTemplateInvoker(request *model.GetDevServerJobTemplateRequest) *GetDevServerJobTemplateInvoker {
	requestDef := GenReqDefForGetDevServerJobTemplate()
	return &GetDevServerJobTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetDevServerOperation 查询Operation详情
//
// 查询Operation详情接口用于获取指定Operation的详细信息。该接口适用于以下场景：当用户需要了解某个Operation的具体执行情况和状态，以便进行故障排查或操作审计时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询Operation详情的权限，且指定的Operation已存在。查询操作完成后，接口将返回指定Operation的详细信息，包括Operation ID、操作类型、执行状态、开始时间、结束时间、操作结果等。若用户无权限操作、指定的Operation不存在或Operation ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetDevServerOperation(request *model.GetDevServerOperationRequest) (*model.GetDevServerOperationResponse, error) {
	requestDef := GenReqDefForGetDevServerOperation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetDevServerOperationResponse), nil
	}
}

// GetDevServerOperationInvoker 查询Operation详情
func (c *ModelArtsClient) GetDevServerOperationInvoker(request *model.GetDevServerOperationRequest) *GetDevServerOperationInvoker {
	requestDef := GenReqDefForGetDevServerOperation()
	return &GetDevServerOperationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetHyperCluster 查询Hyper Cluster实例详情
//
// 查询Hyper Cluster实例详情接口用于获取指定Hyper Cluster实例的详细信息。该接口适用于以下场景：当用户需要了解某个超节点网络的具体配置和状态，以便进行管理和监控时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询Hyper Cluster详情的权限，且指定的超节点网络已存在。查询操作完成后，接口将返回指定超节点网络的详细信息，包括ID、名称、子网信息等。若用户无权限操作、指定的超节点网络不存在或ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetHyperCluster(request *model.GetHyperClusterRequest) (*model.GetHyperClusterResponse, error) {
	requestDef := GenReqDefForGetHyperCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetHyperClusterResponse), nil
	}
}

// GetHyperClusterInvoker 查询Hyper Cluster实例详情
func (c *ModelArtsClient) GetHyperClusterInvoker(request *model.GetHyperClusterRequest) *GetHyperClusterInvoker {
	requestDef := GenReqDefForGetHyperCluster()
	return &GetHyperClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetHyperinstance 查询指定超节点实例详情
//
// 查询指定超节点实例详情接口用于获取特定Lite Server超节点实例的详细信息。该接口适用于以下场景：当用户需要查看某个具体超节点实例的配置、状态和使用情况时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询超节点实例的权限，且指定的超节点实例已存在。查询操作完成后，接口将返回指定超节点实例的详细信息，包括实例ID、操作系统、运行状态、资源使用情况等。若用户无权限操作、指定的超节点实例不存在或实例ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetHyperinstance(request *model.GetHyperinstanceRequest) (*model.GetHyperinstanceResponse, error) {
	requestDef := GenReqDefForGetHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetHyperinstanceResponse), nil
	}
}

// GetHyperinstanceInvoker 查询指定超节点实例详情
func (c *ModelArtsClient) GetHyperinstanceInvoker(request *model.GetHyperinstanceRequest) *GetHyperinstanceInvoker {
	requestDef := GenReqDefForGetHyperinstance()
	return &GetHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetScaleEvaluationsDevServer 查询Lite Server超节点扩缩容支持规格列表及容量测算
//
// 查询Lite Server超节点扩缩容支持规格列表及容量测算接口用于获取Lite Server超节点支持的扩缩容规格列表，并进行容量测算。该接口适用于以下场景：当用户需要了解Lite Server超节点支持的扩缩容选项，以便在调整超节点资源时选择合适的规格，并评估扩缩容后的资源需求时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询超节点扩缩容规格的权限，且指定的超节点已存在。查询操作完成后，接口将返回支持的扩缩容规格列表及容量测算结果，包括规格ID、CPU、内存、存储等详细配置和扩缩容后的资源使用情况。若用户无权限操作、指定的超节点不存在或系统中没有可用的扩缩容规格，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetScaleEvaluationsDevServer(request *model.GetScaleEvaluationsDevServerRequest) (*model.GetScaleEvaluationsDevServerResponse, error) {
	requestDef := GenReqDefForGetScaleEvaluationsDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetScaleEvaluationsDevServerResponse), nil
	}
}

// GetScaleEvaluationsDevServerInvoker 查询Lite Server超节点扩缩容支持规格列表及容量测算
func (c *ModelArtsClient) GetScaleEvaluationsDevServerInvoker(request *model.GetScaleEvaluationsDevServerRequest) *GetScaleEvaluationsDevServerInvoker {
	requestDef := GenReqDefForGetScaleEvaluationsDevServer()
	return &GetScaleEvaluationsDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GetTopologies 查询实例的Tor信息
//
// 查询实例的Tor信息接口用于获取指定实例的Top-of-Rack（Tor）交换机相关信息。该接口适用于以下场景：当用户需要了解实例连接的Tor交换机的详细信息，以便进行网络配置时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询实例Tor信息的权限，且指定的实例已存在。查询操作完成后，接口将返回指定实例的Tor信息。若用户无权限操作、指定的实例不存在或实例未连接到Tor交换机，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) GetTopologies(request *model.GetTopologiesRequest) (*model.GetTopologiesResponse, error) {
	requestDef := GenReqDefForGetTopologies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GetTopologiesResponse), nil
	}
}

// GetTopologiesInvoker 查询实例的Tor信息
func (c *ModelArtsClient) GetTopologiesInvoker(request *model.GetTopologiesRequest) *GetTopologiesInvoker {
	requestDef := GenReqDefForGetTopologies()
	return &GetTopologiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllDevServers 查询租户Lite Server列表
//
// 查询租户Lite Server列表接口用于获取指定租户的所有Lite Server实例信息。该接口适用于以下场景：当用户需要查看其租户下所有Lite Server实例的详细信息，以便进行管理和监控时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询租户Lite Server列表的权限。查询操作完成后，接口将返回租户下所有Lite Server实例的详细信息，包括实例ID、名称、状态、资源配置等。若用户无权限操作或租户下没有Lite Server实例，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListAllDevServers(request *model.ListAllDevServersRequest) (*model.ListAllDevServersResponse, error) {
	requestDef := GenReqDefForListAllDevServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllDevServersResponse), nil
	}
}

// ListAllDevServersInvoker 查询租户Lite Server列表
func (c *ModelArtsClient) ListAllDevServersInvoker(request *model.ListAllDevServersRequest) *ListAllDevServersInvoker {
	requestDef := GenReqDefForListAllDevServers()
	return &ListAllDevServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllHyperinstances 查询租户Hyperinstance列表
//
// 查询租户Hyperinstance列表接口用于获取指定租户的所有Hyperinstance实例信息。该接口适用于以下场景：当用户需要查看其租户下所有Hyperinstance实例的详细信息，以便进行管理和监控时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询租户Hyperinstance列表的权限。查询操作完成后，接口将返回租户下所有Hyperinstance实例的详细信息，包括实例ID、名称、状态、资源配置等。若用户无权限操作或租户下没有Hyperinstance实例，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListAllHyperinstances(request *model.ListAllHyperinstancesRequest) (*model.ListAllHyperinstancesResponse, error) {
	requestDef := GenReqDefForListAllHyperinstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllHyperinstancesResponse), nil
	}
}

// ListAllHyperinstancesInvoker 查询租户Hyperinstance列表
func (c *ModelArtsClient) ListAllHyperinstancesInvoker(request *model.ListAllHyperinstancesRequest) *ListAllHyperinstancesInvoker {
	requestDef := GenReqDefForListAllHyperinstances()
	return &ListAllHyperinstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServerFlavors 查询规格列表
//
// 查询规格列表接口用于获取系统中所有可用的资源规格信息。该接口适用于以下场景：当用户需要了解可用的资源规格，以便在创建或调整Lite Server实例时选择合适的配置时，可以通过此接口获取规格列表。使用该接口的前提条件是用户已登录并具有查询规格的权限。查询操作完成后，接口将返回所有可用的资源规格信息，包括规格ID、CPU、内存、存储等详细配置。若用户无权限操作或系统中没有可用的资源规格，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServerFlavors(request *model.ListDevServerFlavorsRequest) (*model.ListDevServerFlavorsResponse, error) {
	requestDef := GenReqDefForListDevServerFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServerFlavorsResponse), nil
	}
}

// ListDevServerFlavorsInvoker 查询规格列表
func (c *ModelArtsClient) ListDevServerFlavorsInvoker(request *model.ListDevServerFlavorsRequest) *ListDevServerFlavorsInvoker {
	requestDef := GenReqDefForListDevServerFlavors()
	return &ListDevServerFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServerImages 查询Lite Server镜像列表
//
// 查询Lite Server镜像列表接口用于获取系统中所有可用的Lite Server镜像信息。该接口适用于以下场景：当用户需要了解可用的Lite Server镜像，以便在创建或调整Lite Server实例时选择合适的镜像时，可以通过此接口获取镜像列表。使用该接口的前提条件是用户已登录并具有查询镜像列表的权限。查询操作完成后，接口将返回所有可用的Lite Server镜像信息，包括镜像ID、名称、架构类型等。若用户无权限操作或系统中没有可用的镜像，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServerImages(request *model.ListDevServerImagesRequest) (*model.ListDevServerImagesResponse, error) {
	requestDef := GenReqDefForListDevServerImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServerImagesResponse), nil
	}
}

// ListDevServerImagesInvoker 查询Lite Server镜像列表
func (c *ModelArtsClient) ListDevServerImagesInvoker(request *model.ListDevServerImagesRequest) *ListDevServerImagesInvoker {
	requestDef := GenReqDefForListDevServerImages()
	return &ListDevServerImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServerJobTemplates 获取Lite Server Job模板列表
//
// 获取Lite Server Job模板列表接口用于获取可用的Lite Server Job模板列表。该接口适用于以下场景：当用户需要查看可用的Job模板，以便选择合适的模板来创建新的Lite Server任务时，可以通过此接口获取模板列表。查询操作完成后，接口将返回所有可用的Lite Server Job模板列表，包括模板ID、名称、描述等信息。若系统中无可用模板，接口将返回相应的信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServerJobTemplates(request *model.ListDevServerJobTemplatesRequest) (*model.ListDevServerJobTemplatesResponse, error) {
	requestDef := GenReqDefForListDevServerJobTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServerJobTemplatesResponse), nil
	}
}

// ListDevServerJobTemplatesInvoker 获取Lite Server Job模板列表
func (c *ModelArtsClient) ListDevServerJobTemplatesInvoker(request *model.ListDevServerJobTemplatesRequest) *ListDevServerJobTemplatesInvoker {
	requestDef := GenReqDefForListDevServerJobTemplates()
	return &ListDevServerJobTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServerJobs 查询Lite Server Job列表
//
// 查询Lite Server Job列表接口用于获取Lite Server Job的列表信息，并支持按照状态、ID等相关字段进行过滤。该接口适用于以下场景：当用户需要查看多个Lite Server Job的概要信息，例如在监控作业状态、排查问题或进行日常管理时，可以通过此接口获取符合过滤条件的Job列表。使用该接口的前提条件是用户具有查看权限。查询操作完成后，接口将返回符合条件的Lite Server Job列表，包括每个Job的ID、状态、创建时间等基本信息。若用户无权限操作或请求参数不正确，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServerJobs(request *model.ListDevServerJobsRequest) (*model.ListDevServerJobsResponse, error) {
	requestDef := GenReqDefForListDevServerJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServerJobsResponse), nil
	}
}

// ListDevServerJobsInvoker 查询Lite Server Job列表
func (c *ModelArtsClient) ListDevServerJobsInvoker(request *model.ListDevServerJobsRequest) *ListDevServerJobsInvoker {
	requestDef := GenReqDefForListDevServerJobs()
	return &ListDevServerJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServerPublicIP 查询已绑定的EIP
//
// 查询已绑定的EIP接口用于获取已绑定到Lite Server服务器上的弹性公网IP（EIP）信息。该接口适用于以下场景：当用户需要查看Lite Server服务器上已绑定的EIP及其详细信息时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询EIP的权限，且指定的Lite Server服务器已存在。查询操作完成后，接口将返回已绑定到Lite Server服务器上的EIP的详细信息，包括EIP地址、绑定时间、状态等。若Lite Server服务器不存在、未绑定EIP或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServerPublicIP(request *model.ListDevServerPublicIpRequest) (*model.ListDevServerPublicIpResponse, error) {
	requestDef := GenReqDefForListDevServerPublicIP()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServerPublicIpResponse), nil
	}
}

// ListDevServerPublicIPInvoker 查询已绑定的EIP
func (c *ModelArtsClient) ListDevServerPublicIPInvoker(request *model.ListDevServerPublicIpRequest) *ListDevServerPublicIPInvoker {
	requestDef := GenReqDefForListDevServerPublicIP()
	return &ListDevServerPublicIPInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDevServers 查询用户所有Lite Server实例列表
//
// 查询用户所有Lite Server实例列表接口用于获取用户名下所有Lite Server实例的详细信息。该接口适用于以下场景：用户需要查看其所有Lite Server实例的状态、配置等信息，以便进行资源管理和监控。使用该接口的前提条件是用户已登录且具有查看Lite Server实例的权限。调用此接口后，系统将返回用户名下所有Lite Server实例的列表，包括实例ID、名称、状态、创建时间等信息。若用户无权限或未登录，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListDevServers(request *model.ListDevServersRequest) (*model.ListDevServersResponse, error) {
	requestDef := GenReqDefForListDevServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevServersResponse), nil
	}
}

// ListDevServersInvoker 查询用户所有Lite Server实例列表
func (c *ModelArtsClient) ListDevServersInvoker(request *model.ListDevServersRequest) *ListDevServersInvoker {
	requestDef := GenReqDefForListDevServers()
	return &ListDevServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHyperCluster 查询Hyper Cluster详情列表
//
// 查询Hyper Cluster详情列表接口用于获取所有Hyper Cluster的详细信息。该接口适用于以下场景：当用户需要了解系统中所有超节点网络的配置和状态时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询Hyper Cluster详情的权限。查询操作完成后，接口将返回所有超节点网络的详细信息，包括ID、名称、子网信息等。若用户无权限操作或系统中没有Hyper Cluster，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListHyperCluster(request *model.ListHyperClusterRequest) (*model.ListHyperClusterResponse, error) {
	requestDef := GenReqDefForListHyperCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHyperClusterResponse), nil
	}
}

// ListHyperClusterInvoker 查询Hyper Cluster详情列表
func (c *ModelArtsClient) ListHyperClusterInvoker(request *model.ListHyperClusterRequest) *ListHyperClusterInvoker {
	requestDef := GenReqDefForListHyperCluster()
	return &ListHyperClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHyperinstanceClustersCapacity 查询超节点hyperinstance-clusters逻辑容量测算结果
//
// 查询超节点hyperinstance-clusters逻辑容量测算结果接口用于获取指定超节点集群的逻辑容量测算结果。该接口适用于以下场景：当用户需要了解超节点集群的资源使用情况和容量规划，以便进行资源管理和优化时，可以通过此接口获取逻辑容量测算结果。使用该接口的前提条件是用户已登录并具有查询超节点集群逻辑容量的权限，且指定的超节点集群已存在。查询操作完成后，接口将返回指定超节点集群的逻辑容量测算结果，包括可用容量信息。若用户无权限操作、指定的超节点集群不存在或集群ID无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListHyperinstanceClustersCapacity(request *model.ListHyperinstanceClustersCapacityRequest) (*model.ListHyperinstanceClustersCapacityResponse, error) {
	requestDef := GenReqDefForListHyperinstanceClustersCapacity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHyperinstanceClustersCapacityResponse), nil
	}
}

// ListHyperinstanceClustersCapacityInvoker 查询超节点hyperinstance-clusters逻辑容量测算结果
func (c *ModelArtsClient) ListHyperinstanceClustersCapacityInvoker(request *model.ListHyperinstanceClustersCapacityRequest) *ListHyperinstanceClustersCapacityInvoker {
	requestDef := GenReqDefForListHyperinstanceClustersCapacity()
	return &ListHyperinstanceClustersCapacityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHyperinstances 查询用户所有超节点实例详情
//
// 查询用户所有超节点实例详情接口用于获取用户所有Lite Server超节点实例的详细信息。该接口适用于以下场景：当用户需要查看其所有超节点实例的配置、状态和使用情况时，可以通过此接口获取相关信息。使用该接口的前提条件是用户已登录并具有查询超节点实例的权限。查询操作完成后，接口将返回所有超节点实例的详细信息，包括实例ID、操作系统、运行状态、资源使用情况等。若用户无权限操作或没有超节点实例，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListHyperinstances(request *model.ListHyperinstancesRequest) (*model.ListHyperinstancesResponse, error) {
	requestDef := GenReqDefForListHyperinstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHyperinstancesResponse), nil
	}
}

// ListHyperinstancesInvoker 查询用户所有超节点实例详情
func (c *ModelArtsClient) ListHyperinstancesInvoker(request *model.ListHyperinstancesRequest) *ListHyperinstancesInvoker {
	requestDef := GenReqDefForListHyperinstances()
	return &ListHyperinstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// QueryHyperinstanceTags 查询Lite Server超节点标签
//
// 查询Lite Server超节点标签接口用于获取Lite Server超节点上的所有标签信息。该接口适用于以下场景：当用户需要查看或管理Lite Server超节点的标签时，可以通过此接口查询指定超节点上的所有标签。使用该接口的前提条件是Lite Server超节点已存在，用户具有查询标签的权限。查询操作完成后，接口将返回超节点上的所有标签信息，包括标签名称和相关属性。若Lite Server超节点不存在或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) QueryHyperinstanceTags(request *model.QueryHyperinstanceTagsRequest) (*model.QueryHyperinstanceTagsResponse, error) {
	requestDef := GenReqDefForQueryHyperinstanceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.QueryHyperinstanceTagsResponse), nil
	}
}

// QueryHyperinstanceTagsInvoker 查询Lite Server超节点标签
func (c *ModelArtsClient) QueryHyperinstanceTagsInvoker(request *model.QueryHyperinstanceTagsRequest) *QueryHyperinstanceTagsInvoker {
	requestDef := GenReqDefForQueryHyperinstanceTags()
	return &QueryHyperinstanceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootDevServer 重启Lite Server实例
//
// 重启Lite Server实例接口用于重启正在运行的Lite Server实例。该接口适用于以下场景：当用户需要重启实例以应用配置更改、解决运行问题或进行系统维护时，可以通过此接口重启指定的Lite Server实例。使用该接口的前提条件是Lite Server实例已创建且处于运行状态，用户具有重启实例的权限。重启操作完成后，Lite Server实例将重新启动并进入运行状态，用户可以继续使用实例提供的服务。若Lite Server实例不存在、已处于停止状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) RebootDevServer(request *model.RebootDevServerRequest) (*model.RebootDevServerResponse, error) {
	requestDef := GenReqDefForRebootDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootDevServerResponse), nil
	}
}

// RebootDevServerInvoker 重启Lite Server实例
func (c *ModelArtsClient) RebootDevServerInvoker(request *model.RebootDevServerRequest) *RebootDevServerInvoker {
	requestDef := GenReqDefForRebootDevServer()
	return &RebootDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ReinstallDevServerOS 重装Lite Server服务器操作系统镜像
//
// 重装Lite Server服务器操作系统镜像接口用于重新安装Lite Server服务器的操作系统镜像。该接口适用于以下场景：当用户需要更新操作系统版本、修复系统故障或重新配置系统环境时，可以通过此接口重装指定的Lite Server服务器操作系统镜像。使用该接口的前提条件是Lite Server服务器已存在且处于停止状态，用户具有重装操作系统的权限。重装操作完成后，Lite Server服务器将安装新的操作系统镜像，并重新进入运行状态，若Lite Server服务器不存在、已处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ReinstallDevServerOS(request *model.ReinstallDevServerOsRequest) (*model.ReinstallDevServerOsResponse, error) {
	requestDef := GenReqDefForReinstallDevServerOS()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ReinstallDevServerOsResponse), nil
	}
}

// ReinstallDevServerOSInvoker 重装Lite Server服务器操作系统镜像
func (c *ModelArtsClient) ReinstallDevServerOSInvoker(request *model.ReinstallDevServerOsRequest) *ReinstallDevServerOSInvoker {
	requestDef := GenReqDefForReinstallDevServerOS()
	return &ReinstallDevServerOSInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ScaleDownHyperinstance 缩容Lite Server超节点
//
// 缩容Lite Server超节点接口用于减少Lite Server超节点的资源容量。该接口适用于以下场景：当用户需要降低Lite Server超节点的资源使用，以节省成本或优化资源分配时，可以通过此接口进行缩容。使用该接口的前提条件是用户已登录并具有缩容超节点的权限，且指定的超节点已存在且处于运行状态。缩容操作完成后，超节点的资源容量将根据指定的规格进行调整，用户可以立即使用减少后的资源。若用户无权限操作、指定的超节点不存在、超节点已处于最小容量或指定的缩容规格无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ScaleDownHyperinstance(request *model.ScaleDownHyperinstanceRequest) (*model.ScaleDownHyperinstanceResponse, error) {
	requestDef := GenReqDefForScaleDownHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ScaleDownHyperinstanceResponse), nil
	}
}

// ScaleDownHyperinstanceInvoker 缩容Lite Server超节点
func (c *ModelArtsClient) ScaleDownHyperinstanceInvoker(request *model.ScaleDownHyperinstanceRequest) *ScaleDownHyperinstanceInvoker {
	requestDef := GenReqDefForScaleDownHyperinstance()
	return &ScaleDownHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ScaleUpHyperinstance 扩容Lite Server超节点
//
// 扩容Lite Server超节点接口用于增加Lite Server超节点的资源容量。该接口适用于以下场景：当用户需要提升Lite Server超节点的性能，以支持更多的负载或更大的数据处理需求时，可以通过此接口进行扩容。使用该接口的前提条件是用户已登录并具有扩容超节点的权限，且指定的超节点已存在且处于运行状态。扩容操作完成后，超节点的资源容量将根据指定的规格进行调整，用户可以立即使用增加的资源。若用户无权限操作、指定的超节点不存在、超节点已处于最大容量或指定的扩容规格无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ScaleUpHyperinstance(request *model.ScaleUpHyperinstanceRequest) (*model.ScaleUpHyperinstanceResponse, error) {
	requestDef := GenReqDefForScaleUpHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ScaleUpHyperinstanceResponse), nil
	}
}

// ScaleUpHyperinstanceInvoker 扩容Lite Server超节点
func (c *ModelArtsClient) ScaleUpHyperinstanceInvoker(request *model.ScaleUpHyperinstanceRequest) *ScaleUpHyperinstanceInvoker {
	requestDef := GenReqDefForScaleUpHyperinstance()
	return &ScaleUpHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDevServer 查询Lite Server实例详情
//
// 查询Lite Server实例详情接口用于获取指定Lite Server实例的详细信息。该接口适用于以下场景：用户需要查看特定Lite Server实例的配置、状态、网络信息等详细数据，以便进行故障排查、资源管理和监控。使用该接口的前提条件是用户已登录且具有查看Lite Server实例的权限，并且需要提供有效的实例ID。查询操作完成后，系统将返回指定Lite Server实例的详细信息，包括实例ID、名称、状态、配置、网络配置等。若用户无权限、实例ID无效或实例不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowDevServer(request *model.ShowDevServerRequest) (*model.ShowDevServerResponse, error) {
	requestDef := GenReqDefForShowDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDevServerResponse), nil
	}
}

// ShowDevServerInvoker 查询Lite Server实例详情
func (c *ModelArtsClient) ShowDevServerInvoker(request *model.ShowDevServerRequest) *ShowDevServerInvoker {
	requestDef := GenReqDefForShowDevServer()
	return &ShowDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartDevServer 启动Lite Server实例
//
// 启动Lite Server实例接口用于启动已创建但未运行的Lite Server实例。该接口适用于以下场景：当用户需要开始使用Lite Server实例进行开发或测试时，可以通过此接口启动指定的Lite Server实例。使用该接口的前提条件是Lite Server实例已创建且处于停止状态，用户具有启动实例的权限。若Lite Server实例不存在、已处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StartDevServer(request *model.StartDevServerRequest) (*model.StartDevServerResponse, error) {
	requestDef := GenReqDefForStartDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartDevServerResponse), nil
	}
}

// StartDevServerInvoker 启动Lite Server实例
func (c *ModelArtsClient) StartDevServerInvoker(request *model.StartDevServerRequest) *StartDevServerInvoker {
	requestDef := GenReqDefForStartDevServer()
	return &StartDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartHyperinstance 启动Lite Server超节点服务器
//
// 启动Lite Server超节点服务器接口用于启动已创建但未运行的Lite Server超节点服务器。该接口适用于以下场景：当用户需要开始使用Lite Server超节点服务器进行开发或测试时，可以通过此接口启动指定的超节点服务器。使用该接口的前提条件是Lite Server超节点服务器已创建且处于停止状态，用户具有启动超节点服务器的权限。启动操作完成后，超节点服务器将进入运行状态，用户可以访问和使用服务器提供的服务。若Lite Server超节点服务器不存在、已处于运行状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StartHyperinstance(request *model.StartHyperinstanceRequest) (*model.StartHyperinstanceResponse, error) {
	requestDef := GenReqDefForStartHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartHyperinstanceResponse), nil
	}
}

// StartHyperinstanceInvoker 启动Lite Server超节点服务器
func (c *ModelArtsClient) StartHyperinstanceInvoker(request *model.StartHyperinstanceRequest) *StartHyperinstanceInvoker {
	requestDef := GenReqDefForStartHyperinstance()
	return &StartHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopDevServer 停止Lite Server实例
//
// 停止Lite Server实例接口用于停止正在运行的Lite Server实例。该接口适用于以下场景：当用户需要停止Lite Server实例，以节省资源或进行维护时，可以通过此接口停止指定的Lite Server实例。使用该接口的前提条件是Lite Server实例已创建且处于运行状态，用户具有停止实例的权限。若Lite Server实例不存在、已处于停止状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopDevServer(request *model.StopDevServerRequest) (*model.StopDevServerResponse, error) {
	requestDef := GenReqDefForStopDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopDevServerResponse), nil
	}
}

// StopDevServerInvoker 停止Lite Server实例
func (c *ModelArtsClient) StopDevServerInvoker(request *model.StopDevServerRequest) *StopDevServerInvoker {
	requestDef := GenReqDefForStopDevServer()
	return &StopDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopHyperinstance 停止Lite Server超节点服务器
//
// 停止Lite Server超节点服务器接口用于停止正在运行的Lite Server超节点服务器。该接口适用于以下场景：当用户需要暂停使用Lite Server超节点服务器，以节省资源或进行维护时，可以通过此接口停止指定的超节点服务器。使用该接口的前提条件是Lite Server超节点服务器已创建且处于运行状态或者停止失败状态，用户具有停止超节点服务器的权限。停止操作完成后，超节点服务器将进入停止状态，不再提供服务。若Lite Server超节点服务器不存在、已处于停止状态或用户无权限操作，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopHyperinstance(request *model.StopHyperinstanceRequest) (*model.StopHyperinstanceResponse, error) {
	requestDef := GenReqDefForStopHyperinstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopHyperinstanceResponse), nil
	}
}

// StopHyperinstanceInvoker 停止Lite Server超节点服务器
func (c *ModelArtsClient) StopHyperinstanceInvoker(request *model.StopHyperinstanceRequest) *StopHyperinstanceInvoker {
	requestDef := GenReqDefForStopHyperinstance()
	return &StopHyperinstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SyncDevServers 实时同步用户指定Lite Server实例状态
//
// 实时同步用户Lite Server实例状态接口用于实时获取并同步用户Lite Server实例的当前状态。该接口适用于以下场景：用户需要实时监控其Lite Server实例的运行状态，确保实例正常运行或及时发现并处理异常情况。使用该接口的前提条件是用户已登录并具有相应的权限，且Lite Server实例已创建并处于运行状态。接口调用成功后，将返回Lite Server实例的最新状态信息，包括但不限于实例ID、运行状态、资源使用情况等。若用户无权限操作或Lite Server实例不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) SyncDevServers(request *model.SyncDevServersRequest) (*model.SyncDevServersResponse, error) {
	requestDef := GenReqDefForSyncDevServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SyncDevServersResponse), nil
	}
}

// SyncDevServersInvoker 实时同步用户指定Lite Server实例状态
func (c *ModelArtsClient) SyncDevServersInvoker(request *model.SyncDevServersRequest) *SyncDevServersInvoker {
	requestDef := GenReqDefForSyncDevServers()
	return &SyncDevServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDevServer 修改Lite Server实例名称
//
// 修改DevServer实例名称接口用于更改已创建的DevServer实例的名称。该接口适用于以下场景：当用户需要对DevServer实例进行重命名以更好地反映实例的功能或用途时，或者在实例名称不再符合当前项目命名规范时进行更新。使用该接口的前提条件是DevServer实例已存在且用户具有对该实例的管理权限。修改操作完成后，实例的新名称将立即生效，并在所有相关视图和记录中更新。若DevServer实例不存在、用户无权限操作或新名称不符合命名规则，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateDevServer(request *model.UpdateDevServerRequest) (*model.UpdateDevServerResponse, error) {
	requestDef := GenReqDefForUpdateDevServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDevServerResponse), nil
	}
}

// UpdateDevServerInvoker 修改Lite Server实例名称
func (c *ModelArtsClient) UpdateDevServerInvoker(request *model.UpdateDevServerRequest) *UpdateDevServerInvoker {
	requestDef := GenReqDefForUpdateDevServer()
	return &UpdateDevServerInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateImage 通过运行的实例保存成容器镜像
//
// 通过运行的实例保存成容器镜像接口用于将正在运行的实例保存为容器镜像。该接口适用于以下场景：用户需要保存当前运行环境以便后续使用或开发时，可通过此接口将实例保存为镜像。使用该接口的前提条件是用户已登录系统并具有访问目标实例的权限，同时实例必须处于运行状态。调用该接口后，系统将保存实例的当前状态为容器镜像，包括安装的依赖包和插件。若用户无权限访问指定实例或实例未运行，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateImage(request *model.CreateImageRequest) (*model.CreateImageResponse, error) {
	requestDef := GenReqDefForCreateImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageResponse), nil
	}
}

// CreateImageInvoker 通过运行的实例保存成容器镜像
func (c *ModelArtsClient) CreateImageInvoker(request *model.CreateImageRequest) *CreateImageInvoker {
	requestDef := GenReqDefForCreateImage()
	return &CreateImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotebook 创建Notebook实例
//
// 创建Notebook实例接口用于根据指定的参数创建一个新的Notebook实例。该接口适用于以下场景：用户需要为特定任务或项目创建Notebook实例时，可通过此接口指定实例规格、AI引擎镜像和存储配置。使用该接口的前提条件是用户已登录系统并具有创建Notebook实例的权限，同时需提供有效的创建参数。调用该接口后，系统将异步创建Notebook实例，用户可通过查询接口获取实例状态。创建完成后，用户可通过网页或SSH客户端访问Notebook实例。若用户无权限创建实例或参数无效，接口将返回相应的错误信息。异常情况包括：若系统资源不足，或创建操作失败，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateNotebook(request *model.CreateNotebookRequest) (*model.CreateNotebookResponse, error) {
	requestDef := GenReqDefForCreateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotebookResponse), nil
	}
}

// CreateNotebookInvoker 创建Notebook实例
func (c *ModelArtsClient) CreateNotebookInvoker(request *model.CreateNotebookRequest) *CreateNotebookInvoker {
	requestDef := GenReqDefForCreateNotebook()
	return &CreateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNotebookTags 添加资源标签
//
// 添加资源标签接口用于为指定的Notebook实例添加标签信息。该接口适用于以下场景：用户需要为Notebook实例添加标签信息，可通过此接口添加一个或多个标签。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限。调用该接口后，系统将为指定的Notebook实例添加标签，若标签的key已存在，则覆盖原有的value值。若用户无权限操作指定Notebook实例或输入的参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateNotebookTags(request *model.CreateNotebookTagsRequest) (*model.CreateNotebookTagsResponse, error) {
	requestDef := GenReqDefForCreateNotebookTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNotebookTagsResponse), nil
	}
}

// CreateNotebookTagsInvoker 添加资源标签
func (c *ModelArtsClient) CreateNotebookTagsInvoker(request *model.CreateNotebookTagsRequest) *CreateNotebookTagsInvoker {
	requestDef := GenReqDefForCreateNotebookTags()
	return &CreateNotebookTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotebook 删除Notebook实例
//
// 删除Notebook实例接口用于移除已创建的Notebook实例及其相关资源。该接口适用于以下场景：用户需要清理不再使用的Notebook实例时，可通过此接口删除指定的Notebook实例，包括其容器和所有存储资源。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限。调用该接口后，系统将删除指定的Notebook实例及其相关资源。若用户无权限操作指定实例或Notebook实例未停止，接口将返回相应的错误信息。异常情况包括：若指定的Notebook实例不存在，或删除操作失败，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteNotebook(request *model.DeleteNotebookRequest) (*model.DeleteNotebookResponse, error) {
	requestDef := GenReqDefForDeleteNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotebookResponse), nil
	}
}

// DeleteNotebookInvoker 删除Notebook实例
func (c *ModelArtsClient) DeleteNotebookInvoker(request *model.DeleteNotebookRequest) *DeleteNotebookInvoker {
	requestDef := GenReqDefForDeleteNotebook()
	return &DeleteNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNotebookTags 删除资源标签
//
// 删除资源标签接口用于移除指定Notebook实例的标签信息。该接口适用于以下场景：用户需要清理或重新组织Notebook实例的标签时，可通过此接口删除单个或多个标签。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限。调用该接口后，系统将删除指定的标签，若标签不存在则不进行操作。若用户无权限操作指定Notebook实例或输入的参数无效，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteNotebookTags(request *model.DeleteNotebookTagsRequest) (*model.DeleteNotebookTagsResponse, error) {
	requestDef := GenReqDefForDeleteNotebookTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNotebookTagsResponse), nil
	}
}

// DeleteNotebookTagsInvoker 删除资源标签
func (c *ModelArtsClient) DeleteNotebookTagsInvoker(request *model.DeleteNotebookTagsRequest) *DeleteNotebookTagsInvoker {
	requestDef := GenReqDefForDeleteNotebookTags()
	return &DeleteNotebookTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllNotebooks 查询所有Notebook实例列表
//
// 查询所有Notebook实例列表接口用于获取所有已创建的Notebook实例信息。该接口适用于以下场景：用户需要全面了解当前系统中所有Notebook实例的状态、资源使用情况或管理多个Notebook实例时，可通过此接口获取相关信息。使用该接口的前提条件是用户已创建Notebook实例，并且具有相应的查询权限。调用成功后，系统将返回所有Notebook实例的列表，包含实例ID、状态、创建时间等详细信息。若用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListAllNotebooks(request *model.ListAllNotebooksRequest) (*model.ListAllNotebooksResponse, error) {
	requestDef := GenReqDefForListAllNotebooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllNotebooksResponse), nil
	}
}

// ListAllNotebooksInvoker 查询所有Notebook实例列表
func (c *ModelArtsClient) ListAllNotebooksInvoker(request *model.ListAllNotebooksRequest) *ListAllNotebooksInvoker {
	requestDef := GenReqDefForListAllNotebooks()
	return &ListAllNotebooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthoringClusters 查询用户所有Notebook资源池实例详情
//
// 查询用户所有Notebook资源池实例详情接口用于获取用户关联的所有Notebook资源池实例的详细信息。该接口适用于以下场景：当用户创建Notebook示例需要选择资源池时，可通过此接口获取所有资源池实例列表信息。使用该接口的前提条件是用户已注册并登录系统，且具有查看资源池实例的权限。调用成功后，接口将返回包含所有资源池实例的详细信息列表，包括实例名称、状态、节点规格等。若用户未登录、无权限访问或系统内部出现错误，接口将返回相应的错误信息，如未认证、无权限或服务不可用等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListAuthoringClusters(request *model.ListAuthoringClustersRequest) (*model.ListAuthoringClustersResponse, error) {
	requestDef := GenReqDefForListAuthoringClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthoringClustersResponse), nil
	}
}

// ListAuthoringClustersInvoker 查询用户所有Notebook资源池实例详情
func (c *ModelArtsClient) ListAuthoringClustersInvoker(request *model.ListAuthoringClustersRequest) *ListAuthoringClustersInvoker {
	requestDef := GenReqDefForListAuthoringClusters()
	return &ListAuthoringClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFeatures 查询当前用户指定特性的开关及配额
//
// 查询当前用户指定特性的开关及配额接口用于获取指定特性在当前用户下的开关状态及配额信息。该接口适用于以下场景：当用户需要了解特定特性是否已开启、查看配额限制或监控已使用的资源情况时，可通过此接口查询相关信息。使用该接口的前提条件是用户已登录且具有查询权限，同时指定的特性必须存在。调用该接口后，系统将返回该特性是否已开启、配额总量及已使用的资源情况等详细信息。若用户无权限查询、特性不存在或系统出现异常，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListFeatures(request *model.ListFeaturesRequest) (*model.ListFeaturesResponse, error) {
	requestDef := GenReqDefForListFeatures()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFeaturesResponse), nil
	}
}

// ListFeaturesInvoker 查询当前用户指定特性的开关及配额
func (c *ModelArtsClient) ListFeaturesInvoker(request *model.ListFeaturesRequest) *ListFeaturesInvoker {
	requestDef := GenReqDefForListFeatures()
	return &ListFeaturesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlavors 查询Notebook支持的有效规格列表
//
// 查询Notebook支持的有效规格列表接口用于获取运行Notebook实例时可使用的规格选项。该接口适用于以下场景：用户需要了解Notebook实例支持的配置选项时，可通过此接口查询可用的规格列表。使用该接口的前提条件是用户已登录系统并具有访问目标Notebook实例的权限。调用该接口后，系统将返回Notebook实例支持的有效规格列表，包括内存、CPU等配置信息。若用户无权限访问指定实例或Notebook实例未运行，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListFlavors(request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	requestDef := GenReqDefForListFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlavorsResponse), nil
	}
}

// ListFlavorsInvoker 查询Notebook支持的有效规格列表
func (c *ModelArtsClient) ListFlavorsInvoker(request *model.ListFlavorsRequest) *ListFlavorsInvoker {
	requestDef := GenReqDefForListFlavors()
	return &ListFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNotebooks 查询Notebook实例列表
//
// 查询Notebook实例列表接口用于获取满足特定条件的Notebook实例信息。该接口适用于以下场景：用户管理多个Notebook实例或查看特定状态的Notebook实例时，可通过此接口获取相关信息。使用该接口的前提条件是用户已创建Notebook实例，并且具有相应的查询权限。调用成功后，系统将返回符合条件的Notebook实例列表，包含实例ID、状态、创建时间等详细信息。若用户无权限访问或查询条件不明确，接口将返回相应的错误信息或空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListNotebooks(request *model.ListNotebooksRequest) (*model.ListNotebooksResponse, error) {
	requestDef := GenReqDefForListNotebooks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNotebooksResponse), nil
	}
}

// ListNotebooksInvoker 查询Notebook实例列表
func (c *ModelArtsClient) ListNotebooksInvoker(request *model.ListNotebooksRequest) *ListNotebooksInvoker {
	requestDef := GenReqDefForListNotebooks()
	return &ListNotebooksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenewLease Notebook时长续约
//
// Notebook时长续约接口用于延长运行中的Notebook实例的运行时间。该接口适用于以下场景：用户需要延长Notebook实例的使用时间以完成长时间任务时，可通过此接口延长指定实例的运行时间。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限，同时Notebook实例必须处于运行状态。调用该接口后，系统将延长指定Notebook实例的运行时间，用户可继续使用。若用户无权限操作指定实例或Notebook实例未运行，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) RenewLease(request *model.RenewLeaseRequest) (*model.RenewLeaseResponse, error) {
	requestDef := GenReqDefForRenewLease()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewLeaseResponse), nil
	}
}

// RenewLeaseInvoker Notebook时长续约
func (c *ModelArtsClient) RenewLeaseInvoker(request *model.RenewLeaseRequest) *RenewLeaseInvoker {
	requestDef := GenReqDefForRenewLease()
	return &RenewLeaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCluster 查询Notebook资源池详情
//
// 查询Notebook资源池详情接口用于获取资源池的详细信息。该接口适用于以下场景：当用户需要创建Notebook实例作业时，可通过此接口查询指定集群的详细信息。使用该接口的前提条件是集群已成功纳管且用户具有相应的访问权限。调用该接口后，系统将返回集群的实例ID、名称、Flavor规格、实例状态以及实例可打开的URL等详细信息。若集群不存在、未被纳管或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowCluster(request *model.ShowClusterRequest) (*model.ShowClusterResponse, error) {
	requestDef := GenReqDefForShowCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterResponse), nil
	}
}

// ShowClusterInvoker 查询Notebook资源池详情
func (c *ModelArtsClient) ShowClusterInvoker(request *model.ShowClusterRequest) *ShowClusterInvoker {
	requestDef := GenReqDefForShowCluster()
	return &ShowClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLease 查询运行中的Notebook可用时长
//
// 查询运行中的Notebook可用时长接口用于获取正在运行的Notebook实例的剩余可用时间。该接口适用于以下场景：用户需要了解Notebook实例的剩余运行时间以合理安排任务时，可通过此接口查询指定实例的可用时长。使用该接口的前提条件是用户已登录系统并具有访问目标Notebook实例的权限，同时Notebook实例必须处于运行状态。调用该接口后，系统将返回指定Notebook实例的可用时长信息。若用户无权限访问指定实例或Notebook实例未运行，接口将返回相应的错误信息。异常情况包括：若指定的Notebook实例不存在，或查询操作失败，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowLease(request *model.ShowLeaseRequest) (*model.ShowLeaseResponse, error) {
	requestDef := GenReqDefForShowLease()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLeaseResponse), nil
	}
}

// ShowLeaseInvoker 查询运行中的Notebook可用时长
func (c *ModelArtsClient) ShowLeaseInvoker(request *model.ShowLeaseRequest) *ShowLeaseInvoker {
	requestDef := GenReqDefForShowLease()
	return &ShowLeaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebook 查询Notebook实例详情
//
// 查询Notebook实例详情接口用于获取指定Notebook实例的详细信息。该接口适用于以下场景：用户需要查看特定Notebook实例的详细配置、运行状态或获取访问链接时，可通过此接口获取相关信息。使用该接口的前提条件是Notebook实例已存在且用户具有相应的查询权限。调用成功后，系统将返回实例ID、名称、规格、镜像、实例状态和实例可打开的URL等详细信息。若实例不存在或用户无权限访问，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNotebook(request *model.ShowNotebookRequest) (*model.ShowNotebookResponse, error) {
	requestDef := GenReqDefForShowNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookResponse), nil
	}
}

// ShowNotebookInvoker 查询Notebook实例详情
func (c *ModelArtsClient) ShowNotebookInvoker(request *model.ShowNotebookRequest) *ShowNotebookInvoker {
	requestDef := GenReqDefForShowNotebook()
	return &ShowNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNotebookTags 查询Notebook资源类型下的标签
//
// 查询Notebook资源类型下的标签接口用于获取用户当前project下Notebook实例的标签信息。该接口适用于以下场景：用户需要管理或统计Notebook资源时，可通过此接口查询特定标签或所有标签的Notebook实例。使用该接口的前提条件是用户已登录系统并具有访问权限，同时可指定工作空间或默认查询所有工作空间。调用该接口后，系统将返回指定Notebook实例的标签列表，包括标签名称、标签值等信息。若用户无权限，则返回相应的异常信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowNotebookTags(request *model.ShowNotebookTagsRequest) (*model.ShowNotebookTagsResponse, error) {
	requestDef := GenReqDefForShowNotebookTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNotebookTagsResponse), nil
	}
}

// ShowNotebookTagsInvoker 查询Notebook资源类型下的标签
func (c *ModelArtsClient) ShowNotebookTagsInvoker(request *model.ShowNotebookTagsRequest) *ShowNotebookTagsInvoker {
	requestDef := GenReqDefForShowNotebookTags()
	return &ShowNotebookTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSwitchableFlavors 查询Notebook支持的可切换规格列表
//
// 查询Notebook支持的可切换规格列表接口用于获取创建Notebook实例时可选择的规格选项。该接口适用于以下场景：用户需要了解Notebook实例支持的配置选项时，可通过此接口查询可用的规格列表。使用该接口的前提条件是用户已登录系统并具有创建Notebook实例的权限。调用该接口后，系统将返回Notebook实例支持的可切换规格列表，包括内存、CPU等配置信息。若用户无权限创建Notebook实例，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowSwitchableFlavors(request *model.ShowSwitchableFlavorsRequest) (*model.ShowSwitchableFlavorsResponse, error) {
	requestDef := GenReqDefForShowSwitchableFlavors()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSwitchableFlavorsResponse), nil
	}
}

// ShowSwitchableFlavorsInvoker 查询Notebook支持的可切换规格列表
func (c *ModelArtsClient) ShowSwitchableFlavorsInvoker(request *model.ShowSwitchableFlavorsRequest) *ShowSwitchableFlavorsInvoker {
	requestDef := GenReqDefForShowSwitchableFlavors()
	return &ShowSwitchableFlavorsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StartNotebook 启动Notebook实例
//
// 启动Notebook实例接口用于启动已创建的Notebook实例。该接口适用于以下场景：用户需要开始运行Notebook实例以进行数据处理、模型训练或开发时，可通过此接口启动指定的Notebook实例。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限，同时Notebook实例必须处于停止状态且配置正确。调用该接口后，系统将启动指定的Notebook实例，用户可开始使用。若用户无权限操作指定实例或Notebook实例未停止，接口将返回相应的错误信息。异常情况包括：若指定的Notebook实例不存在，或启动操作失败，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StartNotebook(request *model.StartNotebookRequest) (*model.StartNotebookResponse, error) {
	requestDef := GenReqDefForStartNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StartNotebookResponse), nil
	}
}

// StartNotebookInvoker 启动Notebook实例
func (c *ModelArtsClient) StartNotebookInvoker(request *model.StartNotebookRequest) *StartNotebookInvoker {
	requestDef := GenReqDefForStartNotebook()
	return &StartNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopNotebook 停止Notebook实例
//
// 停止Notebook实例接口用于停止正在运行的Notebook实例。该接口适用于以下场景：用户需要释放Notebook实例占用的资源或结束当前运行的任务时，可通过此接口停止指定的Notebook实例。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限，同时Notebook实例必须处于运行状态。调用该接口后，系统将停止指定的Notebook实例，释放相关资源。若用户无权限操作指定实例或Notebook实例未运行，接口将返回相应的错误信息。异常情况包括：若指定的Notebook实例不存在，或停止操作失败，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) StopNotebook(request *model.StopNotebookRequest) (*model.StopNotebookResponse, error) {
	requestDef := GenReqDefForStopNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopNotebookResponse), nil
	}
}

// StopNotebookInvoker 停止Notebook实例
func (c *ModelArtsClient) StopNotebookInvoker(request *model.StopNotebookRequest) *StopNotebookInvoker {
	requestDef := GenReqDefForStopNotebook()
	return &StopNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNotebook 更新Notebook实例
//
// 更新Notebook实例接口用于修改Notebook实例的配置信息，包括名称、描述、规格和镜像等。该接口适用于以下场景：用户需要调整Notebook实例的配置以适应新的需求时，可通过此接口更新实例的详细信息。使用该接口的前提条件是用户已登录系统并具有操作目标Notebook实例的权限，同时Notebook实例必须处于停止状态。调用该接口后，系统将更新指定Notebook实例的配置信息。若用户无权限操作指定实例或Notebook实例未停止，接口将返回相应的错误信息。异常情况包括：若指定的Notebook实例不存在，或更新参数无效，接口将返回相应的错误提示。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateNotebook(request *model.UpdateNotebookRequest) (*model.UpdateNotebookResponse, error) {
	requestDef := GenReqDefForUpdateNotebook()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNotebookResponse), nil
	}
}

// UpdateNotebookInvoker 更新Notebook实例
func (c *ModelArtsClient) UpdateNotebookInvoker(request *model.UpdateNotebookRequest) *UpdateNotebookInvoker {
	requestDef := GenReqDefForUpdateNotebook()
	return &UpdateNotebookInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflow 新建Workflow工作流
//
// 创建Workflow工作流。[可参考[如何开发Workflow](https://support.huaweicloud.com/usermanual-standard-modelarts/modelarts_workflow_0292.html)，创建工作流。](tag:hc)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflow(request *model.CreateWorkflowRequest) (*model.CreateWorkflowResponse, error) {
	requestDef := GenReqDefForCreateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowResponse), nil
	}
}

// CreateWorkflowInvoker 新建Workflow工作流
func (c *ModelArtsClient) CreateWorkflowInvoker(request *model.CreateWorkflowRequest) *CreateWorkflowInvoker {
	requestDef := GenReqDefForCreateWorkflow()
	return &CreateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowPurchasePool 创建在线服务包
//
// 计费工作流购买资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowPurchasePool(request *model.CreateWorkflowPurchasePoolRequest) (*model.CreateWorkflowPurchasePoolResponse, error) {
	requestDef := GenReqDefForCreateWorkflowPurchasePool()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowPurchasePoolResponse), nil
	}
}

// CreateWorkflowPurchasePoolInvoker 创建在线服务包
func (c *ModelArtsClient) CreateWorkflowPurchasePoolInvoker(request *model.CreateWorkflowPurchasePoolRequest) *CreateWorkflowPurchasePoolInvoker {
	requestDef := GenReqDefForCreateWorkflowPurchasePool()
	return &CreateWorkflowPurchasePoolInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowServiceAuth 在线服务鉴权
//
// 计费工作流在线服务鉴权。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowServiceAuth(request *model.CreateWorkflowServiceAuthRequest) (*model.CreateWorkflowServiceAuthResponse, error) {
	requestDef := GenReqDefForCreateWorkflowServiceAuth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowServiceAuthResponse), nil
	}
}

// CreateWorkflowServiceAuthInvoker 在线服务鉴权
func (c *ModelArtsClient) CreateWorkflowServiceAuthInvoker(request *model.CreateWorkflowServiceAuthRequest) *CreateWorkflowServiceAuthInvoker {
	requestDef := GenReqDefForCreateWorkflowServiceAuth()
	return &CreateWorkflowServiceAuthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflow 删除Workflow工作流
//
// 通过ID删除Workflow工作流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteWorkflow(request *model.DeleteWorkflowRequest) (*model.DeleteWorkflowResponse, error) {
	requestDef := GenReqDefForDeleteWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowResponse), nil
	}
}

// DeleteWorkflowInvoker 删除Workflow工作流
func (c *ModelArtsClient) DeleteWorkflowInvoker(request *model.DeleteWorkflowRequest) *DeleteWorkflowInvoker {
	requestDef := GenReqDefForDeleteWorkflow()
	return &DeleteWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflows 获取Workflow工作流列表
//
// 展示Workflow工作流列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListWorkflows(request *model.ListWorkflowsRequest) (*model.ListWorkflowsResponse, error) {
	requestDef := GenReqDefForListWorkflows()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowsResponse), nil
	}
}

// ListWorkflowsInvoker 获取Workflow工作流列表
func (c *ModelArtsClient) ListWorkflowsInvoker(request *model.ListWorkflowsRequest) *ListWorkflowsInvoker {
	requestDef := GenReqDefForListWorkflows()
	return &ListWorkflowsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflow 查询Workflow工作流
//
// 通过ID查询Workflow工作流详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflow(request *model.ShowWorkflowRequest) (*model.ShowWorkflowResponse, error) {
	requestDef := GenReqDefForShowWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowResponse), nil
	}
}

// ShowWorkflowInvoker 查询Workflow工作流
func (c *ModelArtsClient) ShowWorkflowInvoker(request *model.ShowWorkflowRequest) *ShowWorkflowInvoker {
	requestDef := GenReqDefForShowWorkflow()
	return &ShowWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowLabels Workflow列表所有标签
//
// Workflow列表所有标签接口用于获取指定项目下所有工作流的标签信息。
// 该接口适用于以下场景：当用户需要了解项目中所有工作流的标签配置，以便进行资源管理和筛选时，可以通过此接口获取标签列表。使用该接口的前提条件是用户已登录并具有查看工作流标签的权限。响应消息体中包含每个工作流的标签信息，如标签键和值。若用户无权限或项目下无工作流，接口将返回相应的错误信息或空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowLabels(request *model.ShowWorkflowLabelsRequest) (*model.ShowWorkflowLabelsResponse, error) {
	requestDef := GenReqDefForShowWorkflowLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowLabelsResponse), nil
	}
}

// ShowWorkflowLabelsInvoker Workflow列表所有标签
func (c *ModelArtsClient) ShowWorkflowLabelsInvoker(request *model.ShowWorkflowLabelsRequest) *ShowWorkflowLabelsInvoker {
	requestDef := GenReqDefForShowWorkflowLabels()
	return &ShowWorkflowLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowsOverview 总览Workflow工作流
//
// 获取Workflow工作流统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowsOverview(request *model.ShowWorkflowsOverviewRequest) (*model.ShowWorkflowsOverviewResponse, error) {
	requestDef := GenReqDefForShowWorkflowsOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowsOverviewResponse), nil
	}
}

// ShowWorkflowsOverviewInvoker 总览Workflow工作流
func (c *ModelArtsClient) ShowWorkflowsOverviewInvoker(request *model.ShowWorkflowsOverviewRequest) *ShowWorkflowsOverviewInvoker {
	requestDef := GenReqDefForShowWorkflowsOverview()
	return &ShowWorkflowsOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowsTodolist 查询Workflow待办事项
//
// 获取Workflow待办列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowsTodolist(request *model.ShowWorkflowsTodolistRequest) (*model.ShowWorkflowsTodolistResponse, error) {
	requestDef := GenReqDefForShowWorkflowsTodolist()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowsTodolistResponse), nil
	}
}

// ShowWorkflowsTodolistInvoker 查询Workflow待办事项
func (c *ModelArtsClient) ShowWorkflowsTodolistInvoker(request *model.ShowWorkflowsTodolistRequest) *ShowWorkflowsTodolistInvoker {
	requestDef := GenReqDefForShowWorkflowsTodolist()
	return &ShowWorkflowsTodolistInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflow 修改Workflow工作流
//
// 更新Workflow工作流信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkflow(request *model.UpdateWorkflowRequest) (*model.UpdateWorkflowResponse, error) {
	requestDef := GenReqDefForUpdateWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowResponse), nil
	}
}

// UpdateWorkflowInvoker 修改Workflow工作流
func (c *ModelArtsClient) UpdateWorkflowInvoker(request *model.UpdateWorkflowRequest) *UpdateWorkflowInvoker {
	requestDef := GenReqDefForUpdateWorkflow()
	return &UpdateWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowExecution 新建Workflow Execution
//
// 创建Workflow Execution。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowExecution(request *model.CreateWorkflowExecutionRequest) (*model.CreateWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForCreateWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowExecutionResponse), nil
	}
}

// CreateWorkflowExecutionInvoker 新建Workflow Execution
func (c *ModelArtsClient) CreateWorkflowExecutionInvoker(request *model.CreateWorkflowExecutionRequest) *CreateWorkflowExecutionInvoker {
	requestDef := GenReqDefForCreateWorkflowExecution()
	return &CreateWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowExecutionsActions 管理Workflow Execution
//
// 本接口支持对Workflow Execution进行停止或重跑操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowExecutionsActions(request *model.CreateWorkflowExecutionsActionsRequest) (*model.CreateWorkflowExecutionsActionsResponse, error) {
	requestDef := GenReqDefForCreateWorkflowExecutionsActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowExecutionsActionsResponse), nil
	}
}

// CreateWorkflowExecutionsActionsInvoker 管理Workflow Execution
func (c *ModelArtsClient) CreateWorkflowExecutionsActionsInvoker(request *model.CreateWorkflowExecutionsActionsRequest) *CreateWorkflowExecutionsActionsInvoker {
	requestDef := GenReqDefForCreateWorkflowExecutionsActions()
	return &CreateWorkflowExecutionsActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowStepExecutionsActions 管理Workflow StepExecution
//
// 本接口支持对Workflow StepExecution进行重试、停止和继续操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowStepExecutionsActions(request *model.CreateWorkflowStepExecutionsActionsRequest) (*model.CreateWorkflowStepExecutionsActionsResponse, error) {
	requestDef := GenReqDefForCreateWorkflowStepExecutionsActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowStepExecutionsActionsResponse), nil
	}
}

// CreateWorkflowStepExecutionsActionsInvoker 管理Workflow StepExecution
func (c *ModelArtsClient) CreateWorkflowStepExecutionsActionsInvoker(request *model.CreateWorkflowStepExecutionsActionsRequest) *CreateWorkflowStepExecutionsActionsInvoker {
	requestDef := GenReqDefForCreateWorkflowStepExecutionsActions()
	return &CreateWorkflowStepExecutionsActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflowExecution 删除Workflow Execution
//
// 通过ID删除Workflow Execution。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteWorkflowExecution(request *model.DeleteWorkflowExecutionRequest) (*model.DeleteWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForDeleteWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowExecutionResponse), nil
	}
}

// DeleteWorkflowExecutionInvoker 删除Workflow Execution
func (c *ModelArtsClient) DeleteWorkflowExecutionInvoker(request *model.DeleteWorkflowExecutionRequest) *DeleteWorkflowExecutionInvoker {
	requestDef := GenReqDefForDeleteWorkflowExecution()
	return &DeleteWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListExecutionLabels 获取Workflow Execution列表的所有标签
//
// 获取Workflow Execution列表的所有标签接口用于查询指定工作流执行记录中的所有标签。
// 该接口适用于以下场景：当用户需要查看工作流执行记录的标签信息，以便进行分类、筛选或统计时，可以通过此接口获取所有标签的列表。使用该接口的前提条件是用户已登录且具有查看工作流执行记录的权限。接口响应消息体中包含每个标签的详细信息，如标签键和标签值。若用户无权限操作或指定的工作流执行记录不存在，接口将返回相应的错误信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListExecutionLabels(request *model.ListExecutionLabelsRequest) (*model.ListExecutionLabelsResponse, error) {
	requestDef := GenReqDefForListExecutionLabels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListExecutionLabelsResponse), nil
	}
}

// ListExecutionLabelsInvoker 获取Workflow Execution列表的所有标签
func (c *ModelArtsClient) ListExecutionLabelsInvoker(request *model.ListExecutionLabelsRequest) *ListExecutionLabelsInvoker {
	requestDef := GenReqDefForListExecutionLabels()
	return &ListExecutionLabelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowExecutions 获取Execution列表
//
// 查询Workflow下的执行记录列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListWorkflowExecutions(request *model.ListWorkflowExecutionsRequest) (*model.ListWorkflowExecutionsResponse, error) {
	requestDef := GenReqDefForListWorkflowExecutions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowExecutionsResponse), nil
	}
}

// ListWorkflowExecutionsInvoker 获取Execution列表
func (c *ModelArtsClient) ListWorkflowExecutionsInvoker(request *model.ListWorkflowExecutionsRequest) *ListWorkflowExecutionsInvoker {
	requestDef := GenReqDefForListWorkflowExecutions()
	return &ListWorkflowExecutionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkflowStepExecution 获取StepExecution列表
//
// 查询指定工作流中各步骤的执行情况。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ListWorkflowStepExecution(request *model.ListWorkflowStepExecutionRequest) (*model.ListWorkflowStepExecutionResponse, error) {
	requestDef := GenReqDefForListWorkflowStepExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkflowStepExecutionResponse), nil
	}
}

// ListWorkflowStepExecutionInvoker 获取StepExecution列表
func (c *ModelArtsClient) ListWorkflowStepExecutionInvoker(request *model.ListWorkflowStepExecutionRequest) *ListWorkflowStepExecutionInvoker {
	requestDef := GenReqDefForListWorkflowStepExecution()
	return &ListWorkflowStepExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowExecution 查询Workflow Execution
//
// 通过ID查询Workflow Execution详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowExecution(request *model.ShowWorkflowExecutionRequest) (*model.ShowWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForShowWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowExecutionResponse), nil
	}
}

// ShowWorkflowExecutionInvoker 查询Workflow Execution
func (c *ModelArtsClient) ShowWorkflowExecutionInvoker(request *model.ShowWorkflowExecutionRequest) *ShowWorkflowExecutionInvoker {
	requestDef := GenReqDefForShowWorkflowExecution()
	return &ShowWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowStepExecutionMetrics 获取Workflow工作流节点度量信息
//
// 获取Workflow工作流节点的度量信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowStepExecutionMetrics(request *model.ShowWorkflowStepExecutionMetricsRequest) (*model.ShowWorkflowStepExecutionMetricsResponse, error) {
	requestDef := GenReqDefForShowWorkflowStepExecutionMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowStepExecutionMetricsResponse), nil
	}
}

// ShowWorkflowStepExecutionMetricsInvoker 获取Workflow工作流节点度量信息
func (c *ModelArtsClient) ShowWorkflowStepExecutionMetricsInvoker(request *model.ShowWorkflowStepExecutionMetricsRequest) *ShowWorkflowStepExecutionMetricsInvoker {
	requestDef := GenReqDefForShowWorkflowStepExecutionMetrics()
	return &ShowWorkflowStepExecutionMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflowExecution 更新Workflow Execution
//
// 通过ID更新Workflow Exectuion。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkflowExecution(request *model.UpdateWorkflowExecutionRequest) (*model.UpdateWorkflowExecutionResponse, error) {
	requestDef := GenReqDefForUpdateWorkflowExecution()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowExecutionResponse), nil
	}
}

// UpdateWorkflowExecutionInvoker 更新Workflow Execution
func (c *ModelArtsClient) UpdateWorkflowExecutionInvoker(request *model.UpdateWorkflowExecutionRequest) *UpdateWorkflowExecutionInvoker {
	requestDef := GenReqDefForUpdateWorkflowExecution()
	return &UpdateWorkflowExecutionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowSchedule 创建工作流定时调度
//
// 创建Workflow定时调度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowSchedule(request *model.CreateWorkflowScheduleRequest) (*model.CreateWorkflowScheduleResponse, error) {
	requestDef := GenReqDefForCreateWorkflowSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowScheduleResponse), nil
	}
}

// CreateWorkflowScheduleInvoker 创建工作流定时调度
func (c *ModelArtsClient) CreateWorkflowScheduleInvoker(request *model.CreateWorkflowScheduleRequest) *CreateWorkflowScheduleInvoker {
	requestDef := GenReqDefForCreateWorkflowSchedule()
	return &CreateWorkflowScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflowScheduleId 删除工作流定时调度信息
//
// 删除工作流调度信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteWorkflowScheduleId(request *model.DeleteWorkflowScheduleIdRequest) (*model.DeleteWorkflowScheduleIdResponse, error) {
	requestDef := GenReqDefForDeleteWorkflowScheduleId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowScheduleIdResponse), nil
	}
}

// DeleteWorkflowScheduleIdInvoker 删除工作流定时调度信息
func (c *ModelArtsClient) DeleteWorkflowScheduleIdInvoker(request *model.DeleteWorkflowScheduleIdRequest) *DeleteWorkflowScheduleIdInvoker {
	requestDef := GenReqDefForDeleteWorkflowScheduleId()
	return &DeleteWorkflowScheduleIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowSchedule 查询工作流定时调度详情
//
// 查询工作流调度详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowSchedule(request *model.ShowWorkflowScheduleRequest) (*model.ShowWorkflowScheduleResponse, error) {
	requestDef := GenReqDefForShowWorkflowSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowScheduleResponse), nil
	}
}

// ShowWorkflowScheduleInvoker 查询工作流定时调度详情
func (c *ModelArtsClient) ShowWorkflowScheduleInvoker(request *model.ShowWorkflowScheduleRequest) *ShowWorkflowScheduleInvoker {
	requestDef := GenReqDefForShowWorkflowSchedule()
	return &ShowWorkflowScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowScheduleList 查询工作流定时调度列表
//
// 查询工作流定时调度列表接口用于获取指定项目下所有工作流的定时调度信息。
// 该接口适用于以下场景：当用户需要查看项目中所有工作流的定时调度配置，以便进行任务管理和调度优化时，可以通过此接口获取定时调度列表。使用该接口的前提条件是用户已登录并具有查看工作流定时调度的权限。响应消息体中包含每个工作流的定时调度信息，如调度ID、调度时间、状态等。若用户无权限或项目下无工作流定时调度，接口将返回相应的错误信息或空列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowScheduleList(request *model.ShowWorkflowScheduleListRequest) (*model.ShowWorkflowScheduleListResponse, error) {
	requestDef := GenReqDefForShowWorkflowScheduleList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowScheduleListResponse), nil
	}
}

// ShowWorkflowScheduleListInvoker 查询工作流定时调度列表
func (c *ModelArtsClient) ShowWorkflowScheduleListInvoker(request *model.ShowWorkflowScheduleListRequest) *ShowWorkflowScheduleListInvoker {
	requestDef := GenReqDefForShowWorkflowScheduleList()
	return &ShowWorkflowScheduleListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflowSchedule 更新工作流定时调度信息
//
// 更新WorkflowSchedule信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkflowSchedule(request *model.UpdateWorkflowScheduleRequest) (*model.UpdateWorkflowScheduleResponse, error) {
	requestDef := GenReqDefForUpdateWorkflowSchedule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowScheduleResponse), nil
	}
}

// UpdateWorkflowScheduleInvoker 更新工作流定时调度信息
func (c *ModelArtsClient) UpdateWorkflowScheduleInvoker(request *model.UpdateWorkflowScheduleRequest) *UpdateWorkflowScheduleInvoker {
	requestDef := GenReqDefForUpdateWorkflowSchedule()
	return &UpdateWorkflowScheduleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkflowSubscriptions 新建消息订阅Subscription
//
// 为Workflow工作流添加消息订阅功能。工作流已订阅的事件发生时，会产生消息提醒。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) CreateWorkflowSubscriptions(request *model.CreateWorkflowSubscriptionsRequest) (*model.CreateWorkflowSubscriptionsResponse, error) {
	requestDef := GenReqDefForCreateWorkflowSubscriptions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkflowSubscriptionsResponse), nil
	}
}

// CreateWorkflowSubscriptionsInvoker 新建消息订阅Subscription
func (c *ModelArtsClient) CreateWorkflowSubscriptionsInvoker(request *model.CreateWorkflowSubscriptionsRequest) *CreateWorkflowSubscriptionsInvoker {
	requestDef := GenReqDefForCreateWorkflowSubscriptions()
	return &CreateWorkflowSubscriptionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkflowSubscription 删除消息订阅Subscription
//
// 删除已订阅的消息订阅Subscription。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) DeleteWorkflowSubscription(request *model.DeleteWorkflowSubscriptionRequest) (*model.DeleteWorkflowSubscriptionResponse, error) {
	requestDef := GenReqDefForDeleteWorkflowSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkflowSubscriptionResponse), nil
	}
}

// DeleteWorkflowSubscriptionInvoker 删除消息订阅Subscription
func (c *ModelArtsClient) DeleteWorkflowSubscriptionInvoker(request *model.DeleteWorkflowSubscriptionRequest) *DeleteWorkflowSubscriptionInvoker {
	requestDef := GenReqDefForDeleteWorkflowSubscription()
	return &DeleteWorkflowSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkflowSubscription 查询消息订阅Subscription详情
//
// 查询Workflow工作流已订阅的订阅信息详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) ShowWorkflowSubscription(request *model.ShowWorkflowSubscriptionRequest) (*model.ShowWorkflowSubscriptionResponse, error) {
	requestDef := GenReqDefForShowWorkflowSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkflowSubscriptionResponse), nil
	}
}

// ShowWorkflowSubscriptionInvoker 查询消息订阅Subscription详情
func (c *ModelArtsClient) ShowWorkflowSubscriptionInvoker(request *model.ShowWorkflowSubscriptionRequest) *ShowWorkflowSubscriptionInvoker {
	requestDef := GenReqDefForShowWorkflowSubscription()
	return &ShowWorkflowSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkflowSubscription 更新消息订阅Subscription
//
// 更新Workflow工作流已订阅的订阅信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *ModelArtsClient) UpdateWorkflowSubscription(request *model.UpdateWorkflowSubscriptionRequest) (*model.UpdateWorkflowSubscriptionResponse, error) {
	requestDef := GenReqDefForUpdateWorkflowSubscription()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkflowSubscriptionResponse), nil
	}
}

// UpdateWorkflowSubscriptionInvoker 更新消息订阅Subscription
func (c *ModelArtsClient) UpdateWorkflowSubscriptionInvoker(request *model.UpdateWorkflowSubscriptionRequest) *UpdateWorkflowSubscriptionInvoker {
	requestDef := GenReqDefForUpdateWorkflowSubscription()
	return &UpdateWorkflowSubscriptionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
