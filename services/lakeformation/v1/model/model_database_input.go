package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatabaseInput 数据库信息
type DatabaseInput struct {

	// 数据库名称。只能包含中文、字母、数字和下划线，且长度为1~128个字符。
	DatabaseName string `json:"database_name"`

	// 数据库所有者。长度为0~128个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerType *DatabaseInputOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务分为一期和二期，一期响应Body无该参数。
	OwnerAuthSourceType *DatabaseInputOwnerAuthSourceType `json:"owner_auth_source_type,omitempty"`

	// 数据库描述信息。由用户创建数据库时输入，最大长度为4000个字符。
	Description *string `json:"description,omitempty"`

	// 数据库路径地址。例如obs://location/uri/
	Location *string `json:"location,omitempty"`

	// 标签信息
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表路径列表
	TableLocationList *[]string `json:"table_location_list,omitempty"`

	// 函数路径列表
	FunctionLocationList *[]string `json:"function_location_list,omitempty"`
}

func (o DatabaseInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseInput struct{}"
	}

	return strings.Join([]string{"DatabaseInput", string(data)}, " ")
}

type DatabaseInputOwnerType struct {
	value string
}

type DatabaseInputOwnerTypeEnum struct {
	USER  DatabaseInputOwnerType
	ROLE  DatabaseInputOwnerType
	GROUP DatabaseInputOwnerType
}

func GetDatabaseInputOwnerTypeEnum() DatabaseInputOwnerTypeEnum {
	return DatabaseInputOwnerTypeEnum{
		USER: DatabaseInputOwnerType{
			value: "USER",
		},
		ROLE: DatabaseInputOwnerType{
			value: "ROLE",
		},
		GROUP: DatabaseInputOwnerType{
			value: "GROUP",
		},
	}
}

func (c DatabaseInputOwnerType) Value() string {
	return c.value
}

func (c DatabaseInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInputOwnerType) UnmarshalJSON(b []byte) error {
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

type DatabaseInputOwnerAuthSourceType struct {
	value string
}

type DatabaseInputOwnerAuthSourceTypeEnum struct {
	IAM         DatabaseInputOwnerAuthSourceType
	SAML        DatabaseInputOwnerAuthSourceType
	LDAP        DatabaseInputOwnerAuthSourceType
	LOCAL       DatabaseInputOwnerAuthSourceType
	AGENTTENANT DatabaseInputOwnerAuthSourceType
	OTHER       DatabaseInputOwnerAuthSourceType
}

func GetDatabaseInputOwnerAuthSourceTypeEnum() DatabaseInputOwnerAuthSourceTypeEnum {
	return DatabaseInputOwnerAuthSourceTypeEnum{
		IAM: DatabaseInputOwnerAuthSourceType{
			value: "IAM",
		},
		SAML: DatabaseInputOwnerAuthSourceType{
			value: "SAML",
		},
		LDAP: DatabaseInputOwnerAuthSourceType{
			value: "LDAP",
		},
		LOCAL: DatabaseInputOwnerAuthSourceType{
			value: "LOCAL",
		},
		AGENTTENANT: DatabaseInputOwnerAuthSourceType{
			value: "AGENTTENANT",
		},
		OTHER: DatabaseInputOwnerAuthSourceType{
			value: "OTHER",
		},
	}
}

func (c DatabaseInputOwnerAuthSourceType) Value() string {
	return c.value
}

func (c DatabaseInputOwnerAuthSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatabaseInputOwnerAuthSourceType) UnmarshalJSON(b []byte) error {
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
