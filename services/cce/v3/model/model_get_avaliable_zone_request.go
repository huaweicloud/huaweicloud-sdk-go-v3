package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetAvaliableZoneRequest Request Object
type GetAvaliableZoneRequest struct {

	// **参数解释**： 该参数用于按所在区域显示可用区名称 **取值范围**： - zh-cn: 显示中文名称，例如：“可用区1” - en-us: 显示英文名称，例如：“AZ1”
	Locale *GetAvaliableZoneRequestLocale `json:"locale,omitempty"`
}

func (o GetAvaliableZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAvaliableZoneRequest struct{}"
	}

	return strings.Join([]string{"GetAvaliableZoneRequest", string(data)}, " ")
}

type GetAvaliableZoneRequestLocale struct {
	value string
}

type GetAvaliableZoneRequestLocaleEnum struct {
	ZH_CN GetAvaliableZoneRequestLocale
	EN_US GetAvaliableZoneRequestLocale
}

func GetGetAvaliableZoneRequestLocaleEnum() GetAvaliableZoneRequestLocaleEnum {
	return GetAvaliableZoneRequestLocaleEnum{
		ZH_CN: GetAvaliableZoneRequestLocale{
			value: "zh-cn",
		},
		EN_US: GetAvaliableZoneRequestLocale{
			value: "en-us",
		},
	}
}

func (c GetAvaliableZoneRequestLocale) Value() string {
	return c.value
}

func (c GetAvaliableZoneRequestLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetAvaliableZoneRequestLocale) UnmarshalJSON(b []byte) error {
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
