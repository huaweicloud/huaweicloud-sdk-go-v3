package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunFormatConverterReq struct {
	File *ConvertFile `json:"file"`

	// 待转换文件格式，支持PDB、SDF、MOL2、SMI。
	InFormat string `json:"in_format"`

	// 转换后文件格式，支持PDB、SDF、MOL2、SMI。
	OutFormat string `json:"out_format"`
}

func (o RunFormatConverterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFormatConverterReq struct{}"
	}

	return strings.Join([]string{"RunFormatConverterReq", string(data)}, " ")
}
