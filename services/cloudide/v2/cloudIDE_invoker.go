package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudide/v2/model"
)

type AddExtensionEvaluationInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddExtensionEvaluationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddExtensionEvaluationInvoker) Invoke() (*model.AddExtensionEvaluationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddExtensionEvaluationResponse), nil
	}
}

type AddExtensionEvaluationReplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddExtensionEvaluationReplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddExtensionEvaluationReplyInvoker) Invoke() (*model.AddExtensionEvaluationReplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddExtensionEvaluationReplyResponse), nil
	}
}

type AddExtensionStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddExtensionStarInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddExtensionStarInvoker) Invoke() (*model.AddExtensionStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddExtensionStarResponse), nil
	}
}

type CheckMaliciousExtensionEvaluationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckMaliciousExtensionEvaluationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckMaliciousExtensionEvaluationInvoker) Invoke() (*model.CheckMaliciousExtensionEvaluationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckMaliciousExtensionEvaluationResponse), nil
	}
}

type CreateExtensionAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExtensionAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateExtensionAuthorizationInvoker) Invoke() (*model.CreateExtensionAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExtensionAuthorizationResponse), nil
	}
}

type DeleteEvaluationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEvaluationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEvaluationInvoker) Invoke() (*model.DeleteEvaluationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEvaluationResponse), nil
	}
}

type DeleteEvaluationReplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEvaluationReplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEvaluationReplyInvoker) Invoke() (*model.DeleteEvaluationReplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEvaluationReplyResponse), nil
	}
}

type ListExtensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExtensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExtensionsInvoker) Invoke() (*model.ListExtensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExtensionsResponse), nil
	}
}

type ListProjectTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectTemplatesInvoker) Invoke() (*model.ListProjectTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTemplatesResponse), nil
	}
}

type ListPublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublisherInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPublisherInvoker) Invoke() (*model.ListPublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublisherResponse), nil
	}
}

type ListStacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStacksInvoker) Invoke() (*model.ListStacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStacksResponse), nil
	}
}

type PublishExtensionInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishExtensionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PublishExtensionInvoker) Invoke() (*model.PublishExtensionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishExtensionResponse), nil
	}
}

type ShowAccountStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccountStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccountStatusInvoker) Invoke() (*model.ShowAccountStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccountStatusResponse), nil
	}
}

type ShowCategoryListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCategoryListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCategoryListInvoker) Invoke() (*model.ShowCategoryListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCategoryListResponse), nil
	}
}

type ShowExtensionAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtensionAuthorizationInvoker) Invoke() (*model.ShowExtensionAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionAuthorizationResponse), nil
	}
}

type ShowExtensionDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtensionDetailInvoker) Invoke() (*model.ShowExtensionDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionDetailResponse), nil
	}
}

type ShowExtensionEvaluationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionEvaluationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtensionEvaluationInvoker) Invoke() (*model.ShowExtensionEvaluationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionEvaluationResponse), nil
	}
}

type ShowExtensionEvaluationStarInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionEvaluationStarInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtensionEvaluationStarInvoker) Invoke() (*model.ShowExtensionEvaluationStarResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionEvaluationStarResponse), nil
	}
}

type ShowExtensionTestingResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExtensionTestingResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExtensionTestingResultInvoker) Invoke() (*model.ShowExtensionTestingResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExtensionTestingResultResponse), nil
	}
}

type ShowPriceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPriceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPriceInvoker) Invoke() (*model.ShowPriceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPriceResponse), nil
	}
}

type UploadExtensionFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadExtensionFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadExtensionFileInvoker) Invoke() (*model.UploadExtensionFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadExtensionFileResponse), nil
	}
}

type UploadFilePublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadFilePublisherInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadFilePublisherInvoker) Invoke() (*model.UploadFilePublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadFilePublisherResponse), nil
	}
}

type CreateAcceptanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAcceptanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAcceptanceInvoker) Invoke() (*model.CreateAcceptanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAcceptanceResponse), nil
	}
}

type CreateApplyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateApplyInvoker) Invoke() (*model.CreateApplyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplyResponse), nil
	}
}

type CreateEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEventInvoker) Invoke() (*model.CreateEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEventResponse), nil
	}
}

type CreateLoginInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLoginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLoginInvoker) Invoke() (*model.CreateLoginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLoginResponse), nil
	}
}

type CreateRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRequestInvoker) Invoke() (*model.CreateRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRequestResponse), nil
	}
}

type ShowResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResultInvoker) Invoke() (*model.ShowResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResultResponse), nil
	}
}

type StartChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartChatInvoker) Invoke() (*model.StartChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartChatResponse), nil
	}
}

type SyncChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncChatInvoker) Invoke() (*model.SyncChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncChatResponse), nil
	}
}

type SyncGetChatResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *SyncGetChatResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SyncGetChatResultInvoker) Invoke() (*model.SyncGetChatResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SyncGetChatResultResponse), nil
	}
}

type CheckInstanceAccessInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckInstanceAccessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckInstanceAccessInvoker) Invoke() (*model.CheckInstanceAccessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckInstanceAccessResponse), nil
	}
}

type CheckNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckNameInvoker) Invoke() (*model.CheckNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckNameResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateInstanceBy3rdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceBy3rdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceBy3rdInvoker) Invoke() (*model.CreateInstanceBy3rdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceBy3rdResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
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

type ListOrgInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrgInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrgInstancesInvoker) Invoke() (*model.ListOrgInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrgInstancesResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowInstanceStatusInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceStatusInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceStatusInfoInvoker) Invoke() (*model.ShowInstanceStatusInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceStatusInfoResponse), nil
	}
}

type StartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartInstanceInvoker) Invoke() (*model.StartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResponse), nil
	}
}

type StopInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInstanceInvoker) Invoke() (*model.StopInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInstanceResponse), nil
	}
}

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
	}
}

type UpdateInstanceActivityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceActivityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceActivityInvoker) Invoke() (*model.UpdateInstanceActivityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceActivityResponse), nil
	}
}
