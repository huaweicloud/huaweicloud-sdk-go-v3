package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultBindRules struct {

	// 按tags过滤自动绑定的资源  最小长度：1  最大长度：5
	Tags *[]BindRulesTags `json:"tags,omitempty"`
}

func (o VaultBindRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultBindRules struct{}"
	}

	return strings.Join([]string{"VaultBindRules", string(data)}, " ")
}
