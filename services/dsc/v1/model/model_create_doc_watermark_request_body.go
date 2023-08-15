package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type CreateDocWatermarkRequestBody struct {

	// 要嵌入水印的文档类型
	DocType *def.MultiPart `json:"doc_type"`

	// 输入文件有密码时，读取文件的密码， 最大支持长度256。如果Office文档有读密码或域控的权限密码，请输入读密码，或者有读权限的域控密码。
	FilePassword *def.MultiPart `json:"file_password,omitempty"`

	// 添加水印后给文件设置密码， 最大支持长度256。默认不加文档密码。
	MarkedFilePassword *def.MultiPart `json:"marked_file_password,omitempty"`

	// 添加水印后给文件设置只读密码， 最大支持长度256。默认不加只读密码。
	ReadonlyPassword *def.MultiPart `json:"readonly_password,omitempty"`

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkRequestBody", string(data)}, " ")
}

func (o *CreateDocWatermarkRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
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

func (c CreateDocWatermarkRequestBodyDocType) Value() string {
	return c.value
}

func (c CreateDocWatermarkRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDocWatermarkRequestBodyDocType) UnmarshalJSON(b []byte) error {
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

func (c CreateDocWatermarkRequestBodyVisibleType) Value() string {
	return c.value
}

func (c CreateDocWatermarkRequestBodyVisibleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDocWatermarkRequestBodyVisibleType) UnmarshalJSON(b []byte) error {
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
