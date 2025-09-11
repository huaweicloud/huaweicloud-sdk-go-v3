package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentEditRequest struct {

	// CPU阈值,0-100之间
	CpuThreshold int32 `json:"cpu_threshold"`

	// 内存阈值，0-100之间
	MemThreshold int32 `json:"mem_threshold"`
}

func (o AgentEditRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentEditRequest struct{}"
	}

	return strings.Join([]string{"AgentEditRequest", string(data)}, " ")
}
