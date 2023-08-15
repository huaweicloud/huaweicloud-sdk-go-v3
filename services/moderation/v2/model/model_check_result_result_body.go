package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckResultResultBody
type CheckResultResultBody struct {

	// 任务标识。
	JobId *string `json:"job_id,omitempty"`

	// 任务状态。  created：已创建  running：正在处理  finish：已完成  failed：处理失败
	Status *string `json:"status,omitempty"`

	// 任务创建的时间。例如：2018-01-02T15:03:04Z
	CreateTime *string `json:"create_time,omitempty"`

	// 任务最近更新的时间。例如：2018-01-02T15:03:04Z
	UpdateTime *string `json:"update_time,omitempty"`

	// 图片内容检测的结果列表。
	Items *[]CheckResultItemsBody `json:"items,omitempty"`
}

func (o CheckResultResultBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResultResultBody struct{}"
	}

	return strings.Join([]string{"CheckResultResultBody", string(data)}, " ")
}
