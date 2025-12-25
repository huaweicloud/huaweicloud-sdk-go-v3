package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/versatile/v1/model"
)

type SearchKnowledgeBaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchKnowledgeBaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchKnowledgeBaseInvoker) Invoke() (*model.SearchKnowledgeBaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchKnowledgeBaseResponse), nil
	}
}

type RunAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAgentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunAgentInvoker) Invoke() (*model.RunAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAgentResponse), nil
	}
}

type RunWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunWorkflowInvoker) Invoke() (*model.RunWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunWorkflowResponse), nil
	}
}

type UploadAgentFileInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadAgentFileInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadAgentFileInvoker) Invoke() (*model.UploadAgentFileResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadAgentFileResponse), nil
	}
}
