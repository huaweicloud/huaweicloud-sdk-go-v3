package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AsrVocabularyTypeEnum 热词表类型。 > MOBVOI:使用的语音识别服务为MOBVOI时选择此类型
type AsrVocabularyTypeEnum struct {
	value string
}

type AsrVocabularyTypeEnumEnum struct {
	MOBVOI AsrVocabularyTypeEnum
}

func GetAsrVocabularyTypeEnumEnum() AsrVocabularyTypeEnumEnum {
	return AsrVocabularyTypeEnumEnum{
		MOBVOI: AsrVocabularyTypeEnum{
			value: "MOBVOI",
		},
	}
}

func (c AsrVocabularyTypeEnum) Value() string {
	return c.value
}

func (c AsrVocabularyTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AsrVocabularyTypeEnum) UnmarshalJSON(b []byte) error {
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
