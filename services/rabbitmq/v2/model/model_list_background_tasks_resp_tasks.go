package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListBackgroundTasksRespTasks struct {

	// 任务ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 任务名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 用户名。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 任务参数。
	Params *string `json:"params,omitempty" xml:"params"`

	// 任务状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 启动时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 结束时间。
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at"`
}

func (o ListBackgroundTasksRespTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTasksRespTasks struct{}"
	}

	return strings.Join([]string{"ListBackgroundTasksRespTasks", string(data)}, " ")
}
