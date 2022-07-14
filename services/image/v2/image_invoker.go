package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/model"
)

type RunCelebrityRecognitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCelebrityRecognitionInvoker) Invoke() (*model.RunCelebrityRecognitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunCelebrityRecognitionResponse), nil
	}
}

type RunImageMainObjectDetectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageMainObjectDetectionInvoker) Invoke() (*model.RunImageMainObjectDetectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageMainObjectDetectionResponse), nil
	}
}

type RunImageMediaTaggingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageMediaTaggingInvoker) Invoke() (*model.RunImageMediaTaggingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageMediaTaggingResponse), nil
	}
}

type RunImageTaggingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageTaggingInvoker) Invoke() (*model.RunImageTaggingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageTaggingResponse), nil
	}
}
