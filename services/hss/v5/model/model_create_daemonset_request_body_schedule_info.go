package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDaemonsetRequestBodyScheduleInfo 节点调度信息
type CreateDaemonsetRequestBodyScheduleInfo struct {

	// 节点选择器
	NodeSelector *[]string `json:"node_selector,omitempty"`

	// pod容忍度
	PodTolerances *[]string `json:"pod_tolerances,omitempty"`
}

func (o CreateDaemonsetRequestBodyScheduleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDaemonsetRequestBodyScheduleInfo struct{}"
	}

	return strings.Join([]string{"CreateDaemonsetRequestBodyScheduleInfo", string(data)}, " ")
}
