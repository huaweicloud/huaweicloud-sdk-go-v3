package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 设置IP白名单结构体
type ModifyIpWhitelistBody struct {

	// 是否启用白名单（true/false）。
	EnableWhitelist bool `json:"enable_whitelist"`

	// IP白名单分组列表。
	Whitelist []Whitelist `json:"whitelist"`
}

func (o ModifyIpWhitelistBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyIpWhitelistBody struct{}"
	}

	return strings.Join([]string{"ModifyIpWhitelistBody", string(data)}, " ")
}
