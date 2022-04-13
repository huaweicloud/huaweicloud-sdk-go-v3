package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type PutUpdateVocabReq struct {
	// 热词表名，不可重复。内容限制为字母，数字，下中划线和井号，长度不超过32字节。

	Name string `json:"name"`
	// 热词表描述，长度不超过255字节。

	Description *string `json:"description,omitempty"`
	// 热词表语言类型。 language取值范围： chinese_mandarin  汉语普通话

	Language PutUpdateVocabReqLanguage `json:"language"`
	// 支持中英混编热词，单个热词只能由英文字母和unicode编码的汉字组成，不能有其他符号，包括空格。  单词库支持热词数上限1024。 单个热词长度上限32字节。

	Contents []string `json:"contents"`
}

func (o PutUpdateVocabReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutUpdateVocabReq struct{}"
	}

	return strings.Join([]string{"PutUpdateVocabReq", string(data)}, " ")
}

type PutUpdateVocabReqLanguage struct {
	value string
}

type PutUpdateVocabReqLanguageEnum struct {
	CHINESE_MANDARIN PutUpdateVocabReqLanguage
}

func GetPutUpdateVocabReqLanguageEnum() PutUpdateVocabReqLanguageEnum {
	return PutUpdateVocabReqLanguageEnum{
		CHINESE_MANDARIN: PutUpdateVocabReqLanguage{
			value: "chinese_mandarin",
		},
	}
}

func (c PutUpdateVocabReqLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PutUpdateVocabReqLanguage) UnmarshalJSON(b []byte) error {
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
