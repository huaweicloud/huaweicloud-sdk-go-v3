package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateLdapConfigRequestBody ldap配置信息
type UpdateLdapConfigRequestBody struct {

	// ldap服务器的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
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

	// ldap备节点的url，固定格式为 ldap://{ip_address}:{port_number} 或 ldaps://{ip_address}:{port_number}，例如ldap://192.168.xx.xx:60000
	BackupUrl *string `json:"backup_url,omitempty"`

	// ldap的schema，不填写则默认为RFC2307
	Schema *string `json:"schema,omitempty"`

	// ldap搜索的超时时间，单位为秒。不填写则默认为3秒
	SearchTimeout *int32 `json:"search_timeout,omitempty"`

	// 访问ldap服务器失败后是否允许使用本地用户鉴权
	AllowLocalUser *UpdateLdapConfigRequestBodyAllowLocalUser `json:"allow_local_user,omitempty"`
}

func (o UpdateLdapConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLdapConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLdapConfigRequestBody", string(data)}, " ")
}

type UpdateLdapConfigRequestBodyAllowLocalUser struct {
	value string
}

type UpdateLdapConfigRequestBodyAllowLocalUserEnum struct {
	YES UpdateLdapConfigRequestBodyAllowLocalUser
	NO  UpdateLdapConfigRequestBodyAllowLocalUser
}

func GetUpdateLdapConfigRequestBodyAllowLocalUserEnum() UpdateLdapConfigRequestBodyAllowLocalUserEnum {
	return UpdateLdapConfigRequestBodyAllowLocalUserEnum{
		YES: UpdateLdapConfigRequestBodyAllowLocalUser{
			value: "Yes",
		},
		NO: UpdateLdapConfigRequestBodyAllowLocalUser{
			value: "No",
		},
	}
}

func (c UpdateLdapConfigRequestBodyAllowLocalUser) Value() string {
	return c.value
}

func (c UpdateLdapConfigRequestBodyAllowLocalUser) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLdapConfigRequestBodyAllowLocalUser) UnmarshalJSON(b []byte) error {
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
