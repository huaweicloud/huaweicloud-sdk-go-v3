package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsbuild/v3/model"
)

type CreateBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBuildJobInvoker) Invoke() (*model.CreateBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBuildJobResponse), nil
	}
}

type CreateTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplatesInvoker) Invoke() (*model.CreateTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplatesResponse), nil
	}
}

type DeleteBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBuildJobInvoker) Invoke() (*model.DeleteBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBuildJobResponse), nil
	}
}

type DeleteTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplatesInvoker) Invoke() (*model.DeleteTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplatesResponse), nil
	}
}

type DisableBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableBuildJobInvoker) Invoke() (*model.DisableBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableBuildJobResponse), nil
	}
}

type DisableNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableNoticeInvoker) Invoke() (*model.DisableNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableNoticeResponse), nil
	}
}

type DownloadBuildLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadBuildLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadBuildLogInvoker) Invoke() (*model.DownloadBuildLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadBuildLogResponse), nil
	}
}

type DownloadKeystoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadKeystoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadKeystoreInvoker) Invoke() (*model.DownloadKeystoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadKeystoreResponse), nil
	}
}

type DownloadRealTimeLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadRealTimeLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadRealTimeLogInvoker) Invoke() (*model.DownloadRealTimeLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadRealTimeLogResponse), nil
	}
}

type DownloadTaskLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadTaskLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadTaskLogInvoker) Invoke() (*model.DownloadTaskLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadTaskLogResponse), nil
	}
}

type EnableBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableBuildJobInvoker) Invoke() (*model.EnableBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableBuildJobResponse), nil
	}
}

type ListBuildInfoRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBuildInfoRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBuildInfoRecordInvoker) Invoke() (*model.ListBuildInfoRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBuildInfoRecordResponse), nil
	}
}

type ListJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobConfigInvoker) Invoke() (*model.ListJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobConfigResponse), nil
	}
}

type ListNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNoticeInvoker) Invoke() (*model.ListNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNoticeResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type RunJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunJobInvoker) Invoke() (*model.RunJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunJobResponse), nil
	}
}

type ShowHistoryDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryDetailsInvoker) Invoke() (*model.ShowHistoryDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryDetailsResponse), nil
	}
}

type ShowJobListByProjectIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobListByProjectIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobListByProjectIdInvoker) Invoke() (*model.ShowJobListByProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobListByProjectIdResponse), nil
	}
}

type ShowJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobStatusInvoker) Invoke() (*model.ShowJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobStatusResponse), nil
	}
}

type ShowJobSuccessRatioInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobSuccessRatioInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobSuccessRatioInvoker) Invoke() (*model.ShowJobSuccessRatioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobSuccessRatioResponse), nil
	}
}

type ShowLastHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLastHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLastHistoryInvoker) Invoke() (*model.ShowLastHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLastHistoryResponse), nil
	}
}

type ShowListHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListHistoryInvoker) Invoke() (*model.ShowListHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListHistoryResponse), nil
	}
}

type ShowListPeriodHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListPeriodHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListPeriodHistoryInvoker) Invoke() (*model.ShowListPeriodHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListPeriodHistoryResponse), nil
	}
}

type ShowOutputInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOutputInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOutputInfoInvoker) Invoke() (*model.ShowOutputInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOutputInfoResponse), nil
	}
}

type ShowRecordDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordDetailInvoker) Invoke() (*model.ShowRecordDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordDetailResponse), nil
	}
}

type StopBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopBuildJobInvoker) Invoke() (*model.StopBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopBuildJobResponse), nil
	}
}

type UpdateBuildJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBuildJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateBuildJobInvoker) Invoke() (*model.UpdateBuildJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBuildJobResponse), nil
	}
}

type UpdateNoticeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNoticeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNoticeInvoker) Invoke() (*model.UpdateNoticeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNoticeResponse), nil
	}
}

type ListRelatedProjectInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelatedProjectInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRelatedProjectInfoInvoker) Invoke() (*model.ListRelatedProjectInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelatedProjectInfoResponse), nil
	}
}

type ShowProjectPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectPermissionInvoker) Invoke() (*model.ShowProjectPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectPermissionResponse), nil
	}
}

type ShowRelatedProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRelatedProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRelatedProjectInvoker) Invoke() (*model.ShowRelatedProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRelatedProjectResponse), nil
	}
}

type ShowSummaryBuildJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSummaryBuildJobInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSummaryBuildJobInfoInvoker) Invoke() (*model.ShowSummaryBuildJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSummaryBuildJobInfoResponse), nil
	}
}

type ShowUserOverPackageQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowUserOverPackageQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowUserOverPackageQuotaInvoker) Invoke() (*model.ShowUserOverPackageQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowUserOverPackageQuotaResponse), nil
	}
}

type ShowDockerfileTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDockerfileTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDockerfileTemplateInvoker) Invoke() (*model.ShowDockerfileTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDockerfileTemplateResponse), nil
	}
}

type ShowImageTemplateListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageTemplateListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowImageTemplateListInvoker) Invoke() (*model.ShowImageTemplateListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageTemplateListResponse), nil
	}
}

type CheckJobCountIsTopLimitInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckJobCountIsTopLimitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckJobCountIsTopLimitInvoker) Invoke() (*model.CheckJobCountIsTopLimitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckJobCountIsTopLimitResponse), nil
	}
}

type CheckJobNameIsExistsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckJobNameIsExistsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckJobNameIsExistsInvoker) Invoke() (*model.CheckJobNameIsExistsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckJobNameIsExistsResponse), nil
	}
}

type CheckWebhookUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckWebhookUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckWebhookUrlInvoker) Invoke() (*model.CheckWebhookUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckWebhookUrlResponse), nil
	}
}

type ClearRecyclingJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ClearRecyclingJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ClearRecyclingJobsInvoker) Invoke() (*model.ClearRecyclingJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ClearRecyclingJobsResponse), nil
	}
}

type CopyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyJobInvoker) Invoke() (*model.CopyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyJobResponse), nil
	}
}

type DeleteRecyclingJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecyclingJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecyclingJobsInvoker) Invoke() (*model.DeleteRecyclingJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecyclingJobsResponse), nil
	}
}

type DeleteTheJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTheJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTheJobInvoker) Invoke() (*model.DeleteTheJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTheJobResponse), nil
	}
}

type DisableTheJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableTheJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableTheJobInvoker) Invoke() (*model.DisableTheJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableTheJobResponse), nil
	}
}

type ExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteJobInvoker) Invoke() (*model.ExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteJobResponse), nil
	}
}

type ListBuildParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBuildParameterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBuildParameterInvoker) Invoke() (*model.ListBuildParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBuildParameterResponse), nil
	}
}

type ListJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobInvoker) Invoke() (*model.ListJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobResponse), nil
	}
}

type ListProjectJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProjectJobsInvoker) Invoke() (*model.ListProjectJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectJobsResponse), nil
	}
}

type ListRecyclingJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecyclingJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecyclingJobInvoker) Invoke() (*model.ListRecyclingJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecyclingJobResponse), nil
	}
}

type ListUpdateJobHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUpdateJobHistoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUpdateJobHistoryInvoker) Invoke() (*model.ListUpdateJobHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUpdateJobHistoryResponse), nil
	}
}

type RestoreRecyclingJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreRecyclingJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreRecyclingJobsInvoker) Invoke() (*model.RestoreRecyclingJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreRecyclingJobsResponse), nil
	}
}

type SetKeepTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetKeepTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetKeepTimeInvoker) Invoke() (*model.SetKeepTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetKeepTimeResponse), nil
	}
}

type ShowBuildParamsListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildParamsListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildParamsListInvoker) Invoke() (*model.ShowBuildParamsListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildParamsListResponse), nil
	}
}

type ShowCopyNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCopyNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCopyNameInvoker) Invoke() (*model.ShowCopyNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCopyNameResponse), nil
	}
}

type ShowDefaultBuildParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDefaultBuildParametersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDefaultBuildParametersInvoker) Invoke() (*model.ShowDefaultBuildParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDefaultBuildParametersResponse), nil
	}
}

type ShowDefaultProjectPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDefaultProjectPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDefaultProjectPermissionInvoker) Invoke() (*model.ShowDefaultProjectPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDefaultProjectPermissionResponse), nil
	}
}

type ShowDisableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDisableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDisableInvoker) Invoke() (*model.ShowDisableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDisableResponse), nil
	}
}

type ShowJobConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobConfigInvoker) Invoke() (*model.ShowJobConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobConfigResponse), nil
	}
}

type ShowJobConfigDiffInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobConfigDiffInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobConfigDiffInvoker) Invoke() (*model.ShowJobConfigDiffResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobConfigDiffResponse), nil
	}
}

type ShowJobInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInfoInvoker) Invoke() (*model.ShowJobInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobInfoResponse), nil
	}
}

type ShowJobNoticeConfigInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobNoticeConfigInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobNoticeConfigInfoInvoker) Invoke() (*model.ShowJobNoticeConfigInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobNoticeConfigInfoResponse), nil
	}
}

type ShowJobRolePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobRolePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobRolePermissionInvoker) Invoke() (*model.ShowJobRolePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobRolePermissionResponse), nil
	}
}

type ShowJobStepStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobStepStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobStepStatusInvoker) Invoke() (*model.ShowJobStepStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobStepStatusResponse), nil
	}
}

type ShowJobSystemParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobSystemParametersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobSystemParametersInvoker) Invoke() (*model.ShowJobSystemParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobSystemParametersResponse), nil
	}
}

type ShowRunningStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRunningStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRunningStatusInvoker) Invoke() (*model.ShowRunningStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRunningStatusResponse), nil
	}
}

type UpdateNewJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNewJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNewJobInvoker) Invoke() (*model.UpdateNewJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNewJobResponse), nil
	}
}

type DeleteKeystoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKeystoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKeystoreInvoker) Invoke() (*model.DeleteKeystoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKeystoreResponse), nil
	}
}

type DeleteKeystorePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKeystorePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteKeystorePermissionInvoker) Invoke() (*model.DeleteKeystorePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKeystorePermissionResponse), nil
	}
}

type DownloadKeystoreByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadKeystoreByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadKeystoreByNameInvoker) Invoke() (*model.DownloadKeystoreByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadKeystoreByNameResponse), nil
	}
}

type ListKeystoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeystoreInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKeystoreInvoker) Invoke() (*model.ListKeystoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeystoreResponse), nil
	}
}

type ListKeystoreSearchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeystoreSearchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListKeystoreSearchInvoker) Invoke() (*model.ListKeystoreSearchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeystoreSearchResponse), nil
	}
}

type ShowKeystorePermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKeystorePermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowKeystorePermissionInvoker) Invoke() (*model.ShowKeystorePermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKeystorePermissionResponse), nil
	}
}

type DownloadLogByRecordIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadLogByRecordIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadLogByRecordIdInvoker) Invoke() (*model.DownloadLogByRecordIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadLogByRecordIdResponse), nil
	}
}

type ShowFlowGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFlowGraphInvoker) Invoke() (*model.ShowFlowGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowGraphResponse), nil
	}
}

type ShowRecordInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecordInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecordInfoInvoker) Invoke() (*model.ShowRecordInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecordInfoResponse), nil
	}
}

type StopJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopJobInvoker) Invoke() (*model.StopJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobResponse), nil
	}
}

type ListBriefRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBriefRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBriefRecordInvoker) Invoke() (*model.ListBriefRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBriefRecordResponse), nil
	}
}

type ListBuildInfoRecordByJobIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBuildInfoRecordByJobIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBuildInfoRecordByJobIdInvoker) Invoke() (*model.ListBuildInfoRecordByJobIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBuildInfoRecordByJobIdResponse), nil
	}
}

type ListRecordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordsInvoker) Invoke() (*model.ListRecordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordsResponse), nil
	}
}

type ShowBuildInfoRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildInfoRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildInfoRecordInvoker) Invoke() (*model.ShowBuildInfoRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildInfoRecordResponse), nil
	}
}

type ShowBuildRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildRecordInvoker) Invoke() (*model.ShowBuildRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildRecordResponse), nil
	}
}

type ShowBuildRecordBuildScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildRecordBuildScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildRecordBuildScriptInvoker) Invoke() (*model.ShowBuildRecordBuildScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildRecordBuildScriptResponse), nil
	}
}

type ShowBuildRecordFlowGraphInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildRecordFlowGraphInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildRecordFlowGraphInvoker) Invoke() (*model.ShowBuildRecordFlowGraphResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildRecordFlowGraphResponse), nil
	}
}

type ShowBuildRecordFullStagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBuildRecordFullStagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBuildRecordFullStagesInvoker) Invoke() (*model.ShowBuildRecordFullStagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBuildRecordFullStagesResponse), nil
	}
}

type ShowJobBuildRecordDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobBuildRecordDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobBuildRecordDetailInvoker) Invoke() (*model.ShowJobBuildRecordDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobBuildRecordDetailResponse), nil
	}
}

type ShowJobTotalInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobTotalInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobTotalInvoker) Invoke() (*model.ShowJobTotalResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobTotalResponse), nil
	}
}

type DownloadJunitCoverageZipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadJunitCoverageZipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadJunitCoverageZipInvoker) Invoke() (*model.DownloadJunitCoverageZipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadJunitCoverageZipResponse), nil
	}
}

type ListJunitCoverageSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJunitCoverageSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJunitCoverageSummaryInvoker) Invoke() (*model.ListJunitCoverageSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJunitCoverageSummaryResponse), nil
	}
}

type ListRepoBranchInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepoBranchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepoBranchInvoker) Invoke() (*model.ListRepoBranchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepoBranchResponse), nil
	}
}

type ListRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRepositoryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRepositoryInvoker) Invoke() (*model.ListRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRepositoryResponse), nil
	}
}

type ShowJobBuildSuccessRatioInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobBuildSuccessRatioInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobBuildSuccessRatioInvoker) Invoke() (*model.ShowJobBuildSuccessRatioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobBuildSuccessRatioResponse), nil
	}
}

type ShowJobBuildTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobBuildTimeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobBuildTimeInvoker) Invoke() (*model.ShowJobBuildTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobBuildTimeResponse), nil
	}
}

type ShowReportSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReportSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowReportSummaryInvoker) Invoke() (*model.ShowReportSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReportSummaryResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type ListCustomTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCustomTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCustomTemplateInvoker) Invoke() (*model.ListCustomTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCustomTemplateResponse), nil
	}
}

type ListOfficialTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOfficialTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOfficialTemplateInvoker) Invoke() (*model.ListOfficialTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOfficialTemplateResponse), nil
	}
}

type ListRecommendOfficialTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecommendOfficialTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecommendOfficialTemplateInvoker) Invoke() (*model.ListRecommendOfficialTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecommendOfficialTemplateResponse), nil
	}
}

type ShowYamlTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowYamlTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowYamlTemplateInvoker) Invoke() (*model.ShowYamlTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowYamlTemplateResponse), nil
	}
}
