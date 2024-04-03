package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/pangulargemodels/v2/model"
)

type PanguLargeModelsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewPanguLargeModelsClient(hcClient *httpclient.HcHttpClient) *PanguLargeModelsClient {
	return &PanguLargeModelsClient{HcClient: hcClient}
}

func PanguLargeModelsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ExecuteChatCompletion 对话问答
//
// 基于对话问答功能，用户可以与模型进行自然而流畅的对话和交流。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *PanguLargeModelsClient) ExecuteChatCompletion(request *model.ExecuteChatCompletionRequest) (*model.ExecuteChatCompletionResponse, error) {
	requestDef := GenReqDefForExecuteChatCompletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteChatCompletionResponse), nil
	}
}

// ExecuteChatCompletionInvoker 对话问答
func (c *PanguLargeModelsClient) ExecuteChatCompletionInvoker(request *model.ExecuteChatCompletionRequest) *ExecuteChatCompletionInvoker {
	requestDef := GenReqDefForExecuteChatCompletion()
	return &ExecuteChatCompletionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExecuteTextCompletion 通用文本
//
// 给定一个提示和一些参数，模型会根据这些信息生成一个或多个预测的补全，还可以返回每个位置上不同词语的概率。它可以用来做文本生成、自动写作、代码补全等任务。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *PanguLargeModelsClient) ExecuteTextCompletion(request *model.ExecuteTextCompletionRequest) (*model.ExecuteTextCompletionResponse, error) {
	requestDef := GenReqDefForExecuteTextCompletion()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExecuteTextCompletionResponse), nil
	}
}

// ExecuteTextCompletionInvoker 通用文本
func (c *PanguLargeModelsClient) ExecuteTextCompletionInvoker(request *model.ExecuteTextCompletionRequest) *ExecuteTextCompletionInvoker {
	requestDef := GenReqDefForExecuteTextCompletion()
	return &ExecuteTextCompletionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
