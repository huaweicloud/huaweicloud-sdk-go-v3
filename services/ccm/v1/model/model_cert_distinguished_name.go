package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertDistinguishedName struct {

	// 证书通用名称（CN）。
	CommonName string `json:"common_name"`

	// 国家编码，需符合正则\"**[A-Za-z]{2}**\"。若不传入，则默认继承父CA对应的值。
	Country *string `json:"country,omitempty"`

	// 省市名称。若不传入，则默认继承父CA对应的值。
	State *string `json:"state,omitempty"`

	// 地区名称。若不传入，则默认继承父CA对应的值。
	Locality *string `json:"locality,omitempty"`

	// 组织名称。若不传入，则默认继承父CA对应的值。
	Organization *string `json:"organization,omitempty"`

	// 组织单元名称。若不传入，则默认继承父CA对应的值。
	OrganizationalUnit *string `json:"organizational_unit,omitempty"`
}

func (o CertDistinguishedName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertDistinguishedName struct{}"
	}

	return strings.Join([]string{"CertDistinguishedName", string(data)}, " ")
}
