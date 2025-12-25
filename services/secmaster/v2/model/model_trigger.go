package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Trigger struct {

	// 累计次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`

	// UUID
	JobId *string `json:"job_id,omitempty"`

	Mode *Mode `json:"mode,omitempty"`

	Operator *Operator `json:"operator,omitempty"`

	Severity *Serverity `json:"severity,omitempty"`
}

func (o Trigger) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Trigger struct{}"
	}

	return strings.Join([]string{"Trigger", string(data)}, " ")
}
