package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeExecution 节点执行记录
type NodeExecution struct {

	// 流程节点执行状态  最小长度：1  最大长度：32  枚举值：  success  fail  running  timeout  cancel
	Status *string `json:"status,omitempty"`

	// 函数执行时的入参
	Input *interface{} `json:"input,omitempty"`

	// 函数执行结果
	Output *interface{} `json:"output,omitempty"`

	// 节点启动时间，UTC毫秒时间戳格式  最小值：0  最大值：99999999999999999
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 节点结束时间，UTC毫秒时间戳格式  最小值：0  最大值：99999999999999999
	EndTime *int64 `json:"end_time,omitempty"`

	// 节点错误信息，仅在节点出错时非空
	ErrorMessage *string `json:"error_message,omitempty"`

	// 流程节点请求ID
	RequestId *string `json:"request_id,omitempty"`
}

func (o NodeExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExecution struct{}"
	}

	return strings.Join([]string{"NodeExecution", string(data)}, " ")
}
