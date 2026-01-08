package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptByDesktopTagRequest Request Object
type ExecuteScriptByDesktopTagRequest struct {
	Body *ExecuteScriptByDesktopTagReq `json:"body,omitempty"`
}

func (o ExecuteScriptByDesktopTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptByDesktopTagRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScriptByDesktopTagRequest", string(data)}, " ")
}
