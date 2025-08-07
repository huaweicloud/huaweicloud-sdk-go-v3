package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumHostAccessStatusRequestBody **参数解释：** 域名接入状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpdatePremiumHostAccessStatusRequestBody struct {

	// **参数解释：** 域名接入状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AccessStatus *int32 `json:"access_status,omitempty"`
}

func (o UpdatePremiumHostAccessStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostAccessStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostAccessStatusRequestBody", string(data)}, " ")
}
