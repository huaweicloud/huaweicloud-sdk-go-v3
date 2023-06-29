package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceDbInfo 源数据库信息。
type SourceDbInfo struct {

	// 连接字符串。
	ConnectionString string `json:"connection_string"`

	// 用户名。
	UserName string `json:"user_name"`

	// 用户密码。
	Password string `json:"password"`

	// 数据库类型。
	SourceDbType SourceDbInfoSourceDbType `json:"source_db_type"`

	// 数据库版本。 （注意：该字段的值是数据库类型source_db_type对应的版本，当前支持以下组合： ORACLE-11g； ORACLE-12c； ORACLE-18c； ORACLE-19c。）
	SourceDbVersion SourceDbInfoSourceDbVersion `json:"source_db_version"`
}

func (o SourceDbInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceDbInfo struct{}"
	}

	return strings.Join([]string{"SourceDbInfo", string(data)}, " ")
}

type SourceDbInfoSourceDbType struct {
	value string
}

type SourceDbInfoSourceDbTypeEnum struct {
	ORACLE SourceDbInfoSourceDbType
}

func GetSourceDbInfoSourceDbTypeEnum() SourceDbInfoSourceDbTypeEnum {
	return SourceDbInfoSourceDbTypeEnum{
		ORACLE: SourceDbInfoSourceDbType{
			value: "ORACLE",
		},
	}
}

func (c SourceDbInfoSourceDbType) Value() string {
	return c.value
}

func (c SourceDbInfoSourceDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceDbInfoSourceDbType) UnmarshalJSON(b []byte) error {
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

type SourceDbInfoSourceDbVersion struct {
	value string
}

type SourceDbInfoSourceDbVersionEnum struct {
	E_11G SourceDbInfoSourceDbVersion
	E_12C SourceDbInfoSourceDbVersion
	E_18C SourceDbInfoSourceDbVersion
	E_19C SourceDbInfoSourceDbVersion
}

func GetSourceDbInfoSourceDbVersionEnum() SourceDbInfoSourceDbVersionEnum {
	return SourceDbInfoSourceDbVersionEnum{
		E_11G: SourceDbInfoSourceDbVersion{
			value: "11g",
		},
		E_12C: SourceDbInfoSourceDbVersion{
			value: "12c",
		},
		E_18C: SourceDbInfoSourceDbVersion{
			value: "18c",
		},
		E_19C: SourceDbInfoSourceDbVersion{
			value: "19c",
		},
	}
}

func (c SourceDbInfoSourceDbVersion) Value() string {
	return c.value
}

func (c SourceDbInfoSourceDbVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceDbInfoSourceDbVersion) UnmarshalJSON(b []byte) error {
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
