package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AppConfigModifyRequestV2 struct {
	// 应用配置类型： - variable：模板变量 - password：密码 - certificate：证书

	ConfigType AppConfigModifyRequestV2ConfigType `json:"config_type"`
	// 应用配置值： - config_type = variable：config_value为模板变量的值 - config_type = password：config_value为密码值 - config_type = certificate：config_value需要包含证书public_key（必填），私钥private_key（必填）和密码passphrase（非必填），格式如：\"{\\\\\"public_key\\\\\": \\\"\\,\\\\\"private_key\\\\\":\\\\\"\\\\\",\\\\\"passphrase\\\\\":\\\\\"\\\\\"}\"

	ConfigValue *string `json:"config_value,omitempty"`
	// 应用配置描述

	Description *string `json:"description,omitempty"`
}

func (o AppConfigModifyRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppConfigModifyRequestV2 struct{}"
	}

	return strings.Join([]string{"AppConfigModifyRequestV2", string(data)}, " ")
}

type AppConfigModifyRequestV2ConfigType struct {
	value string
}

type AppConfigModifyRequestV2ConfigTypeEnum struct {
	VARIABLE    AppConfigModifyRequestV2ConfigType
	PASSWORD    AppConfigModifyRequestV2ConfigType
	CERTIFICATE AppConfigModifyRequestV2ConfigType
}

func GetAppConfigModifyRequestV2ConfigTypeEnum() AppConfigModifyRequestV2ConfigTypeEnum {
	return AppConfigModifyRequestV2ConfigTypeEnum{
		VARIABLE: AppConfigModifyRequestV2ConfigType{
			value: "variable",
		},
		PASSWORD: AppConfigModifyRequestV2ConfigType{
			value: "password",
		},
		CERTIFICATE: AppConfigModifyRequestV2ConfigType{
			value: "certificate",
		},
	}
}

func (c AppConfigModifyRequestV2ConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppConfigModifyRequestV2ConfigType) UnmarshalJSON(b []byte) error {
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
