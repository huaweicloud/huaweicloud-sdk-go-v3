package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RtcUser struct {

	// 域名
	Domain *string `json:"domain,omitempty"`

	// 应用标识
	App *string `json:"app,omitempty"`

	// 房间ID
	RoomId *string `json:"room_id,omitempty"`

	// 用户id
	Uid *string `json:"uid,omitempty"`

	// 会话id
	Session *string `json:"session,omitempty"`

	// 用户状态   - FAIL： 加入失败   - ONLINE：在线   - OFFLINE：离开
	State *RtcUserState `json:"state,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 用户接入IP
	Ip *string `json:"ip,omitempty"`

	// 用户接入IP所在省份
	Region *string `json:"region,omitempty"`

	// 用户接入IP所在运营商
	Isp *string `json:"isp,omitempty"`

	// 用户设备型号
	DeviceModel *string `json:"device_model,omitempty"`

	// 用户设备平台
	Platform *string `json:"platform,omitempty"`

	// 用户sdk版本
	Sdk *string `json:"sdk,omitempty"`

	// 用户加入房间时间。格式为：YYYY-MM-DDThh:mm:ssZ
	JoinTime *string `json:"join_time,omitempty"`

	// 用户离开房间时间。格式为：YYYY-MM-DDThh:mm:ssZ，若用户未离开，则返回 “-”
	LeaveTime *string `json:"leave_time,omitempty"`
}

func (o RtcUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcUser struct{}"
	}

	return strings.Join([]string{"RtcUser", string(data)}, " ")
}

type RtcUserState struct {
	value string
}

type RtcUserStateEnum struct {
	FAIL    RtcUserState
	ONLINE  RtcUserState
	OFFLINE RtcUserState
}

func GetRtcUserStateEnum() RtcUserStateEnum {
	return RtcUserStateEnum{
		FAIL: RtcUserState{
			value: "FAIL",
		},
		ONLINE: RtcUserState{
			value: "ONLINE",
		},
		OFFLINE: RtcUserState{
			value: "OFFLINE",
		},
	}
}

func (c RtcUserState) Value() string {
	return c.value
}

func (c RtcUserState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RtcUserState) UnmarshalJSON(b []byte) error {
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
