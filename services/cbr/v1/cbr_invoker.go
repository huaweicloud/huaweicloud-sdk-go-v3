package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbr/v1/model"
)

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
