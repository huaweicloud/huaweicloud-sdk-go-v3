package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchConfirmLiveCommandsReq 批量确认命令。
type BatchConfirmLiveCommandsReq struct {

	// 命令名称。 - REWRITE_PLAY_SCRIPT: 动态编辑未播放剧本。 - REWRITE_INTERACTION_RULES: 动态修改互动规则。
	Command *BatchConfirmLiveCommandsReqCommand `json:"command,omitempty"`

	// 确认操作。 * confirm: 确认。 * reject: 拒绝。
	Action *BatchConfirmLiveCommandsReqAction `json:"action,omitempty"`

	// 命令ID列表。不填则为全部未播放的插入命令均清理。
	CommandIds *[]string `json:"command_ids,omitempty"`
}

func (o BatchConfirmLiveCommandsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmLiveCommandsReq struct{}"
	}

	return strings.Join([]string{"BatchConfirmLiveCommandsReq", string(data)}, " ")
}

type BatchConfirmLiveCommandsReqCommand struct {
	value string
}

type BatchConfirmLiveCommandsReqCommandEnum struct {
	REWRITE_PLAY_SCRIPT       BatchConfirmLiveCommandsReqCommand
	REWRITE_INTERACTION_RULES BatchConfirmLiveCommandsReqCommand
}

func GetBatchConfirmLiveCommandsReqCommandEnum() BatchConfirmLiveCommandsReqCommandEnum {
	return BatchConfirmLiveCommandsReqCommandEnum{
		REWRITE_PLAY_SCRIPT: BatchConfirmLiveCommandsReqCommand{
			value: "REWRITE_PLAY_SCRIPT",
		},
		REWRITE_INTERACTION_RULES: BatchConfirmLiveCommandsReqCommand{
			value: "REWRITE_INTERACTION_RULES",
		},
	}
}

func (c BatchConfirmLiveCommandsReqCommand) Value() string {
	return c.value
}

func (c BatchConfirmLiveCommandsReqCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchConfirmLiveCommandsReqCommand) UnmarshalJSON(b []byte) error {
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

type BatchConfirmLiveCommandsReqAction struct {
	value string
}

type BatchConfirmLiveCommandsReqActionEnum struct {
	CONFIRM BatchConfirmLiveCommandsReqAction
	REJECT  BatchConfirmLiveCommandsReqAction
}

func GetBatchConfirmLiveCommandsReqActionEnum() BatchConfirmLiveCommandsReqActionEnum {
	return BatchConfirmLiveCommandsReqActionEnum{
		CONFIRM: BatchConfirmLiveCommandsReqAction{
			value: "confirm",
		},
		REJECT: BatchConfirmLiveCommandsReqAction{
			value: "reject",
		},
	}
}

func (c BatchConfirmLiveCommandsReqAction) Value() string {
	return c.value
}

func (c BatchConfirmLiveCommandsReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchConfirmLiveCommandsReqAction) UnmarshalJSON(b []byte) error {
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
