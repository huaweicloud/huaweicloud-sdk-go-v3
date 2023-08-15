package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cms/v1/model"
)

type CreateAutoLaunchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoLaunchGroupInvoker) Invoke() (*model.CreateAutoLaunchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAutoLaunchGroupResponse), nil
	}
}

type DeleteAutoLaunchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAutoLaunchGroupInvoker) Invoke() (*model.DeleteAutoLaunchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAutoLaunchGroupResponse), nil
	}
}

type ListAutoLaunchGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAutoLaunchGroupsInvoker) Invoke() (*model.ListAutoLaunchGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAutoLaunchGroupsResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListSupplyRecommendationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSupplyRecommendationInvoker) Invoke() (*model.ListSupplyRecommendationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSupplyRecommendationResponse), nil
	}
}

type ShowAutoLaunchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAutoLaunchGroupInvoker) Invoke() (*model.ShowAutoLaunchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAutoLaunchGroupResponse), nil
	}
}

type UpdateAutoLaunchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAutoLaunchGroupInvoker) Invoke() (*model.UpdateAutoLaunchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoLaunchGroupResponse), nil
	}
}
