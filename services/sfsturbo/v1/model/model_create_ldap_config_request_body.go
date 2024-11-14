package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateLdapConfigRequestBody ldap配置信息
type CreateLdapConfigRequestBody struct {

	// ldap服务器的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
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

	// ldap备节点的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
	BackupUrl *string `json:"backup_url,omitempty"`

	// ldap的schema，不填写则默认为RFC2307
	Schema *string `json:"schema,omitempty"`

	// ldap搜索的超时时间，单位为秒。不填写则默认为3秒
	SearchTimeout *int32 `json:"search_timeout,omitempty"`

	// 访问ldap服务器失败后是否允许使用本地用户鉴权
	AllowLocalUser *CreateLdapConfigRequestBodyAllowLocalUser `json:"allow_local_user,omitempty"`
}

func (o CreateLdapConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLdapConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLdapConfigRequestBody", string(data)}, " ")
}

type CreateLdapConfigRequestBodyAllowLocalUser struct {
	value string
}

type CreateLdapConfigRequestBodyAllowLocalUserEnum struct {
	YES CreateLdapConfigRequestBodyAllowLocalUser
	NO  CreateLdapConfigRequestBodyAllowLocalUser
}

func GetCreateLdapConfigRequestBodyAllowLocalUserEnum() CreateLdapConfigRequestBodyAllowLocalUserEnum {
	return CreateLdapConfigRequestBodyAllowLocalUserEnum{
		YES: CreateLdapConfigRequestBodyAllowLocalUser{
			value: "Yes",
		},
		NO: CreateLdapConfigRequestBodyAllowLocalUser{
			value: "No",
		},
	}
}

func (c CreateLdapConfigRequestBodyAllowLocalUser) Value() string {
	return c.value
}

func (c CreateLdapConfigRequestBodyAllowLocalUser) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLdapConfigRequestBodyAllowLocalUser) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
