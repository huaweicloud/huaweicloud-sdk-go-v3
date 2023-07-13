package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigBaseLogInfoCreate 日志接入基础日志信息。
type AccessConfigBaseLogInfoCreate struct {

	// 日志组ID
	LogGroupId string `json:"log_group_id"`

	// 日志流ID
	LogStreamId string `json:"log_stream_id"`
}

func (o AccessConfigBaseLogInfoCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigBaseLogInfoCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigBaseLogInfoCreate", string(data)}, " ")
}
