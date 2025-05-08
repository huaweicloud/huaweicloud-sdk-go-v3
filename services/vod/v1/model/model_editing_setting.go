package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EditingSetting struct {

	// 分辨率自适应策略, 选值：OPEN, CLOSE（默认OPEN） 输入为OPEN时，width表示长边，height表示短边 输入为CLOSE时，width表示宽度，height表示长度
	ResolutionAdaptation *EditingSettingResolutionAdaptation `json:"resolution_adaptation,omitempty"`

	// 分辨率上采样开关, 选值：ON, OFF（默认OFF）
	ResolutionUpsample *EditingSettingResolutionUpsample `json:"resolution_upsample,omitempty"`

	// 输出封装格式，HLS、MP4（默认MP4）、MP3、MOV、FLV、AVI
	Format *EditingSettingFormat `json:"format,omitempty"`

	// 输出宽或长边，整型，输入小数向下取整，默认0，按源  - 当 width、height 均为 0，则分辨率取片源分辨率； - 当 width 为 0，height 非 0，则 width 按片源分辨率比例缩放； - 当 width 非 0，height 为 0，则 height 按片源分辨率比例缩放； - 当 width、height 均非 0，则分辨率按用户指定。
	Width *int32 `json:"width,omitempty"`

	// 输出高或短边，整型，输入小数向下取整，默认0，按源  - 当 Width、Height 均为 0，则分辨率取片源分辨率； - 当 Width 为 0，Height 非 0，则 Width 按片源分辨率比例缩放； - 当 Width 非 0，Height 为 0，则 Height 按片源分辨率比例缩放； - 当 Width、Height 均非 0，则分辨率按用户指定。
	Height *int32 `json:"height,omitempty"`

	// 输出参考基准，可选，默认为空  - NONE 输出分辨率按输入的第一个片源为主，码率按输出分辨率自适应，帧率固定输出25fps - MAX_BITRATE 按码率最大的输入片源参数为基准 - MAX_RESOLUTION 按分辨率最大的输入片源参数为基准
	Reference *EditingSettingReference `json:"reference,omitempty"`

	// 视频编码格式。 取值如下： - 1：VIDEO_CODEC_H264 - 2：VIDEO_CODEC_H265
	VideoCodec *int32 `json:"video_codec,omitempty"`
}

func (o EditingSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingSetting struct{}"
	}

	return strings.Join([]string{"EditingSetting", string(data)}, " ")
}

type EditingSettingResolutionAdaptation struct {
	value string
}

type EditingSettingResolutionAdaptationEnum struct {
	OPEN  EditingSettingResolutionAdaptation
	CLOSE EditingSettingResolutionAdaptation
}

func GetEditingSettingResolutionAdaptationEnum() EditingSettingResolutionAdaptationEnum {
	return EditingSettingResolutionAdaptationEnum{
		OPEN: EditingSettingResolutionAdaptation{
			value: "OPEN",
		},
		CLOSE: EditingSettingResolutionAdaptation{
			value: "CLOSE",
		},
	}
}

func (c EditingSettingResolutionAdaptation) Value() string {
	return c.value
}

func (c EditingSettingResolutionAdaptation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditingSettingResolutionAdaptation) UnmarshalJSON(b []byte) error {
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

type EditingSettingResolutionUpsample struct {
	value string
}

type EditingSettingResolutionUpsampleEnum struct {
	ON  EditingSettingResolutionUpsample
	OFF EditingSettingResolutionUpsample
}

func GetEditingSettingResolutionUpsampleEnum() EditingSettingResolutionUpsampleEnum {
	return EditingSettingResolutionUpsampleEnum{
		ON: EditingSettingResolutionUpsample{
			value: "ON",
		},
		OFF: EditingSettingResolutionUpsample{
			value: "OFF",
		},
	}
}

func (c EditingSettingResolutionUpsample) Value() string {
	return c.value
}

func (c EditingSettingResolutionUpsample) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditingSettingResolutionUpsample) UnmarshalJSON(b []byte) error {
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

type EditingSettingFormat struct {
	value string
}

type EditingSettingFormatEnum struct {
	MP4 EditingSettingFormat
	HLS EditingSettingFormat
	MP3 EditingSettingFormat
	MOV EditingSettingFormat
	AVI EditingSettingFormat
	FLV EditingSettingFormat
}

func GetEditingSettingFormatEnum() EditingSettingFormatEnum {
	return EditingSettingFormatEnum{
		MP4: EditingSettingFormat{
			value: "MP4",
		},
		HLS: EditingSettingFormat{
			value: "HLS",
		},
		MP3: EditingSettingFormat{
			value: "MP3",
		},
		MOV: EditingSettingFormat{
			value: "MOV",
		},
		AVI: EditingSettingFormat{
			value: "AVI",
		},
		FLV: EditingSettingFormat{
			value: "FLV",
		},
	}
}

func (c EditingSettingFormat) Value() string {
	return c.value
}

func (c EditingSettingFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditingSettingFormat) UnmarshalJSON(b []byte) error {
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

type EditingSettingReference struct {
	value string
}

type EditingSettingReferenceEnum struct {
	NONE           EditingSettingReference
	MAX_BITRATE    EditingSettingReference
	MAX_RESOLUTION EditingSettingReference
}

func GetEditingSettingReferenceEnum() EditingSettingReferenceEnum {
	return EditingSettingReferenceEnum{
		NONE: EditingSettingReference{
			value: "NONE",
		},
		MAX_BITRATE: EditingSettingReference{
			value: "MAX_BITRATE",
		},
		MAX_RESOLUTION: EditingSettingReference{
			value: "MAX_RESOLUTION",
		},
	}
}

func (c EditingSettingReference) Value() string {
	return c.value
}

func (c EditingSettingReference) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditingSettingReference) UnmarshalJSON(b []byte) error {
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
