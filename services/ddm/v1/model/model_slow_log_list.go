package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowLogList struct {

	// 执行慢sql的DDM账号名称。
	Users *string `json:"users,omitempty" xml:"users"`

	// 慢sql所属逻辑库的名称。
	Database *string `json:"database,omitempty" xml:"database"`

	// 慢sql执行语法。
	QuerySample *string `json:"querySample,omitempty" xml:"querySample"`

	// DDM慢sql开始执行时间。
	LogTime *string `json:"logTime,omitempty" xml:"logTime"`

	// 慢sql的执行时长，精确到毫秒。
	Time *string `json:"time,omitempty" xml:"time"`

	// 逻辑库物理分片名称。
	Shards *string `json:"shards,omitempty" xml:"shards"`

	// 慢sql影响行数。
	RowsExamined *string `json:"rowsExamined,omitempty" xml:"rowsExamined"`
}

func (o SlowLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogList struct{}"
	}

	return strings.Join([]string{"SlowLogList", string(data)}, " ")
}
