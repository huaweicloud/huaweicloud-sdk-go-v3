package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandSimilarityGraphTaskResponse Response Object
type CreateDrugLigandSimilarityGraphTaskResponse struct {

	// 任务ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrugLigandSimilarityGraphTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandSimilarityGraphTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandSimilarityGraphTaskResponse", string(data)}, " ")
}
