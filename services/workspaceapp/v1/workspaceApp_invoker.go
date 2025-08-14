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

type BindAppWarehouseBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *BindAppWarehouseBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BindAppWarehouseBucketInvoker) Invoke() (*model.BindAppWarehouseBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BindAppWarehouseBucketResponse), nil
	}
}

type CreateBucketOrAclInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBucketOrAclInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBucketOrAclInvoker) Invoke() (*model.CreateBucketOrAclResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBucketOrAclResponse), nil
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

type DeleteWarehouseAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWarehouseAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWarehouseAppInvoker) Invoke() (*model.DeleteWarehouseAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWarehouseAppResponse), nil
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

type ShowAppWarehouseBucketInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppWarehouseBucketInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppWarehouseBucketInvoker) Invoke() (*model.ShowAppWarehouseBucketResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppWarehouseBucketResponse), nil
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

type BatchDisableAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisableAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisableAppInvoker) Invoke() (*model.BatchDisableAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisableAppResponse), nil
	}
}

type BatchEnableAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableAppInvoker) Invoke() (*model.BatchEnableAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableAppResponse), nil
	}
}

type DeleteAppIconInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppIconInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppIconInvoker) Invoke() (*model.DeleteAppIconResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppIconResponse), nil
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

type ShowAppDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppDetailInvoker) Invoke() (*model.ShowAppDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppDetailResponse), nil
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

type UpdatePreBootPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePreBootPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePreBootPolicyInvoker) Invoke() (*model.UpdatePreBootPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePreBootPolicyResponse), nil
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

type InitializeTenantInvoker struct {
	*invoker.BaseInvoker
}

func (i *InitializeTenantInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InitializeTenantInvoker) Invoke() (*model.InitializeTenantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InitializeTenantResponse), nil
	}
}

type ListCorpConfigInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCorpConfigInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCorpConfigInfoInvoker) Invoke() (*model.ListCorpConfigInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCorpConfigInfoResponse), nil
	}
}

type ListTenantProfileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantProfileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantProfileInvoker) Invoke() (*model.ListTenantProfileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantProfileResponse), nil
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

type DeleteAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppGroupInvoker) Invoke() (*model.DeleteAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppGroupResponse), nil
	}
}

type DisassociateAppGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateAppGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateAppGroupInvoker) Invoke() (*model.DisassociateAppGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateAppGroupResponse), nil
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

type ShowAppGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppGroupDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppGroupDetailInvoker) Invoke() (*model.ShowAppGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppGroupDetailResponse), nil
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

type CreateOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrderInvoker) Invoke() (*model.CreateOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrderResponse), nil
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

type ShowSessionTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSessionTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSessionTypesInvoker) Invoke() (*model.ShowSessionTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSessionTypesResponse), nil
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

type ListAzInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAzInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAzInvoker) Invoke() (*model.ListAzResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAzResponse), nil
	}
}

type BatchDeleteCloudStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteCloudStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteCloudStorageInvoker) Invoke() (*model.BatchDeleteCloudStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteCloudStorageResponse), nil
	}
}

type CreateCloudStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCloudStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCloudStorageInvoker) Invoke() (*model.CreateCloudStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCloudStorageResponse), nil
	}
}

type CreateUserFolderAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserFolderAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserFolderAssignmentInvoker) Invoke() (*model.CreateUserFolderAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserFolderAssignmentResponse), nil
	}
}

type DeleteCloudStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCloudStorageInvoker) Invoke() (*model.DeleteCloudStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudStorageResponse), nil
	}
}

type DeleteCloudStorageAttachmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCloudStorageAttachmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCloudStorageAttachmentInvoker) Invoke() (*model.DeleteCloudStorageAttachmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCloudStorageAttachmentResponse), nil
	}
}

type ListCloudStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudStorageInvoker) Invoke() (*model.ListCloudStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudStorageResponse), nil
	}
}

type ListCloudStorageAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudStorageAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCloudStorageAssignmentInvoker) Invoke() (*model.ListCloudStorageAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudStorageAssignmentResponse), nil
	}
}

type ListProjectConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectConfigsInvoker) Invoke() (*model.ListProjectConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectConfigsResponse), nil
	}
}

type ShowProjectConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectConfigInvoker) Invoke() (*model.ShowProjectConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectConfigResponse), nil
	}
}

type UpdateCloudUserFolderAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCloudUserFolderAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCloudUserFolderAssignmentInvoker) Invoke() (*model.UpdateCloudUserFolderAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCloudUserFolderAssignmentResponse), nil
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

type ShowImageServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageServerInvoker) Invoke() (*model.ShowImageServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageServerResponse), nil
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

type BatchDeleteAppSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppSubJobsInvoker) Invoke() (*model.BatchDeleteAppSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppSubJobsResponse), nil
	}
}

type BatchDeleteImageSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteImageSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteImageSubJobsInvoker) Invoke() (*model.BatchDeleteImageSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteImageSubJobsResponse), nil
	}
}

type CountSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountSubJobsInvoker) Invoke() (*model.CountSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountSubJobsResponse), nil
	}
}

type ListImageJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageJobsInvoker) Invoke() (*model.ListImageJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageJobsResponse), nil
	}
}

type ListImageSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImageSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListImageSubJobsInvoker) Invoke() (*model.ListImageSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImageSubJobsResponse), nil
	}
}

type ListSubJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubJobsInvoker) Invoke() (*model.ListSubJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubJobsResponse), nil
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

type ListAuthorizationMailRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorizationMailRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuthorizationMailRecordInvoker) Invoke() (*model.ListAuthorizationMailRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizationMailRecordResponse), nil
	}
}

type SendAuthorizationMailInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendAuthorizationMailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendAuthorizationMailInvoker) Invoke() (*model.SendAuthorizationMailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendAuthorizationMailResponse), nil
	}
}

type SendAuthorizedMailInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendAuthorizedMailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendAuthorizedMailInvoker) Invoke() (*model.SendAuthorizedMailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendAuthorizedMailResponse), nil
	}
}

type BatchDeletePersistentStorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePersistentStorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePersistentStorageInvoker) Invoke() (*model.BatchDeletePersistentStorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePersistentStorageResponse), nil
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

type ListSfs3StorageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSfs3StorageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSfs3StorageInvoker) Invoke() (*model.ListSfs3StorageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSfs3StorageResponse), nil
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

type ListPolicyGroupDetailInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyGroupDetailInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyGroupDetailInfoInvoker) Invoke() (*model.ListPolicyGroupDetailInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyGroupDetailInfoResponse), nil
	}
}

type ListPolicyOfPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyOfPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyOfPolicyGroupInvoker) Invoke() (*model.ListPolicyOfPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyOfPolicyGroupResponse), nil
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

type ShowPolicyGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyGroupInvoker) Invoke() (*model.ShowPolicyGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyGroupResponse), nil
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

type CreateOrUpdateScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateScalingPolicyInvoker) Invoke() (*model.CreateOrUpdateScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateScalingPolicyResponse), nil
	}
}

type DeleteScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScalingPolicyInvoker) Invoke() (*model.DeleteScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScalingPolicyResponse), nil
	}
}

type ShowScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScalingPolicyInvoker) Invoke() (*model.ShowScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScalingPolicyResponse), nil
	}
}

type BatchDeleteScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteScheduleTaskInvoker) Invoke() (*model.BatchDeleteScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteScheduleTaskResponse), nil
	}
}

type CreateScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScheduleTaskInvoker) Invoke() (*model.CreateScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduleTaskResponse), nil
	}
}

type DeleteScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteScheduleTaskInvoker) Invoke() (*model.DeleteScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleTaskResponse), nil
	}
}

type ListFutureExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFutureExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFutureExecutionsInvoker) Invoke() (*model.ListFutureExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFutureExecutionsResponse), nil
	}
}

type ListScheduleTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScheduleTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListScheduleTasksInvoker) Invoke() (*model.ListScheduleTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScheduleTasksResponse), nil
	}
}

type ListTaskExecuteDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskExecuteDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskExecuteDetailInvoker) Invoke() (*model.ListTaskExecuteDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskExecuteDetailResponse), nil
	}
}

type ListTaskExecuteHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskExecuteHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskExecuteHistoryInvoker) Invoke() (*model.ListTaskExecuteHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskExecuteHistoryResponse), nil
	}
}

type ShowScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowScheduleTaskInvoker) Invoke() (*model.ShowScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScheduleTaskResponse), nil
	}
}

type UpdateScheduleTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduleTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateScheduleTaskInvoker) Invoke() (*model.UpdateScheduleTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduleTaskResponse), nil
	}
}

type BatchChangeServerImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeServerImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchChangeServerImageInvoker) Invoke() (*model.BatchChangeServerImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeServerImageResponse), nil
	}
}

type BatchChangeServerMaintainModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeServerMaintainModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchChangeServerMaintainModeInvoker) Invoke() (*model.BatchChangeServerMaintainModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeServerMaintainModeResponse), nil
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

type BatchReinstallServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchReinstallServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchReinstallServerInvoker) Invoke() (*model.BatchReinstallServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchReinstallServerResponse), nil
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

type BatchUpgradeHdaVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpgradeHdaVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpgradeHdaVersionInvoker) Invoke() (*model.BatchUpgradeHdaVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpgradeHdaVersionResponse), nil
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

type DeleteServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServerInvoker) Invoke() (*model.DeleteServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerResponse), nil
	}
}

type ListAccessAgentLatestVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessAgentLatestVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessAgentLatestVersionInvoker) Invoke() (*model.ListAccessAgentLatestVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessAgentLatestVersionResponse), nil
	}
}

type ListServerHdaDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerHdaDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerHdaDetailsInvoker) Invoke() (*model.ListServerHdaDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerHdaDetailsResponse), nil
	}
}

type ListServerHdaUpgradeRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerHdaUpgradeRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerHdaUpgradeRecordsInvoker) Invoke() (*model.ListServerHdaUpgradeRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerHdaUpgradeRecordsResponse), nil
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

type ShowAccessAgentLatestVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessAgentLatestVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessAgentLatestVersionInvoker) Invoke() (*model.ShowAccessAgentLatestVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessAgentLatestVersionResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ShowServerMetricDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerMetricDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerMetricDataInvoker) Invoke() (*model.ShowServerMetricDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerMetricDataResponse), nil
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

type ListTenantServerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantServerGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantServerGroupsInvoker) Invoke() (*model.ListTenantServerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantServerGroupsResponse), nil
	}
}

type ShowServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerGroupInvoker) Invoke() (*model.ShowServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerGroupResponse), nil
	}
}

type ShowServerGroupRestrictInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerGroupRestrictInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerGroupRestrictInvoker) Invoke() (*model.ShowServerGroupRestrictResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerGroupRestrictResponse), nil
	}
}

type ShowServerGroupStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerGroupStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerGroupStateInvoker) Invoke() (*model.ShowServerGroupStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerGroupStateResponse), nil
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

type BatchCreateServerGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateServerGroupTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateServerGroupTagsInvoker) Invoke() (*model.BatchCreateServerGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateServerGroupTagsResponse), nil
	}
}

type BatchDeleteServerGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteServerGroupTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteServerGroupTagsInvoker) Invoke() (*model.BatchDeleteServerGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteServerGroupTagsResponse), nil
	}
}

type CreateServerGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServerGroupTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateServerGroupTagsInvoker) Invoke() (*model.CreateServerGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServerGroupTagsResponse), nil
	}
}

type DeleteServerGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerGroupTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteServerGroupTagsInvoker) Invoke() (*model.DeleteServerGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerGroupTagsResponse), nil
	}
}

type ListServerGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerGroupTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerGroupTagInvoker) Invoke() (*model.ListServerGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerGroupTagResponse), nil
	}
}

type ShowServerGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerGroupTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerGroupTagInvoker) Invoke() (*model.ShowServerGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerGroupTagResponse), nil
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
