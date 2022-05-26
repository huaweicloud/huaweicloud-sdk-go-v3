package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudpipeline/v2/model"
)

type BatchShowPipelinesStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowPipelinesStatusInvoker) Invoke() (*model.BatchShowPipelinesStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowPipelinesStatusResponse), nil
	}
}

type CreatePipelineByTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipelineByTemplateInvoker) Invoke() (*model.CreatePipelineByTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipelineByTemplateResponse), nil
	}
}

type ListPipelineSimpleInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineSimpleInfoInvoker) Invoke() (*model.ListPipelineSimpleInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineSimpleInfoResponse), nil
	}
}

type ListPipleineBuildResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipleineBuildResultInvoker) Invoke() (*model.ListPipleineBuildResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipleineBuildResultResponse), nil
	}
}

type ListTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplatesInvoker) Invoke() (*model.ListTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplatesResponse), nil
	}
}

type RegisterAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterAgentInvoker) Invoke() (*model.RegisterAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterAgentResponse), nil
	}
}

type RemovePipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemovePipelineInvoker) Invoke() (*model.RemovePipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemovePipelineResponse), nil
	}
}

type ShowAgentStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentStatusInvoker) Invoke() (*model.ShowAgentStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentStatusResponse), nil
	}
}

type ShowInstanceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceStatusInvoker) Invoke() (*model.ShowInstanceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceStatusResponse), nil
	}
}

type ShowPipleineStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipleineStatusInvoker) Invoke() (*model.ShowPipleineStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipleineStatusResponse), nil
	}
}

type ShowTemplateDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateDetailInvoker) Invoke() (*model.ShowTemplateDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateDetailResponse), nil
	}
}

type StartNewPipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartNewPipelineInvoker) Invoke() (*model.StartNewPipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartNewPipelineResponse), nil
	}
}

type StopPipelineNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopPipelineNewInvoker) Invoke() (*model.StopPipelineNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopPipelineNewResponse), nil
	}
}
