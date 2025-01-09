package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScriptRequest Request Object
type DeleteScriptRequest struct {

	// 脚本ID。
	ScriptId string `json:"script_id"`
}

func (o DeleteScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScriptRequest struct{}"
	}

	return strings.Join([]string{"DeleteScriptRequest", string(data)}, " ")
}
