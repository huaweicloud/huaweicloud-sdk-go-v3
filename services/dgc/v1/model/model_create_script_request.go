package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptRequest Request Object
type CreateScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ScriptInfo `json:"body,omitempty"`
}

func (o CreateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateScriptRequest", string(data)}, " ")
}
