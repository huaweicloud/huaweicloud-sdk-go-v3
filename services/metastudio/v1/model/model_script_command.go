package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScriptCommand 数字人话术命令信息。
type ScriptCommand struct {

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 命令ID。
	CommandId *string `json:"command_id,omitempty"`

	// 命令时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CommandTime *string `json:"command_time,omitempty"`

	// 直播剧本列表。
	SceneScripts *[]LivePlayingScriptInfo `json:"scene_scripts,omitempty"`
}

func (o ScriptCommand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptCommand struct{}"
	}

	return strings.Join([]string{"ScriptCommand", string(data)}, " ")
}
