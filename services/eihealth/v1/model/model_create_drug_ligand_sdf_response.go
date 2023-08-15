package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandSdfResponse Response Object
type CreateDrugLigandSdfResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrugLigandSdfResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandSdfResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandSdfResponse", string(data)}, " ")
}
