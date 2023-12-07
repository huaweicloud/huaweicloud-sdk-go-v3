package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchWindow struct {

	// 批量推送条数[1,10000]
	Count *int32 `json:"count,omitempty"`

	// 重试次数
	Time *int32 `json:"time,omitempty"`

	// 批量推送间隔[0,15]，单位秒
	Interval *int32 `json:"interval,omitempty"`
}

func (o BatchWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchWindow struct{}"
	}

	return strings.Join([]string{"BatchWindow", string(data)}, " ")
}
