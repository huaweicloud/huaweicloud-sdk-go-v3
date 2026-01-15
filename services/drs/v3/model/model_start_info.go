package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInfo 启动任务请求体。
type StartInfo struct {

	// 任务id。
	JobId string `json:"job_id"`

	// 是否支持只初始化任务。仅支持白名单用户使用，需要提交工单申请才能使用。
	IsOnlyInitTask *bool `json:"is_only_init_task,omitempty"`

	// 是否在任务结束时自动创建对比任务，不填默认设置为true。
	IsAutoCreateCompare *bool `json:"is_auto_create_compare,omitempty"`

	// 任务启动时间，时间戳格式精确到毫秒，例如：1679966489593，取值为空代表立即启动。
	StartTime *string `json:"start_time,omitempty"`
}

func (o StartInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInfo struct{}"
	}

	return strings.Join([]string{"StartInfo", string(data)}, " ")
}
