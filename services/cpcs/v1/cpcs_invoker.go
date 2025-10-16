package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cpcs/v1/model"
)

type AddClusterPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *AddClusterPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AddClusterPortInvoker) Invoke() (*model.AddClusterPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AddClusterPortResponse), nil
	}
}

type AssociateAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *AssociateAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AssociateAppsInvoker) Invoke() (*model.AssociateAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssociateAppsResponse), nil
	}
}

type AuthorizeAccessKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *AuthorizeAccessKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *AuthorizeAccessKeysInvoker) Invoke() (*model.AuthorizeAccessKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizeAccessKeysResponse), nil
	}
}

type BatchDisableAccessKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchDisableAccessKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchDisableAccessKeysInvoker) Invoke() (*model.BatchDisableAccessKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchDisableAccessKeysResponse), nil
	}
}

type BatchEnableAccessKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *BatchEnableAccessKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *BatchEnableAccessKeysInvoker) Invoke() (*model.BatchEnableAccessKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.BatchEnableAccessKeysResponse), nil
	}
}

type CancelAuthorizeAccessKeysInvoker struct {
	*invoker.BaseInvoker
}

func (i *CancelAuthorizeAccessKeysInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CancelAuthorizeAccessKeysInvoker) Invoke() (*model.CancelAuthorizeAccessKeysResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CancelAuthorizeAccessKeysResponse), nil
	}
}

type CheckClusterPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckClusterPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckClusterPortInvoker) Invoke() (*model.CheckClusterPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckClusterPortResponse), nil
	}
}

type CreateAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppInvoker) Invoke() (*model.CreateAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppResponse), nil
	}
}

type CreateAppAccessKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAppAccessKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAppAccessKeyInvoker) Invoke() (*model.CreateAppAccessKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAppAccessKeyResponse), nil
	}
}

type CreateClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateClusterInvoker) Invoke() (*model.CreateClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateClusterResponse), nil
	}
}

type DeleteAccessKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAccessKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAccessKeyInvoker) Invoke() (*model.DeleteAccessKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAccessKeyResponse), nil
	}
}

type DeleteAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteAppInvoker) Invoke() (*model.DeleteAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteAppResponse), nil
	}
}

type DeleteCcspClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCcspClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteCcspClusterInvoker) Invoke() (*model.DeleteCcspClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCcspClusterResponse), nil
	}
}

type DeleteClusterPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteClusterPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteClusterPortInvoker) Invoke() (*model.DeleteClusterPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteClusterPortResponse), nil
	}
}

type DisableCcspInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableCcspInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableCcspInstanceInvoker) Invoke() (*model.DisableCcspInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableCcspInstanceResponse), nil
	}
}

type DisassociateAppsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisassociateAppsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisassociateAppsInvoker) Invoke() (*model.DisassociateAppsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisassociateAppsResponse), nil
	}
}

type EnableCcspInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableCcspInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableCcspInstanceInvoker) Invoke() (*model.EnableCcspInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableCcspInstanceResponse), nil
	}
}

type ListCcspTenantImagesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCcspTenantImagesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListCcspTenantImagesInvoker) Invoke() (*model.ListCcspTenantImagesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCcspTenantImagesResponse), nil
	}
}

type ListClusterPortInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListClusterPortInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListClusterPortInvoker) Invoke() (*model.ListClusterPortResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListClusterPortResponse), nil
	}
}

type ShowAccessKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAccessKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAccessKeyInvoker) Invoke() (*model.ShowAccessKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAccessKeyResponse), nil
	}
}

type ShowAppAccessKeyListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppAccessKeyListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppAccessKeyListInvoker) Invoke() (*model.ShowAppAccessKeyListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppAccessKeyListResponse), nil
	}
}

type ShowAppListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAppListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAppListInvoker) Invoke() (*model.ShowAppListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAppListResponse), nil
	}
}

type ShowAssociationListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssociationListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssociationListInvoker) Invoke() (*model.ShowAssociationListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssociationListResponse), nil
	}
}

type ShowAuditLogInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAuditLogInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAuditLogInvoker) Invoke() (*model.ShowAuditLogResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAuditLogResponse), nil
	}
}

type ShowAvailableAzInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAvailableAzInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAvailableAzInvoker) Invoke() (*model.ShowAvailableAzResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAvailableAzResponse), nil
	}
}

type ShowCcspClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCcspClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCcspClusterInvoker) Invoke() (*model.ShowCcspClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCcspClusterResponse), nil
	}
}

type ShowCcspClusterListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCcspClusterListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCcspClusterListInvoker) Invoke() (*model.ShowCcspClusterListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCcspClusterListResponse), nil
	}
}

type ShowCcspInstanceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowCcspInstanceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowCcspInstanceInfoInvoker) Invoke() (*model.ShowCcspInstanceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowCcspInstanceInfoResponse), nil
	}
}

type ShowClusterAccessKeyListInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterAccessKeyListInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterAccessKeyListInvoker) Invoke() (*model.ShowClusterAccessKeyListResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterAccessKeyListResponse), nil
	}
}

type ShowClusterUriInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowClusterUriInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowClusterUriInvoker) Invoke() (*model.ShowClusterUriResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowClusterUriResponse), nil
	}
}

type ShowResourceDetailAccessKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceDetailAccessKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceDetailAccessKeyInvoker) Invoke() (*model.ShowResourceDetailAccessKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceDetailAccessKeyResponse), nil
	}
}

type ShowResourceDetailCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceDetailCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceDetailCertificateInvoker) Invoke() (*model.ShowResourceDetailCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceDetailCertificateResponse), nil
	}
}

type ShowResourceInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceInfoInvoker) Invoke() (*model.ShowResourceInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceInfoResponse), nil
	}
}

type ShowStatisticCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticCertificateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticCertificateInvoker) Invoke() (*model.ShowStatisticCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticCertificateResponse), nil
	}
}

type ShowStatisticInterfaceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticInterfaceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticInterfaceInvoker) Invoke() (*model.ShowStatisticInterfaceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticInterfaceResponse), nil
	}
}

type ShowStatisticResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticResourceInvoker) Invoke() (*model.ShowStatisticResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticResourceResponse), nil
	}
}

type ShowStatisticSecretKeyInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatisticSecretKeyInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatisticSecretKeyInvoker) Invoke() (*model.ShowStatisticSecretKeyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatisticSecretKeyResponse), nil
	}
}

type ShowStatusAppInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatusAppInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatusAppInvoker) Invoke() (*model.ShowStatusAppResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatusAppResponse), nil
	}
}

type ShowStatusClusterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatusClusterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatusClusterInvoker) Invoke() (*model.ShowStatusClusterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatusClusterResponse), nil
	}
}

type ShowStatusInstanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatusInstanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatusInstanceInvoker) Invoke() (*model.ShowStatusInstanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatusInstanceResponse), nil
	}
}

type ShowStatusServiceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowStatusServiceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowStatusServiceInvoker) Invoke() (*model.ShowStatusServiceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowStatusServiceResponse), nil
	}
}

type ShowVmMonitorInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowVmMonitorInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowVmMonitorInvoker) Invoke() (*model.ShowVmMonitorResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowVmMonitorResponse), nil
	}
}

type SwitchCpcsTokenInvoker struct {
	*invoker.BaseInvoker
}

func (i *SwitchCpcsTokenInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *SwitchCpcsTokenInvoker) Invoke() (*model.SwitchCpcsTokenResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SwitchCpcsTokenResponse), nil
	}
}
