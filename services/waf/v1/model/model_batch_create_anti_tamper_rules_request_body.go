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

	// 策略id
	PolicyIds *[]string `json:"policy_ids,omitempty"`
}

func (o BatchCreateAntiTamperRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAntiTamperRulesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateAntiTamperRulesRequestBody", string(data)}, " ")
}
