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

type BatchCreateSecurityGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateSecurityGroupTagsInvoker) Invoke() (*model.BatchCreateSecurityGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateSecurityGroupTagsResponse), nil
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

type BatchDeleteSecurityGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteSecurityGroupTagsInvoker) Invoke() (*model.BatchDeleteSecurityGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteSecurityGroupTagsResponse), nil
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

type CreateFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFlowLogInvoker) Invoke() (*model.CreateFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFlowLogResponse), nil
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

type CreateSecurityGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSecurityGroupTagInvoker) Invoke() (*model.CreateSecurityGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSecurityGroupTagResponse), nil
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

type DeleteFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFlowLogInvoker) Invoke() (*model.DeleteFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFlowLogResponse), nil
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

type DeleteSecurityGroupTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSecurityGroupTagInvoker) Invoke() (*model.DeleteSecurityGroupTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSecurityGroupTagResponse), nil
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

type ListFlowLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlowLogsInvoker) Invoke() (*model.ListFlowLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlowLogsResponse), nil
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

type ListSecurityGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupTagsInvoker) Invoke() (*model.ListSecurityGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupTagsResponse), nil
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

type ListSecurityGroupsByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSecurityGroupsByTagsInvoker) Invoke() (*model.ListSecurityGroupsByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSecurityGroupsByTagsResponse), nil
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

type ShowFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFlowLogInvoker) Invoke() (*model.ShowFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFlowLogResponse), nil
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

type ShowSecurityGroupTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSecurityGroupTagsInvoker) Invoke() (*model.ShowSecurityGroupTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSecurityGroupTagsResponse), nil
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

type UpdateFlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFlowLogInvoker) Invoke() (*model.UpdateFlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFlowLogResponse), nil
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

type NeutronAddRouterInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronAddRouterInterfaceInvoker) Invoke() (*model.NeutronAddRouterInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronAddRouterInterfaceResponse), nil
	}
}

type NeutronCreateNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateNetworkInvoker) Invoke() (*model.NeutronCreateNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateNetworkResponse), nil
	}
}

type NeutronCreatePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreatePortInvoker) Invoke() (*model.NeutronCreatePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreatePortResponse), nil
	}
}

type NeutronCreateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateRouterInvoker) Invoke() (*model.NeutronCreateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateRouterResponse), nil
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

type NeutronCreateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronCreateSubnetInvoker) Invoke() (*model.NeutronCreateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronCreateSubnetResponse), nil
	}
}

type NeutronDeleteNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteNetworkInvoker) Invoke() (*model.NeutronDeleteNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteNetworkResponse), nil
	}
}

type NeutronDeletePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeletePortInvoker) Invoke() (*model.NeutronDeletePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeletePortResponse), nil
	}
}

type NeutronDeleteRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteRouterInvoker) Invoke() (*model.NeutronDeleteRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteRouterResponse), nil
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

type NeutronDeleteSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronDeleteSubnetInvoker) Invoke() (*model.NeutronDeleteSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronDeleteSubnetResponse), nil
	}
}

type NeutronListNetworksInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListNetworksInvoker) Invoke() (*model.NeutronListNetworksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListNetworksResponse), nil
	}
}

type NeutronListPortsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListPortsInvoker) Invoke() (*model.NeutronListPortsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListPortsResponse), nil
	}
}

type NeutronListRoutersInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListRoutersInvoker) Invoke() (*model.NeutronListRoutersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListRoutersResponse), nil
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

type NeutronListSubnetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronListSubnetsInvoker) Invoke() (*model.NeutronListSubnetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronListSubnetsResponse), nil
	}
}

type NeutronRemoveRouterInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronRemoveRouterInterfaceInvoker) Invoke() (*model.NeutronRemoveRouterInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronRemoveRouterInterfaceResponse), nil
	}
}

type NeutronShowNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowNetworkInvoker) Invoke() (*model.NeutronShowNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowNetworkResponse), nil
	}
}

type NeutronShowPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowPortInvoker) Invoke() (*model.NeutronShowPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowPortResponse), nil
	}
}

type NeutronShowRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowRouterInvoker) Invoke() (*model.NeutronShowRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowRouterResponse), nil
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

type NeutronShowSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronShowSubnetInvoker) Invoke() (*model.NeutronShowSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronShowSubnetResponse), nil
	}
}

type NeutronUpdateNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateNetworkInvoker) Invoke() (*model.NeutronUpdateNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateNetworkResponse), nil
	}
}

type NeutronUpdatePortInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdatePortInvoker) Invoke() (*model.NeutronUpdatePortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdatePortResponse), nil
	}
}

type NeutronUpdateRouterInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateRouterInvoker) Invoke() (*model.NeutronUpdateRouterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateRouterResponse), nil
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

type NeutronUpdateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *NeutronUpdateSubnetInvoker) Invoke() (*model.NeutronUpdateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NeutronUpdateSubnetResponse), nil
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

type ListApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionInvoker) Invoke() (*model.ListApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionResponse), nil
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
