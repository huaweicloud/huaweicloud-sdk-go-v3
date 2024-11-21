package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHotQuestionRequest Request Object
type ListHotQuestionRequest struct {

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

	// 排序方式。 * asc：升序 * desc：降序  默认asc升序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 机器人ID。
	RobotId string `json:"robot_id"`

	// 智能交互语言  * CN:中文  * EN:英文
	Language *ListHotQuestionRequestLanguage `json:"language,omitempty"`
}

func (o ListHotQuestionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotQuestionRequest struct{}"
	}

	return strings.Join([]string{"ListHotQuestionRequest", string(data)}, " ")
}

type ListHotQuestionRequestLanguage struct {
	value string
}

type ListHotQuestionRequestLanguageEnum struct {
	CN ListHotQuestionRequestLanguage
	EN ListHotQuestionRequestLanguage
}

func GetListHotQuestionRequestLanguageEnum() ListHotQuestionRequestLanguageEnum {
	return ListHotQuestionRequestLanguageEnum{
		CN: ListHotQuestionRequestLanguage{
			value: "CN",
		},
		EN: ListHotQuestionRequestLanguage{
			value: "EN",
		},
	}
}

func (c ListHotQuestionRequestLanguage) Value() string {
	return c.value
}

func (c ListHotQuestionRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotQuestionRequestLanguage) UnmarshalJSON(b []byte) error {
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
