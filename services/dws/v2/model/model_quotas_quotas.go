package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotasQuotas **参数解释**： 配额列表对象。 **取值范围**： 不涉及。
type QuotasQuotas struct {

	// **参数解释**： 资源列表对象。 **取值范围**： 不涉及。
	Resources *[]QuotasResource `json:"resources,omitempty"`
}

func (o QuotasQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotasQuotas struct{}"
	}

	return strings.Join([]string{"QuotasQuotas", string(data)}, " ")
}
