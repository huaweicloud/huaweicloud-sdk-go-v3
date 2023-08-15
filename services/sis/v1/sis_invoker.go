package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/sis/v1/model"
)

type CollectTranscriberJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectTranscriberJobInvoker) Invoke() (*model.CollectTranscriberJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectTranscriberJobResponse), nil
	}
}

type CreateVocabularyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVocabularyInvoker) Invoke() (*model.CreateVocabularyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVocabularyResponse), nil
	}
}

type DeleteVocabularyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVocabularyInvoker) Invoke() (*model.DeleteVocabularyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVocabularyResponse), nil
	}
}

type PushTranscriberJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *PushTranscriberJobsInvoker) Invoke() (*model.PushTranscriberJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PushTranscriberJobsResponse), nil
	}
}

type RecognizeFlashAsrInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeFlashAsrInvoker) Invoke() (*model.RecognizeFlashAsrResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeFlashAsrResponse), nil
	}
}

type RecognizeShortAudioInvoker struct {
	*invoker.BaseInvoker
}

func (i *RecognizeShortAudioInvoker) Invoke() (*model.RecognizeShortAudioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecognizeShortAudioResponse), nil
	}
}

type RunAudioAssessmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAudioAssessmentInvoker) Invoke() (*model.RunAudioAssessmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAudioAssessmentResponse), nil
	}
}

type RunMultiModalAssessmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunMultiModalAssessmentInvoker) Invoke() (*model.RunMultiModalAssessmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunMultiModalAssessmentResponse), nil
	}
}

type RunTtsInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTtsInvoker) Invoke() (*model.RunTtsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTtsResponse), nil
	}
}

type ShowVocabulariesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVocabulariesInvoker) Invoke() (*model.ShowVocabulariesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVocabulariesResponse), nil
	}
}

type ShowVocabularyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVocabularyInvoker) Invoke() (*model.ShowVocabularyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVocabularyResponse), nil
	}
}

type UpdateVocabularyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVocabularyInvoker) Invoke() (*model.UpdateVocabularyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVocabularyResponse), nil
	}
}
