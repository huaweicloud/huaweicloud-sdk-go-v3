package model

import (
	"encoding/json"

	"strings"
)

type OperateLog struct {
	// 操作指令

	Oper *string `json:"oper,omitempty"`
	// 操作时间

	OperateTime *string `json:"operate_time,omitempty"`
}

func (o OperateLog) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OperateLog struct{}"
	}

	return strings.Join([]string{"OperateLog", string(data)}, " ")
}
