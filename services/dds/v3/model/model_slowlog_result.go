package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowlogResult struct {
	// 节点名称。

	NodeName string `json:"node_name"`
	// 执行语法。

	QuerySample string `json:"query_sample"`
	// 语句类型。

	Type string `json:"type"`
	// 执行时间。

	Time string `json:"time"`
	// 等待锁时间。

	LockTime string `json:"lock_time"`
	// 角色所在数据库名称。

	RowsSent string `json:"rows_sent"`
	// 扫描的行数量。

	RowsExamined string `json:"rows_examined"`
	// 所属数据库。

	Database string `json:"database"`
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
