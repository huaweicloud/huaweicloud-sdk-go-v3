package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

type AddServerGroupMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServerGroupMemberInvoker) Invoke() (*model.AddServerGroupMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServerGroupMemberResponse), nil
	}
}

type AssociateServerVirtualIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateServerVirtualIpInvoker) Invoke() (*model.AssociateServerVirtualIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateServerVirtualIpResponse), nil
	}
}

type AttachServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachServerVolumeInvoker) Invoke() (*model.AttachServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachServerVolumeResponse), nil
	}
}

type BatchAddServerNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddServerNicsInvoker) Invoke() (*model.BatchAddServerNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddServerNicsResponse), nil
	}
}

type BatchAttachSharableVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAttachSharableVolumesInvoker) Invoke() (*model.BatchAttachSharableVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAttachSharableVolumesResponse), nil
	}
}

type BatchCreateServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateServerTagsInvoker) Invoke() (*model.BatchCreateServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateServerTagsResponse), nil
	}
}

type BatchDeleteServerNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteServerNicsInvoker) Invoke() (*model.BatchDeleteServerNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteServerNicsResponse), nil
	}
}

type BatchDeleteServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteServerTagsInvoker) Invoke() (*model.BatchDeleteServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteServerTagsResponse), nil
	}
}

type BatchRebootServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebootServersInvoker) Invoke() (*model.BatchRebootServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebootServersResponse), nil
	}
}

type BatchResetServersPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchResetServersPasswordInvoker) Invoke() (*model.BatchResetServersPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchResetServersPasswordResponse), nil
	}
}

type BatchStartServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartServersInvoker) Invoke() (*model.BatchStartServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartServersResponse), nil
	}
}

type BatchStopServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopServersInvoker) Invoke() (*model.BatchStopServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopServersResponse), nil
	}
}

type BatchUpdateServersNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateServersNameInvoker) Invoke() (*model.BatchUpdateServersNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateServersNameResponse), nil
	}
}

type ChangeServerChargeModeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerChargeModeInvoker) Invoke() (*model.ChangeServerChargeModeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerChargeModeResponse), nil
	}
}

type ChangeServerNetworkInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerNetworkInterfaceInvoker) Invoke() (*model.ChangeServerNetworkInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerNetworkInterfaceResponse), nil
	}
}

type ChangeServerOsWithCloudInitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerOsWithCloudInitInvoker) Invoke() (*model.ChangeServerOsWithCloudInitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerOsWithCloudInitResponse), nil
	}
}

type ChangeServerOsWithoutCloudInitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerOsWithoutCloudInitInvoker) Invoke() (*model.ChangeServerOsWithoutCloudInitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerOsWithoutCloudInitResponse), nil
	}
}

type ChangeVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeVpcInvoker) Invoke() (*model.ChangeVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeVpcResponse), nil
	}
}

type CreatePostPaidServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidServersInvoker) Invoke() (*model.CreatePostPaidServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidServersResponse), nil
	}
}

type CreateServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServerGroupInvoker) Invoke() (*model.CreateServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServerGroupResponse), nil
	}
}

type CreateServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateServersInvoker) Invoke() (*model.CreateServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateServersResponse), nil
	}
}

type DeleteServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerGroupInvoker) Invoke() (*model.DeleteServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerGroupResponse), nil
	}
}

type DeleteServerGroupMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerGroupMemberInvoker) Invoke() (*model.DeleteServerGroupMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerGroupMemberResponse), nil
	}
}

type DeleteServerMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerMetadataInvoker) Invoke() (*model.DeleteServerMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerMetadataResponse), nil
	}
}

type DeleteServerPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerPasswordInvoker) Invoke() (*model.DeleteServerPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerPasswordResponse), nil
	}
}

type DeleteServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServersInvoker) Invoke() (*model.DeleteServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServersResponse), nil
	}
}

type DetachServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachServerVolumeInvoker) Invoke() (*model.DetachServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachServerVolumeResponse), nil
	}
}

type DisassociateServerVirtualIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateServerVirtualIpInvoker) Invoke() (*model.DisassociateServerVirtualIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateServerVirtualIpResponse), nil
	}
}

type ListCloudServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCloudServersInvoker) Invoke() (*model.ListCloudServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCloudServersResponse), nil
	}
}

type ListFlavorSellPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFlavorSellPoliciesInvoker) Invoke() (*model.ListFlavorSellPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFlavorSellPoliciesResponse), nil
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

type ListResizeFlavorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResizeFlavorsInvoker) Invoke() (*model.ListResizeFlavorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResizeFlavorsResponse), nil
	}
}

type ListServerBlockDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerBlockDevicesInvoker) Invoke() (*model.ListServerBlockDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerBlockDevicesResponse), nil
	}
}

type ListServerGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerGroupsInvoker) Invoke() (*model.ListServerGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerGroupsResponse), nil
	}
}

type ListServerInterfacesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerInterfacesInvoker) Invoke() (*model.ListServerInterfacesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerInterfacesResponse), nil
	}
}

type ListServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServerTagsInvoker) Invoke() (*model.ListServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServerTagsResponse), nil
	}
}

type ListServersByTagInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ListServersByTagInvoker) Invoke() (*model.ListServersByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersByTagResponse), nil
	}
}

type ListServersDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersDetailsInvoker) Invoke() (*model.ListServersDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersDetailsResponse), nil
	}
}

type MigrateServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateServerInvoker) Invoke() (*model.MigrateServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateServerResponse), nil
	}
}

type NovaAssociateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaAssociateSecurityGroupInvoker) Invoke() (*model.NovaAssociateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaAssociateSecurityGroupResponse), nil
	}
}

type NovaAttachInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaAttachInterfaceInvoker) Invoke() (*model.NovaAttachInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaAttachInterfaceResponse), nil
	}
}

type NovaCreateKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaCreateKeypairInvoker) Invoke() (*model.NovaCreateKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaCreateKeypairResponse), nil
	}
}

type NovaCreateServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaCreateServersInvoker) Invoke() (*model.NovaCreateServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaCreateServersResponse), nil
	}
}

type NovaDeleteKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaDeleteKeypairInvoker) Invoke() (*model.NovaDeleteKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaDeleteKeypairResponse), nil
	}
}

type NovaDeleteServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaDeleteServerInvoker) Invoke() (*model.NovaDeleteServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaDeleteServerResponse), nil
	}
}

type NovaDisassociateSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaDisassociateSecurityGroupInvoker) Invoke() (*model.NovaDisassociateSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaDisassociateSecurityGroupResponse), nil
	}
}

type NovaListAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaListAvailabilityZonesInvoker) Invoke() (*model.NovaListAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaListAvailabilityZonesResponse), nil
	}
}

type NovaListKeypairsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaListKeypairsInvoker) Invoke() (*model.NovaListKeypairsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaListKeypairsResponse), nil
	}
}

type NovaListServerSecurityGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaListServerSecurityGroupsInvoker) Invoke() (*model.NovaListServerSecurityGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaListServerSecurityGroupsResponse), nil
	}
}

type NovaListServersDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaListServersDetailsInvoker) Invoke() (*model.NovaListServersDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaListServersDetailsResponse), nil
	}
}

type NovaShowKeypairInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaShowKeypairInvoker) Invoke() (*model.NovaShowKeypairResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaShowKeypairResponse), nil
	}
}

type NovaShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaShowServerInvoker) Invoke() (*model.NovaShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaShowServerResponse), nil
	}
}

type NovaShowServerInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaShowServerInterfaceInvoker) Invoke() (*model.NovaShowServerInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaShowServerInterfaceResponse), nil
	}
}

type RegisterServerMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterServerMonitorInvoker) Invoke() (*model.RegisterServerMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterServerMonitorResponse), nil
	}
}

type ReinstallServerWithCloudInitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallServerWithCloudInitInvoker) Invoke() (*model.ReinstallServerWithCloudInitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallServerWithCloudInitResponse), nil
	}
}

type ReinstallServerWithoutCloudInitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallServerWithoutCloudInitInvoker) Invoke() (*model.ReinstallServerWithoutCloudInitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallServerWithoutCloudInitResponse), nil
	}
}

type ResetServerPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetServerPasswordInvoker) Invoke() (*model.ResetServerPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetServerPasswordResponse), nil
	}
}

type ResizePostPaidServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizePostPaidServerInvoker) Invoke() (*model.ResizePostPaidServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizePostPaidServerResponse), nil
	}
}

type ResizeServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeServerInvoker) Invoke() (*model.ResizeServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeServerResponse), nil
	}
}

type ShowResetPasswordFlagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResetPasswordFlagInvoker) Invoke() (*model.ShowResetPasswordFlagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResetPasswordFlagResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ShowServerBlockDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerBlockDeviceInvoker) Invoke() (*model.ShowServerBlockDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerBlockDeviceResponse), nil
	}
}

type ShowServerGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerGroupInvoker) Invoke() (*model.ShowServerGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerGroupResponse), nil
	}
}

type ShowServerLimitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerLimitsInvoker) Invoke() (*model.ShowServerLimitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerLimitsResponse), nil
	}
}

type ShowServerPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerPasswordInvoker) Invoke() (*model.ShowServerPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerPasswordResponse), nil
	}
}

type ShowServerRemoteConsoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerRemoteConsoleInvoker) Invoke() (*model.ShowServerRemoteConsoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerRemoteConsoleResponse), nil
	}
}

type ShowServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerTagsInvoker) Invoke() (*model.ShowServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerTagsResponse), nil
	}
}

type UpdateServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerInvoker) Invoke() (*model.UpdateServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerResponse), nil
	}
}

type UpdateServerAutoTerminateTimeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerAutoTerminateTimeInvoker) Invoke() (*model.UpdateServerAutoTerminateTimeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerAutoTerminateTimeResponse), nil
	}
}

type UpdateServerBlockDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerBlockDeviceInvoker) Invoke() (*model.UpdateServerBlockDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerBlockDeviceResponse), nil
	}
}

type UpdateServerMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateServerMetadataInvoker) Invoke() (*model.UpdateServerMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateServerMetadataResponse), nil
	}
}

type NovaListVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaListVersionsInvoker) Invoke() (*model.NovaListVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaListVersionsResponse), nil
	}
}

type NovaShowVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *NovaShowVersionInvoker) Invoke() (*model.NovaShowVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.NovaShowVersionResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}
