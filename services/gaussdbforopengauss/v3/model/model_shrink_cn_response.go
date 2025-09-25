package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkCnResponse Response Object
type ShrinkCnResponse struct {

	// **参数解释**： 修改实例名称的任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 订单ID。包周期实例时仅返回订单ID。退订到期后的包周期实例返回空。 **取值范围**： 不涉及。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkCnResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkCnResponse struct{}"
	}

	return strings.Join([]string{"ShrinkCnResponse", string(data)}, " ")
}
