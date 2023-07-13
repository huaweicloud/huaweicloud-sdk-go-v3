package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentExecutionResult struct {

	// 组件名称。
	ComponentName *string `json:"component_name,omitempty"`

	// 组件ID。
	ComponentId *string `json:"component_id,omitempty"`

	// 组件执行启停的结果：failed/success。
	Status *string `json:"status,omitempty"`

	// 组件执行启停的错误信息，如果执行结果为成功，则为空。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ComponentExecutionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentExecutionResult struct{}"
	}

	return strings.Join([]string{"ComponentExecutionResult", string(data)}, " ")
}
