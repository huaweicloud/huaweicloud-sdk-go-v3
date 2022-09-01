package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelScriptRequest struct {
	ScriptName string `json:"script_name" xml:"script_name"`

	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o CancelScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScriptRequest struct{}"
	}

	return strings.Join([]string{"CancelScriptRequest", string(data)}, " ")
}
