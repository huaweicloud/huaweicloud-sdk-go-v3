package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseP2cVgw struct {

	// P2C VPN网关ID
	Id *string `json:"id,omitempty"`
}

func (o ResponseP2cVgw) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseP2cVgw struct{}"
	}

	return strings.Join([]string{"ResponseP2cVgw", string(data)}, " ")
}
