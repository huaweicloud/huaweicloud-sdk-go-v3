package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParseDeadLockRequestBody struct {

	// 死锁唯一标识
	DeadLockId string `json:"dead_lock_id"`
}

func (o ParseDeadLockRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseDeadLockRequestBody struct{}"
	}

	return strings.Join([]string{"ParseDeadLockRequestBody", string(data)}, " ")
}
