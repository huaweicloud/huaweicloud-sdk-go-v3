package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderIdResponse Response Object
type CreateOrderIdResponse struct {

	// **参数解释**：订单ID。 **取值范围**：不涉及。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOrderIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderIdResponse struct{}"
	}

	return strings.Join([]string{"CreateOrderIdResponse", string(data)}, " ")
}
