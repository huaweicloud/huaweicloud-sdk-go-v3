package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelinesPageLatestRunStageStatusList struct {

	// 阶段名
	Name *string `json:"name,omitempty"`

	// 序列号
	Sequence *int32 `json:"sequence,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 阶段ID
	Id *string `json:"id,omitempty"`
}

func (o ListPipelinesPageLatestRunStageStatusList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPageLatestRunStageStatusList struct{}"
	}

	return strings.Join([]string{"ListPipelinesPageLatestRunStageStatusList", string(data)}, " ")
}
