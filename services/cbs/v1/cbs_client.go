package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbs/v1/model"
)

type CbsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCbsClient(hcClient *httpclient.HcHttpClient) *CbsClient {
	return &CbsClient{HcClient: hcClient}
}

func CbsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// CollectHotQuestions 热点问题统计
//
// 获取完全匹配的热点标准问题列表。
// 默认按照完全匹配标准问题被问及的频次降序排序。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) CollectHotQuestions(request *model.CollectHotQuestionsRequest) (*model.CollectHotQuestionsResponse, error) {
	requestDef := GenReqDefForCollectHotQuestions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectHotQuestionsResponse), nil
	}
}

// CollectHotQuestionsInvoker 热点问题统计
func (c *CbsClient) CollectHotQuestionsInvoker(request *model.CollectHotQuestionsRequest) *CollectHotQuestionsInvoker {
	requestDef := GenReqDefForCollectHotQuestions()
	return &CollectHotQuestionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectKeyWords 关键词统计
//
// 用户问关键词统计。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) CollectKeyWords(request *model.CollectKeyWordsRequest) (*model.CollectKeyWordsResponse, error) {
	requestDef := GenReqDefForCollectKeyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectKeyWordsResponse), nil
	}
}

// CollectKeyWordsInvoker 关键词统计
func (c *CbsClient) CollectKeyWordsInvoker(request *model.CollectKeyWordsRequest) *CollectKeyWordsInvoker {
	requestDef := GenReqDefForCollectKeyWords()
	return &CollectKeyWordsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectReplyRates 问答统计
//
// 指定领域获取指定时间范围内的问题答复率，支持按周期统计。
// 如果领域未指定则表示获取所有领域的问题答复率。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) CollectReplyRates(request *model.CollectReplyRatesRequest) (*model.CollectReplyRatesResponse, error) {
	requestDef := GenReqDefForCollectReplyRates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectReplyRatesResponse), nil
	}
}

// CollectReplyRatesInvoker 问答统计
func (c *CbsClient) CollectReplyRatesInvoker(request *model.CollectReplyRatesRequest) *CollectReplyRatesInvoker {
	requestDef := GenReqDefForCollectReplyRates()
	return &CollectReplyRatesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CollectSessionStats 访问统计
//
// 获取用户会话统计信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) CollectSessionStats(request *model.CollectSessionStatsRequest) (*model.CollectSessionStatsResponse, error) {
	requestDef := GenReqDefForCollectSessionStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectSessionStatsResponse), nil
	}
}

// CollectSessionStatsInvoker 访问统计
func (c *CbsClient) CollectSessionStatsInvoker(request *model.CollectSessionStatsRequest) *CollectSessionStatsInvoker {
	requestDef := GenReqDefForCollectSessionStats()
	return &CollectSessionStatsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateSession 开启会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口创建会话。该接口仅支持老用户，新用户请优先使用问答机器人API接口进行调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) CreateSession(request *model.CreateSessionRequest) (*model.CreateSessionResponse, error) {
	requestDef := GenReqDefForCreateSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSessionResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// CreateSessionInvoker 开启会话
func (c *CbsClient) CreateSessionInvoker(request *model.CreateSessionRequest) *CreateSessionInvoker {
	requestDef := GenReqDefForCreateSession()
	return &CreateSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteSession 关闭会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口关闭会话。该接口即将下线，请优先使用问答机器人API接口进行调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) DeleteSession(request *model.DeleteSessionRequest) (*model.DeleteSessionResponse, error) {
	requestDef := GenReqDefForDeleteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSessionResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// DeleteSessionInvoker 关闭会话
func (c *CbsClient) DeleteSessionInvoker(request *model.DeleteSessionRequest) *DeleteSessionInvoker {
	requestDef := GenReqDefForDeleteSession()
	return &DeleteSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteComposeVideo 合成视频(按包周期收费)
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteComposeVideo(request *model.ExecuteComposeVideoRequest) (*model.ExecuteComposeVideoResponse, error) {
	requestDef := GenReqDefForExecuteComposeVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteComposeVideoResponse), nil
	}
}

// ExecuteComposeVideoInvoker 合成视频(按包周期收费)
func (c *CbsClient) ExecuteComposeVideoInvoker(request *model.ExecuteComposeVideoRequest) *ExecuteComposeVideoInvoker {
	requestDef := GenReqDefForExecuteComposeVideo()
	return &ExecuteComposeVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteComposeVideoOndemand 合成视频(按需收费)
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteComposeVideoOndemand(request *model.ExecuteComposeVideoOndemandRequest) (*model.ExecuteComposeVideoOndemandResponse, error) {
	requestDef := GenReqDefForExecuteComposeVideoOndemand()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteComposeVideoOndemandResponse), nil
	}
}

