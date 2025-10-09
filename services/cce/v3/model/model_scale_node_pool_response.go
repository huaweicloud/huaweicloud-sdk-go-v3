package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleNodePoolResponse Response Object
type ScaleNodePoolResponse struct {

	// **参数解释**： 订单ID，仅扩容包周期节点时返回 **取值范围**： 不涉及
	OrderID        *string `json:"orderID,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ScaleNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleNodePoolResponse struct{}"
	}

	return strings.Join([]string{"ScaleNodePoolResponse", string(data)}, " ")
}
