package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/config/v1/model"
)

type CreateAggregationAuthorizationInvoker struct {
	*invoker.BaseInvoker
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

func (i *CreateConformancePackInvoker) Invoke() (*model.CreateConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConformancePackResponse), nil
	}
}

type DeleteConformancePackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConformancePackInvoker) Invoke() (*model.DeleteConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConformancePackResponse), nil
	}
}

type ListBuiltInConformancePackTemplatesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListConformancePacksInvoker) Invoke() (*model.ListConformancePacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConformancePacksResponse), nil
	}
}

type ShowBuiltInConformancePackTemplateInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowConformancePackInvoker) Invoke() (*model.ShowConformancePackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConformancePackResponse), nil
	}
}

type ShowResourceHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceHistoryInvoker) Invoke() (*model.ShowResourceHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceHistoryResponse), nil
	}
}

type CreateOrganizationPolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeletePolicyAssignmentInvoker) Invoke() (*model.DeletePolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyAssignmentResponse), nil
	}
}

type DisablePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListPolicyStatesByResourceIdInvoker) Invoke() (*model.ListPolicyStatesByResourceIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyStatesByResourceIdResponse), nil
	}
}

type RunEvaluationByPolicyAssignmentIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunEvaluationByPolicyAssignmentIdInvoker) Invoke() (*model.RunEvaluationByPolicyAssignmentIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunEvaluationByPolicyAssignmentIdResponse), nil
	}
}

type ShowBuiltInPolicyDefinitionInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowPolicyAssignmentInvoker) Invoke() (*model.ShowPolicyAssignmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyAssignmentResponse), nil
	}
}

type UpdatePolicyAssignmentInvoker struct {
	*invoker.BaseInvoker
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

func (i *CollectAllResourcesSummaryInvoker) Invoke() (*model.CollectAllResourcesSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectAllResourcesSummaryResponse), nil
	}
}

type CountAllResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountAllResourcesInvoker) Invoke() (*model.CountAllResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountAllResourcesResponse), nil
	}
}

type ListAllResourcesInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListResourcesInvoker) Invoke() (*model.ListResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesResponse), nil
	}
}

type ShowResourceByIdInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowResourceDetailInvoker) Invoke() (*model.ShowResourceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceDetailResponse), nil
	}
}

type CreateTrackerConfigInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowTrackerConfigInvoker) Invoke() (*model.ShowTrackerConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrackerConfigResponse), nil
	}
}
