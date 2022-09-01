package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateScriptRequest struct {
	ScriptName string `json:"script_name" xml:"script_name"`

	Body *ScriptInfo `json:"body,omitempty" xml:"body"`
}

func (o UpdateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateScriptRequest", string(data)}, " ")
}
