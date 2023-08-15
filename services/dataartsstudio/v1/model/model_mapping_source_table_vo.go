package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MappingSourceTableVo struct {

	// 表id
	Table1Id int64 `json:"table1_id"`

	// 表id
	Table2Id *int64 `json:"table2_id,omitempty"`

	// 表1名称
	Table1Name string `json:"table1_name"`

	// 表2名称
	Table2Name *string `json:"table2_name,omitempty"`

	// 关联类型，左外连接，右外连接，内连接，全连接
	JoinType MappingSourceTableVoJoinType `json:"join_type"`

	// on条件
	JoinFields []MappingJoinFieldVo `json:"join_fields"`
}

func (o MappingSourceTableVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingSourceTableVo struct{}"
	}

	return strings.Join([]string{"MappingSourceTableVo", string(data)}, " ")
}

type MappingSourceTableVoJoinType struct {
	value string
}

type MappingSourceTableVoJoinTypeEnum struct {
	LEFT  MappingSourceTableVoJoinType
	RIGHT MappingSourceTableVoJoinType
	INNER MappingSourceTableVoJoinType
	FULL  MappingSourceTableVoJoinType
}

func GetMappingSourceTableVoJoinTypeEnum() MappingSourceTableVoJoinTypeEnum {
	return MappingSourceTableVoJoinTypeEnum{
		LEFT: MappingSourceTableVoJoinType{
			value: "LEFT",
		},
		RIGHT: MappingSourceTableVoJoinType{
			value: "RIGHT",
		},
		INNER: MappingSourceTableVoJoinType{
			value: "INNER",
		},
		FULL: MappingSourceTableVoJoinType{
			value: "FULL",
		},
	}
}

func (c MappingSourceTableVoJoinType) Value() string {
	return c.value
}

func (c MappingSourceTableVoJoinType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MappingSourceTableVoJoinType) UnmarshalJSON(b []byte) error {
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
