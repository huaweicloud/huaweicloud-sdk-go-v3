package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateAddonPrecheckResponse Response Object
type BatchCreateAddonPrecheckResponse struct {

	// **参数解释：** API类型 **约束限制：** 该值不可修改 **取值范围：** 固定值\"AddonCheck\" **默认取值：** AddonCheck
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本 **约束限制：** 该值不可修改 **取值范围：** 固定值\"v3\" **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Spec *AddonCheckSpec `json:"spec,omitempty"`

	Status         *AddonCheckStatus `json:"status,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o BatchCreateAddonPrecheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAddonPrecheckResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateAddonPrecheckResponse", string(data)}, " ")
}
