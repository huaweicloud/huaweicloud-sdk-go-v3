package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SaveTtscVocabularyConfigsRequestBody 保存租户级tts扩展词表配置请求
type SaveTtscVocabularyConfigsRequestBody struct {

	// TTSS支持配置的词表类型 * CHINESE_G2P:拼音 * PHONETIC_SYMBOL:音标 * CONTINUUM:连读 * ALIAS:别名 * SAY_AS:数字英文读法
	Type SaveTtscVocabularyConfigsRequestBodyType `json:"type"`

	// 映射键
	Key *string `json:"key,omitempty"`

	// 映射值
	Value *string `json:"value,omitempty"`
}

func (o SaveTtscVocabularyConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTtscVocabularyConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"SaveTtscVocabularyConfigsRequestBody", string(data)}, " ")
}

type SaveTtscVocabularyConfigsRequestBodyType struct {
	value string
}

type SaveTtscVocabularyConfigsRequestBodyTypeEnum struct {
	CHINESE_G2_P    SaveTtscVocabularyConfigsRequestBodyType
	PHONETIC_SYMBOL SaveTtscVocabularyConfigsRequestBodyType
	CONTINUUM       SaveTtscVocabularyConfigsRequestBodyType
	ALIAS           SaveTtscVocabularyConfigsRequestBodyType
	SAY_AS          SaveTtscVocabularyConfigsRequestBodyType
}

func GetSaveTtscVocabularyConfigsRequestBodyTypeEnum() SaveTtscVocabularyConfigsRequestBodyTypeEnum {
	return SaveTtscVocabularyConfigsRequestBodyTypeEnum{
		CHINESE_G2_P: SaveTtscVocabularyConfigsRequestBodyType{
			value: "CHINESE_G2P",
		},
		PHONETIC_SYMBOL: SaveTtscVocabularyConfigsRequestBodyType{
			value: "PHONETIC_SYMBOL",
		},
		CONTINUUM: SaveTtscVocabularyConfigsRequestBodyType{
			value: "CONTINUUM",
		},
		ALIAS: SaveTtscVocabularyConfigsRequestBodyType{
			value: "ALIAS",
		},
		SAY_AS: SaveTtscVocabularyConfigsRequestBodyType{
			value: "SAY_AS",
		},
	}
}

func (c SaveTtscVocabularyConfigsRequestBodyType) Value() string {
	return c.value
}

func (c SaveTtscVocabularyConfigsRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SaveTtscVocabularyConfigsRequestBodyType) UnmarshalJSON(b []byte) error {
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
