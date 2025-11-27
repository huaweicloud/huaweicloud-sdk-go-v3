package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControlTask 具体任务信息。
type AccessControlTask struct {

	// 任务id。
	Id *string `json:"id,omitempty"`

	// url列表。
	Urls *[]string `json:"urls,omitempty"`

	// 任务状态，状态类型：processing：处理中；succeed：完成；failed：失败。
	Status *string `json:"status,omitempty"`

	// 任务类型，unban：解禁任务；ban：封禁任务。
	Action *string `json:"action,omitempty"`

	// 创建时间戳（毫秒）。
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o AccessControlTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControlTask struct{}"
	}

	return strings.Join([]string{"AccessControlTask", string(data)}, " ")
}
