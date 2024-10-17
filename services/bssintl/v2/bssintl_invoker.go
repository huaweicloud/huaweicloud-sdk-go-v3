package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bssintl/v2/model"
)

type AutoRenewalResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AutoRenewalResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AutoRenewalResourcesInvoker) Invoke() (*model.AutoRenewalResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoRenewalResourcesResponse), nil
	}
}

type CancelAutoRenewalResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAutoRenewalResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAutoRenewalResourcesInvoker) Invoke() (*model.CancelAutoRenewalResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAutoRenewalResourcesResponse), nil
	}
}

type CancelCustomerOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelCustomerOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelCustomerOrderInvoker) Invoke() (*model.CancelCustomerOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelCustomerOrderResponse), nil
	}
}

type CancelResourcesSubscriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelResourcesSubscriptionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelResourcesSubscriptionInvoker) Invoke() (*model.CancelResourcesSubscriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelResourcesSubscriptionResponse), nil
	}
}

type ChangeEnterpriseRealnameAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeEnterpriseRealnameAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeEnterpriseRealnameAuthenticationInvoker) Invoke() (*model.ChangeEnterpriseRealnameAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeEnterpriseRealnameAuthenticationResponse), nil
	}
}

type CheckUserIdentityInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckUserIdentityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckUserIdentityInvoker) Invoke() (*model.CheckUserIdentityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckUserIdentityResponse), nil
	}
}

type CreateEnterpriseProjectAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnterpriseProjectAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnterpriseProjectAuthInvoker) Invoke() (*model.CreateEnterpriseProjectAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnterpriseProjectAuthResponse), nil
	}
}

type CreateEnterpriseRealnameAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnterpriseRealnameAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnterpriseRealnameAuthenticationInvoker) Invoke() (*model.CreateEnterpriseRealnameAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnterpriseRealnameAuthenticationResponse), nil
	}
}

type CreatePersonalRealnameAuthInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePersonalRealnameAuthInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePersonalRealnameAuthInvoker) Invoke() (*model.CreatePersonalRealnameAuthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePersonalRealnameAuthResponse), nil
	}
}

type CreateSubCustomerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubCustomerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubCustomerInvoker) Invoke() (*model.CreateSubCustomerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubCustomerResponse), nil
	}
}

type FreezeSubCustomersInvoker struct {
	*invoker.BaseInvoker
}

func (i *FreezeSubCustomersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *FreezeSubCustomersInvoker) Invoke() (*model.FreezeSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.FreezeSubCustomersResponse), nil
	}
}

type ListConversionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConversionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConversionsInvoker) Invoke() (*model.ListConversionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConversionsResponse), nil
	}
}

type ListCostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCostsInvoker) Invoke() (*model.ListCostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCostsResponse), nil
	}
}

type ListCustomerOnDemandResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerOnDemandResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerOnDemandResourcesInvoker) Invoke() (*model.ListCustomerOnDemandResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerOnDemandResourcesResponse), nil
	}
}

type ListCustomerOrdersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerOrdersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerOrdersInvoker) Invoke() (*model.ListCustomerOrdersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerOrdersResponse), nil
	}
}

type ListCustomerselfResourceRecordDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerselfResourceRecordDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerselfResourceRecordDetailsInvoker) Invoke() (*model.ListCustomerselfResourceRecordDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerselfResourceRecordDetailsResponse), nil
	}
}

type ListCustomerselfResourceRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerselfResourceRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerselfResourceRecordsInvoker) Invoke() (*model.ListCustomerselfResourceRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerselfResourceRecordsResponse), nil
	}
}

type ListFreeResourceInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFreeResourceInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFreeResourceInfosInvoker) Invoke() (*model.ListFreeResourceInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFreeResourceInfosResponse), nil
	}
}

type ListFreeResourceUsagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFreeResourceUsagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFreeResourceUsagesInvoker) Invoke() (*model.ListFreeResourceUsagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFreeResourceUsagesResponse), nil
	}
}

type ListFreeResourcesUsageRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFreeResourcesUsageRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFreeResourcesUsageRecordsInvoker) Invoke() (*model.ListFreeResourcesUsageRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFreeResourcesUsageRecordsResponse), nil
	}
}

type ListIndirectPartnersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIndirectPartnersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIndirectPartnersInvoker) Invoke() (*model.ListIndirectPartnersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIndirectPartnersResponse), nil
	}
}

type ListInvoicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInvoicesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInvoicesInvoker) Invoke() (*model.ListInvoicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInvoicesResponse), nil
	}
}

type ListMeasureUnitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMeasureUnitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMeasureUnitsInvoker) Invoke() (*model.ListMeasureUnitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMeasureUnitsResponse), nil
	}
}

type ListMonthlyExpendituresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMonthlyExpendituresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMonthlyExpendituresInvoker) Invoke() (*model.ListMonthlyExpendituresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMonthlyExpendituresResponse), nil
	}
}

type ListOnDemandResourceRatingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOnDemandResourceRatingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOnDemandResourceRatingsInvoker) Invoke() (*model.ListOnDemandResourceRatingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOnDemandResourceRatingsResponse), nil
	}
}

type ListOrderDiscountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrderDiscountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrderDiscountsInvoker) Invoke() (*model.ListOrderDiscountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrderDiscountsResponse), nil
	}
}

type ListPayPerUseCustomerResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPayPerUseCustomerResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPayPerUseCustomerResourcesInvoker) Invoke() (*model.ListPayPerUseCustomerResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPayPerUseCustomerResourcesResponse), nil
	}
}

type ListPostpaidBillSumInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPostpaidBillSumInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPostpaidBillSumInvoker) Invoke() (*model.ListPostpaidBillSumResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPostpaidBillSumResponse), nil
	}
}

type ListRateOnPeriodDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRateOnPeriodDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRateOnPeriodDetailInvoker) Invoke() (*model.ListRateOnPeriodDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRateOnPeriodDetailResponse), nil
	}
}

type ListRenewRateOnPeriodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRenewRateOnPeriodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRenewRateOnPeriodInvoker) Invoke() (*model.ListRenewRateOnPeriodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRenewRateOnPeriodResponse), nil
	}
}

type ListResourceTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceTypesInvoker) Invoke() (*model.ListResourceTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceTypesResponse), nil
	}
}

type ListServiceResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceResourcesInvoker) Invoke() (*model.ListServiceResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceResourcesResponse), nil
	}
}

type ListServiceTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServiceTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServiceTypesInvoker) Invoke() (*model.ListServiceTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServiceTypesResponse), nil
	}
}

type ListSubCustomerBudgetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubCustomerBudgetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubCustomerBudgetInvoker) Invoke() (*model.ListSubCustomerBudgetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomerBudgetResponse), nil
	}
}

type ListSubCustomerCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubCustomerCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubCustomerCouponsInvoker) Invoke() (*model.ListSubCustomerCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomerCouponsResponse), nil
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

type ListUsageTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsageTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsageTypesInvoker) Invoke() (*model.ListUsageTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsageTypesResponse), nil
	}
}

type PayOrdersInvoker struct {
	*invoker.BaseInvoker
}

func (i *PayOrdersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PayOrdersInvoker) Invoke() (*model.PayOrdersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PayOrdersResponse), nil
	}
}

type RenewalResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *RenewalResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RenewalResourcesInvoker) Invoke() (*model.RenewalResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RenewalResourcesResponse), nil
	}
}

type SendVerificationMessageCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendVerificationMessageCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendVerificationMessageCodeInvoker) Invoke() (*model.SendVerificationMessageCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendVerificationMessageCodeResponse), nil
	}
}

type ShowCustomerAccountBalancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerAccountBalancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomerAccountBalancesInvoker) Invoke() (*model.ShowCustomerAccountBalancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerAccountBalancesResponse), nil
	}
}

type ShowCustomerOrderDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerOrderDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomerOrderDetailsInvoker) Invoke() (*model.ShowCustomerOrderDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerOrderDetailsResponse), nil
	}
}

type ShowPartnerConsumptionQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPartnerConsumptionQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPartnerConsumptionQuotaInvoker) Invoke() (*model.ShowPartnerConsumptionQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPartnerConsumptionQuotaResponse), nil
	}
}

type ShowRealnameAuthenticationReviewResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRealnameAuthenticationReviewResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRealnameAuthenticationReviewResultInvoker) Invoke() (*model.ShowRealnameAuthenticationReviewResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRealnameAuthenticationReviewResultResponse), nil
	}
}

type ShowRefundOrderDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRefundOrderDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRefundOrderDetailsInvoker) Invoke() (*model.ShowRefundOrderDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRefundOrderDetailsResponse), nil
	}
}

type ShowSubCustomerBudgetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubCustomerBudgetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSubCustomerBudgetInvoker) Invoke() (*model.ShowSubCustomerBudgetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubCustomerBudgetResponse), nil
	}
}

type UnfreezeSubCustomersInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnfreezeSubCustomersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnfreezeSubCustomersInvoker) Invoke() (*model.UnfreezeSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnfreezeSubCustomersResponse), nil
	}
}

type UpdatePeriodToOnDemandInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePeriodToOnDemandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePeriodToOnDemandInvoker) Invoke() (*model.UpdatePeriodToOnDemandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePeriodToOnDemandResponse), nil
	}
}

type UpdateSubCustomerBudgetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubCustomerBudgetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubCustomerBudgetInvoker) Invoke() (*model.UpdateSubCustomerBudgetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubCustomerBudgetResponse), nil
	}
}
