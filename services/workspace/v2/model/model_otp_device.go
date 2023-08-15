package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type OtpDevice struct {

	// 用户otp 信息id。
	Id *string `json:"id,omitempty"`

	// 用户id。
	UserId *string `json:"user_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户otp设备状态 UNREGISTER: 未绑定 REGISTERED：已绑定
	Status *OtpDeviceStatus `json:"status,omitempty"`

	// 用户otp设备绑定时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o OtpDevice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OtpDevice struct{}"
	}

	return strings.Join([]string{"OtpDevice", string(data)}, " ")
}

type OtpDeviceStatus struct {
	value string
}

type OtpDeviceStatusEnum struct {
	UNREGISTER OtpDeviceStatus
	REGISTERED OtpDeviceStatus
}

func GetOtpDeviceStatusEnum() OtpDeviceStatusEnum {
	return OtpDeviceStatusEnum{
		UNREGISTER: OtpDeviceStatus{
			value: "UNREGISTER",
		},
		REGISTERED: OtpDeviceStatus{
			value: "REGISTERED",
		},
	}
}

func (c OtpDeviceStatus) Value() string {
	return c.value
}

func (c OtpDeviceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OtpDeviceStatus) UnmarshalJSON(b []byte) error {
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
