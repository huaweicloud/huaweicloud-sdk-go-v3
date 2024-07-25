package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsstudio/v1/model"
)

type DataArtsStudioClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDataArtsStudioClient(hcClient *httpclient.HcHttpClient) *DataArtsStudioClient {
	return &DataArtsStudioClient{HcClient: hcClient}
}

func DataArtsStudioClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AddDesignEntityTags 添加标签
//
// 根据资产（表或属性）的ID给资产打上标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AddDesignEntityTags(request *model.AddDesignEntityTagsRequest) (*model.AddDesignEntityTagsResponse, error) {
	requestDef := GenReqDefForAddDesignEntityTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDesignEntityTagsResponse), nil
	}
}

// AddDesignEntityTagsInvoker 添加标签
func (c *DataArtsStudioClient) AddDesignEntityTagsInvoker(request *model.AddDesignEntityTagsRequest) *AddDesignEntityTagsInvoker {
	requestDef := GenReqDefForAddDesignEntityTags()
	return &AddDesignEntityTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddTagToAsset 标签关联到资产
//
// 标签关联到资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AddTagToAsset(request *model.AddTagToAssetRequest) (*model.AddTagToAssetResponse, error) {
	requestDef := GenReqDefForAddTagToAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddTagToAssetResponse), nil
	}
}

// AddTagToAssetInvoker 标签关联到资产
func (c *DataArtsStudioClient) AddTagToAssetInvoker(request *model.AddTagToAssetRequest) *AddTagToAssetInvoker {
	requestDef := GenReqDefForAddTagToAsset()
	return &AddTagToAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddWorkSpaceUsers 添加工作空间用户
//
// 添加工作空间用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AddWorkSpaceUsers(request *model.AddWorkSpaceUsersRequest) (*model.AddWorkSpaceUsersResponse, error) {
	requestDef := GenReqDefForAddWorkSpaceUsers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddWorkSpaceUsersResponse), nil
	}
}

// AddWorkSpaceUsersInvoker 添加工作空间用户
func (c *DataArtsStudioClient) AddWorkSpaceUsersInvoker(request *model.AddWorkSpaceUsersRequest) *AddWorkSpaceUsersInvoker {
	requestDef := GenReqDefForAddWorkSpaceUsers()
	return &AddWorkSpaceUsersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateClassificationToEntity 资产关联分类
//
// 将一个分类关联到一个或多个指定guid的资产上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AssociateClassificationToEntity(request *model.AssociateClassificationToEntityRequest) (*model.AssociateClassificationToEntityResponse, error) {
	requestDef := GenReqDefForAssociateClassificationToEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateClassificationToEntityResponse), nil
	}
}

// AssociateClassificationToEntityInvoker 资产关联分类
func (c *DataArtsStudioClient) AssociateClassificationToEntityInvoker(request *model.AssociateClassificationToEntityRequest) *AssociateClassificationToEntityInvoker {
	requestDef := GenReqDefForAssociateClassificationToEntity()
	return &AssociateClassificationToEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateSecurityLevelToEntitie 资产关联密级
//
// 关联资产到密级，资产关联指定密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AssociateSecurityLevelToEntitie(request *model.AssociateSecurityLevelToEntitieRequest) (*model.AssociateSecurityLevelToEntitieResponse, error) {
	requestDef := GenReqDefForAssociateSecurityLevelToEntitie()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateSecurityLevelToEntitieResponse), nil
	}
}

// AssociateSecurityLevelToEntitieInvoker 资产关联密级
func (c *DataArtsStudioClient) AssociateSecurityLevelToEntitieInvoker(request *model.AssociateSecurityLevelToEntitieRequest) *AssociateSecurityLevelToEntitieInvoker {
	requestDef := GenReqDefForAssociateSecurityLevelToEntitie()
	return &AssociateSecurityLevelToEntitieInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchApproveApply 审核申请
//
// 审核申请。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchApproveApply(request *model.BatchApproveApplyRequest) (*model.BatchApproveApplyResponse, error) {
	requestDef := GenReqDefForBatchApproveApply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchApproveApplyResponse), nil
	}
}

// BatchApproveApplyInvoker 审核申请
func (c *DataArtsStudioClient) BatchApproveApplyInvoker(request *model.BatchApproveApplyRequest) *BatchApproveApplyInvoker {
	requestDef := GenReqDefForBatchApproveApply()
	return &BatchApproveApplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateClassificationToEntities 批量资产关联分类
//
// 批量资产关联分类：只支持对数据表的列和OBS对象添加分类
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchAssociateClassificationToEntities(request *model.BatchAssociateClassificationToEntitiesRequest) (*model.BatchAssociateClassificationToEntitiesResponse, error) {
	requestDef := GenReqDefForBatchAssociateClassificationToEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateClassificationToEntitiesResponse), nil
	}
}

// BatchAssociateClassificationToEntitiesInvoker 批量资产关联分类
func (c *DataArtsStudioClient) BatchAssociateClassificationToEntitiesInvoker(request *model.BatchAssociateClassificationToEntitiesRequest) *BatchAssociateClassificationToEntitiesInvoker {
	requestDef := GenReqDefForBatchAssociateClassificationToEntities()
	return &BatchAssociateClassificationToEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchAssociateSecurityLevelToEntities 批量资产关联密级
//
// 批量资产关联密级：单个密级关联到多个资产上
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchAssociateSecurityLevelToEntities(request *model.BatchAssociateSecurityLevelToEntitiesRequest) (*model.BatchAssociateSecurityLevelToEntitiesResponse, error) {
	requestDef := GenReqDefForBatchAssociateSecurityLevelToEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAssociateSecurityLevelToEntitiesResponse), nil
	}
}

// BatchAssociateSecurityLevelToEntitiesInvoker 批量资产关联密级
func (c *DataArtsStudioClient) BatchAssociateSecurityLevelToEntitiesInvoker(request *model.BatchAssociateSecurityLevelToEntitiesRequest) *BatchAssociateSecurityLevelToEntitiesInvoker {
	requestDef := GenReqDefForBatchAssociateSecurityLevelToEntities()
	return &BatchAssociateSecurityLevelToEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecurityDataClassificationRule 批量删除识别规则接口
//
// 批量删除识别规则接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteSecurityDataClassificationRule(request *model.BatchDeleteSecurityDataClassificationRuleRequest) (*model.BatchDeleteSecurityDataClassificationRuleResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecurityDataClassificationRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecurityDataClassificationRuleResponse), nil
	}
}

// BatchDeleteSecurityDataClassificationRuleInvoker 批量删除识别规则接口
func (c *DataArtsStudioClient) BatchDeleteSecurityDataClassificationRuleInvoker(request *model.BatchDeleteSecurityDataClassificationRuleRequest) *BatchDeleteSecurityDataClassificationRuleInvoker {
	requestDef := GenReqDefForBatchDeleteSecurityDataClassificationRule()
	return &BatchDeleteSecurityDataClassificationRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecurityDynamicMaskingPolicies 批量删除动态脱敏策略
//
// 批量删除动态脱敏策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteSecurityDynamicMaskingPolicies(request *model.BatchDeleteSecurityDynamicMaskingPoliciesRequest) (*model.BatchDeleteSecurityDynamicMaskingPoliciesResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecurityDynamicMaskingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecurityDynamicMaskingPoliciesResponse), nil
	}
}

// BatchDeleteSecurityDynamicMaskingPoliciesInvoker 批量删除动态脱敏策略
func (c *DataArtsStudioClient) BatchDeleteSecurityDynamicMaskingPoliciesInvoker(request *model.BatchDeleteSecurityDynamicMaskingPoliciesRequest) *BatchDeleteSecurityDynamicMaskingPoliciesInvoker {
	requestDef := GenReqDefForBatchDeleteSecurityDynamicMaskingPolicies()
	return &BatchDeleteSecurityDynamicMaskingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecurityPermissionSetMembers 批量删除权限集成员
//
// 批量删除权限集成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteSecurityPermissionSetMembers(request *model.BatchDeleteSecurityPermissionSetMembersRequest) (*model.BatchDeleteSecurityPermissionSetMembersResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecurityPermissionSetMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecurityPermissionSetMembersResponse), nil
	}
}

// BatchDeleteSecurityPermissionSetMembersInvoker 批量删除权限集成员
func (c *DataArtsStudioClient) BatchDeleteSecurityPermissionSetMembersInvoker(request *model.BatchDeleteSecurityPermissionSetMembersRequest) *BatchDeleteSecurityPermissionSetMembersInvoker {
	requestDef := GenReqDefForBatchDeleteSecurityPermissionSetMembers()
	return &BatchDeleteSecurityPermissionSetMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecurityPermissionSetPermissions 删除权限集的权限
//
// 删除权限集的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteSecurityPermissionSetPermissions(request *model.BatchDeleteSecurityPermissionSetPermissionsRequest) (*model.BatchDeleteSecurityPermissionSetPermissionsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecurityPermissionSetPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecurityPermissionSetPermissionsResponse), nil
	}
}

// BatchDeleteSecurityPermissionSetPermissionsInvoker 删除权限集的权限
func (c *DataArtsStudioClient) BatchDeleteSecurityPermissionSetPermissionsInvoker(request *model.BatchDeleteSecurityPermissionSetPermissionsRequest) *BatchDeleteSecurityPermissionSetPermissionsInvoker {
	requestDef := GenReqDefForBatchDeleteSecurityPermissionSetPermissions()
	return &BatchDeleteSecurityPermissionSetPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteSecuritySecrecyLevels 批量删除密级
//
// 批量删除密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteSecuritySecrecyLevels(request *model.BatchDeleteSecuritySecrecyLevelsRequest) (*model.BatchDeleteSecuritySecrecyLevelsResponse, error) {
	requestDef := GenReqDefForBatchDeleteSecuritySecrecyLevels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteSecuritySecrecyLevelsResponse), nil
	}
}

// BatchDeleteSecuritySecrecyLevelsInvoker 批量删除密级
func (c *DataArtsStudioClient) BatchDeleteSecuritySecrecyLevelsInvoker(request *model.BatchDeleteSecuritySecrecyLevelsRequest) *BatchDeleteSecuritySecrecyLevelsInvoker {
	requestDef := GenReqDefForBatchDeleteSecuritySecrecyLevels()
	return &BatchDeleteSecuritySecrecyLevelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteTemplates 批量删除规则模板
//
// 批量删除规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchDeleteTemplates(request *model.BatchDeleteTemplatesRequest) (*model.BatchDeleteTemplatesResponse, error) {
	requestDef := GenReqDefForBatchDeleteTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteTemplatesResponse), nil
	}
}

// BatchDeleteTemplatesInvoker 批量删除规则模板
func (c *DataArtsStudioClient) BatchDeleteTemplatesInvoker(request *model.BatchDeleteTemplatesRequest) *BatchDeleteTemplatesInvoker {
	requestDef := GenReqDefForBatchDeleteTemplates()
	return &BatchDeleteTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchOffline 批量下线
//
// 批量下线。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchOffline(request *model.BatchOfflineRequest) (*model.BatchOfflineResponse, error) {
	requestDef := GenReqDefForBatchOffline()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchOfflineResponse), nil
	}
}

// BatchOfflineInvoker 批量下线
func (c *DataArtsStudioClient) BatchOfflineInvoker(request *model.BatchOfflineRequest) *BatchOfflineInvoker {
	requestDef := GenReqDefForBatchOffline()
	return &BatchOfflineInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchPublish 批量发布
//
// 批量发布。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchPublish(request *model.BatchPublishRequest) (*model.BatchPublishResponse, error) {
	requestDef := GenReqDefForBatchPublish()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchPublishResponse), nil
	}
}

// BatchPublishInvoker 批量发布
func (c *DataArtsStudioClient) BatchPublishInvoker(request *model.BatchPublishRequest) *BatchPublishInvoker {
	requestDef := GenReqDefForBatchPublish()
	return &BatchPublishInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchSyncMetadata 元数据实时同步接口(邀测)
//
// 元数据实时同步接口，支持批量。该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchSyncMetadata(request *model.BatchSyncMetadataRequest) (*model.BatchSyncMetadataResponse, error) {
	requestDef := GenReqDefForBatchSyncMetadata()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchSyncMetadataResponse), nil
	}
}

// BatchSyncMetadataInvoker 元数据实时同步接口(邀测)
func (c *DataArtsStudioClient) BatchSyncMetadataInvoker(request *model.BatchSyncMetadataRequest) *BatchSyncMetadataInvoker {
	requestDef := GenReqDefForBatchSyncMetadata()
	return &BatchSyncMetadataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateSecurityDlfDataWareHouses 批量更新数据开发连接细粒度认证状态
//
// 批量更新数据开发连接细粒度认证状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) BatchUpdateSecurityDlfDataWareHouses(request *model.BatchUpdateSecurityDlfDataWareHousesRequest) (*model.BatchUpdateSecurityDlfDataWareHousesResponse, error) {
	requestDef := GenReqDefForBatchUpdateSecurityDlfDataWareHouses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateSecurityDlfDataWareHousesResponse), nil
	}
}

// BatchUpdateSecurityDlfDataWareHousesInvoker 批量更新数据开发连接细粒度认证状态
func (c *DataArtsStudioClient) BatchUpdateSecurityDlfDataWareHousesInvoker(request *model.BatchUpdateSecurityDlfDataWareHousesRequest) *BatchUpdateSecurityDlfDataWareHousesInvoker {
	requestDef := GenReqDefForBatchUpdateSecurityDlfDataWareHouses()
	return &BatchUpdateSecurityDlfDataWareHousesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelFactoryPackages 撤销任务包
//
// 撤销任务包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CancelFactoryPackages(request *model.CancelFactoryPackagesRequest) (*model.CancelFactoryPackagesResponse, error) {
	requestDef := GenReqDefForCancelFactoryPackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelFactoryPackagesResponse), nil
	}
}

// CancelFactoryPackagesInvoker 撤销任务包
func (c *DataArtsStudioClient) CancelFactoryPackagesInvoker(request *model.CancelFactoryPackagesRequest) *CancelFactoryPackagesInvoker {
	requestDef := GenReqDefForCancelFactoryPackages()
	return &CancelFactoryPackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeCatalog 修改流程架构
//
// 修改流程架构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ChangeCatalog(request *model.ChangeCatalogRequest) (*model.ChangeCatalogResponse, error) {
	requestDef := GenReqDefForChangeCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeCatalogResponse), nil
	}
}

// ChangeCatalogInvoker 修改流程架构
func (c *DataArtsStudioClient) ChangeCatalogInvoker(request *model.ChangeCatalogRequest) *ChangeCatalogInvoker {
	requestDef := GenReqDefForChangeCatalog()
	return &ChangeCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeResource 规格变更接口
//
// 规格变更接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ChangeResource(request *model.ChangeResourceRequest) (*model.ChangeResourceResponse, error) {
	requestDef := GenReqDefForChangeResource()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeResourceResponse), nil
	}
}

// ChangeResourceInvoker 规格变更接口
func (c *DataArtsStudioClient) ChangeResourceInvoker(request *model.ChangeResourceRequest) *ChangeResourceInvoker {
	requestDef := GenReqDefForChangeResource()
	return &ChangeResourceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeSubjects 修改或删除主题层级
//
// 修改或删除主题层级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ChangeSubjects(request *model.ChangeSubjectsRequest) (*model.ChangeSubjectsResponse, error) {
	requestDef := GenReqDefForChangeSubjects()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeSubjectsResponse), nil
	}
}

// ChangeSubjectsInvoker 修改或删除主题层级
func (c *DataArtsStudioClient) ChangeSubjectsInvoker(request *model.ChangeSubjectsRequest) *ChangeSubjectsInvoker {
	requestDef := GenReqDefForChangeSubjects()
	return &ChangeSubjectsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckDimensionStatus 查看逆向维度表任务
//
// 查看逆向维度表任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CheckDimensionStatus(request *model.CheckDimensionStatusRequest) (*model.CheckDimensionStatusResponse, error) {
	requestDef := GenReqDefForCheckDimensionStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckDimensionStatusResponse), nil
	}
}

// CheckDimensionStatusInvoker 查看逆向维度表任务
func (c *DataArtsStudioClient) CheckDimensionStatusInvoker(request *model.CheckDimensionStatusRequest) *CheckDimensionStatusInvoker {
	requestDef := GenReqDefForCheckDimensionStatus()
	return &CheckDimensionStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CheckFactLogicTableStatus 查看逆向事实表任务
//
// 查看逆向事实表任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CheckFactLogicTableStatus(request *model.CheckFactLogicTableStatusRequest) (*model.CheckFactLogicTableStatusResponse, error) {
	requestDef := GenReqDefForCheckFactLogicTableStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CheckFactLogicTableStatusResponse), nil
	}
}

// CheckFactLogicTableStatusInvoker 查看逆向事实表任务
func (c *DataArtsStudioClient) CheckFactLogicTableStatusInvoker(request *model.CheckFactLogicTableStatusRequest) *CheckFactLogicTableStatusInvoker {
	requestDef := GenReqDefForCheckFactLogicTableStatus()
	return &CheckFactLogicTableStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmApprovals 审批单处理
//
// 审批驳回/通过，单个或多个action-id&#x3D;reject/resolve。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ConfirmApprovals(request *model.ConfirmApprovalsRequest) (*model.ConfirmApprovalsResponse, error) {
	requestDef := GenReqDefForConfirmApprovals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmApprovalsResponse), nil
	}
}

// ConfirmApprovalsInvoker 审批单处理
func (c *DataArtsStudioClient) ConfirmApprovalsInvoker(request *model.ConfirmApprovalsRequest) *ConfirmApprovalsInvoker {
	requestDef := GenReqDefForConfirmApprovals()
	return &ConfirmApprovalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ConfirmMessage 处理消息
//
// 对收到的通知消息进行确认，可以在指定的时间范围内选择何时进行处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ConfirmMessage(request *model.ConfirmMessageRequest) (*model.ConfirmMessageResponse, error) {
	requestDef := GenReqDefForConfirmMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ConfirmMessageResponse), nil
	}
}

// ConfirmMessageInvoker 处理消息
func (c *DataArtsStudioClient) ConfirmMessageInvoker(request *model.ConfirmMessageRequest) *ConfirmMessageInvoker {
	requestDef := GenReqDefForConfirmMessage()
	return &ConfirmMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountAllModels 关系建模统计信息
//
// 关系建模页面，外层的统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CountAllModels(request *model.CountAllModelsRequest) (*model.CountAllModelsResponse, error) {
	requestDef := GenReqDefForCountAllModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountAllModelsResponse), nil
	}
}

// CountAllModelsInvoker 关系建模统计信息
func (c *DataArtsStudioClient) CountAllModelsInvoker(request *model.CountAllModelsRequest) *CountAllModelsInvoker {
	requestDef := GenReqDefForCountAllModels()
	return &CountAllModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountOverviews 总览统计信息
//
// 总览统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CountOverviews(request *model.CountOverviewsRequest) (*model.CountOverviewsResponse, error) {
	requestDef := GenReqDefForCountOverviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountOverviewsResponse), nil
	}
}

// CountOverviewsInvoker 总览统计信息
func (c *DataArtsStudioClient) CountOverviewsInvoker(request *model.CountOverviewsRequest) *CountOverviewsInvoker {
	requestDef := GenReqDefForCountOverviews()
	return &CountOverviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountStandards 标准覆盖率统计信息
//
// 查看某个数据标准在所有模型字段中的覆盖率，即使用该标准的字段占总字段的百分比。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CountStandards(request *model.CountStandardsRequest) (*model.CountStandardsResponse, error) {
	requestDef := GenReqDefForCountStandards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountStandardsResponse), nil
	}
}

// CountStandardsInvoker 标准覆盖率统计信息
func (c *DataArtsStudioClient) CountStandardsInvoker(request *model.CountStandardsRequest) *CountStandardsInvoker {
	requestDef := GenReqDefForCountStandards()
	return &CountStandardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountTableModels 模型统计信息
//
// 单个模型中的统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CountTableModels(request *model.CountTableModelsRequest) (*model.CountTableModelsResponse, error) {
	requestDef := GenReqDefForCountTableModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountTableModelsResponse), nil
	}
}

// CountTableModelsInvoker 模型统计信息
func (c *DataArtsStudioClient) CountTableModelsInvoker(request *model.CountTableModelsRequest) *CountTableModelsInvoker {
	requestDef := GenReqDefForCountTableModels()
	return &CountTableModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建应用
//
// 创建应用。
// 支持创建APP， IAM应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用
func (c *DataArtsStudioClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApprover 创建审批人
//
// 创建审批人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateApprover(request *model.CreateApproverRequest) (*model.CreateApproverResponse, error) {
	requestDef := GenReqDefForCreateApprover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApproverResponse), nil
	}
}

// CreateApproverInvoker 创建审批人
func (c *DataArtsStudioClient) CreateApproverInvoker(request *model.CreateApproverRequest) *CreateApproverInvoker {
	requestDef := GenReqDefForCreateApprover()
	return &CreateApproverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBizMetric 创建业务指标
//
// 创建业务指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateBizMetric(request *model.CreateBizMetricRequest) (*model.CreateBizMetricResponse, error) {
	requestDef := GenReqDefForCreateBizMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBizMetricResponse), nil
	}
}

// CreateBizMetricInvoker 创建业务指标
func (c *DataArtsStudioClient) CreateBizMetricInvoker(request *model.CreateBizMetricRequest) *CreateBizMetricInvoker {
	requestDef := GenReqDefForCreateBizMetric()
	return &CreateBizMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCatalog 创建流程架构
//
// 创建流程架构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateCatalog(request *model.CreateCatalogRequest) (*model.CreateCatalogResponse, error) {
	requestDef := GenReqDefForCreateCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCatalogResponse), nil
	}
}

// CreateCatalogInvoker 创建流程架构
func (c *DataArtsStudioClient) CreateCatalogInvoker(request *model.CreateCatalogRequest) *CreateCatalogInvoker {
	requestDef := GenReqDefForCreateCatalog()
	return &CreateCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCodeTable 创建码表
//
// 创建码表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateCodeTable(request *model.CreateCodeTableRequest) (*model.CreateCodeTableResponse, error) {
	requestDef := GenReqDefForCreateCodeTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCodeTableResponse), nil
	}
}

// CreateCodeTableInvoker 创建码表
func (c *DataArtsStudioClient) CreateCodeTableInvoker(request *model.CreateCodeTableRequest) *CreateCodeTableInvoker {
	requestDef := GenReqDefForCreateCodeTable()
	return &CreateCodeTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateConnections 创建数据连接
//
// 创建数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateConnections(request *model.CreateConnectionsRequest) (*model.CreateConnectionsResponse, error) {
	requestDef := GenReqDefForCreateConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateConnectionsResponse), nil
	}
}

// CreateConnectionsInvoker 创建数据连接
func (c *DataArtsStudioClient) CreateConnectionsInvoker(request *model.CreateConnectionsRequest) *CreateConnectionsInvoker {
	requestDef := GenReqDefForCreateConnections()
	return &CreateConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesignAggregationLogicTable 新建汇总表
//
// 根据入参，手动创建汇总表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateDesignAggregationLogicTable(request *model.CreateDesignAggregationLogicTableRequest) (*model.CreateDesignAggregationLogicTableResponse, error) {
	requestDef := GenReqDefForCreateDesignAggregationLogicTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesignAggregationLogicTableResponse), nil
	}
}

// CreateDesignAggregationLogicTableInvoker 新建汇总表
func (c *DataArtsStudioClient) CreateDesignAggregationLogicTableInvoker(request *model.CreateDesignAggregationLogicTableRequest) *CreateDesignAggregationLogicTableInvoker {
	requestDef := GenReqDefForCreateDesignAggregationLogicTable()
	return &CreateDesignAggregationLogicTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDesignAtomicIndex 新建原子指标
//
// 新建单个原子指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateDesignAtomicIndex(request *model.CreateDesignAtomicIndexRequest) (*model.CreateDesignAtomicIndexResponse, error) {
	requestDef := GenReqDefForCreateDesignAtomicIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDesignAtomicIndexResponse), nil
	}
}

// CreateDesignAtomicIndexInvoker 新建原子指标
func (c *DataArtsStudioClient) CreateDesignAtomicIndexInvoker(request *model.CreateDesignAtomicIndexRequest) *CreateDesignAtomicIndexInvoker {
	requestDef := GenReqDefForCreateDesignAtomicIndex()
	return &CreateDesignAtomicIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateDirectory 创建目录
//
// 创建目录（数据标准、码表）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateDirectory(request *model.CreateDirectoryRequest) (*model.CreateDirectoryResponse, error) {
	requestDef := GenReqDefForCreateDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDirectoryResponse), nil
	}
}

