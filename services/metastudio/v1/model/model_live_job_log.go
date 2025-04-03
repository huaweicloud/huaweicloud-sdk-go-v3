package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveJobLog 直播任务记录。
type LiveJobLog struct {

	// 直播互动记录文件地址
	InteractionRecordsUrl *string `json:"interaction_records_url,omitempty"`

	// 任务配置记录文件地址
	JobConfigRecordsUrl *string `json:"job_config_records_url,omitempty"`

	// 剧本播放记录文件地址
	ScriptsRecordsUrl *string `json:"scripts_records_url,omitempty"`

	// 命令接受记录文件地址
	CommandRevicedRecordsUrl *string `json:"command_reviced_records_url,omitempty"`

	// 命令执行记录文件地址
	CommandExecRecordsUrl *string `json:"command_exec_records_url,omitempty"`
}

func (o LiveJobLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveJobLog struct{}"
	}

	return strings.Join([]string{"LiveJobLog", string(data)}, " ")
}
