package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ShowVocabularyResponse struct {
	// 调用成功返回热词表ID，调用失败时无此字段。

	VocabularyId *string `json:"vocabulary_id,omitempty"`
	// 调用成功返回热词表名，调用失败时无此字段。

	Name *string `json:"name,omitempty"`
	// 调用成功返回热词表描述，调用失败时无此字段。

	Description *string `json:"description,omitempty"`
	// 调用成功返回热词表语言类型，调用失败时无此字段。

	Language *ShowVocabularyResponseLanguage `json:"language,omitempty"`
	// 调用成功返回热词列表，调用失败时无此字段。

	Contents       *[]string `json:"contents,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVocabularyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabularyResponse struct{}"
	}

	return strings.Join([]string{"ShowVocabularyResponse", string(data)}, " ")
}

type ShowVocabularyResponseLanguage struct {
	value string
}

type ShowVocabularyResponseLanguageEnum struct {
	CHINESE_MANDARIN ShowVocabularyResponseLanguage
	ENGLISH          ShowVocabularyResponseLanguage
}

func GetShowVocabularyResponseLanguageEnum() ShowVocabularyResponseLanguageEnum {
	return ShowVocabularyResponseLanguageEnum{
		CHINESE_MANDARIN: ShowVocabularyResponseLanguage{
			value: "chinese-mandarin",
		},
		ENGLISH: ShowVocabularyResponseLanguage{
			value: " english",
		},
	}
}

func (c ShowVocabularyResponseLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVocabularyResponseLanguage) UnmarshalJSON(b []byte) error {
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
