package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableDisplaySetting 表展示设置
type TableDisplaySetting struct {

	// 表格列展示列表
	Columns *[]TableColumnDisplaySetting `json:"columns,omitempty"`

	// **参数解释**: 表展示设置 - TABLE 表展示 - RAW 原始数据展示  **约束限制** 不涉及 **取值范围**: - TABLE - RAW  **默认值** 不涉及
	Format *TableDisplaySettingFormat `json:"format,omitempty"`
}

func (o TableDisplaySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableDisplaySetting struct{}"
	}

	return strings.Join([]string{"TableDisplaySetting", string(data)}, " ")
}

type TableDisplaySettingFormat struct {
	value string
}

type TableDisplaySettingFormatEnum struct {
	TABLE TableDisplaySettingFormat
	RAW   TableDisplaySettingFormat
}

func GetTableDisplaySettingFormatEnum() TableDisplaySettingFormatEnum {
	return TableDisplaySettingFormatEnum{
		TABLE: TableDisplaySettingFormat{
			value: "TABLE",
		},
		RAW: TableDisplaySettingFormat{
			value: "RAW",
		},
	}
}

func (c TableDisplaySettingFormat) Value() string {
	return c.value
}

func (c TableDisplaySettingFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableDisplaySettingFormat) UnmarshalJSON(b []byte) error {
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
