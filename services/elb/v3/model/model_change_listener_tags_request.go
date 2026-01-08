package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeListenerTagsRequest Request Object
type ChangeListenerTagsRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId string `json:"listener_id"`

	Body *ChangeListenerTagsRequestBody `json:"body,omitempty"`
}

func (o ChangeListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ChangeListenerTagsRequest", string(data)}, " ")
}
