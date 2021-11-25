package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建独享模式域名的请求
type CreatePremiumHostRequestBody struct {
	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   -对外协议HTTPS时为必填参数

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名   - 对外协议为HTTP时不需要填写   -对外协议HTTPS时为必填参数

	Certificatename *string `json:"certificatename,omitempty"`
	// 防护域名或IP（可带端口）

	Hostname string `json:"hostname"`
	// 是否使用代理

	Proxy bool `json:"proxy"`
	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id

	Policyid *string `json:"policyid,omitempty"`
	// 独享模式回源服务器配置

	Server []PremiumWafServer `json:"server"`
}

func (o CreatePremiumHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequestBody", string(data)}, " ")
}
