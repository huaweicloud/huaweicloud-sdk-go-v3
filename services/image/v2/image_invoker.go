package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/image/v2/model"
)

type CreateImageHighresolutionMattingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateImageHighresolutionMattingTaskInvoker) Invoke() (*model.CreateImageHighresolutionMattingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateImageHighresolutionMattingTaskResponse), nil
	}
}

type CreateVideoObjectMaskingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVideoObjectMaskingTaskInvoker) Invoke() (*model.CreateVideoObjectMaskingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVideoObjectMaskingTaskResponse), nil
	}
}

type CreateVideoTaggingMediaTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVideoTaggingMediaTaskInvoker) Invoke() (*model.CreateVideoTaggingMediaTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVideoTaggingMediaTaskResponse), nil
	}
}

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

type ShowImageHighresolutionMattingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageHighresolutionMattingTaskInvoker) Invoke() (*model.ShowImageHighresolutionMattingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageHighresolutionMattingTaskResponse), nil
	}
}

type ShowVideoObjectMaskingTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVideoObjectMaskingTaskInvoker) Invoke() (*model.ShowVideoObjectMaskingTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVideoObjectMaskingTaskResponse), nil
	}
}

type ShowVideoTaggingMediaTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVideoTaggingMediaTaskInvoker) Invoke() (*model.ShowVideoTaggingMediaTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVideoTaggingMediaTaskResponse), nil
	}
}
