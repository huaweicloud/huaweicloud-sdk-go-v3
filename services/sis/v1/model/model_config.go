package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Config struct {

	// 支持语音的格式。  audio_format取值范围：  pcm16k16bit  16k16bit单通道录音数据。  pcm8k16bit   8k16bit单通道录音数据。  ulaw16k8bit  16k8bit ulaw 单通道录音数据。  ulaw8k8bit   8k8bit ulaw 单通道录音数据。  alaw16k8bit  16k8bit alaw 单通道录音数据。  alaw8k8bit   8k8bit alaw 单通道录音数据。  mp3  mp3格式音频。目前仅支持单通道的音频。  aac  aac格式音频。目前仅支持单通道的音频。  wav  带wav封装头的格式，从封装头中自动确定格式，目前仅支持8k/16k采样率、单通道、pcm, alaw, ulaw三种编码格式  amr  AMR窄带(8k) 压缩录音数据。  amrwb  AMR 宽带(16k) 压缩录音数据。  auto 由引擎自动判断音频数据的格式并解码，支持自动判断wav，mp3，amr/amrwb，aac，m4a，wma格式
	AudioFormat ConfigAudioFormat `json:"audio_format"`

	// 所使用的模型特征串。通常是 “语种_采样率_领域”的形式。  采样率需要与音频采样率保持一致。  当前支持如下模型特征串：  chinese_16k_general  支持采样率为16k的中文普通话语音识别，同时可识别一些简单的方言。格式仅支持pcm16k16bit、mp3、wav，区域仅支持cn-north-4。  chinese_16k_travel 支持采样率为16k的中文普通话语音识别，采用新一代端到端识别算法，并针对网约车质检场景进行了优化。格式仅支持pcm16k16bit、mp3、wav和aac，区域支持cn-east-3和cn-north-4（强烈推荐使用）。  chinese_8k_common  支持采样率为8k的中文普通话语音识别。  chinese_16k_common  支持采样率为16k的中文普通话语音识别。  sichuan_16k_common  支持采样率为16k的中文普通话与四川话方言识别。区域仅支持cn-north-4。  cantonese_16k_common  支持采样率为16k的粤语方言识别。区域仅支持cn-north-4。  shanghai_16k_common  支持采样率为16k的上海话方言识别，区域仅支持cn-north-4。  english_8k_common  支持采样率为16k的英文识别，区域仅支持cn-east-3，格式仅支持wav。  english_16k_common  支持采样率为16k的英文识别，区域仅支持cn-east-3，格式仅支持wav。
	Property ConfigProperty `json:"property"`

	// 表示是否在识别结果中添加标点，取值为“yes”和“no”，缺省为“no”。
	AddPunc *ConfigAddPunc `json:"add_punc,omitempty"`

	// 热词表id，不使用则不填写。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// 表示是否将语音中的数字识别为阿拉伯数字，取值为“yes” 和 “no”，缺省为“yes”。
	DigitNorm *ConfigDigitNorm `json:"digit_norm,omitempty"`

	// 表示是否在识别结果中输出分词结果信息，取值为“yes”和“no”，默认为“no”。
	NeedWordInfo *ConfigNeedWordInfo `json:"need_word_info,omitempty"`
}

func (o Config) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}

type ConfigAudioFormat struct {
	value string
}

type ConfigAudioFormatEnum struct {
	PCM16K16BIT ConfigAudioFormat
	PCM8K16BIT  ConfigAudioFormat
	ULAW16K8BIT ConfigAudioFormat
	ULAW8K8BIT  ConfigAudioFormat
	ALAW16K8BIT ConfigAudioFormat
	ALAW8K8BIT  ConfigAudioFormat
	MP3         ConfigAudioFormat
	AAC         ConfigAudioFormat
	WAV         ConfigAudioFormat
	AMR         ConfigAudioFormat
	AMRWB       ConfigAudioFormat
	AUTO        ConfigAudioFormat
}

