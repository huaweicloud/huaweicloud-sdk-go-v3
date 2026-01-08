package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneListenerRequestBody 复制监听器接口请求体。
type CloneListenerRequestBody struct {

	// **参数解释**：复制后的新监听器相关配置。 **约束限制**：不涉及
	TargetListenerParams []CloneListenerOption `json:"target_listener_params"`
}

func (o CloneListenerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneListenerRequestBody struct{}"
	}

	return strings.Join([]string{"CloneListenerRequestBody", string(data)}, " ")
}
