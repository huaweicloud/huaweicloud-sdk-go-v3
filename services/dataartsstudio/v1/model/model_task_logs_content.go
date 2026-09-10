package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskLogsContent struct {

	// 请求是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息。
	Message *string `json:"message,omitempty"`

	// 日志路径前缀。
	Prefix *string `json:"prefix,omitempty"`

	// 日志时间。
	Time *string `json:"time,omitempty"`

	// 文件列表。
	Files *[]TaskLogFile `json:"files,omitempty"`
}

func (o TaskLogsContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskLogsContent struct{}"
	}

	return strings.Join([]string{"TaskLogsContent", string(data)}, " ")
}
