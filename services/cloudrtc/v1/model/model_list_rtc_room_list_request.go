package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListRtcRoomListRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty" xml:"X-Sdk-Date"`

	// 使用AK/SK方式认证时必选，携带项目ID信息，与路径参数中的项目ID相同。
	XProjectId *string `json:"X-Project-Id,omitempty" xml:"X-Project-Id"`

	// 应用标识
	App string `json:"app" xml:"app"`

	// 房间id
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	// 房间状态，取值如下： - RUNNING：开启中 - CLOSED：已关闭
	State *ListRtcRoomListRequestState `json:"state,omitempty" xml:"state"`

	// 查询起始时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T06:00:00Z，不写默认读取过去1小时数据数据。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 查询结束时间。UTC时间，格式：YYYY-MM-DDThh:mm:ssZ，如2020-04-23T07:00:00Z，不写默认为当前时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 查询结果条数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListRtcRoomListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRoomListRequest struct{}"
	}

	return strings.Join([]string{"ListRtcRoomListRequest", string(data)}, " ")
}

type ListRtcRoomListRequestState struct {
	value string
}

type ListRtcRoomListRequestStateEnum struct {
	RUNNING ListRtcRoomListRequestState
	CLOSED  ListRtcRoomListRequestState
}

func GetListRtcRoomListRequestStateEnum() ListRtcRoomListRequestStateEnum {
	return ListRtcRoomListRequestStateEnum{
		RUNNING: ListRtcRoomListRequestState{
			value: "RUNNING",
		},
		CLOSED: ListRtcRoomListRequestState{
			value: "CLOSED",
		},
	}
}

func (c ListRtcRoomListRequestState) Value() string {
	return c.value
}

func (c ListRtcRoomListRequestState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRtcRoomListRequestState) UnmarshalJSON(b []byte) error {
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
