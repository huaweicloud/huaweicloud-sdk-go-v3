package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aidatalake/v2/model"
)

type CancelAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAuraDataframeJobInvoker) Invoke() (*model.CancelAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAuraDataframeJobResponse), nil
	}
}

type CreateAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuraDataframeJobInvoker) Invoke() (*model.CreateAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuraDataframeJobResponse), nil
	}
}

type ListAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraDataframeJobInvoker) Invoke() (*model.ListAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraDataframeJobResponse), nil
	}
}

type ShowAuraDataframeJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraDataframeJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraDataframeJobInvoker) Invoke() (*model.ShowAuraDataframeJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraDataframeJobResponse), nil
	}
}

type CancelAuraSqlStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAuraSqlStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAuraSqlStatementInvoker) Invoke() (*model.CancelAuraSqlStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAuraSqlStatementResponse), nil
	}
}

type CloseAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseAuraSqlSessionInvoker) Invoke() (*model.CloseAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseAuraSqlSessionResponse), nil
	}
}

type CreateAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAuraSqlSessionInvoker) Invoke() (*model.CreateAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAuraSqlSessionResponse), nil
	}
}

type ExecuteAuraSqlStatementInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteAuraSqlStatementInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteAuraSqlStatementInvoker) Invoke() (*model.ExecuteAuraSqlStatementResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteAuraSqlStatementResponse), nil
	}
}

type ListAuraSessionStatementRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraSessionStatementRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraSessionStatementRecordsInvoker) Invoke() (*model.ListAuraSessionStatementRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraSessionStatementRecordsResponse), nil
	}
}

type ListAuraSqlSessionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraSqlSessionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraSqlSessionsInvoker) Invoke() (*model.ListAuraSqlSessionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraSqlSessionsResponse), nil
	}
}

type ListAuraStatementRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAuraStatementRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAuraStatementRecordsInvoker) Invoke() (*model.ListAuraStatementRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAuraStatementRecordsResponse), nil
	}
}

type ShowAuraSqlSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraSqlSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraSqlSessionInvoker) Invoke() (*model.ShowAuraSqlSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraSqlSessionResponse), nil
	}
}

type ShowAuraSqlStatementResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraSqlStatementResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraSqlStatementResultInvoker) Invoke() (*model.ShowAuraSqlStatementResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraSqlStatementResultResponse), nil
	}
}

type ShowAuraStatementOperatorMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraStatementOperatorMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraStatementOperatorMetricsInvoker) Invoke() (*model.ShowAuraStatementOperatorMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraStatementOperatorMetricsResponse), nil
	}
}

type ShowAuraStatementQueryMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuraStatementQueryMetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuraStatementQueryMetricsInvoker) Invoke() (*model.ShowAuraStatementQueryMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuraStatementQueryMetricsResponse), nil
	}
}

type ListFlinkJobAutoscalingHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobAutoscalingHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobAutoscalingHistoriesInvoker) Invoke() (*model.ListFlinkJobAutoscalingHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobAutoscalingHistoriesResponse), nil
	}
}

type ShowFlinkJobAutoscalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobAutoscalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobAutoscalingPolicyInvoker) Invoke() (*model.ShowFlinkJobAutoscalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobAutoscalingPolicyResponse), nil
	}
}

type UpdateFlinkJobAutoscalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlinkJobAutoscalingPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFlinkJobAutoscalingPolicyInvoker) Invoke() (*model.UpdateFlinkJobAutoscalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlinkJobAutoscalingPolicyResponse), nil
	}
}

type CancelFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelFlinkJobInvoker) Invoke() (*model.CancelFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelFlinkJobResponse), nil
	}
}

type ListFlinkJobEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobEventsInvoker) Invoke() (*model.ListFlinkJobEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobEventsResponse), nil
	}
}

type ListFlinkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkJobsInvoker) Invoke() (*model.ListFlinkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkJobsResponse), nil
	}
}

type RunFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunFlinkJobInvoker) Invoke() (*model.RunFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunFlinkJobResponse), nil
	}
}

type ShowFlinkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobInvoker) Invoke() (*model.ShowFlinkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobResponse), nil
	}
}

type ShowFlinkJobStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkJobStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkJobStateInvoker) Invoke() (*model.ShowFlinkJobStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkJobStateResponse), nil
	}
}

type CreateFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFlinkSavepointInvoker) Invoke() (*model.CreateFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlinkSavepointResponse), nil
	}
}

type DeleteFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFlinkSavepointInvoker) Invoke() (*model.DeleteFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlinkSavepointResponse), nil
	}
}

type ListFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFlinkSavepointInvoker) Invoke() (*model.ListFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlinkSavepointResponse), nil
	}
}

type ShowFlinkSavepointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlinkSavepointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlinkSavepointInvoker) Invoke() (*model.ShowFlinkSavepointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlinkSavepointResponse), nil
	}
}

type CreateFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFederationInvoker) Invoke() (*model.CreateFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFederationResponse), nil
	}
}

type DeleteFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFederationInvoker) Invoke() (*model.DeleteFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFederationResponse), nil
	}
}

type ListFederationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFederationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFederationsInvoker) Invoke() (*model.ListFederationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFederationsResponse), nil
	}
}

type ShowFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFederationInvoker) Invoke() (*model.ShowFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFederationResponse), nil
	}
}

type UpdateFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFederationInvoker) Invoke() (*model.UpdateFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFederationResponse), nil
	}
}

type CancelRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelRayJobInvoker) Invoke() (*model.CancelRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelRayJobResponse), nil
	}
}

type ListRayJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRayJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRayJobsInvoker) Invoke() (*model.ListRayJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRayJobsResponse), nil
	}
}

type RunRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunRayJobInvoker) Invoke() (*model.RunRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRayJobResponse), nil
	}
}

type ShowRayJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRayJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRayJobInvoker) Invoke() (*model.ShowRayJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRayJobResponse), nil
	}
}

type CloseRaySessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloseRaySessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloseRaySessionInvoker) Invoke() (*model.CloseRaySessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloseRaySessionResponse), nil
	}
}

type CreateRaySessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRaySessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRaySessionInvoker) Invoke() (*model.CreateRaySessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRaySessionResponse), nil
	}
}

type CancelSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSparkJobInvoker) Invoke() (*model.CancelSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkJobResponse), nil
	}
}

type ListSparkJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSparkJobsInvoker) Invoke() (*model.ListSparkJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkJobsResponse), nil
	}
}

type RunSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSparkJobInvoker) Invoke() (*model.RunSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSparkJobResponse), nil
	}
}

type ShowSparkJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkJobInvoker) Invoke() (*model.ShowSparkJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobResponse), nil
	}
}

type ShowSparkJobStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkJobStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkJobStateInvoker) Invoke() (*model.ShowSparkJobStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkJobStateResponse), nil
	}
}

type ShowSparkOperatorStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkOperatorStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkOperatorStateInvoker) Invoke() (*model.ShowSparkOperatorStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkOperatorStateResponse), nil
	}
}

type CancelSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelSparkSqlInvoker) Invoke() (*model.CancelSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelSparkSqlResponse), nil
	}
}

type ListSparkSqlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSparkSqlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSparkSqlsInvoker) Invoke() (*model.ListSparkSqlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSparkSqlsResponse), nil
	}
}

type PreviewSparkSqlResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *PreviewSparkSqlResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PreviewSparkSqlResultInvoker) Invoke() (*model.PreviewSparkSqlResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PreviewSparkSqlResultResponse), nil
	}
}

type RestartSparkSqlClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartSparkSqlClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartSparkSqlClusterInvoker) Invoke() (*model.RestartSparkSqlClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartSparkSqlClusterResponse), nil
	}
}

type RunSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSparkSqlInvoker) Invoke() (*model.RunSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSparkSqlResponse), nil
	}
}

type ShowSparkSqlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkSqlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkSqlInvoker) Invoke() (*model.ShowSparkSqlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkSqlResponse), nil
	}
}

type ShowSparkSqlStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSparkSqlStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSparkSqlStateInvoker) Invoke() (*model.ShowSparkSqlStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSparkSqlStateResponse), nil
	}
}
