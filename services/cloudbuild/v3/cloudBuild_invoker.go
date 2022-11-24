package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudbuild/v3/model"
)

type DownloadKeystoreInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadKeystoreInvoker) Invoke() (*model.DownloadKeystoreResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadKeystoreResponse), nil
	}
}

type RunJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunJobInvoker) Invoke() (*model.RunJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunJobResponse), nil
	}
}

type ShowHistoryDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryDetailsInvoker) Invoke() (*model.ShowHistoryDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryDetailsResponse), nil
	}
}

type ShowJobListByProjectIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobListByProjectIdInvoker) Invoke() (*model.ShowJobListByProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobListByProjectIdResponse), nil
	}
}

type ShowJobStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobStatusInvoker) Invoke() (*model.ShowJobStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobStatusResponse), nil
	}
}

type ShowJobSuccessRatioInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobSuccessRatioInvoker) Invoke() (*model.ShowJobSuccessRatioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobSuccessRatioResponse), nil
	}
}

type ShowLastHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLastHistoryInvoker) Invoke() (*model.ShowLastHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLastHistoryResponse), nil
	}
}

type ShowListHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListHistoryInvoker) Invoke() (*model.ShowListHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListHistoryResponse), nil
	}
}

type ShowListPeriodHistoryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListPeriodHistoryInvoker) Invoke() (*model.ShowListPeriodHistoryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListPeriodHistoryResponse), nil
	}
}
