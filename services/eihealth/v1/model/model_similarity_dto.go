package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimilarityDto struct {

	// 配体对
	LigandIds []string `json:"ligand_ids"`
}

func (o SimilarityDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimilarityDto struct{}"
	}

	return strings.Join([]string{"SimilarityDto", string(data)}, " ")
}
