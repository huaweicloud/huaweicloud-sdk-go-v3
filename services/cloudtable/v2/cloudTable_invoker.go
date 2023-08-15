package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cloudtable/v2/model"
)

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

type EnableComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableComponentInvoker) Invoke() (*model.EnableComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableComponentResponse), nil
	}
}

type ExpandClusterComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandClusterComponentInvoker) Invoke() (*model.ExpandClusterComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandClusterComponentResponse), nil
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

type RebootCloudTableClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootCloudTableClusterInvoker) Invoke() (*model.RebootCloudTableClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootCloudTableClusterResponse), nil
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

type ShowClusterSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterSettingInvoker) Invoke() (*model.ShowClusterSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterSettingResponse), nil
	}
}

type UpdateClusterSettingInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterSettingInvoker) Invoke() (*model.UpdateClusterSettingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterSettingResponse), nil
	}
}
