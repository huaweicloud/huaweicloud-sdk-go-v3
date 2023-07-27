package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigFormatSingleCreate 日志接入格式单行日志
type AccessConfigFormatSingleCreate struct {

	// 单行日志。system：系统时间，wildcard：时间通配符。
	Mode AccessConfigFormatSingleCreateMode `json:"mode"`

	// 日志时间。 当mode为”system”，则填入当前时间戳。 当mode为\"wildcard\"，则时间通配符：用日志打印时间来标识一条日志数据，通过时间通配符来匹配日志，每条日志的行首显示日志的打印时间；如果日志中的时间格式为：2019-01-01 23:59:59，时间通配符应该填写为：YYYY-MM-DD hh:mm:ss；如果日志中的时间格式为：19-1-1 23:59:59，时间通配符应该填写为：YY-M-D hh:mm:ss
	Value string `json:"value"`
}

func (o AccessConfigFormatSingleCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigFormatSingleCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigFormatSingleCreate", string(data)}, " ")
}

type AccessConfigFormatSingleCreateMode struct {
	value string
}

type AccessConfigFormatSingleCreateModeEnum struct {
	SYSTEM   AccessConfigFormatSingleCreateMode
	WILDCARD AccessConfigFormatSingleCreateMode
}

func GetAccessConfigFormatSingleCreateModeEnum() AccessConfigFormatSingleCreateModeEnum {
	return AccessConfigFormatSingleCreateModeEnum{
		SYSTEM: AccessConfigFormatSingleCreateMode{
			value: "system",
		},
		WILDCARD: AccessConfigFormatSingleCreateMode{
			value: "wildcard",
		},
	}
}

func (c AccessConfigFormatSingleCreateMode) Value() string {
	return c.value
}

func (c AccessConfigFormatSingleCreateMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigFormatSingleCreateMode) UnmarshalJSON(b []byte) error {
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
