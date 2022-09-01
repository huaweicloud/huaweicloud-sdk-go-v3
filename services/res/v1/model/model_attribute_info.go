package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 综合排序信息
type AttributeInfo struct {

	// 属性匹配对。
	RankFeaturePairs *[]RankFeaturePair `json:"rank_feature_pairs,omitempty" xml:"rank_feature_pairs"`

	// 属性权重。
	NumericalAttrs *[]NumericalAttr `json:"numerical_attrs,omitempty" xml:"numerical_attrs"`

	// 统计方式： - ORDER，顺序 - ABS，绝对值
	NumStatisticsType *AttributeInfoNumStatisticsType `json:"num_statistics_type,omitempty" xml:"num_statistics_type"`
}

func (o AttributeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributeInfo struct{}"
	}

	return strings.Join([]string{"AttributeInfo", string(data)}, " ")
}

type AttributeInfoNumStatisticsType struct {
	value string
}

type AttributeInfoNumStatisticsTypeEnum struct {
	ORDER AttributeInfoNumStatisticsType
	ABS   AttributeInfoNumStatisticsType
}

func GetAttributeInfoNumStatisticsTypeEnum() AttributeInfoNumStatisticsTypeEnum {
	return AttributeInfoNumStatisticsTypeEnum{
		ORDER: AttributeInfoNumStatisticsType{
			value: "ORDER",
		},
		ABS: AttributeInfoNumStatisticsType{
			value: "ABS",
		},
	}
}

func (c AttributeInfoNumStatisticsType) Value() string {
	return c.value
}

func (c AttributeInfoNumStatisticsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttributeInfoNumStatisticsType) UnmarshalJSON(b []byte) error {
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
