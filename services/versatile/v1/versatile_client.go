package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/versatile/v1/model"
)

type VersatileClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewVersatileClient(hcClient *httpclient.HcHttpClient) *VersatileClient {
	return &VersatileClient{HcClient: hcClient}
}

func VersatileClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// SearchKnowledgeBase 知识库检索
//
// 提供多知识库并行检索能力，支持语义、关键词、混合及FAQ四种检索模式，并允许自定义相似度阈值与返回结果数量，实现精准高效的信息匹配。
//
// **适用场景**：
// - 同时从多个知识库或文档集合中搜索相关内容。
// - 在预设的问答列表中快速定位最相关的答案（FAQ检索）。
// - 通过混合模式或调整阈值，兼顾搜索结果的准确性和全面性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VersatileClient) SearchKnowledgeBase(request *model.SearchKnowledgeBaseRequest) (*model.SearchKnowledgeBaseResponse, error) {
	requestDef := GenReqDefForSearchKnowledgeBase()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SearchKnowledgeBaseResponse), nil
	}
}

// SearchKnowledgeBaseInvoker 知识库检索
func (c *VersatileClient) SearchKnowledgeBaseInvoker(request *model.SearchKnowledgeBaseRequest) *SearchKnowledgeBaseInvoker {
	requestDef := GenReqDefForSearchKnowledgeBase()
	return &SearchKnowledgeBaseInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunAgent 运行一个智能体
//
// 运行一个智能体，支持流式和非流式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VersatileClient) RunAgent(request *model.RunAgentRequest) (*model.RunAgentResponse, error) {
	requestDef := GenReqDefForRunAgent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunAgentResponse), nil
	}
}

// RunAgentInvoker 运行一个智能体
func (c *VersatileClient) RunAgentInvoker(request *model.RunAgentRequest) *RunAgentInvoker {
	requestDef := GenReqDefForRunAgent()
	return &RunAgentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RunWorkflow 运行一个工作流
//
// 运行一个工作流，支持流式和非流式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VersatileClient) RunWorkflow(request *model.RunWorkflowRequest) (*model.RunWorkflowResponse, error) {
	requestDef := GenReqDefForRunWorkflow()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RunWorkflowResponse), nil
	}
}

// RunWorkflowInvoker 运行一个工作流
func (c *VersatileClient) RunWorkflowInvoker(request *model.RunWorkflowRequest) *RunWorkflowInvoker {
	requestDef := GenReqDefForRunWorkflow()
	return &RunWorkflowInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UploadAgentFile 上传文件
//
// 上传文件，以供agent或者工作流使用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *VersatileClient) UploadAgentFile(request *model.UploadAgentFileRequest) (*model.UploadAgentFileResponse, error) {
	requestDef := GenReqDefForUploadAgentFile()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UploadAgentFileResponse), nil
	}
}

// UploadAgentFileInvoker 上传文件
func (c *VersatileClient) UploadAgentFileInvoker(request *model.UploadAgentFileRequest) *UploadAgentFileInvoker {
	requestDef := GenReqDefForUploadAgentFile()
	return &UploadAgentFileInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
