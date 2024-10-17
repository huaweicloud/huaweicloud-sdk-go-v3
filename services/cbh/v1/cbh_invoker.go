package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v1/model"
)

type ChangeInstanceNetworkInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstanceNetworkInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeInstanceNetworkInvoker) Invoke() (*model.ChangeInstanceNetworkResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceNetworkResponse), nil
	}
}

type ChangeInstanceOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstanceOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeInstanceOrderInvoker) Invoke() (*model.ChangeInstanceOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceOrderResponse), nil
	}
}

type CreateCbhInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCbhInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateCbhInvoker) Invoke() (*model.CreateCbhResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCbhResponse), nil
	}
}

type CreateInstanceOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceOrderInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateInstanceOrderInvoker) Invoke() (*model.CreateInstanceOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceOrderResponse), nil
	}
}

type InstallCbhEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallCbhEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *InstallCbhEipInvoker) Invoke() (*model.InstallCbhEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallCbhEipResponse), nil
	}
}

type ListCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCbhInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCbhInstanceInvoker) Invoke() (*model.ListCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCbhInstanceResponse), nil
	}
}

type ListQuotaStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListQuotaStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListQuotaStatusInvoker) Invoke() (*model.ListQuotaStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListQuotaStatusResponse), nil
	}
}

type ResetLoginMethodInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetLoginMethodInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetLoginMethodInvoker) Invoke() (*model.ResetLoginMethodResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetLoginMethodResponse), nil
	}
}

type ResetPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ResetPasswordInvoker) Invoke() (*model.ResetPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPasswordResponse), nil
	}
}

type RestartCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RestartCbhInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RestartCbhInstanceInvoker) Invoke() (*model.RestartCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RestartCbhInstanceResponse), nil
	}
}

type SearchQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *SearchQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SearchQuotaInvoker) Invoke() (*model.SearchQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SearchQuotaResponse), nil
	}
}

type ShowAvailableZoneInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAvailableZoneInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAvailableZoneInfoInvoker) Invoke() (*model.ShowAvailableZoneInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAvailableZoneInfoResponse), nil
	}
}

type ShowNetworkConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNetworkConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowNetworkConfigurationInvoker) Invoke() (*model.ShowNetworkConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNetworkConfigurationResponse), nil
	}
}

type StartCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StartCbhInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StartCbhInstanceInvoker) Invoke() (*model.StartCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StartCbhInstanceResponse), nil
	}
}

type StopCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *StopCbhInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *StopCbhInstanceInvoker) Invoke() (*model.StopCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopCbhInstanceResponse), nil
	}
}

type UninstallCbhEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallCbhEipInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UninstallCbhEipInvoker) Invoke() (*model.UninstallCbhEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallCbhEipResponse), nil
	}
}

type UpgradeCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeCbhInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpgradeCbhInstanceInvoker) Invoke() (*model.UpgradeCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeCbhInstanceResponse), nil
	}
}

type LoginCbhInvoker struct {
	*invoker.BaseInvoker
}

func (i *LoginCbhInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LoginCbhInvoker) Invoke() (*model.LoginCbhResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoginCbhResponse), nil
	}
}
