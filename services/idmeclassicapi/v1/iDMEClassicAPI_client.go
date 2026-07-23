package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idmeclassicapi/v1/model"
)

type IDMEClassicAPIClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIDMEClassicAPIClient(hcClient *httpclient.HcHttpClient) *IDMEClassicAPIClient {
	return &IDMEClassicAPIClient{HcClient: hcClient}
}

func IDMEClassicAPIClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddTag 绑定标签
//
// 本接口用于为指定数据模型的数据实例绑定预定义的标签。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“标签管理”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
// 3. 已在应用运行态的“基础数据管理 &gt; 标签”中创建目标标签。具体操作请参见[标签](https://support.huaweicloud.com/usermanual-idme/idme_clientog_0096.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) AddTag(request *model.AddTagRequest) (*model.AddTagResponse, error) {
	requestDef := GenReqDefForAddTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTagResponse), nil
	}
}

// AddTagInvoker 绑定标签
func (c *IDMEClassicAPIClient) AddTagInvoker(request *model.AddTagRequest) *AddTagInvoker {
	requestDef := GenReqDefForAddTag()
	return &AddTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddToCategory 添加数据分类
//
// 本接口用于将数据实例添加到指定分类。当需要对零部件、设备、物料等数据实例进行归类管理（如按功能模块、产品线、工艺类型等维度分类）时，可调用本接口建立实例与分类的关联关系。
// - 当实例不存在时，系统将抛出异常。
// - 当实例与指定分类已存在关联关系时，系统将不重复添加，返回值为0。
// - 当实例与指定分类首次建立关联时，返回值为1。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) AddToCategory(request *model.AddToCategoryRequest) (*model.AddToCategoryResponse, error) {
	requestDef := GenReqDefForAddToCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddToCategoryResponse), nil
	}
}

// AddToCategoryInvoker 添加数据分类
func (c *IDMEClassicAPIClient) AddToCategoryInvoker(request *model.AddToCategoryRequest) *AddToCategoryInvoker {
	requestDef := GenReqDefForAddToCategory()
	return &AddToCategoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAddChildNode 批量添加实例的子节点
//
// 本接口用于批量为指定数据实例添加子节点，建立父子层级关联关系。当BOM（物料清单）需要新增子装配节点、组织架构需要新增下属团队或产品分类需要新增下级类目时，可调用本接口批量完成节点挂载。
// - 调用本接口时，需在parentId中指定父节点实例ID，在childList中传入待添加为子节点的实例ID列表。
// - 添加操作执行后，子节点将发生以下变更：
//   - 父节点（parentNode）字段被设置为指定的父节点。
//   - 根节点（rootNode）字段被设置为父节点所在树的根节点。
// - 全路径（fullPath）和原始全路径（rawFullPath）字段被更新为包含父节点路径的完整路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchAddChildNode(request *model.BatchAddChildNodeRequest) (*model.BatchAddChildNodeResponse, error) {
	requestDef := GenReqDefForBatchAddChildNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddChildNodeResponse), nil
	}
}

// BatchAddChildNodeInvoker 批量添加实例的子节点
func (c *IDMEClassicAPIClient) BatchAddChildNodeInvoker(request *model.BatchAddChildNodeRequest) *BatchAddChildNodeInvoker {
	requestDef := GenReqDefForBatchAddChildNode()
	return &BatchAddChildNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckin 批量检入M-V模型数据实例
//
// 本接口用于根据主对象ID批量检入M-V模型数据实例。已检入的数据实例会生成一个新的迭代版本，并将数据存储至系统中。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckin(request *model.BatchCheckinRequest) (*model.BatchCheckinResponse, error) {
	requestDef := GenReqDefForBatchCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckinResponse), nil
	}
}

// BatchCheckinInvoker 批量检入M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchCheckinInvoker(request *model.BatchCheckinRequest) *BatchCheckinInvoker {
	requestDef := GenReqDefForBatchCheckin()
	return &BatchCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckout 批量检出M-V模型数据实例
//
// 本接口用于根据主对象ID（masterId）列表，对指定的多个Master-Branch-Version（M-V）模型实例批量执行检出（Check-out）操作。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckout(request *model.BatchCheckoutRequest) (*model.BatchCheckoutResponse, error) {
	requestDef := GenReqDefForBatchCheckout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutResponse), nil
	}
}

// BatchCheckoutInvoker 批量检出M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchCheckoutInvoker(request *model.BatchCheckoutRequest) *BatchCheckoutInvoker {
	requestDef := GenReqDefForBatchCheckout()
	return &BatchCheckoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutAndUpdate 批量检出并更新M-V模型
//
// 本接口用于根据主对象ID（masterId）列表，对指定的多个Master-Branch-Version（M-V）模型实例批量执行检出（Check-out）并同步更新（Update）的原子操作。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutAndUpdate(request *model.BatchCheckoutAndUpdateRequest) (*model.BatchCheckoutAndUpdateResponse, error) {
	requestDef := GenReqDefForBatchCheckoutAndUpdate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutAndUpdateResponse), nil
	}
}

// BatchCheckoutAndUpdateInvoker 批量检出并更新M-V模型
func (c *IDMEClassicAPIClient) BatchCheckoutAndUpdateInvoker(request *model.BatchCheckoutAndUpdateRequest) *BatchCheckoutAndUpdateInvoker {
	requestDef := GenReqDefForBatchCheckoutAndUpdate()
	return &BatchCheckoutAndUpdateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutUndo 批量撤销检出M-V模型数据实例
//
// 本接口用于批量撤销指定M-V模型实例的检出，将实例数据批量还原至检出前的内容。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutUndo(request *model.BatchCheckoutUndoRequest) (*model.BatchCheckoutUndoResponse, error) {
	requestDef := GenReqDefForBatchCheckoutUndo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutUndoResponse), nil
	}
}

// BatchCheckoutUndoInvoker 批量撤销检出M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchCheckoutUndoInvoker(request *model.BatchCheckoutUndoRequest) *BatchCheckoutUndoInvoker {
	requestDef := GenReqDefForBatchCheckoutUndo()
	return &BatchCheckoutUndoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCheckoutUndoByAdmin 管理员批量撤销检出M-V模型数据实例
//
// 本接口用于管理员批量撤销Master-Branch-Version（M-V）模型实例的检出状态，将实例数据批量还原至检出前的最后检入版本内容。适用于数据治理中大规模清理长期锁定实例、批量处理离职员工未检入数据、系统维护前统一释放编辑锁等管理场景。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCheckoutUndoByAdmin(request *model.BatchCheckoutUndoByAdminRequest) (*model.BatchCheckoutUndoByAdminResponse, error) {
	requestDef := GenReqDefForBatchCheckoutUndoByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCheckoutUndoByAdminResponse), nil
	}
}

// BatchCheckoutUndoByAdminInvoker 管理员批量撤销检出M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchCheckoutUndoByAdminInvoker(request *model.BatchCheckoutUndoByAdminRequest) *BatchCheckoutUndoByAdminInvoker {
	requestDef := GenReqDefForBatchCheckoutUndoByAdmin()
	return &BatchCheckoutUndoByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateShareDocs 批量创建分享结构化文档
//
// 本接口用于批量创建结构化文档的分享记录，实现文档的跨用户或跨团队协作共享。当需要将产品设计说明书、工艺卡片、BOM清单等结构化文档共享给团队成员、上下游合作伙伴或全体用户时，可调用本接口一次性完成批量分享配置。
// - 调用本接口时，需在params数组中传入多个分享记录，每个元素对应一条文档分享配置。
// - 分享时可指定被分享用户的权限级别：
//   - 当auth_type为read时，被分享用户仅可查看文档内容。
//   - 当auth_type为write时，被分享用户可查看和编辑文档内容。
// - 当shared_user_id和shared_user_name均设置为all时，表示将文档分享给所有用户。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCreateShareDocs(request *model.BatchCreateShareDocsRequest) (*model.BatchCreateShareDocsResponse, error) {
	requestDef := GenReqDefForBatchCreateShareDocs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateShareDocsResponse), nil
	}
}

// BatchCreateShareDocsInvoker 批量创建分享结构化文档
func (c *IDMEClassicAPIClient) BatchCreateShareDocsInvoker(request *model.BatchCreateShareDocsRequest) *BatchCreateShareDocsInvoker {
	requestDef := GenReqDefForBatchCreateShareDocs()
	return &BatchCreateShareDocsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateUsingPost 批量创建实例
//
// 本接口用于在指定的数据模型（数据实体或关系实体）下，批量创建多个数据实例。调用成功后，将返回所有成功创建的实例信息列表。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCreateUsingPost(request *model.BatchCreateUsingPostRequest) (*model.BatchCreateUsingPostResponse, error) {
	requestDef := GenReqDefForBatchCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateUsingPostResponse), nil
	}
}

// BatchCreateUsingPostInvoker 批量创建实例
func (c *IDMEClassicAPIClient) BatchCreateUsingPostInvoker(request *model.BatchCreateUsingPostRequest) *BatchCreateUsingPostInvoker {
	requestDef := GenReqDefForBatchCreateUsingPost()
	return &BatchCreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateView 批量创建多维视图
//
// 本接口用于批量为指定的一个或多个M-V模型数据实例创建新的多维视图。多维视图允许用户从不同的预设维度（如设计视图、工艺视图、采购视图等）来观察和管理同一个数据对象。在同一个数据实例下，视图的标识（item）必须是唯一的。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“多维视图&amp;多维分支”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchCreateView(request *model.BatchCreateViewRequest) (*model.BatchCreateViewResponse, error) {
	requestDef := GenReqDefForBatchCreateView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateViewResponse), nil
	}
}

// BatchCreateViewInvoker 批量创建多维视图
func (c *IDMEClassicAPIClient) BatchCreateViewInvoker(request *model.BatchCreateViewRequest) *BatchCreateViewInvoker {
	requestDef := GenReqDefForBatchCreateView()
	return &BatchCreateViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteBranch 批量删除最新大版本下的所有小版本
//
// 本接口用于根据主对象ID（masterIds）列表，批量删除指定M-V模型实例中最新分支（Branch）下的所有迭代版本（Iteration）。
// 此操作为物理删除，执行后数据将无法恢复，请谨慎使用。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteBranch(request *model.BatchDeleteBranchRequest) (*model.BatchDeleteBranchResponse, error) {
	requestDef := GenReqDefForBatchDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteBranchResponse), nil
	}
}

