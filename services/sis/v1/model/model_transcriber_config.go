package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TranscriberConfig struct {

	// 支持语音的格式。  audioformat取值范围:  auto  自动判断，系统会自动判断并支持WAV（内部支持pcm/ulaw/alaw/adpcm编码格式）、MP3、M4A、ogg-speex、ogg-opus、AMR、wma等格式，相应的文件后缀名为.wav, .mp3, .m4a, .spx, .opus, .amr 和.wma, 不区分大小写。支持双声道的音频。  pcm16k16bit  16k16bit裸音频录音数据。  pcm8k16bit   8k16bit裸音频录音数据。  ulaw16k8bit  16k8bit ulaw 裸音频录音数据。  ulaw8k8bit   8k8bit ulaw 裸音频录音数据。  alaw16k8bit  16k8bit alaw 裸音频录音数据。  alaw8k8bit   8k8bit alaw 裸音频录音数据。  vox8k4bit    8k4bit vox 裸音频录音数据。  v3_8k4bit    8k4bit v3 裸音频录音数据。
	AudioFormat *TranscriberConfigAudioFormat `json:"audio_format,omitempty"`

	// 所使用的模型特征串。通常是“语种_采样率_领域”的形式，例如chinese_8k_common。  采样率需要与音频采样率保持一致。  当前支持如下模型特征串：  chinese_8k_general （最新8k 通用版e2e模型，推荐使用）  chinese_16k_media (音视频领域，区域仅支持cn-north-4，cn-east-3)  chinese_8k_common  chinese_16k_conversation  chinese_8k_bank（银行领域，区域仅支持cn-north-4）  chinese_8k_insurance（保险领域，区域仅支持cn-north-4）  sichuan_8k_common （四川话识别，区域支持cn-north-4，cn-east-3）
	Property TranscriberConfigProperty `json:"property"`

	// 是否加标点， 可以为 yes, no(缺省)。
	AddPunc *TranscriberConfigAddPunc `json:"add_punc,omitempty"`

	NeedAnalysisInfo *AnalysisInfo `json:"need_analysis_info,omitempty"`

	// 热词表id，不使用则不填写。
	VocabularyId *string `json:"vocabulary_id,omitempty"`

	// 表示是否将语音中的数字识别为阿拉伯数字，取值为yes 、 no，缺省为yes。
	DigitNorm *TranscriberConfigDigitNorm `json:"digit_norm,omitempty"`

	// 用于录音文件识表示回调 url，即用户用于接收识别结果的服务器地址，不支持IP地址方式调用，url长度小于2048字节。服务请求方法为POST。  如果用户使用回调方式获取识别结果，需填写该参数，处理成功后用户服务器需返回状态码“200”。  如果用户使用轮询方式获取识别结果，则无需填写该参数。别结果的回调url，不使用则不填写。
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 表示是否在识别结果中输出分词结果信息，取值为“yes”和“no”，默认为“no”。
	NeedWordInfo *TranscriberConfigNeedWordInfo `json:"need_word_info,omitempty"`
}

func (o TranscriberConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscriberConfig struct{}"
	}

	return strings.Join([]string{"TranscriberConfig", string(data)}, " ")
}

type TranscriberConfigAudioFormat struct {
	value string
}

type TranscriberConfigAudioFormatEnum struct {
	AUTO        TranscriberConfigAudioFormat
	PCM16K16BIT TranscriberConfigAudioFormat
	PCM8K16BIT  TranscriberConfigAudioFormat
	ULAW16K8BIT TranscriberConfigAudioFormat
	ULAW8K8BIT  TranscriberConfigAudioFormat
	ALAW16K8BIT TranscriberConfigAudioFormat
	ALAW8K8BIT  TranscriberConfigAudioFormat
	VOX8K4BIT   TranscriberConfigAudioFormat
	V3_8K4BIT   TranscriberConfigAudioFormat
}

func GetTranscriberConfigAudioFormatEnum() TranscriberConfigAudioFormatEnum {
	return TranscriberConfigAudioFormatEnum{
		AUTO: TranscriberConfigAudioFormat{
			value: "auto",
		},
		PCM16K16BIT: TranscriberConfigAudioFormat{
			value: "pcm16k16bit",
		},
		PCM8K16BIT: TranscriberConfigAudioFormat{
			value: "pcm8k16bit",
		},
		ULAW16K8BIT: TranscriberConfigAudioFormat{
			value: "ulaw16k8bit",
		},
		ULAW8K8BIT: TranscriberConfigAudioFormat{
			value: "ulaw8k8bit",
		},
		ALAW16K8BIT: TranscriberConfigAudioFormat{
			value: "alaw16k8bit",
		},
		ALAW8K8BIT: TranscriberConfigAudioFormat{
			value: "alaw8k8bit",
		},
		VOX8K4BIT: TranscriberConfigAudioFormat{
			value: "vox8k4bit",
		},
		V3_8K4BIT: TranscriberConfigAudioFormat{
			value: "v3_8k4bit",
		},
	}
}

