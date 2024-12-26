package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveScriptCommandsResponse Response Object
type ListSmartLiveScriptCommandsResponse struct {

	// 数字人直播任务总数。
	Count *int32 `json:"count,omitempty"`

	// 数字人话术命令列表。
	ScriptsCommands *[]ScriptCommand `json:"scripts_commands,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartLiveScriptCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveScriptCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListSmartLiveScriptCommandsResponse", string(data)}, " ")
}
