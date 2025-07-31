package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployTemplateInfoScheduleInfo 节点调度信息
type DeployTemplateInfoScheduleInfo struct {

	// 节点选择器
	NodeSelector *[]string `json:"node_selector,omitempty"`

	// pod容忍度
	PodTolerances *[]string `json:"pod_tolerances,omitempty"`
}

func (o DeployTemplateInfoScheduleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployTemplateInfoScheduleInfo struct{}"
	}

	return strings.Join([]string{"DeployTemplateInfoScheduleInfo", string(data)}, " ")
}
