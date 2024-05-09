package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/drs/v3/model"
)

type BatchChangeDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchChangeDataInvoker) Invoke() (*model.BatchChangeDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchChangeDataResponse), nil
	}
}

type BatchCheckJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckJobsInvoker) Invoke() (*model.BatchCheckJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckJobsResponse), nil
	}
}

type BatchCheckResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckResultsInvoker) Invoke() (*model.BatchCheckResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckResultsResponse), nil
	}
}

type BatchCreateJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateJobsInvoker) Invoke() (*model.BatchCreateJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateJobsResponse), nil
	}
}

type BatchDeleteJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteJobsInvoker) Invoke() (*model.BatchDeleteJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteJobsResponse), nil
	}
}

type BatchListJobDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListJobDetailsInvoker) Invoke() (*model.BatchListJobDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListJobDetailsResponse), nil
	}
}

type BatchListJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListJobStatusInvoker) Invoke() (*model.BatchListJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListJobStatusResponse), nil
	}
}

type BatchListProgressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListProgressesInvoker) Invoke() (*model.BatchListProgressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListProgressesResponse), nil
	}
}

type BatchListRposAndRtosInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListRposAndRtosInvoker) Invoke() (*model.BatchListRposAndRtosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListRposAndRtosResponse), nil
	}
}

type BatchListStructDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListStructDetailInvoker) Invoke() (*model.BatchListStructDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListStructDetailResponse), nil
	}
}

type BatchListStructProcessInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListStructProcessInvoker) Invoke() (*model.BatchListStructProcessResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListStructProcessResponse), nil
	}
}

type BatchResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResetPasswordInvoker) Invoke() (*model.BatchResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResetPasswordResponse), nil
	}
}

type BatchRestoreTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRestoreTaskInvoker) Invoke() (*model.BatchRestoreTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRestoreTaskResponse), nil
	}
}

type BatchSetDefinerInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetDefinerInvoker) Invoke() (*model.BatchSetDefinerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetDefinerResponse), nil
	}
}

type BatchSetObjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetObjectsInvoker) Invoke() (*model.BatchSetObjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetObjectsResponse), nil
	}
}

type BatchSetPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetPolicyInvoker) Invoke() (*model.BatchSetPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetPolicyResponse), nil
	}
}

type BatchSetSmnInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetSmnInvoker) Invoke() (*model.BatchSetSmnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetSmnResponse), nil
	}
}

type BatchSetSpeedInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSetSpeedInvoker) Invoke() (*model.BatchSetSpeedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSetSpeedResponse), nil
	}
}

type BatchShowParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowParamsInvoker) Invoke() (*model.BatchShowParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowParamsResponse), nil
	}
}

type BatchStartJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartJobsInvoker) Invoke() (*model.BatchStartJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartJobsResponse), nil
	}
}

type BatchStopJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopJobsInvoker) Invoke() (*model.BatchStopJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopJobsResponse), nil
	}
}

type BatchSwitchoverInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSwitchoverInvoker) Invoke() (*model.BatchSwitchoverResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSwitchoverResponse), nil
	}
}

type BatchUpdateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateJobInvoker) Invoke() (*model.BatchUpdateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateJobResponse), nil
	}
}

type BatchUpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateUserInvoker) Invoke() (*model.BatchUpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateUserResponse), nil
	}
}

type BatchValidateClustersConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchValidateClustersConnectionsInvoker) Invoke() (*model.BatchValidateClustersConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchValidateClustersConnectionsResponse), nil
	}
}

type BatchValidateConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchValidateConnectionsInvoker) Invoke() (*model.BatchValidateConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchValidateConnectionsResponse), nil
	}
}

type CreateCompareResultFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCompareResultFileInvoker) Invoke() (*model.CreateCompareResultFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCompareResultFileResponse), nil
	}
}

type CreateCompareTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCompareTaskInvoker) Invoke() (*model.CreateCompareTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCompareTaskResponse), nil
	}
}

type CreateDataLevelTableCompareJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataLevelTableCompareJobInvoker) Invoke() (*model.CreateDataLevelTableCompareJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataLevelTableCompareJobResponse), nil
	}
}

type CreateObjectLevelCompareJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateObjectLevelCompareJobInvoker) Invoke() (*model.CreateObjectLevelCompareJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateObjectLevelCompareJobResponse), nil
	}
}

type DeleteCompareJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCompareJobInvoker) Invoke() (*model.DeleteCompareJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCompareJobResponse), nil
	}
}

type DownloadCompareResultFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadCompareResultFileInvoker) Invoke() (*model.DownloadCompareResultFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadCompareResultFileResponse), nil
	}
}

type ListAvailableNodeTypesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableNodeTypesInvoker) Invoke() (*model.ListAvailableNodeTypesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableNodeTypesResponse), nil
	}
}

type ListAvailableZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableZoneInvoker) Invoke() (*model.ListAvailableZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableZoneResponse), nil
	}
}

type ListCompareResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCompareResultInvoker) Invoke() (*model.ListCompareResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCompareResultResponse), nil
	}
}

type ListContentCompareDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContentCompareDetailInvoker) Invoke() (*model.ListContentCompareDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContentCompareDetailResponse), nil
	}
}

type ListContentCompareDifferenceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContentCompareDifferenceInvoker) Invoke() (*model.ListContentCompareDifferenceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContentCompareDifferenceResponse), nil
	}
}

type ListContentCompareOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListContentCompareOverviewInvoker) Invoke() (*model.ListContentCompareOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListContentCompareOverviewResponse), nil
	}
}

type ListDataCompareDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataCompareDetailInvoker) Invoke() (*model.ListDataCompareDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataCompareDetailResponse), nil
	}
}

type ListDataCompareOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataCompareOverviewInvoker) Invoke() (*model.ListDataCompareOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataCompareOverviewResponse), nil
	}
}

type ListDataLevelTableCompareJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDataLevelTableCompareJobsInvoker) Invoke() (*model.ListDataLevelTableCompareJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDataLevelTableCompareJobsResponse), nil
	}
}

type ListObejectLevelCompareDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObejectLevelCompareDetailInvoker) Invoke() (*model.ListObejectLevelCompareDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObejectLevelCompareDetailResponse), nil
	}
}

type ListObejectLevelCompareOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListObejectLevelCompareOverviewInvoker) Invoke() (*model.ListObejectLevelCompareOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListObejectLevelCompareOverviewResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ShowJobListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobListInvoker) Invoke() (*model.ShowJobListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobListResponse), nil
	}
}

type ShowMonitoringDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMonitoringDataInvoker) Invoke() (*model.ShowMonitoringDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMonitoringDataResponse), nil
	}
}

type ShowQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInvoker) Invoke() (*model.ShowQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasResponse), nil
	}
}

type StartPromptlyDataLevelTableCompareJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartPromptlyDataLevelTableCompareJobInvoker) Invoke() (*model.StartPromptlyDataLevelTableCompareJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartPromptlyDataLevelTableCompareJobResponse), nil
	}
}

type UpdateParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateParamsInvoker) Invoke() (*model.UpdateParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateParamsResponse), nil
	}
}

type UpdateTuningParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTuningParamsInvoker) Invoke() (*model.UpdateTuningParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTuningParamsResponse), nil
	}
}
