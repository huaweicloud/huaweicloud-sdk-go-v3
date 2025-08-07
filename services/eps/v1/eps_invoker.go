package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/model"
)

type CreateEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *CreateEnterpriseProjectInvoker) Invoke() (*model.CreateEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnterpriseProjectResponse), nil
	}
}

type DeleteEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DeleteEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DeleteEnterpriseProjectInvoker) Invoke() (*model.DeleteEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeleteEnterpriseProjectResponse), nil
	}
}

type DisableEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *DisableEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *DisableEnterpriseProjectInvoker) Invoke() (*model.DisableEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.DisableEnterpriseProjectResponse), nil
	}
}

type EnableEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *EnableEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *EnableEnterpriseProjectInvoker) Invoke() (*model.EnableEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnableEnterpriseProjectResponse), nil
	}
}

type ListApiVersionsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListApiVersionsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListApiVersionsInvoker) Invoke() (*model.ListApiVersionsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListApiVersionsResponse), nil
	}
}

type ListEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListEnterpriseProjectInvoker) Invoke() (*model.ListEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseProjectResponse), nil
	}
}

type ListMigrationRecordInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListMigrationRecordInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListMigrationRecordInvoker) Invoke() (*model.ListMigrationRecordResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListMigrationRecordResponse), nil
	}
}

type ListProvidersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListProvidersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListProvidersInvoker) Invoke() (*model.ListProvidersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListProvidersResponse), nil
	}
}

type ListResourceMappingInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListResourceMappingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListResourceMappingInvoker) Invoke() (*model.ListResourceMappingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListResourceMappingResponse), nil
	}
}

type MigrateResourceInvoker struct {
	*invoker.BaseInvoker
}

func (i *MigrateResourceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *MigrateResourceInvoker) Invoke() (*model.MigrateResourceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrateResourceResponse), nil
	}
}

type ShowApiVersionInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowApiVersionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}

type ShowAssociatedResourcesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAssociatedResourcesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowAssociatedResourcesInvoker) Invoke() (*model.ShowAssociatedResourcesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAssociatedResourcesResponse), nil
	}
}

type ShowEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEnterpriseProjectInvoker) Invoke() (*model.ShowEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseProjectResponse), nil
	}
}

type ShowEnterpriseProjectQuotaInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEnterpriseProjectQuotaInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEnterpriseProjectQuotaInvoker) Invoke() (*model.ShowEnterpriseProjectQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseProjectQuotaResponse), nil
	}
}

type ShowEpConfigsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowEpConfigsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowEpConfigsInvoker) Invoke() (*model.ShowEpConfigsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEpConfigsResponse), nil
	}
}

type ShowResourceBindEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceBindEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceBindEnterpriseProjectInvoker) Invoke() (*model.ShowResourceBindEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceBindEnterpriseProjectResponse), nil
	}
}

type UpdateEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *UpdateEnterpriseProjectInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *UpdateEnterpriseProjectInvoker) Invoke() (*model.UpdateEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnterpriseProjectResponse), nil
	}
}
