package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/pangulargemodels/v1/model"
)

type ExecuteChatCompletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteChatCompletionInvoker) Invoke() (*model.ExecuteChatCompletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteChatCompletionResponse), nil
	}
}

type ExecuteTextCompletionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteTextCompletionInvoker) Invoke() (*model.ExecuteTextCompletionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteTextCompletionResponse), nil
	}
}
