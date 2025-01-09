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

	// 播音设置音量。 不设置音量取值为：Do Not Set Volume。 音量设置，最小值为10，最大值为100，中间取值，间隔为5的递增序列。如：10、15、20、25等。
	PlayVolume *string `json:"play_volume,omitempty"`

	// 播音设置音量线性系数。取值为： 不设置：Do Not Set Volume Ratio。 低：Low。 中：Middle。 高：High。
	PlayVolumeRatio *PoliciesAudioPlayVolumeRatio `json:"play_volume_ratio,omitempty"`

	// 录音设置音量。 不设置取值为：Do Not Set Volume。 音量设置，最小值为10，最大值为100，中间取值，间隔为5的递增序列。如：10、15、20、25等。
	RecordVolume *string `json:"record_volume,omitempty"`

	// 录音设置音量线性系数。取值为： 不设置：Do Not Set Volume Ratio。 低：Low。 中：Middle。 高：High。
	RecordVolumeRatio *PoliciesAudioRecordVolumeRatio `json:"record_volume_ratio,omitempty"`

	// 音频传输方式。取值为： 可靠传输：Reliable Transmission。 实时传输：Real Time Transmission。
	AudioTransmissionMode *PoliciesAudioAudioTransmissionMode `json:"audio_transmission_mode,omitempty"`

	// 是否开启播音重定向。取值为： false：表示关闭。 true：表示开启。
	PlayRedirectionEnable *bool `json:"play_redirection_enable,omitempty"`

	// 播音场景。取值为： 无损播音：LossLess。 语音通话：Speech Call。 音乐播音：Music Play。 自动识别：Automatic Identification。
	PlayClassification *PoliciesAudioPlayClassification `json:"play_classification,omitempty"`

	// 播音质量。取值为： 低：Low。 中：Middle。 高：High
	PlayQuality *PoliciesAudioPlayQuality `json:"play_quality,omitempty"`

	// 播音降噪。取值为： 不开启降噪：Disable Denoising。 开启降噪，最小值为-100，最大值为-5，中间取值，间隔为5的递增序列。如：-100、-95、-90、-85等。
	PlayDenoising *string `json:"play_denoising,omitempty"`

	// 播音增益。取值为：不开启增益：Disable AGC。开启增益，最小值为4000，最大值为32000，中间取值，从10000开始间隔为1000的递增序列。如：4000、6000、8000、10000、11000、12000、13000等。
	PlayAgc *string `json:"play_agc,omitempty"`

	// 播音设置冗余。取值为： 不开启冗余：Disable CRC。 开启冗余：Enable CRC。
	PlayCrc *PoliciesAudioPlayCrc `json:"play_crc,omitempty"`

	// 播音设置模式。取值为： 播音设备共享模式：Play Device In Shared Mode。 播音设备独占模式：Play Device In Exclusive Mode。
	PlayDeviceMode *PoliciesAudioPlayDeviceMode `json:"play_device_mode,omitempty"`

	// 播音控制时延阈值。取值为：最小值为160，最大值为2500。中间取值，400以下为40的递增序列，1000以下为50的递增序列。从高到低为：2500、2000、1800、1500、1200、1000、950、900、850等。
	PlayDelayThreshold *string `json:"play_delay_threshold,omitempty"`

	// 播音检测振幅阈值。取值为：2048、4096、5120、6144、7168、8192。
	PlayAmplitudeThreshold *string `json:"play_amplitude_threshold,omitempty"`

	// 播音音乐预充数据。取值为：不预充：Do Not Prefill Data。预充取值：最小值为50，最大值为2000，250以前为50的递增序列，500以前为100的递增序列。从高到低取值如：2000、1500、1000、500、400、300、250。
	PlayPrefillData *string `json:"play_prefill_data,omitempty"`

	// 是否开启录音重定向。取值为： false：表示关闭。 true：表示开启。
	RecordRedirectionEnable *bool `json:"record_redirection_enable,omitempty"`

	// 录音场景。取值为： 无损录音：LossLess。 语音通话：Speech Call。 音乐录音：Music Record。 自动识别：Automatic Identification。
	RecordClassification *PoliciesAudioRecordClassification `json:"record_classification,omitempty"`

	// 录音质量。取值为： 低：Low。 中：Middle。 高：High。
	RecordQuality *PoliciesAudioRecordQuality `json:"record_quality,omitempty"`

	// 录音降噪。取值为： 不开启降噪：Disable Denoising。 开启降噪，最小值为-100，最大值为-5，中间取值，间隔为5的递增序列。如：-100、-95、-90、-85等。
	RecordDenoising *string `json:"record_denoising,omitempty"`

	// 录音增益。取值为：不开启增益：Disable AGC。开启增益，最小值为4000，最大值为32000，中间取值，从10000开始间隔为1000的递增序列。如：4000、6000、8000、10000、11000、12000、13000等。
	RecordAgc *string `json:"record_agc,omitempty"`

	// 录音设置冗余。取值为： 不开启冗余：Disable CRC。 开启冗余：Enable CRC。
	RecordCrc *PoliciesAudioRecordCrc `json:"record_crc,omitempty"`

	// 录音设置模式。取值为： 播音设备共享模式：Record Device In Shared Mode。 播音设备独占模式：Record Device In Exclusive Mode。
	RecordDeviceMode *PoliciesAudioRecordDeviceMode `json:"record_device_mode,omitempty"`

	// 录音控制时延阈值。取值为：最小值为160，最大值为2500。中间取值，400以下为40的递增序列，1000以下为50的递增序列。从高到低为：2500、2000、1800、1500、1200、1000、950、900、850等。
	RecordDelayThreshold *string `json:"record_delay_threshold,omitempty"`

	// 录音检测振幅阈值。取值为：2048、4096、5120、6144、7168、8192。
	RecordAmplitudeThreshold *string `json:"record_amplitude_threshold,omitempty"`
}

