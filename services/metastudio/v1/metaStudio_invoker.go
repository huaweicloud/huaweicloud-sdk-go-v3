package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/metastudio/v1/model"
)

type CreateActiveCodeInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowAgencyInvoker) Invoke() (*model.ShowAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgencyResponse), nil
	}
}

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

type BatchExecuteAssetActionInvoker struct {
	*invoker.BaseInvoker
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

type ShowAssetReplicationInfoInvoker struct {
	*invoker.BaseInvoker
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

type CreateLargeFileInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateHotWordsSwitchInvoker) Invoke() (*model.UpdateHotWordsSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHotWordsSwitchResponse), nil
	}
}

type CreateIntentAndQuestionInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateKnowledgeIntentInvoker) Invoke() (*model.UpdateKnowledgeIntentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeIntentResponse), nil
	}
}

type CreateBatchKnowledgeQuestionInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateKnowledgeSkillInvoker) Invoke() (*model.UpdateKnowledgeSkillResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateKnowledgeSkillResponse), nil
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

type CreateProductInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateProductInvoker) Invoke() (*model.UpdateProductResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProductResponse), nil
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

type ListJobOperationLogInvoker struct {
	*invoker.BaseInvoker
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

type CreateWelcomeSpeechInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateWelcomeSpeechSwitchInvoker) Invoke() (*model.UpdateWelcomeSpeechSwitchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWelcomeSpeechSwitchResponse), nil
	}
}
