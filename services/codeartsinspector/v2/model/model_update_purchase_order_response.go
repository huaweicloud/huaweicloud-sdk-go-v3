package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePurchaseOrderResponse Response Object
type UpdatePurchaseOrderResponse struct {

	// order_id
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePurchaseOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePurchaseOrderResponse struct{}"
	}

	return strings.Join([]string{"UpdatePurchaseOrderResponse", string(data)}, " ")
}
