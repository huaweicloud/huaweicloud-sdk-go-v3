package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskUrls SSH连接地址信息。
type TaskUrls struct {

	// 训练作业的任务ID。
	Task *string `json:"task,omitempty"`

	// 训练作业SSH连接地址。
	Url *string `json:"url,omitempty"`
}

func (o TaskUrls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskUrls struct{}"
	}

	return strings.Join([]string{"TaskUrls", string(data)}, " ")
}
