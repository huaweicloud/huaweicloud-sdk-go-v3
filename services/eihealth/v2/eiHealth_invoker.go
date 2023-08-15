package v2

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/eihealth/v2/model"
)

type ShowAdmetWithCustomPropsInvoker struct {
	*invoker.BaseInvoker
}

func (i *ShowAdmetWithCustomPropsInvoker) Invoke() (*model.ShowAdmetWithCustomPropsResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ShowAdmetWithCustomPropsResponse), nil
	}
}
