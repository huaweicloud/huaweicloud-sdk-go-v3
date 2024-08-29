package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRobotRequest Request Object
type ListRobotRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 智能交互对话房间ID。
	RoomId *string `json:"room_id,omitempty"`

	// 交互对接类型  * LIVE:直播交互  * CHAT:智能交互
	RobotType *ListRobotRequestRobotType `json:"robot_type,omitempty"`
}

func (o ListRobotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRobotRequest struct{}"
	}

	return strings.Join([]string{"ListRobotRequest", string(data)}, " ")
}

type ListRobotRequestRobotType struct {
	value string
}

type ListRobotRequestRobotTypeEnum struct {
	LIVE ListRobotRequestRobotType
	CHAT ListRobotRequestRobotType
}

func GetListRobotRequestRobotTypeEnum() ListRobotRequestRobotTypeEnum {
	return ListRobotRequestRobotTypeEnum{
		LIVE: ListRobotRequestRobotType{
			value: "LIVE",
		},
		CHAT: ListRobotRequestRobotType{
			value: "CHAT",
		},
	}
}

func (c ListRobotRequestRobotType) Value() string {
	return c.value
}

func (c ListRobotRequestRobotType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRobotRequestRobotType) UnmarshalJSON(b []byte) error {
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
