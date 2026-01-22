package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasRespQuotas **参数解释**： 配额信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type QuotasRespQuotas struct {

	// **参数解释**： 配额列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Resources *[]QuotaResourceEntity `json:"resources,omitempty"`
}

func (o QuotasRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasRespQuotas struct{}"
	}

	return strings.Join([]string{"QuotasRespQuotas", string(data)}, " ")
}
