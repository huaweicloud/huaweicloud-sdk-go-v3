package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 域名请求体
type CreateHostRequestBody struct {

	// 域名（域名只能由字母、数字、-、_和.组成，长度不能超过64个字符，如www.domain.com）
	Hostname string `json:"hostname"`

	// 防护域名初始绑定的策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// 源站信息
	Server []CloudWafServer `json:"server"`

	// 证书id，通过查询证书列表接口（ListCertificates）接口获取证书id   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificateid *string `json:"certificateid,omitempty"`

	// 证书名   - 对外协议为HTTP时不需要填写   - 对外协议HTTPS时为必填参数
	Certificatename *string `json:"certificatename,omitempty"`

	// 是否使用代理
	Proxy bool `json:"proxy"`

	// 域名描述
	Description *string `json:"description,omitempty"`
}

func (o CreateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostRequestBody", string(data)}, " ")
}
