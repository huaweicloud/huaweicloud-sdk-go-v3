package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
)

type BatchAddSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddSharedTagsInvoker) Invoke() (*model.BatchAddSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddSharedTagsResponse), nil
	}
}

type ChangeSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) Invoke() (*model.ChangeSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSecurityGroupResponse), nil
	}
}

type ChangeShareNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeShareNameInvoker) Invoke() (*model.ChangeShareNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeShareNameResponse), nil
	}
}

type CreateFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFsDirInvoker) Invoke() (*model.CreateFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFsDirResponse), nil
	}
}

type CreateFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFsDirQuotaInvoker) Invoke() (*model.CreateFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFsDirQuotaResponse), nil
	}
}

type CreateShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareInvoker) Invoke() (*model.CreateShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareResponse), nil
	}
}

type CreateSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSharedTagInvoker) Invoke() (*model.CreateSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSharedTagResponse), nil
	}
}

type DeleteFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFsDirInvoker) Invoke() (*model.DeleteFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFsDirResponse), nil
	}
}

type DeleteFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFsDirQuotaInvoker) Invoke() (*model.DeleteFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFsDirQuotaResponse), nil
	}
}

type DeleteShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareInvoker) Invoke() (*model.DeleteShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareResponse), nil
	}
}

type DeleteSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSharedTagInvoker) Invoke() (*model.DeleteSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSharedTagResponse), nil
	}
}

type ExpandShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandShareInvoker) Invoke() (*model.ExpandShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandShareResponse), nil
	}
}

type ListSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedTagsInvoker) Invoke() (*model.ListSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedTagsResponse), nil
	}
}

type ListSharesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharesInvoker) Invoke() (*model.ListSharesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharesResponse), nil
	}
}

type ShowFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsDirInvoker) Invoke() (*model.ShowFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsDirResponse), nil
	}
}

type ShowFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsDirQuotaInvoker) Invoke() (*model.ShowFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsDirQuotaResponse), nil
	}
}

type ShowShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShareInvoker) Invoke() (*model.ShowShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShareResponse), nil
	}
}

type ShowSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSharedTagsInvoker) Invoke() (*model.ShowSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSharedTagsResponse), nil
	}
}

type UpdateFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFsDirQuotaInvoker) Invoke() (*model.UpdateFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFsDirQuotaResponse), nil
	}
}
