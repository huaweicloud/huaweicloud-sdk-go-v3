package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 授权结果
type AuthResult struct {

	// 授权结果 - SUCCESS：授权成功 - SKIPPED：跳过 - FAILED：授权失败
	Status *AuthResultStatus `json:"status,omitempty" xml:"status"`

	// 授权失败错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 授权失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 授权失败的API名称
	ApiName *string `json:"api_name,omitempty" xml:"api_name"`

	// 授权失败的APP名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`
}

func (o AuthResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthResult struct{}"
	}

	return strings.Join([]string{"AuthResult", string(data)}, " ")
}

type AuthResultStatus struct {
	value string
}

type AuthResultStatusEnum struct {
	SUCCESS AuthResultStatus
	SKIPPED AuthResultStatus
	FAILED  AuthResultStatus
}

func GetAuthResultStatusEnum() AuthResultStatusEnum {
	return AuthResultStatusEnum{
		SUCCESS: AuthResultStatus{
			value: "SUCCESS",
		},
		SKIPPED: AuthResultStatus{
			value: "SKIPPED",
		},
		FAILED: AuthResultStatus{
			value: "FAILED",
		},
	}
}

func (c AuthResultStatus) Value() string {
	return c.value
}

func (c AuthResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthResultStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
