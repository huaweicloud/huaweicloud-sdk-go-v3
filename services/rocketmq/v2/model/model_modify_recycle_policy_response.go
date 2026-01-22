package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRecyclePolicyResponse Response Object
type ModifyRecyclePolicyResponse struct {

	// **参数解释**： 保留天数。 **约束限制**： 不涉及。 **取值范围**： 1~7。 **默认取值**： 不涉及。
	RetentionDays *int32 `json:"retention_days,omitempty"`

	// **参数解释**： 是否使用回收站。 **约束限制**： 不涉及。 **取值范围**： - true：使用回收站。 - false：不使用回收站。 **默认取值**： 不涉及。
	DefaultUseRecycle *bool `json:"default_use_recycle,omitempty"`

	// **参数解释**： 回收实例列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RecycleInstances *[]InstanceRecycleInfo `json:"recycle_instances,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ModifyRecyclePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRecyclePolicyResponse struct{}"
	}

	return strings.Join([]string{"ModifyRecyclePolicyResponse", string(data)}, " ")
}
