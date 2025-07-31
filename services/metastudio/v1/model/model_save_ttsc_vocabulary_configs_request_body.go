package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SaveTtscVocabularyConfigsRequestBody 保存租户级tts自定义读法配置请求
type SaveTtscVocabularyConfigsRequestBody struct {

	// 支持配置的自定义读法类型。当前读法类型会映射为SSML标签，详见[文本驱动SSML定义](metastudio_02_0038.xml)。  包含如下选项： * CHINESE_G2P：拼音 * PHONETIC_SYMBOL：音标 * CONTINUUM：连读 * ALIAS：别名 * SAY_AS：数字/英文的读法。不同value值有不同的读法，详情如下所示。   数字的读法包括：   - date：读日期   - number：读数字   - figure：读数值   - telephone：读电话    英文的读法包括：   - spell：读字母   - english：读单词
	Type SaveTtscVocabularyConfigsRequestBodyType `json:"type"`

	// 原始词。
	Key *string `json:"key,omitempty"`

	// 自定义读法。其中，音标的读法请参考[词典](https://www.youdao.com/)。
	Value *string `json:"value,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`
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
