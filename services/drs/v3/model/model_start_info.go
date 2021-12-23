package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动任务请求体。
type StartInfo struct {
	// 任务id。

	JobId string `json:"job_id"`
	// 任务启动时间，时间戳格式精确到秒，例如：1614078283，取值为空代表立即启动。

	StartTime *string `json:"start_time,omitempty"`
}

func (o StartInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInfo struct{}"
	}

	return strings.Join([]string{"StartInfo", string(data)}, " ")
}
