package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsdeploy/v2/model"
)

type CreateAppGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppGroupsInvoker) Invoke() (*model.CreateAppGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppGroupsResponse), nil
	}
}

type DeleteAppGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppGroupsInvoker) Invoke() (*model.DeleteAppGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppGroupsResponse), nil
	}
}

type ListAppGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAppGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAppGroupsInvoker) Invoke() (*model.ListAppGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAppGroupsResponse), nil
	}
}

type MoveAppGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *MoveAppGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MoveAppGroupsInvoker) Invoke() (*model.MoveAppGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MoveAppGroupsResponse), nil
	}
}

type MoveAppToGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *MoveAppToGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MoveAppToGroupInvoker) Invoke() (*model.MoveAppToGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MoveAppToGroupResponse), nil
	}
}

type UpdateAppGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppGroupsInvoker) Invoke() (*model.UpdateAppGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppGroupsResponse), nil
	}
}

type BatchUpdateApplicationPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdateApplicationPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdateApplicationPermissionsInvoker) Invoke() (*model.BatchUpdateApplicationPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdateApplicationPermissionsResponse), nil
	}
}

type BatchUpdatePermissionLevelInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchUpdatePermissionLevelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchUpdatePermissionLevelInvoker) Invoke() (*model.BatchUpdatePermissionLevelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchUpdatePermissionLevelResponse), nil
	}
}

type CheckCanCreateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckCanCreateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckCanCreateInvoker) Invoke() (*model.CheckCanCreateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckCanCreateResponse), nil
	}
}

type ListApplicationPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApplicationPermissionsInvoker) Invoke() (*model.ListApplicationPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationPermissionsResponse), nil
	}
}

type BatchDeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteAppInvoker) Invoke() (*model.BatchDeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteAppResponse), nil
	}
}

type CheckIsDuplicateAppNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckIsDuplicateAppNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckIsDuplicateAppNameInvoker) Invoke() (*model.CheckIsDuplicateAppNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckIsDuplicateAppNameResponse), nil
	}
}

type CopyApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyApplicationInvoker) Invoke() (*model.CopyApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyApplicationResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateDeployTaskByTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeployTaskByTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeployTaskByTemplateInvoker) Invoke() (*model.CreateDeployTaskByTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeployTaskByTemplateResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type DeleteDeployTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeployTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeployTaskInvoker) Invoke() (*model.DeleteDeployTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeployTaskResponse), nil
	}
}

type ListAllAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAllAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAllAppInvoker) Invoke() (*model.ListAllAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAllAppResponse), nil
	}
}

type ListDeployTaskHistoryByDateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeployTaskHistoryByDateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeployTaskHistoryByDateInvoker) Invoke() (*model.ListDeployTaskHistoryByDateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeployTaskHistoryByDateResponse), nil
	}
}

type ListDeployTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDeployTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListDeployTasksInvoker) Invoke() (*model.ListDeployTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDeployTasksResponse), nil
	}
}

type ShowAppDetailByIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppDetailByIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppDetailByIdInvoker) Invoke() (*model.ShowAppDetailByIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppDetailByIdResponse), nil
	}
}

type ShowDeployTaskDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeployTaskDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeployTaskDetailInvoker) Invoke() (*model.ShowDeployTaskDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeployTaskDetailResponse), nil
	}
}

type ShowExecutionParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecutionParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowExecutionParamsInvoker) Invoke() (*model.ShowExecutionParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecutionParamsResponse), nil
	}
}

type StartDeployTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartDeployTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartDeployTaskInvoker) Invoke() (*model.StartDeployTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartDeployTaskResponse), nil
	}
}

type UpdateAppDisableStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppDisableStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppDisableStatusInvoker) Invoke() (*model.UpdateAppDisableStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppDisableStatusResponse), nil
	}
}

type UpdateAppInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateAppInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateAppInfoInvoker) Invoke() (*model.UpdateAppInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateAppInfoResponse), nil
	}
}

type CreateEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnvironmentInvoker) Invoke() (*model.CreateEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentResponse), nil
	}
}

type DeleteEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEnvironmentInvoker) Invoke() (*model.DeleteEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentResponse), nil
	}
}

type DeleteHostFromEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostFromEnvironmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostFromEnvironmentInvoker) Invoke() (*model.DeleteHostFromEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostFromEnvironmentResponse), nil
	}
}

type ImportHostToEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportHostToEnvironmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ImportHostToEnvironmentInvoker) Invoke() (*model.ImportHostToEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportHostToEnvironmentResponse), nil
	}
}

type ListEnvironmentHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvironmentHostsInvoker) Invoke() (*model.ListEnvironmentHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentHostsResponse), nil
	}
}

type ListEnvironmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvironmentsInvoker) Invoke() (*model.ListEnvironmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentsResponse), nil
	}
}

type ShowEnvironmentDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnvironmentDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEnvironmentDetailInvoker) Invoke() (*model.ShowEnvironmentDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnvironmentDetailResponse), nil
	}
}

type UpdateEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnvironmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnvironmentInvoker) Invoke() (*model.UpdateEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnvironmentResponse), nil
	}
}

type ListEnvironmentPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnvironmentPermissionsInvoker) Invoke() (*model.ListEnvironmentPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentPermissionsResponse), nil
	}
}

type UpdateEnvironmentPermissionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnvironmentPermissionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnvironmentPermissionInvoker) Invoke() (*model.UpdateEnvironmentPermissionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnvironmentPermissionResponse), nil
	}
}

type BatchDeleteHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteHostsInvoker) Invoke() (*model.BatchDeleteHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteHostsResponse), nil
	}
}

type CopyHostsToTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CopyHostsToTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CopyHostsToTargetInvoker) Invoke() (*model.CopyHostsToTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CopyHostsToTargetResponse), nil
	}
}

type CreateDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeploymentHostInvoker) Invoke() (*model.CreateDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentHostResponse), nil
	}
}

type CreateHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHostInvoker) Invoke() (*model.CreateHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostResponse), nil
	}
}

type DeleteDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeploymentHostInvoker) Invoke() (*model.DeleteDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentHostResponse), nil
	}
}

type DeleteHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostInvoker) Invoke() (*model.DeleteHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostResponse), nil
	}
}

type ListHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostsInvoker) Invoke() (*model.ListHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostsResponse), nil
	}
}

type ListNewHostsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNewHostsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListNewHostsInvoker) Invoke() (*model.ListNewHostsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNewHostsResponse), nil
	}
}

type ShowDeploymentHostDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentHostDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeploymentHostDetailInvoker) Invoke() (*model.ShowDeploymentHostDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentHostDetailResponse), nil
	}
}

type ShowHostDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostDetailInvoker) Invoke() (*model.ShowHostDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostDetailResponse), nil
	}
}

type UpdateDeploymentHostInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentHostInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeploymentHostInvoker) Invoke() (*model.UpdateDeploymentHostResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentHostResponse), nil
	}
}

type UpdateHostInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostInfoInvoker) Invoke() (*model.UpdateHostInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostInfoResponse), nil
	}
}

type CreateDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDeploymentGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDeploymentGroupInvoker) Invoke() (*model.CreateDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDeploymentGroupResponse), nil
	}
}

type CreateHostClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHostClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHostClusterInvoker) Invoke() (*model.CreateHostClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHostClusterResponse), nil
	}
}

type DeleteDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeploymentGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDeploymentGroupInvoker) Invoke() (*model.DeleteDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeploymentGroupResponse), nil
	}
}

type DeleteHostClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteHostClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteHostClusterInvoker) Invoke() (*model.DeleteHostClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteHostClusterResponse), nil
	}
}

type ListAssociateEnvironmentsInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAssociateEnvironmentsInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAssociateEnvironmentsInfosInvoker) Invoke() (*model.ListAssociateEnvironmentsInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAssociateEnvironmentsInfosResponse), nil
	}
}

type ListHostClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostClustersInvoker) Invoke() (*model.ListHostClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostClustersResponse), nil
	}
}

type ListHostGroupBaseInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupBaseInfosInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostGroupBaseInfosInvoker) Invoke() (*model.ListHostGroupBaseInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupBaseInfosResponse), nil
	}
}

type ListHostGroupsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostGroupsInvoker) Invoke() (*model.ListHostGroupsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupsResponse), nil
	}
}

type ShowDeploymentGroupDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentGroupDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeploymentGroupDetailInvoker) Invoke() (*model.ShowDeploymentGroupDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentGroupDetailResponse), nil
	}
}

type ShowHostClusterDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHostClusterDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHostClusterDetailInvoker) Invoke() (*model.ShowHostClusterDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHostClusterDetailResponse), nil
	}
}

type UpdateDeploymentGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeploymentGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDeploymentGroupInvoker) Invoke() (*model.UpdateDeploymentGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeploymentGroupResponse), nil
	}
}

type UpdateHostClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostClusterInvoker) Invoke() (*model.UpdateHostClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostClusterResponse), nil
	}
}

type CheckWhetherHostGroupCanBeCreatedInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckWhetherHostGroupCanBeCreatedInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckWhetherHostGroupCanBeCreatedInvoker) Invoke() (*model.CheckWhetherHostGroupCanBeCreatedResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckWhetherHostGroupCanBeCreatedResponse), nil
	}
}

type ListHostGroupPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHostGroupPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHostGroupPermissionsInvoker) Invoke() (*model.ListHostGroupPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHostGroupPermissionsResponse), nil
	}
}

type UpdateHostGroupPermissionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHostGroupPermissionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHostGroupPermissionsInvoker) Invoke() (*model.UpdateHostGroupPermissionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHostGroupPermissionsResponse), nil
	}
}

type ListTaskSuccessRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTaskSuccessRateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTaskSuccessRateInvoker) Invoke() (*model.ListTaskSuccessRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTaskSuccessRateResponse), nil
	}
}

type ShowProjectSuccessRateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectSuccessRateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowProjectSuccessRateInvoker) Invoke() (*model.ShowProjectSuccessRateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectSuccessRateResponse), nil
	}
}
