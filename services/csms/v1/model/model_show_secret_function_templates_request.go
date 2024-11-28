package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecretFunctionTemplatesRequest Request Object
type ShowSecretFunctionTemplatesRequest struct {

	// 凭据类型。
	SecretType ShowSecretFunctionTemplatesRequestSecretType `json:"secret_type"`

	// 凭据轮转账号类型。 - SingleUser：单用户模式轮转 - MultiUser：双用户模式轮转
	SecretSubType ShowSecretFunctionTemplatesRequestSecretSubType `json:"secret_sub_type"`

	// 数据库类型。凭据类型为RDS-FG时为必填参数，可传入mysql、postgresql、sqlserver。其余凭据类型不支持。
	Engine *ShowSecretFunctionTemplatesRequestEngine `json:"engine,omitempty"`
}

func (o ShowSecretFunctionTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecretFunctionTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ShowSecretFunctionTemplatesRequest", string(data)}, " ")
}

type ShowSecretFunctionTemplatesRequestSecretType struct {
	value string
}

type ShowSecretFunctionTemplatesRequestSecretTypeEnum struct {
	RDS_FG      ShowSecretFunctionTemplatesRequestSecretType
	GAUSS_DB_FG ShowSecretFunctionTemplatesRequestSecretType
}

func GetShowSecretFunctionTemplatesRequestSecretTypeEnum() ShowSecretFunctionTemplatesRequestSecretTypeEnum {
	return ShowSecretFunctionTemplatesRequestSecretTypeEnum{
		RDS_FG: ShowSecretFunctionTemplatesRequestSecretType{
			value: "RDS-FG",
		},
		GAUSS_DB_FG: ShowSecretFunctionTemplatesRequestSecretType{
			value: "GaussDB-FG",
		},
	}
}

func (c ShowSecretFunctionTemplatesRequestSecretType) Value() string {
	return c.value
}

func (c ShowSecretFunctionTemplatesRequestSecretType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecretFunctionTemplatesRequestSecretType) UnmarshalJSON(b []byte) error {
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

type ShowSecretFunctionTemplatesRequestSecretSubType struct {
	value string
}

type ShowSecretFunctionTemplatesRequestSecretSubTypeEnum struct {
	SINGLE_USER ShowSecretFunctionTemplatesRequestSecretSubType
	MULTI_USER  ShowSecretFunctionTemplatesRequestSecretSubType
}

func GetShowSecretFunctionTemplatesRequestSecretSubTypeEnum() ShowSecretFunctionTemplatesRequestSecretSubTypeEnum {
	return ShowSecretFunctionTemplatesRequestSecretSubTypeEnum{
		SINGLE_USER: ShowSecretFunctionTemplatesRequestSecretSubType{
			value: "SingleUser",
		},
		MULTI_USER: ShowSecretFunctionTemplatesRequestSecretSubType{
			value: "MultiUser",
		},
	}
}

func (c ShowSecretFunctionTemplatesRequestSecretSubType) Value() string {
	return c.value
}

func (c ShowSecretFunctionTemplatesRequestSecretSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecretFunctionTemplatesRequestSecretSubType) UnmarshalJSON(b []byte) error {
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

type ShowSecretFunctionTemplatesRequestEngine struct {
	value string
}

type ShowSecretFunctionTemplatesRequestEngineEnum struct {
	MYSQL      ShowSecretFunctionTemplatesRequestEngine
	POSTGRESQL ShowSecretFunctionTemplatesRequestEngine
	SQLSERVER  ShowSecretFunctionTemplatesRequestEngine
}

func GetShowSecretFunctionTemplatesRequestEngineEnum() ShowSecretFunctionTemplatesRequestEngineEnum {
	return ShowSecretFunctionTemplatesRequestEngineEnum{
		MYSQL: ShowSecretFunctionTemplatesRequestEngine{
			value: "mysql",
		},
		POSTGRESQL: ShowSecretFunctionTemplatesRequestEngine{
			value: "postgresql",
		},
		SQLSERVER: ShowSecretFunctionTemplatesRequestEngine{
			value: "sqlserver",
		},
	}
}

func (c ShowSecretFunctionTemplatesRequestEngine) Value() string {
	return c.value
}

func (c ShowSecretFunctionTemplatesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecretFunctionTemplatesRequestEngine) UnmarshalJSON(b []byte) error {
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
