package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rgc/v1/model"
)

type CreateBestPracticeDetectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateBestPracticeDetectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateBestPracticeDetectInvoker) Invoke() (*model.CreateBestPracticeDetectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateBestPracticeDetectResponse), nil
	}
}

type ShowBestPracticeAccountInfoInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBestPracticeAccountInfoInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBestPracticeAccountInfoInvoker) Invoke() (*model.ShowBestPracticeAccountInfoResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBestPracticeAccountInfoResponse), nil
	}
}

type ShowBestPracticeDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBestPracticeDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBestPracticeDetailsInvoker) Invoke() (*model.ShowBestPracticeDetailsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBestPracticeDetailsResponse), nil
	}
}

type ShowBestPracticeOverviewInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBestPracticeOverviewInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBestPracticeOverviewInvoker) Invoke() (*model.ShowBestPracticeOverviewResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBestPracticeOverviewResponse), nil
	}
}

type ShowBestPracticeStatusInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowBestPracticeStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowBestPracticeStatusInvoker) Invoke() (*model.ShowBestPracticeStatusResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowBestPracticeStatusResponse), nil
	}
}

type DisableControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *EnableControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableControlInvoker) Invoke() (*model.EnableControlResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableControlResponse), nil
	}
}

type ListConfigRuleCompliancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListConfigRuleCompliancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListConfigRuleCompliancesInvoker) Invoke() (*model.ListConfigRuleCompliancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListConfigRuleCompliancesResponse), nil
	}
}

type ListControlViolationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlViolationsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListControlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListControlsForAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListControlsForAccountInvoker) Invoke() (*model.ListControlsForAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsForAccountResponse), nil
	}
}

type ListControlsForOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListControlsForOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListControlsForOrganizationalUnitInvoker) Invoke() (*model.ListControlsForOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListControlsForOrganizationalUnitResponse), nil
	}
}

type ListDriftDetailsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDriftDetailsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListEnabledControlsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnabledControlsInvoker) Invoke() (*model.ListEnabledControlsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnabledControlsResponse), nil
	}
}

type ListExternalConfigRuleCompliancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListExternalConfigRuleCompliancesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListExternalConfigRuleCompliancesInvoker) Invoke() (*model.ListExternalConfigRuleCompliancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListExternalConfigRuleCompliancesResponse), nil
	}
}

type ShowComplianceStatusForAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComplianceStatusForAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowComplianceStatusForAccountInvoker) Invoke() (*model.ShowComplianceStatusForAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComplianceStatusForAccountResponse), nil
	}
}

type ShowComplianceStatusForOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComplianceStatusForOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowComplianceStatusForOrganizationalUnitInvoker) Invoke() (*model.ShowComplianceStatusForOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComplianceStatusForOrganizationalUnitResponse), nil
	}
}

type ShowControlInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowControlOperateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowControlOperateInvoker) Invoke() (*model.ShowControlOperateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlOperateResponse), nil
	}
}

type ShowControlsForAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlsForAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowControlsForAccountInvoker) Invoke() (*model.ShowControlsForAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlsForAccountResponse), nil
	}
}

type ShowControlsForOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowControlsForOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowControlsForOrganizationalUnitInvoker) Invoke() (*model.ShowControlsForOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowControlsForOrganizationalUnitResponse), nil
	}
}

type CheckLaunchInvoker struct {
	*invoker.BaseInvoker
}

func (i *CheckLaunchInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CheckLaunchInvoker) Invoke() (*model.CheckLaunchResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CheckLaunchResponse), nil
	}
}

type DeleteLandingZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteLandingZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteLandingZoneInvoker) Invoke() (*model.DeleteLandingZoneResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteLandingZoneResponse), nil
	}
}

type SetupLandingZoneInvoker struct {
	*invoker.BaseInvoker
}

func (i *SetupLandingZoneInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowAvailableUpdatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowHomeRegionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowLandingZoneConfigurationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowLandingZoneIdentityCenterInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowLandingZoneStatusInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *CreateAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateAccountInvoker) Invoke() (*model.CreateAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAccountResponse), nil
	}
}

type CreateManagedOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateManagedOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateManagedOrganizationalUnitInvoker) Invoke() (*model.CreateManagedOrganizationalUnitResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateManagedOrganizationalUnitResponse), nil
	}
}

type DeleteManagedOrganizationalUnitsInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteManagedOrganizationalUnitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *DeregisterOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *EnrollAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListManagedAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListManagedAccountsForParentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ListManagedOrganizationalUnitsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListManagedOrganizationalUnitsInvoker) Invoke() (*model.ListManagedOrganizationalUnitsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListManagedOrganizationalUnitsResponse), nil
	}
}

type ListOperationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListOperationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListOperationInvoker) Invoke() (*model.ListOperationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListOperationResponse), nil
	}
}

type ReRegisterOrganizationalUnitInvoker struct {
	*invoker.BaseInvoker
}

func (i *ReRegisterOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *RegisterOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowManagedAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowManagedAccountInvoker) Invoke() (*model.ShowManagedAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManagedAccountResponse), nil
	}
}

type ShowManagedAccountTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManagedAccountTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowManagedAccountTemplateInvoker) Invoke() (*model.ShowManagedAccountTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowManagedAccountTemplateResponse), nil
	}
}

type ShowManagedCoreAccountInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowManagedCoreAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowManagedOrganizationalUnitInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *ShowOperationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UnEnrollAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
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

func (i *UpdateManagedAccountInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateManagedAccountInvoker) Invoke() (*model.UpdateManagedAccountResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateManagedAccountResponse), nil
	}
}

type CreateTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateTemplateInvoker) Invoke() (*model.CreateTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTemplateResponse), nil
	}
}

type DeleteTemplateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTemplateInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteTemplateInvoker) Invoke() (*model.DeleteTemplateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTemplateResponse), nil
	}
}

type ListPredefinedTemplatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListPredefinedTemplatesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListPredefinedTemplatesInvoker) Invoke() (*model.ListPredefinedTemplatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListPredefinedTemplatesResponse), nil
	}
}

type ShowTemplateDeployParamsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowTemplateDeployParamsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowTemplateDeployParamsInvoker) Invoke() (*model.ShowTemplateDeployParamsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowTemplateDeployParamsResponse), nil
	}
}
