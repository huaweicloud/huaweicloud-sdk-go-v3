package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BizAppParam struct {

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用名称.字符集长度2-64，仅支持字符集：中文字符、英文字母、数字、下划线、中划线、点
	DisplayName *string `json:"display_name,omitempty"`

	// 应用关联的企业项目id。企业级用户必传
	EpsId *string `json:"eps_id,omitempty"`

	// 唯一标识.字符集长度2-64，仅支持字符集：英文字母、数字、下划线、中划线、点
	Name string `json:"name"`

	// 前端默认是CONSOLE，不需要传参。rest接口无参数是API，有参数只能是：SERVICE_DISCOVERY
	RegisterType *BizAppParamRegisterType `json:"register_type,omitempty"`
}

func (o BizAppParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizAppParam struct{}"
	}

	return strings.Join([]string{"BizAppParam", string(data)}, " ")
}

type BizAppParamRegisterType struct {
	value string
}

type BizAppParamRegisterTypeEnum struct {
	API                      BizAppParamRegisterType
	CONSOLESERVICE_DISCOVERY BizAppParamRegisterType
}

func GetBizAppParamRegisterTypeEnum() BizAppParamRegisterTypeEnum {
	return BizAppParamRegisterTypeEnum{
		API: BizAppParamRegisterType{
			value: "API",
		},
		CONSOLESERVICE_DISCOVERY: BizAppParamRegisterType{
			value: "CONSOLESERVICE_DISCOVERY",
		},
	}
}

func (c BizAppParamRegisterType) Value() string {
	return c.value
}

func (c BizAppParamRegisterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BizAppParamRegisterType) UnmarshalJSON(b []byte) error {
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
