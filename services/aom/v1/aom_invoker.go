package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aom/v1/model"
)

type CreateFastExecuteScriptInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFastExecuteScriptInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFastExecuteScriptInvoker) Invoke() (*model.CreateFastExecuteScriptResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFastExecuteScriptResponse), nil
	}
}

type CreateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowInvoker) Invoke() (*model.CreateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowResponse), nil
	}
}

type ExecuteWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExecuteWorkflowInvoker) Invoke() (*model.ExecuteWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteWorkflowResponse), nil
	}
}

type ListAllJobByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllJobByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllJobByNameInvoker) Invoke() (*model.ListAllJobByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllJobByNameResponse), nil
	}
}

type ListAllScriptByNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllScriptByNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllScriptByNameInvoker) Invoke() (*model.ListAllScriptByNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllScriptByNameResponse), nil
	}
}

type ListAllVersionByVersionIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllVersionByVersionIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllVersionByVersionIdInvoker) Invoke() (*model.ListAllVersionByVersionIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllVersionByVersionIdResponse), nil
	}
}

type ListTemplateByJobIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTemplateByJobIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTemplateByJobIdInvoker) Invoke() (*model.ListTemplateByJobIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTemplateByJobIdResponse), nil
	}
}

type ListWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowInvoker) Invoke() (*model.ListWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowResponse), nil
	}
}

type ListWorkflowExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowExecutionsInvoker) Invoke() (*model.ListWorkflowExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowExecutionsResponse), nil
	}
}

type SearchTemplateByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchTemplateByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchTemplateByIdInvoker) Invoke() (*model.SearchTemplateByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchTemplateByIdResponse), nil
	}
}

type SearchWorkflowExecutionDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchWorkflowExecutionDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchWorkflowExecutionDetailInvoker) Invoke() (*model.SearchWorkflowExecutionDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchWorkflowExecutionDetailResponse), nil
	}
}

type StartPausingWorkflowExecutionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartPausingWorkflowExecutionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartPausingWorkflowExecutionsInvoker) Invoke() (*model.StartPausingWorkflowExecutionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartPausingWorkflowExecutionsResponse), nil
	}
}

type StopExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopExecutionInvoker) Invoke() (*model.StopExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopExecutionResponse), nil
	}
}

type UpdateWorkflowTriggerStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowTriggerStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowTriggerStatusInvoker) Invoke() (*model.UpdateWorkflowTriggerStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowTriggerStatusResponse), nil
	}
}
