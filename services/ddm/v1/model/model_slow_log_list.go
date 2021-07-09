package model

import (
	"encoding/json"

	"strings"
)

type SlowLogList struct {
	// 执行慢sql的DDM账号名称。

	Users *string `json:"users,omitempty"`
	// 慢sql所属逻辑库的名称。

	Database *string `json:"database,omitempty"`
	// 慢sql执行语法。

	QuerySample *string `json:"querySample,omitempty"`
	// DDM慢sql开始执行时间。

	LogTime *string `json:"logTime,omitempty"`
	// 慢sql的执行时长，精确到毫秒。

	Time *string `json:"time,omitempty"`
	// 逻辑库物理分片名称。

	Shards *string `json:"shards,omitempty"`
	// 慢sql影响行数。

	RowsExamined *string `json:"rowsExamined,omitempty"`
}

func (o SlowLogList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SlowLogList struct{}"
	}

	return strings.Join([]string{"SlowLogList", string(data)}, " ")
}
