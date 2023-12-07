package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/codeartsartifact/v2/model"
)

type ShowProjectReleaseFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowProjectReleaseFilesInvoker) Invoke() (*model.ShowProjectReleaseFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowProjectReleaseFilesResponse), nil
	}
}

type ShowReleaseProjectFilesInvoker struct {
	*invoker.BaseInvoker
}

// Deprecated: This function is deprecated and will be removed in the future versions.
func (i *ShowReleaseProjectFilesInvoker) Invoke() (*model.ShowReleaseProjectFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReleaseProjectFilesResponse), nil
	}
}
