package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPostalAddressRequest struct {

	// 偏移量。默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的个数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPostalAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostalAddressRequest struct{}"
	}

	return strings.Join([]string{"ListPostalAddressRequest", string(data)}, " ")
}
