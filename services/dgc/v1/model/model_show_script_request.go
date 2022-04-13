package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowScriptRequest struct {
	ScriptName string `json:"script_name"`
}

func (o ShowScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptRequest", string(data)}, " ")
}
