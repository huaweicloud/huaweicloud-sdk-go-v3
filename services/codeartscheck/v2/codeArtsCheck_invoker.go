package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartscheck/v2/model"
)

type CheckParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckParametersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckParametersInvoker) Invoke() (*model.CheckParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckParametersResponse), nil
	}
}

type CheckRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRecordInvoker) Invoke() (*model.CheckRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRecordResponse), nil
	}
}

type CheckRulesetParametersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckRulesetParametersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckRulesetParametersInvoker) Invoke() (*model.CheckRulesetParametersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckRulesetParametersResponse), nil
	}
}

type CreateRulesetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRulesetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRulesetInvoker) Invoke() (*model.CreateRulesetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRulesetResponse), nil
	}
}

type CreateTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTaskInvoker) Invoke() (*model.CreateTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTaskResponse), nil
	}
}

type DeleteRulesetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRulesetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRulesetInvoker) Invoke() (*model.DeleteRulesetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRulesetResponse), nil
	}
}

type DeleteTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTaskInvoker) Invoke() (*model.DeleteTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTaskResponse), nil
	}
}

type ListRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRulesInvoker) Invoke() (*model.ListRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRulesResponse), nil
	}
}

type ListRulesetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRulesetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRulesetsInvoker) Invoke() (*model.ListRulesetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRulesetsResponse), nil
	}
}

type ListTaskParameterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskParameterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskParameterInvoker) Invoke() (*model.ListTaskParameterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskParameterResponse), nil
	}
}

type ListTaskRulesetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskRulesetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskRulesetInvoker) Invoke() (*model.ListTaskRulesetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskRulesetResponse), nil
	}
}

type ListTemplateRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateRulesInvoker) Invoke() (*model.ListTemplateRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateRulesResponse), nil
	}
}

type RunTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTaskInvoker) Invoke() (*model.RunTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTaskResponse), nil
	}
}

type SetDefaulTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetDefaulTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetDefaulTemplateInvoker) Invoke() (*model.SetDefaulTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetDefaulTemplateResponse), nil
	}
}

type ShowProgressDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProgressDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProgressDetailInvoker) Invoke() (*model.ShowProgressDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProgressDetailResponse), nil
	}
}

type ShowTaskCmetricsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskCmetricsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskCmetricsInvoker) Invoke() (*model.ShowTaskCmetricsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskCmetricsResponse), nil
	}
}

type ShowTaskDefectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskDefectsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskDefectsInvoker) Invoke() (*model.ShowTaskDefectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskDefectsResponse), nil
	}
}

type ShowTaskDefectsStatisticInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskDefectsStatisticInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskDefectsStatisticInvoker) Invoke() (*model.ShowTaskDefectsStatisticResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskDefectsStatisticResponse), nil
	}
}

type ShowTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskDetailInvoker) Invoke() (*model.ShowTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskDetailResponse), nil
	}
}

type ShowTaskListByProjectIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskListByProjectIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskListByProjectIdInvoker) Invoke() (*model.ShowTaskListByProjectIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskListByProjectIdResponse), nil
	}
}

type ShowTaskPathTreeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskPathTreeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskPathTreeInvoker) Invoke() (*model.ShowTaskPathTreeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskPathTreeResponse), nil
	}
}

type ShowTaskSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTaskSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTaskSettingsInvoker) Invoke() (*model.ShowTaskSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTaskSettingsResponse), nil
	}
}

type ShowTasklogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTasklogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTasklogInvoker) Invoke() (*model.ShowTasklogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTasklogResponse), nil
	}
}

type ShowTasksRulesetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTasksRulesetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTasksRulesetsInvoker) Invoke() (*model.ShowTasksRulesetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTasksRulesetsResponse), nil
	}
}

type StopTaskByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopTaskByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopTaskByIdInvoker) Invoke() (*model.StopTaskByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopTaskByIdResponse), nil
	}
}

type UpdateDefectStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDefectStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDefectStatusInvoker) Invoke() (*model.UpdateDefectStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDefectStatusResponse), nil
	}
}

type UpdateIgnorePathInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIgnorePathInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIgnorePathInvoker) Invoke() (*model.UpdateIgnorePathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIgnorePathResponse), nil
	}
}

type UpdateTaskRulesetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskRulesetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskRulesetInvoker) Invoke() (*model.UpdateTaskRulesetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskRulesetResponse), nil
	}
}

type UpdateTaskSettingsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTaskSettingsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTaskSettingsInvoker) Invoke() (*model.UpdateTaskSettingsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTaskSettingsResponse), nil
	}
}
