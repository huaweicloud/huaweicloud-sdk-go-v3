package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpn/v5/model"
)

type CheckClientCaCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckClientCaCertificateInvoker) Invoke() (*model.CheckClientCaCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckClientCaCertificateResponse), nil
	}
}

type DeleteClientCaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClientCaInvoker) Invoke() (*model.DeleteClientCaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClientCaResponse), nil
	}
}

type ImportClientCaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportClientCaInvoker) Invoke() (*model.ImportClientCaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportClientCaResponse), nil
	}
}

type ShowClientCaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClientCaInvoker) Invoke() (*model.ShowClientCaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClientCaResponse), nil
	}
}

type UpdateClientCaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateClientCaInvoker) Invoke() (*model.UpdateClientCaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateClientCaResponse), nil
	}
}

type CreateConnectionMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateConnectionMonitorInvoker) Invoke() (*model.CreateConnectionMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateConnectionMonitorResponse), nil
	}
}

type DeleteConnectionMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteConnectionMonitorInvoker) Invoke() (*model.DeleteConnectionMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteConnectionMonitorResponse), nil
	}
}

type ListConnectionMonitorsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConnectionMonitorsInvoker) Invoke() (*model.ListConnectionMonitorsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConnectionMonitorsResponse), nil
	}
}

type ShowConnectionMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowConnectionMonitorInvoker) Invoke() (*model.ShowConnectionMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowConnectionMonitorResponse), nil
	}
}

type CreateCgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCgwInvoker) Invoke() (*model.CreateCgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCgwResponse), nil
	}
}

type DeleteCgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCgwInvoker) Invoke() (*model.DeleteCgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCgwResponse), nil
	}
}

type ListCgwsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCgwsInvoker) Invoke() (*model.ListCgwsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCgwsResponse), nil
	}
}

type ShowCgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCgwInvoker) Invoke() (*model.ShowCgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCgwResponse), nil
	}
}

type UpdateCgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCgwInvoker) Invoke() (*model.UpdateCgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCgwResponse), nil
	}
}

type ListP2cVgwAvailabilityZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListP2cVgwAvailabilityZonesInvoker) Invoke() (*model.ListP2cVgwAvailabilityZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListP2cVgwAvailabilityZonesResponse), nil
	}
}

type ListP2cVgwConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListP2cVgwConnectionsInvoker) Invoke() (*model.ListP2cVgwConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListP2cVgwConnectionsResponse), nil
	}
}

type ListP2cVgwsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListP2cVgwsInvoker) Invoke() (*model.ListP2cVgwsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListP2cVgwsResponse), nil
	}
}

type ShowP2cVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowP2cVgwInvoker) Invoke() (*model.ShowP2cVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowP2cVgwResponse), nil
	}
}

type UpdateP2cVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateP2cVgwInvoker) Invoke() (*model.UpdateP2cVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateP2cVgwResponse), nil
	}
}

type BatchCreateResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateResourceTagsInvoker) Invoke() (*model.BatchCreateResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateResourceTagsResponse), nil
	}
}

type BatchDeleteResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteResourceTagsInvoker) Invoke() (*model.BatchDeleteResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteResourceTagsResponse), nil
	}
}

type CountResourcesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountResourcesByTagsInvoker) Invoke() (*model.CountResourcesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountResourcesByTagsResponse), nil
	}
}

type ListProjectTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectTagsInvoker) Invoke() (*model.ListProjectTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectTagsResponse), nil
	}
}

type ListResourcesByTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourcesByTagsInvoker) Invoke() (*model.ListResourcesByTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourcesByTagsResponse), nil
	}
}

type ShowResourceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceTagsInvoker) Invoke() (*model.ShowResourceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceTagsResponse), nil
	}
}

type CreateVpnAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpnAccessPolicyInvoker) Invoke() (*model.CreateVpnAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpnAccessPolicyResponse), nil
	}
}

type DeleteVpnAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpnAccessPolicyInvoker) Invoke() (*model.DeleteVpnAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpnAccessPolicyResponse), nil
	}
}

type ListVpnAccessPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnAccessPoliciesInvoker) Invoke() (*model.ListVpnAccessPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnAccessPoliciesResponse), nil
	}
}

type ShowVpnAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpnAccessPolicyInvoker) Invoke() (*model.ShowVpnAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpnAccessPolicyResponse), nil
	}
}

type UpdateVpnAccessPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnAccessPolicyInvoker) Invoke() (*model.UpdateVpnAccessPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnAccessPolicyResponse), nil
	}
}

type CreateVpnConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpnConnectionInvoker) Invoke() (*model.CreateVpnConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpnConnectionResponse), nil
	}
}

type DeleteVpnConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpnConnectionInvoker) Invoke() (*model.DeleteVpnConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpnConnectionResponse), nil
	}
}

type ListVpnConnectionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnConnectionsInvoker) Invoke() (*model.ListVpnConnectionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnConnectionsResponse), nil
	}
}

type ShowVpnConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpnConnectionInvoker) Invoke() (*model.ShowVpnConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpnConnectionResponse), nil
	}
}

type UpdateVpnConnectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnConnectionInvoker) Invoke() (*model.UpdateVpnConnectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnConnectionResponse), nil
	}
}

type CreateVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVgwInvoker) Invoke() (*model.CreateVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVgwResponse), nil
	}
}

type DeleteVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVgwInvoker) Invoke() (*model.DeleteVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVgwResponse), nil
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

type ListVgwsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVgwsInvoker) Invoke() (*model.ListVgwsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVgwsResponse), nil
	}
}

type ShowVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVgwInvoker) Invoke() (*model.ShowVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVgwResponse), nil
	}
}

type UpdatePostpaidVgwSpecificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePostpaidVgwSpecificationInvoker) Invoke() (*model.UpdatePostpaidVgwSpecificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePostpaidVgwSpecificationResponse), nil
	}
}

type UpdateVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVgwInvoker) Invoke() (*model.UpdateVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVgwResponse), nil
	}
}

type CreateVgwCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVgwCertificateInvoker) Invoke() (*model.CreateVgwCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVgwCertificateResponse), nil
	}
}

type ShowVpnGatewayCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpnGatewayCertificateInvoker) Invoke() (*model.ShowVpnGatewayCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpnGatewayCertificateResponse), nil
	}
}

type UpdateVgwCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVgwCertificateInvoker) Invoke() (*model.UpdateVgwCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVgwCertificateResponse), nil
	}
}

type ShowQuotasInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInfoInvoker) Invoke() (*model.ShowQuotasInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasInfoResponse), nil
	}
}

type CreateVpnServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpnServerInvoker) Invoke() (*model.CreateVpnServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpnServerResponse), nil
	}
}

type ExportClientConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportClientConfigInvoker) Invoke() (*model.ExportClientConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportClientConfigResponse), nil
	}
}

type ListVpnServersByProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnServersByProjectInvoker) Invoke() (*model.ListVpnServersByProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnServersByProjectResponse), nil
	}
}

type ListVpnServersByVgwInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnServersByVgwInvoker) Invoke() (*model.ListVpnServersByVgwResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnServersByVgwResponse), nil
	}
}

type UpdateVpnServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnServerInvoker) Invoke() (*model.UpdateVpnServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnServerResponse), nil
	}
}

type CreateVpnUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpnUserInvoker) Invoke() (*model.CreateVpnUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpnUserResponse), nil
	}
}

type DeleteVpnUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpnUserInvoker) Invoke() (*model.DeleteVpnUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpnUserResponse), nil
	}
}

type ListVpnUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnUsersInvoker) Invoke() (*model.ListVpnUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnUsersResponse), nil
	}
}

type ResetVpnUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetVpnUserPasswordInvoker) Invoke() (*model.ResetVpnUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetVpnUserPasswordResponse), nil
	}
}

type ShowVpnUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpnUserInvoker) Invoke() (*model.ShowVpnUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpnUserResponse), nil
	}
}

type UpdateVpnUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnUserInvoker) Invoke() (*model.UpdateVpnUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnUserResponse), nil
	}
}

type UpdateVpnUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnUserPasswordInvoker) Invoke() (*model.UpdateVpnUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnUserPasswordResponse), nil
	}
}

type AddVpnUsersToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVpnUsersToGroupInvoker) Invoke() (*model.AddVpnUsersToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVpnUsersToGroupResponse), nil
	}
}

type CreateVpnUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVpnUserGroupInvoker) Invoke() (*model.CreateVpnUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVpnUserGroupResponse), nil
	}
}

type DeleteVpnUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVpnUserGroupInvoker) Invoke() (*model.DeleteVpnUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVpnUserGroupResponse), nil
	}
}

type ListVpnUserGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnUserGroupsInvoker) Invoke() (*model.ListVpnUserGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnUserGroupsResponse), nil
	}
}

type ListVpnUsersInGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVpnUsersInGroupInvoker) Invoke() (*model.ListVpnUsersInGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVpnUsersInGroupResponse), nil
	}
}

type RemoveVpnUsersFromGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveVpnUsersFromGroupInvoker) Invoke() (*model.RemoveVpnUsersFromGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveVpnUsersFromGroupResponse), nil
	}
}

type ShowVpnUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVpnUserGroupInvoker) Invoke() (*model.ShowVpnUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVpnUserGroupResponse), nil
	}
}

type UpdateVpnUserGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVpnUserGroupInvoker) Invoke() (*model.UpdateVpnUserGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVpnUserGroupResponse), nil
	}
}
