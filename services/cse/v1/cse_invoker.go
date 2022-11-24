package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
)

type CreateEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEngineInvoker) Invoke() (*model.CreateEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEngineResponse), nil
	}
}

type DeleteEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEngineInvoker) Invoke() (*model.DeleteEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEngineResponse), nil
	}
}

type DownloadKieInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadKieInvoker) Invoke() (*model.DownloadKieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadKieResponse), nil
	}
}

type ListEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnginesInvoker) Invoke() (*model.ListEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnginesResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type RetryEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryEngineInvoker) Invoke() (*model.RetryEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryEngineResponse), nil
	}
}

type ShowEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineInvoker) Invoke() (*model.ShowEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineResponse), nil
	}
}

type ShowEngineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineJobInvoker) Invoke() (*model.ShowEngineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineJobResponse), nil
	}
}

type UpgradeEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeEngineInvoker) Invoke() (*model.UpgradeEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeEngineResponse), nil
	}
}

type UploadKieInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadKieInvoker) Invoke() (*model.UploadKieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadKieResponse), nil
	}
}
