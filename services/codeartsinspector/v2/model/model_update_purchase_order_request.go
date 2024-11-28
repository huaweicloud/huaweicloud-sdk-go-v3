package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePurchaseOrderRequest Request Object
type UpdatePurchaseOrderRequest struct {

	// servicename,购买vss服务时使用\"webscan\"
	Service string `json:"service"`

	Body *UpdateCbcOrderRequestBody `json:"body,omitempty"`
}

func (o UpdatePurchaseOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePurchaseOrderRequest struct{}"
	}

	return strings.Join([]string{"UpdatePurchaseOrderRequest", string(data)}, " ")
}
