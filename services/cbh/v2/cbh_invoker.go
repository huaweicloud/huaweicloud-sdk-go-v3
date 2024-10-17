package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v2/model"
)

type BatchCreateInstanceTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateInstanceTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateInstanceTagInvoker) Invoke() (*model.BatchCreateInstanceTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateInstanceTagResponse), nil
	}
}

type ChangeInstanceTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstanceTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeInstanceTypeInvoker) Invoke() (*model.ChangeInstanceTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceTypeResponse), nil
	}
}

type CountInstancesByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *CountInstancesByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CountInstancesByTagInvoker) Invoke() (*model.CountInstancesByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CountInstancesByTagResponse), nil
	}
}

type CreateInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceInvoker) Invoke() (*model.CreateInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceResponse), nil
	}
}

type DeleteInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstanceInvoker) Invoke() (*model.DeleteInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstanceResponse), nil
	}
}

type InstallInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallInstanceEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InstallInstanceEipInvoker) Invoke() (*model.InstallInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallInstanceEipResponse), nil
	}
}

type ListAvailableZonesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAvailableZonesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAvailableZonesInvoker) Invoke() (*model.ListAvailableZonesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAvailableZonesResponse), nil
	}
}

type ListInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesInvoker) Invoke() (*model.ListInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesResponse), nil
	}
}

type ListInstancesByTagInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListInstancesByTagInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListInstancesByTagInvoker) Invoke() (*model.ListInstancesByTagResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListInstancesByTagResponse), nil
	}
}

type ListSpecificationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListSpecificationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListSpecificationsInvoker) Invoke() (*model.ListSpecificationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListSpecificationsResponse), nil
	}
}

type ListTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListTagsInvoker) Invoke() (*model.ListTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTagsResponse), nil
	}
}

type LoginInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *LoginInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LoginInstanceInvoker) Invoke() (*model.LoginInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginInstanceResponse), nil
	}
}

type LoginInstanceAdminInvoker struct {
	*invoker.BaseInvoker
}

func (i *LoginInstanceAdminInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LoginInstanceAdminInvoker) Invoke() (*model.LoginInstanceAdminResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginInstanceAdminResponse), nil
	}
}

type RebootInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootInstanceInvoker) Invoke() (*model.RebootInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootInstanceResponse), nil
	}
}

type RegisterAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RegisterAuthorizationInvoker) Invoke() (*model.RegisterAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterAuthorizationResponse), nil
	}
}

type ResetInstanceLoginMethodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetInstanceLoginMethodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetInstanceLoginMethodInvoker) Invoke() (*model.ResetInstanceLoginMethodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetInstanceLoginMethodResponse), nil
	}
}

type ResetInstancePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetInstancePasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetInstancePasswordInvoker) Invoke() (*model.ResetInstancePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetInstancePasswordResponse), nil
	}
}

type ResizeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResizeInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResizeInstanceInvoker) Invoke() (*model.ResizeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResizeInstanceResponse), nil
	}
}

type RollbackInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RollbackInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RollbackInstanceInvoker) Invoke() (*model.RollbackInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RollbackInstanceResponse), nil
	}
}

type ShowAuthorizationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuthorizationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuthorizationInvoker) Invoke() (*model.ShowAuthorizationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuthorizationResponse), nil
	}
}

type ShowEcsQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEcsQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEcsQuotaInvoker) Invoke() (*model.ShowEcsQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEcsQuotaResponse), nil
	}
}

type ShowInstanceStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceStatusInvoker) Invoke() (*model.ShowInstanceStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceStatusResponse), nil
	}
}

type ShowInstanceTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowInstanceTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowInstanceTagsInvoker) Invoke() (*model.ShowInstanceTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowInstanceTagsResponse), nil
	}
}

type ShowOmUrlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOmUrlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowOmUrlInvoker) Invoke() (*model.ShowOmUrlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOmUrlResponse), nil
	}
}

type ShowQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotaInvoker) Invoke() (*model.ShowQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotaResponse), nil
	}
}

type StartInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartInstanceInvoker) Invoke() (*model.StartInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartInstanceResponse), nil
	}
}

type StopInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopInstanceInvoker) Invoke() (*model.StopInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopInstanceResponse), nil
	}
}

type SwitchInstanceVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchInstanceVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchInstanceVpcInvoker) Invoke() (*model.SwitchInstanceVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchInstanceVpcResponse), nil
	}
}

type UninstallInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallInstanceEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UninstallInstanceEipInvoker) Invoke() (*model.UninstallInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallInstanceEipResponse), nil
	}
}

type UpdateInstanceSecurityGroupInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateInstanceSecurityGroupInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateInstanceSecurityGroupInvoker) Invoke() (*model.UpdateInstanceSecurityGroupResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateInstanceSecurityGroupResponse), nil
	}
}

type UpgradeInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeInstanceInvoker) Invoke() (*model.UpgradeInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeInstanceResponse), nil
	}
}
