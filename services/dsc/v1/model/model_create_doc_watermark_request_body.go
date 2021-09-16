package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type CreateDocWatermarkRequestBody struct {
	// 要嵌入水印的文档类型
	DocType *def.MultiPart `json:"doc_type"`

	// 解密文件的密码， 最大支持长度256。添加水印后的文件不带密码。如果Office文档有读密码或域控的权限密码，请输入读密码，或者有读权限的域控密码。
	FilePassword *def.MultiPart `json:"file_password,omitempty"`

	// 明水印内容，与“blind_watermark”字段至少有一个不为空
	VisibleWatermark *def.MultiPart `json:"visible_watermark,omitempty"`

	// 明水印字体大小，取值为[1,100]，默认值50
	FontSize *def.MultiPart `json:"font_size,omitempty"`

	// 明水印旋转角度，逆时针方向，取值为[0,90]，默认值45
	Rotation *def.MultiPart `json:"rotation,omitempty"`

	// 明水印的透明度，取值[0,1]，默认值为0.3；
	Opacity *def.MultiPart `json:"opacity,omitempty"`

	// 暗水印内容，与“visible_watermark”字段至少有一个不为空
	BlindWatermark *def.MultiPart `json:"blind_watermark,omitempty"`

	// 要添加水印的文档
	File *def.FilePart `json:"file"`

	// 图形水印的字节流。图形文件的格式必须为“png”或“jpg”，否则返回参数错误；图像文件大小不超过1MB；在分段的请求体“Content-Disposition”部分，参数“name”的值必须为“image_mark”。
	ImageMark *def.FilePart `json:"image_mark,omitempty"`

	// 该字段为空时，默认为**TEXT**类型。  当该字段为**IMAGE**时: - 请求的表单中必须包含名为“image”的图像文件，图像格式必须为“png”或“jpg”，否则返回参数错误； - 图像文件大小不超过1MB； - “visible_watermark”，“font_size”，“rotation”和“opacity”字段无效。
	VisibleType *def.MultiPart `json:"visible_type,omitempty"`
}

func (o CreateDocWatermarkRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkRequestBody", string(data)}, " ")
}

type CreateDocWatermarkRequestBodyDocType struct {
	value string
}

type CreateDocWatermarkRequestBodyDocTypeEnum struct {
	WORD  CreateDocWatermarkRequestBodyDocType
	EXCEL CreateDocWatermarkRequestBodyDocType
	PDF   CreateDocWatermarkRequestBodyDocType
	PPT   CreateDocWatermarkRequestBodyDocType
}

func GetCreateDocWatermarkRequestBodyDocTypeEnum() CreateDocWatermarkRequestBodyDocTypeEnum {
	return CreateDocWatermarkRequestBodyDocTypeEnum{
		WORD: CreateDocWatermarkRequestBodyDocType{
			value: "WORD",
		},
		EXCEL: CreateDocWatermarkRequestBodyDocType{
			value: "EXCEL",
		},
		PDF: CreateDocWatermarkRequestBodyDocType{
			value: "PDF",
		},
		PPT: CreateDocWatermarkRequestBodyDocType{
			value: "PPT",
		},
	}
}

func (c CreateDocWatermarkRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateDocWatermarkRequestBodyDocType) UnmarshalJSON(b []byte) error {
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

type CreateDocWatermarkRequestBodyVisibleType struct {
	value string
}

type CreateDocWatermarkRequestBodyVisibleTypeEnum struct {
	TEXT  CreateDocWatermarkRequestBodyVisibleType
	IMAGE CreateDocWatermarkRequestBodyVisibleType
}

func GetCreateDocWatermarkRequestBodyVisibleTypeEnum() CreateDocWatermarkRequestBodyVisibleTypeEnum {
	return CreateDocWatermarkRequestBodyVisibleTypeEnum{
		TEXT: CreateDocWatermarkRequestBodyVisibleType{
			value: "TEXT",
		},
		IMAGE: CreateDocWatermarkRequestBodyVisibleType{
			value: "IMAGE",
		},
	}
}

func (c CreateDocWatermarkRequestBodyVisibleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateDocWatermarkRequestBodyVisibleType) UnmarshalJSON(b []byte) error {
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
