package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpValue struct {

	// IP地址信息。
	Ip *string `json:"ip,omitempty"`
}

func (o IpValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpValue struct{}"
	}

	return strings.Join([]string{"IpValue", string(data)}, " ")
}
