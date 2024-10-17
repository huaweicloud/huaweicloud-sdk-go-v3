package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"
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

type ClaimEnterpriseMultiAccountCouponInvoker struct {
	*invoker.BaseInvoker
}

func (i *ClaimEnterpriseMultiAccountCouponInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClaimEnterpriseMultiAccountCouponInvoker) Invoke() (*model.ClaimEnterpriseMultiAccountCouponResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClaimEnterpriseMultiAccountCouponResponse), nil
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

type CreatePartnerCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePartnerCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePartnerCouponsInvoker) Invoke() (*model.CreatePartnerCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePartnerCouponsResponse), nil
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

type CreateSubEnterpriseAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubEnterpriseAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSubEnterpriseAccountInvoker) Invoke() (*model.CreateSubEnterpriseAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubEnterpriseAccountResponse), nil
	}
}

type ListCitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCitiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCitiesInvoker) Invoke() (*model.ListCitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCitiesResponse), nil
	}
}

type ListConsumeSubCustomersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConsumeSubCustomersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConsumeSubCustomersInvoker) Invoke() (*model.ListConsumeSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConsumeSubCustomersResponse), nil
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

type ListCountiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCountiesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCountiesInvoker) Invoke() (*model.ListCountiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCountiesResponse), nil
	}
}

type ListCouponQuotasRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCouponQuotasRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCouponQuotasRecordsInvoker) Invoke() (*model.ListCouponQuotasRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCouponQuotasRecordsResponse), nil
	}
}

type ListCustomerAccountChangeRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerAccountChangeRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerAccountChangeRecordsInvoker) Invoke() (*model.ListCustomerAccountChangeRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerAccountChangeRecordsResponse), nil
	}
}

type ListCustomerBillsFeeRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerBillsFeeRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerBillsFeeRecordsInvoker) Invoke() (*model.ListCustomerBillsFeeRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerBillsFeeRecordsResponse), nil
	}
}

type ListCustomerBillsMonthlyBreakDownInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomerBillsMonthlyBreakDownInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomerBillsMonthlyBreakDownInvoker) Invoke() (*model.ListCustomerBillsMonthlyBreakDownResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomerBillsMonthlyBreakDownResponse), nil
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

type ListCustomersBalancesDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomersBalancesDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomersBalancesDetailInvoker) Invoke() (*model.ListCustomersBalancesDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomersBalancesDetailResponse), nil
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

type ListEnterpriseMultiAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseMultiAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseMultiAccountInvoker) Invoke() (*model.ListEnterpriseMultiAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseMultiAccountResponse), nil
	}
}

type ListEnterpriseOrganizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseOrganizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseOrganizationsInvoker) Invoke() (*model.ListEnterpriseOrganizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseOrganizationsResponse), nil
	}
}

type ListEnterpriseSubCustomersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseSubCustomersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseSubCustomersInvoker) Invoke() (*model.ListEnterpriseSubCustomersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseSubCustomersResponse), nil
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

type ListIncentiveDiscountPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIncentiveDiscountPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIncentiveDiscountPoliciesInvoker) Invoke() (*model.ListIncentiveDiscountPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIncentiveDiscountPoliciesResponse), nil
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

type ListIssuedCouponQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssuedCouponQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssuedCouponQuotasInvoker) Invoke() (*model.ListIssuedCouponQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssuedCouponQuotasResponse), nil
	}
}

type ListIssuedPartnerCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIssuedPartnerCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIssuedPartnerCouponsInvoker) Invoke() (*model.ListIssuedPartnerCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIssuedPartnerCouponsResponse), nil
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

type ListMultiAccountRetrieveCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMultiAccountRetrieveCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMultiAccountRetrieveCouponsInvoker) Invoke() (*model.ListMultiAccountRetrieveCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMultiAccountRetrieveCouponsResponse), nil
	}
}

type ListMultiAccountTransferCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMultiAccountTransferCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMultiAccountTransferCouponsInvoker) Invoke() (*model.ListMultiAccountTransferCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMultiAccountTransferCouponsResponse), nil
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

type ListOrderCouponsByOrderIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrderCouponsByOrderIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrderCouponsByOrderIdInvoker) Invoke() (*model.ListOrderCouponsByOrderIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrderCouponsByOrderIdResponse), nil
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

type ListPartnerAccountChangeRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartnerAccountChangeRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartnerAccountChangeRecordsInvoker) Invoke() (*model.ListPartnerAccountChangeRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartnerAccountChangeRecordsResponse), nil
	}
}

type ListPartnerAdjustRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartnerAdjustRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartnerAdjustRecordsInvoker) Invoke() (*model.ListPartnerAdjustRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartnerAdjustRecordsResponse), nil
	}
}

type ListPartnerBalancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartnerBalancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartnerBalancesInvoker) Invoke() (*model.ListPartnerBalancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartnerBalancesResponse), nil
	}
}

type ListPartnerCouponsRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPartnerCouponsRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPartnerCouponsRecordInvoker) Invoke() (*model.ListPartnerCouponsRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPartnerCouponsRecordResponse), nil
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

type ListProvincesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProvincesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProvincesInvoker) Invoke() (*model.ListProvincesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProvincesResponse), nil
	}
}

type ListQuotaCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaCouponsInvoker) Invoke() (*model.ListQuotaCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaCouponsResponse), nil
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

type ListResourceUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceUsageInvoker) Invoke() (*model.ListResourceUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceUsageResponse), nil
	}
}

type ListResourceUsageSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceUsageSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceUsageSummaryInvoker) Invoke() (*model.ListResourceUsageSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceUsageSummaryResponse), nil
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

type ListStoredValueCardsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoredValueCardsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStoredValueCardsInvoker) Invoke() (*model.ListStoredValueCardsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoredValueCardsResponse), nil
	}
}

type ListSubCustomerBillDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubCustomerBillDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubCustomerBillDetailInvoker) Invoke() (*model.ListSubCustomerBillDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomerBillDetailResponse), nil
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

type ListSubCustomerNewTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubCustomerNewTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubCustomerNewTagInvoker) Invoke() (*model.ListSubCustomerNewTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubCustomerNewTagResponse), nil
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

type ListSubcustomerMonthlyBillsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubcustomerMonthlyBillsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSubcustomerMonthlyBillsInvoker) Invoke() (*model.ListSubcustomerMonthlyBillsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubcustomerMonthlyBillsResponse), nil
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

type ReclaimCouponQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimCouponQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimCouponQuotasInvoker) Invoke() (*model.ReclaimCouponQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimCouponQuotasResponse), nil
	}
}

type ReclaimEnterpriseMultiAccountCouponInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimEnterpriseMultiAccountCouponInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimEnterpriseMultiAccountCouponInvoker) Invoke() (*model.ReclaimEnterpriseMultiAccountCouponResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimEnterpriseMultiAccountCouponResponse), nil
	}
}

type ReclaimIndirectPartnerAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimIndirectPartnerAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimIndirectPartnerAccountInvoker) Invoke() (*model.ReclaimIndirectPartnerAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimIndirectPartnerAccountResponse), nil
	}
}

type ReclaimPartnerCouponsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimPartnerCouponsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimPartnerCouponsInvoker) Invoke() (*model.ReclaimPartnerCouponsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimPartnerCouponsResponse), nil
	}
}

type ReclaimSubEnterpriseAmountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimSubEnterpriseAmountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimSubEnterpriseAmountInvoker) Invoke() (*model.ReclaimSubEnterpriseAmountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimSubEnterpriseAmountResponse), nil
	}
}

type ReclaimToPartnerAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReclaimToPartnerAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReclaimToPartnerAccountInvoker) Invoke() (*model.ReclaimToPartnerAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReclaimToPartnerAccountResponse), nil
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

type SendSmsVerificationCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendSmsVerificationCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendSmsVerificationCodeInvoker) Invoke() (*model.SendSmsVerificationCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendSmsVerificationCodeResponse), nil
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

type ShowCustomerMonthlySumInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCustomerMonthlySumInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCustomerMonthlySumInvoker) Invoke() (*model.ShowCustomerMonthlySumResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCustomerMonthlySumResponse), nil
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

type ShowMultiAccountTransferAmountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMultiAccountTransferAmountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMultiAccountTransferAmountInvoker) Invoke() (*model.ShowMultiAccountTransferAmountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMultiAccountTransferAmountResponse), nil
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

type UpdateCouponQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCouponQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCouponQuotasInvoker) Invoke() (*model.UpdateCouponQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCouponQuotasResponse), nil
	}
}

type UpdateCustomerAccountAmountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCustomerAccountAmountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateCustomerAccountAmountInvoker) Invoke() (*model.UpdateCustomerAccountAmountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCustomerAccountAmountResponse), nil
	}
}

type UpdateIndirectPartnerAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIndirectPartnerAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIndirectPartnerAccountInvoker) Invoke() (*model.UpdateIndirectPartnerAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIndirectPartnerAccountResponse), nil
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

type UpdateSubEnterpriseAmountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubEnterpriseAmountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSubEnterpriseAmountInvoker) Invoke() (*model.UpdateSubEnterpriseAmountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubEnterpriseAmountResponse), nil
	}
}
