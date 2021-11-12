package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAntiTamperRulesRequestBody struct {
	// 防护网站（查询云模式防护域名列表获取防护域名，响应体的hostname字段）

	Hostname *string `json:"hostname,omitempty"`
	// 防篡改的url

	Url *string `json:"url,omitempty"`
	// 规则描述

	Description *string `json:"description,omitempty"`
}

func (o CreateAntiTamperRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAntiTamperRulesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAntiTamperRulesRequestBody", string(data)}, " ")
}
