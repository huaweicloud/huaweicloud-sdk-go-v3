package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmSmartLiveRoomReq 确认直播间剧本请求
type ConfirmSmartLiveRoomReq struct {

	// 确认操作。 * confirm: 确认。 * reject: 拒绝。
	Action *ConfirmSmartLiveRoomReqAction `json:"action,omitempty"`

	// 剧本版本。从查询直播间详情接口中获取。
	ScriptVersion *string `json:"script_version,omitempty"`
}

func (o ConfirmSmartLiveRoomReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSmartLiveRoomReq struct{}"
	}

	return strings.Join([]string{"ConfirmSmartLiveRoomReq", string(data)}, " ")
}

type ConfirmSmartLiveRoomReqAction struct {
	value string
}

type ConfirmSmartLiveRoomReqActionEnum struct {
	CONFIRM ConfirmSmartLiveRoomReqAction
	REJECT  ConfirmSmartLiveRoomReqAction
}

func GetConfirmSmartLiveRoomReqActionEnum() ConfirmSmartLiveRoomReqActionEnum {
	return ConfirmSmartLiveRoomReqActionEnum{
		CONFIRM: ConfirmSmartLiveRoomReqAction{
			value: "confirm",
		},
		REJECT: ConfirmSmartLiveRoomReqAction{
			value: "reject",
		},
	}
}

func (c ConfirmSmartLiveRoomReqAction) Value() string {
	return c.value
}

func (c ConfirmSmartLiveRoomReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmSmartLiveRoomReqAction) UnmarshalJSON(b []byte) error {
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
