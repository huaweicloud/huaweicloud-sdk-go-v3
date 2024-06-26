package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShootScriptDetail struct {

	// 脚本类型，即视频制作的驱动方式。默认TEXT * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	ScriptType *ShootScriptDetailScriptType `json:"script_type,omitempty"`

	TextConfig *TextConfig `json:"text_config,omitempty"`

	// 背景配置。
	BackgroundConfig *[]BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	// 剧本场景缩略图url。
	ThumbnailUrl *string `json:"thumbnail_url,omitempty"`
}

func (o ShootScriptDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptDetail struct{}"
	}

	return strings.Join([]string{"ShootScriptDetail", string(data)}, " ")
}

type ShootScriptDetailScriptType struct {
	value string
}

type ShootScriptDetailScriptTypeEnum struct {
	TEXT  ShootScriptDetailScriptType
	AUDIO ShootScriptDetailScriptType
}

func GetShootScriptDetailScriptTypeEnum() ShootScriptDetailScriptTypeEnum {
	return ShootScriptDetailScriptTypeEnum{
		TEXT: ShootScriptDetailScriptType{
			value: "TEXT",
		},
		AUDIO: ShootScriptDetailScriptType{
			value: "AUDIO",
		},
	}
}

func (c ShootScriptDetailScriptType) Value() string {
	return c.value
}

func (c ShootScriptDetailScriptType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShootScriptDetailScriptType) UnmarshalJSON(b []byte) error {
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
