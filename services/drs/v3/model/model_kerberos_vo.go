package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// kerberos认证需要的信息
type KerberosVo struct {

	// krb5配置文件
	Krb5ConfFile *string `json:"krb5_conf_file,omitempty" xml:"krb5_conf_file"`

	// key文件
	KeyTabFile *string `json:"key_tab_file,omitempty" xml:"key_tab_file"`

	// 域名
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// Kerberos用户对象
	UserPrincipal *string `json:"user_principal,omitempty" xml:"user_principal"`
}

func (o KerberosVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KerberosVo struct{}"
	}

	return strings.Join([]string{"KerberosVo", string(data)}, " ")
}
