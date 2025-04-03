package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenterportalapi/v1/model"
)

type ListAccountsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountsInvoker) Invoke() (*model.ListAccountsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountsResponse), nil
	}
}

type ListAccountAgenciesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAccountAgenciesInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ListAccountAgenciesInvoker) Invoke() (*model.ListAccountAgenciesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAccountAgenciesResponse), nil
	}
}

type GetAgencyCredentialsInvoker struct {
	*invoker.BaseInvoker
}

func (i *GetAgencyCredentialsInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *GetAgencyCredentialsInvoker) Invoke() (*model.GetAgencyCredentialsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.GetAgencyCredentialsResponse), nil
	}
}

type LogoutInvoker struct {
	*invoker.BaseInvoker
}

func (i *LogoutInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *LogoutInvoker) Invoke() (*model.LogoutResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogoutResponse), nil
	}
}
