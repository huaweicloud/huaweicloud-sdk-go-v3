package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VodEditingSetting struct {

	// 分辨率自适应策略, 选值：true, false（默认true） 输入为true时，width表示长边，height表示短边 输入为false时，width表示宽度，height表示长度
	ResolutionAdaptation *bool `json:"resolution_adaptation,omitempty"`

	// 分辨率上采样开关, 选值：true, false（默认false），若为false则按照原片源输出，分辨率不会上浮。
	ResolutionUpsample *bool `json:"resolution_upsample,omitempty"`

	// 输出封装格式，HLS、MP4（默认MP4）、MP3、MOV、FLV、AVI。 不支持将视频文件输出成音频封装格式。
	Format *VodEditingSettingFormat `json:"format,omitempty"`

	// 输出宽或长边，整型，输入小数向下取整，默认0，按源  - 当width、height均为0，则分辨率取片源分辨率； - 当width为0，height非0，则width按片源分辨率比例缩放； - 当width非0，height为0，则height按片源分辨率比例缩放； - 当width、height均非0，则分辨率按用户指定。 - 当视频编码为H.264时，则width最小值为32，最大值为4096。 - 当视频编码为H.265，则width最小值为160，最大值为4096。
	Width *int32 `json:"width,omitempty"`

	// 输出高或短边，整型，输入小数向下取整，默认0，按源  - 当Width、Height均为0，则分辨率取片源分辨率； - 当Width为0，Height非0，则Width按片源分辨率比例缩放； - 当Width非0，Height为0，则Height按片源分辨率比例缩放； - 当Width、Height 均非0，则分辨率按用户指定。 - 当视频编码为H.264时，则height最小值为32，最大值为2880。 - 当视频编码为H.265，则height最小值为160，最大值为2880。
	Height *int32 `json:"height,omitempty"`

	// 输出参考基准，可选，默认为NONE  - NONE 输出分辨率按输入的第一个片源为主，码率按输出分辨率自适应 - MAX_BITRATE 按码率最大的输入片源参数为基准 - MAX_RESOLUTION 按分辨率最大的输入片源参数为基准
	Reference *VodEditingSettingReference `json:"reference,omitempty"`

	// 视频编码格式。 取值如下： - 1：VIDEO_CODEC_H264 - 2：VIDEO_CODEC_H265
	VideoCodec *int32 `json:"video_codec,omitempty"`
}

func (o VodEditingSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VodEditingSetting struct{}"
	}

	return strings.Join([]string{"VodEditingSetting", string(data)}, " ")
}

type VodEditingSettingFormat struct {
	value string
}

type VodEditingSettingFormatEnum struct {
	MP4 VodEditingSettingFormat
	HLS VodEditingSettingFormat
	MP3 VodEditingSettingFormat
	MOV VodEditingSettingFormat
	AVI VodEditingSettingFormat
	FLV VodEditingSettingFormat
}

func GetVodEditingSettingFormatEnum() VodEditingSettingFormatEnum {
	return VodEditingSettingFormatEnum{
		MP4: VodEditingSettingFormat{
			value: "MP4",
		},
		HLS: VodEditingSettingFormat{
			value: "HLS",
		},
		MP3: VodEditingSettingFormat{
			value: "MP3",
		},
		MOV: VodEditingSettingFormat{
			value: "MOV",
		},
		AVI: VodEditingSettingFormat{
			value: "AVI",
		},
		FLV: VodEditingSettingFormat{
			value: "FLV",
		},
	}
}

func (c VodEditingSettingFormat) Value() string {
	return c.value
}

func (c VodEditingSettingFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VodEditingSettingFormat) UnmarshalJSON(b []byte) error {
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

type VodEditingSettingReference struct {
	value string
}

type VodEditingSettingReferenceEnum struct {
	NONE           VodEditingSettingReference
	MAX_BITRATE    VodEditingSettingReference
	MAX_RESOLUTION VodEditingSettingReference
}

func GetVodEditingSettingReferenceEnum() VodEditingSettingReferenceEnum {
	return VodEditingSettingReferenceEnum{
		NONE: VodEditingSettingReference{
			value: "NONE",
		},
		MAX_BITRATE: VodEditingSettingReference{
			value: "MAX_BITRATE",
		},
		MAX_RESOLUTION: VodEditingSettingReference{
			value: "MAX_RESOLUTION",
		},
	}
}

func (c VodEditingSettingReference) Value() string {
	return c.value
}

func (c VodEditingSettingReference) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VodEditingSettingReference) UnmarshalJSON(b []byte) error {
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
