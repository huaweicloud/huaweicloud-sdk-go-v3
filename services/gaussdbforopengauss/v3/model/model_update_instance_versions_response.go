package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceVersionsResponse Response Object
type UpdateInstanceVersionsResponse struct {

	// **参数解释**: 任务ID。按需实例时仅返回任务ID。 **取值范围**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 订单ID。包周期实例时仅返回订单ID。 **取值范围**: 不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceVersionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceVersionsResponse", string(data)}, " ")
}
