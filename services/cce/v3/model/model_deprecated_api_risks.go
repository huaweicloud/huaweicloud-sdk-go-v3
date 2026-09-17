package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeprecatedApiRisks 废弃API风险来源
type DeprecatedApiRisks struct {

	// **参数解释：** 请求路径，如/apis/policy/v1beta1/podsecuritypolicies。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释：** 客户端信息。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UserAgent *string `json:"userAgent,omitempty"`
}

func (o DeprecatedApiRisks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeprecatedApiRisks struct{}"
	}

	return strings.Join([]string{"DeprecatedApiRisks", string(data)}, " ")
}