// CreateDirectoryInvoker 创建目录
func (c *DataArtsStudioClient) CreateDirectoryInvoker(request *model.CreateDirectoryRequest) *CreateDirectoryInvoker {
	requestDef := GenReqDefForCreateDirectory()
	return &CreateDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFactoryEnv 创建环境变量
//
// 创建环境变量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateFactoryEnv(request *model.CreateFactoryEnvRequest) (*model.CreateFactoryEnvResponse, error) {
	requestDef := GenReqDefForCreateFactoryEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFactoryEnvResponse), nil
	}
}

// CreateFactoryEnvInvoker 创建环境变量
func (c *DataArtsStudioClient) CreateFactoryEnvInvoker(request *model.CreateFactoryEnvRequest) *CreateFactoryEnvInvoker {
	requestDef := GenReqDefForCreateFactoryEnv()
	return &CreateFactoryEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFactoryJob 创建作业
//
// 创建作业
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateFactoryJob(request *model.CreateFactoryJobRequest) (*model.CreateFactoryJobResponse, error) {
	requestDef := GenReqDefForCreateFactoryJob()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFactoryJobResponse), nil
	}
}

// CreateFactoryJobInvoker 创建作业
func (c *DataArtsStudioClient) CreateFactoryJobInvoker(request *model.CreateFactoryJobRequest) *CreateFactoryJobInvoker {
	requestDef := GenReqDefForCreateFactoryJob()
	return &CreateFactoryJobInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateFactorySupplementDataInstance 创建补数据实例
//
// 创建补数据实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateFactorySupplementDataInstance(request *model.CreateFactorySupplementDataInstanceRequest) (*model.CreateFactorySupplementDataInstanceResponse, error) {
	requestDef := GenReqDefForCreateFactorySupplementDataInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateFactorySupplementDataInstanceResponse), nil
	}
}

// CreateFactorySupplementDataInstanceInvoker 创建补数据实例
func (c *DataArtsStudioClient) CreateFactorySupplementDataInstanceInvoker(request *model.CreateFactorySupplementDataInstanceRequest) *CreateFactorySupplementDataInstanceInvoker {
	requestDef := GenReqDefForCreateFactorySupplementDataInstance()
	return &CreateFactorySupplementDataInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateManagerWorkSpace 创建工作空间
//
// 创建工作空间
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateManagerWorkSpace(request *model.CreateManagerWorkSpaceRequest) (*model.CreateManagerWorkSpaceResponse, error) {
	requestDef := GenReqDefForCreateManagerWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateManagerWorkSpaceResponse), nil
	}
}

// CreateManagerWorkSpaceInvoker 创建工作空间
func (c *DataArtsStudioClient) CreateManagerWorkSpaceInvoker(request *model.CreateManagerWorkSpaceRequest) *CreateManagerWorkSpaceInvoker {
	requestDef := GenReqDefForCreateManagerWorkSpace()
	return &CreateManagerWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateAsset 添加或修改资产
//
// 添加或修改资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateOrUpdateAsset(request *model.CreateOrUpdateAssetRequest) (*model.CreateOrUpdateAssetResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateAssetResponse), nil
	}
}

// CreateOrUpdateAssetInvoker 添加或修改资产
func (c *DataArtsStudioClient) CreateOrUpdateAssetInvoker(request *model.CreateOrUpdateAssetRequest) *CreateOrUpdateAssetInvoker {
	requestDef := GenReqDefForCreateOrUpdateAsset()
	return &CreateOrUpdateAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateOrUpdateEntities 创建或修改资产(邀测)
//
// 创建或修改资产，该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateOrUpdateEntities(request *model.CreateOrUpdateEntitiesRequest) (*model.CreateOrUpdateEntitiesResponse, error) {
	requestDef := GenReqDefForCreateOrUpdateEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateOrUpdateEntitiesResponse), nil
	}
}

// CreateOrUpdateEntitiesInvoker 创建或修改资产(邀测)
func (c *DataArtsStudioClient) CreateOrUpdateEntitiesInvoker(request *model.CreateOrUpdateEntitiesRequest) *CreateOrUpdateEntitiesInvoker {
	requestDef := GenReqDefForCreateOrUpdateEntities()
	return &CreateOrUpdateEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityAssignedQueue 分配队列资源给指定空间
//
// 分配队列资源给指定空间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityAssignedQueue(request *model.CreateSecurityAssignedQueueRequest) (*model.CreateSecurityAssignedQueueResponse, error) {
	requestDef := GenReqDefForCreateSecurityAssignedQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityAssignedQueueResponse), nil
	}
}

// CreateSecurityAssignedQueueInvoker 分配队列资源给指定空间
func (c *DataArtsStudioClient) CreateSecurityAssignedQueueInvoker(request *model.CreateSecurityAssignedQueueRequest) *CreateSecurityAssignedQueueInvoker {
	requestDef := GenReqDefForCreateSecurityAssignedQueue()
	return &CreateSecurityAssignedQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityDataClassificationRule 创建识别规则
//
// 创建识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityDataClassificationRule(request *model.CreateSecurityDataClassificationRuleRequest) (*model.CreateSecurityDataClassificationRuleResponse, error) {
	requestDef := GenReqDefForCreateSecurityDataClassificationRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityDataClassificationRuleResponse), nil
	}
}

// CreateSecurityDataClassificationRuleInvoker 创建识别规则
func (c *DataArtsStudioClient) CreateSecurityDataClassificationRuleInvoker(request *model.CreateSecurityDataClassificationRuleRequest) *CreateSecurityDataClassificationRuleInvoker {
	requestDef := GenReqDefForCreateSecurityDataClassificationRule()
	return &CreateSecurityDataClassificationRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityDynamicMaskingPolicy 创建数据脱敏策略
//
// 创建动态数据脱敏策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityDynamicMaskingPolicy(request *model.CreateSecurityDynamicMaskingPolicyRequest) (*model.CreateSecurityDynamicMaskingPolicyResponse, error) {
	requestDef := GenReqDefForCreateSecurityDynamicMaskingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityDynamicMaskingPolicyResponse), nil
	}
}

// CreateSecurityDynamicMaskingPolicyInvoker 创建数据脱敏策略
func (c *DataArtsStudioClient) CreateSecurityDynamicMaskingPolicyInvoker(request *model.CreateSecurityDynamicMaskingPolicyRequest) *CreateSecurityDynamicMaskingPolicyInvoker {
	requestDef := GenReqDefForCreateSecurityDynamicMaskingPolicy()
	return &CreateSecurityDynamicMaskingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityPermissionSet 创建权限集
//
// 创建权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityPermissionSet(request *model.CreateSecurityPermissionSetRequest) (*model.CreateSecurityPermissionSetResponse, error) {
	requestDef := GenReqDefForCreateSecurityPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPermissionSetResponse), nil
	}
}

// CreateSecurityPermissionSetInvoker 创建权限集
func (c *DataArtsStudioClient) CreateSecurityPermissionSetInvoker(request *model.CreateSecurityPermissionSetRequest) *CreateSecurityPermissionSetInvoker {
	requestDef := GenReqDefForCreateSecurityPermissionSet()
	return &CreateSecurityPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityPermissionSetMember 添加权限集成员
//
// 添加权限集成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityPermissionSetMember(request *model.CreateSecurityPermissionSetMemberRequest) (*model.CreateSecurityPermissionSetMemberResponse, error) {
	requestDef := GenReqDefForCreateSecurityPermissionSetMember()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPermissionSetMemberResponse), nil
	}
}

// CreateSecurityPermissionSetMemberInvoker 添加权限集成员
func (c *DataArtsStudioClient) CreateSecurityPermissionSetMemberInvoker(request *model.CreateSecurityPermissionSetMemberRequest) *CreateSecurityPermissionSetMemberInvoker {
	requestDef := GenReqDefForCreateSecurityPermissionSetMember()
	return &CreateSecurityPermissionSetMemberInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecurityPermissionSetPermission 添加权限集的权限
//
// 添加权限集的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecurityPermissionSetPermission(request *model.CreateSecurityPermissionSetPermissionRequest) (*model.CreateSecurityPermissionSetPermissionResponse, error) {
	requestDef := GenReqDefForCreateSecurityPermissionSetPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecurityPermissionSetPermissionResponse), nil
	}
}

// CreateSecurityPermissionSetPermissionInvoker 添加权限集的权限
func (c *DataArtsStudioClient) CreateSecurityPermissionSetPermissionInvoker(request *model.CreateSecurityPermissionSetPermissionRequest) *CreateSecurityPermissionSetPermissionInvoker {
	requestDef := GenReqDefForCreateSecurityPermissionSetPermission()
	return &CreateSecurityPermissionSetPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSecuritySecrecyLevel 创建密级
//
// 创建密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSecuritySecrecyLevel(request *model.CreateSecuritySecrecyLevelRequest) (*model.CreateSecuritySecrecyLevelResponse, error) {
	requestDef := GenReqDefForCreateSecuritySecrecyLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSecuritySecrecyLevelResponse), nil
	}
}

// CreateSecuritySecrecyLevelInvoker 创建密级
func (c *DataArtsStudioClient) CreateSecuritySecrecyLevelInvoker(request *model.CreateSecuritySecrecyLevelRequest) *CreateSecuritySecrecyLevelInvoker {
	requestDef := GenReqDefForCreateSecuritySecrecyLevel()
	return &CreateSecuritySecrecyLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateServiceCatalog 创建服务目录
//
// 创建服务目录。 根目录编号为0。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateServiceCatalog(request *model.CreateServiceCatalogRequest) (*model.CreateServiceCatalogResponse, error) {
	requestDef := GenReqDefForCreateServiceCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateServiceCatalogResponse), nil
	}
}

// CreateServiceCatalogInvoker 创建服务目录
func (c *DataArtsStudioClient) CreateServiceCatalogInvoker(request *model.CreateServiceCatalogRequest) *CreateServiceCatalogInvoker {
	requestDef := GenReqDefForCreateServiceCatalog()
	return &CreateServiceCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStandard 创建数据标准
//
// 创建数据标准。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateStandard(request *model.CreateStandardRequest) (*model.CreateStandardResponse, error) {
	requestDef := GenReqDefForCreateStandard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStandardResponse), nil
	}
}

// CreateStandardInvoker 创建数据标准
func (c *DataArtsStudioClient) CreateStandardInvoker(request *model.CreateStandardRequest) *CreateStandardInvoker {
	requestDef := GenReqDefForCreateStandard()
	return &CreateStandardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateStandardTemplate 创建数据标准模板
//
// 创建当前工作空间下的数据标准模板自定义项。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateStandardTemplate(request *model.CreateStandardTemplateRequest) (*model.CreateStandardTemplateResponse, error) {
	requestDef := GenReqDefForCreateStandardTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateStandardTemplateResponse), nil
	}
}

// CreateStandardTemplateInvoker 创建数据标准模板
func (c *DataArtsStudioClient) CreateStandardTemplateInvoker(request *model.CreateStandardTemplateRequest) *CreateStandardTemplateInvoker {
	requestDef := GenReqDefForCreateStandardTemplate()
	return &CreateStandardTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubject 创建主题
//
// 创建主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSubject(request *model.CreateSubjectRequest) (*model.CreateSubjectResponse, error) {
	requestDef := GenReqDefForCreateSubject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubjectResponse), nil
	}
}

// CreateSubjectInvoker 创建主题
func (c *DataArtsStudioClient) CreateSubjectInvoker(request *model.CreateSubjectRequest) *CreateSubjectInvoker {
	requestDef := GenReqDefForCreateSubject()
	return &CreateSubjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateSubjectNew 创建主题(新)
//
// 创建主题(新)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateSubjectNew(request *model.CreateSubjectNewRequest) (*model.CreateSubjectNewResponse, error) {
	requestDef := GenReqDefForCreateSubjectNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSubjectNewResponse), nil
	}
}

// CreateSubjectNewInvoker 创建主题(新)
func (c *DataArtsStudioClient) CreateSubjectNewInvoker(request *model.CreateSubjectNewRequest) *CreateSubjectNewInvoker {
	requestDef := GenReqDefForCreateSubjectNew()
	return &CreateSubjectNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTableModel 创建表模型
//
// 在关系建模中创建一个表模型，包括逻辑实体和物理表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateTableModel(request *model.CreateTableModelRequest) (*model.CreateTableModelResponse, error) {
	requestDef := GenReqDefForCreateTableModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTableModelResponse), nil
	}
}

// CreateTableModelInvoker 创建表模型
func (c *DataArtsStudioClient) CreateTableModelInvoker(request *model.CreateTableModelRequest) *CreateTableModelInvoker {
	requestDef := GenReqDefForCreateTableModel()
	return &CreateTableModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTask 创建采集任务
//
// 创建采集任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateTask(request *model.CreateTaskRequest) (*model.CreateTaskResponse, error) {
	requestDef := GenReqDefForCreateTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTaskResponse), nil
	}
}

// CreateTaskInvoker 创建采集任务
func (c *DataArtsStudioClient) CreateTaskInvoker(request *model.CreateTaskRequest) *CreateTaskInvoker {
	requestDef := GenReqDefForCreateTask()
	return &CreateTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTemplate 创建规则模板
//
// 创建规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateTemplate(request *model.CreateTemplateRequest) (*model.CreateTemplateResponse, error) {
	requestDef := GenReqDefForCreateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTemplateResponse), nil
	}
}

// CreateTemplateInvoker 创建规则模板
func (c *DataArtsStudioClient) CreateTemplateInvoker(request *model.CreateTemplateRequest) *CreateTemplateInvoker {
	requestDef := GenReqDefForCreateTemplate()
	return &CreateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateWorkspace 新建模型工作区
//
// 新建模型工作区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateWorkspace(request *model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	requestDef := GenReqDefForCreateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateWorkspaceResponse), nil
	}
}

// CreateWorkspaceInvoker 新建模型工作区
func (c *DataArtsStudioClient) CreateWorkspaceInvoker(request *model.CreateWorkspaceRequest) *CreateWorkspaceInvoker {
	requestDef := GenReqDefForCreateWorkspace()
	return &CreateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugDataconnection 测试创建数据连接
//
// 测试创建数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DebugDataconnection(request *model.DebugDataconnectionRequest) (*model.DebugDataconnectionResponse, error) {
	requestDef := GenReqDefForDebugDataconnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugDataconnectionResponse), nil
	}
}

// DebugDataconnectionInvoker 测试创建数据连接
func (c *DataArtsStudioClient) DebugDataconnectionInvoker(request *model.DebugDataconnectionRequest) *DebugDataconnectionInvoker {
	requestDef := GenReqDefForDebugDataconnection()
	return &DebugDataconnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugSecurityDlfDataWareHouses 测试数据开发连接细粒度连通性
//
// 测试数据开发连接细粒度连通性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DebugSecurityDlfDataWareHouses(request *model.DebugSecurityDlfDataWareHousesRequest) (*model.DebugSecurityDlfDataWareHousesResponse, error) {
	requestDef := GenReqDefForDebugSecurityDlfDataWareHouses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugSecurityDlfDataWareHousesResponse), nil
	}
}

// DebugSecurityDlfDataWareHousesInvoker 测试数据开发连接细粒度连通性
func (c *DataArtsStudioClient) DebugSecurityDlfDataWareHousesInvoker(request *model.DebugSecurityDlfDataWareHousesRequest) *DebugSecurityDlfDataWareHousesInvoker {
	requestDef := GenReqDefForDebugSecurityDlfDataWareHouses()
	return &DebugSecurityDlfDataWareHousesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *DataArtsStudioClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApprover 删除审批人
//
// 删除审批人。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteApprover(request *model.DeleteApproverRequest) (*model.DeleteApproverResponse, error) {
	requestDef := GenReqDefForDeleteApprover()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApproverResponse), nil
	}
}

// DeleteApproverInvoker 删除审批人
func (c *DataArtsStudioClient) DeleteApproverInvoker(request *model.DeleteApproverRequest) *DeleteApproverInvoker {
	requestDef := GenReqDefForDeleteApprover()
	return &DeleteApproverInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAsset 删除资产
//
// 删除资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteAsset(request *model.DeleteAssetRequest) (*model.DeleteAssetResponse, error) {
	requestDef := GenReqDefForDeleteAsset()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAssetResponse), nil
	}
}

// DeleteAssetInvoker 删除资产
func (c *DataArtsStudioClient) DeleteAssetInvoker(request *model.DeleteAssetRequest) *DeleteAssetInvoker {
	requestDef := GenReqDefForDeleteAsset()
	return &DeleteAssetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBizMetric 删除业务指标
//
// 删除业务指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteBizMetric(request *model.DeleteBizMetricRequest) (*model.DeleteBizMetricResponse, error) {
	requestDef := GenReqDefForDeleteBizMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBizMetricResponse), nil
	}
}

// DeleteBizMetricInvoker 删除业务指标
func (c *DataArtsStudioClient) DeleteBizMetricInvoker(request *model.DeleteBizMetricRequest) *DeleteBizMetricInvoker {
	requestDef := GenReqDefForDeleteBizMetric()
	return &DeleteBizMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCatalog 删除流程架构
//
// 删除流程架构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteCatalog(request *model.DeleteCatalogRequest) (*model.DeleteCatalogResponse, error) {
	requestDef := GenReqDefForDeleteCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCatalogResponse), nil
	}
}

// DeleteCatalogInvoker 删除流程架构
func (c *DataArtsStudioClient) DeleteCatalogInvoker(request *model.DeleteCatalogRequest) *DeleteCatalogInvoker {
	requestDef := GenReqDefForDeleteCatalog()
	return &DeleteCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteClassificationFromEntities 移除资产关联的分类
//
// 移除资产关联分类：
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteClassificationFromEntities(request *model.DeleteClassificationFromEntitiesRequest) (*model.DeleteClassificationFromEntitiesResponse, error) {
	requestDef := GenReqDefForDeleteClassificationFromEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClassificationFromEntitiesResponse), nil
	}
}

// DeleteClassificationFromEntitiesInvoker 移除资产关联的分类
func (c *DataArtsStudioClient) DeleteClassificationFromEntitiesInvoker(request *model.DeleteClassificationFromEntitiesRequest) *DeleteClassificationFromEntitiesInvoker {
	requestDef := GenReqDefForDeleteClassificationFromEntities()
	return &DeleteClassificationFromEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCodeTable 删除码表
//
// 删除码表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteCodeTable(request *model.DeleteCodeTableRequest) (*model.DeleteCodeTableResponse, error) {
	requestDef := GenReqDefForDeleteCodeTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCodeTableResponse), nil
	}
}

// DeleteCodeTableInvoker 删除码表
func (c *DataArtsStudioClient) DeleteCodeTableInvoker(request *model.DeleteCodeTableRequest) *DeleteCodeTableInvoker {
	requestDef := GenReqDefForDeleteCodeTable()
	return &DeleteCodeTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDataconnection 删除数据连接
//
// 删除数据连接
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteDataconnection(request *model.DeleteDataconnectionRequest) (*model.DeleteDataconnectionResponse, error) {
	requestDef := GenReqDefForDeleteDataconnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDataconnectionResponse), nil
	}
}

// DeleteDataconnectionInvoker 删除数据连接
func (c *DataArtsStudioClient) DeleteDataconnectionInvoker(request *model.DeleteDataconnectionRequest) *DeleteDataconnectionInvoker {
	requestDef := GenReqDefForDeleteDataconnection()
	return &DeleteDataconnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesignAggregationLogicTable 删除汇总表
//
// 批量删除汇总表，只能删除状态为草稿、已线下、已驳回的表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteDesignAggregationLogicTable(request *model.DeleteDesignAggregationLogicTableRequest) (*model.DeleteDesignAggregationLogicTableResponse, error) {
	requestDef := GenReqDefForDeleteDesignAggregationLogicTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesignAggregationLogicTableResponse), nil
	}
}

// DeleteDesignAggregationLogicTableInvoker 删除汇总表
func (c *DataArtsStudioClient) DeleteDesignAggregationLogicTableInvoker(request *model.DeleteDesignAggregationLogicTableRequest) *DeleteDesignAggregationLogicTableInvoker {
	requestDef := GenReqDefForDeleteDesignAggregationLogicTable()
	return &DeleteDesignAggregationLogicTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesignAtomicIndex 删除原子指标
//
// 批量删除原子指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteDesignAtomicIndex(request *model.DeleteDesignAtomicIndexRequest) (*model.DeleteDesignAtomicIndexResponse, error) {
	requestDef := GenReqDefForDeleteDesignAtomicIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesignAtomicIndexResponse), nil
	}
}

// DeleteDesignAtomicIndexInvoker 删除原子指标
func (c *DataArtsStudioClient) DeleteDesignAtomicIndexInvoker(request *model.DeleteDesignAtomicIndexRequest) *DeleteDesignAtomicIndexInvoker {
	requestDef := GenReqDefForDeleteDesignAtomicIndex()
	return &DeleteDesignAtomicIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDesignLatestApproval 删除实体下展
//
// 当已发布的实体被编辑时，其会生成下展，该接口用于删除实体的下展信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteDesignLatestApproval(request *model.DeleteDesignLatestApprovalRequest) (*model.DeleteDesignLatestApprovalResponse, error) {
	requestDef := GenReqDefForDeleteDesignLatestApproval()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDesignLatestApprovalResponse), nil
	}
}

// DeleteDesignLatestApprovalInvoker 删除实体下展
func (c *DataArtsStudioClient) DeleteDesignLatestApprovalInvoker(request *model.DeleteDesignLatestApprovalRequest) *DeleteDesignLatestApprovalInvoker {
	requestDef := GenReqDefForDeleteDesignLatestApproval()
	return &DeleteDesignLatestApprovalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDirectory 删除目录
//
// 删除目录（数据标准、码表）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteDirectory(request *model.DeleteDirectoryRequest) (*model.DeleteDirectoryResponse, error) {
	requestDef := GenReqDefForDeleteDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDirectoryResponse), nil
	}
}

// DeleteDirectoryInvoker 删除目录
func (c *DataArtsStudioClient) DeleteDirectoryInvoker(request *model.DeleteDirectoryRequest) *DeleteDirectoryInvoker {
	requestDef := GenReqDefForDeleteDirectory()
	return &DeleteDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityAssignedQueue 删除当前空间下分配的队列资源
//
// 删除当前空间下分配的队列资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSecurityAssignedQueue(request *model.DeleteSecurityAssignedQueueRequest) (*model.DeleteSecurityAssignedQueueResponse, error) {
	requestDef := GenReqDefForDeleteSecurityAssignedQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityAssignedQueueResponse), nil
	}
}

// DeleteSecurityAssignedQueueInvoker 删除当前空间下分配的队列资源
func (c *DataArtsStudioClient) DeleteSecurityAssignedQueueInvoker(request *model.DeleteSecurityAssignedQueueRequest) *DeleteSecurityAssignedQueueInvoker {
	requestDef := GenReqDefForDeleteSecurityAssignedQueue()
	return &DeleteSecurityAssignedQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityDataClassificationRule 删除识别规则
//
// 删除识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSecurityDataClassificationRule(request *model.DeleteSecurityDataClassificationRuleRequest) (*model.DeleteSecurityDataClassificationRuleResponse, error) {
	requestDef := GenReqDefForDeleteSecurityDataClassificationRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityDataClassificationRuleResponse), nil
	}
}

// DeleteSecurityDataClassificationRuleInvoker 删除识别规则
func (c *DataArtsStudioClient) DeleteSecurityDataClassificationRuleInvoker(request *model.DeleteSecurityDataClassificationRuleRequest) *DeleteSecurityDataClassificationRuleInvoker {
	requestDef := GenReqDefForDeleteSecurityDataClassificationRule()
	return &DeleteSecurityDataClassificationRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityLevelFromEntity 移除资产关联密级
//
// 移除资产关联密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSecurityLevelFromEntity(request *model.DeleteSecurityLevelFromEntityRequest) (*model.DeleteSecurityLevelFromEntityResponse, error) {
	requestDef := GenReqDefForDeleteSecurityLevelFromEntity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityLevelFromEntityResponse), nil
	}
}

// DeleteSecurityLevelFromEntityInvoker 移除资产关联密级
func (c *DataArtsStudioClient) DeleteSecurityLevelFromEntityInvoker(request *model.DeleteSecurityLevelFromEntityRequest) *DeleteSecurityLevelFromEntityInvoker {
	requestDef := GenReqDefForDeleteSecurityLevelFromEntity()
	return &DeleteSecurityLevelFromEntityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecurityPermissionSet 删除权限集
//
// 删除权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSecurityPermissionSet(request *model.DeleteSecurityPermissionSetRequest) (*model.DeleteSecurityPermissionSetResponse, error) {
	requestDef := GenReqDefForDeleteSecurityPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecurityPermissionSetResponse), nil
	}
}

