package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SsoConfigurationDto the struct of SSOConfiguration
type SsoConfigurationDto struct {

	// MFA生效模式
	MfaMode *SsoConfigurationDtoMfaMode `json:"mfa_mode,omitempty"`

	// 未注册MFA情况下，可选择的登录行为
	NoMfaSigninBehavior *SsoConfigurationDtoNoMfaSigninBehavior `json:"no_mfa_signin_behavior,omitempty"`

	// 没有密码情况下登录的行为
	NoPasswordSigninBehavior *SsoConfigurationDtoNoPasswordSigninBehavior `json:"no_password_signin_behavior,omitempty"`

	// 允许的MFA类型
	AllowedMfaTypes *[]SsoConfigurationDtoAllowedMfaTypes `json:"allowed_mfa_types,omitempty"`

	SessionConfiguration *SessionConfigurationDto `json:"session_configuration,omitempty"`
}

func (o SsoConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SsoConfigurationDto struct{}"
	}

	return strings.Join([]string{"SsoConfigurationDto", string(data)}, " ")
}

type SsoConfigurationDtoMfaMode struct {
	value string
}

type SsoConfigurationDtoMfaModeEnum struct {
	CONTEXT_AWARE SsoConfigurationDtoMfaMode
	DISABLED      SsoConfigurationDtoMfaMode
	ALWAYS_ON     SsoConfigurationDtoMfaMode
}

func GetSsoConfigurationDtoMfaModeEnum() SsoConfigurationDtoMfaModeEnum {
	return SsoConfigurationDtoMfaModeEnum{
		CONTEXT_AWARE: SsoConfigurationDtoMfaMode{
			value: "CONTEXT_AWARE",
		},
		DISABLED: SsoConfigurationDtoMfaMode{
			value: "DISABLED",
		},
		ALWAYS_ON: SsoConfigurationDtoMfaMode{
			value: "ALWAYS_ON",
		},
	}
}

func (c SsoConfigurationDtoMfaMode) Value() string {
	return c.value
}

func (c SsoConfigurationDtoMfaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SsoConfigurationDtoMfaMode) UnmarshalJSON(b []byte) error {
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

type SsoConfigurationDtoNoMfaSigninBehavior struct {
	value string
}

type SsoConfigurationDtoNoMfaSigninBehaviorEnum struct {
	ALLOWED_WITH_ENROLLMENT SsoConfigurationDtoNoMfaSigninBehavior
	ALLOWED                 SsoConfigurationDtoNoMfaSigninBehavior
	EMAIL_OTP               SsoConfigurationDtoNoMfaSigninBehavior
	BLOCKED                 SsoConfigurationDtoNoMfaSigninBehavior
}

func GetSsoConfigurationDtoNoMfaSigninBehaviorEnum() SsoConfigurationDtoNoMfaSigninBehaviorEnum {
	return SsoConfigurationDtoNoMfaSigninBehaviorEnum{
		ALLOWED_WITH_ENROLLMENT: SsoConfigurationDtoNoMfaSigninBehavior{
			value: "ALLOWED_WITH_ENROLLMENT",
		},
		ALLOWED: SsoConfigurationDtoNoMfaSigninBehavior{
			value: "ALLOWED",
		},
		EMAIL_OTP: SsoConfigurationDtoNoMfaSigninBehavior{
			value: "EMAIL_OTP",
		},
		BLOCKED: SsoConfigurationDtoNoMfaSigninBehavior{
			value: "BLOCKED",
		},
	}
}

func (c SsoConfigurationDtoNoMfaSigninBehavior) Value() string {
	return c.value
}

func (c SsoConfigurationDtoNoMfaSigninBehavior) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SsoConfigurationDtoNoMfaSigninBehavior) UnmarshalJSON(b []byte) error {
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

type SsoConfigurationDtoNoPasswordSigninBehavior struct {
	value string
}

type SsoConfigurationDtoNoPasswordSigninBehaviorEnum struct {
	BLOCKED   SsoConfigurationDtoNoPasswordSigninBehavior
	EMAIL_OTP SsoConfigurationDtoNoPasswordSigninBehavior
}

func GetSsoConfigurationDtoNoPasswordSigninBehaviorEnum() SsoConfigurationDtoNoPasswordSigninBehaviorEnum {
	return SsoConfigurationDtoNoPasswordSigninBehaviorEnum{
		BLOCKED: SsoConfigurationDtoNoPasswordSigninBehavior{
			value: "BLOCKED",
		},
		EMAIL_OTP: SsoConfigurationDtoNoPasswordSigninBehavior{
			value: "EMAIL_OTP",
		},
	}
}

func (c SsoConfigurationDtoNoPasswordSigninBehavior) Value() string {
	return c.value
}

func (c SsoConfigurationDtoNoPasswordSigninBehavior) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SsoConfigurationDtoNoPasswordSigninBehavior) UnmarshalJSON(b []byte) error {
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

type SsoConfigurationDtoAllowedMfaTypes struct {
	value string
}

type SsoConfigurationDtoAllowedMfaTypesEnum struct {
	TOTP                  SsoConfigurationDtoAllowedMfaTypes
	WEBAUTHN              SsoConfigurationDtoAllowedMfaTypes
	WEBAUTHN_SECURITY_KEY SsoConfigurationDtoAllowedMfaTypes
}

func GetSsoConfigurationDtoAllowedMfaTypesEnum() SsoConfigurationDtoAllowedMfaTypesEnum {
	return SsoConfigurationDtoAllowedMfaTypesEnum{
		TOTP: SsoConfigurationDtoAllowedMfaTypes{
			value: "TOTP",
		},
		WEBAUTHN: SsoConfigurationDtoAllowedMfaTypes{
			value: "WEBAUTHN",
		},
		WEBAUTHN_SECURITY_KEY: SsoConfigurationDtoAllowedMfaTypes{
			value: "WEBAUTHN_SECURITY_KEY",
		},
	}
}

func (c SsoConfigurationDtoAllowedMfaTypes) Value() string {
	return c.value
}

func (c SsoConfigurationDtoAllowedMfaTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SsoConfigurationDtoAllowedMfaTypes) UnmarshalJSON(b []byte) error {
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
