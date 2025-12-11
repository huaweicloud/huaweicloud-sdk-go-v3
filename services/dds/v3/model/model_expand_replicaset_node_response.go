package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandReplicasetNodeResponse Response Object
type ExpandReplicasetNodeResponse struct {

	// **参数解释：** 任务ID，仅按需实例返回该参数。 **取值范围：** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释：** 订单ID，仅包周期实例返回该参数。 **取值范围：** 不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandReplicasetNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandReplicasetNodeResponse struct{}"
	}

	return strings.Join([]string{"ExpandReplicasetNodeResponse", string(data)}, " ")
}
