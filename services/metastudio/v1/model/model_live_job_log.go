package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveJobLog 直播任务记录。
type LiveJobLog struct {

	// 直播互动记录文件地址
	InteractionRecordsUrl *string `json:"interaction_records_url,omitempty"`
}

func (o LiveJobLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveJobLog struct{}"
	}

	return strings.Join([]string{"LiveJobLog", string(data)}, " ")
}
