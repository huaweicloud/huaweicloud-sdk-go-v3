package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddExtensionEvaluationRequest Request Object
type AddExtensionEvaluationRequest struct {
	Body *Evaluation `json:"body,omitempty"`
}

func (o AddExtensionEvaluationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddExtensionEvaluationRequest struct{}"
	}

	return strings.Join([]string{"AddExtensionEvaluationRequest", string(data)}, " ")
}
