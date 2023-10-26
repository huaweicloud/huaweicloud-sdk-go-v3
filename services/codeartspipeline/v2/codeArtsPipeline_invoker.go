package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartspipeline/v2/model"
)

type AcceptManualReviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptManualReviewInvoker) Invoke() (*model.AcceptManualReviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptManualReviewResponse), nil
	}
}

type BatchShowPipelinesLatestStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowPipelinesLatestStatusInvoker) Invoke() (*model.BatchShowPipelinesLatestStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowPipelinesLatestStatusResponse), nil
	}
}

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

type CreatePipelineByTemplateIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipelineByTemplateIdInvoker) Invoke() (*model.CreatePipelineByTemplateIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipelineByTemplateIdResponse), nil
	}
}

type CreatePipelineNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipelineNewInvoker) Invoke() (*model.CreatePipelineNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipelineNewResponse), nil
	}
}

type DeletePipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipelineInvoker) Invoke() (*model.DeletePipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipelineResponse), nil
	}
}

type ListPipelineRunsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineRunsInvoker) Invoke() (*model.ListPipelineRunsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineRunsResponse), nil
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

type ListPipelineTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelineTemplatesInvoker) Invoke() (*model.ListPipelineTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelineTemplatesResponse), nil
	}
}

type ListPipelinesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPipelinesInvoker) Invoke() (*model.ListPipelinesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPipelinesResponse), nil
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

type RejectManualReviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *RejectManualReviewInvoker) Invoke() (*model.RejectManualReviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RejectManualReviewResponse), nil
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

type RetryPipelineRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryPipelineRunInvoker) Invoke() (*model.RetryPipelineRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryPipelineRunResponse), nil
	}
}

type RunPipelineInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunPipelineInvoker) Invoke() (*model.RunPipelineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunPipelineResponse), nil
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

type ShowPipelineRunDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineRunDetailInvoker) Invoke() (*model.ShowPipelineRunDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineRunDetailResponse), nil
	}
}

type ShowPipelineTemplateDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineTemplateDetailInvoker) Invoke() (*model.ShowPipelineTemplateDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineTemplateDetailResponse), nil
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

type StopPipelineRunInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopPipelineRunInvoker) Invoke() (*model.StopPipelineRunResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopPipelineRunResponse), nil
	}
}
