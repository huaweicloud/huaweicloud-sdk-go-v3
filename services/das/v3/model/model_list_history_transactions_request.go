package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryTransactionsRequest Request Object
type ListHistoryTransactionsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 采集开始时间（Unix timestamp，毫秒）
	StartAt int64 `json:"start_at"`

	// 采集结束时间（Unix timestamp，毫秒）
	EndAt int64 `json:"end_at"`

	// 页数
	PageNum *int32 `json:"page_num,omitempty"`

	// 页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段
	Order *string `json:"order,omitempty"`

	// 升序|降序
	OrderBy *string `json:"order_by,omitempty"`

	// 持续时间下限
	LastSecMin *int64 `json:"last_sec_min,omitempty"`

	// 持续时间上限
	LastSecMax *int64 `json:"last_sec_max,omitempty"`
}

func (o ListHistoryTransactionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryTransactionsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryTransactionsRequest", string(data)}, " ")
}
