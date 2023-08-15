package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/imagesearch/v1/model"
)

type RunAddPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAddPictureInvoker) Invoke() (*model.RunAddPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAddPictureResponse), nil
	}
}

type RunCheckPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCheckPictureInvoker) Invoke() (*model.RunCheckPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCheckPictureResponse), nil
	}
}

type RunCreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateInstanceInvoker) Invoke() (*model.RunCreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCreateInstanceResponse), nil
	}
}

type RunDeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDeleteInstanceInvoker) Invoke() (*model.RunDeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDeleteInstanceResponse), nil
	}
}

type RunDeletePictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDeletePictureInvoker) Invoke() (*model.RunDeletePictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDeletePictureResponse), nil
	}
}

type RunModifyPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunModifyPictureInvoker) Invoke() (*model.RunModifyPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunModifyPictureResponse), nil
	}
}

type RunQueryInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryInstanceInvoker) Invoke() (*model.RunQueryInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryInstanceResponse), nil
	}
}

type RunSearchPictureInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSearchPictureInvoker) Invoke() (*model.RunSearchPictureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSearchPictureResponse), nil
	}
}
