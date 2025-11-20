package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePartitionResponse Response Object
type UpdatePartitionResponse struct {

	// **参数解释**：  API类型  **约束限制**：  不允许修改 **取值范围**：  不涉及  **默认取值**：  Partition
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *PartitionMetadata `json:"metadata,omitempty"`

	Spec           *PartitionSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdatePartitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePartitionResponse struct{}"
	}

	return strings.Join([]string{"UpdatePartitionResponse", string(data)}, " ")
}