func GetConfigAudioFormatEnum() ConfigAudioFormatEnum {
	return ConfigAudioFormatEnum{
		PCM16K16BIT: ConfigAudioFormat{
			value: "pcm16k16bit",
		},
		PCM8K16BIT: ConfigAudioFormat{
			value: "pcm8k16bit",
		},
		ULAW16K8BIT: ConfigAudioFormat{
			value: "ulaw16k8bit",
		},
		ULAW8K8BIT: ConfigAudioFormat{
			value: "ulaw8k8bit",
		},
		ALAW16K8BIT: ConfigAudioFormat{
			value: "alaw16k8bit",
		},
		ALAW8K8BIT: ConfigAudioFormat{
			value: "alaw8k8bit",
		},
		MP3: ConfigAudioFormat{
			value: "mp3",
		},
		AAC: ConfigAudioFormat{
			value: "aac",
		},
		WAV: ConfigAudioFormat{
			value: "wav",
		},
		AMR: ConfigAudioFormat{
			value: "amr",
		},
		AMRWB: ConfigAudioFormat{
			value: "amrwb",
		},
		AUTO: ConfigAudioFormat{
			value: "auto",
		},
	}
}

func (c ConfigAudioFormat) Value() string {
	return c.value
}

func (c ConfigAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigAudioFormat) UnmarshalJSON(b []byte) error {
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

type ConfigProperty struct {
	value string
}

type ConfigPropertyEnum struct {
	CHINESE_16K_GENERAL  ConfigProperty
	CHINESE_16K_TRAVEL   ConfigProperty
	CHINESE_8K_COMMON    ConfigProperty
	CHINESE_16K_COMMON   ConfigProperty
	SICHUAN_16K_COMMON   ConfigProperty
	CANTONESE_16K_COMMON ConfigProperty
	SHANGHAI_16K_COMMON  ConfigProperty
	ENGLISH_8K_COMMON    ConfigProperty
	ENGLISH_16K_COMMON   ConfigProperty
}

func GetConfigPropertyEnum() ConfigPropertyEnum {
	return ConfigPropertyEnum{
		CHINESE_16K_GENERAL: ConfigProperty{
			value: "chinese_16k_general",
		},
		CHINESE_16K_TRAVEL: ConfigProperty{
			value: "chinese_16k_travel",
		},
		CHINESE_8K_COMMON: ConfigProperty{
			value: "chinese_8k_common",
		},
		CHINESE_16K_COMMON: ConfigProperty{
			value: "chinese_16k_common",
		},
		SICHUAN_16K_COMMON: ConfigProperty{
			value: "sichuan_16k_common",
		},
		CANTONESE_16K_COMMON: ConfigProperty{
			value: "cantonese_16k_common",
		},
		SHANGHAI_16K_COMMON: ConfigProperty{
			value: "shanghai_16k_common",
		},
		ENGLISH_8K_COMMON: ConfigProperty{
			value: "english_8k_common",
		},
		ENGLISH_16K_COMMON: ConfigProperty{
			value: "english_16k_common",
		},
	}
}

func (c ConfigProperty) Value() string {
	return c.value
}

func (c ConfigProperty) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigProperty) UnmarshalJSON(b []byte) error {
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

type ConfigAddPunc struct {
	value string
}

type ConfigAddPuncEnum struct {
	YES ConfigAddPunc
	NO  ConfigAddPunc
}

func GetConfigAddPuncEnum() ConfigAddPuncEnum {
	return ConfigAddPuncEnum{
		YES: ConfigAddPunc{
			value: "yes",
		},
		NO: ConfigAddPunc{
			value: "no",
		},
	}
}

func (c ConfigAddPunc) Value() string {
	return c.value
}

func (c ConfigAddPunc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigAddPunc) UnmarshalJSON(b []byte) error {
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

type ConfigDigitNorm struct {
	value string
}

type ConfigDigitNormEnum struct {
	YES ConfigDigitNorm
	NO  ConfigDigitNorm
}

func GetConfigDigitNormEnum() ConfigDigitNormEnum {
	return ConfigDigitNormEnum{
		YES: ConfigDigitNorm{
			value: "yes",
		},
		NO: ConfigDigitNorm{
			value: "no",
		},
	}
}

func (c ConfigDigitNorm) Value() string {
	return c.value
}

func (c ConfigDigitNorm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigDigitNorm) UnmarshalJSON(b []byte) error {
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

type ConfigNeedWordInfo struct {
	value string
}

type ConfigNeedWordInfoEnum struct {
	YES ConfigNeedWordInfo
	NO  ConfigNeedWordInfo
}

func GetConfigNeedWordInfoEnum() ConfigNeedWordInfoEnum {
	return ConfigNeedWordInfoEnum{
		YES: ConfigNeedWordInfo{
			value: "yes",
		},
		NO: ConfigNeedWordInfo{
			value: "no",
		},
	}
}

func (c ConfigNeedWordInfo) Value() string {
	return c.value
}

func (c ConfigNeedWordInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigNeedWordInfo) UnmarshalJSON(b []byte) error {
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
