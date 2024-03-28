package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

type AddAgentPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAgentPathInvoker) Invoke() (*model.AddAgentPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAgentPathResponse), nil
	}
}

type AddMemberInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddMemberInvoker) Invoke() (*model.AddMemberResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddMemberResponse), nil
	}
}

type AddVaultResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVaultResourceInvoker) Invoke() (*model.AddVaultResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVaultResourceResponse), nil
	}
}

type AssociateVaultPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateVaultPolicyInvoker) Invoke() (*model.AssociateVaultPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateVaultPolicyResponse), nil
	}
}

type BatchCreateAndDeleteVaultTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateAndDeleteVaultTagsInvoker) Invoke() (*model.BatchCreateAndDeleteVaultTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateAndDeleteVaultTagsResponse), nil
	}
}

type BatchUpdateVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateVaultInvoker) Invoke() (*model.BatchUpdateVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateVaultResponse), nil
	}
}

type CheckAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckAgentInvoker) Invoke() (*model.CheckAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckAgentResponse), nil
	}
}

type CopyBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyBackupInvoker) Invoke() (*model.CopyBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyBackupResponse), nil
	}
}

type CopyCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyCheckpointInvoker) Invoke() (*model.CopyCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyCheckpointResponse), nil
	}
}

type CreateCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCheckpointInvoker) Invoke() (*model.CreateCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCheckpointResponse), nil
	}
}

type CreateOrganizationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateOrganizationPolicyInvoker) Invoke() (*model.CreateOrganizationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateOrganizationPolicyResponse), nil
	}
}

type CreatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePolicyInvoker) Invoke() (*model.CreatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePolicyResponse), nil
	}
}

type CreatePostPaidVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePostPaidVaultInvoker) Invoke() (*model.CreatePostPaidVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePostPaidVaultResponse), nil
	}
}

type CreateVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVaultInvoker) Invoke() (*model.CreateVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVaultResponse), nil
	}
}

type CreateVaultTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVaultTagsInvoker) Invoke() (*model.CreateVaultTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVaultTagsResponse), nil
	}
}

type DeleteBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackupInvoker) Invoke() (*model.DeleteBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackupResponse), nil
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

type DeleteOrganizationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOrganizationPolicyInvoker) Invoke() (*model.DeleteOrganizationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOrganizationPolicyResponse), nil
	}
}

type DeletePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePolicyInvoker) Invoke() (*model.DeletePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePolicyResponse), nil
	}
}

type DeleteVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVaultInvoker) Invoke() (*model.DeleteVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVaultResponse), nil
	}
}

type DeleteVaultTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVaultTagInvoker) Invoke() (*model.DeleteVaultTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVaultTagResponse), nil
	}
}

type DisassociateVaultPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateVaultPolicyInvoker) Invoke() (*model.DisassociateVaultPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateVaultPolicyResponse), nil
	}
}

type ImportBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportBackupInvoker) Invoke() (*model.ImportBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportBackupResponse), nil
	}
}

type ImportCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportCheckpointInvoker) Invoke() (*model.ImportCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportCheckpointResponse), nil
	}
}

type ListAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgentInvoker) Invoke() (*model.ListAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgentResponse), nil
	}
}

type ListBackupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackupsInvoker) Invoke() (*model.ListBackupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackupsResponse), nil
	}
}

type ListDomainProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainProjectsInvoker) Invoke() (*model.ListDomainProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainProjectsResponse), nil
	}
}

type ListExternalVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExternalVaultInvoker) Invoke() (*model.ListExternalVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExternalVaultResponse), nil
	}
}

type ListOpLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOpLogsInvoker) Invoke() (*model.ListOpLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOpLogsResponse), nil
	}
}

type ListOrganizationPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationPoliciesInvoker) Invoke() (*model.ListOrganizationPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationPoliciesResponse), nil
	}
}

type ListOrganizationPolicyDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOrganizationPolicyDetailInvoker) Invoke() (*model.ListOrganizationPolicyDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOrganizationPolicyDetailResponse), nil
	}
}

type ListPoliciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPoliciesInvoker) Invoke() (*model.ListPoliciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPoliciesResponse), nil
	}
}

type ListProjectsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProjectsInvoker) Invoke() (*model.ListProjectsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProjectsResponse), nil
	}
}

type ListProtectableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProtectableInvoker) Invoke() (*model.ListProtectableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProtectableResponse), nil
	}
}

type ListVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVaultInvoker) Invoke() (*model.ListVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVaultResponse), nil
	}
}

type MigrateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateDomainInvoker) Invoke() (*model.MigrateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateDomainResponse), nil
	}
}

type MigrateVaultResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateVaultResourceInvoker) Invoke() (*model.MigrateVaultResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateVaultResourceResponse), nil
	}
}

type RegisterAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterAgentInvoker) Invoke() (*model.RegisterAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterAgentResponse), nil
	}
}

type RemoveAgentPathInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveAgentPathInvoker) Invoke() (*model.RemoveAgentPathResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveAgentPathResponse), nil
	}
}

type RemoveVaultResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RemoveVaultResourceInvoker) Invoke() (*model.RemoveVaultResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RemoveVaultResourceResponse), nil
	}
}

type RestoreBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestoreBackupInvoker) Invoke() (*model.RestoreBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestoreBackupResponse), nil
	}
}

type SetVaultResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetVaultResourceInvoker) Invoke() (*model.SetVaultResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetVaultResourceResponse), nil
	}
}

type ShowAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAgentInvoker) Invoke() (*model.ShowAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAgentResponse), nil
	}
}

type ShowBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackupInvoker) Invoke() (*model.ShowBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackupResponse), nil
	}
}

type ShowCheckpointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCheckpointInvoker) Invoke() (*model.ShowCheckpointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCheckpointResponse), nil
	}
}

type ShowDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDomainInvoker) Invoke() (*model.ShowDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDomainResponse), nil
	}
}

type ShowMemberDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMemberDetailInvoker) Invoke() (*model.ShowMemberDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMemberDetailResponse), nil
	}
}

type ShowMembersDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMembersDetailInvoker) Invoke() (*model.ShowMembersDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMembersDetailResponse), nil
	}
}

type ShowMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMetadataInvoker) Invoke() (*model.ShowMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMetadataResponse), nil
	}
}

type ShowMigrateStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowMigrateStatusInvoker) Invoke() (*model.ShowMigrateStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowMigrateStatusResponse), nil
	}
}

type ShowOpLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOpLogInvoker) Invoke() (*model.ShowOpLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOpLogResponse), nil
	}
}

type ShowOrganizationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOrganizationPolicyInvoker) Invoke() (*model.ShowOrganizationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOrganizationPolicyResponse), nil
	}
}

type ShowPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPolicyInvoker) Invoke() (*model.ShowPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPolicyResponse), nil
	}
}

type ShowProtectableInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProtectableInvoker) Invoke() (*model.ShowProtectableResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProtectableResponse), nil
	}
}

type ShowReplicationCapabilitiesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReplicationCapabilitiesInvoker) Invoke() (*model.ShowReplicationCapabilitiesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReplicationCapabilitiesResponse), nil
	}
}

type ShowStorageUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStorageUsageInvoker) Invoke() (*model.ShowStorageUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStorageUsageResponse), nil
	}
}

type ShowSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSummaryInvoker) Invoke() (*model.ShowSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSummaryResponse), nil
	}
}

type ShowVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVaultInvoker) Invoke() (*model.ShowVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVaultResponse), nil
	}
}

type ShowVaultProjectTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVaultProjectTagInvoker) Invoke() (*model.ShowVaultProjectTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVaultProjectTagResponse), nil
	}
}

type ShowVaultResourceInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVaultResourceInstancesInvoker) Invoke() (*model.ShowVaultResourceInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVaultResourceInstancesResponse), nil
	}
}

type ShowVaultTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVaultTagInvoker) Invoke() (*model.ShowVaultTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVaultTagResponse), nil
	}
}

type UnregisterAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnregisterAgentInvoker) Invoke() (*model.UnregisterAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnregisterAgentResponse), nil
	}
}

type UpdateAgentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAgentInvoker) Invoke() (*model.UpdateAgentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAgentResponse), nil
	}
}

type UpdateBackupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBackupInvoker) Invoke() (*model.UpdateBackupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBackupResponse), nil
	}
}

type UpdateMemberStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateMemberStatusInvoker) Invoke() (*model.UpdateMemberStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateMemberStatusResponse), nil
	}
}

type UpdateOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOrderInvoker) Invoke() (*model.UpdateOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOrderResponse), nil
	}
}

type UpdateOrganizationPolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateOrganizationPolicyInvoker) Invoke() (*model.UpdateOrganizationPolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateOrganizationPolicyResponse), nil
	}
}

type UpdatePolicyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePolicyInvoker) Invoke() (*model.UpdatePolicyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePolicyResponse), nil
	}
}

type UpdateVaultInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVaultInvoker) Invoke() (*model.UpdateVaultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVaultResponse), nil
	}
}
