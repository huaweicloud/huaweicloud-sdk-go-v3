package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableDisplaySettingDto 表展示设置数据传输对象
type IsapTableDisplaySettingDto struct {

	// 表格列展示列表
	Columns *[]IsapTableColumnDisplaySettingDto `json:"columns,omitempty"`

	// **参数解释**: 表展示设置 - TABLE 表展示 - RAW 原始数据展示  **约束限制** 不涉及 **取值范围**: - TABLE - RAW  **默认值** 不涉及
	Format *IsapTableDisplaySettingDtoFormat `json:"format,omitempty"`
}

func (o IsapTableDisplaySettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapTableDisplaySettingDto struct{}"
	}

	return strings.Join([]string{"IsapTableDisplaySettingDto", string(data)}, " ")
}

type IsapTableDisplaySettingDtoFormat struct {
	value string
}

type IsapTableDisplaySettingDtoFormatEnum struct {
	TABLE IsapTableDisplaySettingDtoFormat
	RAW   IsapTableDisplaySettingDtoFormat
}

func GetIsapTableDisplaySettingDtoFormatEnum() IsapTableDisplaySettingDtoFormatEnum {
	return IsapTableDisplaySettingDtoFormatEnum{
		TABLE: IsapTableDisplaySettingDtoFormat{
			value: "TABLE",
		},
		RAW: IsapTableDisplaySettingDtoFormat{
			value: "RAW",
		},
	}
}

func (c IsapTableDisplaySettingDtoFormat) Value() string {
	return c.value
}

func (c IsapTableDisplaySettingDtoFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableDisplaySettingDtoFormat) UnmarshalJSON(b []byte) error {
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
