package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandSimilarityGraphLigandDto 创建配体相似度图任务配体
type CreateLigandSimilarityGraphLigandDto struct {

	// 配体分子唯一名字，受体中的建议使用\"{氨基酸}:{链}:{编号}\"
	Name string `json:"name"`

	// 分子SMILES表达式
	Smiles string `json:"smiles"`

	// 配体是否为主要配体，在中心模式下，必须指定1个主要配体
	Main *bool `json:"main,omitempty"`
}

func (o CreateLigandSimilarityGraphLigandDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandSimilarityGraphLigandDto struct{}"
	}

	return strings.Join([]string{"CreateLigandSimilarityGraphLigandDto", string(data)}, " ")
}