// ExecuteComposeVideoOndemandInvoker 合成视频(按需收费)
func (c *CbsClient) ExecuteComposeVideoOndemandInvoker(request *model.ExecuteComposeVideoOndemandRequest) *ExecuteComposeVideoOndemandInvoker {
	requestDef := GenReqDefForExecuteComposeVideoOndemand()
	return &ExecuteComposeVideoOndemandInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteCreateVideo 创建视频
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteCreateVideo(request *model.ExecuteCreateVideoRequest) (*model.ExecuteCreateVideoResponse, error) {
	requestDef := GenReqDefForExecuteCreateVideo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteCreateVideoResponse), nil
	}
}

// ExecuteCreateVideoInvoker 创建视频
func (c *CbsClient) ExecuteCreateVideoInvoker(request *model.ExecuteCreateVideoRequest) *ExecuteCreateVideoInvoker {
	requestDef := GenReqDefForExecuteCreateVideo()
	return &ExecuteCreateVideoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDeleteVideoById 删除视频
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteDeleteVideoById(request *model.ExecuteDeleteVideoByIdRequest) (*model.ExecuteDeleteVideoByIdResponse, error) {
	requestDef := GenReqDefForExecuteDeleteVideoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeleteVideoByIdResponse), nil
	}
}

// ExecuteDeleteVideoByIdInvoker 删除视频
func (c *CbsClient) ExecuteDeleteVideoByIdInvoker(request *model.ExecuteDeleteVideoByIdRequest) *ExecuteDeleteVideoByIdInvoker {
	requestDef := GenReqDefForExecuteDeleteVideoById()
	return &ExecuteDeleteVideoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteDeleteimageById 删除图片
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteDeleteimageById(request *model.ExecuteDeleteimageByIdRequest) (*model.ExecuteDeleteimageByIdResponse, error) {
	requestDef := GenReqDefForExecuteDeleteimageById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteDeleteimageByIdResponse), nil
	}
}

// ExecuteDeleteimageByIdInvoker 删除图片
func (c *CbsClient) ExecuteDeleteimageByIdInvoker(request *model.ExecuteDeleteimageByIdRequest) *ExecuteDeleteimageByIdInvoker {
	requestDef := GenReqDefForExecuteDeleteimageById()
	return &ExecuteDeleteimageByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetCharacterInfoById 获取形象详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetCharacterInfoById(request *model.ExecuteGetCharacterInfoByIdRequest) (*model.ExecuteGetCharacterInfoByIdResponse, error) {
	requestDef := GenReqDefForExecuteGetCharacterInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetCharacterInfoByIdResponse), nil
	}
}

// ExecuteGetCharacterInfoByIdInvoker 获取形象详情
func (c *CbsClient) ExecuteGetCharacterInfoByIdInvoker(request *model.ExecuteGetCharacterInfoByIdRequest) *ExecuteGetCharacterInfoByIdInvoker {
	requestDef := GenReqDefForExecuteGetCharacterInfoById()
	return &ExecuteGetCharacterInfoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetCharacters 获取形象列表
//
// TODO:
//
// 本期不做形象进度
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetCharacters(request *model.ExecuteGetCharactersRequest) (*model.ExecuteGetCharactersResponse, error) {
	requestDef := GenReqDefForExecuteGetCharacters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetCharactersResponse), nil
	}
}

// ExecuteGetCharactersInvoker 获取形象列表
func (c *CbsClient) ExecuteGetCharactersInvoker(request *model.ExecuteGetCharactersRequest) *ExecuteGetCharactersInvoker {
	requestDef := GenReqDefForExecuteGetCharacters()
	return &ExecuteGetCharactersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetFramsListByImagesId 获取播报框
//
// 获取指定图片可用的播报框列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetFramsListByImagesId(request *model.ExecuteGetFramsListByImagesIdRequest) (*model.ExecuteGetFramsListByImagesIdResponse, error) {
	requestDef := GenReqDefForExecuteGetFramsListByImagesId()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetFramsListByImagesIdResponse), nil
	}
}

// ExecuteGetFramsListByImagesIdInvoker 获取播报框
func (c *CbsClient) ExecuteGetFramsListByImagesIdInvoker(request *model.ExecuteGetFramsListByImagesIdRequest) *ExecuteGetFramsListByImagesIdInvoker {
	requestDef := GenReqDefForExecuteGetFramsListByImagesId()
	return &ExecuteGetFramsListByImagesIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetImagesList 获取图片列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetImagesList(request *model.ExecuteGetImagesListRequest) (*model.ExecuteGetImagesListResponse, error) {
	requestDef := GenReqDefForExecuteGetImagesList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetImagesListResponse), nil
	}
}

