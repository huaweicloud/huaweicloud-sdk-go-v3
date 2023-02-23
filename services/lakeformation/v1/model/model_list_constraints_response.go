package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListConstraintsResponse struct {

	// 限制类型
	ConstraintType *ListConstraintsResponseConstraintType `json:"constraint_type,omitempty"`

	// 外键列表
	ForeignKeys *[]ForeignKey `json:"foreign_keys,omitempty"`

	// 主键列表
	PrimaryKeys *[]PrimaryKey `json:"primary_keys,omitempty"`

	// 非空限制列表
	NotNullConstraints *[]NotNullConstraint `json:"not_null_constraints,omitempty"`

	// 检查限制列表
	CheckConstraints *[]CheckConstraint `json:"check_constraints,omitempty"`

	// 唯一值限制列表
	UniqueConstraints *[]UniqueConstraint `json:"unique_constraints,omitempty"`

	// 默认限制列表
	DefaultConstraints *[]DefaultConstraint `json:"default_constraints,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ListConstraintsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConstraintsResponse struct{}"
	}

	return strings.Join([]string{"ListConstraintsResponse", string(data)}, " ")
}

type ListConstraintsResponseConstraintType struct {
	value string
}

type ListConstraintsResponseConstraintTypeEnum struct {
	CHECK_CSTR    ListConstraintsResponseConstraintType
	DEFAULT_CSTR  ListConstraintsResponseConstraintType
	NOT_NULL_CSTR ListConstraintsResponseConstraintType
	UNIQUE_CSTR   ListConstraintsResponseConstraintType
	PRIMARY_KEY   ListConstraintsResponseConstraintType
	FOREIGN_KEY   ListConstraintsResponseConstraintType
}

func GetListConstraintsResponseConstraintTypeEnum() ListConstraintsResponseConstraintTypeEnum {
	return ListConstraintsResponseConstraintTypeEnum{
		CHECK_CSTR: ListConstraintsResponseConstraintType{
			value: "CHECK_CSTR",
		},
		DEFAULT_CSTR: ListConstraintsResponseConstraintType{
			value: "DEFAULT_CSTR",
		},
		NOT_NULL_CSTR: ListConstraintsResponseConstraintType{
			value: "NOT_NULL_CSTR",
		},
		UNIQUE_CSTR: ListConstraintsResponseConstraintType{
			value: "UNIQUE_CSTR",
		},
		PRIMARY_KEY: ListConstraintsResponseConstraintType{
			value: "PRIMARY_KEY",
		},
		FOREIGN_KEY: ListConstraintsResponseConstraintType{
			value: "FOREIGN_KEY",
		},
	}
}

func (c ListConstraintsResponseConstraintType) Value() string {
	return c.value
}

func (c ListConstraintsResponseConstraintType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConstraintsResponseConstraintType) UnmarshalJSON(b []byte) error {
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
