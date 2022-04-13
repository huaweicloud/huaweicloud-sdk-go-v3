package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type RtcServerRoomInfo struct {
	// 域名

	Domain *string `json:"domain,omitempty"`
	// 应用标识

	App *string `json:"app,omitempty"`
	// 房间ID

	RoomId *string `json:"room_id,omitempty"`
	// 房间状态，取值如下：  - RUNNING：开启中  - CLOSED：已关闭

	State *RtcServerRoomInfoState `json:"state,omitempty"`
	// 房间持续时长

	Duration *int32 `json:"duration,omitempty"`
	// 房间开始时间，即第一个用户加入房间时间，UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z

	StartTime *string `json:"start_time,omitempty"`
	// 房间关闭时间，即最后一个room_uuid关闭的时间，UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z，若房间未关闭，则返回 “-”

	EndTime *string `json:"end_time,omitempty"`
}

func (o RtcServerRoomInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcServerRoomInfo struct{}"
	}

	return strings.Join([]string{"RtcServerRoomInfo", string(data)}, " ")
}

type RtcServerRoomInfoState struct {
	value string
}

type RtcServerRoomInfoStateEnum struct {
	RUNNING RtcServerRoomInfoState
	CLOSED  RtcServerRoomInfoState
}

func GetRtcServerRoomInfoStateEnum() RtcServerRoomInfoStateEnum {
	return RtcServerRoomInfoStateEnum{
		RUNNING: RtcServerRoomInfoState{
			value: "RUNNING",
		},
		CLOSED: RtcServerRoomInfoState{
			value: "CLOSED",
		},
	}
}

func (c RtcServerRoomInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RtcServerRoomInfoState) UnmarshalJSON(b []byte) error {
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
