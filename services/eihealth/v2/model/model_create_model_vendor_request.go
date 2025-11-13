package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelVendorRequest Request Object
type CreateModelVendorRequest struct {
	Body *CreateModelVendorReq `json:"body,omitempty"`
}

func (o CreateModelVendorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelVendorRequest struct{}"
	}

	return strings.Join([]string{"CreateModelVendorRequest", string(data)}, " ")
}
