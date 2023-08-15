package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandSimilarityGraphTaskResultDto 配体相似度图任务结果
type LigandSimilarityGraphTaskResultDto struct {

	// 配体相似度图任务结果对列表
	Pairs []LigandSimilarityGraphTaskResultPairDto `json:"pairs"`
}

func (o LigandSimilarityGraphTaskResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandSimilarityGraphTaskResultDto struct{}"
	}

	return strings.Join([]string{"LigandSimilarityGraphTaskResultDto", string(data)}, " ")
}
