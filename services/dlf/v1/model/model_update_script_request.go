package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	ScriptName string `json:"script_name"`

	Body *ScriptInfo `json:"body,omitempty"`
}

func (o UpdateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateScriptRequest", string(data)}, " ")
}
