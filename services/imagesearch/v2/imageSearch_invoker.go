package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/imagesearch/v2/model"
)

type RunAddDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAddDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunAddDataInvoker) Invoke() (*model.RunAddDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAddDataResponse), nil
	}
}

type RunCheckDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCheckDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunCheckDataInvoker) Invoke() (*model.RunCheckDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCheckDataResponse), nil
	}
}

type RunDeleteDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDeleteDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDeleteDataInvoker) Invoke() (*model.RunDeleteDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDeleteDataResponse), nil
	}
}

type RunSearchInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSearchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSearchInvoker) Invoke() (*model.RunSearchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSearchResponse), nil
	}
}

type RunUpdateDataInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunUpdateDataInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunUpdateDataInvoker) Invoke() (*model.RunUpdateDataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunUpdateDataResponse), nil
	}
}
