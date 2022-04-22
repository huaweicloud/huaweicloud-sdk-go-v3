package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type AudioConfig struct {

	// 语音的格式。不填写此字段，则默认为auto。注意音频不论何种格式，均要求采样率在16000Hz以上。  auto  自动判断，系统会自动判断并支持WAV（内部支持pcm/ulaw/alaw编码格式）、MP3、M4A、ogg-opus、AMR等格式。推荐使用此取值。  wav  wav格式。  aac  aac格式。  mp3  mp3格式。  amr  amr格式。  m4a  m4a格式。  opus  ogg-opus格式。
	AudioFormat *AudioConfigAudioFormat `json:"audio_format,omitempty"`

	// 评测语言  en_gb  英语-英式口音。
	Language AudioConfigLanguage `json:"language"`

	// 评测模式  word 单词模式  sentence 句子模式
	Mode AudioConfigMode `json:"mode"`
}

func (o AudioConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioConfig struct{}"
	}

	return strings.Join([]string{"AudioConfig", string(data)}, " ")
}

type AudioConfigAudioFormat struct {
	value string
}

type AudioConfigAudioFormatEnum struct {
	AUTO AudioConfigAudioFormat
	WAV  AudioConfigAudioFormat
	AAC  AudioConfigAudioFormat
	MP3  AudioConfigAudioFormat
	AMR  AudioConfigAudioFormat
	M4A  AudioConfigAudioFormat
	OPUS AudioConfigAudioFormat
}

func GetAudioConfigAudioFormatEnum() AudioConfigAudioFormatEnum {
	return AudioConfigAudioFormatEnum{
		AUTO: AudioConfigAudioFormat{
			value: "auto",
		},
		WAV: AudioConfigAudioFormat{
			value: "wav",
		},
		AAC: AudioConfigAudioFormat{
			value: "aac",
		},
		MP3: AudioConfigAudioFormat{
			value: "mp3",
		},
		AMR: AudioConfigAudioFormat{
			value: "amr",
		},
		M4A: AudioConfigAudioFormat{
			value: "m4a",
		},
		OPUS: AudioConfigAudioFormat{
			value: "opus",
		},
	}
}

func (c AudioConfigAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioConfigAudioFormat) UnmarshalJSON(b []byte) error {
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

type AudioConfigLanguage struct {
	value string
}

type AudioConfigLanguageEnum struct {
	EN_GB AudioConfigLanguage
}

func GetAudioConfigLanguageEnum() AudioConfigLanguageEnum {
	return AudioConfigLanguageEnum{
		EN_GB: AudioConfigLanguage{
			value: "en_gb",
		},
	}
}

func (c AudioConfigLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioConfigLanguage) UnmarshalJSON(b []byte) error {
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

type AudioConfigMode struct {
	value string
}

type AudioConfigModeEnum struct {
	WORD     AudioConfigMode
	SENTENCE AudioConfigMode
}

func GetAudioConfigModeEnum() AudioConfigModeEnum {
	return AudioConfigModeEnum{
		WORD: AudioConfigMode{
			value: "word",
		},
		SENTENCE: AudioConfigMode{
			value: "sentence",
		},
	}
}

func (c AudioConfigMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AudioConfigMode) UnmarshalJSON(b []byte) error {
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