// BatchDeleteBranchInvoker 批量删除最新大版本下的所有小版本
func (c *IDMEClassicAPIClient) BatchDeleteBranchInvoker(request *model.BatchDeleteBranchRequest) *BatchDeleteBranchInvoker {
	requestDef := GenReqDefForBatchDeleteBranch()
	return &BatchDeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLatestVersion 批量删除版本对象下最新分支的最新版本实例数据
//
// 本接口用于根据多个主对象（Master）ID，批量永久删除指定M-V模型实体下最新分支的最新版本数据实例。此操作为物理删除，执行后数据将无法恢复，请谨慎使用。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLatestVersion(request *model.BatchDeleteLatestVersionRequest) (*model.BatchDeleteLatestVersionResponse, error) {
	requestDef := GenReqDefForBatchDeleteLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLatestVersionResponse), nil
	}
}

// BatchDeleteLatestVersionInvoker 批量删除版本对象下最新分支的最新版本实例数据
func (c *IDMEClassicAPIClient) BatchDeleteLatestVersionInvoker(request *model.BatchDeleteLatestVersionRequest) *BatchDeleteLatestVersionInvoker {
	requestDef := GenReqDefForBatchDeleteLatestVersion()
	return &BatchDeleteLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLogicalBranch 批量软删除最新大版本下的所有小版本
//
// 本接口用于根据主对象ID（masterIds）列表，批量软删除指定M-V模型实例中最新分支（Branch）下的所有迭代版本（Iteration）。
// 软删除操作并非物理删除，而是将目标分支下的所有版本实例标记为已删除状态（rdmDeleteFlag置为1），并将数据转存至XDM应用的XDMLogicDeleteData内置模型中，以便在需要时进行数据恢复或归档查询。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLogicalBranch(request *model.BatchDeleteLogicalBranchRequest) (*model.BatchDeleteLogicalBranchResponse, error) {
	requestDef := GenReqDefForBatchDeleteLogicalBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLogicalBranchResponse), nil
	}
}

// BatchDeleteLogicalBranchInvoker 批量软删除最新大版本下的所有小版本
func (c *IDMEClassicAPIClient) BatchDeleteLogicalBranchInvoker(request *model.BatchDeleteLogicalBranchRequest) *BatchDeleteLogicalBranchInvoker {
	requestDef := GenReqDefForBatchDeleteLogicalBranch()
	return &BatchDeleteLogicalBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLogicalLatestVersion 批量软删除版本对象下最新分支的最新版本实例数据
//
// 本接口用于根据主对象ID（masterIds）列表，批量软删除指定M-V模型实体下最新分支的最新版本数据实例。
// 软删除操作并非物理删除，而是将目标实例标记为已删除状态（rdmDeleteFlag置为1），并将数据转存至XDM应用的XDMLogicDeleteData内置模型中，以便在需要时进行数据恢复或归档查询。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLogicalLatestVersion(request *model.BatchDeleteLogicalLatestVersionRequest) (*model.BatchDeleteLogicalLatestVersionResponse, error) {
	requestDef := GenReqDefForBatchDeleteLogicalLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLogicalLatestVersionResponse), nil
	}
}

// BatchDeleteLogicalLatestVersionInvoker 批量软删除版本对象下最新分支的最新版本实例数据
func (c *IDMEClassicAPIClient) BatchDeleteLogicalLatestVersionInvoker(request *model.BatchDeleteLogicalLatestVersionRequest) *BatchDeleteLogicalLatestVersionInvoker {
	requestDef := GenReqDefForBatchDeleteLogicalLatestVersion()
	return &BatchDeleteLogicalLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteLogicalUsingPost 批量软删除实例
//
// 本接口用于根据数据实例的唯一编码（id）列表，批量对指定数据模型下的多个数据实例执行软删除（逻辑删除）操作。
// 软删除操作不会从数据库中物理移除数据，而是将所有实例标记为已删除状态，并转存至XDM应用的XDMLogicDeleteData内置模型中。如需彻底删除数据，请使用[批量删除实例 - BatchDeleteUsingPost](BatchDeleteUsingPost.xml)接口。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteLogicalUsingPost(request *model.BatchDeleteLogicalUsingPostRequest) (*model.BatchDeleteLogicalUsingPostResponse, error) {
	requestDef := GenReqDefForBatchDeleteLogicalUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteLogicalUsingPostResponse), nil
	}
}

// BatchDeleteLogicalUsingPostInvoker 批量软删除实例
func (c *IDMEClassicAPIClient) BatchDeleteLogicalUsingPostInvoker(request *model.BatchDeleteLogicalUsingPostRequest) *BatchDeleteLogicalUsingPostInvoker {
	requestDef := GenReqDefForBatchDeleteLogicalUsingPost()
	return &BatchDeleteLogicalUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteShareDocs 批量删除结构化文档分享权限
//
// 本接口用于批量删除结构化文档的分享权限记录，撤销已共享文档的访问权限。当项目结束、人员岗位变动或文档权限需要回收时，可调用本接口一次性撤销多个用户的文档分享权限。
// - 调用本接口时，需在ids中传入待删除的分享权限记录ID列表，系统将对列表中的分享记录执行批量删除。
// - 分享权限删除后，被分享用户将无法再通过原分享渠道访问该文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteShareDocs(request *model.BatchDeleteShareDocsRequest) (*model.BatchDeleteShareDocsResponse, error) {
	requestDef := GenReqDefForBatchDeleteShareDocs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteShareDocsResponse), nil
	}
}

// BatchDeleteShareDocsInvoker 批量删除结构化文档分享权限
func (c *IDMEClassicAPIClient) BatchDeleteShareDocsInvoker(request *model.BatchDeleteShareDocsRequest) *BatchDeleteShareDocsInvoker {
	requestDef := GenReqDefForBatchDeleteShareDocs()
	return &BatchDeleteShareDocsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteStructuredDocument 批量删除结构化文档
//
// 本接口用于批量删除指定数据模型下的结构化文档。当产品下线、项目归档或文档批量清理时，可调用本接口一次性删除多个结构化文档，提升数据管理效率。
// - 调用本接口时，需在ids中传入待删除文档的ID列表，系统将对列表中的文档执行批量删除操作。
// - 当is_check设置为true时，系统将在删除前检查当前用户是否具备文档删除权限，无权限的文档将被跳过删除。
// - 当is_check设置为false时，系统将直接执行删除，不做权限检查。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteStructuredDocument(request *model.BatchDeleteStructuredDocumentRequest) (*model.BatchDeleteStructuredDocumentResponse, error) {
	requestDef := GenReqDefForBatchDeleteStructuredDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteStructuredDocumentResponse), nil
	}
}

// BatchDeleteStructuredDocumentInvoker 批量删除结构化文档
func (c *IDMEClassicAPIClient) BatchDeleteStructuredDocumentInvoker(request *model.BatchDeleteStructuredDocumentRequest) *BatchDeleteStructuredDocumentInvoker {
	requestDef := GenReqDefForBatchDeleteStructuredDocument()
	return &BatchDeleteStructuredDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteUsingPost 批量删除实例
//
// 本接口用于根据数据实例的唯一编码（id）列表，批量删除指定数据模型中的多个数据实例。
// 删除操作不可逆，实例删除后将无法恢复。请在调用前务必确认目标实例ID列表，谨慎操作。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchDeleteUsingPost(request *model.BatchDeleteUsingPostRequest) (*model.BatchDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForBatchDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteUsingPostResponse), nil
	}
}

// BatchDeleteUsingPostInvoker 批量删除实例
func (c *IDMEClassicAPIClient) BatchDeleteUsingPostInvoker(request *model.BatchDeleteUsingPostRequest) *BatchDeleteUsingPostInvoker {
	requestDef := GenReqDefForBatchDeleteUsingPost()
	return &BatchDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchExecuteRevise 批量修订M-V模型数据实例
//
// 本接口用于批量修订一个或多个指定的M-V模型数据实例。修订操作会基于当前实例创建一个新的修订版本，并自动将新实例的version字段更新为下一个修订版本号（例如从A升级到B）。修订后的新实例默认处于“已检入”状态。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchExecuteRevise(request *model.BatchExecuteReviseRequest) (*model.BatchExecuteReviseResponse, error) {
	requestDef := GenReqDefForBatchExecuteRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchExecuteReviseResponse), nil
	}
}

// BatchExecuteReviseInvoker 批量修订M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchExecuteReviseInvoker(request *model.BatchExecuteReviseRequest) *BatchExecuteReviseInvoker {
	requestDef := GenReqDefForBatchExecuteRevise()
	return &BatchExecuteReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchRemoveChildNode 批量移除实例的子节点
//
// 本接口用于批量移除指定数据实例的子节点，解除父子节点间的层级关联关系。当BOM（物料清单）结构发生变更需要拆解某个装配节点、组织架构调整需要撤销某个部门的下属团队归属，或产品分类体系重组需要移除下级类目时，可调用本接口批量解除子节点关联。
// - 调用本接口时，需在parentId中指定父节点实例ID，在childList中传入待移除的子节点实例ID列表。
// - 移除操作执行后，被移除的子节点将发生以下变更：
//   - 父节点（parent）字段被置为空。
//   - 根节点（root）字段被置为空。
// - 当父节点的所有子节点均被移除后，该父节点将被置为叶子节点（即不再拥有子节点）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchRemoveChildNode(request *model.BatchRemoveChildNodeRequest) (*model.BatchRemoveChildNodeResponse, error) {
	requestDef := GenReqDefForBatchRemoveChildNode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchRemoveChildNodeResponse), nil
	}
}

// BatchRemoveChildNodeInvoker 批量移除实例的子节点
func (c *IDMEClassicAPIClient) BatchRemoveChildNodeInvoker(request *model.BatchRemoveChildNodeRequest) *BatchRemoveChildNodeInvoker {
	requestDef := GenReqDefForBatchRemoveChildNode()
	return &BatchRemoveChildNodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchShowGetUsingPost 批量查询实例
//
// 本接口用于根据数据实例的系统唯一标识（id）列表，批量查询指定数据模型下多个实例的完整详细信息。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchShowGetUsingPost(request *model.BatchShowGetUsingPostRequest) (*model.BatchShowGetUsingPostResponse, error) {
	requestDef := GenReqDefForBatchShowGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchShowGetUsingPostResponse), nil
	}
}

