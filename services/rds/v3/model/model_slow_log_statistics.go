package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 慢日志信息。
type SlowLogStatistics struct {

	// 执行次数。
	Count string `json:"count" xml:"count"`

	// 平均执行时间。
	Time string `json:"time" xml:"time"`

	// 平均等待锁时间。
	LockTime string `json:"lockTime" xml:"lockTime"`

	// 平均结果行数量。
	RowsSent int64 `json:"rowsSent" xml:"rowsSent"`

	// 平均扫描的行数量。
	RowsExamined int64 `json:"rowsExamined" xml:"rowsExamined"`

	// 所属数据库。
	Database string `json:"database" xml:"database"`

	// 帐号。
	Users string `json:"users" xml:"users"`

	// 执行语法。
	QuerySample string `json:"querySample" xml:"querySample"`

	// 语句类型。
	Type string `json:"type" xml:"type"`

	// IP地址。
	ClientIP string `json:"clientIP" xml:"clientIP"`
}

func (o SlowLogStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogStatistics struct{}"
	}

	return strings.Join([]string{"SlowLogStatistics", string(data)}, " ")
}
