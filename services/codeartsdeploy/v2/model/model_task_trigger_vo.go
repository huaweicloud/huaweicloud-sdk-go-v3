package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskTriggerVo 部署任务触发场景配置
type TaskTriggerVo struct {

	// 部署任务允许执行的场景。其中0：所有执行请求均可，1：只允许流水线触发
	TriggerSource *string `json:"trigger_source,omitempty"`

	// 当任务只允许流水线触发执行时，流水线传递的来源信息，当前只有CloudArtifact
	ArtifactSourceSystem *string `json:"artifact_source_system,omitempty"`

	// 当任务只允许流水线触发执行时，对应流水线源的制品仓库类型（generic、docker）
	ArtifactType *string `json:"artifact_type,omitempty"`
}

func (o TaskTriggerVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTriggerVo struct{}"
	}

	return strings.Join([]string{"TaskTriggerVo", string(data)}, " ")
}
