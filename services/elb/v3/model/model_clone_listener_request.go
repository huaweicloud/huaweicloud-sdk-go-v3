package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneListenerRequest Request Object
type CloneListenerRequest struct {

	// **参数解释**：被复制监听器ID（UUID）。  **约束限制**：不涉及  **取值范围**：标准的UUID格式，长度为36个字符。  **默认取值**：不涉及
	ListenerId string `json:"listener_id"`

	Body *CloneListenerRequestBody `json:"body,omitempty"`
}

func (o CloneListenerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneListenerRequest struct{}"
	}

	return strings.Join([]string{"CloneListenerRequest", string(data)}, " ")
}
