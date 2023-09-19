package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

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
