package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/clouddc/v1/model"
)

type BatchCreateIrackTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateIrackTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateIrackTagsInvoker) Invoke() (*model.BatchCreateIrackTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateIrackTagsResponse), nil
	}
}

type BatchCreateTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchCreateTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchCreateTagsInvoker) Invoke() (*model.BatchCreateTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchCreateTagsResponse), nil
	}
}

type BatchDeleteIrackTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteIrackTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteIrackTagsInvoker) Invoke() (*model.BatchDeleteIrackTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteIrackTagsResponse), nil
	}
}

type BatchDeleteTagsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDeleteTagsInvoker) Invoke() (*model.BatchDeleteTagsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDeleteTagsResponse), nil
	}
}

type ChangeInstancePasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeInstancePasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeInstancePasswordInvoker) Invoke() (*model.ChangeInstancePasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeInstancePasswordResponse), nil
	}
}

type ChangeServerPowerStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeServerPowerStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeServerPowerStateInvoker) Invoke() (*model.ChangeServerPowerStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeServerPowerStateResponse), nil
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

type DeleteInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteInstancesInvoker) Invoke() (*model.DeleteInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteInstancesResponse), nil
	}
}

type DownloadServerLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DownloadServerLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DownloadServerLogsInvoker) Invoke() (*model.DownloadServerLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DownloadServerLogsResponse), nil
	}
}

type ExportServerLogsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExportServerLogsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ExportServerLogsInvoker) Invoke() (*model.ExportServerLogsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExportServerLogsResponse), nil
	}
}

type ListAlarmsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAlarmsInvoker) Invoke() (*model.ListAlarmsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmsResponse), nil
	}
}

type ListEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEventsInvoker) Invoke() (*model.ListEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEventsResponse), nil
	}
}

type ListIDcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIDcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIDcsInvoker) Invoke() (*model.ListIDcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIDcsResponse), nil
	}
}

type ListIRacksInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIRacksInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIRacksInvoker) Invoke() (*model.ListIRacksResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIRacksResponse), nil
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

type ListServersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListServersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListServersInvoker) Invoke() (*model.ListServersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListServersResponse), nil
	}
}

type ModifyInstanceIpInvoker struct {
	*invoker.BaseInvoker
}

func (i *ModifyInstanceIpInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ModifyInstanceIpInvoker) Invoke() (*model.ModifyInstanceIpResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ModifyInstanceIpResponse), nil
	}
}

type ReinstallOsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReinstallOsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ReinstallOsInvoker) Invoke() (*model.ReinstallOsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReinstallOsResponse), nil
	}
}

type RunInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunInstancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunInstancesInvoker) Invoke() (*model.RunInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunInstancesResponse), nil
	}
}

type ShowAlarmSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmSummaryInvoker) Invoke() (*model.ShowAlarmSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmSummaryResponse), nil
	}
}

type ShowAlarmTrendInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAlarmTrendInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAlarmTrendInvoker) Invoke() (*model.ShowAlarmTrendResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAlarmTrendResponse), nil
	}
}

type ShowEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEventInvoker) Invoke() (*model.ShowEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEventResponse), nil
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

type ShowLogsExportStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLogsExportStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLogsExportStatusInvoker) Invoke() (*model.ShowLogsExportStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLogsExportStatusResponse), nil
	}
}

type ShowRemoteConsoleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRemoteConsoleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRemoteConsoleInvoker) Invoke() (*model.ShowRemoteConsoleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRemoteConsoleResponse), nil
	}
}

type ShowServerInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerInvoker) Invoke() (*model.ShowServerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerResponse), nil
	}
}

type ShowServerFirmwareAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerFirmwareAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerFirmwareAttributesInvoker) Invoke() (*model.ShowServerFirmwareAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerFirmwareAttributesResponse), nil
	}
}

type ShowServerHardwareAttributesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerHardwareAttributesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerHardwareAttributesInvoker) Invoke() (*model.ShowServerHardwareAttributesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerHardwareAttributesResponse), nil
	}
}

type ShowServerStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowServerStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowServerStatusInvoker) Invoke() (*model.ShowServerStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowServerStatusResponse), nil
	}
}

type UpdateIDcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIDcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIDcsInvoker) Invoke() (*model.UpdateIDcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIDcsResponse), nil
	}
}

type UpdateIRackInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIRackInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIRackInvoker) Invoke() (*model.UpdateIRackResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIRackResponse), nil
	}
}
