package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPacifyWordsRequest Request Object
type ListPacifyWordsRequest struct {

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

	// 智能交互语言  * CN:中文  * EN:英文
	Language *ListPacifyWordsRequestLanguage `json:"language,omitempty"`

	// 安抚话术类型 > 0:通用安抚话术, 1:基于意图匹配安抚话术
	PacifyWordsType *int32 `json:"pacify_words_type,omitempty"`

	// 安抚话术意图
	Intent *string `json:"intent,omitempty"`
}

func (o ListPacifyWordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPacifyWordsRequest struct{}"
	}

	return strings.Join([]string{"ListPacifyWordsRequest", string(data)}, " ")
}

type ListPacifyWordsRequestLanguage struct {
	value string
}

type ListPacifyWordsRequestLanguageEnum struct {
	CN ListPacifyWordsRequestLanguage
	EN ListPacifyWordsRequestLanguage
}

func GetListPacifyWordsRequestLanguageEnum() ListPacifyWordsRequestLanguageEnum {
	return ListPacifyWordsRequestLanguageEnum{
		CN: ListPacifyWordsRequestLanguage{
			value: "CN",
		},
		EN: ListPacifyWordsRequestLanguage{
			value: "EN",
		},
	}
}

func (c ListPacifyWordsRequestLanguage) Value() string {
	return c.value
}

func (c ListPacifyWordsRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPacifyWordsRequestLanguage) UnmarshalJSON(b []byte) error {
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
