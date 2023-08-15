package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskSourceUsageEstimate 计算资源
type TaskSourceUsageEstimate struct {

	// cpu消耗
	Cpu float32 `json:"cpu"`

	// 内存消耗
	Memory int32 `json:"memory"`
}

func (o TaskSourceUsageEstimate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSourceUsageEstimate struct{}"
	}

	return strings.Join([]string{"TaskSourceUsageEstimate", string(data)}, " ")
}
