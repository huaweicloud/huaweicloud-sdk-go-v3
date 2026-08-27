package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceReadonlyRequestBody 设置/解除只读。
type SetInstanceReadonlyRequestBody struct {

	// **参数解释**：  设置或者解除实例只读。  **约束限制**：  不涉及。  **取值范围**：  - true: 设置只读。 - false: 解除只读。 **默认取值**：  不涉及。
	Readonly bool `json:"readonly"`
}

func (o SetInstanceReadonlyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceReadonlyRequestBody struct{}"
	}

	return strings.Join([]string{"SetInstanceReadonlyRequestBody", string(data)}, " ")
}
