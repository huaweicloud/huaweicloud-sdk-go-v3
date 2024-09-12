package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotedge/v2/model"
)

type CreateEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEdgeNodeInvoker) Invoke() (*model.CreateEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEdgeNodeResponse), nil
	}
}

type CreateInstallCmdInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateInstallCmdInvoker) Invoke() (*model.CreateInstallCmdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateInstallCmdResponse), nil
	}
}

type DeleteEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEdgeNodeInvoker) Invoke() (*model.DeleteEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeNodeResponse), nil
	}
}

type ListEdgeNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEdgeNodesInvoker) Invoke() (*model.ListEdgeNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEdgeNodesResponse), nil
	}
}

type ShowEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeNodeInvoker) Invoke() (*model.ShowEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeNodeResponse), nil
	}
}

type ShowEdgeNodeHostsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeNodeHostsInfoInvoker) Invoke() (*model.ShowEdgeNodeHostsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeNodeHostsInfoResponse), nil
	}
}

type UpdateEdgeNodeInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEdgeNodeInvoker) Invoke() (*model.UpdateEdgeNodeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeNodeResponse), nil
	}
}

type ExecuteDeviceControlsReleaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDeviceControlsReleaseInvoker) Invoke() (*model.ExecuteDeviceControlsReleaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDeviceControlsReleaseResponse), nil
	}
}

type ExecuteDeviceControlsSetInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteDeviceControlsSetInvoker) Invoke() (*model.ExecuteDeviceControlsSetResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteDeviceControlsSetResponse), nil
	}
}

type SetDeviceControlDefaultValuesInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetDeviceControlDefaultValuesInvoker) Invoke() (*model.SetDeviceControlDefaultValuesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetDeviceControlDefaultValuesResponse), nil
	}
}

type AddDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddDeviceInvoker) Invoke() (*model.AddDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddDeviceResponse), nil
	}
}

type DeleteDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDeviceInvoker) Invoke() (*model.DeleteDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDeviceResponse), nil
	}
}

type ListDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDevicesInvoker) Invoke() (*model.ListDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDevicesResponse), nil
	}
}

type ShowProductConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProductConfigInvoker) Invoke() (*model.ShowProductConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProductConfigResponse), nil
	}
}

type UpdateDeviceInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDeviceInvoker) Invoke() (*model.UpdateDeviceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDeviceResponse), nil
	}
}

type AddAppConfigsTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddAppConfigsTemplatesInvoker) Invoke() (*model.AddAppConfigsTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddAppConfigsTemplatesResponse), nil
	}
}

type AddGeneralAppConfigsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGeneralAppConfigsTemplateInvoker) Invoke() (*model.AddGeneralAppConfigsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGeneralAppConfigsTemplateResponse), nil
	}
}

type BatchListAppConfigsTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListAppConfigsTemplatesInvoker) Invoke() (*model.BatchListAppConfigsTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListAppConfigsTemplatesResponse), nil
	}
}

type DeleteAppConfigsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppConfigsTemplateInvoker) Invoke() (*model.DeleteAppConfigsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppConfigsTemplateResponse), nil
	}
}

type ShowAppConfigsTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppConfigsTemplateInvoker) Invoke() (*model.ShowAppConfigsTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppConfigsTemplateResponse), nil
	}
}

type BatchListEdgeAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListEdgeAppsInvoker) Invoke() (*model.BatchListEdgeAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListEdgeAppsResponse), nil
	}
}

type CreateEdgeAppInvoker struct {
	*invoker.BaseInvoker
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

func (i *DeleteEdgeAppInvoker) Invoke() (*model.DeleteEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEdgeAppResponse), nil
	}
}

type ShowEdgeAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEdgeAppInvoker) Invoke() (*model.ShowEdgeAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEdgeAppResponse), nil
	}
}

type BatchListEdgeAppVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListEdgeAppVersionsInvoker) Invoke() (*model.BatchListEdgeAppVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListEdgeAppVersionsResponse), nil
	}
}

type CreateEdgeApplicationVersionInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateEdgeApplicationVersionStateInvoker) Invoke() (*model.UpdateEdgeApplicationVersionStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEdgeApplicationVersionStateResponse), nil
	}
}

type BatchListDcDsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListDcDsInvoker) Invoke() (*model.BatchListDcDsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListDcDsResponse), nil
	}
}

type CreateDsInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDsInvoker) Invoke() (*model.CreateDsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDsResponse), nil
	}
}

type DeleteDcDsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDcDsInvoker) Invoke() (*model.DeleteDcDsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDcDsResponse), nil
	}
}

type ShowDcDsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDcDsInvoker) Invoke() (*model.ShowDcDsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDcDsResponse), nil
	}
}

type SynchronizeDcConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *SynchronizeDcConfigsInvoker) Invoke() (*model.SynchronizeDcConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SynchronizeDcConfigsResponse), nil
	}
}

type UpdateDcDsInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDcDsInvoker) Invoke() (*model.UpdateDcDsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDcDsResponse), nil
	}
}

type BatchListDcDevicesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListDcDevicesInvoker) Invoke() (*model.BatchListDcDevicesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListDcDevicesResponse), nil
	}
}

type BatchListDcPointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListDcPointsInvoker) Invoke() (*model.BatchListDcPointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListDcPointsResponse), nil
	}
}

type CreateDcPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDcPointInvoker) Invoke() (*model.CreateDcPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDcPointResponse), nil
	}
}

type DeleteDcPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDcPointInvoker) Invoke() (*model.DeleteDcPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDcPointResponse), nil
	}
}

type DeleteDcPointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDcPointsInvoker) Invoke() (*model.DeleteDcPointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDcPointsResponse), nil
	}
}

type ShowDcPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowDcPointInvoker) Invoke() (*model.ShowDcPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowDcPointResponse), nil
	}
}

type UpdateDcPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateDcPointInvoker) Invoke() (*model.UpdateDcPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateDcPointResponse), nil
	}
}

type CreateExternalEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateExternalEntityInvoker) Invoke() (*model.CreateExternalEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateExternalEntityResponse), nil
	}
}

type DeleteExternalEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteExternalEntityInvoker) Invoke() (*model.DeleteExternalEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteExternalEntityResponse), nil
	}
}

type ListExternalEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExternalEntityInvoker) Invoke() (*model.ListExternalEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExternalEntityResponse), nil
	}
}

type UpdateExternalEntityInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateExternalEntityInvoker) Invoke() (*model.UpdateExternalEntityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateExternalEntityResponse), nil
	}
}

type BatchListModulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListModulesInvoker) Invoke() (*model.BatchListModulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListModulesResponse), nil
	}
}

type CreateModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateModuleInvoker) Invoke() (*model.CreateModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateModuleResponse), nil
	}
}

type DeleteModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteModuleInvoker) Invoke() (*model.DeleteModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteModuleResponse), nil
	}
}

type InvokeModuleMsgInvoker struct {
	*invoker.BaseInvoker
}

func (i *InvokeModuleMsgInvoker) Invoke() (*model.InvokeModuleMsgResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.InvokeModuleMsgResponse), nil
	}
}

type ShowModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModuleInvoker) Invoke() (*model.ShowModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModuleResponse), nil
	}
}

type UpdateModuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModuleInvoker) Invoke() (*model.UpdateModuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModuleResponse), nil
	}
}

type UpdateModuleStateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModuleStateInvoker) Invoke() (*model.UpdateModuleStateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModuleStateResponse), nil
	}
}

type ShowModuleShadowInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowModuleShadowInvoker) Invoke() (*model.ShowModuleShadowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowModuleShadowResponse), nil
	}
}

type UpdateModuleShadowInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateModuleShadowInvoker) Invoke() (*model.UpdateModuleShadowResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateModuleShadowResponse), nil
	}
}

type ListRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListRoutesInvoker) Invoke() (*model.ListRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListRoutesResponse), nil
	}
}

type UpdateRoutesInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateRoutesInvoker) Invoke() (*model.UpdateRoutesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateRoutesResponse), nil
	}
}

type AddGeneralOtTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddGeneralOtTemplateInvoker) Invoke() (*model.AddGeneralOtTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddGeneralOtTemplateResponse), nil
	}
}

type AddOtTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddOtTemplatesInvoker) Invoke() (*model.AddOtTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddOtTemplatesResponse), nil
	}
}

type BatchListOtTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchListOtTemplatesInvoker) Invoke() (*model.BatchListOtTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchListOtTemplatesResponse), nil
	}
}

type DeleteOtTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteOtTemplateInvoker) Invoke() (*model.DeleteOtTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteOtTemplateResponse), nil
	}
}

type ShowOtTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOtTemplateInvoker) Invoke() (*model.ShowOtTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOtTemplateResponse), nil
	}
}

type ImportPointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ImportPointsInvoker) Invoke() (*model.ImportPointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImportPointsResponse), nil
	}
}

type ShowPointTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPointTemplateInvoker) Invoke() (*model.ShowPointTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPointTemplateResponse), nil
	}
}

type ShowPointsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowPointsInvoker) Invoke() (*model.ShowPointsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowPointsResponse), nil
	}
}

type CreateScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateScheduleInvoker) Invoke() (*model.CreateScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateScheduleResponse), nil
	}
}

type DeleteScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteScheduleInvoker) Invoke() (*model.DeleteScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteScheduleResponse), nil
	}
}

type UpdateScheduleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateScheduleInvoker) Invoke() (*model.UpdateScheduleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateScheduleResponse), nil
	}
}

type BatchConfirmConfigsNewInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchConfirmConfigsNewInvoker) Invoke() (*model.BatchConfirmConfigsNewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchConfirmConfigsNewResponse), nil
	}
}

type BatchImportConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchImportConfigsInvoker) Invoke() (*model.BatchImportConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchImportConfigsResponse), nil
	}
}

type DeleteIaConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteIaConfigInvoker) Invoke() (*model.DeleteIaConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteIaConfigResponse), nil
	}
}

type ListIaConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIaConfigsInvoker) Invoke() (*model.ListIaConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIaConfigsResponse), nil
	}
}

type ShowIaConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIaConfigInvoker) Invoke() (*model.ShowIaConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIaConfigResponse), nil
	}
}

type UpdateIaConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIaConfigInvoker) Invoke() (*model.UpdateIaConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIaConfigResponse), nil
	}
}

type BatchAssociateNaToNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchAssociateNaToNodesInvoker) Invoke() (*model.BatchAssociateNaToNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchAssociateNaToNodesResponse), nil
	}
}

type DeleteNaInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteNaInvoker) Invoke() (*model.DeleteNaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteNaResponse), nil
	}
}

type ListNaAuthorizedNodesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNaAuthorizedNodesInvoker) Invoke() (*model.ListNaAuthorizedNodesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNaAuthorizedNodesResponse), nil
	}
}

type ListNasInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListNasInvoker) Invoke() (*model.ListNasResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListNasResponse), nil
	}
}

type ShowNaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowNaInvoker) Invoke() (*model.ShowNaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowNaResponse), nil
	}
}

type UpdateNaInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateNaInvoker) Invoke() (*model.UpdateNaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateNaResponse), nil
	}
}