// ExecuteGetImagesListInvoker 获取图片列表
func (c *CbsClient) ExecuteGetImagesListInvoker(request *model.ExecuteGetImagesListRequest) *ExecuteGetImagesListInvoker {
	requestDef := GenReqDefForExecuteGetImagesList()
	return &ExecuteGetImagesListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetVideoInfoById 获取视频详情
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetVideoInfoById(request *model.ExecuteGetVideoInfoByIdRequest) (*model.ExecuteGetVideoInfoByIdResponse, error) {
	requestDef := GenReqDefForExecuteGetVideoInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetVideoInfoByIdResponse), nil
	}
}

// ExecuteGetVideoInfoByIdInvoker 获取视频详情
func (c *CbsClient) ExecuteGetVideoInfoByIdInvoker(request *model.ExecuteGetVideoInfoByIdRequest) *ExecuteGetVideoInfoByIdInvoker {
	requestDef := GenReqDefForExecuteGetVideoInfoById()
	return &ExecuteGetVideoInfoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteGetVideosList 获取视频列表
//
// 该接口用于获取视频列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteGetVideosList(request *model.ExecuteGetVideosListRequest) (*model.ExecuteGetVideosListResponse, error) {
	requestDef := GenReqDefForExecuteGetVideosList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteGetVideosListResponse), nil
	}
}

// ExecuteGetVideosListInvoker 获取视频列表
func (c *CbsClient) ExecuteGetVideosListInvoker(request *model.ExecuteGetVideosListRequest) *ExecuteGetVideosListInvoker {
	requestDef := GenReqDefForExecuteGetVideosList()
	return &ExecuteGetVideosListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecutePostCreateImages 创建图片
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecutePostCreateImages(request *model.ExecutePostCreateImagesRequest) (*model.ExecutePostCreateImagesResponse, error) {
	requestDef := GenReqDefForExecutePostCreateImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecutePostCreateImagesResponse), nil
	}
}

// ExecutePostCreateImagesInvoker 创建图片
func (c *CbsClient) ExecutePostCreateImagesInvoker(request *model.ExecutePostCreateImagesRequest) *ExecutePostCreateImagesInvoker {
	requestDef := GenReqDefForExecutePostCreateImages()
	return &ExecutePostCreateImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteQaChat 问答机器人会话
//
// 用户调用该接口和机器人进行聊天。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteQaChat(request *model.ExecuteQaChatRequest) (*model.ExecuteQaChatResponse, error) {
	requestDef := GenReqDefForExecuteQaChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteQaChatResponse), nil
	}
}

// ExecuteQaChatInvoker 问答机器人会话
func (c *CbsClient) ExecuteQaChatInvoker(request *model.ExecuteQaChatRequest) *ExecuteQaChatInvoker {
	requestDef := GenReqDefForExecuteQaChat()
	return &ExecuteQaChatInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ExecuteSession 处理会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口与机器人进行会话。该接口即将下线，请优先使用问答机器人API接口进行调用。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteSession(request *model.ExecuteSessionRequest) (*model.ExecuteSessionResponse, error) {
	requestDef := GenReqDefForExecuteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSessionResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// ExecuteSessionInvoker 处理会话
func (c *CbsClient) ExecuteSessionInvoker(request *model.ExecuteSessionRequest) *ExecuteSessionInvoker {
	requestDef := GenReqDefForExecuteSession()
	return &ExecuteSessionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteUpdateImageName 修改图片名
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteUpdateImageName(request *model.ExecuteUpdateImageNameRequest) (*model.ExecuteUpdateImageNameResponse, error) {
	requestDef := GenReqDefForExecuteUpdateImageName()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUpdateImageNameResponse), nil
	}
}

// ExecuteUpdateImageNameInvoker 修改图片名
func (c *CbsClient) ExecuteUpdateImageNameInvoker(request *model.ExecuteUpdateImageNameRequest) *ExecuteUpdateImageNameInvoker {
	requestDef := GenReqDefForExecuteUpdateImageName()
	return &ExecuteUpdateImageNameInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteUpdateVideoById 更新视频名
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteUpdateVideoById(request *model.ExecuteUpdateVideoByIdRequest) (*model.ExecuteUpdateVideoByIdResponse, error) {
	requestDef := GenReqDefForExecuteUpdateVideoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUpdateVideoByIdResponse), nil
	}
}

// ExecuteUpdateVideoByIdInvoker 更新视频名
func (c *CbsClient) ExecuteUpdateVideoByIdInvoker(request *model.ExecuteUpdateVideoByIdRequest) *ExecuteUpdateVideoByIdInvoker {
	requestDef := GenReqDefForExecuteUpdateVideoById()
	return &ExecuteUpdateVideoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteUpdateVideoInfoById 配置视频
//
// 通过该接口配置视频
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteUpdateVideoInfoById(request *model.ExecuteUpdateVideoInfoByIdRequest) (*model.ExecuteUpdateVideoInfoByIdResponse, error) {
	requestDef := GenReqDefForExecuteUpdateVideoInfoById()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUpdateVideoInfoByIdResponse), nil
	}
}

