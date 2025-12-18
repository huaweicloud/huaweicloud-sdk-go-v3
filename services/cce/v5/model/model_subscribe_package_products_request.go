package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribePackageProductsRequest Request Object
type SubscribePackageProductsRequest struct {
	Body *SubscribePackageProductsRequestBody `json:"body,omitempty"`
}

func (o SubscribePackageProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribePackageProductsRequest struct{}"
	}

	return strings.Join([]string{"SubscribePackageProductsRequest", string(data)}, " ")
}
