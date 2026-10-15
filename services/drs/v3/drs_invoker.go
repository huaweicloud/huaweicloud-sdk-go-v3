package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/drs/v3/model"
)

type BatchChangeDataInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchChangeDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchCheckJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchCheckResultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchCreateJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchDeleteJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchListJobDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchListJobStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

func (i *BatchListProgressesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchListRposAndRtosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchListStructDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchListStructProcessInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchResetPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchRestoreTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSetDefinerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSetObjectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSetPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSetSmnInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSetSpeedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchShowParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchStartJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchStopJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchSwitchoverInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchUpdateJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchUpdateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchValidateClustersConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *BatchValidateConnectionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateCompareResultFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateCompareTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateDataLevelTableCompareJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateObjectLevelCompareJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteCompareJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DownloadCompareResultFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListAvailableNodeTypesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

func (i *ListAvailableZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListCompareResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListContentCompareDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListContentCompareDifferenceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListContentCompareOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDataCompareDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDataCompareOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListDataLevelTableCompareJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListObejectLevelCompareDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListObejectLevelCompareOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type ModifyGroupAndStreamInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ModifyGroupAndStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ModifyGroupAndStreamInvoker) Invoke() (*model.ModifyGroupAndStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyGroupAndStreamResponse), nil
	}
}

type QueryTuningParamsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *QueryTuningParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *QueryTuningParamsInvoker) Invoke() (*model.QueryTuningParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.QueryTuningParamsResponse), nil
	}
}

type SelectGroupAndStreamInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SelectGroupAndStreamInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *SelectGroupAndStreamInvoker) Invoke() (*model.SelectGroupAndStreamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SelectGroupAndStreamResponse), nil
	}
}

type ShowJobListInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowJobListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowMonitoringDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

func (i *ShowQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *StartPromptlyDataLevelTableCompareJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateTuningParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *UpdateTuningParamsInvoker) Invoke() (*model.UpdateTuningParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTuningParamsResponse), nil
	}
}
