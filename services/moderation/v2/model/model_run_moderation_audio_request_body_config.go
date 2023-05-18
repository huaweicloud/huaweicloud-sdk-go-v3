package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 语音文件配置信息
type RunModerationAudioRequestBodyConfig struct {

	// 支持的语音格式。 枚举值： - pcm16k16bit - pcm8k16bit - ulaw16k8bit - ulaw8k8bit - alaw16k8bit - alaw8k8bit - mp3 - aac - wav - amr - amrwb
	Format RunModerationAudioRequestBodyConfigFormat `json:"format"`

	// 所使用的模型特征串。通常是 “语种_采样率_领域”的形式。 采样率需要与音频采样率保持一致。 当前支持如下模型特征串：   chinese_8k_common   chinese_16k_common
	Property RunModerationAudioRequestBodyConfigProperty `json:"property"`
}

func (o RunModerationAudioRequestBodyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModerationAudioRequestBodyConfig struct{}"
	}

	return strings.Join([]string{"RunModerationAudioRequestBodyConfig", string(data)}, " ")
}

type RunModerationAudioRequestBodyConfigFormat struct {
	value string
}

type RunModerationAudioRequestBodyConfigFormatEnum struct {
	PCM16K16BIT RunModerationAudioRequestBodyConfigFormat
	PCM8K16BIT  RunModerationAudioRequestBodyConfigFormat
	ULAW16K8BIT RunModerationAudioRequestBodyConfigFormat
	ULAW8K8BIT  RunModerationAudioRequestBodyConfigFormat
	ALAW16K8BIT RunModerationAudioRequestBodyConfigFormat
	ALAW8K8BIT  RunModerationAudioRequestBodyConfigFormat
	MP3         RunModerationAudioRequestBodyConfigFormat
	AAC         RunModerationAudioRequestBodyConfigFormat
	WAV         RunModerationAudioRequestBodyConfigFormat
	AMR         RunModerationAudioRequestBodyConfigFormat
	AMRWB       RunModerationAudioRequestBodyConfigFormat
}

func GetRunModerationAudioRequestBodyConfigFormatEnum() RunModerationAudioRequestBodyConfigFormatEnum {
	return RunModerationAudioRequestBodyConfigFormatEnum{
		PCM16K16BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "pcm16k16bit",
		},
		PCM8K16BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "pcm8k16bit",
		},
		ULAW16K8BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "ulaw16k8bit",
		},
		ULAW8K8BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "ulaw8k8bit",
		},
		ALAW16K8BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "alaw16k8bit",
		},
		ALAW8K8BIT: RunModerationAudioRequestBodyConfigFormat{
			value: "alaw8k8bit",
		},
		MP3: RunModerationAudioRequestBodyConfigFormat{
			value: "mp3",
		},
		AAC: RunModerationAudioRequestBodyConfigFormat{
			value: "aac",
		},
		WAV: RunModerationAudioRequestBodyConfigFormat{
			value: "wav",
		},
		AMR: RunModerationAudioRequestBodyConfigFormat{
			value: "amr",
		},
		AMRWB: RunModerationAudioRequestBodyConfigFormat{
			value: "amrwb",
		},
	}
}

func (c RunModerationAudioRequestBodyConfigFormat) Value() string {
	return c.value
}

func (c RunModerationAudioRequestBodyConfigFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunModerationAudioRequestBodyConfigFormat) UnmarshalJSON(b []byte) error {
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

type RunModerationAudioRequestBodyConfigProperty struct {
	value string
}

type RunModerationAudioRequestBodyConfigPropertyEnum struct {
	CHINESE_8K_COMMON  RunModerationAudioRequestBodyConfigProperty
	CHINESE_16K_COMMON RunModerationAudioRequestBodyConfigProperty
}

func GetRunModerationAudioRequestBodyConfigPropertyEnum() RunModerationAudioRequestBodyConfigPropertyEnum {
	return RunModerationAudioRequestBodyConfigPropertyEnum{
		CHINESE_8K_COMMON: RunModerationAudioRequestBodyConfigProperty{
			value: "chinese_8k_common",
		},
		CHINESE_16K_COMMON: RunModerationAudioRequestBodyConfigProperty{
			value: "chinese_16k_common",
		},
	}
}

func (c RunModerationAudioRequestBodyConfigProperty) Value() string {
	return c.value
}

func (c RunModerationAudioRequestBodyConfigProperty) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunModerationAudioRequestBodyConfigProperty) UnmarshalJSON(b []byte) error {
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
