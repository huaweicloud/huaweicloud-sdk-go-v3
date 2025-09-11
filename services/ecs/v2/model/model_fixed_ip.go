package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FixedIp struct {
	SubnetId *string `json:"subnet_id,omitempty"`

	IpAddress *string `json:"ip_address,omitempty"`
}

func (o FixedIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixedIp struct{}"
	}

	return strings.Join([]string{"FixedIp", string(data)}, " ")
}
