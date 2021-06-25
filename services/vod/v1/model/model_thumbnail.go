package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 截图参数
type Thumbnail struct {
	// 采样类型。支持三种采样方式（当前只支持“time”）： “percent”：根据视频时长的百分比间隔采样 “time”：根据时间间隔采样 “dots” : 指定时间点截图

	Type ThumbnailType `json:"type"`
	// 根据视频时长百分比间隔采样时的百分比值

	Percent *int32 `json:"percent,omitempty"`
	// 根据时间间隔采样时的时间间隔值

	Time *int32 `json:"time,omitempty"`
	// 指定时间截图时的时间点数组

	Dots *[]int32 `json:"dots,omitempty"`
	// 该值表示指定第几张截图作为封面(从1开始)

	CoverPosition *int32 `json:"cover_position,omitempty"`
	// 截图文件格式，枚举值（jpg，png，webP）。当前只支持jpg。1 : jpg。

	Format *ThumbnailFormat `json:"format,omitempty"`
	// 纵横比（保留,图像缩放方式）。0：自适应（保持原有宽高比）。1：16:9

	AspectRatio *ThumbnailAspectRatio `json:"aspect_ratio,omitempty"`
	// 截图最长边的尺寸（单位：像素）（宽边尺寸按照该尺寸与原始视频像素等比缩放计算）

	MaxLength *int32 `json:"max_length,omitempty"`
}

func (o Thumbnail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Thumbnail struct{}"
	}

	return strings.Join([]string{"Thumbnail", string(data)}, " ")
}

type ThumbnailType struct {
	value string
}

type ThumbnailTypeEnum struct {
	TIME    ThumbnailType
	PERCENT ThumbnailType
	DOTS    ThumbnailType
}

func GetThumbnailTypeEnum() ThumbnailTypeEnum {
	return ThumbnailTypeEnum{
		TIME: ThumbnailType{
			value: "time",
		},
		PERCENT: ThumbnailType{
			value: "percent",
		},
		DOTS: ThumbnailType{
			value: "dots",
		},
	}
}

func (c ThumbnailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ThumbnailType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ThumbnailFormat struct {
	value int32
}

type ThumbnailFormatEnum struct {
	E_1 ThumbnailFormat
}

func GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatEnum{
		E_1: ThumbnailFormat{
			value: 1,
		},
	}
}

func (c ThumbnailFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ThumbnailFormat) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ThumbnailAspectRatio struct {
	value int32
}

type ThumbnailAspectRatioEnum struct {
	E_0 ThumbnailAspectRatio
	E_1 ThumbnailAspectRatio
}

func GetThumbnailAspectRatioEnum() ThumbnailAspectRatioEnum {
	return ThumbnailAspectRatioEnum{
		E_0: ThumbnailAspectRatio{
			value: 0,
		}, E_1: ThumbnailAspectRatio{
			value: 1,
		},
	}
}

func (c ThumbnailAspectRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ThumbnailAspectRatio) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