func (o PoliciesAudio) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesAudio struct{}"
	}

	return strings.Join([]string{"PoliciesAudio", string(data)}, " ")
}

type PoliciesAudioPlayVolumeRatio struct {
	value string
}

type PoliciesAudioPlayVolumeRatioEnum struct {
	DO_NOT_SET_VOLUME_RATIO PoliciesAudioPlayVolumeRatio
	LOW                     PoliciesAudioPlayVolumeRatio
	MIDDLE                  PoliciesAudioPlayVolumeRatio
	HIGH                    PoliciesAudioPlayVolumeRatio
}

func GetPoliciesAudioPlayVolumeRatioEnum() PoliciesAudioPlayVolumeRatioEnum {
	return PoliciesAudioPlayVolumeRatioEnum{
		DO_NOT_SET_VOLUME_RATIO: PoliciesAudioPlayVolumeRatio{
			value: "Do Not Set Volume Ratio",
		},
		LOW: PoliciesAudioPlayVolumeRatio{
			value: "Low",
		},
		MIDDLE: PoliciesAudioPlayVolumeRatio{
			value: "Middle",
		},
		HIGH: PoliciesAudioPlayVolumeRatio{
			value: "High",
		},
	}
}

func (c PoliciesAudioPlayVolumeRatio) Value() string {
	return c.value
}

func (c PoliciesAudioPlayVolumeRatio) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioPlayVolumeRatio) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioRecordVolumeRatio struct {
	value string
}

type PoliciesAudioRecordVolumeRatioEnum struct {
	DO_NOT_SET_VOLUME_RATIO PoliciesAudioRecordVolumeRatio
	LOW                     PoliciesAudioRecordVolumeRatio
	MIDDLE                  PoliciesAudioRecordVolumeRatio
	HIGH                    PoliciesAudioRecordVolumeRatio
}

func GetPoliciesAudioRecordVolumeRatioEnum() PoliciesAudioRecordVolumeRatioEnum {
	return PoliciesAudioRecordVolumeRatioEnum{
		DO_NOT_SET_VOLUME_RATIO: PoliciesAudioRecordVolumeRatio{
			value: "Do Not Set Volume Ratio",
		},
		LOW: PoliciesAudioRecordVolumeRatio{
			value: "Low",
		},
		MIDDLE: PoliciesAudioRecordVolumeRatio{
			value: "Middle",
		},
		HIGH: PoliciesAudioRecordVolumeRatio{
			value: "High",
		},
	}
}

func (c PoliciesAudioRecordVolumeRatio) Value() string {
	return c.value
}

func (c PoliciesAudioRecordVolumeRatio) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioRecordVolumeRatio) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioAudioTransmissionMode struct {
	value string
}

type PoliciesAudioAudioTransmissionModeEnum struct {
	REAL_TIME_TRANSMISSION PoliciesAudioAudioTransmissionMode
	RELIABLE_TRANSMISSION  PoliciesAudioAudioTransmissionMode
}

func GetPoliciesAudioAudioTransmissionModeEnum() PoliciesAudioAudioTransmissionModeEnum {
	return PoliciesAudioAudioTransmissionModeEnum{
		REAL_TIME_TRANSMISSION: PoliciesAudioAudioTransmissionMode{
			value: "Real Time Transmission",
		},
		RELIABLE_TRANSMISSION: PoliciesAudioAudioTransmissionMode{
			value: "Reliable Transmission",
		},
	}
}

func (c PoliciesAudioAudioTransmissionMode) Value() string {
	return c.value
}

func (c PoliciesAudioAudioTransmissionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioAudioTransmissionMode) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioPlayQuality struct {
	value string
}

type PoliciesAudioPlayQualityEnum struct {
	LOW    PoliciesAudioPlayQuality
	MIDDLE PoliciesAudioPlayQuality
	HIGH   PoliciesAudioPlayQuality
}

