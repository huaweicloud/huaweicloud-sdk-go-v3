package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/nlp/v2/model"
)

type NlpClient struct {
	HcClient *http_client.HcHttpClient
}

func NewNlpClient(hcClient *http_client.HcHttpClient) *NlpClient {
	return &NlpClient{HcClient: hcClient}
}

func NlpClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//属性级情感分析，针对手机领域的用户评论进行属性级情感分析。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunAspectSentiment(request *model.RunAspectSentimentRequest) (*model.RunAspectSentimentResponse, error) {
	requestDef := GenReqDefForRunAspectSentiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAspectSentimentResponse), nil
	}
}

//属性级情感分析（高级版），针对手机、汽车领域的用户评论进行属性级情感分析。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunAspectSentimentAdvance(request *model.RunAspectSentimentAdvanceRequest) (*model.RunAspectSentimentAdvanceResponse, error) {
	requestDef := GenReqDefForRunAspectSentimentAdvance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAspectSentimentAdvanceResponse), nil
	}
}

//针对广告领域的自动分类，判断是否是广告。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunClassification(request *model.RunClassificationRequest) (*model.RunClassificationResponse, error) {
	requestDef := GenReqDefForRunClassification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunClassificationResponse), nil
	}
}

//识别句子中词汇与词汇之间的相互依存关系。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunDependencyParser(request *model.RunDependencyParserRequest) (*model.RunDependencyParserResponse, error) {
	requestDef := GenReqDefForRunDependencyParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDependencyParserResponse), nil
	}
}

//文档分类接口，输入文档内容，输出文档的标签和置信度，支持多个标签。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunDocClassification(request *model.RunDocClassificationRequest) (*model.RunDocClassificationResponse, error) {
	requestDef := GenReqDefForRunDocClassification()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDocClassificationResponse), nil
	}
}

//领域情感分析，针对未知领域，电商，汽车领域的用户评论进行情感分析。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunDomainSentiment(request *model.RunDomainSentimentRequest) (*model.RunDomainSentimentResponse, error) {
	requestDef := GenReqDefForRunDomainSentiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunDomainSentimentResponse), nil
	}
}

//针对通用领域的文本进行实体链接分析，识别出其中的实体，并返回实体相关信息。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunEntityLinking(request *model.RunEntityLinkingRequest) (*model.RunEntityLinkingResponse, error) {
	requestDef := GenReqDefForRunEntityLinking()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEntityLinkingResponse), nil
	}
}

//实体级情感分析，本产品适用于金融方面公司实体正负面新闻的分析。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunEntitySentiment(request *model.RunEntitySentimentRequest) (*model.RunEntitySentimentResponse, error) {
	requestDef := GenReqDefForRunEntitySentiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEntitySentimentResponse), nil
	}
}

//事件抽取是指从自然语言文本中抽取指定类型的事件以及相关实体信息，并形成结构化数据输出的文本处理技术。 目前只支持金融公告中会议召开、聘任、辞职、股票增持、股票减持5类事件以及相关要素的抽取。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunEventExtraction(request *model.RunEventExtractionRequest) (*model.RunEventExtractionResponse, error) {
	requestDef := GenReqDefForRunEventExtraction()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunEventExtractionResponse), nil
	}
}

//文档翻译接口，用于翻译文档格式文件。由于文档翻译会需要较长的时间，因此识别是异步的，也即接口分为创建翻译任务和查询任务状态两个接口。创建翻译任务接口创建任务完成后返回，然后用户通过调用查询任务状态接口来获得翻译状态和临时URL。 用户可以使用临时URL下载翻译好的文件，每个临时URL有效期为10分种。翻译结果会保存24小时（从翻译完成的时间算起）。24小时后如果再访问，将会返回 \\\"task id is not found\\\"错误。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunFileTranslation(request *model.RunFileTranslationRequest) (*model.RunFileTranslationResponse, error) {
	requestDef := GenReqDefForRunFileTranslation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunFileTranslationResponse), nil
	}
}

//该接口用于获取文档翻译识别状态以及临时url，临时url可以用与获取翻译后的文档，每个临时url有效期为十分钟。
func (c *NlpClient) RunGetFileTranslationResult(request *model.RunGetFileTranslationResultRequest) (*model.RunGetFileTranslationResultResponse, error) {
	requestDef := GenReqDefForRunGetFileTranslationResult()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunGetFileTranslationResultResponse), nil
	}
}

