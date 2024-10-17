package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dris/v1/model"
)

type AddForwardingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddForwardingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddForwardingConfigsInvoker) Invoke() (*model.AddForwardingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddForwardingConfigsResponse), nil
	}
}

type DeleteForwardingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteForwardingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteForwardingConfigInvoker) Invoke() (*model.DeleteForwardingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteForwardingConfigResponse), nil
	}
}

type ShowForwardingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowForwardingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowForwardingConfigInvoker) Invoke() (*model.ShowForwardingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowForwardingConfigResponse), nil
	}
}

type ShowForwardingConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowForwardingConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowForwardingConfigsInvoker) Invoke() (*model.ShowForwardingConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowForwardingConfigsResponse), nil
	}
}

type UpdateForwardingConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateForwardingConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateForwardingConfigInvoker) Invoke() (*model.UpdateForwardingConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateForwardingConfigResponse), nil
	}
}

type ListEdgeFlowsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeFlowsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEdgeFlowsInvoker) Invoke() (*model.ListEdgeFlowsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeFlowsResponse), nil
	}
}

type ShowHistoryTrafficEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHistoryTrafficEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowHistoryTrafficEventsInvoker) Invoke() (*model.ShowHistoryTrafficEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHistoryTrafficEventsResponse), nil
	}
}

type SendImmediateEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *SendImmediateEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SendImmediateEventInvoker) Invoke() (*model.SendImmediateEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendImmediateEventResponse), nil
	}
}

type CreateDataChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDataChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateDataChannelInvoker) Invoke() (*model.CreateDataChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDataChannelResponse), nil
	}
}

type DeleteDataChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDataChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteDataChannelInvoker) Invoke() (*model.DeleteDataChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDataChannelResponse), nil
	}
}

type ShowDataChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDataChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDataChannelInvoker) Invoke() (*model.ShowDataChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDataChannelResponse), nil
	}
}

type UpdateDataChannelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDataChannelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateDataChannelInvoker) Invoke() (*model.UpdateDataChannelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDataChannelResponse), nil
	}
}

type CreateV2xEdgeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2xEdgeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2xEdgeInvoker) Invoke() (*model.CreateV2xEdgeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2xEdgeResponse), nil
	}
}

type DeleteV2XEdgeByV2xEdgeIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2XEdgeByV2xEdgeIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2XEdgeByV2xEdgeIdInvoker) Invoke() (*model.DeleteV2XEdgeByV2xEdgeIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2XEdgeByV2xEdgeIdResponse), nil
	}
}

type ListV2xEdgesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2xEdgesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2xEdgesInvoker) Invoke() (*model.ListV2xEdgesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2xEdgesResponse), nil
	}
}

type ShowDeploymentCodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDeploymentCodeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowDeploymentCodeInvoker) Invoke() (*model.ShowDeploymentCodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDeploymentCodeResponse), nil
	}
}

type ShowV2xEdgeDetailInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2xEdgeDetailInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2xEdgeDetailInvoker) Invoke() (*model.ShowV2xEdgeDetailResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2xEdgeDetailResponse), nil
	}
}

type UpdateV2xEdgeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2xEdgeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2xEdgeInvoker) Invoke() (*model.UpdateV2xEdgeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2xEdgeResponse), nil
	}
}

type BatchShowIpcsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowIpcsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowIpcsInvoker) Invoke() (*model.BatchShowIpcsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowIpcsResponse), nil
	}
}

type ShowIpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIpcInvoker) Invoke() (*model.ShowIpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIpcResponse), nil
	}
}

type BatchShowRadarsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowRadarsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowRadarsInvoker) Invoke() (*model.BatchShowRadarsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowRadarsResponse), nil
	}
}

type BatchShowRsusInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowRsusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowRsusInvoker) Invoke() (*model.BatchShowRsusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowRsusResponse), nil
	}
}

type CreateRsuInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRsuInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRsuInvoker) Invoke() (*model.CreateRsuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRsuResponse), nil
	}
}

type DeleteRsuInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRsuInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRsuInvoker) Invoke() (*model.DeleteRsuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRsuResponse), nil
	}
}

type UpdateRsuInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRsuInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRsuInvoker) Invoke() (*model.UpdateRsuResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRsuResponse), nil
	}
}

type BatchShowTrafficControllersInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowTrafficControllersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowTrafficControllersInvoker) Invoke() (*model.BatchShowTrafficControllersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowTrafficControllersResponse), nil
	}
}

type CreateTrafficControllerInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrafficControllerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrafficControllerInvoker) Invoke() (*model.CreateTrafficControllerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrafficControllerResponse), nil
	}
}

type DeleteTrafficControllerInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrafficControllerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrafficControllerInvoker) Invoke() (*model.DeleteTrafficControllerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrafficControllerResponse), nil
	}
}

type UpdateTrafficControllerInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrafficControllerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTrafficControllerInvoker) Invoke() (*model.UpdateTrafficControllerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrafficControllerResponse), nil
	}
}

type BatchShowTrafficEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowTrafficEventsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowTrafficEventsInvoker) Invoke() (*model.BatchShowTrafficEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowTrafficEventsResponse), nil
	}
}

type CreateTrafficEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTrafficEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTrafficEventInvoker) Invoke() (*model.CreateTrafficEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTrafficEventResponse), nil
	}
}

type DeleteTrafficEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTrafficEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTrafficEventInvoker) Invoke() (*model.DeleteTrafficEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTrafficEventResponse), nil
	}
}

type ShowTrafficEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTrafficEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTrafficEventInvoker) Invoke() (*model.ShowTrafficEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTrafficEventResponse), nil
	}
}

type UpdateTrafficEventInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTrafficEventInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateTrafficEventInvoker) Invoke() (*model.UpdateTrafficEventResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTrafficEventResponse), nil
	}
}

type CreateV2xEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateV2xEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateV2xEdgeAppInvoker) Invoke() (*model.CreateV2xEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateV2xEdgeAppResponse), nil
	}
}

type DeleteV2XEdgeAppByEdgeAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteV2XEdgeAppByEdgeAppIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteV2XEdgeAppByEdgeAppIdInvoker) Invoke() (*model.DeleteV2XEdgeAppByEdgeAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteV2XEdgeAppByEdgeAppIdResponse), nil
	}
}

type ListV2xEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListV2xEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListV2xEdgeAppInvoker) Invoke() (*model.ListV2xEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListV2xEdgeAppResponse), nil
	}
}

type ShowV2XEdgeAppDetailByEdgeAppIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowV2XEdgeAppDetailByEdgeAppIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowV2XEdgeAppDetailByEdgeAppIdInvoker) Invoke() (*model.ShowV2XEdgeAppDetailByEdgeAppIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowV2XEdgeAppDetailByEdgeAppIdResponse), nil
	}
}

type UpdateV2xEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateV2xEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateV2xEdgeAppInvoker) Invoke() (*model.UpdateV2xEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateV2xEdgeAppResponse), nil
	}
}

type BatchShowVehiclesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowVehiclesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowVehiclesInvoker) Invoke() (*model.BatchShowVehiclesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowVehiclesResponse), nil
	}
}

type CreateVehicleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVehicleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateVehicleInvoker) Invoke() (*model.CreateVehicleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVehicleResponse), nil
	}
}

type DeleteVehicleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVehicleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVehicleInvoker) Invoke() (*model.DeleteVehicleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVehicleResponse), nil
	}
}

type UpdateVehicleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVehicleInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVehicleInvoker) Invoke() (*model.UpdateVehicleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVehicleResponse), nil
	}
}

type BatchShowEdgeAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowEdgeAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowEdgeAppsInvoker) Invoke() (*model.BatchShowEdgeAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowEdgeAppsResponse), nil
	}
}

type CreateEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeAppInvoker) Invoke() (*model.CreateEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeAppResponse), nil
	}
}

type DeleteEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeAppInvoker) Invoke() (*model.DeleteEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeAppResponse), nil
	}
}

type UpdateEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeAppInvoker) Invoke() (*model.UpdateEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeAppResponse), nil
	}
}

type BatchShowEdgeAppVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchShowEdgeAppVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchShowEdgeAppVersionsInvoker) Invoke() (*model.BatchShowEdgeAppVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchShowEdgeAppVersionsResponse), nil
	}
}

type CreateEdgeApplicationVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeApplicationVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEdgeApplicationVersionInvoker) Invoke() (*model.CreateEdgeApplicationVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeApplicationVersionResponse), nil
	}
}

type DeleteEdgeApplicationVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeApplicationVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEdgeApplicationVersionInvoker) Invoke() (*model.DeleteEdgeApplicationVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeApplicationVersionResponse), nil
	}
}

type ShowEdgeApplicationVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeApplicationVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEdgeApplicationVersionInvoker) Invoke() (*model.ShowEdgeApplicationVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeApplicationVersionResponse), nil
	}
}

type UpdateEdgeApplicationVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeApplicationVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeApplicationVersionInvoker) Invoke() (*model.UpdateEdgeApplicationVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeApplicationVersionResponse), nil
	}
}

type UpdateEdgeApplicationVersionStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeApplicationVersionStateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEdgeApplicationVersionStateInvoker) Invoke() (*model.UpdateEdgeApplicationVersionStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeApplicationVersionStateResponse), nil
	}
}

type CreateRsuModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateRsuModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateRsuModelInvoker) Invoke() (*model.CreateRsuModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateRsuModelResponse), nil
	}
}

type DeleteRsuModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteRsuModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteRsuModelInvoker) Invoke() (*model.DeleteRsuModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteRsuModelResponse), nil
	}
}

type ListRsuModelsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRsuModelsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListRsuModelsInvoker) Invoke() (*model.ListRsuModelsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRsuModelsResponse), nil
	}
}

type ShowRsuModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowRsuModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowRsuModelInvoker) Invoke() (*model.ShowRsuModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowRsuModelResponse), nil
	}
}

type UpdateRsuModelInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRsuModelInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateRsuModelInvoker) Invoke() (*model.UpdateRsuModelResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRsuModelResponse), nil
	}
}
