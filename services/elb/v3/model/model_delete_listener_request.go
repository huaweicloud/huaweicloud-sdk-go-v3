package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteListenerRequest Request Object
type DeleteListenerRequest struct {

	// **参数解释**：监听器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId string `json:"listener_id"`
}

func (o DeleteListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerRequest", string(data)}, " ")
}