// BatchShowGetUsingPostInvoker 批量查询实例
func (c *IDMEClassicAPIClient) BatchShowGetUsingPostInvoker(request *model.BatchShowGetUsingPostRequest) *BatchShowGetUsingPostInvoker {
	requestDef := GenReqDefForBatchShowGetUsingPost()
	return &BatchShowGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAndCheckin 批量更新并检入M-V模型数据实例
//
// 本接口用于批量更新一个或多个指定的M-V模型数据实例，并在更新成功后自动执行检入（Check-in）操作，将实例从“检出”状态切换为“已检入”状态。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateAndCheckin(request *model.BatchUpdateAndCheckinRequest) (*model.BatchUpdateAndCheckinResponse, error) {
	requestDef := GenReqDefForBatchUpdateAndCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAndCheckinResponse), nil
	}
}

// BatchUpdateAndCheckinInvoker 批量更新并检入M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchUpdateAndCheckinInvoker(request *model.BatchUpdateAndCheckinRequest) *BatchUpdateAndCheckinInvoker {
	requestDef := GenReqDefForBatchUpdateAndCheckin()
	return &BatchUpdateAndCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAndRevise 批量修订并更新M-V模型数据实例
//
// 根据主对象ID批量修订并更新M-V模型数据实例，即修订后实例的“version.修订版本”更新为新的修订版本，并同时更新该实例的信息。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateAndRevise(request *model.BatchUpdateAndReviseRequest) (*model.BatchUpdateAndReviseResponse, error) {
	requestDef := GenReqDefForBatchUpdateAndRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAndReviseResponse), nil
	}
}

// BatchUpdateAndReviseInvoker 批量修订并更新M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchUpdateAndReviseInvoker(request *model.BatchUpdateAndReviseRequest) *BatchUpdateAndReviseInvoker {
	requestDef := GenReqDefForBatchUpdateAndRevise()
	return &BatchUpdateAndReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateByAdmin 管理员批量更新M-V模型数据实例
//
// 本接口用于管理员批量强制更新Master-Branch-Version（M-V）模型数据实例，适用于数据治理中的大规模紧急数据修正、批量版本回滚、全量属性修正等管理场景。对于批量中的每个实例，若其唯一编码不存在，则该实例不做任何更新操作，不影响其他实例的正常更新。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateByAdmin(request *model.BatchUpdateByAdminRequest) (*model.BatchUpdateByAdminResponse, error) {
	requestDef := GenReqDefForBatchUpdateByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateByAdminResponse), nil
	}
}

// BatchUpdateByAdminInvoker 管理员批量更新M-V模型数据实例
func (c *IDMEClassicAPIClient) BatchUpdateByAdminInvoker(request *model.BatchUpdateByAdminRequest) *BatchUpdateByAdminInvoker {
	requestDef := GenReqDefForBatchUpdateByAdmin()
	return &BatchUpdateByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateDocument 批量更新结构化文档
//
// 本接口用于批量更新结构化文档的属性信息。
// - 调用本接口时，需在params数组中传入多个文档的更新信息，每个元素对应一个待更新的文档。
// - 每个文档更新项中，id为必填参数，用于唯一标识目标文档；其余字段为可选，仅更新传入的字段。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateDocument(request *model.BatchUpdateDocumentRequest) (*model.BatchUpdateDocumentResponse, error) {
	requestDef := GenReqDefForBatchUpdateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateDocumentResponse), nil
	}
}

// BatchUpdateDocumentInvoker 批量更新结构化文档
func (c *IDMEClassicAPIClient) BatchUpdateDocumentInvoker(request *model.BatchUpdateDocumentRequest) *BatchUpdateDocumentInvoker {
	requestDef := GenReqDefForBatchUpdateDocument()
	return &BatchUpdateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateUsingPost 批量更新实例
//
// 本接口用于批量更新指定数据模型中的多个实例数据。如果请求列表中某个实例的唯一编码（id）在系统中不存在，系统将跳过该实例，不执行任何更新操作，且不会因此导致整个批量请求失败。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateUsingPost(request *model.BatchUpdateUsingPostRequest) (*model.BatchUpdateUsingPostResponse, error) {
	requestDef := GenReqDefForBatchUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateUsingPostResponse), nil
	}
}

// BatchUpdateUsingPostInvoker 批量更新实例
func (c *IDMEClassicAPIClient) BatchUpdateUsingPostInvoker(request *model.BatchUpdateUsingPostRequest) *BatchUpdateUsingPostInvoker {
	requestDef := GenReqDefForBatchUpdateUsingPost()
	return &BatchUpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateVersion 批量升级M-V模型实例的版本号
//
// 本接口用于根据数据实例的唯一标识（id），批量更新/升级指定Master-Branch-Version（M-V）模型实例的版本号信息（包括修订版本version和迭代版本iteration）。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) BatchUpdateVersion(request *model.BatchUpdateVersionRequest) (*model.BatchUpdateVersionResponse, error) {
	requestDef := GenReqDefForBatchUpdateVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateVersionResponse), nil
	}
}

// BatchUpdateVersionInvoker 批量升级M-V模型实例的版本号
func (c *IDMEClassicAPIClient) BatchUpdateVersionInvoker(request *model.BatchUpdateVersionRequest) *BatchUpdateVersionInvoker {
	requestDef := GenReqDefForBatchUpdateVersion()
	return &BatchUpdateVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Checkin 检入M-V模型数据实例
//
// 根据主对象ID检入M-V模型数据实例。已检入的数据实例会生成一个新的迭代版本，并将数据存储至系统中。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) Checkin(request *model.CheckinRequest) (*model.CheckinResponse, error) {
	requestDef := GenReqDefForCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckinResponse), nil
	}
}

// CheckinInvoker 检入M-V模型数据实例
func (c *IDMEClassicAPIClient) CheckinInvoker(request *model.CheckinRequest) *CheckinInvoker {
	requestDef := GenReqDefForCheckin()
	return &CheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Checkout 检出M-V模型数据实例
//
// 本接口用于根据主对象ID（masterId），对指定的M-V模型数据实例执行检出（Check-out）操作。
// 调用本接口后，系统会基于原实例（已检入版本）完全复制生成一个新的数据实例（工作副本），将其状态修改为“工作中”，并对该主对象加锁，阻止其他用户再次检出，直到当前用户执行“检入”或“撤销检出”。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) Checkout(request *model.CheckoutRequest) (*model.CheckoutResponse, error) {
	requestDef := GenReqDefForCheckout()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutResponse), nil
	}
}

// CheckoutInvoker 检出M-V模型数据实例
func (c *IDMEClassicAPIClient) CheckoutInvoker(request *model.CheckoutRequest) *CheckoutInvoker {
	requestDef := GenReqDefForCheckout()
	return &CheckoutInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutAndUpdate 检出并更新M-V模型
//
// 本接口用于根据主对象ID（masterId），对指定的Master-Branch-Version（M-V）模型实例执行检出（Check-out）并同步更新（Update）的原子操作。
// 本接口将“检出”与“更新”合并为一次网络请求，系统在生成新工作副本并加锁的同时，直接将请求中指定的属性值应用到该工作副本上。这大幅减少了网络交互次数，避免了“先检出空副本，再发请求更新”带来的数据不一致风险与性能开销。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutAndUpdate(request *model.CheckoutAndUpdateRequest) (*model.CheckoutAndUpdateResponse, error) {
	requestDef := GenReqDefForCheckoutAndUpdate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutAndUpdateResponse), nil
	}
}

// CheckoutAndUpdateInvoker 检出并更新M-V模型
func (c *IDMEClassicAPIClient) CheckoutAndUpdateInvoker(request *model.CheckoutAndUpdateRequest) *CheckoutAndUpdateInvoker {
	requestDef := GenReqDefForCheckoutAndUpdate()
	return &CheckoutAndUpdateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutUndo 撤销检出M-V模型数据实例
//
// 本接口用于撤销指定M-V模型实例的检出状态。
// 执行撤销检出后，系统会彻底丢弃当前工作副本中的所有未提交修改，将实例数据还原至最近一次检入（Check-in）的历史版本状态，并释放该对象的并发锁，以便自己或他人重新检出。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutUndo(request *model.CheckoutUndoRequest) (*model.CheckoutUndoResponse, error) {
	requestDef := GenReqDefForCheckoutUndo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutUndoResponse), nil
	}
}

// CheckoutUndoInvoker 撤销检出M-V模型数据实例
func (c *IDMEClassicAPIClient) CheckoutUndoInvoker(request *model.CheckoutUndoRequest) *CheckoutUndoInvoker {
	requestDef := GenReqDefForCheckoutUndo()
	return &CheckoutUndoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckoutUndoByAdmin 管理员撤销检出M-V模型数据实例
//
// 本接口用于管理员强制撤销Master-Branch-Version（M-V）模型实例的检出状态，将实例数据还原至检出前的最后检入版本内容。适用于数据治理中因用户离职、会话超时、误操作等导致的实例长期锁定场景。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CheckoutUndoByAdmin(request *model.CheckoutUndoByAdminRequest) (*model.CheckoutUndoByAdminResponse, error) {
	requestDef := GenReqDefForCheckoutUndoByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckoutUndoByAdminResponse), nil
	}
}

// CheckoutUndoByAdminInvoker 管理员撤销检出M-V模型数据实例
func (c *IDMEClassicAPIClient) CheckoutUndoByAdminInvoker(request *model.CheckoutUndoByAdminRequest) *CheckoutUndoByAdminInvoker {
	requestDef := GenReqDefForCheckoutUndoByAdmin()
	return &CheckoutUndoByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectHistoryData 根据时间范围获取模型的历史记录数
//
// 本接口用于根据指定的时间范围，获取数据模型的历史操作统计记录数，包括创建实例、更新实例、删除实例及软删除实例的数量。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“系统版本”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CollectHistoryData(request *model.CollectHistoryDataRequest) (*model.CollectHistoryDataResponse, error) {
	requestDef := GenReqDefForCollectHistoryData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectHistoryDataResponse), nil
	}
}

// CollectHistoryDataInvoker 根据时间范围获取模型的历史记录数
func (c *IDMEClassicAPIClient) CollectHistoryDataInvoker(request *model.CollectHistoryDataRequest) *CollectHistoryDataInvoker {
	requestDef := GenReqDefForCollectHistoryData()
	return &CollectHistoryDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareBusinessVersion 对比M-V模型实例
//
// 本接口用于根据主对象ID（id），获取指定Master-Branch-Version（M-V）模型下两个不同版本实例的完整数据快照，以便进行属性与关系的差异对比。
// **建议：**
// 本接口为基础对比功能。为获得更优的对比体验和更细粒度的控制（如仅对比属性或仅对比关系），推荐使用数据建模引擎（xDM Foundation，简称xDM-F）新增的差异对比接口，即instance-attrs-comparison（属性对比）和instance-relation-comparison（关系对比）接口。更多内容可在应用运行态的“数据服务管理 &gt; 全量数据服务 &gt; 系统管理API &gt; 属性对比API”中查看。
//
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CompareBusinessVersion(request *model.CompareBusinessVersionRequest) (*model.CompareBusinessVersionResponse, error) {
	requestDef := GenReqDefForCompareBusinessVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareBusinessVersionResponse), nil
	}
}

