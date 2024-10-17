package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/model"
)

type RunCelebrityRecognitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunCelebrityRecognitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunImageMainObjectDetectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RunImageMediaTaggingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunImageMediaTaggingInvoker) Invoke() (*model.RunImageMediaTaggingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageMediaTaggingResponse), nil
	}
}

type RunImageMediaTaggingDetInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageMediaTaggingDetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunImageMediaTaggingDetInvoker) Invoke() (*model.RunImageMediaTaggingDetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageMediaTaggingDetResponse), nil
	}
}

type RunImageTaggingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageTaggingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunImageTaggingInvoker) Invoke() (*model.RunImageTaggingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageTaggingResponse), nil
	}
}

type RunRecaptureDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRecaptureDetectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunRecaptureDetectInvoker) Invoke() (*model.RunRecaptureDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRecaptureDetectResponse), nil
	}
}
