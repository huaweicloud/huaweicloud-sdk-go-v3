package v1

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/cae/v1/model"
)

type CreateAgencyInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateAgencyInvoker) Invoke() (*model.CreateAgencyResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateAgencyResponse), nil
	}
}

type ListAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAgenciesInvoker) Invoke() (*model.ListAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAgenciesResponse), nil
	}
}

type CreateApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateApplicationInvoker) Invoke() (*model.CreateApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateApplicationResponse), nil
	}
}

type DeleteApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteApplicationInvoker) Invoke() (*model.DeleteApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteApplicationResponse), nil
	}
}

type ListApplicationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApplicationsInvoker) Invoke() (*model.ListApplicationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApplicationsResponse), nil
	}
}

type ShowApplicationInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApplicationInvoker) Invoke() (*model.ShowApplicationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApplicationResponse), nil
	}
}

type CreateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateCertificateInvoker) Invoke() (*model.CreateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateCertificateResponse), nil
	}
}

type DeleteCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteCertificateInvoker) Invoke() (*model.DeleteCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteCertificateResponse), nil
	}
}

type ListCertificatesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListCertificatesInvoker) Invoke() (*model.ListCertificatesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListCertificatesResponse), nil
	}
}

type UpdateCertificateInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateCertificateInvoker) Invoke() (*model.UpdateCertificateResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateCertificateResponse), nil
	}
}

type CreateComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComponentInvoker) Invoke() (*model.CreateComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComponentResponse), nil
	}
}

type DeleteComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComponentInvoker) Invoke() (*model.DeleteComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComponentResponse), nil
	}
}

type ExecuteActionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ExecuteActionInvoker) Invoke() (*model.ExecuteActionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteActionResponse), nil
	}
}

type ListComponentEventsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentEventsInvoker) Invoke() (*model.ListComponentEventsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentEventsResponse), nil
	}
}

type ListComponentInstancesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentInstancesInvoker) Invoke() (*model.ListComponentInstancesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentInstancesResponse), nil
	}
}

type ListComponentSnapshotsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentSnapshotsInvoker) Invoke() (*model.ListComponentSnapshotsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentSnapshotsResponse), nil
	}
}

type ListComponentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentsInvoker) Invoke() (*model.ListComponentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentsResponse), nil
	}
}

type ShowComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowComponentInvoker) Invoke() (*model.ShowComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowComponentResponse), nil
	}
}

type UpdateComponentInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateComponentInvoker) Invoke() (*model.UpdateComponentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateComponentResponse), nil
	}
}

type CreateComponentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateComponentConfigurationInvoker) Invoke() (*model.CreateComponentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateComponentConfigurationResponse), nil
	}
}

type DeleteComponentConfigurationInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteComponentConfigurationInvoker) Invoke() (*model.DeleteComponentConfigurationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteComponentConfigurationResponse), nil
	}
}

type ListComponentConfigurationsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListComponentConfigurationsInvoker) Invoke() (*model.ListComponentConfigurationsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListComponentConfigurationsResponse), nil
	}
}

type CreateDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateDomainInvoker) Invoke() (*model.CreateDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateDomainResponse), nil
	}
}

type DeleteDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteDomainInvoker) Invoke() (*model.DeleteDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteDomainResponse), nil
	}
}

type ListDomainsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListDomainsInvoker) Invoke() (*model.ListDomainsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListDomainsResponse), nil
	}
}

type CreateEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnvironmentInvoker) Invoke() (*model.CreateEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnvironmentResponse), nil
	}
}

type DeleteEnvironmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnvironmentInvoker) Invoke() (*model.DeleteEnvironmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnvironmentResponse), nil
	}
}

type ListEnvironmentsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnvironmentsInvoker) Invoke() (*model.ListEnvironmentsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnvironmentsResponse), nil
	}
}

type RetryJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *RetryJobInvoker) Invoke() (*model.RetryJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetryJobResponse), nil
	}
}

type ShowJobInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowJobInvoker) Invoke() (*model.ShowJobResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowJobResponse), nil
	}
}

type CreateTimerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateTimerRuleInvoker) Invoke() (*model.CreateTimerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateTimerRuleResponse), nil
	}
}

type DeleteTimerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteTimerRuleInvoker) Invoke() (*model.DeleteTimerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteTimerRuleResponse), nil
	}
}

type ListTimerRulesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListTimerRulesInvoker) Invoke() (*model.ListTimerRulesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListTimerRulesResponse), nil
	}
}

type ShowExecutionResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowExecutionResultInvoker) Invoke() (*model.ShowExecutionResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowExecutionResultResponse), nil
	}
}

type UpdateTimerRuleInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateTimerRuleInvoker) Invoke() (*model.UpdateTimerRuleResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateTimerRuleResponse), nil
	}
}

type CreateVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateVolumeInvoker) Invoke() (*model.CreateVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateVolumeResponse), nil
	}
}

type DeleteVolumeInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteVolumeInvoker) Invoke() (*model.DeleteVolumeResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteVolumeResponse), nil
	}
}

type ListVolumesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListVolumesInvoker) Invoke() (*model.ListVolumesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListVolumesResponse), nil
	}
}
