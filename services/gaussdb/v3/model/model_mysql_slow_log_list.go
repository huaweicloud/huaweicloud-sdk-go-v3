package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlSlowLogList struct {

	// 节点ID。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 执行次数。
	Count *string `json:"count,omitempty" xml:"count"`

	// 执行时间。
	Time *string `json:"time,omitempty" xml:"time"`

	// 等待锁时间。
	LockTime *string `json:"lock_time,omitempty" xml:"lock_time"`

	// 结果行数量。
	RowsSent *string `json:"rows_sent,omitempty" xml:"rows_sent"`

	// 扫描的行数量。
	RowsExamined *string `json:"rows_examined,omitempty" xml:"rows_examined"`

	// 所属数据库。
	Database *string `json:"database,omitempty" xml:"database"`

	// 账号。
	Users *string `json:"users,omitempty" xml:"users"`

	// 执行语法。
	QuerySample *string `json:"query_sample,omitempty" xml:"query_sample"`

	// 语句类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 发生时间,UTC时间
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// IP地址。
	ClientIp *string `json:"client_ip,omitempty" xml:"client_ip"`
}

func (o MysqlSlowLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlSlowLogList struct{}"
	}

	return strings.Join([]string{"MysqlSlowLogList", string(data)}, " ")
}
