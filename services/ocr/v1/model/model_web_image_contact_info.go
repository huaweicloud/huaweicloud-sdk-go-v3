package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebImageContactInfo
type WebImageContactInfo struct {

	// 传入contact_info时的返回，为姓名。
	Name *string `json:"name,omitempty"`

	// 传入contact_info时的返回，联系电话。
	Phone *string `json:"phone,omitempty"`

	// 传入contact_info时的返回，省。
	Province *string `json:"province,omitempty"`

	// 传入contact_info时的返回，市。
	City *string `json:"city,omitempty"`

	// 传入contact_info时的返回，县区。
	District *string `json:"district,omitempty"`

	// 传入contact_info时的返回，详细地址（不含省市区）。
	DetailAddress *string `json:"detail_address,omitempty"`
}

func (o WebImageContactInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageContactInfo struct{}"
	}

	return strings.Join([]string{"WebImageContactInfo", string(data)}, " ")
}
