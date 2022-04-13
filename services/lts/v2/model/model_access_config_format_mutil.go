package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 日志接入格式多行日志
type AccessConfigFormatMutil struct {
	// 单行日志。time：日志时间，regular：正则模式。

	Mode AccessConfigFormatMutilMode `json:"mode"`
	// 日志时间。 当mode为\"regular\"，则输入正则表达式 当mode为\"time\"，则时间通配符：用日志打印时间来标识一条日志数据，通过时间通配符来匹配日志，每条日志的行首显示日志的打印时间；如果日志中的时间格式为：2019-01-01 23:59:59，时间通配符应该填写为：YYYY-MM-DD hh:mm:ss；如果日志中的时间格式为：19-1-1 23:59:59，时间通配符应该填写为：YY-M-D hh:mm:ss

	Value string `json:"value"`
}

func (o AccessConfigFormatMutil) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatMutil struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatMutil", string(data)}, " ")
}

type AccessConfigFormatMutilMode struct {
	value string
}

type AccessConfigFormatMutilModeEnum struct {
	TIME    AccessConfigFormatMutilMode
	REGULAR AccessConfigFormatMutilMode
}

func GetAccessConfigFormatMutilModeEnum() AccessConfigFormatMutilModeEnum {
	return AccessConfigFormatMutilModeEnum{
		TIME: AccessConfigFormatMutilMode{
			value: "time",
		},
		REGULAR: AccessConfigFormatMutilMode{
			value: "regular",
		},
	}
}

func (c AccessConfigFormatMutilMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigFormatMutilMode) UnmarshalJSON(b []byte) error {
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
