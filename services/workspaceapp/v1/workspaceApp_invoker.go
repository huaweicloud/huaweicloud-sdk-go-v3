package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/workspaceapp/v1/model"
)

type ListPublishedAppInvoker struct {
	*invoker.BaseInvoker
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

func (i *PublishAppInvoker) Invoke() (*model.PublishAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishAppResponse), nil
	}
}

type UnpublishAppInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateAppInvoker) Invoke() (*model.UpdateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppResponse), nil
	}
}

type BatchDeleteAppGroupInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProductInvoker) Invoke() (*model.ListProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductResponse), nil
	}
}

type AddAppGroupAuthorizationInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAvailabilityZoneInvoker) Invoke() (*model.ListAvailabilityZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZoneResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type CreatePersistentStorageInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateAppServersInvoker) Invoke() (*model.CreateAppServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppServersResponse), nil
	}
}

type ListServersInvoker struct {
	*invoker.BaseInvoker
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

func (i *ReinstallServerInvoker) Invoke() (*model.ReinstallServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallServerResponse), nil
	}
}

type UpdateServerInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAppConnectionInvoker) Invoke() (*model.ListAppConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppConnectionResponse), nil
	}
}

type ListUserConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUserConnectionInvoker) Invoke() (*model.ListUserConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUserConnectionResponse), nil
	}
}

type ListVolumeTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumeTypeInvoker) Invoke() (*model.ListVolumeTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumeTypeResponse), nil
	}
}
