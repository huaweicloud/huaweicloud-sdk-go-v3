package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/osm/v2/model"
)

type CheckHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckHostsInvoker) Invoke() (*model.CheckHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckHostsResponse), nil
	}
}

type CheckNeedVerifyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckNeedVerifyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckNeedVerifyInvoker) Invoke() (*model.CheckNeedVerifyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckNeedVerifyResponse), nil
	}
}

type CheckVerifyCodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckVerifyCodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckVerifyCodesInvoker) Invoke() (*model.CheckVerifyCodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckVerifyCodesResponse), nil
	}
}

type ConfirmAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ConfirmAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ConfirmAuthorizationsInvoker) Invoke() (*model.ConfirmAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmAuthorizationsResponse), nil
	}
}

type CreateAskQuestionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAskQuestionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAskQuestionInvoker) Invoke() (*model.CreateAskQuestionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAskQuestionResponse), nil
	}
}

type CreateCaseExtendsParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCaseExtendsParamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCaseExtendsParamInvoker) Invoke() (*model.CreateCaseExtendsParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCaseExtendsParamResponse), nil
	}
}

type CreateCaseLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCaseLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCaseLabelsInvoker) Invoke() (*model.CreateCaseLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCaseLabelsResponse), nil
	}
}

type CreateCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCasesInvoker) Invoke() (*model.CreateCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCasesResponse), nil
	}
}

type CreateDiagnoseFeedbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnoseFeedbackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDiagnoseFeedbackInvoker) Invoke() (*model.CreateDiagnoseFeedbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnoseFeedbackResponse), nil
	}
}

type CreateDiagnoseJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDiagnoseJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDiagnoseJobInvoker) Invoke() (*model.CreateDiagnoseJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDiagnoseJobResponse), nil
	}
}

type CreateEvaluateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEvaluateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEvaluateInvoker) Invoke() (*model.CreateEvaluateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEvaluateResponse), nil
	}
}

type CreateFeedbackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFeedbackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFeedbackInvoker) Invoke() (*model.CreateFeedbackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFeedbackResponse), nil
	}
}

type CreateLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLabelsInvoker) Invoke() (*model.CreateLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLabelsResponse), nil
	}
}

type CreateMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMessagesInvoker) Invoke() (*model.CreateMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMessagesResponse), nil
	}
}

type CreatePrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivilegesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePrivilegesInvoker) Invoke() (*model.CreatePrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivilegesResponse), nil
	}
}

type CreateQaAskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQaAskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQaAskInvoker) Invoke() (*model.CreateQaAskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQaAskResponse), nil
	}
}

type CreateQaFeedbacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQaFeedbacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQaFeedbacksInvoker) Invoke() (*model.CreateQaFeedbacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQaFeedbacksResponse), nil
	}
}

type CreateQuestionInSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateQuestionInSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateQuestionInSessionInvoker) Invoke() (*model.CreateQuestionInSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateQuestionInSessionResponse), nil
	}
}

type CreateRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRelationsInvoker) Invoke() (*model.CreateRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRelationsResponse), nil
	}
}

type CreateScoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateScoresInvoker) Invoke() (*model.CreateScoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScoresResponse), nil
	}
}

type CreateSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSessionInvoker) Invoke() (*model.CreateSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSessionResponse), nil
	}
}

type DeleteAccessoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccessoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccessoriesInvoker) Invoke() (*model.DeleteAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccessoriesResponse), nil
	}
}

type DeleteCaseLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCaseLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCaseLabelsInvoker) Invoke() (*model.DeleteCaseLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCaseLabelsResponse), nil
	}
}

type DeleteLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLabelsInvoker) Invoke() (*model.DeleteLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLabelsResponse), nil
	}
}

type DeleteRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRelationInvoker) Invoke() (*model.DeleteRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRelationResponse), nil
	}
}

type DownloadAccessoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadAccessoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadAccessoriesInvoker) Invoke() (*model.DownloadAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadAccessoriesResponse), nil
	}
}

type DownloadCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadCasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadCasesInvoker) Invoke() (*model.DownloadCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadCasesResponse), nil
	}
}

type DownloadImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadImagesInvoker) Invoke() (*model.DownloadImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadImagesResponse), nil
	}
}

type ListAccessoryAccessUrlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccessoryAccessUrlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccessoryAccessUrlsInvoker) Invoke() (*model.ListAccessoryAccessUrlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccessoryAccessUrlsResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type ListAreaCodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAreaCodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAreaCodesInvoker) Invoke() (*model.ListAreaCodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAreaCodesResponse), nil
	}
}

type ListArticlesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListArticlesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListArticlesInvoker) Invoke() (*model.ListArticlesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListArticlesResponse), nil
	}
}

type ListAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuthorizationsInvoker) Invoke() (*model.ListAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuthorizationsResponse), nil
	}
}

type ListCaseCategoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseCategoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseCategoriesInvoker) Invoke() (*model.ListCaseCategoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseCategoriesResponse), nil
	}
}

type ListCaseCcEmailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseCcEmailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseCcEmailsInvoker) Invoke() (*model.ListCaseCcEmailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseCcEmailsResponse), nil
	}
}

type ListCaseCountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseCountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseCountsInvoker) Invoke() (*model.ListCaseCountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseCountsResponse), nil
	}
}

type ListCaseLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseLabelsInvoker) Invoke() (*model.ListCaseLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseLabelsResponse), nil
	}
}

type ListCaseLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseLimitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseLimitsInvoker) Invoke() (*model.ListCaseLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseLimitsResponse), nil
	}
}

type ListCaseOperateLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseOperateLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseOperateLogsInvoker) Invoke() (*model.ListCaseOperateLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseOperateLogsResponse), nil
	}
}

type ListCaseQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseQuotasInvoker) Invoke() (*model.ListCaseQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseQuotasResponse), nil
	}
}

type ListCaseTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCaseTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCaseTemplatesInvoker) Invoke() (*model.ListCaseTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCaseTemplatesResponse), nil
	}
}

type ListCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCasesInvoker) Invoke() (*model.ListCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCasesResponse), nil
	}
}

type ListCustomersRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomersRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomersRegionsInvoker) Invoke() (*model.ListCustomersRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomersRegionsResponse), nil
	}
}

type ListDiagnoseItemsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnoseItemsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnoseItemsInvoker) Invoke() (*model.ListDiagnoseItemsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnoseItemsResponse), nil
	}
}

type ListDiagnoseJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnoseJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnoseJobInvoker) Invoke() (*model.ListDiagnoseJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnoseJobResponse), nil
	}
}

type ListDiagnoseRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnoseRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnoseRecordsInvoker) Invoke() (*model.ListDiagnoseRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnoseRecordsResponse), nil
	}
}

type ListDiagnoseResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDiagnoseResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDiagnoseResourcesInvoker) Invoke() (*model.ListDiagnoseResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDiagnoseResourcesResponse), nil
	}
}

type ListExtendsParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExtendsParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExtendsParamsInvoker) Invoke() (*model.ListExtendsParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExtendsParamsResponse), nil
	}
}

type ListFeedbackOptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeedbackOptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeedbackOptionInvoker) Invoke() (*model.ListFeedbackOptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeedbackOptionResponse), nil
	}
}

type ListHasVerifiedContactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHasVerifiedContactsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHasVerifiedContactsInvoker) Invoke() (*model.ListHasVerifiedContactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHasVerifiedContactsResponse), nil
	}
}

type ListHistoryOperateLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistoryOperateLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistoryOperateLogsInvoker) Invoke() (*model.ListHistoryOperateLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistoryOperateLogsResponse), nil
	}
}

type ListHistorySessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHistorySessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHistorySessionsInvoker) Invoke() (*model.ListHistorySessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHistorySessionsResponse), nil
	}
}

type ListLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLabelsInvoker) Invoke() (*model.ListLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLabelsResponse), nil
	}
}

type ListMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMessagesInvoker) Invoke() (*model.ListMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMessagesResponse), nil
	}
}

type ListMoreInstantMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMoreInstantMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMoreInstantMessagesInvoker) Invoke() (*model.ListMoreInstantMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMoreInstantMessagesResponse), nil
	}
}

type ListNewInstantMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNewInstantMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNewInstantMessagesInvoker) Invoke() (*model.ListNewInstantMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNewInstantMessagesResponse), nil
	}
}

type ListNoticesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNoticesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNoticesInvoker) Invoke() (*model.ListNoticesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNoticesResponse), nil
	}
}

type ListOrderIncidentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrderIncidentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrderIncidentInvoker) Invoke() (*model.ListOrderIncidentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrderIncidentResponse), nil
	}
}

type ListPrivilegesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivilegesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPrivilegesInvoker) Invoke() (*model.ListPrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivilegesResponse), nil
	}
}

type ListProblemTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProblemTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProblemTypesInvoker) Invoke() (*model.ListProblemTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProblemTypesResponse), nil
	}
}

type ListProductCategoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProductCategoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProductCategoriesInvoker) Invoke() (*model.ListProductCategoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductCategoriesResponse), nil
	}
}

type ListRecommendWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecommendWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecommendWordsInvoker) Invoke() (*model.ListRecommendWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecommendWordsResponse), nil
	}
}

type ListRegionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegionsInvoker) Invoke() (*model.ListRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegionsResponse), nil
	}
}

type ListRelationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRelationInvoker) Invoke() (*model.ListRelationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelationResponse), nil
	}
}

type ListSatisfactionDimensionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSatisfactionDimensionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSatisfactionDimensionsInvoker) Invoke() (*model.ListSatisfactionDimensionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSatisfactionDimensionsResponse), nil
	}
}

type ListSeveritiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSeveritiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSeveritiesInvoker) Invoke() (*model.ListSeveritiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSeveritiesResponse), nil
	}
}

type ListSubCustomersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubCustomersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubCustomersInvoker) Invoke() (*model.ListSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomersResponse), nil
	}
}

type ListToolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListToolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListToolsInvoker) Invoke() (*model.ListToolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListToolsResponse), nil
	}
}

type ListTransportHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTransportHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTransportHistoriesInvoker) Invoke() (*model.ListTransportHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTransportHistoriesResponse), nil
	}
}

type ListUnreadNewInstantMessagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUnreadNewInstantMessagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUnreadNewInstantMessagesInvoker) Invoke() (*model.ListUnreadNewInstantMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnreadNewInstantMessagesResponse), nil
	}
}

type RevokeMessageInvoker struct {
	*invoker.BaseInvoker
}

func (i *RevokeMessageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RevokeMessageInvoker) Invoke() (*model.RevokeMessageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RevokeMessageResponse), nil
	}
}

type SendVerifyCodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendVerifyCodesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendVerifyCodesInvoker) Invoke() (*model.SendVerifyCodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendVerifyCodesResponse), nil
	}
}

type ShowAccessoryLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessoryLimitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessoryLimitsInvoker) Invoke() (*model.ShowAccessoryLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessoryLimitsResponse), nil
	}
}

type ShowAssociatedQuestionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssociatedQuestionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssociatedQuestionsInvoker) Invoke() (*model.ShowAssociatedQuestionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssociatedQuestionsResponse), nil
	}
}

type ShowAuthorizationDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthorizationDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthorizationDetailInvoker) Invoke() (*model.ShowAuthorizationDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthorizationDetailResponse), nil
	}
}

type ShowCaseDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCaseDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCaseDetailInvoker) Invoke() (*model.ShowCaseDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCaseDetailResponse), nil
	}
}

type ShowCaseExtendsParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCaseExtendsParamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCaseExtendsParamInvoker) Invoke() (*model.ShowCaseExtendsParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCaseExtendsParamResponse), nil
	}
}

type ShowCaseStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCaseStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCaseStatusInvoker) Invoke() (*model.ShowCaseStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCaseStatusResponse), nil
	}
}

type ShowConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigurationInvoker) Invoke() (*model.ShowConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationResponse), nil
	}
}

type ShowCustomerPrivilegePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerPrivilegePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomerPrivilegePolicyInvoker) Invoke() (*model.ShowCustomerPrivilegePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerPrivilegePolicyResponse), nil
	}
}

type ShowDownloadAccessoryUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDownloadAccessoryUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDownloadAccessoryUrlInvoker) Invoke() (*model.ShowDownloadAccessoryUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDownloadAccessoryUrlResponse), nil
	}
}

type ShowLatestPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestPublishedAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestPublishedAgreementInvoker) Invoke() (*model.ShowLatestPublishedAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestPublishedAgreementResponse), nil
	}
}

type ShowLoginTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoginTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoginTypeInvoker) Invoke() (*model.ShowLoginTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoginTypeResponse), nil
	}
}

type ShowPartnersCasesPrivilegeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartnersCasesPrivilegeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartnersCasesPrivilegeInvoker) Invoke() (*model.ShowPartnersCasesPrivilegeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartnersCasesPrivilegeResponse), nil
	}
}

type ShowPartnersServiceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartnersServiceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartnersServiceInfoInvoker) Invoke() (*model.ShowPartnersServiceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartnersServiceInfoResponse), nil
	}
}

type ShowQaPairDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQaPairDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQaPairDetailInvoker) Invoke() (*model.ShowQaPairDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQaPairDetailResponse), nil
	}
}

type ShowQaPairsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQaPairsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQaPairsInvoker) Invoke() (*model.ShowQaPairsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQaPairsResponse), nil
	}
}

type ShowSignedLatestPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSignedLatestPublishedAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSignedLatestPublishedAgreementInvoker) Invoke() (*model.ShowSignedLatestPublishedAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSignedLatestPublishedAgreementResponse), nil
	}
}

type ShowThemeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowThemeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowThemeInvoker) Invoke() (*model.ShowThemeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowThemeResponse), nil
	}
}

type SignPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *SignPublishedAgreementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SignPublishedAgreementInvoker) Invoke() (*model.SignPublishedAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SignPublishedAgreementResponse), nil
	}
}

type UpdateAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAuthorizationsInvoker) Invoke() (*model.UpdateAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAuthorizationsResponse), nil
	}
}

type UpdateCaseContactInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCaseContactInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCaseContactInfoInvoker) Invoke() (*model.UpdateCaseContactInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCaseContactInfoResponse), nil
	}
}

type UpdateCasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCasesInvoker) Invoke() (*model.UpdateCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCasesResponse), nil
	}
}

type UpdateLabelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLabelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLabelsInvoker) Invoke() (*model.UpdateLabelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLabelsResponse), nil
	}
}

type UpdateNewInstantMessagesReadInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNewInstantMessagesReadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNewInstantMessagesReadInvoker) Invoke() (*model.UpdateNewInstantMessagesReadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNewInstantMessagesReadResponse), nil
	}
}

type UploadJsonAccessoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadJsonAccessoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadJsonAccessoriesInvoker) Invoke() (*model.UploadJsonAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadJsonAccessoriesResponse), nil
	}
}