// CompareBusinessVersionInvoker 对比M-V模型实例
func (c *IDMEClassicAPIClient) CompareBusinessVersionInvoker(request *model.CompareBusinessVersionRequest) *CompareBusinessVersionInvoker {
	requestDef := GenReqDefForCompareBusinessVersion()
	return &CompareBusinessVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CompareVersion 数据实例指定版本对比
//
// 本接口用于对比指定数据实例两个版本之间的属性差异和关系差异，返回基础版本对象及差异对比结果。
// 建议使用数据建模引擎（xDM Foundation，简称xDM-F）新增的差异对比功能，即使用instance-attrs-comparison（属性对比）和instance-relation-comparison（关系对比）接口。更多内容可在应用运行态的“数据服务管理 &gt; 全量数据服务 &gt; 系统管理API &gt; 属性对比API”中查看。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“系统版本”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CompareVersion(request *model.CompareVersionRequest) (*model.CompareVersionResponse, error) {
	requestDef := GenReqDefForCompareVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CompareVersionResponse), nil
	}
}

// CompareVersionInvoker 数据实例指定版本对比
func (c *IDMEClassicAPIClient) CompareVersionInvoker(request *model.CompareVersionRequest) *CompareVersionInvoker {
	requestDef := GenReqDefForCompareVersion()
	return &CompareVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountUsingPost 统计指定数据模型的实例总数
//
// 本接口用于根据指定的查询条件，统计指定数据模型中满足条件的实例总数。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CountUsingPost(request *model.CountUsingPostRequest) (*model.CountUsingPostResponse, error) {
	requestDef := GenReqDefForCountUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountUsingPostResponse), nil
	}
}

// CountUsingPostInvoker 统计指定数据模型的实例总数
func (c *IDMEClassicAPIClient) CountUsingPostInvoker(request *model.CountUsingPostRequest) *CountUsingPostInvoker {
	requestDef := GenReqDefForCountUsingPost()
	return &CountUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDocument 创建结构化文档
//
// 本接口用于在指定数据模型下创建结构化文档。结构化文档是工业数据管理中用于承载设计说明书、工艺卡片、BOM清单等工业文档的载体，支持目录（directory）、Page文档（pageDocument）、Board文档（boardDocument）、Mind文档（mindDocument）、Draw文档（drawDocument）等多种文档类型。
// - 创建结构化文档时，需指定文档标题和文档类型。
// - 如需将文档关联到具体的数据模型实例，可通过instance_id参数指定实例ID。
// - 如需将文档挂载到指定目录下，可通过parent_document_id参数指定父文档ID。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CreateDocument(request *model.CreateDocumentRequest) (*model.CreateDocumentResponse, error) {
	requestDef := GenReqDefForCreateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocumentResponse), nil
	}
}

// CreateDocumentInvoker 创建结构化文档
func (c *IDMEClassicAPIClient) CreateDocumentInvoker(request *model.CreateDocumentRequest) *CreateDocumentInvoker {
	requestDef := GenReqDefForCreateDocument()
	return &CreateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateMultiView 创建M-V模型视图数据实例
//
// 本接口用于勾选了“多维视图&amp;多维分支”的MV模型创建数据实例。创建实例时给视图属性赋不同视图属性表示实例所属的视图。视图属性为NULL表示的是默认视图，若创建、更新、修订等MV模型特有接口不指定视图参数则表示操作的视图属性为NULL的视图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CreateMultiView(request *model.CreateMultiViewRequest) (*model.CreateMultiViewResponse, error) {
	requestDef := GenReqDefForCreateMultiView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateMultiViewResponse), nil
	}
}

// CreateMultiViewInvoker 创建M-V模型视图数据实例
func (c *IDMEClassicAPIClient) CreateMultiViewInvoker(request *model.CreateMultiViewRequest) *CreateMultiViewInvoker {
	requestDef := GenReqDefForCreateMultiView()
	return &CreateMultiViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateUsingPost 创建实例
//
// 本接口用于在指定的数据模型（数据实体或关系实体）下创建数据实例。调用成功后，将返回新创建实例的详细信息。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CreateUsingPost(request *model.CreateUsingPostRequest) (*model.CreateUsingPostResponse, error) {
	requestDef := GenReqDefForCreateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateUsingPostResponse), nil
	}
}

// CreateUsingPostInvoker 创建实例
func (c *IDMEClassicAPIClient) CreateUsingPostInvoker(request *model.CreateUsingPostRequest) *CreateUsingPostInvoker {
	requestDef := GenReqDefForCreateUsingPost()
	return &CreateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateView 创建多维视图
//
// 本接口用于为指定的M-V模型数据实例创建一个新的多维视图。多维视图允许用户从不同的预设维度（如设计视图、工艺视图、采购视图等）来观察和管理同一个数据对象，每个视图下的属性或关系可以独立配置和演进。在同一个数据实例下，视图的标识（item）必须是唯一的。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“多维视图&amp;多维分支”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) CreateView(request *model.CreateViewRequest) (*model.CreateViewResponse, error) {
	requestDef := GenReqDefForCreateView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateViewResponse), nil
	}
}

// CreateViewInvoker 创建多维视图
func (c *IDMEClassicAPIClient) CreateViewInvoker(request *model.CreateViewRequest) *CreateViewInvoker {
	requestDef := GenReqDefForCreateView()
	return &CreateViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBranch 删除最新大版本下的所有小版本
//
// 本接口用于根据主对象ID（masterId）和修订版本号（version），删除指定M-V模型实例中最新分支（Branch）下的所有迭代版本（Iteration）。
// 其中：
// - 分支（Branch）：对应M-V模型中的大版本（如版本A、版本B），代表一次完整的修订基线。
// - 迭代（Iteration）：对应分支内的小版本（如A.1、A.2），代表同一修订基线下的多次迭代更新。
// - 本接口将级联删除指定分支下的所有迭代版本，包括该分支的最终版本及其历史迭代记录。
//
// 此操作为物理删除，执行后数据将无法恢复，请谨慎使用。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteBranch(request *model.DeleteBranchRequest) (*model.DeleteBranchResponse, error) {
	requestDef := GenReqDefForDeleteBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBranchResponse), nil
	}
}

// DeleteBranchInvoker 删除最新大版本下的所有小版本
func (c *IDMEClassicAPIClient) DeleteBranchInvoker(request *model.DeleteBranchRequest) *DeleteBranchInvoker {
	requestDef := GenReqDefForDeleteBranch()
	return &DeleteBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteByConditionMultiView 条件删除M-V模型数据实例
//
// 根据用户指定条件删除MV模型数据实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteByConditionMultiView(request *model.DeleteByConditionMultiViewRequest) (*model.DeleteByConditionMultiViewResponse, error) {
	requestDef := GenReqDefForDeleteByConditionMultiView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteByConditionMultiViewResponse), nil
	}
}

// DeleteByConditionMultiViewInvoker 条件删除M-V模型数据实例
func (c *IDMEClassicAPIClient) DeleteByConditionMultiViewInvoker(request *model.DeleteByConditionMultiViewRequest) *DeleteByConditionMultiViewInvoker {
	requestDef := GenReqDefForDeleteByConditionMultiView()
	return &DeleteByConditionMultiViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteByConditionUsingPost 根据指定条件删除实例
//
// 本接口用于根据指定的过滤条件，批量删除指定数据模型下满足条件的所有数据实例。
// 删除操作不可逆，实例删除后将无法恢复。请在调用前务必先使用[分页查询实例 - ShowFindUsingPost](ShowFindUsingPost.xml)接口验证条件表达式的准确性。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteByConditionUsingPost(request *model.DeleteByConditionUsingPostRequest) (*model.DeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteByConditionUsingPostResponse), nil
	}
}

// DeleteByConditionUsingPostInvoker 根据指定条件删除实例
func (c *IDMEClassicAPIClient) DeleteByConditionUsingPostInvoker(request *model.DeleteByConditionUsingPostRequest) *DeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForDeleteByConditionUsingPost()
	return &DeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLatestVersion 删除版本对象下最新分支的最新版本实例数据
//
// 本接口用于根据主对象ID（masterId），精准删除指定Master-Branch-Version（M-V）模型在特定分支下的最新（末端）版本实例。
// 此操作为物理删除，执行后数据将无法恢复，请谨慎使用。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLatestVersion(request *model.DeleteLatestVersionRequest) (*model.DeleteLatestVersionResponse, error) {
	requestDef := GenReqDefForDeleteLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLatestVersionResponse), nil
	}
}

// DeleteLatestVersionInvoker 删除版本对象下最新分支的最新版本实例数据
func (c *IDMEClassicAPIClient) DeleteLatestVersionInvoker(request *model.DeleteLatestVersionRequest) *DeleteLatestVersionInvoker {
	requestDef := GenReqDefForDeleteLatestVersion()
	return &DeleteLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogicalBranch 软删除M-V模型实例下最新分支的所有小版本数据
//
// 本接口用于根据主对象ID（masterId）和修订版本号（version），软删除指定M-V模型实例中最新分支（Branch）下的所有迭代版本（Iteration）。
// 软删除操作并非物理删除，而是将目标分支下的所有版本实例标记为已删除状态（rdmDeleteFlag置为1），并将数据转存至XDM应用的XDMLogicDeleteData内置模型中，以便在需要时进行数据恢复或归档查询。
//
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLogicalBranch(request *model.DeleteLogicalBranchRequest) (*model.DeleteLogicalBranchResponse, error) {
	requestDef := GenReqDefForDeleteLogicalBranch()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogicalBranchResponse), nil
	}
}

