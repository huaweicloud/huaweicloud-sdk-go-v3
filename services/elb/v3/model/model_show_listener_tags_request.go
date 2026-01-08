package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListenerTagsRequest Request Object
type ShowListenerTagsRequest struct {

	// **参数解释**：监听器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerTagsRequest", string(data)}, " ")
}
