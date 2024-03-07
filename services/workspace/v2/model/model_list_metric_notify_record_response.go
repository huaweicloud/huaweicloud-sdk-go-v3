package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricNotifyRecordResponse Response Object
type ListMetricNotifyRecordResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 通知记录
	Items          *[]DesktopMetricNotifyRecord `json:"items,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMetricNotifyRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricNotifyRecordResponse struct{}"
	}

	return strings.Join([]string{"ListMetricNotifyRecordResponse", string(data)}, " ")
}
