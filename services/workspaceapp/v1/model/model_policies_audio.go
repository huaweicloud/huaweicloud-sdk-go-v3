package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesAudio 音频。
type PoliciesAudio struct {

	// 是否开启音频重定向。取值为： false：表示关闭。 true：表示开启。
	AudioRedirectionEnable *bool `json:"audio_redirection_enable,omitempty"`

	// 是否开启播音重定向。取值为： false：表示关闭。 true：表示开启。
	PlayRedirectionEnable *bool `json:"play_redirection_enable,omitempty"`

	// 播音场景。取值为： 无损播音：LossLess。 语音通话：Speech Call。 音乐播音：Music Play。 自动识别：Automatic Identification。
	PlayClassification *PoliciesAudioPlayClassification `json:"play_classification,omitempty"`

	// 是否开启录音重定向。取值为： false：表示关闭。 true：表示开启。
	RecordRedirectionEnable *bool `json:"record_redirection_enable,omitempty"`

	// 录音场景。取值为： 无损录音：LossLess。 语音通话：Speech Call。 音乐录音：Music Record。 自动识别：Automatic Identification。
	RecordClassification *PoliciesAudioRecordClassification `json:"record_classification,omitempty"`
}

func (o PoliciesAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesAudio struct{}"
	}

	return strings.Join([]string{"PoliciesAudio", string(data)}, " ")
}

type PoliciesAudioPlayClassification struct {
	value string
}

type PoliciesAudioPlayClassificationEnum struct {
	LOSS_LESS                PoliciesAudioPlayClassification
	SPEECH_CALL              PoliciesAudioPlayClassification
	MUSIC_PLAY               PoliciesAudioPlayClassification
	AUTOMATIC_IDENTIFICATION PoliciesAudioPlayClassification
}

func GetPoliciesAudioPlayClassificationEnum() PoliciesAudioPlayClassificationEnum {
	return PoliciesAudioPlayClassificationEnum{
		LOSS_LESS: PoliciesAudioPlayClassification{
			value: "LossLess",
		},
		SPEECH_CALL: PoliciesAudioPlayClassification{
			value: "Speech Call",
		},
		MUSIC_PLAY: PoliciesAudioPlayClassification{
			value: "Music Play",
		},
		AUTOMATIC_IDENTIFICATION: PoliciesAudioPlayClassification{
			value: "Automatic Identification",
		},
	}
}

func (c PoliciesAudioPlayClassification) Value() string {
	return c.value
}

func (c PoliciesAudioPlayClassification) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioPlayClassification) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioRecordClassification struct {
	value string
}

type PoliciesAudioRecordClassificationEnum struct {
	LOSS_LESS                PoliciesAudioRecordClassification
	SPEECH_CALL              PoliciesAudioRecordClassification
	MUSIC_RECORD             PoliciesAudioRecordClassification
	AUTOMATIC_IDENTIFICATION PoliciesAudioRecordClassification
}

func GetPoliciesAudioRecordClassificationEnum() PoliciesAudioRecordClassificationEnum {
	return PoliciesAudioRecordClassificationEnum{
		LOSS_LESS: PoliciesAudioRecordClassification{
			value: "LossLess",
		},
		SPEECH_CALL: PoliciesAudioRecordClassification{
			value: "Speech Call",
		},
		MUSIC_RECORD: PoliciesAudioRecordClassification{
			value: "Music Record",
		},
		AUTOMATIC_IDENTIFICATION: PoliciesAudioRecordClassification{
			value: "Automatic Identification",
		},
	}
}

func (c PoliciesAudioRecordClassification) Value() string {
	return c.value
}

func (c PoliciesAudioRecordClassification) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioRecordClassification) UnmarshalJSON(b []byte) error {
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
