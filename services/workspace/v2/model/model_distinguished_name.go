package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DistinguishedName 证书DN。
type DistinguishedName struct {

	// ca名称。
	CommonName string `json:"common_name"`

	// 国家/地区。
	Country string `json:"country"`

	// 省份/州。
	State string `json:"state"`

	// 城市。
	Locality string `json:"locality"`

	// 公司名称。
	Organization string `json:"organization"`

	// 部门名称。
	OrganizationalUnit string `json:"organizational_unit"`
}

func (o DistinguishedName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DistinguishedName struct{}"
	}

	return strings.Join([]string{"DistinguishedName", string(data)}, " ")
}
