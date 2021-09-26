package model

import (
	"encoding/json"

	"strings"
)

type ConcatInfo struct {
	// 拼接任务输入源地址。

	Inputs *[]ObsObjInfo `json:"inputs,omitempty"`
}

func (o ConcatInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConcatInfo struct{}"
	}

	return strings.Join([]string{"ConcatInfo", string(data)}, " ")
}