// DeleteSecurityPermissionSetInvoker 删除权限集
func (c *DataArtsStudioClient) DeleteSecurityPermissionSetInvoker(request *model.DeleteSecurityPermissionSetRequest) *DeleteSecurityPermissionSetInvoker {
	requestDef := GenReqDefForDeleteSecurityPermissionSet()
	return &DeleteSecurityPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSecuritySecrecyLevel 删除指定的id的密级
//
// 删除指定的id的密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSecuritySecrecyLevel(request *model.DeleteSecuritySecrecyLevelRequest) (*model.DeleteSecuritySecrecyLevelResponse, error) {
	requestDef := GenReqDefForDeleteSecuritySecrecyLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSecuritySecrecyLevelResponse), nil
	}
}

// DeleteSecuritySecrecyLevelInvoker 删除指定的id的密级
func (c *DataArtsStudioClient) DeleteSecuritySecrecyLevelInvoker(request *model.DeleteSecuritySecrecyLevelRequest) *DeleteSecuritySecrecyLevelInvoker {
	requestDef := GenReqDefForDeleteSecuritySecrecyLevel()
	return &DeleteSecuritySecrecyLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceCatalog 批量删除目录
//
// 批量删除服务目录。
// 删除目录的同时会删除其下的所有子目录，不支持删除带有API的目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteServiceCatalog(request *model.DeleteServiceCatalogRequest) (*model.DeleteServiceCatalogResponse, error) {
	requestDef := GenReqDefForDeleteServiceCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceCatalogResponse), nil
	}
}

// DeleteServiceCatalogInvoker 批量删除目录
func (c *DataArtsStudioClient) DeleteServiceCatalogInvoker(request *model.DeleteServiceCatalogRequest) *DeleteServiceCatalogInvoker {
	requestDef := GenReqDefForDeleteServiceCatalog()
	return &DeleteServiceCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStandard 删除数据标准
//
// 删除数据标准。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteStandard(request *model.DeleteStandardRequest) (*model.DeleteStandardResponse, error) {
	requestDef := GenReqDefForDeleteStandard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStandardResponse), nil
	}
}

// DeleteStandardInvoker 删除数据标准
func (c *DataArtsStudioClient) DeleteStandardInvoker(request *model.DeleteStandardRequest) *DeleteStandardInvoker {
	requestDef := GenReqDefForDeleteStandard()
	return &DeleteStandardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteStandardTemplate 删除数据标准模板
//
// 删除数据标准模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteStandardTemplate(request *model.DeleteStandardTemplateRequest) (*model.DeleteStandardTemplateResponse, error) {
	requestDef := GenReqDefForDeleteStandardTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteStandardTemplateResponse), nil
	}
}

// DeleteStandardTemplateInvoker 删除数据标准模板
func (c *DataArtsStudioClient) DeleteStandardTemplateInvoker(request *model.DeleteStandardTemplateRequest) *DeleteStandardTemplateInvoker {
	requestDef := GenReqDefForDeleteStandardTemplate()
	return &DeleteStandardTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubject 删除主题
//
// 删除主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSubject(request *model.DeleteSubjectRequest) (*model.DeleteSubjectResponse, error) {
	requestDef := GenReqDefForDeleteSubject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubjectResponse), nil
	}
}

// DeleteSubjectInvoker 删除主题
func (c *DataArtsStudioClient) DeleteSubjectInvoker(request *model.DeleteSubjectRequest) *DeleteSubjectInvoker {
	requestDef := GenReqDefForDeleteSubject()
	return &DeleteSubjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteSubjectNew 删除主题(新)
//
// 删除主题(新)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteSubjectNew(request *model.DeleteSubjectNewRequest) (*model.DeleteSubjectNewResponse, error) {
	requestDef := GenReqDefForDeleteSubjectNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSubjectNewResponse), nil
	}
}

// DeleteSubjectNewInvoker 删除主题(新)
func (c *DataArtsStudioClient) DeleteSubjectNewInvoker(request *model.DeleteSubjectNewRequest) *DeleteSubjectNewInvoker {
	requestDef := GenReqDefForDeleteSubjectNew()
	return &DeleteSubjectNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTableModel 删除表模型
//
// 在关系建模中删除一个表模型及其属性，包括逻辑实体和物理表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteTableModel(request *model.DeleteTableModelRequest) (*model.DeleteTableModelResponse, error) {
	requestDef := GenReqDefForDeleteTableModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTableModelResponse), nil
	}
}

// DeleteTableModelInvoker 删除表模型
func (c *DataArtsStudioClient) DeleteTableModelInvoker(request *model.DeleteTableModelRequest) *DeleteTableModelInvoker {
	requestDef := GenReqDefForDeleteTableModel()
	return &DeleteTableModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTaskInfo 删除单个采集任务
//
// 删除单个采集任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteTaskInfo(request *model.DeleteTaskInfoRequest) (*model.DeleteTaskInfoResponse, error) {
	requestDef := GenReqDefForDeleteTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTaskInfoResponse), nil
	}
}

// DeleteTaskInfoInvoker 删除单个采集任务
func (c *DataArtsStudioClient) DeleteTaskInfoInvoker(request *model.DeleteTaskInfoRequest) *DeleteTaskInfoInvoker {
	requestDef := GenReqDefForDeleteTaskInfo()
	return &DeleteTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkspaces 删除模型工作区
//
// 删除模型工作区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteWorkspaces(request *model.DeleteWorkspacesRequest) (*model.DeleteWorkspacesResponse, error) {
	requestDef := GenReqDefForDeleteWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkspacesResponse), nil
	}
}

// DeleteWorkspacesInvoker 删除模型工作区
func (c *DataArtsStudioClient) DeleteWorkspacesInvoker(request *model.DeleteWorkspacesRequest) *DeleteWorkspacesInvoker {
	requestDef := GenReqDefForDeleteWorkspaces()
	return &DeleteWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteWorkspaceusers 删除工作空间用户
//
// 删除工作空间用户
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteWorkspaceusers(request *model.DeleteWorkspaceusersRequest) (*model.DeleteWorkspaceusersResponse, error) {
	requestDef := GenReqDefForDeleteWorkspaceusers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteWorkspaceusersResponse), nil
	}
}

// DeleteWorkspaceusersInvoker 删除工作空间用户
func (c *DataArtsStudioClient) DeleteWorkspaceusersInvoker(request *model.DeleteWorkspaceusersRequest) *DeleteWorkspaceusersInvoker {
	requestDef := GenReqDefForDeleteWorkspaceusers()
	return &DeleteWorkspaceusersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeployFactoryPackages 发布任务包
//
// 发布任务包
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeployFactoryPackages(request *model.DeployFactoryPackagesRequest) (*model.DeployFactoryPackagesResponse, error) {
	requestDef := GenReqDefForDeployFactoryPackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeployFactoryPackagesResponse), nil
	}
}

// DeployFactoryPackagesInvoker 发布任务包
func (c *DataArtsStudioClient) DeployFactoryPackagesInvoker(request *model.DeployFactoryPackagesRequest) *DeployFactoryPackagesInvoker {
	requestDef := GenReqDefForDeployFactoryPackages()
	return &DeployFactoryPackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTaskAction 启动、调度、停止采集任务
//
// 启动、调度、停止采集任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExecuteTaskAction(request *model.ExecuteTaskActionRequest) (*model.ExecuteTaskActionResponse, error) {
	requestDef := GenReqDefForExecuteTaskAction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTaskActionResponse), nil
	}
}

// ExecuteTaskActionInvoker 启动、调度、停止采集任务
func (c *DataArtsStudioClient) ExecuteTaskActionInvoker(request *model.ExecuteTaskActionRequest) *ExecuteTaskActionInvoker {
	requestDef := GenReqDefForExecuteTaskAction()
	return &ExecuteTaskActionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDesignModelTableDdl 导出模型中表的DDL语句
//
// 根据模型ID导出指定表的DDL语句。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDesignModelTableDdl(request *model.ExportDesignModelTableDdlRequest) (*model.ExportDesignModelTableDdlResponse, error) {
	requestDef := GenReqDefForExportDesignModelTableDdl()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDesignModelTableDdlResponse), nil
	}
}

// ExportDesignModelTableDdlInvoker 导出模型中表的DDL语句
func (c *DataArtsStudioClient) ExportDesignModelTableDdlInvoker(request *model.ExportDesignModelTableDdlRequest) *ExportDesignModelTableDdlInvoker {
	requestDef := GenReqDefForExportDesignModelTableDdl()
	return &ExportDesignModelTableDdlInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDesignModels 导出业务数据
//
// 根据请求参数，导出业务数据，可以导出：码表、数据标准、原子指标、衍生指标、复合指标、汇总表、业务指标、主题、流程、逻辑模型、物理模型、维度、事实表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDesignModels(request *model.ExportDesignModelsRequest) (*model.ExportDesignModelsResponse, error) {
	requestDef := GenReqDefForExportDesignModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDesignModelsResponse), nil
	}
}

// ExportDesignModelsInvoker 导出业务数据
func (c *DataArtsStudioClient) ExportDesignModelsInvoker(request *model.ExportDesignModelsRequest) *ExportDesignModelsInvoker {
	requestDef := GenReqDefForExportDesignModels()
	return &ExportDesignModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDesignResult 获取excel导出结果
//
// 根据请求导出业务数据（/export-model）时返回的uuid，获取excel导出结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDesignResult(request *model.ExportDesignResultRequest) (*model.ExportDesignResultResponse, error) {
	requestDef := GenReqDefForExportDesignResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDesignResultResponse), nil
	}
}

// ExportDesignResultInvoker 获取excel导出结果
func (c *DataArtsStudioClient) ExportDesignResultInvoker(request *model.ExportDesignResultRequest) *ExportDesignResultInvoker {
	requestDef := GenReqDefForExportDesignResult()
	return &ExportDesignResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportCatalogs 导入主题
//
// 用于导入主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ImportCatalogs(request *model.ImportCatalogsRequest) (*model.ImportCatalogsResponse, error) {
	requestDef := GenReqDefForImportCatalogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportCatalogsResponse), nil
	}
}

// ImportCatalogsInvoker 导入主题
func (c *DataArtsStudioClient) ImportCatalogsInvoker(request *model.ImportCatalogsRequest) *ImportCatalogsInvoker {
	requestDef := GenReqDefForImportCatalogs()
	return &ImportCatalogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportLineage 血缘导入
//
// 血缘查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ImportLineage(request *model.ImportLineageRequest) (*model.ImportLineageResponse, error) {
	requestDef := GenReqDefForImportLineage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportLineageResponse), nil
	}
}

// ImportLineageInvoker 血缘导入
func (c *DataArtsStudioClient) ImportLineageInvoker(request *model.ImportLineageRequest) *ImportLineageInvoker {
	requestDef := GenReqDefForImportLineage()
	return &ImportLineageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportModels 导入模型，关系建模，维度建模，码表，业务指标以及流程架构
//
// 导入模型，关系建模，维度建模，码表，业务指标以及流程架构。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ImportModels(request *model.ImportModelsRequest) (*model.ImportModelsResponse, error) {
	requestDef := GenReqDefForImportModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportModelsResponse), nil
	}
}

// ImportModelsInvoker 导入模型，关系建模，维度建模，码表，业务指标以及流程架构
func (c *DataArtsStudioClient) ImportModelsInvoker(request *model.ImportModelsRequest) *ImportModelsInvoker {
	requestDef := GenReqDefForImportModels()
	return &ImportModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportResult 查询导入结果
//
// 查询导入excel的处理结果（其中参数uuid获取为：/design/models/action或/design/catalogs/action接口返回结果）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ImportResult(request *model.ImportResultRequest) (*model.ImportResultResponse, error) {
	requestDef := GenReqDefForImportResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportResultResponse), nil
	}
}

// ImportResultInvoker 查询导入结果
func (c *DataArtsStudioClient) ImportResultInvoker(request *model.ImportResultRequest) *ImportResultInvoker {
	requestDef := GenReqDefForImportResult()
	return &ImportResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// InitializeStandardTemplate 初始化数据标准模板
//
// 初始化数据标准模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) InitializeStandardTemplate(request *model.InitializeStandardTemplateRequest) (*model.InitializeStandardTemplateResponse, error) {
	requestDef := GenReqDefForInitializeStandardTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.InitializeStandardTemplateResponse), nil
	}
}

// InitializeStandardTemplateInvoker 初始化数据标准模板
func (c *DataArtsStudioClient) InitializeStandardTemplateInvoker(request *model.InitializeStandardTemplateRequest) *InitializeStandardTemplateInvoker {
	requestDef := GenReqDefForInitializeStandardTemplate()
	return &InitializeStandardTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAggregationLogicTables 查找汇总表
//
// 通过中英文名称、创建者、审核人、状态、修改时间分页查找汇总表信息，中英文名称支持模糊查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListAggregationLogicTables(request *model.ListAggregationLogicTablesRequest) (*model.ListAggregationLogicTablesResponse, error) {
	requestDef := GenReqDefForListAggregationLogicTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAggregationLogicTablesResponse), nil
	}
}

// ListAggregationLogicTablesInvoker 查找汇总表
func (c *DataArtsStudioClient) ListAggregationLogicTablesInvoker(request *model.ListAggregationLogicTablesRequest) *ListAggregationLogicTablesInvoker {
	requestDef := GenReqDefForListAggregationLogicTables()
	return &ListAggregationLogicTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllCatalogList 获取当前目录下的所有类型列表
//
// 获取当前目录下所有类型列表（包括api和目录，均以目录的数据格式形式展示）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListAllCatalogList(request *model.ListAllCatalogListRequest) (*model.ListAllCatalogListResponse, error) {
	requestDef := GenReqDefForListAllCatalogList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllCatalogListResponse), nil
	}
}

// ListAllCatalogListInvoker 获取当前目录下的所有类型列表
func (c *DataArtsStudioClient) ListAllCatalogListInvoker(request *model.ListAllCatalogListRequest) *ListAllCatalogListInvoker {
	requestDef := GenReqDefForListAllCatalogList()
	return &ListAllCatalogListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllStandards 获取数据标准集合
//
// 根据查询条件分页获取数据标准集合，按修改时间降序排序。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListAllStandards(request *model.ListAllStandardsRequest) (*model.ListAllStandardsResponse, error) {
	requestDef := GenReqDefForListAllStandards()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllStandardsResponse), nil
	}
}

// ListAllStandardsInvoker 获取数据标准集合
func (c *DataArtsStudioClient) ListAllStandardsInvoker(request *model.ListAllStandardsRequest) *ListAllStandardsInvoker {
	requestDef := GenReqDefForListAllStandards()
	return &ListAllStandardsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAllTables 查询多种类型的表信息
//
// 从信息架构中查询多种类型的表信息，包括逻辑实体、物理表、维度表、事实表、汇总表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListAllTables(request *model.ListAllTablesRequest) (*model.ListAllTablesResponse, error) {
	requestDef := GenReqDefForListAllTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTablesResponse), nil
	}
}

// ListAllTablesInvoker 查询多种类型的表信息
func (c *DataArtsStudioClient) ListAllTablesInvoker(request *model.ListAllTablesRequest) *ListAllTablesInvoker {
	requestDef := GenReqDefForListAllTables()
	return &ListAllTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiCatalogList 获取当前目录下的api列表
//
// 获取当前目录下的api列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApiCatalogList(request *model.ListApiCatalogListRequest) (*model.ListApiCatalogListResponse, error) {
	requestDef := GenReqDefForListApiCatalogList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiCatalogListResponse), nil
	}
}

// ListApiCatalogListInvoker 获取当前目录下的api列表
func (c *DataArtsStudioClient) ListApiCatalogListInvoker(request *model.ListApiCatalogListRequest) *ListApiCatalogListInvoker {
	requestDef := GenReqDefForListApiCatalogList()
	return &ListApiCatalogListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApiTopN 查询指定api 应用调用topN
//
// 查询指定api 应用调用topN。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApiTopN(request *model.ListApiTopNRequest) (*model.ListApiTopNResponse, error) {
	requestDef := GenReqDefForListApiTopN()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiTopNResponse), nil
	}
}

// ListApiTopNInvoker 查询指定api 应用调用topN
func (c *DataArtsStudioClient) ListApiTopNInvoker(request *model.ListApiTopNRequest) *ListApiTopNInvoker {
	requestDef := GenReqDefForListApiTopN()
	return &ListApiTopNInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApicGroups 获取网关分组
//
// 获取网关分组。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApicGroups(request *model.ListApicGroupsRequest) (*model.ListApicGroupsResponse, error) {
	requestDef := GenReqDefForListApicGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApicGroupsResponse), nil
	}
}

// ListApicGroupsInvoker 获取网关分组
func (c *DataArtsStudioClient) ListApicGroupsInvoker(request *model.ListApicGroupsRequest) *ListApicGroupsInvoker {
	requestDef := GenReqDefForListApicGroups()
	return &ListApicGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApicInstances 获取网关实例(专享版)
//
// 获取网关实例(专享版)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApicInstances(request *model.ListApicInstancesRequest) (*model.ListApicInstancesResponse, error) {
	requestDef := GenReqDefForListApicInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApicInstancesResponse), nil
	}
}

// ListApicInstancesInvoker 获取网关实例(专享版)
func (c *DataArtsStudioClient) ListApicInstancesInvoker(request *model.ListApicInstancesRequest) *ListApicInstancesInvoker {
	requestDef := GenReqDefForListApicInstances()
	return &ListApicInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApisTop 查询api 服务调用topN
//
// 查询api 服务调用topN。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApisTop(request *model.ListApisTopRequest) (*model.ListApisTopResponse, error) {
	requestDef := GenReqDefForListApisTop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisTopResponse), nil
	}
}

// ListApisTopInvoker 查询api 服务调用topN
func (c *DataArtsStudioClient) ListApisTopInvoker(request *model.ListApisTopRequest) *ListApisTopInvoker {
	requestDef := GenReqDefForListApisTop()
	return &ListApisTopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApply 查询申请列表
//
// 查询申请列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApply(request *model.ListApplyRequest) (*model.ListApplyResponse, error) {
	requestDef := GenReqDefForListApply()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApplyResponse), nil
	}
}

// ListApplyInvoker 查询申请列表
func (c *DataArtsStudioClient) ListApplyInvoker(request *model.ListApplyRequest) *ListApplyInvoker {
	requestDef := GenReqDefForListApply()
	return &ListApplyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApprovers 查询审批人列表
//
// 查询审批人列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApprovers(request *model.ListApproversRequest) (*model.ListApproversResponse, error) {
	requestDef := GenReqDefForListApprovers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApproversResponse), nil
	}
}

// ListApproversInvoker 查询审批人列表
func (c *DataArtsStudioClient) ListApproversInvoker(request *model.ListApproversRequest) *ListApproversInvoker {
	requestDef := GenReqDefForListApprovers()
	return &ListApproversInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApps 查询应用列表
//
// 查询应用列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApps(request *model.ListAppsRequest) (*model.ListAppsResponse, error) {
	requestDef := GenReqDefForListApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsResponse), nil
	}
}

// ListAppsInvoker 查询应用列表
func (c *DataArtsStudioClient) ListAppsInvoker(request *model.ListAppsRequest) *ListAppsInvoker {
	requestDef := GenReqDefForListApps()
	return &ListAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAppsTop 查询app 服务使用topN
//
// 查询app 服务使用topN。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListAppsTop(request *model.ListAppsTopRequest) (*model.ListAppsTopResponse, error) {
	requestDef := GenReqDefForListAppsTop()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAppsTopResponse), nil
	}
}

// ListAppsTopInvoker 查询app 服务使用topN
func (c *DataArtsStudioClient) ListAppsTopInvoker(request *model.ListAppsTopRequest) *ListAppsTopInvoker {
	requestDef := GenReqDefForListAppsTop()
	return &ListAppsTopInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBizMetricDimensions 查看指标维度信息
//
// 查看指标维度信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListBizMetricDimensions(request *model.ListBizMetricDimensionsRequest) (*model.ListBizMetricDimensionsResponse, error) {
	requestDef := GenReqDefForListBizMetricDimensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBizMetricDimensionsResponse), nil
	}
}

// ListBizMetricDimensionsInvoker 查看指标维度信息
func (c *DataArtsStudioClient) ListBizMetricDimensionsInvoker(request *model.ListBizMetricDimensionsRequest) *ListBizMetricDimensionsInvoker {
	requestDef := GenReqDefForListBizMetricDimensions()
	return &ListBizMetricDimensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBizMetricOwners 查看指标责任人信息
//
// 查看指标责任人信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListBizMetricOwners(request *model.ListBizMetricOwnersRequest) (*model.ListBizMetricOwnersResponse, error) {
	requestDef := GenReqDefForListBizMetricOwners()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBizMetricOwnersResponse), nil
	}
}

// ListBizMetricOwnersInvoker 查看指标责任人信息
func (c *DataArtsStudioClient) ListBizMetricOwnersInvoker(request *model.ListBizMetricOwnersRequest) *ListBizMetricOwnersInvoker {
	requestDef := GenReqDefForListBizMetricOwners()
	return &ListBizMetricOwnersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBizMetrics 查询业务指标信息
//
// 通过名称、创建者、修改时间分页查找业务指标信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListBizMetrics(request *model.ListBizMetricsRequest) (*model.ListBizMetricsResponse, error) {
	requestDef := GenReqDefForListBizMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBizMetricsResponse), nil
	}
}

// ListBizMetricsInvoker 查询业务指标信息
func (c *DataArtsStudioClient) ListBizMetricsInvoker(request *model.ListBizMetricsRequest) *ListBizMetricsInvoker {
	requestDef := GenReqDefForListBizMetrics()
	return &ListBizMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBusiness 获取主题树信息
//
// 获取数据资产主题树信息l1，l2，l3。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListBusiness(request *model.ListBusinessRequest) (*model.ListBusinessResponse, error) {
	requestDef := GenReqDefForListBusiness()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBusinessResponse), nil
	}
}

// ListBusinessInvoker 获取主题树信息
func (c *DataArtsStudioClient) ListBusinessInvoker(request *model.ListBusinessRequest) *ListBusinessInvoker {
	requestDef := GenReqDefForListBusiness()
	return &ListBusinessInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogList 获取当前目录下的目录列表（全量）
//
// 获取当前目录下的目录列表（全量数据，不分页，推荐仅用于生成目录树等无法分页的场景）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListCatalogList(request *model.ListCatalogListRequest) (*model.ListCatalogListResponse, error) {
	requestDef := GenReqDefForListCatalogList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogListResponse), nil
	}
}

// ListCatalogListInvoker 获取当前目录下的目录列表（全量）
func (c *DataArtsStudioClient) ListCatalogListInvoker(request *model.ListCatalogListRequest) *ListCatalogListInvoker {
	requestDef := GenReqDefForListCatalogList()
	return &ListCatalogListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCatalogTree 获取所有流程架构目录树
//
// 获取所有目录树。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListCatalogTree(request *model.ListCatalogTreeRequest) (*model.ListCatalogTreeResponse, error) {
	requestDef := GenReqDefForListCatalogTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCatalogTreeResponse), nil
	}
}

// ListCatalogTreeInvoker 获取所有流程架构目录树
func (c *DataArtsStudioClient) ListCatalogTreeInvoker(request *model.ListCatalogTreeRequest) *ListCatalogTreeInvoker {
	requestDef := GenReqDefForListCatalogTree()
	return &ListCatalogTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCategory 获取作业目录
//
// 获取作业目录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListCategory(request *model.ListCategoryRequest) (*model.ListCategoryResponse, error) {
	requestDef := GenReqDefForListCategory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCategoryResponse), nil
	}
}

// ListCategoryInvoker 获取作业目录
func (c *DataArtsStudioClient) ListCategoryInvoker(request *model.ListCategoryRequest) *ListCategoryInvoker {
	requestDef := GenReqDefForListCategory()
	return &ListCategoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListColumns 获取数据源中表的字段
//
// 获取数据源中表的字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListColumns(request *model.ListColumnsRequest) (*model.ListColumnsResponse, error) {
	requestDef := GenReqDefForListColumns()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListColumnsResponse), nil
	}
}

// ListColumnsInvoker 获取数据源中表的字段
func (c *DataArtsStudioClient) ListColumnsInvoker(request *model.ListColumnsRequest) *ListColumnsInvoker {
	requestDef := GenReqDefForListColumns()
	return &ListColumnsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCompoundMetrics 查找复合指标
//
// 通过中英文名称、创建者、审核人、状态、修改时间、l3Id分页查找复合指标信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListCompoundMetrics(request *model.ListCompoundMetricsRequest) (*model.ListCompoundMetricsResponse, error) {
	requestDef := GenReqDefForListCompoundMetrics()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCompoundMetricsResponse), nil
	}
}

