package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbs/v1/model"
)

type CollectHotQuestionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectHotQuestionsInvoker) Invoke() (*model.CollectHotQuestionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectHotQuestionsResponse), nil
	}
}

type CollectKeyWordsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectKeyWordsInvoker) Invoke() (*model.CollectKeyWordsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectKeyWordsResponse), nil
	}
}

type CollectReplyRatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectReplyRatesInvoker) Invoke() (*model.CollectReplyRatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectReplyRatesResponse), nil
	}
}

type CollectSessionStatsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectSessionStatsInvoker) Invoke() (*model.CollectSessionStatsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectSessionStatsResponse), nil
	}
}

type CreateSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSessionInvoker) Invoke() (*model.CreateSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSessionResponse), nil
	}
}

type DeleteSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSessionInvoker) Invoke() (*model.DeleteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSessionResponse), nil
	}
}

type ExecuteQaChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteQaChatInvoker) Invoke() (*model.ExecuteQaChatResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteQaChatResponse), nil
	}
}

type ExecuteSessionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteSessionInvoker) Invoke() (*model.ExecuteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSessionResponse), nil
	}
}

type ListSuggestionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSuggestionsInvoker) Invoke() (*model.ListSuggestionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSuggestionsResponse), nil
	}
}

type TagLaborInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagLaborInvoker) Invoke() (*model.TagLaborResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagLaborResponse), nil
	}
}

type TagSatisfactionInvoker struct {
	*invoker.BaseInvoker
}

func (i *TagSatisfactionInvoker) Invoke() (*model.TagSatisfactionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagSatisfactionResponse), nil
	}
}
