package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowScriptRequest struct {
	ScriptName string `json:"script_name" xml:"script_name"`
}

func (o ShowScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptRequest", string(data)}, " ")
}
