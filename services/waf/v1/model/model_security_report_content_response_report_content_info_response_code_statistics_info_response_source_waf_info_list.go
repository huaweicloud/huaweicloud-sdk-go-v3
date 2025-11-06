package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoResponseSourceWafInfoList struct {

	// **参数解释：** 响应码标识（如504表示网关超时）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 时间线数据，按时间顺序排列的WAF响应码数量。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timeline *[]SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoTimeline `json:"timeline,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoResponseSourceWafInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoResponseSourceWafInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoResponseSourceWafInfoList", string(data)}, " ")
}
