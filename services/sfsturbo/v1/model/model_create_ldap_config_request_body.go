package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLdapConfigRequestBody ldap配置信息
type CreateLdapConfigRequestBody struct {

	// ldap服务器的url
	Url string `json:"url"`

	// 数据库中的域
	BaseDn string `json:"base_dn"`

	// 用户区别名
	UserDn *string `json:"user_dn,omitempty"`

	// ldap认证密码
	Password *string `json:"password,omitempty"`

	// 一般不涉及。仅在SFSTurbo支持多VPC的场景下，需要指定LDAP服务器可连通的VPC的id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 过滤条件。保留字段，暂不支持
	FilterCondition *string `json:"filter_condition,omitempty"`
}

func (o CreateLdapConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLdapConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLdapConfigRequestBody", string(data)}, " ")
}
