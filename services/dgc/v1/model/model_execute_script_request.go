package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

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
