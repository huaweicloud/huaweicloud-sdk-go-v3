package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertDistinguishedName struct {

	// 证书通用名称（CN），名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、\"*\"、字母、数字、汉字组成，长度不能超过64位字符。
	CommonName string `json:"common_name"`

	// 国家编码，只能由英文组成，长度为2位字符。若不传入，则默认继承父CA对应的值。
	Country *string `json:"country,omitempty"`

	// 省市名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过128位字符。若不传入，则默认继承父CA对应的值。
	State *string `json:"state,omitempty"`

	// 地区名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过128位字符。若不传入，则默认继承父CA对应的值。
	Locality *string `json:"locality,omitempty"`

	// 组织名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过64位字符。若不传入，则默认继承父CA对应的值。
	Organization *string `json:"organization,omitempty"`

	// 组织单元名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过64位字符。若不传入，则默认继承父CA对应的值。
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
}

func (o CertDistinguishedName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDistinguishedName struct{}"
	}

	return strings.Join([]string{"CertDistinguishedName", string(data)}, " ")
}
