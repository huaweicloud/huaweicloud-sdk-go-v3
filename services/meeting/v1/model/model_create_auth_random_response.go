package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAuthRandomResponse Response Object
type CreateAuthRandomResponse struct {

	// 鉴权随机数
	Random *string `json:"random,omitempty"`

	// 会议主题
	Subject *string `json:"subject,omitempty"`

	// 会议类型模型。 * COMMON：MCU会议 * RTC：MMR会议
	ConfMode *CreateAuthRandomResponseConfMode `json:"conf_mode,omitempty"`

	// 是否为网络研讨会
	Webinar *bool `json:"webinar,omitempty"`

	// 是否需要密码
	NeedPassword *bool `json:"need_password,omitempty"`

	// 是否支持小程序
	SupportApplets *bool `json:"support_applets,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateAuthRandomResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthRandomResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthRandomResponse", string(data)}, " ")
}

type CreateAuthRandomResponseConfMode struct {
	value string
}

type CreateAuthRandomResponseConfModeEnum struct {
	COMMON CreateAuthRandomResponseConfMode
	RTC    CreateAuthRandomResponseConfMode
}

func GetCreateAuthRandomResponseConfModeEnum() CreateAuthRandomResponseConfModeEnum {
	return CreateAuthRandomResponseConfModeEnum{
		COMMON: CreateAuthRandomResponseConfMode{
			value: "COMMON",
		},
		RTC: CreateAuthRandomResponseConfMode{
			value: "RTC",
		},
	}
}

func (c CreateAuthRandomResponseConfMode) Value() string {
	return c.value
}

func (c CreateAuthRandomResponseConfMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAuthRandomResponseConfMode) UnmarshalJSON(b []byte) error {
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
