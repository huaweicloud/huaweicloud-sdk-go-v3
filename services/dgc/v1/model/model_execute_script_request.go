package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteScriptRequest struct {
	ScriptName string `json:"script_name"`

	Body *ExecuteScriptReq `json:"body,omitempty"`
}

func (o ExecuteScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScriptRequest", string(data)}, " ")
}
