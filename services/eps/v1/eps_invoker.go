package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eps/v1/model"
)

type CreateEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
}

func (i *CreateEnterpriseProjectInvoker) Invoke() (*model.CreateEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.CreateEnterpriseProjectResponse), nil
	}
}

type DisableEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
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

func (i *ListEnterpriseProjectInvoker) Invoke() (*model.ListEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListEnterpriseProjectResponse), nil
	}
}

type MigrateResourceInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowApiVersionInvoker) Invoke() (*model.ShowApiVersionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowApiVersionResponse), nil
	}
}

type ShowEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
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

func (i *ShowEnterpriseProjectQuotaInvoker) Invoke() (*model.ShowEnterpriseProjectQuotaResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowEnterpriseProjectQuotaResponse), nil
	}
}

type ShowResourceBindEnterpriseProjectInvoker struct {
	*invoker.BaseInvoker
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

func (i *UpdateEnterpriseProjectInvoker) Invoke() (*model.UpdateEnterpriseProjectResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.UpdateEnterpriseProjectResponse), nil
	}
}
