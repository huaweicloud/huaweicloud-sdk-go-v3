package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DistinguishedName struct {

	// 证书通用名称（CN），名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、\"*\"、字母、数字、汉字组成，长度不能超过64位字符。
	CommonName string `json:"common_name"`

	// 国家编码，只能由英文组成，长度为2位字符。
	Country string `json:"country"`

	// 省市名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过128位字符。
	State string `json:"state"`

	// 地区名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过128位字符。
	Locality string `json:"locality"`

	// 组织名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过64位字符。
	Organization string `json:"organization"`

	// 组织单元名称，名称只能由\"-\"、\"_\"、\" \"、\".\"、\",\"、字母、数字、汉字组成，长度不能超过64位字符。
	OrganizationalUnit string `json:"organizational_unit"`
}

func (o DistinguishedName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinguishedName struct{}"
	}

	return strings.Join([]string{"DistinguishedName", string(data)}, " ")
}
