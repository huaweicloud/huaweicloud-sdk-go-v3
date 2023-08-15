package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineLatestRunStageStatusList struct {

	// 阶段名
	Name *string `json:"name,omitempty"`

	// 阶段序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 阶段状态
	Status *string `json:"status,omitempty"`

	// 阶段开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 阶段结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o PipelineLatestRunStageStatusList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineLatestRunStageStatusList struct{}"
	}

	return strings.Join([]string{"PipelineLatestRunStageStatusList", string(data)}, " ")
}
