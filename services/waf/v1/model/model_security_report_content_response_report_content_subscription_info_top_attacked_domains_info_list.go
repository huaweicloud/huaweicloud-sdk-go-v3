package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedDomainsInfoList struct {

	// **参数解释：** 域名标识，包含域名及端口（如*:80表示所有域名的80端口）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该域名被攻击的总次数。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`

	// **参数解释：** 域名的Web标签，用于标识域名所属业务类型。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WebTag *string `json:"web_tag,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedDomainsInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedDomainsInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoTopAttackedDomainsInfoList", string(data)}, " ")
}
