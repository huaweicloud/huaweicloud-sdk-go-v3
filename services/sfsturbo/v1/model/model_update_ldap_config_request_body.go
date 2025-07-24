package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLdapConfigRequestBody LDAP配置信息
type UpdateLdapConfigRequestBody struct {

	// LDAP服务器的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
	Url *string `json:"url,omitempty"`

	// 数据库中的域
	BaseDn *string `json:"base_dn,omitempty"`

	// 用户区别名
	UserDn *string `json:"user_dn,omitempty"`

	// LDAP认证密码
	Password *string `json:"password,omitempty"`

	// vpc的id
	VpcId *string `json:"vpc_id,omitempty"`

	// 过滤条件。保留字段，暂不支持
	FilterCondition *string `json:"filter_condition,omitempty"`

	// LDAP备节点的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
	BackupUrl *string `json:"backup_url,omitempty"`

	// LDAP的schema，不填写则默认为RFC2307
	Schema *string `json:"schema,omitempty"`

	// LDAP搜索的超时时间，单位为秒。不填写则默认为3秒
	SearchTimeout *int32 `json:"search_timeout,omitempty"`
}

func (o UpdateLdapConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLdapConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLdapConfigRequestBody", string(data)}, " ")
}
