package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImageSpritePara struct {

	// 采样类型，取值： - PERCENT：按百分比 - TIME：按时间间隔
	SampleType ImageSpriteParaSampleType `json:"sample_type"`

	// 采样间隔。 -当 sample_type 为 PERCENT 时，指定采样间隔的百分比，范围(0,100]。 -当 sample_type 为 TIME 时，指定采样间隔的时间，单位为秒,范围(0,360000]。
	SampleInterval int32 `json:"sample_interval"`

	// 雪碧图中小图的行数，行数*列数不得超过 100。
	RowCount int32 `json:"row_count"`

	// 雪碧图中小图的列数，行数*列数不得超过 100。
	ColumnCount int32 `json:"column_count"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [96, 4096]，单位：px。 - 当 width、height 均为 0，则分辨率同源； - 当 width 为 0，Height 非 0，则 Width 按比例缩放； - 当 width 非 0，Height 为 0，则 Height 按比例缩放； - 当 width、Height 均非 0，则分辨率按用户指定
	Width *int32 `json:"width,omitempty"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [96, 4096]，单位：px。 - 当 Width、Height 均为 0，则分辨率同源； - 当 Width 为 0，Height 非 0，则 Width 按比例缩放； - 当 Width 非 0，Height 为 0，则 Height 按比例缩放； - 当 Width、Height 均非 0，则分辨率按用户指定。
	Height *int32 `json:"height,omitempty"`

	// 分辨率自适应，可选值： - open：开启，此时，Width 代表视频的长边，Height 表示视频的短边； - close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。
	ResolutionAdaptive *ImageSpriteParaResolutionAdaptive `json:"resolution_adaptive,omitempty"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式： - stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“； - black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。 - white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。
	FillType *ImageSpriteParaFillType `json:"fill_type,omitempty"`

	// 图片格式，取值为 jpg、png。默认为 jpg。
	Format *ImageSpriteParaFormat `json:"format,omitempty"`
}

func (o ImageSpritePara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSpritePara struct{}"
	}

	return strings.Join([]string{"ImageSpritePara", string(data)}, " ")
}

type ImageSpriteParaSampleType struct {
	value string
}

type ImageSpriteParaSampleTypeEnum struct {
	PERCENT ImageSpriteParaSampleType
	TIME    ImageSpriteParaSampleType
}

func GetImageSpriteParaSampleTypeEnum() ImageSpriteParaSampleTypeEnum {
	return ImageSpriteParaSampleTypeEnum{
		PERCENT: ImageSpriteParaSampleType{
			value: "PERCENT",
		},
		TIME: ImageSpriteParaSampleType{
			value: "TIME",
		},
	}
}

func (c ImageSpriteParaSampleType) Value() string {
	return c.value
}

func (c ImageSpriteParaSampleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSpriteParaSampleType) UnmarshalJSON(b []byte) error {
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

type ImageSpriteParaResolutionAdaptive struct {
	value string
}

type ImageSpriteParaResolutionAdaptiveEnum struct {
	OPEN  ImageSpriteParaResolutionAdaptive
	CLOSE ImageSpriteParaResolutionAdaptive
}

func GetImageSpriteParaResolutionAdaptiveEnum() ImageSpriteParaResolutionAdaptiveEnum {
	return ImageSpriteParaResolutionAdaptiveEnum{
		OPEN: ImageSpriteParaResolutionAdaptive{
			value: "open",
		},
		CLOSE: ImageSpriteParaResolutionAdaptive{
			value: "close",
		},
	}
}

func (c ImageSpriteParaResolutionAdaptive) Value() string {
	return c.value
}

func (c ImageSpriteParaResolutionAdaptive) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSpriteParaResolutionAdaptive) UnmarshalJSON(b []byte) error {
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

type ImageSpriteParaFillType struct {
	value string
}

type ImageSpriteParaFillTypeEnum struct {
	STRETCH ImageSpriteParaFillType
	BLACK   ImageSpriteParaFillType
	WHITE   ImageSpriteParaFillType
}

func GetImageSpriteParaFillTypeEnum() ImageSpriteParaFillTypeEnum {
	return ImageSpriteParaFillTypeEnum{
		STRETCH: ImageSpriteParaFillType{
			value: "stretch",
		},
		BLACK: ImageSpriteParaFillType{
			value: "black",
		},
		WHITE: ImageSpriteParaFillType{
			value: "white",
		},
	}
}

func (c ImageSpriteParaFillType) Value() string {
	return c.value
}

func (c ImageSpriteParaFillType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSpriteParaFillType) UnmarshalJSON(b []byte) error {
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

type ImageSpriteParaFormat struct {
	value string
}

type ImageSpriteParaFormatEnum struct {
	JPG ImageSpriteParaFormat
	PNG ImageSpriteParaFormat
}

func GetImageSpriteParaFormatEnum() ImageSpriteParaFormatEnum {
	return ImageSpriteParaFormatEnum{
		JPG: ImageSpriteParaFormat{
			value: "jpg",
		},
		PNG: ImageSpriteParaFormat{
			value: "png",
		},
	}
}

func (c ImageSpriteParaFormat) Value() string {
	return c.value
}

func (c ImageSpriteParaFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImageSpriteParaFormat) UnmarshalJSON(b []byte) error {
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
