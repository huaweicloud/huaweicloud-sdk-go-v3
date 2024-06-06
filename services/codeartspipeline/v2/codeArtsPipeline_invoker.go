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

type BatchMovePipelineToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchMovePipelineToGroupInvoker) Invoke() (*model.BatchMovePipelineToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchMovePipelineToGroupResponse), nil
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

type CreateBasicPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBasicPluginInvoker) Invoke() (*model.CreateBasicPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBasicPluginResponse), nil
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

type CreatePipelineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipelineGroupInvoker) Invoke() (*model.CreatePipelineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipelineGroupResponse), nil
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

type CreatePipelineTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePipelineTemplateInvoker) Invoke() (*model.CreatePipelineTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePipelineTemplateResponse), nil
	}
}

type CreatePluginDraftInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePluginDraftInvoker) Invoke() (*model.CreatePluginDraftResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePluginDraftResponse), nil
	}
}

type CreatePluginVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePluginVersionInvoker) Invoke() (*model.CreatePluginVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePluginVersionResponse), nil
	}
}

type CreatePublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublisherInvoker) Invoke() (*model.CreatePublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublisherResponse), nil
	}
}

type CreateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRuleInvoker) Invoke() (*model.CreateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRuleResponse), nil
	}
}

type CreateStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStrategyInvoker) Invoke() (*model.CreateStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStrategyResponse), nil
	}
}

type DeleteBasicPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBasicPluginInvoker) Invoke() (*model.DeleteBasicPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBasicPluginResponse), nil
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

type DeletePipelineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipelineGroupInvoker) Invoke() (*model.DeletePipelineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipelineGroupResponse), nil
	}
}

type DeletePipelineTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePipelineTemplateInvoker) Invoke() (*model.DeletePipelineTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePipelineTemplateResponse), nil
	}
}

type DeletePluginDraftInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePluginDraftInvoker) Invoke() (*model.DeletePluginDraftResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePluginDraftResponse), nil
	}
}

type DeletePublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublisherInvoker) Invoke() (*model.DeletePublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublisherResponse), nil
	}
}

type DeleteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleInvoker) Invoke() (*model.DeleteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleResponse), nil
	}
}

type DeleteStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteStrategyInvoker) Invoke() (*model.DeleteStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteStrategyResponse), nil
	}
}

type ListAvailablePublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailablePublisherInvoker) Invoke() (*model.ListAvailablePublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailablePublisherResponse), nil
	}
}

type ListBasePluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBasePluginsInvoker) Invoke() (*model.ListBasePluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBasePluginsResponse), nil
	}
}

type ListBasePluginsNewPostInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBasePluginsNewPostInvoker) Invoke() (*model.ListBasePluginsNewPostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBasePluginsNewPostResponse), nil
	}
}

type ListPLuginVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPLuginVersionInvoker) Invoke() (*model.ListPLuginVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPLuginVersionResponse), nil
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

type ListPluginVersionNumberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginVersionNumberInvoker) Invoke() (*model.ListPluginVersionNumberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginVersionNumberResponse), nil
	}
}

type ListPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPluginsInvoker) Invoke() (*model.ListPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPluginsResponse), nil
	}
}

type ListProjectStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectStrategyInvoker) Invoke() (*model.ListProjectStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectStrategyResponse), nil
	}
}

type ListPublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublisherInvoker) Invoke() (*model.ListPublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublisherResponse), nil
	}
}

type ListRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleInvoker) Invoke() (*model.ListRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleResponse), nil
	}
}

type ListStagePluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStagePluginsInvoker) Invoke() (*model.ListStagePluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStagePluginsResponse), nil
	}
}

type ListStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStrategyInvoker) Invoke() (*model.ListStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStrategyResponse), nil
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

type PublishPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishPluginInvoker) Invoke() (*model.PublishPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishPluginResponse), nil
	}
}

type PublishPluginBindInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishPluginBindInvoker) Invoke() (*model.PublishPluginBindResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishPluginBindResponse), nil
	}
}

