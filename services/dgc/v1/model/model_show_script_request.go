package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScriptRequest Request Object
type ShowScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	ScriptName string `json:"script_name"`
}

func (o ShowScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptRequest", string(data)}, " ")
}
