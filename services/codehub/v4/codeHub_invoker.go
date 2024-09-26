package v4

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codehub/v4/model"
)

type AssociateGroupUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateGroupUserGroupInvoker) Invoke() (*model.AssociateGroupUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateGroupUserGroupResponse), nil
	}
}

type AssociateRepositoryUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRepositoryUserGroupInvoker) Invoke() (*model.AssociateRepositoryUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRepositoryUserGroupResponse), nil
	}
}

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) Invoke() (*model.CreateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupResponse), nil
	}
}

type DeleteGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type ListManageableGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManageableGroupsInvoker) Invoke() (*model.ListManageableGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManageableGroupsResponse), nil
	}
}

type ShowGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGroupInvoker) Invoke() (*model.ShowGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGroupResponse), nil
	}
}

type AddTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTrustedIpAddressInvoker) Invoke() (*model.AddTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTrustedIpAddressResponse), nil
	}
}

type DeleteTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrustedIpAddressInvoker) Invoke() (*model.DeleteTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrustedIpAddressResponse), nil
	}
}

type ListTrustedIpAddressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTrustedIpAddressesInvoker) Invoke() (*model.ListTrustedIpAddressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTrustedIpAddressesResponse), nil
	}
}

type LockRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *LockRepositoryInvoker) Invoke() (*model.LockRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LockRepositoryResponse), nil
	}
}

type UnlockRepositoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnlockRepositoryInvoker) Invoke() (*model.UnlockRepositoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnlockRepositoryResponse), nil
	}
}

type UpdateTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrustedIpAddressInvoker) Invoke() (*model.UpdateTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrustedIpAddressResponse), nil
	}
}

type AddTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddTenantTrustedIpAddressInvoker) Invoke() (*model.AddTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddTenantTrustedIpAddressResponse), nil
	}
}

type DeleteTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTenantTrustedIpAddressInvoker) Invoke() (*model.DeleteTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTenantTrustedIpAddressResponse), nil
	}
}

type ListTenantTrustedIpAddressesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTenantTrustedIpAddressesInvoker) Invoke() (*model.ListTenantTrustedIpAddressesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTenantTrustedIpAddressesResponse), nil
	}
}

type UpdateTenantTrustedIpAddressInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTenantTrustedIpAddressInvoker) Invoke() (*model.UpdateTenantTrustedIpAddressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTenantTrustedIpAddressResponse), nil
	}
}
