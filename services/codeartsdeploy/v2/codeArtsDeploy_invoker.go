package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/codeartsdeploy/v2/model"
)

type CreateDeployTaskByTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeployTaskByTemplateInvoker) Invoke() (*model.CreateDeployTaskByTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeployTaskByTemplateResponse), nil
	}
}

type DeleteDeployTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeployTaskInvoker) Invoke() (*model.DeleteDeployTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeployTaskResponse), nil
	}
}

type ListDeployTaskHistoryByDateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeployTaskHistoryByDateInvoker) Invoke() (*model.ListDeployTaskHistoryByDateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeployTaskHistoryByDateResponse), nil
	}
}

type ListDeployTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeployTasksInvoker) Invoke() (*model.ListDeployTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeployTasksResponse), nil
	}
}

type ShowDeployTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeployTaskDetailInvoker) Invoke() (*model.ShowDeployTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeployTaskDetailResponse), nil
	}
}

type StartDeployTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDeployTaskInvoker) Invoke() (*model.StartDeployTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDeployTaskResponse), nil
	}
}

type CreateDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentHostInvoker) Invoke() (*model.CreateDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentHostResponse), nil
	}
}

type DeleteDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentHostInvoker) Invoke() (*model.DeleteDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentHostResponse), nil
	}
}

type ListHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsInvoker) Invoke() (*model.ListHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsResponse), nil
	}
}

type ShowDeploymentHostDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentHostDetailInvoker) Invoke() (*model.ShowDeploymentHostDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentHostDetailResponse), nil
	}
}

type UpdateDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentHostInvoker) Invoke() (*model.UpdateDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentHostResponse), nil
	}
}

type CreateDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentGroupInvoker) Invoke() (*model.CreateDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentGroupResponse), nil
	}
}

type DeleteDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentGroupInvoker) Invoke() (*model.DeleteDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentGroupResponse), nil
	}
}

type ListHostGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupsInvoker) Invoke() (*model.ListHostGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupsResponse), nil
	}
}

type ShowDeploymentGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentGroupDetailInvoker) Invoke() (*model.ShowDeploymentGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentGroupDetailResponse), nil
	}
}

type UpdateDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentGroupInvoker) Invoke() (*model.UpdateDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentGroupResponse), nil
	}
}

type ListTaskSuccessRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskSuccessRateInvoker) Invoke() (*model.ListTaskSuccessRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskSuccessRateResponse), nil
	}
}

type ShowProjectSuccessRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectSuccessRateInvoker) Invoke() (*model.ShowProjectSuccessRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectSuccessRateResponse), nil
	}
}
