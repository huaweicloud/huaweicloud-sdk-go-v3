package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomerIpsRequest Request Object
type CreateCustomerIpsRequest struct {
	Body *CustomerIpsSaveDto `json:"body,omitempty"`
}

func (o CreateCustomerIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomerIpsRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomerIpsRequest", string(data)}, " ")
}