// ListCompoundMetricsInvoker 查找复合指标
func (c *DataArtsStudioClient) ListCompoundMetricsInvoker(request *model.ListCompoundMetricsRequest) *ListCompoundMetricsInvoker {
	requestDef := GenReqDefForListCompoundMetrics()
	return &ListCompoundMetricsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCondition 查找业务限定
//
// 通过中英文名称、描述、创建者、审核人、限定分组id、修改时间状态分页查找限定信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListCondition(request *model.ListConditionRequest) (*model.ListConditionResponse, error) {
	requestDef := GenReqDefForListCondition()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConditionResponse), nil
	}
}

// ListConditionInvoker 查找业务限定
func (c *DataArtsStudioClient) ListConditionInvoker(request *model.ListConditionRequest) *ListConditionInvoker {
	requestDef := GenReqDefForListCondition()
	return &ListConditionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListConsistencyTask 获取对账作业列表
//
// 获取对账作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListConsistencyTask(request *model.ListConsistencyTaskRequest) (*model.ListConsistencyTaskResponse, error) {
	requestDef := GenReqDefForListConsistencyTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListConsistencyTaskResponse), nil
	}
}

// ListConsistencyTaskInvoker 获取对账作业列表
func (c *DataArtsStudioClient) ListConsistencyTaskInvoker(request *model.ListConsistencyTaskRequest) *ListConsistencyTaskInvoker {
	requestDef := GenReqDefForListConsistencyTask()
	return &ListConsistencyTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataArtsStudioInstances 获取实例列表
//
// 获取实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataArtsStudioInstances(request *model.ListDataArtsStudioInstancesRequest) (*model.ListDataArtsStudioInstancesResponse, error) {
	requestDef := GenReqDefForListDataArtsStudioInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataArtsStudioInstancesResponse), nil
	}
}

// ListDataArtsStudioInstancesInvoker 获取实例列表
func (c *DataArtsStudioClient) ListDataArtsStudioInstancesInvoker(request *model.ListDataArtsStudioInstancesRequest) *ListDataArtsStudioInstancesInvoker {
	requestDef := GenReqDefForListDataArtsStudioInstances()
	return &ListDataArtsStudioInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataServiceInstancesDetail 查询集群详情信息列表
//
// 查询集群详情信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataServiceInstancesDetail(request *model.ListDataServiceInstancesDetailRequest) (*model.ListDataServiceInstancesDetailResponse, error) {
	requestDef := GenReqDefForListDataServiceInstancesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataServiceInstancesDetailResponse), nil
	}
}

// ListDataServiceInstancesDetailInvoker 查询集群详情信息列表
func (c *DataArtsStudioClient) ListDataServiceInstancesDetailInvoker(request *model.ListDataServiceInstancesDetailRequest) *ListDataServiceInstancesDetailInvoker {
	requestDef := GenReqDefForListDataServiceInstancesDetail()
	return &ListDataServiceInstancesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataServiceInstancesOverview 查询集群概览信息列表
//
// 查询集群概览信息列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataServiceInstancesOverview(request *model.ListDataServiceInstancesOverviewRequest) (*model.ListDataServiceInstancesOverviewResponse, error) {
	requestDef := GenReqDefForListDataServiceInstancesOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataServiceInstancesOverviewResponse), nil
	}
}

// ListDataServiceInstancesOverviewInvoker 查询集群概览信息列表
func (c *DataArtsStudioClient) ListDataServiceInstancesOverviewInvoker(request *model.ListDataServiceInstancesOverviewRequest) *ListDataServiceInstancesOverviewInvoker {
	requestDef := GenReqDefForListDataServiceInstancesOverview()
	return &ListDataServiceInstancesOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataServiceMarketApis 查询服务目录API列表
//
// 查询服务目录API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataServiceMarketApis(request *model.ListDataServiceMarketApisRequest) (*model.ListDataServiceMarketApisResponse, error) {
	requestDef := GenReqDefForListDataServiceMarketApis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataServiceMarketApisResponse), nil
	}
}

// ListDataServiceMarketApisInvoker 查询服务目录API列表
func (c *DataArtsStudioClient) ListDataServiceMarketApisInvoker(request *model.ListDataServiceMarketApisRequest) *ListDataServiceMarketApisInvoker {
	requestDef := GenReqDefForListDataServiceMarketApis()
	return &ListDataServiceMarketApisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataTables 获取数据源中的表
//
// 获取数据源中的表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataTables(request *model.ListDataTablesRequest) (*model.ListDataTablesResponse, error) {
	requestDef := GenReqDefForListDataTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataTablesResponse), nil
	}
}

// ListDataTablesInvoker 获取数据源中的表
func (c *DataArtsStudioClient) ListDataTablesInvoker(request *model.ListDataTablesRequest) *ListDataTablesInvoker {
	requestDef := GenReqDefForListDataTables()
	return &ListDataTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDatabases 获取数据库列表
//
// 获取数据库列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDatabases(request *model.ListDatabasesRequest) (*model.ListDatabasesResponse, error) {
	requestDef := GenReqDefForListDatabases()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDatabasesResponse), nil
	}
}

// ListDatabasesInvoker 获取数据库列表
func (c *DataArtsStudioClient) ListDatabasesInvoker(request *model.ListDatabasesRequest) *ListDatabasesInvoker {
	requestDef := GenReqDefForListDatabases()
	return &ListDatabasesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDataconnections 查询数据连接列表
//
// 查询数据连接列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDataconnections(request *model.ListDataconnectionsRequest) (*model.ListDataconnectionsResponse, error) {
	requestDef := GenReqDefForListDataconnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDataconnectionsResponse), nil
	}
}

// ListDataconnectionsInvoker 查询数据连接列表
func (c *DataArtsStudioClient) ListDataconnectionsInvoker(request *model.ListDataconnectionsRequest) *ListDataconnectionsInvoker {
	requestDef := GenReqDefForListDataconnections()
	return &ListDataconnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDerivativeIndexes 查找衍生指标
//
// 通过中英文名称、创建者、审核人、状态、修改时间、l3Id分页查找衍生指标信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDerivativeIndexes(request *model.ListDerivativeIndexesRequest) (*model.ListDerivativeIndexesResponse, error) {
	requestDef := GenReqDefForListDerivativeIndexes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDerivativeIndexesResponse), nil
	}
}

// ListDerivativeIndexesInvoker 查找衍生指标
func (c *DataArtsStudioClient) ListDerivativeIndexesInvoker(request *model.ListDerivativeIndexesRequest) *ListDerivativeIndexesInvoker {
	requestDef := GenReqDefForListDerivativeIndexes()
	return &ListDerivativeIndexesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDimensionGroups 查看维度颗粒度
//
// 查询维度颗粒度，依据tableId查询涉及所有维度，不传tableId查询所有维度组颗粒度。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDimensionGroups(request *model.ListDimensionGroupsRequest) (*model.ListDimensionGroupsResponse, error) {
	requestDef := GenReqDefForListDimensionGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDimensionGroupsResponse), nil
	}
}

// ListDimensionGroupsInvoker 查看维度颗粒度
func (c *DataArtsStudioClient) ListDimensionGroupsInvoker(request *model.ListDimensionGroupsRequest) *ListDimensionGroupsInvoker {
	requestDef := GenReqDefForListDimensionGroups()
	return &ListDimensionGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDimensionLogicTables 查找维度表
//
// 通过中英文名称、创建者、审核人、状态、修改时间分页查找维度表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDimensionLogicTables(request *model.ListDimensionLogicTablesRequest) (*model.ListDimensionLogicTablesResponse, error) {
	requestDef := GenReqDefForListDimensionLogicTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDimensionLogicTablesResponse), nil
	}
}

// ListDimensionLogicTablesInvoker 查找维度表
func (c *DataArtsStudioClient) ListDimensionLogicTablesInvoker(request *model.ListDimensionLogicTablesRequest) *ListDimensionLogicTablesInvoker {
	requestDef := GenReqDefForListDimensionLogicTables()
	return &ListDimensionLogicTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDimensions 查找维度
//
// 通过中英文名称、描述、创建者、审核人、状态、l3Id、衍生指标idList、修改时间分页查找维度信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDimensions(request *model.ListDimensionsRequest) (*model.ListDimensionsResponse, error) {
	requestDef := GenReqDefForListDimensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDimensionsResponse), nil
	}
}

// ListDimensionsInvoker 查找维度
func (c *DataArtsStudioClient) ListDimensionsInvoker(request *model.ListDimensionsRequest) *ListDimensionsInvoker {
	requestDef := GenReqDefForListDimensions()
	return &ListDimensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDirectories 获取所有目录
//
// 获取所有目录(数据标准、码表)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListDirectories(request *model.ListDirectoriesRequest) (*model.ListDirectoriesResponse, error) {
	requestDef := GenReqDefForListDirectories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDirectoriesResponse), nil
	}
}

// ListDirectoriesInvoker 获取所有目录
func (c *DataArtsStudioClient) ListDirectoriesInvoker(request *model.ListDirectoriesRequest) *ListDirectoriesInvoker {
	requestDef := GenReqDefForListDirectories()
	return &ListDirectoriesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactLogicTables 查找事实表
//
// 通过中英文名称、创建者、审核人、状态、修改时间分页查找事实表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactLogicTables(request *model.ListFactLogicTablesRequest) (*model.ListFactLogicTablesResponse, error) {
	requestDef := GenReqDefForListFactLogicTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactLogicTablesResponse), nil
	}
}

// ListFactLogicTablesInvoker 查找事实表
func (c *DataArtsStudioClient) ListFactLogicTablesInvoker(request *model.ListFactLogicTablesRequest) *ListFactLogicTablesInvoker {
	requestDef := GenReqDefForListFactLogicTables()
	return &ListFactLogicTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryAlarmInfo 查询告警通知记录
//
// 查询告警通知记录
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryAlarmInfo(request *model.ListFactoryAlarmInfoRequest) (*model.ListFactoryAlarmInfoResponse, error) {
	requestDef := GenReqDefForListFactoryAlarmInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryAlarmInfoResponse), nil
	}
}

// ListFactoryAlarmInfoInvoker 查询告警通知记录
func (c *DataArtsStudioClient) ListFactoryAlarmInfoInvoker(request *model.ListFactoryAlarmInfoRequest) *ListFactoryAlarmInfoInvoker {
	requestDef := GenReqDefForListFactoryAlarmInfo()
	return &ListFactoryAlarmInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryJobInstancesByName 查询指定作业的实例列表
//
// 查询指定作业的实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryJobInstancesByName(request *model.ListFactoryJobInstancesByNameRequest) (*model.ListFactoryJobInstancesByNameResponse, error) {
	requestDef := GenReqDefForListFactoryJobInstancesByName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryJobInstancesByNameResponse), nil
	}
}

// ListFactoryJobInstancesByNameInvoker 查询指定作业的实例列表
func (c *DataArtsStudioClient) ListFactoryJobInstancesByNameInvoker(request *model.ListFactoryJobInstancesByNameRequest) *ListFactoryJobInstancesByNameInvoker {
	requestDef := GenReqDefForListFactoryJobInstancesByName()
	return &ListFactoryJobInstancesByNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryJobs 查询作业列表
//
// 查询作业列表清单
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryJobs(request *model.ListFactoryJobsRequest) (*model.ListFactoryJobsResponse, error) {
	requestDef := GenReqDefForListFactoryJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryJobsResponse), nil
	}
}

// ListFactoryJobsInvoker 查询作业列表
func (c *DataArtsStudioClient) ListFactoryJobsInvoker(request *model.ListFactoryJobsRequest) *ListFactoryJobsInvoker {
	requestDef := GenReqDefForListFactoryJobs()
	return &ListFactoryJobsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryReleasePackages 查询发布包列表
//
// 查询发布包列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryReleasePackages(request *model.ListFactoryReleasePackagesRequest) (*model.ListFactoryReleasePackagesResponse, error) {
	requestDef := GenReqDefForListFactoryReleasePackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryReleasePackagesResponse), nil
	}
}

// ListFactoryReleasePackagesInvoker 查询发布包列表
func (c *DataArtsStudioClient) ListFactoryReleasePackagesInvoker(request *model.ListFactoryReleasePackagesRequest) *ListFactoryReleasePackagesInvoker {
	requestDef := GenReqDefForListFactoryReleasePackages()
	return &ListFactoryReleasePackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryTaskCompletion 查询任务完成情况
//
// 查询任务完成情况
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryTaskCompletion(request *model.ListFactoryTaskCompletionRequest) (*model.ListFactoryTaskCompletionResponse, error) {
	requestDef := GenReqDefForListFactoryTaskCompletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryTaskCompletionResponse), nil
	}
}

// ListFactoryTaskCompletionInvoker 查询任务完成情况
func (c *DataArtsStudioClient) ListFactoryTaskCompletionInvoker(request *model.ListFactoryTaskCompletionRequest) *ListFactoryTaskCompletionInvoker {
	requestDef := GenReqDefForListFactoryTaskCompletion()
	return &ListFactoryTaskCompletionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFactoryTaskOverview 查询实例运行状态
//
// 查询实例运行状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListFactoryTaskOverview(request *model.ListFactoryTaskOverviewRequest) (*model.ListFactoryTaskOverviewResponse, error) {
	requestDef := GenReqDefForListFactoryTaskOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFactoryTaskOverviewResponse), nil
	}
}

// ListFactoryTaskOverviewInvoker 查询实例运行状态
func (c *DataArtsStudioClient) ListFactoryTaskOverviewInvoker(request *model.ListFactoryTaskOverviewRequest) *ListFactoryTaskOverviewInvoker {
	requestDef := GenReqDefForListFactoryTaskOverview()
	return &ListFactoryTaskOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstances 获取任务执行结果列表
//
// 获取任务执行结果列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 获取任务执行结果列表
func (c *DataArtsStudioClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListManagerWorkSpaces 获取工作空间列表
//
// 获取工作空间列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListManagerWorkSpaces(request *model.ListManagerWorkSpacesRequest) (*model.ListManagerWorkSpacesResponse, error) {
	requestDef := GenReqDefForListManagerWorkSpaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListManagerWorkSpacesResponse), nil
	}
}

// ListManagerWorkSpacesInvoker 获取工作空间列表
func (c *DataArtsStudioClient) ListManagerWorkSpacesInvoker(request *model.ListManagerWorkSpacesRequest) *ListManagerWorkSpacesInvoker {
	requestDef := GenReqDefForListManagerWorkSpaces()
	return &ListManagerWorkSpacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMessage 查询消息列表
//
// 查询审核中心的通知消息列表。与申请不同，通知类消息，无法驳回，仅能在指定的时间范围内作出处理。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListMessage(request *model.ListMessageRequest) (*model.ListMessageResponse, error) {
	requestDef := GenReqDefForListMessage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMessageResponse), nil
	}
}

// ListMessageInvoker 查询消息列表
func (c *DataArtsStudioClient) ListMessageInvoker(request *model.ListMessageRequest) *ListMessageInvoker {
	requestDef := GenReqDefForListMessage()
	return &ListMessageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListMetricRelations 获取指标关联信息
//
// 获取当前指标图谱。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListMetricRelations(request *model.ListMetricRelationsRequest) (*model.ListMetricRelationsResponse, error) {
	requestDef := GenReqDefForListMetricRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListMetricRelationsResponse), nil
	}
}

// ListMetricRelationsInvoker 获取指标关联信息
func (c *DataArtsStudioClient) ListMetricRelationsInvoker(request *model.ListMetricRelationsRequest) *ListMetricRelationsInvoker {
	requestDef := GenReqDefForListMetricRelations()
	return &ListMetricRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQualityTask 获取质量作业列表
//
// 获取质量作业列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListQualityTask(request *model.ListQualityTaskRequest) (*model.ListQualityTaskResponse, error) {
	requestDef := GenReqDefForListQualityTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQualityTaskResponse), nil
	}
}

// ListQualityTaskInvoker 获取质量作业列表
func (c *DataArtsStudioClient) ListQualityTaskInvoker(request *model.ListQualityTaskRequest) *ListQualityTaskInvoker {
	requestDef := GenReqDefForListQualityTask()
	return &ListQualityTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQualityTaskLists 获取质量作业列表V1
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListQualityTaskLists(request *model.ListQualityTaskListsRequest) (*model.ListQualityTaskListsResponse, error) {
	requestDef := GenReqDefForListQualityTaskLists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQualityTaskListsResponse), nil
	}
}

// ListQualityTaskListsInvoker 获取质量作业列表V1
func (c *DataArtsStudioClient) ListQualityTaskListsInvoker(request *model.ListQualityTaskListsRequest) *ListQualityTaskListsInvoker {
	requestDef := GenReqDefForListQualityTaskLists()
	return &ListQualityTaskListsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQualityTemplates 获取规则模板列表
//
// 分页获取规则模板列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListQualityTemplates(request *model.ListQualityTemplatesRequest) (*model.ListQualityTemplatesResponse, error) {
	requestDef := GenReqDefForListQualityTemplates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQualityTemplatesResponse), nil
	}
}

// ListQualityTemplatesInvoker 获取规则模板列表
func (c *DataArtsStudioClient) ListQualityTemplatesInvoker(request *model.ListQualityTemplatesRequest) *ListQualityTemplatesInvoker {
	requestDef := GenReqDefForListQualityTemplates()
	return &ListQualityTemplatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRelations 查询关系
//
// 通过关系名称(支持模糊查询)、创建人、开始时间、结束时间等分页查找关系信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListRelations(request *model.ListRelationsRequest) (*model.ListRelationsResponse, error) {
	requestDef := GenReqDefForListRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRelationsResponse), nil
	}
}

// ListRelationsInvoker 查询关系
func (c *DataArtsStudioClient) ListRelationsInvoker(request *model.ListRelationsRequest) *ListRelationsInvoker {
	requestDef := GenReqDefForListRelations()
	return &ListRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSchemas 获取schemas
//
// 获取schemas,目前只有DWS和采用postgresql驱动的RDS数据源支持schema,请在调用前确认该数据源是否支持schema字段
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSchemas(request *model.ListSchemasRequest) (*model.ListSchemasResponse, error) {
	requestDef := GenReqDefForListSchemas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSchemasResponse), nil
	}
}

// ListSchemasInvoker 获取schemas
func (c *DataArtsStudioClient) ListSchemasInvoker(request *model.ListSchemasRequest) *ListSchemasInvoker {
	requestDef := GenReqDefForListSchemas()
	return &ListSchemasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityAssignedQueues 查询当前空间下分配的队列资源
//
// 查询当前空间下分配的队列资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityAssignedQueues(request *model.ListSecurityAssignedQueuesRequest) (*model.ListSecurityAssignedQueuesResponse, error) {
	requestDef := GenReqDefForListSecurityAssignedQueues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityAssignedQueuesResponse), nil
	}
}

// ListSecurityAssignedQueuesInvoker 查询当前空间下分配的队列资源
func (c *DataArtsStudioClient) ListSecurityAssignedQueuesInvoker(request *model.ListSecurityAssignedQueuesRequest) *ListSecurityAssignedQueuesInvoker {
	requestDef := GenReqDefForListSecurityAssignedQueues()
	return &ListSecurityAssignedQueuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDataClassificationRuleGroups 查询规则组列表
//
// 查询规则组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDataClassificationRuleGroups(request *model.ListSecurityDataClassificationRuleGroupsRequest) (*model.ListSecurityDataClassificationRuleGroupsResponse, error) {
	requestDef := GenReqDefForListSecurityDataClassificationRuleGroups()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDataClassificationRuleGroupsResponse), nil
	}
}

// ListSecurityDataClassificationRuleGroupsInvoker 查询规则组列表
func (c *DataArtsStudioClient) ListSecurityDataClassificationRuleGroupsInvoker(request *model.ListSecurityDataClassificationRuleGroupsRequest) *ListSecurityDataClassificationRuleGroupsInvoker {
	requestDef := GenReqDefForListSecurityDataClassificationRuleGroups()
	return &ListSecurityDataClassificationRuleGroupsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDataClassificationRules 查询识别规则列表
//
// 查询识别规则列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDataClassificationRules(request *model.ListSecurityDataClassificationRulesRequest) (*model.ListSecurityDataClassificationRulesResponse, error) {
	requestDef := GenReqDefForListSecurityDataClassificationRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDataClassificationRulesResponse), nil
	}
}

// ListSecurityDataClassificationRulesInvoker 查询识别规则列表
func (c *DataArtsStudioClient) ListSecurityDataClassificationRulesInvoker(request *model.ListSecurityDataClassificationRulesRequest) *ListSecurityDataClassificationRulesInvoker {
	requestDef := GenReqDefForListSecurityDataClassificationRules()
	return &ListSecurityDataClassificationRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDatasourceActions 查询数据操作信息
//
// 查询数据操作信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDatasourceActions(request *model.ListSecurityDatasourceActionsRequest) (*model.ListSecurityDatasourceActionsResponse, error) {
	requestDef := GenReqDefForListSecurityDatasourceActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDatasourceActionsResponse), nil
	}
}

// ListSecurityDatasourceActionsInvoker 查询数据操作信息
func (c *DataArtsStudioClient) ListSecurityDatasourceActionsInvoker(request *model.ListSecurityDatasourceActionsRequest) *ListSecurityDatasourceActionsInvoker {
	requestDef := GenReqDefForListSecurityDatasourceActions()
	return &ListSecurityDatasourceActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDatasourceConfigurations 查询数据源可配置权限
//
// 查询数据源可配置权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDatasourceConfigurations(request *model.ListSecurityDatasourceConfigurationsRequest) (*model.ListSecurityDatasourceConfigurationsResponse, error) {
	requestDef := GenReqDefForListSecurityDatasourceConfigurations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDatasourceConfigurationsResponse), nil
	}
}

// ListSecurityDatasourceConfigurationsInvoker 查询数据源可配置权限
func (c *DataArtsStudioClient) ListSecurityDatasourceConfigurationsInvoker(request *model.ListSecurityDatasourceConfigurationsRequest) *ListSecurityDatasourceConfigurationsInvoker {
	requestDef := GenReqDefForListSecurityDatasourceConfigurations()
	return &ListSecurityDatasourceConfigurationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDatasourceUrls 查询url信息
//
// 查询url信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDatasourceUrls(request *model.ListSecurityDatasourceUrlsRequest) (*model.ListSecurityDatasourceUrlsResponse, error) {
	requestDef := GenReqDefForListSecurityDatasourceUrls()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDatasourceUrlsResponse), nil
	}
}

// ListSecurityDatasourceUrlsInvoker 查询url信息
func (c *DataArtsStudioClient) ListSecurityDatasourceUrlsInvoker(request *model.ListSecurityDatasourceUrlsRequest) *ListSecurityDatasourceUrlsInvoker {
	requestDef := GenReqDefForListSecurityDatasourceUrls()
	return &ListSecurityDatasourceUrlsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDlfDataWareHouses 查询数据开发细粒度连接列表（全量）
//
// 查询数据开发细粒度连接列表（全量）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDlfDataWareHouses(request *model.ListSecurityDlfDataWareHousesRequest) (*model.ListSecurityDlfDataWareHousesResponse, error) {
	requestDef := GenReqDefForListSecurityDlfDataWareHouses()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDlfDataWareHousesResponse), nil
	}
}

// ListSecurityDlfDataWareHousesInvoker 查询数据开发细粒度连接列表（全量）
func (c *DataArtsStudioClient) ListSecurityDlfDataWareHousesInvoker(request *model.ListSecurityDlfDataWareHousesRequest) *ListSecurityDlfDataWareHousesInvoker {
	requestDef := GenReqDefForListSecurityDlfDataWareHouses()
	return &ListSecurityDlfDataWareHousesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityDynamicMaskingPolicies 查询动态数据脱敏策略列表
//
// 查询动态数据脱敏策略列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityDynamicMaskingPolicies(request *model.ListSecurityDynamicMaskingPoliciesRequest) (*model.ListSecurityDynamicMaskingPoliciesResponse, error) {
	requestDef := GenReqDefForListSecurityDynamicMaskingPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityDynamicMaskingPoliciesResponse), nil
	}
}

// ListSecurityDynamicMaskingPoliciesInvoker 查询动态数据脱敏策略列表
func (c *DataArtsStudioClient) ListSecurityDynamicMaskingPoliciesInvoker(request *model.ListSecurityDynamicMaskingPoliciesRequest) *ListSecurityDynamicMaskingPoliciesInvoker {
	requestDef := GenReqDefForListSecurityDynamicMaskingPolicies()
	return &ListSecurityDynamicMaskingPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityMemberSyncTasks 查询用户同步列表
//
// 查询用户同步列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityMemberSyncTasks(request *model.ListSecurityMemberSyncTasksRequest) (*model.ListSecurityMemberSyncTasksResponse, error) {
	requestDef := GenReqDefForListSecurityMemberSyncTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityMemberSyncTasksResponse), nil
	}
}

