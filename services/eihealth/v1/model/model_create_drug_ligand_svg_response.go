package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandSvgResponse Response Object
type CreateDrugLigandSvgResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrugLigandSvgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandSvgResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandSvgResponse", string(data)}, " ")
}
