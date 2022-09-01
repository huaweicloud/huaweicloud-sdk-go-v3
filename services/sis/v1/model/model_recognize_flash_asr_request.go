package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RecognizeFlashAsrRequest struct {

	// 所使用的模型特征串。通常是 “语种_采样率_领域”的形式。 采样率需要与音频采样率保持一致。 当前支持如下模型特征串： chinese_8k_common chinese_16k_conversation
	Property RecognizeFlashAsrRequestProperty `json:"property" xml:"property"`

	// 音频格式，audio_format取值范围： wav,mp3,m4a,aac,opus
	AudioFormat RecognizeFlashAsrRequestAudioFormat `json:"audio_format" xml:"audio_format"`

	// 是否加标点， 可以为 yes, 默认no
	AddPunc *RecognizeFlashAsrRequestAddPunc `json:"add_punc,omitempty" xml:"add_punc"`

	// 是否将音频中的数字使用阿拉伯数字的形式呈现，取值为yes，no，默认为yes
	DigitNorm *RecognizeFlashAsrRequestDigitNorm `json:"digit_norm,omitempty" xml:"digit_norm"`

	// 表示是否在识别结果中输出分词结果信息，取值为yes，no，默认no
	NeedWordInfo *RecognizeFlashAsrRequestNeedWordInfo `json:"need_word_info,omitempty" xml:"need_word_info"`

	// 热词表id
	VocabularyId *string `json:"vocabulary_id,omitempty" xml:"vocabulary_id"`

	// obs桶名
	ObsBucketName *string `json:"obs_bucket_name,omitempty" xml:"obs_bucket_name"`

	// obs对象key，经过urlencode编码，长度不超过1024个字符
	ObsObjectKey *string `json:"obs_object_key,omitempty" xml:"obs_object_key"`

	// 表示是否在识别中只识别首个声道的音频数据，取值为“yes”和“no”，默认为“no”。
	FirstChannelOnly *RecognizeFlashAsrRequestFirstChannelOnly `json:"first_channel_only,omitempty" xml:"first_channel_only"`
}

func (o RecognizeFlashAsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeFlashAsrRequest struct{}"
	}

	return strings.Join([]string{"RecognizeFlashAsrRequest", string(data)}, " ")
}

type RecognizeFlashAsrRequestProperty struct {
	value string
}

type RecognizeFlashAsrRequestPropertyEnum struct {
	CHINESE_8K_COMMON        RecognizeFlashAsrRequestProperty
	CHINESE_16K_CONVERSATION RecognizeFlashAsrRequestProperty
}

func GetRecognizeFlashAsrRequestPropertyEnum() RecognizeFlashAsrRequestPropertyEnum {
	return RecognizeFlashAsrRequestPropertyEnum{
		CHINESE_8K_COMMON: RecognizeFlashAsrRequestProperty{
			value: "chinese_8k_common",
		},
		CHINESE_16K_CONVERSATION: RecognizeFlashAsrRequestProperty{
			value: "chinese_16k_conversation",
		},
	}
}

