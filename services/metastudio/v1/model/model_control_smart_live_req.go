package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ControlSmartLiveReq 控制命令。
type ControlSmartLiveReq struct {

	// 命令名称。 - INSERT_PLAY_SCRIPT: 插入表演脚本。用于互动回复。数字人不变，背景不变。params结构定义：ShootScript - REWRITE_PLAY_SCRIPT: 动态编辑未播放剧本。params结构定义：scence_scripts - INSERT_PLAY_ADUIO: 插入驱动音频。用于音频直接驱动。数字人不变，背景不变。params结构定义：PlayAudioInfo
	Command ControlSmartLiveReqCommand `json:"command"`

	// 命令参数。
	Params *interface{} `json:"params,omitempty"`
}

func (o ControlSmartLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlSmartLiveReq struct{}"
	}

	return strings.Join([]string{"ControlSmartLiveReq", string(data)}, " ")
}

type ControlSmartLiveReqCommand struct {
	value string
}

type ControlSmartLiveReqCommandEnum struct {
	INSERT_PLAY_SCRIPT  ControlSmartLiveReqCommand
	REWRITE_PLAY_SCRIPT ControlSmartLiveReqCommand
	INSERT_PLAY_AUDIO   ControlSmartLiveReqCommand
}

func GetControlSmartLiveReqCommandEnum() ControlSmartLiveReqCommandEnum {
	return ControlSmartLiveReqCommandEnum{
		INSERT_PLAY_SCRIPT: ControlSmartLiveReqCommand{
			value: "INSERT_PLAY_SCRIPT",
		},
		REWRITE_PLAY_SCRIPT: ControlSmartLiveReqCommand{
			value: "REWRITE_PLAY_SCRIPT",
		},
		INSERT_PLAY_AUDIO: ControlSmartLiveReqCommand{
			value: "INSERT_PLAY_AUDIO",
		},
	}
}

func (c ControlSmartLiveReqCommand) Value() string {
	return c.value
}

func (c ControlSmartLiveReqCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ControlSmartLiveReqCommand) UnmarshalJSON(b []byte) error {
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