type PublishPluginDraftInvoker struct {
	*invoker.BaseInvoker
}

func (i *PublishPluginDraftInvoker) Invoke() (*model.PublishPluginDraftResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishPluginDraftResponse), nil
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

type ShowBasicPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBasicPluginInvoker) Invoke() (*model.ShowBasicPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBasicPluginResponse), nil
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

type ShowPipelineArtifactsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineArtifactsInvoker) Invoke() (*model.ShowPipelineArtifactsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineArtifactsResponse), nil
	}
}

type ShowPipelineDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineDetailInvoker) Invoke() (*model.ShowPipelineDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineDetailResponse), nil
	}
}

type ShowPipelineGroupTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineGroupTreeInvoker) Invoke() (*model.ShowPipelineGroupTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineGroupTreeResponse), nil
	}
}

type ShowPipelineLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPipelineLogInvoker) Invoke() (*model.ShowPipelineLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPipelineLogResponse), nil
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

type ShowPluginInputsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginInputsInvoker) Invoke() (*model.ShowPluginInputsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginInputsResponse), nil
	}
}

type ShowPluginMetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginMetricsInvoker) Invoke() (*model.ShowPluginMetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginMetricsResponse), nil
	}
}

type ShowPluginOutputsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginOutputsInvoker) Invoke() (*model.ShowPluginOutputsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginOutputsResponse), nil
	}
}

type ShowPluginVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginVersionInvoker) Invoke() (*model.ShowPluginVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginVersionResponse), nil
	}
}

type ShowProjectStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectStrategyInvoker) Invoke() (*model.ShowProjectStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectStrategyResponse), nil
	}
}

type ShowPublisherInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublisherInvoker) Invoke() (*model.ShowPublisherResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublisherResponse), nil
	}
}

type ShowRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRuleInvoker) Invoke() (*model.ShowRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRuleResponse), nil
	}
}

type ShowStepOutputsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStepOutputsInvoker) Invoke() (*model.ShowStepOutputsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStepOutputsResponse), nil
	}
}

type ShowStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStrategyInvoker) Invoke() (*model.ShowStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStrategyResponse), nil
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

type SwitchStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchStrategyInvoker) Invoke() (*model.SwitchStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchStrategyResponse), nil
	}
}

type UpdateBasicPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBasicPluginInvoker) Invoke() (*model.UpdateBasicPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBasicPluginResponse), nil
	}
}

type UpdatePipelineGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipelineGroupInvoker) Invoke() (*model.UpdatePipelineGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipelineGroupResponse), nil
	}
}

type UpdatePipelineInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipelineInfoInvoker) Invoke() (*model.UpdatePipelineInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipelineInfoResponse), nil
	}
}

type UpdatePipelineTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePipelineTemplateInvoker) Invoke() (*model.UpdatePipelineTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePipelineTemplateResponse), nil
	}
}

type UpdatePluginBaseInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePluginBaseInfoInvoker) Invoke() (*model.UpdatePluginBaseInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginBaseInfoResponse), nil
	}
}

type UpdatePluginDraftInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePluginDraftInvoker) Invoke() (*model.UpdatePluginDraftResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePluginDraftResponse), nil
	}
}

type UpdateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRuleInvoker) Invoke() (*model.UpdateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRuleResponse), nil
	}
}

type UpdateStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateStrategyInvoker) Invoke() (*model.UpdateStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateStrategyResponse), nil
	}
}

type UploadBasicPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadBasicPluginInvoker) Invoke() (*model.UploadBasicPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadBasicPluginResponse), nil
	}
}

type UploadPluginIconInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadPluginIconInvoker) Invoke() (*model.UploadPluginIconResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadPluginIconResponse), nil
	}
}

type UploadPublisherIconInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadPublisherIconInvoker) Invoke() (*model.UploadPublisherIconResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadPublisherIconResponse), nil
	}
}
