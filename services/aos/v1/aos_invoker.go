package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/aos/v1/model"
)

type ApplyExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ApplyExecutionPlanInvoker) Invoke() (*model.ApplyExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ApplyExecutionPlanResponse), nil
	}
}

type CreateExecutionPlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExecutionPlanInvoker) Invoke() (*model.CreateExecutionPlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExecutionPlanResponse), nil
	}
}

type EstimateExecutionPlanPriceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EstimateExecutionPlanPriceInvoker) Invoke() (*model.EstimateExecutionPlanPriceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EstimateExecutionPlanPriceResponse), nil
	}
}

type GetStackTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetStackTemplateInvoker) Invoke() (*model.GetStackTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetStackTemplateResponse), nil
	}
}

type ListExecutionPlansInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExecutionPlansInvoker) Invoke() (*model.ListExecutionPlansResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExecutionPlansResponse), nil
	}
}

type ListStackOutputsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStackOutputsInvoker) Invoke() (*model.ListStackOutputsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStackOutputsResponse), nil
	}
}

type ParseTemplateVariablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ParseTemplateVariablesInvoker) Invoke() (*model.ParseTemplateVariablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ParseTemplateVariablesResponse), nil
	}
}

type ContinueRollbackStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *ContinueRollbackStackInvoker) Invoke() (*model.ContinueRollbackStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContinueRollbackStackResponse), nil
	}
}

type CreateStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateStackInvoker) Invoke() (*model.CreateStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateStackResponse), nil
	}
}

type DeployStackInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeployStackInvoker) Invoke() (*model.DeployStackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployStackResponse), nil
	}
}

type ListStacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListStacksInvoker) Invoke() (*model.ListStacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListStacksResponse), nil
	}
}
