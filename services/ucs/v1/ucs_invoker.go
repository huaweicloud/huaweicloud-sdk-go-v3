package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ucs/v1/model"
)

type CreateAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAddonInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAddonInstanceInvoker) Invoke() (*model.CreateAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAddonInstanceResponse), nil
	}
}

type CreateClusterGroupPolicyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterGroupPolicyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterGroupPolicyInstanceInvoker) Invoke() (*model.CreateClusterGroupPolicyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterGroupPolicyInstanceResponse), nil
	}
}

type CreateConfigSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConfigSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateConfigSetInvoker) Invoke() (*model.CreateConfigSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConfigSetResponse), nil
	}
}

type CreateFederationCertInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFederationCertInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFederationCertInvoker) Invoke() (*model.CreateFederationCertResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFederationCertResponse), nil
	}
}

type CreateFederationConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFederationConnectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFederationConnectionInvoker) Invoke() (*model.CreateFederationConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFederationConnectionResponse), nil
	}
}

type CreateProxyUnitedWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateProxyUnitedWorkloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateProxyUnitedWorkloadInvoker) Invoke() (*model.CreateProxyUnitedWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateProxyUnitedWorkloadResponse), nil
	}
}

type CreateRecordSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRecordSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRecordSetInvoker) Invoke() (*model.CreateRecordSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRecordSetResponse), nil
	}
}

type CreateRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRepoInvoker) Invoke() (*model.CreateRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRepoResponse), nil
	}
}

type CreateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRuleInvoker) Invoke() (*model.CreateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRuleResponse), nil
	}
}

type DeleteAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAddonInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAddonInstanceInvoker) Invoke() (*model.DeleteAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAddonInstanceResponse), nil
	}
}

type DeleteClusterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClusterGroupInvoker) Invoke() (*model.DeleteClusterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterGroupResponse), nil
	}
}

type DeleteConfigSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConfigSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteConfigSetInvoker) Invoke() (*model.DeleteConfigSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConfigSetResponse), nil
	}
}

type DeletePolicyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePolicyInstanceInvoker) Invoke() (*model.DeletePolicyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyInstanceResponse), nil
	}
}

type DeleteProxyUnitedWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteProxyUnitedWorkloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteProxyUnitedWorkloadInvoker) Invoke() (*model.DeleteProxyUnitedWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteProxyUnitedWorkloadResponse), nil
	}
}

type DeleteRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRepoInvoker) Invoke() (*model.DeleteRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRepoResponse), nil
	}
}

type DeleteRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRuleInvoker) Invoke() (*model.DeleteRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRuleResponse), nil
	}
}

type DisableClusterGroupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableClusterGroupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableClusterGroupPolicyInvoker) Invoke() (*model.DisableClusterGroupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableClusterGroupPolicyResponse), nil
	}
}

type DisableFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableFederationInvoker) Invoke() (*model.DisableFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableFederationResponse), nil
	}
}

type DownloadFederationKubeconfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadFederationKubeconfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadFederationKubeconfigInvoker) Invoke() (*model.DownloadFederationKubeconfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadFederationKubeconfigResponse), nil
	}
}

type EnableClusterGroupPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableClusterGroupPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableClusterGroupPolicyInvoker) Invoke() (*model.EnableClusterGroupPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableClusterGroupPolicyResponse), nil
	}
}

type EnableFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableFederationInvoker) Invoke() (*model.EnableFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableFederationResponse), nil
	}
}

type JoinGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *JoinGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *JoinGroupInvoker) Invoke() (*model.JoinGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.JoinGroupResponse), nil
	}
}

type LeaveGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *LeaveGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LeaveGroupInvoker) Invoke() (*model.LeaveGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveGroupResponse), nil
	}
}

type ListAddonInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddonInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAddonInstancesInvoker) Invoke() (*model.ListAddonInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddonInstancesResponse), nil
	}
}

type ListAddonTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAddonTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAddonTemplatesInvoker) Invoke() (*model.ListAddonTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAddonTemplatesResponse), nil
	}
}

type ListClusterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterGroupInvoker) Invoke() (*model.ListClusterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterGroupResponse), nil
	}
}

type ListConfigSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigSetsInvoker) Invoke() (*model.ListConfigSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigSetsResponse), nil
	}
}

type ListPolicyDefinitionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyDefinitionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyDefinitionsInvoker) Invoke() (*model.ListPolicyDefinitionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyDefinitionsResponse), nil
	}
}

type ListPolicyInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyInstancesInvoker) Invoke() (*model.ListPolicyInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyInstancesResponse), nil
	}
}

type ListPolicyJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPolicyJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPolicyJobsInvoker) Invoke() (*model.ListPolicyJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPolicyJobsResponse), nil
	}
}

type ListRecordSetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecordSetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecordSetsInvoker) Invoke() (*model.ListRecordSetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecordSetsResponse), nil
	}
}

type ListReposInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReposInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListReposInvoker) Invoke() (*model.ListReposResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReposResponse), nil
	}
}

type ListRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRuleInvoker) Invoke() (*model.ListRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRuleResponse), nil
	}
}

type ListServerConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServerConfigsInvoker) Invoke() (*model.ListServerConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerConfigsResponse), nil
	}
}

type RegisterClusterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterClusterGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterClusterGroupInvoker) Invoke() (*model.RegisterClusterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterClusterGroupResponse), nil
	}
}

type ShowAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAddonInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAddonInstanceInvoker) Invoke() (*model.ShowAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAddonInstanceResponse), nil
	}
}

type ShowClusterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterGroupInvoker) Invoke() (*model.ShowClusterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterGroupResponse), nil
	}
}

type ShowConfigSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConfigSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowConfigSetInvoker) Invoke() (*model.ShowConfigSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConfigSetResponse), nil
	}
}

type ShowFederationProgressInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFederationProgressInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFederationProgressInvoker) Invoke() (*model.ShowFederationProgressResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFederationProgressResponse), nil
	}
}

type ShowGitopsStatisticsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGitopsStatisticsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGitopsStatisticsInvoker) Invoke() (*model.ShowGitopsStatisticsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGitopsStatisticsResponse), nil
	}
}

type ShowPolicyDefinitionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyDefinitionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyDefinitionInvoker) Invoke() (*model.ShowPolicyDefinitionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyDefinitionResponse), nil
	}
}

type ShowPolicyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyInstanceInvoker) Invoke() (*model.ShowPolicyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyInstanceResponse), nil
	}
}

type ShowPolicyJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPolicyJobInvoker) Invoke() (*model.ShowPolicyJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyJobResponse), nil
	}
}

type ShowProxyUnitedWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProxyUnitedWorkloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProxyUnitedWorkloadInvoker) Invoke() (*model.ShowProxyUnitedWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProxyUnitedWorkloadResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type UpdateAddonInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAddonInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAddonInstanceInvoker) Invoke() (*model.UpdateAddonInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAddonInstanceResponse), nil
	}
}

type UpdateClusterGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterGroupInvoker) Invoke() (*model.UpdateClusterGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterGroupResponse), nil
	}
}

type UpdateClusterGroupAssociatedClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedClustersInvoker) Invoke() (*model.UpdateClusterGroupAssociatedClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterGroupAssociatedClustersResponse), nil
	}
}

type UpdateClusterGroupAssociatedRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedRulesInvoker) Invoke() (*model.UpdateClusterGroupAssociatedRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterGroupAssociatedRulesResponse), nil
	}
}

type UpdateClusterGroupAssociatedZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterGroupAssociatedZonesInvoker) Invoke() (*model.UpdateClusterGroupAssociatedZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterGroupAssociatedZonesResponse), nil
	}
}

type UpdateConfigSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateConfigSetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateConfigSetInvoker) Invoke() (*model.UpdateConfigSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateConfigSetResponse), nil
	}
}

type UpdatePolicyDefinationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyDefinationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyDefinationInvoker) Invoke() (*model.UpdatePolicyDefinationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyDefinationResponse), nil
	}
}

type UpdatePolicyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePolicyInstanceInvoker) Invoke() (*model.UpdatePolicyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyInstanceResponse), nil
	}
}

type UpdateProxyUnitedWorkloadInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateProxyUnitedWorkloadInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateProxyUnitedWorkloadInvoker) Invoke() (*model.UpdateProxyUnitedWorkloadResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateProxyUnitedWorkloadResponse), nil
	}
}

type UpdateRepoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRepoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRepoInvoker) Invoke() (*model.UpdateRepoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRepoResponse), nil
	}
}

type UpdateRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRuleInvoker) Invoke() (*model.UpdateRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRuleResponse), nil
	}
}

type UpgradeFederationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeFederationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeFederationInvoker) Invoke() (*model.UpgradeFederationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeFederationResponse), nil
	}
}

type CreateClusterConfInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterConfInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterConfInvoker) Invoke() (*model.CreateClusterConfResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterConfResponse), nil
	}
}

type CreateClusterKubeconfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterKubeconfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterKubeconfigInvoker) Invoke() (*model.CreateClusterKubeconfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterKubeconfigResponse), nil
	}
}

type CreateClusterPolicyInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterPolicyInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterPolicyInstanceInvoker) Invoke() (*model.CreateClusterPolicyInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterPolicyInstanceResponse), nil
	}
}

type DeleteClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClusterInvoker) Invoke() (*model.DeleteClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterResponse), nil
	}
}

type DisableClusterPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableClusterPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableClusterPolicyInvoker) Invoke() (*model.DisableClusterPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableClusterPolicyResponse), nil
	}
}

type DisableGitOpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableGitOpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableGitOpsInvoker) Invoke() (*model.DisableGitOpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableGitOpsResponse), nil
	}
}

type EnableClusterPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableClusterPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableClusterPolicyInvoker) Invoke() (*model.EnableClusterPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableClusterPolicyResponse), nil
	}
}

type EnableGitOpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableGitOpsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableGitOpsInvoker) Invoke() (*model.EnableGitOpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableGitOpsResponse), nil
	}
}

type ListManagedClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedClustersInvoker) Invoke() (*model.ListManagedClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedClustersResponse), nil
	}
}

type ListRegisteredClusterVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRegisteredClusterVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRegisteredClusterVersionsInvoker) Invoke() (*model.ListRegisteredClusterVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRegisteredClusterVersionsResponse), nil
	}
}

type RegisterClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterClusterInvoker) Invoke() (*model.RegisterClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterClusterResponse), nil
	}
}

type RetryClusterActivationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryClusterActivationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RetryClusterActivationInvoker) Invoke() (*model.RetryClusterActivationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryClusterActivationResponse), nil
	}
}

type ShowClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterInvoker) Invoke() (*model.ShowClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterResponse), nil
	}
}

type ShowClusterAccessInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterAccessInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterAccessInfoInvoker) Invoke() (*model.ShowClusterAccessInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterAccessInfoResponse), nil
	}
}

type ShowClusterListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterListInvoker) Invoke() (*model.ShowClusterListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterListResponse), nil
	}
}

type ShowGitOpsStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowGitOpsStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowGitOpsStatusInvoker) Invoke() (*model.ShowGitOpsStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowGitOpsStatusResponse), nil
	}
}

type UpdateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterInvoker) Invoke() (*model.UpdateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterResponse), nil
	}
}

type UpdateClusterRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClusterRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateClusterRulesInvoker) Invoke() (*model.UpdateClusterRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClusterRulesResponse), nil
	}
}
