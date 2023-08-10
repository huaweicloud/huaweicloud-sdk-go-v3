package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLigandSimilarityGraphTaskReq 创建配体相似度图任务请求体
type CreateLigandSimilarityGraphTaskReq struct {
	Mode *LigandSimilarityGraphMode `json:"mode"`

	// 配体列表
	Ligands []CreateLigandSimilarityGraphLigandDto `json:"ligands"`
}

func (o CreateLigandSimilarityGraphTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLigandSimilarityGraphTaskReq struct{}"
	}

	return strings.Join([]string{"CreateLigandSimilarityGraphTaskReq", string(data)}, " ")
}
