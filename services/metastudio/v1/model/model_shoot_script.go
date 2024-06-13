package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShootScript 表演脚本。
type ShootScript struct {

	// 脚本类型，即视频制作的驱动方式。默认TEXT * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	ScriptType *ShootScriptScriptType `json:"script_type,omitempty"`

	TextConfig *TextConfig `json:"text_config,omitempty"`

	// 背景配置。
	BackgroundConfig *[]BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`
}

func (o ShootScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScript struct{}"
	}

	return strings.Join([]string{"ShootScript", string(data)}, " ")
}

type ShootScriptScriptType struct {
	value string
}

type ShootScriptScriptTypeEnum struct {
	TEXT  ShootScriptScriptType
	AUDIO ShootScriptScriptType
}

func GetShootScriptScriptTypeEnum() ShootScriptScriptTypeEnum {
	return ShootScriptScriptTypeEnum{
		TEXT: ShootScriptScriptType{
			value: "TEXT",
		},
		AUDIO: ShootScriptScriptType{
			value: "AUDIO",
		},
	}
}

func (c ShootScriptScriptType) Value() string {
	return c.value
}

func (c ShootScriptScriptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShootScriptScriptType) UnmarshalJSON(b []byte) error {
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
