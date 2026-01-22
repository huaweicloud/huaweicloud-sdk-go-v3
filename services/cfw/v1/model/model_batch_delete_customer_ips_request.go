package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomerIpsRequest Request Object
type BatchDeleteCustomerIpsRequest struct {
	Body *IpsCustomerBatchDto `json:"body,omitempty"`
}

func (o BatchDeleteCustomerIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomerIpsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomerIpsRequest", string(data)}, " ")
}
