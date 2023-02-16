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

type RunDeleteCustomTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDeleteCustomTagsInvoker) Invoke() (*model.RunDeleteCustomTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDeleteCustomTagsResponse), nil
	}
}

type RunImageDescriptionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageDescriptionInvoker) Invoke() (*model.RunImageDescriptionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageDescriptionResponse), nil
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

type RunImageMediaTaggingDetInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageMediaTaggingDetInvoker) Invoke() (*model.RunImageMediaTaggingDetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageMediaTaggingDetResponse), nil
	}
}

type RunImageSuperResolutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunImageSuperResolutionInvoker) Invoke() (*model.RunImageSuperResolutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunImageSuperResolutionResponse), nil
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

type RunQueryCustomTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunQueryCustomTagsInvoker) Invoke() (*model.RunQueryCustomTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunQueryCustomTagsResponse), nil
	}
}

type RunRecaptureDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunRecaptureDetectInvoker) Invoke() (*model.RunRecaptureDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunRecaptureDetectResponse), nil
	}
}
