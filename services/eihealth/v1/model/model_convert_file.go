package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConvertFile 待转换文件。
type ConvertFile struct {
	Source *DrugFileSource `json:"source"`

	// 文件URL，当数据源为外部网络数据时为https地址，为用户私有数据中心时为项目路径，为公共数据场景时为obs地址。
	Url *string `json:"url,omitempty"`

	// 文件格式，仅支持PDB、SDF、MOL2、SMI，仅数据源为RAW时提供。
	Format *string `json:"format,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供。
	Data *string `json:"data,omitempty"`
}

func (o ConvertFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConvertFile struct{}"
	}

	return strings.Join([]string{"ConvertFile", string(data)}, " ")
}
