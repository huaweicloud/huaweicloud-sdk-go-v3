package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/config/v1/model"
)

type CreateAggregationAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAggregationAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAggregationAuthorizationInvoker) Invoke() (*model.CreateAggregationAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAggregationAuthorizationResponse), nil
	}
}

type CreateConfigurationAggregatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigurationAggregatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigurationAggregatorInvoker) Invoke() (*model.CreateConfigurationAggregatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigurationAggregatorResponse), nil
	}
}

type DeleteAggregationAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAggregationAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAggregationAuthorizationInvoker) Invoke() (*model.DeleteAggregationAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAggregationAuthorizationResponse), nil
	}
}

type DeleteConfigurationAggregatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigurationAggregatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConfigurationAggregatorInvoker) Invoke() (*model.DeleteConfigurationAggregatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigurationAggregatorResponse), nil
	}
}

type DeletePendingAggregationRequestInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePendingAggregationRequestInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePendingAggregationRequestInvoker) Invoke() (*model.DeletePendingAggregationRequestResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePendingAggregationRequestResponse), nil
	}
}

type ListAggregateComplianceByPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAggregateComplianceByPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAggregateComplianceByPolicyAssignmentInvoker) Invoke() (*model.ListAggregateComplianceByPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAggregateComplianceByPolicyAssignmentResponse), nil
	}
}

type ListAggregateDiscoveredResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAggregateDiscoveredResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAggregateDiscoveredResourcesInvoker) Invoke() (*model.ListAggregateDiscoveredResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAggregateDiscoveredResourcesResponse), nil
	}
}

type ListAggregationAuthorizationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAggregationAuthorizationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAggregationAuthorizationsInvoker) Invoke() (*model.ListAggregationAuthorizationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAggregationAuthorizationsResponse), nil
	}
}

type ListConfigurationAggregatorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigurationAggregatorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigurationAggregatorsInvoker) Invoke() (*model.ListConfigurationAggregatorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigurationAggregatorsResponse), nil
	}
}

type ListPendingAggregationRequestsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPendingAggregationRequestsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPendingAggregationRequestsInvoker) Invoke() (*model.ListPendingAggregationRequestsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPendingAggregationRequestsResponse), nil
	}
}

type RunAggregateResourceQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAggregateResourceQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunAggregateResourceQueryInvoker) Invoke() (*model.RunAggregateResourceQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAggregateResourceQueryResponse), nil
	}
}

type ShowAggregateComplianceDetailsByPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregateComplianceDetailsByPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregateComplianceDetailsByPolicyAssignmentInvoker) Invoke() (*model.ShowAggregateComplianceDetailsByPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregateComplianceDetailsByPolicyAssignmentResponse), nil
	}
}

type ShowAggregateDiscoveredResourceCountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregateDiscoveredResourceCountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregateDiscoveredResourceCountsInvoker) Invoke() (*model.ShowAggregateDiscoveredResourceCountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregateDiscoveredResourceCountsResponse), nil
	}
}

type ShowAggregatePolicyAssignmentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregatePolicyAssignmentDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregatePolicyAssignmentDetailInvoker) Invoke() (*model.ShowAggregatePolicyAssignmentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregatePolicyAssignmentDetailResponse), nil
	}
}

type ShowAggregatePolicyStateComplianceSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregatePolicyStateComplianceSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregatePolicyStateComplianceSummaryInvoker) Invoke() (*model.ShowAggregatePolicyStateComplianceSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregatePolicyStateComplianceSummaryResponse), nil
	}
}

type ShowAggregateResourceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAggregateResourceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAggregateResourceConfigInvoker) Invoke() (*model.ShowAggregateResourceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAggregateResourceConfigResponse), nil
	}
}

type ShowConfigurationAggregatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationAggregatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigurationAggregatorInvoker) Invoke() (*model.ShowConfigurationAggregatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationAggregatorResponse), nil
	}
}

type ShowConfigurationAggregatorSourcesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigurationAggregatorSourcesStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigurationAggregatorSourcesStatusInvoker) Invoke() (*model.ShowConfigurationAggregatorSourcesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigurationAggregatorSourcesStatusResponse), nil
	}
}

type UpdateConfigurationAggregatorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigurationAggregatorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConfigurationAggregatorInvoker) Invoke() (*model.UpdateConfigurationAggregatorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigurationAggregatorResponse), nil
	}
}

type CollectConformancePackComplianceSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectConformancePackComplianceSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectConformancePackComplianceSummaryInvoker) Invoke() (*model.CollectConformancePackComplianceSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectConformancePackComplianceSummaryResponse), nil
	}
}

type CreateConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConformancePackInvoker) Invoke() (*model.CreateConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConformancePackResponse), nil
	}
}

type CreateOrganizationConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrganizationConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrganizationConformancePackInvoker) Invoke() (*model.CreateOrganizationConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrganizationConformancePackResponse), nil
	}
}

type DeleteConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConformancePackInvoker) Invoke() (*model.DeleteConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConformancePackResponse), nil
	}
}

type DeleteOrganizationConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOrganizationConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOrganizationConformancePackInvoker) Invoke() (*model.DeleteOrganizationConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOrganizationConformancePackResponse), nil
	}
}

type ListBuiltInConformancePackTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBuiltInConformancePackTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBuiltInConformancePackTemplatesInvoker) Invoke() (*model.ListBuiltInConformancePackTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBuiltInConformancePackTemplatesResponse), nil
	}
}

type ListConformancePackComplianceByPackIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConformancePackComplianceByPackIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConformancePackComplianceByPackIdInvoker) Invoke() (*model.ListConformancePackComplianceByPackIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConformancePackComplianceByPackIdResponse), nil
	}
}

type ListConformancePackComplianceDetailsByPackIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConformancePackComplianceDetailsByPackIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConformancePackComplianceDetailsByPackIdInvoker) Invoke() (*model.ListConformancePackComplianceDetailsByPackIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConformancePackComplianceDetailsByPackIdResponse), nil
	}
}

type ListConformancePackComplianceScoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConformancePackComplianceScoresInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConformancePackComplianceScoresInvoker) Invoke() (*model.ListConformancePackComplianceScoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConformancePackComplianceScoresResponse), nil
	}
}

type ListConformancePacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConformancePacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConformancePacksInvoker) Invoke() (*model.ListConformancePacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConformancePacksResponse), nil
	}
}

type ListOrganizationConformancePackStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationConformancePackStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationConformancePackStatusesInvoker) Invoke() (*model.ListOrganizationConformancePackStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationConformancePackStatusesResponse), nil
	}
}

type ListOrganizationConformancePacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationConformancePacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationConformancePacksInvoker) Invoke() (*model.ListOrganizationConformancePacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationConformancePacksResponse), nil
	}
}

type ShowBuiltInConformancePackTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuiltInConformancePackTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuiltInConformancePackTemplateInvoker) Invoke() (*model.ShowBuiltInConformancePackTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuiltInConformancePackTemplateResponse), nil
	}
}

type ShowConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConformancePackInvoker) Invoke() (*model.ShowConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConformancePackResponse), nil
	}
}

type ShowOrganizationConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationConformancePackInvoker) Invoke() (*model.ShowOrganizationConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationConformancePackResponse), nil
	}
}

type ShowOrganizationConformancePackDetailedStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationConformancePackDetailedStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationConformancePackDetailedStatusesInvoker) Invoke() (*model.ShowOrganizationConformancePackDetailedStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationConformancePackDetailedStatusesResponse), nil
	}
}

type UpdateConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConformancePackInvoker) Invoke() (*model.UpdateConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConformancePackResponse), nil
	}
}

type UpdateOrganizationConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOrganizationConformancePackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOrganizationConformancePackInvoker) Invoke() (*model.UpdateOrganizationConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOrganizationConformancePackResponse), nil
	}
}

type ShowResourceHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceHistoryInvoker) Invoke() (*model.ShowResourceHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceHistoryResponse), nil
	}
}

type BatchCreateRemediationExceptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateRemediationExceptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateRemediationExceptionsInvoker) Invoke() (*model.BatchCreateRemediationExceptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateRemediationExceptionsResponse), nil
	}
}

type BatchDeleteRemediationExceptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteRemediationExceptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteRemediationExceptionsInvoker) Invoke() (*model.BatchDeleteRemediationExceptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteRemediationExceptionsResponse), nil
	}
}

type CollectRemediationExecutionStatusesSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectRemediationExecutionStatusesSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectRemediationExecutionStatusesSummaryInvoker) Invoke() (*model.CollectRemediationExecutionStatusesSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectRemediationExecutionStatusesSummaryResponse), nil
	}
}

type CreateOrUpdateRemediationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrUpdateRemediationConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrUpdateRemediationConfigurationInvoker) Invoke() (*model.CreateOrUpdateRemediationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrUpdateRemediationConfigurationResponse), nil
	}
}

type CreateOrganizationPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrganizationPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateOrganizationPolicyAssignmentInvoker) Invoke() (*model.CreateOrganizationPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrganizationPolicyAssignmentResponse), nil
	}
}

type CreatePolicyAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePolicyAssignmentsInvoker) Invoke() (*model.CreatePolicyAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyAssignmentsResponse), nil
	}
}

type DeleteOrganizationPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOrganizationPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteOrganizationPolicyAssignmentInvoker) Invoke() (*model.DeleteOrganizationPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOrganizationPolicyAssignmentResponse), nil
	}
}

type DeletePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyAssignmentInvoker) Invoke() (*model.DeletePolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyAssignmentResponse), nil
	}
}

type DeleteRemediationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRemediationConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRemediationConfigurationInvoker) Invoke() (*model.DeleteRemediationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRemediationConfigurationResponse), nil
	}
}

type DisablePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisablePolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisablePolicyAssignmentInvoker) Invoke() (*model.DisablePolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisablePolicyAssignmentResponse), nil
	}
}

type EnablePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnablePolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnablePolicyAssignmentInvoker) Invoke() (*model.EnablePolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnablePolicyAssignmentResponse), nil
	}
}

type ListBuiltInPolicyDefinitionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBuiltInPolicyDefinitionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBuiltInPolicyDefinitionsInvoker) Invoke() (*model.ListBuiltInPolicyDefinitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBuiltInPolicyDefinitionsResponse), nil
	}
}

type ListOrganizationPolicyAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationPolicyAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOrganizationPolicyAssignmentsInvoker) Invoke() (*model.ListOrganizationPolicyAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationPolicyAssignmentsResponse), nil
	}
}

type ListPolicyAssignmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyAssignmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyAssignmentsInvoker) Invoke() (*model.ListPolicyAssignmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyAssignmentsResponse), nil
	}
}

type ListPolicyStatesByAssignmentIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyStatesByAssignmentIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyStatesByAssignmentIdInvoker) Invoke() (*model.ListPolicyStatesByAssignmentIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyStatesByAssignmentIdResponse), nil
	}
}

type ListPolicyStatesByDomainIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyStatesByDomainIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyStatesByDomainIdInvoker) Invoke() (*model.ListPolicyStatesByDomainIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyStatesByDomainIdResponse), nil
	}
}

type ListPolicyStatesByResourceIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyStatesByResourceIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyStatesByResourceIdInvoker) Invoke() (*model.ListPolicyStatesByResourceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyStatesByResourceIdResponse), nil
	}
}

type ListRemediationExceptionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRemediationExceptionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRemediationExceptionsInvoker) Invoke() (*model.ListRemediationExceptionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRemediationExceptionsResponse), nil
	}
}

type ListRemediationExecutionStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRemediationExecutionStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRemediationExecutionStatusesInvoker) Invoke() (*model.ListRemediationExecutionStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRemediationExecutionStatusesResponse), nil
	}
}

type RunEvaluationByPolicyAssignmentIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunEvaluationByPolicyAssignmentIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunEvaluationByPolicyAssignmentIdInvoker) Invoke() (*model.RunEvaluationByPolicyAssignmentIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunEvaluationByPolicyAssignmentIdResponse), nil
	}
}

type RunRemediationExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRemediationExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunRemediationExecutionInvoker) Invoke() (*model.RunRemediationExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRemediationExecutionResponse), nil
	}
}

type ShowBuiltInPolicyDefinitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuiltInPolicyDefinitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuiltInPolicyDefinitionInvoker) Invoke() (*model.ShowBuiltInPolicyDefinitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuiltInPolicyDefinitionResponse), nil
	}
}

type ShowEvaluationStateByAssignmentIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEvaluationStateByAssignmentIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEvaluationStateByAssignmentIdInvoker) Invoke() (*model.ShowEvaluationStateByAssignmentIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEvaluationStateByAssignmentIdResponse), nil
	}
}

type ShowOrganizationPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentInvoker) Invoke() (*model.ShowOrganizationPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationPolicyAssignmentResponse), nil
	}
}

type ShowOrganizationPolicyAssignmentDetailedStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentDetailedStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentDetailedStatusInvoker) Invoke() (*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationPolicyAssignmentDetailedStatusResponse), nil
	}
}

type ShowOrganizationPolicyAssignmentStatusesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentStatusesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOrganizationPolicyAssignmentStatusesInvoker) Invoke() (*model.ShowOrganizationPolicyAssignmentStatusesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationPolicyAssignmentStatusesResponse), nil
	}
}

type ShowPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyAssignmentInvoker) Invoke() (*model.ShowPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyAssignmentResponse), nil
	}
}

type ShowRemediationConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRemediationConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRemediationConfigurationInvoker) Invoke() (*model.ShowRemediationConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRemediationConfigurationResponse), nil
	}
}

type UpdateOrganizationPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOrganizationPolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateOrganizationPolicyAssignmentInvoker) Invoke() (*model.UpdateOrganizationPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOrganizationPolicyAssignmentResponse), nil
	}
}

type UpdatePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyAssignmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyAssignmentInvoker) Invoke() (*model.UpdatePolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyAssignmentResponse), nil
	}
}

type UpdatePolicyStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyStateInvoker) Invoke() (*model.UpdatePolicyStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyStateResponse), nil
	}
}

type CreateStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateStoredQueryInvoker) Invoke() (*model.CreateStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStoredQueryResponse), nil
	}
}

type DeleteStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteStoredQueryInvoker) Invoke() (*model.DeleteStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStoredQueryResponse), nil
	}
}

type ListSchemasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSchemasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSchemasInvoker) Invoke() (*model.ListSchemasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSchemasResponse), nil
	}
}

type ListStoredQueriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStoredQueriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStoredQueriesInvoker) Invoke() (*model.ListStoredQueriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStoredQueriesResponse), nil
	}
}

type RunQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunQueryInvoker) Invoke() (*model.RunQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryResponse), nil
	}
}

type ShowStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStoredQueryInvoker) Invoke() (*model.ShowStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStoredQueryResponse), nil
	}
}

type UpdateStoredQueryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStoredQueryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateStoredQueryInvoker) Invoke() (*model.UpdateStoredQueryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStoredQueryResponse), nil
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

type ShowResourceRelationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceRelationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceRelationsInvoker) Invoke() (*model.ShowResourceRelationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceRelationsResponse), nil
	}
}

type ShowResourceRelationsDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceRelationsDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceRelationsDetailInvoker) Invoke() (*model.ShowResourceRelationsDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceRelationsDetailResponse), nil
	}
}

type CollectAllResourcesSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectAllResourcesSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectAllResourcesSummaryInvoker) Invoke() (*model.CollectAllResourcesSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

type CollectTrackedResourcesSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectTrackedResourcesSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CollectTrackedResourcesSummaryInvoker) Invoke() (*model.CollectTrackedResourcesSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectTrackedResourcesSummaryResponse), nil
	}
}

type CountAllResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountAllResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountAllResourcesInvoker) Invoke() (*model.CountAllResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountAllResourcesResponse), nil
	}
}

type CountTrackedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountTrackedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountTrackedResourcesInvoker) Invoke() (*model.CountTrackedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountTrackedResourcesResponse), nil
	}
}

type ListAllResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllResourcesInvoker) Invoke() (*model.ListAllResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllResourcesResponse), nil
	}
}

type ListAllTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllTagsInvoker) Invoke() (*model.ListAllTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTagsResponse), nil
	}
}

type ListProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProvidersInvoker) Invoke() (*model.ListProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProvidersResponse), nil
	}
}

type ListResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
	}
}

type ListTrackedResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrackedResourceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrackedResourceTagsInvoker) Invoke() (*model.ListTrackedResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrackedResourceTagsResponse), nil
	}
}

type ListTrackedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrackedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTrackedResourcesInvoker) Invoke() (*model.ListTrackedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrackedResourcesResponse), nil
	}
}

type ShowResourceByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceByIdInvoker) Invoke() (*model.ShowResourceByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceByIdResponse), nil
	}
}

type ShowResourceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceDetailInvoker) Invoke() (*model.ShowResourceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceDetailResponse), nil
	}
}

type ShowTrackedResourceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrackedResourceDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrackedResourceDetailInvoker) Invoke() (*model.ShowTrackedResourceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrackedResourceDetailResponse), nil
	}
}

type CountResourcesByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourcesByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountResourcesByTagInvoker) Invoke() (*model.CountResourcesByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourcesByTagResponse), nil
	}
}

type ListResourcesByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourcesByTagInvoker) Invoke() (*model.ListResourcesByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesByTagResponse), nil
	}
}

type ListTagsForResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsForResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsForResourceInvoker) Invoke() (*model.ListTagsForResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsForResourceResponse), nil
	}
}

type ListTagsForResourceTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsForResourceTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsForResourceTypeInvoker) Invoke() (*model.ListTagsForResourceTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsForResourceTypeResponse), nil
	}
}

type TagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagResourceInvoker) Invoke() (*model.TagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagResourceResponse), nil
	}
}

type UnTagResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnTagResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UnTagResourceInvoker) Invoke() (*model.UnTagResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnTagResourceResponse), nil
	}
}

type CreateTrackerConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrackerConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrackerConfigInvoker) Invoke() (*model.CreateTrackerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrackerConfigResponse), nil
	}
}

type DeleteTrackerConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrackerConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrackerConfigInvoker) Invoke() (*model.DeleteTrackerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrackerConfigResponse), nil
	}
}

type ShowTrackerConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrackerConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrackerConfigInvoker) Invoke() (*model.ShowTrackerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrackerConfigResponse), nil
	}
}
