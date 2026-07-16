package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BillingInfo struct {

	// **参数解释**：计费码。 **取值范围**：不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**：计费单元。 **取值范围**：不涉及。
	UnitNum *int32 `json:"unit_num,omitempty"`
}

func (o BillingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingInfo struct{}"
	}

	return strings.Join([]string{"BillingInfo", string(data)}, " ")
}
