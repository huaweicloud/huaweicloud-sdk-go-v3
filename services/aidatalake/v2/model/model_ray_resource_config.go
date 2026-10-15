package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayResourceConfig **参数解释**：Ray节点资源配置。 **约束限制**：仅提交作业到RayJob端点时配置。
type RayResourceConfig struct {
	HeadResourceSpec *RayHeadResourceSpec `json:"head_resource_spec"`

	// **参数解释**：worker资源配置。 **约束限制**：仅提交作业到RayJob端点时配置。
	WorkerResourceSpec []RayWorkerResourceSpec `json:"worker_resource_spec"`
}

func (o RayResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayResourceConfig struct{}"
	}

	return strings.Join([]string{"RayResourceConfig", string(data)}, " ")
}
