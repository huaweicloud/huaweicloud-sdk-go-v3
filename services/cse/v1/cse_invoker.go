package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
)

type CreateEngineInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateGovernancePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateMicroserviceRouteRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteGovernancePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteMicroserviceRouteRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DownloadKieInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListEnginesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListFlavorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGovernancePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGovernancePolicyByPolicyIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListGovernancePolicysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListMicroserviceRouteRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ResizeEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RetryEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowEngineJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowEngineQuotasInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateGovernancePolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpgradeEngineInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpgradeEngineConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UploadKieInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UploadKieInvoker) Invoke() (*model.UploadKieResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadKieResponse), nil
	}
}

type CreateHttp2RpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHttp2RpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHttp2RpcInvoker) Invoke() (*model.CreateHttp2RpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHttp2RpcResponse), nil
	}
}

type CreatePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePluginInvoker) Invoke() (*model.CreatePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePluginResponse), nil
	}
}

type DeleteHttp2RpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHttp2RpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHttp2RpcInvoker) Invoke() (*model.DeleteHttp2RpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHttp2RpcResponse), nil
	}
}

type DeletePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePluginInvoker) Invoke() (*model.DeletePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePluginResponse), nil
	}
}

type ModifyHttp2RpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyHttp2RpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyHttp2RpcInvoker) Invoke() (*model.ModifyHttp2RpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyHttp2RpcResponse), nil
	}
}

type ModifyPluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyPluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyPluginInvoker) Invoke() (*model.ModifyPluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyPluginResponse), nil
	}
}

type ShowHttp2RpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHttp2RpcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHttp2RpcsInvoker) Invoke() (*model.ShowHttp2RpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHttp2RpcsResponse), nil
	}
}

type ShowPluginsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPluginsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPluginsInvoker) Invoke() (*model.ShowPluginsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPluginsResponse), nil
	}
}

type ShowSinglePluginInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSinglePluginInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSinglePluginInvoker) Invoke() (*model.ShowSinglePluginResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSinglePluginResponse), nil
	}
}

type CreateNacosNamespacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateNacosNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteNacosNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListNacosNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateNacosNamespacesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateNacosNamespacesInvoker) Invoke() (*model.UpdateNacosNamespacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNacosNamespacesResponse), nil
	}
}
