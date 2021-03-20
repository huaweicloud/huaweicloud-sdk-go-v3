package model

import (
	"encoding/json"

	"strings"
)

type ExecuteScriptReq struct {
	// 脚本的执行参数

	Params *string `json:"params,omitempty"`
}

func (o ExecuteScriptReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScriptReq struct{}"
	}

	return strings.Join([]string{"ExecuteScriptReq", string(data)}, " ")
}
