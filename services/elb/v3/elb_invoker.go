package v3

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type BatchCreateMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateMembersInvoker) Invoke() (*model.BatchCreateMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateMembersResponse), nil
	}
}

type BatchDeleteMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMembersInvoker) Invoke() (*model.BatchDeleteMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMembersResponse), nil
	}
}

type BatchUpdateMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateMembersInvoker) Invoke() (*model.BatchUpdateMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateMembersResponse), nil
	}
}

type BatchUpdatePoliciesPriorityInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdatePoliciesPriorityInvoker) Invoke() (*model.BatchUpdatePoliciesPriorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

type ChangeLoadbalancerChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeLoadbalancerChargeModeInvoker) Invoke() (*model.ChangeLoadbalancerChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type CreateHealthMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHealthMonitorInvoker) Invoke() (*model.CreateHealthMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHealthMonitorResponse), nil
	}
}

type CreateL7PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateL7PolicyInvoker) Invoke() (*model.CreateL7PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateL7PolicyResponse), nil
	}
}

type CreateL7RuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateL7RuleInvoker) Invoke() (*model.CreateL7RuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateL7RuleResponse), nil
	}
}

type CreateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateListenerInvoker) Invoke() (*model.CreateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateListenerResponse), nil
	}
}

type CreateLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLoadBalancerInvoker) Invoke() (*model.CreateLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLoadBalancerResponse), nil
	}
}

type CreateLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLogtankInvoker) Invoke() (*model.CreateLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogtankResponse), nil
	}
}

type CreateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMemberInvoker) Invoke() (*model.CreateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMemberResponse), nil
	}
}

type CreatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePoolInvoker) Invoke() (*model.CreatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePoolResponse), nil
	}
}

type CreateSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityPolicyInvoker) Invoke() (*model.CreateSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityPolicyResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type DeleteHealthMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHealthMonitorInvoker) Invoke() (*model.DeleteHealthMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHealthMonitorResponse), nil
	}
}

type DeleteL7PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteL7PolicyInvoker) Invoke() (*model.DeleteL7PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteL7PolicyResponse), nil
	}
}

type DeleteL7RuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteL7RuleInvoker) Invoke() (*model.DeleteL7RuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteL7RuleResponse), nil
	}
}

type DeleteListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteListenerInvoker) Invoke() (*model.DeleteListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteListenerResponse), nil
	}
}

type DeleteListenerForceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteListenerForceInvoker) Invoke() (*model.DeleteListenerForceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteListenerForceResponse), nil
	}
}

type DeleteLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadBalancerInvoker) Invoke() (*model.DeleteLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadBalancerResponse), nil
	}
}

type DeleteLoadBalancerForceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadBalancerForceInvoker) Invoke() (*model.DeleteLoadBalancerForceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadBalancerForceResponse), nil
	}
}

type DeleteLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLogtankInvoker) Invoke() (*model.DeleteLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogtankResponse), nil
	}
}

type DeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInvoker) Invoke() (*model.DeleteMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMemberResponse), nil
	}
}

type DeletePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePoolInvoker) Invoke() (*model.DeletePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoolResponse), nil
	}
}

type DeleteSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityPolicyInvoker) Invoke() (*model.DeleteSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityPolicyResponse), nil
	}
}

type ListAllMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllMembersInvoker) Invoke() (*model.ListAllMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllMembersResponse), nil
	}
}

type ListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailabilityZonesInvoker) Invoke() (*model.ListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailabilityZonesResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
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

type ListHealthMonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHealthMonitorsInvoker) Invoke() (*model.ListHealthMonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHealthMonitorsResponse), nil
	}
}

type ListL7PoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListL7PoliciesInvoker) Invoke() (*model.ListL7PoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListL7PoliciesResponse), nil
	}
}

type ListL7RulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListL7RulesInvoker) Invoke() (*model.ListL7RulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListL7RulesResponse), nil
	}
}

type ListListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenersInvoker) Invoke() (*model.ListListenersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenersResponse), nil
	}
}

type ListLoadBalancersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadBalancersInvoker) Invoke() (*model.ListLoadBalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadBalancersResponse), nil
	}
}

type ListLogtanksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogtanksInvoker) Invoke() (*model.ListLogtanksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogtanksResponse), nil
	}
}

type ListMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMembersInvoker) Invoke() (*model.ListMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMembersResponse), nil
	}
}

type ListPoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoolsInvoker) Invoke() (*model.ListPoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoolsResponse), nil
	}
}

type ListQuotaDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaDetailsInvoker) Invoke() (*model.ListQuotaDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaDetailsResponse), nil
	}
}

type ListSecurityPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityPoliciesInvoker) Invoke() (*model.ListSecurityPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityPoliciesResponse), nil
	}
}

type ListSystemSecurityPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSystemSecurityPoliciesInvoker) Invoke() (*model.ListSystemSecurityPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}

type ShowFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlavorInvoker) Invoke() (*model.ShowFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlavorResponse), nil
	}
}

type ShowHealthMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHealthMonitorInvoker) Invoke() (*model.ShowHealthMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthMonitorResponse), nil
	}
}

type ShowL7PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowL7PolicyInvoker) Invoke() (*model.ShowL7PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowL7PolicyResponse), nil
	}
}

type ShowL7RuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowL7RuleInvoker) Invoke() (*model.ShowL7RuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowL7RuleResponse), nil
	}
}

type ShowListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListenerInvoker) Invoke() (*model.ShowListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerResponse), nil
	}
}

type ShowLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerInvoker) Invoke() (*model.ShowLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerResponse), nil
	}
}

type ShowLoadBalancerStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerStatusInvoker) Invoke() (*model.ShowLoadBalancerStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

type ShowLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogtankInvoker) Invoke() (*model.ShowLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogtankResponse), nil
	}
}

type ShowMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberInvoker) Invoke() (*model.ShowMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberResponse), nil
	}
}

type ShowPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolInvoker) Invoke() (*model.ShowPoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPoolResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type ShowSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityPolicyInvoker) Invoke() (*model.ShowSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityPolicyResponse), nil
	}
}

type UpdateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateInvoker) Invoke() (*model.UpdateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateResponse), nil
	}
}

type UpdateHealthMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHealthMonitorInvoker) Invoke() (*model.UpdateHealthMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHealthMonitorResponse), nil
	}
}

type UpdateL7PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateL7PolicyInvoker) Invoke() (*model.UpdateL7PolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateL7PolicyResponse), nil
	}
}

type UpdateL7RuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateL7RuleInvoker) Invoke() (*model.UpdateL7RuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateL7RuleResponse), nil
	}
}

type UpdateListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateListenerInvoker) Invoke() (*model.UpdateListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateListenerResponse), nil
	}
}

type UpdateLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLoadBalancerInvoker) Invoke() (*model.UpdateLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLoadBalancerResponse), nil
	}
}

type UpdateLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLogtankInvoker) Invoke() (*model.UpdateLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLogtankResponse), nil
	}
}

type UpdateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberInvoker) Invoke() (*model.UpdateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberResponse), nil
	}
}

type UpdatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePoolInvoker) Invoke() (*model.UpdatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePoolResponse), nil
	}
}

type UpdateSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityPolicyInvoker) Invoke() (*model.UpdateSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityPolicyResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type BatchDeleteIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteIpListInvoker) Invoke() (*model.BatchDeleteIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteIpListResponse), nil
	}
}

type CountPreoccupyIpNumInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountPreoccupyIpNumInvoker) Invoke() (*model.CountPreoccupyIpNumResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountPreoccupyIpNumResponse), nil
	}
}

type CreateIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateIpGroupInvoker) Invoke() (*model.CreateIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateIpGroupResponse), nil
	}
}

type DeleteIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIpGroupInvoker) Invoke() (*model.DeleteIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIpGroupResponse), nil
	}
}

type ListIpGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIpGroupsInvoker) Invoke() (*model.ListIpGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIpGroupsResponse), nil
	}
}

type ShowIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpGroupInvoker) Invoke() (*model.ShowIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpGroupResponse), nil
	}
}

type UpdateIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpGroupInvoker) Invoke() (*model.UpdateIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpGroupResponse), nil
	}
}

type UpdateIpListInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpListInvoker) Invoke() (*model.UpdateIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpListResponse), nil
	}
}
