package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScriptRequest Request Object
type CancelScriptRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	ScriptName string `json:"script_name"`

	InstanceId string `json:"instance_id"`
}

func (o CancelScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScriptRequest struct{}"
	}

	return strings.Join([]string{"CancelScriptRequest", string(data)}, " ")
}
