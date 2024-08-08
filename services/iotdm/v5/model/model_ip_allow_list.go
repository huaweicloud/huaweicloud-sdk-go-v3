package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpAllowList struct {

	// **参数说明**：白名单ip地址
	Address string `json:"address"`

	// **参数说明**：描述
	Description *string `json:"description,omitempty"`
}

func (o IpAllowList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpAllowList struct{}"
	}

	return strings.Join([]string{"IpAllowList", string(data)}, " ")
}
