package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoOverviewStatisticsInfoList struct {

	// **参数解释：** 统计维度标识（如ACCESS表示访问类统计）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该统计维度的总数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 拦截数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 不涉及
	BlockNum *int64 `json:"block_num,omitempty"`

	// **参数解释：** 仅记录数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 不涉及
	OnlyLogNum *int64 `json:"only_log_num,omitempty"`

	// **参数解释：** TOP域名列表，按统计数量排序的域名信息。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopDomains *[]SecurityReportContentResponseReportContentInfoTopDomains `json:"top_domains,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoOverviewStatisticsInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoOverviewStatisticsInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoOverviewStatisticsInfoList", string(data)}, " ")
}
