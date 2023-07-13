package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigFormatMutilCreate 日志接入格式多行日志
type AccessConfigFormatMutilCreate struct {

	// 单行日志。time：日志时间，regular：正则模式。
	Mode *AccessConfigFormatMutilCreateMode `json:"mode,omitempty"`

	// 日志时间。 当mode为\"regular\"，则输入正则表达式 当mode为\"time\"，则时间通配符：用日志打印时间来标识一条日志数据，通过时间通配符来匹配日志，每条日志的行首显示日志的打印时间；如果日志中的时间格式为：2019-01-01 23:59:59，时间通配符应该填写为：YYYY-MM-DD hh:mm:ss；如果日志中的时间格式为：19-1-1 23:59:59，时间通配符应该填写为：YY-M-D hh:mm:ss
	Value *string `json:"value,omitempty"`
}

func (o AccessConfigFormatMutilCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatMutilCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatMutilCreate", string(data)}, " ")
}

type AccessConfigFormatMutilCreateMode struct {
	value string
}

type AccessConfigFormatMutilCreateModeEnum struct {
	TIME    AccessConfigFormatMutilCreateMode
	REGULAR AccessConfigFormatMutilCreateMode
}

func GetAccessConfigFormatMutilCreateModeEnum() AccessConfigFormatMutilCreateModeEnum {
	return AccessConfigFormatMutilCreateModeEnum{
		TIME: AccessConfigFormatMutilCreateMode{
			value: "time",
		},
		REGULAR: AccessConfigFormatMutilCreateMode{
			value: "regular",
		},
	}
}

func (c AccessConfigFormatMutilCreateMode) Value() string {
	return c.value
}

func (c AccessConfigFormatMutilCreateMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigFormatMutilCreateMode) UnmarshalJSON(b []byte) error {
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
