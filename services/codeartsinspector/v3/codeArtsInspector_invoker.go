package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsinspector/v3/model"
)

type AddGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddGroupInvoker) Invoke() (*model.AddGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGroupResponse), nil
	}
}

type DeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type ListGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsInvoker) Invoke() (*model.ListGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsResponse), nil
	}
}

type ListHostResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostResultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostResultsInvoker) Invoke() (*model.ListHostResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostResultsResponse), nil
	}
}

type BatchStartHostTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartHostTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchStartHostTasksInvoker) Invoke() (*model.BatchStartHostTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartHostTasksResponse), nil
	}
}

type BatchCreateHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateHostsInvoker) Invoke() (*model.BatchCreateHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateHostsResponse), nil
	}
}

type DeleteHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostInvoker) Invoke() (*model.DeleteHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostResponse), nil
	}
}

type ListHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostsInvoker) Invoke() (*model.ListHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsResponse), nil
	}
}

type DownloadTaskReportInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadTaskReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ExecuteGenerateReportInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListBusinessRisksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPortResultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowReportStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowResultsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateFalsePositiveInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CancelTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListTaskHistoriesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *AuthorizeDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListDomainsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowDomainSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateDomainSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDomainSettingsInvoker) Invoke() (*model.UpdateDomainSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDomainSettingsResponse), nil
	}
}
