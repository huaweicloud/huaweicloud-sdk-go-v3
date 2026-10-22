package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2AgencyRequest Request Object
type ShowV2AgencyRequest struct {

	// **参数解释**：委托名称。 **约束限制**：固定值ai_datalake_admin_trust。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AgencyName string `json:"agency_name"`
}

func (o ShowV2AgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2AgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowV2AgencyRequest", string(data)}, " ")
}
