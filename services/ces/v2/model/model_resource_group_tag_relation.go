package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceGroupTagRelation 标签的匹配规则
type ResourceGroupTagRelation struct {

	// TMS标签键规范。
	Key string `json:"key"`

	// tag操作符，含义是标签key与value的关系。   include表示包含   prefix表示前缀   suffix表示后缀   notInclude表示不包含   equal表示相等   当operator为equal，value为空字符串时表示为全部   all表示全部
	Operator *ResourceGroupTagRelationOperator `json:"operator,omitempty"`

	// TMS标签值规范。
	Value *string `json:"value,omitempty"`
}

func (o ResourceGroupTagRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupTagRelation struct{}"
	}

	return strings.Join([]string{"ResourceGroupTagRelation", string(data)}, " ")
}

type ResourceGroupTagRelationOperator struct {
	value string
}

type ResourceGroupTagRelationOperatorEnum struct {
	INCLUDE     ResourceGroupTagRelationOperator
	PREFIX      ResourceGroupTagRelationOperator
	SUFFIX      ResourceGroupTagRelationOperator
	NOT_INCLUDE ResourceGroupTagRelationOperator
	EQUAL       ResourceGroupTagRelationOperator
	ALL         ResourceGroupTagRelationOperator
}

func GetResourceGroupTagRelationOperatorEnum() ResourceGroupTagRelationOperatorEnum {
	return ResourceGroupTagRelationOperatorEnum{
		INCLUDE: ResourceGroupTagRelationOperator{
			value: "include",
		},
		PREFIX: ResourceGroupTagRelationOperator{
			value: "prefix",
		},
		SUFFIX: ResourceGroupTagRelationOperator{
			value: "suffix",
		},
		NOT_INCLUDE: ResourceGroupTagRelationOperator{
			value: "notInclude",
		},
		EQUAL: ResourceGroupTagRelationOperator{
			value: "equal",
		},
		ALL: ResourceGroupTagRelationOperator{
			value: "all",
		},
	}
}

func (c ResourceGroupTagRelationOperator) Value() string {
	return c.value
}

func (c ResourceGroupTagRelationOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceGroupTagRelationOperator) UnmarshalJSON(b []byte) error {
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
