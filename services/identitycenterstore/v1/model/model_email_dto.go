package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EmailDto The email address associated with the user.
type EmailDto struct {

	// 一个布尔值，表示这是否为用户的主电子邮箱
	Primary bool `json:"primary"`

	// 表示电子邮箱类型的字符串
	Type string `json:"type"`

	// 包含电子邮箱地址的字符串
	Value string `json:"value"`

	// 电子邮箱地址的验证状态
	VerificationStatus *EmailDtoVerificationStatus `json:"verification_status,omitempty"`
}

func (o EmailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailDto struct{}"
	}

	return strings.Join([]string{"EmailDto", string(data)}, " ")
}

type EmailDtoVerificationStatus struct {
	value string
}

type EmailDtoVerificationStatusEnum struct {
	NOT_VERIFIED EmailDtoVerificationStatus
	VERIFIED     EmailDtoVerificationStatus
}

func GetEmailDtoVerificationStatusEnum() EmailDtoVerificationStatusEnum {
	return EmailDtoVerificationStatusEnum{
		NOT_VERIFIED: EmailDtoVerificationStatus{
			value: "NOT_VERIFIED",
		},
		VERIFIED: EmailDtoVerificationStatus{
			value: "VERIFIED",
		},
	}
}

func (c EmailDtoVerificationStatus) Value() string {
	return c.value
}

func (c EmailDtoVerificationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EmailDtoVerificationStatus) UnmarshalJSON(b []byte) error {
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
