package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VideoAssetMeta 视频元数据，自动提取获得。
type VideoAssetMeta struct {

	// **参数解释**： 视频编码格式。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及
	VideoCodec *string `json:"video_codec,omitempty"`

	// **参数解释**： 视频画面宽度。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Width *int32 `json:"width,omitempty"`

	// **参数解释**： 视频画面高度。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Height *int32 `json:"height,omitempty"`

	// **参数解释**： 视频帧率。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及
	FrameRate *string `json:"frame_rate,omitempty"`

	// **参数解释**： 视频平均码率,单位kbps。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	VideoBitRate *int32 `json:"video_bit_rate,omitempty"`

	// **参数解释**： 时长,单位秒。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**： 音频编码格式。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及
	AudioCodec *string `json:"audio_codec,omitempty"`

	// **参数解释**： 音频平均码率,单位kbps。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	AudioBitRate *int32 `json:"audio_bit_rate,omitempty"`

	// **参数解释**： 音频声道数。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	AudioChannels *int32 `json:"audio_channels,omitempty"`

	// **参数解释**： 采样率,HZ。 **约束限制**： 用户无需填写，系统自行提取。 **默认取值**： 不涉及
	Sample *int32 `json:"sample,omitempty"`

	// **参数解释**： 横向画面或纵向画面。 **约束限制**： 用户无需填写，系统自行提取。 **取值范围**： * Horizontal：横向 * Vertical：纵向  **默认取值**： 不涉及
	Mode *VideoAssetMetaMode `json:"mode,omitempty"`

	// **参数解释**： 视频转码状态。 **约束限制**： 用户无需填写，系统自行填写。 **取值范围**： * WAITING：等待 * TRANSCODING：转码中 * FAILED：失败 * SUCCEEDED：成功  **默认取值**： 不涉及
	VideoTranscodingStatus *VideoAssetMetaVideoTranscodingStatus `json:"video_transcoding_status,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o VideoAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoAssetMeta struct{}"
	}

	return strings.Join([]string{"VideoAssetMeta", string(data)}, " ")
}

type VideoAssetMetaMode struct {
	value string
}

type VideoAssetMetaModeEnum struct {
	HORIZONTAL VideoAssetMetaMode
	VERTICAL   VideoAssetMetaMode
}

func GetVideoAssetMetaModeEnum() VideoAssetMetaModeEnum {
	return VideoAssetMetaModeEnum{
		HORIZONTAL: VideoAssetMetaMode{
			value: "Horizontal",
		},
		VERTICAL: VideoAssetMetaMode{
			value: "Vertical",
		},
	}
}

func (c VideoAssetMetaMode) Value() string {
	return c.value
}

func (c VideoAssetMetaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoAssetMetaMode) UnmarshalJSON(b []byte) error {
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

type VideoAssetMetaVideoTranscodingStatus struct {
	value string
}

type VideoAssetMetaVideoTranscodingStatusEnum struct {
	WAITING     VideoAssetMetaVideoTranscodingStatus
	TRANSCODING VideoAssetMetaVideoTranscodingStatus
	FAILED      VideoAssetMetaVideoTranscodingStatus
	SUCCEEDED   VideoAssetMetaVideoTranscodingStatus
}

func GetVideoAssetMetaVideoTranscodingStatusEnum() VideoAssetMetaVideoTranscodingStatusEnum {
	return VideoAssetMetaVideoTranscodingStatusEnum{
		WAITING: VideoAssetMetaVideoTranscodingStatus{
			value: "WAITING",
		},
		TRANSCODING: VideoAssetMetaVideoTranscodingStatus{
			value: "TRANSCODING",
		},
		FAILED: VideoAssetMetaVideoTranscodingStatus{
			value: "FAILED",
		},
		SUCCEEDED: VideoAssetMetaVideoTranscodingStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c VideoAssetMetaVideoTranscodingStatus) Value() string {
	return c.value
}

func (c VideoAssetMetaVideoTranscodingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoAssetMetaVideoTranscodingStatus) UnmarshalJSON(b []byte) error {
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
