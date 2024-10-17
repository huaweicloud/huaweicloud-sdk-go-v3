package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/workspaceapp/v1/model"
)

type AuthorizeObsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeObsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizeObsInvoker) Invoke() (*model.AuthorizeObsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeObsResponse), nil
	}
}

type BatchDeleteWarehouseAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteWarehouseAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteWarehouseAppInvoker) Invoke() (*model.BatchDeleteWarehouseAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteWarehouseAppResponse), nil
	}
}

type CreateWarehouseAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWarehouseAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWarehouseAppInvoker) Invoke() (*model.CreateWarehouseAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWarehouseAppResponse), nil
	}
}

type ListWarehouseAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWarehouseAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWarehouseAppsInvoker) Invoke() (*model.ListWarehouseAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWarehouseAppsResponse), nil
	}
}

type UpdateWarehouseAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWarehouseAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWarehouseAppInvoker) Invoke() (*model.UpdateWarehouseAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWarehouseAppResponse), nil
	}
}

type UploadWarehouseAppIconInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadWarehouseAppIconInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadWarehouseAppIconInvoker) Invoke() (*model.UploadWarehouseAppIconResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadWarehouseAppIconResponse), nil
	}
}

type ListPublishedAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublishedAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublishedAppInvoker) Invoke() (*model.ListPublishedAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublishedAppResponse), nil
	}
}

type PublishAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishAppInvoker) Invoke() (*model.PublishAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishAppResponse), nil
	}
}

type ShowPublishableAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublishableAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublishableAppInvoker) Invoke() (*model.ShowPublishableAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublishableAppResponse), nil
	}
}

type UnpublishAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnpublishAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnpublishAppInvoker) Invoke() (*model.UnpublishAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnpublishAppResponse), nil
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

type UploadAppIconInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAppIconInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAppIconInvoker) Invoke() (*model.UploadAppIconResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAppIconResponse), nil
	}
}

type BatchDeleteAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppGroupInvoker) Invoke() (*model.BatchDeleteAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppGroupResponse), nil
	}
}

type CreateAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppGroupInvoker) Invoke() (*model.CreateAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppGroupResponse), nil
	}
}

type ListAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppGroupInvoker) Invoke() (*model.ListAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppGroupResponse), nil
	}
}

type UpdateAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppGroupInvoker) Invoke() (*model.UpdateAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppGroupResponse), nil
	}
}

type ListProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductInvoker) Invoke() (*model.ListProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductResponse), nil
	}
}

type ListSessionTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSessionTypeInvoker) Invoke() (*model.ListSessionTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionTypeResponse), nil
	}
}

type AddAppGroupAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAppGroupAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddAppGroupAuthorizationInvoker) Invoke() (*model.AddAppGroupAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAppGroupAuthorizationResponse), nil
	}
}

type BatchDeleteAppGroupAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppGroupAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppGroupAuthorizationInvoker) Invoke() (*model.BatchDeleteAppGroupAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppGroupAuthorizationResponse), nil
	}
}

type ListAppGroupAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppGroupAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppGroupAuthorizationInvoker) Invoke() (*model.ListAppGroupAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppGroupAuthorizationResponse), nil
	}
}

type ListAvailabilityZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailabilityZoneInvoker) Invoke() (*model.ListAvailabilityZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneResponse), nil
	}
}

type AttachImageServerAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachImageServerAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AttachImageServerAppInvoker) Invoke() (*model.AttachImageServerAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachImageServerAppResponse), nil
	}
}

type BatchDeleteImageServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteImageServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteImageServerInvoker) Invoke() (*model.BatchDeleteImageServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteImageServerResponse), nil
	}
}

type CreateImageServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateImageServerInvoker) Invoke() (*model.CreateImageServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageServerResponse), nil
	}
}

type ListImageServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageServersInvoker) Invoke() (*model.ListImageServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageServersResponse), nil
	}
}

type ListLatestAttachedServerAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLatestAttachedServerAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLatestAttachedServerAppInvoker) Invoke() (*model.ListLatestAttachedServerAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLatestAttachedServerAppResponse), nil
	}
}

type RecreateServerImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecreateServerImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RecreateServerImageInvoker) Invoke() (*model.RecreateServerImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecreateServerImageResponse), nil
	}
}

type UpdateImageServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateImageServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateImageServerInvoker) Invoke() (*model.UpdateImageServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateImageServerResponse), nil
	}
}

type ShowImageJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageJobInvoker) Invoke() (*model.ShowImageJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageJobResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type CreateOrUpdateStoragePolicyStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateStoragePolicyStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateStoragePolicyStatementInvoker) Invoke() (*model.CreateOrUpdateStoragePolicyStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateStoragePolicyStatementResponse), nil
	}
}

type CreatePersistentStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePersistentStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePersistentStorageInvoker) Invoke() (*model.CreatePersistentStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePersistentStorageResponse), nil
	}
}

type CreateShareFolderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareFolderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShareFolderInvoker) Invoke() (*model.CreateShareFolderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareFolderResponse), nil
	}
}

type DeletePersistentStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePersistentStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePersistentStorageInvoker) Invoke() (*model.DeletePersistentStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePersistentStorageResponse), nil
	}
}

type DeleteStorageClaimInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStorageClaimInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStorageClaimInvoker) Invoke() (*model.DeleteStorageClaimResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStorageClaimResponse), nil
	}
}

type DeleteUserStorageAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserStorageAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserStorageAttachmentInvoker) Invoke() (*model.DeleteUserStorageAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserStorageAttachmentResponse), nil
	}
}

type ListPersistentStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPersistentStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPersistentStorageInvoker) Invoke() (*model.ListPersistentStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPersistentStorageResponse), nil
	}
}

type ListShareFolderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListShareFolderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListShareFolderInvoker) Invoke() (*model.ListShareFolderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListShareFolderResponse), nil
	}
}

type ListStorageAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStorageAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStorageAssignmentInvoker) Invoke() (*model.ListStorageAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStorageAssignmentResponse), nil
	}
}

type ListStoragePolicyStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoragePolicyStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStoragePolicyStatementInvoker) Invoke() (*model.ListStoragePolicyStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoragePolicyStatementResponse), nil
	}
}

type UpdateShareFolderAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateShareFolderAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateShareFolderAssignmentInvoker) Invoke() (*model.UpdateShareFolderAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateShareFolderAssignmentResponse), nil
	}
}

type UpdateUserFolderAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserFolderAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserFolderAssignmentInvoker) Invoke() (*model.UpdateUserFolderAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserFolderAssignmentResponse), nil
	}
}

type CreatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyGroupInvoker) Invoke() (*model.CreatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyGroupResponse), nil
	}
}

type CreatePolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyTemplateInvoker) Invoke() (*model.CreatePolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyTemplateResponse), nil
	}
}

type DeletePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyGroupInvoker) Invoke() (*model.DeletePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyGroupResponse), nil
	}
}

type DeletePolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyTemplateInvoker) Invoke() (*model.DeletePolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyTemplateResponse), nil
	}
}

type ListPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyGroupInvoker) Invoke() (*model.ListPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupResponse), nil
	}
}

type ListPolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyTemplateInvoker) Invoke() (*model.ListPolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyTemplateResponse), nil
	}
}

type ListTargetsOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTargetsOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTargetsOfPolicyGroupInvoker) Invoke() (*model.ListTargetsOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTargetsOfPolicyGroupResponse), nil
	}
}

type ShowOriginalPolicyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOriginalPolicyInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOriginalPolicyInfoInvoker) Invoke() (*model.ShowOriginalPolicyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOriginalPolicyInfoResponse), nil
	}
}

type UpdatePolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyGroupInvoker) Invoke() (*model.UpdatePolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyGroupResponse), nil
	}
}

type UpdatePolicyTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyTemplateInvoker) Invoke() (*model.UpdatePolicyTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyTemplateResponse), nil
	}
}

type CheckQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckQuotaInvoker) Invoke() (*model.CheckQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckQuotaResponse), nil
	}
}

type BatchDeleteServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteServerInvoker) Invoke() (*model.BatchDeleteServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteServerResponse), nil
	}
}

type BatchMigrateHostsServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchMigrateHostsServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchMigrateHostsServerInvoker) Invoke() (*model.BatchMigrateHostsServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchMigrateHostsServerResponse), nil
	}
}

type BatchRebootServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebootServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRebootServerInvoker) Invoke() (*model.BatchRebootServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebootServerResponse), nil
	}
}

type BatchRejoinDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRejoinDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRejoinDomainInvoker) Invoke() (*model.BatchRejoinDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRejoinDomainResponse), nil
	}
}

type BatchStartServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartServerInvoker) Invoke() (*model.BatchStartServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartServerResponse), nil
	}
}

type BatchStopServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStopServerInvoker) Invoke() (*model.BatchStopServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopServerResponse), nil
	}
}

type BatchUpdateTsviInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateTsviInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateTsviInvoker) Invoke() (*model.BatchUpdateTsviResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateTsviResponse), nil
	}
}

type ChangeServerImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeServerImageInvoker) Invoke() (*model.ChangeServerImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerImageResponse), nil
	}
}

type CreateAppServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppServersInvoker) Invoke() (*model.CreateAppServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppServersResponse), nil
	}
}

type ListServerMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerMetricDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerMetricDataInvoker) Invoke() (*model.ListServerMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerMetricDataResponse), nil
	}
}

type ListServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersInvoker) Invoke() (*model.ListServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersResponse), nil
	}
}

type ReinstallServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReinstallServerInvoker) Invoke() (*model.ReinstallServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallServerResponse), nil
	}
}

type ShowServerVncInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerVncInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerVncInvoker) Invoke() (*model.ShowServerVncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerVncResponse), nil
	}
}

type UpdateServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServerInvoker) Invoke() (*model.UpdateServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerResponse), nil
	}
}

type CreateServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServerGroupInvoker) Invoke() (*model.CreateServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServerGroupResponse), nil
	}
}

type DeleteServerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServerGroupsInvoker) Invoke() (*model.DeleteServerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerGroupsResponse), nil
	}
}

type ListServerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerGroupsInvoker) Invoke() (*model.ListServerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerGroupsResponse), nil
	}
}

type UpdateServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateServerGroupInvoker) Invoke() (*model.UpdateServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerGroupResponse), nil
	}
}

type ListAppConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppConnectionInvoker) Invoke() (*model.ListAppConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppConnectionResponse), nil
	}
}

type ListSessionByUserNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionByUserNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSessionByUserNameInvoker) Invoke() (*model.ListSessionByUserNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionByUserNameResponse), nil
	}
}

type ListSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSessionsInvoker) Invoke() (*model.ListSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSessionsResponse), nil
	}
}

type ListUserConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUserConnectionInvoker) Invoke() (*model.ListUserConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserConnectionResponse), nil
	}
}

type LogoffUserSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *LogoffUserSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LogoffUserSessionInvoker) Invoke() (*model.LogoffUserSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogoffUserSessionResponse), nil
	}
}

type ListVolumeTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumeTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVolumeTypeInvoker) Invoke() (*model.ListVolumeTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumeTypeResponse), nil
	}
}