func GetPoliciesAudioPlayQualityEnum() PoliciesAudioPlayQualityEnum {
	return PoliciesAudioPlayQualityEnum{
		LOW: PoliciesAudioPlayQuality{
			value: "Low",
		},
		MIDDLE: PoliciesAudioPlayQuality{
			value: "Middle",
		},
		HIGH: PoliciesAudioPlayQuality{
			value: "High",
		},
	}
}

func (c PoliciesAudioPlayQuality) Value() string {
	return c.value
}

func (c PoliciesAudioPlayQuality) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioPlayQuality) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioPlayCrc struct {
	value string
}

type PoliciesAudioPlayCrcEnum struct {
	DISABLE_CRC PoliciesAudioPlayCrc
	ENABLE_CRC  PoliciesAudioPlayCrc
}

func GetPoliciesAudioPlayCrcEnum() PoliciesAudioPlayCrcEnum {
	return PoliciesAudioPlayCrcEnum{
		DISABLE_CRC: PoliciesAudioPlayCrc{
			value: "Disable CRC",
		},
		ENABLE_CRC: PoliciesAudioPlayCrc{
			value: "Enable CRC",
		},
	}
}

func (c PoliciesAudioPlayCrc) Value() string {
	return c.value
}

func (c PoliciesAudioPlayCrc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioPlayCrc) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioPlayDeviceMode struct {
	value string
}

type PoliciesAudioPlayDeviceModeEnum struct {
	PLAY_DEVICE_IN_SHARED_MODE    PoliciesAudioPlayDeviceMode
	PLAY_DEVICE_IN_EXCLUSIVE_MODE PoliciesAudioPlayDeviceMode
}

func GetPoliciesAudioPlayDeviceModeEnum() PoliciesAudioPlayDeviceModeEnum {
	return PoliciesAudioPlayDeviceModeEnum{
		PLAY_DEVICE_IN_SHARED_MODE: PoliciesAudioPlayDeviceMode{
			value: "Play Device In Shared Mode",
		},
		PLAY_DEVICE_IN_EXCLUSIVE_MODE: PoliciesAudioPlayDeviceMode{
			value: "Play Device In Exclusive Mode",
		},
	}
}

func (c PoliciesAudioPlayDeviceMode) Value() string {
	return c.value
}

func (c PoliciesAudioPlayDeviceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioPlayDeviceMode) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioRecordQuality struct {
	value string
}

type PoliciesAudioRecordQualityEnum struct {
	LOW    PoliciesAudioRecordQuality
	MIDDLE PoliciesAudioRecordQuality
	HIGH   PoliciesAudioRecordQuality
}

func GetPoliciesAudioRecordQualityEnum() PoliciesAudioRecordQualityEnum {
	return PoliciesAudioRecordQualityEnum{
		LOW: PoliciesAudioRecordQuality{
			value: "Low",
		},
		MIDDLE: PoliciesAudioRecordQuality{
			value: "Middle",
		},
		HIGH: PoliciesAudioRecordQuality{
			value: "High",
		},
	}
}

func (c PoliciesAudioRecordQuality) Value() string {
	return c.value
}

func (c PoliciesAudioRecordQuality) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioRecordQuality) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioRecordCrc struct {
	value string
}

type PoliciesAudioRecordCrcEnum struct {
	DISABLE_CRC PoliciesAudioRecordCrc
	ENABLE_CRC  PoliciesAudioRecordCrc
}

func GetPoliciesAudioRecordCrcEnum() PoliciesAudioRecordCrcEnum {
	return PoliciesAudioRecordCrcEnum{
		DISABLE_CRC: PoliciesAudioRecordCrc{
			value: "Disable CRC",
		},
		ENABLE_CRC: PoliciesAudioRecordCrc{
			value: "Enable CRC",
		},
	}
}

func (c PoliciesAudioRecordCrc) Value() string {
	return c.value
}

func (c PoliciesAudioRecordCrc) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioRecordCrc) UnmarshalJSON(b []byte) error {
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

type PoliciesAudioRecordDeviceMode struct {
	value string
}

type PoliciesAudioRecordDeviceModeEnum struct {
	RECORD_DEVICE_IN_SHARED_MODE    PoliciesAudioRecordDeviceMode
	RECORD_DEVICE_IN_EXCLUSIVE_MODE PoliciesAudioRecordDeviceMode
}

func GetPoliciesAudioRecordDeviceModeEnum() PoliciesAudioRecordDeviceModeEnum {
	return PoliciesAudioRecordDeviceModeEnum{
		RECORD_DEVICE_IN_SHARED_MODE: PoliciesAudioRecordDeviceMode{
			value: "Record Device In Shared Mode",
		},
		RECORD_DEVICE_IN_EXCLUSIVE_MODE: PoliciesAudioRecordDeviceMode{
			value: "Record Device In Exclusive Mode",
		},
	}
}

func (c PoliciesAudioRecordDeviceMode) Value() string {
	return c.value
}

func (c PoliciesAudioRecordDeviceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesAudioRecordDeviceMode) UnmarshalJSON(b []byte) error {
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
