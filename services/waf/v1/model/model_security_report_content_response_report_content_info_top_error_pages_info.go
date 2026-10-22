package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseReportContentInfoTopErrorPagesInfo **参数解释：** TOP异常URL信息，包含返回502、500、404等错误的URL列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseReportContentInfoTopErrorPagesInfo struct {

	// **参数解释：** TOP返回502错误的URL列表，按错误次数排序。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Abnormal502InfoList *[]SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal502InfoList `json:"abnormal_502_info_list,omitempty"`

	// **参数解释：** TOP返回500错误的URL列表，按错误次数排序。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Abnormal500InfoList *[]SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal500InfoList `json:"abnormal_500_info_list,omitempty"`

	// **参数解释：** TOP返回404错误的URL列表，按错误次数排序。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Abnormal404InfoList *[]SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal404InfoList `json:"abnormal_404_info_list,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoTopErrorPagesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoTopErrorPagesInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoTopErrorPagesInfo", string(data)}, " ")
}
