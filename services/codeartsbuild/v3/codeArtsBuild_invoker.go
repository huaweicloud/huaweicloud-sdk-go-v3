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
