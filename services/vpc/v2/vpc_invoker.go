package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

type AcceptVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *AcceptVpcPeeringInvoker) Invoke() (*model.AcceptVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AcceptVpcPeeringResponse), nil
	}
}

type AssociateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateRouteTableInvoker) Invoke() (*model.AssociateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateRouteTableResponse), nil
	}
}

type BatchCreateSubnetTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSubnetTagsInvoker) Invoke() (*model.BatchCreateSubnetTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSubnetTagsResponse), nil
	}
}

type BatchDeleteSubnetTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSubnetTagsInvoker) Invoke() (*model.BatchDeleteSubnetTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSubnetTagsResponse), nil
	}
}

type CreatePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePortInvoker) Invoke() (*model.CreatePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePortResponse), nil
	}
}

type CreateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRouteTableInvoker) Invoke() (*model.CreateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRouteTableResponse), nil
	}
}

type CreateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityGroupInvoker) Invoke() (*model.CreateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityGroupResponse), nil
	}
}

type CreateSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityGroupRuleInvoker) Invoke() (*model.CreateSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityGroupRuleResponse), nil
	}
}

type CreateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubnetInvoker) Invoke() (*model.CreateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubnetResponse), nil
	}
}

type CreateSubnetTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSubnetTagInvoker) Invoke() (*model.CreateSubnetTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSubnetTagResponse), nil
	}
}

type CreateVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcPeeringInvoker) Invoke() (*model.CreateVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcPeeringResponse), nil
	}
}

type DeletePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePortInvoker) Invoke() (*model.DeletePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePortResponse), nil
	}
}

type DeleteRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRouteTableInvoker) Invoke() (*model.DeleteRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRouteTableResponse), nil
	}
}

type DeleteSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityGroupInvoker) Invoke() (*model.DeleteSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityGroupResponse), nil
	}
}

type DeleteSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityGroupRuleInvoker) Invoke() (*model.DeleteSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityGroupRuleResponse), nil
	}
}

type DeleteSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubnetInvoker) Invoke() (*model.DeleteSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubnetResponse), nil
	}
}

type DeleteSubnetTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSubnetTagInvoker) Invoke() (*model.DeleteSubnetTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSubnetTagResponse), nil
	}
}

type DeleteVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcPeeringInvoker) Invoke() (*model.DeleteVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcPeeringResponse), nil
	}
}

type DisassociateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateRouteTableInvoker) Invoke() (*model.DisassociateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateRouteTableResponse), nil
	}
}

type ListPortsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPortsInvoker) Invoke() (*model.ListPortsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPortsResponse), nil
	}
}

type ListRouteTablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRouteTablesInvoker) Invoke() (*model.ListRouteTablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRouteTablesResponse), nil
	}
}

type ListSecurityGroupRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupRulesInvoker) Invoke() (*model.ListSecurityGroupRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupRulesResponse), nil
	}
}

type ListSecurityGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupsInvoker) Invoke() (*model.ListSecurityGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupsResponse), nil
	}
}

type ListSubnetTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubnetTagsInvoker) Invoke() (*model.ListSubnetTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubnetTagsResponse), nil
	}
}

type ListSubnetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubnetsInvoker) Invoke() (*model.ListSubnetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubnetsResponse), nil
	}
}

type ListSubnetsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSubnetsByTagsInvoker) Invoke() (*model.ListSubnetsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSubnetsByTagsResponse), nil
	}
}

type ListVpcPeeringsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcPeeringsInvoker) Invoke() (*model.ListVpcPeeringsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcPeeringsResponse), nil
	}
}

type RejectVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *RejectVpcPeeringInvoker) Invoke() (*model.RejectVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RejectVpcPeeringResponse), nil
	}
}

type ShowPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPortInvoker) Invoke() (*model.ShowPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPortResponse), nil
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

type ShowRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRouteTableInvoker) Invoke() (*model.ShowRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRouteTableResponse), nil
	}
}

type ShowSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityGroupInvoker) Invoke() (*model.ShowSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityGroupResponse), nil
	}
}

type ShowSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityGroupRuleInvoker) Invoke() (*model.ShowSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityGroupRuleResponse), nil
	}
}

type ShowSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubnetInvoker) Invoke() (*model.ShowSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubnetResponse), nil
	}
}

type ShowSubnetTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSubnetTagsInvoker) Invoke() (*model.ShowSubnetTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSubnetTagsResponse), nil
	}
}

type ShowVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcPeeringInvoker) Invoke() (*model.ShowVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcPeeringResponse), nil
	}
}

type UpdatePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePortInvoker) Invoke() (*model.UpdatePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePortResponse), nil
	}
}

type UpdateRouteTableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRouteTableInvoker) Invoke() (*model.UpdateRouteTableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRouteTableResponse), nil
	}
}

type UpdateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateSubnetInvoker) Invoke() (*model.UpdateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateSubnetResponse), nil
	}
}

type UpdateVpcPeeringInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcPeeringInvoker) Invoke() (*model.UpdateVpcPeeringResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcPeeringResponse), nil
	}
}

type CreatePrivateipInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePrivateipInvoker) Invoke() (*model.CreatePrivateipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePrivateipResponse), nil
	}
}

type DeletePrivateipInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePrivateipInvoker) Invoke() (*model.DeletePrivateipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePrivateipResponse), nil
	}
}

type ListPrivateipsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPrivateipsInvoker) Invoke() (*model.ListPrivateipsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPrivateipsResponse), nil
	}
}

type ShowNetworkIpAvailabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkIpAvailabilitiesInvoker) Invoke() (*model.ShowNetworkIpAvailabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkIpAvailabilitiesResponse), nil
	}
}

type ShowPrivateipInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPrivateipInvoker) Invoke() (*model.ShowPrivateipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPrivateipResponse), nil
	}
}

type NeutronCreateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateSecurityGroupInvoker) Invoke() (*model.NeutronCreateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateSecurityGroupResponse), nil
	}
}

type NeutronCreateSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateSecurityGroupRuleInvoker) Invoke() (*model.NeutronCreateSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateSecurityGroupRuleResponse), nil
	}
}

type NeutronDeleteSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteSecurityGroupInvoker) Invoke() (*model.NeutronDeleteSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteSecurityGroupResponse), nil
	}
}

type NeutronDeleteSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteSecurityGroupRuleInvoker) Invoke() (*model.NeutronDeleteSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteSecurityGroupRuleResponse), nil
	}
}

type NeutronListSecurityGroupRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListSecurityGroupRulesInvoker) Invoke() (*model.NeutronListSecurityGroupRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListSecurityGroupRulesResponse), nil
	}
}

type NeutronListSecurityGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListSecurityGroupsInvoker) Invoke() (*model.NeutronListSecurityGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListSecurityGroupsResponse), nil
	}
}

type NeutronShowSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowSecurityGroupInvoker) Invoke() (*model.NeutronShowSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowSecurityGroupResponse), nil
	}
}

type NeutronShowSecurityGroupRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowSecurityGroupRuleInvoker) Invoke() (*model.NeutronShowSecurityGroupRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowSecurityGroupRuleResponse), nil
	}
}

type NeutronUpdateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateSecurityGroupInvoker) Invoke() (*model.NeutronUpdateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateSecurityGroupResponse), nil
	}
}

type NeutronAddFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronAddFirewallRuleInvoker) Invoke() (*model.NeutronAddFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronAddFirewallRuleResponse), nil
	}
}

type NeutronCreateFirewallGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateFirewallGroupInvoker) Invoke() (*model.NeutronCreateFirewallGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateFirewallGroupResponse), nil
	}
}

type NeutronCreateFirewallPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateFirewallPolicyInvoker) Invoke() (*model.NeutronCreateFirewallPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateFirewallPolicyResponse), nil
	}
}

type NeutronCreateFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateFirewallRuleInvoker) Invoke() (*model.NeutronCreateFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateFirewallRuleResponse), nil
	}
}

type NeutronDeleteFirewallGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteFirewallGroupInvoker) Invoke() (*model.NeutronDeleteFirewallGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteFirewallGroupResponse), nil
	}
}

type NeutronDeleteFirewallPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteFirewallPolicyInvoker) Invoke() (*model.NeutronDeleteFirewallPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteFirewallPolicyResponse), nil
	}
}

type NeutronDeleteFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteFirewallRuleInvoker) Invoke() (*model.NeutronDeleteFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteFirewallRuleResponse), nil
	}
}

type NeutronListFirewallGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListFirewallGroupsInvoker) Invoke() (*model.NeutronListFirewallGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListFirewallGroupsResponse), nil
	}
}

type NeutronListFirewallPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListFirewallPoliciesInvoker) Invoke() (*model.NeutronListFirewallPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListFirewallPoliciesResponse), nil
	}
}

type NeutronListFirewallRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListFirewallRulesInvoker) Invoke() (*model.NeutronListFirewallRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListFirewallRulesResponse), nil
	}
}

type NeutronRemoveFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronRemoveFirewallRuleInvoker) Invoke() (*model.NeutronRemoveFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronRemoveFirewallRuleResponse), nil
	}
}

type NeutronShowFirewallGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowFirewallGroupInvoker) Invoke() (*model.NeutronShowFirewallGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowFirewallGroupResponse), nil
	}
}

type NeutronShowFirewallPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowFirewallPolicyInvoker) Invoke() (*model.NeutronShowFirewallPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowFirewallPolicyResponse), nil
	}
}

type NeutronShowFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowFirewallRuleInvoker) Invoke() (*model.NeutronShowFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowFirewallRuleResponse), nil
	}
}

type NeutronUpdateFirewallGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateFirewallGroupInvoker) Invoke() (*model.NeutronUpdateFirewallGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateFirewallGroupResponse), nil
	}
}

type NeutronUpdateFirewallPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateFirewallPolicyInvoker) Invoke() (*model.NeutronUpdateFirewallPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateFirewallPolicyResponse), nil
	}
}

type NeutronUpdateFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateFirewallRuleInvoker) Invoke() (*model.NeutronUpdateFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateFirewallRuleResponse), nil
	}
}

type BatchCreateVpcTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateVpcTagsInvoker) Invoke() (*model.BatchCreateVpcTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateVpcTagsResponse), nil
	}
}

type BatchDeleteVpcTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteVpcTagsInvoker) Invoke() (*model.BatchDeleteVpcTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteVpcTagsResponse), nil
	}
}

type CreateVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcInvoker) Invoke() (*model.CreateVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcResponse), nil
	}
}

type CreateVpcResourceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcResourceTagInvoker) Invoke() (*model.CreateVpcResourceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcResourceTagResponse), nil
	}
}

type CreateVpcRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpcRouteInvoker) Invoke() (*model.CreateVpcRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpcRouteResponse), nil
	}
}

type DeleteVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcInvoker) Invoke() (*model.DeleteVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcResponse), nil
	}
}

type DeleteVpcRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcRouteInvoker) Invoke() (*model.DeleteVpcRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcRouteResponse), nil
	}
}

type DeleteVpcTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpcTagInvoker) Invoke() (*model.DeleteVpcTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpcTagResponse), nil
	}
}

type ListVpcRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcRoutesInvoker) Invoke() (*model.ListVpcRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcRoutesResponse), nil
	}
}

type ListVpcTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcTagsInvoker) Invoke() (*model.ListVpcTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcTagsResponse), nil
	}
}

type ListVpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcsInvoker) Invoke() (*model.ListVpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcsResponse), nil
	}
}

type ListVpcsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpcsByTagsInvoker) Invoke() (*model.ListVpcsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpcsByTagsResponse), nil
	}
}

type ShowVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcInvoker) Invoke() (*model.ShowVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcResponse), nil
	}
}

type ShowVpcRouteInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcRouteInvoker) Invoke() (*model.ShowVpcRouteResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcRouteResponse), nil
	}
}

type ShowVpcTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpcTagsInvoker) Invoke() (*model.ShowVpcTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpcTagsResponse), nil
	}
}

type UpdateVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpcInvoker) Invoke() (*model.UpdateVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpcResponse), nil
	}
}
