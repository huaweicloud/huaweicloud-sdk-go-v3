package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetConsumeOffsetReq struct {

	// 重置的主题。
	Topic string `json:"topic"`

	// 重置的时间。
	Timestamp float32 `json:"timestamp"`
}

func (o ResetConsumeOffsetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetReq struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetReq", string(data)}, " ")
}
