package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagFilter struct {

	// **参数解释**： 资源ID，返回为集群ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 资源名，返回为集群名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags *string `json:"tags,omitempty"`

	// **参数解释**： EPS标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SysTags *string `json:"sys_tags,omitempty"`
}

func (o TagFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFilter struct{}"
	}

	return strings.Join([]string{"TagFilter", string(data)}, " ")
}