// DeleteLogicalBranchInvoker 软删除M-V模型实例下最新分支的所有小版本数据
func (c *IDMEClassicAPIClient) DeleteLogicalBranchInvoker(request *model.DeleteLogicalBranchRequest) *DeleteLogicalBranchInvoker {
	requestDef := GenReqDefForDeleteLogicalBranch()
	return &DeleteLogicalBranchInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteLogicalLatestVersion 软删除版本对象下最新分支的最新版本实例数据
//
// 本接口用于根据主对象ID，软删除指定M-V模型实体下最新分支的最新版本数据实例。
// 软删除操作并非从数据库中物理移除数据，而是将目标实例标记为已删除状态（rdmDeleteFlag置为1），并将该实例的数据转存至XDM应用的XDMLogicDeleteData内置模型中，以便在需要时进行数据恢复或归档查询。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteLogicalLatestVersion(request *model.DeleteLogicalLatestVersionRequest) (*model.DeleteLogicalLatestVersionResponse, error) {
	requestDef := GenReqDefForDeleteLogicalLatestVersion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteLogicalLatestVersionResponse), nil
	}
}

// DeleteLogicalLatestVersionInvoker 软删除版本对象下最新分支的最新版本实例数据
func (c *IDMEClassicAPIClient) DeleteLogicalLatestVersionInvoker(request *model.DeleteLogicalLatestVersionRequest) *DeleteLogicalLatestVersionInvoker {
	requestDef := GenReqDefForDeleteLogicalLatestVersion()
	return &DeleteLogicalLatestVersionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteMultiView 删除M-V模型指定视图版本
//
// 本接口用于勾选了“多维视图&amp;多维分支”的MV模型删除指定的数据实例。若模型配置视图属性，且调用createView接口创建了多维视图。需要在接口参数中给视图属性指定要删除的视图。若入参中书体属性为null，会删除视图属性为null的视图。本接口为同步接口，调用完成后立即返回结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteMultiView(request *model.DeleteMultiViewRequest) (*model.DeleteMultiViewResponse, error) {
	requestDef := GenReqDefForDeleteMultiView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteMultiViewResponse), nil
	}
}

// DeleteMultiViewInvoker 删除M-V模型指定视图版本
func (c *IDMEClassicAPIClient) DeleteMultiViewInvoker(request *model.DeleteMultiViewRequest) *DeleteMultiViewInvoker {
	requestDef := GenReqDefForDeleteMultiView()
	return &DeleteMultiViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTarget 通过目标模型删除关系实体的数据实例
//
// 本接口用于通过目标模型删除关系实体的数据实例，即解除源模型实例与目标模型实例之间的关联关系。
// - 调用本接口时，需在sourceId中指定源模型实例ID，在targetType中指定目标模型的英文名称，系统将删除该源实例与指定目标模型之间的所有关联关系实例。
// - 当latestOnly设置为true时，仅删除源实例关联的最新版本目标模型实例的关系（仅对M-V模型实体生效）。
// - 当latestOnly设置为false时，删除源实例关联的所有版本目标模型实例的关系。
// - 本接口仅执行关系表的DELETE操作，目标端实体的业务数据不受任何影响。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteTarget(request *model.DeleteTargetRequest) (*model.DeleteTargetResponse, error) {
	requestDef := GenReqDefForDeleteTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTargetResponse), nil
	}
}

// DeleteTargetInvoker 通过目标模型删除关系实体的数据实例
func (c *IDMEClassicAPIClient) DeleteTargetInvoker(request *model.DeleteTargetRequest) *DeleteTargetInvoker {
	requestDef := GenReqDefForDeleteTarget()
	return &DeleteTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteUsingPost 删除实例
//
// 本接口用于根据数据实例的唯一编码（id），删除指定数据模型中的一个数据实例。
// 删除操作不可逆，实例删除后将无法恢复。请在调用前务必确认目标实例，谨慎操作。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DeleteUsingPost(request *model.DeleteUsingPostRequest) (*model.DeleteUsingPostResponse, error) {
	requestDef := GenReqDefForDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteUsingPostResponse), nil
	}
}

// DeleteUsingPostInvoker 删除实例
func (c *IDMEClassicAPIClient) DeleteUsingPostInvoker(request *model.DeleteUsingPostRequest) *DeleteUsingPostInvoker {
	requestDef := GenReqDefForDeleteUsingPost()
	return &DeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableDataInstance 失效模型数据实例
//
// 本接口用于批量将指定数据模型的数据实例标记为失效状态，并同步返回实际失效成功的实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) DisableDataInstance(request *model.DisableDataInstanceRequest) (*model.DisableDataInstanceResponse, error) {
	requestDef := GenReqDefForDisableDataInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableDataInstanceResponse), nil
	}
}

// DisableDataInstanceInvoker 失效模型数据实例
func (c *IDMEClassicAPIClient) DisableDataInstanceInvoker(request *model.DisableDataInstanceRequest) *DisableDataInstanceInvoker {
	requestDef := GenReqDefForDisableDataInstance()
	return &DisableDataInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableDataInstance 生效模型数据实例
//
// 本接口用于批量将指定数据模型中处于失效状态的数据实例重新标记为生效状态，并同步返回实际生效成功的实例数量。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) EnableDataInstance(request *model.EnableDataInstanceRequest) (*model.EnableDataInstanceResponse, error) {
	requestDef := GenReqDefForEnableDataInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableDataInstanceResponse), nil
	}
}

// EnableDataInstanceInvoker 生效模型数据实例
func (c *IDMEClassicAPIClient) EnableDataInstanceInvoker(request *model.EnableDataInstanceRequest) *EnableDataInstanceInvoker {
	requestDef := GenReqDefForEnableDataInstance()
	return &EnableDataInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteRevise 修订M-V模型数据实例
//
// 本接口用于对指定的Master-Branch-Version（M-V）模型实例执行修订（Revise）操作。
// 调用本接口后，系统会基于指定的主对象（Master）及其当前最新修订版本，复制并生成一个新的修订版本实例。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ExecuteRevise(request *model.ExecuteReviseRequest) (*model.ExecuteReviseResponse, error) {
	requestDef := GenReqDefForExecuteRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteReviseResponse), nil
	}
}

// ExecuteReviseInvoker 修订M-V模型数据实例
func (c *IDMEClassicAPIClient) ExecuteReviseInvoker(request *model.ExecuteReviseRequest) *ExecuteReviseInvoker {
	requestDef := GenReqDefForExecuteRevise()
	return &ExecuteReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// GenerateBusinessCode 生成新模型业务编码
//
// 本接口用于根据指定数据模型在应用设计态预配置的“业务编码生成器”规则，自动生成符合企业命名规范的业务流水码。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) GenerateBusinessCode(request *model.GenerateBusinessCodeRequest) (*model.GenerateBusinessCodeResponse, error) {
	requestDef := GenReqDefForGenerateBusinessCode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.GenerateBusinessCodeResponse), nil
	}
}

// GenerateBusinessCodeInvoker 生成新模型业务编码
func (c *IDMEClassicAPIClient) GenerateBusinessCodeInvoker(request *model.GenerateBusinessCodeRequest) *GenerateBusinessCodeInvoker {
	requestDef := GenReqDefForGenerateBusinessCode()
	return &GenerateBusinessCodeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllVersions 获取指定M-V模型实例的版本列表
//
// 本接口用于根据主对象ID（masterId），分页获取指定Master-Branch-Version（M-V）模型实例的所有版本信息（包含对应版本下的属性信息）。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListAllVersions(request *model.ListAllVersionsRequest) (*model.ListAllVersionsResponse, error) {
	requestDef := GenReqDefForListAllVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllVersionsResponse), nil
	}
}

// ListAllVersionsInvoker 获取指定M-V模型实例的版本列表
func (c *IDMEClassicAPIClient) ListAllVersionsInvoker(request *model.ListAllVersionsRequest) *ListAllVersionsInvoker {
	requestDef := GenReqDefForListAllVersions()
	return &ListAllVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBatchQueryRelatedObjects 批量查询关系实体关联模型的信息
//
// 本接口用于批量查询指定关系实体所关联的源模型或目标模型的所有实例信息，返回结果包含实例的具体属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListBatchQueryRelatedObjects(request *model.ListBatchQueryRelatedObjectsRequest) (*model.ListBatchQueryRelatedObjectsResponse, error) {
	requestDef := GenReqDefForListBatchQueryRelatedObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBatchQueryRelatedObjectsResponse), nil
	}
}

// ListBatchQueryRelatedObjectsInvoker 批量查询关系实体关联模型的信息
func (c *IDMEClassicAPIClient) ListBatchQueryRelatedObjectsInvoker(request *model.ListBatchQueryRelatedObjectsRequest) *ListBatchQueryRelatedObjectsInvoker {
	requestDef := GenReqDefForListBatchQueryRelatedObjects()
	return &ListBatchQueryRelatedObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGetAllParentList 获取所有父节点
//
// 本接口用于获取指定数据实例的所有父节点列表，同时返回各父节点的列表属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListGetAllParentList(request *model.ListGetAllParentListRequest) (*model.ListGetAllParentListResponse, error) {
	requestDef := GenReqDefForListGetAllParentList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGetAllParentListResponse), nil
	}
}

// ListGetAllParentListInvoker 获取所有父节点
func (c *IDMEClassicAPIClient) ListGetAllParentListInvoker(request *model.ListGetAllParentListRequest) *ListGetAllParentListInvoker {
	requestDef := GenReqDefForListGetAllParentList()
	return &ListGetAllParentListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGetChildList 获取子节点
//
// 本接口用于分页获取指定数据实例的直接子节点列表，同时返回各子节点的列表属性。当需要展开BOM（物料清单）的某个装配节点查看下级零部件、浏览组织架构中某个部门的下属团队或查看产品分类的下一级类目时，可调用本接口获取直接子节点信息。
// 调用本接口时，需在parentId中传入目标数据实例的ID，系统将返回该实例的直接子节点（仅返回一层，不向下递归）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListGetChildList(request *model.ListGetChildListRequest) (*model.ListGetChildListResponse, error) {
	requestDef := GenReqDefForListGetChildList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGetChildListResponse), nil
	}
}

// ListGetChildListInvoker 获取子节点
func (c *IDMEClassicAPIClient) ListGetChildListInvoker(request *model.ListGetChildListRequest) *ListGetChildListInvoker {
	requestDef := GenReqDefForListGetChildList()
	return &ListGetChildListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListHistoryData 分页查询模型历史版本信息
//
// 本接口用于分页查询指定数据实例的历史版本信息。系统以数据实例的最后修改时间作为查询条件，根据您指定的时间范围返回该实例的历史版本变更记录。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“系统版本”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListHistoryData(request *model.ListHistoryDataRequest) (*model.ListHistoryDataResponse, error) {
	requestDef := GenReqDefForListHistoryData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListHistoryDataResponse), nil
	}
}