func (c TranscriberConfigAudioFormat) Value() string {
	return c.value
}

func (c TranscriberConfigAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TranscriberConfigAudioFormat) UnmarshalJSON(b []byte) error {
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

type TranscriberConfigProperty struct {
	value string
}

type TranscriberConfigPropertyEnum struct {
	CHINESE_16K_MEDIA        TranscriberConfigProperty
	CHINESE_8K_COMMON        TranscriberConfigProperty
	CHINESE_16K_CONVERSATION TranscriberConfigProperty
	CHINESE_8K_BANK          TranscriberConfigProperty
	CHINESE_8K_INSURANCE     TranscriberConfigProperty
	SICHUAN_8K_COMMON        TranscriberConfigProperty
	CHINESE_8K_GENERAL       TranscriberConfigProperty
}

func GetTranscriberConfigPropertyEnum() TranscriberConfigPropertyEnum {
	return TranscriberConfigPropertyEnum{
		CHINESE_16K_MEDIA: TranscriberConfigProperty{
			value: "chinese_16k_media",
		},
		CHINESE_8K_COMMON: TranscriberConfigProperty{
			value: "chinese_8k_common",
		},
		CHINESE_16K_CONVERSATION: TranscriberConfigProperty{
			value: "chinese_16k_conversation",
		},
		CHINESE_8K_BANK: TranscriberConfigProperty{
			value: "chinese_8k_bank",
		},
		CHINESE_8K_INSURANCE: TranscriberConfigProperty{
			value: "chinese_8k_insurance",
		},
		SICHUAN_8K_COMMON: TranscriberConfigProperty{
			value: "sichuan_8k_common",
		},
		CHINESE_8K_GENERAL: TranscriberConfigProperty{
			value: "chinese_8k_general",
		},
	}
}

func (c TranscriberConfigProperty) Value() string {
	return c.value
}

func (c TranscriberConfigProperty) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TranscriberConfigProperty) UnmarshalJSON(b []byte) error {
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

type TranscriberConfigAddPunc struct {
	value string
}

type TranscriberConfigAddPuncEnum struct {
	YES TranscriberConfigAddPunc
	NO  TranscriberConfigAddPunc
}

func GetTranscriberConfigAddPuncEnum() TranscriberConfigAddPuncEnum {
	return TranscriberConfigAddPuncEnum{
		YES: TranscriberConfigAddPunc{
			value: "yes",
		},
		NO: TranscriberConfigAddPunc{
			value: "no",
		},
	}
}

func (c TranscriberConfigAddPunc) Value() string {
	return c.value
}

func (c TranscriberConfigAddPunc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TranscriberConfigAddPunc) UnmarshalJSON(b []byte) error {
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

type TranscriberConfigDigitNorm struct {
	value string
}

type TranscriberConfigDigitNormEnum struct {
	YES TranscriberConfigDigitNorm
	NO  TranscriberConfigDigitNorm
}

func GetTranscriberConfigDigitNormEnum() TranscriberConfigDigitNormEnum {
	return TranscriberConfigDigitNormEnum{
		YES: TranscriberConfigDigitNorm{
			value: "yes",
		},
		NO: TranscriberConfigDigitNorm{
			value: "no",
		},
	}
}

func (c TranscriberConfigDigitNorm) Value() string {
	return c.value
}

func (c TranscriberConfigDigitNorm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TranscriberConfigDigitNorm) UnmarshalJSON(b []byte) error {
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

type TranscriberConfigNeedWordInfo struct {
	value string
}

type TranscriberConfigNeedWordInfoEnum struct {
	YES TranscriberConfigNeedWordInfo
	NO  TranscriberConfigNeedWordInfo
}

func GetTranscriberConfigNeedWordInfoEnum() TranscriberConfigNeedWordInfoEnum {
	return TranscriberConfigNeedWordInfoEnum{
		YES: TranscriberConfigNeedWordInfo{
			value: "yes",
		},
		NO: TranscriberConfigNeedWordInfo{
			value: "no",
		},
	}
}

func (c TranscriberConfigNeedWordInfo) Value() string {
	return c.value
}

func (c TranscriberConfigNeedWordInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TranscriberConfigNeedWordInfo) UnmarshalJSON(b []byte) error {
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