// ExecuteUpdateVideoInfoByIdInvoker 配置视频
func (c *CbsClient) ExecuteUpdateVideoInfoByIdInvoker(request *model.ExecuteUpdateVideoInfoByIdRequest) *ExecuteUpdateVideoInfoByIdInvoker {
	requestDef := GenReqDefForExecuteUpdateVideoInfoById()
	return &ExecuteUpdateVideoInfoByIdInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteUploadImage 上传播报插图
//
// 上传图片并生成图片链接，图片需小于10m；
// 同一个视频同时最多支持50张插图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteUploadImage(request *model.ExecuteUploadImageRequest) (*model.ExecuteUploadImageResponse, error) {
	requestDef := GenReqDefForExecuteUploadImage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUploadImageResponse), nil
	}
}

// ExecuteUploadImageInvoker 上传播报插图
func (c *CbsClient) ExecuteUploadImageInvoker(request *model.ExecuteUploadImageRequest) *ExecuteUploadImageInvoker {
	requestDef := GenReqDefForExecuteUploadImage()
	return &ExecuteUploadImageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteUploadPpt 通过pdf上传多张插图
//
// 当前仅支持上传PDF，如有PPT请将PPT转化为PDF再进行上传，文件需小于10m；
// 该接口会将pdf每一页转换图片，并生成链接；
// 同一个视频同时最多支持50张插图。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ExecuteUploadPpt(request *model.ExecuteUploadPptRequest) (*model.ExecuteUploadPptResponse, error) {
	requestDef := GenReqDefForExecuteUploadPpt()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteUploadPptResponse), nil
	}
}

// ExecuteUploadPptInvoker 通过pdf上传多张插图
func (c *CbsClient) ExecuteUploadPptInvoker(request *model.ExecuteUploadPptRequest) *ExecuteUploadPptInvoker {
	requestDef := GenReqDefForExecuteUploadPpt()
	return &ExecuteUploadPptInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSuggestions 获取问题提示
//
// 获取用户输入问题的提示问题列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) ListSuggestions(request *model.ListSuggestionsRequest) (*model.ListSuggestionsResponse, error) {
	requestDef := GenReqDefForListSuggestions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSuggestionsResponse), nil
	}
}

// ListSuggestionsInvoker 获取问题提示
func (c *CbsClient) ListSuggestionsInvoker(request *model.ListSuggestionsRequest) *ListSuggestionsInvoker {
	requestDef := GenReqDefForListSuggestions()
	return &ListSuggestionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagLabor 标记为转人工
//
// 智能问答返回的结果后，用户是否转人工。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) TagLabor(request *model.TagLaborRequest) (*model.TagLaborResponse, error) {
	requestDef := GenReqDefForTagLabor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagLaborResponse), nil
	}
}

// TagLaborInvoker 标记为转人工
func (c *CbsClient) TagLaborInvoker(request *model.TagLaborRequest) *TagLaborInvoker {
	requestDef := GenReqDefForTagLabor()
	return &TagLaborInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagSatisfaction 问答满意评价
//
// 用户提出问题后，对智能问答返回的结果是否满意。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) TagSatisfaction(request *model.TagSatisfactionRequest) (*model.TagSatisfactionResponse, error) {
	requestDef := GenReqDefForTagSatisfaction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagSatisfactionResponse), nil
	}
}

// TagSatisfactionInvoker 问答满意评价
func (c *CbsClient) TagSatisfactionInvoker(request *model.TagSatisfactionRequest) *TagSatisfactionInvoker {
	requestDef := GenReqDefForTagSatisfaction()
	return &TagSatisfactionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PostRequests PostRequests
//
// 问答服务的输入为用户提问，输出是与输入最匹配的Top N(默认为top5)个知识点，知识点按得分从高到低排序。
//
// 说明：
// 返回知识点如果含有答案字段（answer），则表示返回匹配成功结果，如果没有答案字段，则表示推荐结果。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CbsClient) PostRequests(request *model.PostRequestsRequest) (*model.PostRequestsResponse, error) {
	requestDef := GenReqDefForPostRequests()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PostRequestsResponse), nil
	}
}

// Deprecated: This function is deprecated and will be removed in the future versions.
// PostRequestsInvoker PostRequests
func (c *CbsClient) PostRequestsInvoker(request *model.PostRequestsRequest) *PostRequestsInvoker {
	requestDef := GenReqDefForPostRequests()
	return &PostRequestsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
