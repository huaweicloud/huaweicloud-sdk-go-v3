package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseDrugReceptorInfoResponse Response Object
type ParseDrugReceptorInfoResponse struct {

	// 受体中的氨基酸残基列表
	Residues *[]ResidueDto `json:"residues,omitempty"`

	// 受体中的配体列表
	Ligands        *[]ReceptorLigandInfoDto `json:"ligands,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ParseDrugReceptorInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseDrugReceptorInfoResponse struct{}"
	}

	return strings.Join([]string{"ParseDrugReceptorInfoResponse", string(data)}, " ")
}
