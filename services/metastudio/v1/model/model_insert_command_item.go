package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InsertCommandItem 插入播放命令列表
type InsertCommandItem struct {

	// 控制命令ID
	CommandId *string `json:"command_id,omitempty"`

	// 命令名称。 - INSERT_PLAY_SCRIPT: 插入表演脚本。用于互动回复。数字人不变，背景不变。params结构定义：[PlayTextInfo](metastudio_02_0014.xml) - INSERT_PLAY_AUDIO:插入驱动音频。用于音频直接驱动。数字人不变，背景不变。params结构定义：[PlayAudioInfo](metastudio_02_0014.xml) - REWRITE_INTERACTION_RULES: 修改互动规则。params结构定义：[PlayAudioInfo](metastudio_02_0014.xml)
	Command *InsertCommandItemCommand `json:"command,omitempty"`

	// 命令参数。
	Params *interface{} `json:"params,omitempty"`

	// 命令来源。 * EXTERNAL： 外部命令 * AUTO: 系统自动触发
	Source *InsertCommandItemSource `json:"source,omitempty"`
}

func (o InsertCommandItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsertCommandItem struct{}"
	}

	return strings.Join([]string{"InsertCommandItem", string(data)}, " ")
}

type InsertCommandItemCommand struct {
	value string
}

type InsertCommandItemCommandEnum struct {
	INSERT_PLAY_SCRIPT        InsertCommandItemCommand
	INSERT_PLAY_AUDIO         InsertCommandItemCommand
	REWRITE_INTERACTION_RULES InsertCommandItemCommand
}

func GetInsertCommandItemCommandEnum() InsertCommandItemCommandEnum {
	return InsertCommandItemCommandEnum{
		INSERT_PLAY_SCRIPT: InsertCommandItemCommand{
			value: "INSERT_PLAY_SCRIPT",
		},
		INSERT_PLAY_AUDIO: InsertCommandItemCommand{
			value: "INSERT_PLAY_AUDIO",
		},
		REWRITE_INTERACTION_RULES: InsertCommandItemCommand{
			value: "REWRITE_INTERACTION_RULES",
		},
	}
}

func (c InsertCommandItemCommand) Value() string {
	return c.value
}

func (c InsertCommandItemCommand) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InsertCommandItemCommand) UnmarshalJSON(b []byte) error {
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

type InsertCommandItemSource struct {
	value string
}

type InsertCommandItemSourceEnum struct {
	EXTERNAL InsertCommandItemSource
	AUTO     InsertCommandItemSource
}

func GetInsertCommandItemSourceEnum() InsertCommandItemSourceEnum {
	return InsertCommandItemSourceEnum{
		EXTERNAL: InsertCommandItemSource{
			value: "EXTERNAL",
		},
		AUTO: InsertCommandItemSource{
			value: "AUTO",
		},
	}
}

func (c InsertCommandItemSource) Value() string {
	return c.value
}

func (c InsertCommandItemSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InsertCommandItemSource) UnmarshalJSON(b []byte) error {
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
