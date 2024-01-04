package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLdapConfigRequestBody ldap配置信息
type UpdateLdapConfigRequestBody struct {

	// ldap服务器的url
	Url *string `json:"url,omitempty"`

	// 数据库中的域
	BaseDn *string `json:"base_dn,omitempty"`

	// 用户区别名
	UserDn *string `json:"user_dn,omitempty"`

	// ldap认证密码
	Password *string `json:"password,omitempty"`

	// vpc的id
	VpcId *string `json:"vpc_id,omitempty"`

	// 过滤条件。保留字段，暂不支持
	FilterCondition *string `json:"filter_condition,omitempty"`
}

func (o UpdateLdapConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLdapConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLdapConfigRequestBody", string(data)}, " ")
}
