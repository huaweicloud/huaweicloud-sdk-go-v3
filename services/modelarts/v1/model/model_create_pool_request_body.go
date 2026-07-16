package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolRequestBody 资源池创建请求体。
type CreatePoolRequestBody struct {

	// **参数解释**：API版本。 **约束限制**：不涉及。 **取值范围**：可选值如下： - v2 **默认取值**：不涉及。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Pool：资源池 **默认取值**：不涉及。
	Kind string `json:"kind"`

	Metadata *PoolMetadataCreation `json:"metadata"`

	Spec *PoolSpecCreation `json:"spec,omitempty"`
}

func (o CreatePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePoolRequestBody", string(data)}, " ")
}
