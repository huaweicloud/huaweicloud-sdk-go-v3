package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShootScript 表演脚本。
type ShootScript struct {

	// **参数解释**： 脚本类型，即视频制作的驱动方式 **约束限制**： 不涉及 **取值范围** * TEXT: 文本驱动，即通过TTS合成语音 * AUDIO: 语音驱动
	ScriptType *ShootScriptScriptType `json:"script_type,omitempty"`

	TextConfig *TextConfig `json:"text_config,omitempty"`

	// 语音驱动时，音频时长，单位秒。 > * 创建剧本时此参数可以不设置，音频文件上传成功后，通过更新剧本接口设置 > * 查询剧本详情时，返回音频时长，用于预估视频时长
	AudioDuration *float32 `json:"audio_duration,omitempty"`

	// 语音驱动时的动作配置。
	AudioDriveActionConfig *[]AudioDriveActionConfig `json:"audio_drive_action_config,omitempty"`

	// 语音驱动音频文件外部下载URL。  > * 只支持分身数字人视频制作  > * 需要先申请开通白名单后，才允许通过外部URL的音频文件来驱动分身数字人视频。  > * 音频文件需要存放在华为云OBS
	AudioDriveFileExternalUrl *string `json:"audio_drive_file_external_url,omitempty"`

	// 背景配置。
	BackgroundConfig *[]BackgroundConfigInfo `json:"background_config,omitempty"`

	// 图层配置。
	LayerConfig *[]LayerConfig `json:"layer_config,omitempty"`

	AudioConfig *AudioInfo `json:"audio_config,omitempty"`
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
