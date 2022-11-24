package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iec/v1/model"
)

type AddNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddNicsInvoker) Invoke() (*model.AddNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddNicsResponse), nil
	}
}

type AssociateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateSubnetInvoker) Invoke() (*model.AssociateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateSubnetResponse), nil
	}
}

type BatchRebootInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebootInstanceInvoker) Invoke() (*model.BatchRebootInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebootInstanceResponse), nil
	}
}

type BatchStartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartInstanceInvoker) Invoke() (*model.BatchStartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartInstanceResponse), nil
	}
}

type BatchStopInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopInstanceInvoker) Invoke() (*model.BatchStopInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopInstanceResponse), nil
	}
}

type ChangeOsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeOsInvoker) Invoke() (*model.ChangeOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeOsResponse), nil
	}
}

type CreateDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentInvoker) Invoke() (*model.CreateDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentResponse), nil
	}
}

type CreateKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateKeypairInvoker) Invoke() (*model.CreateKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateKeypairResponse), nil
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

type CreateRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRoutesInvoker) Invoke() (*model.CreateRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRoutesResponse), nil
	}
}

type CreateRoutetableInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRoutetableInvoker) Invoke() (*model.CreateRoutetableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRoutetableResponse), nil
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

type DeleteDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentInvoker) Invoke() (*model.DeleteDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentResponse), nil
	}
}

type DeleteEdgeCloudInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeCloudInvoker) Invoke() (*model.DeleteEdgeCloudResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeCloudResponse), nil
	}
}

type DeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstancesInvoker) Invoke() (*model.DeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesResponse), nil
	}
}

type DeleteKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteKeypairInvoker) Invoke() (*model.DeleteKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteKeypairResponse), nil
	}
}

type DeleteNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNicsInvoker) Invoke() (*model.DeleteNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNicsResponse), nil
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

type DeleteRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRoutesInvoker) Invoke() (*model.DeleteRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRoutesResponse), nil
	}
}

type DeleteRoutetableInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRoutetableInvoker) Invoke() (*model.DeleteRoutetableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRoutetableResponse), nil
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

type DisassociateSubnetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateSubnetInvoker) Invoke() (*model.DisassociateSubnetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateSubnetResponse), nil
	}
}

type ExecuteDeploymentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDeploymentInvoker) Invoke() (*model.ExecuteDeploymentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDeploymentResponse), nil
	}
}

type ExpandEdgecloudInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandEdgecloudInvoker) Invoke() (*model.ExpandEdgecloudResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandEdgecloudResponse), nil
	}
}

type ListBandwidthsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBandwidthsInvoker) Invoke() (*model.ListBandwidthsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBandwidthsResponse), nil
	}
}

type ListDeploymentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeploymentsInvoker) Invoke() (*model.ListDeploymentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeploymentsResponse), nil
	}
}

type ListEdgeCloudInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeCloudInvoker) Invoke() (*model.ListEdgeCloudResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeCloudResponse), nil
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

type ListImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListImagesInvoker) Invoke() (*model.ListImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListImagesResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListKeypairsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListKeypairsInvoker) Invoke() (*model.ListKeypairsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListKeypairsResponse), nil
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

type ListQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaInvoker) Invoke() (*model.ListQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaResponse), nil
	}
}

type ListRelatedRoutetablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRelatedRoutetablesInvoker) Invoke() (*model.ListRelatedRoutetablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRelatedRoutetablesResponse), nil
	}
}

type ListRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRoutesInvoker) Invoke() (*model.ListRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRoutesResponse), nil
	}
}

type ListRoutetablesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRoutetablesInvoker) Invoke() (*model.ListRoutetablesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRoutetablesResponse), nil
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

type ListSitesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSitesInvoker) Invoke() (*model.ListSitesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSitesResponse), nil
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

type ShowBandwidthInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBandwidthInvoker) Invoke() (*model.ShowBandwidthResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBandwidthResponse), nil
	}
}

type ShowEdgeCloudInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeCloudInvoker) Invoke() (*model.ShowEdgeCloudResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeCloudResponse), nil
	}
}

type ShowImageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowImageInvoker) Invoke() (*model.ShowImageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowImageResponse), nil
	}
}

type ShowInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceInvoker) Invoke() (*model.ShowInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceResponse), nil
	}
}

type ShowKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowKeypairInvoker) Invoke() (*model.ShowKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowKeypairResponse), nil
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

type ShowRoutetableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRoutetableInvoker) Invoke() (*model.ShowRoutetableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRoutetableResponse), nil
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

type ShowVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVolumeInvoker) Invoke() (*model.ShowVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVolumeResponse), nil
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

type UpdateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceInvoker) Invoke() (*model.UpdateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceResponse), nil
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

type UpdateRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoutesInvoker) Invoke() (*model.UpdateRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoutesResponse), nil
	}
}

type UpdateRoutetableInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoutetableInvoker) Invoke() (*model.UpdateRoutetableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoutetableResponse), nil
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

type CreateFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFirewallInvoker) Invoke() (*model.CreateFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFirewallResponse), nil
	}
}

type DeleteFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFirewallInvoker) Invoke() (*model.DeleteFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFirewallResponse), nil
	}
}

type ListFirewallsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFirewallsInvoker) Invoke() (*model.ListFirewallsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFirewallsResponse), nil
	}
}

type ShowFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFirewallInvoker) Invoke() (*model.ShowFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFirewallResponse), nil
	}
}

type UpdateFirewallInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFirewallInvoker) Invoke() (*model.UpdateFirewallResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFirewallResponse), nil
	}
}

type UpdateFirewallRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFirewallRuleInvoker) Invoke() (*model.UpdateFirewallRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFirewallRuleResponse), nil
	}
}

type CreatePublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePublicIpInvoker) Invoke() (*model.CreatePublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePublicIpResponse), nil
	}
}

type DeletePublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePublicIpInvoker) Invoke() (*model.DeletePublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePublicIpResponse), nil
	}
}

type ListPublicIpsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPublicIpsInvoker) Invoke() (*model.ListPublicIpsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPublicIpsResponse), nil
	}
}

type ShowPublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPublicIpInvoker) Invoke() (*model.ShowPublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPublicIpResponse), nil
	}
}

type UpdatePublicIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePublicIpInvoker) Invoke() (*model.UpdatePublicIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePublicIpResponse), nil
	}
}
