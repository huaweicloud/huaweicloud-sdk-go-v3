package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoTopDomains struct {

	// **参数解释：** 域名标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该域名在对应统计维度下的数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 域名标识，包含域名及关联标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Host *string `json:"host,omitempty"`

	// **参数解释：** 域名的Web标签。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WebTag *string `json:"web_tag,omitempty"`

	// **参数解释：** 该域名拦截数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 不涉及
	BlockNum *int64 `json:"block_num,omitempty"`

	// **参数解释：** 该域名仅记录数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 不涉及
	OnlyLogNum *int64 `json:"only_log_num,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoTopDomains) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoTopDomains struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoTopDomains", string(data)}, " ")
}
