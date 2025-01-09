package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScriptRequest Request Object
type UpdateScriptRequest struct {

	// 脚本ID。
	ScriptId string `json:"script_id"`

	Body *UpdateScriptReq `json:"body,omitempty"`
}

func (o UpdateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptRequest struct{}"
	}

	return strings.Join([]string{"UpdateScriptRequest", string(data)}, " ")
}
