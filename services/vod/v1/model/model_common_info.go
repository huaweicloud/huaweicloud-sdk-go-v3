package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CommonInfo struct {

	// 高清低码开关<br/>
	Pvc *bool `json:"pvc,omitempty"`

	// 视频编码格式<br/>
	VideoCodec *CommonInfoVideoCodec `json:"video_codec,omitempty"`

	// 音频编码格式<br/> AAC：AAC格式 (default)<br/> HEAAC1：HEAAC1格式<br/> HEAAC2：HEAAC2格式<br/> MP3：MP3格式<br/>
	AudioCodec *CommonInfoAudioCodec `json:"audio_codec,omitempty"`

	// 黑边剪裁类型<br/>
	IsBlackCut *bool `json:"is_black_cut,omitempty"`

	// 格式<br/>
	Format CommonInfoFormat `json:"format"`

	// 分片时长(默认为5秒)<br/>
	HlsInterval *int32 `json:"hls_interval,omitempty"`

	// 上采样<br/>
	Upsample *bool `json:"upsample,omitempty"`

	// SHORT：短边自适应<br/> LONG：长边自适应<br/> NONE：宽高自适应<br/>
	Adaptation *CommonInfoAdaptation `json:"adaptation,omitempty"`

	// 编码质量等级，取值[0,2] 0表示当前现网方式默认方式，1表示转码效率优先，2表示转码质量优先。<br/>
	Preset *int32 `json:"preset,omitempty"`

	// I帧最大间隔，取值范围：[2，10]。默认值：5，单位秒。<br/>
	MaxIframesInterval *int32 `json:"max_iframes_interval,omitempty"`

	// 转码后音频是否独立存储。<br/>
	HlsAudioSeparate *bool `json:"hls_audio_separate,omitempty"`

	// 分片的封装格式，目前支持TS和FMP4，默认TS格式。
	HlsSegmentType *CommonInfoHlsSegmentType `json:"hls_segment_type,omitempty"`
}

func (o CommonInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonInfo struct{}"
	}

	return strings.Join([]string{"CommonInfo", string(data)}, " ")
}

type CommonInfoVideoCodec struct {
	value string
}

type CommonInfoVideoCodecEnum struct {
	H264 CommonInfoVideoCodec
	H265 CommonInfoVideoCodec
}

func GetCommonInfoVideoCodecEnum() CommonInfoVideoCodecEnum {
	return CommonInfoVideoCodecEnum{
		H264: CommonInfoVideoCodec{
			value: "H264",
		},
		H265: CommonInfoVideoCodec{
			value: "H265",
		},
	}
}

func (c CommonInfoVideoCodec) Value() string {
	return c.value
}

func (c CommonInfoVideoCodec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonInfoVideoCodec) UnmarshalJSON(b []byte) error {
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

type CommonInfoAudioCodec struct {
	value string
}

type CommonInfoAudioCodecEnum struct {
	AAC    CommonInfoAudioCodec
	HEAAC1 CommonInfoAudioCodec
	HEAAC2 CommonInfoAudioCodec
	MP3    CommonInfoAudioCodec
}

func GetCommonInfoAudioCodecEnum() CommonInfoAudioCodecEnum {
	return CommonInfoAudioCodecEnum{
		AAC: CommonInfoAudioCodec{
			value: "AAC",
		},
		HEAAC1: CommonInfoAudioCodec{
			value: "HEAAC1",
		},
		HEAAC2: CommonInfoAudioCodec{
			value: "HEAAC2",
		},
		MP3: CommonInfoAudioCodec{
			value: "MP3",
		},
	}
}

func (c CommonInfoAudioCodec) Value() string {
	return c.value
}

func (c CommonInfoAudioCodec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonInfoAudioCodec) UnmarshalJSON(b []byte) error {
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

type CommonInfoFormat struct {
	value string
}

type CommonInfoFormatEnum struct {
	MP4      CommonInfoFormat
	HLS      CommonInfoFormat
	DASH     CommonInfoFormat
	DASH_HLS CommonInfoFormat
	MP3      CommonInfoFormat
	ADTS     CommonInfoFormat
	UNKNOW   CommonInfoFormat
}

func GetCommonInfoFormatEnum() CommonInfoFormatEnum {
	return CommonInfoFormatEnum{
		MP4: CommonInfoFormat{
			value: "MP4",
		},
		HLS: CommonInfoFormat{
			value: "HLS",
		},
		DASH: CommonInfoFormat{
			value: "DASH",
		},
		DASH_HLS: CommonInfoFormat{
			value: "DASH_HLS",
		},
		MP3: CommonInfoFormat{
			value: "MP3",
		},
		ADTS: CommonInfoFormat{
			value: "ADTS",
		},
		UNKNOW: CommonInfoFormat{
			value: "UNKNOW",
		},
	}
}

func (c CommonInfoFormat) Value() string {
	return c.value
}

func (c CommonInfoFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonInfoFormat) UnmarshalJSON(b []byte) error {
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

type CommonInfoAdaptation struct {
	value string
}

type CommonInfoAdaptationEnum struct {
	SHORT CommonInfoAdaptation
	LONG  CommonInfoAdaptation
	NONE  CommonInfoAdaptation
}

func GetCommonInfoAdaptationEnum() CommonInfoAdaptationEnum {
	return CommonInfoAdaptationEnum{
		SHORT: CommonInfoAdaptation{
			value: "SHORT",
		},
		LONG: CommonInfoAdaptation{
			value: "LONG",
		},
		NONE: CommonInfoAdaptation{
			value: "NONE",
		},
	}
}

func (c CommonInfoAdaptation) Value() string {
	return c.value
}

func (c CommonInfoAdaptation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonInfoAdaptation) UnmarshalJSON(b []byte) error {
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

type CommonInfoHlsSegmentType struct {
	value string
}

type CommonInfoHlsSegmentTypeEnum struct {
	TS   CommonInfoHlsSegmentType
	FMP4 CommonInfoHlsSegmentType
}

func GetCommonInfoHlsSegmentTypeEnum() CommonInfoHlsSegmentTypeEnum {
	return CommonInfoHlsSegmentTypeEnum{
		TS: CommonInfoHlsSegmentType{
			value: "TS",
		},
		FMP4: CommonInfoHlsSegmentType{
			value: "FMP4",
		},
	}
}

func (c CommonInfoHlsSegmentType) Value() string {
	return c.value
}

func (c CommonInfoHlsSegmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonInfoHlsSegmentType) UnmarshalJSON(b []byte) error {
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
