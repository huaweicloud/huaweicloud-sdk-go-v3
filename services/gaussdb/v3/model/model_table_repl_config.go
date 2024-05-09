package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableReplConfig 表同步配置。
type TableReplConfig struct {

	// 表同步类型。include_tables:白名单,exclude_tables:黑名单。
	ReplType *TableReplConfigReplType `json:"repl_type,omitempty"`

	// 表同步范围。all:全量同步，part:部分同步。
	ReplScope *TableReplConfigReplScope `json:"repl_scope,omitempty"`

	// 白名单或黑名单的表范围。
	Tables *[]string `json:"tables,omitempty"`
}

func (o TableReplConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableReplConfig struct{}"
	}

	return strings.Join([]string{"TableReplConfig", string(data)}, " ")
}

type TableReplConfigReplType struct {
	value string
}

type TableReplConfigReplTypeEnum struct {
	INCLUDE_TABLES TableReplConfigReplType
	EXCLUDE_TABLES TableReplConfigReplType
}

func GetTableReplConfigReplTypeEnum() TableReplConfigReplTypeEnum {
	return TableReplConfigReplTypeEnum{
		INCLUDE_TABLES: TableReplConfigReplType{
			value: "include_tables",
		},
		EXCLUDE_TABLES: TableReplConfigReplType{
			value: "exclude_tables",
		},
	}
}

func (c TableReplConfigReplType) Value() string {
	return c.value
}

func (c TableReplConfigReplType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableReplConfigReplType) UnmarshalJSON(b []byte) error {
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

type TableReplConfigReplScope struct {
	value string
}

type TableReplConfigReplScopeEnum struct {
	ALL  TableReplConfigReplScope
	PART TableReplConfigReplScope
}

func GetTableReplConfigReplScopeEnum() TableReplConfigReplScopeEnum {
	return TableReplConfigReplScopeEnum{
		ALL: TableReplConfigReplScope{
			value: "all",
		},
		PART: TableReplConfigReplScope{
			value: "part",
		},
	}
}

func (c TableReplConfigReplScope) Value() string {
	return c.value
}

func (c TableReplConfigReplScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableReplConfigReplScope) UnmarshalJSON(b []byte) error {
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
