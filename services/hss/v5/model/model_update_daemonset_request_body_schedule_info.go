package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDaemonsetRequestBodyScheduleInfo 节点调度信息
type UpdateDaemonsetRequestBodyScheduleInfo struct {

	// 节点选择器
	NodeSelector *[]string `json:"node_selector,omitempty"`

	// pod容忍度
	PodTolerances *[]string `json:"pod_tolerances,omitempty"`
}

func (o UpdateDaemonsetRequestBodyScheduleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDaemonsetRequestBodyScheduleInfo struct{}"
	}

	return strings.Join([]string{"UpdateDaemonsetRequestBodyScheduleInfo", string(data)}, " ")
}
