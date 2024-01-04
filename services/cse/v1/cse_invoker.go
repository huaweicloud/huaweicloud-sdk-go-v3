package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
)

type CreateEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEngineInvoker) Invoke() (*model.CreateEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEngineResponse), nil
	}
}

type CreateGovernancePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGovernancePolicyInvoker) Invoke() (*model.CreateGovernancePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGovernancePolicyResponse), nil
	}
}

type CreateMicroserviceRouteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMicroserviceRouteRuleInvoker) Invoke() (*model.CreateMicroserviceRouteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMicroserviceRouteRuleResponse), nil
	}
}

type DeleteEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEngineInvoker) Invoke() (*model.DeleteEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEngineResponse), nil
	}
}

type DeleteGovernancePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteGovernancePolicyInvoker) Invoke() (*model.DeleteGovernancePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteGovernancePolicyResponse), nil
	}
}

type DeleteMicroserviceRouteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMicroserviceRouteRuleInvoker) Invoke() (*model.DeleteMicroserviceRouteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMicroserviceRouteRuleResponse), nil
	}
}

type DownloadKieInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadKieInvoker) Invoke() (*model.DownloadKieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadKieResponse), nil
	}
}

type ListEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnginesInvoker) Invoke() (*model.ListEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnginesResponse), nil
	}
}

type ListFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorsInvoker) Invoke() (*model.ListFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorsResponse), nil
	}
}

type ListGovernancePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGovernancePolicyInvoker) Invoke() (*model.ListGovernancePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGovernancePolicyResponse), nil
	}
}

type ListGovernancePolicyByPolicyIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGovernancePolicyByPolicyIdInvoker) Invoke() (*model.ListGovernancePolicyByPolicyIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGovernancePolicyByPolicyIdResponse), nil
	}
}

type ListGovernancePolicysInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGovernancePolicysInvoker) Invoke() (*model.ListGovernancePolicysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGovernancePolicysResponse), nil
	}
}

type ListMicroserviceRouteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMicroserviceRouteRuleInvoker) Invoke() (*model.ListMicroserviceRouteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMicroserviceRouteRuleResponse), nil
	}
}

type ResizeEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeEngineInvoker) Invoke() (*model.ResizeEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeEngineResponse), nil
	}
}

type RetryEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryEngineInvoker) Invoke() (*model.RetryEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryEngineResponse), nil
	}
}

type ShowEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineInvoker) Invoke() (*model.ShowEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineResponse), nil
	}
}

type ShowEngineJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineJobInvoker) Invoke() (*model.ShowEngineJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineJobResponse), nil
	}
}

type ShowEngineQuotasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEngineQuotasInvoker) Invoke() (*model.ShowEngineQuotasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEngineQuotasResponse), nil
	}
}

type UpdateGovernancePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateGovernancePolicyInvoker) Invoke() (*model.UpdateGovernancePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateGovernancePolicyResponse), nil
	}
}

type UpgradeEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeEngineInvoker) Invoke() (*model.UpgradeEngineResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeEngineResponse), nil
	}
}

type UpgradeEngineConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeEngineConfigInvoker) Invoke() (*model.UpgradeEngineConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeEngineConfigResponse), nil
	}
}

type UploadKieInvoker struct {
	*invoker.BaseInvoker
}

func (i *UploadKieInvoker) Invoke() (*model.UploadKieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadKieResponse), nil
	}
}

type CreateNacosNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNacosNamespacesInvoker) Invoke() (*model.CreateNacosNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateNacosNamespacesResponse), nil
	}
}

type DeleteNacosNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNacosNamespacesInvoker) Invoke() (*model.DeleteNacosNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNacosNamespacesResponse), nil
	}
}

type ListNacosNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNacosNamespacesInvoker) Invoke() (*model.ListNacosNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNacosNamespacesResponse), nil
	}
}

type UpdateNacosNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNacosNamespacesInvoker) Invoke() (*model.UpdateNacosNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNacosNamespacesResponse), nil
	}
}
