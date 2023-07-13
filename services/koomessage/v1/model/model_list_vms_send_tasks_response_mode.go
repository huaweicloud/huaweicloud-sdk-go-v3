package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsSendTasksResponseMode 查询发送任务列表响应体。
type ListVmsSendTasksResponseMode struct {

	// 智能信息基础版任务查询列表。
	AimBasicSendTasks *[]VmsSendTask `json:"aim_basic_send_tasks,omitempty"`

	PageInfo *Page `json:"page_info,omitempty"`
}

func (o ListVmsSendTasksResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsSendTasksResponseMode struct{}"
	}

	return strings.Join([]string{"ListVmsSendTasksResponseMode", string(data)}, " ")
}
