package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPostalAddressResponse struct {
	// 邮寄地址的个数，只有成功的时候才返回。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 邮寄地址，具体参见表2。

	PostalAddress  *[]CustomerPostalAddressV2 `json:"postal_address,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListPostalAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostalAddressResponse struct{}"
	}

	return strings.Join([]string{"ListPostalAddressResponse", string(data)}, " ")
}
