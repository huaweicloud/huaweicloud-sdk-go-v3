package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/mrs/v1/model"
)

type BatchCreateClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateClusterTagsInvoker) Invoke() (*model.BatchCreateClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateClusterTagsResponse), nil
	}
}

type BatchDeleteClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteClusterTagsInvoker) Invoke() (*model.BatchDeleteClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteClusterTagsResponse), nil
	}
}

type CreateAndExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAndExecuteJobInvoker) Invoke() (*model.CreateAndExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAndExecuteJobResponse), nil
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

type CreateClusterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterTagInvoker) Invoke() (*model.CreateClusterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterTagResponse), nil
	}
}

type CreateScalingPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScalingPolicyInvoker) Invoke() (*model.CreateScalingPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScalingPolicyResponse), nil
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

type DeleteClusterTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterTagInvoker) Invoke() (*model.DeleteClusterTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterTagResponse), nil
	}
}

type DeleteJobExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteJobExecutionInvoker) Invoke() (*model.DeleteJobExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteJobExecutionResponse), nil
	}
}

type ListAllTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllTagsInvoker) Invoke() (*model.ListAllTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllTagsResponse), nil
	}
}

type ListClusterTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterTagsInvoker) Invoke() (*model.ListClusterTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterTagsResponse), nil
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

type ListClustersByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClustersByTagsInvoker) Invoke() (*model.ListClustersByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClustersByTagsResponse), nil
	}
}

type ListExecuteJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecuteJobInvoker) Invoke() (*model.ListExecuteJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecuteJobResponse), nil
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

type ShowClusterDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterDetailsInvoker) Invoke() (*model.ShowClusterDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterDetailsResponse), nil
	}
}

type ShowJobExesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobExesInvoker) Invoke() (*model.ShowJobExesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobExesResponse), nil
	}
}

type UpdateClusterScalingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterScalingInvoker) Invoke() (*model.UpdateClusterScalingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterScalingResponse), nil
	}
}

type ListAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableZonesInvoker) Invoke() (*model.ListAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableZonesResponse), nil
	}
}
