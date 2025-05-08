package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ObjectImageSpriteTask struct {

	// 采样类型，取值： - PERCENT：按百分比 - TIME：按时间间隔
	SampleType ObjectImageSpriteTaskSampleType `json:"sample_type"`

	// 采样间隔。 -当 sample_type 为 PERCENT 时，指定采样间隔的百分比，(0<sample_interval<=100)。 -当 sample_type 为 TIME 时，指定采样间隔的时间，单位为秒(>0)。
	SampleInterval int32 `json:"sample_interval"`

	// 雪碧图中小图的行数，行数*列数不得超过 100。
	RowCount int32 `json:"row_count"`

	// 雪碧图中小图的列数，行数*列数不得超过 100。
	ColumnCount int32 `json:"column_count"`

	// 雪碧图中小图的宽度（或长边）的最大值，取值范围：0 和 [96, 4096]，单位：px。 - 当 width、height 均为 0，则分辨率同源； - 当 width 为 0，Height 非 0，则 Width 按比例缩放； - 当 width 非 0，Height 为 0，则 Height 按比例缩放； - 当 width、Height 均非 0，则分辨率按用户指定
	Width *int32 `json:"width,omitempty"`

	// 雪碧图中小图的高度（或短边）的最大值，取值范围：0 和 [96, 4096]，单位：px。 - 当 width、height 均为 0，则分辨率同源； - 当 width 为 0，height 非 0，则 width 按比例缩放； - 当 width 非 0，height 为 0，则 height 按比例缩放； - 当 width、height 均非 0，则分辨率按用户指定。
	Height *int32 `json:"height,omitempty"`

	// 分辨率自适应，可选值： - open：开启，此时，Width 代表视频的长边，Height 表示视频的短边； - close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。
	ResolutionAdaptive *ObjectImageSpriteTaskResolutionAdaptive `json:"resolution_adaptive,omitempty"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式： - stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“； - black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。 - white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。
	FillType *ObjectImageSpriteTaskFillType `json:"fill_type,omitempty"`

	// 图片格式，取值为 jpg、png。默认为 jpg。
	Format *ObjectImageSpriteTaskFormat `json:"format,omitempty"`

	// 截取雪碧图后，雪碧图图片文件的输出文件名，如果不填，则默认为：{inputName}_imageSprite_{雪碧图id}_{number}.{format}.{雪碧图id}和{number}从0开始递增
	OutputObjectName *string `json:"output_object_name,omitempty"`

	// 截取雪碧图后，Web VTT 文件的输出路径，只能为相对路径。如果不填，则默认为相对路径：{inputName}_imageSprite_{雪碧图_id}.vtt
	WebvttObjectName *string `json:"webvtt_object_name,omitempty"`

	Output *ObsInfo `json:"output,omitempty"`
}

func (o ObjectImageSpriteTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectImageSpriteTask struct{}"
	}

	return strings.Join([]string{"ObjectImageSpriteTask", string(data)}, " ")
}

type ObjectImageSpriteTaskSampleType struct {
	value string
}

type ObjectImageSpriteTaskSampleTypeEnum struct {
	PERCENT ObjectImageSpriteTaskSampleType
	TIME    ObjectImageSpriteTaskSampleType
}

func GetObjectImageSpriteTaskSampleTypeEnum() ObjectImageSpriteTaskSampleTypeEnum {
	return ObjectImageSpriteTaskSampleTypeEnum{
		PERCENT: ObjectImageSpriteTaskSampleType{
			value: "PERCENT",
		},
		TIME: ObjectImageSpriteTaskSampleType{
			value: "TIME",
		},
	}
}

func (c ObjectImageSpriteTaskSampleType) Value() string {
	return c.value
}

func (c ObjectImageSpriteTaskSampleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectImageSpriteTaskSampleType) UnmarshalJSON(b []byte) error {
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

type ObjectImageSpriteTaskResolutionAdaptive struct {
	value string
}

type ObjectImageSpriteTaskResolutionAdaptiveEnum struct {
	OPEN  ObjectImageSpriteTaskResolutionAdaptive
	CLOSE ObjectImageSpriteTaskResolutionAdaptive
}

func GetObjectImageSpriteTaskResolutionAdaptiveEnum() ObjectImageSpriteTaskResolutionAdaptiveEnum {
	return ObjectImageSpriteTaskResolutionAdaptiveEnum{
		OPEN: ObjectImageSpriteTaskResolutionAdaptive{
			value: "open",
		},
		CLOSE: ObjectImageSpriteTaskResolutionAdaptive{
			value: "close",
		},
	}
}

func (c ObjectImageSpriteTaskResolutionAdaptive) Value() string {
	return c.value
}

func (c ObjectImageSpriteTaskResolutionAdaptive) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectImageSpriteTaskResolutionAdaptive) UnmarshalJSON(b []byte) error {
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

type ObjectImageSpriteTaskFillType struct {
	value string
}

type ObjectImageSpriteTaskFillTypeEnum struct {
	STRETCH ObjectImageSpriteTaskFillType
	BLACK   ObjectImageSpriteTaskFillType
	WHITE   ObjectImageSpriteTaskFillType
	GAUSS   ObjectImageSpriteTaskFillType
}

func GetObjectImageSpriteTaskFillTypeEnum() ObjectImageSpriteTaskFillTypeEnum {
	return ObjectImageSpriteTaskFillTypeEnum{
		STRETCH: ObjectImageSpriteTaskFillType{
			value: "stretch",
		},
		BLACK: ObjectImageSpriteTaskFillType{
			value: "black",
		},
		WHITE: ObjectImageSpriteTaskFillType{
			value: "white",
		},
		GAUSS: ObjectImageSpriteTaskFillType{
			value: "gauss",
		},
	}
}

func (c ObjectImageSpriteTaskFillType) Value() string {
	return c.value
}

func (c ObjectImageSpriteTaskFillType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectImageSpriteTaskFillType) UnmarshalJSON(b []byte) error {
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

type ObjectImageSpriteTaskFormat struct {
	value string
}

type ObjectImageSpriteTaskFormatEnum struct {
	JPG ObjectImageSpriteTaskFormat
	PNG ObjectImageSpriteTaskFormat
}

func GetObjectImageSpriteTaskFormatEnum() ObjectImageSpriteTaskFormatEnum {
	return ObjectImageSpriteTaskFormatEnum{
		JPG: ObjectImageSpriteTaskFormat{
			value: "jpg",
		},
		PNG: ObjectImageSpriteTaskFormat{
			value: "png",
		},
	}
}

func (c ObjectImageSpriteTaskFormat) Value() string {
	return c.value
}

func (c ObjectImageSpriteTaskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectImageSpriteTaskFormat) UnmarshalJSON(b []byte) error {
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
