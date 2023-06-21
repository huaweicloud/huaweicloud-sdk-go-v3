package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建任务请求体。
type CreateVmsTaskResponseMode struct {

	// 智能信息基础版下发结果返回码。
	RetCode *string `json:"ret_code,omitempty"`

	// 智能信息基础版下发任务批次ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息基础版下发描述信息。
	Desc *string `json:"desc,omitempty"`
}

func (o CreateVmsTaskResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVmsTaskResponseMode struct{}"
	}

	return strings.Join([]string{"CreateVmsTaskResponseMode", string(data)}, " ")
}
