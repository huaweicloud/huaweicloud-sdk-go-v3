package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPacifyWordsSwitchRequest Request Object
type ShowPacifyWordsSwitchRequest struct {

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
	Language ShowPacifyWordsSwitchRequestLanguage `json:"language"`
}

func (o ShowPacifyWordsSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPacifyWordsSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowPacifyWordsSwitchRequest", string(data)}, " ")
}

type ShowPacifyWordsSwitchRequestLanguage struct {
	value string
}

type ShowPacifyWordsSwitchRequestLanguageEnum struct {
	CN ShowPacifyWordsSwitchRequestLanguage
	EN ShowPacifyWordsSwitchRequestLanguage
}

func GetShowPacifyWordsSwitchRequestLanguageEnum() ShowPacifyWordsSwitchRequestLanguageEnum {
	return ShowPacifyWordsSwitchRequestLanguageEnum{
		CN: ShowPacifyWordsSwitchRequestLanguage{
			value: "CN",
		},
		EN: ShowPacifyWordsSwitchRequestLanguage{
			value: "EN",
		},
	}
}

func (c ShowPacifyWordsSwitchRequestLanguage) Value() string {
	return c.value
}

func (c ShowPacifyWordsSwitchRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPacifyWordsSwitchRequestLanguage) UnmarshalJSON(b []byte) error {
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
