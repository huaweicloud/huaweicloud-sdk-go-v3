package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPostalAddressResponse struct {

	// 邮寄地址的个数，只有成功的时候才返回。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 邮寄地址，具体参见表2。
	PostalAddress  *[]CustomerPostalAddressV2 `json:"postal_address,omitempty" xml:"postal_address"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListPostalAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostalAddressResponse struct{}"
	}

	return strings.Join([]string{"ListPostalAddressResponse", string(data)}, " ")
}
