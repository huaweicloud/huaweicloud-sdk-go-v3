package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomerIpsResponse Response Object
type UpdateCustomerIpsResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateCustomerIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomerIpsResponse struct{}"
	}

	return strings.Join([]string{"UpdateCustomerIpsResponse", string(data)}, " ")
}
