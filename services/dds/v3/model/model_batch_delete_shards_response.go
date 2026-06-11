package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteShardsResponse Response Object
type BatchDeleteShardsResponse struct {

	// **参数解释：** 任务ID，仅按需实例返回该参数。 **取值范围：** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释：** 订单ID，仅包周期实例返回该参数。 **取值范围：** 不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteShardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteShardsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteShardsResponse", string(data)}, " ")
}
