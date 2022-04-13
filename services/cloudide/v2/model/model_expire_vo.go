package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpireVo struct {
	// 过期时间。UNIX时间戳，单位毫秒。eg:1635905480465

	Expiration *int64 `json:"expiration,omitempty"`
	// CloudIDE实例id

	InstanceId *string `json:"instance_id,omitempty"`
	// CloudIDE实例自动休眠时长，单位‘分钟’

	Interval *int64 `json:"interval,omitempty"`
}

func (o ExpireVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpireVo struct{}"
	}

	return strings.Join([]string{"ExpireVo", string(data)}, " ")
}
