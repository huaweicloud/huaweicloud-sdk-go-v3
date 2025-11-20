package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePartitionResponse Response Object
type CreatePartitionResponse struct {

	// **参数解释**：  API类型  **约束限制**：  不允许修改 **取值范围**：  不涉及  **默认取值**：  Partition
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *PartitionMetadata `json:"metadata,omitempty"`

	Spec           *PartitionSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreatePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionResponse struct{}"
	}

	return strings.Join([]string{"CreatePartitionResponse", string(data)}, " ")
}
