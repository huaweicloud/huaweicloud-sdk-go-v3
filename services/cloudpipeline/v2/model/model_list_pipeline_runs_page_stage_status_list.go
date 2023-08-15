package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelineRunsPageStageStatusList struct {

	// 阶段名称
	Name *string `json:"name,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListPipelineRunsPageStageStatusList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsPageStageStatusList struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsPageStageStatusList", string(data)}, " ")
}
