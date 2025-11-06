package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsIndexConfigInfo struct {
	FullTextIndex *LtsFullTextIndexInfo `json:"fullTextIndex"`

	// 字段索引配置
	Fields *[]LtsFieldsInfo `json:"fields,omitempty"`

	// 是否开启可视化
	SqlAnalysisEnable *bool `json:"sqlAnalysisEnable,omitempty"`

	// 日志流id
	LogStreamId string `json:"logStreamId"`

	// **参数解释：** 快速分析采样日志条数。 **约束限制：** 不涉及。 **取值范围：** 最小值：100000 最大值：10000000 **默认取值：** 100000
	FastAnalysisSampleCount *int64 `json:"fastAnalysisSampleCount,omitempty"`
}

func (o LtsIndexConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsIndexConfigInfo struct{}"
	}

	return strings.Join([]string{"LtsIndexConfigInfo", string(data)}, " ")
}