// ListSecurityMemberSyncTasksInvoker 查询用户同步列表
func (c *DataArtsStudioClient) ListSecurityMemberSyncTasksInvoker(request *model.ListSecurityMemberSyncTasksRequest) *ListSecurityMemberSyncTasksInvoker {
	requestDef := GenReqDefForListSecurityMemberSyncTasks()
	return &ListSecurityMemberSyncTasksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityPermissionSetMembers 查询权限集成员列表
//
// 查询权限集成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityPermissionSetMembers(request *model.ListSecurityPermissionSetMembersRequest) (*model.ListSecurityPermissionSetMembersResponse, error) {
	requestDef := GenReqDefForListSecurityPermissionSetMembers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPermissionSetMembersResponse), nil
	}
}

// ListSecurityPermissionSetMembersInvoker 查询权限集成员列表
func (c *DataArtsStudioClient) ListSecurityPermissionSetMembersInvoker(request *model.ListSecurityPermissionSetMembersRequest) *ListSecurityPermissionSetMembersInvoker {
	requestDef := GenReqDefForListSecurityPermissionSetMembers()
	return &ListSecurityPermissionSetMembersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityPermissionSetPermissions 查询权限集的权限列表
//
// 查询权限集的权限列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityPermissionSetPermissions(request *model.ListSecurityPermissionSetPermissionsRequest) (*model.ListSecurityPermissionSetPermissionsResponse, error) {
	requestDef := GenReqDefForListSecurityPermissionSetPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPermissionSetPermissionsResponse), nil
	}
}

// ListSecurityPermissionSetPermissionsInvoker 查询权限集的权限列表
func (c *DataArtsStudioClient) ListSecurityPermissionSetPermissionsInvoker(request *model.ListSecurityPermissionSetPermissionsRequest) *ListSecurityPermissionSetPermissionsInvoker {
	requestDef := GenReqDefForListSecurityPermissionSetPermissions()
	return &ListSecurityPermissionSetPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityPermissionSets 查询权限集列表
//
// 查询权限集列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityPermissionSets(request *model.ListSecurityPermissionSetsRequest) (*model.ListSecurityPermissionSetsResponse, error) {
	requestDef := GenReqDefForListSecurityPermissionSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityPermissionSetsResponse), nil
	}
}

// ListSecurityPermissionSetsInvoker 查询权限集列表
func (c *DataArtsStudioClient) ListSecurityPermissionSetsInvoker(request *model.ListSecurityPermissionSetsRequest) *ListSecurityPermissionSetsInvoker {
	requestDef := GenReqDefForListSecurityPermissionSets()
	return &ListSecurityPermissionSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecurityRoleActions 查询角色对一组库、表的权限交集
//
// 查询角色对一组库、表的权限交集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecurityRoleActions(request *model.ListSecurityRoleActionsRequest) (*model.ListSecurityRoleActionsResponse, error) {
	requestDef := GenReqDefForListSecurityRoleActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecurityRoleActionsResponse), nil
	}
}

// ListSecurityRoleActionsInvoker 查询角色对一组库、表的权限交集
func (c *DataArtsStudioClient) ListSecurityRoleActionsInvoker(request *model.ListSecurityRoleActionsRequest) *ListSecurityRoleActionsInvoker {
	requestDef := GenReqDefForListSecurityRoleActions()
	return &ListSecurityRoleActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecuritySecrecyLevels 获取密级
//
// 获取密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecuritySecrecyLevels(request *model.ListSecuritySecrecyLevelsRequest) (*model.ListSecuritySecrecyLevelsResponse, error) {
	requestDef := GenReqDefForListSecuritySecrecyLevels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecuritySecrecyLevelsResponse), nil
	}
}

// ListSecuritySecrecyLevelsInvoker 获取密级
func (c *DataArtsStudioClient) ListSecuritySecrecyLevelsInvoker(request *model.ListSecuritySecrecyLevelsRequest) *ListSecuritySecrecyLevelsInvoker {
	requestDef := GenReqDefForListSecuritySecrecyLevels()
	return &ListSecuritySecrecyLevelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSecuritySensitiveDataOverviews 查询敏感数据发现概览结果(以分类和密级为单位)
//
// 查询敏感数据发现概览结果(以分类和密级为单位)
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSecuritySensitiveDataOverviews(request *model.ListSecuritySensitiveDataOverviewsRequest) (*model.ListSecuritySensitiveDataOverviewsResponse, error) {
	requestDef := GenReqDefForListSecuritySensitiveDataOverviews()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSecuritySensitiveDataOverviewsResponse), nil
	}
}

// ListSecuritySensitiveDataOverviewsInvoker 查询敏感数据发现概览结果(以分类和密级为单位)
func (c *DataArtsStudioClient) ListSecuritySensitiveDataOverviewsInvoker(request *model.ListSecuritySensitiveDataOverviewsRequest) *ListSecuritySensitiveDataOverviewsInvoker {
	requestDef := GenReqDefForListSecuritySensitiveDataOverviews()
	return &ListSecuritySensitiveDataOverviewsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSubjectLevels 获取主题层级
//
// 获取主题层级。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListSubjectLevels(request *model.ListSubjectLevelsRequest) (*model.ListSubjectLevelsResponse, error) {
	requestDef := GenReqDefForListSubjectLevels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSubjectLevelsResponse), nil
	}
}

// ListSubjectLevelsInvoker 获取主题层级
func (c *DataArtsStudioClient) ListSubjectLevelsInvoker(request *model.ListSubjectLevelsRequest) *ListSubjectLevelsInvoker {
	requestDef := GenReqDefForListSubjectLevels()
	return &ListSubjectLevelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableModelRelations 查询模型下所有关系
//
// 查询模型下所有关系。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListTableModelRelations(request *model.ListTableModelRelationsRequest) (*model.ListTableModelRelationsResponse, error) {
	requestDef := GenReqDefForListTableModelRelations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableModelRelationsResponse), nil
	}
}

// ListTableModelRelationsInvoker 查询模型下所有关系
func (c *DataArtsStudioClient) ListTableModelRelationsInvoker(request *model.ListTableModelRelationsRequest) *ListTableModelRelationsInvoker {
	requestDef := GenReqDefForListTableModelRelations()
	return &ListTableModelRelationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTableModels 查找表模型列表
//
// 通过中英文名称、创建者、审核人、状态、修改时间分页查找关系建模中的表模型信息，包括逻辑实体、物理表和其属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListTableModels(request *model.ListTableModelsRequest) (*model.ListTableModelsResponse, error) {
	requestDef := GenReqDefForListTableModels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTableModelsResponse), nil
	}
}

// ListTableModelsInvoker 查找表模型列表
func (c *DataArtsStudioClient) ListTableModelsInvoker(request *model.ListTableModelsRequest) *ListTableModelsInvoker {
	requestDef := GenReqDefForListTableModels()
	return &ListTableModelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaceRoles 获取工作空间用户角色
//
// 获取工作空间用户角色
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListWorkspaceRoles(request *model.ListWorkspaceRolesRequest) (*model.ListWorkspaceRolesResponse, error) {
	requestDef := GenReqDefForListWorkspaceRoles()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspaceRolesResponse), nil
	}
}

// ListWorkspaceRolesInvoker 获取工作空间用户角色
func (c *DataArtsStudioClient) ListWorkspaceRolesInvoker(request *model.ListWorkspaceRolesRequest) *ListWorkspaceRolesInvoker {
	requestDef := GenReqDefForListWorkspaceRoles()
	return &ListWorkspaceRolesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaces 获取模型
//
// 获取当前空间下的全部模型信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListWorkspaces(request *model.ListWorkspacesRequest) (*model.ListWorkspacesResponse, error) {
	requestDef := GenReqDefForListWorkspaces()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspacesResponse), nil
	}
}

// ListWorkspacesInvoker 获取模型
func (c *DataArtsStudioClient) ListWorkspacesInvoker(request *model.ListWorkspacesRequest) *ListWorkspacesInvoker {
	requestDef := GenReqDefForListWorkspaces()
	return &ListWorkspacesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListWorkspaceusers 获取工作空间用户信息
//
// 获取工作空间用户信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListWorkspaceusers(request *model.ListWorkspaceusersRequest) (*model.ListWorkspaceusersResponse, error) {
	requestDef := GenReqDefForListWorkspaceusers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWorkspaceusersResponse), nil
	}
}

// ListWorkspaceusersInvoker 获取工作空间用户信息
func (c *DataArtsStudioClient) ListWorkspaceusersInvoker(request *model.ListWorkspaceusersRequest) *ListWorkspaceusersInvoker {
	requestDef := GenReqDefForListWorkspaceusers()
	return &ListWorkspaceusersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateApi 批量移动api至新目录
//
// 批量移动api至新目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) MigrateApi(request *model.MigrateApiRequest) (*model.MigrateApiResponse, error) {
	requestDef := GenReqDefForMigrateApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateApiResponse), nil
	}
}

// MigrateApiInvoker 批量移动api至新目录
func (c *DataArtsStudioClient) MigrateApiInvoker(request *model.MigrateApiRequest) *MigrateApiInvoker {
	requestDef := GenReqDefForMigrateApi()
	return &MigrateApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// MigrateCatalog 移动当前目录至新目录
//
// 移动当前目录至新目录。
// 移动目录的的同时会移动其下的所有子目录与api。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) MigrateCatalog(request *model.MigrateCatalogRequest) (*model.MigrateCatalogResponse, error) {
	requestDef := GenReqDefForMigrateCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.MigrateCatalogResponse), nil
	}
}

// MigrateCatalogInvoker 移动当前目录至新目录
func (c *DataArtsStudioClient) MigrateCatalogInvoker(request *model.MigrateCatalogRequest) *MigrateCatalogInvoker {
	requestDef := GenReqDefForMigrateCatalog()
	return &MigrateCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ModifyCustomizedFields 修改自定义项
//
// 修改自定义项（包括表自定义项、属性自定义项、主题自定义项、业务指标自定义项）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ModifyCustomizedFields(request *model.ModifyCustomizedFieldsRequest) (*model.ModifyCustomizedFieldsResponse, error) {
	requestDef := GenReqDefForModifyCustomizedFields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ModifyCustomizedFieldsResponse), nil
	}
}

// ModifyCustomizedFieldsInvoker 修改自定义项
func (c *DataArtsStudioClient) ModifyCustomizedFieldsInvoker(request *model.ModifyCustomizedFieldsRequest) *ModifyCustomizedFieldsInvoker {
	requestDef := GenReqDefForModifyCustomizedFields()
	return &ModifyCustomizedFieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ParseUserBehavior 用户行为分析
//
// 用户行为分析
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ParseUserBehavior(request *model.ParseUserBehaviorRequest) (*model.ParseUserBehaviorResponse, error) {
	requestDef := GenReqDefForParseUserBehavior()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ParseUserBehaviorResponse), nil
	}
}

// ParseUserBehaviorInvoker 用户行为分析
func (c *DataArtsStudioClient) ParseUserBehaviorInvoker(request *model.ParseUserBehaviorRequest) *ParseUserBehaviorInvoker {
	requestDef := GenReqDefForParseUserBehavior()
	return &ParseUserBehaviorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PayForDgcOneKey DataArtsStudio实例一键购买接口
//
// # DataArtsStudio实例一键购买接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) PayForDgcOneKey(request *model.PayForDgcOneKeyRequest) (*model.PayForDgcOneKeyResponse, error) {
	requestDef := GenReqDefForPayForDgcOneKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PayForDgcOneKeyResponse), nil
	}
}

// PayForDgcOneKeyInvoker DataArtsStudio实例一键购买接口
func (c *DataArtsStudioClient) PayForDgcOneKeyInvoker(request *model.PayForDgcOneKeyRequest) *PayForDgcOneKeyInvoker {
	requestDef := GenReqDefForPayForDgcOneKey()
	return &PayForDgcOneKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveDesignEntityTags 删除标签
//
// 根据资产（表或属性）的ID删除资产标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) RemoveDesignEntityTags(request *model.RemoveDesignEntityTagsRequest) (*model.RemoveDesignEntityTagsResponse, error) {
	requestDef := GenReqDefForRemoveDesignEntityTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveDesignEntityTagsResponse), nil
	}
}

// RemoveDesignEntityTagsInvoker 删除标签
func (c *DataArtsStudioClient) RemoveDesignEntityTagsInvoker(request *model.RemoveDesignEntityTagsRequest) *RemoveDesignEntityTagsInvoker {
	requestDef := GenReqDefForRemoveDesignEntityTags()
	return &RemoveDesignEntityTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RemoveDesignQualityInfos 清空质量规则
//
// 清空表的质量规则。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) RemoveDesignQualityInfos(request *model.RemoveDesignQualityInfosRequest) (*model.RemoveDesignQualityInfosResponse, error) {
	requestDef := GenReqDefForRemoveDesignQualityInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RemoveDesignQualityInfosResponse), nil
	}
}

// RemoveDesignQualityInfosInvoker 清空质量规则
func (c *DataArtsStudioClient) RemoveDesignQualityInfosInvoker(request *model.RemoveDesignQualityInfosRequest) *RemoveDesignQualityInfosInvoker {
	requestDef := GenReqDefForRemoveDesignQualityInfos()
	return &RemoveDesignQualityInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RenewDataProfile 指定字段采集概要
//
// 指定字段采集概要信息接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) RenewDataProfile(request *model.RenewDataProfileRequest) (*model.RenewDataProfileResponse, error) {
	requestDef := GenReqDefForRenewDataProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RenewDataProfileResponse), nil
	}
}

// RenewDataProfileInvoker 指定字段采集概要
func (c *DataArtsStudioClient) RenewDataProfileInvoker(request *model.RenewDataProfileRequest) *RenewDataProfileInvoker {
	requestDef := GenReqDefForRenewDataProfile()
	return &RenewDataProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ResetLinkAttributeAndStandard 关联属性与数据标准
//
// 关联属性与数据标准。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ResetLinkAttributeAndStandard(request *model.ResetLinkAttributeAndStandardRequest) (*model.ResetLinkAttributeAndStandardResponse, error) {
	requestDef := GenReqDefForResetLinkAttributeAndStandard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetLinkAttributeAndStandardResponse), nil
	}
}

// ResetLinkAttributeAndStandardInvoker 关联属性与数据标准
func (c *DataArtsStudioClient) ResetLinkAttributeAndStandardInvoker(request *model.ResetLinkAttributeAndStandardRequest) *ResetLinkAttributeAndStandardInvoker {
	requestDef := GenReqDefForResetLinkAttributeAndStandard()
	return &ResetLinkAttributeAndStandardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RollbackApproval 撤回审批单
//
// 撤回审批单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) RollbackApproval(request *model.RollbackApprovalRequest) (*model.RollbackApprovalResponse, error) {
	requestDef := GenReqDefForRollbackApproval()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RollbackApprovalResponse), nil
	}
}

// RollbackApprovalInvoker 撤回审批单
func (c *DataArtsStudioClient) RollbackApprovalInvoker(request *model.RollbackApprovalRequest) *RollbackApprovalInvoker {
	requestDef := GenReqDefForRollbackApproval()
	return &RollbackApprovalInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchApprovals 获取审批单
//
// 获取审批单。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchApprovals(request *model.SearchApprovalsRequest) (*model.SearchApprovalsResponse, error) {
	requestDef := GenReqDefForSearchApprovals()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchApprovalsResponse), nil
	}
}

// SearchApprovalsInvoker 获取审批单
func (c *DataArtsStudioClient) SearchApprovalsInvoker(request *model.SearchApprovalsRequest) *SearchApprovalsInvoker {
	requestDef := GenReqDefForSearchApprovals()
	return &SearchApprovalsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchAtomicIndexes 查找原子指标
//
// 通过中英文名称、创建者、审核人、状态、修改时间分页查找原子指标信息看，中英文名称支持模糊查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchAtomicIndexes(request *model.SearchAtomicIndexesRequest) (*model.SearchAtomicIndexesResponse, error) {
	requestDef := GenReqDefForSearchAtomicIndexes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAtomicIndexesResponse), nil
	}
}

// SearchAtomicIndexesInvoker 查找原子指标
func (c *DataArtsStudioClient) SearchAtomicIndexesInvoker(request *model.SearchAtomicIndexesRequest) *SearchAtomicIndexesInvoker {
	requestDef := GenReqDefForSearchAtomicIndexes()
	return &SearchAtomicIndexesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchAuthorizeApp 查询API已授权的APP
//
// 查询API已授权的APP。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchAuthorizeApp(request *model.SearchAuthorizeAppRequest) (*model.SearchAuthorizeAppResponse, error) {
	requestDef := GenReqDefForSearchAuthorizeApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchAuthorizeAppResponse), nil
	}
}

// SearchAuthorizeAppInvoker 查询API已授权的APP
func (c *DataArtsStudioClient) SearchAuthorizeAppInvoker(request *model.SearchAuthorizeAppRequest) *SearchAuthorizeAppInvoker {
	requestDef := GenReqDefForSearchAuthorizeApp()
	return &SearchAuthorizeAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchBindApi 查询APP已拥有授权的API
//
// 查询APP已拥有授权的API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchBindApi(request *model.SearchBindApiRequest) (*model.SearchBindApiResponse, error) {
	requestDef := GenReqDefForSearchBindApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchBindApiResponse), nil
	}
}

// SearchBindApiInvoker 查询APP已拥有授权的API
func (c *DataArtsStudioClient) SearchBindApiInvoker(request *model.SearchBindApiRequest) *SearchBindApiInvoker {
	requestDef := GenReqDefForSearchBindApi()
	return &SearchBindApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCatalogs 查询流程架构列表
//
// 查询流程架构列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchCatalogs(request *model.SearchCatalogsRequest) (*model.SearchCatalogsResponse, error) {
	requestDef := GenReqDefForSearchCatalogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCatalogsResponse), nil
	}
}

// SearchCatalogsInvoker 查询流程架构列表
func (c *DataArtsStudioClient) SearchCatalogsInvoker(request *model.SearchCatalogsRequest) *SearchCatalogsInvoker {
	requestDef := GenReqDefForSearchCatalogs()
	return &SearchCatalogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCodeTableValues 查看码表字段值
//
// 查看码表字段值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchCodeTableValues(request *model.SearchCodeTableValuesRequest) (*model.SearchCodeTableValuesResponse, error) {
	requestDef := GenReqDefForSearchCodeTableValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCodeTableValuesResponse), nil
	}
}

// SearchCodeTableValuesInvoker 查看码表字段值
func (c *DataArtsStudioClient) SearchCodeTableValuesInvoker(request *model.SearchCodeTableValuesRequest) *SearchCodeTableValuesInvoker {
	requestDef := GenReqDefForSearchCodeTableValues()
	return &SearchCodeTableValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCodeTables 查询码表列表
//
// 查询码表列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchCodeTables(request *model.SearchCodeTablesRequest) (*model.SearchCodeTablesResponse, error) {
	requestDef := GenReqDefForSearchCodeTables()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCodeTablesResponse), nil
	}
}

// SearchCodeTablesInvoker 查询码表列表
func (c *DataArtsStudioClient) SearchCodeTablesInvoker(request *model.SearchCodeTablesRequest) *SearchCodeTablesInvoker {
	requestDef := GenReqDefForSearchCodeTables()
	return &SearchCodeTablesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchCustomizedFields 查询自定义项
//
// 查询自定义项（包括表自定义项、属性自定义项、主题自定义项、业务指标自定义项）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchCustomizedFields(request *model.SearchCustomizedFieldsRequest) (*model.SearchCustomizedFieldsResponse, error) {
	requestDef := GenReqDefForSearchCustomizedFields()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchCustomizedFieldsResponse), nil
	}
}

// SearchCustomizedFieldsInvoker 查询自定义项
func (c *DataArtsStudioClient) SearchCustomizedFieldsInvoker(request *model.SearchCustomizedFieldsRequest) *SearchCustomizedFieldsInvoker {
	requestDef := GenReqDefForSearchCustomizedFields()
	return &SearchCustomizedFieldsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDesignLatestApprovalDiff 获取下展信息与已发布实体的差异
//
// 当已发布的实体被编辑时，其会生成下展，该接口用于获取下展信息与已发布实体的差异。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchDesignLatestApprovalDiff(request *model.SearchDesignLatestApprovalDiffRequest) (*model.SearchDesignLatestApprovalDiffResponse, error) {
	requestDef := GenReqDefForSearchDesignLatestApprovalDiff()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDesignLatestApprovalDiffResponse), nil
	}
}

// SearchDesignLatestApprovalDiffInvoker 获取下展信息与已发布实体的差异
func (c *DataArtsStudioClient) SearchDesignLatestApprovalDiffInvoker(request *model.SearchDesignLatestApprovalDiffRequest) *SearchDesignLatestApprovalDiffInvoker {
	requestDef := GenReqDefForSearchDesignLatestApprovalDiff()
	return &SearchDesignLatestApprovalDiffInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDwByType 获取数据连接信息
//
// 获取指定类型下的数据连接信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchDwByType(request *model.SearchDwByTypeRequest) (*model.SearchDwByTypeResponse, error) {
	requestDef := GenReqDefForSearchDwByType()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDwByTypeResponse), nil
	}
}

// SearchDwByTypeInvoker 获取数据连接信息
func (c *DataArtsStudioClient) SearchDwByTypeInvoker(request *model.SearchDwByTypeRequest) *SearchDwByTypeInvoker {
	requestDef := GenReqDefForSearchDwByType()
	return &SearchDwByTypeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchFieldsForRelation 查询目的表和字段(待下线)
//
// 查询目的表和字段(待下线)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchFieldsForRelation(request *model.SearchFieldsForRelationRequest) (*model.SearchFieldsForRelationResponse, error) {
	requestDef := GenReqDefForSearchFieldsForRelation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchFieldsForRelationResponse), nil
	}
}

// SearchFieldsForRelationInvoker 查询目的表和字段(待下线)
func (c *DataArtsStudioClient) SearchFieldsForRelationInvoker(request *model.SearchFieldsForRelationRequest) *SearchFieldsForRelationInvoker {
	requestDef := GenReqDefForSearchFieldsForRelation()
	return &SearchFieldsForRelationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchIdByPath 通过路径获取id
//
// 通过路径获取id。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchIdByPath(request *model.SearchIdByPathRequest) (*model.SearchIdByPathResponse, error) {
	requestDef := GenReqDefForSearchIdByPath()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchIdByPathResponse), nil
	}
}

// SearchIdByPathInvoker 通过路径获取id
func (c *DataArtsStudioClient) SearchIdByPathInvoker(request *model.SearchIdByPathRequest) *SearchIdByPathInvoker {
	requestDef := GenReqDefForSearchIdByPath()
	return &SearchIdByPathInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchSgcComputeDimensions 获取计算维度成本列表信息
//
// 获取计算维度成本列表信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchSgcComputeDimensions(request *model.SearchSgcComputeDimensionsRequest) (*model.SearchSgcComputeDimensionsResponse, error) {
	requestDef := GenReqDefForSearchSgcComputeDimensions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchSgcComputeDimensionsResponse), nil
	}
}

// SearchSgcComputeDimensionsInvoker 获取计算维度成本列表信息
func (c *DataArtsStudioClient) SearchSgcComputeDimensionsInvoker(request *model.SearchSgcComputeDimensionsRequest) *SearchSgcComputeDimensionsInvoker {
	requestDef := GenReqDefForSearchSgcComputeDimensions()
	return &SearchSgcComputeDimensionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchSubject 查找主题列表
//
// 通过名称（支持模糊查询）、创建者、责任人、状态、修改时间分页查找主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchSubject(request *model.SearchSubjectRequest) (*model.SearchSubjectResponse, error) {
	requestDef := GenReqDefForSearchSubject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchSubjectResponse), nil
	}
}

// SearchSubjectInvoker 查找主题列表
func (c *DataArtsStudioClient) SearchSubjectInvoker(request *model.SearchSubjectRequest) *SearchSubjectInvoker {
	requestDef := GenReqDefForSearchSubject()
	return &SearchSubjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchSubjectNew 查找主题列表(新)
//
// 查找主题(新)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchSubjectNew(request *model.SearchSubjectNewRequest) (*model.SearchSubjectNewResponse, error) {
	requestDef := GenReqDefForSearchSubjectNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchSubjectNewResponse), nil
	}
}

// SearchSubjectNewInvoker 查找主题列表(新)
func (c *DataArtsStudioClient) SearchSubjectNewInvoker(request *model.SearchSubjectNewRequest) *SearchSubjectNewInvoker {
	requestDef := GenReqDefForSearchSubjectNew()
	return &SearchSubjectNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchVersions 查找版本信息
//
// 通过名称、创建者、修改时间查找版本信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchVersions(request *model.SearchVersionsRequest) (*model.SearchVersionsResponse, error) {
	requestDef := GenReqDefForSearchVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchVersionsResponse), nil
	}
}

