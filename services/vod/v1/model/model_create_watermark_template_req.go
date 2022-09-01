package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateWatermarkTemplateReq struct {

	// 水印模板名称。
	Name string `json:"name" xml:"name"`

	// 水印类型，当前只支持Image（图片水印）。
	WatermarkType *CreateWatermarkTemplateReqWatermarkType `json:"watermark_type,omitempty" xml:"watermark_type"`

	// type设置为Image时有效。  目前包括： - Original：只做简单缩放，不做其他处理 - Transparent：图片底色透明 - Grayed：彩色图片变灰
	ImageProcess *CreateWatermarkTemplateReqImageProcess `json:"image_process,omitempty" xml:"image_process"`

	// 水印图片相对输出视频的水平偏移量，默认值是0。
	Dx *string `json:"dx,omitempty" xml:"dx"`

	// 水印图片相对输出视频的垂直偏移量，默认值是0。
	Dy *string `json:"dy,omitempty" xml:"dy"`

	// 水印的位置。
	Position *CreateWatermarkTemplateReqPosition `json:"position,omitempty" xml:"position"`

	// 水印图片宽。
	Width *string `json:"width,omitempty" xml:"width"`

	// 水印图片高。
	Height *string `json:"height,omitempty" xml:"height"`

	// 水印开始时间，与\"timeline_duration\"配合使用。 取值范围:[0, END)。\"END\"表示视频结束时间。 单位:秒。
	TimelineStart *string `json:"timeline_start,omitempty" xml:"timeline_start"`

	// 水印持续时间，与\"timeline_start\"配合使用。 取值范围:(0,END-开始时间]。\"END\"表示视频结束时间。 单位:秒。 默认:END。
	TimelineDuration *string `json:"timeline_duration,omitempty" xml:"timeline_duration"`

	// 水印图片格式类型。
	Type string `json:"type" xml:"type"`

	// 水印图片MD5值。
	Md5 *string `json:"md5,omitempty" xml:"md5"`
}

func (o CreateWatermarkTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateReq", string(data)}, " ")
}

type CreateWatermarkTemplateReqWatermarkType struct {
	value string
}

type CreateWatermarkTemplateReqWatermarkTypeEnum struct {
	IMAGE CreateWatermarkTemplateReqWatermarkType
	TEXT  CreateWatermarkTemplateReqWatermarkType
}

func GetCreateWatermarkTemplateReqWatermarkTypeEnum() CreateWatermarkTemplateReqWatermarkTypeEnum {
	return CreateWatermarkTemplateReqWatermarkTypeEnum{
		IMAGE: CreateWatermarkTemplateReqWatermarkType{
			value: "IMAGE",
		},
		TEXT: CreateWatermarkTemplateReqWatermarkType{
			value: "TEXT",
		},
	}
}

func (c CreateWatermarkTemplateReqWatermarkType) Value() string {
	return c.value
}

func (c CreateWatermarkTemplateReqWatermarkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqWatermarkType) UnmarshalJSON(b []byte) error {
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

type CreateWatermarkTemplateReqImageProcess struct {
	value string
}

type CreateWatermarkTemplateReqImageProcessEnum struct {
	ORIGINAL    CreateWatermarkTemplateReqImageProcess
	TRANSPARENT CreateWatermarkTemplateReqImageProcess
	GRAYED      CreateWatermarkTemplateReqImageProcess
}

func GetCreateWatermarkTemplateReqImageProcessEnum() CreateWatermarkTemplateReqImageProcessEnum {
	return CreateWatermarkTemplateReqImageProcessEnum{
		ORIGINAL: CreateWatermarkTemplateReqImageProcess{
			value: "ORIGINAL",
		},
		TRANSPARENT: CreateWatermarkTemplateReqImageProcess{
			value: "TRANSPARENT",
		},
		GRAYED: CreateWatermarkTemplateReqImageProcess{
			value: "GRAYED",
		},
	}
}

func (c CreateWatermarkTemplateReqImageProcess) Value() string {
	return c.value
}

func (c CreateWatermarkTemplateReqImageProcess) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqImageProcess) UnmarshalJSON(b []byte) error {
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

type CreateWatermarkTemplateReqPosition struct {
	value string
}

type CreateWatermarkTemplateReqPositionEnum struct {
	TOPRIGHT    CreateWatermarkTemplateReqPosition
	TOPLEFT     CreateWatermarkTemplateReqPosition
	BOTTOMRIGHT CreateWatermarkTemplateReqPosition
	BOTTOMLEFT  CreateWatermarkTemplateReqPosition
}

func GetCreateWatermarkTemplateReqPositionEnum() CreateWatermarkTemplateReqPositionEnum {
	return CreateWatermarkTemplateReqPositionEnum{
		TOPRIGHT: CreateWatermarkTemplateReqPosition{
			value: "TOPRIGHT",
		},
		TOPLEFT: CreateWatermarkTemplateReqPosition{
			value: "TOPLEFT",
		},
		BOTTOMRIGHT: CreateWatermarkTemplateReqPosition{
			value: "BOTTOMRIGHT",
		},
		BOTTOMLEFT: CreateWatermarkTemplateReqPosition{
			value: "BOTTOMLEFT",
		},
	}
}

func (c CreateWatermarkTemplateReqPosition) Value() string {
	return c.value
}

func (c CreateWatermarkTemplateReqPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWatermarkTemplateReqPosition) UnmarshalJSON(b []byte) error {
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
