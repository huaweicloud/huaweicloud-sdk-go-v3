package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerRef **参数解释**：关联的监听器。  **默认取值**：不涉及
type ListenerRef struct {

	// **参数解释**：监听器ID。  **取值范围**：不涉及
	Id string `json:"id"`
}

func (o ListenerRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerRef struct{}"
	}

	return strings.Join([]string{"ListenerRef", string(data)}, " ")
}