// SearchVersionsInvoker 查找版本信息
func (c *DataArtsStudioClient) SearchVersionsInvoker(request *model.SearchVersionsRequest) *SearchVersionsInvoker {
	requestDef := GenReqDefForSearchVersions()
	return &SearchVersionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SetFactoryJobTags 设置作业标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SetFactoryJobTags(request *model.SetFactoryJobTagsRequest) (*model.SetFactoryJobTagsResponse, error) {
	requestDef := GenReqDefForSetFactoryJobTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetFactoryJobTagsResponse), nil
	}
}

// SetFactoryJobTagsInvoker 设置作业标签
func (c *DataArtsStudioClient) SetFactoryJobTagsInvoker(request *model.SetFactoryJobTagsRequest) *SetFactoryJobTagsInvoker {
	requestDef := GenReqDefForSetFactoryJobTags()
	return &SetFactoryJobTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAggregationLogicTableById 查看汇总表详情
//
// 通过ID查看汇总表的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAggregationLogicTableById(request *model.ShowAggregationLogicTableByIdRequest) (*model.ShowAggregationLogicTableByIdResponse, error) {
	requestDef := GenReqDefForShowAggregationLogicTableById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAggregationLogicTableByIdResponse), nil
	}
}

// ShowAggregationLogicTableByIdInvoker 查看汇总表详情
func (c *DataArtsStudioClient) ShowAggregationLogicTableByIdInvoker(request *model.ShowAggregationLogicTableByIdRequest) *ShowAggregationLogicTableByIdInvoker {
	requestDef := GenReqDefForShowAggregationLogicTableById()
	return &ShowAggregationLogicTableByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApiDashboard 查询指定api 仪表板数据详情
//
// 查询指定api 仪表板数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApiDashboard(request *model.ShowApiDashboardRequest) (*model.ShowApiDashboardResponse, error) {
	requestDef := GenReqDefForShowApiDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiDashboardResponse), nil
	}
}

// ShowApiDashboardInvoker 查询指定api 仪表板数据详情
func (c *DataArtsStudioClient) ShowApiDashboardInvoker(request *model.ShowApiDashboardRequest) *ShowApiDashboardInvoker {
	requestDef := GenReqDefForShowApiDashboard()
	return &ShowApiDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApisDashboard 查询api 仪表板数据详情
//
// 查询api 仪表板数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApisDashboard(request *model.ShowApisDashboardRequest) (*model.ShowApisDashboardResponse, error) {
	requestDef := GenReqDefForShowApisDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApisDashboardResponse), nil
	}
}

// ShowApisDashboardInvoker 查询api 仪表板数据详情
func (c *DataArtsStudioClient) ShowApisDashboardInvoker(request *model.ShowApisDashboardRequest) *ShowApisDashboardInvoker {
	requestDef := GenReqDefForShowApisDashboard()
	return &ShowApisDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApisDetail 查询api 统计数据详情
//
// 查询api 统计数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApisDetail(request *model.ShowApisDetailRequest) (*model.ShowApisDetailResponse, error) {
	requestDef := GenReqDefForShowApisDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApisDetailResponse), nil
	}
}

// ShowApisDetailInvoker 查询api 统计数据详情
func (c *DataArtsStudioClient) ShowApisDetailInvoker(request *model.ShowApisDetailRequest) *ShowApisDetailInvoker {
	requestDef := GenReqDefForShowApisDetail()
	return &ShowApisDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApisOverview 查询统计用户相关的总览开发指标
//
// 查询统计用户相关的总览开发指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApisOverview(request *model.ShowApisOverviewRequest) (*model.ShowApisOverviewResponse, error) {
	requestDef := GenReqDefForShowApisOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApisOverviewResponse), nil
	}
}

// ShowApisOverviewInvoker 查询统计用户相关的总览开发指标
func (c *DataArtsStudioClient) ShowApisOverviewInvoker(request *model.ShowApisOverviewRequest) *ShowApisOverviewInvoker {
	requestDef := GenReqDefForShowApisOverview()
	return &ShowApisOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppInfo 查询应用详情
//
// 查询应用详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAppInfo(request *model.ShowAppInfoRequest) (*model.ShowAppInfoResponse, error) {
	requestDef := GenReqDefForShowAppInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppInfoResponse), nil
	}
}

// ShowAppInfoInvoker 查询应用详情
func (c *DataArtsStudioClient) ShowAppInfoInvoker(request *model.ShowAppInfoRequest) *ShowAppInfoInvoker {
	requestDef := GenReqDefForShowAppInfo()
	return &ShowAppInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApplyDetail 获取申请详情
//
// 获取申请详情。
// 此功能仅用作信息详情展示，不用做业务处理，因此不展示编号等后台参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApplyDetail(request *model.ShowApplyDetailRequest) (*model.ShowApplyDetailResponse, error) {
	requestDef := GenReqDefForShowApplyDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApplyDetailResponse), nil
	}
}

// ShowApplyDetailInvoker 获取申请详情
func (c *DataArtsStudioClient) ShowApplyDetailInvoker(request *model.ShowApplyDetailRequest) *ShowApplyDetailInvoker {
	requestDef := GenReqDefForShowApplyDetail()
	return &ShowApplyDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppsDashboard 查询app 仪表板数据详情
//
// 查询app 仪表板数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAppsDashboard(request *model.ShowAppsDashboardRequest) (*model.ShowAppsDashboardResponse, error) {
	requestDef := GenReqDefForShowAppsDashboard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppsDashboardResponse), nil
	}
}

// ShowAppsDashboardInvoker 查询app 仪表板数据详情
func (c *DataArtsStudioClient) ShowAppsDashboardInvoker(request *model.ShowAppsDashboardRequest) *ShowAppsDashboardInvoker {
	requestDef := GenReqDefForShowAppsDashboard()
	return &ShowAppsDashboardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppsDetail 查询app 统计数据详情
//
// 查询app 统计数据详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAppsDetail(request *model.ShowAppsDetailRequest) (*model.ShowAppsDetailResponse, error) {
	requestDef := GenReqDefForShowAppsDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppsDetailResponse), nil
	}
}

// ShowAppsDetailInvoker 查询app 统计数据详情
func (c *DataArtsStudioClient) ShowAppsDetailInvoker(request *model.ShowAppsDetailRequest) *ShowAppsDetailInvoker {
	requestDef := GenReqDefForShowAppsDetail()
	return &ShowAppsDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppsOverview 查询统计用户相关的总览调用指标
//
// 查询统计用户相关的总览调用指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAppsOverview(request *model.ShowAppsOverviewRequest) (*model.ShowAppsOverviewResponse, error) {
	requestDef := GenReqDefForShowAppsOverview()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppsOverviewResponse), nil
	}
}

// ShowAppsOverviewInvoker 查询统计用户相关的总览调用指标
func (c *DataArtsStudioClient) ShowAppsOverviewInvoker(request *model.ShowAppsOverviewRequest) *ShowAppsOverviewInvoker {
	requestDef := GenReqDefForShowAppsOverview()
	return &ShowAppsOverviewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAtomicIndexById 查看原子指标详情
//
// 通过ID获取原子指标详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowAtomicIndexById(request *model.ShowAtomicIndexByIdRequest) (*model.ShowAtomicIndexByIdResponse, error) {
	requestDef := GenReqDefForShowAtomicIndexById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAtomicIndexByIdResponse), nil
	}
}

// ShowAtomicIndexByIdInvoker 查看原子指标详情
func (c *DataArtsStudioClient) ShowAtomicIndexByIdInvoker(request *model.ShowAtomicIndexByIdRequest) *ShowAtomicIndexByIdInvoker {
	requestDef := GenReqDefForShowAtomicIndexById()
	return &ShowAtomicIndexByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBizCatalogDetail 查找流程架构详情
//
// 查找流程架构详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowBizCatalogDetail(request *model.ShowBizCatalogDetailRequest) (*model.ShowBizCatalogDetailResponse, error) {
	requestDef := GenReqDefForShowBizCatalogDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBizCatalogDetailResponse), nil
	}
}

// ShowBizCatalogDetailInvoker 查找流程架构详情
func (c *DataArtsStudioClient) ShowBizCatalogDetailInvoker(request *model.ShowBizCatalogDetailRequest) *ShowBizCatalogDetailInvoker {
	requestDef := GenReqDefForShowBizCatalogDetail()
	return &ShowBizCatalogDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBizMetricById 查看指标详情
//
// 通过ID查看指标的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowBizMetricById(request *model.ShowBizMetricByIdRequest) (*model.ShowBizMetricByIdResponse, error) {
	requestDef := GenReqDefForShowBizMetricById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBizMetricByIdResponse), nil
	}
}

// ShowBizMetricByIdInvoker 查看指标详情
func (c *DataArtsStudioClient) ShowBizMetricByIdInvoker(request *model.ShowBizMetricByIdRequest) *ShowBizMetricByIdInvoker {
	requestDef := GenReqDefForShowBizMetricById()
	return &ShowBizMetricByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBusinessAssets 查询业务资产
//
// 业务资产查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowBusinessAssets(request *model.ShowBusinessAssetsRequest) (*model.ShowBusinessAssetsResponse, error) {
	requestDef := GenReqDefForShowBusinessAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBusinessAssetsResponse), nil
	}
}

// ShowBusinessAssetsInvoker 查询业务资产
func (c *DataArtsStudioClient) ShowBusinessAssetsInvoker(request *model.ShowBusinessAssetsRequest) *ShowBusinessAssetsInvoker {
	requestDef := GenReqDefForShowBusinessAssets()
	return &ShowBusinessAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBusinessAssetsStatistic 获取业务资产统计信息
//
// 获取业务资产统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowBusinessAssetsStatistic(request *model.ShowBusinessAssetsStatisticRequest) (*model.ShowBusinessAssetsStatisticResponse, error) {
	requestDef := GenReqDefForShowBusinessAssetsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBusinessAssetsStatisticResponse), nil
	}
}

// ShowBusinessAssetsStatisticInvoker 获取业务资产统计信息
func (c *DataArtsStudioClient) ShowBusinessAssetsStatisticInvoker(request *model.ShowBusinessAssetsStatisticRequest) *ShowBusinessAssetsStatisticInvoker {
	requestDef := GenReqDefForShowBusinessAssetsStatistic()
	return &ShowBusinessAssetsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBusinessAssetsTree 查询业务资产目录树
//
// 逐级查询业务资产目录树,包含数据规范同步过来的业务对象和逻辑实体。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowBusinessAssetsTree(request *model.ShowBusinessAssetsTreeRequest) (*model.ShowBusinessAssetsTreeResponse, error) {
	requestDef := GenReqDefForShowBusinessAssetsTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBusinessAssetsTreeResponse), nil
	}
}

// ShowBusinessAssetsTreeInvoker 查询业务资产目录树
func (c *DataArtsStudioClient) ShowBusinessAssetsTreeInvoker(request *model.ShowBusinessAssetsTreeRequest) *ShowBusinessAssetsTreeInvoker {
	requestDef := GenReqDefForShowBusinessAssetsTree()
	return &ShowBusinessAssetsTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCatalogDetail 查询服务目录
//
// 查询服务目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowCatalogDetail(request *model.ShowCatalogDetailRequest) (*model.ShowCatalogDetailResponse, error) {
	requestDef := GenReqDefForShowCatalogDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCatalogDetailResponse), nil
	}
}

// ShowCatalogDetailInvoker 查询服务目录
func (c *DataArtsStudioClient) ShowCatalogDetailInvoker(request *model.ShowCatalogDetailRequest) *ShowCatalogDetailInvoker {
	requestDef := GenReqDefForShowCatalogDetail()
	return &ShowCatalogDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCodeTableById 查看码表详情
//
// 通过ID查看码表的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowCodeTableById(request *model.ShowCodeTableByIdRequest) (*model.ShowCodeTableByIdResponse, error) {
	requestDef := GenReqDefForShowCodeTableById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCodeTableByIdResponse), nil
	}
}

// ShowCodeTableByIdInvoker 查看码表详情
func (c *DataArtsStudioClient) ShowCodeTableByIdInvoker(request *model.ShowCodeTableByIdRequest) *ShowCodeTableByIdInvoker {
	requestDef := GenReqDefForShowCodeTableById()
	return &ShowCodeTableByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCompoundMetricById 查看复合指标详情
//
// 通过ID获取复合指标详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowCompoundMetricById(request *model.ShowCompoundMetricByIdRequest) (*model.ShowCompoundMetricByIdResponse, error) {
	requestDef := GenReqDefForShowCompoundMetricById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCompoundMetricByIdResponse), nil
	}
}

// ShowCompoundMetricByIdInvoker 查看复合指标详情
func (c *DataArtsStudioClient) ShowCompoundMetricByIdInvoker(request *model.ShowCompoundMetricByIdRequest) *ShowCompoundMetricByIdInvoker {
	requestDef := GenReqDefForShowCompoundMetricById()
	return &ShowCompoundMetricByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConditionById 查看限定详情
//
// 通过ID查看限定详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowConditionById(request *model.ShowConditionByIdRequest) (*model.ShowConditionByIdResponse, error) {
	requestDef := GenReqDefForShowConditionById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConditionByIdResponse), nil
	}
}

// ShowConditionByIdInvoker 查看限定详情
func (c *DataArtsStudioClient) ShowConditionByIdInvoker(request *model.ShowConditionByIdRequest) *ShowConditionByIdInvoker {
	requestDef := GenReqDefForShowConditionById()
	return &ShowConditionByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowConsistencyTaskDetail 获取对账作业详情
//
// 获取对账作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowConsistencyTaskDetail(request *model.ShowConsistencyTaskDetailRequest) (*model.ShowConsistencyTaskDetailResponse, error) {
	requestDef := GenReqDefForShowConsistencyTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowConsistencyTaskDetailResponse), nil
	}
}

// ShowConsistencyTaskDetailInvoker 获取对账作业详情
func (c *DataArtsStudioClient) ShowConsistencyTaskDetailInvoker(request *model.ShowConsistencyTaskDetailRequest) *ShowConsistencyTaskDetailInvoker {
	requestDef := GenReqDefForShowConsistencyTaskDetail()
	return &ShowConsistencyTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataDetail 资产详情(邀测)
//
// 资产详情接口，该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDataDetail(request *model.ShowDataDetailRequest) (*model.ShowDataDetailResponse, error) {
	requestDef := GenReqDefForShowDataDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataDetailResponse), nil
	}
}

// ShowDataDetailInvoker 资产详情(邀测)
func (c *DataArtsStudioClient) ShowDataDetailInvoker(request *model.ShowDataDetailRequest) *ShowDataDetailInvoker {
	requestDef := GenReqDefForShowDataDetail()
	return &ShowDataDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataProfile 资产信息
//
// 查询概要
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDataProfile(request *model.ShowDataProfileRequest) (*model.ShowDataProfileResponse, error) {
	requestDef := GenReqDefForShowDataProfile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataProfileResponse), nil
	}
}

// ShowDataProfileInvoker 资产信息
func (c *DataArtsStudioClient) ShowDataProfileInvoker(request *model.ShowDataProfileRequest) *ShowDataProfileInvoker {
	requestDef := GenReqDefForShowDataProfile()
	return &ShowDataProfileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataServiceInstance 查询集群详情信息
//
// 查询集群详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDataServiceInstance(request *model.ShowDataServiceInstanceRequest) (*model.ShowDataServiceInstanceResponse, error) {
	requestDef := GenReqDefForShowDataServiceInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataServiceInstanceResponse), nil
	}
}

// ShowDataServiceInstanceInvoker 查询集群详情信息
func (c *DataArtsStudioClient) ShowDataServiceInstanceInvoker(request *model.ShowDataServiceInstanceRequest) *ShowDataServiceInstanceInvoker {
	requestDef := GenReqDefForShowDataServiceInstance()
	return &ShowDataServiceInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataSets 资产搜索
//
// 资产搜索
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDataSets(request *model.ShowDataSetsRequest) (*model.ShowDataSetsResponse, error) {
	requestDef := GenReqDefForShowDataSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataSetsResponse), nil
	}
}

// ShowDataSetsInvoker 资产搜索
func (c *DataArtsStudioClient) ShowDataSetsInvoker(request *model.ShowDataSetsRequest) *ShowDataSetsInvoker {
	requestDef := GenReqDefForShowDataSets()
	return &ShowDataSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDataconnection 查询单个数据连接信息
//
// 查询单个数据连接信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDataconnection(request *model.ShowDataconnectionRequest) (*model.ShowDataconnectionResponse, error) {
	requestDef := GenReqDefForShowDataconnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDataconnectionResponse), nil
	}
}

// ShowDataconnectionInvoker 查询单个数据连接信息
func (c *DataArtsStudioClient) ShowDataconnectionInvoker(request *model.ShowDataconnectionRequest) *ShowDataconnectionInvoker {
	requestDef := GenReqDefForShowDataconnection()
	return &ShowDataconnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDatamapLineage 资产血缘(邀测)
//
// 资产血缘接口，该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDatamapLineage(request *model.ShowDatamapLineageRequest) (*model.ShowDatamapLineageResponse, error) {
	requestDef := GenReqDefForShowDatamapLineage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatamapLineageResponse), nil
	}
}

// ShowDatamapLineageInvoker 资产血缘(邀测)
func (c *DataArtsStudioClient) ShowDatamapLineageInvoker(request *model.ShowDatamapLineageRequest) *ShowDatamapLineageInvoker {
	requestDef := GenReqDefForShowDatamapLineage()
	return &ShowDatamapLineageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDerivativeIndexById 查看衍生指标详情
//
// 通过ID获取衍生详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDerivativeIndexById(request *model.ShowDerivativeIndexByIdRequest) (*model.ShowDerivativeIndexByIdResponse, error) {
	requestDef := GenReqDefForShowDerivativeIndexById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDerivativeIndexByIdResponse), nil
	}
}

// ShowDerivativeIndexByIdInvoker 查看衍生指标详情
func (c *DataArtsStudioClient) ShowDerivativeIndexByIdInvoker(request *model.ShowDerivativeIndexByIdRequest) *ShowDerivativeIndexByIdInvoker {
	requestDef := GenReqDefForShowDerivativeIndexById()
	return &ShowDerivativeIndexByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDimensionById 查看维度详情
//
// 通过ID查看维度详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDimensionById(request *model.ShowDimensionByIdRequest) (*model.ShowDimensionByIdResponse, error) {
	requestDef := GenReqDefForShowDimensionById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDimensionByIdResponse), nil
	}
}

// ShowDimensionByIdInvoker 查看维度详情
func (c *DataArtsStudioClient) ShowDimensionByIdInvoker(request *model.ShowDimensionByIdRequest) *ShowDimensionByIdInvoker {
	requestDef := GenReqDefForShowDimensionById()
	return &ShowDimensionByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowDimensionLogicTableById 查看维度表详情
//
// 通过ID查看维度表的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowDimensionLogicTableById(request *model.ShowDimensionLogicTableByIdRequest) (*model.ShowDimensionLogicTableByIdResponse, error) {
	requestDef := GenReqDefForShowDimensionLogicTableById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDimensionLogicTableByIdResponse), nil
	}
}

// ShowDimensionLogicTableByIdInvoker 查看维度表详情
func (c *DataArtsStudioClient) ShowDimensionLogicTableByIdInvoker(request *model.ShowDimensionLogicTableByIdRequest) *ShowDimensionLogicTableByIdInvoker {
	requestDef := GenReqDefForShowDimensionLogicTableById()
	return &ShowDimensionLogicTableByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEntities 查询技术资产
//
// 查询技术资产
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowEntities(request *model.ShowEntitiesRequest) (*model.ShowEntitiesResponse, error) {
	requestDef := GenReqDefForShowEntities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEntitiesResponse), nil
	}
}

// ShowEntitiesInvoker 查询技术资产
func (c *DataArtsStudioClient) ShowEntitiesInvoker(request *model.ShowEntitiesRequest) *ShowEntitiesInvoker {
	requestDef := GenReqDefForShowEntities()
	return &ShowEntitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowEntityInfoByGuid 根据guid获取资产详情
//
// 根据表guid可以获取表的详情信息，表的详情信息包含column的信息，也可以根据column的guid直接获取column的信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowEntityInfoByGuid(request *model.ShowEntityInfoByGuidRequest) (*model.ShowEntityInfoByGuidResponse, error) {
	requestDef := GenReqDefForShowEntityInfoByGuid()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowEntityInfoByGuidResponse), nil
	}
}

// ShowEntityInfoByGuidInvoker 根据guid获取资产详情
func (c *DataArtsStudioClient) ShowEntityInfoByGuidInvoker(request *model.ShowEntityInfoByGuidRequest) *ShowEntityInfoByGuidInvoker {
	requestDef := GenReqDefForShowEntityInfoByGuid()
	return &ShowEntityInfoByGuidInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactLogicTableById 查看事实表详情
//
// 通过ID查看事实表的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowFactLogicTableById(request *model.ShowFactLogicTableByIdRequest) (*model.ShowFactLogicTableByIdResponse, error) {
	requestDef := GenReqDefForShowFactLogicTableById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactLogicTableByIdResponse), nil
	}
}

// ShowFactLogicTableByIdInvoker 查看事实表详情
func (c *DataArtsStudioClient) ShowFactLogicTableByIdInvoker(request *model.ShowFactLogicTableByIdRequest) *ShowFactLogicTableByIdInvoker {
	requestDef := GenReqDefForShowFactLogicTableById()
	return &ShowFactLogicTableByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactoryEnv 查询环境变量信息
//
// 查询环境变量信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowFactoryEnv(request *model.ShowFactoryEnvRequest) (*model.ShowFactoryEnvResponse, error) {
	requestDef := GenReqDefForShowFactoryEnv()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactoryEnvResponse), nil
	}
}

// ShowFactoryEnvInvoker 查询环境变量信息
func (c *DataArtsStudioClient) ShowFactoryEnvInvoker(request *model.ShowFactoryEnvRequest) *ShowFactoryEnvInvoker {
	requestDef := GenReqDefForShowFactoryEnv()
	return &ShowFactoryEnvInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactoryPackageDetail 查询指定发布包详情
//
// 查询指定发布包详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowFactoryPackageDetail(request *model.ShowFactoryPackageDetailRequest) (*model.ShowFactoryPackageDetailResponse, error) {
	requestDef := GenReqDefForShowFactoryPackageDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactoryPackageDetailResponse), nil
	}
}

// ShowFactoryPackageDetailInvoker 查询指定发布包详情
func (c *DataArtsStudioClient) ShowFactoryPackageDetailInvoker(request *model.ShowFactoryPackageDetailRequest) *ShowFactoryPackageDetailInvoker {
	requestDef := GenReqDefForShowFactoryPackageDetail()
	return &ShowFactoryPackageDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowFactorySupplementData 查询补数据实例
//
// 查询补数据实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowFactorySupplementData(request *model.ShowFactorySupplementDataRequest) (*model.ShowFactorySupplementDataResponse, error) {
	requestDef := GenReqDefForShowFactorySupplementData()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowFactorySupplementDataResponse), nil
	}
}

// ShowFactorySupplementDataInvoker 查询补数据实例
func (c *DataArtsStudioClient) ShowFactorySupplementDataInvoker(request *model.ShowFactorySupplementDataRequest) *ShowFactorySupplementDataInvoker {
	requestDef := GenReqDefForShowFactorySupplementData()
	return &ShowFactorySupplementDataInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlossaryList 查询标签列表
//
// 查询标签列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowGlossaryList(request *model.ShowGlossaryListRequest) (*model.ShowGlossaryListResponse, error) {
	requestDef := GenReqDefForShowGlossaryList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlossaryListResponse), nil
	}
}

// ShowGlossaryListInvoker 查询标签列表
func (c *DataArtsStudioClient) ShowGlossaryListInvoker(request *model.ShowGlossaryListRequest) *ShowGlossaryListInvoker {
	requestDef := GenReqDefForShowGlossaryList()
	return &ShowGlossaryListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceInfos 产出信息(邀测)
//
// 查询表相关的作业算子运行实例信息，该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowInstanceInfos(request *model.ShowInstanceInfosRequest) (*model.ShowInstanceInfosResponse, error) {
	requestDef := GenReqDefForShowInstanceInfos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceInfosResponse), nil
	}
}

