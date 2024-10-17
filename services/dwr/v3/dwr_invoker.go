package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dwr/v3/model"
)

type AcceptServiceContractInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptServiceContractInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AcceptServiceContractInvoker) Invoke() (*model.AcceptServiceContractResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptServiceContractResponse), nil
	}
}

type AsyncInvokeApiStartWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *AsyncInvokeApiStartWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AsyncInvokeApiStartWorkflowInvoker) Invoke() (*model.AsyncInvokeApiStartWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AsyncInvokeApiStartWorkflowResponse), nil
	}
}

type CheckWorkflowAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckWorkflowAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckWorkflowAuthenticationInvoker) Invoke() (*model.CheckWorkflowAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckWorkflowAuthenticationResponse), nil
	}
}

type CreateMyActionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMyActionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMyActionTemplateInvoker) Invoke() (*model.CreateMyActionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMyActionTemplateResponse), nil
	}
}

type CreateWorkflowAuthenticationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateWorkflowAuthenticationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateWorkflowAuthenticationInvoker) Invoke() (*model.CreateWorkflowAuthenticationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateWorkflowAuthenticationResponse), nil
	}
}

type DeleteMyActionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMyActionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMyActionTemplateInvoker) Invoke() (*model.DeleteMyActionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMyActionTemplateResponse), nil
	}
}

type ListMyActionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMyActionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMyActionTemplateInvoker) Invoke() (*model.ListMyActionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMyActionTemplateResponse), nil
	}
}

type ListSystemTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemTemplatesInvoker) Invoke() (*model.ListSystemTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemTemplatesResponse), nil
	}
}

type ListWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowInstanceInvoker) Invoke() (*model.ListWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowInstanceResponse), nil
	}
}

type RestoreWorkflowExecutionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreWorkflowExecutionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreWorkflowExecutionInvoker) Invoke() (*model.RestoreWorkflowExecutionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreWorkflowExecutionResponse), nil
	}
}

type ShowPublicActionListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicActionListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicActionListInvoker) Invoke() (*model.ShowPublicActionListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicActionListResponse), nil
	}
}

type ShowPublicTemplateInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicTemplateInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPublicTemplateInfoInvoker) Invoke() (*model.ShowPublicTemplateInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicTemplateInfoResponse), nil
	}
}

type ShowServiceContractInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServiceContractInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServiceContractInvoker) Invoke() (*model.ShowServiceContractResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServiceContractResponse), nil
	}
}

type ShowSystemTemplateDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSystemTemplateDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSystemTemplateDetailInvoker) Invoke() (*model.ShowSystemTemplateDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSystemTemplateDetailResponse), nil
	}
}

type ShowThirdTemplateInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowThirdTemplateInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowThirdTemplateInfoInvoker) Invoke() (*model.ShowThirdTemplateInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowThirdTemplateInfoResponse), nil
	}
}

type ShowWorkflowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowInstanceInvoker) Invoke() (*model.ShowWorkflowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowInstanceResponse), nil
	}
}

type UpdateMyActionTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMyActionTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMyActionTemplateInvoker) Invoke() (*model.UpdateMyActionTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMyActionTemplateResponse), nil
	}
}

type UpdateMyActionTemplateToDeprecatedInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMyActionTemplateToDeprecatedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateMyActionTemplateToDeprecatedInvoker) Invoke() (*model.UpdateMyActionTemplateToDeprecatedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMyActionTemplateToDeprecatedResponse), nil
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

type DeleteWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteWorkflowInvoker) Invoke() (*model.DeleteWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWorkflowResponse), nil
	}
}

type ListWorkflowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListWorkflowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListWorkflowsInvoker) Invoke() (*model.ListWorkflowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListWorkflowsResponse), nil
	}
}

type ShowWorkflowInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWorkflowInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowWorkflowInfoInvoker) Invoke() (*model.ShowWorkflowInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWorkflowInfoResponse), nil
	}
}

type UpdateWorkflowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateWorkflowInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateWorkflowInvoker) Invoke() (*model.UpdateWorkflowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateWorkflowResponse), nil
	}
}
