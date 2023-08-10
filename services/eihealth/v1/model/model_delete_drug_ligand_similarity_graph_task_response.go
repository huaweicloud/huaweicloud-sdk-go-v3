package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugLigandSimilarityGraphTaskResponse Response Object
type DeleteDrugLigandSimilarityGraphTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDrugLigandSimilarityGraphTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugLigandSimilarityGraphTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteDrugLigandSimilarityGraphTaskResponse", string(data)}, " ")
}
