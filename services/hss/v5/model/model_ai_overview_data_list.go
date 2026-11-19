package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiOverviewDataList AI Defender概览数据项（按产品类型分组）
type AiOverviewDataList struct {

	// **参数解释**： 产品类型 **取值范围**： - AgentArts - OfficeAce - OpenClaw
	Source *string `json:"source,omitempty"`

	// **参数解释**: 维度统计列表（type=alarm/risk时返回） **取值范围**: 不涉及
	DimensionStats *[]AiOverviewDimensionStats `json:"dimension_stats,omitempty"`

	SourceInfo *AiOverviewSourceInfo `json:"source_info,omitempty"`
}

func (o AiOverviewDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOverviewDataList struct{}"
	}

	return strings.Join([]string{"AiOverviewDataList", string(data)}, " ")
}
