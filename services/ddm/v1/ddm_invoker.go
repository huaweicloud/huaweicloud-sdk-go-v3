package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ddm/v1/model"
)

type CreateDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDatabaseInvoker) Invoke() (*model.CreateDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDatabaseResponse), nil
	}
}

type CreateGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateGroupInvoker) Invoke() (*model.CreateGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateGroupResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type CreateUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateUsersInvoker) Invoke() (*model.CreateUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateUsersResponse), nil
	}
}

type DeleteDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDatabaseInvoker) Invoke() (*model.DeleteDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDatabaseResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type DeleteUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteUserInvoker) Invoke() (*model.DeleteUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteUserResponse), nil
	}
}

type ExecuteKillLogicalProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteKillLogicalProcessesInvoker) Invoke() (*model.ExecuteKillLogicalProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteKillLogicalProcessesResponse), nil
	}
}

type ExecuteKillPhysicalProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteKillPhysicalProcessesInvoker) Invoke() (*model.ExecuteKillPhysicalProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteKillPhysicalProcessesResponse), nil
	}
}

type ExpandInstanceNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandInstanceNodesInvoker) Invoke() (*model.ExpandInstanceNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandInstanceNodesResponse), nil
	}
}

type ListAvailableRdsListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableRdsListInvoker) Invoke() (*model.ListAvailableRdsListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableRdsListResponse), nil
	}
}

type ListDatabasesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDatabasesInvoker) Invoke() (*model.ListDatabasesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDatabasesResponse), nil
	}
}

type ListEnginesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnginesInvoker) Invoke() (*model.ListEnginesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnginesResponse), nil
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

type ListGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListGroupInvoker) Invoke() (*model.ListGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListGroupResponse), nil
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

type ListNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNodesInvoker) Invoke() (*model.ListNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNodesResponse), nil
	}
}

type ListReadWriteRatioInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListReadWriteRatioInvoker) Invoke() (*model.ListReadWriteRatioResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListReadWriteRatioResponse), nil
	}
}

type ListSlowLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSlowLogInvoker) Invoke() (*model.ListSlowLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSlowLogResponse), nil
	}
}

type ListUsersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListUsersInvoker) Invoke() (*model.ListUsersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListUsersResponse), nil
	}
}

type RebuildConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebuildConfigInvoker) Invoke() (*model.RebuildConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebuildConfigResponse), nil
	}
}

type ResetAdministratorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetAdministratorInvoker) Invoke() (*model.ResetAdministratorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetAdministratorResponse), nil
	}
}

type ResetUserPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetUserPasswordInvoker) Invoke() (*model.ResetUserPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetUserPasswordResponse), nil
	}
}

type ResizeFlavorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeFlavorInvoker) Invoke() (*model.ResizeFlavorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeFlavorResponse), nil
	}
}

type RestartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartInstanceInvoker) Invoke() (*model.RestartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartInstanceResponse), nil
	}
}

type ShowDatabaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDatabaseInvoker) Invoke() (*model.ShowDatabaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDatabaseResponse), nil
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

type ShowInstanceParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceParamInvoker) Invoke() (*model.ShowInstanceParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceParamResponse), nil
	}
}

type ShowLogicalProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogicalProcessesInvoker) Invoke() (*model.ShowLogicalProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogicalProcessesResponse), nil
	}
}

type ShowNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNodeInvoker) Invoke() (*model.ShowNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNodeResponse), nil
	}
}

type ShowPhysicalProcessesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPhysicalProcessesInvoker) Invoke() (*model.ShowPhysicalProcessesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPhysicalProcessesResponse), nil
	}
}

type ShowProcessesAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProcessesAuditLogInvoker) Invoke() (*model.ShowProcessesAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProcessesAuditLogResponse), nil
	}
}

type ShrinkInstanceNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShrinkInstanceNodesInvoker) Invoke() (*model.ShrinkInstanceNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShrinkInstanceNodesResponse), nil
	}
}

type UpdateDatabaseInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDatabaseInfoInvoker) Invoke() (*model.UpdateDatabaseInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDatabaseInfoResponse), nil
	}
}

type UpdateInstanceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceNameInvoker) Invoke() (*model.UpdateInstanceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceNameResponse), nil
	}
}

type UpdateInstanceParamInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceParamInvoker) Invoke() (*model.UpdateInstanceParamResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceParamResponse), nil
	}
}

type UpdateInstanceSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceSecurityGroupInvoker) Invoke() (*model.UpdateInstanceSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceSecurityGroupResponse), nil
	}
}

type UpdateReadAndWriteStrategyInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateReadAndWriteStrategyInvoker) Invoke() (*model.UpdateReadAndWriteStrategyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateReadAndWriteStrategyResponse), nil
	}
}

type UpdateUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateUserInvoker) Invoke() (*model.UpdateUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateUserResponse), nil
	}
}

type ValidateWeakPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateWeakPasswordInvoker) Invoke() (*model.ValidateWeakPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateWeakPasswordResponse), nil
	}
}
