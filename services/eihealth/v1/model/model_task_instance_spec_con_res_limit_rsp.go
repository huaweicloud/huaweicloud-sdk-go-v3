package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceSpecConResLimitRsp struct {

	// CPU限制值
	Cpu *string `json:"cpu,omitempty"`

	// 内存限制值
	Memory *string `json:"memory,omitempty"`
}

func (o TaskInstanceSpecConResLimitRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceSpecConResLimitRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceSpecConResLimitRsp", string(data)}, " ")
}
