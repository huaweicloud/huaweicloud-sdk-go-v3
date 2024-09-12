package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/meeting/v1/model"
)

type AddAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAppIdInvoker) Invoke() (*model.AddAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAppIdResponse), nil
	}
}

type AddCorpInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddCorpInvoker) Invoke() (*model.AddCorpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddCorpResponse), nil
	}
}

type AddCorpAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddCorpAdminInvoker) Invoke() (*model.AddCorpAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddCorpAdminResponse), nil
	}
}

type AddDepartmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDepartmentInvoker) Invoke() (*model.AddDepartmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDepartmentResponse), nil
	}
}

type AddDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDeviceInvoker) Invoke() (*model.AddDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDeviceResponse), nil
	}
}

type AddMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddMaterialInvoker) Invoke() (*model.AddMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddMaterialResponse), nil
	}
}

type AddProgramInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddProgramInvoker) Invoke() (*model.AddProgramResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddProgramResponse), nil
	}
}

type AddPublicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddPublicationInvoker) Invoke() (*model.AddPublicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddPublicationResponse), nil
	}
}

type AddResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddResourceInvoker) Invoke() (*model.AddResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddResourceResponse), nil
	}
}

type AddToPersonalSpaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddToPersonalSpaceInvoker) Invoke() (*model.AddToPersonalSpaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddToPersonalSpaceResponse), nil
	}
}

type AddUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddUserInvoker) Invoke() (*model.AddUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddUserResponse), nil
	}
}

type AllowClientRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *AllowClientRecordInvoker) Invoke() (*model.AllowClientRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowClientRecordResponse), nil
	}
}

type AllowGuestUnmuteInvoker struct {
	*invoker.BaseInvoker
}

func (i *AllowGuestUnmuteInvoker) Invoke() (*model.AllowGuestUnmuteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowGuestUnmuteResponse), nil
	}
}

type AllowWaitingParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *AllowWaitingParticipantInvoker) Invoke() (*model.AllowWaitingParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllowWaitingParticipantResponse), nil
	}
}

type AssociateVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateVmrInvoker) Invoke() (*model.AssociateVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateVmrResponse), nil
	}
}

type BatchDeleteCorpAdminsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteCorpAdminsInvoker) Invoke() (*model.BatchDeleteCorpAdminsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteCorpAdminsResponse), nil
	}
}

type BatchDeleteDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteDevicesInvoker) Invoke() (*model.BatchDeleteDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteDevicesResponse), nil
	}
}

type BatchDeleteMaterialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMaterialsInvoker) Invoke() (*model.BatchDeleteMaterialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMaterialsResponse), nil
	}
}

type BatchDeleteProgramsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteProgramsInvoker) Invoke() (*model.BatchDeleteProgramsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteProgramsResponse), nil
	}
}

type BatchDeletePublicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePublicationsInvoker) Invoke() (*model.BatchDeletePublicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePublicationsResponse), nil
	}
}

type BatchDeleteUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteUsersInvoker) Invoke() (*model.BatchDeleteUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteUsersResponse), nil
	}
}

type BatchHandInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchHandInvoker) Invoke() (*model.BatchHandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchHandResponse), nil
	}
}

type BatchSearchAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSearchAppIdInvoker) Invoke() (*model.BatchSearchAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSearchAppIdResponse), nil
	}
}

type BatchShowUserDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowUserDetailsInvoker) Invoke() (*model.BatchShowUserDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowUserDetailsResponse), nil
	}
}

type BatchUpdateDevicesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateDevicesStatusInvoker) Invoke() (*model.BatchUpdateDevicesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateDevicesStatusResponse), nil
	}
}

type BatchUpdateUserStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateUserStatusInvoker) Invoke() (*model.BatchUpdateUserStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateUserStatusResponse), nil
	}
}

type BroadcastParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *BroadcastParticipantInvoker) Invoke() (*model.BroadcastParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BroadcastParticipantResponse), nil
	}
}

type CancelBroadcastInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelBroadcastInvoker) Invoke() (*model.CancelBroadcastResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelBroadcastResponse), nil
	}
}

type CancelMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelMeetingInvoker) Invoke() (*model.CancelMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelMeetingResponse), nil
	}
}

type CancelRecurringMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRecurringMeetingInvoker) Invoke() (*model.CancelRecurringMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRecurringMeetingResponse), nil
	}
}

type CancelRecurringSubMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRecurringSubMeetingInvoker) Invoke() (*model.CancelRecurringSubMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRecurringSubMeetingResponse), nil
	}
}

type CheckSlideVerifyCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckSlideVerifyCodeInvoker) Invoke() (*model.CheckSlideVerifyCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckSlideVerifyCodeResponse), nil
	}
}

type CheckTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTokenInvoker) Invoke() (*model.CheckTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTokenResponse), nil
	}
}

type CheckVeriCodeForUpdateUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckVeriCodeForUpdateUserInfoInvoker) Invoke() (*model.CheckVeriCodeForUpdateUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckVeriCodeForUpdateUserInfoResponse), nil
	}
}

type CheckVerifyCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckVerifyCodeInvoker) Invoke() (*model.CheckVerifyCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckVerifyCodeResponse), nil
	}
}

type CreateAnonymousAuthRandomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAnonymousAuthRandomInvoker) Invoke() (*model.CreateAnonymousAuthRandomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAnonymousAuthRandomResponse), nil
	}
}

type CreateAuthRandomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuthRandomInvoker) Invoke() (*model.CreateAuthRandomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuthRandomResponse), nil
	}
}

type CreateConfTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfTokenInvoker) Invoke() (*model.CreateConfTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfTokenResponse), nil
	}
}

type CreateMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMeetingInvoker) Invoke() (*model.CreateMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMeetingResponse), nil
	}
}

type CreatePortalRefNonceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePortalRefNonceInvoker) Invoke() (*model.CreatePortalRefNonceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePortalRefNonceResponse), nil
	}
}

type CreateRecurringMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecurringMeetingInvoker) Invoke() (*model.CreateRecurringMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecurringMeetingResponse), nil
	}
}

type CreateVisionActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVisionActiveCodeInvoker) Invoke() (*model.CreateVisionActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVisionActiveCodeResponse), nil
	}
}

type CreateWebSocketTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWebSocketTokenInvoker) Invoke() (*model.CreateWebSocketTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWebSocketTokenResponse), nil
	}
}

type CreateWebinarInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWebinarInvoker) Invoke() (*model.CreateWebinarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWebinarResponse), nil
	}
}

type DeleteAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppIdInvoker) Invoke() (*model.DeleteAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppIdResponse), nil
	}
}

type DeleteAttendeesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAttendeesInvoker) Invoke() (*model.DeleteAttendeesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAttendeesResponse), nil
	}
}

type DeleteCorpInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCorpInvoker) Invoke() (*model.DeleteCorpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCorpResponse), nil
	}
}

type DeleteCorpVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCorpVmrInvoker) Invoke() (*model.DeleteCorpVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCorpVmrResponse), nil
	}
}

type DeleteDepartmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDepartmentInvoker) Invoke() (*model.DeleteDepartmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDepartmentResponse), nil
	}
}

type DeleteLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLayoutInvoker) Invoke() (*model.DeleteLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLayoutResponse), nil
	}
}

type DeleteRecordingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecordingsInvoker) Invoke() (*model.DeleteRecordingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecordingsResponse), nil
	}
}

type DeleteResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInvoker) Invoke() (*model.DeleteResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceResponse), nil
	}
}

type DeleteTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTokenInvoker) Invoke() (*model.DeleteTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTokenResponse), nil
	}
}

type DeleteVisionActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVisionActiveCodeInvoker) Invoke() (*model.DeleteVisionActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVisionActiveCodeResponse), nil
	}
}

type DeleteWebHookConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWebHookConfigInvoker) Invoke() (*model.DeleteWebHookConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWebHookConfigResponse), nil
	}
}

type DeleteWebinarInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWebinarInvoker) Invoke() (*model.DeleteWebinarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWebinarResponse), nil
	}
}

type DisassociateVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateVmrInvoker) Invoke() (*model.DisassociateVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateVmrResponse), nil
	}
}

type HandInvoker struct {
	*invoker.BaseInvoker
}

func (i *HandInvoker) Invoke() (*model.HandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HandResponse), nil
	}
}

type HangUpInvoker struct {
	*invoker.BaseInvoker
}

func (i *HangUpInvoker) Invoke() (*model.HangUpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.HangUpResponse), nil
	}
}

type InviteOperateVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteOperateVideoInvoker) Invoke() (*model.InviteOperateVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteOperateVideoResponse), nil
	}
}

type InviteParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteParticipantInvoker) Invoke() (*model.InviteParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteParticipantResponse), nil
	}
}

type InviteShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteShareInvoker) Invoke() (*model.InviteShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteShareResponse), nil
	}
}

type InviteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteUserInvoker) Invoke() (*model.InviteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteUserResponse), nil
	}
}

type InviteWithPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *InviteWithPwdInvoker) Invoke() (*model.InviteWithPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteWithPwdResponse), nil
	}
}

type ListHistoryWebinarsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryWebinarsInvoker) Invoke() (*model.ListHistoryWebinarsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryWebinarsResponse), nil
	}
}

type ListNetworkQualityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNetworkQualityInvoker) Invoke() (*model.ListNetworkQualityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNetworkQualityResponse), nil
	}
}

type ListOngoingWebinarsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOngoingWebinarsInvoker) Invoke() (*model.ListOngoingWebinarsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOngoingWebinarsResponse), nil
	}
}

type ListUpComingWebinarsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUpComingWebinarsInvoker) Invoke() (*model.ListUpComingWebinarsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUpComingWebinarsResponse), nil
	}
}

type LiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *LiveInvoker) Invoke() (*model.LiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveResponse), nil
	}
}

type LockMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *LockMeetingInvoker) Invoke() (*model.LockMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LockMeetingResponse), nil
	}
}

type LockViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *LockViewInvoker) Invoke() (*model.LockViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LockViewResponse), nil
	}
}

type MoveToWaitingRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *MoveToWaitingRoomInvoker) Invoke() (*model.MoveToWaitingRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MoveToWaitingRoomResponse), nil
	}
}

type MuteMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *MuteMeetingInvoker) Invoke() (*model.MuteMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MuteMeetingResponse), nil
	}
}

type MuteParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *MuteParticipantInvoker) Invoke() (*model.MuteParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MuteParticipantResponse), nil
	}
}

type ProlongMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ProlongMeetingInvoker) Invoke() (*model.ProlongMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProlongMeetingResponse), nil
	}
}

type RecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecordInvoker) Invoke() (*model.RecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordResponse), nil
	}
}

type RenameParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenameParticipantInvoker) Invoke() (*model.RenameParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenameParticipantResponse), nil
	}
}

type ResetActivecodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetActivecodeInvoker) Invoke() (*model.ResetActivecodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetActivecodeResponse), nil
	}
}

type ResetAppKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetAppKeyInvoker) Invoke() (*model.ResetAppKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetAppKeyResponse), nil
	}
}

type ResetPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdInvoker) Invoke() (*model.ResetPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPwdResponse), nil
	}
}

type ResetPwdByAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdByAdminInvoker) Invoke() (*model.ResetPwdByAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPwdByAdminResponse), nil
	}
}

type ResetVisionActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetVisionActiveCodeInvoker) Invoke() (*model.ResetVisionActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetVisionActiveCodeResponse), nil
	}
}

type ResumeSimultaneousInterpretationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResumeSimultaneousInterpretationInvoker) Invoke() (*model.ResumeSimultaneousInterpretationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResumeSimultaneousInterpretationResponse), nil
	}
}

type RollcallParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollcallParticipantInvoker) Invoke() (*model.RollcallParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollcallParticipantResponse), nil
	}
}

type SaveLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveLayoutInvoker) Invoke() (*model.SaveLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveLayoutResponse), nil
	}
}

type SearchAttendanceRecordsOfHisMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchAttendanceRecordsOfHisMeetingInvoker) Invoke() (*model.SearchAttendanceRecordsOfHisMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchAttendanceRecordsOfHisMeetingResponse), nil
	}
}

type SearchCorpInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpInvoker) Invoke() (*model.SearchCorpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpResponse), nil
	}
}

type SearchCorpAdminsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpAdminsInvoker) Invoke() (*model.SearchCorpAdminsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpAdminsResponse), nil
	}
}

type SearchCorpDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpDirInvoker) Invoke() (*model.SearchCorpDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpDirResponse), nil
	}
}

type SearchCorpExternalDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpExternalDirInvoker) Invoke() (*model.SearchCorpExternalDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpExternalDirResponse), nil
	}
}

type SearchCorpResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpResourcesInvoker) Invoke() (*model.SearchCorpResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpResourcesResponse), nil
	}
}

type SearchCorpVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCorpVmrInvoker) Invoke() (*model.SearchCorpVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCorpVmrResponse), nil
	}
}

type SearchCtlRecordsOfHisMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchCtlRecordsOfHisMeetingInvoker) Invoke() (*model.SearchCtlRecordsOfHisMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchCtlRecordsOfHisMeetingResponse), nil
	}
}

type SearchDepartmentByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDepartmentByNameInvoker) Invoke() (*model.SearchDepartmentByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDepartmentByNameResponse), nil
	}
}

type SearchDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchDevicesInvoker) Invoke() (*model.SearchDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchDevicesResponse), nil
	}
}

type SearchHisMeetingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchHisMeetingsInvoker) Invoke() (*model.SearchHisMeetingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchHisMeetingsResponse), nil
	}
}

type SearchMaterialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchMaterialsInvoker) Invoke() (*model.SearchMaterialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchMaterialsResponse), nil
	}
}

type SearchMeetingFileListInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchMeetingFileListInvoker) Invoke() (*model.SearchMeetingFileListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchMeetingFileListResponse), nil
	}
}

type SearchMeetingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchMeetingsInvoker) Invoke() (*model.SearchMeetingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchMeetingsResponse), nil
	}
}

type SearchMemberVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchMemberVmrInvoker) Invoke() (*model.SearchMemberVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchMemberVmrResponse), nil
	}
}

type SearchOnlineMeetingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchOnlineMeetingsInvoker) Invoke() (*model.SearchOnlineMeetingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchOnlineMeetingsResponse), nil
	}
}

type SearchProgramsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchProgramsInvoker) Invoke() (*model.SearchProgramsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchProgramsResponse), nil
	}
}

type SearchPublicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchPublicationsInvoker) Invoke() (*model.SearchPublicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchPublicationsResponse), nil
	}
}

type SearchRecordingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchRecordingsInvoker) Invoke() (*model.SearchRecordingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchRecordingsResponse), nil
	}
}

type SearchResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceInvoker) Invoke() (*model.SearchResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceResponse), nil
	}
}

type SearchResourceOpRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchResourceOpRecordInvoker) Invoke() (*model.SearchResourceOpRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchResourceOpRecordResponse), nil
	}
}

type SearchUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchUsersInvoker) Invoke() (*model.SearchUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchUsersResponse), nil
	}
}

type SearchVisionActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchVisionActiveCodeInvoker) Invoke() (*model.SearchVisionActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchVisionActiveCodeResponse), nil
	}
}

type SendSlideVerifyCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendSlideVerifyCodeInvoker) Invoke() (*model.SendSlideVerifyCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendSlideVerifyCodeResponse), nil
	}
}

type SendVeriCodeForChangePwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendVeriCodeForChangePwdInvoker) Invoke() (*model.SendVeriCodeForChangePwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendVeriCodeForChangePwdResponse), nil
	}
}

type SendVeriCodeForUpdateUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendVeriCodeForUpdateUserInfoInvoker) Invoke() (*model.SendVeriCodeForUpdateUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendVeriCodeForUpdateUserInfoResponse), nil
	}
}

type SetAttendeeLanChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetAttendeeLanChannelInvoker) Invoke() (*model.SetAttendeeLanChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetAttendeeLanChannelResponse), nil
	}
}

type SetCohostInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetCohostInvoker) Invoke() (*model.SetCohostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetCohostResponse), nil
	}
}

type SetCustomMultiPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetCustomMultiPictureInvoker) Invoke() (*model.SetCustomMultiPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetCustomMultiPictureResponse), nil
	}
}

type SetHostViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetHostViewInvoker) Invoke() (*model.SetHostViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetHostViewResponse), nil
	}
}

type SetInterpreterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetInterpreterGroupInvoker) Invoke() (*model.SetInterpreterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetInterpreterGroupResponse), nil
	}
}

type SetMultiPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetMultiPictureInvoker) Invoke() (*model.SetMultiPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetMultiPictureResponse), nil
	}
}

type SetParticipantViewInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetParticipantViewInvoker) Invoke() (*model.SetParticipantViewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetParticipantViewResponse), nil
	}
}

type SetProfileImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetProfileImageInvoker) Invoke() (*model.SetProfileImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetProfileImageResponse), nil
	}
}

type SetRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetRoleInvoker) Invoke() (*model.SetRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetRoleResponse), nil
	}
}

type SetSsoConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetSsoConfigInvoker) Invoke() (*model.SetSsoConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetSsoConfigResponse), nil
	}
}

type SetUserProfileImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetUserProfileImageInvoker) Invoke() (*model.SetUserProfileImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetUserProfileImageResponse), nil
	}
}

type SetWebHookConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetWebHookConfigInvoker) Invoke() (*model.SetWebHookConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetWebHookConfigResponse), nil
	}
}

type ShowConfOrgInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfOrgInvoker) Invoke() (*model.ShowConfOrgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfOrgResponse), nil
	}
}

type ShowCorpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCorpInvoker) Invoke() (*model.ShowCorpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCorpResponse), nil
	}
}

type ShowCorpAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCorpAdminInvoker) Invoke() (*model.ShowCorpAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCorpAdminResponse), nil
	}
}

type ShowCorpBasicInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCorpBasicInfoInvoker) Invoke() (*model.ShowCorpBasicInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCorpBasicInfoResponse), nil
	}
}

type ShowCorpResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCorpResourceInvoker) Invoke() (*model.ShowCorpResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCorpResourceResponse), nil
	}
}

type ShowDepartmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDepartmentInvoker) Invoke() (*model.ShowDepartmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDepartmentResponse), nil
	}
}

type ShowDeptAndChildDeptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeptAndChildDeptInvoker) Invoke() (*model.ShowDeptAndChildDeptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeptAndChildDeptResponse), nil
	}
}

type ShowDeviceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceDetailInvoker) Invoke() (*model.ShowDeviceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceDetailResponse), nil
	}
}

type ShowDeviceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceStatusInvoker) Invoke() (*model.ShowDeviceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceStatusResponse), nil
	}
}

type ShowDeviceTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeviceTypesInvoker) Invoke() (*model.ShowDeviceTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeviceTypesResponse), nil
	}
}

type ShowHisMeetingDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHisMeetingDetailInvoker) Invoke() (*model.ShowHisMeetingDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHisMeetingDetailResponse), nil
	}
}

type ShowLayoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLayoutInvoker) Invoke() (*model.ShowLayoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLayoutResponse), nil
	}
}

type ShowMeetingDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMeetingDetailInvoker) Invoke() (*model.ShowMeetingDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMeetingDetailResponse), nil
	}
}

type ShowMeetingFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMeetingFileInvoker) Invoke() (*model.ShowMeetingFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMeetingFileResponse), nil
	}
}

type ShowMeetingFileListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMeetingFileListInvoker) Invoke() (*model.ShowMeetingFileListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMeetingFileListResponse), nil
	}
}

type ShowMyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMyInfoInvoker) Invoke() (*model.ShowMyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMyInfoResponse), nil
	}
}

type ShowOnlineMeetingDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOnlineMeetingDetailInvoker) Invoke() (*model.ShowOnlineMeetingDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOnlineMeetingDetailResponse), nil
	}
}

type ShowOrgResInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrgResInvoker) Invoke() (*model.ShowOrgResResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrgResResponse), nil
	}
}

type ShowProgramInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProgramInvoker) Invoke() (*model.ShowProgramResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProgramResponse), nil
	}
}

type ShowPublicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicationInvoker) Invoke() (*model.ShowPublicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicationResponse), nil
	}
}

type ShowRealTimeInfoOfMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRealTimeInfoOfMeetingInvoker) Invoke() (*model.ShowRealTimeInfoOfMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRealTimeInfoOfMeetingResponse), nil
	}
}

type ShowRecordInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordInfoInvoker) Invoke() (*model.ShowRecordInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordInfoResponse), nil
	}
}

type ShowRecordingDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordingDetailInvoker) Invoke() (*model.ShowRecordingDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordingDetailResponse), nil
	}
}

type ShowRecordingFileDownloadUrlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordingFileDownloadUrlsInvoker) Invoke() (*model.ShowRecordingFileDownloadUrlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordingFileDownloadUrlsResponse), nil
	}
}

type ShowRegionInfoOfMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRegionInfoOfMeetingInvoker) Invoke() (*model.ShowRegionInfoOfMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRegionInfoOfMeetingResponse), nil
	}
}

type ShowRoomSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRoomSettingInvoker) Invoke() (*model.ShowRoomSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRoomSettingResponse), nil
	}
}

type ShowSpResInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpResInvoker) Invoke() (*model.ShowSpResResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpResResponse), nil
	}
}

type ShowSpResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpResourceInvoker) Invoke() (*model.ShowSpResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpResourceResponse), nil
	}
}

type ShowSsoConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSsoConfigInvoker) Invoke() (*model.ShowSsoConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSsoConfigResponse), nil
	}
}

type ShowUserDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserDetailInvoker) Invoke() (*model.ShowUserDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserDetailResponse), nil
	}
}

type ShowWebHookConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebHookConfigInvoker) Invoke() (*model.ShowWebHookConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebHookConfigResponse), nil
	}
}

type ShowWebinarInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWebinarInvoker) Invoke() (*model.ShowWebinarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWebinarResponse), nil
	}
}

type StartMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartMeetingInvoker) Invoke() (*model.StartMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartMeetingResponse), nil
	}
}

type StopMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopMeetingInvoker) Invoke() (*model.StopMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopMeetingResponse), nil
	}
}

type SwitchModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchModeInvoker) Invoke() (*model.SwitchModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchModeResponse), nil
	}
}

type UpdateAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppIdInvoker) Invoke() (*model.UpdateAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppIdResponse), nil
	}
}

type UpdateContactInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateContactInvoker) Invoke() (*model.UpdateContactResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateContactResponse), nil
	}
}

type UpdateCorpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCorpInvoker) Invoke() (*model.UpdateCorpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCorpResponse), nil
	}
}

type UpdateCorpBasicInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCorpBasicInfoInvoker) Invoke() (*model.UpdateCorpBasicInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCorpBasicInfoResponse), nil
	}
}

type UpdateDepartmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDepartmentInvoker) Invoke() (*model.UpdateDepartmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDepartmentResponse), nil
	}
}

type UpdateDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceInvoker) Invoke() (*model.UpdateDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceResponse), nil
	}
}

type UpdateMaterialInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMaterialInvoker) Invoke() (*model.UpdateMaterialResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMaterialResponse), nil
	}
}

type UpdateMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMeetingInvoker) Invoke() (*model.UpdateMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMeetingResponse), nil
	}
}

type UpdateMemberVmrInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberVmrInvoker) Invoke() (*model.UpdateMemberVmrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberVmrResponse), nil
	}
}

type UpdateMyInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMyInfoInvoker) Invoke() (*model.UpdateMyInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMyInfoResponse), nil
	}
}

type UpdateProgramInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProgramInvoker) Invoke() (*model.UpdateProgramResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProgramResponse), nil
	}
}

type UpdatePublicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicationInvoker) Invoke() (*model.UpdatePublicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicationResponse), nil
	}
}

type UpdatePwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePwdInvoker) Invoke() (*model.UpdatePwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePwdResponse), nil
	}
}

type UpdateRecurringMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecurringMeetingInvoker) Invoke() (*model.UpdateRecurringMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecurringMeetingResponse), nil
	}
}

type UpdateRecurringSubMeetingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecurringSubMeetingInvoker) Invoke() (*model.UpdateRecurringSubMeetingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecurringSubMeetingResponse), nil
	}
}

type UpdateResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceInvoker) Invoke() (*model.UpdateResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceResponse), nil
	}
}

type UpdateRoomSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoomSettingInvoker) Invoke() (*model.UpdateRoomSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoomSettingResponse), nil
	}
}

type UpdateStartedConfConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStartedConfConfigInvoker) Invoke() (*model.UpdateStartedConfConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStartedConfConfigResponse), nil
	}
}

type UpdateTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTokenInvoker) Invoke() (*model.UpdateTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTokenResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type UpdateWebHookConfigStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWebHookConfigStatusInvoker) Invoke() (*model.UpdateWebHookConfigStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWebHookConfigStatusResponse), nil
	}
}

type UpdateWebinarInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWebinarInvoker) Invoke() (*model.UpdateWebinarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWebinarResponse), nil
	}
}

type UploadFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFileInvoker) Invoke() (*model.UploadFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFileResponse), nil
	}
}

type SearchQosHistoryMeetingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchQosHistoryMeetingsInvoker) Invoke() (*model.SearchQosHistoryMeetingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchQosHistoryMeetingsResponse), nil
	}
}

type SearchQosOnlineMeetingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchQosOnlineMeetingsInvoker) Invoke() (*model.SearchQosOnlineMeetingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchQosOnlineMeetingsResponse), nil
	}
}

type SearchQosParticipantDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchQosParticipantDetailInvoker) Invoke() (*model.SearchQosParticipantDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchQosParticipantDetailResponse), nil
	}
}

type SearchQosParticipantsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchQosParticipantsInvoker) Invoke() (*model.SearchQosParticipantsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchQosParticipantsResponse), nil
	}
}

type SetQosThresholdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetQosThresholdInvoker) Invoke() (*model.SetQosThresholdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetQosThresholdResponse), nil
	}
}

type ShowQosThresholdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQosThresholdInvoker) Invoke() (*model.ShowQosThresholdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQosThresholdResponse), nil
	}
}

type SearchStatisticConferenceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchStatisticConferenceInfoInvoker) Invoke() (*model.SearchStatisticConferenceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchStatisticConferenceInfoResponse), nil
	}
}

type SearchStatisticConferenceParticipantInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchStatisticConferenceParticipantInvoker) Invoke() (*model.SearchStatisticConferenceParticipantResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchStatisticConferenceParticipantResponse), nil
	}
}

type SearchStatisticResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchStatisticResourceInfoInvoker) Invoke() (*model.SearchStatisticResourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchStatisticResourceInfoResponse), nil
	}
}

type SearchStatisticUserInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchStatisticUserInfoInvoker) Invoke() (*model.SearchStatisticUserInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchStatisticUserInfoResponse), nil
	}
}
