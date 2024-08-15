package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdm/v1/model"
)

type CreateAndStartRandomClusterJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAndStartRandomClusterJobInvoker) Invoke() (*model.CreateAndStartRandomClusterJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAndStartRandomClusterJobResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
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

type CreateLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLinkInvoker) Invoke() (*model.CreateLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLinkResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
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

type DeleteLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLinkInvoker) Invoke() (*model.DeleteLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLinkResponse), nil
	}
}

type ListClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersInvoker) Invoke() (*model.ListClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersResponse), nil
	}
}

type ModifyClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyClusterInvoker) Invoke() (*model.ModifyClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyClusterResponse), nil
	}
}

type RestartClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartClusterInvoker) Invoke() (*model.RestartClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartClusterResponse), nil
	}
}

type ShowAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAvailabilityZonesInvoker) Invoke() (*model.ShowAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAvailabilityZonesResponse), nil
	}
}

type ShowClusterDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterDetailInvoker) Invoke() (*model.ShowClusterDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterDetailResponse), nil
	}
}

type ShowClusterEnterpriseProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterEnterpriseProjectsInvoker) Invoke() (*model.ShowClusterEnterpriseProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterEnterpriseProjectsResponse), nil
	}
}

type ShowDatastoresInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatastoresInvoker) Invoke() (*model.ShowDatastoresResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatastoresResponse), nil
	}
}

type ShowEnterpriseProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnterpriseProjectsInvoker) Invoke() (*model.ShowEnterpriseProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseProjectsResponse), nil
	}
}

type ShowFlavorDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlavorDetailInvoker) Invoke() (*model.ShowFlavorDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlavorDetailResponse), nil
	}
}

type ShowFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlavorsInvoker) Invoke() (*model.ShowFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlavorsResponse), nil
	}
}

type ShowInstanceDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceDetailInvoker) Invoke() (*model.ShowInstanceDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceDetailResponse), nil
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

type ShowJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobsInvoker) Invoke() (*model.ShowJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobsResponse), nil
	}
}

type ShowLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLinkInvoker) Invoke() (*model.ShowLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLinkResponse), nil
	}
}

type ShowSubmissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubmissionsInvoker) Invoke() (*model.ShowSubmissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubmissionsResponse), nil
	}
}

type StartClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartClusterInvoker) Invoke() (*model.StartClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartClusterResponse), nil
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

type StopClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopClusterInvoker) Invoke() (*model.StopClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopClusterResponse), nil
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

type UpdateLinkInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLinkInvoker) Invoke() (*model.UpdateLinkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLinkResponse), nil
	}
}
