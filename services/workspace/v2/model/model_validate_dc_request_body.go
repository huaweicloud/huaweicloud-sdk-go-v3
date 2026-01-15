package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDcRequestBody 校验域控请求体。
type ValidateDcRequestBody struct {

	// 域名。
	Domain string `json:"domain"`

	// 域控制器IP地址。
	DcIp string `json:"dc_ip"`

	// 域控制器名称。
	DcName string `json:"dc_name"`

	// 域管理员账号。
	DomainAdminAccount *string `json:"domain_admin_account,omitempty"`

	// 域管理员账号密码。
	DomainPassword *string `json:"domain_password,omitempty"`

	// 是否开启LDAPS。
	IsLdapsEnable *bool `json:"is_ldaps_enable,omitempty"`

	// LDAPS用的密钥证书。
	LdapsCertificate *string `json:"ldaps_certificate,omitempty"`
}

func (o ValidateDcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDcRequestBody struct{}"
	}

	return strings.Join([]string{"ValidateDcRequestBody", string(data)}, " ")
}
