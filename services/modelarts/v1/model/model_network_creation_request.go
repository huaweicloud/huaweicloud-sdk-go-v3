package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkCreationRequest 网络创建信息，最终用户不感知子网。
type NetworkCreationRequest struct {

	// **参数解释**：API版本。 **约束限制**：不涉及。 **取值范围**：可选值如下： - v1 **默认取值**：不涉及。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Network：网络 **默认取值**：不涉及。
	Kind string `json:"kind"`

	Metadata *NetworkMetadataCreation `json:"metadata"`

	Spec *NetworkSpec `json:"spec"`
}

func (o NetworkCreationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkCreationRequest struct{}"
	}

	return strings.Join([]string{"NetworkCreationRequest", string(data)}, " ")
}
