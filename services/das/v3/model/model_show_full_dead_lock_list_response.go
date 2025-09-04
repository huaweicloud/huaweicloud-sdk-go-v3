package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFullDeadLockListResponse Response Object
type ShowFullDeadLockListResponse struct {

	// 死锁数量
	Total *int32 `json:"total,omitempty"`

	// 死锁具体信息
	FullDeadlockList *[]FullDeadLockListRespFullDeadlockList `json:"full_deadlock_list,omitempty"`
	HttpStatusCode   int                                     `json:"-"`
}

func (o ShowFullDeadLockListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullDeadLockListResponse struct{}"
	}

	return strings.Join([]string{"ShowFullDeadLockListResponse", string(data)}, " ")
}
