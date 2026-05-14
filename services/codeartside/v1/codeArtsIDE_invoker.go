package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartside/v1/model"
)

type ShowLatestUpgradableReleaseInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowLatestUpgradableReleaseInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowLatestUpgradableReleaseInvoker) Invoke() (*model.ShowLatestUpgradableReleaseResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowLatestUpgradableReleaseResponse), nil
	}
}

type ValidateWhitelistUserInvoker struct {
	*invoker.BaseInvoker
}

func (i *ValidateWhitelistUserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ValidateWhitelistUserInvoker) Invoke() (*model.ValidateWhitelistUserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ValidateWhitelistUserResponse), nil
	}
}
