package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTrustedIpAddressRequest Request Object
type AddTrustedIpAddressRequest struct {

	// 仓库id
	Id int32 `json:"id"`

	Body *AddTrustedIpAddressRequestBody `json:"body,omitempty"`
}

func (o AddTrustedIpAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrustedIpAddressRequest struct{}"
	}

	return strings.Join([]string{"AddTrustedIpAddressRequest", string(data)}, " ")
}
