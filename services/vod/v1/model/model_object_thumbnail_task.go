package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ObjectThumbnailTask struct {

	// 采样类型。 取值如下： PERCENT：根据视频时长百分比间隔采样时的百分比值。 DOTS：指定时间点截图。选择同步截图时，需指定此类型。
	Type ObjectThumbnailTaskType `json:"type"`

	// 根据视频时长百分比间隔采样时的百分比值。
	Percent *int32 `json:"percent,omitempty"`

	// 指定时间截图的时间点数组 例如输入[1,3,5]，分别截取视频第1秒、第3秒、第5秒位置的图像帧 最多支持10个时间点
	Dots *[]int32 `json:"dots,omitempty"`

	// 截图输出文件名。 - 如果只抽一张图（即：按DOTS方式，指定1个时间点）则按该指定文件名输出图片。 - 如果抽多张图（即：按DOTS方式指定多个时间点或按TIME间隔截图）则输出图片名在该指定文件名基础上在增加时间点（示例：output_filename_10.jpg）。 - 如果按照PERCENT截图则按照output_filename_0.jpg,output_filename_1.jpg顺序命名
	OutputFilename *string `json:"output_filename,omitempty"`

	// 截图文件格式 取值如下：jpg、png
	Format *ObjectThumbnailTaskFormat `json:"format,omitempty"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式： - stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“； - black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。 - white：留白，保持视频宽高比不变，边缘剩余部分使用白色填充。
	FillType *ObjectThumbnailTaskFillType `json:"fill_type,omitempty"`

	// 图片宽度 取值范围： - [96,3840] - 0：自适应，保持原有宽高 单位：px
	Width *int32 `json:"width,omitempty"`

	// 图片高度 取值范围： - [96,2160] - 0：自适应，保持原有宽高 单位：px
	Height *int32 `json:"height,omitempty"`

	Output *ObsInfo `json:"output,omitempty"`
}

func (o ObjectThumbnailTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectThumbnailTask struct{}"
	}

	return strings.Join([]string{"ObjectThumbnailTask", string(data)}, " ")
}

type ObjectThumbnailTaskType struct {
	value string
}

type ObjectThumbnailTaskTypeEnum struct {
	PERCENT ObjectThumbnailTaskType
	DOTS    ObjectThumbnailTaskType
}

func GetObjectThumbnailTaskTypeEnum() ObjectThumbnailTaskTypeEnum {
	return ObjectThumbnailTaskTypeEnum{
		PERCENT: ObjectThumbnailTaskType{
			value: "PERCENT",
		},
		DOTS: ObjectThumbnailTaskType{
			value: "DOTS",
		},
	}
}

func (c ObjectThumbnailTaskType) Value() string {
	return c.value
}

func (c ObjectThumbnailTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectThumbnailTaskType) UnmarshalJSON(b []byte) error {
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

type ObjectThumbnailTaskFormat struct {
	value string
}

type ObjectThumbnailTaskFormatEnum struct {
	JPG ObjectThumbnailTaskFormat
	PNG ObjectThumbnailTaskFormat
}

func GetObjectThumbnailTaskFormatEnum() ObjectThumbnailTaskFormatEnum {
	return ObjectThumbnailTaskFormatEnum{
		JPG: ObjectThumbnailTaskFormat{
			value: "jpg",
		},
		PNG: ObjectThumbnailTaskFormat{
			value: "png",
		},
	}
}

func (c ObjectThumbnailTaskFormat) Value() string {
	return c.value
}

func (c ObjectThumbnailTaskFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectThumbnailTaskFormat) UnmarshalJSON(b []byte) error {
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

type ObjectThumbnailTaskFillType struct {
	value string
}

type ObjectThumbnailTaskFillTypeEnum struct {
	STRETCH ObjectThumbnailTaskFillType
	BLACK   ObjectThumbnailTaskFillType
	WHITE   ObjectThumbnailTaskFillType
	GAUSS   ObjectThumbnailTaskFillType
}

func GetObjectThumbnailTaskFillTypeEnum() ObjectThumbnailTaskFillTypeEnum {
	return ObjectThumbnailTaskFillTypeEnum{
		STRETCH: ObjectThumbnailTaskFillType{
			value: "stretch",
		},
		BLACK: ObjectThumbnailTaskFillType{
			value: "black",
		},
		WHITE: ObjectThumbnailTaskFillType{
			value: "white",
		},
		GAUSS: ObjectThumbnailTaskFillType{
			value: "gauss",
		},
	}
}

func (c ObjectThumbnailTaskFillType) Value() string {
	return c.value
}

func (c ObjectThumbnailTaskFillType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectThumbnailTaskFillType) UnmarshalJSON(b []byte) error {
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
