package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 镜像导出请求体
type ExportImageRequestBody struct {

	// 目的文件的URL，格式：<bucket>:<file>。 说明：此处的OBS桶和镜像文件的存储类别必须是OBS标准存储。
	BucketUrl string `json:"bucket_url"`

	// 文件格式，支持qcow2、vhd、zvhd和vmdk。
	FileFormat ExportImageRequestBodyFileFormat `json:"file_format"`

	// 是否使用快速导出，取值为true或者false。 说明：若使用快速导出，则无法指定file_format参数。
	IsQuickExport *bool `json:"is_quick_export,omitempty"`
}

func (o ExportImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportImageRequestBody struct{}"
	}

	return strings.Join([]string{"ExportImageRequestBody", string(data)}, " ")
}

type ExportImageRequestBodyFileFormat struct {
	value string
}

type ExportImageRequestBodyFileFormatEnum struct {
	QCOW2 ExportImageRequestBodyFileFormat
	VHD   ExportImageRequestBodyFileFormat
	ZVHD  ExportImageRequestBodyFileFormat
	VMDK  ExportImageRequestBodyFileFormat
}

func GetExportImageRequestBodyFileFormatEnum() ExportImageRequestBodyFileFormatEnum {
	return ExportImageRequestBodyFileFormatEnum{
		QCOW2: ExportImageRequestBodyFileFormat{
			value: "qcow2",
		},
		VHD: ExportImageRequestBodyFileFormat{
			value: "vhd",
		},
		ZVHD: ExportImageRequestBodyFileFormat{
			value: "zvhd",
		},
		VMDK: ExportImageRequestBodyFileFormat{
			value: "vmdk",
		},
	}
}

func (c ExportImageRequestBodyFileFormat) Value() string {
	return c.value
}

func (c ExportImageRequestBodyFileFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportImageRequestBodyFileFormat) UnmarshalJSON(b []byte) error {
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
