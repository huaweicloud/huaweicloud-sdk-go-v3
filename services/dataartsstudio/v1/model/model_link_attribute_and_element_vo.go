package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LinkAttributeAndElementVo 属性ID列表。
type LinkAttributeAndElementVo struct {

	// 属性ID列表，填写String类型替代Long类型。
	Ids []string `json:"ids"`

	// 关联的数据标准的ID，填写String类型替代Long类型。
	StandRowId string `json:"stand_row_id"`

	// 表ID，填写String类型替代Long类型。
	TableId string `json:"table_id"`

	// 表类型，默认是TABLE_MODEL。 枚举值：   - TABLE_MODEL: 关系模型（逻辑模型/物理模型）   - AGGREGATION_LOGIC_TABLE: 汇总表   - FACT_LOGIC_TABLE: 事实表   - DIMENSION: 维度   - DIMENSION_LOGIC_TABLE: 维度表
	BizType LinkAttributeAndElementVoBizType `json:"biz_type"`
}

func (o LinkAttributeAndElementVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkAttributeAndElementVo struct{}"
	}

	return strings.Join([]string{"LinkAttributeAndElementVo", string(data)}, " ")
}

type LinkAttributeAndElementVoBizType struct {
	value string
}

type LinkAttributeAndElementVoBizTypeEnum struct {
	TABLE_MODEL             LinkAttributeAndElementVoBizType
	AGGREGATION_LOGIC_TABLE LinkAttributeAndElementVoBizType
	FACT_LOGIC_TABLE        LinkAttributeAndElementVoBizType
	DIMENSION               LinkAttributeAndElementVoBizType
	DIMENSION_LOGIC_TABLE   LinkAttributeAndElementVoBizType
}

func GetLinkAttributeAndElementVoBizTypeEnum() LinkAttributeAndElementVoBizTypeEnum {
	return LinkAttributeAndElementVoBizTypeEnum{
		TABLE_MODEL: LinkAttributeAndElementVoBizType{
			value: "TABLE_MODEL",
		},
		AGGREGATION_LOGIC_TABLE: LinkAttributeAndElementVoBizType{
			value: "AGGREGATION_LOGIC_TABLE",
		},
		FACT_LOGIC_TABLE: LinkAttributeAndElementVoBizType{
			value: "FACT_LOGIC_TABLE",
		},
		DIMENSION: LinkAttributeAndElementVoBizType{
			value: "DIMENSION",
		},
		DIMENSION_LOGIC_TABLE: LinkAttributeAndElementVoBizType{
			value: "DIMENSION_LOGIC_TABLE",
		},
	}
}

func (c LinkAttributeAndElementVoBizType) Value() string {
	return c.value
}

func (c LinkAttributeAndElementVoBizType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LinkAttributeAndElementVoBizType) UnmarshalJSON(b []byte) error {
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
