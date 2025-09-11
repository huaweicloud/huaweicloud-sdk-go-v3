package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchReplicaResponse Response Object
type SwitchReplicaResponse struct {

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 订单ID。 **取值范围**: 不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchReplicaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchReplicaResponse struct{}"
	}

	return strings.Join([]string{"SwitchReplicaResponse", string(data)}, " ")
}
