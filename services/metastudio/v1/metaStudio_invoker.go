package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type CreateDialogUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDialogUrlInvoker) Invoke() (*model.CreateDialogUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDialogUrlResponse), nil
	}
}

type ShowSmartChatJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSmartChatJobInvoker) Invoke() (*model.ShowSmartChatJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSmartChatJobResponse), nil
	}
}

type StartSmartChatJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartSmartChatJobInvoker) Invoke() (*model.StartSmartChatJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartSmartChatJobResponse), nil
	}
}

type StopSmartChatJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopSmartChatJobInvoker) Invoke() (*model.StopSmartChatJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSmartChatJobResponse), nil
	}
}

type CreateDigitalAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDigitalAssetInvoker) Invoke() (*model.CreateDigitalAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDigitalAssetResponse), nil
	}
}

type DeleteAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAssetInvoker) Invoke() (*model.DeleteAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAssetResponse), nil
	}
}

type ListAssetSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetSummaryInvoker) Invoke() (*model.ListAssetSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetSummaryResponse), nil
	}
}

type ListAssetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetsInvoker) Invoke() (*model.ListAssetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssetsResponse), nil
	}
}

type RestoreAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreAssetInvoker) Invoke() (*model.RestoreAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreAssetResponse), nil
	}
}

type ShowAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetInvoker) Invoke() (*model.ShowAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetResponse), nil
	}
}

type UpdateDigitalAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDigitalAssetInvoker) Invoke() (*model.UpdateDigitalAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDigitalAssetResponse), nil
	}
}

type CreateDigitalHumanBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDigitalHumanBusinessCardInvoker) Invoke() (*model.CreateDigitalHumanBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDigitalHumanBusinessCardResponse), nil
	}
}

type DeleteDigitalHumanBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDigitalHumanBusinessCardInvoker) Invoke() (*model.DeleteDigitalHumanBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDigitalHumanBusinessCardResponse), nil
	}
}

type ListDigitalHumanBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDigitalHumanBusinessCardInvoker) Invoke() (*model.ListDigitalHumanBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDigitalHumanBusinessCardResponse), nil
	}
}

type ShowDigitalHumanBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDigitalHumanBusinessCardInvoker) Invoke() (*model.ShowDigitalHumanBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDigitalHumanBusinessCardResponse), nil
	}
}

type UpdateDigitalHumanBusinessCardInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDigitalHumanBusinessCardInvoker) Invoke() (*model.UpdateDigitalHumanBusinessCardResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDigitalHumanBusinessCardResponse), nil
	}
}

type ListDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDigitalHumanVideoInvoker) Invoke() (*model.ListDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDigitalHumanVideoResponse), nil
	}
}

type Cancel2DDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *Cancel2DDigitalHumanVideoInvoker) Invoke() (*model.Cancel2DDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cancel2DDigitalHumanVideoResponse), nil
	}
}

type Create2DDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *Create2DDigitalHumanVideoInvoker) Invoke() (*model.Create2DDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Create2DDigitalHumanVideoResponse), nil
	}
}

type Show2DDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *Show2DDigitalHumanVideoInvoker) Invoke() (*model.Show2DDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Show2DDigitalHumanVideoResponse), nil
	}
}

type CancelPhotoDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelPhotoDigitalHumanVideoInvoker) Invoke() (*model.CancelPhotoDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelPhotoDigitalHumanVideoResponse), nil
	}
}

type CreatePhotoDetectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePhotoDetectionInvoker) Invoke() (*model.CreatePhotoDetectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePhotoDetectionResponse), nil
	}
}

type CreatePhotoDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePhotoDigitalHumanVideoInvoker) Invoke() (*model.CreatePhotoDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePhotoDigitalHumanVideoResponse), nil
	}
}

type ShowPhotoDetectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPhotoDetectionInvoker) Invoke() (*model.ShowPhotoDetectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPhotoDetectionResponse), nil
	}
}

type ShowPhotoDigitalHumanVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPhotoDigitalHumanVideoInvoker) Invoke() (*model.ShowPhotoDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPhotoDigitalHumanVideoResponse), nil
	}
}

type ConfirmFileUploadInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmFileUploadInvoker) Invoke() (*model.ConfirmFileUploadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmFileUploadResponse), nil
	}
}

type CreateFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFileInvoker) Invoke() (*model.CreateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFileResponse), nil
	}
}

type DeleteFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFileInvoker) Invoke() (*model.DeleteFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFileResponse), nil
	}
}

type CreateOnceCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOnceCodeInvoker) Invoke() (*model.CreateOnceCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOnceCodeResponse), nil
	}
}

type CreatePictureModelingByUrlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePictureModelingByUrlJobInvoker) Invoke() (*model.CreatePictureModelingByUrlJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePictureModelingByUrlJobResponse), nil
	}
}

type CreatePictureModelingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePictureModelingJobInvoker) Invoke() (*model.CreatePictureModelingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePictureModelingJobResponse), nil
	}
}

type ListPictureModelingJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPictureModelingJobsInvoker) Invoke() (*model.ListPictureModelingJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPictureModelingJobsResponse), nil
	}
}

type ShowPictureModelingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPictureModelingJobInvoker) Invoke() (*model.ShowPictureModelingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPictureModelingJobResponse), nil
	}
}

type CreateRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRobotInvoker) Invoke() (*model.CreateRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRobotResponse), nil
	}
}

type DeleteRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRobotInvoker) Invoke() (*model.DeleteRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRobotResponse), nil
	}
}

type ListRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRobotInvoker) Invoke() (*model.ListRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRobotResponse), nil
	}
}

type ShowRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRobotInvoker) Invoke() (*model.ShowRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRobotResponse), nil
	}
}

type UpdateRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRobotInvoker) Invoke() (*model.UpdateRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRobotResponse), nil
	}
}

type CreateSmartChatRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSmartChatRoomInvoker) Invoke() (*model.CreateSmartChatRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSmartChatRoomResponse), nil
	}
}

type DeleteSmartChatRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSmartChatRoomInvoker) Invoke() (*model.DeleteSmartChatRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSmartChatRoomResponse), nil
	}
}

type ListSmartChatRoomsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartChatRoomsInvoker) Invoke() (*model.ListSmartChatRoomsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartChatRoomsResponse), nil
	}
}

type ShowSmartChatRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSmartChatRoomInvoker) Invoke() (*model.ShowSmartChatRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSmartChatRoomResponse), nil
	}
}

type UpdateSmartChatRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSmartChatRoomInvoker) Invoke() (*model.UpdateSmartChatRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSmartChatRoomResponse), nil
	}
}

type ExecuteSmartLiveCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSmartLiveCommandInvoker) Invoke() (*model.ExecuteSmartLiveCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSmartLiveCommandResponse), nil
	}
}

type ListSmartLiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartLiveInvoker) Invoke() (*model.ListSmartLiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveResponse), nil
	}
}

type ListSmartLiveJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartLiveJobsInvoker) Invoke() (*model.ListSmartLiveJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveJobsResponse), nil
	}
}

type LiveEventReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *LiveEventReportInvoker) Invoke() (*model.LiveEventReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LiveEventReportResponse), nil
	}
}

type ShowSmartLiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSmartLiveInvoker) Invoke() (*model.ShowSmartLiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSmartLiveResponse), nil
	}
}

type StartSmartLiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartSmartLiveInvoker) Invoke() (*model.StartSmartLiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartSmartLiveResponse), nil
	}
}

type StopSmartLiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopSmartLiveInvoker) Invoke() (*model.StopSmartLiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSmartLiveResponse), nil
	}
}

type CheckTextLanguageInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckTextLanguageInvoker) Invoke() (*model.CheckTextLanguageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckTextLanguageResponse), nil
	}
}

type CreateInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInteractionRuleGroupInvoker) Invoke() (*model.CreateInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInteractionRuleGroupResponse), nil
	}
}

type CreateSmartLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSmartLiveRoomInvoker) Invoke() (*model.CreateSmartLiveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSmartLiveRoomResponse), nil
	}
}

type DeleteInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInteractionRuleGroupInvoker) Invoke() (*model.DeleteInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInteractionRuleGroupResponse), nil
	}
}

type DeleteSmartLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSmartLiveRoomInvoker) Invoke() (*model.DeleteSmartLiveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSmartLiveRoomResponse), nil
	}
}

type ListInteractionRuleGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInteractionRuleGroupsInvoker) Invoke() (*model.ListInteractionRuleGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInteractionRuleGroupsResponse), nil
	}
}

type ListSmartLiveRoomsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartLiveRoomsInvoker) Invoke() (*model.ListSmartLiveRoomsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveRoomsResponse), nil
	}
}

type ShowSmartLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSmartLiveRoomInvoker) Invoke() (*model.ShowSmartLiveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSmartLiveRoomResponse), nil
	}
}

type UpdateInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInteractionRuleGroupInvoker) Invoke() (*model.UpdateInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInteractionRuleGroupResponse), nil
	}
}

type UpdateSmartLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSmartLiveRoomInvoker) Invoke() (*model.UpdateSmartLiveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSmartLiveRoomResponse), nil
	}
}

type ListStylesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStylesInvoker) Invoke() (*model.ListStylesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStylesResponse), nil
	}
}

type CommitVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitVoiceTrainingJobInvoker) Invoke() (*model.CommitVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommitVoiceTrainingJobResponse), nil
	}
}

type ConfirmTrainingSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmTrainingSegmentInvoker) Invoke() (*model.ConfirmTrainingSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmTrainingSegmentResponse), nil
	}
}

type CreateTrainingAdvanceJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainingAdvanceJobInvoker) Invoke() (*model.CreateTrainingAdvanceJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainingAdvanceJobResponse), nil
	}
}

type CreateTrainingBasicJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainingBasicJobInvoker) Invoke() (*model.CreateTrainingBasicJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainingBasicJobResponse), nil
	}
}

type CreateTrainingMiddleJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrainingMiddleJobInvoker) Invoke() (*model.CreateTrainingMiddleJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrainingMiddleJobResponse), nil
	}
}

type DeleteVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVoiceTrainingJobInvoker) Invoke() (*model.DeleteVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVoiceTrainingJobResponse), nil
	}
}

type ListVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVoiceTrainingJobInvoker) Invoke() (*model.ListVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVoiceTrainingJobResponse), nil
	}
}

type ShowJobAuditResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobAuditResultInvoker) Invoke() (*model.ShowJobAuditResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobAuditResultResponse), nil
	}
}

type ShowJobUploadingAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobUploadingAddressInvoker) Invoke() (*model.ShowJobUploadingAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobUploadingAddressResponse), nil
	}
}

type ShowTrainingSegmentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingSegmentInfoInvoker) Invoke() (*model.ShowTrainingSegmentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrainingSegmentInfoResponse), nil
	}
}

type ShowVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVoiceTrainingJobInvoker) Invoke() (*model.ShowVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVoiceTrainingJobResponse), nil
	}
}

type Create2dModelTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *Create2dModelTrainingJobInvoker) Invoke() (*model.Create2dModelTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Create2dModelTrainingJobResponse), nil
	}
}

type Delete2dModelTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *Delete2dModelTrainingJobInvoker) Invoke() (*model.Delete2dModelTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Delete2dModelTrainingJobResponse), nil
	}
}

type Execute2dModelTrainingCommandByUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *Execute2dModelTrainingCommandByUserInvoker) Invoke() (*model.Execute2dModelTrainingCommandByUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Execute2dModelTrainingCommandByUserResponse), nil
	}
}

type List2dModelTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *List2dModelTrainingJobInvoker) Invoke() (*model.List2dModelTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.List2dModelTrainingJobResponse), nil
	}
}

type Show2dModelTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *Show2dModelTrainingJobInvoker) Invoke() (*model.Show2dModelTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Show2dModelTrainingJobResponse), nil
	}
}

type Update2dModelTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *Update2dModelTrainingJobInvoker) Invoke() (*model.Update2dModelTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.Update2dModelTrainingJobResponse), nil
	}
}

type CreateFacialAnimationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFacialAnimationsInvoker) Invoke() (*model.CreateFacialAnimationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFacialAnimationsResponse), nil
	}
}

type CreateTtsaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTtsaInvoker) Invoke() (*model.CreateTtsaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTtsaResponse), nil
	}
}

type ListFacialAnimationsDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFacialAnimationsDataInvoker) Invoke() (*model.ListFacialAnimationsDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFacialAnimationsDataResponse), nil
	}
}

type ListTtsaDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTtsaDataInvoker) Invoke() (*model.ListTtsaDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTtsaDataResponse), nil
	}
}

type ListTtsaJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTtsaJobsInvoker) Invoke() (*model.ListTtsaJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTtsaJobsResponse), nil
	}
}

type CreateTtsAuditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTtsAuditionInvoker) Invoke() (*model.CreateTtsAuditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTtsAuditionResponse), nil
	}
}

type ShowTtsAuditionFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTtsAuditionFileInvoker) Invoke() (*model.ShowTtsAuditionFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTtsAuditionFileResponse), nil
	}
}

type CreateVideoMotionCaptureJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVideoMotionCaptureJobInvoker) Invoke() (*model.CreateVideoMotionCaptureJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVideoMotionCaptureJobResponse), nil
	}
}

type ExecuteVideoMotionCaptureCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteVideoMotionCaptureCommandInvoker) Invoke() (*model.ExecuteVideoMotionCaptureCommandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteVideoMotionCaptureCommandResponse), nil
	}
}

type ListVideoMotionCaptureJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVideoMotionCaptureJobsInvoker) Invoke() (*model.ListVideoMotionCaptureJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVideoMotionCaptureJobsResponse), nil
	}
}

type ShowVideoMotionCaptureJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVideoMotionCaptureJobInvoker) Invoke() (*model.ShowVideoMotionCaptureJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVideoMotionCaptureJobResponse), nil
	}
}

type StopVideoMotionCaptureJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopVideoMotionCaptureJobInvoker) Invoke() (*model.StopVideoMotionCaptureJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopVideoMotionCaptureJobResponse), nil
	}
}

type CopyVideoScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyVideoScriptsInvoker) Invoke() (*model.CopyVideoScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyVideoScriptsResponse), nil
	}
}

type CreateVideoScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVideoScriptsInvoker) Invoke() (*model.CreateVideoScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVideoScriptsResponse), nil
	}
}

type DeleteVideoScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVideoScriptInvoker) Invoke() (*model.DeleteVideoScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVideoScriptResponse), nil
	}
}

type ListVideoScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVideoScriptsInvoker) Invoke() (*model.ListVideoScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVideoScriptsResponse), nil
	}
}

type ShowVideoScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVideoScriptInvoker) Invoke() (*model.ShowVideoScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVideoScriptResponse), nil
	}
}

type UpdateVideoScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVideoScriptInvoker) Invoke() (*model.UpdateVideoScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVideoScriptResponse), nil
	}
}
