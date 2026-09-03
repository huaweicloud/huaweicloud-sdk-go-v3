package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTrainingQuotasRequest Request Object
type ModifyTrainingQuotasRequest struct {
	Body *ModifyTrainingQuotaRequest `json:"body,omitempty"`
}

func (o ModifyTrainingQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTrainingQuotasRequest struct{}"
	}

	return strings.Join([]string{"ModifyTrainingQuotasRequest", string(data)}, " ")
}
