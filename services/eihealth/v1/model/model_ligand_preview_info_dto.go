package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandPreviewInfoDto 配体预览信息，若成功则除reason字段均有，若失败也有可能存在formula和smiles字段
type LigandPreviewInfoDto struct {

	// 配体索引（从0起编号）
	Index int32 `json:"index"`

	// 配体名称，若无名称则自动命名，格式为UNK+索引（从1起编号）
	Name string `json:"name"`

	// 解析是否成功
	Success bool `json:"success"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	// 配体分子的化学式
	Formula *string `json:"formula,omitempty"`

	// 标识原始数据是否为3D
	Is3d *bool `json:"is_3d,omitempty"`

	Structure *LigandStructureDto `json:"structure,omitempty"`

	// 解析失败的理由
	Reason *string `json:"reason,omitempty"`
}

func (o LigandPreviewInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandPreviewInfoDto struct{}"
	}

	return strings.Join([]string{"LigandPreviewInfoDto", string(data)}, " ")
}
