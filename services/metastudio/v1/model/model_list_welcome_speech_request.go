package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWelcomeSpeechRequest Request Object
type ListWelcomeSpeechRequest struct {

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

	// 智能交互语言  * CN:中文  * EN:英文  * ESP：西班牙语（仅海外站点支持）  * por：葡萄牙语（仅海外站点支持）  * Arabic：阿拉伯语（仅海外站点支持）  * Thai：泰语（仅海外站点支持）  * fr：法语（仅海外站点支持）
	Language *ListWelcomeSpeechRequestLanguage `json:"language,omitempty"`
}

func (o ListWelcomeSpeechRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWelcomeSpeechRequest struct{}"
	}

	return strings.Join([]string{"ListWelcomeSpeechRequest", string(data)}, " ")
}

type ListWelcomeSpeechRequestLanguage struct {
	value string
}

type ListWelcomeSpeechRequestLanguageEnum struct {
	CN     ListWelcomeSpeechRequestLanguage
	EN     ListWelcomeSpeechRequestLanguage
	ESP    ListWelcomeSpeechRequestLanguage
	POR    ListWelcomeSpeechRequestLanguage
	ARABIC ListWelcomeSpeechRequestLanguage
	THAI   ListWelcomeSpeechRequestLanguage
	FR     ListWelcomeSpeechRequestLanguage
}

func GetListWelcomeSpeechRequestLanguageEnum() ListWelcomeSpeechRequestLanguageEnum {
	return ListWelcomeSpeechRequestLanguageEnum{
		CN: ListWelcomeSpeechRequestLanguage{
			value: "CN",
		},
		EN: ListWelcomeSpeechRequestLanguage{
			value: "EN",
		},
		ESP: ListWelcomeSpeechRequestLanguage{
			value: "ESP",
		},
		POR: ListWelcomeSpeechRequestLanguage{
			value: "por",
		},
		ARABIC: ListWelcomeSpeechRequestLanguage{
			value: "Arabic",
		},
		THAI: ListWelcomeSpeechRequestLanguage{
			value: "Thai",
		},
		FR: ListWelcomeSpeechRequestLanguage{
			value: "fr",
		},
	}
}

func (c ListWelcomeSpeechRequestLanguage) Value() string {
	return c.value
}

func (c ListWelcomeSpeechRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWelcomeSpeechRequestLanguage) UnmarshalJSON(b []byte) error {
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
