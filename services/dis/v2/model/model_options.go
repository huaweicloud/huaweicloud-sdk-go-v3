package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Options DWS容错性选项（用于指定外表数据的各类参数）。
type Options struct {

	// 数据入库时，数据源文件中某行的最后一个字段缺失时，请选择是直接将字段设为Null，还是在错误表中报错提示。  取值范围：   - true/on - false/off  缺省值：false/off
	FillMissingFields *OptionsFillMissingFields `json:"fill_missing_fields,omitempty"`

	// 数据源文件中的字段比外表定义列数多时，是否忽略多出的列。该参数只在数据导入过程中使用。  取值范围：  - true/on - false/off  缺省值：false/off
	IgnoreExtraData *OptionsIgnoreExtraData `json:"ignore_extra_data,omitempty"`

	// 导入非法字符容错参数。是将非法字符按照转换规则转换后入库，还是报错中止导入。  取值范围：  - true/on - false/off  缺省值：false/off
	CompatibleIllegalChars *OptionsCompatibleIllegalChars `json:"compatible_illegal_chars,omitempty"`

	// 指定本次数据导入允许出现的数据格式错误个数，当导入过程中出现的数据格式错误未达到限定值时，本次数据导入可以成功。  取值范围：  - 整型值 - unlimited（无限制）  缺省值为0，有错误信息立即返回。
	RejectLimit *string `json:"reject_limit,omitempty"`

	// 用于记录数据格式错误信息的错误表表名。并行导入结束后查询此错误信息表，能够获取详细的错误信息。
	ErrorTableName *string `json:"error_table_name,omitempty"`
}

func (o Options) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Options struct{}"
	}

	return strings.Join([]string{"Options", string(data)}, " ")
}

type OptionsFillMissingFields struct {
	value string
}

type OptionsFillMissingFieldsEnum struct {
	TRUE_ON   OptionsFillMissingFields
	FALSE_OFF OptionsFillMissingFields
}

func GetOptionsFillMissingFieldsEnum() OptionsFillMissingFieldsEnum {
	return OptionsFillMissingFieldsEnum{
		TRUE_ON: OptionsFillMissingFields{
			value: "true/on",
		},
		FALSE_OFF: OptionsFillMissingFields{
			value: "false/off",
		},
	}
}

func (c OptionsFillMissingFields) Value() string {
	return c.value
}

func (c OptionsFillMissingFields) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OptionsFillMissingFields) UnmarshalJSON(b []byte) error {
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

type OptionsIgnoreExtraData struct {
	value string
}

type OptionsIgnoreExtraDataEnum struct {
	TRUE_ON   OptionsIgnoreExtraData
	FALSE_OFF OptionsIgnoreExtraData
}

func GetOptionsIgnoreExtraDataEnum() OptionsIgnoreExtraDataEnum {
	return OptionsIgnoreExtraDataEnum{
		TRUE_ON: OptionsIgnoreExtraData{
			value: "true/on",
		},
		FALSE_OFF: OptionsIgnoreExtraData{
			value: "false/off",
		},
	}
}

func (c OptionsIgnoreExtraData) Value() string {
	return c.value
}

func (c OptionsIgnoreExtraData) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OptionsIgnoreExtraData) UnmarshalJSON(b []byte) error {
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

type OptionsCompatibleIllegalChars struct {
	value string
}

type OptionsCompatibleIllegalCharsEnum struct {
	TRUE_ON   OptionsCompatibleIllegalChars
	FALSE_OFF OptionsCompatibleIllegalChars
}

func GetOptionsCompatibleIllegalCharsEnum() OptionsCompatibleIllegalCharsEnum {
	return OptionsCompatibleIllegalCharsEnum{
		TRUE_ON: OptionsCompatibleIllegalChars{
			value: "true/on",
		},
		FALSE_OFF: OptionsCompatibleIllegalChars{
			value: "false/off",
		},
	}
}

func (c OptionsCompatibleIllegalChars) Value() string {
	return c.value
}

func (c OptionsCompatibleIllegalChars) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OptionsCompatibleIllegalChars) UnmarshalJSON(b []byte) error {
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
