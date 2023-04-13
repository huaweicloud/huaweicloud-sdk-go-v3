package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 控制命令。
type ControlDigitalHumanLiveReq struct {

	// 命令名称。 - BODY_POS_RESET: 视觉驱动复位 - HIPS_POS_ADJUST: 模型位移调整 - EYE_POS: 眼神锁定/解锁 - SKELETON_ROTATION_ADJUST: 骨骼旋转 - LOCK_SKELETONS: 骨骼锁定 - UNLOCK_SKELETONS: 骨骼解锁
	Command ControlDigitalHumanLiveReqCommand `json:"command"`

	// 命令参数。
	Params *interface{} `json:"params,omitempty"`
}

func (o ControlDigitalHumanLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlDigitalHumanLiveReq struct{}"
	}

	return strings.Join([]string{"ControlDigitalHumanLiveReq", string(data)}, " ")
}

type ControlDigitalHumanLiveReqCommand struct {
	value string
}

type ControlDigitalHumanLiveReqCommandEnum struct {
	BODY_POS_RESET           ControlDigitalHumanLiveReqCommand
	HIPS_POS_ADJUST          ControlDigitalHumanLiveReqCommand
	EYE_POS                  ControlDigitalHumanLiveReqCommand
	SKELETON_ROTATION_ADJUST ControlDigitalHumanLiveReqCommand
	LOCK_SKELETONS           ControlDigitalHumanLiveReqCommand
	UNLOCK_SKELETONS         ControlDigitalHumanLiveReqCommand
}

func GetControlDigitalHumanLiveReqCommandEnum() ControlDigitalHumanLiveReqCommandEnum {
	return ControlDigitalHumanLiveReqCommandEnum{
		BODY_POS_RESET: ControlDigitalHumanLiveReqCommand{
			value: "BODY_POS_RESET",
		},
		HIPS_POS_ADJUST: ControlDigitalHumanLiveReqCommand{
			value: "HIPS_POS_ADJUST",
		},
		EYE_POS: ControlDigitalHumanLiveReqCommand{
			value: "EYE_POS",
		},
		SKELETON_ROTATION_ADJUST: ControlDigitalHumanLiveReqCommand{
			value: "SKELETON_ROTATION_ADJUST",
		},
		LOCK_SKELETONS: ControlDigitalHumanLiveReqCommand{
			value: "LOCK_SKELETONS",
		},
		UNLOCK_SKELETONS: ControlDigitalHumanLiveReqCommand{
			value: "UNLOCK_SKELETONS",
		},
	}
}

func (c ControlDigitalHumanLiveReqCommand) Value() string {
	return c.value
}

func (c ControlDigitalHumanLiveReqCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ControlDigitalHumanLiveReqCommand) UnmarshalJSON(b []byte) error {
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
