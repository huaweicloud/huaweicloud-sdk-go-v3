package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RtcUser struct {

	// 域名
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 应用标识
	App *string `json:"app,omitempty" xml:"app"`

	// 房间ID
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	// 用户id
	Uid *string `json:"uid,omitempty" xml:"uid"`

	// 会话id
	Session *string `json:"session,omitempty" xml:"session"`

	// 用户状态   - FAIL： 加入失败   - ONLINE：在线   - OFFLINE：离开
	State *RtcUserState `json:"state,omitempty" xml:"state"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 用户接入IP
	Ip *string `json:"ip,omitempty" xml:"ip"`

	// 用户接入IP所在省份
	Region *string `json:"region,omitempty" xml:"region"`

	// 用户接入IP所在运营商
	Isp *string `json:"isp,omitempty" xml:"isp"`

	// 用户设备型号
	DeviceModel *string `json:"device_model,omitempty" xml:"device_model"`

	// 用户设备平台
	Platform *string `json:"platform,omitempty" xml:"platform"`

	// 用户sdk版本
	Sdk *string `json:"sdk,omitempty" xml:"sdk"`

	// 用户加入房间时间。格式为：YYYY-MM-DDThh:mm:ssZ
	JoinTime *string `json:"join_time,omitempty" xml:"join_time"`

	// 用户离开房间时间。格式为：YYYY-MM-DDThh:mm:ssZ，若用户未离开，则返回 “-”
	LeaveTime *string `json:"leave_time,omitempty" xml:"leave_time"`
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
