package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VocabInfo
type VocabInfo struct {

	// 热词表ID。
	VocabularyId string `json:"vocabulary_id"`

	// 热词表名。
	Name string `json:"name"`

	// 热词表语言类型。
	Language VocabInfoLanguage `json:"language"`

	// 热词表描述。
	Description string `json:"description"`
}

func (o VocabInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VocabInfo struct{}"
	}

	return strings.Join([]string{"VocabInfo", string(data)}, " ")
}

type VocabInfoLanguage struct {
	value string
}

type VocabInfoLanguageEnum struct {
	CHINESE_MANDARIN VocabInfoLanguage
	ENGLISH          VocabInfoLanguage
}

func GetVocabInfoLanguageEnum() VocabInfoLanguageEnum {
	return VocabInfoLanguageEnum{
		CHINESE_MANDARIN: VocabInfoLanguage{
			value: "chinese_mandarin",
		},
		ENGLISH: VocabInfoLanguage{
			value: "english",
		},
	}
}

func (c VocabInfoLanguage) Value() string {
	return c.value
}

func (c VocabInfoLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VocabInfoLanguage) UnmarshalJSON(b []byte) error {
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
