package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterscim/v1/model"
)

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteGroupInvoker) Invoke() (*model.DeleteGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGroupResponse), nil
	}
}

type GetGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetGroupInvoker) Invoke() (*model.GetGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetGroupResponse), nil
	}
}

type ListGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListGroupsInvoker) Invoke() (*model.ListGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupsResponse), nil
	}
}

type PatchGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *PatchGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PatchGroupInvoker) Invoke() (*model.PatchGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PatchGroupResponse), nil
	}
}

type ServiceProviderConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ServiceProviderConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ServiceProviderConfigInvoker) Invoke() (*model.ServiceProviderConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServiceProviderConfigResponse), nil
	}
}

type CreateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateUserInvoker) Invoke() (*model.CreateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUserResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type GetUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetUserInvoker) Invoke() (*model.GetUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetUserResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type PatchUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *PatchUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PatchUserInvoker) Invoke() (*model.PatchUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PatchUserResponse), nil
	}
}

type PutUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *PutUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *PutUserInvoker) Invoke() (*model.PutUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PutUserResponse), nil
	}
}
