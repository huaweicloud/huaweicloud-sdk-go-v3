package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cms/v1/model"
)

type CreateAutoLaunchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAutoLaunchGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteAutoLaunchGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAutoLaunchGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSupplyRecommendationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowAutoLaunchGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateAutoLaunchGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAutoLaunchGroupInvoker) Invoke() (*model.UpdateAutoLaunchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAutoLaunchGroupResponse), nil
	}
}
