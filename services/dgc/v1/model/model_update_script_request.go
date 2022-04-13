package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateScriptRequest struct {
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
