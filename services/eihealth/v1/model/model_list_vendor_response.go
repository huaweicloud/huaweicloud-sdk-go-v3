package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVendorResponse Response Object
type ListVendorResponse struct {

	// 供应商个数
	Count *int32 `json:"count,omitempty"`

	// 供应商列表
	Vendors        *[]VendorDto `json:"vendors,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListVendorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVendorResponse struct{}"
	}

	return strings.Join([]string{"ListVendorResponse", string(data)}, " ")
}
