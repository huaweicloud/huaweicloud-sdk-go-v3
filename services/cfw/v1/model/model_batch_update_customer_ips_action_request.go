package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateCustomerIpsActionRequest Request Object
type BatchUpdateCustomerIpsActionRequest struct {
	Body *IpsCustomerBatchDto `json:"body,omitempty"`
}

func (o BatchUpdateCustomerIpsActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCustomerIpsActionRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateCustomerIpsActionRequest", string(data)}, " ")
}