//给定一段文本，抽取其中最能够反映文本主题或者意思的词汇。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunKeywordExtract(request *model.RunKeywordExtractRequest) (*model.RunKeywordExtractResponse, error) {
	requestDef := GenReqDefForRunKeywordExtract()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunKeywordExtractResponse), nil
	}
}

//对于用户输入的文本，返回识别出的所属语种。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunLanguageDetection(request *model.RunLanguageDetectionRequest) (*model.RunLanguageDetectionResponse, error) {
	requestDef := GenReqDefForRunLanguageDetection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunLanguageDetectionResponse), nil
	}
}

//多粒度分词：给定一个句子输入，输出不同粒度的所有单词的层次结构。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunMultiGrainedSegment(request *model.RunMultiGrainedSegmentRequest) (*model.RunMultiGrainedSegmentResponse, error) {
	requestDef := GenReqDefForRunMultiGrainedSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunMultiGrainedSegmentResponse), nil
	}
}

//基础版命名实体识别，对文本进行命名实体识别分析，目前支持人名、地名、时间、组织机构类实体的识别。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunNer(request *model.RunNerRequest) (*model.RunNerResponse, error) {
	requestDef := GenReqDefForRunNer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunNerResponse), nil
	}
}

//领域版本命名实体识别，对文本进行命名实体识别分析，目前支持人名、地名、组织机构、时间点、日期、百分比、货币额度、序数词、计量规格词、民族、职业、邮箱12类实体的识别。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunNerDomain(request *model.RunNerDomainRequest) (*model.RunNerDomainResponse, error) {
	requestDef := GenReqDefForRunNerDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunNerDomainResponse), nil
	}
}

//根据用户的输入生成诗歌。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunPoem(request *model.RunPoemRequest) (*model.RunPoemResponse, error) {
	requestDef := GenReqDefForRunPoem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunPoemResponse), nil
	}
}

//对文本进行分词和词性标注处理。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSegment(request *model.RunSegmentRequest) (*model.RunSegmentResponse, error) {
	requestDef := GenReqDefForRunSegment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSegmentResponse), nil
	}
}

//针对天气、报时、新闻、笑话、翻译、提醒、闹钟、音乐8个领域进行意图理解，对用户的问题进行领域识别并提取领域内的参数。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSemanticParser(request *model.RunSemanticParserRequest) (*model.RunSemanticParserResponse, error) {
	requestDef := GenReqDefForRunSemanticParser()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSemanticParserResponse), nil
	}
}

//输入句子，返回对应的句向量。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSentenceEmbedding(request *model.RunSentenceEmbeddingRequest) (*model.RunSentenceEmbeddingResponse, error) {
	requestDef := GenReqDefForRunSentenceEmbedding()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSentenceEmbeddingResponse), nil
	}
}

//通用情感分析，针对通用领域的用户评论进行情感分析。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSentiment(request *model.RunSentimentRequest) (*model.RunSentimentResponse, error) {
	requestDef := GenReqDefForRunSentiment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSentimentResponse), nil
	}
}

//对文本生成摘要。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSummary(request *model.RunSummaryRequest) (*model.RunSummaryResponse, error) {
	requestDef := GenReqDefForRunSummary()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSummaryResponse), nil
	}
}

//对文本生成摘要。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunSummaryDomain(request *model.RunSummaryDomainRequest) (*model.RunSummaryDomainResponse, error) {
	requestDef := GenReqDefForRunSummaryDomain()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunSummaryDomainResponse), nil
	}
}

//文本相似度服务，对文本对进行相似度计算。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunTextSimilarity(request *model.RunTextSimilarityRequest) (*model.RunTextSimilarityResponse, error) {
	requestDef := GenReqDefForRunTextSimilarity()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTextSimilarityResponse), nil
	}
}

//文本相似度服务高级版，对文本对进行相似度计算。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunTextSimilarityAdvance(request *model.RunTextSimilarityAdvanceRequest) (*model.RunTextSimilarityAdvanceResponse, error) {
	requestDef := GenReqDefForRunTextSimilarityAdvance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTextSimilarityAdvanceResponse), nil
	}
}

//对于用户输入原始语种的文本，转换为目标语种的文本。 在使用本API之前， 需要您完成服务申请， 具体操作流程请参见[申请服务](https://support.huaweicloud.com/api-nlp/nlp_03_0004.html)章节。
func (c *NlpClient) RunTextTranslation(request *model.RunTextTranslationRequest) (*model.RunTextTranslationResponse, error) {
	requestDef := GenReqDefForRunTextTranslation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunTextTranslationResponse), nil
	}
}
