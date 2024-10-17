package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/idme/v1/model"
)

type CreateXdmApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateXdmApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateXdmApplicationInvoker) Invoke() (*model.CreateXdmApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateXdmApplicationResponse), nil
	}
}

type DeleteXdmApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteXdmApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteXdmApplicationInvoker) Invoke() (*model.DeleteXdmApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteXdmApplicationResponse), nil
	}
}

type DeployApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeployApplicationInvoker) Invoke() (*model.DeployApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployApplicationResponse), nil
	}
}

type ListAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppsInvoker) Invoke() (*model.ListAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppsResponse), nil
	}
}

type ListEnvsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvsInvoker) Invoke() (*model.ListEnvsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvsResponse), nil
	}
}

type ModifyApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyApplicationInvoker) Invoke() (*model.ModifyApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyApplicationResponse), nil
	}
}

type UninstallInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UninstallInvoker) Invoke() (*model.UninstallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallResponse), nil
	}
}
