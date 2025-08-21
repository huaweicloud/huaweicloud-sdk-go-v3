package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustedIpAddressesResponse Response Object
type ListTrustedIpAddressesResponse struct {
	Body *[]TrustedIpAddressDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTrustedIpAddressesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedIpAddressesResponse struct{}"
	}

	return strings.Join([]string{"ListTrustedIpAddressesResponse", string(data)}, " ")
}
