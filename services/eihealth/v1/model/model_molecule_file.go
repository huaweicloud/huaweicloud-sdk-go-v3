package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MoleculeFile struct {

	// 文件来源，支持用户私有数据中心、公共数据和源数据
	Source string `json:"source"`

	// 文件URL，用户私有数据中心为项目路径、公共数据场景为obs地址
	Url *string `json:"url,omitempty"`

	// 文件格式，支持PDB、SDF、MOL2、SMI，仅数据源为RAW时提供
	Format *string `json:"format,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供
	Data *string `json:"data,omitempty"`
}

func (o MoleculeFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoleculeFile struct{}"
	}

	return strings.Join([]string{"MoleculeFile", string(data)}, " ")
}
