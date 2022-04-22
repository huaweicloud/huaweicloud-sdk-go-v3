package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppConfigCreateRequestV2 struct {

	// 应用配置类型： - variable：模板变量 - password：密码 - certificate：证书
	ConfigType AppConfigCreateRequestV2ConfigType `json:"config_type"`

	// 应用配置值： - config_type = variable：config_value为模板变量的值 - config_type = password：config_value为密码值 - config_type = certificate：config_value需要包含证书public_key（必填），私钥private_key（必填）和密码passphrase（非必填），格式如：\"{\\\\\"public_key\\\\\": \\\"\\,\\\\\"private_key\\\\\":\\\\\"\\\\\",\\\\\"passphrase\\\\\":\\\\\"\\\\\"}\"
	ConfigValue *string `json:"config_value,omitempty"`

	// 应用配置描述
	Description *string `json:"description,omitempty"`

	// 应用配置名称。  支持英文大小写字符、数字、下划线、中划线、点和@，且只能以英文字母开头。
	ConfigName *string `json:"config_name,omitempty"`
}

func (o AppConfigCreateRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppConfigCreateRequestV2 struct{}"
	}

	return strings.Join([]string{"AppConfigCreateRequestV2", string(data)}, " ")
}

type AppConfigCreateRequestV2ConfigType struct {
	value string
}

type AppConfigCreateRequestV2ConfigTypeEnum struct {
	VARIABLE    AppConfigCreateRequestV2ConfigType
	PASSWORD    AppConfigCreateRequestV2ConfigType
	CERTIFICATE AppConfigCreateRequestV2ConfigType
}

func GetAppConfigCreateRequestV2ConfigTypeEnum() AppConfigCreateRequestV2ConfigTypeEnum {
	return AppConfigCreateRequestV2ConfigTypeEnum{
		VARIABLE: AppConfigCreateRequestV2ConfigType{
			value: "variable",
		},
		PASSWORD: AppConfigCreateRequestV2ConfigType{
			value: "password",
		},
		CERTIFICATE: AppConfigCreateRequestV2ConfigType{
			value: "certificate",
		},
	}
}

func (c AppConfigCreateRequestV2ConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppConfigCreateRequestV2ConfigType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
