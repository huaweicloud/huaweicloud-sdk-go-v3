package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据加工信息
type TransformationInfo struct {

	// - 生成加工规则值为contentConditionalFilter - 生成配置规则值为configConditionalFilter
	TransformationType TransformationInfoTransformationType `json:"transformation_type" xml:"transformation_type"`

	// 过滤条件，生成加工规则值为sql条件语句，生成配置规则值为config。长度限制256。
	Value string `json:"value" xml:"value"`
}

func (o TransformationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransformationInfo struct{}"
	}

	return strings.Join([]string{"TransformationInfo", string(data)}, " ")
}

type TransformationInfoTransformationType struct {
	value string
}

type TransformationInfoTransformationTypeEnum struct {
	CONTENT_CONDITIONAL_FILTER TransformationInfoTransformationType
	CONFIG_CONDITIONAL_FILTER  TransformationInfoTransformationType
}

func GetTransformationInfoTransformationTypeEnum() TransformationInfoTransformationTypeEnum {
	return TransformationInfoTransformationTypeEnum{
		CONTENT_CONDITIONAL_FILTER: TransformationInfoTransformationType{
			value: "contentConditionalFilter",
		},
		CONFIG_CONDITIONAL_FILTER: TransformationInfoTransformationType{
			value: "configConditionalFilter",
		},
	}
}

func (c TransformationInfoTransformationType) Value() string {
	return c.value
}

func (c TransformationInfoTransformationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransformationInfoTransformationType) UnmarshalJSON(b []byte) error {
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
