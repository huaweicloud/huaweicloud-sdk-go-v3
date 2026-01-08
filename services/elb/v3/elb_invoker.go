package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/elb/v3/model"
)

type BatchAddAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddAvailableZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddAvailableZonesInvoker) Invoke() (*model.BatchAddAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddAvailableZonesResponse), nil
	}
}

type BatchCreateLoadBalancersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateLoadBalancersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateLoadBalancersInvoker) Invoke() (*model.BatchCreateLoadBalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateLoadBalancersResponse), nil
	}
}

type BatchCreateMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateMembersInvoker) Invoke() (*model.BatchCreateMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateMembersResponse), nil
	}
}

type BatchDeleteCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteCertificatesInvoker) Invoke() (*model.BatchDeleteCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteCertificatesResponse), nil
	}
}

type BatchDeleteListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteListenersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteListenersInvoker) Invoke() (*model.BatchDeleteListenersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteListenersResponse), nil
	}
}

type BatchDeleteLoadbalancersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteLoadbalancersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteLoadbalancersInvoker) Invoke() (*model.BatchDeleteLoadbalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteLoadbalancersResponse), nil
	}
}

type BatchDeleteMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteMembersInvoker) Invoke() (*model.BatchDeleteMembersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteMembersResponse), nil
	}
}

type BatchDeletePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeletePoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeletePoolsInvoker) Invoke() (*model.BatchDeletePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeletePoolsResponse), nil
	}
}

type BatchDisableDomainIPsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisableDomainIPsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisableDomainIPsInvoker) Invoke() (*model.BatchDisableDomainIPsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisableDomainIPsResponse), nil
	}
}

type BatchEnableDomainIPsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableDomainIPsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableDomainIPsInvoker) Invoke() (*model.BatchEnableDomainIPsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableDomainIPsResponse), nil
	}
}

type BatchRemoveAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRemoveAvailableZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchRemoveAvailableZonesInvoker) Invoke() (*model.BatchRemoveAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRemoveAvailableZonesResponse), nil
	}
}

type BatchUpdateMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchUpdatePoliciesPriorityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdatePoliciesPriorityInvoker) Invoke() (*model.BatchUpdatePoliciesPriorityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePoliciesPriorityResponse), nil
	}
}

type ChangeListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeListenerTagsInvoker) Invoke() (*model.ChangeListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeListenerTagsResponse), nil
	}
}

type ChangeLoadbalancerChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeLoadbalancerChargeModeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeLoadbalancerChargeModeInvoker) Invoke() (*model.ChangeLoadbalancerChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeLoadbalancerChargeModeResponse), nil
	}
}

type ChangeLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeLoadbalancerTagsInvoker) Invoke() (*model.ChangeLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeLoadbalancerTagsResponse), nil
	}
}

type CloneListenerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloneListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloneListenerInvoker) Invoke() (*model.CloneListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloneListenerResponse), nil
	}
}

type CloneLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CloneLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CloneLoadbalancerInvoker) Invoke() (*model.CloneLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloneLoadbalancerResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type CreateCertificatePrivateKeyEchoInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificatePrivateKeyEchoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCertificatePrivateKeyEchoInvoker) Invoke() (*model.CreateCertificatePrivateKeyEchoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificatePrivateKeyEchoResponse), nil
	}
}

type CreateHealthMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHealthMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateL7PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateL7RuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateLoadBalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLogtankInvoker) Invoke() (*model.CreateLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLogtankResponse), nil
	}
}

type CreateMasterSlavePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMasterSlavePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMasterSlavePoolInvoker) Invoke() (*model.CreateMasterSlavePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMasterSlavePoolResponse), nil
	}
}

type CreateMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMemberInvoker) Invoke() (*model.CreateMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMemberResponse), nil
	}
}

type CreateMemberHealthCheckJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateMemberHealthCheckJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateMemberHealthCheckJobInvoker) Invoke() (*model.CreateMemberHealthCheckJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateMemberHealthCheckJobResponse), nil
	}
}

type CreatePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateSecurityPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteHealthMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteL7PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteL7RuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteListenerForceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteLoadBalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLoadBalancerInvoker) Invoke() (*model.DeleteLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadBalancerResponse), nil
	}
}

type DeleteLoadBalancerCascadeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadBalancerCascadeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLoadBalancerCascadeInvoker) Invoke() (*model.DeleteLoadBalancerCascadeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLoadBalancerCascadeResponse), nil
	}
}

type DeleteLoadBalancerForceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLoadBalancerForceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLogtankInvoker) Invoke() (*model.DeleteLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLogtankResponse), nil
	}
}

type DeleteMasterSlavePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMasterSlavePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteMasterSlavePoolInvoker) Invoke() (*model.DeleteMasterSlavePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteMasterSlavePoolResponse), nil
	}
}

type DeleteMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeletePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePoolInvoker) Invoke() (*model.DeletePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoolResponse), nil
	}
}

type DeletePoolCascadeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePoolCascadeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePoolCascadeInvoker) Invoke() (*model.DeletePoolCascadeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePoolCascadeResponse), nil
	}
}

type DeleteRecycleLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRecycleLoadBalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRecycleLoadBalancerInvoker) Invoke() (*model.DeleteRecycleLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRecycleLoadBalancerResponse), nil
	}
}

type DeleteSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSecurityPolicyInvoker) Invoke() (*model.DeleteSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityPolicyResponse), nil
	}
}

type ListAllL7RulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllL7RulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllL7RulesInvoker) Invoke() (*model.ListAllL7RulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllL7RulesResponse), nil
	}
}

type ListAllMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListAvailabilityZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListCertificatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type ListDomainIPsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainIPsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDomainIPsInvoker) Invoke() (*model.ListDomainIPsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainIPsResponse), nil
	}
}

type ListFeatureConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFeatureConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFeatureConfigsInvoker) Invoke() (*model.ListFeatureConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFeatureConfigsResponse), nil
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

type ListHealthMonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHealthMonitorsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHealthMonitorsInvoker) Invoke() (*model.ListHealthMonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHealthMonitorsResponse), nil
	}
}

type ListJobsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListJobsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListJobsInvoker) Invoke() (*model.ListJobsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListJobsResponse), nil
	}
}

type ListL7PoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListL7PoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListL7RulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListL7RulesInvoker) Invoke() (*model.ListL7RulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListL7RulesResponse), nil
	}
}

type ListListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListListenerTagsInvoker) Invoke() (*model.ListListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListListenerTagsResponse), nil
	}
}

type ListListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListListenersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListLoadBalancersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadBalancersInvoker) Invoke() (*model.ListLoadBalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadBalancersResponse), nil
	}
}

type ListLoadbalancerFeatureInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadbalancerFeatureInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadbalancerFeatureInvoker) Invoke() (*model.ListLoadbalancerFeatureResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadbalancerFeatureResponse), nil
	}
}

type ListLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLoadbalancerTagsInvoker) Invoke() (*model.ListLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLoadbalancerTagsResponse), nil
	}
}

type ListLogtanksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListLogtanksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListLogtanksInvoker) Invoke() (*model.ListLogtanksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListLogtanksResponse), nil
	}
}

type ListMasterSlavePoolsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMasterSlavePoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMasterSlavePoolsInvoker) Invoke() (*model.ListMasterSlavePoolsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMasterSlavePoolsResponse), nil
	}
}

type ListMembersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMembersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListPoolsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListQuotaDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaDetailsInvoker) Invoke() (*model.ListQuotaDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaDetailsResponse), nil
	}
}

type ListRecycleBinLoadBalancersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRecycleBinLoadBalancersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRecycleBinLoadBalancersInvoker) Invoke() (*model.ListRecycleBinLoadBalancersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRecycleBinLoadBalancersResponse), nil
	}
}

type ListSecurityPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListSystemSecurityPoliciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSystemSecurityPoliciesInvoker) Invoke() (*model.ListSystemSecurityPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSystemSecurityPoliciesResponse), nil
	}
}

type RestoreLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestoreLoadbalancerInvoker) Invoke() (*model.RestoreLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreLoadbalancerResponse), nil
	}
}

type ShowCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificateInvoker) Invoke() (*model.ShowCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificateResponse), nil
	}
}

type ShowCertificatePrivateKeyEchoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCertificatePrivateKeyEchoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCertificatePrivateKeyEchoInvoker) Invoke() (*model.ShowCertificatePrivateKeyEchoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCertificatePrivateKeyEchoResponse), nil
	}
}

type ShowFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlavorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowHealthMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHealthMonitorInvoker) Invoke() (*model.ShowHealthMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHealthMonitorResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type ShowL7PolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowL7PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowL7RuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListenerInvoker) Invoke() (*model.ShowListenerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerResponse), nil
	}
}

type ShowListenerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowListenerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowListenerTagsInvoker) Invoke() (*model.ShowListenerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowListenerTagsResponse), nil
	}
}

type ShowLoadBalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadBalancerInvoker) Invoke() (*model.ShowLoadBalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerResponse), nil
	}
}

type ShowLoadBalancerPortsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerPortsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadBalancerPortsInvoker) Invoke() (*model.ShowLoadBalancerPortsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerPortsResponse), nil
	}
}

type ShowLoadBalancerStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadBalancerStatusInvoker) Invoke() (*model.ShowLoadBalancerStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerStatusResponse), nil
	}
}

type ShowLoadBalancerTopologyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadBalancerTopologyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadBalancerTopologyInvoker) Invoke() (*model.ShowLoadBalancerTopologyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadBalancerTopologyResponse), nil
	}
}

type ShowLoadbalancerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLoadbalancerTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLoadbalancerTagsInvoker) Invoke() (*model.ShowLoadbalancerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLoadbalancerTagsResponse), nil
	}
}

type ShowLogtankInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLogtankInvoker) Invoke() (*model.ShowLogtankResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogtankResponse), nil
	}
}

type ShowMasterSlavePoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMasterSlavePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMasterSlavePoolInvoker) Invoke() (*model.ShowMasterSlavePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMasterSlavePoolResponse), nil
	}
}

type ShowMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMemberInvoker) Invoke() (*model.ShowMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberResponse), nil
	}
}

type ShowMemberHealthCheckJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberHealthCheckJobInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowMemberHealthCheckJobInvoker) Invoke() (*model.ShowMemberHealthCheckJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberHealthCheckJobResponse), nil
	}
}

type ShowPoolInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

type ShowRecycleBinInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRecycleBinInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRecycleBinInvoker) Invoke() (*model.ShowRecycleBinResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRecycleBinResponse), nil
	}
}

type ShowSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateHealthMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateL7PolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateL7RuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateListenerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateLoadBalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateLogtankInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateMemberInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdatePoolInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePoolInvoker) Invoke() (*model.UpdatePoolResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePoolResponse), nil
	}
}

type UpdateRecycleBinEnableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecycleBinEnableInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecycleBinEnableInvoker) Invoke() (*model.UpdateRecycleBinEnableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecycleBinEnableResponse), nil
	}
}

type UpdateRecycleBinPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRecycleBinPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRecycleBinPolicyInvoker) Invoke() (*model.UpdateRecycleBinPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRecycleBinPolicyResponse), nil
	}
}

type UpdateSecurityPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSecurityPolicyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSecurityPolicyInvoker) Invoke() (*model.UpdateSecurityPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSecurityPolicyResponse), nil
	}
}

type UpdateSystemDefaultDomainConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSystemDefaultDomainConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateSystemDefaultDomainConfigInvoker) Invoke() (*model.UpdateSystemDefaultDomainConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSystemDefaultDomainConfigResponse), nil
	}
}

type UpdateUserDefinedDomainConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserDefinedDomainConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateUserDefinedDomainConfigInvoker) Invoke() (*model.UpdateUserDefinedDomainConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserDefinedDomainConfigResponse), nil
	}
}

type UpgradeLoadbalancerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeLoadbalancerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeLoadbalancerInvoker) Invoke() (*model.UpgradeLoadbalancerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeLoadbalancerResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *BatchDeleteIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CountPreoccupyIpNumInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeleteIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListIpGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpGroupInvoker) Invoke() (*model.ShowIpGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpGroupResponse), nil
	}
}

type ShowIpGroupRelatedListenersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpGroupRelatedListenersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpGroupRelatedListenersInvoker) Invoke() (*model.ShowIpGroupRelatedListenersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpGroupRelatedListenersResponse), nil
	}
}

type UpdateIpGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIpGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateIpListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIpListInvoker) Invoke() (*model.UpdateIpListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIpListResponse), nil
	}
}
