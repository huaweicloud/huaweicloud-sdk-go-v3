package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDesktopPoolScriptRequest Request Object
type ExecuteDesktopPoolScriptRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *ExecuteDesktopPoolScriptsReq `json:"body,omitempty"`
}

func (o ExecuteDesktopPoolScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDesktopPoolScriptRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDesktopPoolScriptRequest", string(data)}, " ")
}
