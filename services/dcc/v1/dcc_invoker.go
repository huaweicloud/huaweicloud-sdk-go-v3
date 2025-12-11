package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcc/v1/model"
)

type ShowResourceClustersInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowResourceClustersInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *ShowResourceClustersInvoker) Invoke() (*model.ShowResourceClustersResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowResourceClustersResponse), nil
	}
}
