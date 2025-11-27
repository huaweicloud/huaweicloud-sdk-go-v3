package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVendorAccountRequest Request Object
type CreateVendorAccountRequest struct {
	Body *VendorAccountCreateRequest `json:"body,omitempty"`
}

func (o CreateVendorAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVendorAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateVendorAccountRequest", string(data)}, " ")
}
