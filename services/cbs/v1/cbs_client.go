package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbs/v1/model"
)

type CbsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCbsClient(hcClient *http_client.HcHttpClient) *CbsClient {
	return &CbsClient{HcClient: hcClient}
}

func CbsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 热点问题统计
//
// 获取完全匹配的热点标准问题列表。
// 默认按照完全匹配标准问题被问及的频次降序排序。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CollectHotQuestions(request *model.CollectHotQuestionsRequest) (*model.CollectHotQuestionsResponse, error) {
	requestDef := GenReqDefForCollectHotQuestions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectHotQuestionsResponse), nil
	}
}

// 关键词统计
//
// 用户问关键词统计。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CollectKeyWords(request *model.CollectKeyWordsRequest) (*model.CollectKeyWordsResponse, error) {
	requestDef := GenReqDefForCollectKeyWords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectKeyWordsResponse), nil
	}
}

// 问答统计
//
// 指定领域获取指定时间范围内的问题答复率，支持按周期统计。
// 如果领域未指定则表示获取所有领域的问题答复率。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CollectReplyRates(request *model.CollectReplyRatesRequest) (*model.CollectReplyRatesResponse, error) {
	requestDef := GenReqDefForCollectReplyRates()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectReplyRatesResponse), nil
	}
}

// 访问统计
//
// 获取用户会话统计信息。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CollectSessionStats(request *model.CollectSessionStatsRequest) (*model.CollectSessionStatsResponse, error) {
	requestDef := GenReqDefForCollectSessionStats()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CollectSessionStatsResponse), nil
	}
}

// 开启会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口创建会话。该接口仅支持老用户，新用户请优先使用问答机器人API接口进行调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CreateSession(request *model.CreateSessionRequest) (*model.CreateSessionResponse, error) {
	requestDef := GenReqDefForCreateSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSessionResponse), nil
	}
}

// 发起会话
//
// 发起话务机器人会话。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) CreateTbSession(request *model.CreateTbSessionRequest) (*model.CreateTbSessionResponse, error) {
	requestDef := GenReqDefForCreateTbSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTbSessionResponse), nil
	}
}

// 关闭会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口关闭会话。该接口即将下线，请优先使用问答机器人API接口进行调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) DeleteSession(request *model.DeleteSessionRequest) (*model.DeleteSessionResponse, error) {
	requestDef := GenReqDefForDeleteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSessionResponse), nil
	}
}

// 结束会话
//
// 结束话务机器人会话。如果会话持续10分钟无会话请求则被清理。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) DeleteTbSession(request *model.DeleteTbSessionRequest) (*model.DeleteTbSessionResponse, error) {
	requestDef := GenReqDefForDeleteTbSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTbSessionResponse), nil
	}
}

// 问答机器人会话
//
// 用户调用该接口和机器人进行聊天。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) ExecuteQaChat(request *model.ExecuteQaChatRequest) (*model.ExecuteQaChatResponse, error) {
	requestDef := GenReqDefForExecuteQaChat()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteQaChatResponse), nil
	}
}

// 处理会话
//
// 问答会话API由开启会话、处理会话、关闭会话三个接口组成。用户可通过调用该接口与机器人进行会话。该接口即将下线，请优先使用问答机器人API接口进行调用。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) ExecuteSession(request *model.ExecuteSessionRequest) (*model.ExecuteSessionResponse, error) {
	requestDef := GenReqDefForExecuteSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteSessionResponse), nil
	}
}

// 进行会话
//
// 进行话务机器人会话。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) ExecuteTbSession(request *model.ExecuteTbSessionRequest) (*model.ExecuteTbSessionResponse, error) {
	requestDef := GenReqDefForExecuteTbSession()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTbSessionResponse), nil
	}
}

// 获取问题提示
//
// 获取用户输入问题的提示问题列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) ListSuggestions(request *model.ListSuggestionsRequest) (*model.ListSuggestionsResponse, error) {
	requestDef := GenReqDefForListSuggestions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSuggestionsResponse), nil
	}
}

// 标记为转人工
//
// 智能问答返回的结果后，用户是否转人工。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) TagLabor(request *model.TagLaborRequest) (*model.TagLaborResponse, error) {
	requestDef := GenReqDefForTagLabor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagLaborResponse), nil
	}
}

// 问答满意评价
//
// 用户提出问题后，对智能问答返回的结果是否满意。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CbsClient) TagSatisfaction(request *model.TagSatisfactionRequest) (*model.TagSatisfactionResponse, error) {
	requestDef := GenReqDefForTagSatisfaction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagSatisfactionResponse), nil
	}
}
