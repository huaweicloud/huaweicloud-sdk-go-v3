package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrugLigandSimilarityGraphTaskResponse Response Object
type ShowDrugLigandSimilarityGraphTaskResponse struct {

	// 任务状态
	Status *string `json:"status,omitempty"`

	Result         *LigandSimilarityGraphTaskResultDto `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowDrugLigandSimilarityGraphTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrugLigandSimilarityGraphTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowDrugLigandSimilarityGraphTaskResponse", string(data)}, " ")
}
