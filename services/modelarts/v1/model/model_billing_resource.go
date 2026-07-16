package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BillingResource struct {

	// **参数解释：** 计费码。 **取值范围：** 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释：** 计费单元。 **取值范围：** 不涉及。
	UnitNum *int32 `json:"unit_num,omitempty"`
}

func (o BillingResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingResource struct{}"
	}

	return strings.Join([]string{"BillingResource", string(data)}, " ")
}
