package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PocketFragment 靶点口袋分子优化片段
type PocketFragment struct {
	Source *DrugFileSource `json:"source,omitempty"`

	// 文件URL，当数据源为外部网络数据时为https地址；用户私有数据中心为项目路径、公共数据场景为obs地址
	Url *string `json:"url,omitempty"`

	// 文件格式，支持PDB、SDF、MOL2、SMI，仅数据源为RAW时提供
	Format *string `json:"format,omitempty"`

	// 原始配体名称，仅RAW类型时用于配体名称标识
	Name *string `json:"name,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供
	Data *string `json:"data,omitempty"`

	Edited *EditedLigand `json:"edited,omitempty"`

	LabelSites *LabelSite `json:"label_sites,omitempty"`
}

func (o PocketFragment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PocketFragment struct{}"
	}

	return strings.Join([]string{"PocketFragment", string(data)}, " ")
}
