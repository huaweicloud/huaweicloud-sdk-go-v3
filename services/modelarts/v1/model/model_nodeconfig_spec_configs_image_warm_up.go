package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeconfigSpecConfigsImageWarmUp struct {

	// **参数解释**：镜像预热任务ID。 **约束限制**：系统自动生成。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：镜像预热任务名称。 **约束限制**：由4~32位小写字母、数字或\"-\"组成。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：镜像预热任务的创建时间。 **约束限制**：系统自动填充。 **取值范围**：RFC3339格式时间。 **默认取值**：不涉及。
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**：镜像类型。 **约束限制**：不涉及。 **取值范围**：SWR。 **默认取值**：不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**：需要预热的镜像地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Image *string `json:"image,omitempty"`
}

func (o NodeconfigSpecConfigsImageWarmUp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpecConfigsImageWarmUp struct{}"
	}

	return strings.Join([]string{"NodeconfigSpecConfigsImageWarmUp", string(data)}, " ")
}
