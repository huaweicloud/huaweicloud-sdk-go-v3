package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DistinguishedName struct {

	// 证书通用名称（CN）。
	CommonName string `json:"common_name"`

	// 国家编码，需符合正则\"**[A-Za-z]{2}**\"。
	Country string `json:"country"`

	// 省市名称。
	State string `json:"state"`

	// 地区名称。
	Locality string `json:"locality"`

	// 组织名称。
	Organization string `json:"organization"`

	// 组织单元名称。
	OrganizationalUnit string `json:"organizational_unit"`
}

func (o DistinguishedName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinguishedName struct{}"
	}

	return strings.Join([]string{"DistinguishedName", string(data)}, " ")
}
