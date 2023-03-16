package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cbh/v1/model"
)

type ChangeInstanceNetworkInvoker struct {
	*invoker.BaseInvoker
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

func (i *ChangeInstanceOrderInvoker) Invoke() (*model.ChangeInstanceOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstanceOrderResponse), nil
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

type CreateInstanceOrderInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstanceOrderInvoker) Invoke() (*model.CreateInstanceOrderResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstanceOrderResponse), nil
	}
}

type InstallInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *InstallInstanceEipInvoker) Invoke() (*model.InstallInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstallInstanceEipResponse), nil
	}
}

type ListCbhInstanceInvoker struct {
	*invoker.BaseInvoker
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

func (i *StopCbhInstanceInvoker) Invoke() (*model.StopCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.StopCbhInstanceResponse), nil
	}
}

type UninstallInstanceEipInvoker struct {
	*invoker.BaseInvoker
}

func (i *UninstallInstanceEipInvoker) Invoke() (*model.UninstallInstanceEipResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UninstallInstanceEipResponse), nil
	}
}

type UpgradeCbhInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpgradeCbhInstanceInvoker) Invoke() (*model.UpgradeCbhInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpgradeCbhInstanceResponse), nil
	}
}
