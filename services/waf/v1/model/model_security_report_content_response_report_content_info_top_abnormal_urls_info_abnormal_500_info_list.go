package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoTopAbnormalUrlsInfoAbnormal500InfoList struct {

	// **参数解释：** 返回500错误的URL路径。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该URL返回500错误的总次数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 该URL所属的域名。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Host *string `json:"host,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoTopAbnormalUrlsInfoAbnormal500InfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoTopAbnormalUrlsInfoAbnormal500InfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoTopAbnormalUrlsInfoAbnormal500InfoList", string(data)}, " ")
}
