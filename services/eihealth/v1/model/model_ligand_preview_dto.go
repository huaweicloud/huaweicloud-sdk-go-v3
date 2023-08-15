package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandPreviewDto 配体预览信息
type LigandPreviewDto struct {

	// 配体索引（从0起编号）
	Index int32 `json:"index"`

	File *DrugFile `json:"file"`

	// 配体名称，若无名称则自动命名，格式为UNK+索引（从1起编号）
	Name string `json:"name"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`
}

func (o LigandPreviewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandPreviewDto struct{}"
	}

	return strings.Join([]string{"LigandPreviewDto", string(data)}, " ")
}
