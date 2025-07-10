package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHotWordsRequest Request Object
type ListHotWordsRequest struct {

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

	// 应用ID。
	RobotId string `json:"robot_id"`

	// sis服务所在区域
	Region *int32 `json:"region,omitempty"`

	// 智能交互语言 * zh_CN：简体中文（已下线，请使用CN） * en_US：英语（已下线，请使用EN） * CN：中文。 * EN：英文。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	Language *ListHotWordsRequestLanguage `json:"language,omitempty"`

	// * SIS:SIS热词 * MOBVOI:奇妙问ASR热词
	HotWordsType *ListHotWordsRequestHotWordsType `json:"hot_words_type,omitempty"`
}

func (o ListHotWordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotWordsRequest struct{}"
	}

	return strings.Join([]string{"ListHotWordsRequest", string(data)}, " ")
}

type ListHotWordsRequestLanguage struct {
	value string
}

type ListHotWordsRequestLanguageEnum struct {
	ZH_CN  ListHotWordsRequestLanguage
	EN_US  ListHotWordsRequestLanguage
	CN     ListHotWordsRequestLanguage
	EN     ListHotWordsRequestLanguage
	ESP    ListHotWordsRequestLanguage
	POR    ListHotWordsRequestLanguage
	ARABIC ListHotWordsRequestLanguage
	THAI   ListHotWordsRequestLanguage
}

func GetListHotWordsRequestLanguageEnum() ListHotWordsRequestLanguageEnum {
	return ListHotWordsRequestLanguageEnum{
		ZH_CN: ListHotWordsRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ListHotWordsRequestLanguage{
			value: "en_US",
		},
		CN: ListHotWordsRequestLanguage{
			value: "CN",
		},
		EN: ListHotWordsRequestLanguage{
			value: "EN",
		},
		ESP: ListHotWordsRequestLanguage{
			value: "ESP",
		},
		POR: ListHotWordsRequestLanguage{
			value: "por",
		},
		ARABIC: ListHotWordsRequestLanguage{
			value: "Arabic",
		},
		THAI: ListHotWordsRequestLanguage{
			value: "Thai",
		},
	}
}

func (c ListHotWordsRequestLanguage) Value() string {
	return c.value
}

func (c ListHotWordsRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotWordsRequestLanguage) UnmarshalJSON(b []byte) error {
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

type ListHotWordsRequestHotWordsType struct {
	value string
}

type ListHotWordsRequestHotWordsTypeEnum struct {
	SIS    ListHotWordsRequestHotWordsType
	MOBVOI ListHotWordsRequestHotWordsType
}

func GetListHotWordsRequestHotWordsTypeEnum() ListHotWordsRequestHotWordsTypeEnum {
	return ListHotWordsRequestHotWordsTypeEnum{
		SIS: ListHotWordsRequestHotWordsType{
			value: "SIS",
		},
		MOBVOI: ListHotWordsRequestHotWordsType{
			value: "MOBVOI",
		},
	}
}

func (c ListHotWordsRequestHotWordsType) Value() string {
	return c.value
}

func (c ListHotWordsRequestHotWordsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotWordsRequestHotWordsType) UnmarshalJSON(b []byte) error {
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
