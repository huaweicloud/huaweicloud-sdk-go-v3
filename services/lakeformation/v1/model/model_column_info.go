package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ColumnInfo column input when grant policy
type ColumnInfo struct {

	// 列名
	ColumnName []string `json:"column_name"`

	// 是否排除：Include包含,Exclude排除
	Filter ColumnInfoFilter `json:"filter"`
}

func (o ColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnInfo struct{}"
	}

	return strings.Join([]string{"ColumnInfo", string(data)}, " ")
}

type ColumnInfoFilter struct {
	value string
}

type ColumnInfoFilterEnum struct {
	INCLUDE ColumnInfoFilter
	EXCLUDE ColumnInfoFilter
}

func GetColumnInfoFilterEnum() ColumnInfoFilterEnum {
	return ColumnInfoFilterEnum{
		INCLUDE: ColumnInfoFilter{
			value: "Include",
		},
		EXCLUDE: ColumnInfoFilter{
			value: "Exclude",
		},
	}
}

func (c ColumnInfoFilter) Value() string {
	return c.value
}

func (c ColumnInfoFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ColumnInfoFilter) UnmarshalJSON(b []byte) error {
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
