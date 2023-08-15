package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpAddressGroupForceResponse Response Object
type DeleteIpAddressGroupForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteIpAddressGroupForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpAddressGroupForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpAddressGroupForceResponse", string(data)}, " ")
}
