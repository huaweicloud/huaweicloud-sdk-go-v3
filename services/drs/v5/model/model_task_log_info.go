package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志查询结果信息体。
type TaskLogInfo struct {

	// 日志级别。
	Level *string `json:"level,omitempty"`

	// 日志信息。
	Message *string `json:"message,omitempty"`

	// 日志时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o TaskLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskLogInfo struct{}"
	}

	return strings.Join([]string{"TaskLogInfo", string(data)}, " ")
}
