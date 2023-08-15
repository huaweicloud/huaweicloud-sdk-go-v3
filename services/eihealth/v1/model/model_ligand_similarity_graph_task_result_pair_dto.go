package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandSimilarityGraphTaskResultPairDto 配体相似度图任务结果对，成功则返回similarity，失败则返回reason
type LigandSimilarityGraphTaskResultPairDto struct {

	// 两个配体名称
	Ligands *[]string `json:"ligands,omitempty"`

	// 相似度计算是否成功
	Success bool `json:"success"`

	// 配体对之间的相似度
	Similarity *float32 `json:"similarity,omitempty"`

	// 相似度计算失败的理由
	Reason *string `json:"reason,omitempty"`
}

func (o LigandSimilarityGraphTaskResultPairDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandSimilarityGraphTaskResultPairDto struct{}"
	}

	return strings.Join([]string{"LigandSimilarityGraphTaskResultPairDto", string(data)}, " ")
}
