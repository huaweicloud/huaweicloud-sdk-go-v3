package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
)

type BatchAddSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAddSharedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchAddSharedTagsInvoker) Invoke() (*model.BatchAddSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAddSharedTagsResponse), nil
	}
}

type ChangeSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeSecurityGroupInvoker) Invoke() (*model.ChangeSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeSecurityGroupResponse), nil
	}
}

type ChangeShareNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeShareNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeShareNameInvoker) Invoke() (*model.ChangeShareNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeShareNameResponse), nil
	}
}

type CreateBackendTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBackendTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBackendTargetInvoker) Invoke() (*model.CreateBackendTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBackendTargetResponse), nil
	}
}

type CreateFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFsDirInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFsDirInvoker) Invoke() (*model.CreateFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFsDirResponse), nil
	}
}

type CreateFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFsDirQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFsDirQuotaInvoker) Invoke() (*model.CreateFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFsDirQuotaResponse), nil
	}
}

type CreateFsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateFsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateFsTaskInvoker) Invoke() (*model.CreateFsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateFsTaskResponse), nil
	}
}

type CreateHpcCacheTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateHpcCacheTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateHpcCacheTaskInvoker) Invoke() (*model.CreateHpcCacheTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateHpcCacheTaskResponse), nil
	}
}

type CreateLdapConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateLdapConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateLdapConfigInvoker) Invoke() (*model.CreateLdapConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateLdapConfigResponse), nil
	}
}

type CreatePermRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreatePermRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreatePermRuleInvoker) Invoke() (*model.CreatePermRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreatePermRuleResponse), nil
	}
}

type CreateShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateShareInvoker) Invoke() (*model.CreateShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateShareResponse), nil
	}
}

type CreateSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateSharedTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateSharedTagInvoker) Invoke() (*model.CreateSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateSharedTagResponse), nil
	}
}

type DeleteBackendTargetInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteBackendTargetInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteBackendTargetInvoker) Invoke() (*model.DeleteBackendTargetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteBackendTargetResponse), nil
	}
}

type DeleteFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFsDirInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFsDirInvoker) Invoke() (*model.DeleteFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFsDirResponse), nil
	}
}

type DeleteFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFsDirQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFsDirQuotaInvoker) Invoke() (*model.DeleteFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFsDirQuotaResponse), nil
	}
}

type DeleteFsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteFsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteFsTaskInvoker) Invoke() (*model.DeleteFsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteFsTaskResponse), nil
	}
}

type DeleteLdapConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLdapConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLdapConfigInvoker) Invoke() (*model.DeleteLdapConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLdapConfigResponse), nil
	}
}

type DeletePermRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeletePermRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeletePermRuleInvoker) Invoke() (*model.DeletePermRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeletePermRuleResponse), nil
	}
}

type DeleteShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteShareInvoker) Invoke() (*model.DeleteShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteShareResponse), nil
	}
}

type DeleteSharedTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteSharedTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteSharedTagInvoker) Invoke() (*model.DeleteSharedTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteSharedTagResponse), nil
	}
}

type ExpandShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExpandShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExpandShareInvoker) Invoke() (*model.ExpandShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExpandShareResponse), nil
	}
}

type ListBackendTargetsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBackendTargetsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListBackendTargetsInvoker) Invoke() (*model.ListBackendTargetsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBackendTargetsResponse), nil
	}
}

type ListFsTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListFsTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListFsTasksInvoker) Invoke() (*model.ListFsTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListFsTasksResponse), nil
	}
}

type ListHpcCacheTasksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListHpcCacheTasksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListHpcCacheTasksInvoker) Invoke() (*model.ListHpcCacheTasksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListHpcCacheTasksResponse), nil
	}
}

type ListPermRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPermRulesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPermRulesInvoker) Invoke() (*model.ListPermRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPermRulesResponse), nil
	}
}

type ListSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharedTagsInvoker) Invoke() (*model.ListSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharedTagsResponse), nil
	}
}

type ListSharesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSharesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSharesInvoker) Invoke() (*model.ListSharesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSharesResponse), nil
	}
}

type SetHpcCacheBackendInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetHpcCacheBackendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SetHpcCacheBackendInvoker) Invoke() (*model.SetHpcCacheBackendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetHpcCacheBackendResponse), nil
	}
}

type ShowBackendTargetInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBackendTargetInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBackendTargetInfoInvoker) Invoke() (*model.ShowBackendTargetInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBackendTargetInfoResponse), nil
	}
}

type ShowFsDirInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsDirInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFsDirInvoker) Invoke() (*model.ShowFsDirResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsDirResponse), nil
	}
}

type ShowFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsDirQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFsDirQuotaInvoker) Invoke() (*model.ShowFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsDirQuotaResponse), nil
	}
}

type ShowFsDirUsageInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsDirUsageInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFsDirUsageInvoker) Invoke() (*model.ShowFsDirUsageResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsDirUsageResponse), nil
	}
}

type ShowFsTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowFsTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowFsTaskInvoker) Invoke() (*model.ShowFsTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowFsTaskResponse), nil
	}
}

type ShowHpcCacheTaskInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHpcCacheTaskInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHpcCacheTaskInvoker) Invoke() (*model.ShowHpcCacheTaskResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHpcCacheTaskResponse), nil
	}
}

type ShowJobDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowJobDetailInvoker) Invoke() (*model.ShowJobDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobDetailResponse), nil
	}
}

type ShowLdapConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLdapConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLdapConfigInvoker) Invoke() (*model.ShowLdapConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLdapConfigResponse), nil
	}
}

type ShowPermRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPermRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowPermRuleInvoker) Invoke() (*model.ShowPermRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPermRuleResponse), nil
	}
}

type ShowShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowShareInvoker) Invoke() (*model.ShowShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowShareResponse), nil
	}
}

type ShowSharedTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSharedTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowSharedTagsInvoker) Invoke() (*model.ShowSharedTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSharedTagsResponse), nil
	}
}

type UpdateFsDirQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateFsDirQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateFsDirQuotaInvoker) Invoke() (*model.UpdateFsDirQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateFsDirQuotaResponse), nil
	}
}

type UpdateHpcShareInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateHpcShareInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateHpcShareInvoker) Invoke() (*model.UpdateHpcShareResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateHpcShareResponse), nil
	}
}

type UpdateLdapConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateLdapConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateLdapConfigInvoker) Invoke() (*model.UpdateLdapConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateLdapConfigResponse), nil
	}
}

type UpdatePermRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdatePermRuleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdatePermRuleInvoker) Invoke() (*model.UpdatePermRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdatePermRuleResponse), nil
	}
}
