package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderResponse Response Object
type ShowOrderResponse struct {

	// **参数解释**：订单关联的资源数量。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：订单关联的资源信息列表。
	Items          *[]OrderDetailItem `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderResponse struct{}"
	}

	return strings.Join([]string{"ShowOrderResponse", string(data)}, " ")
}
