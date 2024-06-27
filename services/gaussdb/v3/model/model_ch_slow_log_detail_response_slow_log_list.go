package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChSlowLogDetailResponseSlowLogList struct {

	// ClickHouse实例节点ID。
	NodeId string `json:"node_id"`

	// 数据库语句执行时间。
	Time string `json:"time"`

	// 数据库语句等待锁时间。
	LockTime string `json:"lock_time"`

	// 数据库语句执行结果行数。
	RowsSent int32 `json:"rows_sent"`

	// 数据库语句扫描行数。
	RowsExamined int32 `json:"rows_examined"`

	// 所属数据库名。
	Database string `json:"database"`

	// 执行语句账号。
	Users string `json:"users"`

	// 数据库执行语句。
	QuerySample string `json:"query_sample"`

	// 数据库语句类型。
	Type string `json:"type"`

	// IP地址。
	ClientIp string `json:"client_ip"`

	// 数据库语句发生时间。
	StartTime string `json:"start_time"`

	// 日志单行序列号，第一次查询时不需要此参数，后续分页查询时需要使用，可从上次查询的返回信息中获取。
	LineNum string `json:"line_num"`

	// 慢日志数量。
	Count string `json:"count"`
}

func (o ChSlowLogDetailResponseSlowLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChSlowLogDetailResponseSlowLogList struct{}"
	}

	return strings.Join([]string{"ChSlowLogDetailResponseSlowLogList", string(data)}, " ")
}
