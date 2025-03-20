package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/moderation/v3/model"
)

type BatchCheckImageSyncInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCheckImageSyncInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCheckImageSyncInvoker) Invoke() (*model.BatchCheckImageSyncResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCheckImageSyncResponse), nil
	}
}

type CheckImageModerationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckImageModerationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckImageModerationInvoker) Invoke() (*model.CheckImageModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckImageModerationResponse), nil
	}
}

type RunCloseAudioStreamModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCloseAudioStreamModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunCloseAudioStreamModerationJobInvoker) Invoke() (*model.RunCloseAudioStreamModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCloseAudioStreamModerationJobResponse), nil
	}
}

type RunCreateAudioModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateAudioModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunCreateAudioModerationJobInvoker) Invoke() (*model.RunCreateAudioModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCreateAudioModerationJobResponse), nil
	}
}

type RunCreateAudioStreamModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateAudioStreamModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunCreateAudioStreamModerationJobInvoker) Invoke() (*model.RunCreateAudioStreamModerationJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCreateAudioStreamModerationJobResponse), nil
	}
}

type RunCreateVideoModerationJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCreateVideoModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunQueryAudioModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunQueryVideoModerationJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunTextModerationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTextModerationInvoker) Invoke() (*model.RunTextModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTextModerationResponse), nil
	}
}
