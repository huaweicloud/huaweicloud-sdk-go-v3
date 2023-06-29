package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestLockReqBody 锁定会议请求。
type RestLockReqBody struct {

	// - 0: 解锁 - 1: 锁定
	IsLock int32 `json:"isLock"`
}

func (o RestLockReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestLockReqBody struct{}"
	}

	return strings.Join([]string{"RestLockReqBody", string(data)}, " ")
}
