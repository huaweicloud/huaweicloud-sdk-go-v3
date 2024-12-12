package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/smsapi/v1/model"
)

type BatchSendDiffSmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSendDiffSmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSendDiffSmsInvoker) Invoke() (*model.BatchSendDiffSmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSendDiffSmsResponse), nil
	}
}

type BatchSendSmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchSendSmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchSendSmsInvoker) Invoke() (*model.BatchSendSmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchSendSmsResponse), nil
	}
}
