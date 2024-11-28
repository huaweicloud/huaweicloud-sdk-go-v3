package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePurchaseOrderResponse Response Object
type CreatePurchaseOrderResponse struct {

	// order_id
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePurchaseOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePurchaseOrderResponse struct{}"
	}

	return strings.Join([]string{"CreatePurchaseOrderResponse", string(data)}, " ")
}
