package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpGroupItem struct {

	// ip或者网段
	Ip string `json:"ip"`

	// ip的描述
	Description string `json:"description"`
}

func (o IpGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupItem struct{}"
	}

	return strings.Join([]string{"IpGroupItem", string(data)}, " ")
}
