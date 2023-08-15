package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVendorRequest Request Object
type ShowVendorRequest struct {
}

func (o ShowVendorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVendorRequest struct{}"
	}

	return strings.Join([]string{"ShowVendorRequest", string(data)}, " ")
}
