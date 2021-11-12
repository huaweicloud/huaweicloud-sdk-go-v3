package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 命令耗时统计
type CommandTimeTaken struct {
	// 调用次数

	CallsSum int32 `json:"calls_sum"`
	// 耗时总数

	UsecSum float64 `json:"usec_sum"`
	// 命令名称

	CommandName string `json:"command_name"`
	// 耗时占比

	PerUsec string `json:"per_usec"`
	// 每次调用平均耗时

	AverageUsec float64 `json:"average_usec"`
}

func (o CommandTimeTaken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandTimeTaken struct{}"
	}

	return strings.Join([]string{"CommandTimeTaken", string(data)}, " ")
}
