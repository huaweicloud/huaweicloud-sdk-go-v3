package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVendorRequest Request Object
type ListVendorRequest struct {
}

func (o ListVendorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVendorRequest struct{}"
	}

	return strings.Join([]string{"ListVendorRequest", string(data)}, " ")
}
