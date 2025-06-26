package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadResource **参数解释**： 资源池资源项。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadResource struct {

	// **参数解释**： 资源名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceName string `json:"resource_name"`

	// **参数解释**： 资源属性值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceValue int32 `json:"resource_value"`
}

func (o WorkloadResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadResource struct{}"
	}

	return strings.Join([]string{"WorkloadResource", string(data)}, " ")
}
