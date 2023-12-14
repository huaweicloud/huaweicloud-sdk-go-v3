package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenTsdbMetric struct {

	// - Constant表示metric为常量value的值。 - String表示metric为通道内用户数据对应JSON属性的取值，且该JOSN属性的取值为String。
	Type OpenTsdbMetricType `json:"type"`

	// 常量或通道内用户数据的JSON属性名称。  取值范围：1～32，只能包含英文字母、数字和点。
	Value string `json:"value"`
}

func (o OpenTsdbMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenTsdbMetric struct{}"
	}

	return strings.Join([]string{"OpenTsdbMetric", string(data)}, " ")
}

type OpenTsdbMetricType struct {
	value string
}

type OpenTsdbMetricTypeEnum struct {
	CONSTANT OpenTsdbMetricType
	STRING   OpenTsdbMetricType
}

func GetOpenTsdbMetricTypeEnum() OpenTsdbMetricTypeEnum {
	return OpenTsdbMetricTypeEnum{
		CONSTANT: OpenTsdbMetricType{
			value: "Constant",
		},
		STRING: OpenTsdbMetricType{
			value: "String",
		},
	}
}

func (c OpenTsdbMetricType) Value() string {
	return c.value
}

func (c OpenTsdbMetricType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenTsdbMetricType) UnmarshalJSON(b []byte) error {
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
