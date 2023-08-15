package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugLigandPreviewTaskResponse Response Object
type CreateDrugLigandPreviewTaskResponse struct {

	// 任务ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDrugLigandPreviewTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandPreviewTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandPreviewTaskResponse", string(data)}, " ")
}
