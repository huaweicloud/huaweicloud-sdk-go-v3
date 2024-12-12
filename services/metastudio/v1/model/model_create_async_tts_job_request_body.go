package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAsyncTtsJobRequestBody 文本转语音任务请求。
type CreateAsyncTtsJobRequestBody struct {

	// 待合成文本
	Text string `json:"text"`

	// 音色ID，获取方式详见[获取音色ID](metastudio_02_0054.xml)。
	VoiceAssetId string `json:"voice_asset_id"`

	// 语速。 * 当取值为“100”时，表示一个成年人正常的语速，约为250字/分钟。 * 50表示0.5倍语速，100表示正常语速，200表示2倍语速。
	Speed *int32 `json:"speed,omitempty"`

	// 音高。
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量。
	Volume *int32 `json:"volume,omitempty"`

	// 输出音频文件格式。默认WAV。 * WAV：wav格式。 * MP3：mp3格式。
	AudioFormat *CreateAsyncTtsJobRequestBodyAudioFormat `json:"audio_format,omitempty"`

	// 是否需要时间戳。false为不需要，true为需要返回时间戳信息。默认值为false。
	NeedTimestamp *bool `json:"need_timestamp,omitempty"`

	// 异常时是否返回静默音频流
	SilenceFlag *bool `json:"silence_flag,omitempty"`

	// 异常时返回的静默音频流时长，单位毫秒。
	SilenceTimeMs *int32 `json:"silence_time_ms,omitempty"`
}

func (o CreateAsyncTtsJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAsyncTtsJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAsyncTtsJobRequestBody", string(data)}, " ")
}

type CreateAsyncTtsJobRequestBodyAudioFormat struct {
	value string
}

type CreateAsyncTtsJobRequestBodyAudioFormatEnum struct {
	WAV CreateAsyncTtsJobRequestBodyAudioFormat
	MP3 CreateAsyncTtsJobRequestBodyAudioFormat
}

func GetCreateAsyncTtsJobRequestBodyAudioFormatEnum() CreateAsyncTtsJobRequestBodyAudioFormatEnum {
	return CreateAsyncTtsJobRequestBodyAudioFormatEnum{
		WAV: CreateAsyncTtsJobRequestBodyAudioFormat{
			value: "WAV",
		},
		MP3: CreateAsyncTtsJobRequestBodyAudioFormat{
			value: "MP3",
		},
	}
}

func (c CreateAsyncTtsJobRequestBodyAudioFormat) Value() string {
	return c.value
}

func (c CreateAsyncTtsJobRequestBodyAudioFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAsyncTtsJobRequestBodyAudioFormat) UnmarshalJSON(b []byte) error {
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
