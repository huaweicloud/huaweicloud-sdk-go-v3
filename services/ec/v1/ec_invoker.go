package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ec/v1/model"
)

type CreateEcnAccessPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEcnAccessPointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEcnAccessPointInvoker) Invoke() (*model.CreateEcnAccessPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEcnAccessPointResponse), nil
	}
}

type DeleteEcnAccessPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEcnAccessPointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEcnAccessPointInvoker) Invoke() (*model.DeleteEcnAccessPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEcnAccessPointResponse), nil
	}
}

type ListEcnAccessPointByEcnIdInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcnAccessPointByEcnIdInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcnAccessPointByEcnIdInvoker) Invoke() (*model.ListEcnAccessPointByEcnIdResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcnAccessPointByEcnIdResponse), nil
	}
}

type UpdateEcnAccessPointInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEcnAccessPointInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEcnAccessPointInvoker) Invoke() (*model.UpdateEcnAccessPointResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEcnAccessPointResponse), nil
	}
}

type AddEcnWithIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddEcnWithIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddEcnWithIegInvoker) Invoke() (*model.AddEcnWithIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddEcnWithIegResponse), nil
	}
}

type DeleteEcnWithIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEcnWithIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEcnWithIegInvoker) Invoke() (*model.DeleteEcnWithIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEcnWithIegResponse), nil
	}
}

type ListEcnInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcnInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcnInvoker) Invoke() (*model.ListEcnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcnResponse), nil
	}
}

type ListEcnWithIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcnWithIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcnWithIegInvoker) Invoke() (*model.ListEcnWithIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcnWithIegResponse), nil
	}
}

type ShowEcnInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEcnInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEcnInfoInvoker) Invoke() (*model.ShowEcnInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEcnInfoResponse), nil
	}
}

type ShowEcnWithIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEcnWithIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEcnWithIegInvoker) Invoke() (*model.ShowEcnWithIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEcnWithIegResponse), nil
	}
}

type UpdateEcnInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEcnInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEcnInvoker) Invoke() (*model.UpdateEcnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEcnResponse), nil
	}
}

type CreateEquipmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEquipmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEquipmentInvoker) Invoke() (*model.CreateEquipmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEquipmentResponse), nil
	}
}

type DeleteEquipmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEquipmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEquipmentInvoker) Invoke() (*model.DeleteEquipmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEquipmentResponse), nil
	}
}

type GenerateInitialConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *GenerateInitialConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GenerateInitialConfigurationInvoker) Invoke() (*model.GenerateInitialConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GenerateInitialConfigurationResponse), nil
	}
}

type ListEquipmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEquipmentsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEquipmentsInvoker) Invoke() (*model.ListEquipmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEquipmentsResponse), nil
	}
}

type RebootEquipmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RebootEquipmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RebootEquipmentInvoker) Invoke() (*model.RebootEquipmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RebootEquipmentResponse), nil
	}
}

type ShowEquipmentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentInfoInvoker) Invoke() (*model.ShowEquipmentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentInfoResponse), nil
	}
}

type ShowEquipmentSpecificConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentSpecificConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentSpecificConfigInvoker) Invoke() (*model.ShowEquipmentSpecificConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentSpecificConfigResponse), nil
	}
}

type UpdateEquipmentEsnInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentEsnInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentEsnInvoker) Invoke() (*model.UpdateEquipmentEsnResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentEsnResponse), nil
	}
}

type UpdateEquipmentInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentInfoInvoker) Invoke() (*model.UpdateEquipmentInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentInfoResponse), nil
	}
}

type CreateEquipmentLanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEquipmentLanConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEquipmentLanConfigInvoker) Invoke() (*model.CreateEquipmentLanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEquipmentLanConfigResponse), nil
	}
}

type DeleteEquipmentLanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEquipmentLanConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEquipmentLanConfigInvoker) Invoke() (*model.DeleteEquipmentLanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEquipmentLanConfigResponse), nil
	}
}

type ListEquipmentInterfaceNameInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEquipmentInterfaceNameInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEquipmentInterfaceNameInvoker) Invoke() (*model.ListEquipmentInterfaceNameResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEquipmentInterfaceNameResponse), nil
	}
}

type ShowEquipmentDnsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentDnsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentDnsInfoInvoker) Invoke() (*model.ShowEquipmentDnsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentDnsInfoResponse), nil
	}
}

type ShowEquipmentLanInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentLanInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentLanInfoInvoker) Invoke() (*model.ShowEquipmentLanInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentLanInfoResponse), nil
	}
}

type UpdateEquipmentDnsInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentDnsInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentDnsInfoInvoker) Invoke() (*model.UpdateEquipmentDnsInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentDnsInfoResponse), nil
	}
}

type UpdateEquipmentLanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentLanConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentLanConfigInvoker) Invoke() (*model.UpdateEquipmentLanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentLanConfigResponse), nil
	}
}

type ShowEquipmentOspfInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentOspfInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentOspfInvoker) Invoke() (*model.ShowEquipmentOspfResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentOspfResponse), nil
	}
}

type UpdateEquipmentOspfInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentOspfInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentOspfInvoker) Invoke() (*model.UpdateEquipmentOspfResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentOspfResponse), nil
	}
}

type CreateEquipmentStaticRouteConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEquipmentStaticRouteConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEquipmentStaticRouteConfigInvoker) Invoke() (*model.CreateEquipmentStaticRouteConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEquipmentStaticRouteConfigResponse), nil
	}
}

type DeleteEquipmentStaticRouteConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEquipmentStaticRouteConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEquipmentStaticRouteConfigInvoker) Invoke() (*model.DeleteEquipmentStaticRouteConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEquipmentStaticRouteConfigResponse), nil
	}
}

type ShowEquipmentStaticRouteInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentStaticRouteInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentStaticRouteInfoInvoker) Invoke() (*model.ShowEquipmentStaticRouteInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentStaticRouteInfoResponse), nil
	}
}

type UpdateEquipmentStaticRouteConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentStaticRouteConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentStaticRouteConfigInvoker) Invoke() (*model.UpdateEquipmentStaticRouteConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentStaticRouteConfigResponse), nil
	}
}

type ShowEquipmentWanInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentWanInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentWanInfoInvoker) Invoke() (*model.ShowEquipmentWanInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentWanInfoResponse), nil
	}
}

type UpdateEquipmentWanConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentWanConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentWanConfigInvoker) Invoke() (*model.UpdateEquipmentWanConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentWanConfigResponse), nil
	}
}

type ShowEquipmentWlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEquipmentWlanInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEquipmentWlanInvoker) Invoke() (*model.ShowEquipmentWlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEquipmentWlanResponse), nil
	}
}

type UpdateEquipmentWlanInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEquipmentWlanInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEquipmentWlanInvoker) Invoke() (*model.UpdateEquipmentWlanResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEquipmentWlanResponse), nil
	}
}

type AddEcnWithErInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddEcnWithErInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddEcnWithErInvoker) Invoke() (*model.AddEcnWithErResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddEcnWithErResponse), nil
	}
}

type DeleteEcnWithErInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEcnWithErInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEcnWithErInvoker) Invoke() (*model.DeleteEcnWithErResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEcnWithErResponse), nil
	}
}

type ListEcnWithErInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcnWithErInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcnWithErInvoker) Invoke() (*model.ListEcnWithErResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcnWithErResponse), nil
	}
}

type ChangeIegPasswordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ChangeIegPasswordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ChangeIegPasswordInvoker) Invoke() (*model.ChangeIegPasswordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChangeIegPasswordResponse), nil
	}
}

type ListIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListIegInvoker) Invoke() (*model.ListIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListIegResponse), nil
	}
}

type ShowIegInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowIegInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowIegInfoInvoker) Invoke() (*model.ShowIegInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowIegInfoResponse), nil
	}
}

type SwitchEquipmentHaTypeInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchEquipmentHaTypeInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchEquipmentHaTypeInvoker) Invoke() (*model.SwitchEquipmentHaTypeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchEquipmentHaTypeResponse), nil
	}
}

type UpdateIegInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateIegInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateIegInvoker) Invoke() (*model.UpdateIegResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateIegResponse), nil
	}
}

type ShowQuotasInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowQuotasInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowQuotasInfoInvoker) Invoke() (*model.ShowQuotasInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowQuotasInfoResponse), nil
	}
}

type AddEcnWithVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddEcnWithVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddEcnWithVpcInvoker) Invoke() (*model.AddEcnWithVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddEcnWithVpcResponse), nil
	}
}

type DeleteEcnWithVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEcnWithVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEcnWithVpcInvoker) Invoke() (*model.DeleteEcnWithVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEcnWithVpcResponse), nil
	}
}

type ListEcnWithVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEcnWithVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEcnWithVpcInvoker) Invoke() (*model.ListEcnWithVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEcnWithVpcResponse), nil
	}
}

type UpdateEcnWithVpcInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEcnWithVpcInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEcnWithVpcInvoker) Invoke() (*model.UpdateEcnWithVpcResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEcnWithVpcResponse), nil
	}
}

type AddVrrpConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddVrrpConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddVrrpConfigInvoker) Invoke() (*model.AddVrrpConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddVrrpConfigResponse), nil
	}
}

type DeleteVrrpConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVrrpConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteVrrpConfigInvoker) Invoke() (*model.DeleteVrrpConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVrrpConfigResponse), nil
	}
}

type ShowVrrpConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVrrpConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVrrpConfigInvoker) Invoke() (*model.ShowVrrpConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVrrpConfigResponse), nil
	}
}

type UpdateVrrpConfigInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateVrrpConfigInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateVrrpConfigInvoker) Invoke() (*model.UpdateVrrpConfigResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateVrrpConfigResponse), nil
	}
}
