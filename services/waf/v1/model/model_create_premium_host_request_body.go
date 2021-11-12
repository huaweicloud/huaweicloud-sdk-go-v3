package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建独享模式域名的请求
type CreatePremiumHostRequestBody struct {
	// 证书id

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名称

	Certificatename *string `json:"certificatename,omitempty"`
	// 防护域名或IP（可带端口）

	Hostname string `json:"hostname"`
	// 是否使用代理

	Proxy *bool `json:"proxy,omitempty"`
	// 防护域名初始绑定的策略ID

	Policyid *string `json:"policyid,omitempty"`
	// 独享模式回源服务器配置

	Server *[]PremiumWafServer `json:"server,omitempty"`
}

func (o CreatePremiumHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequestBody", string(data)}, " ")
}