// ListHistoryDataInvoker 分页查询模型历史版本信息
func (c *IDMEClassicAPIClient) ListHistoryDataInvoker(request *model.ListHistoryDataRequest) *ListHistoryDataInvoker {
	requestDef := GenReqDefForListHistoryData()
	return &ListHistoryDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryDocuments 查询结构化文档
//
// 本接口用于查询指定数据模型下的结构化文档列表。
// - 当请求中传入instance_id时，返回与该实例关联的结构化文档列表。
// - 当请求中传入type时，返回指定类型的结构化文档列表。
// - 当请求中同时传入instance_id和type时，返回同时满足两个条件的结构化文档列表。
// - 当请求中不传入任何筛选条件时，返回该数据模型下所有的结构化文档列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryDocuments(request *model.ListQueryDocumentsRequest) (*model.ListQueryDocumentsResponse, error) {
	requestDef := GenReqDefForListQueryDocuments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryDocumentsResponse), nil
	}
}

// ListQueryDocumentsInvoker 查询结构化文档
func (c *IDMEClassicAPIClient) ListQueryDocumentsInvoker(request *model.ListQueryDocumentsRequest) *ListQueryDocumentsInvoker {
	requestDef := GenReqDefForListQueryDocuments()
	return &ListQueryDocumentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryRelatedObjects 查询关系实体关联模型的信息
//
// 本接口用于分页查询指定关系实体所关联的源模型或目标模型的所有实例信息，返回结果包含关联实例的具体属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryRelatedObjects(request *model.ListQueryRelatedObjectsRequest) (*model.ListQueryRelatedObjectsResponse, error) {
	requestDef := GenReqDefForListQueryRelatedObjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryRelatedObjectsResponse), nil
	}
}

// ListQueryRelatedObjectsInvoker 查询关系实体关联模型的信息
func (c *IDMEClassicAPIClient) ListQueryRelatedObjectsInvoker(request *model.ListQueryRelatedObjectsRequest) *ListQueryRelatedObjectsInvoker {
	requestDef := GenReqDefForListQueryRelatedObjects()
	return &ListQueryRelatedObjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryRelationship 查询关系实体的数据实例
//
// 本接口用于分页查询关系实体的数据实例。通过指定数据实例ID及对应的关系角色（源/目标），或数据模型的英文名称及对应的关系角色（源类型/目标类型），返回匹配的关系实体数据实例信息。
// 如果对应的关系实体存在“参考对象”类型属性，且参考的数据模型为抽象模型，返回信息仅返回对应模型的英文名称和ID。如果参考的数据模型为实体模型，返回空。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryRelationship(request *model.ListQueryRelationshipRequest) (*model.ListQueryRelationshipResponse, error) {
	requestDef := GenReqDefForListQueryRelationship()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryRelationshipResponse), nil
	}
}

// ListQueryRelationshipInvoker 查询关系实体的数据实例
func (c *IDMEClassicAPIClient) ListQueryRelationshipInvoker(request *model.ListQueryRelationshipRequest) *ListQueryRelationshipInvoker {
	requestDef := GenReqDefForListQueryRelationship()
	return &ListQueryRelationshipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryShareDocs 查询结构化文档分享授权列表
//
// 本接口用于查询指定结构化文档的分享授权列表，获取该文档的所有分享记录信息。当需要审计文档的共享范围、回收过期权限或排查文档访问异常时，可调用本接口获取完整的分享授权信息。
// - 调用本接口时，需在params中传入目标结构化文档的ID。
// - 返回结果中包含该文档的所有分享记录，每条记录包含被分享用户、分享用户、权限类型等详细信息。
// - 返回结果中的id字段为分享权限记录的唯一标识，可用于[批量删除结构化文档分享权限 - BatchDeleteShareDocs](BatchDeleteShareDocs.xml)接口删除指定分享权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryShareDocs(request *model.ListQueryShareDocsRequest) (*model.ListQueryShareDocsResponse, error) {
	requestDef := GenReqDefForListQueryShareDocs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryShareDocsResponse), nil
	}
}

// ListQueryShareDocsInvoker 查询结构化文档分享授权列表
func (c *IDMEClassicAPIClient) ListQueryShareDocsInvoker(request *model.ListQueryShareDocsRequest) *ListQueryShareDocsInvoker {
	requestDef := GenReqDefForListQueryShareDocs()
	return &ListQueryShareDocsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryTarget 通过源模型实例ID查询关联的目标模型实例
//
// 本接口用于通过源模型的数据实例ID，分页查询并返回与该实例关联的目标模型数据实例信息。返回的实例信息包含目标模板的“列表属性”（即模型中标记为列表展示的属性）。
// 如果目标模型存在“参考对象”类型的属性，且参考的数据模型为抽象模型，返回信息仅返回对应模型的英文名称和ID。如果参考的数据模型为实体模型，返回空。
// 如果目标对象是M-V模型，可通过设置latestOnly&#x3D;true仅返回源实例关联的最新版本目标对象；默认返回所有版本。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryTarget(request *model.ListQueryTargetRequest) (*model.ListQueryTargetResponse, error) {
	requestDef := GenReqDefForListQueryTarget()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryTargetResponse), nil
	}
}

// ListQueryTargetInvoker 通过源模型实例ID查询关联的目标模型实例
func (c *IDMEClassicAPIClient) ListQueryTargetInvoker(request *model.ListQueryTargetRequest) *ListQueryTargetInvoker {
	requestDef := GenReqDefForListQueryTarget()
	return &ListQueryTargetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQueryUsingPost 根据“列表属性”为“是”的属性查询实例
//
// 当数据模型中存在“列表属性”为“是”的属性时，可通过此接口查询数据模型中的实例数据，且支持分页。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 目标数据模型在应用设计态中，必须至少存在一个“列表属性”为“是”的属性。
// 2. 已在应用设计态完成数据模型的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 3. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListQueryUsingPost(request *model.ListQueryUsingPostRequest) (*model.ListQueryUsingPostResponse, error) {
	requestDef := GenReqDefForListQueryUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQueryUsingPostResponse), nil
	}
}

// ListQueryUsingPostInvoker 根据“列表属性”为“是”的属性查询实例
func (c *IDMEClassicAPIClient) ListQueryUsingPostInvoker(request *model.ListQueryUsingPostRequest) *ListQueryUsingPostInvoker {
	requestDef := GenReqDefForListQueryUsingPost()
	return &ListQueryUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSelectUsingPost 查询实例的指定属性
//
// 本接口用于根据查询条件，按需返回数据模型实例的指定属性信息。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListSelectUsingPost(request *model.ListSelectUsingPostRequest) (*model.ListSelectUsingPostResponse, error) {
	requestDef := GenReqDefForListSelectUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSelectUsingPostResponse), nil
	}
}

// ListSelectUsingPostInvoker 查询实例的指定属性
func (c *IDMEClassicAPIClient) ListSelectUsingPostInvoker(request *model.ListSelectUsingPostRequest) *ListSelectUsingPostInvoker {
	requestDef := GenReqDefForListSelectUsingPost()
	return &ListSelectUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListUsingPost 查询实例的基础属性
//
// 本接口用于根据查询条件，分页检索数据模型实例的基础属性与系统级元数据信息。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ListUsingPost(request *model.ListUsingPostRequest) (*model.ListUsingPostResponse, error) {
	requestDef := GenReqDefForListUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListUsingPostResponse), nil
	}
}

// ListUsingPostInvoker 查询实例的基础属性
func (c *IDMEClassicAPIClient) ListUsingPostInvoker(request *model.ListUsingPostRequest) *ListUsingPostInvoker {
	requestDef := GenReqDefForListUsingPost()
	return &ListUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Refresh 刷新树形节点
//
// 本接口用于刷新指定数据实例对应的节点全路径。在调用本接口前请确保数据模型具有“树形结构”功能。
// 调用本接口时，如果未指定数据实例或指定的数据实例为父节点，则刷新整棵树的所有节点全路径。
// 本接口为异步接口，调用后立即返回结果，但实际上整棵树的结构还在刷新中，需适时等待后（刷新时间视该树的大小而定）再查看该树的具体数据是否已刷新。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) Refresh(request *model.RefreshRequest) (*model.RefreshResponse, error) {
	requestDef := GenReqDefForRefresh()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RefreshResponse), nil
	}
}

// RefreshInvoker 刷新树形节点
func (c *IDMEClassicAPIClient) RefreshInvoker(request *model.RefreshRequest) *RefreshInvoker {
	requestDef := GenReqDefForRefresh()
	return &RefreshInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveFromCategory 移除数据分类
//
// 本接口用于数据实例从指定分类中移除。当零部件变更产品线、设备调整工艺归属或物料重新归类时，可调用本接口解除实例与分类的关联关系。
// - 当实例不存在时，系统将抛出异常。
// - 当实例与指定分类不存在关联关系时，返回值为0。
// - 当实例与指定分类存在关联关系时，返回值为成功移除的关联数据分类个数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) RemoveFromCategory(request *model.RemoveFromCategoryRequest) (*model.RemoveFromCategoryResponse, error) {
	requestDef := GenReqDefForRemoveFromCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveFromCategoryResponse), nil
	}
}

// RemoveFromCategoryInvoker 移除数据分类
func (c *IDMEClassicAPIClient) RemoveFromCategoryInvoker(request *model.RemoveFromCategoryRequest) *RemoveFromCategoryInvoker {
	requestDef := GenReqDefForRemoveFromCategory()
	return &RemoveFromCategoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveTag 解绑标签
//
// 本接口用于为指定数据模型的数据实例解绑标签，移除实例与标签之间的关联关系。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“标签管理”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) RemoveTag(request *model.RemoveTagRequest) (*model.RemoveTagResponse, error) {
	requestDef := GenReqDefForRemoveTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveTagResponse), nil
	}
}

// RemoveTagInvoker 解绑标签
func (c *IDMEClassicAPIClient) RemoveTagInvoker(request *model.RemoveTagRequest) *RemoveTagInvoker {
	requestDef := GenReqDefForRemoveTag()
	return &RemoveTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveAllUsingPost 根据唯一键为“是”的属性更新实例数据
//
// 本接口用于根据数据模型中“唯一键”为“是”的属性（业务唯一键），对数据实例执行全量保存或更新操作。
// 更新（Update）：若系统中已存在匹配该唯一键的实例，则更新其所有字段。
// 创建（Insert）：若系统中不存在匹配该唯一键的实例，则自动创建一条新实例。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveAllUsingPost(request *model.SaveAllUsingPostRequest) (*model.SaveAllUsingPostResponse, error) {
	requestDef := GenReqDefForSaveAllUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveAllUsingPostResponse), nil
	}
}

