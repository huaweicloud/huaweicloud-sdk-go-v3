package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 表列限制条件输入模型
type TableConstraintsInput struct {

	// 限制类型
	ConstraintType *TableConstraintsInputConstraintType `json:"constraint_type,omitempty"`

	// 外键列表
	ForeignKeys *[]ForeignKeyInput `json:"foreign_keys,omitempty"`

	// 主键列表
	PrimaryKeys *[]PrimaryKeyInput `json:"primary_keys,omitempty"`

	// 非空限制列表
	NotNullConstraints *[]NotNullConstraintInput `json:"not_null_constraints,omitempty"`

	// 检查限制列表
	CheckConstraints *[]CheckConstraintInput `json:"check_constraints,omitempty"`

	// 唯一值限制列表
	UniqueConstraints *[]UniqueConstraintInput `json:"unique_constraints,omitempty"`

	// 默认限制列表
	DefaultConstraints *[]DefaultConstraintInput `json:"default_constraints,omitempty"`
}

func (o TableConstraintsInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableConstraintsInput struct{}"
	}

	return strings.Join([]string{"TableConstraintsInput", string(data)}, " ")
}

type TableConstraintsInputConstraintType struct {
	value string
}

type TableConstraintsInputConstraintTypeEnum struct {
	CHECK_CSTR    TableConstraintsInputConstraintType
	DEFAULT_CSTR  TableConstraintsInputConstraintType
	NOT_NULL_CSTR TableConstraintsInputConstraintType
	UNIQUE_CSTR   TableConstraintsInputConstraintType
	PRIMARY_KEY   TableConstraintsInputConstraintType
	FOREIGN_KEY   TableConstraintsInputConstraintType
}

func GetTableConstraintsInputConstraintTypeEnum() TableConstraintsInputConstraintTypeEnum {
	return TableConstraintsInputConstraintTypeEnum{
		CHECK_CSTR: TableConstraintsInputConstraintType{
			value: "CHECK_CSTR",
		},
		DEFAULT_CSTR: TableConstraintsInputConstraintType{
			value: "DEFAULT_CSTR",
		},
		NOT_NULL_CSTR: TableConstraintsInputConstraintType{
			value: "NOT_NULL_CSTR",
		},
		UNIQUE_CSTR: TableConstraintsInputConstraintType{
			value: "UNIQUE_CSTR",
		},
		PRIMARY_KEY: TableConstraintsInputConstraintType{
			value: "PRIMARY_KEY",
		},
		FOREIGN_KEY: TableConstraintsInputConstraintType{
			value: "FOREIGN_KEY",
		},
	}
}

func (c TableConstraintsInputConstraintType) Value() string {
	return c.value
}

func (c TableConstraintsInputConstraintType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableConstraintsInputConstraintType) UnmarshalJSON(b []byte) error {
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
