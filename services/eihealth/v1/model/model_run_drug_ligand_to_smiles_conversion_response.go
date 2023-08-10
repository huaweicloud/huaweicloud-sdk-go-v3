package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDrugLigandToSmilesConversionResponse Response Object
type RunDrugLigandToSmilesConversionResponse struct {

	// 分子SMILES表达式
	Smiles         *string `json:"smiles,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunDrugLigandToSmilesConversionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDrugLigandToSmilesConversionResponse struct{}"
	}

	return strings.Join([]string{"RunDrugLigandToSmilesConversionResponse", string(data)}, " ")
}
