package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建独享模式域名的请求体
type CreatePremiumHostRequestBody struct {

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty" xml:"certificateid"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty" xml:"certificatename"`

	// 防护域名或IP（可带端口）
	Hostname string `json:"hostname" xml:"hostname"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy bool `json:"proxy" xml:"proxy"`

	// 防护域名初始绑定的防护策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty" xml:"policyid"`

	// 防护域名的源站服务器配置信息
	Server []PremiumWafServer `json:"server" xml:"server"`

	BlockPage *BlockPage `json:"block_page,omitempty" xml:"block_page"`

	// 防护域名备注
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o CreatePremiumHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePremiumHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePremiumHostRequestBody", string(data)}, " ")
}
