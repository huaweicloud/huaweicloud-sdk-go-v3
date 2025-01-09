package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowLogs struct {

	// 执行慢sql的DDM账号名称。
	Users *string `json:"users,omitempty"`

	// 慢sql所属逻辑库的名称。
	Database *string `json:"database,omitempty"`

	// 慢sql执行语法。
	QuerySample *string `json:"query_sample,omitempty"`

	// DDM慢sql开始执行时间。
	LogTime *string `json:"log_time,omitempty"`

	// 慢sql的执行时长，精确到毫秒。
	Time *string `json:"time,omitempty"`

	// 逻辑库物理分片名称。
	Shards *string `json:"shards,omitempty"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 慢sql影响行数。
	RowsExamined *string `json:"rows_examined,omitempty"`

	// 客户端ip，该IP地址可能涉及个人数据，建议用户依据实际IP地址的敏感性做查询后脱敏处理。
	Host *string `json:"host,omitempty"`
}

func (o SlowLogs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogs struct{}"
	}

	return strings.Join([]string{"SlowLogs", string(data)}, " ")
}
