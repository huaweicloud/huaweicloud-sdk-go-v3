package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HotWordsTypeEnum 热词类型。 > SIS:使用的语音识别服务为SIS时选此类型
type HotWordsTypeEnum struct {
	value string
}

type HotWordsTypeEnumEnum struct {
	SIS HotWordsTypeEnum
}

func GetHotWordsTypeEnumEnum() HotWordsTypeEnumEnum {
	return HotWordsTypeEnumEnum{
		SIS: HotWordsTypeEnum{
			value: "SIS",
		},
	}
}

func (c HotWordsTypeEnum) Value() string {
	return c.value
}

func (c HotWordsTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HotWordsTypeEnum) UnmarshalJSON(b []byte) error {
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