func (c RecognizeFlashAsrRequestProperty) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestProperty) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestProperty) UnmarshalJSON(b []byte) error {
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

type RecognizeFlashAsrRequestAudioFormat struct {
	value string
}

type RecognizeFlashAsrRequestAudioFormatEnum struct {
	WAV  RecognizeFlashAsrRequestAudioFormat
	MP3  RecognizeFlashAsrRequestAudioFormat
	M4A  RecognizeFlashAsrRequestAudioFormat
	AAC  RecognizeFlashAsrRequestAudioFormat
	OPUS RecognizeFlashAsrRequestAudioFormat
}

func GetRecognizeFlashAsrRequestAudioFormatEnum() RecognizeFlashAsrRequestAudioFormatEnum {
	return RecognizeFlashAsrRequestAudioFormatEnum{
		WAV: RecognizeFlashAsrRequestAudioFormat{
			value: "wav",
		},
		MP3: RecognizeFlashAsrRequestAudioFormat{
			value: "mp3",
		},
		M4A: RecognizeFlashAsrRequestAudioFormat{
			value: "m4a",
		},
		AAC: RecognizeFlashAsrRequestAudioFormat{
			value: "aac",
		},
		OPUS: RecognizeFlashAsrRequestAudioFormat{
			value: "opus",
		},
	}
}

func (c RecognizeFlashAsrRequestAudioFormat) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestAudioFormat) UnmarshalJSON(b []byte) error {
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

type RecognizeFlashAsrRequestAddPunc struct {
	value string
}

type RecognizeFlashAsrRequestAddPuncEnum struct {
	YES RecognizeFlashAsrRequestAddPunc
	NO  RecognizeFlashAsrRequestAddPunc
}

func GetRecognizeFlashAsrRequestAddPuncEnum() RecognizeFlashAsrRequestAddPuncEnum {
	return RecognizeFlashAsrRequestAddPuncEnum{
		YES: RecognizeFlashAsrRequestAddPunc{
			value: "yes",
		},
		NO: RecognizeFlashAsrRequestAddPunc{
			value: "no",
		},
	}
}

func (c RecognizeFlashAsrRequestAddPunc) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestAddPunc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestAddPunc) UnmarshalJSON(b []byte) error {
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

type RecognizeFlashAsrRequestDigitNorm struct {
	value string
}

type RecognizeFlashAsrRequestDigitNormEnum struct {
	YES RecognizeFlashAsrRequestDigitNorm
	NO  RecognizeFlashAsrRequestDigitNorm
}

func GetRecognizeFlashAsrRequestDigitNormEnum() RecognizeFlashAsrRequestDigitNormEnum {
	return RecognizeFlashAsrRequestDigitNormEnum{
		YES: RecognizeFlashAsrRequestDigitNorm{
			value: "yes",
		},
		NO: RecognizeFlashAsrRequestDigitNorm{
			value: "no",
		},
	}
}

func (c RecognizeFlashAsrRequestDigitNorm) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestDigitNorm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestDigitNorm) UnmarshalJSON(b []byte) error {
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

type RecognizeFlashAsrRequestNeedWordInfo struct {
	value string
}

type RecognizeFlashAsrRequestNeedWordInfoEnum struct {
	YES RecognizeFlashAsrRequestNeedWordInfo
	NO  RecognizeFlashAsrRequestNeedWordInfo
}

func GetRecognizeFlashAsrRequestNeedWordInfoEnum() RecognizeFlashAsrRequestNeedWordInfoEnum {
	return RecognizeFlashAsrRequestNeedWordInfoEnum{
		YES: RecognizeFlashAsrRequestNeedWordInfo{
			value: "yes",
		},
		NO: RecognizeFlashAsrRequestNeedWordInfo{
			value: "no",
		},
	}
}

func (c RecognizeFlashAsrRequestNeedWordInfo) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestNeedWordInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestNeedWordInfo) UnmarshalJSON(b []byte) error {
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

type RecognizeFlashAsrRequestFirstChannelOnly struct {
	value string
}

type RecognizeFlashAsrRequestFirstChannelOnlyEnum struct {
	YES RecognizeFlashAsrRequestFirstChannelOnly
	NO  RecognizeFlashAsrRequestFirstChannelOnly
}

func GetRecognizeFlashAsrRequestFirstChannelOnlyEnum() RecognizeFlashAsrRequestFirstChannelOnlyEnum {
	return RecognizeFlashAsrRequestFirstChannelOnlyEnum{
		YES: RecognizeFlashAsrRequestFirstChannelOnly{
			value: "yes",
		},
		NO: RecognizeFlashAsrRequestFirstChannelOnly{
			value: "no",
		},
	}
}

func (c RecognizeFlashAsrRequestFirstChannelOnly) Value() string {
	return c.value
}

func (c RecognizeFlashAsrRequestFirstChannelOnly) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecognizeFlashAsrRequestFirstChannelOnly) UnmarshalJSON(b []byte) error {
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
