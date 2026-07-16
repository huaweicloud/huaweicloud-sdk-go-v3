package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetworkMetadata struct {

	// **参数解释**：网络资源的ID。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：网络的创建时间。例如\"2025-11-01T03:49:41Z\"。 **取值范围**：不涉及。
	CreationTimestamp string `json:"creationTimestamp"`

	Labels *NetworkMetadataLabels `json:"labels"`

	Annotations *NetworkMetadataAnnotations `json:"annotations,omitempty"`
}

func (o NetworkMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkMetadata struct{}"
	}

	return strings.Join([]string{"NetworkMetadata", string(data)}, " ")
}
