package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 错误日志信息。
type ErrorLog struct {
	// 日期时间UTC时间。

	Time string `json:"time"`
	// 日志级别。

	Level string `json:"level"`
	// 错误日志内容。

	Content string `json:"content"`
}

func (o ErrorLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorLog struct{}"
	}

	return strings.Join([]string{"ErrorLog", string(data)}, " ")
}
