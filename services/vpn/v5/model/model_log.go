package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Log struct {

	// 时间戳
	Time *int64 `json:"time,omitempty"`

	// 日志信息
	RawMessage *string `json:"raw_message,omitempty"`
}

func (o Log) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Log struct{}"
	}

	return strings.Join([]string{"Log", string(data)}, " ")
}
