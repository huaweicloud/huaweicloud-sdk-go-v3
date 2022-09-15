package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/osm/v2/model"
)

type CheckHostsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ConfirmAuthorizationsInvoker) Invoke() (*model.ConfirmAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ConfirmAuthorizationsResponse), nil
	}
}

type CreateCaseExtendsParamInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateCasesInvoker) Invoke() (*model.CreateCasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCasesResponse), nil
	}
}

type CreateLabelsInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreatePrivilegesInvoker) Invoke() (*model.CreatePrivilegesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivilegesResponse), nil
	}
}

type CreateRelationsInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateScoresInvoker) Invoke() (*model.CreateScoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScoresResponse), nil
	}
}

type DeleteAccessoriesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListAreaCodesInvoker) Invoke() (*model.ListAreaCodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAreaCodesResponse), nil
	}
}

type ListAuthorizationsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListCustomersRegionsInvoker) Invoke() (*model.ListCustomersRegionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomersRegionsResponse), nil
	}
}

type ListExtendsParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExtendsParamsInvoker) Invoke() (*model.ListExtendsParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExtendsParamsResponse), nil
	}
}

type ListHasVerifiedContactsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListNewInstantMessagesInvoker) Invoke() (*model.ListNewInstantMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNewInstantMessagesResponse), nil
	}
}

type ListPrivilegesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListProductCategoriesInvoker) Invoke() (*model.ListProductCategoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProductCategoriesResponse), nil
	}
}

type ListRegionsInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListSubCustomersInvoker) Invoke() (*model.ListSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomersResponse), nil
	}
}

type ListTransportHistoriesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListUnreadNewInstantMessagesInvoker) Invoke() (*model.ListUnreadNewInstantMessagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUnreadNewInstantMessagesResponse), nil
	}
}

type SendVerifyCodesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowAccessoryLimitsInvoker) Invoke() (*model.ShowAccessoryLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessoryLimitsResponse), nil
	}
}

type ShowAuthorizationDetailInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowCaseStatusInvoker) Invoke() (*model.ShowCaseStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCaseStatusResponse), nil
	}
}

type ShowCustomerPrivilegePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerPrivilegePolicyInvoker) Invoke() (*model.ShowCustomerPrivilegePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerPrivilegePolicyResponse), nil
	}
}

type ShowLatestPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestPublishedAgreementInvoker) Invoke() (*model.ShowLatestPublishedAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestPublishedAgreementResponse), nil
	}
}

type ShowPartnersCasesPrivilegeInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowPartnersServiceInfoInvoker) Invoke() (*model.ShowPartnersServiceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartnersServiceInfoResponse), nil
	}
}

type ShowSignedLatestPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSignedLatestPublishedAgreementInvoker) Invoke() (*model.ShowSignedLatestPublishedAgreementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSignedLatestPublishedAgreementResponse), nil
	}
}

type SignPublishedAgreementInvoker struct {
	*invoker.BaseInvoker
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

func (i *UploadJsonAccessoriesInvoker) Invoke() (*model.UploadJsonAccessoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadJsonAccessoriesResponse), nil
	}
}
