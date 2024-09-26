package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrustedIpAddressResponse Response Object
type DeleteTrustedIpAddressResponse struct {

	// 状态码
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrustedIpAddressResponse", string(data)}, " ")
}
