package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v3/model"
)

type RunCreateAudioModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateAudioModerationJobInvoker) Invoke() (*model.RunCreateAudioModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCreateAudioModerationJobResponse), nil
	}
}

type RunQueryAudioModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryAudioModerationJobInvoker) Invoke() (*model.RunQueryAudioModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryAudioModerationJobResponse), nil
	}
}

type RunTextModerationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTextModerationInvoker) Invoke() (*model.RunTextModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTextModerationResponse), nil
	}
}
