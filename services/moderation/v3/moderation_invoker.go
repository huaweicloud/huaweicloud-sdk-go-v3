package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v3/model"
)

type CheckImageModerationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckImageModerationInvoker) Invoke() (*model.CheckImageModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckImageModerationResponse), nil
	}
}

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

type RunCreateVideoModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateVideoModerationJobInvoker) Invoke() (*model.RunCreateVideoModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCreateVideoModerationJobResponse), nil
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

type RunQueryVideoModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryVideoModerationJobInvoker) Invoke() (*model.RunQueryVideoModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryVideoModerationJobResponse), nil
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
