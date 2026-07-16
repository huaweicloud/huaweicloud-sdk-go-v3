package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingQuotasResponse Response Object
type ShowTrainingQuotasResponse struct {

	// **参数解释**：训练作业配额组。
	Quotas         *[]TrainingQuotaResponse `json:"quotas,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowTrainingQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingQuotasResponse", string(data)}, " ")
}
