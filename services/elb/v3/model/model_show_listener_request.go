package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListenerRequest Request Object
type ShowListenerRequest struct {

	// **参数解释**：监听器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	ListenerId string `json:"listener_id"`
}

func (o ShowListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerRequest struct{}"
	}

	return strings.Join([]string{"ShowListenerRequest", string(data)}, " ")
}
