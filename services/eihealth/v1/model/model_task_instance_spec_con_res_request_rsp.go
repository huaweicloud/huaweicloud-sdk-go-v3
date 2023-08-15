package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceSpecConResRequestRsp struct {

	// CPU申请值
	Cpu *string `json:"cpu,omitempty"`

	// 内存申请值
	Memory *string `json:"memory,omitempty"`
}

func (o TaskInstanceSpecConResRequestRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceSpecConResRequestRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceSpecConResRequestRsp", string(data)}, " ")
}
