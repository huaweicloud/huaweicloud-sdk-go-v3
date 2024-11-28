package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePurchaseOrderRequest Request Object
type CreatePurchaseOrderRequest struct {

	// servicename,购买vss服务时使用\"webscan\"
	Service string `json:"service"`

	Body *CreateCbcOrderRequestBody `json:"body,omitempty"`
}

func (o CreatePurchaseOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePurchaseOrderRequest struct{}"
	}

	return strings.Join([]string{"CreatePurchaseOrderRequest", string(data)}, " ")
}
