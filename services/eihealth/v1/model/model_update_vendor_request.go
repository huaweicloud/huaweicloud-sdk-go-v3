package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVendorRequest Request Object
type UpdateVendorRequest struct {
	Body *UpdateVendorRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UpdateVendorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVendorRequest struct{}"
	}

	return strings.Join([]string{"UpdateVendorRequest", string(data)}, " ")
}
