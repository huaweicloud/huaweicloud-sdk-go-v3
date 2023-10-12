package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsdeploy/v2/model"
)

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

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

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
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

type ListAllAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllAppInvoker) Invoke() (*model.ListAllAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllAppResponse), nil
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

type ShowAppDetailByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppDetailByIdInvoker) Invoke() (*model.ShowAppDetailByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppDetailByIdResponse), nil
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

type ShowExecutionParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecutionParamsInvoker) Invoke() (*model.ShowExecutionParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecutionParamsResponse), nil
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

type CreateEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentInvoker) Invoke() (*model.CreateEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentResponse), nil
	}
}

type DeleteEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentInvoker) Invoke() (*model.DeleteEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentResponse), nil
	}
}

type DeleteHostFromEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostFromEnvironmentInvoker) Invoke() (*model.DeleteHostFromEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostFromEnvironmentResponse), nil
	}
}

type ImportHostToEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportHostToEnvironmentInvoker) Invoke() (*model.ImportHostToEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportHostToEnvironmentResponse), nil
	}
}

type ListEnvironmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentsInvoker) Invoke() (*model.ListEnvironmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentsResponse), nil
	}
}

type ShowEnvironmentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvironmentDetailInvoker) Invoke() (*model.ShowEnvironmentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvironmentDetailResponse), nil
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

type CreateHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostInvoker) Invoke() (*model.CreateHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostResponse), nil
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

type ListNewHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNewHostsInvoker) Invoke() (*model.ListNewHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNewHostsResponse), nil
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

type ShowHostDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostDetailInvoker) Invoke() (*model.ShowHostDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostDetailResponse), nil
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

type CreateHostClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostClusterInvoker) Invoke() (*model.CreateHostClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostClusterResponse), nil
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

type ListHostClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostClustersInvoker) Invoke() (*model.ListHostClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostClustersResponse), nil
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

type ShowHostClusterDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostClusterDetailInvoker) Invoke() (*model.ShowHostClusterDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostClusterDetailResponse), nil
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
