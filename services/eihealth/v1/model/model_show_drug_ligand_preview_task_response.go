package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrugLigandPreviewTaskResponse Response Object
type ShowDrugLigandPreviewTaskResponse struct {

	// 任务状态
	Status *string `json:"status,omitempty"`

	Result         *LigandPreviewTaskResultDto `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowDrugLigandPreviewTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrugLigandPreviewTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowDrugLigandPreviewTaskResponse", string(data)}, " ")
}
