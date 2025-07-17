package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAsrVocabularyRequest Request Object
type ListAsrVocabularyRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 词表类型。 > MOBVOI:使用的语音识别服务为MOBVOI时选择此类型
	VocabularyType *ListAsrVocabularyRequestVocabularyType `json:"vocabulary_type,omitempty"`

	// 智能交互语言 * CN：中文。 * EN：英文。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	Language *ListAsrVocabularyRequestLanguage `json:"language,omitempty"`
}

func (o ListAsrVocabularyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsrVocabularyRequest struct{}"
	}

	return strings.Join([]string{"ListAsrVocabularyRequest", string(data)}, " ")
}

type ListAsrVocabularyRequestVocabularyType struct {
	value string
}

type ListAsrVocabularyRequestVocabularyTypeEnum struct {
	MOBVOI ListAsrVocabularyRequestVocabularyType
}

func GetListAsrVocabularyRequestVocabularyTypeEnum() ListAsrVocabularyRequestVocabularyTypeEnum {
	return ListAsrVocabularyRequestVocabularyTypeEnum{
		MOBVOI: ListAsrVocabularyRequestVocabularyType{
			value: "MOBVOI",
		},
	}
}

func (c ListAsrVocabularyRequestVocabularyType) Value() string {
	return c.value
}

func (c ListAsrVocabularyRequestVocabularyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAsrVocabularyRequestVocabularyType) UnmarshalJSON(b []byte) error {
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

type ListAsrVocabularyRequestLanguage struct {
	value string
}

type ListAsrVocabularyRequestLanguageEnum struct {
	CN     ListAsrVocabularyRequestLanguage
	EN     ListAsrVocabularyRequestLanguage
	ESP    ListAsrVocabularyRequestLanguage
	POR    ListAsrVocabularyRequestLanguage
	ARABIC ListAsrVocabularyRequestLanguage
	THAI   ListAsrVocabularyRequestLanguage
}

func GetListAsrVocabularyRequestLanguageEnum() ListAsrVocabularyRequestLanguageEnum {
	return ListAsrVocabularyRequestLanguageEnum{
		CN: ListAsrVocabularyRequestLanguage{
			value: "CN",
		},
		EN: ListAsrVocabularyRequestLanguage{
			value: "EN",
		},
		ESP: ListAsrVocabularyRequestLanguage{
			value: "ESP",
		},
		POR: ListAsrVocabularyRequestLanguage{
			value: "por",
		},
		ARABIC: ListAsrVocabularyRequestLanguage{
			value: "Arabic",
		},
		THAI: ListAsrVocabularyRequestLanguage{
			value: "Thai",
		},
	}
}

func (c ListAsrVocabularyRequestLanguage) Value() string {
	return c.value
}

func (c ListAsrVocabularyRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAsrVocabularyRequestLanguage) UnmarshalJSON(b []byte) error {
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
