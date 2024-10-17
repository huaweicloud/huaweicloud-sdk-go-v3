package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dataartsstudio/v1/model"
)

type AddDesignEntityTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDesignEntityTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddDesignEntityTagsInvoker) Invoke() (*model.AddDesignEntityTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDesignEntityTagsResponse), nil
	}
}

type AddTagToAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTagToAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddTagToAssetInvoker) Invoke() (*model.AddTagToAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTagToAssetResponse), nil
	}
}

type AddWorkSpaceUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddWorkSpaceUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddWorkSpaceUsersInvoker) Invoke() (*model.AddWorkSpaceUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddWorkSpaceUsersResponse), nil
	}
}

type AssociateClassificationToEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateClassificationToEntityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateClassificationToEntityInvoker) Invoke() (*model.AssociateClassificationToEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateClassificationToEntityResponse), nil
	}
}

type AssociateSecurityLevelToEntitieInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateSecurityLevelToEntitieInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateSecurityLevelToEntitieInvoker) Invoke() (*model.AssociateSecurityLevelToEntitieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateSecurityLevelToEntitieResponse), nil
	}
}

type BatchApproveApplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchApproveApplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchApproveApplyInvoker) Invoke() (*model.BatchApproveApplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchApproveApplyResponse), nil
	}
}

type BatchAssociateClassificationToEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateClassificationToEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAssociateClassificationToEntitiesInvoker) Invoke() (*model.BatchAssociateClassificationToEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateClassificationToEntitiesResponse), nil
	}
}

type BatchAssociateSecurityLevelToEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateSecurityLevelToEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAssociateSecurityLevelToEntitiesInvoker) Invoke() (*model.BatchAssociateSecurityLevelToEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateSecurityLevelToEntitiesResponse), nil
	}
}

type BatchDeleteSecurityDataClassificationRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecurityDataClassificationRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSecurityDataClassificationRuleInvoker) Invoke() (*model.BatchDeleteSecurityDataClassificationRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecurityDataClassificationRuleResponse), nil
	}
}

type BatchDeleteSecurityDynamicMaskingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecurityDynamicMaskingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSecurityDynamicMaskingPoliciesInvoker) Invoke() (*model.BatchDeleteSecurityDynamicMaskingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecurityDynamicMaskingPoliciesResponse), nil
	}
}

type BatchDeleteSecurityPermissionSetMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecurityPermissionSetMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSecurityPermissionSetMembersInvoker) Invoke() (*model.BatchDeleteSecurityPermissionSetMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecurityPermissionSetMembersResponse), nil
	}
}

type BatchDeleteSecurityPermissionSetPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecurityPermissionSetPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSecurityPermissionSetPermissionsInvoker) Invoke() (*model.BatchDeleteSecurityPermissionSetPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecurityPermissionSetPermissionsResponse), nil
	}
}

type BatchDeleteSecuritySecrecyLevelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecuritySecrecyLevelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteSecuritySecrecyLevelsInvoker) Invoke() (*model.BatchDeleteSecuritySecrecyLevelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecuritySecrecyLevelsResponse), nil
	}
}

type BatchDeleteTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTemplatesInvoker) Invoke() (*model.BatchDeleteTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTemplatesResponse), nil
	}
}

type BatchOfflineInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchOfflineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchOfflineInvoker) Invoke() (*model.BatchOfflineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchOfflineResponse), nil
	}
}

type BatchPublishInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchPublishInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchPublishInvoker) Invoke() (*model.BatchPublishResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchPublishResponse), nil
	}
}

type BatchSyncMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSyncMetadataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSyncMetadataInvoker) Invoke() (*model.BatchSyncMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSyncMetadataResponse), nil
	}
}

type BatchUpdateSecurityDlfDataWareHousesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateSecurityDlfDataWareHousesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateSecurityDlfDataWareHousesInvoker) Invoke() (*model.BatchUpdateSecurityDlfDataWareHousesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateSecurityDlfDataWareHousesResponse), nil
	}
}

type CancelFactoryPackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelFactoryPackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelFactoryPackagesInvoker) Invoke() (*model.CancelFactoryPackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelFactoryPackagesResponse), nil
	}
}

type ChangeCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeCatalogInvoker) Invoke() (*model.ChangeCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeCatalogResponse), nil
	}
}

type ChangeResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeResourceInvoker) Invoke() (*model.ChangeResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeResourceResponse), nil
	}
}

type ChangeSubjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSubjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSubjectsInvoker) Invoke() (*model.ChangeSubjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSubjectsResponse), nil
	}
}

type CheckDimensionStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckDimensionStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckDimensionStatusInvoker) Invoke() (*model.CheckDimensionStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckDimensionStatusResponse), nil
	}
}

type CheckFactLogicTableStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckFactLogicTableStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckFactLogicTableStatusInvoker) Invoke() (*model.CheckFactLogicTableStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckFactLogicTableStatusResponse), nil
	}
}

type ConfirmApprovalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmApprovalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmApprovalsInvoker) Invoke() (*model.ConfirmApprovalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmApprovalsResponse), nil
	}
}

type ConfirmMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmMessageInvoker) Invoke() (*model.ConfirmMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmMessageResponse), nil
	}
}

type CountAllModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountAllModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountAllModelsInvoker) Invoke() (*model.CountAllModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountAllModelsResponse), nil
	}
}

type CountOverviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountOverviewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountOverviewsInvoker) Invoke() (*model.CountOverviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountOverviewsResponse), nil
	}
}

type CountStandardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountStandardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountStandardsInvoker) Invoke() (*model.CountStandardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountStandardsResponse), nil
	}
}

type CountTableModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountTableModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountTableModelsInvoker) Invoke() (*model.CountTableModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountTableModelsResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateApproverInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApproverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApproverInvoker) Invoke() (*model.CreateApproverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApproverResponse), nil
	}
}

type CreateBizMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBizMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBizMetricInvoker) Invoke() (*model.CreateBizMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBizMetricResponse), nil
	}
}

type CreateCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCatalogInvoker) Invoke() (*model.CreateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCatalogResponse), nil
	}
}

type CreateCodeTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCodeTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCodeTableInvoker) Invoke() (*model.CreateCodeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCodeTableResponse), nil
	}
}

type CreateConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConnectionsInvoker) Invoke() (*model.CreateConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionsResponse), nil
	}
}

type CreateDesignAggregationLogicTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesignAggregationLogicTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesignAggregationLogicTableInvoker) Invoke() (*model.CreateDesignAggregationLogicTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesignAggregationLogicTableResponse), nil
	}
}

type CreateDesignAtomicIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDesignAtomicIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDesignAtomicIndexInvoker) Invoke() (*model.CreateDesignAtomicIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDesignAtomicIndexResponse), nil
	}
}

type CreateDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDirectoryInvoker) Invoke() (*model.CreateDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDirectoryResponse), nil
	}
}

type CreateFactoryEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFactoryEnvInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFactoryEnvInvoker) Invoke() (*model.CreateFactoryEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFactoryEnvResponse), nil
	}
}

type CreateFactoryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFactoryJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFactoryJobInvoker) Invoke() (*model.CreateFactoryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFactoryJobResponse), nil
	}
}

type CreateFactorySupplementDataInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFactorySupplementDataInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFactorySupplementDataInstanceInvoker) Invoke() (*model.CreateFactorySupplementDataInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFactorySupplementDataInstanceResponse), nil
	}
}

type CreateLineageInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLineageInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLineageInfoInvoker) Invoke() (*model.CreateLineageInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLineageInfoResponse), nil
	}
}

type CreateManagerWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManagerWorkSpaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateManagerWorkSpaceInvoker) Invoke() (*model.CreateManagerWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManagerWorkSpaceResponse), nil
	}
}

type CreateOrUpdateAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateAssetInvoker) Invoke() (*model.CreateOrUpdateAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateAssetResponse), nil
	}
}

type CreateOrUpdateEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateEntitiesInvoker) Invoke() (*model.CreateOrUpdateEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateEntitiesResponse), nil
	}
}

type CreateSecurityAssignedQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityAssignedQueueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityAssignedQueueInvoker) Invoke() (*model.CreateSecurityAssignedQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityAssignedQueueResponse), nil
	}
}

type CreateSecurityDataClassificationRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityDataClassificationRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityDataClassificationRuleInvoker) Invoke() (*model.CreateSecurityDataClassificationRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityDataClassificationRuleResponse), nil
	}
}

type CreateSecurityDynamicMaskingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityDynamicMaskingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityDynamicMaskingPolicyInvoker) Invoke() (*model.CreateSecurityDynamicMaskingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityDynamicMaskingPolicyResponse), nil
	}
}

type CreateSecurityPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityPermissionSetInvoker) Invoke() (*model.CreateSecurityPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityPermissionSetResponse), nil
	}
}

type CreateSecurityPermissionSetMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityPermissionSetMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityPermissionSetMemberInvoker) Invoke() (*model.CreateSecurityPermissionSetMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityPermissionSetMemberResponse), nil
	}
}

type CreateSecurityPermissionSetPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityPermissionSetPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecurityPermissionSetPermissionInvoker) Invoke() (*model.CreateSecurityPermissionSetPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityPermissionSetPermissionResponse), nil
	}
}

type CreateSecuritySecrecyLevelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecuritySecrecyLevelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSecuritySecrecyLevelInvoker) Invoke() (*model.CreateSecuritySecrecyLevelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecuritySecrecyLevelResponse), nil
	}
}

type CreateServiceCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServiceCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServiceCatalogInvoker) Invoke() (*model.CreateServiceCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServiceCatalogResponse), nil
	}
}

type CreateStandardInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStandardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStandardInvoker) Invoke() (*model.CreateStandardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStandardResponse), nil
	}
}

type CreateStandardTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStandardTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStandardTemplateInvoker) Invoke() (*model.CreateStandardTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStandardTemplateResponse), nil
	}
}

type CreateSubjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubjectInvoker) Invoke() (*model.CreateSubjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubjectResponse), nil
	}
}

type CreateSubjectNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubjectNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubjectNewInvoker) Invoke() (*model.CreateSubjectNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubjectNewResponse), nil
	}
}

type CreateTableModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTableModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTableModelInvoker) Invoke() (*model.CreateTableModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTableModelResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type CreateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkspaceInvoker) Invoke() (*model.CreateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkspaceResponse), nil
	}
}

type DebugDataconnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DebugDataconnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugDataconnectionInvoker) Invoke() (*model.DebugDataconnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugDataconnectionResponse), nil
	}
}

type DebugSecurityDlfDataWareHousesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DebugSecurityDlfDataWareHousesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugSecurityDlfDataWareHousesInvoker) Invoke() (*model.DebugSecurityDlfDataWareHousesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugSecurityDlfDataWareHousesResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteApproverInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApproverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApproverInvoker) Invoke() (*model.DeleteApproverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApproverResponse), nil
	}
}

type DeleteAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAssetInvoker) Invoke() (*model.DeleteAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetResponse), nil
	}
}

type DeleteBizMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBizMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBizMetricInvoker) Invoke() (*model.DeleteBizMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBizMetricResponse), nil
	}
}

type DeleteCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCatalogInvoker) Invoke() (*model.DeleteCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCatalogResponse), nil
	}
}

type DeleteClassificationFromEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClassificationFromEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClassificationFromEntitiesInvoker) Invoke() (*model.DeleteClassificationFromEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClassificationFromEntitiesResponse), nil
	}
}

type DeleteCodeTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCodeTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCodeTableInvoker) Invoke() (*model.DeleteCodeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCodeTableResponse), nil
	}
}

type DeleteDataconnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataconnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataconnectionInvoker) Invoke() (*model.DeleteDataconnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataconnectionResponse), nil
	}
}

type DeleteDesignAggregationLogicTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesignAggregationLogicTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesignAggregationLogicTableInvoker) Invoke() (*model.DeleteDesignAggregationLogicTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesignAggregationLogicTableResponse), nil
	}
}

type DeleteDesignAtomicIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesignAtomicIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesignAtomicIndexInvoker) Invoke() (*model.DeleteDesignAtomicIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesignAtomicIndexResponse), nil
	}
}

type DeleteDesignLatestApprovalInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDesignLatestApprovalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDesignLatestApprovalInvoker) Invoke() (*model.DeleteDesignLatestApprovalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDesignLatestApprovalResponse), nil
	}
}

type DeleteDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDirectoryInvoker) Invoke() (*model.DeleteDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDirectoryResponse), nil
	}
}

type DeleteSecurityAssignedQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityAssignedQueueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityAssignedQueueInvoker) Invoke() (*model.DeleteSecurityAssignedQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityAssignedQueueResponse), nil
	}
}

type DeleteSecurityDataClassificationRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityDataClassificationRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityDataClassificationRuleInvoker) Invoke() (*model.DeleteSecurityDataClassificationRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityDataClassificationRuleResponse), nil
	}
}

type DeleteSecurityLevelFromEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityLevelFromEntityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityLevelFromEntityInvoker) Invoke() (*model.DeleteSecurityLevelFromEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityLevelFromEntityResponse), nil
	}
}

type DeleteSecurityPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityPermissionSetInvoker) Invoke() (*model.DeleteSecurityPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityPermissionSetResponse), nil
	}
}

type DeleteSecuritySecrecyLevelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecuritySecrecyLevelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecuritySecrecyLevelInvoker) Invoke() (*model.DeleteSecuritySecrecyLevelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecuritySecrecyLevelResponse), nil
	}
}

type DeleteServiceCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServiceCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServiceCatalogInvoker) Invoke() (*model.DeleteServiceCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServiceCatalogResponse), nil
	}
}

type DeleteStandardInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStandardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStandardInvoker) Invoke() (*model.DeleteStandardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStandardResponse), nil
	}
}

type DeleteStandardTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStandardTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStandardTemplateInvoker) Invoke() (*model.DeleteStandardTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStandardTemplateResponse), nil
	}
}

type DeleteSubjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubjectInvoker) Invoke() (*model.DeleteSubjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubjectResponse), nil
	}
}

type DeleteSubjectNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubjectNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSubjectNewInvoker) Invoke() (*model.DeleteSubjectNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubjectNewResponse), nil
	}
}

type DeleteTableModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTableModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTableModelInvoker) Invoke() (*model.DeleteTableModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTableModelResponse), nil
	}
}

type DeleteTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskInfoInvoker) Invoke() (*model.DeleteTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskInfoResponse), nil
	}
}

type DeleteWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkspacesInvoker) Invoke() (*model.DeleteWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkspacesResponse), nil
	}
}

type DeleteWorkspaceusersInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkspaceusersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkspaceusersInvoker) Invoke() (*model.DeleteWorkspaceusersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkspaceusersResponse), nil
	}
}

type DeployFactoryPackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployFactoryPackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeployFactoryPackagesInvoker) Invoke() (*model.DeployFactoryPackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployFactoryPackagesResponse), nil
	}
}

type ExecuteTaskActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTaskActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteTaskActionInvoker) Invoke() (*model.ExecuteTaskActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTaskActionResponse), nil
	}
}

type ExportDesignModelTableDdlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDesignModelTableDdlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDesignModelTableDdlInvoker) Invoke() (*model.ExportDesignModelTableDdlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDesignModelTableDdlResponse), nil
	}
}

type ExportDesignModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDesignModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDesignModelsInvoker) Invoke() (*model.ExportDesignModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDesignModelsResponse), nil
	}
}

type ExportDesignResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDesignResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDesignResultInvoker) Invoke() (*model.ExportDesignResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDesignResultResponse), nil
	}
}

type ImportCatalogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCatalogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportCatalogsInvoker) Invoke() (*model.ImportCatalogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCatalogsResponse), nil
	}
}

type ImportLineageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportLineageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportLineageInvoker) Invoke() (*model.ImportLineageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportLineageResponse), nil
	}
}

type ImportModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportModelsInvoker) Invoke() (*model.ImportModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportModelsResponse), nil
	}
}

type ImportResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportResultInvoker) Invoke() (*model.ImportResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportResultResponse), nil
	}
}

type InitializeStandardTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *InitializeStandardTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InitializeStandardTemplateInvoker) Invoke() (*model.InitializeStandardTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InitializeStandardTemplateResponse), nil
	}
}

type ListAggregationLogicTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAggregationLogicTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAggregationLogicTablesInvoker) Invoke() (*model.ListAggregationLogicTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAggregationLogicTablesResponse), nil
	}
}

type ListAllCatalogListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllCatalogListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllCatalogListInvoker) Invoke() (*model.ListAllCatalogListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllCatalogListResponse), nil
	}
}

type ListAllStandardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllStandardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllStandardsInvoker) Invoke() (*model.ListAllStandardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllStandardsResponse), nil
	}
}

type ListAllTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllTablesInvoker) Invoke() (*model.ListAllTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTablesResponse), nil
	}
}

type ListApiCatalogListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiCatalogListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiCatalogListInvoker) Invoke() (*model.ListApiCatalogListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiCatalogListResponse), nil
	}
}

type ListApiTopNInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiTopNInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiTopNInvoker) Invoke() (*model.ListApiTopNResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiTopNResponse), nil
	}
}

type ListApicGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApicGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApicGroupsInvoker) Invoke() (*model.ListApicGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApicGroupsResponse), nil
	}
}

type ListApicInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApicInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApicInstancesInvoker) Invoke() (*model.ListApicInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApicInstancesResponse), nil
	}
}

type ListApisTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisTopInvoker) Invoke() (*model.ListApisTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisTopResponse), nil
	}
}

type ListApplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplyInvoker) Invoke() (*model.ListApplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplyResponse), nil
	}
}

type ListApproversInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApproversInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApproversInvoker) Invoke() (*model.ListApproversResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApproversResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListAppsTopInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsTopInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsTopInvoker) Invoke() (*model.ListAppsTopResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsTopResponse), nil
	}
}

type ListBizMetricDimensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBizMetricDimensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBizMetricDimensionsInvoker) Invoke() (*model.ListBizMetricDimensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBizMetricDimensionsResponse), nil
	}
}

type ListBizMetricOwnersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBizMetricOwnersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBizMetricOwnersInvoker) Invoke() (*model.ListBizMetricOwnersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBizMetricOwnersResponse), nil
	}
}

type ListBizMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBizMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBizMetricsInvoker) Invoke() (*model.ListBizMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBizMetricsResponse), nil
	}
}

type ListBusinessInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBusinessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBusinessInvoker) Invoke() (*model.ListBusinessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBusinessResponse), nil
	}
}

type ListCatalogListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCatalogListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCatalogListInvoker) Invoke() (*model.ListCatalogListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCatalogListResponse), nil
	}
}

type ListCatalogTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCatalogTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCatalogTreeInvoker) Invoke() (*model.ListCatalogTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCatalogTreeResponse), nil
	}
}

type ListCategoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCategoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCategoryInvoker) Invoke() (*model.ListCategoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCategoryResponse), nil
	}
}

type ListColumnsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListColumnsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListColumnsInvoker) Invoke() (*model.ListColumnsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListColumnsResponse), nil
	}
}

type ListCompoundMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCompoundMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCompoundMetricsInvoker) Invoke() (*model.ListCompoundMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompoundMetricsResponse), nil
	}
}

type ListConditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConditionInvoker) Invoke() (*model.ListConditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConditionResponse), nil
	}
}

type ListConsistencyTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsistencyTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConsistencyTaskInvoker) Invoke() (*model.ListConsistencyTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsistencyTaskResponse), nil
	}
}

type ListDataArtsStudioInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataArtsStudioInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataArtsStudioInstancesInvoker) Invoke() (*model.ListDataArtsStudioInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataArtsStudioInstancesResponse), nil
	}
}

type ListDataServiceInstancesDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataServiceInstancesDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataServiceInstancesDetailInvoker) Invoke() (*model.ListDataServiceInstancesDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataServiceInstancesDetailResponse), nil
	}
}

type ListDataServiceInstancesOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataServiceInstancesOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataServiceInstancesOverviewInvoker) Invoke() (*model.ListDataServiceInstancesOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataServiceInstancesOverviewResponse), nil
	}
}

type ListDataServiceMarketApisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataServiceMarketApisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataServiceMarketApisInvoker) Invoke() (*model.ListDataServiceMarketApisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataServiceMarketApisResponse), nil
	}
}

type ListDataTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataTablesInvoker) Invoke() (*model.ListDataTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataTablesResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ListDataconnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataconnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDataconnectionsInvoker) Invoke() (*model.ListDataconnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataconnectionsResponse), nil
	}
}

type ListDerivativeIndexesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDerivativeIndexesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDerivativeIndexesInvoker) Invoke() (*model.ListDerivativeIndexesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDerivativeIndexesResponse), nil
	}
}

type ListDimensionGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDimensionGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDimensionGroupsInvoker) Invoke() (*model.ListDimensionGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDimensionGroupsResponse), nil
	}
}

type ListDimensionLogicTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDimensionLogicTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDimensionLogicTablesInvoker) Invoke() (*model.ListDimensionLogicTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDimensionLogicTablesResponse), nil
	}
}

type ListDimensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDimensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDimensionsInvoker) Invoke() (*model.ListDimensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDimensionsResponse), nil
	}
}

type ListDirectoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDirectoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDirectoriesInvoker) Invoke() (*model.ListDirectoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDirectoriesResponse), nil
	}
}

type ListFactLogicTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactLogicTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactLogicTablesInvoker) Invoke() (*model.ListFactLogicTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactLogicTablesResponse), nil
	}
}

type ListFactoryAlarmInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryAlarmInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryAlarmInfoInvoker) Invoke() (*model.ListFactoryAlarmInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryAlarmInfoResponse), nil
	}
}

type ListFactoryJobInstancesByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryJobInstancesByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryJobInstancesByNameInvoker) Invoke() (*model.ListFactoryJobInstancesByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryJobInstancesByNameResponse), nil
	}
}

type ListFactoryJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryJobsInvoker) Invoke() (*model.ListFactoryJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryJobsResponse), nil
	}
}

type ListFactoryReleasePackagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryReleasePackagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryReleasePackagesInvoker) Invoke() (*model.ListFactoryReleasePackagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryReleasePackagesResponse), nil
	}
}

type ListFactoryTaskCompletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryTaskCompletionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryTaskCompletionInvoker) Invoke() (*model.ListFactoryTaskCompletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryTaskCompletionResponse), nil
	}
}

type ListFactoryTaskOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFactoryTaskOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFactoryTaskOverviewInvoker) Invoke() (*model.ListFactoryTaskOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFactoryTaskOverviewResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListManagerWorkSpacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagerWorkSpacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagerWorkSpacesInvoker) Invoke() (*model.ListManagerWorkSpacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagerWorkSpacesResponse), nil
	}
}

type ListMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessageInvoker) Invoke() (*model.ListMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessageResponse), nil
	}
}

type ListMetricRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMetricRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMetricRelationsInvoker) Invoke() (*model.ListMetricRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMetricRelationsResponse), nil
	}
}

type ListQualityTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQualityTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQualityTaskInvoker) Invoke() (*model.ListQualityTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQualityTaskResponse), nil
	}
}

type ListQualityTaskListsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQualityTaskListsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQualityTaskListsInvoker) Invoke() (*model.ListQualityTaskListsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQualityTaskListsResponse), nil
	}
}

type ListQualityTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQualityTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQualityTemplatesInvoker) Invoke() (*model.ListQualityTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQualityTemplatesResponse), nil
	}
}

type ListRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRelationsInvoker) Invoke() (*model.ListRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationsResponse), nil
	}
}

type ListSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSchemasInvoker) Invoke() (*model.ListSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSchemasResponse), nil
	}
}

type ListSecurityAssignedQueuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityAssignedQueuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityAssignedQueuesInvoker) Invoke() (*model.ListSecurityAssignedQueuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityAssignedQueuesResponse), nil
	}
}

type ListSecurityDataClassificationRuleGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDataClassificationRuleGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDataClassificationRuleGroupsInvoker) Invoke() (*model.ListSecurityDataClassificationRuleGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDataClassificationRuleGroupsResponse), nil
	}
}

type ListSecurityDataClassificationRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDataClassificationRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDataClassificationRulesInvoker) Invoke() (*model.ListSecurityDataClassificationRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDataClassificationRulesResponse), nil
	}
}

type ListSecurityDatasourceActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDatasourceActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDatasourceActionsInvoker) Invoke() (*model.ListSecurityDatasourceActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDatasourceActionsResponse), nil
	}
}

type ListSecurityDatasourceConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDatasourceConfigurationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDatasourceConfigurationsInvoker) Invoke() (*model.ListSecurityDatasourceConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDatasourceConfigurationsResponse), nil
	}
}

type ListSecurityDatasourceUrlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDatasourceUrlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDatasourceUrlsInvoker) Invoke() (*model.ListSecurityDatasourceUrlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDatasourceUrlsResponse), nil
	}
}

type ListSecurityDlfDataWareHousesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDlfDataWareHousesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDlfDataWareHousesInvoker) Invoke() (*model.ListSecurityDlfDataWareHousesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDlfDataWareHousesResponse), nil
	}
}

type ListSecurityDynamicMaskingPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityDynamicMaskingPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityDynamicMaskingPoliciesInvoker) Invoke() (*model.ListSecurityDynamicMaskingPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityDynamicMaskingPoliciesResponse), nil
	}
}

type ListSecurityMemberSyncTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityMemberSyncTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityMemberSyncTasksInvoker) Invoke() (*model.ListSecurityMemberSyncTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityMemberSyncTasksResponse), nil
	}
}

type ListSecurityPermissionSetMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityPermissionSetMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityPermissionSetMembersInvoker) Invoke() (*model.ListSecurityPermissionSetMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityPermissionSetMembersResponse), nil
	}
}

type ListSecurityPermissionSetPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityPermissionSetPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityPermissionSetPermissionsInvoker) Invoke() (*model.ListSecurityPermissionSetPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityPermissionSetPermissionsResponse), nil
	}
}

type ListSecurityPermissionSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityPermissionSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityPermissionSetsInvoker) Invoke() (*model.ListSecurityPermissionSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityPermissionSetsResponse), nil
	}
}

type ListSecurityRoleActionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityRoleActionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecurityRoleActionsInvoker) Invoke() (*model.ListSecurityRoleActionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityRoleActionsResponse), nil
	}
}

type ListSecuritySecrecyLevelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecuritySecrecyLevelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecuritySecrecyLevelsInvoker) Invoke() (*model.ListSecuritySecrecyLevelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecuritySecrecyLevelsResponse), nil
	}
}

type ListSecuritySensitiveDataOverviewsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecuritySensitiveDataOverviewsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSecuritySensitiveDataOverviewsInvoker) Invoke() (*model.ListSecuritySensitiveDataOverviewsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecuritySensitiveDataOverviewsResponse), nil
	}
}

type ListSubjectLevelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubjectLevelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubjectLevelsInvoker) Invoke() (*model.ListSubjectLevelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubjectLevelsResponse), nil
	}
}

type ListTableModelRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableModelRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTableModelRelationsInvoker) Invoke() (*model.ListTableModelRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableModelRelationsResponse), nil
	}
}

type ListTableModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTableModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTableModelsInvoker) Invoke() (*model.ListTableModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTableModelsResponse), nil
	}
}

type ListWorkspaceRolesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspaceRolesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspaceRolesInvoker) Invoke() (*model.ListWorkspaceRolesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspaceRolesResponse), nil
	}
}

type ListWorkspacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspacesInvoker) Invoke() (*model.ListWorkspacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspacesResponse), nil
	}
}

type ListWorkspaceusersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkspaceusersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkspaceusersInvoker) Invoke() (*model.ListWorkspaceusersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkspaceusersResponse), nil
	}
}

type MigrateApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MigrateApiInvoker) Invoke() (*model.MigrateApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateApiResponse), nil
	}
}

type MigrateCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MigrateCatalogInvoker) Invoke() (*model.MigrateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateCatalogResponse), nil
	}
}

type ModifyCustomizedFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyCustomizedFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyCustomizedFieldsInvoker) Invoke() (*model.ModifyCustomizedFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyCustomizedFieldsResponse), nil
	}
}

type ParseUserBehaviorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseUserBehaviorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ParseUserBehaviorInvoker) Invoke() (*model.ParseUserBehaviorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseUserBehaviorResponse), nil
	}
}

type PayForDgcOneKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *PayForDgcOneKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PayForDgcOneKeyInvoker) Invoke() (*model.PayForDgcOneKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PayForDgcOneKeyResponse), nil
	}
}

type RemoveDesignEntityTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveDesignEntityTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveDesignEntityTagsInvoker) Invoke() (*model.RemoveDesignEntityTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveDesignEntityTagsResponse), nil
	}
}

type RemoveDesignQualityInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveDesignQualityInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RemoveDesignQualityInfosInvoker) Invoke() (*model.RemoveDesignQualityInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveDesignQualityInfosResponse), nil
	}
}

type RenewDataProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenewDataProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenewDataProfileInvoker) Invoke() (*model.RenewDataProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenewDataProfileResponse), nil
	}
}

type ResetLinkAttributeAndStandardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetLinkAttributeAndStandardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetLinkAttributeAndStandardInvoker) Invoke() (*model.ResetLinkAttributeAndStandardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetLinkAttributeAndStandardResponse), nil
	}
}

type RollbackApprovalInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollbackApprovalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RollbackApprovalInvoker) Invoke() (*model.RollbackApprovalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollbackApprovalResponse), nil
	}
}

type SearchApprovalsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchApprovalsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchApprovalsInvoker) Invoke() (*model.SearchApprovalsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchApprovalsResponse), nil
	}
}

type SearchAtomicIndexesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchAtomicIndexesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchAtomicIndexesInvoker) Invoke() (*model.SearchAtomicIndexesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchAtomicIndexesResponse), nil
	}
}

type SearchAuthorizeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchAuthorizeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchAuthorizeAppInvoker) Invoke() (*model.SearchAuthorizeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchAuthorizeAppResponse), nil
	}
}

type SearchBindApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchBindApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchBindApiInvoker) Invoke() (*model.SearchBindApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchBindApiResponse), nil
	}
}

type SearchCatalogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCatalogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchCatalogsInvoker) Invoke() (*model.SearchCatalogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCatalogsResponse), nil
	}
}

type SearchCodeTableValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCodeTableValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchCodeTableValuesInvoker) Invoke() (*model.SearchCodeTableValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCodeTableValuesResponse), nil
	}
}

type SearchCodeTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCodeTablesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchCodeTablesInvoker) Invoke() (*model.SearchCodeTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCodeTablesResponse), nil
	}
}

type SearchCustomizedFieldsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCustomizedFieldsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchCustomizedFieldsInvoker) Invoke() (*model.SearchCustomizedFieldsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCustomizedFieldsResponse), nil
	}
}

type SearchDesignLatestApprovalDiffInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDesignLatestApprovalDiffInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchDesignLatestApprovalDiffInvoker) Invoke() (*model.SearchDesignLatestApprovalDiffResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDesignLatestApprovalDiffResponse), nil
	}
}

type SearchDwByTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDwByTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchDwByTypeInvoker) Invoke() (*model.SearchDwByTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDwByTypeResponse), nil
	}
}

type SearchFieldsForRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchFieldsForRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchFieldsForRelationInvoker) Invoke() (*model.SearchFieldsForRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchFieldsForRelationResponse), nil
	}
}

type SearchIdByPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchIdByPathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchIdByPathInvoker) Invoke() (*model.SearchIdByPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchIdByPathResponse), nil
	}
}

type SearchSgcComputeDimensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSgcComputeDimensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchSgcComputeDimensionsInvoker) Invoke() (*model.SearchSgcComputeDimensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSgcComputeDimensionsResponse), nil
	}
}

type SearchSubjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSubjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchSubjectInvoker) Invoke() (*model.SearchSubjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSubjectResponse), nil
	}
}

type SearchSubjectNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchSubjectNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchSubjectNewInvoker) Invoke() (*model.SearchSubjectNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchSubjectNewResponse), nil
	}
}

type SearchVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchVersionsInvoker) Invoke() (*model.SearchVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchVersionsResponse), nil
	}
}

type SetFactoryJobTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetFactoryJobTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetFactoryJobTagsInvoker) Invoke() (*model.SetFactoryJobTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetFactoryJobTagsResponse), nil
	}
}

type ShowAggregationLogicTableByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregationLogicTableByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregationLogicTableByIdInvoker) Invoke() (*model.ShowAggregationLogicTableByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregationLogicTableByIdResponse), nil
	}
}

type ShowApiDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiDashboardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiDashboardInvoker) Invoke() (*model.ShowApiDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiDashboardResponse), nil
	}
}

type ShowApisDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApisDashboardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApisDashboardInvoker) Invoke() (*model.ShowApisDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApisDashboardResponse), nil
	}
}

type ShowApisDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApisDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApisDetailInvoker) Invoke() (*model.ShowApisDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApisDetailResponse), nil
	}
}

type ShowApisOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApisOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApisOverviewInvoker) Invoke() (*model.ShowApisOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApisOverviewResponse), nil
	}
}

type ShowAppInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppInfoInvoker) Invoke() (*model.ShowAppInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppInfoResponse), nil
	}
}

type ShowApplyDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplyDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApplyDetailInvoker) Invoke() (*model.ShowApplyDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplyDetailResponse), nil
	}
}

type ShowAppsDashboardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppsDashboardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppsDashboardInvoker) Invoke() (*model.ShowAppsDashboardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppsDashboardResponse), nil
	}
}

type ShowAppsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppsDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppsDetailInvoker) Invoke() (*model.ShowAppsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppsDetailResponse), nil
	}
}

type ShowAppsOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppsOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppsOverviewInvoker) Invoke() (*model.ShowAppsOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppsOverviewResponse), nil
	}
}

type ShowAtomicIndexByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAtomicIndexByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAtomicIndexByIdInvoker) Invoke() (*model.ShowAtomicIndexByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAtomicIndexByIdResponse), nil
	}
}

type ShowBizCatalogDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBizCatalogDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBizCatalogDetailInvoker) Invoke() (*model.ShowBizCatalogDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBizCatalogDetailResponse), nil
	}
}

type ShowBizMetricByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBizMetricByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBizMetricByIdInvoker) Invoke() (*model.ShowBizMetricByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBizMetricByIdResponse), nil
	}
}

type ShowBusinessAssetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBusinessAssetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBusinessAssetsInvoker) Invoke() (*model.ShowBusinessAssetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBusinessAssetsResponse), nil
	}
}

type ShowBusinessAssetsStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBusinessAssetsStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBusinessAssetsStatisticInvoker) Invoke() (*model.ShowBusinessAssetsStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBusinessAssetsStatisticResponse), nil
	}
}

type ShowBusinessAssetsTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBusinessAssetsTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBusinessAssetsTreeInvoker) Invoke() (*model.ShowBusinessAssetsTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBusinessAssetsTreeResponse), nil
	}
}

type ShowCatalogDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCatalogDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCatalogDetailInvoker) Invoke() (*model.ShowCatalogDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCatalogDetailResponse), nil
	}
}

type ShowCodeTableByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCodeTableByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCodeTableByIdInvoker) Invoke() (*model.ShowCodeTableByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCodeTableByIdResponse), nil
	}
}

type ShowCompoundMetricByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCompoundMetricByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCompoundMetricByIdInvoker) Invoke() (*model.ShowCompoundMetricByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCompoundMetricByIdResponse), nil
	}
}

type ShowConditionByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConditionByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConditionByIdInvoker) Invoke() (*model.ShowConditionByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConditionByIdResponse), nil
	}
}

type ShowConsistencyTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConsistencyTaskDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConsistencyTaskDetailInvoker) Invoke() (*model.ShowConsistencyTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConsistencyTaskDetailResponse), nil
	}
}

type ShowDataDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataDetailInvoker) Invoke() (*model.ShowDataDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataDetailResponse), nil
	}
}

type ShowDataProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataProfileInvoker) Invoke() (*model.ShowDataProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataProfileResponse), nil
	}
}

type ShowDataServiceInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataServiceInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataServiceInstanceInvoker) Invoke() (*model.ShowDataServiceInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataServiceInstanceResponse), nil
	}
}

type ShowDataSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataSetsInvoker) Invoke() (*model.ShowDataSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataSetsResponse), nil
	}
}

type ShowDataconnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataconnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataconnectionInvoker) Invoke() (*model.ShowDataconnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataconnectionResponse), nil
	}
}

type ShowDatamapLineageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatamapLineageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDatamapLineageInvoker) Invoke() (*model.ShowDatamapLineageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatamapLineageResponse), nil
	}
}

type ShowDerivativeIndexByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDerivativeIndexByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDerivativeIndexByIdInvoker) Invoke() (*model.ShowDerivativeIndexByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDerivativeIndexByIdResponse), nil
	}
}

type ShowDimensionByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDimensionByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDimensionByIdInvoker) Invoke() (*model.ShowDimensionByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDimensionByIdResponse), nil
	}
}

type ShowDimensionLogicTableByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDimensionLogicTableByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDimensionLogicTableByIdInvoker) Invoke() (*model.ShowDimensionLogicTableByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDimensionLogicTableByIdResponse), nil
	}
}

type ShowEntitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEntitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEntitiesInvoker) Invoke() (*model.ShowEntitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEntitiesResponse), nil
	}
}

type ShowEntityInfoByGuidInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEntityInfoByGuidInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEntityInfoByGuidInvoker) Invoke() (*model.ShowEntityInfoByGuidResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEntityInfoByGuidResponse), nil
	}
}

type ShowFactLogicTableByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactLogicTableByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFactLogicTableByIdInvoker) Invoke() (*model.ShowFactLogicTableByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactLogicTableByIdResponse), nil
	}
}

type ShowFactoryEnvInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactoryEnvInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFactoryEnvInvoker) Invoke() (*model.ShowFactoryEnvResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactoryEnvResponse), nil
	}
}

type ShowFactoryPackageDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactoryPackageDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFactoryPackageDetailInvoker) Invoke() (*model.ShowFactoryPackageDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactoryPackageDetailResponse), nil
	}
}

type ShowFactorySupplementDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFactorySupplementDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFactorySupplementDataInvoker) Invoke() (*model.ShowFactorySupplementDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFactorySupplementDataResponse), nil
	}
}

type ShowGlossaryListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGlossaryListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGlossaryListInvoker) Invoke() (*model.ShowGlossaryListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGlossaryListResponse), nil
	}
}

type ShowInstanceInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInfosInvoker) Invoke() (*model.ShowInstanceInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceInfosResponse), nil
	}
}

type ShowInstanceLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceLogInvoker) Invoke() (*model.ShowInstanceLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceLogResponse), nil
	}
}

type ShowInstanceResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceResultInvoker) Invoke() (*model.ShowInstanceResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResultResponse), nil
	}
}

type ShowLineageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLineageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLineageInvoker) Invoke() (*model.ShowLineageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLineageResponse), nil
	}
}

type ShowLineageBulkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLineageBulkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLineageBulkInvoker) Invoke() (*model.ShowLineageBulkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLineageBulkResponse), nil
	}
}

type ShowMessageDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMessageDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMessageDetailInvoker) Invoke() (*model.ShowMessageDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMessageDetailResponse), nil
	}
}

type ShowMetricAssetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricAssetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetricAssetsInvoker) Invoke() (*model.ShowMetricAssetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricAssetsResponse), nil
	}
}

type ShowMetricTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetricTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMetricTreeInvoker) Invoke() (*model.ShowMetricTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetricTreeResponse), nil
	}
}

type ShowNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNodesInvoker) Invoke() (*model.ShowNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodesResponse), nil
	}
}

type ShowPathByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPathByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPathByIdInvoker) Invoke() (*model.ShowPathByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPathByIdResponse), nil
	}
}

type ShowPathObjectByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPathObjectByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPathObjectByIdInvoker) Invoke() (*model.ShowPathObjectByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPathObjectByIdResponse), nil
	}
}

type ShowQualityTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQualityTaskDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQualityTaskDetailInvoker) Invoke() (*model.ShowQualityTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQualityTaskDetailResponse), nil
	}
}

type ShowRelationByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRelationByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRelationByIdInvoker) Invoke() (*model.ShowRelationByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRelationByIdResponse), nil
	}
}

type ShowSecurityDataClassificationRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityDataClassificationRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityDataClassificationRuleInvoker) Invoke() (*model.ShowSecurityDataClassificationRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityDataClassificationRuleResponse), nil
	}
}

type ShowSecurityDataClassificationRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityDataClassificationRuleGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityDataClassificationRuleGroupInvoker) Invoke() (*model.ShowSecurityDataClassificationRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityDataClassificationRuleGroupResponse), nil
	}
}

type ShowSecurityDynamicMaskingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityDynamicMaskingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityDynamicMaskingPolicyInvoker) Invoke() (*model.ShowSecurityDynamicMaskingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityDynamicMaskingPolicyResponse), nil
	}
}

type ShowSecurityMemberSyncTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityMemberSyncTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityMemberSyncTaskInvoker) Invoke() (*model.ShowSecurityMemberSyncTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityMemberSyncTaskResponse), nil
	}
}

type ShowSecurityPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecurityPermissionSetInvoker) Invoke() (*model.ShowSecurityPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityPermissionSetResponse), nil
	}
}

type ShowSecuritySecrecyLevelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecuritySecrecyLevelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSecuritySecrecyLevelInvoker) Invoke() (*model.ShowSecuritySecrecyLevelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecuritySecrecyLevelResponse), nil
	}
}

type ShowStandardByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStandardByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStandardByIdInvoker) Invoke() (*model.ShowStandardByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStandardByIdResponse), nil
	}
}

type ShowStandardTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStandardTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStandardTemplateInvoker) Invoke() (*model.ShowStandardTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStandardTemplateResponse), nil
	}
}

type ShowTableModelByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTableModelByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTableModelByIdInvoker) Invoke() (*model.ShowTableModelByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTableModelByIdResponse), nil
	}
}

type ShowTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTagsInvoker) Invoke() (*model.ShowTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTagsResponse), nil
	}
}

type ShowTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInfoInvoker) Invoke() (*model.ShowTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskInfoResponse), nil
	}
}

type ShowTaskListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskListInvoker) Invoke() (*model.ShowTaskListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskListResponse), nil
	}
}

type ShowTechnicalAssetsStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTechnicalAssetsStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTechnicalAssetsStatisticInvoker) Invoke() (*model.ShowTechnicalAssetsStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTechnicalAssetsStatisticResponse), nil
	}
}

type ShowTemplatesDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplatesDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplatesDetailInvoker) Invoke() (*model.ShowTemplatesDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplatesDetailResponse), nil
	}
}

type ShowUnrelatedTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUnrelatedTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUnrelatedTableInvoker) Invoke() (*model.ShowUnrelatedTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUnrelatedTableResponse), nil
	}
}

type ShowWorkSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkSpaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkSpaceInvoker) Invoke() (*model.ShowWorkSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkSpaceResponse), nil
	}
}

type ShowWorkspaceDetailByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkspaceDetailByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkspaceDetailByIdInvoker) Invoke() (*model.ShowWorkspaceDetailByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkspaceDetailByIdResponse), nil
	}
}

type StopFactorySupplementDataInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopFactorySupplementDataInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopFactorySupplementDataInstanceInvoker) Invoke() (*model.StopFactorySupplementDataInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopFactorySupplementDataInstanceResponse), nil
	}
}

type UpdateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type UpdateBizMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBizMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBizMetricInvoker) Invoke() (*model.UpdateBizMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBizMetricResponse), nil
	}
}

type UpdateCatalogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCatalogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCatalogInvoker) Invoke() (*model.UpdateCatalogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCatalogResponse), nil
	}
}

type UpdateCodeTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCodeTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCodeTableInvoker) Invoke() (*model.UpdateCodeTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCodeTableResponse), nil
	}
}

type UpdateCodeTableValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCodeTableValuesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCodeTableValuesInvoker) Invoke() (*model.UpdateCodeTableValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCodeTableValuesResponse), nil
	}
}

type UpdateDataconnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataconnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataconnectionInvoker) Invoke() (*model.UpdateDataconnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataconnectionResponse), nil
	}
}

type UpdateDesignAggregationLogicTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesignAggregationLogicTableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesignAggregationLogicTableInvoker) Invoke() (*model.UpdateDesignAggregationLogicTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesignAggregationLogicTableResponse), nil
	}
}

type UpdateDesignAtomicIndexInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesignAtomicIndexInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesignAtomicIndexInvoker) Invoke() (*model.UpdateDesignAtomicIndexResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesignAtomicIndexResponse), nil
	}
}

type UpdateDesignTableQualityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDesignTableQualityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDesignTableQualityInvoker) Invoke() (*model.UpdateDesignTableQualityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDesignTableQualityResponse), nil
	}
}

type UpdateDirectoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDirectoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDirectoryInvoker) Invoke() (*model.UpdateDirectoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDirectoryResponse), nil
	}
}

type UpdateFactoryJobNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFactoryJobNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFactoryJobNameInvoker) Invoke() (*model.UpdateFactoryJobNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFactoryJobNameResponse), nil
	}
}

type UpdateSecurityAssignedQueueInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityAssignedQueueInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityAssignedQueueInvoker) Invoke() (*model.UpdateSecurityAssignedQueueResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityAssignedQueueResponse), nil
	}
}

type UpdateSecurityDataClassificationRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityDataClassificationRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityDataClassificationRuleInvoker) Invoke() (*model.UpdateSecurityDataClassificationRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityDataClassificationRuleResponse), nil
	}
}

type UpdateSecurityDynamicMaskingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityDynamicMaskingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityDynamicMaskingPolicyInvoker) Invoke() (*model.UpdateSecurityDynamicMaskingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityDynamicMaskingPolicyResponse), nil
	}
}

type UpdateSecurityPermissionSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityPermissionSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityPermissionSetInvoker) Invoke() (*model.UpdateSecurityPermissionSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityPermissionSetResponse), nil
	}
}

type UpdateSecurityPermissionSetPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityPermissionSetPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityPermissionSetPermissionInvoker) Invoke() (*model.UpdateSecurityPermissionSetPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityPermissionSetPermissionResponse), nil
	}
}

type UpdateSecurityRuleEnableStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityRuleEnableStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityRuleEnableStatusInvoker) Invoke() (*model.UpdateSecurityRuleEnableStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityRuleEnableStatusResponse), nil
	}
}

type UpdateSecuritySecrecyLevelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecuritySecrecyLevelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecuritySecrecyLevelInvoker) Invoke() (*model.UpdateSecuritySecrecyLevelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecuritySecrecyLevelResponse), nil
	}
}

type UpdateStandardInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStandardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStandardInvoker) Invoke() (*model.UpdateStandardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStandardResponse), nil
	}
}

type UpdateStandardTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStandardTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStandardTemplateInvoker) Invoke() (*model.UpdateStandardTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStandardTemplateResponse), nil
	}
}

type UpdateSubjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubjectInvoker) Invoke() (*model.UpdateSubjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubjectResponse), nil
	}
}

type UpdateSubjectNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubjectNewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubjectNewInvoker) Invoke() (*model.UpdateSubjectNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubjectNewResponse), nil
	}
}

type UpdateTableModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTableModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTableModelInvoker) Invoke() (*model.UpdateTableModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTableModelResponse), nil
	}
}

type UpdateTaskInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskInfoInvoker) Invoke() (*model.UpdateTaskInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskInfoResponse), nil
	}
}

type UpdateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTemplateInvoker) Invoke() (*model.UpdateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTemplateResponse), nil
	}
}

type UpdateWorkSpaceUserOrGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkSpaceUserOrGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkSpaceUserOrGroupInvoker) Invoke() (*model.UpdateWorkSpaceUserOrGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkSpaceUserOrGroupResponse), nil
	}
}

type UpdateWorkspaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkspaceInvoker) Invoke() (*model.UpdateWorkspaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkspaceResponse), nil
	}
}

type AuthorizeActionApiToInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeActionApiToInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizeActionApiToInstanceInvoker) Invoke() (*model.AuthorizeActionApiToInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeActionApiToInstanceResponse), nil
	}
}

type AuthorizeApiToInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeApiToInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizeApiToInstanceInvoker) Invoke() (*model.AuthorizeApiToInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeApiToInstanceResponse), nil
	}
}

type CreateApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApiInvoker) Invoke() (*model.CreateApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApiResponse), nil
	}
}

type DebugApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *DebugApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DebugApiInvoker) Invoke() (*model.DebugApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DebugApiResponse), nil
	}
}

type DeleteApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApiInvoker) Invoke() (*model.DeleteApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApiResponse), nil
	}
}

type ExecuteApiToInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteApiToInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteApiToInstanceInvoker) Invoke() (*model.ExecuteApiToInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteApiToInstanceResponse), nil
	}
}

type ExportDataServiceExcelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDataServiceExcelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDataServiceExcelInvoker) Invoke() (*model.ExportDataServiceExcelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDataServiceExcelResponse), nil
	}
}

type ExportDataServiceExcelTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDataServiceExcelTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDataServiceExcelTemplateInvoker) Invoke() (*model.ExportDataServiceExcelTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDataServiceExcelTemplateResponse), nil
	}
}

type ExportDataServiceZipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportDataServiceZipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportDataServiceZipInvoker) Invoke() (*model.ExportDataServiceZipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportDataServiceZipResponse), nil
	}
}

type ImportDataServiceExcelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportDataServiceExcelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportDataServiceExcelInvoker) Invoke() (*model.ImportDataServiceExcelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportDataServiceExcelResponse), nil
	}
}

type ListApisInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApisInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApisInvoker) Invoke() (*model.ListApisResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApisResponse), nil
	}
}

type ListInstanceListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstanceListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstanceListInvoker) Invoke() (*model.ListInstanceListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstanceListResponse), nil
	}
}

type PublishApiToInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishApiToInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishApiToInstanceInvoker) Invoke() (*model.PublishApiToInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishApiToInstanceResponse), nil
	}
}

type SearchDebugInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDebugInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchDebugInfoInvoker) Invoke() (*model.SearchDebugInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDebugInfoResponse), nil
	}
}

type SearchPublishInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchPublishInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchPublishInfoInvoker) Invoke() (*model.SearchPublishInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchPublishInfoResponse), nil
	}
}

type ShowApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiInvoker) Invoke() (*model.ShowApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiResponse), nil
	}
}

type UpdateApiInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateApiInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateApiInvoker) Invoke() (*model.UpdateApiResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateApiResponse), nil
	}
}