// ShowInstanceInfosInvoker 产出信息(邀测)
func (c *DataArtsStudioClient) ShowInstanceInfosInvoker(request *model.ShowInstanceInfosRequest) *ShowInstanceInfosInvoker {
	requestDef := GenReqDefForShowInstanceInfos()
	return &ShowInstanceInfosInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceLog 获取任务日志
//
// 获取任务日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowInstanceLog(request *model.ShowInstanceLogRequest) (*model.ShowInstanceLogResponse, error) {
	requestDef := GenReqDefForShowInstanceLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceLogResponse), nil
	}
}

// ShowInstanceLogInvoker 获取任务日志
func (c *DataArtsStudioClient) ShowInstanceLogInvoker(request *model.ShowInstanceLogRequest) *ShowInstanceLogInvoker {
	requestDef := GenReqDefForShowInstanceLog()
	return &ShowInstanceLogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInstanceResult 获取实例结果
//
// 获取实例结果
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowInstanceResult(request *model.ShowInstanceResultRequest) (*model.ShowInstanceResultResponse, error) {
	requestDef := GenReqDefForShowInstanceResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInstanceResultResponse), nil
	}
}

// ShowInstanceResultInvoker 获取实例结果
func (c *DataArtsStudioClient) ShowInstanceResultInvoker(request *model.ShowInstanceResultRequest) *ShowInstanceResultInvoker {
	requestDef := GenReqDefForShowInstanceResult()
	return &ShowInstanceResultInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLineage 血缘查询
//
// 血缘查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowLineage(request *model.ShowLineageRequest) (*model.ShowLineageResponse, error) {
	requestDef := GenReqDefForShowLineage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLineageResponse), nil
	}
}

// ShowLineageInvoker 血缘查询
func (c *DataArtsStudioClient) ShowLineageInvoker(request *model.ShowLineageRequest) *ShowLineageInvoker {
	requestDef := GenReqDefForShowLineage()
	return &ShowLineageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowLineageBulk 批量血缘(邀测)
//
// 批量血缘接口，根据作业算子分页批量查询血缘。该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowLineageBulk(request *model.ShowLineageBulkRequest) (*model.ShowLineageBulkResponse, error) {
	requestDef := GenReqDefForShowLineageBulk()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowLineageBulkResponse), nil
	}
}

// ShowLineageBulkInvoker 批量血缘(邀测)
func (c *DataArtsStudioClient) ShowLineageBulkInvoker(request *model.ShowLineageBulkRequest) *ShowLineageBulkInvoker {
	requestDef := GenReqDefForShowLineageBulk()
	return &ShowLineageBulkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMessageDetail 获取消息详情
//
// 获取消息详情。此功能仅用作信息详情展示，不用做业务处理，因此不展示编号等后台参数。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowMessageDetail(request *model.ShowMessageDetailRequest) (*model.ShowMessageDetailResponse, error) {
	requestDef := GenReqDefForShowMessageDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMessageDetailResponse), nil
	}
}

// ShowMessageDetailInvoker 获取消息详情
func (c *DataArtsStudioClient) ShowMessageDetailInvoker(request *model.ShowMessageDetailRequest) *ShowMessageDetailInvoker {
	requestDef := GenReqDefForShowMessageDetail()
	return &ShowMessageDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricAssets 查询指标资产
//
// 指标资产查询接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowMetricAssets(request *model.ShowMetricAssetsRequest) (*model.ShowMetricAssetsResponse, error) {
	requestDef := GenReqDefForShowMetricAssets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricAssetsResponse), nil
	}
}

// ShowMetricAssetsInvoker 查询指标资产
func (c *DataArtsStudioClient) ShowMetricAssetsInvoker(request *model.ShowMetricAssetsRequest) *ShowMetricAssetsInvoker {
	requestDef := GenReqDefForShowMetricAssets()
	return &ShowMetricAssetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowMetricTree 查询指标资产目录树
//
// 查询指标资产目录树
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowMetricTree(request *model.ShowMetricTreeRequest) (*model.ShowMetricTreeResponse, error) {
	requestDef := GenReqDefForShowMetricTree()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowMetricTreeResponse), nil
	}
}

// ShowMetricTreeInvoker 查询指标资产目录树
func (c *DataArtsStudioClient) ShowMetricTreeInvoker(request *model.ShowMetricTreeRequest) *ShowMetricTreeInvoker {
	requestDef := GenReqDefForShowMetricTree()
	return &ShowMetricTreeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNodes 表关联作业算子列表(邀测)
//
// 查询表相关的作业算子列表，该接口功能处于邀测阶段，后续将随功能公测将逐步开放。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowNodes(request *model.ShowNodesRequest) (*model.ShowNodesResponse, error) {
	requestDef := GenReqDefForShowNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNodesResponse), nil
	}
}

// ShowNodesInvoker 表关联作业算子列表(邀测)
func (c *DataArtsStudioClient) ShowNodesInvoker(request *model.ShowNodesRequest) *ShowNodesInvoker {
	requestDef := GenReqDefForShowNodes()
	return &ShowNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPathById 通过id获取路径
//
// 通过id获取路径。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowPathById(request *model.ShowPathByIdRequest) (*model.ShowPathByIdResponse, error) {
	requestDef := GenReqDefForShowPathById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPathByIdResponse), nil
	}
}

// ShowPathByIdInvoker 通过id获取路径
func (c *DataArtsStudioClient) ShowPathByIdInvoker(request *model.ShowPathByIdRequest) *ShowPathByIdInvoker {
	requestDef := GenReqDefForShowPathById()
	return &ShowPathByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowPathObjectById 通过id获取路径对象
//
// 通过目录id获取路径对象。
// 通过目录id获取从根目录至当前目录链路上每一层的路径信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowPathObjectById(request *model.ShowPathObjectByIdRequest) (*model.ShowPathObjectByIdResponse, error) {
	requestDef := GenReqDefForShowPathObjectById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPathObjectByIdResponse), nil
	}
}

// ShowPathObjectByIdInvoker 通过id获取路径对象
func (c *DataArtsStudioClient) ShowPathObjectByIdInvoker(request *model.ShowPathObjectByIdRequest) *ShowPathObjectByIdInvoker {
	requestDef := GenReqDefForShowPathObjectById()
	return &ShowPathObjectByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowQualityTaskDetail 获取质量作业详情
//
// 获取质量作业详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowQualityTaskDetail(request *model.ShowQualityTaskDetailRequest) (*model.ShowQualityTaskDetailResponse, error) {
	requestDef := GenReqDefForShowQualityTaskDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowQualityTaskDetailResponse), nil
	}
}

// ShowQualityTaskDetailInvoker 获取质量作业详情
func (c *DataArtsStudioClient) ShowQualityTaskDetailInvoker(request *model.ShowQualityTaskDetailRequest) *ShowQualityTaskDetailInvoker {
	requestDef := GenReqDefForShowQualityTaskDetail()
	return &ShowQualityTaskDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowRelationById 查看关系详情
//
// 通过ID获取关系详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowRelationById(request *model.ShowRelationByIdRequest) (*model.ShowRelationByIdResponse, error) {
	requestDef := GenReqDefForShowRelationById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRelationByIdResponse), nil
	}
}

// ShowRelationByIdInvoker 查看关系详情
func (c *DataArtsStudioClient) ShowRelationByIdInvoker(request *model.ShowRelationByIdRequest) *ShowRelationByIdInvoker {
	requestDef := GenReqDefForShowRelationById()
	return &ShowRelationByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityDataClassificationRule 查询特定识别规则
//
// 查询特定识别规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecurityDataClassificationRule(request *model.ShowSecurityDataClassificationRuleRequest) (*model.ShowSecurityDataClassificationRuleResponse, error) {
	requestDef := GenReqDefForShowSecurityDataClassificationRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityDataClassificationRuleResponse), nil
	}
}

// ShowSecurityDataClassificationRuleInvoker 查询特定识别规则
func (c *DataArtsStudioClient) ShowSecurityDataClassificationRuleInvoker(request *model.ShowSecurityDataClassificationRuleRequest) *ShowSecurityDataClassificationRuleInvoker {
	requestDef := GenReqDefForShowSecurityDataClassificationRule()
	return &ShowSecurityDataClassificationRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityDataClassificationRuleGroup 查询规则组
//
// 查询规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecurityDataClassificationRuleGroup(request *model.ShowSecurityDataClassificationRuleGroupRequest) (*model.ShowSecurityDataClassificationRuleGroupResponse, error) {
	requestDef := GenReqDefForShowSecurityDataClassificationRuleGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityDataClassificationRuleGroupResponse), nil
	}
}

// ShowSecurityDataClassificationRuleGroupInvoker 查询规则组
func (c *DataArtsStudioClient) ShowSecurityDataClassificationRuleGroupInvoker(request *model.ShowSecurityDataClassificationRuleGroupRequest) *ShowSecurityDataClassificationRuleGroupInvoker {
	requestDef := GenReqDefForShowSecurityDataClassificationRuleGroup()
	return &ShowSecurityDataClassificationRuleGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityDynamicMaskingPolicy 查询某个脱敏策略的详细信息
//
// 查询某个脱敏策略的详细信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecurityDynamicMaskingPolicy(request *model.ShowSecurityDynamicMaskingPolicyRequest) (*model.ShowSecurityDynamicMaskingPolicyResponse, error) {
	requestDef := GenReqDefForShowSecurityDynamicMaskingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityDynamicMaskingPolicyResponse), nil
	}
}

// ShowSecurityDynamicMaskingPolicyInvoker 查询某个脱敏策略的详细信息
func (c *DataArtsStudioClient) ShowSecurityDynamicMaskingPolicyInvoker(request *model.ShowSecurityDynamicMaskingPolicyRequest) *ShowSecurityDynamicMaskingPolicyInvoker {
	requestDef := GenReqDefForShowSecurityDynamicMaskingPolicy()
	return &ShowSecurityDynamicMaskingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityMemberSyncTask 查询单个用户同步任务
//
// 查询单个用户同步任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecurityMemberSyncTask(request *model.ShowSecurityMemberSyncTaskRequest) (*model.ShowSecurityMemberSyncTaskResponse, error) {
	requestDef := GenReqDefForShowSecurityMemberSyncTask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityMemberSyncTaskResponse), nil
	}
}

// ShowSecurityMemberSyncTaskInvoker 查询单个用户同步任务
func (c *DataArtsStudioClient) ShowSecurityMemberSyncTaskInvoker(request *model.ShowSecurityMemberSyncTaskRequest) *ShowSecurityMemberSyncTaskInvoker {
	requestDef := GenReqDefForShowSecurityMemberSyncTask()
	return &ShowSecurityMemberSyncTaskInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecurityPermissionSet 查询权限集
//
// 查询权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecurityPermissionSet(request *model.ShowSecurityPermissionSetRequest) (*model.ShowSecurityPermissionSetResponse, error) {
	requestDef := GenReqDefForShowSecurityPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecurityPermissionSetResponse), nil
	}
}

// ShowSecurityPermissionSetInvoker 查询权限集
func (c *DataArtsStudioClient) ShowSecurityPermissionSetInvoker(request *model.ShowSecurityPermissionSetRequest) *ShowSecurityPermissionSetInvoker {
	requestDef := GenReqDefForShowSecurityPermissionSet()
	return &ShowSecurityPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowSecuritySecrecyLevel 根据指定的id查询密级
//
// 根据指定的id查询密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowSecuritySecrecyLevel(request *model.ShowSecuritySecrecyLevelRequest) (*model.ShowSecuritySecrecyLevelResponse, error) {
	requestDef := GenReqDefForShowSecuritySecrecyLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowSecuritySecrecyLevelResponse), nil
	}
}

// ShowSecuritySecrecyLevelInvoker 根据指定的id查询密级
func (c *DataArtsStudioClient) ShowSecuritySecrecyLevelInvoker(request *model.ShowSecuritySecrecyLevelRequest) *ShowSecuritySecrecyLevelInvoker {
	requestDef := GenReqDefForShowSecuritySecrecyLevel()
	return &ShowSecuritySecrecyLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStandardById 查看数据标准详情
//
// 通过ID获取数据标准详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowStandardById(request *model.ShowStandardByIdRequest) (*model.ShowStandardByIdResponse, error) {
	requestDef := GenReqDefForShowStandardById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStandardByIdResponse), nil
	}
}

// ShowStandardByIdInvoker 查看数据标准详情
func (c *DataArtsStudioClient) ShowStandardByIdInvoker(request *model.ShowStandardByIdRequest) *ShowStandardByIdInvoker {
	requestDef := GenReqDefForShowStandardById()
	return &ShowStandardByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowStandardTemplate 查询数据标准模板
//
// 查询当前工作空间下的数据标准模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowStandardTemplate(request *model.ShowStandardTemplateRequest) (*model.ShowStandardTemplateResponse, error) {
	requestDef := GenReqDefForShowStandardTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowStandardTemplateResponse), nil
	}
}

// ShowStandardTemplateInvoker 查询数据标准模板
func (c *DataArtsStudioClient) ShowStandardTemplateInvoker(request *model.ShowStandardTemplateRequest) *ShowStandardTemplateInvoker {
	requestDef := GenReqDefForShowStandardTemplate()
	return &ShowStandardTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTableModelById 查看表模型详情
//
// 通过ID获取模型表详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTableModelById(request *model.ShowTableModelByIdRequest) (*model.ShowTableModelByIdResponse, error) {
	requestDef := GenReqDefForShowTableModelById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTableModelByIdResponse), nil
	}
}

// ShowTableModelByIdInvoker 查看表模型详情
func (c *DataArtsStudioClient) ShowTableModelByIdInvoker(request *model.ShowTableModelByIdRequest) *ShowTableModelByIdInvoker {
	requestDef := GenReqDefForShowTableModelById()
	return &ShowTableModelByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTags 搜索查询标签分页展示
//
// 搜索查询标签分页展示
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTags(request *model.ShowTagsRequest) (*model.ShowTagsResponse, error) {
	requestDef := GenReqDefForShowTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTagsResponse), nil
	}
}

// ShowTagsInvoker 搜索查询标签分页展示
func (c *DataArtsStudioClient) ShowTagsInvoker(request *model.ShowTagsRequest) *ShowTagsInvoker {
	requestDef := GenReqDefForShowTags()
	return &ShowTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskInfo 查询采集任务详情
//
// 查询采集任务详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTaskInfo(request *model.ShowTaskInfoRequest) (*model.ShowTaskInfoResponse, error) {
	requestDef := GenReqDefForShowTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskInfoResponse), nil
	}
}

// ShowTaskInfoInvoker 查询采集任务详情
func (c *DataArtsStudioClient) ShowTaskInfoInvoker(request *model.ShowTaskInfoRequest) *ShowTaskInfoInvoker {
	requestDef := GenReqDefForShowTaskInfo()
	return &ShowTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTaskList 查询采集任务列表
//
// 查询采集任务列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTaskList(request *model.ShowTaskListRequest) (*model.ShowTaskListResponse, error) {
	requestDef := GenReqDefForShowTaskList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTaskListResponse), nil
	}
}

// ShowTaskListInvoker 查询采集任务列表
func (c *DataArtsStudioClient) ShowTaskListInvoker(request *model.ShowTaskListRequest) *ShowTaskListInvoker {
	requestDef := GenReqDefForShowTaskList()
	return &ShowTaskListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTechnicalAssetsStatistic 获取技术资产统计信息
//
// 获取技术资产统计信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTechnicalAssetsStatistic(request *model.ShowTechnicalAssetsStatisticRequest) (*model.ShowTechnicalAssetsStatisticResponse, error) {
	requestDef := GenReqDefForShowTechnicalAssetsStatistic()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTechnicalAssetsStatisticResponse), nil
	}
}

// ShowTechnicalAssetsStatisticInvoker 获取技术资产统计信息
func (c *DataArtsStudioClient) ShowTechnicalAssetsStatisticInvoker(request *model.ShowTechnicalAssetsStatisticRequest) *ShowTechnicalAssetsStatisticInvoker {
	requestDef := GenReqDefForShowTechnicalAssetsStatistic()
	return &ShowTechnicalAssetsStatisticInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowTemplatesDetail 获取规则模板详情
//
// 获取规则模板详情
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowTemplatesDetail(request *model.ShowTemplatesDetailRequest) (*model.ShowTemplatesDetailResponse, error) {
	requestDef := GenReqDefForShowTemplatesDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTemplatesDetailResponse), nil
	}
}

// ShowTemplatesDetailInvoker 获取规则模板详情
func (c *DataArtsStudioClient) ShowTemplatesDetailInvoker(request *model.ShowTemplatesDetailRequest) *ShowTemplatesDetailInvoker {
	requestDef := GenReqDefForShowTemplatesDetail()
	return &ShowTemplatesDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowUnrelatedTable 无血缘关系表查询
//
// 无血缘关系表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowUnrelatedTable(request *model.ShowUnrelatedTableRequest) (*model.ShowUnrelatedTableResponse, error) {
	requestDef := GenReqDefForShowUnrelatedTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowUnrelatedTableResponse), nil
	}
}

// ShowUnrelatedTableInvoker 无血缘关系表查询
func (c *DataArtsStudioClient) ShowUnrelatedTableInvoker(request *model.ShowUnrelatedTableRequest) *ShowUnrelatedTableInvoker {
	requestDef := GenReqDefForShowUnrelatedTable()
	return &ShowUnrelatedTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkSpace 获取单个工作空间信息
//
// 获取单个工作空间信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowWorkSpace(request *model.ShowWorkSpaceRequest) (*model.ShowWorkSpaceResponse, error) {
	requestDef := GenReqDefForShowWorkSpace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkSpaceResponse), nil
	}
}

// ShowWorkSpaceInvoker 获取单个工作空间信息
func (c *DataArtsStudioClient) ShowWorkSpaceInvoker(request *model.ShowWorkSpaceRequest) *ShowWorkSpaceInvoker {
	requestDef := GenReqDefForShowWorkSpace()
	return &ShowWorkSpaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowWorkspaceDetailById 查询模型详情
//
// 查询物理模型或逻辑模型的工作区空间详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowWorkspaceDetailById(request *model.ShowWorkspaceDetailByIdRequest) (*model.ShowWorkspaceDetailByIdResponse, error) {
	requestDef := GenReqDefForShowWorkspaceDetailById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowWorkspaceDetailByIdResponse), nil
	}
}

// ShowWorkspaceDetailByIdInvoker 查询模型详情
func (c *DataArtsStudioClient) ShowWorkspaceDetailByIdInvoker(request *model.ShowWorkspaceDetailByIdRequest) *ShowWorkspaceDetailByIdInvoker {
	requestDef := GenReqDefForShowWorkspaceDetailById()
	return &ShowWorkspaceDetailByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// StopFactorySupplementDataInstance 停止补数据实例
//
// 停止补数据实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) StopFactorySupplementDataInstance(request *model.StopFactorySupplementDataInstanceRequest) (*model.StopFactorySupplementDataInstanceResponse, error) {
	requestDef := GenReqDefForStopFactorySupplementDataInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.StopFactorySupplementDataInstanceResponse), nil
	}
}

// StopFactorySupplementDataInstanceInvoker 停止补数据实例
func (c *DataArtsStudioClient) StopFactorySupplementDataInstanceInvoker(request *model.StopFactorySupplementDataInstanceRequest) *StopFactorySupplementDataInstanceInvoker {
	requestDef := GenReqDefForStopFactorySupplementDataInstance()
	return &StopFactorySupplementDataInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApp 更新应用
//
// 更新应用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateApp(request *model.UpdateAppRequest) (*model.UpdateAppResponse, error) {
	requestDef := GenReqDefForUpdateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAppResponse), nil
	}
}

// UpdateAppInvoker 更新应用
func (c *DataArtsStudioClient) UpdateAppInvoker(request *model.UpdateAppRequest) *UpdateAppInvoker {
	requestDef := GenReqDefForUpdateApp()
	return &UpdateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBizMetric 更新业务指标
//
// 更新业务指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateBizMetric(request *model.UpdateBizMetricRequest) (*model.UpdateBizMetricResponse, error) {
	requestDef := GenReqDefForUpdateBizMetric()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBizMetricResponse), nil
	}
}

// UpdateBizMetricInvoker 更新业务指标
func (c *DataArtsStudioClient) UpdateBizMetricInvoker(request *model.UpdateBizMetricRequest) *UpdateBizMetricInvoker {
	requestDef := GenReqDefForUpdateBizMetric()
	return &UpdateBizMetricInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCatalog 更新服务目录
//
// 更新服务目录。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateCatalog(request *model.UpdateCatalogRequest) (*model.UpdateCatalogResponse, error) {
	requestDef := GenReqDefForUpdateCatalog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCatalogResponse), nil
	}
}

// UpdateCatalogInvoker 更新服务目录
func (c *DataArtsStudioClient) UpdateCatalogInvoker(request *model.UpdateCatalogRequest) *UpdateCatalogInvoker {
	requestDef := GenReqDefForUpdateCatalog()
	return &UpdateCatalogInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCodeTable 修改码表
//
// 修改码表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateCodeTable(request *model.UpdateCodeTableRequest) (*model.UpdateCodeTableResponse, error) {
	requestDef := GenReqDefForUpdateCodeTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCodeTableResponse), nil
	}
}

// UpdateCodeTableInvoker 修改码表
func (c *DataArtsStudioClient) UpdateCodeTableInvoker(request *model.UpdateCodeTableRequest) *UpdateCodeTableInvoker {
	requestDef := GenReqDefForUpdateCodeTable()
	return &UpdateCodeTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCodeTableValues 编辑码表字段值
//
// 编辑码表字段值。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateCodeTableValues(request *model.UpdateCodeTableValuesRequest) (*model.UpdateCodeTableValuesResponse, error) {
	requestDef := GenReqDefForUpdateCodeTableValues()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCodeTableValuesResponse), nil
	}
}

// UpdateCodeTableValuesInvoker 编辑码表字段值
func (c *DataArtsStudioClient) UpdateCodeTableValuesInvoker(request *model.UpdateCodeTableValuesRequest) *UpdateCodeTableValuesInvoker {
	requestDef := GenReqDefForUpdateCodeTableValues()
	return &UpdateCodeTableValuesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDataconnection 更新数据连接信息
//
// 更新数据连接信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateDataconnection(request *model.UpdateDataconnectionRequest) (*model.UpdateDataconnectionResponse, error) {
	requestDef := GenReqDefForUpdateDataconnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDataconnectionResponse), nil
	}
}

// UpdateDataconnectionInvoker 更新数据连接信息
func (c *DataArtsStudioClient) UpdateDataconnectionInvoker(request *model.UpdateDataconnectionRequest) *UpdateDataconnectionInvoker {
	requestDef := GenReqDefForUpdateDataconnection()
	return &UpdateDataconnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesignAggregationLogicTable 更新汇总表
//
// 更新汇总表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateDesignAggregationLogicTable(request *model.UpdateDesignAggregationLogicTableRequest) (*model.UpdateDesignAggregationLogicTableResponse, error) {
	requestDef := GenReqDefForUpdateDesignAggregationLogicTable()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesignAggregationLogicTableResponse), nil
	}
}

// UpdateDesignAggregationLogicTableInvoker 更新汇总表
func (c *DataArtsStudioClient) UpdateDesignAggregationLogicTableInvoker(request *model.UpdateDesignAggregationLogicTableRequest) *UpdateDesignAggregationLogicTableInvoker {
	requestDef := GenReqDefForUpdateDesignAggregationLogicTable()
	return &UpdateDesignAggregationLogicTableInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesignAtomicIndex 更新原子指标
//
// 更新单个原子指标。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateDesignAtomicIndex(request *model.UpdateDesignAtomicIndexRequest) (*model.UpdateDesignAtomicIndexResponse, error) {
	requestDef := GenReqDefForUpdateDesignAtomicIndex()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesignAtomicIndexResponse), nil
	}
}

// UpdateDesignAtomicIndexInvoker 更新原子指标
func (c *DataArtsStudioClient) UpdateDesignAtomicIndexInvoker(request *model.UpdateDesignAtomicIndexRequest) *UpdateDesignAtomicIndexInvoker {
	requestDef := GenReqDefForUpdateDesignAtomicIndex()
	return &UpdateDesignAtomicIndexInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDesignTableQuality 更新表的异常数据输出配置
//
// 更新表的异常数据输出配置，包括是否生成异常数据、设置异常数据数据库或Schema、设置异常表表前缀/表后缀。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateDesignTableQuality(request *model.UpdateDesignTableQualityRequest) (*model.UpdateDesignTableQualityResponse, error) {
	requestDef := GenReqDefForUpdateDesignTableQuality()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDesignTableQualityResponse), nil
	}
}

