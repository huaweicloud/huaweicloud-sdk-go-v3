package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PlatformAuthorizationInfo 直播平台授权信息
type PlatformAuthorizationInfo struct {

	// 授权状态。 * AUTHORIZED: 已授权 * UNAUTHORIZED: 未授权
	AuthorizeState *PlatformAuthorizationInfoAuthorizeState `json:"authorize_state,omitempty"`

	// 授权时间
	AuthorizedTime *string `json:"authorized_time,omitempty"`

	// 过期时间
	ExpireTime *string `json:"expire_time,omitempty"`

	// 授权账号信息。 美团平台对应：opBizCode
	Account *string `json:"account,omitempty"`
}

func (o PlatformAuthorizationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlatformAuthorizationInfo struct{}"
	}

	return strings.Join([]string{"PlatformAuthorizationInfo", string(data)}, " ")
}

type PlatformAuthorizationInfoAuthorizeState struct {
	value string
}

type PlatformAuthorizationInfoAuthorizeStateEnum struct {
	AUTHORIZED   PlatformAuthorizationInfoAuthorizeState
	UNAUTHORIZED PlatformAuthorizationInfoAuthorizeState
}

func GetPlatformAuthorizationInfoAuthorizeStateEnum() PlatformAuthorizationInfoAuthorizeStateEnum {
	return PlatformAuthorizationInfoAuthorizeStateEnum{
		AUTHORIZED: PlatformAuthorizationInfoAuthorizeState{
			value: "AUTHORIZED",
		},
		UNAUTHORIZED: PlatformAuthorizationInfoAuthorizeState{
			value: "UNAUTHORIZED",
		},
	}
}

func (c PlatformAuthorizationInfoAuthorizeState) Value() string {
	return c.value
}

func (c PlatformAuthorizationInfoAuthorizeState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PlatformAuthorizationInfoAuthorizeState) UnmarshalJSON(b []byte) error {
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
