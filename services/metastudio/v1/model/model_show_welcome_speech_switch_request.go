package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWelcomeSpeechSwitchRequest Request Object
type ShowWelcomeSpeechSwitchRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 智能交互语言  * CN:中文  * EN:英文
	Language *ShowWelcomeSpeechSwitchRequestLanguage `json:"language,omitempty"`
}

func (o ShowWelcomeSpeechSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWelcomeSpeechSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowWelcomeSpeechSwitchRequest", string(data)}, " ")
}

type ShowWelcomeSpeechSwitchRequestLanguage struct {
	value string
}

type ShowWelcomeSpeechSwitchRequestLanguageEnum struct {
	CN ShowWelcomeSpeechSwitchRequestLanguage
	EN ShowWelcomeSpeechSwitchRequestLanguage
}

func GetShowWelcomeSpeechSwitchRequestLanguageEnum() ShowWelcomeSpeechSwitchRequestLanguageEnum {
	return ShowWelcomeSpeechSwitchRequestLanguageEnum{
		CN: ShowWelcomeSpeechSwitchRequestLanguage{
			value: "CN",
		},
		EN: ShowWelcomeSpeechSwitchRequestLanguage{
			value: "EN",
		},
	}
}

func (c ShowWelcomeSpeechSwitchRequestLanguage) Value() string {
	return c.value
}

func (c ShowWelcomeSpeechSwitchRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWelcomeSpeechSwitchRequestLanguage) UnmarshalJSON(b []byte) error {
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
