package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeadLockDetailResponse Response Object
type ListDeadLockDetailResponse struct {

	// 死锁明细列表
	DetailList *[]DeadLockDetail `json:"detail_list,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDeadLockDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeadLockDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDeadLockDetailResponse", string(data)}, " ")
}
