package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowlogResult struct {
	// 执行时间。

	Time string `json:"time"`
	// 所属数据库。

	Database string `json:"database"`
	// 执行语法。

	QuerySample string `json:"query_sample"`
	// 语句类型。

	Type string `json:"type"`
	// 发生时间，UTC时间。

	StartTime string `json:"start_time"`
}

func (o SlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogResult struct{}"
	}

	return strings.Join([]string{"SlowlogResult", string(data)}, " ")
}
