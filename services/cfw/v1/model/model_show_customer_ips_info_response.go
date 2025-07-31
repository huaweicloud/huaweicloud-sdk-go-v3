package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomerIpsInfoResponse Response Object
type ShowCustomerIpsInfoResponse struct {
	Data           *CustomerIpsVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowCustomerIpsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerIpsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerIpsInfoResponse", string(data)}, " ")
}
