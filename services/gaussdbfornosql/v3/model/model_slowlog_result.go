package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowlogResult struct {

	// 执行时间。
	Time string `json:"time" xml:"time"`

	// 所属数据库。
	Database string `json:"database" xml:"database"`

	// 执行语法。
	QuerySample string `json:"query_sample" xml:"query_sample"`

	// 语句类型。
	Type string `json:"type" xml:"type"`

	// 发生时间，UTC时间。
	StartTime string `json:"start_time" xml:"start_time"`
}

func (o SlowlogResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogResult struct{}"
	}

	return strings.Join([]string{"SlowlogResult", string(data)}, " ")
}
