package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetadata 资源池的metadata信息。
type PoolMetadata struct {

	// **参数解释**：资源池的ID。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：资源池的创建时间。例如\"2025-11-01T03:49:41Z\"。 **取值范围**：不涉及。
	CreationTimestamp string `json:"creationTimestamp"`

	Labels *PoolMetaLabels `json:"labels"`

	Annotations *PoolMetaAnnotations `json:"annotations,omitempty"`
}

func (o PoolMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetadata struct{}"
	}

	return strings.Join([]string{"PoolMetadata", string(data)}, " ")
}
