package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnvParam struct {

	// 环境关联组件id；id长度不能超过36位，由大小写字母、数字组成。创建环境必传，修改环境时非必选
	ComponentId string `json:"component_id"`

	// 描述：最大255字符
	Description *string `json:"description,omitempty"`

	// 显示名：字符集长度2-64，仅支持字符集：中文字符、英文字母、数字、下划线、中划线、点
	EnvName string `json:"env_name"`

	// 环境类型，取值：DEV、TEST、PRE、ONLINE，不区分大小写
	EnvType EnvParamEnvType `json:"env_type"`

	// OS类型，取值：LINUX、WINDOWS。创建环境必传，不可修改
	OsType EnvParamOsType `json:"os_type"`

	// 环境关联region。创建环境必传，不可修改
	Region *string `json:"region,omitempty"`

	// 注册类型，取值：API、SERVICE_DISCOVERY、CONSOLE，默认值：API
	RegisterType *EnvParamRegisterType `json:"register_type,omitempty"`
}

func (o EnvParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvParam struct{}"
	}

	return strings.Join([]string{"EnvParam", string(data)}, " ")
}

type EnvParamEnvType struct {
	value string
}

type EnvParamEnvTypeEnum struct {
	DEV    EnvParamEnvType
	TEST   EnvParamEnvType
	PRE    EnvParamEnvType
	ONLINE EnvParamEnvType
}

func GetEnvParamEnvTypeEnum() EnvParamEnvTypeEnum {
	return EnvParamEnvTypeEnum{
		DEV: EnvParamEnvType{
			value: "DEV",
		},
		TEST: EnvParamEnvType{
			value: "TEST",
		},
		PRE: EnvParamEnvType{
			value: "PRE",
		},
		ONLINE: EnvParamEnvType{
			value: "ONLINE",
		},
	}
}

func (c EnvParamEnvType) Value() string {
	return c.value
}

func (c EnvParamEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvParamEnvType) UnmarshalJSON(b []byte) error {
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

type EnvParamOsType struct {
	value string
}

type EnvParamOsTypeEnum struct {
	LINUX   EnvParamOsType
	WINDOWS EnvParamOsType
}

func GetEnvParamOsTypeEnum() EnvParamOsTypeEnum {
	return EnvParamOsTypeEnum{
		LINUX: EnvParamOsType{
			value: "LINUX",
		},
		WINDOWS: EnvParamOsType{
			value: "WINDOWS",
		},
	}
}

func (c EnvParamOsType) Value() string {
	return c.value
}

func (c EnvParamOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvParamOsType) UnmarshalJSON(b []byte) error {
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

type EnvParamRegisterType struct {
	value string
}

type EnvParamRegisterTypeEnum struct {
	API               EnvParamRegisterType
	CONSOLE           EnvParamRegisterType
	SERVICE_DISCOVERY EnvParamRegisterType
}

func GetEnvParamRegisterTypeEnum() EnvParamRegisterTypeEnum {
	return EnvParamRegisterTypeEnum{
		API: EnvParamRegisterType{
			value: "API",
		},
		CONSOLE: EnvParamRegisterType{
			value: "CONSOLE",
		},
		SERVICE_DISCOVERY: EnvParamRegisterType{
			value: "SERVICE_DISCOVERY",
		},
	}
}

func (c EnvParamRegisterType) Value() string {
	return c.value
}

func (c EnvParamRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvParamRegisterType) UnmarshalJSON(b []byte) error {
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
