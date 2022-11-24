package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bms/v1/model"
)

type AddServerNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddServerNicsInvoker) Invoke() (*model.AddServerNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddServerNicsResponse), nil
	}
}

type AttachBaremetalServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *AttachBaremetalServerVolumeInvoker) Invoke() (*model.AttachBaremetalServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AttachBaremetalServerVolumeResponse), nil
	}
}

type BatchCreateBaremetalServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateBaremetalServerTagsInvoker) Invoke() (*model.BatchCreateBaremetalServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateBaremetalServerTagsResponse), nil
	}
}

type BatchDeleteBaremetalServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteBaremetalServerTagsInvoker) Invoke() (*model.BatchDeleteBaremetalServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteBaremetalServerTagsResponse), nil
	}
}

type BatchRebootBaremetalServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchRebootBaremetalServersInvoker) Invoke() (*model.BatchRebootBaremetalServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchRebootBaremetalServersResponse), nil
	}
}

type BatchStartBaremetalServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStartBaremetalServersInvoker) Invoke() (*model.BatchStartBaremetalServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStartBaremetalServersResponse), nil
	}
}

type BatchStopBaremetalServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchStopBaremetalServersInvoker) Invoke() (*model.BatchStopBaremetalServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchStopBaremetalServersResponse), nil
	}
}

type ChangeBaremetalServerNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeBaremetalServerNameInvoker) Invoke() (*model.ChangeBaremetalServerNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeBaremetalServerNameResponse), nil
	}
}

type ChangeBaremetalServerOsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeBaremetalServerOsInvoker) Invoke() (*model.ChangeBaremetalServerOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeBaremetalServerOsResponse), nil
	}
}

type CreateBareMetalServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBareMetalServersInvoker) Invoke() (*model.CreateBareMetalServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBareMetalServersResponse), nil
	}
}

type DeleteServerNicsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteServerNicsInvoker) Invoke() (*model.DeleteServerNicsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteServerNicsResponse), nil
	}
}

type DeleteWindowsBareMetalServerPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteWindowsBareMetalServerPasswordInvoker) Invoke() (*model.DeleteWindowsBareMetalServerPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteWindowsBareMetalServerPasswordResponse), nil
	}
}

type DetachBaremetalServerVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DetachBaremetalServerVolumeInvoker) Invoke() (*model.DetachBaremetalServerVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DetachBaremetalServerVolumeResponse), nil
	}
}

type ListBareMetalServerDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBareMetalServerDetailsInvoker) Invoke() (*model.ListBareMetalServerDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBareMetalServerDetailsResponse), nil
	}
}

type ListBareMetalServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBareMetalServersInvoker) Invoke() (*model.ListBareMetalServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBareMetalServersResponse), nil
	}
}

type ListBaremetalFlavorDetailExtendsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListBaremetalFlavorDetailExtendsInvoker) Invoke() (*model.ListBaremetalFlavorDetailExtendsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListBaremetalFlavorDetailExtendsResponse), nil
	}
}

type ReinstallBaremetalServerOsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallBaremetalServerOsInvoker) Invoke() (*model.ReinstallBaremetalServerOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallBaremetalServerOsResponse), nil
	}
}

type ResetPwdOneClickInvoker struct {
	*invoker.BaseInvoker
}

func (i *ResetPwdOneClickInvoker) Invoke() (*model.ResetPwdOneClickResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ResetPwdOneClickResponse), nil
	}
}

type ShowBaremetalServerInterfaceAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaremetalServerInterfaceAttachmentsInvoker) Invoke() (*model.ShowBaremetalServerInterfaceAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaremetalServerInterfaceAttachmentsResponse), nil
	}
}

type ShowBaremetalServerTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaremetalServerTagsInvoker) Invoke() (*model.ShowBaremetalServerTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaremetalServerTagsResponse), nil
	}
}

type ShowBaremetalServerVolumeInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBaremetalServerVolumeInfoInvoker) Invoke() (*model.ShowBaremetalServerVolumeInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBaremetalServerVolumeInfoResponse), nil
	}
}

type ShowResetPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResetPwdInvoker) Invoke() (*model.ShowResetPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResetPwdResponse), nil
	}
}

type ShowServerRemoteConsoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerRemoteConsoleInvoker) Invoke() (*model.ShowServerRemoteConsoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerRemoteConsoleResponse), nil
	}
}

type ShowTenantQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTenantQuotaInvoker) Invoke() (*model.ShowTenantQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTenantQuotaResponse), nil
	}
}

type ShowWindowsBaremetalServerPwdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowWindowsBaremetalServerPwdInvoker) Invoke() (*model.ShowWindowsBaremetalServerPwdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowWindowsBaremetalServerPwdResponse), nil
	}
}

type UpdateBaremetalServerInterfaceAttachmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBaremetalServerInterfaceAttachmentsInvoker) Invoke() (*model.UpdateBaremetalServerInterfaceAttachmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBaremetalServerInterfaceAttachmentsResponse), nil
	}
}

type UpdateBaremetalServerMetadataInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateBaremetalServerMetadataInvoker) Invoke() (*model.UpdateBaremetalServerMetadataResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateBaremetalServerMetadataResponse), nil
	}
}

type ShowSpecifiedVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowSpecifiedVersionInvoker) Invoke() (*model.ShowSpecifiedVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowSpecifiedVersionResponse), nil
	}
}

type ShowJobInfosInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInfosInvoker) Invoke() (*model.ShowJobInfosResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobInfosResponse), nil
	}
}