// UpdateDesignTableQualityInvoker 更新表的异常数据输出配置
func (c *DataArtsStudioClient) UpdateDesignTableQualityInvoker(request *model.UpdateDesignTableQualityRequest) *UpdateDesignTableQualityInvoker {
	requestDef := GenReqDefForUpdateDesignTableQuality()
	return &UpdateDesignTableQualityInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDirectory 修改目录
//
// 修改目录（数据标准、码表）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateDirectory(request *model.UpdateDirectoryRequest) (*model.UpdateDirectoryResponse, error) {
	requestDef := GenReqDefForUpdateDirectory()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDirectoryResponse), nil
	}
}

// UpdateDirectoryInvoker 修改目录
func (c *DataArtsStudioClient) UpdateDirectoryInvoker(request *model.UpdateDirectoryRequest) *UpdateDirectoryInvoker {
	requestDef := GenReqDefForUpdateDirectory()
	return &UpdateDirectoryInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateFactoryJobName 修改作业名称
//
// 修改作业名称
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateFactoryJobName(request *model.UpdateFactoryJobNameRequest) (*model.UpdateFactoryJobNameResponse, error) {
	requestDef := GenReqDefForUpdateFactoryJobName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateFactoryJobNameResponse), nil
	}
}

// UpdateFactoryJobNameInvoker 修改作业名称
func (c *DataArtsStudioClient) UpdateFactoryJobNameInvoker(request *model.UpdateFactoryJobNameRequest) *UpdateFactoryJobNameInvoker {
	requestDef := GenReqDefForUpdateFactoryJobName()
	return &UpdateFactoryJobNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityAssignedQueue 修改当前空间下分配的队列资源
//
// 修改当前空间下分配的队列资源。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityAssignedQueue(request *model.UpdateSecurityAssignedQueueRequest) (*model.UpdateSecurityAssignedQueueResponse, error) {
	requestDef := GenReqDefForUpdateSecurityAssignedQueue()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityAssignedQueueResponse), nil
	}
}

// UpdateSecurityAssignedQueueInvoker 修改当前空间下分配的队列资源
func (c *DataArtsStudioClient) UpdateSecurityAssignedQueueInvoker(request *model.UpdateSecurityAssignedQueueRequest) *UpdateSecurityAssignedQueueInvoker {
	requestDef := GenReqDefForUpdateSecurityAssignedQueue()
	return &UpdateSecurityAssignedQueueInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityDataClassificationRule 修改识别规则接口
//
// 修改识别规则接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityDataClassificationRule(request *model.UpdateSecurityDataClassificationRuleRequest) (*model.UpdateSecurityDataClassificationRuleResponse, error) {
	requestDef := GenReqDefForUpdateSecurityDataClassificationRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityDataClassificationRuleResponse), nil
	}
}

// UpdateSecurityDataClassificationRuleInvoker 修改识别规则接口
func (c *DataArtsStudioClient) UpdateSecurityDataClassificationRuleInvoker(request *model.UpdateSecurityDataClassificationRuleRequest) *UpdateSecurityDataClassificationRuleInvoker {
	requestDef := GenReqDefForUpdateSecurityDataClassificationRule()
	return &UpdateSecurityDataClassificationRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityDynamicMaskingPolicy 更新动态数据脱敏策略
//
// 更新动态数据脱敏策略
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityDynamicMaskingPolicy(request *model.UpdateSecurityDynamicMaskingPolicyRequest) (*model.UpdateSecurityDynamicMaskingPolicyResponse, error) {
	requestDef := GenReqDefForUpdateSecurityDynamicMaskingPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityDynamicMaskingPolicyResponse), nil
	}
}

// UpdateSecurityDynamicMaskingPolicyInvoker 更新动态数据脱敏策略
func (c *DataArtsStudioClient) UpdateSecurityDynamicMaskingPolicyInvoker(request *model.UpdateSecurityDynamicMaskingPolicyRequest) *UpdateSecurityDynamicMaskingPolicyInvoker {
	requestDef := GenReqDefForUpdateSecurityDynamicMaskingPolicy()
	return &UpdateSecurityDynamicMaskingPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityPermissionSet 更新权限集
//
// 更新权限集
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityPermissionSet(request *model.UpdateSecurityPermissionSetRequest) (*model.UpdateSecurityPermissionSetResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPermissionSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPermissionSetResponse), nil
	}
}

// UpdateSecurityPermissionSetInvoker 更新权限集
func (c *DataArtsStudioClient) UpdateSecurityPermissionSetInvoker(request *model.UpdateSecurityPermissionSetRequest) *UpdateSecurityPermissionSetInvoker {
	requestDef := GenReqDefForUpdateSecurityPermissionSet()
	return &UpdateSecurityPermissionSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityPermissionSetPermission 更新权限集的权限
//
// 更新权限集的权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityPermissionSetPermission(request *model.UpdateSecurityPermissionSetPermissionRequest) (*model.UpdateSecurityPermissionSetPermissionResponse, error) {
	requestDef := GenReqDefForUpdateSecurityPermissionSetPermission()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityPermissionSetPermissionResponse), nil
	}
}

// UpdateSecurityPermissionSetPermissionInvoker 更新权限集的权限
func (c *DataArtsStudioClient) UpdateSecurityPermissionSetPermissionInvoker(request *model.UpdateSecurityPermissionSetPermissionRequest) *UpdateSecurityPermissionSetPermissionInvoker {
	requestDef := GenReqDefForUpdateSecurityPermissionSetPermission()
	return &UpdateSecurityPermissionSetPermissionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecurityRuleEnableStatus 修改识别规则状态接口
//
// 修改识别规则状态接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecurityRuleEnableStatus(request *model.UpdateSecurityRuleEnableStatusRequest) (*model.UpdateSecurityRuleEnableStatusResponse, error) {
	requestDef := GenReqDefForUpdateSecurityRuleEnableStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecurityRuleEnableStatusResponse), nil
	}
}

// UpdateSecurityRuleEnableStatusInvoker 修改识别规则状态接口
func (c *DataArtsStudioClient) UpdateSecurityRuleEnableStatusInvoker(request *model.UpdateSecurityRuleEnableStatusRequest) *UpdateSecurityRuleEnableStatusInvoker {
	requestDef := GenReqDefForUpdateSecurityRuleEnableStatus()
	return &UpdateSecurityRuleEnableStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSecuritySecrecyLevel 根据指定的id修改密级
//
// 根据指定的id修改密级
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSecuritySecrecyLevel(request *model.UpdateSecuritySecrecyLevelRequest) (*model.UpdateSecuritySecrecyLevelResponse, error) {
	requestDef := GenReqDefForUpdateSecuritySecrecyLevel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSecuritySecrecyLevelResponse), nil
	}
}

// UpdateSecuritySecrecyLevelInvoker 根据指定的id修改密级
func (c *DataArtsStudioClient) UpdateSecuritySecrecyLevelInvoker(request *model.UpdateSecuritySecrecyLevelRequest) *UpdateSecuritySecrecyLevelInvoker {
	requestDef := GenReqDefForUpdateSecuritySecrecyLevel()
	return &UpdateSecuritySecrecyLevelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStandard 修改数据标准
//
// 修改数据标准。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateStandard(request *model.UpdateStandardRequest) (*model.UpdateStandardResponse, error) {
	requestDef := GenReqDefForUpdateStandard()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStandardResponse), nil
	}
}

// UpdateStandardInvoker 修改数据标准
func (c *DataArtsStudioClient) UpdateStandardInvoker(request *model.UpdateStandardRequest) *UpdateStandardInvoker {
	requestDef := GenReqDefForUpdateStandard()
	return &UpdateStandardInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateStandardTemplate 修改数据标准模板
//
// 修改数据标准模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateStandardTemplate(request *model.UpdateStandardTemplateRequest) (*model.UpdateStandardTemplateResponse, error) {
	requestDef := GenReqDefForUpdateStandardTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateStandardTemplateResponse), nil
	}
}

// UpdateStandardTemplateInvoker 修改数据标准模板
func (c *DataArtsStudioClient) UpdateStandardTemplateInvoker(request *model.UpdateStandardTemplateRequest) *UpdateStandardTemplateInvoker {
	requestDef := GenReqDefForUpdateStandardTemplate()
	return &UpdateStandardTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubject 修改主题
//
// 修改主题。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSubject(request *model.UpdateSubjectRequest) (*model.UpdateSubjectResponse, error) {
	requestDef := GenReqDefForUpdateSubject()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubjectResponse), nil
	}
}

// UpdateSubjectInvoker 修改主题
func (c *DataArtsStudioClient) UpdateSubjectInvoker(request *model.UpdateSubjectRequest) *UpdateSubjectInvoker {
	requestDef := GenReqDefForUpdateSubject()
	return &UpdateSubjectInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateSubjectNew 修改主题(新)
//
// 修改主题(新)。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateSubjectNew(request *model.UpdateSubjectNewRequest) (*model.UpdateSubjectNewResponse, error) {
	requestDef := GenReqDefForUpdateSubjectNew()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateSubjectNewResponse), nil
	}
}

// UpdateSubjectNewInvoker 修改主题(新)
func (c *DataArtsStudioClient) UpdateSubjectNewInvoker(request *model.UpdateSubjectNewRequest) *UpdateSubjectNewInvoker {
	requestDef := GenReqDefForUpdateSubjectNew()
	return &UpdateSubjectNewInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTableModel 更新表模型
//
// 在关系建模中更新一个表模型及其属性，包括逻辑实体和物理表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateTableModel(request *model.UpdateTableModelRequest) (*model.UpdateTableModelResponse, error) {
	requestDef := GenReqDefForUpdateTableModel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTableModelResponse), nil
	}
}

// UpdateTableModelInvoker 更新表模型
func (c *DataArtsStudioClient) UpdateTableModelInvoker(request *model.UpdateTableModelRequest) *UpdateTableModelInvoker {
	requestDef := GenReqDefForUpdateTableModel()
	return &UpdateTableModelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTaskInfo 编辑采集任务
//
// 编辑采集任务
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateTaskInfo(request *model.UpdateTaskInfoRequest) (*model.UpdateTaskInfoResponse, error) {
	requestDef := GenReqDefForUpdateTaskInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTaskInfoResponse), nil
	}
}

// UpdateTaskInfoInvoker 编辑采集任务
func (c *DataArtsStudioClient) UpdateTaskInfoInvoker(request *model.UpdateTaskInfoRequest) *UpdateTaskInfoInvoker {
	requestDef := GenReqDefForUpdateTaskInfo()
	return &UpdateTaskInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateTemplate 更新规则模板
//
// 更新规则模板
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateTemplate(request *model.UpdateTemplateRequest) (*model.UpdateTemplateResponse, error) {
	requestDef := GenReqDefForUpdateTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateTemplateResponse), nil
	}
}

// UpdateTemplateInvoker 更新规则模板
func (c *DataArtsStudioClient) UpdateTemplateInvoker(request *model.UpdateTemplateRequest) *UpdateTemplateInvoker {
	requestDef := GenReqDefForUpdateTemplate()
	return &UpdateTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkSpaceUserOrGroup 编辑工作空间用户或用户组
//
// 编辑工作空间用户或用户组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateWorkSpaceUserOrGroup(request *model.UpdateWorkSpaceUserOrGroupRequest) (*model.UpdateWorkSpaceUserOrGroupResponse, error) {
	requestDef := GenReqDefForUpdateWorkSpaceUserOrGroup()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkSpaceUserOrGroupResponse), nil
	}
}

// UpdateWorkSpaceUserOrGroupInvoker 编辑工作空间用户或用户组
func (c *DataArtsStudioClient) UpdateWorkSpaceUserOrGroupInvoker(request *model.UpdateWorkSpaceUserOrGroupRequest) *UpdateWorkSpaceUserOrGroupInvoker {
	requestDef := GenReqDefForUpdateWorkSpaceUserOrGroup()
	return &UpdateWorkSpaceUserOrGroupInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateWorkspace 更新模型工作区
//
// 更新模型工作区。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateWorkspace(request *model.UpdateWorkspaceRequest) (*model.UpdateWorkspaceResponse, error) {
	requestDef := GenReqDefForUpdateWorkspace()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateWorkspaceResponse), nil
	}
}

// UpdateWorkspaceInvoker 更新模型工作区
func (c *DataArtsStudioClient) UpdateWorkspaceInvoker(request *model.UpdateWorkspaceRequest) *UpdateWorkspaceInvoker {
	requestDef := GenReqDefForUpdateWorkspace()
	return &UpdateWorkspaceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeActionApiToInstance API授权操作(授权/取消授权/申请/续约)
//
// - API主动授权：  API审核人可发起，API主动授权成功后，在有效期内，APP即可访问该API。API授权包含授权和续约两部分功能。
//   - 授权：授权会给予APP在有效期内访问API的权利。
//   - 续约：续约会更新授权有效期，仅支持延长有效期，不能减少。
//
// - API解除授权：  API审核人可发起，解除API对APP的授权关系。解除授权后，APP将不再能够调用该API。API解除已授权的APP关系，需要为APP预留至少2天的准备时间。
// - APP解除授权：  APP所有者可发起，解除API对APP的授权关系。解除授权后，APP将不再能够调用该API。APP解除自己的授权关系，无需预留准备时间。
// - APP申请授权：  APP所有者可发起，APP申请API后，待API的审核人完成审核，APP即可访问该API。授权会给予APP在有效期内访问API的权利，需要API审核。
// - APP申请续约：  APP所有者可发起，续约会更新授权有效期，仅支持延长有效期，不能减少，需要API审核。
// &gt; * 申请自己的API推荐采用API主动授权/续约，无需审核。
// &gt; * 自己的应用推荐采用APP解除授权，无需预留准备时间。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AuthorizeActionApiToInstance(request *model.AuthorizeActionApiToInstanceRequest) (*model.AuthorizeActionApiToInstanceResponse, error) {
	requestDef := GenReqDefForAuthorizeActionApiToInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeActionApiToInstanceResponse), nil
	}
}

// AuthorizeActionApiToInstanceInvoker API授权操作(授权/取消授权/申请/续约)
func (c *DataArtsStudioClient) AuthorizeActionApiToInstanceInvoker(request *model.AuthorizeActionApiToInstanceRequest) *AuthorizeActionApiToInstanceInvoker {
	requestDef := GenReqDefForAuthorizeActionApiToInstance()
	return &AuthorizeActionApiToInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeApiToInstance 批量授权API(专享版)
//
// APP创建成功后，还不能访问API，如果想要访问某个API，需要将该API授权给APP。API主动授权成功后，在有效期内，APP即可访问该API。
// API授权包含授权和续约两部分功能。
// - 授权：授权会给予APP在有效期内访问API的权利。
// - 续约：续约会更新授权有效期，仅支持延长有效期，不能减少。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) AuthorizeApiToInstance(request *model.AuthorizeApiToInstanceRequest) (*model.AuthorizeApiToInstanceResponse, error) {
	requestDef := GenReqDefForAuthorizeApiToInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeApiToInstanceResponse), nil
	}
}

// AuthorizeApiToInstanceInvoker 批量授权API(专享版)
func (c *DataArtsStudioClient) AuthorizeApiToInstanceInvoker(request *model.AuthorizeApiToInstanceRequest) *AuthorizeApiToInstanceInvoker {
	requestDef := GenReqDefForAuthorizeApiToInstance()
	return &AuthorizeApiToInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApi 创建API
//
// 创建API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) CreateApi(request *model.CreateApiRequest) (*model.CreateApiResponse, error) {
	requestDef := GenReqDefForCreateApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateApiResponse), nil
	}
}

// CreateApiInvoker 创建API
func (c *DataArtsStudioClient) CreateApiInvoker(request *model.CreateApiRequest) *CreateApiInvoker {
	requestDef := GenReqDefForCreateApi()
	return &CreateApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DebugApi 调试API
//
// 调试API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DebugApi(request *model.DebugApiRequest) (*model.DebugApiResponse, error) {
	requestDef := GenReqDefForDebugApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DebugApiResponse), nil
	}
}

// DebugApiInvoker 调试API
func (c *DataArtsStudioClient) DebugApiInvoker(request *model.DebugApiRequest) *DebugApiInvoker {
	requestDef := GenReqDefForDebugApi()
	return &DebugApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApi 批量删除API
//
// 批量删除API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) DeleteApi(request *model.DeleteApiRequest) (*model.DeleteApiResponse, error) {
	requestDef := GenReqDefForDeleteApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteApiResponse), nil
	}
}

// DeleteApiInvoker 批量删除API
func (c *DataArtsStudioClient) DeleteApiInvoker(request *model.DeleteApiRequest) *DeleteApiInvoker {
	requestDef := GenReqDefForDeleteApi()
	return &DeleteApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteApiToInstance API操作(下线/停用/恢复)
//
// - 下线API。将已发布的API下线。下线后，所有授权关系都会被解除，API将无法再被调用。
// - 停用API。将已发布的API临时下线。下线后，授权关系会保留，停用期间API将无法再被调用。
// - 恢复API。将已停用的API恢复使用。恢复后， API重新提供调用。
// &gt; * 恢复请求的发起者若非审核人，需要API的审核人完成申请的审核。
// &gt; * 下线/停用请求的发起者，必须为API的审核人。
// &gt; * 下线/停用功能需要为已授权的应用预留充分的准备时间，需至少提前2天发起请求。若需要立即执行下线/停用，需要发起请求后，无有效的授权应用或是有效的授权应用均处理完消息（立即执行，或定期后完成执行）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExecuteApiToInstance(request *model.ExecuteApiToInstanceRequest) (*model.ExecuteApiToInstanceResponse, error) {
	requestDef := GenReqDefForExecuteApiToInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteApiToInstanceResponse), nil
	}
}

// ExecuteApiToInstanceInvoker API操作(下线/停用/恢复)
func (c *DataArtsStudioClient) ExecuteApiToInstanceInvoker(request *model.ExecuteApiToInstanceRequest) *ExecuteApiToInstanceInvoker {
	requestDef := GenReqDefForExecuteApiToInstance()
	return &ExecuteApiToInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDataServiceExcel 导出包含API信息的excel文件
//
// 导出包含API信息的excel文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDataServiceExcel(request *model.ExportDataServiceExcelRequest) (*model.ExportDataServiceExcelResponse, error) {
	requestDef := GenReqDefForExportDataServiceExcel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDataServiceExcelResponse), nil
	}
}

// ExportDataServiceExcelInvoker 导出包含API信息的excel文件
func (c *DataArtsStudioClient) ExportDataServiceExcelInvoker(request *model.ExportDataServiceExcelRequest) *ExportDataServiceExcelInvoker {
	requestDef := GenReqDefForExportDataServiceExcel()
	return &ExportDataServiceExcelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDataServiceExcelTemplate 下载excel模板
//
// 下载excel模板。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDataServiceExcelTemplate(request *model.ExportDataServiceExcelTemplateRequest) (*model.ExportDataServiceExcelTemplateResponse, error) {
	requestDef := GenReqDefForExportDataServiceExcelTemplate()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDataServiceExcelTemplateResponse), nil
	}
}

// ExportDataServiceExcelTemplateInvoker 下载excel模板
func (c *DataArtsStudioClient) ExportDataServiceExcelTemplateInvoker(request *model.ExportDataServiceExcelTemplateRequest) *ExportDataServiceExcelTemplateInvoker {
	requestDef := GenReqDefForExportDataServiceExcelTemplate()
	return &ExportDataServiceExcelTemplateInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExportDataServiceZip 全量导出包含API的excel压缩文件
//
// 全量导出包含API的excel压缩文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ExportDataServiceZip(request *model.ExportDataServiceZipRequest) (*model.ExportDataServiceZipResponse, error) {
	requestDef := GenReqDefForExportDataServiceZip()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExportDataServiceZipResponse), nil
	}
}

// ExportDataServiceZipInvoker 全量导出包含API的excel压缩文件
func (c *DataArtsStudioClient) ExportDataServiceZipInvoker(request *model.ExportDataServiceZipRequest) *ExportDataServiceZipInvoker {
	requestDef := GenReqDefForExportDataServiceZip()
	return &ExportDataServiceZipInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ImportDataServiceExcel 导入包含API信息的excel文件
//
// 导入包含API信息的excel文件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ImportDataServiceExcel(request *model.ImportDataServiceExcelRequest) (*model.ImportDataServiceExcelResponse, error) {
	requestDef := GenReqDefForImportDataServiceExcel()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ImportDataServiceExcelResponse), nil
	}
}

// ImportDataServiceExcelInvoker 导入包含API信息的excel文件
func (c *DataArtsStudioClient) ImportDataServiceExcelInvoker(request *model.ImportDataServiceExcelRequest) *ImportDataServiceExcelInvoker {
	requestDef := GenReqDefForImportDataServiceExcel()
	return &ImportDataServiceExcelInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListApis 查询API列表
//
// 查询API列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListApis(request *model.ListApisRequest) (*model.ListApisResponse, error) {
	requestDef := GenReqDefForListApis()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApisResponse), nil
	}
}

// ListApisInvoker 查询API列表
func (c *DataArtsStudioClient) ListApisInvoker(request *model.ListApisRequest) *ListApisInvoker {
	requestDef := GenReqDefForListApis()
	return &ListApisInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInstanceList 查看API不同操作对应的实例信息(专享版)
//
// 查看API不同操作对应的实例信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ListInstanceList(request *model.ListInstanceListRequest) (*model.ListInstanceListResponse, error) {
	requestDef := GenReqDefForListInstanceList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstanceListResponse), nil
	}
}

// ListInstanceListInvoker 查看API不同操作对应的实例信息(专享版)
func (c *DataArtsStudioClient) ListInstanceListInvoker(request *model.ListInstanceListRequest) *ListInstanceListInvoker {
	requestDef := GenReqDefForListInstanceList()
	return &ListInstanceListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PublishApiToInstance 发布API
//
// 发布API。API只有发布后，才能够被调用。API发布时，可以将API发送至指定网关。
// - 共享版，必须发送至API网关共享版。
// - 专享版，可以依据自身需要，选择将API发送至API网关专享版、ROMA-APIC、或不发布网关。
// &gt; 发布请求的发起者若非审核人，需要API的审核人完成申请的审核。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) PublishApiToInstance(request *model.PublishApiToInstanceRequest) (*model.PublishApiToInstanceResponse, error) {
	requestDef := GenReqDefForPublishApiToInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PublishApiToInstanceResponse), nil
	}
}

// PublishApiToInstanceInvoker 发布API
func (c *DataArtsStudioClient) PublishApiToInstanceInvoker(request *model.PublishApiToInstanceRequest) *PublishApiToInstanceInvoker {
	requestDef := GenReqDefForPublishApiToInstance()
	return &PublishApiToInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchDebugInfo 查看API调试信息(专享版)
//
// 查看API在不同集群上的调试信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchDebugInfo(request *model.SearchDebugInfoRequest) (*model.SearchDebugInfoResponse, error) {
	requestDef := GenReqDefForSearchDebugInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchDebugInfoResponse), nil
	}
}

// SearchDebugInfoInvoker 查看API调试信息(专享版)
func (c *DataArtsStudioClient) SearchDebugInfoInvoker(request *model.SearchDebugInfoRequest) *SearchDebugInfoInvoker {
	requestDef := GenReqDefForSearchDebugInfo()
	return &SearchDebugInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// SearchPublishInfo 查看API发布信息(专享版)
//
// 查看API在不同集群上的发布信息。
// API在集群上进行过操作后会存在发布信息，例如调试、注册类发布等。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) SearchPublishInfo(request *model.SearchPublishInfoRequest) (*model.SearchPublishInfoResponse, error) {
	requestDef := GenReqDefForSearchPublishInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchPublishInfoResponse), nil
	}
}

// SearchPublishInfoInvoker 查看API发布信息(专享版)
func (c *DataArtsStudioClient) SearchPublishInfoInvoker(request *model.SearchPublishInfoRequest) *SearchPublishInfoInvoker {
	requestDef := GenReqDefForSearchPublishInfo()
	return &SearchPublishInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowApi 查询API信息
//
// 查询API信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) ShowApi(request *model.ShowApiRequest) (*model.ShowApiResponse, error) {
	requestDef := GenReqDefForShowApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiResponse), nil
	}
}

// ShowApiInvoker 查询API信息
func (c *DataArtsStudioClient) ShowApiInvoker(request *model.ShowApiRequest) *ShowApiInvoker {
	requestDef := GenReqDefForShowApi()
	return &ShowApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateApi 更新API
//
// 更新API。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DataArtsStudioClient) UpdateApi(request *model.UpdateApiRequest) (*model.UpdateApiResponse, error) {
	requestDef := GenReqDefForUpdateApi()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateApiResponse), nil
	}
}

// UpdateApiInvoker 更新API
func (c *DataArtsStudioClient) UpdateApiInvoker(request *model.UpdateApiRequest) *UpdateApiInvoker {
	requestDef := GenReqDefForUpdateApi()
	return &UpdateApiInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
