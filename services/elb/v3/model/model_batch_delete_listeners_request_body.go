package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteListenersRequestBody This is a auto create Body Object
type BatchDeleteListenersRequestBody struct {

	// **参数解释**：待删除的监听器id列表。 **约束限制**：一次最多删除10个监听器。 **取值范围**：不涉及 **默认取值**：不涉及
	ListenerIds []string `json:"listener_ids"`
}

func (o BatchDeleteListenersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenersRequestBody", string(data)}, " ")
}
