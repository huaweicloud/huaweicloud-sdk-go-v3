package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebsiteExtendProperties 其它属性
type WebsiteExtendProperties struct {

	// mac地址
	MacAddr *string `json:"mac_addr,omitempty"`
}

func (o WebsiteExtendProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebsiteExtendProperties struct{}"
	}

	return strings.Join([]string{"WebsiteExtendProperties", string(data)}, " ")
}
