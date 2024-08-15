package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicScriptPropertiesModel 公共脚本附加属性
type PublicScriptPropertiesModel struct {

	// 风险等级 LOW:低风险 MEDIUM:中风险 HIGH:高风险
	RiskLevel PublicScriptPropertiesModelRiskLevel `json:"risk_level"`

	// 脚本版本号
	Version string `json:"version"`
}

func (o PublicScriptPropertiesModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicScriptPropertiesModel struct{}"
	}

	return strings.Join([]string{"PublicScriptPropertiesModel", string(data)}, " ")
}

type PublicScriptPropertiesModelRiskLevel struct {
	value string
}

type PublicScriptPropertiesModelRiskLevelEnum struct {
	LOW    PublicScriptPropertiesModelRiskLevel
	MEDIUM PublicScriptPropertiesModelRiskLevel
	HIGH   PublicScriptPropertiesModelRiskLevel
}

func GetPublicScriptPropertiesModelRiskLevelEnum() PublicScriptPropertiesModelRiskLevelEnum {
	return PublicScriptPropertiesModelRiskLevelEnum{
		LOW: PublicScriptPropertiesModelRiskLevel{
			value: "LOW",
		},
		MEDIUM: PublicScriptPropertiesModelRiskLevel{
			value: "MEDIUM",
		},
		HIGH: PublicScriptPropertiesModelRiskLevel{
			value: "HIGH",
		},
	}
}

func (c PublicScriptPropertiesModelRiskLevel) Value() string {
	return c.value
}

func (c PublicScriptPropertiesModelRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicScriptPropertiesModelRiskLevel) UnmarshalJSON(b []byte) error {
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
