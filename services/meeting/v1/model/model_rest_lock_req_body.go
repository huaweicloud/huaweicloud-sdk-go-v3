package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 锁定会议消息体。
type RestLockReqBody struct {
	// - 0: 解锁。 - 1: 锁定。

	IsLock int32 `json:"isLock"`
}

func (o RestLockReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestLockReqBody struct{}"
	}

	return strings.Join([]string{"RestLockReqBody", string(data)}, " ")
}
