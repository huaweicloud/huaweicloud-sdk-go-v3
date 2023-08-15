package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/vss/v3/model"
)

type DownloadTaskReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadTaskReportInvoker) Invoke() (*model.DownloadTaskReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadTaskReportResponse), nil
	}
}

type ExecuteGenerateReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGenerateReportInvoker) Invoke() (*model.ExecuteGenerateReportResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGenerateReportResponse), nil
	}
}

type ListBusinessRisksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBusinessRisksInvoker) Invoke() (*model.ListBusinessRisksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBusinessRisksResponse), nil
	}
}

type ListPortResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortResultsInvoker) Invoke() (*model.ListPortResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortResultsResponse), nil
	}
}

type ShowReportStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReportStatusInvoker) Invoke() (*model.ShowReportStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReportStatusResponse), nil
	}
}

type ShowResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResultsInvoker) Invoke() (*model.ShowResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResultsResponse), nil
	}
}

type UpdateFalsePositiveInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFalsePositiveInvoker) Invoke() (*model.UpdateFalsePositiveResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFalsePositiveResponse), nil
	}
}

type CancelTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelTasksInvoker) Invoke() (*model.CancelTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelTasksResponse), nil
	}
}

type CreateTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTasksInvoker) Invoke() (*model.CreateTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTasksResponse), nil
	}
}

type ListTaskHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskHistoriesInvoker) Invoke() (*model.ListTaskHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskHistoriesResponse), nil
	}
}

type ShowTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTasksInvoker) Invoke() (*model.ShowTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTasksResponse), nil
	}
}

type AuthorizeDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeDomainsInvoker) Invoke() (*model.AuthorizeDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeDomainsResponse), nil
	}
}

type CreateDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainsInvoker) Invoke() (*model.CreateDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainsResponse), nil
	}
}

type DeleteDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainsInvoker) Invoke() (*model.DeleteDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainsResponse), nil
	}
}

type ListDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainsInvoker) Invoke() (*model.ListDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainsResponse), nil
	}
}

type ShowDomainSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainSettingsInvoker) Invoke() (*model.ShowDomainSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainSettingsResponse), nil
	}
}

type UpdateDomainSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDomainSettingsInvoker) Invoke() (*model.UpdateDomainSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainSettingsResponse), nil
	}
}
