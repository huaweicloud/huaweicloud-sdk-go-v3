package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OBS文档嵌入水印请求体
type CreateDocWatermarkByAddressRequestBody struct {
	// 项目所在region的id，如北京一为：cn-north-1。

	RegionId string `json:"region_id"`
	// 待添加水印的文档地址，当前只支持华为云OBS对象，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。

	SrcFile string `json:"src_file"`
	// 待嵌入水印的文档类型。

	DocType CreateDocWatermarkByAddressRequestBodyDocType `json:"doc_type"`
	// 添加水印后的文档存放地址，格式和要求同src_file字段，不设置时，默认取src_file的值，即添加水印后覆盖原文件。

	DstFile *string `json:"dst_file,omitempty"`
	// 暗文字水印内容，与“visible_watermark”字段至少有一个不为空

	BlindWatermark *string `json:"blind_watermark,omitempty"`
	// 明文字水印内容，与暗水印“blind_watermark”字段至少有一个不为空。

	VisibleWatermark *string `json:"visible_watermark,omitempty"`
	// 待嵌入的图形明水印文件的地址, 字段格式要求同src_file字段，图形文件的格式必须为“png”或“jpg”，否则返回参数错误；图像文件大小不超过1MB

	ImageMark *string `json:"image_mark,omitempty"`
	// 该字段控制明水印嵌入文字还是图片。默认为**TEXT**类型，需填写visible_watermark字段设置明文字水印； 当该字段为**IMAGE**时，需填写image_watermark字段设置水印图片地址此时 ，“visible_watermark”，“font_size”，“rotation”和“opacity”字段无效。

	VisibleType *CreateDocWatermarkByAddressRequestBodyVisibleType `json:"visible_type,omitempty"`
	// 待加水印文件有密码时，读取文件的密码， 最大支持长度256。如果Office文档有读密码或域控的权限密码，请输入读密码，或者有读权限的域控密码。

	FilePassword *string `json:"file_password,omitempty"`
	// 添加水印后给文件设置密码， 最大支持长度256。默认不加文档密码。

	MarkedFilePassword *string `json:"marked_file_password,omitempty"`
	// 添加水印后给文件设置只读密码， 最大支持长度256。默认不加只读密码。

	ReadonlyPassword *string `json:"readonly_password,omitempty"`
	// 明水印字体大小，取值为[1,100]，默认值50

	Front *int32 `json:"front,omitempty"`
	// 明水印旋转角度，逆时针方向，取值为[0,90]，默认值45。

	Rotation *int32 `json:"rotation,omitempty"`
	// 明水印的透明度，取值[0,1]，默认值为0.3；

	Opacity *float32 `json:"opacity,omitempty"`
}

func (o CreateDocWatermarkByAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocWatermarkByAddressRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDocWatermarkByAddressRequestBody", string(data)}, " ")
}

type CreateDocWatermarkByAddressRequestBodyDocType struct {
	value string
}

type CreateDocWatermarkByAddressRequestBodyDocTypeEnum struct {
	WORD  CreateDocWatermarkByAddressRequestBodyDocType
	EXCEL CreateDocWatermarkByAddressRequestBodyDocType
	PDF   CreateDocWatermarkByAddressRequestBodyDocType
	PPT   CreateDocWatermarkByAddressRequestBodyDocType
}

func GetCreateDocWatermarkByAddressRequestBodyDocTypeEnum() CreateDocWatermarkByAddressRequestBodyDocTypeEnum {
	return CreateDocWatermarkByAddressRequestBodyDocTypeEnum{
		WORD: CreateDocWatermarkByAddressRequestBodyDocType{
			value: "WORD",
		},
		EXCEL: CreateDocWatermarkByAddressRequestBodyDocType{
			value: "EXCEL",
		},
		PDF: CreateDocWatermarkByAddressRequestBodyDocType{
			value: "PDF",
		},
		PPT: CreateDocWatermarkByAddressRequestBodyDocType{
			value: "PPT",
		},
	}
}

func (c CreateDocWatermarkByAddressRequestBodyDocType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDocWatermarkByAddressRequestBodyDocType) UnmarshalJSON(b []byte) error {
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

type CreateDocWatermarkByAddressRequestBodyVisibleType struct {
	value string
}

type CreateDocWatermarkByAddressRequestBodyVisibleTypeEnum struct {
	TEXT  CreateDocWatermarkByAddressRequestBodyVisibleType
	IMAGE CreateDocWatermarkByAddressRequestBodyVisibleType
}

func GetCreateDocWatermarkByAddressRequestBodyVisibleTypeEnum() CreateDocWatermarkByAddressRequestBodyVisibleTypeEnum {
	return CreateDocWatermarkByAddressRequestBodyVisibleTypeEnum{
		TEXT: CreateDocWatermarkByAddressRequestBodyVisibleType{
			value: "TEXT",
		},
		IMAGE: CreateDocWatermarkByAddressRequestBodyVisibleType{
			value: "IMAGE",
		},
	}
}

func (c CreateDocWatermarkByAddressRequestBodyVisibleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDocWatermarkByAddressRequestBodyVisibleType) UnmarshalJSON(b []byte) error {
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
