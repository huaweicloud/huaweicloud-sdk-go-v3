package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmSmarLiveRoomReq 确认直播间剧本请求
type ConfirmSmarLiveRoomReq struct {

	// 确认操作。 * confirm: 确认。 * reject: 拒绝。
	Action *ConfirmSmarLiveRoomReqAction `json:"action,omitempty"`

	// 剧本版本。从查询直播间详情接口中获取。
	ScriptVersion *string `json:"script_version,omitempty"`
}

func (o ConfirmSmarLiveRoomReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmSmarLiveRoomReq struct{}"
	}

	return strings.Join([]string{"ConfirmSmarLiveRoomReq", string(data)}, " ")
}

type ConfirmSmarLiveRoomReqAction struct {
	value string
}

type ConfirmSmarLiveRoomReqActionEnum struct {
	CONFIRM ConfirmSmarLiveRoomReqAction
	REJECT  ConfirmSmarLiveRoomReqAction
}

func GetConfirmSmarLiveRoomReqActionEnum() ConfirmSmarLiveRoomReqActionEnum {
	return ConfirmSmarLiveRoomReqActionEnum{
		CONFIRM: ConfirmSmarLiveRoomReqAction{
			value: "confirm",
		},
		REJECT: ConfirmSmarLiveRoomReqAction{
			value: "reject",
		},
	}
}

func (c ConfirmSmarLiveRoomReqAction) Value() string {
	return c.value
}

func (c ConfirmSmarLiveRoomReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmSmarLiveRoomReqAction) UnmarshalJSON(b []byte) error {
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
