package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateScriptRequest struct {
	Body *ScriptInfo `json:"body,omitempty"`
}

func (o CreateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateScriptRequest", string(data)}, " ")
}
