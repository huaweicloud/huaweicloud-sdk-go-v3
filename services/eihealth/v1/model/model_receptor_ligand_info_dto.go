package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReceptorLigandInfoDto 受体中配体信息，若成功则除reason字段均有，若失败也有可能存在formula和smiles字段
type ReceptorLigandInfoDto struct {

	// 配体索引（从0起编号）
	Index int32 `json:"index"`

	// 配体名称，即配体所在的残基表示
	Name string `json:"name"`

	// 解析是否成功
	Success bool `json:"success"`

	// 分子SMILES表达式
	Smiles *string `json:"smiles,omitempty"`

	// 配体分子的化学式
	Formula *string `json:"formula,omitempty"`

	Structure *LigandStructureDto `json:"structure,omitempty"`

	BoundingBox *BoundingBoxDto `json:"bounding_box,omitempty"`
}

func (o ReceptorLigandInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReceptorLigandInfoDto struct{}"
	}

	return strings.Join([]string{"ReceptorLigandInfoDto", string(data)}, " ")
}
