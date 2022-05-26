package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type ListAlarmHistoriesInvoker struct {
	*invoker.BaseInvoker
}

func (i *ListAlarmHistoriesInvoker) Invoke() (*model.ListAlarmHistoriesResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.ListAlarmHistoriesResponse), nil
	}
}
