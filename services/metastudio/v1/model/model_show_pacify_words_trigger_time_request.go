package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPacifyWordsTriggerTimeRequest Request Object
type ShowPacifyWordsTriggerTimeRequest struct {

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

	// 智能交互语言 * CN：中文。 * EN：英文。 * ESP：西班牙语（仅海外站点支持） * por：葡萄牙语（仅海外站点支持） * Arabic：阿拉伯语（仅海外站点支持） * Thai：泰语（仅海外站点支持）
	Language ShowPacifyWordsTriggerTimeRequestLanguage `json:"language"`
}

func (o ShowPacifyWordsTriggerTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPacifyWordsTriggerTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowPacifyWordsTriggerTimeRequest", string(data)}, " ")
}

type ShowPacifyWordsTriggerTimeRequestLanguage struct {
	value string
}

type ShowPacifyWordsTriggerTimeRequestLanguageEnum struct {
	CN     ShowPacifyWordsTriggerTimeRequestLanguage
	EN     ShowPacifyWordsTriggerTimeRequestLanguage
	ESP    ShowPacifyWordsTriggerTimeRequestLanguage
	POR    ShowPacifyWordsTriggerTimeRequestLanguage
	ARABIC ShowPacifyWordsTriggerTimeRequestLanguage
	THAI   ShowPacifyWordsTriggerTimeRequestLanguage
}

func GetShowPacifyWordsTriggerTimeRequestLanguageEnum() ShowPacifyWordsTriggerTimeRequestLanguageEnum {
	return ShowPacifyWordsTriggerTimeRequestLanguageEnum{
		CN: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "CN",
		},
		EN: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "EN",
		},
		ESP: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "ESP",
		},
		POR: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "por",
		},
		ARABIC: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "Arabic",
		},
		THAI: ShowPacifyWordsTriggerTimeRequestLanguage{
			value: "Thai",
		},
	}
}

func (c ShowPacifyWordsTriggerTimeRequestLanguage) Value() string {
	return c.value
}

func (c ShowPacifyWordsTriggerTimeRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPacifyWordsTriggerTimeRequestLanguage) UnmarshalJSON(b []byte) error {
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
