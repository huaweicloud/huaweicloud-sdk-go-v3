package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RelationType 关系类型，只读。 枚举值：   - ONE: 表示每条子（父）逻辑实体数据在父（子）逻辑实体中有且只有一条数据与之对应。   - ZERO_OR_ONE: 表示每条子（父）逻辑实体数据在父（子）逻辑实体中最多有一条数据与之对应。   - ZERO_OR_N: 表示每条子（父）逻辑实体数据在父（子）逻辑实体中可能有多条数据与之对应。   - ONE_OR_N: 表示每条子（父）逻辑实体数据在父（子）逻辑实体中至少有一条数据与之对应。
type RelationType struct {
	value string
}

type RelationTypeEnum struct {
	ONE         RelationType
	ZERO_OR_ONE RelationType
	ZERO_OR_N   RelationType
	ONE_OR_N    RelationType
}

func GetRelationTypeEnum() RelationTypeEnum {
	return RelationTypeEnum{
		ONE: RelationType{
			value: "ONE",
		},
		ZERO_OR_ONE: RelationType{
			value: "ZERO_OR_ONE",
		},
		ZERO_OR_N: RelationType{
			value: "ZERO_OR_N",
		},
		ONE_OR_N: RelationType{
			value: "ONE_OR_N",
		},
	}
}

func (c RelationType) Value() string {
	return c.value
}

func (c RelationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RelationType) UnmarshalJSON(b []byte) error {
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
