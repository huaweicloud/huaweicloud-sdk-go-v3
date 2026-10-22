package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal404InfoList struct {

	// **参数解释：** 返回404错误的URL路径。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该URL返回404错误的总次数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 该URL所属的域名。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Host *string `json:"host,omitempty"`

	// **参数解释：** 域名的Web标签。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WebTag *string `json:"web_tag,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal404InfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal404InfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoTopErrorPagesInfoAbnormal404InfoList", string(data)}, " ")
}
