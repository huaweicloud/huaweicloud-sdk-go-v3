package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
)

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
