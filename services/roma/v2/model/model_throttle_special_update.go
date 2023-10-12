package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleSpecialUpdate struct {

	// 流控时间内特殊对象能够访问API的最大次数限制
	CallLimits int32 `json:"call_limits"`
}

func (o ThrottleSpecialUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleSpecialUpdate struct{}"
	}

	return strings.Join([]string{"ThrottleSpecialUpdate", string(data)}, " ")
}
