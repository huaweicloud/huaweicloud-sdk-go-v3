package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务实例信息体。
type JobNodeInfo struct {
	Spec *JobNodeSpecInfo `json:"spec"`

	Vpc *JobNodeVpcInfo `json:"vpc,omitempty"`
}

func (o JobNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobNodeInfo struct{}"
	}

	return strings.Join([]string{"JobNodeInfo", string(data)}, " ")
}
