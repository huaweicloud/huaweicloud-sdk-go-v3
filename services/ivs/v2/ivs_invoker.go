package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ivs/v2/model"
)

type DetectExtentionByIdCardImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectExtentionByIdCardImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectExtentionByIdCardImageInvoker) Invoke() (*model.DetectExtentionByIdCardImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectExtentionByIdCardImageResponse), nil
	}
}

type DetectExtentionByNameAndIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectExtentionByNameAndIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectExtentionByNameAndIdInvoker) Invoke() (*model.DetectExtentionByNameAndIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectExtentionByNameAndIdResponse), nil
	}
}

type DetectStandardByIdCardImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectStandardByIdCardImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectStandardByIdCardImageInvoker) Invoke() (*model.DetectStandardByIdCardImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectStandardByIdCardImageResponse), nil
	}
}

type DetectStandardByNameAndIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectStandardByNameAndIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectStandardByNameAndIdInvoker) Invoke() (*model.DetectStandardByNameAndIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectStandardByNameAndIdResponse), nil
	}
}

type DetectStandardByVideoAndIdCardImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectStandardByVideoAndIdCardImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectStandardByVideoAndIdCardImageInvoker) Invoke() (*model.DetectStandardByVideoAndIdCardImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectStandardByVideoAndIdCardImageResponse), nil
	}
}

type DetectStandardByVideoAndNameAndIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetectStandardByVideoAndNameAndIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DetectStandardByVideoAndNameAndIdInvoker) Invoke() (*model.DetectStandardByVideoAndNameAndIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetectStandardByVideoAndNameAndIdResponse), nil
	}
}
