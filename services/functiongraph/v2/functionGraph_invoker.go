package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/model"
)

type AsyncInvokeFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *AsyncInvokeFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AsyncInvokeFunctionInvoker) Invoke() (*model.AsyncInvokeFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AsyncInvokeFunctionResponse), nil
	}
}

type BatchDeleteFunctionTriggersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteFunctionTriggersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteFunctionTriggersInvoker) Invoke() (*model.BatchDeleteFunctionTriggersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteFunctionTriggersResponse), nil
	}
}

type BatchDeleteWorkflowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteWorkflowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteWorkflowsInvoker) Invoke() (*model.BatchDeleteWorkflowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteWorkflowsResponse), nil
	}
}

type CancelAsyncInvocationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAsyncInvocationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAsyncInvocationInvoker) Invoke() (*model.CancelAsyncInvocationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAsyncInvocationResponse), nil
	}
}

type CreateCallbackWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCallbackWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCallbackWorkflowInvoker) Invoke() (*model.CreateCallbackWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCallbackWorkflowResponse), nil
	}
}

type CreateDependencyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDependencyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDependencyVersionInvoker) Invoke() (*model.CreateDependencyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDependencyVersionResponse), nil
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

type CreateFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFunctionInvoker) Invoke() (*model.CreateFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFunctionResponse), nil
	}
}

type CreateFunctionAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFunctionAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFunctionAppInvoker) Invoke() (*model.CreateFunctionAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFunctionAppResponse), nil
	}
}

type CreateFunctionTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFunctionTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFunctionTriggerInvoker) Invoke() (*model.CreateFunctionTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFunctionTriggerResponse), nil
	}
}

type CreateFunctionVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFunctionVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFunctionVersionInvoker) Invoke() (*model.CreateFunctionVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFunctionVersionResponse), nil
	}
}

type CreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTagsInvoker) Invoke() (*model.CreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTagsResponse), nil
	}
}

type CreateVersionAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVersionAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVersionAliasInvoker) Invoke() (*model.CreateVersionAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVersionAliasResponse), nil
	}
}

type CreateVpcEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVpcEndpointInvoker) Invoke() (*model.CreateVpcEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcEndpointResponse), nil
	}
}

type CreateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowInvoker) Invoke() (*model.CreateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowResponse), nil
	}
}

type DeleteDependencyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDependencyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDependencyVersionInvoker) Invoke() (*model.DeleteDependencyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDependencyVersionResponse), nil
	}
}

type DeleteEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEventInvoker) Invoke() (*model.DeleteEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEventResponse), nil
	}
}

type DeleteFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFunctionInvoker) Invoke() (*model.DeleteFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFunctionResponse), nil
	}
}

type DeleteFunctionAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFunctionAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFunctionAppInvoker) Invoke() (*model.DeleteFunctionAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFunctionAppResponse), nil
	}
}

type DeleteFunctionAsyncInvokeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFunctionAsyncInvokeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFunctionAsyncInvokeConfigInvoker) Invoke() (*model.DeleteFunctionAsyncInvokeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFunctionAsyncInvokeConfigResponse), nil
	}
}

type DeleteFunctionTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFunctionTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFunctionTriggerInvoker) Invoke() (*model.DeleteFunctionTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFunctionTriggerResponse), nil
	}
}

type DeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTagsInvoker) Invoke() (*model.DeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTagsResponse), nil
	}
}

type DeleteVersionAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVersionAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVersionAliasInvoker) Invoke() (*model.DeleteVersionAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVersionAliasResponse), nil
	}
}

type DeleteVpcEndpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcEndpointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVpcEndpointInvoker) Invoke() (*model.DeleteVpcEndpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcEndpointResponse), nil
	}
}

type EnableAsyncStatusLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableAsyncStatusLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableAsyncStatusLogInvoker) Invoke() (*model.EnableAsyncStatusLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableAsyncStatusLogResponse), nil
	}
}

type EnableLtsLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableLtsLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableLtsLogsInvoker) Invoke() (*model.EnableLtsLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableLtsLogsResponse), nil
	}
}

type ExportFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportFunctionInvoker) Invoke() (*model.ExportFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportFunctionResponse), nil
	}
}

type ImportFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportFunctionInvoker) Invoke() (*model.ImportFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportFunctionResponse), nil
	}
}

type InvokeFunctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *InvokeFunctionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InvokeFunctionInvoker) Invoke() (*model.InvokeFunctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvokeFunctionResponse), nil
	}
}

type ListActiveAsyncInvocationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListActiveAsyncInvocationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListActiveAsyncInvocationsInvoker) Invoke() (*model.ListActiveAsyncInvocationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListActiveAsyncInvocationsResponse), nil
	}
}

type ListAppTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppTemplatesInvoker) Invoke() (*model.ListAppTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppTemplatesResponse), nil
	}
}

type ListAsyncInvocationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAsyncInvocationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAsyncInvocationsInvoker) Invoke() (*model.ListAsyncInvocationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAsyncInvocationsResponse), nil
	}
}

type ListBridgeFunctionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBridgeFunctionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBridgeFunctionsInvoker) Invoke() (*model.ListBridgeFunctionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBridgeFunctionsResponse), nil
	}
}

type ListBridgeVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBridgeVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBridgeVersionsInvoker) Invoke() (*model.ListBridgeVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBridgeVersionsResponse), nil
	}
}

type ListDependenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDependenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDependenciesInvoker) Invoke() (*model.ListDependenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDependenciesResponse), nil
	}
}

type ListDependencyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDependencyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDependencyVersionInvoker) Invoke() (*model.ListDependencyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDependencyVersionResponse), nil
	}
}

type ListEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventsInvoker) Invoke() (*model.ListEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventsResponse), nil
	}
}

type ListFunctionApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionApplicationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionApplicationsInvoker) Invoke() (*model.ListFunctionApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionApplicationsResponse), nil
	}
}

type ListFunctionAsMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionAsMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionAsMetricInvoker) Invoke() (*model.ListFunctionAsMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionAsMetricResponse), nil
	}
}

type ListFunctionAsyncInvokeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionAsyncInvokeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionAsyncInvokeConfigInvoker) Invoke() (*model.ListFunctionAsyncInvokeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionAsyncInvokeConfigResponse), nil
	}
}

type ListFunctionReservedInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionReservedInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionReservedInstancesInvoker) Invoke() (*model.ListFunctionReservedInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionReservedInstancesResponse), nil
	}
}

type ListFunctionStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionStatisticsInvoker) Invoke() (*model.ListFunctionStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionStatisticsResponse), nil
	}
}

type ListFunctionTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionTagsInvoker) Invoke() (*model.ListFunctionTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionTagsResponse), nil
	}
}

type ListFunctionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionTemplateInvoker) Invoke() (*model.ListFunctionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionTemplateResponse), nil
	}
}

type ListFunctionTriggersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionTriggersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionTriggersInvoker) Invoke() (*model.ListFunctionTriggersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionTriggersResponse), nil
	}
}

type ListFunctionVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionVersionsInvoker) Invoke() (*model.ListFunctionVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionVersionsResponse), nil
	}
}

type ListFunctionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFunctionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFunctionsInvoker) Invoke() (*model.ListFunctionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFunctionsResponse), nil
	}
}

type ListQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotasInvoker) Invoke() (*model.ListQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotasResponse), nil
	}
}

type ListReservedInstanceConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReservedInstanceConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReservedInstanceConfigsInvoker) Invoke() (*model.ListReservedInstanceConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReservedInstanceConfigsResponse), nil
	}
}

type ListStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListStatisticsInvoker) Invoke() (*model.ListStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStatisticsResponse), nil
	}
}

type ListVersionAliasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVersionAliasesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListVersionAliasesInvoker) Invoke() (*model.ListVersionAliasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVersionAliasesResponse), nil
	}
}

type ListWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowInvoker) Invoke() (*model.ListWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowResponse), nil
	}
}

type ListWorkflowExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) Invoke() (*model.ListWorkflowExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowExecutionsResponse), nil
	}
}

type RetryWorkFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryWorkFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryWorkFlowInvoker) Invoke() (*model.RetryWorkFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryWorkFlowResponse), nil
	}
}

type ShowAppTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppTemplateInvoker) Invoke() (*model.ShowAppTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppTemplateResponse), nil
	}
}

type ShowDependencyVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDependencyVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDependencyVersionInvoker) Invoke() (*model.ShowDependencyVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDependencyVersionResponse), nil
	}
}

type ShowEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventInvoker) Invoke() (*model.ShowEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventResponse), nil
	}
}

type ShowFuncReservedInstanceMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFuncReservedInstanceMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFuncReservedInstanceMetricsInvoker) Invoke() (*model.ShowFuncReservedInstanceMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFuncReservedInstanceMetricsResponse), nil
	}
}

type ShowFuncSnapshotStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFuncSnapshotStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFuncSnapshotStateInvoker) Invoke() (*model.ShowFuncSnapshotStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFuncSnapshotStateResponse), nil
	}
}

type ShowFunctionAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionAppInvoker) Invoke() (*model.ShowFunctionAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionAppResponse), nil
	}
}

type ShowFunctionAsyncInvokeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionAsyncInvokeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionAsyncInvokeConfigInvoker) Invoke() (*model.ShowFunctionAsyncInvokeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionAsyncInvokeConfigResponse), nil
	}
}

type ShowFunctionCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionCodeInvoker) Invoke() (*model.ShowFunctionCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionCodeResponse), nil
	}
}

type ShowFunctionConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionConfigInvoker) Invoke() (*model.ShowFunctionConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionConfigResponse), nil
	}
}

type ShowFunctionMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionMetricsInvoker) Invoke() (*model.ShowFunctionMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionMetricsResponse), nil
	}
}

type ShowFunctionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionTemplateInvoker) Invoke() (*model.ShowFunctionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionTemplateResponse), nil
	}
}

type ShowFunctionTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFunctionTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFunctionTriggerInvoker) Invoke() (*model.ShowFunctionTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFunctionTriggerResponse), nil
	}
}

type ShowLtsLogDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLtsLogDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLtsLogDetailsInvoker) Invoke() (*model.ShowLtsLogDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLtsLogDetailsResponse), nil
	}
}

type ShowProjectAsyncStatusLogInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectAsyncStatusLogInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectAsyncStatusLogInfoInvoker) Invoke() (*model.ShowProjectAsyncStatusLogInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectAsyncStatusLogInfoResponse), nil
	}
}

type ShowProjectTagsListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectTagsListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectTagsListInvoker) Invoke() (*model.ShowProjectTagsListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectTagsListResponse), nil
	}
}

type ShowResInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResInstanceInfoInvoker) Invoke() (*model.ShowResInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResInstanceInfoResponse), nil
	}
}

type ShowTenantMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTenantMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTenantMetricInvoker) Invoke() (*model.ShowTenantMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTenantMetricResponse), nil
	}
}

type ShowTracingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTracingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTracingInvoker) Invoke() (*model.ShowTracingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTracingResponse), nil
	}
}

type ShowVersionAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVersionAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVersionAliasInvoker) Invoke() (*model.ShowVersionAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVersionAliasResponse), nil
	}
}

type ShowWorkFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkFlowInvoker) Invoke() (*model.ShowWorkFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkFlowResponse), nil
	}
}

type ShowWorkFlowMetricInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkFlowMetricInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkFlowMetricInvoker) Invoke() (*model.ShowWorkFlowMetricResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkFlowMetricResponse), nil
	}
}

type ShowWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowExecutionInvoker) Invoke() (*model.ShowWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowExecutionResponse), nil
	}
}

type ShowWorkflowExecutionForPageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowExecutionForPageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowExecutionForPageInvoker) Invoke() (*model.ShowWorkflowExecutionForPageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowExecutionForPageResponse), nil
	}
}

type StartSyncWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartSyncWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartSyncWorkflowExecutionInvoker) Invoke() (*model.StartSyncWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartSyncWorkflowExecutionResponse), nil
	}
}

type StartWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartWorkflowExecutionInvoker) Invoke() (*model.StartWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartWorkflowExecutionResponse), nil
	}
}

type StopWorkFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopWorkFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopWorkFlowInvoker) Invoke() (*model.StopWorkFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopWorkFlowResponse), nil
	}
}

type UpdateEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEventInvoker) Invoke() (*model.UpdateEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEventResponse), nil
	}
}

type UpdateFuncSnapshotInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFuncSnapshotInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFuncSnapshotInvoker) Invoke() (*model.UpdateFuncSnapshotResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFuncSnapshotResponse), nil
	}
}

type UpdateFunctionAsyncInvokeConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionAsyncInvokeConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionAsyncInvokeConfigInvoker) Invoke() (*model.UpdateFunctionAsyncInvokeConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionAsyncInvokeConfigResponse), nil
	}
}

type UpdateFunctionCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionCodeInvoker) Invoke() (*model.UpdateFunctionCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionCodeResponse), nil
	}
}

type UpdateFunctionCollectStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionCollectStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionCollectStateInvoker) Invoke() (*model.UpdateFunctionCollectStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionCollectStateResponse), nil
	}
}

type UpdateFunctionConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionConfigInvoker) Invoke() (*model.UpdateFunctionConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionConfigResponse), nil
	}
}

type UpdateFunctionMaxInstanceConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionMaxInstanceConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionMaxInstanceConfigInvoker) Invoke() (*model.UpdateFunctionMaxInstanceConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionMaxInstanceConfigResponse), nil
	}
}

type UpdateFunctionReservedInstancesCountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFunctionReservedInstancesCountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFunctionReservedInstancesCountInvoker) Invoke() (*model.UpdateFunctionReservedInstancesCountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFunctionReservedInstancesCountResponse), nil
	}
}

type UpdateTracingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTracingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTracingInvoker) Invoke() (*model.UpdateTracingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTracingResponse), nil
	}
}

type UpdateTriggerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTriggerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTriggerInvoker) Invoke() (*model.UpdateTriggerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTriggerResponse), nil
	}
}

type UpdateVersionAliasInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVersionAliasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVersionAliasInvoker) Invoke() (*model.UpdateVersionAliasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVersionAliasResponse), nil
	}
}

type UpdateWorkFlowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkFlowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkFlowInvoker) Invoke() (*model.UpdateWorkFlowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkFlowResponse), nil
	}
}
