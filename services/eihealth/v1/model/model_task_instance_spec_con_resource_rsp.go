package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskInstanceSpecConResourceRsp struct {
	Limits *TaskInstanceSpecConResLimitRsp `json:"limits,omitempty"`

	Requests *TaskInstanceSpecConResRequestRsp `json:"requests,omitempty"`
}

func (o TaskInstanceSpecConResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInstanceSpecConResourceRsp struct{}"
	}

	return strings.Join([]string{"TaskInstanceSpecConResourceRsp", string(data)}, " ")
}
