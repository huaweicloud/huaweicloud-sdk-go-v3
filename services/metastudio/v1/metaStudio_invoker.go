package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type CreateActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateActiveCodeInvoker) Invoke() (*model.CreateActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateActiveCodeResponse), nil
	}
}

type DeleteActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteActiveCodeInvoker) Invoke() (*model.DeleteActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteActiveCodeResponse), nil
	}
}

type ListActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListActiveCodeInvoker) Invoke() (*model.ListActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListActiveCodeResponse), nil
	}
}

type ResetActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetActiveCodeInvoker) Invoke() (*model.ResetActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetActiveCodeResponse), nil
	}
}

type ShowActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowActiveCodeInvoker) Invoke() (*model.ShowActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowActiveCodeResponse), nil
	}
}

type UpdateActiveCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateActiveCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateActiveCodeInvoker) Invoke() (*model.UpdateActiveCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateActiveCodeResponse), nil
	}
}

type CreateAgencyWithRoleTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyWithRoleTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAgencyWithRoleTypeInvoker) Invoke() (*model.CreateAgencyWithRoleTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyWithRoleTypeResponse), nil
	}
}

type DeleteAgencyWithRoleTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAgencyWithRoleTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAgencyWithRoleTypeInvoker) Invoke() (*model.DeleteAgencyWithRoleTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAgencyWithRoleTypeResponse), nil
	}
}

type ShowAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgencyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAgencyInvoker) Invoke() (*model.ShowAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type DownloadResultFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadResultFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadResultFileInvoker) Invoke() (*model.DownloadResultFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadResultFileResponse), nil
	}
}

type DownloadTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadTemplateInvoker) Invoke() (*model.DownloadTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadTemplateResponse), nil
	}
}

type ExportResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportResourceInvoker) Invoke() (*model.ExportResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportResourceResponse), nil
	}
}

type ImportResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportResourceInvoker) Invoke() (*model.ImportResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportResourceResponse), nil
	}
}

type SearchTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchTaskInvoker) Invoke() (*model.SearchTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchTaskResponse), nil
	}
}

type ShowTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskInvoker) Invoke() (*model.ShowTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskResponse), nil
	}
}

type CreateDialogUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDialogUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDialogUrlInvoker) Invoke() (*model.CreateDialogUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDialogUrlResponse), nil
	}
}

type ListSmartChatJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartChatJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmartChatJobInvoker) Invoke() (*model.ListSmartChatJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartChatJobResponse), nil
	}
}

type ShowSmartChatJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSmartChatJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *StartSmartChatJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *StopSmartChatJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopSmartChatJobInvoker) Invoke() (*model.StopSmartChatJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSmartChatJobResponse), nil
	}
}

type CreateDialogReportConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDialogReportConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDialogReportConfigInvoker) Invoke() (*model.CreateDialogReportConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDialogReportConfigResponse), nil
	}
}

type DeleteDialogReportConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDialogReportConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDialogReportConfigInvoker) Invoke() (*model.DeleteDialogReportConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDialogReportConfigResponse), nil
	}
}

type ShowDialogReportConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDialogReportConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDialogReportConfigInvoker) Invoke() (*model.ShowDialogReportConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDialogReportConfigResponse), nil
	}
}

type UpdateDialogReportConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDialogReportConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDialogReportConfigInvoker) Invoke() (*model.UpdateDialogReportConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDialogReportConfigResponse), nil
	}
}

type BatchExecuteAssetActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchExecuteAssetActionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchExecuteAssetActionInvoker) Invoke() (*model.BatchExecuteAssetActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchExecuteAssetActionResponse), nil
	}
}

type CreateAssetByReplicationInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAssetByReplicationInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAssetByReplicationInfoInvoker) Invoke() (*model.CreateAssetByReplicationInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAssetByReplicationInfoResponse), nil
	}
}

type CreateDigitalAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDigitalAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ListAssetSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssetSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAssetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RestoreAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetInvoker) Invoke() (*model.ShowAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetResponse), nil
	}
}

type ShowAssetReplicationInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssetReplicationInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssetReplicationInfoInvoker) Invoke() (*model.ShowAssetReplicationInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssetReplicationInfoResponse), nil
	}
}

type UpdateDigitalAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDigitalAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateDigitalHumanBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteDigitalHumanBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDigitalHumanBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDigitalHumanBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateDigitalHumanBusinessCardInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Cancel2DDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Create2DDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Show2DDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CancelPhotoDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreatePhotoDetectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreatePhotoDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPhotoDetectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPhotoDigitalHumanVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPhotoDigitalHumanVideoInvoker) Invoke() (*model.ShowPhotoDigitalHumanVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPhotoDigitalHumanVideoResponse), nil
	}
}

type CreateDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDocumentInvoker) Invoke() (*model.CreateDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDocumentResponse), nil
	}
}

type DeleteDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDocumentInvoker) Invoke() (*model.DeleteDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDocumentResponse), nil
	}
}

type DownloadDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadDocumentInvoker) Invoke() (*model.DownloadDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadDocumentResponse), nil
	}
}

type ListDocumentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDocumentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDocumentInfoInvoker) Invoke() (*model.ListDocumentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDocumentInfoResponse), nil
	}
}

type ShowDocumentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDocumentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDocumentInfoInvoker) Invoke() (*model.ShowDocumentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDocumentInfoResponse), nil
	}
}

type UpdateDocumentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDocumentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDocumentInvoker) Invoke() (*model.UpdateDocumentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDocumentResponse), nil
	}
}

type ListDocumentSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDocumentSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDocumentSegmentInvoker) Invoke() (*model.ListDocumentSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDocumentSegmentResponse), nil
	}
}

type PreviewDocumentSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *PreviewDocumentSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PreviewDocumentSegmentInvoker) Invoke() (*model.PreviewDocumentSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PreviewDocumentSegmentResponse), nil
	}
}

type StartDocumentSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDocumentSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDocumentSegmentInvoker) Invoke() (*model.StartDocumentSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDocumentSegmentResponse), nil
	}
}

type UpdateDocumentSegmentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDocumentSegmentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDocumentSegmentInfoInvoker) Invoke() (*model.UpdateDocumentSegmentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDocumentSegmentInfoResponse), nil
	}
}

type UpdateDocumentSegmentParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDocumentSegmentParamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDocumentSegmentParamInvoker) Invoke() (*model.UpdateDocumentSegmentParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDocumentSegmentParamResponse), nil
	}
}

type ConfirmFileUploadInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmFileUploadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFileInvoker) Invoke() (*model.CreateFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFileResponse), nil
	}
}

type CreateLargeFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLargeFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLargeFileInvoker) Invoke() (*model.CreateLargeFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLargeFileResponse), nil
	}
}

type DeleteFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFileInvoker) Invoke() (*model.DeleteFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFileResponse), nil
	}
}

type CreateHotQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHotQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHotQuestionInvoker) Invoke() (*model.CreateHotQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHotQuestionResponse), nil
	}
}

type DeleteHotQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHotQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHotQuestionInvoker) Invoke() (*model.DeleteHotQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHotQuestionResponse), nil
	}
}

type ListHotQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHotQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHotQuestionInvoker) Invoke() (*model.ListHotQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHotQuestionResponse), nil
	}
}

type ShowHotQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHotQuestionInvoker) Invoke() (*model.ShowHotQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotQuestionResponse), nil
	}
}

type UpdateHotQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHotQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHotQuestionInvoker) Invoke() (*model.UpdateHotQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHotQuestionResponse), nil
	}
}

type CreateHotWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHotWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHotWordsInvoker) Invoke() (*model.CreateHotWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHotWordsResponse), nil
	}
}

type DeleteHotWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHotWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHotWordsInvoker) Invoke() (*model.DeleteHotWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHotWordsResponse), nil
	}
}

type ListHotWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHotWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHotWordsInvoker) Invoke() (*model.ListHotWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHotWordsResponse), nil
	}
}

type ShowHotWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHotWordsInvoker) Invoke() (*model.ShowHotWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotWordsResponse), nil
	}
}

type ShowHotWordsSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHotWordsSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHotWordsSwitchInvoker) Invoke() (*model.ShowHotWordsSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHotWordsSwitchResponse), nil
	}
}

type UpdateHotWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHotWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHotWordsInvoker) Invoke() (*model.UpdateHotWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHotWordsResponse), nil
	}
}

type UpdateHotWordsSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHotWordsSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHotWordsSwitchInvoker) Invoke() (*model.UpdateHotWordsSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHotWordsSwitchResponse), nil
	}
}

type CreateInstructionLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstructionLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstructionLibraryInvoker) Invoke() (*model.CreateInstructionLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstructionLibraryResponse), nil
	}
}

type DeleteInstructionLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstructionLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstructionLibraryInvoker) Invoke() (*model.DeleteInstructionLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstructionLibraryResponse), nil
	}
}

type ListInstructionLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstructionLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstructionLibraryInvoker) Invoke() (*model.ListInstructionLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstructionLibraryResponse), nil
	}
}

type ShowInstructionLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstructionLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstructionLibraryInvoker) Invoke() (*model.ShowInstructionLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstructionLibraryResponse), nil
	}
}

type UpdateInstructionLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstructionLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstructionLibraryInvoker) Invoke() (*model.UpdateInstructionLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstructionLibraryResponse), nil
	}
}

type CreateInstructionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstructionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstructionInvoker) Invoke() (*model.CreateInstructionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstructionResponse), nil
	}
}

type DeleteInstructionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstructionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstructionInvoker) Invoke() (*model.DeleteInstructionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstructionResponse), nil
	}
}

type ListInstructionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstructionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstructionInvoker) Invoke() (*model.ListInstructionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstructionResponse), nil
	}
}

type ShowInstructionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstructionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstructionInvoker) Invoke() (*model.ShowInstructionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstructionResponse), nil
	}
}

type UpdateInstructionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstructionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstructionInvoker) Invoke() (*model.UpdateInstructionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstructionResponse), nil
	}
}

type CreateInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInteractionRuleGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInteractionRuleGroupInvoker) Invoke() (*model.CreateInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInteractionRuleGroupResponse), nil
	}
}

type DeleteInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInteractionRuleGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInteractionRuleGroupInvoker) Invoke() (*model.DeleteInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInteractionRuleGroupResponse), nil
	}
}

type ListInteractionRuleGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInteractionRuleGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInteractionRuleGroupsInvoker) Invoke() (*model.ListInteractionRuleGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInteractionRuleGroupsResponse), nil
	}
}

type UpdateInteractionRuleGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInteractionRuleGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInteractionRuleGroupInvoker) Invoke() (*model.UpdateInteractionRuleGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInteractionRuleGroupResponse), nil
	}
}

type CreateInteractiveChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInteractiveChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInteractiveChatInvoker) Invoke() (*model.CreateInteractiveChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInteractiveChatResponse), nil
	}
}

type CreateIntentAndQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIntentAndQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateIntentAndQuestionInvoker) Invoke() (*model.CreateIntentAndQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIntentAndQuestionResponse), nil
	}
}

type CreateKnowledgeIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKnowledgeIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKnowledgeIntentInvoker) Invoke() (*model.CreateKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKnowledgeIntentResponse), nil
	}
}

type DeleteKnowledgeIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKnowledgeIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKnowledgeIntentInvoker) Invoke() (*model.DeleteKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKnowledgeIntentResponse), nil
	}
}

type ListKnowledgeIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKnowledgeIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKnowledgeIntentInvoker) Invoke() (*model.ListKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKnowledgeIntentResponse), nil
	}
}

type ShowKnowledgeIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKnowledgeIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKnowledgeIntentInvoker) Invoke() (*model.ShowKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKnowledgeIntentResponse), nil
	}
}

type UpdateKnowledgeIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKnowledgeIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKnowledgeIntentInvoker) Invoke() (*model.UpdateKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeIntentResponse), nil
	}
}

type CheckRecallKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRecallKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRecallKnowledgeLibraryInvoker) Invoke() (*model.CheckRecallKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRecallKnowledgeLibraryResponse), nil
	}
}

type CreateKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKnowledgeLibraryInvoker) Invoke() (*model.CreateKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKnowledgeLibraryResponse), nil
	}
}

type DeleteKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKnowledgeLibraryInvoker) Invoke() (*model.DeleteKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKnowledgeLibraryResponse), nil
	}
}

type ListKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKnowledgeLibraryInvoker) Invoke() (*model.ListKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKnowledgeLibraryResponse), nil
	}
}

type ShowKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKnowledgeLibraryInvoker) Invoke() (*model.ShowKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKnowledgeLibraryResponse), nil
	}
}

type UpdateKnowledgeLibraryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKnowledgeLibraryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKnowledgeLibraryInvoker) Invoke() (*model.UpdateKnowledgeLibraryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeLibraryResponse), nil
	}
}

type CreateBatchKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBatchKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBatchKnowledgeQuestionInvoker) Invoke() (*model.CreateBatchKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBatchKnowledgeQuestionResponse), nil
	}
}

type CreateKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKnowledgeQuestionInvoker) Invoke() (*model.CreateKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKnowledgeQuestionResponse), nil
	}
}

type DeleteKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKnowledgeQuestionInvoker) Invoke() (*model.DeleteKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKnowledgeQuestionResponse), nil
	}
}

type ListKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKnowledgeQuestionInvoker) Invoke() (*model.ListKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKnowledgeQuestionResponse), nil
	}
}

type ShowKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKnowledgeQuestionInvoker) Invoke() (*model.ShowKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKnowledgeQuestionResponse), nil
	}
}

type UpdateBatchKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBatchKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBatchKnowledgeQuestionInvoker) Invoke() (*model.UpdateBatchKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBatchKnowledgeQuestionResponse), nil
	}
}

type UpdateKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKnowledgeQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKnowledgeQuestionInvoker) Invoke() (*model.UpdateKnowledgeQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeQuestionResponse), nil
	}
}

type CreateKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateKnowledgeSkillInvoker) Invoke() (*model.CreateKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKnowledgeSkillResponse), nil
	}
}

type DeleteKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKnowledgeSkillInvoker) Invoke() (*model.DeleteKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKnowledgeSkillResponse), nil
	}
}

type ExportKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportKnowledgeSkillInvoker) Invoke() (*model.ExportKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportKnowledgeSkillResponse), nil
	}
}

type ListKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKnowledgeSkillInvoker) Invoke() (*model.ListKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKnowledgeSkillResponse), nil
	}
}

type ShowKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKnowledgeSkillInvoker) Invoke() (*model.ShowKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKnowledgeSkillResponse), nil
	}
}

type UpdateKnowledgeSkillInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateKnowledgeSkillInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateKnowledgeSkillInvoker) Invoke() (*model.UpdateKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeSkillResponse), nil
	}
}

type CreateLivePlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLivePlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLivePlatformInvoker) Invoke() (*model.CreateLivePlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLivePlatformResponse), nil
	}
}

type DeleteLivePlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLivePlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLivePlatformInvoker) Invoke() (*model.DeleteLivePlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLivePlatformResponse), nil
	}
}

type ListLivePlatformProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLivePlatformProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLivePlatformProductsInvoker) Invoke() (*model.ListLivePlatformProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLivePlatformProductsResponse), nil
	}
}

type ListLivePlatformsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLivePlatformsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLivePlatformsInvoker) Invoke() (*model.ListLivePlatformsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLivePlatformsResponse), nil
	}
}

type ShowLivePlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLivePlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLivePlatformInvoker) Invoke() (*model.ShowLivePlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLivePlatformResponse), nil
	}
}

type UpdateLivePlatformInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLivePlatformInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLivePlatformInvoker) Invoke() (*model.UpdateLivePlatformResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLivePlatformResponse), nil
	}
}

type CreateLlmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLlmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLlmConfigInvoker) Invoke() (*model.CreateLlmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLlmConfigResponse), nil
	}
}

type DeleteLlmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLlmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLlmConfigInvoker) Invoke() (*model.DeleteLlmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLlmConfigResponse), nil
	}
}

type ListLlmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLlmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLlmConfigInvoker) Invoke() (*model.ListLlmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLlmConfigResponse), nil
	}
}

type ShowLlmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLlmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLlmConfigInvoker) Invoke() (*model.ShowLlmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLlmConfigResponse), nil
	}
}

type UpdateLlmConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLlmConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLlmConfigInvoker) Invoke() (*model.UpdateLlmConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLlmConfigResponse), nil
	}
}

type CreateMcpServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMcpServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMcpServerInvoker) Invoke() (*model.CreateMcpServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMcpServerResponse), nil
	}
}

type DeleteMcpServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMcpServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMcpServerInvoker) Invoke() (*model.DeleteMcpServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMcpServerResponse), nil
	}
}

type ListMcpServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMcpServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMcpServerInvoker) Invoke() (*model.ListMcpServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMcpServerResponse), nil
	}
}

type ShowMcpServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMcpServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMcpServerInvoker) Invoke() (*model.ShowMcpServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMcpServerResponse), nil
	}
}

type UpdateMcpServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMcpServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMcpServerInvoker) Invoke() (*model.UpdateMcpServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMcpServerResponse), nil
	}
}

type CreateOnceCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOnceCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOnceCodeInvoker) Invoke() (*model.CreateOnceCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOnceCodeResponse), nil
	}
}

type CreateMetaStudioOrdersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMetaStudioOrdersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMetaStudioOrdersInvoker) Invoke() (*model.CreateMetaStudioOrdersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMetaStudioOrdersResponse), nil
	}
}

type BatchDeletePacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePacifyWordsInvoker) Invoke() (*model.BatchDeletePacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePacifyWordsResponse), nil
	}
}

type CreatePacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePacifyWordsInvoker) Invoke() (*model.CreatePacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePacifyWordsResponse), nil
	}
}

type DeletePacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePacifyWordsInvoker) Invoke() (*model.DeletePacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePacifyWordsResponse), nil
	}
}

type ListPacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPacifyWordsInvoker) Invoke() (*model.ListPacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPacifyWordsResponse), nil
	}
}

type ShowPacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPacifyWordsInvoker) Invoke() (*model.ShowPacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPacifyWordsResponse), nil
	}
}

type ShowPacifyWordsIntentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPacifyWordsIntentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPacifyWordsIntentInvoker) Invoke() (*model.ShowPacifyWordsIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPacifyWordsIntentResponse), nil
	}
}

type ShowPacifyWordsSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPacifyWordsSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPacifyWordsSwitchInvoker) Invoke() (*model.ShowPacifyWordsSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPacifyWordsSwitchResponse), nil
	}
}

type ShowPacifyWordsTriggerTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPacifyWordsTriggerTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPacifyWordsTriggerTimeInvoker) Invoke() (*model.ShowPacifyWordsTriggerTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPacifyWordsTriggerTimeResponse), nil
	}
}

type UpdatePacifyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePacifyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePacifyWordsInvoker) Invoke() (*model.UpdatePacifyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePacifyWordsResponse), nil
	}
}

type UpdatePacifyWordsSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePacifyWordsSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePacifyWordsSwitchInvoker) Invoke() (*model.UpdatePacifyWordsSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePacifyWordsSwitchResponse), nil
	}
}

type UpdatePacifyWordsTriggerTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePacifyWordsTriggerTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePacifyWordsTriggerTimeInvoker) Invoke() (*model.UpdatePacifyWordsTriggerTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePacifyWordsTriggerTimeResponse), nil
	}
}

type CreatePictureModelingByUrlJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePictureModelingByUrlJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreatePictureModelingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPictureModelingJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowPictureModelingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPictureModelingJobInvoker) Invoke() (*model.ShowPictureModelingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPictureModelingJobResponse), nil
	}
}

type CreatePluginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePluginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePluginConfigInvoker) Invoke() (*model.CreatePluginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePluginConfigResponse), nil
	}
}

type DeletePluginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePluginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePluginConfigInvoker) Invoke() (*model.DeletePluginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePluginConfigResponse), nil
	}
}

type ListPluginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPluginConfigInvoker) Invoke() (*model.ListPluginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginConfigResponse), nil
	}
}

type ShowPluginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPluginConfigInvoker) Invoke() (*model.ShowPluginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginConfigResponse), nil
	}
}

type ShowPluginConfigDefaultInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginConfigDefaultInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPluginConfigDefaultInfoInvoker) Invoke() (*model.ShowPluginConfigDefaultInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginConfigDefaultInfoResponse), nil
	}
}

type UpdatePluginConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePluginConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePluginConfigInvoker) Invoke() (*model.UpdatePluginConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginConfigResponse), nil
	}
}

type CreateProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProductInvoker) Invoke() (*model.CreateProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProductResponse), nil
	}
}

type DeleteProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProductInvoker) Invoke() (*model.DeleteProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProductResponse), nil
	}
}

type ListProductsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductsInvoker) Invoke() (*model.ListProductsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductsResponse), nil
	}
}

type SetProductAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetProductAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetProductAssetInvoker) Invoke() (*model.SetProductAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetProductAssetResponse), nil
	}
}

type ShowProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProductInvoker) Invoke() (*model.ShowProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductResponse), nil
	}
}

type UpdateProductInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProductInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProductInvoker) Invoke() (*model.UpdateProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProductResponse), nil
	}
}

type CreateQuestionAnswerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQuestionAnswerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQuestionAnswerInvoker) Invoke() (*model.CreateQuestionAnswerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQuestionAnswerResponse), nil
	}
}

type DeleteQuestionAnswerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteQuestionAnswerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteQuestionAnswerInvoker) Invoke() (*model.DeleteQuestionAnswerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteQuestionAnswerResponse), nil
	}
}

type ListQuestionAnswerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuestionAnswerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuestionAnswerInvoker) Invoke() (*model.ListQuestionAnswerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuestionAnswerResponse), nil
	}
}

type ShowQuestionAnswerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuestionAnswerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuestionAnswerInvoker) Invoke() (*model.ShowQuestionAnswerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuestionAnswerResponse), nil
	}
}

type UpdateQuestionAnswerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateQuestionAnswerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateQuestionAnswerInvoker) Invoke() (*model.UpdateQuestionAnswerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateQuestionAnswerResponse), nil
	}
}

type CreateRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRobotInvoker) Invoke() (*model.UpdateRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRobotResponse), nil
	}
}

type ValidateRobotInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateRobotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateRobotInvoker) Invoke() (*model.ValidateRobotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateRobotResponse), nil
	}
}

type CreateRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRoleInvoker) Invoke() (*model.CreateRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRoleResponse), nil
	}
}

type DeleteRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRoleInvoker) Invoke() (*model.DeleteRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRoleResponse), nil
	}
}

type ListRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRoleInvoker) Invoke() (*model.ListRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRoleResponse), nil
	}
}

type ShowRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRoleInvoker) Invoke() (*model.ShowRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRoleResponse), nil
	}
}

type UpdateRoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRoleInvoker) Invoke() (*model.UpdateRoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoleResponse), nil
	}
}

type CreateSmartChatRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSmartChatRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteSmartChatRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSmartChatRoomsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowSmartChatRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateSmartChatRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSmartChatRoomInvoker) Invoke() (*model.UpdateSmartChatRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSmartChatRoomResponse), nil
	}
}

type BatchConfirmLiveCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchConfirmLiveCommandsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchConfirmLiveCommandsInvoker) Invoke() (*model.BatchConfirmLiveCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchConfirmLiveCommandsResponse), nil
	}
}

type ExecuteSmartLiveCommandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSmartLiveCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSmartLiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSmartLiveJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmartLiveJobsInvoker) Invoke() (*model.ListSmartLiveJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveJobsResponse), nil
	}
}

type ListSmartLiveRuleCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartLiveRuleCommandsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmartLiveRuleCommandsInvoker) Invoke() (*model.ListSmartLiveRuleCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveRuleCommandsResponse), nil
	}
}

type ListSmartLiveScriptCommandsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSmartLiveScriptCommandsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSmartLiveScriptCommandsInvoker) Invoke() (*model.ListSmartLiveScriptCommandsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSmartLiveScriptCommandsResponse), nil
	}
}

type LiveEventReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *LiveEventReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowSmartLiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *StartSmartLiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *StopSmartLiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopSmartLiveInvoker) Invoke() (*model.StopSmartLiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopSmartLiveResponse), nil
	}
}

type ConfirmSmarLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmSmarLiveRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmSmarLiveRoomInvoker) Invoke() (*model.ConfirmSmarLiveRoomResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmSmarLiveRoomResponse), nil
	}
}

type CreateSmartLiveRoomInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSmartLiveRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteSmartLiveRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSmartLiveRoomsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowSmartLiveRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateSmartLiveRoomInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListStylesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStylesInvoker) Invoke() (*model.ListStylesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStylesResponse), nil
	}
}

type CreateSubtitleFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubtitleFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubtitleFileInvoker) Invoke() (*model.CreateSubtitleFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubtitleFileResponse), nil
	}
}

type ShowSubtitleFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubtitleFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSubtitleFileInvoker) Invoke() (*model.ShowSubtitleFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubtitleFileResponse), nil
	}
}

type CountTenantResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountTenantResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountTenantResourcesInvoker) Invoke() (*model.CountTenantResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountTenantResourcesResponse), nil
	}
}

type ListTenantResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTenantResourcesInvoker) Invoke() (*model.ListTenantResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantResourcesResponse), nil
	}
}

type ShowResourceUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceUsageInvoker) Invoke() (*model.ShowResourceUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceUsageResponse), nil
	}
}

type SignAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SignAgreementInvoker) Invoke() (*model.SignAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignAgreementResponse), nil
	}
}

type SignSpecialAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignSpecialAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SignSpecialAgreementInvoker) Invoke() (*model.SignSpecialAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignSpecialAgreementResponse), nil
	}
}

type CommitVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CommitVoiceTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ConfirmTrainingSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateTrainingAdvanceJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateTrainingBasicJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateTrainingMiddleJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteVoiceTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVoiceTrainingJobInvoker) Invoke() (*model.DeleteVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVoiceTrainingJobResponse), nil
	}
}

type ListJobOperationLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobOperationLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobOperationLogInvoker) Invoke() (*model.ListJobOperationLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobOperationLogResponse), nil
	}
}

type ListVoiceTrainingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVoiceTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVoiceTrainingJobInvoker) Invoke() (*model.ListVoiceTrainingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVoiceTrainingJobResponse), nil
	}
}

type SetJobBatchNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetJobBatchNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetJobBatchNameInvoker) Invoke() (*model.SetJobBatchNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetJobBatchNameResponse), nil
	}
}

type ShowJobAuditResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobAuditResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowJobUploadingAddressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobUploadingAddressInvoker) Invoke() (*model.ShowJobUploadingAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobUploadingAddressResponse), nil
	}
}

type ShowTenantDurationCfgInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTenantDurationCfgInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTenantDurationCfgInvoker) Invoke() (*model.ShowTenantDurationCfgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTenantDurationCfgResponse), nil
	}
}

type ShowTrainingSegmentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrainingSegmentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowVoiceTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Create2dModelTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Delete2dModelTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Execute2dModelTrainingCommandByUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *List2dModelTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Show2dModelTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *Update2dModelTrainingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateFacialAnimationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateTtsaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListFacialAnimationsDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListTtsaDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListTtsaJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTtsaJobsInvoker) Invoke() (*model.ListTtsaJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTtsaJobsResponse), nil
	}
}

type CheckVoiceAssetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckVoiceAssetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckVoiceAssetInvoker) Invoke() (*model.CheckVoiceAssetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckVoiceAssetResponse), nil
	}
}

type ShowTtsJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTtsJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTtsJobInvoker) Invoke() (*model.ShowTtsJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTtsJobResponse), nil
	}
}

type ShowTtsPhoneticSymbolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTtsPhoneticSymbolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTtsPhoneticSymbolInvoker) Invoke() (*model.ShowTtsPhoneticSymbolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTtsPhoneticSymbolResponse), nil
	}
}

type CreateAsyncTtsJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAsyncTtsJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAsyncTtsJobInvoker) Invoke() (*model.CreateAsyncTtsJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAsyncTtsJobResponse), nil
	}
}

type CreateTtsAuditionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTtsAuditionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTtsAuditionInvoker) Invoke() (*model.CreateTtsAuditionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTtsAuditionResponse), nil
	}
}

type ShowAsyncTtsJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAsyncTtsJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAsyncTtsJobInvoker) Invoke() (*model.ShowAsyncTtsJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAsyncTtsJobResponse), nil
	}
}

type ShowTtsAuditionFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTtsAuditionFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTtsAuditionFileInvoker) Invoke() (*model.ShowTtsAuditionFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTtsAuditionFileResponse), nil
	}
}

type CreateTtscVocabularyConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTtscVocabularyConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTtscVocabularyConfigsInvoker) Invoke() (*model.CreateTtscVocabularyConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTtscVocabularyConfigsResponse), nil
	}
}

type CreateTtscVocabularyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTtscVocabularyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTtscVocabularyGroupsInvoker) Invoke() (*model.CreateTtscVocabularyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTtscVocabularyGroupsResponse), nil
	}
}

type DeleteTtscVocabularyConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTtscVocabularyConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTtscVocabularyConfigsInvoker) Invoke() (*model.DeleteTtscVocabularyConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTtscVocabularyConfigsResponse), nil
	}
}

type DeleteTtscVocabularyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTtscVocabularyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTtscVocabularyGroupsInvoker) Invoke() (*model.DeleteTtscVocabularyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTtscVocabularyGroupsResponse), nil
	}
}

type ListTtscVocabularyConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTtscVocabularyConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTtscVocabularyConfigsInvoker) Invoke() (*model.ListTtscVocabularyConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTtscVocabularyConfigsResponse), nil
	}
}

type ListTtscVocabularyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTtscVocabularyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTtscVocabularyGroupsInvoker) Invoke() (*model.ListTtscVocabularyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTtscVocabularyGroupsResponse), nil
	}
}

type SaveTtscTenantConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveTtscTenantConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveTtscTenantConfigsInvoker) Invoke() (*model.SaveTtscTenantConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveTtscTenantConfigsResponse), nil
	}
}

type SaveTtscVocabularyConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SaveTtscVocabularyConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SaveTtscVocabularyConfigsInvoker) Invoke() (*model.SaveTtscVocabularyConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SaveTtscVocabularyConfigsResponse), nil
	}
}

type SetTtscGroupAssetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetTtscGroupAssetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetTtscGroupAssetsInvoker) Invoke() (*model.SetTtscGroupAssetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetTtscGroupAssetsResponse), nil
	}
}

type ShowVocabularySwitchConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVocabularySwitchConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVocabularySwitchConfigsInvoker) Invoke() (*model.ShowVocabularySwitchConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVocabularySwitchConfigsResponse), nil
	}
}

type UpdateTtscVocabularyGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTtscVocabularyGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTtscVocabularyGroupsInvoker) Invoke() (*model.UpdateTtscVocabularyGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTtscVocabularyGroupsResponse), nil
	}
}

type CreateVideoMotionCaptureJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVideoMotionCaptureJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExecuteVideoMotionCaptureCommandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListVideoMotionCaptureJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowVideoMotionCaptureJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *StopVideoMotionCaptureJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CopyVideoScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateVideoScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteVideoScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListVideoScriptsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowVideoScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateVideoScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVideoScriptInvoker) Invoke() (*model.UpdateVideoScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVideoScriptResponse), nil
	}
}

type CreateWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWelcomeSpeechInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWelcomeSpeechInvoker) Invoke() (*model.CreateWelcomeSpeechResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWelcomeSpeechResponse), nil
	}
}

type DeleteWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWelcomeSpeechInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWelcomeSpeechInvoker) Invoke() (*model.DeleteWelcomeSpeechResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWelcomeSpeechResponse), nil
	}
}

type ListWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWelcomeSpeechInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWelcomeSpeechInvoker) Invoke() (*model.ListWelcomeSpeechResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWelcomeSpeechResponse), nil
	}
}

type ShowWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWelcomeSpeechInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWelcomeSpeechInvoker) Invoke() (*model.ShowWelcomeSpeechResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWelcomeSpeechResponse), nil
	}
}

type ShowWelcomeSpeechSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWelcomeSpeechSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWelcomeSpeechSwitchInvoker) Invoke() (*model.ShowWelcomeSpeechSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWelcomeSpeechSwitchResponse), nil
	}
}

type UpdateWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWelcomeSpeechInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWelcomeSpeechInvoker) Invoke() (*model.UpdateWelcomeSpeechResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWelcomeSpeechResponse), nil
	}
}

type UpdateWelcomeSpeechSwitchInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWelcomeSpeechSwitchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWelcomeSpeechSwitchInvoker) Invoke() (*model.UpdateWelcomeSpeechSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWelcomeSpeechSwitchResponse), nil
	}
}
