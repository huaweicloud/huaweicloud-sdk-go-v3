package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudartifact/v2/model"
)

type ShowReleaseProjectFilesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowReleaseProjectFilesInvoker) Invoke() (*model.ShowReleaseProjectFilesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowReleaseProjectFilesResponse), nil
	}
}
