package v5

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpn/v5/model"
)

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
