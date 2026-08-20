package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LabelParam 标签创建/编辑参数
type LabelParam struct {

	// 标签所属工作项类型，对应工作项的type字段，枚举类型。不推荐使用此参数，建议使用category_types参数。
	LabelType *LabelParamLabelType `json:"label_type,omitempty"`

	// 标签颜色，作为更新参数时非必填。
	Color string `json:"color"`

	// 标签标题。 1~30个字符。
	Title string `json:"title"`

	// 标签所属工作项类型编码。
	CategoryTypes []string `json:"category_types"`
}

func (o LabelParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelParam struct{}"
	}

	return strings.Join([]string{"LabelParam", string(data)}, " ")
}

type LabelParamLabelType struct {
	value string
}

type LabelParamLabelTypeEnum struct {
	FEATURE         LabelParamLabelType
	RAW_REQUIREMENT LabelParamLabelType
	REQUIREMENT     LabelParamLabelType
	TASK            LabelParamLabelType
	BUG             LabelParamLabelType
}

func GetLabelParamLabelTypeEnum() LabelParamLabelTypeEnum {
	return LabelParamLabelTypeEnum{
		FEATURE: LabelParamLabelType{
			value: "feature",
		},
		RAW_REQUIREMENT: LabelParamLabelType{
			value: "raw requirement",
		},
		REQUIREMENT: LabelParamLabelType{
			value: "requirement",
		},
		TASK: LabelParamLabelType{
			value: "task",
		},
		BUG: LabelParamLabelType{
			value: "bug",
		},
	}
}

func (c LabelParamLabelType) Value() string {
	return c.value
}

func (c LabelParamLabelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LabelParamLabelType) UnmarshalJSON(b []byte) error {
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
