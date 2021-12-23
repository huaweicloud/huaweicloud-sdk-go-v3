package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckTaskJobsItemsBody struct {
	// 任务标识。

	JobId *string `json:"job_id,omitempty"`
	// 任务状态如下：  - created 已创建  - running 正在处理  - finish 已完成  - failed 处理失败

	Status *string `json:"status,omitempty"`
	// 任务创建的时间。例如：2018-01-02T15:03:04Z

	CreateTime *string `json:"create_time,omitempty"`
	// 任务最近更新的时间。例如：2018-01-02T15:03:04Z

	UpdateTime *string `json:"update_time,omitempty"`
}

func (o CheckTaskJobsItemsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTaskJobsItemsBody struct{}"
	}

	return strings.Join([]string{"CheckTaskJobsItemsBody", string(data)}, " ")
}
