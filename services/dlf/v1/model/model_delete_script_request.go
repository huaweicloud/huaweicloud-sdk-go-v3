package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScriptRequest Request Object
type DeleteScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	ScriptName string `json:"script_name"`
}

func (o DeleteScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScriptRequest struct{}"
	}

	return strings.Join([]string{"DeleteScriptRequest", string(data)}, " ")
}
