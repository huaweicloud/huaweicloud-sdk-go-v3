package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbs/v1/model"
)

type CollectHotQuestionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CollectHotQuestionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CollectKeyWordsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CollectReplyRatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CollectSessionStatsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *CreateSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *DeleteSessionInvoker) Invoke() (*model.DeleteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSessionResponse), nil
	}
}

type ExecuteComposeVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteComposeVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteComposeVideoInvoker) Invoke() (*model.ExecuteComposeVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteComposeVideoResponse), nil
	}
}

type ExecuteComposeVideoOndemandInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteComposeVideoOndemandInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteComposeVideoOndemandInvoker) Invoke() (*model.ExecuteComposeVideoOndemandResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteComposeVideoOndemandResponse), nil
	}
}

type ExecuteCreateVideoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteCreateVideoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteCreateVideoInvoker) Invoke() (*model.ExecuteCreateVideoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteCreateVideoResponse), nil
	}
}

type ExecuteDeleteVideoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDeleteVideoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteDeleteVideoByIdInvoker) Invoke() (*model.ExecuteDeleteVideoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDeleteVideoByIdResponse), nil
	}
}

type ExecuteDeleteimageByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDeleteimageByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteDeleteimageByIdInvoker) Invoke() (*model.ExecuteDeleteimageByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDeleteimageByIdResponse), nil
	}
}

type ExecuteGetCharacterInfoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetCharacterInfoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetCharacterInfoByIdInvoker) Invoke() (*model.ExecuteGetCharacterInfoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetCharacterInfoByIdResponse), nil
	}
}

type ExecuteGetCharactersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetCharactersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetCharactersInvoker) Invoke() (*model.ExecuteGetCharactersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetCharactersResponse), nil
	}
}

type ExecuteGetFramsListByImagesIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetFramsListByImagesIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetFramsListByImagesIdInvoker) Invoke() (*model.ExecuteGetFramsListByImagesIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetFramsListByImagesIdResponse), nil
	}
}

type ExecuteGetImagesListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetImagesListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetImagesListInvoker) Invoke() (*model.ExecuteGetImagesListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetImagesListResponse), nil
	}
}

type ExecuteGetVideoInfoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetVideoInfoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetVideoInfoByIdInvoker) Invoke() (*model.ExecuteGetVideoInfoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetVideoInfoByIdResponse), nil
	}
}

type ExecuteGetVideosListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteGetVideosListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteGetVideosListInvoker) Invoke() (*model.ExecuteGetVideosListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteGetVideosListResponse), nil
	}
}

type ExecutePostCreateImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecutePostCreateImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecutePostCreateImagesInvoker) Invoke() (*model.ExecutePostCreateImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecutePostCreateImagesResponse), nil
	}
}

type ExecuteQaChatInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteQaChatInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ExecuteSessionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ExecuteSessionInvoker) Invoke() (*model.ExecuteSessionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteSessionResponse), nil
	}
}

type ExecuteUpdateImageNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUpdateImageNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteUpdateImageNameInvoker) Invoke() (*model.ExecuteUpdateImageNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUpdateImageNameResponse), nil
	}
}

type ExecuteUpdateVideoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUpdateVideoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteUpdateVideoByIdInvoker) Invoke() (*model.ExecuteUpdateVideoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUpdateVideoByIdResponse), nil
	}
}

type ExecuteUpdateVideoInfoByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUpdateVideoInfoByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteUpdateVideoInfoByIdInvoker) Invoke() (*model.ExecuteUpdateVideoInfoByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUpdateVideoInfoByIdResponse), nil
	}
}

type ExecuteUploadImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUploadImageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteUploadImageInvoker) Invoke() (*model.ExecuteUploadImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUploadImageResponse), nil
	}
}

type ExecuteUploadPptInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteUploadPptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteUploadPptInvoker) Invoke() (*model.ExecuteUploadPptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteUploadPptResponse), nil
	}
}

type ListSuggestionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSuggestionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *TagLaborInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *TagSatisfactionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *TagSatisfactionInvoker) Invoke() (*model.TagSatisfactionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.TagSatisfactionResponse), nil
	}
}

type PostRequestsInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *PostRequestsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *PostRequestsInvoker) Invoke() (*model.PostRequestsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostRequestsResponse), nil
	}
}
