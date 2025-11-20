package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateAntiTamperRulesRequestBody struct {

	// 防护网站，查询云模式防护域名列表（ListHost）接口获取防护域名，响应体中的的hostname字段
	Hostname string `json:"hostname"`

	// 防篡改规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"号结尾代表路径前缀
	Url string `json:"url"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateAntiTamperRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAntiTamperRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateAntiTamperRulesRequestBody", string(data)}, " ")
}
