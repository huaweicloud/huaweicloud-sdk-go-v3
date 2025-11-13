package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShortJobType 短任务类型。 * VOICE_ASSESS: 声音质量检测
type ShortJobType struct {
	value string
}

type ShortJobTypeEnum struct {
	VOICE_ASSESS ShortJobType
}

func GetShortJobTypeEnum() ShortJobTypeEnum {
	return ShortJobTypeEnum{
		VOICE_ASSESS: ShortJobType{
			value: "VOICE_ASSESS",
		},
	}
}

func (c ShortJobType) Value() string {
	return c.value
}

func (c ShortJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShortJobType) UnmarshalJSON(b []byte) error {
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
