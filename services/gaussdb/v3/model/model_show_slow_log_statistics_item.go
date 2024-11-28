package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSlowLogStatisticsItem struct {

	// IP地址。
	ClientIp *string `json:"client_ip,omitempty"`

	// 执行次数。
	Count *string `json:"count,omitempty"`

	// 所属数据库。
	Database *string `json:"database,omitempty"`

	// 平均等待锁时间。
	LockTime *string `json:"lock_time,omitempty"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 执行语法。
	QuerySample *string `json:"query_sample,omitempty"`

	// 平均扫描的行数量。
	RowsExamined *int32 `json:"rows_examined,omitempty"`

	// 平均结果行统计数量。
	RowsSent *int32 `json:"rows_sent,omitempty"`

	// 平均执行时间。
	Time *string `json:"time,omitempty"`

	// 语句类型。
	Type *string `json:"type,omitempty"`

	// 账号。
	Users *string `json:"users,omitempty"`
}

func (o ShowSlowLogStatisticsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsItem struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsItem", string(data)}, " ")
}
