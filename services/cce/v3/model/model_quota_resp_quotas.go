package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaRespQuotas **参数解释：** 模板配额 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type QuotaRespQuotas struct {

	// **参数解释：** 资源 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Resources *[]QuotaRespQuotasResources `json:"resources,omitempty"`
}

func (o QuotaRespQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaRespQuotas struct{}"
	}

	return strings.Join([]string{"QuotaRespQuotas", string(data)}, " ")
}
