package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/moderation/v2/model"
)

type RunCheckResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCheckResultInvoker) Invoke() (*model.RunCheckResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCheckResultResponse), nil
	}
}

type RunCheckTaskJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCheckTaskJobsInvoker) Invoke() (*model.RunCheckTaskJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCheckTaskJobsResponse), nil
	}
}

type RunImageBatchModerationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageBatchModerationInvoker) Invoke() (*model.RunImageBatchModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageBatchModerationResponse), nil
	}
}

type RunImageModerationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageModerationInvoker) Invoke() (*model.RunImageModerationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageModerationResponse), nil
	}
}

type RunModerationAudioInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunModerationAudioInvoker) Invoke() (*model.RunModerationAudioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunModerationAudioResponse), nil
	}
}

type RunTaskSumbitInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTaskSumbitInvoker) Invoke() (*model.RunTaskSumbitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTaskSumbitResponse), nil
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
