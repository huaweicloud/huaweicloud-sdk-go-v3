package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubCustomerRequest Request Object
type CreateSubCustomerRequest struct {
	Body *CreateCustomerV2Req `json:"body,omitempty"`
}

func (o CreateSubCustomerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubCustomerRequest struct{}"
	}

	return strings.Join([]string{"CreateSubCustomerRequest", string(data)}, " ")
}
