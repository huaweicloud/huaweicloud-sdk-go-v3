package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/dlf/v1/model"
)

type CancelScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelScriptInvoker) Invoke() (*model.CancelScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelScriptResponse), nil
	}
}

type CreateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionInvoker) Invoke() (*model.CreateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionResponse), nil
	}
}

type CreateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateJobInvoker) Invoke() (*model.CreateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateJobResponse), nil
	}
}

type CreateResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateResourceInvoker) Invoke() (*model.CreateResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateResourceResponse), nil
	}
}

type CreateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScriptInvoker) Invoke() (*model.CreateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScriptResponse), nil
	}
}

type DeleteConnctionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnctionInvoker) Invoke() (*model.DeleteConnctionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnctionResponse), nil
	}
}

type DeleteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobInvoker) Invoke() (*model.DeleteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobResponse), nil
	}
}

type DeleteResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteResourceInvoker) Invoke() (*model.DeleteResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteResourceResponse), nil
	}
}

type DeleteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScriptInvoker) Invoke() (*model.DeleteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScriptResponse), nil
	}
}

type ExecuteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteScriptInvoker) Invoke() (*model.ExecuteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteScriptResponse), nil
	}
}

type ExportConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportConnectionsInvoker) Invoke() (*model.ExportConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportConnectionsResponse), nil
	}
}

type ExportJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportJobInvoker) Invoke() (*model.ExportJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportJobResponse), nil
	}
}

type ExportJobListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportJobListInvoker) Invoke() (*model.ExportJobListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportJobListResponse), nil
	}
}

type ImportConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportConnectionsInvoker) Invoke() (*model.ImportConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportConnectionsResponse), nil
	}
}

type ImportJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportJobInvoker) Invoke() (*model.ImportJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportJobResponse), nil
	}
}

type ListConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionsInvoker) Invoke() (*model.ListConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionsResponse), nil
	}
}

type ListJobInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobInstancesInvoker) Invoke() (*model.ListJobInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobInstancesResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
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

type ListScriptResultsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptResultsInvoker) Invoke() (*model.ListScriptResultsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptResultsResponse), nil
	}
}

type ListScriptsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListScriptsInvoker) Invoke() (*model.ListScriptsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListScriptsResponse), nil
	}
}

type ListSystemTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemTasksInvoker) Invoke() (*model.ListSystemTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemTasksResponse), nil
	}
}

type RestoreJobInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreJobInstanceInvoker) Invoke() (*model.RestoreJobInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreJobInstanceResponse), nil
	}
}

type RunOnceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunOnceInvoker) Invoke() (*model.RunOnceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunOnceResponse), nil
	}
}

type ShowConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectionInvoker) Invoke() (*model.ShowConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectionResponse), nil
	}
}

type ShowDirectoryTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDirectoryTreeInvoker) Invoke() (*model.ShowDirectoryTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDirectoryTreeResponse), nil
	}
}

type ShowFileInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFileInfoInvoker) Invoke() (*model.ShowFileInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFileInfoResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowJobInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInstanceInvoker) Invoke() (*model.ShowJobInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobInstanceResponse), nil
	}
}

type ShowJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobStatusInvoker) Invoke() (*model.ShowJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobStatusResponse), nil
	}
}

type ShowResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceInvoker) Invoke() (*model.ShowResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceResponse), nil
	}
}

type ShowScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowScriptInvoker) Invoke() (*model.ShowScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowScriptResponse), nil
	}
}

type StartJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartJobInvoker) Invoke() (*model.StartJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartJobResponse), nil
	}
}

type StopJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInvoker) Invoke() (*model.StopJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobResponse), nil
	}
}

type StopJobInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopJobInstanceInvoker) Invoke() (*model.StopJobInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopJobInstanceResponse), nil
	}
}

type UpdateConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConnectionInvoker) Invoke() (*model.UpdateConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConnectionResponse), nil
	}
}

type UpdateJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateJobInvoker) Invoke() (*model.UpdateJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateJobResponse), nil
	}
}

type UpdateResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateResourceInvoker) Invoke() (*model.UpdateResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateResourceResponse), nil
	}
}

type UpdateScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScriptInvoker) Invoke() (*model.UpdateScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScriptResponse), nil
	}
}