// SaveAllUsingPostInvoker 根据唯一键为“是”的属性更新实例数据
func (c *IDMEClassicAPIClient) SaveAllUsingPostInvoker(request *model.SaveAllUsingPostRequest) *SaveAllUsingPostInvoker {
	requestDef := GenReqDefForSaveAllUsingPost()
	return &SaveAllUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveAsUsingPost 另存实例数据
//
// 本接口用于对指定数据模型的数据实例执行另存为（SaveAs）操作。系统将以指定的源实例为模板，克隆其数据并生成一条全新的实例记录。
// 在另存过程中，用户可以灵活控制新实例的属性值：
// - 未作特殊指定的属性，将完全保持与源实例一致。
// - 可通过entityToSave为指定属性赋予新的值。
// - 可通过needSetNullAttrs将指定属性强制清空（置为Null）。
// 另存操作不会修改或覆盖源实例，而是生成一个具有全新系统主键（id）的新实例。此功能常用于工业场景中的“图纸/模型版本迭代”、“BOM复制建版”或“基于历史数据快速创建新对象”等业务。
//
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveAsUsingPost(request *model.SaveAsUsingPostRequest) (*model.SaveAsUsingPostResponse, error) {
	requestDef := GenReqDefForSaveAsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveAsUsingPostResponse), nil
	}
}

// SaveAsUsingPostInvoker 另存实例数据
func (c *IDMEClassicAPIClient) SaveAsUsingPostInvoker(request *model.SaveAsUsingPostRequest) *SaveAsUsingPostInvoker {
	requestDef := GenReqDefForSaveAsUsingPost()
	return &SaveAsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SaveUsingPost 根据唯一键为“是”的属性更新实例的指定字段
//
// 本接口用于根据数据模型中“唯一键”为“是”的属性（业务唯一键），对数据实例执行指定字段的保存或更新操作。
//   - 更新（Update）：若系统中已存在匹配该唯一键的实例，则仅更新请求体中传入的指定字段。
//   - 创建（Insert）：若系统中不存在匹配该唯一键的实例，则自动创建一条新实例。
//
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SaveUsingPost(request *model.SaveUsingPostRequest) (*model.SaveUsingPostResponse, error) {
	requestDef := GenReqDefForSaveUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SaveUsingPostResponse), nil
	}
}

// SaveUsingPostInvoker 根据唯一键为“是”的属性更新实例的指定字段
func (c *IDMEClassicAPIClient) SaveUsingPostInvoker(request *model.SaveUsingPostRequest) *SaveUsingPostInvoker {
	requestDef := GenReqDefForSaveUsingPost()
	return &SaveUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFindUsingPost 分页查询实例
//
// 本接口用于分页查询指定数据模型中的数据实例列表。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowFindUsingPost(request *model.ShowFindUsingPostRequest) (*model.ShowFindUsingPostResponse, error) {
	requestDef := GenReqDefForShowFindUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFindUsingPostResponse), nil
	}
}

// ShowFindUsingPostInvoker 分页查询实例
func (c *IDMEClassicAPIClient) ShowFindUsingPostInvoker(request *model.ShowFindUsingPostRequest) *ShowFindUsingPostInvoker {
	requestDef := GenReqDefForShowFindUsingPost()
	return &ShowFindUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetByUniqueKey 根据唯一键为“是”的基本属性查询实例
//
// 当数据模型中存在“唯一键”为“是”的基本属性时，可根据该属性查询实例数据。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetByUniqueKey(request *model.ShowGetByUniqueKeyRequest) (*model.ShowGetByUniqueKeyResponse, error) {
	requestDef := GenReqDefForShowGetByUniqueKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetByUniqueKeyResponse), nil
	}
}

// ShowGetByUniqueKeyInvoker 根据唯一键为“是”的基本属性查询实例
func (c *IDMEClassicAPIClient) ShowGetByUniqueKeyInvoker(request *model.ShowGetByUniqueKeyRequest) *ShowGetByUniqueKeyInvoker {
	requestDef := GenReqDefForShowGetByUniqueKey()
	return &ShowGetByUniqueKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetParent 获取父节点
//
// 本接口用于获取指定数据实例的直接父节点，同时返回父节点的列表属性。当需要查询BOM（物料清单）中某个零部件的直接上级装配节点、查看组织架构中某个部门的直接汇报部门或获取产品分类的直接上级类目时，可调用本接口获取直接父节点信息。
// - 调用本接口时，需在childId中传入目标数据实例的ID，系统将返回该实例的直接父节点（仅返回一层，不向上递归）。
// - 返回结果中包含父节点的完整属性信息，包括创建者、创建时间、租户信息等列表属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetParent(request *model.ShowGetParentRequest) (*model.ShowGetParentResponse, error) {
	requestDef := GenReqDefForShowGetParent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetParentResponse), nil
	}
}

// ShowGetParentInvoker 获取父节点
func (c *IDMEClassicAPIClient) ShowGetParentInvoker(request *model.ShowGetParentRequest) *ShowGetParentInvoker {
	requestDef := GenReqDefForShowGetParent()
	return &ShowGetParentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetRoot 获取根节点
//
// 本接口用于获取指定数据实例所在树形结构的根节点信息。当需要追溯BOM（物料清单）的顶层装配节点、查找组织架构的顶层部门或定位产品分类的顶级类目时，可调用本接口获取目标节点的根节点。
// 调用本接口时，需在id中传入目标数据实例的ID，系统将沿树形结构向上遍历，返回该实例所在树的根节点。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetRoot(request *model.ShowGetRootRequest) (*model.ShowGetRootResponse, error) {
	requestDef := GenReqDefForShowGetRoot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetRootResponse), nil
	}
}

// ShowGetRootInvoker 获取根节点
func (c *IDMEClassicAPIClient) ShowGetRootInvoker(request *model.ShowGetRootRequest) *ShowGetRootInvoker {
	requestDef := GenReqDefForShowGetRoot()
	return &ShowGetRootInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetTokens 获取Token信息
//
// 本接口用于通过文档ID和认证类型获取结构化文档的访问Token。获取Token后，可在Token有效期内对指定结构化文档进行只读或读写操作。
// - 当认证类型为read时，获取的Token仅支持对文档进行查看操作。
// - 当认证类型为write时，获取的Token支持对文档进行编辑、保存等操作。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetTokens(request *model.ShowGetTokensRequest) (*model.ShowGetTokensResponse, error) {
	requestDef := GenReqDefForShowGetTokens()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetTokensResponse), nil
	}
}

// ShowGetTokensInvoker 获取Token信息
func (c *IDMEClassicAPIClient) ShowGetTokensInvoker(request *model.ShowGetTokensRequest) *ShowGetTokensInvoker {
	requestDef := GenReqDefForShowGetTokens()
	return &ShowGetTokensInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGetUsingPost 查询实例详情
//
// 本接口用于根据数据实例的唯一编码（id），查询指定数据模型下单个数据实例的完整详细信息。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowGetUsingPost(request *model.ShowGetUsingPostRequest) (*model.ShowGetUsingPostResponse, error) {
	requestDef := GenReqDefForShowGetUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGetUsingPostResponse), nil
	}
}

// ShowGetUsingPostInvoker 查询实例详情
func (c *IDMEClassicAPIClient) ShowGetUsingPostInvoker(request *model.ShowGetUsingPostRequest) *ShowGetUsingPostInvoker {
	requestDef := GenReqDefForShowGetUsingPost()
	return &ShowGetUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteByConditionUsingPost 根据指定条件软删除实例
//
// 本接口用于根据指定的过滤条件，批量对指定数据模型下满足条件的所有数据实例执行软删除（逻辑删除）操作。
// 通过此接口进行删除操作时，系统会将当前删除的实例转存至XDM应用的XDMLogicDeleteData内置模型中。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPost(request *model.ShowLogicalDeleteByConditionUsingPostRequest) (*model.ShowLogicalDeleteByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteByConditionUsingPostResponse), nil
	}
}

// ShowLogicalDeleteByConditionUsingPostInvoker 根据指定条件软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteByConditionUsingPostInvoker(request *model.ShowLogicalDeleteByConditionUsingPostRequest) *ShowLogicalDeleteByConditionUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteByConditionUsingPost()
	return &ShowLogicalDeleteByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLogicalDeleteUsingPost 软删除实例
//
// 本接口用于根据数据实例的唯一编码（id），对指定数据模型下的单个数据实例执行软删除（逻辑删除）操作。
// 软删除操作不会从数据库中物理移除数据，而是将实例标记为已删除状态，并转存至XDM应用的XDMLogicDeleteData内置模型中。如需彻底删除数据，请使用[删除实例 - DeleteUsingPost](DeleteUsingPost.xml)接口。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPost(request *model.ShowLogicalDeleteUsingPostRequest) (*model.ShowLogicalDeleteUsingPostResponse, error) {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLogicalDeleteUsingPostResponse), nil
	}
}

// ShowLogicalDeleteUsingPostInvoker 软删除实例
func (c *IDMEClassicAPIClient) ShowLogicalDeleteUsingPostInvoker(request *model.ShowLogicalDeleteUsingPostRequest) *ShowLogicalDeleteUsingPostInvoker {
	requestDef := GenReqDefForShowLogicalDeleteUsingPost()
	return &ShowLogicalDeleteUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStaticsPage 分页查询数据实例的统计信息
//
// 本接口用于根据指定的聚合函数与分组条件，对指定数据模型的实例数据进行统计计算，并支持对分组后的统计结果进行分页返回。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowStaticsPage(request *model.ShowStaticsPageRequest) (*model.ShowStaticsPageResponse, error) {
	requestDef := GenReqDefForShowStaticsPage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStaticsPageResponse), nil
	}
}

// ShowStaticsPageInvoker 分页查询数据实例的统计信息
func (c *IDMEClassicAPIClient) ShowStaticsPageInvoker(request *model.ShowStaticsPageRequest) *ShowStaticsPageInvoker {
	requestDef := GenReqDefForShowStaticsPage()
	return &ShowStaticsPageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStaticsUsingPost 查询指定数据模型的实例统计信息
//
// 本接口用于根据指定的聚合函数与分组条件，对指定数据模型的实例数据进行统计计算（如计数、求平均值、求最大/最小值等）。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowStaticsUsingPost(request *model.ShowStaticsUsingPostRequest) (*model.ShowStaticsUsingPostResponse, error) {
	requestDef := GenReqDefForShowStaticsUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStaticsUsingPostResponse), nil
	}
}

