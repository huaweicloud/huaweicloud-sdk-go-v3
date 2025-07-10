package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesAudio 音频。
type PoliciesAudio struct {

	// 是否开启音频重定向。取值为： false：表示关闭。 true：表示开启。
	AudioRedirectionEnable *bool `json:"audio_redirection_enable,omitempty"`

	// 播音设置音量。 不设置音量取值为：Do Not Set Volume。 音量设置，最小值为10，最大值为100，中间取值，间隔为5的递增序列。如：10、15、20、25等。
	PlayVolume *string `json:"play_volume,omitempty"`

	// 播音设置音量线性系数。取值为： 不设置：Do Not Set Volume Ratio。 低：Low。 中：Middle。 高：High。
	PlayVolumeRatio *string `json:"play_volume_ratio,omitempty"`

	// 录音设置音量。 不设置取值为：Do Not Set Volume。 音量设置，最小值为10，最大值为100，中间取值，间隔为5的递增序列。如：10、15、20、25等。
	RecordVolume *string `json:"record_volume,omitempty"`

	// 录音设置音量线性系数。取值为： 不设置：Do Not Set Volume Ratio。 低：Low。 中：Middle。 高：High。
	RecordVolumeRatio *string `json:"record_volume_ratio,omitempty"`

	// 音频传输方式。取值为： 可靠传输：Reliable Transmission。 实时传输：Real Time Transmission。
	AudioTransmissionMode *string `json:"audio_transmission_mode,omitempty"`

	// 是否开启播音重定向。取值为： false：表示关闭。 true：表示开启。
	PlayRedirectionEnable *bool `json:"play_redirection_enable,omitempty"`

	// 播音场景。取值为： 无损播音：LossLess。 语音通话：Speech Call。 音乐播音：Music Play。 自动识别：Automatic Identification。
	PlayClassification *string `json:"play_classification,omitempty"`

	// 播音质量。取值为： 低：Low。 中：Middle。 高：High
	PlayQuality *string `json:"play_quality,omitempty"`

	// 播音降噪。取值为： 不开启降噪：Disable Denoising。 开启降噪，最小值为-100，最大值为-5，中间取值，间隔为5的递增序列。如：-100、-95、-90、-85等。
	PlayDenoising *string `json:"play_denoising,omitempty"`

	// 播音增益。取值为：不开启增益：Disable AGC。开启增益，最小值为4000，最大值为32000，中间取值，从10000开始间隔为1000的递增序列。如：4000、6000、8000、10000、11000、12000、13000等。
	PlayAgc *string `json:"play_agc,omitempty"`

	// 播音设置冗余。取值为： 不开启冗余：Disable CRC。 开启冗余：Enable CRC。
	PlayCrc *string `json:"play_crc,omitempty"`

	// 播音设置模式。取值为： 播音设备共享模式：Play Device In Shared Mode。 播音设备独占模式：Play Device In Exclusive Mode。
	PlayDeviceMode *string `json:"play_device_mode,omitempty"`

	// 播音控制时延阈值。取值为：最小值为160，最大值为2500。中间取值，400以下为40的递增序列，1000以下为50的递增序列。从高到低为：2500、2000、1800、1500、1200、1000、950、900、850等。
	PlayDelayThreshold *string `json:"play_delay_threshold,omitempty"`

	// 播音检测振幅阈值。取值为：2048、4096、5120、6144、7168、8192。
	PlayAmplitudeThreshold *string `json:"play_amplitude_threshold,omitempty"`

	// 播音音乐预充数据。取值为：不预充：Do Not Prefill Data。预充取值：最小值为50，最大值为2000，250以前为50的递增序列，500以前为100的递增序列。从高到低取值如：2000、1500、1000、500、400、300、250。
	PlayPrefillData *string `json:"play_prefill_data,omitempty"`

	// 是否开启录音重定向。取值为： false：表示关闭。 true：表示开启。
	RecordRedirectionEnable *bool `json:"record_redirection_enable,omitempty"`

	// 录音场景。取值为： 无损录音：LossLess。 语音通话：Speech Call。 音乐录音：Music Record。 自动识别：Automatic Identification。
	RecordClassification *string `json:"record_classification,omitempty"`

	// 录音质量。取值为： 低：Low。 中：Middle。 高：High。
	RecordQuality *string `json:"record_quality,omitempty"`

	// 录音降噪。取值为： 不开启降噪：Disable Denoising。 开启降噪，最小值为-100，最大值为-5，中间取值，间隔为5的递增序列。如：-100、-95、-90、-85等。
	RecordDenoising *string `json:"record_denoising,omitempty"`

	// 录音增益。取值为：不开启增益：Disable AGC。开启增益，最小值为4000，最大值为32000，中间取值，从10000开始间隔为1000的递增序列。如：4000、6000、8000、10000、11000、12000、13000等。
	RecordAgc *string `json:"record_agc,omitempty"`

	// 录音设置冗余。取值为： 不开启冗余：Disable CRC。 开启冗余：Enable CRC。
	RecordCrc *string `json:"record_crc,omitempty"`

	// 录音设置模式。取值为： 播音设备共享模式：Record Device In Shared Mode。 播音设备独占模式：Record Device In Exclusive Mode。
	RecordDeviceMode *string `json:"record_device_mode,omitempty"`

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
