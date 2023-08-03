package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RtcUserInfo 接入RTC的用户信息。
type RtcUserInfo struct {

	// 接入RTC的用户类型。 * CAPTURE：直播助手，将摄像头获取视频流推送到RTC房间 * ANIMATION：VDS服务，从RTC房间拉视频流生成动作数据 * RENDER：渲染服务，将动作数据渲染成数字人动画 * PLAYER：普通观看方，可选择原始视频流或者数字人动画视频流观看
	UserType *RtcUserInfoUserType `json:"user_type,omitempty"`

	// RTC用户ID。
	UserId *string `json:"user_id,omitempty"`

	// RTC鉴权token。
	Signature *string `json:"signature,omitempty"`

	// 有效期。时间戳。  单位：秒。
	Ctime *int64 `json:"ctime,omitempty"`
}

func (o RtcUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcUserInfo struct{}"
	}

	return strings.Join([]string{"RtcUserInfo", string(data)}, " ")
}

type RtcUserInfoUserType struct {
	value string
}

type RtcUserInfoUserTypeEnum struct {
	CAPTURE   RtcUserInfoUserType
	ANIMATION RtcUserInfoUserType
	RENDER    RtcUserInfoUserType
	PLAYER    RtcUserInfoUserType
}

func GetRtcUserInfoUserTypeEnum() RtcUserInfoUserTypeEnum {
	return RtcUserInfoUserTypeEnum{
		CAPTURE: RtcUserInfoUserType{
			value: "CAPTURE",
		},
		ANIMATION: RtcUserInfoUserType{
			value: "ANIMATION",
		},
		RENDER: RtcUserInfoUserType{
			value: "RENDER",
		},
		PLAYER: RtcUserInfoUserType{
			value: "PLAYER",
		},
	}
}

func (c RtcUserInfoUserType) Value() string {
	return c.value
}

func (c RtcUserInfoUserType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RtcUserInfoUserType) UnmarshalJSON(b []byte) error {
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
