package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type DisableControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableControlInvoker) Invoke() (*model.DisableControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableControlResponse), nil
	}
}

type EnableControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableControlInvoker) Invoke() (*model.EnableControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableControlResponse), nil
	}
}

type ListConfigRuleComplianceInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigRuleComplianceInvoker) Invoke() (*model.ListConfigRuleComplianceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigRuleComplianceResponse), nil
	}
}

type ListControlViolationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlViolationsInvoker) Invoke() (*model.ListControlViolationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlViolationsResponse), nil
	}
}

type ListControlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlsInvoker) Invoke() (*model.ListControlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsResponse), nil
	}
}

type ListControlsForAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlsForAccountInvoker) Invoke() (*model.ListControlsForAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsForAccountResponse), nil
	}
}

type ListControlsForOrganizationUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlsForOrganizationUnitInvoker) Invoke() (*model.ListControlsForOrganizationUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsForOrganizationUnitResponse), nil
	}
}

type ListDriftDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDriftDetailsInvoker) Invoke() (*model.ListDriftDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDriftDetailsResponse), nil
	}
}

type ListEnabledControlsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnabledControlsInvoker) Invoke() (*model.ListEnabledControlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnabledControlsResponse), nil
	}
}

type ShowComplianceStatusForAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComplianceStatusForAccountInvoker) Invoke() (*model.ShowComplianceStatusForAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComplianceStatusForAccountResponse), nil
	}
}

type ShowComplianceStatusForOrganizationUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComplianceStatusForOrganizationUnitInvoker) Invoke() (*model.ShowComplianceStatusForOrganizationUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComplianceStatusForOrganizationUnitResponse), nil
	}
}

type ShowControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlInvoker) Invoke() (*model.ShowControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlResponse), nil
	}
}

type ShowControlOperateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlOperateInvoker) Invoke() (*model.ShowControlOperateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlOperateResponse), nil
	}
}

type ShowControlsForOrganizationUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlsForOrganizationUnitInvoker) Invoke() (*model.ShowControlsForOrganizationUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlsForOrganizationUnitResponse), nil
	}
}

type CheckLaunchInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckLaunchInvoker) Invoke() (*model.CheckLaunchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckLaunchResponse), nil
	}
}

type SetupLandingZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetupLandingZoneInvoker) Invoke() (*model.SetupLandingZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.SetupLandingZoneResponse), nil
	}
}

type ShowAvailableUpdatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAvailableUpdatesInvoker) Invoke() (*model.ShowAvailableUpdatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAvailableUpdatesResponse), nil
	}
}

type ShowHomeRegionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowHomeRegionInvoker) Invoke() (*model.ShowHomeRegionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowHomeRegionResponse), nil
	}
}

type ShowLandingZoneConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLandingZoneConfigurationInvoker) Invoke() (*model.ShowLandingZoneConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLandingZoneConfigurationResponse), nil
	}
}

type ShowLandingZoneIdentityCenterInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLandingZoneIdentityCenterInvoker) Invoke() (*model.ShowLandingZoneIdentityCenterResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLandingZoneIdentityCenterResponse), nil
	}
}

type ShowLandingZoneStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLandingZoneStatusInvoker) Invoke() (*model.ShowLandingZoneStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLandingZoneStatusResponse), nil
	}
}

type CreateAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAccountInvoker) Invoke() (*model.CreateAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccountResponse), nil
	}
}

type DeleteManagedOrganizationalUnitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteManagedOrganizationalUnitsInvoker) Invoke() (*model.DeleteManagedOrganizationalUnitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteManagedOrganizationalUnitsResponse), nil
	}
}

type DeregisterOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeregisterOrganizationalUnitInvoker) Invoke() (*model.DeregisterOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeregisterOrganizationalUnitResponse), nil
	}
}

type EnrollAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnrollAccountInvoker) Invoke() (*model.EnrollAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnrollAccountResponse), nil
	}
}

type ListManagedAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedAccountsInvoker) Invoke() (*model.ListManagedAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedAccountsResponse), nil
	}
}

type ListManagedAccountsForParentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedAccountsForParentInvoker) Invoke() (*model.ListManagedAccountsForParentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedAccountsForParentResponse), nil
	}
}

type ListManagedOrganizationalUnitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListManagedOrganizationalUnitsInvoker) Invoke() (*model.ListManagedOrganizationalUnitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedOrganizationalUnitsResponse), nil
	}
}

type ReRegisterOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReRegisterOrganizationalUnitInvoker) Invoke() (*model.ReRegisterOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReRegisterOrganizationalUnitResponse), nil
	}
}

type RegisterOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *RegisterOrganizationalUnitInvoker) Invoke() (*model.RegisterOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RegisterOrganizationalUnitResponse), nil
	}
}

type ShowManagedAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManagedAccountInvoker) Invoke() (*model.ShowManagedAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManagedAccountResponse), nil
	}
}

type ShowManagedCoreAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManagedCoreAccountInvoker) Invoke() (*model.ShowManagedCoreAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManagedCoreAccountResponse), nil
	}
}

type ShowManagedOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManagedOrganizationalUnitInvoker) Invoke() (*model.ShowManagedOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManagedOrganizationalUnitResponse), nil
	}
}

type ShowOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowOperationInvoker) Invoke() (*model.ShowOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowOperationResponse), nil
	}
}

type UnEnrollAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UnEnrollAccountInvoker) Invoke() (*model.UnEnrollAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UnEnrollAccountResponse), nil
	}
}

type UpdateManagedAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateManagedAccountInvoker) Invoke() (*model.UpdateManagedAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateManagedAccountResponse), nil
	}
}
