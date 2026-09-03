package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullDeadLocksResponse Response Object
type ListFullDeadLocksResponse struct {

	// 全量死锁列表
	FullDeadLockList *[]FullDeadLock `json:"full_dead_lock_list,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFullDeadLocksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullDeadLocksResponse struct{}"
	}

	return strings.Join([]string{"ListFullDeadLocksResponse", string(data)}, " ")
}