// ShowStaticsUsingPostInvoker 查询指定数据模型的实例统计信息
func (c *IDMEClassicAPIClient) ShowStaticsUsingPostInvoker(request *model.ShowStaticsUsingPostRequest) *ShowStaticsUsingPostInvoker {
	requestDef := GenReqDefForShowStaticsUsingPost()
	return &ShowStaticsUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTag 查询指定数据实例的标签详情
//
// 本接口用于查询指定模型的数据实例已绑定的所有标签详情，返回标签的完整定义信息及所属标签组信息。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“标签管理”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowTag(request *model.ShowTagRequest) (*model.ShowTagResponse, error) {
	requestDef := GenReqDefForShowTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagResponse), nil
	}
}

// ShowTagInvoker 查询指定数据实例的标签详情
func (c *IDMEClassicAPIClient) ShowTagInvoker(request *model.ShowTagRequest) *ShowTagInvoker {
	requestDef := GenReqDefForShowTag()
	return &ShowTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVersionByMaster 获取指定版本的M-V模型实例数据
//
// 本接口用于根据主对象ID、迭代版本和版本号，查询M-V模型实例的详细版本信息。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) ShowVersionByMaster(request *model.ShowVersionByMasterRequest) (*model.ShowVersionByMasterResponse, error) {
	requestDef := GenReqDefForShowVersionByMaster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVersionByMasterResponse), nil
	}
}

// ShowVersionByMasterInvoker 获取指定版本的M-V模型实例数据
func (c *IDMEClassicAPIClient) ShowVersionByMasterInvoker(request *model.ShowVersionByMasterRequest) *ShowVersionByMasterInvoker {
	requestDef := GenReqDefForShowVersionByMaster()
	return &ShowVersionByMasterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SwitchLifecycleTemplate 切换生命周期模板
//
// 本接口用于为指定模型的数据实例切换绑定的生命周期模板，并同时指定实例在新模板中的生命周期状态。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“生命周期管理”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
// 3. 已在运行态创建目标生命周期模板，且模板中包含目标生命周期状态。具体操作请参见[生命周期管理](https://support.huaweicloud.com/usermanual-idme/idme_clientog_0068.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) SwitchLifecycleTemplate(request *model.SwitchLifecycleTemplateRequest) (*model.SwitchLifecycleTemplateResponse, error) {
	requestDef := GenReqDefForSwitchLifecycleTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SwitchLifecycleTemplateResponse), nil
	}
}

// SwitchLifecycleTemplateInvoker 切换生命周期模板
func (c *IDMEClassicAPIClient) SwitchLifecycleTemplateInvoker(request *model.SwitchLifecycleTemplateRequest) *SwitchLifecycleTemplateInvoker {
	requestDef := GenReqDefForSwitchLifecycleTemplate()
	return &SwitchLifecycleTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAndCheckin 更新并检入M-V模型数据实例
//
// 本接口用于更新指定M-V模型实例的数据，并同步完成检入操作，即一步完成“修改-提交”流程。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateAndCheckin(request *model.UpdateAndCheckinRequest) (*model.UpdateAndCheckinResponse, error) {
	requestDef := GenReqDefForUpdateAndCheckin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAndCheckinResponse), nil
	}
}

// UpdateAndCheckinInvoker 更新并检入M-V模型数据实例
func (c *IDMEClassicAPIClient) UpdateAndCheckinInvoker(request *model.UpdateAndCheckinRequest) *UpdateAndCheckinInvoker {
	requestDef := GenReqDefForUpdateAndCheckin()
	return &UpdateAndCheckinInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAndRevise 修订并更新M-V模型数据实例
//
// 本用于修订一个指定的M-V模型数据实例，并在修订生成新版本的同时，更新该新实例的业务数据。
// 修订操作会将version字段升级为新的修订版本号（例如从A升级到B），随后将data参数中指定的属性值应用到新版本实例上。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateAndRevise(request *model.UpdateAndReviseRequest) (*model.UpdateAndReviseResponse, error) {
	requestDef := GenReqDefForUpdateAndRevise()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAndReviseResponse), nil
	}
}

// UpdateAndReviseInvoker 修订并更新M-V模型数据实例
func (c *IDMEClassicAPIClient) UpdateAndReviseInvoker(request *model.UpdateAndReviseRequest) *UpdateAndReviseInvoker {
	requestDef := GenReqDefForUpdateAndRevise()
	return &UpdateAndReviseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateByAdmin 管理员更新M-V模型数据实例
//
// 本接口用于管理员强制更新Master-Branch-Version（M-V）模型数据实例，适用于数据治理中的紧急数据修正、版本回滚、批量属性修正等管理场景。当实例的唯一编码不存在时，系统将不做任何更新操作。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态完成M-V模型实体（即“父模型”为“VersionObject”的数据实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateByAdmin(request *model.UpdateByAdminRequest) (*model.UpdateByAdminResponse, error) {
	requestDef := GenReqDefForUpdateByAdmin()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateByAdminResponse), nil
	}
}

// UpdateByAdminInvoker 管理员更新M-V模型数据实例
func (c *IDMEClassicAPIClient) UpdateByAdminInvoker(request *model.UpdateByAdminRequest) *UpdateByAdminInvoker {
	requestDef := GenReqDefForUpdateByAdmin()
	return &UpdateByAdminInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateByConditionUsingPost 根据指定条件更新实例
//
// 本接口用于根据指定的过滤条件，批量更新指定数据模型下满足条件的所有数据实例。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateByConditionUsingPost(request *model.UpdateByConditionUsingPostRequest) (*model.UpdateByConditionUsingPostResponse, error) {
	requestDef := GenReqDefForUpdateByConditionUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateByConditionUsingPostResponse), nil
	}
}

// UpdateByConditionUsingPostInvoker 根据指定条件更新实例
func (c *IDMEClassicAPIClient) UpdateByConditionUsingPostInvoker(request *model.UpdateByConditionUsingPostRequest) *UpdateByConditionUsingPostInvoker {
	requestDef := GenReqDefForUpdateByConditionUsingPost()
	return &UpdateByConditionUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDocument 更新文档标题
//
// 本接口用于更新结构化文档的标题信息。当文档标题需要随产品版本迭代、工艺变更或项目阶段调整时，可调用本接口修改文档标题。
// 调用本接口时，需通过document_id或instance_id指定目标文档。
// 若同时传入document_id和instance_id，优先以document_id定位目标文档。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateDocument(request *model.UpdateDocumentRequest) (*model.UpdateDocumentResponse, error) {
	requestDef := GenReqDefForUpdateDocument()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDocumentResponse), nil
	}
}

// UpdateDocumentInvoker 更新文档标题
func (c *IDMEClassicAPIClient) UpdateDocumentInvoker(request *model.UpdateDocumentRequest) *UpdateDocumentInvoker {
	requestDef := GenReqDefForUpdateDocument()
	return &UpdateDocumentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateState 设置生命周期的状态
//
// 本接口用于修改或切换数据实例所绑定的生命周期状态，支持单个或批量实例的状态更新。
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“生命周期管理”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
// 3. 已在运行态创建并发布目标生命周期模板，且模板中包含目标生命周期状态。具体操作请参见[生命周期管理](https://support.huaweicloud.com/usermanual-idme/idme_clientog_0068.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateState(request *model.UpdateStateRequest) (*model.UpdateStateResponse, error) {
	requestDef := GenReqDefForUpdateState()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStateResponse), nil
	}
}

// UpdateStateInvoker 设置生命周期的状态
func (c *IDMEClassicAPIClient) UpdateStateInvoker(request *model.UpdateStateRequest) *UpdateStateInvoker {
	requestDef := GenReqDefForUpdateState()
	return &UpdateStateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateUsingPost 更新实例
//
// 本接口用于更新指定数据模型中的一个实例数据。如果请求中指定的实例唯一编码（id）不存在，接口将不会执行任何更新操作，也不会返回错误。
// 在调用本接口前，请确保目标数据模型已满足实例化条件：
// 1. 已在应用设计态完成数据模型（数据实体或关系实体）的创建与发布。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateUsingPost(request *model.UpdateUsingPostRequest) (*model.UpdateUsingPostResponse, error) {
	requestDef := GenReqDefForUpdateUsingPost()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateUsingPostResponse), nil
	}
}

// UpdateUsingPostInvoker 更新实例
func (c *IDMEClassicAPIClient) UpdateUsingPostInvoker(request *model.UpdateUsingPostRequest) *UpdateUsingPostInvoker {
	requestDef := GenReqDefForUpdateUsingPost()
	return &UpdateUsingPostInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateView 更新M-V模型数据实例的多维视图属性
//
// 本接口用于更新指定M-V模型实例的多维视图属性。
// 在多维视图管理中，一个零部件（Part）可能衍生出“研发”、“工艺”、“采购”等多个视图。随着业务流转或数据纠错的需求，可能需要对已存在的视图实例进行属性调整：
// - 变更视图标识（视角切换）：当视图绑定的分类标识（item，如MultiViewItem对象）分配错误或业务标准变更时，可通过此接口将其重新绑定到正确的视图标识上。
// - 视图属性脱敏/置空：在切换视角或修正数据时，通过needSetNull参数强制清空某些不再适用于当前视角的敏感或冗余属性。
//
// 在调用本接口前，请确保目标数据模型满足以下条件：
// 1. 已在应用设计态创建并发布一个具有“多维视图&amp;多维分支”功能的数据模型。具体操作请参见[创建数据实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0007.html)、[创建关系实体](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0022.html)和[应用发布](https://support.huaweicloud.com/usermanual-idme/idme_usermanual_0085.html)。
// 2. 已将对应应用成功部署至运行态。具体操作请参见[部署应用](https://support.huaweicloud.com/consog-idme/idme_consog_0022.html)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IDMEClassicAPIClient) UpdateView(request *model.UpdateViewRequest) (*model.UpdateViewResponse, error) {
	requestDef := GenReqDefForUpdateView()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateViewResponse), nil
	}
}

// UpdateViewInvoker 更新M-V模型数据实例的多维视图属性
func (c *IDMEClassicAPIClient) UpdateViewInvoker(request *model.UpdateViewRequest) *UpdateViewInvoker {
	requestDef := GenReqDefForUpdateView()
	return &UpdateViewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
